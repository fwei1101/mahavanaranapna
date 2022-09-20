package utils

import (
	"bufio"
	"errors"
	"fmt"
	"mahavanaranapna/prompt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadString(hint string) (str string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(hint)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSuffix(str, prompt.READ_STRING_SUFFIX)
	return
}

func ReadInteger(hint string) (ret int64, err error) {
	str := ReadString(hint)
	r, _ := regexp.Compile("^[-+]{0,1}[0-9]+$")
	if r.MatchString(str) {
		ret, _ = strconv.ParseInt(str, 10, 64)
		err = nil
	} else {
		ret = 0
		err = errors.New("Invalid input, cannot be parsed as integer")
	}

	return
}

func SoundBackLn(sound string) {
	fmt.Println(sound)
}

func SoundBack(sound string) {
	fmt.Print(sound)
}
