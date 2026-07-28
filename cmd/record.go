package cmd

import (
	"cmdock/internal/store"
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

var recordCmd = &cobra.Command{
	Use:    "record",
	Short:  "Record a command (called by shell hook)",
	Hidden: true, 
	Run: func(cmd *cobra.Command, args []string) {
		db, err := store.InitDB()
		if err != nil {
			return
		}
		defer db.Close()

		start, _ := strconv.ParseInt(flagStart, 10, 64)
		end, _ := strconv.ParseInt(flagEnd, 10, 64)
		exit, _ := strconv.Atoi(flagExit)

		store.InsertCommand(db, store.Command{
			Command:   flagCmd,
			Directory: flagDir,
			ExitCode:  exit,
			StartTime: start,
			EndTime:   end,
		})
	},
}

func init() {
	recordCmd.Flags().StringVar(&flagCmd, "cmd", "", "Command executed")
	recordCmd.Flags().StringVar(&flagDir, "dir", "", "Directory")
	recordCmd.Flags().StringVar(&flagStart, "start", "", "Start time")
	recordCmd.Flags().StringVar(&flagEnd, "end", "", "End time")
	recordCmd.Flags().StringVar(&flagExit, "exit", "", "Exit code")
	rootCmd.AddCommand(recordCmd)
}

//this is hidden where the commands are inserted in sqllite from shell