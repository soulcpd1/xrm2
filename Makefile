.PHONY: build clean tidy run

APP=xrm

build:
	mkdir -p build
	go build -o build/$(APP) ./cmd/xrm

clean:
	rm -rf build

tidy:
	go mod tidy

run:
	go run ./cmd/xrm
