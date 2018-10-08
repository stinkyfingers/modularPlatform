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
var platform = grpc.loadPackageDefinition(packageDefinition).platform;
var client = new platform.Platform('localhost:10000',
                                       grpc.credentials.createInsecure());



function register() {
	var err = client.RegisterModule({ name:"test_js_module", port: "10000" }, (err, mo) => {
		if (err) console.log("oh no", err);
		console.log(mo)
	});
}

register();