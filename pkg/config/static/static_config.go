package static

import "github.com/ggood/ggood/pkg/types"

// Configuration is the static configuration.
type Configuration struct {
	Log *types.GgoodLog `description:"Ggood log settings." json:"log,omitempty" toml:"log,omitempty" yaml:"log,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
}
