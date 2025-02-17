package sjrw

import (
	"sort"
	"strings"
)

// getEditLineMap returns a map with only the difference changed from the original string
func getEditLineMap[targetType []byte | string](str targetType, editMapFromDiff map[string]map[int]string) map[int]string {
	content := strings.TrimSpace(string(str))
	contentLines := strings.Split(content, "\n")

	indexes := make(map[int]string, 0)

	for i, line := range contentLines {
		indexes[i+1] = line
	}

	rm := editMapFromDiff["rm"]
	RMNUMS := getKey(rm)
	sort.Ints(RMNUMS)

	add := editMapFromDiff["add"]
	ADDNUMS := getKey(add)
	sort.Ints(ADDNUMS)

	// remove
	for _, rmn := range RMNUMS {
		delete(indexes, rmn)
	}

	// add
	for _, addn := range ADDNUMS {
		indexes[addn] = editMapFromDiff["add"][addn]
	}

	return indexes
}

// GetKey returns keySlice of "m"
func getKey(m map[int]string) (keySlice []int) {
	keySlice = make([]int, len(m))
	for k := range m {
		keySlice = append(keySlice, k)
	}
	return keySlice
}
