.PHONY: build pi

default: build

build:
	@echo "=> Building all binaries..."
	godep go build -o bin/echo cmd/echo/main.go
	godep go build -o bin/electra cmd/electra/main.go
	godep go build -o bin/poseidon cmd/poseidon/main.go
	godep go build -o bin/zeus cmd/zeus/main.go
	@echo "=> Done."

pi:
	@echo "=> Building all binaries for Raspberry PI..."
	GOOS=linux GOARCH=arm GOARM=6 godep go build -o bin/echo-linux-arm6 cmd/echo/main.go
	GOOS=linux GOARCH=arm GOARM=6 godep go build -o bin/electra-linux-arm6 cmd/electra/main.go
	GOOS=linux GOARCH=arm GOARM=6 godep go build -o bin/poseidon-linux-arm6 cmd/poseidon/main.go
	@echo "=> Done."

ship: pi
	scp bin/electra-linux-arm6 pi@zeus:~/electra
