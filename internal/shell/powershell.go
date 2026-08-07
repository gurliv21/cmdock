package shell

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func InstallPowerShell() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	profile := powerShellProfilePath(home)

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

func powerShellProfilePath(home string) string {
	if runtime.GOOS == "windows" {
		return filepath.Join(home, "Documents", "PowerShell", "Microsoft.PowerShell_profile.ps1")
	}
	// macOS / Linux (PowerShell Core)
	return filepath.Join(home, ".config", "powershell", "Microsoft.PowerShell_profile.ps1")
}

func powerShellScript() string {
	return "\n" +
		"# >>> cmdock start >>>\n" +
		"\n" +
		"$global:__cmdockLastId = 0\n" +
		"\n" +
		"function global:prompt {\n" +
		"    $history = Get-History -Count 1\n" +
		"\n" +
		"    if ($history -and $history.Id -ne $global:__cmdockLastId) {\n" +
		"        $global:__cmdockLastId = $history.Id\n" +
		"\n" +
		"        if ($history.CommandLine -notmatch \"^cmdock\") {\n" +
		"            $exitCode = if ($?) { 0 } else { 1 }\n" +
		"            $startTs = [DateTimeOffset]$history.StartExecutionTime | ForEach-Object { $_.ToUnixTimeSeconds() }\n" +
		"            $endTs = [DateTimeOffset]$history.EndExecutionTime | ForEach-Object { $_.ToUnixTimeSeconds() }\n" +
		"\n" +
		"            cmdock record `\n" +
		"                --cmd \"$($history.CommandLine)\" `\n" +
		"                --dir \"$PWD\" `\n" +
		"                --start \"$startTs\" `\n" +
		"                --end \"$endTs\" `\n" +
		"                --exit \"$exitCode\" *> $null\n" +
		"        }\n" +
		"    }\n" +
		"\n" +
		"    \"PS $($executionContext.SessionState.Path.CurrentLocation)$('>' * ($nestedPromptLevel + 1)) \"\n" +
		"}\n" +
		"\n" +
		"# <<< cmdock end <<<\n"
}