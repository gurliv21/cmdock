package cmd

import (
	"cmdock/internal/storage"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show command history",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := storage.InitDB()
		if err != nil {
			fmt.Println("Error opening database:", err)
			return
		}
		defer db.Close()

		commands, err := storage.ShowCommands(db)
		if err != nil {
			fmt.Println("Error fetching commands:", err)
			return
		}

		if len(commands) == 0 {
			fmt.Println("No commands logged yet.")
			return
		}

		for _, c := range commands {
			duration := c.EndTime - c.StartTime
			t := time.Unix(c.StartTime, 0).Format("Jan 02 15:04:05")
			fmt.Printf("[%s] (%ds) [exit:%d] %s  →  %s\n",
				t, duration, c.ExitCode, c.Directory, c.Command)
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}