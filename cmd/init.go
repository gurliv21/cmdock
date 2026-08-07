package cmd

import (
	"cmdock/internal/shell"
	"fmt"
	"cmdock/internal/ui"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize cmdock shell integration",
	Run: func(cmd *cobra.Command, args []string){
		err :=initShell()
		if err != nil{
			ui.Error("Error occurred While Init cmdock")
			return
		}
		ui.Success("Init completed.... ")
	},
}

func initShell() error{
	info,err :=shell.Detect()
	if err!=nil{
		return err
	}
	ui.Success(fmt.Sprintf("Running in %s", info.Type))

	switch info.Type{
	case shell.Zsh:
		return shell.InstallZsh()
	case shell.Bash:
		return shell.InstallBash()
	case shell.PowerShell:
		return shell.InstallPowerShell()
	case shell.Fish:
		return shell.InstallFish()	
	default:
		 ui.Error("unsupported shell: %s ")	
		 return nil	
	}

	return nil


}

func init(){ 
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(recordCmd)
	rootCmd.AddCommand(uninstallCmd)
}


// this initalize all the shells 