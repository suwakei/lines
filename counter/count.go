package counter

import (

)

type CntResult struct {
	Filetype string
	Steps int
	Blanks int
	Comments int
	Files int
	Bytes int
}

func Count(file string) (CntResult, error) {

}
