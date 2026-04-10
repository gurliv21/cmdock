package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"cmdock/internal/shell"


)

var initCmd = &cobra.Command{
	Use:"init",
	Short:"Initialize cmdock shell integration",
	Run:func(cmd *cobra.Command, args []string){
		err:= initShell()

		if err != nil{
			fmt.Println("error occured",err)
			return
		}

		fmt.Println("Init completed")

	},
}

func initShell() error{
	home, err := os.UserHomeDir()  // /home/user
	if err!=nil{
		return err
	}

	ZshPath := home + "/.zshrc"

	exists, _ := fileContains(ZshPath,"cmdock start")

	if exists{
		return nil
	}

	file, err :=os.OpenFile(ZshPath, os.O_APPEND| os.O_WRONLY ,0644)

	if err !=nil{
		return err
	}

	defer file.Close()


	script := shell.ZshScript()

	_, err =file.WriteString(script) // write text

	if err!=nil{
		return err
	}

	return nil


}

func fileContains(path, text string) (bool , error){
	data,err := os.ReadFile(path)
	if err!=nil{
		return false, err
	}

	return strings.Contains(string(data), text), nil
}


func init(){
	rootCmd.AddCommand(initCmd)
}