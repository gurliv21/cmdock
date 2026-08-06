package shell

import(
  "cmdock/internal/ui"
  "os"
  "strings"
  "path/filepath"
)
func InstallZsh() error{

		home, err := os.UserHomeDir() // /home/yourName
		if err !=nil{
			return err
		}

		// zshPath := home+ "/.zshrc"

    zshPath := filepath.Join(home, ".zshrc")

    if _, err := os.Stat(zshPath); os.IsNotExist(err) {
		file, err := os.Create(zshPath)
		if err != nil {
			return err
		}
		file.Close()
	}


		exists,_:=fileContains(zshPath,"cmdock start")

		if exists {
			ui.Info("cmdock hook already installed, skipping")
			return nil
	  }

		file,err := os.OpenFile(zshPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

		if err !=nil{
			return err
		}

		defer file.Close()

		script := ZshScript()
		_,err = file.WriteString(script)

		if err !=nil{
			return err
		}

    ui.Info("Restart your terminal or run: source ~/.zshrc")

		return nil  
}

func ZshScript() string {
    return `
# >>> cmdock start >>>
if ! typeset -f __cmdock_preexec >/dev/null 2>&1; then
__cmdock_preexec() {
  export CMD_START_TIME=$(date +%s)
  export CMD_DIR=$(pwd)
  export CMD_COMMAND="$1"
}

__cmdock_precmd() {
  local exit_code=$?
  local end_time=$(date +%s)

  if [[ "$CMD_COMMAND" == cmdock* ]]; then
    return
  fi

  command cmdock record \
    --cmd "$CMD_COMMAND" \
    --dir "$CMD_DIR" \
    --start "$CMD_START_TIME" \
    --end "$end_time" \
    --exit "$exit_code" >/dev/null 2>&1
}

autoload -Uz add-zsh-hook
add-zsh-hook preexec __cmdock_preexec
add-zsh-hook precmd __cmdock_precmd
fi
# <<< cmdock end <<<
`
}

func fileContains(path, text string)(bool,error){
	data, err := os.ReadFile(path)
	if err !=nil{
		return false,err
	}
	return strings.Contains(string(data),text), nil
}
