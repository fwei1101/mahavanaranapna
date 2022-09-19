package prompt

import "strings"

var KNOWN_TUNE = strings.Join([]string{
	ECHO,
	HELP,
	CEASE,
	VOID,
}, "\n")

const (
	ARANARA_FACE = "('~') "
)

const (
	ECHO       = "echo"
	ECHO_CEASE = "stop echo"
)

const (
	HELP = "help"
)

const (
	CEASE = "cease"
)

const (
	VOID    = ""
	BYE     = "(#~#) Bye"
	HI      = "(^o^) Welcome! Type \"help\" to see tunes I recognize"
	UNKNOWN = "(x_x) I do not recognize this tune... Type \"help\" to see what tunes I know"
)

var READ_STRING_SUFFIX = string([]byte{13, 10})
