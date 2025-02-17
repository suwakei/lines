package sjrw

// import (
// 	"bufio"
// 	"io"
// 	"log"
// 	"os"
// 	"sort"
// 	"strings"
// )

// type SjWriter struct {
// }

// WriteFromStr writes str to filepathToWrite
// func (SjWriter) WriteFromStr(str string, filepathToWrite io.ReadWriter) {
// 	from := []byte(str)
// 	var to []byte

// 	f, err := os.OpenFile(filepathToWrite, os.O_RDWR, 0o666)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	reader := bufio.NewReaderSize(f, 24 * 1024)

// 	for {
// 		readByte, _, err := reader.ReadLine()
// 		to = append(to, append(readByte, []byte("\n")...)...)

// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			panic(err)
// 		}

// 	}
// 	mapFromDiff, _ := internal.DiffReturn("from", from, "to", to)

// 	result := internal.GetEditLineMap[string](str, mapFromDiff)

// 	resultkey := internal.GetKey(result)
// 	sort.Ints(resultkey)

// 	var sb strings.Builder
// 	for rk := range resultkey {
// 		sb.WriteString(result[rk] + "\n")
// 	}
// 	writer := bufio.NewWriter(f)
// 	writer.WriteString(sb.String())
// 	writer.Flush()
// }

// func (SjWriter) WriteFromByte(byteSlice []byte, filePath io.Writer) {
// }

// func (SjWriter) WriteFromMap(writemap map[string]any, filePath io.Writer) {
// }

// func (sj *SjWriter) WriteFromTextFile(readFile, writeFile string) {
// }
