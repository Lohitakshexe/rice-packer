package dependency

import (
	"fmt"
	"os/exec"
)

func CheckAndInstall(deps []string) []string {
	var missing []string
	var warnings []string

	for _, dep := range deps {
		_, err := exec.LookPath(dep)
		if err != nil {
			missing = append(missing, dep)
		}
	}

	if len(missing) == 0 {
		return warnings
	}

	fmt.Printf("Missing dependencies detected: %v\n", missing)

	pkgManager, installCmd := detectPackageManager()
	if pkgManager == "" {
		warnings = append(warnings, fmt.Sprintf("Could not detect package manager. Please install manually: %v", missing))
		return warnings
	}

	fmt.Printf("Detected package manager: %s\n", pkgManager)
	fmt.Printf("Attempting to auto-install: %v\n", missing)

	args := append(installCmd, missing...)
	cmd := exec.Command("sudo", args...)
	// In a real TTY we would attach stdin/out/err to allow user to type 'Y' and see progress
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	
	// For this test environment, we'll try to run non-interactively if possible
	// Or we can just run it. We'll capture output for now.
	output, err := cmd.CombinedOutput()
	if err != nil {
		warnings = append(warnings, fmt.Sprintf("Failed to auto-install dependencies: %v\nOutput: %s\nPlease install manually: %v", err, string(output), missing))
	} else {
		fmt.Printf("Successfully installed dependencies!\n")
	}

	return warnings
}

func detectPackageManager() (string, []string) {
	if _, err := exec.LookPath("pacman"); err == nil {
		return "pacman", []string{"pacman", "-S", "--noconfirm"}
	}
	if _, err := exec.LookPath("apt-get"); err == nil {
		return "apt-get", []string{"apt-get", "install", "-y"}
	}
	if _, err := exec.LookPath("dnf"); err == nil {
		return "dnf", []string{"dnf", "install", "-y"}
	}
	if _, err := exec.LookPath("zypper"); err == nil {
		return "zypper", []string{"zypper", "install", "-y"}
	}
	return "", nil
}
