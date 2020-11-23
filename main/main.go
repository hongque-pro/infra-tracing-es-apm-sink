package main

import (
	"flag"
	"github.com/hongque-pro/infra-tracing-es-apm-sink/logging"
	"github.com/spf13/viper"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
	"os"

	"strings"
)

var log = logging.GetLogger("tracing-es-apm-sink")

func main() {

	yml := parseYml()

	factories, err := components()
	if err != nil {
		log.Fatalf("failed to build components: %v", err)
	}

	info := component.ApplicationStartInfo{
		ExeName:  "infra-tracing-sink",
		LongName: "Infra telemetry sink for elastic apm",
		Version:  "1.0.0",
	}

	parameters := &service.Parameters{
		ApplicationStartInfo: info,
		Factories:            factories,
	}

	configFile := flag.String("config", "", "Path to the config file")
	if configFile == nil || strings.TrimSpace(*configFile) == "" {
		parameters.ConfigFactory = staticConfigFactory
		log.Info("Embedded configured (you can use --config [config file]) :\n", yml)
	}

	app, err := service.New(*parameters)
	if err == nil {
		settings, err := loadYml()
		if err == nil {
			setConfig(settings, app, "METRICS_ADDR", "metrics-addr")
		}
	}

	if err != nil {
		log.Fatal("failed to construct the application: %w", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal("application run finished with error: %w", err)
	}
}

func setConfig(settings *viper.Viper, app *service.Application, envKey string, flagKey string) {
	metricsAddr := os.Getenv(envKey)
	if len(metricsAddr) > 0 {
		_ = app.Command().Flags().Set(flagKey, metricsAddr)
	}
}
