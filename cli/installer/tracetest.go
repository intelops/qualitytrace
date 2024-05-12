package installer

import (
	"bytes"
	_ "embed"
	"html/template"

	"fmt"

	cliUI "github.com/intelops/qualityTrace/cli/ui"
)

func configureDemoApp(conf configuration, ui cliUI.UI) configuration {
	conf.set("demo.enable.pokeshop", !conf.Bool("installer.only_qualityTrace"))
	conf.set("demo.enable.otel", false)

	switch conf.String("installer") {
	case "docker-compose":
		conf.set("demo.endpoint.pokeshop.http", "http://demo-api:8081")
		conf.set("demo.endpoint.pokeshop.grpc", "demo-rpc:8082")
		conf.set("demo.endpoint.pokeshop.kafka", "stream:9092")
		conf.set("demo.endpoint.otel.frontend", "http://otel-frontend:8084")
		conf.set("demo.endpoint.otel.product_catalog", "otel-productcatalogservice:3550")
		conf.set("demo.endpoint.otel.cart", "otel-cartservice:7070")
		conf.set("demo.endpoint.otel.checkout", "otel-checkoutservice:5050")
	case "kubernetes":
		conf.set("demo.endpoint.pokeshop.http", "http://demo-pokemon-api.demo")
		conf.set("demo.endpoint.pokeshop.grpc", "demo-pokemon-api.demo:8082")
		conf.set("demo.endpoint.pokeshop.kafka", "stream.demo:9092")
		conf.set("demo.endpoint.otel.frontend", "http://otel-frontend.otel-demo:8084")
		conf.set("demo.endpoint.otel.product_catalog", "otel-productcatalogservice.otel-demo:3550")
		conf.set("demo.endpoint.otel.cart", "otel-cartservice.otel-demo:7070")
		conf.set("demo.endpoint.otel.checkout", "otel-checkoutservice.otel-demo:5050")
	}

	return conf
}

func configureTracetest(conf configuration, ui cliUI.UI) configuration {
	conf = configureBackend(conf, ui)
	conf.set("qualityTrace.analytics", true)

	return conf
}

func configureBackend(conf configuration, ui cliUI.UI) configuration {
	installBackend := !conf.Bool("installer.only_qualityTrace")
	conf.set("qualityTrace.backend.install", installBackend)

	if !installBackend {
		conf.set("qualityTrace.backend.type", "")
		return conf
	}

	// default values
	switch conf.String("installer") {
	case "docker-compose":
		conf.set("qualityTrace.backend.type", "otlp")
		conf.set("qualityTrace.backend.tls.insecure", true)
		conf.set("qualityTrace.backend.endpoint.collector", "http://otel-collector:4317")
		conf.set("qualityTrace.backend.endpoint", "qualityTrace:4317")
	case "kubernetes":
		conf.set("qualityTrace.backend.type", "otlp")
		conf.set("qualityTrace.backend.tls.insecure", true)
		conf.set("qualityTrace.backend.endpoint.collector", "http://otel-collector.qualityTrace:4317")
		conf.set("qualityTrace.backend.endpoint", "qualityTrace:4317")

	default:
		conf.set("qualityTrace.backend.type", "")
	}

	return conf
}

//go:embed templates/config.yaml.tpl
var configTemplate string

func getTracetestConfigFileContents(pHost, pUser, pPasswd string, ui cliUI.UI, config configuration) []byte {
	vals := map[string]string{
		"pHost":   pHost,
		"pUser":   pUser,
		"pPasswd": pPasswd,
	}

	tpl, err := template.New("page").Parse(configTemplate)
	if err != nil {
		ui.Panic(fmt.Errorf("cannot parse config template: %w", err))
	}

	out := &bytes.Buffer{}
	tpl.Execute(out, vals)

	return out.Bytes()
}

//go:embed templates/provision.yaml.tpl
var provisionTemplate string

func getTracetestProvisionFileContents(ui cliUI.UI, config configuration) []byte {
	vals := map[string]string{
		"installBackend":   fmt.Sprintf("%t", config.Bool("qualityTrace.backend.install")),
		"backendType":      config.String("qualityTrace.backend.type"),
		"backendEndpoint":  config.String("qualityTrace.backend.endpoint.query"),
		"backendInsecure":  config.String("qualityTrace.backend.tls.insecure"),
		"backendAddresses": config.String("qualityTrace.backend.addresses"),
		"backendIndex":     config.String("qualityTrace.backend.index"),
		"backendToken":     config.String("qualityTrace.backend.token"),
		"backendRealm":     config.String("qualityTrace.backend.realm"),

		"analyticsEnabled": fmt.Sprintf("%t", config.Bool("qualityTrace.analytics")),

		"enablePokeshopDemo": fmt.Sprintf("%t", config.Bool("demo.enable.pokeshop")),
		"enableOtelDemo":     fmt.Sprintf("%t", config.Bool("demo.enable.otel")),
		"pokeshopHttp":       config.String("demo.endpoint.pokeshop.http"),
		"pokeshopGrpc":       config.String("demo.endpoint.pokeshop.grpc"),
		"pokeshopKafka":      config.String("demo.endpoint.pokeshop.kafka"),
		"otelFrontend":       config.String("demo.endpoint.otel.frontend"),
		"otelProductCatalog": config.String("demo.endpoint.otel.product_catalog"),
		"otelCart":           config.String("demo.endpoint.otel.cart"),
		"otelCheckout":       config.String("demo.endpoint.otel.checkout"),
	}

	tpl, err := template.New("page").Parse(provisionTemplate)
	if err != nil {
		ui.Panic(fmt.Errorf("cannot parse config template: %w", err))
	}

	out := &bytes.Buffer{}
	tpl.Execute(out, vals)

	return out.Bytes()
}
