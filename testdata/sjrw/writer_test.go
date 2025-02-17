package sjrw

// import (
// 	"fmt"
// 	"testing"
// 	"strconv"
// 	"strings"
// 	"path/filepath"

// 	"github.com/suwakei/sjrw/internal"
// )

// func TestWriteFromStr(t *testing.T) {
//   t.Helper()
// 	var sj SjWriter
// 	str :=
//   `{
// 	"add": [
// 	    "git add .",
// 	    "git add README.md",
// 	    "git add .gitignore",
// 		"git add fmfm"
// 	],
// 	"commit": "git commit -m \"first commit\"",
// 	"push": [
// 	    "git push origin main",
// 	    "git push origin master",
// 	    "git push origin dev",
// 	    "git push origin release"
// 	],
// 	"reset": [
// 	    "git reset --soft",
// 	    "git reset --hard HEAD^"
// 	],
// 	"_id": "672d31b2240100375952cb1e",
// 	"index": 1,
// 	"guid": "98329871-09ba-4d38-99e3-8d808395bdfa",
// 	"isActive": false,
// 	"balance": "$2,832.86",
// 	"picture": "http://placehold.it/32x32",
// 	"age": 27,
// 	"eyeColor": "blue",
// 	"name": "Noemi Hays",
// 	"gender": "female",
// 	"company": "OBLIQ",
// 	"email": "noemihays@obliq.com",
// 	"phone": "+1 (940) 540-2100",
// 	"address": "368 Suydam Place, Ezel, Ohio, 7763",
// 	"about": "Pariatur dolore commodo ex aliqua tempor qui sit. Dolor incididunt nulla anim occaecat excepteur consectetur commodo officia voluptate tempor voluptate eiusmod officia ut. Officia labore aliquip fugiat amet aliquip excepteur et qui et laboris aliquip sunt occaecat. Dolore sint amet Lorem ea. Ipsum magna esse amet culpa sunt est nostrud id ut nostrud pariatur eu sit anim. Occaecat Lorem aute elit reprehenderit est reprehenderit velit exercitation qui amet. Enim consequat incididunt est velit ad.\r\n",
// 	"registered": "2023-03-12T06:12:32 -09:00",
// 	"latitude": 44.210044,
// 	"longitude": -165.285857,
// 	"tags": [
// 	    "exercitation",
// 	    "pariatur",
// 	    "amet",
// 	    "qui",
// 	    "deserunt"
// 	],
// 	"friends": [
// 	    {
// 	    "id": 0,
// 	    "name": "Christa Cole"
// 	    },
// 	    {
// 	    "id": 1,
// 	    "name": "Nona Knowles",
// 	    "age": 32,
// 	    "sex": "male",
// 	    "status": "student"
// 	    },
// 	    {
// 	    "id": 2,
// 	    "name": "Keisha Mosley"
// 	    }
// 	],
// 	"greeting": "Hello, Noemi Hays! You have 9 unread messages.",
// 	"favoriteFruit": "strawberry"
// 	}	`
// 	sj.WriteFromStr(str, "C:/Users/ns072/OneDrive/ドキュメント/Github/sjrw/testdata/writetest2.json")
// }

// func TestGetEditLineMap(t *testing.T) {
// 	targetStr, testEditMapFromDiff := writerTestSetup()

// 	fmt.Println("testEditMapFromDiff")
// 	fmt.Println(testEditMapFromDiff)

// 	content := strings.TrimSpace(string(targetStr))
// 	contentLines := strings.Split(content, "\n")

// 	indexes := make(map[int]string)
// 	for i, line := range contentLines {
// 		indexes[i + 1] = line
// 	}

// 	fmt.Println("indexes")
// 	for i := range len(indexes){
// 		fmt.Println(strconv.Itoa(i + 1) + indexes[i + 1])
// 	}

// 	resultIndexes := internal.GetEditLineMap(targetStr, testEditMapFromDiff)

// 	fmt.Print("\n")
// 	fmt.Println("resultIndexes")
// 	for i := range len(resultIndexes){
// 		fmt.Println(strconv.Itoa(i + 1) + resultIndexes[i + 1])
// 	}
// }

// func writerTestSetup() (string, map[string]map[int]string) {
// 	fmt.Println("----writerTestSetup start----")

// 	testEditMapFromDiff := make(map[string]map[int]string)

// 	// prepare two dimenton map for testing
// 	if _, ok := testEditMapFromDiff["rm"]; !ok {
// 		testEditMapFromDiff["rm"] = make(map[int]string)
// 	}

// 	testEditMapFromDiff["rm"][6] = `        "gitmfmf"`
// 	testEditMapFromDiff["rm"][11] = `        "git push origin master",`
// 	testEditMapFromDiff["rm"][17] = `        "git reset --hard HEAD^"`

// 	if _, ok := testEditMapFromDiff["add"]; !ok {
// 		testEditMapFromDiff["add"] = make(map[int]string)
// 	}

// 	testEditMapFromDiff["add"][2] = `    "status": "git status",`

// 	var sj SjReader

// 	str, _ := sj.ReadAsStrFrom("./testdata/writetest.json")

// 	return str, testEditMapFromDiff
// }

// func TestDiff(t *testing.T) {
// 	var jsonPath1 string = "./testdata/difftest.json"
// 	var jsonPath2 string = "./testdata/difftest2.json"
// 	var sj1 SjReader
// 	var sj2 SjReader
// 	s1, _ := sj1.ReadAsBytesFrom(jsonPath1)
// 	s2, _ := sj2.ReadAsBytesFrom(jsonPath2)
// 	editMapFromDiff, _ := internal.DiffReturn(filepath.Base(jsonPath2), s2, filepath.Base(jsonPath1), s1)
// 	fmt.Println(editMapFromDiff)
// }

// func writerTearDown() {

// }
