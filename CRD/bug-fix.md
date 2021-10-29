% go test ./...                                                                                            
generated/clientset/versioned/scheme/register.go:27:2: cannot find package "." in:
        /Users/xiazemin/source/k8s_learn/CRD/code-generator/vendor/k8s.io/apimachinery/pkg/runtime/serializer
generated/clientset/versioned/clientset.go:25:2: cannot find package "." in:
        /Users/xiazemin/source/k8s_learn/CRD/code-generator/vendor/k8s.io/client-go/discovery
generated/clientset/versioned/typed/samplecontroller/v1alpha1/foo.go:30:2: cannot find package "." in:
        /Users/xiazemin/source/k8s_learn/CRD/code-generator/vendor/k8s.io/client-go/rest
generated/clientset/versioned/clientset.go:27:2: cannot find package "." in:
        /Users/xiazemin/source/k8s_learn/CRD/code-generator/vendor/k8s.io/client-go/util/flowcontrol

https://www.jianshu.com/p/fc6744954102

 % go mod tidy


 % go test ./...                   
go: inconsistent vendoring in /Users/xiazemin/source/k8s_learn/CRD/code-generator:
        k8s.io/client-go@v0.22.2: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt

        To ignore the vendor directory, use -mod=readonly or -mod=mod.
        To sync the vendor directory, run:
                go mod vendor

%go mod vendor

% go test ./...   
# code-generator/generated/clientset/versioned/typed/samplecontroller/v1alpha1
generated/clientset/versioned/typed/samplecontroller/v1alpha1/foo.go:50:2: undefined: FooExpansion
# code-generator/generated/informers/externalversions/internalinterfaces
generated/informers/externalversions/internalinterfaces/factory_interfaces.go:31:27: undefined: versioned.Interface
FAIL    code-generator [build failed]
?       code-generator/api/samplecontroller/v1alpha1    [no test files]
?       code-generator/generated/clientset/versioned    [no test files]
?       code-generator/generated/clientset/versioned/fake       [no test files]
?       code-generator/generated/clientset/versioned/scheme     [no test files]
# code-generator/generated/clientset/versioned/typed/samplecontroller/v1alpha1/fake
generated/clientset/versioned/typed/samplecontroller/v1alpha1/fake/fake_foo.go:35:8: undefined: FakeSamplecontrollerV1alpha1


%vi api/samplecontroller/v1alpha1/types.go 
type FooExpansion struct{
}

% sh ./update-codegen.sh                                
./update-codegen.sh: line 26: ../vendor/k8s.io/code-generator/generate-groups.sh: Permission denied

% chmod -R 777 vendor


