# build

build server

    make server

build client

    make client

# use TLS certificate

create TLS certificate and private key

    openssl req -x509 -new -newkey rsa:2048 -nodes -keyout server.key -out server.csr -subj "/CN=*.grpc.example.com"

start server with TLS certificate

    ./helloworld-server -cert_file cert.pem -key_file privkey.pem

start client

    ./helloworld-client -cert_file cert.pem -addr grpc.example.com:50051

# plaintext

start server

    ./helloworld-server

start client

    ./helloworld-client

# build proto

install `protoc-gen-go`

    go get -u -v github.com/golang/protobuf
    go get -u -v github.com/golang/protobuf/proto
    go get -u -v github.com/golang/protobuf/protoc-gen-go

build

    cd internal
    protoc --go_out=plugins=grpc:. proto/helloworld/helloworld.proto
