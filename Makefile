.PHONY: build test clean lint fmt


default:
	all

all:
	fmt test build

run:
	@echo "[INFO] *****************run***********************"
	@go run main.go .

build:
	@echo "[INFO] *****************build***********************"
	@go build -ldflags="-s -w" -trimpath -o ./bin/lines.exe

test:
	@echo "[INFO] *****************test***********************"
	@go test -v .

ben:
	@echo "[INFO] *****************benchmark**********************"
	@go test -bench=. -benchmem

fmt:
	@echo "[INFO] ***********************formatting****************************"
	@go fmt ./...

clean:
	@echo "cleaned bin dir"
	@rm -rf ./bin