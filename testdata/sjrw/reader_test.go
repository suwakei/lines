package sjrw

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/suwakei/sjrw/testdata"
)

const (
	acceptExt  string = ".json"
	acceptExt2 string = ".jsonc"
)

func BenchmarkReadAsStrfrom(b *testing.B) {
	var jsonPath string = "./testdata/readtest.json"
	var sj SjReader

	if _, err := os.Stat(jsonPath); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath)
	}

	f, err := os.OpenFile(jsonPath, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath)
		}
	}()

	if ext := filepath.Ext(f.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sj.ReadAsStrFrom(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkReadAsBytefrom(b *testing.B) {
	var jsonPath string = "./testdata/readtest.json"
	var sj SjReader

	if _, err := os.Stat(jsonPath); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath)
	}

	f, err := os.OpenFile(jsonPath, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath)
		}
	}()

	if ext := filepath.Ext(f.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sj.ReadAsByteFrom(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// func BenchmarkReadAsMapfrom(b *testing.B) {
// 	var jsonPath string = "./testdata/readtest4.json"
// 	var sj SjReader

// if ext := filepath.Ext(jsonPath); ext != acceptExt && ext != acceptExt2 {
// 	log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
// }

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 	_, err := sj.ReadAsMapFrom(jsonPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
// }

func TestReadAsStr(t *testing.T) {
	t.Helper()
	t.Parallel()

	var jsonPath1 string = "./testdata/readtest.json"
	var sj1 SjReader
	if _, err := os.Stat(jsonPath1); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath1)
	}

	f1, err := os.OpenFile(jsonPath1, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath1)
	}

	defer func() {
		err := f1.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath1)
		}
	}()

	if ext := filepath.Ext(f1.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	input1, _ := sj1.ReadAsStrFrom(f1)
	path1 := filepath.Base(f1.Name())

	var jsonPath2 string = "./testdata/readtest2.json"
	var sj2 SjReader

	if _, err := os.Stat(jsonPath2); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath2)
	}

	f2, err := os.OpenFile(jsonPath2, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath2)
	}

	defer func() {
		err := f2.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath2)
		}
	}()

	if ext := filepath.Ext(f2.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	input2, _ := sj2.ReadAsStrFrom(f2)
	path2 := filepath.Base(f2.Name())

	var jsonPath3 string = "./testdata/readtest3.json"
	var sj3 SjReader

	if _, err := os.Stat(jsonPath3); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath3)
	}

	f3, err := os.OpenFile(jsonPath3, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath3)
	}

	defer func() {
		err := f3.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath3)
		}
	}()

	if ext := filepath.Ext(f3.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	input3, _ := sj3.ReadAsStrFrom(f3)
	path3 := filepath.Base(f3.Name())

	var jsonPath4 string = "./testdata/readtest4.json"
	var sj4 SjReader

	if _, err := os.Stat(jsonPath4); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath4)
	}

	f4, err := os.OpenFile(jsonPath4, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath4)
	}

	defer func() {
		err := f4.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath4)
		}
	}()

	if ext := filepath.Ext(f4.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	input4, _ := sj4.ReadAsStrFrom(f4)
	path4 := filepath.Base(f4.Name())

	tests := map[string]struct {
		input    string
		expected string
	}{
		path1: {
			input1,
			testdata.Expected1,
		},

		path2: {
			input2,
			testdata.Expected2,
		},

		path3: {
			input3,
			testdata.Expected3,
		},

		path4: {
			input4,
			testdata.Expected4,
		},
	}

	for tname, tt := range tests {
		tt := tt
		t.Run("testReadAsStrFrom", func(t *testing.T) {
			t.Parallel()
			if tt.input != tt.expected {
				diff := cmp.Diff(tt.expected, tt.input)
				fmt.Printf("----------%q----------\n%s", tname, diff)
				t.Errorf("---------- %q these values are not same----------", tname)
			}
		})
	}
}

func TestReadAsByteFrom(t *testing.T) {
	t.Helper()
	t.Parallel()

	var jsonPath1 string = "./testdata/readtest.json"
	var sj1 SjReader

	if _, err := os.Stat(jsonPath1); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath1)
	}

	f1, err := os.OpenFile(jsonPath1, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath1)
	}

	defer func() {
		err := f1.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath1)
		}
	}()

	if ext := filepath.Ext(f1.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	input1, _ := sj1.ReadAsByteFrom(f1)
	path1 := filepath.Base(f1.Name())

	var jsonPath2 string = "./testdata/readtest2.json"
	var sj2 SjReader

	if _, err := os.Stat(jsonPath2); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath2)
	}

	f2, err := os.OpenFile(jsonPath2, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath2)
	}

	defer func() {
		err := f2.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath2)
		}
	}()

	if ext := filepath.Ext(f2.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	input2, _ := sj2.ReadAsByteFrom(f2)
	path2 := filepath.Base(f2.Name())

	var jsonPath3 string = "./testdata/readtest3.json"
	var sj3 SjReader

	if _, err := os.Stat(jsonPath3); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath3)
	}

	f3, err := os.OpenFile(jsonPath3, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath3)
	}

	defer func() {
		err := f3.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath3)
		}
	}()

	if ext := filepath.Ext(f3.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	input3, _ := sj3.ReadAsByteFrom(f3)
	path3 := filepath.Base(f3.Name())

	var jsonPath4 string = "./testdata/readtest4.json"
	var sj4 SjReader

	if _, err := os.Stat(jsonPath4); err != nil {
		log.Fatalf("this path is not exist \"%s\"", jsonPath4)
	}

	f4, err := os.OpenFile(jsonPath4, os.O_RDONLY, 0o666)

	if err != nil {
		log.Fatalf("could not open file \"%s\"", jsonPath4)
	}

	defer func() {
		err := f4.Close()
		if err != nil {
			log.Fatalf("could not close file \"%s\"", jsonPath4)
		}
	}()

	if ext := filepath.Ext(f4.Name()); ext != acceptExt && ext != acceptExt2 {
		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
	}

	input4, _ := sj4.ReadAsByteFrom(f4)
	path4 := filepath.Base(f4.Name())

	tests := map[string]struct {
		input    []byte
		expected []byte
	}{

		path1: {
			input1,
			[]byte(testdata.Expected1),
		},

		path2: {
			input2,
			[]byte(testdata.Expected2),
		},
		path3: {
			input3,
			[]byte(testdata.Expected3),
		},
		path4: {
			input4,
			[]byte(testdata.Expected4),
		},
	}
	for tname, tt := range tests {
		tt := tt
		t.Run("testReadAsStrFrom", func(t *testing.T) {
			t.Parallel()
			if string(tt.input) != string(tt.expected) {
				diff := cmp.Diff(string(tt.input), string(tt.expected))
				fmt.Printf("----------%q----------\n%s", tname, diff)
				t.Errorf("---------- %q these values are not same----------", tname)
			}
		})
	}
}

// func TestReadAsMapFrom(t *testing.T) {
// 	t.Helper()
// 	t.Parallel()
// 	var jsonPath5 string = "./testdata/readtest5.json"
// 	var sj5 SjReader

// 	if _, err := os.Stat(jsonPath5); err != nil {
// 		log.Fatalf("this path is not exist \"%s\"", jsonPath5)
// 	}

// 	f5, err := os.OpenFile(jsonPath5, os.O_RDONLY, 0o666)

// 	if err != nil {
// 		log.Fatalf("could not open file \"%s\"", jsonPath5)
// 	}

// 	defer func() {
// 		err := f5.Close()
// 		if err != nil {
// 			log.Fatalf("could not close file \"%s\"", jsonPath5)
// 		}
// 	}()

// 	if ext := filepath.Ext(f5.Name()); ext != acceptExt && ext != acceptExt2 {
// 		log.Fatalf("read file %q is not %q or %q file", ext, acceptExt, acceptExt2)
// 	}

// 	m, _ := sj5.ReadAsMapFrom(f5)
// 	fmt.Println(m)
// }
