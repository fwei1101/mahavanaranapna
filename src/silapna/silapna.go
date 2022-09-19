package silapna

import (
	"bufio"
	"fmt"
	prompt "mahavanaranapna/prompt"
	"os"
	"strings"
	"time"
)

func Welcome() {
	fmt.Println(" /=============================================================================================\\")
	fmt.Println("|:                _                                                                            :|")
	fmt.Println("|:    /\\/\\   __ _| |__   __ ___   ____ _ _ __   __ _ _ __ __ _ _ __   __ _ _ __  _ __   __ _   :|")
	fmt.Println("|:   /    \\ / _` | '_ \\ / _` \\ \\ / / _` | '_ \\ / _` | '__/ _` | '_ \\ / _` | '_ \\| '_ \\ / _` |  :|")
	fmt.Println("|:  / /\\/\\ | (_| | | | | (_| |\\ V | (_| | | | | (_| | | | (_| | | | | (_| | |_) | | | | (_| |  :|")
	fmt.Println("|:  \\/    \\/\\__,_|_| |_|\\__,_| \\_/ \\__,_|_| |_|\\__,_|_|  \\__,_|_| |_|\\__,_| .__/|_| |_|\\__,_|  :|")
	fmt.Println("|:                                                                        |_|                  :|")
	fmt.Println(" \\=============================================================================================/")
	fmt.Println("")
}

func Instrument() {
	soundBackLn("Welcome! You can type:")
	soundBackLn(prompt.ECHO)
	soundBackLn(prompt.CEASE)
	soundBackLn(prompt.VOID)

	for {
		tune := readString(prompt.INSTRUMENT)
		switch tune {
		case prompt.CEASE:
			goto cease
		default:
			soundBackLn("Oh... unknown tune...")
		}
	}

cease:
	soundBack(prompt.BYE)
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		soundBack(".")
	}
}

func readString(hint string) (str string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(hint)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSuffix(str, prompt.READ_STRING_SUFFIX)
	return
}

func soundBackLn(sound string) {
	fmt.Println(sound)
}

func soundBack(sound string) {
	fmt.Print(sound)
}
