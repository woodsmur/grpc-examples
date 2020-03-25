# build

build server

    make server

build client

    make client

# use TLS certificate

create TLS certificate and private key

    openssl req -x509 -nodes -newkey rsa:2048 -days 365 \
        -keyout privkey.pem -out cert.pem -subj "/CN=grpc.example.com"

write line to `/etc/hosts`

    127.0.0.1 grpc.example.com

start server with TLS certificate

    ./helloworld-server -cert_file cert.pem -key_file privkey.pem

start client

    ./helloworld-client -cert_file cert.pem -addr grpc.example.com:50051

# plaintext

start server

    ./helloworld-server

start client

    ./helloworld-client
