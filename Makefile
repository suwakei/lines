.PHONY: build test clean lint fmt


default:
	all

all:
	fmt test build

# add -race option
build:
	@echo "[INFO] *****************build***********************"
	@go build -ldflags="-s -w" -trimpath -o ./bin/steps.exe

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