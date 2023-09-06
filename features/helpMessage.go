package features

import (
	"fmt"
	"github.com/fatih/color"
)

var Message string = fmt.Sprintln(`Gocron Project:
-------------------------------------------------------
is a API for learning GO language with example.
-------------------------------------------------------
my name is Sina LalehBakhsh, I hope this API is useful for you
after running program, write your single word about any of GO language.
if your perpuse is more than one word, for convenience searching, just write keywords.
like this:
	map slice
-------------------------------------------------------`)

func HelpMessage(){
	color.HiBlue(Message)
}

