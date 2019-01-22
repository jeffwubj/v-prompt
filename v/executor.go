package v

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/davecgh/go-spew/spew"
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

	cmd := exec.Command("/bin/sh", "-c", "vmrun "+s)
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
	args = unescapeArgs(args)
	spew.Dump(args)
	cmd := exec.Command("/bin/sh", "-c", "vmrun "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}

func unescapeArgs(args []string) []string {
	var res []string
	for _, arg := range args {
		unescape, _ := url.QueryUnescape(arg)
		res = append(res, unescape)
	}
	return res
}
