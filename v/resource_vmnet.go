package v

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jeffwubj/escapeshellchar"
)

func GetHostNetworks() []prompt.Suggest {
	result := ExecuteAndGetResult("listHostNetworks")
	lines := strings.Split(result, "\n")
	var res []prompt.Suggest
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		args := strings.Fields(line)
		vmnetname := args[1]
		vmnettype := args[2]
		line = escapeshellchar.EscapeShellString(line)
		res = append(res, prompt.Suggest{
			Text:        vmnetname,
			Description: vmnettype,
		})
	}
	return res
}
