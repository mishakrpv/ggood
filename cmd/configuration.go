package cmd

import "github.com/ggood/ggood/pkg/config/static"

// GgoodCmdConfiguration wraps the static configuration and extra parameters.
type GgoodCmdConfiguration struct {
	static.Configuration `export:"true"`
}

func NewGgoodCmdConfiguration() *GgoodCmdConfiguration {
	return &GgoodCmdConfiguration{
		Configuration: static.Configuration{},
	}
}
