package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lohitaksh/rice-packer/pkg/manifest"
	"github.com/lohitaksh/rice-packer/pkg/packer"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack [manifest_file]",
	Short: "Pack dotfiles into a .rice archive",
	Long: `Reads a manifest file to gather dotfiles, sanitizes paths, 
and packages them into a distributable .rice archive.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manifestFile := args[0]
		fmt.Printf("Starting pack process using manifest: %s\n", manifestFile)
		
		m, err := manifest.Load(manifestFile)
		if err != nil {
			fmt.Printf("Error loading manifest: %v\n", err)
			os.Exit(1)
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			os.Exit(1)
		}

		outputName := fmt.Sprintf("%s.rice", m.Name)
		outputName = filepath.Clean(outputName)
		// Basic sanitization for spaces in filename
		for i := range outputName {
			if outputName[i] == ' ' {
				outputName = outputName[:i] + "_" + outputName[i+1:]
			}
		}

		err = packer.Pack(m, homeDir, outputName)
		if err != nil {
			fmt.Printf("Error during packing: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("Pack command executed successfully! Created archive: %s\n", outputName)
	},
}

func init() {
	rootCmd.AddCommand(packCmd)
}
