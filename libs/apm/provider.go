package apm

import (
	"context"

	"google.golang.org/grpc"

	"bitbucket.org/finesys/finesys-utility/libs/serror"
)

type (
	GRPCInterceptor struct {
		ServerStream grpc.StreamServerInterceptor
		ServerUnary  grpc.UnaryServerInterceptor
		ClientStream grpc.StreamClientInterceptor
		ClientUnary  grpc.UnaryClientInterceptor
	}

	Provider interface {
		Code() string
		GetGRPCInterceptor() *GRPCInterceptor
		Start() (errx serror.SError)
		Stop() (errx serror.SError)

		StartSpan(name string, opts ...TrackOption) (span APMSpan)
		StartSpanWithContext(ctx context.Context, name string, opts ...TrackOption) (span APMSpan, nctx context.Context)
	}
)
