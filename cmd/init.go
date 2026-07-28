package cmd

import (
	"cmdock/internal/shell"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize cmdock shell integration",
	Run: func(cmd *cobra.Command, args []string){
		err :=initShell()
		if err != nil{
			fmt.Println("Error occurred",err)
			return
		}
		fmt.Println("Init completed. Run: source ~/.zshrc")
	},
}

func initShell() error{
		home, err := os.UserHomeDir() // /home/yourName
		if err !=nil{
			return err
		}

		zshPath := home+ "/.zshrc"

		exists,_:=fileContains(zshPath,"cmdock start")
		if exists {
			fmt.Println("cmdock hook already installed, skipping")
			return nil
	    }
		file,err := os.OpenFile(zshPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err !=nil{
			return err
		}

		defer file.Close()

		script := shell.ZshScript()
		_,err = file.WriteString(script)

		if err !=nil{
			return err
		}

		return nil


}

func fileContains(path, text string)(bool,error){
	data, err := os.ReadFile(path)
	if err !=nil{
		return false,err
	}
	return strings.Contains(string(data),text), nil
}

func init(){
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(recordCmd)
	rootCmd.AddCommand(uninstallCmd)
}


// this initalize all the shells 