package main

import (
	"context"
	"encoding/json"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ggood/ggood/cmd"
	"github.com/ggood/ggood/pkg/config/static"
	"github.com/ggood/ggood/pkg/server"
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

func runCmd(ctx context.Context, staticConfiguration *static.Configuration) error {
	setupLogger(staticConfiguration)

	log.Info().Str("version", version.Version).
		Msgf("Ggood version %s built on %s", version.Version, version.BuildDate)

	jsonConf, err := json.Marshal(staticConfiguration)
	if err != nil {
		log.Error().Err(err).Msg("Could not marshal static configuration")
		log.Debug().Interface("staticConfiguration", staticConfiguration).Msg("Static configuration loaded [struct]")
	} else {
		log.Debug().RawJSON("staticConfiguration", jsonConf).Msg("Static configuration loaded [json]")
	}

	svr, err := setupServer(ctx, staticConfiguration)
	if err != nil {
		return err
	}

	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	svr.Start(ctx)
	defer svr.Close()

	svr.Wait()
	log.Info().Msg("Shutting down")
	return nil
}

func setupServer(_ context.Context, _ *static.Configuration) (*server.Server, error) {
	httpServer := &http.Server{
		Addr: net.JoinHostPort(static.DefaultServerHost, static.DefaultServerPort),
		IdleTimeout: static.DefaultIdleTimeout,
		ReadTimeout: static.DefaultReadTimeout,
		WriteTimeout: static.DefaultWriteTimeout,
	}

	return server.NewServer(httpServer), nil
}
