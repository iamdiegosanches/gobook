build:
	@go build -o bin/gobook

run: build
	@./bin/gobook

test:
	@go test -v ./...

clean:
	@rm -f bin/gobook

.PHONY: build run test clean
