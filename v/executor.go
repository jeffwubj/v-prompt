package v

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	args := strings.Split(s, " ")
	cmd := exec.Command("vmrun", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func ExecuteAndGetResult(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	out := &bytes.Buffer{}
	args := strings.Split(s, " ")
	cmd := exec.Command("vmrun", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}
