package help

import (
	"mahavanaranapna/prompt"
	"mahavanaranapna/utils"
)

func Help() {
	utils.SoundBackLn("The tune I know of:")
	utils.SoundBack(prompt.KNOWN_TUNE)
}
