package env

import (
	"time"

	"github.com/EduardoNGomes/envschema"
)

type Env struct {
	Port                           int           `env:"PORT"`
	Debug                          bool          `env:"DEBUG" default:"false"`
	Timeout                        time.Duration `env:"TIMEOUT" default:"5s"`
	DatabaseURL                    string        `env:"DATABASE_URL"`
	DatabaseIddleConnections       int           `env:"DATABASE_IDLE_CONN" default:"10"`
	DatabaseMaxConnections         int           `env:"DATABASE_MAX_CONN" default:"100"`
	DatabaseMaxConnectionsLifetime int           `env:"DATABASE_MAX_CONN_LIFETIME" default:"30"`
	DatabaseMaxIdleTime            int           `env:"DATABASE_MAX_IDLE_TIME" default:"40"`
}

func New() (*Env, error) {
	return envschema.Validate[Env]()
}
