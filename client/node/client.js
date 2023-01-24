const path = require("path");
const grpc = require("grpc");
protoLoader = require('@grpc/proto-loader')

const PROTO_PATH = path.resolve(__dirname, "../../proto/*.proto");
const url = "localhost:";
const port = "8080";

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true,
    });

const msgServer = grpc.loadPackageDefinition(packageDefinition).MsgService;
const queryServer = grpc.loadPackageDefinition(packageDefinition).queryService;

const msgClient = new msgServer.MsgService(
    `${url}${port}`,
    grpc.credentials.createInsecure()
);

const queryClient = new queryServer.QueryService(
    `${url}${port}`,
    grpc.credentials.createInsecure()
);

module.exports = client;
