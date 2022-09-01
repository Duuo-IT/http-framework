package framework

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/opentracing/opentracing-go"
	"time"
)

//MakeDefaultEntryEndpoint endpoint middleware por defecto
func MakeDefaultEntryEndpoint(service, path, method string, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return EndpointLogMiddleware(service, path, method, logger)(next)
	}
}

// EndpointLogMiddleware loguea tiempo que tomo servir request y error si es que hubo alguno
func EndpointLogMiddleware(service, path, method string, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		logger = log.WithPrefix(logger, "method", method)
		logger = log.WithPrefix(logger, "path", path)
		logger = log.WithPrefix(logger, "serviceName", service)
		logger = log.WithPrefix(logger, "caller", log.DefaultCaller)

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			now := time.Now()

			span := opentracing.SpanFromContext(ctx)

			if span != nil {
				logger = log.WithPrefix(logger, "spanContext", span.Context())
			}

			level.Info(logger).Log("started", now)
			level.Debug(logger).Log("request", request)

			defer func(begin time.Time) {
				level.Info(logger).Log("took", time.Since(begin))

				if err != nil {
					level.Info(logger).Log("Result", "NOK")
					level.Error(logger).Log("endpoint_error", err)
				} else {
					level.Info(logger).Log("Result", "OK")
					level.Debug(logger).Log("response", response)
				}
			}(now)
			return next(ctx, request)
		}
	}
}
