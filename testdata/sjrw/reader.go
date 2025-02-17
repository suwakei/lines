package sjrw

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"unsafe"
)

type SjReader struct {
}

// ReadAsStr returns json content as string
func (SjReader) ReadAsStrFrom(readFile io.Reader) (contentAsStr string, err error) {
	if readFile != nil {
		reader := new(bytes.Buffer)
		_, err := reader.ReadFrom(readFile)
		if err != nil {
			return "", err
		}
		contentAsStr = unsafe.String(unsafe.SliceData(bytes.TrimSpace(reader.Bytes())), len(bytes.TrimSpace(reader.Bytes())))
	}
		return contentAsStr, nil
	}

func (SjReader) ReadAsByteFrom(readFile io.Reader) (contentAsByte []byte, err error) {
	if readFile != nil {
		reader := new(bytes.Buffer)
		_, err := reader.ReadFrom(readFile)
		if err != nil {
			return []byte(""), err
		}
	contentAsByte = bytes.TrimSpace(reader.Bytes())	
	}
	return contentAsByte, nil
}

func (SjReader) ReadAsMapFrom(readFile io.Reader) (contentAsMap map[uint]map[string]any, err error) {
	var jsonRune []rune = make([]rune, 0, 800)

	if readFile != nil {
		reader := bufio.NewReaderSize(readFile, 24*1024)
		runeSlice := make([]rune, 0, 400)

		for {
			readRune, _, err := reader.ReadRune()

			jsonRune = append(jsonRune, append(runeSlice, readRune)...)

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}
		}
	}
	var a assemble
	contentAsMap, err = a.assembleMap(jsonRune)
	if err != nil {
		return nil, err
	}
	return contentAsMap, nil
}
