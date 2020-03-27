.PHONY: all
all: server client

.PHONY: clean
clean:
	rm -rf helloworld-server helloworld-client

server:
	go build -o helloworld-server cmd/helloworld-server/main.go

client:
	go build -o helloworld-client cmd/helloworld-client/main.go


