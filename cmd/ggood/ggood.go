package main

import (
	"context"
	stdlog "log"
	"os"

	"github.com/ggood/ggood/cmd"
	"github.com/ggood/ggood/pkg/config/static"
	"github.com/ggood/ggood/pkg/version"
	"github.com/rs/zerolog/log"
)

func main() {
	// ggood config inits
	gConfig := cmd.NewGgoodCmdConfiguration()

	ctx := context.Background()
	if err := runCmd(ctx, &gConfig.Configuration); err != nil {
		stdlog.Println(err)
		os.Exit(1)
	}
}

func runCmd(_ context.Context, staticConfiguration *static.Configuration) error {
	setupLogger(staticConfiguration)

	log.Info().Str("version", version.Version).
		Msgf("Ggood version %s built on %s", version.Version, version.BuildDate)

	return nil
}
