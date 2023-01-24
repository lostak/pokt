# MUST CALL FROM BASE DIR FOR PATHS
echo -e "\nDelete existing .pb.go files\n"
find . -name "\n*.pb.go" -type f -delete
echo -e "\nBuild proto for tx.proto and protfolio.proto\n"
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/pokt/tx.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/pokt/query.proto
protoc --proto_path=proto/pokt --go_out=types --go_opt=paths=source_relative portfolio.proto 
echo -e "\nMove generated .pb.go files to types and server folders\n"
mv proto/pokt/tx_grpc.pb.go server/tx_grpc.pb.go
mv proto/pokt/tx.pb.go server/tx.pb.go
mv proto/pokt/query_grpc.pb.go server/query_grpc.pb.go
mv proto/pokt/query.pb.go server/query.pb.go