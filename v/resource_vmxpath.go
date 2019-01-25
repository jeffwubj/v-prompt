package v

import (
	"os"
	"path"
	"path/filepath"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jeffwubj/escapeshellchar"
	homedir "github.com/mitchellh/go-homedir"
)

func GetVMXPathesSuggestions() []prompt.Suggest {
	vmxpathes := checkExt(".vmx")
	s := make([]prompt.Suggest, len(vmxpathes))
	for i := range vmxpathes {
		s[i] = prompt.Suggest{
			Text: vmxpathes[i],
			// Description: vmxpathes[i].Status.StartTime.String(),
		}
	}
	return s
}

func getHomeFolder() string {
	home, _ := homedir.Dir()
	return home
}

func checkExt(ext string) []string {
	home := path.Join(getHomeFolder(), "Virtual Machines.localized")
	var files []string
	filepath.Walk(home, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if filepath.Ext(path) == ext {
				path = escapeshellchar.EscapeShellString(path)
				files = append(files, path)
			}
		}
		return nil
	})
	return files
}
