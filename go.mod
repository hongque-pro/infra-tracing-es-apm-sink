module github.com/hongque-pro/infra-tracing-es-apm-sink

go 1.15

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticexporter v0.15.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.7.1
	go.opentelemetry.io/collector v0.15.0
	go.uber.org/multierr v1.5.0
	go.uber.org/zap v1.16.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
