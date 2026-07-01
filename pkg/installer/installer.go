package installer

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lohitaksh/rice-packer/pkg/dependency"
	"github.com/lohitaksh/rice-packer/pkg/hardware"
	"github.com/lohitaksh/rice-packer/pkg/manifest"
	"github.com/pelletier/go-toml/v2"
)

func Install(archivePath string, homeDir string) error {
	// 1. Open the archive
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return fmt.Errorf("failed to open archive: %v", err)
	}
	defer r.Close()

	// Find and parse manifest
	var m *manifest.Manifest
	for _, f := range r.File {
		if f.Name == "manifest.toml" {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return err
			}
			
			m = &manifest.Manifest{}
			err = toml.Unmarshal(data, m)
			if err != nil {
				return fmt.Errorf("failed to parse embedded manifest: %v", err)
			}
			break
		}
	}

	if m == nil {
		return fmt.Errorf("manifest.toml not found in archive")
	}

	// Resolve Dependencies
	warnings := dependency.CheckAndInstall(m.Dependencies)

	// Hardware mapping
	hwMappings := hardware.PromptMappings(m.Hardware)

	// 2. Determine target config dir and backup dir
	configDir := filepath.Join(homeDir, ".config")
	timestamp := time.Now().Format("20060102_150405")
	backupDir := filepath.Join(homeDir, fmt.Sprintf(".config_backup_%s", timestamp))

	fmt.Printf("Backups will be saved to: %s\n", backupDir)

	// 3. Iterate over files in the zip
	for _, f := range r.File {
		// Only process files in the "configs/" directory inside the zip
		if !strings.HasPrefix(f.Name, "configs/") {
			continue
		}

		relPath := strings.TrimPrefix(f.Name, "configs/")
		if relPath == "" {
			continue
		}

		targetPath := filepath.Join(configDir, relPath)
		backupPath := filepath.Join(backupDir, relPath)

		// Create backup directory structure
		if err := os.MkdirAll(filepath.Dir(backupPath), 0755); err != nil {
			return err
		}

		// Backup existing file if it exists
		if _, err := os.Stat(targetPath); err == nil {
			err = copyFile(targetPath, backupPath)
			if err != nil {
				return fmt.Errorf("failed to backup %s: %v", targetPath, err)
			}
		}

		// Now write the new file
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return err
		}

		contentType := http.DetectContentType(content)
		var fileData []byte

		if strings.HasPrefix(contentType, "text/") || strings.HasPrefix(contentType, "application/json") || strings.HasPrefix(contentType, "application/x-sh") {
			// Reinject the target user's home directory
			injectedContent := strings.ReplaceAll(string(content), "$HOME", homeDir)
			
			// Apply hardware mappings
			fileData = hardware.ApplyMappings([]byte(injectedContent), hwMappings)
		} else {
			fileData = content
		}

		err = os.WriteFile(targetPath, fileData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write %s: %v", targetPath, err)
		}
		fmt.Printf("Installed: %s\n", targetPath)
	}

	if len(warnings) > 0 {
		fmt.Println("\n--- WARNINGS ---")
		for _, w := range warnings {
			fmt.Println(w)
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
