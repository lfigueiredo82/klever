# protoc --go_out=../../../../ --proto_path=../api/proto/ ../api/proto/model/*.proto


 protoc --go_out=../ --proto_path=../api/proto/ --go_opt=paths=import \
    --go-grpc_out=../ --go-grpc_opt=paths=import \
    ../api/proto/model/*.proto