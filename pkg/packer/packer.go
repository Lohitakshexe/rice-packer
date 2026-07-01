package packer

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/lohitaksh/rice-packer/pkg/manifest"
	"github.com/pelletier/go-toml/v2"
)

func Pack(m *manifest.Manifest, homeDir string, outputName string) error {
	// Create the zip file
	zipFile, err := os.Create(outputName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// Write the manifest file into the archive
	manifestFile, err := archive.Create("manifest.toml")
	if err != nil {
		return err
	}
	
	manifestData, err := toml.Marshal(m)
	if err != nil {
		return err
	}
	_, err = manifestFile.Write(manifestData)
	if err != nil {
		return err
	}

	configDir := filepath.Join(homeDir, ".config")

	for _, cfg := range m.Configs {
		srcPath := filepath.Join(configDir, cfg)
		
		err := filepath.WalkDir(srcPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			// Read file content
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Detect if the file is binary or text
			contentType := http.DetectContentType(content)
			var fileData []byte

			if strings.HasPrefix(contentType, "text/") || strings.HasPrefix(contentType, "application/json") || strings.HasPrefix(contentType, "application/x-sh") {
				// Sanitize text files: replace absolute home path with $HOME
				sanitizedContent := strings.ReplaceAll(string(content), homeDir, "$HOME")
				fileData = []byte(sanitizedContent)
			} else {
				// Leave binary files alone
				fileData = content
			}

			// Determine path inside the zip file
			relPath, err := filepath.Rel(configDir, path)
			if err != nil {
				return err
			}
			zipPath := filepath.Join("configs", relPath)

			// Create file in zip
			f, err := archive.Create(zipPath)
			if err != nil {
				return err
			}

			_, err = f.Write(fileData)
			return err
		})

		if err != nil {
			fmt.Printf("Warning: failed to pack config %s: %v\n", cfg, err)
		}
	}

	return nil
}
