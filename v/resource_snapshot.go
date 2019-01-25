package v

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jeffwubj/escapeshellchar"
)

func GetSnapshotSuggestions(vmxpath string) []prompt.Suggest {
	result := ExecuteAndGetResult("listSnapshots " + vmxpath)
	lines := strings.Split(result, "\n")
	var res []prompt.Suggest
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		line = escapeshellchar.EscapeShellString(line)
		res = append(res, prompt.Suggest{
			Text:        line,
			Description: "Snapshot name",
		})
	}
	return res
}
