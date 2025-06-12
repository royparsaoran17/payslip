package provider

import (
	"manage-se/internal/appctx"
	"manage-se/internal/provider/dependencies"
	"manage-se/internal/provider/payroll"
	"net/http"
	"time"
)

type Provider struct {
	Payroll Payroll
}

func NewProviders(cfg *appctx.Provider, options ...Option) *Provider {
	dep := defaultDependency()

	for _, opt := range options {
		opt(dep)
	}

	return &Provider{
		Payroll: payroll.NewClient(&cfg.Payroll, dep),
	}
}

func defaultDependency() *dependencies.Dependency {
	client := http.DefaultClient
	client.Timeout = time.Duration(10) * time.Second

	return &dependencies.Dependency{
		HttpClient: http.DefaultClient,
	}
}
