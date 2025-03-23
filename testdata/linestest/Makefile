.PHONY: build test clean lint fmt


default:
	all

all:
	fmt test build

# add -race option
build:
	@echo "[info *****************build***********************]""
	@go build -ldflags="-s -w" -trimpath -o ./bin/steps.exe

test:
	@go test -v .

fmt:
	@echo "[info ***********************formatting****************************]"
	@go fmt ./...

clean:
	@echo "cleaned bin dir"
	@rm -rf ./bin