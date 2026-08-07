package shell

import (
	"os"
	"path/filepath"
	"strings"
)

func InstallPowerShell() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	profile := filepath.Join(
		home,
		"Documents",
		"PowerShell",
		"Microsoft.PowerShell_profile.ps1",
	)

	if err := os.MkdirAll(filepath.Dir(profile), 0755); err != nil {
		return err
	}

	if _, err := os.Stat(profile); os.IsNotExist(err) {
		file, err := os.Create(profile)
		if err != nil {
			return err
		}
		file.Close()
	}

	data, err := os.ReadFile(profile)
	if err != nil {
		return err
	}

	if strings.Contains(string(data), "# >>> cmdock start >>>") {
		return nil
	}

	file, err := os.OpenFile(profile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("\n" + powerShellScript() + "\n")
	return err
}

func powerShellScript() string {
	return `
# >>> cmdock start >>>

Register-EngineEvent PowerShell.OnIdle -Action {
    $history = Get-History -Count 1

    if ($history -and $history.CommandLine -notmatch "^cmdock") {

        cmdock record `
            --cmd "$($history.CommandLine)" `
            --dir "$PWD" `
            --start "$([DateTimeOffset]::Now.ToUnixTimeSeconds())" `
            --end "$([DateTimeOffset]::Now.ToUnixTimeSeconds())" `
            --exit "$LASTEXITCODE" *> $null
    }


# <<< cmdock end <<<
`
}