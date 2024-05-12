module github.com/kform-dev/kform-sdk-go

go 1.22.2

replace github.com/google/cel-go => github.com/google/cel-go v0.16.1

replace k8s.io/client-go => k8s.io/client-go v0.28.6

replace k8s.io/api => k8s.io/api v0.28.6

replace k8s.io/apimachinery => k8s.io/apimachinery v0.28.6

replace k8s.io/component-base => k8s.io/component-base v0.28.6

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.28.6

replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20230717233707-2695361300d9

replace k8s.io/apiserver => k8s.io/apiserver v0.28.6

replace sigs.k8s.io/kustomize/kyaml => sigs.k8s.io/kustomize/kyaml v0.15.0

replace sigs.k8s.io/kustomize/api => sigs.k8s.io/kustomize/api v0.15.0

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.28.6

require (
	github.com/henderiw/logger v0.0.0-20230911123436-8655829b1abe
	github.com/kform-dev/kform-plugin v0.0.0-20240512102710-e5ebed866b1d
	k8s.io/apiextensions-apiserver v0.0.0-00010101000000-000000000000
	k8s.io/apimachinery v0.28.6
	k8s.io/kube-openapi v0.0.0-20230717233707-2695361300d9
)

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/utils v0.0.0-20230406110748-d93618cff8a2 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)
