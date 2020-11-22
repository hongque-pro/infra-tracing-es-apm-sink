package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmodels"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"strings"
)

var defaultConfig = `
receivers:
  kafka:
    protocol_version: 2.0.0
    brokers: ${KAFKA_SERVER:127.0.0.1:9092}
    topic: ${KAFKA_TOPIC:telemetry-spans}
    encoding: otlp_proto
    group_id: ${KAFKA_GROUP_ID:infra-tracing-sink}
    client_id: ${KAFKA_CLIENT_ID:infra-tracing-sink}

exporters:
  elastic:
    apm_server_url: ${ES_APM_URL:http://127.0.0.1:18200}
    secret_token: "hunter2"
  logging:

processors:
  batch:

service:
  pipelines:
    traces:
      receivers: [kafka]
      exporters: [logging,elastic]
      processors: [batch]
`

func staticConfigFactory(v *viper.Viper, factories component.Factories) (*configmodels.Config, error) {
	//参考：https://github.com/open-telemetry/opentelemetry-collector/tree/master/receiver/kafkareceiver
	v.SetConfigType("yml")
	resultYml := parseYml()
	var configReader = strings.NewReader(resultYml)
	err := v.ReadConfig(configReader)
	if err == nil {
		return config.Load(v, factories)
	} else {
		return nil, err
	}

}

func loadYml() (map[string]interface{}, error) {
	content := parseYml()
	var settings = make(map[string]interface{})
	buf := new(bytes.Buffer)
	reader := strings.NewReader(content)
	if _, err := buf.ReadFrom(reader); err == nil {
		if err = yaml.Unmarshal(buf.Bytes(), &settings); err != nil {
			return nil, err
		}
	}
	return settings, nil
}

func parseYml() string {
	var content = defaultConfig
	r := regexp.MustCompile(`\$\{([0-9a-zA-Z_]+)(:\s*(.*))?\}`)
	matches := r.FindAllStringSubmatch(content, -1)

	var actualContent = content
	var errBuilder = &strings.Builder{}
	for _, m := range matches {
		if len(m) == 4 {
			subStr := m[0]
			envName := m[1]
			defaultValue := strings.TrimSpace(m[3])
			v := os.Getenv(envName)
			if strings.TrimSpace(v) == "" && defaultValue != "" {
				v = defaultValue
			}
			if strings.TrimSpace(v) == "" {
				e := fmt.Sprintf("The environment variables '%s' is missing, but config file required it", envName)
				errBuilder.WriteString(fmt.Sprintln(e))
			}
			actualContent = strings.Replace(actualContent, subStr, v, 1)
		}
	}
	if errBuilder.Len() > 0 {
		log.Warn(errBuilder.String())
	}
	return actualContent
}
