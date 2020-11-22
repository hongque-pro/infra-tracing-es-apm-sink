package main

import (
	"flag"
	"github.com/hongque-pro/infra-tracing-es-apm-sink/logging"
	"github.com/spf13/cast"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"

	"strings"
)

var log = logging.GetLogger("tracing-es-apm-sink")

func main() {

	settings, err := loadYml()
	if err == nil {
		configName := ""
		flag.StringVar(&configName, "metrics-addr", cast.ToString(settings["metrics-addr"]), "metrics address")
	}

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
	}

	app, err := service.New(*parameters)

	if err != nil {
		log.Fatal("failed to construct the application: %w", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal("application run finished with error: %w", err)
	}
}
