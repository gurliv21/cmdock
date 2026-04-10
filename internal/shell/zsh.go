package shell

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

  command cmdock log \
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