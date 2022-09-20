package silapna

import (
	"fmt"
	"mahavanaranapna/instruments/echo"
	"mahavanaranapna/instruments/help"
	"mahavanaranapna/prompt"
	"mahavanaranapna/utils"
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

	utils.SoundBackLn(prompt.HI)

	for {
		tune := utils.ReadString(prompt.ARANARA_FACE)
		switch tune {
		case prompt.ECHO:
			echo.Echo()
		case prompt.HELP:
			help.Help()
		case prompt.CEASE:
			goto cease
		default:
			utils.SoundBackLn(prompt.UNKNOWN)
		}
	}

cease:
	utils.SoundBack(prompt.BYE)
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		utils.SoundBack(".")
	}
}
