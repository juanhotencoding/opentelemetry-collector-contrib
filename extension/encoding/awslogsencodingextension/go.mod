module github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension

go 1.23.0

require (
	github.com/aws/aws-lambda-go v1.48.0
	github.com/goccy/go-json v0.10.5
	github.com/klauspost/compress v1.18.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding v0.124.1
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.124.1
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.124.1
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v1.30.0
	go.opentelemetry.io/collector/component/componenttest v0.124.0
	go.opentelemetry.io/collector/confmap v1.30.0
	go.opentelemetry.io/collector/confmap/xconfmap v0.124.0
	go.opentelemetry.io/collector/extension v1.30.0
	go.opentelemetry.io/collector/extension/extensiontest v0.124.0
	go.opentelemetry.io/collector/pdata v1.30.0
	go.opentelemetry.io/collector/semconv v0.124.0
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.124.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.30.0 // indirect
	go.opentelemetry.io/collector/internal/telemetry v0.124.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.124.0 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.10.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/log v0.11.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250212204824-5a70512c5d8b // indirect
	google.golang.org/grpc v1.71.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding => ../

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../../pkg/golden
