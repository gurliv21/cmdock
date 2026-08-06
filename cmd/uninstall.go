package cmd

import(
    "os"
	"path/filepath"
	"cmdock/internal/ui"
	"regexp"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:"uninstall",
	Short:"Removes and Clean up cmdock from you device",
	Run:func(cmd *cobra.Command, args []string){
		err := uninstall()
		if err !=nil{
			ui.Error("Error occurred while unistalling Cmdock")
			return
		}
		ui.Success("Shell integration removed")
		ui.Success("Database deleted")
		ui.Info("Restart your terminal or run: source ~/.zshrc")

	},
}

func uninstall() error{
	home,err:=os.UserHomeDir()
	if err !=nil{
		return err
	}

	//temp shell removal
	zshPath := filepath.Join(home, ".zshrc")

	data, err := os.ReadFile(zshPath)
	if err == nil {
		re := regexp.MustCompile(`(?s)# >>> cmdock start >>>.*?# <<< cmdock end <<<\n?`)
		newContent := re.ReplaceAllString(string(data), "")

		if err := os.WriteFile(zshPath, []byte(newContent), 0644); err != nil {
			return err
		}
	}

	//db removal

	dbPath :=filepath.Join(home,".cmdock.db")
	if err :=os.Remove(dbPath); err !=nil && !os.IsNotExist(err){
		return err
	}

	return nil


}