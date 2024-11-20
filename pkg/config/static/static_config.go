package static

import (
	"time"

	"github.com/ggood/ggood/pkg/types"
)

const (
	// DefaultServerPort defines the default port on which the server will be listening.
	DefaultServerPort = "9312"

	// DefaultServerHost is the default host.
	DefaultServerHost = "localhost"

	// DefaultIdleTimeout before closing an idle connection.
	DefaultIdleTimeout = 180 * time.Second

	// DefaultReadTimeout defines the default maximum duration for reading the entire request, including the body.
	DefaultReadTimeout = 60 * time.Second

	// DefaultWriteTimeout is the maximum duration before timing out writes of the response.
	DefaultWriteTimeout = 60 * time.Second

	// // DefaultAcmeCAServer is the default ACME API endpoint.
	// DefaultAcmeCAServer = "https://acme-v02.api.letsencrypt.org/directory"

	// // DefaultUDPTimeout defines how long to wait by default on an idle session,
	// // before releasing all resources related to that session.
	// DefaultUDPTimeout = 3 * time.Second
)

// Configuration is the static configuration.
type Configuration struct {
	Log *types.GgoodLog `description:"Ggood log settings." json:"log,omitempty" toml:"log,omitempty" yaml:"log,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
}
