module github.com/myk4040okothogodo/ResourceSheduler

go 1.17

replace github.com/myk4040okothogodo/ResourceSheduler/protos/assetsdefinition/protos => ./protos/assetsdefinition/protos

replace github.com/myk4040okothogodo/ResourceSheduler/server => ./server

require (
	github.com/hashicorp/go-hclog v1.1.0
	github.com/myk4040okothogodo/ResourceSheduler/protos/assetsdefinition/protos v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/ResourceSheduler/server v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.44.0
)

require (
	github.com/fatih/color v1.7.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)
