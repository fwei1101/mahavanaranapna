package echo

import (
	"fmt"
	"mahavanaranapna/utils"
)

func Echo() {
	readint, err := utils.ReadInteger("input an int: ")
	fmt.Printf("read %d, type %T, err is %v\n", readint, readint, err)
}
