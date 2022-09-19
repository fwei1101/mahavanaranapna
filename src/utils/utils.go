package utils

import (
	"bufio"
	"fmt"
	"mahavanaranapna/prompt"
	"os"
	"strings"
)

func ReadString(hint string) (str string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(hint)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSuffix(str, prompt.READ_STRING_SUFFIX)
	return
}

func SoundBackLn(sound string) {
	fmt.Println(sound)
}

func SoundBack(sound string) {
	fmt.Print(sound)
}
