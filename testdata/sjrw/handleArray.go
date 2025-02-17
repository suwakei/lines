package sjrw

import (
	"strings"
	"fmt"
)

func (a *assemble) handleArray(inputRune []rune, key string) {
	var (
		curToken   rune
		peekToken  rune
		dc         uint8
		inQuote    bool
		ss         string
		arrVal     any
		rs         []any
		arrBuf strings.Builder
		firstLoop  bool = true
	)
	// preallocate memory
	estimatedLen := arrLength(a.idx, inputRune)
	rs = make([]any, 0, estimatedLen)
	arrBuf.Grow(15)

	for ; a.idx < uint(len(inputRune)); a.idx++ {
		curToken = inputRune[a.idx]
		if a.idx+1 < uint(len(inputRune)) {
			peekToken = inputRune[a.idx+1]
		}

		// FIXME for debug
		fmt.Println("idx", a.idx, "lineCount", a.lineCount, "curToken", string(curToken))
		fmt.Println("idx", a.idx+1, "lineCount", a.lineCount, "peekToken", string(peekToken))

		if firstLoop {
			firstLoop = false
			continue
		}

		switch curToken {
		case SPACE, TAB:
			if inQuote {
				arrBuf.WriteRune(curToken)
			} else if !inQuote {
				a.ignoreSpaceTab(inputRune)
			}

		case DOUBLEQUOTE:
			dc++
            arrBuf.WriteRune(curToken)
            if dc == 1 {
                inQuote = true
            } else if dc == 2 {
                dc = 0
                inQuote = false
            }

		case BACKSLASH:
			arrBuf.WriteRune(curToken)
			if peekToken == DOUBLEQUOTE {
				dc--
			}

		case SLASH:
			if inQuote {
				arrBuf.WriteRune(curToken)
			} else if peekToken == SLASH && dc == 0 {
				a.ignoreComments(inputRune)
			}

		case LBRACKET:
			if inQuote {
				arrBuf.WriteRune(curToken)
			} else if !inQuote {
				a.handleArray(inputRune, key)
				arrBuf.Reset()
			}

		case RBRACKET:
			if inQuote {
				arrBuf.WriteRune(curToken)
			} else if !inQuote {
				ss = arrBuf.String()
				arrVal = determineType(ss)
				rs = append(rs, arrVal)
				arrBuf.Reset()
				a.assembledMap[a.lineCount][key] = rs
				return 
			}

		case COMMA:// TODO commaとlbracket が続いたときいらない処理をしてしまう
			if inQuote {
				arrBuf.WriteRune(curToken)
			} else if !inQuote {
				ss = arrBuf.String()
				arrVal = determineType(ss)
				rs = append(rs, arrVal)
				arrBuf.Reset()
				a.assembledMap[a.lineCount][key] = rs
				return
			}

		case LBRACE:
			if inQuote {
				arrBuf.WriteRune(curToken)
			} else if !inQuote {
				a.handleObject(inputRune, key)
			}

		case lrTOKEN:
			if inQuote {
				arrBuf.WriteRune(curToken)
			} else if !inQuote {
				if peekToken == lnTOKEN {
					continue
				}
				a.lineCount++
			}

		case lnTOKEN:
			if inQuote {
				arrBuf.WriteRune(curToken)
			} else if !inQuote {
				a.lineCount++
			}

		default:
			arrBuf.WriteRune(curToken)
		}
	}
}

func arrLength(idx uint, inputRune []rune) uint {
	var (
		curToken  rune
		peekToken rune
		dc        uint8 = 0
		inQuote   bool = false
		arrLength uint  = 0
		lb        uint8 = 0
		rb        uint8 = 0
	)

	for ; idx < uint(len(inputRune)); idx++ {
		curToken = inputRune[idx]
		if idx+1 < uint(len(inputRune)) {
			peekToken = inputRune[idx+1]
		}

		switch curToken {
		case DOUBLEQUOTE:
			dc++
            if dc == 1 {
                inQuote = true
            } else if dc == 2 {
                dc = 0
                inQuote = false
            }

		case BACKSLASH:
			if peekToken == DOUBLEQUOTE {
				dc--
			}

		case COMMA:
			if inQuote {
				arrLength++
			}

		case LBRACKET:
			if inQuote {
				lb++
			}

		case RBRACKET:
			if inQuote {
				rb++
			}
			if lb == rb {
				return arrLength + 1
			}
		}
	}
	return arrLength
}
