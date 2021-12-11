% go run main.go 
# sigs.k8s.io/json/internal/golang/encoding/json
../../../../go/pkg/mod/sigs.k8s.io/json@v0.0.0-20211020170558-c049b76a60c6/internal/golang/encoding/json/encode.go:1249:12: sf.IsExported undefined (type reflect.StructField has no field or method IsExported)
../../../../go/pkg/mod/sigs.k8s.io/json@v0.0.0-20211020170558-c049b76a60c6/internal/golang/encoding/json/encode.go:1255:18: sf.IsExported undefined (type reflect.StructField has no field or method IsExported)



本机的k8s版本v1.21.5
kubernetes/test/images/agnhost/webhook/main.go


sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 h1:fD1pz4yfdADVNfFmcP2aBEtudwUQ1AlLnRBALr33v3s=
sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6/go.mod h1:p4QtZmO4uMYipTQNzagwnNoseA6OxSUutVw05NhYDRs=

版本太新

k8s.io/apimachinery v0.23.0
k8s.io/apimachinery v0.22.0


go get: downgraded k8s.io/apimachinery v0.23.0 => v0.22.0


% go run main.go                       
# command-line-arguments
./main.go:85:37: undefined: convertAdmissionRequestToV1
./main.go:87:10: undefined: convertAdmissionResponseToV1beta1
./main.go:110:18: undefined: codecs
./main.go:164:42: undefined: alwaysAllowDelayFiveSeconds
./main.go:168:42: undefined: alwaysDeny
./main.go:172:42: undefined: addLabel
./main.go:176:42: undefined: admitPods
./main.go:180:42: undefined: denySpecificAttachment
./main.go:184:42: undefined: mutatePods
./main.go:211:6: func main must have no arguments and no return values

https://github.com/kubernetes/apimachinery/tags

% go get -u k8s.io/apimachinery@v0.22.5-rc.0
go get: downgraded k8s.io/api v0.23.0 => v0.22.5-rc.0
go get: downgraded k8s.io/apimachinery v0.23.0 => v0.22.5-rc.0

% go run main.go                            
main.go:10:2: missing go.sum entry for module providing package k8s.io/api/admission/v1beta1; to add:
        go mod download k8s.io/api

        