package sjrw

import (
	"fmt"
	"strings"
)

// returnObj returns map[string]any
func (a *assemble) handleObject(inputRune []rune, key string) {
	var (
		curToken  rune // The target token.
		peekToken rune // The token for confirmation of next character.

		dc        uint8        // Counter for number of ".
		inQuote   bool         = false
		keyMode   bool  = true // If true, mode which read json key. if false read json value.
		firstLoop bool  = true // First loop flag.

		keyBuf strings.Builder // When in "keyMode" is true, buf for accumulating key token.
		valBuf strings.Builder // When in "keyMode" is false, buf for accumulating value token.
		value  any             = nil
		rm     map[string]any
	)

	// preallocation of memory
	rm = make(map[string]any, mapLength(a.idx, inputRune))
	keyBuf.Grow(20)
	valBuf.Grow(30)

	for ; ; a.idx++ {
		curToken = inputRune[a.idx]
		peekToken = inputRune[a.idx+1]

		// FIXME for debug
		fmt.Println("idx", a.idx, "lineCount", a.lineCount, "curToken", string(curToken))
		fmt.Println("idx", a.idx+1, "lineCount", a.lineCount, "peekToken", string(peekToken))

		if firstLoop {
			firstLoop = false
			continue
		}

		switch curToken {
		case SPACE, TAB:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
				} else if !inQuote {
					a.ignoreSpaceTab(inputRune)
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
				} else if !inQuote {
					a.ignoreSpaceTab(inputRune)
				}
			}

		case DOUBLEQUOTE:
			dc++
			if keyMode {
				keyBuf.WriteRune(curToken)
				if dc == 2 {
					dc = 0
					inQuote = false
					continue
				}
			}

			if !keyMode {
				valBuf.WriteRune(curToken)
				if dc == 2 {
					dc = 0
					inQuote = false
					continue
				}
			}

		case BACKSLASH:
			if keyMode {
				keyBuf.WriteRune(curToken)
				if peekToken == DOUBLEQUOTE {
					dc--
					continue
				}
				continue
			}

			if !keyMode {
				valBuf.WriteRune(curToken)
				if peekToken == DOUBLEQUOTE {
					dc--
					continue
				}
				continue
			}

		case SLASH:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					a.ignoreComments(inputRune)
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					a.ignoreComments(inputRune)
					continue
				}
			}

		case COLON:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
				} else if !inQuote {
					key = keyBuf.String()
					keyBuf.Reset()
					keyMode = false
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
					continue
				}
			}

		case LBRACE:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				}
			}

			if !keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					a.handleObject(inputRune, key)
					continue
				}
			}

		case RBRACE:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					if value != nil {
						rm[key] = value
						value = nil

					} else {
						ss := valBuf.String()
						valBuf.Reset()
						value = nil

						if ss != "" {
							value = determineType(ss)
							rm[key] = value

						} else if ss == "" {
							rm[key] = ss
						}
					}
					a.assembledMap[a.lineCount][key] = rm
					return 
				}
			}

		case LBRACKET:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					a.handleArray(inputRune, key)
					valBuf.Reset()
					continue
				}
			}

		case RBRACKET:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
					continue
				}
			}

		case COMMA:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
				} else if !inQuote {
					if value != nil {
						rm[key] = value
						keyMode = true

					} else {
						ss := valBuf.String()
						valBuf.Reset()
						value = nil

						if ss != "" {
							value = determineType(ss)
							rm[key] = value
							keyMode = true

						} else if ss == "" {
							rm[key] = ss
							keyMode = true
						}
					}
				}
			}

		case lrTOKEN:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					if peekToken == lnTOKEN {
						continue
					}
					a.lineCount++
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					if peekToken == lnTOKEN {
						continue
					}
					a.lineCount++
					continue
				}
			}

		case lnTOKEN:
			if keyMode {
				if inQuote {
					keyBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					a.lineCount++
					continue
				}
			}

			if !keyMode {
				if inQuote {
					valBuf.WriteRune(curToken)
					continue
				} else if !inQuote {
					a.lineCount++
					continue
				}
			}

		default:
			if keyMode {
				keyBuf.WriteRune(curToken)
			}

			if !keyMode {
				valBuf.WriteRune(curToken)
			}
		}
	}
}

func mapLength(idx uint, inputRune []rune) uint {
	var (
		curToken  rune
		peekToken rune
		dc        uint8
		mapLength uint
		lb        uint8
		rb        uint8
		inQuote   bool = false
	)

	for ; lb != rb; idx++ {
		curToken = inputRune[idx]
		peekToken = inputRune[idx+1]

		switch curToken {
		case DOUBLEQUOTE:
			dc++
			if dc == 2 {
				dc = 0
				inQuote = false
			}

		case BACKSLASH:
			if peekToken == DOUBLEQUOTE {
				dc--
			}

		case COLON:
			if !inQuote {
				mapLength++
			}

		case LBRACE:
			if !inQuote {
				lb++
			}

		case RBRACE:
			if !inQuote {
				rb++
			}
		}
	}
	return mapLength + 2
}
