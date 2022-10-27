package parser

import (
	"fmt"
	"os"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	if os.Getenv("MONKEY_PARSER_TRACING_ON") != "" {
		incIdent()
		tracePrint("BEGIN " + msg)
	}
	return msg
}

func untrace(msg string) {
	if os.Getenv("MONKEY_PARSER_TRACING_ON") != "" {
		tracePrint("END " + msg)
		decIdent()
	}
}
