mkdir code-generator && cd code-generator
 go mod init code-generator
 mkdir -p api/samplecontroller/v1alpha1 && cd api/samplecontroller/v1alpha1
 vi doc.go
 vi types.go
 
mkdir hack && cd hack
vi tools.go
vi update-codegen.sh
vi boilerplate.go.txt

go mod vendor
chmod -R 777 vendor

cd hack && sh ./update-codegen.sh
#cd ../ && tree -L 3

cd api/samplecontroller/v1alpha1/
vi register.go

vi main_test.go
go test ./... 
