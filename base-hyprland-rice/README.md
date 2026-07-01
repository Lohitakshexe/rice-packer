# Base Hyprland Rice

This is a modular, dark neon/Catppuccin-inspired base Hyprland rice. It acts as a solid, easily customizable foundation for further ricing.

## Features
- **Hyprland (v0.55+ Compatible)**: Fully updated configuration syntax. Modularized into separate files for keybinds, monitors, autostart, window rules, and animations.
- **Waybar**: A beautiful, floating, rounded status bar with animated modules for battery, network, workspace switching, and system tray.
- **Kitty Terminal**: Clean configuration with a sleek color scheme matching the overall aesthetic.
- **Rofi**: Centered, themed application launcher.
- **Hyprlock & Hypridle**: Native, beautifully blurred lock screen with automatic idle management.
- **Wallpapers**: Built-in script to set and animate wallpapers using `swww`.

## Directory Structure
```
base-hyprland-rice/
├── hypr/
│   ├── hyprland.conf       # Main config, sources the others
│   ├── animations.conf     # Smooth bezier curve animations
│   ├── autostart.conf      # Exec-once commands (waybar, swww, etc)
│   ├── keybinds.conf       # Standard SUPER key bindings
│   ├── monitors.conf       # Display setup
│   ├── windowrules.conf    # Specific app rules (opacity, floats)
│   ├── hyprlock.conf       # Lockscreen UI
│   ├── hypridle.conf       # Idle daemon settings
│   └── scripts/
│       ├── lock.sh         # Manually triggers hyprlock
│       └── wallpaper.sh    # Usage: ./wallpaper.sh /path/to/img
├── kitty/
│   ├── kitty.conf          # Main terminal settings
│   └── theme.conf          # Color palette
├── rofi/
│   ├── config.rasi         # Behavior and modi
│   └── theme.rasi          # UI colors and layout
└── waybar/
    ├── config.jsonc        # Modules and structure
    └── style.css           # CSS styling for the bar
```

## How to Test Without Applying
You can safely test this entire rice in a nested, isolated window without altering your personal system. Because we leverage the `XDG_CONFIG_HOME` environment variable, all the included programs (Waybar, Kitty, Rofi) will automatically load the configurations from this directory instead of your `~/.config/`.

Run this command from your terminal:
```bash
XDG_CONFIG_HOME=/absolute/path/to/base-hyprland-rice Hyprland -c /absolute/path/to/base-hyprland-rice/hypr/hyprland.conf
```
*Note: Replace `/absolute/path/to/base-hyprland-rice` with the actual path where you cloned this folder.*

Once the nested window opens:
- Press `SUPER + Return` to launch the styled Kitty terminal.
- Press `SUPER + Space` to launch the Rofi menu.
- Use Kitty to run `./hypr/scripts/wallpaper.sh /path/to/your/image.jpg` to see the wallpaper transition.

## How to Apply
When you are ready to make this your daily driver:
1. Back up your existing configs in `~/.config/` (hypr, kitty, waybar, rofi).
2. Copy the contents of this folder directly into your `~/.config/` directory:
   ```bash
   cp -r hypr kitty waybar rofi ~/.config/
   ```
3. Log out and log back into a Hyprland session.
