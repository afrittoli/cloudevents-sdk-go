package http

import (
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport/http"
)

func sortOptions(opts ...interface{}) ([]http.Option, []client.Option) {
	tOpts := []http.Option(nil)
	cOpts := []client.Option(nil)

	for _, o := range opts {
		if opt, ok := o.(http.Option); ok {
			tOpts = append(tOpts, opt)
		} else if opt, ok := o.(client.Option); ok {
			cOpts = append(cOpts, opt)
		}
	}
	return tOpts, cOpts
}

// New returns a new HTTP client
func New(opts ...interface{}) (client.Client, error) {
	tOpts, cOpts := sortOptions(opts...)

	t, err := http.New(tOpts...)
	if err != nil {
		return nil, err
	}
	c, err := client.New(t, cOpts...)
	if err != nil {
		return nil, err
	}
	return c, nil
}
