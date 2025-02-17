.PHONY: build test clean static_analysis lint vet fmt chkfmt


default:
	all

all:
	@echo "[info] ***********************exec test fmt bench****************************"
	@t fmt b


b:
	@echo "[info] ***********************benchmark****************************"
	@go test -bench=. -benchmem

bp:
	@echo "[info] ***********************benchmark cpuprof****************************"
	@go test -bench=. -benchmem -cpuprofile cpu.prof

bm:
	@echo "[info] ***********************benchmark memprof****************************"
	@go test -bench=. -benchmem -memprofile mem.prof

bpm:
	@echo "[info] ***********************benchmark memprof cpuprof****************************"
	@go test -bench=. -benchmem -memprofile mem.prof -cpuprofile cpu.prof

tm:
	@echo "[info] ***********************testing ReadAsMapFrom****************************"
	@go test -run TestReadAsMap

ts:
	@echo "[info] ***********************testing ReadAsStrFrom****************************"
	@go test -run TestReadAsStr

tb:
	@echo "[info] ***********************testing ReadAsByteFrom****************************"
	@go test -run TestReadAsByte


t:
	@echo "[info] ***********************testing****************************"
	@go test -v .

fmt:
	@echo "[info] ***********************formatting****************************"
	@go fmt ./...