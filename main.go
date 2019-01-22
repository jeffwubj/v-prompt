package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/jeffwubj/v-prompt/v"
)

func main() {
	fmt.Println("v-prompt")
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	defer fmt.Println("Bye!")
	p := prompt.New(
		v.Executor,
		v.Completer,
		prompt.OptionTitle("kube-prompt: interactive VMware Fusion client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}
