package main

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticexporter"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/service/defaultcomponents"
)

func components() (component.Factories, error) {
	var errs []error

	factories, err := defaultcomponents.Components()
	if err != nil {
		return component.Factories{}, err
	}

	var processors []component.ProcessorFactory

	for _, pr := range factories.Processors {
		processors = append(processors, pr)
	}

	exporters := []component.ExporterFactory{
		elasticexporter.NewFactory(),
	}

	for _, pr := range factories.Processors {
		processors = append(processors, pr)
	}

	for _, exp := range factories.Exporters {
		exporters = append(exporters, exp)
	}

	factories.Processors, err = component.MakeProcessorFactoryMap(processors...)
	factories.Exporters, err = component.MakeExporterFactoryMap(exporters...)

	if err != nil {
		errs = append(errs, err)
	}

	return factories, componenterror.CombineErrors(errs)
}
