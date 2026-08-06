package shell

import (
	"os"
	"path/filepath"
	"strings"
)

func InstallBash() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	bashrc := filepath.Join(home, ".bashrc")

	// Create ~/.bashrc if it doesn't exist.
	if _, err := os.Stat(bashrc); os.IsNotExist(err) {
		file, err := os.Create(bashrc)
		if err != nil {
			return err
		}
		file.Close()
	}

	data, err := os.ReadFile(bashrc)
	if err != nil {
		return err
	}
   
	if strings.Contains(string(data), "# >>> cmdock start >>>") {
		return nil
	}

	file, err := os.OpenFile(bashrc, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("\n" + bashScript() + "\n")
	return err
}

func bashScript() string {
	return `
# >>> cmdock start >>>

__cmdock_preexec() {
    export CMD_START_TIME=$(date +%s)
    export CMD_DIR=$(pwd)
    export CMD_COMMAND="$BASH_COMMAND"
}

__cmdock_precmd() {
    local exit_code=$?
    local end_time=$(date +%s)

    [[ "$CMD_COMMAND" == cmdock* ]] && return

    command cmdock record \
        --cmd "$CMD_COMMAND" \
        --dir "$CMD_DIR" \
        --start "$CMD_START_TIME" \
        --end "$end_time" \
        --exit "$exit_code" >/dev/null 2>&1
}

trap '__cmdock_preexec' DEBUG
PROMPT_COMMAND="__cmdock_precmd${PROMPT_COMMAND:+;$PROMPT_COMMAND}"

# <<< cmdock end <<<
`
}