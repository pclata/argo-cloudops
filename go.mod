module github.com/argoproj-labs/argo-cloudops

go 1.16

require (
	github.com/argoproj/argo-workflows/v3 v3.0.0-rc3
	github.com/aws/aws-sdk-go v1.33.16
	github.com/distribution/distribution v2.7.1+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-test/deep v1.0.7 // indirect
	github.com/google/go-cmp v0.5.2
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-retryablehttp v0.6.7 // indirect
	github.com/hashicorp/hcl v1.0.1-vault // indirect
	github.com/hashicorp/vault/api v1.0.5-0.20201001211907-38d91b749c77
	github.com/hashicorp/vault/sdk v0.1.14-0.20210127182440-8477cfe632c0 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/spf13/cobra v1.1.3
	google.golang.org/grpc v1.33.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.19.6
	k8s.io/apimachinery v0.19.6
)
