// IGNORE - unused code

var PROTO_PATH = __dirname + '/modularPlatform.proto';
var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });

var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
// The protoDescriptor object has the full package hierarchy
var platform = protoDescriptor.platform;


function Stop(module){
}

function RegisterModule(module) {
	console.log("register module ", module.Name)
}

function getServer() {
  var server = new grpc.Server();
  server.addService(platform.Platform.service, {
    Stop: Stop,
    RegisterModule: RegisterModule
  });
  return server;
}
var server = getServer();
server.bind('0.0.0.0:10000', grpc.ServerCredentials.createInsecure());
server.start();