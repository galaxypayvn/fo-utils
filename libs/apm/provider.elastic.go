package apm

import (
	"context"
	"net"
	"os"
	"sync"
	"time"

	"go.elastic.co/apm/module/apmgrpc/v2"
	elasticapm "go.elastic.co/apm/v2"

	"bitbucket.org/finesys/finesys-utility/constants"
	"bitbucket.org/finesys/finesys-utility/libs/serror"
	"bitbucket.org/finesys/finesys-utility/utils/utstring"
)

type (
	ElasticConfig struct {
		Options []elasticapm.TracerOptions
	}

	apmElastic struct {
		Instance  APM
		AgentAddr string
		tracer    *elasticapm.Tracer
		Options   []elasticapm.TracerOptions
		Ticker    *time.Ticker
	}

	apmSpanElastic struct {
		trx        *elasticapm.Transaction
		span       *elasticapm.Span
		mu         sync.Mutex
		Context    context.Context
		IsFinished bool
	}
)

const ProviderCodeElastic = "elastic"

func SetupElastic(cfg ElasticConfig) (errx serror.SError) {
	obj := &apmElastic{
		AgentAddr: net.JoinHostPort(
			utstring.Env(constants.APMProviderElasticAgentHost, "localhost"),
			utstring.Env(constants.APMProviderElasticAgentPort, "8200"),
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
		errx.AddCommentf("while register %s provider into apm instance", ProviderCodeElastic)
		return errx
	}

	return errx
}

func (*apmElastic) Code() string {
	return ProviderCodeElastic
}

func (ox *apmElastic) GetGRPCInterceptor() (resp *GRPCInterceptor) {
	var (
		opts = []apmgrpc.ServerOption{
			apmgrpc.WithRecovery(),
		}
	)

	resp = &GRPCInterceptor{
		ServerStream: apmgrpc.NewStreamServerInterceptor(opts...),
		ServerUnary:  apmgrpc.NewUnaryServerInterceptor(opts...),
		ClientStream: apmgrpc.NewStreamClientInterceptor(),
		ClientUnary:  apmgrpc.NewUnaryClientInterceptor(),
	}

	return resp
}

func (ox *apmElastic) Start() (errx serror.SError) {
	var err error
	svc := ox.Instance.GetService()
	os.Setenv("ELASTIC_APM_SERVER_URL", ox.AgentAddr)
	ox.tracer, err = elasticapm.NewTracerOptions(elasticapm.TracerOptions{
		ServiceName:        svc.Name,
		ServiceVersion:     svc.Version,
		ServiceEnvironment: svc.Namespace,
	})
	if err != nil {
		errx = serror.New(err.Error())
		return
	}

	if ox.Ticker == nil {
		ox.Ticker = time.NewTicker(time.Second * 30)

		go func() {
			for range ox.Ticker.C {
				tx := ox.tracer.StartTransaction("name", "type")
				defer tx.End()
				span := tx.StartSpanOptions("heartbeat", "custom", elasticapm.SpanOptions{})
				time.Sleep(time.Millisecond)
				span.End()
			}
		}()
	}

	return errx
}

func (ox *apmElastic) Stop() (errx serror.SError) {
	ox.Ticker.Stop()
	ox.Ticker = nil

	ox.tracer.Close()

	return errx
}

func (ox *apmElastic) StartSpan(name string, opts ...TrackOption) (span APMSpan) {
	span, _ = ox.StartSpanWithContext(context.Background(), name, opts...)
	return span
}

func (ox *apmElastic) StartSpanWithContext(ctx context.Context, name string, opts ...TrackOption) (span APMSpan, nctx context.Context) {
	var kind string
	for _, v := range opts {
		if v == nil {
			continue
		}

		vx := v(ProviderCodeElastic)
		if cur, ok := vx.(string); ok {
			kind = cur
		}
	}

	trx := ox.tracer.StartTransaction(name, "custom")
	ospan := trx.StartSpanOptions(kind, "custom", elasticapm.SpanOptions{})
	nctx = elasticapm.ContextWithTransaction(context.Background(), trx)

	span = &apmSpanElastic{
		trx:     trx,
		span:    ospan,
		Context: nctx,
	}
	return span, nctx
}

func (ox *apmSpanElastic) SetTag(n string, v interface{}) {
	if ox.IsFinished {
		return
	}
}

func (ox *apmSpanElastic) SetData(n string, v interface{}) {
	if ox.IsFinished {
		return
	}
}

func (ox *apmSpanElastic) GetContext() context.Context {
	return ox.Context
}

func (ox *apmSpanElastic) GetProviderCode() string {
	return ProviderCodeElastic
}

func (ox *apmSpanElastic) Stop(errx serror.SError) {
	if ox.IsFinished {
		return
	}

	ox.mu.Lock()
	defer ox.mu.Unlock()

	ox.IsFinished = true

	ox.span.End()
	ox.trx.End()

	os.Unsetenv("ELASTIC_APM_SERVER_URL")
}
