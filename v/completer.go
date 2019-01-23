package v

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	args := splitArgs(d.TextBeforeCursor())

	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	return argumentsCompleter(args)
}

func splitArgs(text string) []string {
	if strings.Contains(text, ".vmx") {
		firstspace := strings.Index(text, " ")
		firstvmx := strings.Index(text, ".vmx") + 4
		firstArg := strings.TrimSpace(text[0:firstspace])
		secondArg := strings.TrimSpace(text[firstspace:firstvmx])
		remaining := strings.Split(text[firstvmx:], " ")
		res := []string{firstArg, secondArg}
		if len(remaining) > 0 {
			remaining = remaining[1:]
			res = append(res, remaining...)
		}
		return res
	} else {
		if shouldHaveVMXInPath(text) {
			firstspace := strings.Index(text, " ")
			cmd := text[:firstspace]
			return []string{cmd, text[firstspace+1:]}
		} else {
			return strings.Split(text, " ")
		}
	}
}

func shouldHaveVMXInPath(text string) bool {
	args := strings.Split(text, " ")
	if len(args) > 1 {
		cmd := strings.TrimSpace(args[0])
		if cmd != "listHostNetworks" && cmd != "listHostNetworks" &&
			cmd != "setPortForwarding" && cmd != "deletePortForwarding" &&
			cmd != "downloadPhotonVM" {
			return true
		}
	}
	return false
}

var commands = []prompt.Suggest{
	{Text: "start", Description: "Start a VM"},
	{Text: "stop", Description: "Stop a VM"},
	{Text: "reset", Description: "Reset a VM"},
	{Text: "suspend", Description: "Suspend a VM"},
	{Text: "pause", Description: "Pause a VM"},
	{Text: "unpause", Description: "Unpause a VM"},
	{Text: "listSnapshots", Description: "List all snapshots in a VM"},
	{Text: "revertToSnapshot", Description: "Set VM state to a snapshot"},
	{Text: "deleteSnapshot", Description: "Remove a snapshot from a VM"},
	{Text: "listNetworkAdapters", Description: "List all network adapters in a VM"},
	{Text: "addNetworkAdapter", Description: "Add a network adapter on a VM"},
	{Text: "setNetworkAdapter", Description: "Update a network adapter on a VM"},
	{Text: "deleteNetworkAdapter", Description: "Remove a network adapter on a VM"},
	{Text: "listHostNetworks", Description: "List all networks in the host"},
	{Text: "listPortForwardings", Description: "List all available port forwardings on a host network"},
	{Text: "setPortForwarding", Description: "Add or update a port forwarding on a host network"},
	{Text: "deletePortForwarding", Description: "Delete a port forwarding on a host network"},
	{Text: "runProgramInGuest", Description: "Run a program in Guest OS"},
	{Text: "fileExistsInGuest", Description: "Check if a file exists in Guest OS"},
	{Text: "directoryExistsInGuest", Description: "Check if a directory exists in Guest OS"},
	{Text: "setSharedFolderState", Description: "Modify a Host-Guest shared folder"},
	{Text: "addSharedFolder", Description: "Add a Host-Guest shared folder"},
	{Text: "removeSharedFolder", Description: "Remove a Host-Guest shared folder"},
	{Text: "enableSharedFolders", Description: "Enable shared folders in Guest"},
	{Text: "disableSharedFolders", Description: "Disable shared folders in Guest"},
	{Text: "listProcessesInGuest", Description: "List running processes in Guest OS"},
	{Text: "killProcessInGuest", Description: "Kill a process in Guest OS"},
	{Text: "runScriptInGuest", Description: "Run a script in Guest OS"},
	{Text: "deleteFileInGuest", Description: "Delete a file in Guest OS"},
	{Text: "createDirectoryInGuest", Description: "Create a directory in Guest OS"},
	{Text: "deleteDirectoryInGuest", Description: "Delete a directory in Guest OS"},
	{Text: "CreateTempfileInGuest", Description: "Create a temporary file in Guest OS"},
	{Text: "listDirectoryInGuest", Description: "List a directory in Guest OS"},
	{Text: "CopyFileFromHostToGuest", Description: "Copy a file from host OS to guest OS"},
	{Text: "CopyFileFromGuestToHost", Description: "Copy a file from guest OS to host OS"},
	{Text: "renameFileInGuest", Description: "Rename a file in Guest OS"},
	{Text: "typeKeystrokesInGuest", Description: "Type Keystrokes in Guest OS"},
	{Text: "connectNamedDevice", Description: "Connect the named device in the Guest OS"},
	{Text: "disconnectNamedDevice", Description: "Disconnect the named device in the Guest OS"},
	{Text: "captureScreen", Description: "Capture the screen of the VM to a local file"},
	{Text: "writeVariable", Description: "Write a variable in the VM state"},
	{Text: "readVariable", Description: "Read a variable in the VM state"},
	{Text: "getGuestIPAddress", Description: "Gets the IP address of the guest"},
	{Text: "list", Description: "List all running VMs"},
	{Text: "upgradevm", Description: "Upgrade VM file format, virtual hw"},
	{Text: "installTools", Description: "Install Tools in Guest"},
	{Text: "checkToolsState", Description: "Check the current Tools state"},
	{Text: "deleteVM", Description: "Delete a VM"},
	{Text: "clone", Description: "Create a copy of the VM"},
	{Text: "downloadPhotonVM", Description: "Download Photon VM"},

	// Custom command.
	{Text: "exit", Description: "Exit this program"},
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterContains(commands, args[0], true)
	}
	return firstArgumentCompleter(args)
}

func firstArgumentCompleter(args []string) []prompt.Suggest {
	first := args[0]
	switch first {
	case "start":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(GetVMXPathesSuggestions(), second, true)
		} else if len(args) == 3 {
			return []prompt.Suggest{
				{Text: "nogui",
					Description: "Start virtual machine at backend without UI"},
				{Text: "gui",
					Description: "Start virtual machine with Fusion UI"}}
		}
	case "stop":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(GetVMXPathesSuggestions(), second, true)
		} else if len(args) == 3 {
			return []prompt.Suggest{
				{Text: "hard",
					Description: "Power off the virtual machine abruptly with no consideration for work in progress"},
				{Text: "soft",
					Description: "Send a shut down signal to the guest operating system"}}
		}
	case "reset":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(GetVMXPathesSuggestions(), second, true)
		} else if len(args) == 3 {
			return []prompt.Suggest{
				{Text: "hard",
					Description: "Power off the virtual machine abruptly with no consideration for work in progress"},
				{Text: "soft",
					Description: "Shuts down and restarts the guest operating system gracefully. VMware Tools may run scripts during the process"}}
		}
	case "suspend":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(GetVMXPathesSuggestions(), second, true)
		} else if len(args) == 3 {
			return []prompt.Suggest{
				{Text: "hard",
					Description: "Suspend the virtual machine abruptly"},
				{Text: "soft",
					Description: "The guest OS can be made aware of the process, allowing it to sleep/hibernate (if supported)"}}
		}
	case "pause":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(GetVMXPathesSuggestions(), second, true)
		}
	case "unpause":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(GetVMXPathesSuggestions(), second, true)
		}
	case "listSnapshots":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(GetVMXPathesSuggestions(), second, true)
		} else if len(args) == 3 {
			return []prompt.Suggest{
				{Text: "showTree",
					Description: "Show snapshot tree"},
			}
		}
	case "revertToSnapshot":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(GetVMXPathesSuggestions(), second, true)
		} else if len(args) == 3 {
			second := args[1]
			return GetSnapshotSuggestions(second)
		}
	}
	return []prompt.Suggest{}
}
