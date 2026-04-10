package cmd

import (
	"cmdock/internal/storage"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	flagCmd   string
	flagDir   string
	flagStart string
	flagEnd   string
	flagExit  string
)

var logCmd = &cobra.Command{
	Use:    "log",
	Short:  "Log a command (called by shell hook)",
	Hidden: true, 
	Run: func(cmd *cobra.Command, args []string) {
		db, err := storage.InitDB()
		if err != nil {
			return
		}
		defer db.Close()

		start, _ := strconv.ParseInt(flagStart, 10, 64)
		end, _ := strconv.ParseInt(flagEnd, 10, 64)
		exit, _ := strconv.Atoi(flagExit)

		storage.InsertCommand(db, storage.Command{
			Command:   flagCmd,
			Directory: flagDir,
			ExitCode:  exit,
			StartTime: start,
			EndTime:   end,
		})
	},
}

func init() {
	logCmd.Flags().StringVar(&flagCmd, "cmd", "", "Command executed")
	logCmd.Flags().StringVar(&flagDir, "dir", "", "Directory")
	logCmd.Flags().StringVar(&flagStart, "start", "", "Start time unix")
	logCmd.Flags().StringVar(&flagEnd, "end", "", "End time unix")
	logCmd.Flags().StringVar(&flagExit, "exit", "", "Exit code")
	rootCmd.AddCommand(logCmd)
}

