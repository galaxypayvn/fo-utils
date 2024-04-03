package apm

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type TrackOption func(p string) interface{}

func TrackWithTag(key string, value interface{}) TrackOption {
	return func(p string) interface{} {
		switch p {
		case ProviderCodeDataDog:
			return tracer.Tag(key, value)
		case ProviderCodeElastic:
			return key
		}

		return nil
	}
}

func TrackWithResourceName(name string) TrackOption {
	return func(p string) interface{} {
		switch p {
		case ProviderCodeDataDog:
			return tracer.ResourceName(name)
		case ProviderCodeElastic:
			return name
		}

		return nil
	}
}

func TrackWithSpanType(typ string) TrackOption {
	return func(p string) interface{} {
		switch p {
		case ProviderCodeDataDog:
			return tracer.SpanType(typ)
		case ProviderCodeElastic:
			return typ
		}

		return nil
	}
}

func TrackWithServiceName(name string) TrackOption {
	return func(p string) interface{} {
		switch p {
		case ProviderCodeDataDog:
			return tracer.ServiceName(name)
		case ProviderCodeElastic:
			return name
		}

		return nil
	}
}

func TrackWithMeasured() TrackOption {
	return func(p string) interface{} {
		switch p {
		case ProviderCodeDataDog:
			return tracer.Measured()
		case ProviderCodeElastic:
			return ""
		}

		return nil
	}
}
