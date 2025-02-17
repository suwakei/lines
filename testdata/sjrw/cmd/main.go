package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/suwakei/sjrw"
)

func main() {
	var sj sjrw.SjReader

	// カレントディレクトリを変える
	jsonPath5 := "../testdata/singleArr.json"
	f5, oerr := os.OpenFile(jsonPath5, os.O_RDONLY, 0o666)

	if oerr != nil {
		log.Fatal(oerr)
	}

	defer func() {
		err := f5.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath5)
		}
	}()

	ext := filepath.Ext(f5.Name())

	if ext == ".json" || ext == ".jsonc" {
		m, _ := sj.ReadAsMapFrom(f5)
		_ = m
	}
}
