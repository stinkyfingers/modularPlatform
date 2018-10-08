# Modular Platform #

This project is a skeleton for a modular platform, which runs and interacts with internal and/or user-developed modules using gRPC. 

## Directory Structure ##

#### cmd ####
_main.go_ - runs the platform.

_modularPlatform.conf_ - a yaml file that specifies 2 modules for the platform to run.

#### config ####
_config.go_ - parses the config file and populates the parameterized object.

#### platform ####
_server.go_ - runs the gRPC server and has a local "platformServer" object which implements the "PlatformServer" interface defined in the proto file, below.

_module.go_ - contains functions that fetch the module(s) via config and run them each as subprocesses via the "os/exec" package.

_modularPlatform.proto_ - the proto file which defines the PlatformServer service as well as Module and Details objects. Use Google's https://github.com/google/protobuf app to generate the Go protobuf file with a command like: 

```
protoc -I . modularPlatform.proto --go_out=plugins=grpc:.
```

#### modules ####

_go_example_ - This is a crude implementation of a "user" module in Go. The `cmd/main.go` package runs a gRPC client which reaches out to a local gRCP server on port 9999. The proto file is copied from `platform` and the `.pb.go` file generated as above. 


_js_example_ -  This is a crude implementation of a "user" module in NodeJS. The `client.js` module runs a gRPC client which reaches out to a local gRCP server on port  10000. The proto file is copied from `platform`. 

## What Next ##

In real life, there are some things you'd surely want to do which are omitted here for the sake of (some) clarity. 

- Implement TLS on the server, `platform/server.go` and each client (module). 
- Implement streaming data - either client, server, or two-way - between the platform server and each module. This will mean fleshing out the proto file according to your needs significantly. 
- Assuming you are developing a platform with the intent of permitting other developers create their own modules to run, you probably should look at developing SDK's for the various module languages you'd support - there is protobuf support for not only Go and NodeJS, but Java, C++, Python, C#, Ruby, etc. 
- Assure you've perused https://grpc.io/.


