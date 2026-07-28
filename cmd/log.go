package cmd

import(
	"github.com/spf13/cobra"
	"os"
	"cmdock/internal/tui"
)

var logCmd = &cobra.Command{
	Use:"log",
	Short:"Show command history",
	Run:func(cmd *cobra.Command, args []string){
		if err := tui.Run();
		err!=nil{
			os.Exit(1)
		}
	
	
	},
}


// this shows the tui