package sjrw

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SPACE       rune = ' '  // U+0020 32
	TAB         rune = '\t' // U+0009 9
	lnTOKEN     rune = '\n' // U+000A 10
	lrTOKEN     rune = '\r' // U+000D 13
	DOUBLEQUOTE rune = '"'  // U+0022 34
	COLON       rune = ':'  // U+003A 58
	LBRACE      rune = '{'  // U+007B 123
	RBRACE      rune = '}'  // U+007D 125
	LBRACKET    rune = '['  // U+005B 91
	RBRACKET    rune = ']'  // U+005D 93
	COMMA       rune = ','  // U+002C 44
	BACKSLASH   rune = '\\' // U+005C 92
	SLASH       rune = '/'  // U+002F 47
	ASTERISK    rune = '*'  // U+002A 42
)

type assemble struct {
	idx       uint
	lineCount uint // Counter for current number of line.
	assembledMap map[uint]map[string]any
}

// AssembleMap returns map created by input []rune
func (a *assemble) assembleMap(inputRune []rune) (map[uint]map[string]any, error) {
	var (
		curToken   rune                        // The target token.
		peekToken  rune                        // The next target token.
		runeLength uint = uint(len(inputRune)) // The length of input rune slice.

		firstLoop bool = true // First loop flag.
		keyMode   bool = true //  If true, read jsonKey.

		returnedKey   string
	)

	if runeLength == 0 {
		return nil, fmt.Errorf("empty input")
	}

	a.lineCount = 1

	// preallocation of memory.
	initialCap := lnNum(inputRune)
	a.assembledMap = make(map[uint]map[string]any, initialCap)

	for ; a.idx < runeLength-1; a.idx++ {
		curToken = inputRune[a.idx]
		peekToken = inputRune[a.idx+1]

		// TODO for debug
		fmt.Println("idx", a.idx, "lineCount", a.lineCount, "curToken", string(curToken))
		fmt.Println("idx", a.idx+1, "lineCount", a.lineCount, "peekToken", string(peekToken))

		switch {
		case firstLoop:
			a.initializeFirstEntry(a.lineCount, curToken)
			firstLoop = false
			continue

		case curToken == SLASH:
			if err := a.ignoreComments(inputRune); err != nil {
				return nil, err
			}

			continue

		case isNewline(curToken):
			a.lineCount++
			continue

		case isIgnores(curToken):
			continue

		case keyMode:
			returnedKey = a.returnKey(inputRune)
			keyMode = false
			continue

		case !keyMode:
			err := a.handleValue(inputRune, returnedKey, curToken)
			if err != nil {
				return nil, err
			}
			keyMode = true
			continue
		}
	}
	return a.assembledMap, nil
}

func (a *assemble) initializeFirstEntry(lineCount uint, curToken rune) {
	if _, ok := a.assembledMap[lineCount]; !ok {
		a.assembledMap[lineCount] = make(map[string]any, 1)
	}
	a.assembledMap[lineCount][string(curToken)] = ""
}



func (a *assemble) handleValue(inputRune []rune, key string, curToken rune) error {
	switch curToken {
	case LBRACKET:
		a.handleArray(inputRune, key)
	case LBRACE:
		a.handleObject(inputRune, key)
	default:
		a.handlePrimitive(inputRune, key)
	}
	return nil
}

func (a *assemble) ignoreComments(inputRune []rune) error {
	var (
		curToken  rune = inputRune[a.idx]
		peekToken rune = inputRune[a.idx+1]
	)

	if curToken == SLASH && peekToken == SLASH {
		for ; ; a.idx++ {
			peekToken = inputRune[a.idx+1]
			if isNewline(peekToken) {
				break
			}
		}

	} else if curToken == SLASH && peekToken == ASTERISK {
		for ; ; a.idx++ {
			curToken = inputRune[a.idx]
			peekToken = inputRune[a.idx+1]

			if isNewline(curToken) {
				a.lineCount++
				continue
			}

			if curToken == ASTERISK && peekToken == SLASH {
				a.idx += 1
				break
			}
		}
	}
	return nil
}

func (a *assemble) ignoreSpaceTab(inputRune []rune) {
	var (
		curToken rune = inputRune[a.idx]
	)
	for ; ; a.idx++ {
		curToken = inputRune[a.idx]
		if curToken == SPACE || curToken == TAB {
			continue
		}
		break
	}
}

func isNewline(r rune) bool {
	return r == lrTOKEN || r == lnTOKEN
}

func determineType(ss string) any {
	if num, err := strconv.Atoi(ss); err == nil {
		return num

	} else if tr := strings.TrimSpace(ss); tr == "true" || tr == "false" {
		b, _ := strconv.ParseBool(tr)
		return b

	} else {
		return ss
	}
}

// "lnNum" returns the number of "\n" or "\r" from "r".
// this return value used for initializing memory of "initMap"
func lnNum(r []rune) uint {
	var lnCount uint = 0
	var dc uint8 = 0

	for i := 0; i < len(r); i++ {
		if dc > 0 && r[i] == BACKSLASH && r[i+1] == DOUBLEQUOTE {
			dc--
		}

		if r[i] == DOUBLEQUOTE {
			dc++
			if dc == 2 {
				dc = 0
			}
		}

		if dc == 0 && r[i] == lrTOKEN {
			lnCount++
		}
	}
	return lnCount + 1
}

func isIgnores(curToken rune) bool {
	if curToken == SPACE {
		return true
	}
	if curToken == TAB {
		return true
	}
	if curToken == COMMA {
		return true
	}
	if curToken == COLON {
		return true
	}
	return false
}