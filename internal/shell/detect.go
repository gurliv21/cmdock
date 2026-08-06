package shell

import(
	"os"
	"github.com/shirou/gopsutil/v3/process"
	"path/filepath"
	"strings"
)

type Type string

const (
	Unknown Type = "unknown"

	Bash       Type = "bash"
	Zsh        Type = "zsh"
	Fish       Type = "fish"
	PowerShell Type = "powershell"
	Cmd        Type = "cmd"
	Nushell    Type = "nushell"
)

type Info struct{
	Type Type
	Name string
	Path string
}

func Detect()(*Info, error){
	ppid :=os.Getppid()
	proc,err := process.NewProcess(int32(ppid))

	if err !=nil{
		return nil ,err
	}
	name,err := proc.Name()

	if err !=nil{
		return nil ,err
	}

	exe,_ := proc.Exe()

	return &Info{
		Type:classify(name),
		Name:name,
		Path:exe,
	},nil


}

func classify(name string) Type {
	name = strings.ToLower(filepath.Base(name))

	switch {
	case strings.Contains(name, "bash"):
		return Bash

	case strings.Contains(name, "zsh"):
		return Zsh

	case strings.Contains(name, "fish"):
		return Fish

	case strings.Contains(name, "pwsh"),
		strings.Contains(name, "powershell"):
		return PowerShell

	case strings.Contains(name, "cmd"):
		return Cmd

	case name == "nu",
		strings.Contains(name, "nushell"):
		return Nushell

	default:
		return Unknown
	}
}