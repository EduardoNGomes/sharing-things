package env

import (
	"time"

	"github.com/EduardoNGomes/envschema"
)

type Env struct {
	Port    int           `env:"PORT"`
	Debug   bool          `env:"DEBUG" default:"false"`
	Timeout time.Duration `env:"TIMEOUT" default:"5s"`
}

func New() (*Env, error) {
	return envschema.Validate[Env]()
}
