# Rice Packer & Base Hyprland Rice

This repository contains two major components for Linux ricing enthusiasts:
1. **Rice Packer Tool**: A highly concurrent, cross-platform CLI tool built in Go for packaging, mapping, and distributing entire dotfile setups (rices).
2. **Base Hyprland Rice**: A complete, modular, and ready-to-use Hyprland configuration template using Kitty, Waybar, Rofi, and modern Hyprland v0.55+ syntax.

---

## 1. Universal Rice Packer Tool
The Rice Packer is a utility designed to solve the "it works on my machine" problem for dotfiles. It maps hardware requirements (monitor resolutions, GPU drivers), resolves missing dependencies for the target system, and packages everything into a distributable `.rice` archive.

### Features
- **Concurrent Dependency Resolution**: Fast parallel checks for system dependencies.
- **Hardware Mapping**: Generates mapping profiles for monitors and GPUs so rices translate perfectly across different hardware.
- **Unified Archiving**: Packs configs into a standardized format with a `manifest.toml`.

### Build Instructions
You need Go installed on your system. To build the Rice Packer tool:
```bash
git clone https://github.com/Lohitakshexe/rice-packer.git
cd rice-packer
go build -o rice-packer main.go
```

### Usage
**Packing a Rice:**
Create a `manifest.toml` defining your dependencies and config paths, then run:
```bash
./rice-packer pack --config /path/to/manifest.toml --output MySetup.rice
```

**Installing a Rice:**
Take any `.rice` package and apply it to your system (this will check for missing dependencies and prompt to install them):
```bash
./rice-packer install --file MySetup.rice
```

---

## 2. Base Hyprland Rice
Inside the `base-hyprland-rice` directory, you'll find a sleek, modern, Catppuccin/Dark Neon inspired Hyprland configuration. It's meant to be a solid foundation that is completely error-free on the newest versions of Hyprland (v0.55+).

### Components Included:
- **Hyprland**: Modular config (animations, monitors, windowrules separated).
- **Waybar**: Floating, rounded status bar.
- **Kitty**: Themed terminal emulator.
- **Rofi**: Centered application launcher.
- **Hyprlock & Hypridle**: Native lock screen and idle daemon.
- **Scripts**: Utility scripts for wallpapers (using `swww`) and locking.

### How to Test Safely (Nested Mode)
You can test this rice on your current system without overriding your personal `~/.config/` files by using a nested Wayland session and the `XDG_CONFIG_HOME` trick:

```bash
cd rice-packer/base-hyprland-rice
XDG_CONFIG_HOME=$PWD Hyprland -c $PWD/hypr/hyprland.conf
```
*When the window opens, press `SUPER + Return` to launch the themed Kitty, and `SUPER + Space` for Rofi!*

### How to Apply Permanently
When you are ready to make this your permanent setup:
1. Back up your current `~/.config` files.
2. Copy the folders to your system:
```bash
cp -r base-hyprland-rice/hypr ~/.config/
cp -r base-hyprland-rice/kitty ~/.config/
cp -r base-hyprland-rice/waybar ~/.config/
cp -r base-hyprland-rice/rofi ~/.config/
```
3. Restart Hyprland.
