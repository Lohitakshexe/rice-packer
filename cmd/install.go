package cmd

import (
	"fmt"
	"os"

	"github.com/lohitaksh/rice-packer/pkg/installer"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [archive_file]",
	Short: "Install a .rice archive",
	Long: `Safely installs a .rice archive by backing up existing configs 
and applying the new ones with your user paths reinjected.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		archiveFile := args[0]
		fmt.Printf("Starting install process using archive: %s\n", archiveFile)

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			os.Exit(1)
		}

		err = installer.Install(archiveFile, homeDir)
		if err != nil {
			fmt.Printf("Error during installation: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Println("Install command executed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
