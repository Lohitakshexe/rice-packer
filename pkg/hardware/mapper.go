package hardware

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptMappings asks the user for hardware mappings based on the manifest.
func PromptMappings(hardware map[string]string) map[string]string {
	if len(hardware) == 0 {
		return nil
	}

	fmt.Println("\n--- Hardware Mapping ---")
	fmt.Println("This rice contains hardware-specific configurations.")
	fmt.Println("Press Enter to keep the original value, or type a new one.")

	resolved := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	for key, origVal := range hardware {
		fmt.Printf("Original %s: '%s'\n", key, origVal)
		fmt.Printf("Enter your %s (or press Enter to keep '%s'): ", key, origVal)
		
		input, err := reader.ReadString('\n')
		if err != nil {
			// EOF or error, keep original
			resolved[origVal] = origVal
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			resolved[origVal] = origVal
		} else {
			resolved[origVal] = input
		}
	}

	fmt.Println("------------------------\n")
	return resolved
}

// ApplyMappings replaces hardware strings in the content
func ApplyMappings(content []byte, mappings map[string]string) []byte {
	if len(mappings) == 0 {
		return content
	}

	strContent := string(content)
	for orig, newStr := range mappings {
		strContent = strings.ReplaceAll(strContent, orig, newStr)
	}

	return []byte(strContent)
}
