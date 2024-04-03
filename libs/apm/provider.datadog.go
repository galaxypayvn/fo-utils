package apm

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"sync"
	"time"

	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"bitbucket.org/finesys/finesys-utility/constants"
	"bitbucket.org/finesys/finesys-utility/libs/serror"
	"bitbucket.org/finesys/finesys-utility/models"
	"bitbucket.org/finesys/finesys-utility/utils/utarray"
	"bitbucket.org/finesys/finesys-utility/utils/utinterface"
	"bitbucket.org/finesys/finesys-utility/utils/utstring"
)

type (
	DataDogConfig struct {
		Options []tracer.StartOption
	}

	apmDataDog struct {
		Instance  APM
		AgentAddr string
		Options   []tracer.StartOption
		Ticker    *time.Ticker
	}

	apmSpanDataDog struct {
		ddtrace.Span
		mu         sync.Mutex
		Context    context.Context
		IsFinished bool
	}
)

const ProviderCodeDataDog = "datadog"

const (
	TagDataDogHTTPMethod    = ext.HTTPMethod
	TagDataDogHTTPURL       = ext.HTTPURL
	TagDataDogHTTPCode      = ext.HTTPCode
	TagDataDogHTTPRouteType = "http.route.type"

	TagDataDogError        = ext.Error
	TagDataDogErrorDetails = ext.ErrorDetails
	TagDataDogErrorMsg     = ext.ErrorMsg
	TagDataDogErrorStack   = ext.ErrorStack
	TagDataDogErrorType    = ext.ErrorType
)

func SetupDataDog(cfg DataDogConfig) (errx serror.SError) {
	obj := &apmDataDog{
		AgentAddr: net.JoinHostPort(
			utstring.Env(constants.APMProviderDatadogAgentHost, "localhost"),
			utstring.Env(constants.APMProviderDatadogAgentPort, "8126"),
		),
		Options: cfg.Options,
	}

	obj.Instance, errx = GetInstance()
	if errx != nil {
		errx.AddComments("while get instance")
		return errx
	}

	errx = obj.Instance.Register(obj)
	if errx != nil {
		errx.AddCommentf("while register %s provider into apm instance", ProviderCodeDataDog)
		return errx
	}

	return errx
}

func (*apmDataDog) Code() string {
	return ProviderCodeDataDog
}

func (ox *apmDataDog) GetGRPCInterceptor() (resp *GRPCInterceptor) {
	var (
		svc  = ox.Instance.GetService()
		opts = []grpctrace.Option{
			grpctrace.WithServiceName(svc.Key),
			grpctrace.WithAnalytics(true),
			grpctrace.WithMetadataTags(),
			grpctrace.WithRequestTags(),
		}
	)

	resp = &GRPCInterceptor{
		ServerStream: grpctrace.StreamServerInterceptor(opts...),
		ServerUnary:  grpctrace.UnaryServerInterceptor(opts...),
		ClientStream: grpctrace.StreamClientInterceptor(opts...),
		ClientUnary:  grpctrace.UnaryClientInterceptor(opts...),
	}

	return resp
}

func (ox *apmDataDog) Start() (errx serror.SError) {
	svc := ox.Instance.GetService()

	tracer.Start(
		append(
			ox.Options,
			tracer.WithAgentAddr(ox.AgentAddr),
			tracer.WithService(svc.Key),
			tracer.WithServiceVersion(svc.Version),
			tracer.WithEnv(models.Environment()),
		)...,
	)

	if ox.Ticker == nil {
		ox.Ticker = time.NewTicker(time.Second * 30)

		go func() {
			for range ox.Ticker.C {
				span, _ := tracer.StartSpanFromContext(
					context.Background(),
					"heartbeat",
					tracer.SpanType("custom"),
					tracer.ResourceName("heartbeat"),
				)
				time.Sleep(time.Millisecond)
				span.Finish()
			}
		}()
	}

	return errx
}

func (ox *apmDataDog) Stop() (errx serror.SError) {
	ox.Ticker.Stop()
	ox.Ticker = nil

	tracer.Stop()

	return errx
}

func (ox *apmDataDog) StartSpan(name string, opts ...TrackOption) (span APMSpan) {
	span, _ = ox.StartSpanWithContext(context.Background(), name, opts...)
	return span
}

func (ox *apmDataDog) StartSpanWithContext(ctx context.Context, name string, opts ...TrackOption) (span APMSpan, nctx context.Context) {
	var xopts []ddtrace.StartSpanOption
	for _, v := range opts {
		if v == nil {
			continue
		}

		vx := v(ProviderCodeDataDog)
		if cur, ok := vx.(ddtrace.StartSpanOption); ok {
			xopts = append(xopts, cur)
		}
	}

	var ospan ddtrace.Span
	ospan, nctx = tracer.StartSpanFromContext(ctx, name, xopts...)

	span = &apmSpanDataDog{
		Span:    ospan,
		Context: nctx,
	}
	return span, nctx
}

func (ox *apmSpanDataDog) SetTag(n string, v interface{}) {
	if ox.IsFinished {
		return
	}

	ox.Span.SetTag(n, v)
}

func (ox *apmSpanDataDog) SetData(n string, v interface{}) {
	if ox.IsFinished {
		return
	}

	ox.Span.SetBaggageItem(n, utinterface.ToString(v))
}

func (ox *apmSpanDataDog) GetContext() context.Context {
	return ox.Context
}

func (ox *apmSpanDataDog) GetProviderCode() string {
	return ProviderCodeDataDog
}

func (ox *apmSpanDataDog) Stop(errx serror.SError) {
	if ox.IsFinished {
		return
	}

	ox.mu.Lock()
	defer ox.mu.Unlock()

	ox.IsFinished = true

	var (
		xspan = ox.Span
		sts   = "OK"
	)

	if errx != nil {
		if errx.Key() != ErrorKeyTrackManual {
			xspan.SetTag(TagDataDogError, errx)
			xspan.SetTag(TagDataDogErrorDetails, errx.SimpleString())
			xspan.SetTag(TagDataDogErrorMsg, errx.Title())
			xspan.SetTag(TagDataDogErrorStack, errx.StackTraces(0))

			xspan.SetTag(TagDataDogErrorType, fmt.Sprintf("key:%s", errx.Key()))
			if utarray.IsExist(errx.Key(), []string{"-", ""}) {
				var ok bool
				if errx.Code() > 0 {
					xspan.SetTag(TagDataDogErrorType, fmt.Sprintf("code:%d", errx.Code()))
					ok = true
				}

				if !ok {
					oerr := errx.Cause()
					if oerr != nil {
						xspan.SetTag(TagDataDogErrorType, reflect.TypeOf(errx.Cause()).String())
						ok = true
					}
				}

				if !ok {
					xspan.SetTag(TagDataDogErrorType, errx.Type())
				}
			}
		}

		sts = "ERR"
	}

	xspan.SetTag("span.status", sts)
	xspan.Finish()
}
