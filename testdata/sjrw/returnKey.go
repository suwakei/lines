package sjrw

import (
	"strings"
)

func (a *assemble) returnKey(inputRune []rune) (key string) {
	var (
		dc        uint8
		curToken  rune
		peekToken rune
		keyBuf    strings.Builder
	)
	// preallocate memory
	keyBuf.Grow(20)

	for ; ; a.idx++ {
		curToken = inputRune[a.idx]
		peekToken = inputRune[a.idx+1]

		switch curToken {
		case SPACE, TAB:
			if dc > 0 {
				keyBuf.WriteRune(curToken)
			} else if dc == 0 {
				a.ignoreSpaceTab(inputRune)
			}

		case DOUBLEQUOTE:
			keyBuf.WriteRune(curToken)
			dc++
			if dc == 2 {
				dc = 0
			}

		case BACKSLASH:
			keyBuf.WriteRune(curToken)
			if dc > 0 && peekToken == DOUBLEQUOTE {
				dc--
			}

		case SLASH:
			if dc > 0 {
				keyBuf.WriteRune(curToken)
			} else if dc == 0 && peekToken == SLASH {
				a.ignoreComments(inputRune)
			}

		case COLON:
			if dc > 0 {
				keyBuf.WriteRune(curToken)
			} else if dc == 0 {
				key = keyBuf.String()
				keyBuf.Reset()
				return key
			}

		default:
			if dc > 0 {
				keyBuf.WriteRune(curToken)
			}
		}
	}
}
