.PHONY: build build-pi deps-save deps-restore

default: build

build:
	@echo "=> Building all binaries..."
	godep go build -o bin/bell	cmd/bell/main.go
	godep go build -o bin/meter	cmd/meter/main.go
	@echo "=> Done."

build-pi:
	@echo "=> Building all binaries for Raspberry PI..."
	GOOS=linux GOARCH=arm GOARM=6 godep go build -o bin/bell-linux-arm6 cmd/bell/main.go
	GOOS=linux GOARCH=arm GOARM=6 godep go build -o bin/meter-linux-arm6 cmd/meter/main.go
	@echo "=> Done."

deps-save:
	godep save ./...

deps-restore:
	godep restore
