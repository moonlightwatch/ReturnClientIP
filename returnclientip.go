package ReturnClientIP

import (
	"context"
	"net"
	"net/http"
)

type Config struct {
	Active bool
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Active: true,
	}
}

type MethodBlock struct {
	cfg  *Config
	next http.Handler
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &MethodBlock{
		cfg:  config,
		next: next,
	}, nil
}

func (m *MethodBlock) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if m.cfg.Active {
		rw.WriteHeader(http.StatusOK)
		host, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			rw.Write([]byte(req.RemoteAddr))
		} else {
			rw.Write([]byte(host))
		}

		return
	}
	m.next.ServeHTTP(rw, req)
}
