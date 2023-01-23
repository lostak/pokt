# Delete existing .pb.go files
find . -name "*.pb.go" -type f -delete
# Build proto for tx.proto and protfolio.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/pokt/tx.proto
protoc --proto_path=proto/pokt --go_out=types --go_opt=paths=source_relative portfolio.proto 
# Move generated .pb.go files to types and server folders
mv proto/pokt/tx_grpc.pb.go server/tx_grpc.pb.go
mv proto/pokt/tx.pb.go server/tx.pb.go