package framework

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"net/url"
	"time"
)

// HTTPClientBuilder builder de http client
type HTTPClientBuilder interface {
	Build() endpoint.Endpoint
}

type httpClientBuilder struct {
	method         string
	uri            string
	timeout        time.Duration
	encodeRequest  httptransport.EncodeRequestFunc
	decodeResponse httptransport.DecodeResponseFunc
	logger         log.Logger
}

func (h *httpClientBuilder) Build() endpoint.Endpoint {
	url, _ := url.Parse(h.uri)

	opts := []httptransport.ClientOption{
		httptransport.SetClient(&http.Client{Timeout: h.timeout}),
	}

	return httptransport.NewClient(
		h.method,
		url,
		h.encodeRequest,
		h.decodeResponse,
		opts...,
	).Endpoint()
}

// MakeHTTPClientBuilder crea un nuevo http client builder
func MakeHTTPClientBuilder(method, uri string,
	timeout time.Duration,
	encodeRequest httptransport.EncodeRequestFunc,
	decodeResponse httptransport.DecodeResponseFunc,
	logger log.Logger,
) HTTPClientBuilder {
	return &httpClientBuilder{
		method:         method,
		uri:            uri,
		timeout:        timeout,
		encodeRequest:  encodeRequest,
		decodeResponse: decodeResponse,
		logger:         logger,
	}
}
