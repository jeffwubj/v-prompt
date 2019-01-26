package v

import (
	"fmt"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jeffwubj/escapeshellchar"
)

func GetNetworkAdapters(vmxpath string) []prompt.Suggest {
	result := ExecuteAndGetResult("listNetworkAdapters " + vmxpath)
	lines := strings.Split(result, "\n")
	var res []prompt.Suggest
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		args := strings.Fields(line)
		index := args[0]
		vmnettype := args[1]
		vmnetname := args[2]
		line = escapeshellchar.EscapeShellString(line)
		res = append(res, prompt.Suggest{
			Text:        index,
			Description: fmt.Sprintf("Type: %s, VMNET: %s", vmnettype, vmnetname),
		})
	}
	return res
}
