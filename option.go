package zerodriver

func newConfig(opts ...Option) *config {
	var cfg config
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	return &cfg
}

// Option is an option to change logger configuration.
type Option interface {
	apply(*config)
}

type optionFunc struct {
	f func(*config)
}

func (o *optionFunc) apply(opts *config) {
	o.f(opts)
}

func newOptionFunc(f func(cfg *config)) *optionFunc {
	return &optionFunc{
		f: f,
	}
}

// WithServiceName enable adding serviceContext field with given service name.
func WithServiceName(serviceName string) Option {
	return newOptionFunc(func(cfg *config) {
		cfg.serviceName = serviceName
	})
}

// WithReportAllErrors enable adding fields used for error reporting when log level is error or above.
func WithReportAllErrors() Option {
	return newOptionFunc(func(cfg *config) {
		cfg.reportAllErrors = true
	})
}
