package shell

import (
	"os"
	"path/filepath"
	"strings"
)

func InstallFish() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	config := filepath.Join(home, ".config", "fish", "config.fish")

	if err := os.MkdirAll(filepath.Dir(config), 0755); err != nil {
		return err
	}

	if _, err := os.Stat(config); os.IsNotExist(err) {
		file, err := os.Create(config)
		if err != nil {
			return err
		}
		file.Close()
	}

	data, err := os.ReadFile(config)
	if err != nil {
		return err
	}

	if strings.Contains(string(data), "# >>> cmdock start >>>") {
		return nil
	}

	file, err := os.OpenFile(config, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("\n" + fishScript() + "\n")
	return err
}

func fishScript() string {
	return `
# >>> cmdock start >>>

function __cmdock_preexec --on-event fish_preexec
    set -gx CMD_START_TIME (date +%s)
    set -gx CMD_DIR (pwd)
    set -gx CMD_COMMAND $argv
end

function __cmdock_postexec --on-event fish_postexec
    set exit_code $status
    set end_time (date +%s)

    if string match -qr "^cmdock" "$CMD_COMMAND"
        return
    end

    command cmdock record \
        --cmd "$CMD_COMMAND" \
        --dir "$CMD_DIR" \
        --start "$CMD_START_TIME" \
        --end "$end_time" \
        --exit "$exit_code" >/dev/null 2>&1
end

# <<< cmdock end <<<
`
}