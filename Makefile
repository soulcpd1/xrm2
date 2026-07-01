APP=xrm

build:
	go build -o build/$(APP) ./cmd/xrm

clean:
	rm -rf build

tidy:
	go mod tidy

run:
	go run ./cmd/xrm
