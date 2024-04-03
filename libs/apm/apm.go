package apm

import (
	"context"
	"sync"

	"google.golang.org/grpc"

	"bitbucket.org/finesys/finesys-utility/libs/serror"
	"bitbucket.org/finesys/finesys-utility/models"
	"bitbucket.org/finesys/finesys-utility/utils/utfunc"

	log "github.com/sirupsen/logrus"
)

type (
	apm struct {
		svc *models.Service
		mu  sync.Mutex

		IsStarted bool
		Providers map[string]Provider
	}

	APM interface {
		GetService() *models.Service

		Register(pvdr Provider) (errx serror.SError)
		Shutdown()

		GetGRPCServerInterceptor(args *GetGRPCServerInterceptorArguments) (opts []grpc.ServerOption)
		GetGRPCClientInterceptor(args *GetGRPCClientInterceptorArguments) (opts []grpc.DialOption)

		CreateTrackingWithContext(ctx context.Context, name string, opts ...TrackOption) (octx context.Context, apms APMSpans, errx serror.SError)

		Tracking(name string, fn func(APMSpans) serror.SError, opts ...TrackOption) (errx serror.SError)
		TrackingWithContext(ctx context.Context, name string, fn func(context.Context, APMSpans) serror.SError, opts ...TrackOption) (octx context.Context, errx serror.SError)
	}

	APMSpan interface {
		SetTag(n string, v interface{})
		SetData(n string, v interface{})
		GetContext() (ctx context.Context)
		Stop(errx serror.SError)
	}

	APMSpans []APMSpan

	GetGRPCServerInterceptorArguments struct {
		StreamInterceptors []grpc.StreamServerInterceptor
		UnaryInterceptors  []grpc.UnaryServerInterceptor
	}

	GetGRPCClientInterceptorArguments struct {
		StreamInterceptors []grpc.StreamClientInterceptor
		UnaryInterceptors  []grpc.UnaryClientInterceptor
	}
)

var _instance *apm

func SetupAPM(cfg *models.Service) (obj APM, errx serror.SError) {
	if _instance != nil {
		return _instance, errx
	}

	_instance = &apm{
		svc: cfg,

		IsStarted: true,
		Providers: make(map[string]Provider),
	}

	return _instance, errx
}

func GetInstance() (obj APM, errx serror.SError) {
	if _instance == nil {
		errx = serror.Newk(ErrorKeyNotSetup, "APM instance is not setup yet")
		return obj, errx
	}

	return _instance, errx
}

func Register(pvdr Provider) (errx serror.SError) {
	var obj APM
	obj, errx = GetInstance()
	if errx != nil {
		errx.AddComments("while get instance")
		return errx
	}

	errx = obj.Register(pvdr)
	if errx != nil {
		errx.AddComments("while register apm provider")
		return errx
	}

	return errx
}

func (ox *apm) GetService() *models.Service {
	return ox.svc
}

func (ox *apm) Register(pvdr Provider) (errx serror.SError) {
	ox.mu.Lock()
	defer ox.mu.Unlock()

	code := pvdr.Code()
	if _, ok := ox.Providers[code]; ok {
		log.Warnf("Skip registering APM provider %s, detail: already registered", code)
		return errx
	}

	errx = pvdr.Start()
	if errx != nil {
		errx.AddCommentf("while starting provider %s", code)
		return errx
	}

	ox.Providers[code] = pvdr
	return errx
}

func (ox *apm) Shutdown() {
	ox.mu.Lock()
	defer ox.mu.Unlock()

	for k, v := range ox.Providers {
		errx := v.Stop()
		if errx != nil {
			errx.AddCommentf("while stop provider %s", k)
			log.Error(errx)
		}
	}
}

func (ox *apm) GetGRPCServerInterceptor(args *GetGRPCServerInterceptorArguments) (opts []grpc.ServerOption) {
	var (
		streams []grpc.StreamServerInterceptor
		unary   []grpc.UnaryServerInterceptor
	)

	if args != nil {
		streams = args.StreamInterceptors
		unary = args.UnaryInterceptors
	}

	for _, v := range ox.Providers {
		intrc := v.GetGRPCInterceptor()
		if intrc != nil {
			if intrc.ServerStream != nil {
				streams = append(streams, intrc.ServerStream)
			}

			if intrc.ServerUnary != nil {
				unary = append(unary, intrc.ServerUnary)
			}
		}
	}

	if len(streams) > 0 {
		opts = append(opts, grpc.StreamInterceptor(GRPCServerChainStreamInterceptor(streams...)))
	}

	if len(unary) > 0 {
		opts = append(opts, grpc.UnaryInterceptor(GRPCServerChainUnaryInterceptor(unary...)))
	}

	return opts
}

func (ox *apm) GetGRPCClientInterceptor(args *GetGRPCClientInterceptorArguments) (opts []grpc.DialOption) {
	var (
		streams []grpc.StreamClientInterceptor
		unary   []grpc.UnaryClientInterceptor
	)

	if args != nil {
		streams = args.StreamInterceptors
		unary = args.UnaryInterceptors
	}

	for _, v := range ox.Providers {
		intrc := v.GetGRPCInterceptor()
		if intrc != nil {
			if intrc.ClientStream != nil {
				streams = append(streams, intrc.ClientStream)
			}

			if intrc.ClientUnary != nil {
				unary = append(unary, intrc.ClientUnary)
			}
		}
	}

	if len(streams) > 0 {
		opts = append(opts, grpc.WithChainStreamInterceptor(streams...))
	}

	if len(unary) > 0 {
		opts = append(opts, grpc.WithChainUnaryInterceptor(unary...))
	}

	return opts
}

func (ox *apm) CreateTrackingWithContext(ctx context.Context, name string, opts ...TrackOption) (octx context.Context, apms APMSpans, errx serror.SError) {
	for _, v := range ox.Providers {
		var cspan APMSpan
		cspan, ctx = v.StartSpanWithContext(ctx, name, opts...)
		if cspan == nil {
			continue
		}

		apms = append(apms, cspan)
	}

	octx = ctx
	return octx, apms, errx
}

func (ox *apm) Tracking(name string, fn func(APMSpans) serror.SError, opts ...TrackOption) (errx serror.SError) {
	var spans APMSpans
	_, spans, errx = ox.CreateTrackingWithContext(context.Background(), name, opts...)
	if errx != nil {
		errx.AddComments("while create tracking with context")
		return errx
	}

	defer func() {
		spans.Stop(errx)
	}()

	errx = utfunc.Try(func() (errx serror.SError) {
		if fn != nil {
			errx = fn(spans)
		}

		return errx
	})
	if errx != nil {
		errx.AddComments("while trying")
		return errx
	}

	return errx
}

func (ox *apm) TrackingWithContext(ctx context.Context, name string, fn func(context.Context, APMSpans) serror.SError, opts ...TrackOption) (octx context.Context, errx serror.SError) {
	var spans APMSpans
	octx, spans, errx = ox.CreateTrackingWithContext(ctx, name, opts...)
	if errx != nil {
		errx.AddComments("while create tracking with context")
		return octx, errx
	}

	defer func() {
		spans.Stop(errx)
	}()

	errx = utfunc.Try(func() (errx serror.SError) {
		if fn != nil {
			errx = fn(octx, spans)
		}

		return errx
	})
	if errx != nil {
		errx.AddComments("while trying")
		return octx, errx
	}

	return octx, errx
}

func (ox APMSpans) SetTag(n string, v interface{}) {
	for _, o := range ox {
		o.SetTag(n, v)
	}
}

func (ox APMSpans) SetData(n string, v interface{}) {
	for _, o := range ox {
		o.SetData(n, v)
	}
}

func (ox APMSpans) GetContext() (ctx context.Context) {
	if len(ox) <= 0 {
		return context.Background()
	}

	return ox[len(ox)-1].GetContext()
}

func (ox APMSpans) Stop(errx serror.SError) {
	for _, v := range ox {
		v.Stop(errx)
	}
}
