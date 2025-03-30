.PHONY: build test clean lint fmt


default:
	all

all:
	fmt test winamdb


winamdb:
	@echo "[INFO] *****************Windows_build***********************"
	@echo "[INFO] *****************OS=windows ARCH=amd64***********************"
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ./bin/lines.exe

winarmb:
	@echo "[INFO] *****************Windows_build***********************"
	@echo "[INFO] *****************OS=windows ARCH=arm64***********************"
	@GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o ./bin/lines.exe

macamdb:
	@echo "[INFO] *****************macOS_build***********************"
	@echo "[INFO] *****************OS=darwin ARCH=amd64***********************"
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ./bin/lines

macarmb:
	@echo "[INFO] *****************macOS_build***********************"
	@echo "[INFO] *****************OS=darwin ARCH=arm64***********************"
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o ./bin/lines

linuxarmb:
	@echo "[INFO] *****************Linux_build***********************"
	@echo "[INFO] *****************OS=linux ARCH=arm64***********************"
	@GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o ./bin/lines

linuxamdb:
	@echo "[INFO] *****************Linux_build***********************"
	@echo "[INFO] *****************OS=linux ARCH=amd64***********************"
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ./bin/lines

test:
	@echo "[INFO] *****************test***********************"
	@go test -v ./...

ben:
	@echo "[INFO] *****************benchmark**********************"
	@go test -bench=. -benchmem

fmt:
	@echo "[INFO] ***********************formatting****************************"
	@go fmt ./...

clean:
	@echo "cleaned bin dir"
	@rm -rf ./bin