# Universal Rice Packer

The Universal Rice Packer is a cross-platform CLI tool built in Go designed to solve the age-old "it works on my machine" problem for Linux dotfiles (rices). 

Sharing and installing customized desktop environments is notoriously difficult due to differing hardware (monitors, GPUs) and missing software dependencies. This tool automates the process of packaging, mapping, and distributing entire dotfile setups so they work seamlessly on any target machine.

## Why is this important?
Traditionally, installing someone else's "rice" involves manually copying configuration files, hunting down missing fonts and dependencies, and rewriting monitor layouts to fit your physical screens. 

The **Rice Packer** completely eliminates this friction by:
- Creating a standardized `.rice` archive format.
- Automatically mapping the creator's hardware (like monitor resolutions and GPU drivers) to the installer's hardware.
- Resolving and prompting for missing system dependencies during installation.

## How it Works
1. **Manifest Parsing**: It reads a `manifest.toml` file that defines your config paths, required dependencies, and hardware mappings.
2. **Hardware Translation**: When a rice is installed, the tool reads the target machine's hardware and translates the original config's monitor setup to match the new physical displays.
3. **Dependency Resolution**: It checks the host system for required packages (e.g., `waybar`, `rofi`, `kitty`) and handles the installation asynchronously.

---

## Usage

### 1. Build the Tool
Ensure you have Go installed, then clone and build:
```bash
git clone https://github.com/Lohitakshexe/rice-packer.git
cd rice-packer
go build -o rice-packer main.go
```

### 2. Pack a Rice
To share your current setup, create a `manifest.toml` defining your dependencies and config paths, then run:
```bash
./rice-packer pack --config /path/to/manifest.toml --output MySetup.rice
```
This generates a portable `MySetup.rice` file containing your configurations and hardware metadata.

### 3. Install a Rice
To apply a downloaded `.rice` package to your system:
```bash
./rice-packer install --file MySetup.rice
```
The installer will check for missing dependencies, translate hardware mappings, and safely apply the configurations.


