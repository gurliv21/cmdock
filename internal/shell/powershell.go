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
	return "\n" +
		"# >>> cmdock start >>>\n" +
		"\n" +
		"Register-EngineEvent PowerShell.OnIdle -Action {\n" +
		"    $history = Get-History -Count 1\n" +
		"\n" +
		"    if ($history -and $history.CommandLine -notmatch \"^cmdock\") {\n" +
		"\n" +
		"        cmdock record `\n" +
		"--cmd \"$($history.CommandLine)\" `\n" +
		"            --dir \"$PWD\" `\n" +
		"--start \"$([DateTimeOffset]::Now.ToUnixTimeSeconds())\" `\n" +
		"            --end \"$([DateTimeOffset]::Now.ToUnixTimeSeconds())\" `\n" +
		"--exit \"$LASTEXITCODE\" *> $null\n" +
		"    }\n" +
		"}\n" +
		"\n" +
		"# <<< cmdock end <<<\n"
}