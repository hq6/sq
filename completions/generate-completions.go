package main

import (
	"os"

	"github.com/square/exoskeleton"
)

func main() {
	if f, err := os.Create("completions/bash-completion.sh"); err != nil {
		panic(err)
	} else {
		exoskeleton.GenerateCompletionScript("sq", "bash", f)
		f.Close()
	}

	if f, err := os.Create("completions/zsh-completion.sh"); err != nil {
		panic(err)
	} else {
		exoskeleton.GenerateCompletionScript("sq", "zsh", f)
		f.Close()
	}
}
