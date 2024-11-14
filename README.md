# string to ascii art

A Go program that converts strings into ASCII art with various styling options and features.

## Description

This ASCII Art Generator transforms text input into graphical ASCII art representations. It supports multiple banner styles, alignment options, color customization, and can even reverse ASCII art back to text.

## Features

- 🎯 Multiple banner styles (standard, shadow, thinkertoy)
- 📏 Text alignment options (left, center, right, justify)
- 🌈 Color customization with RGB support
- 💾 Output to file functionality
- 🔄 Reverse ASCII art to text conversion
- 🔤 Support for letters, numbers, spaces, and special characters

## Getting Started


### Installation

```bash
# Clone the repository
git clone https://github.com/yassinalmach/string-to-ascii-art
cd string-to-ascii-art

# Build the project
go build
```

## Usage

### Basic Usage
```bash
go run . "Hello World"
```

### With Banner Style
```bash
go run . "Hello World" shadow
```

### Output to File
```bash
go run . --output=example.txt "Hello World" thinkertoy
```

### Text Alignment
```bash
go run . --align=center "Hello World" standard
```

### Color Customization
```bash
# Color specific letters
go run . --color=red He "Hello World"

# Using RGB values
go run . --color=rgb(255,0,0) world "Hello World"

# Color entire text
go run . --color=blue "Hello World"
```

### Reverse ASCII Art
```bash
go run . --reverse=example.txt
```

## Available Colors

- white, black, red, green, blue, yellow, orange, pink, brown, purple
- RGB custom colors (0-255)

## Banner Styles

- `standard`
- `shadow`
- `thinkertoy`

## Options

| Flag | Description |
|------|-------------|
| `--output=<fileName.txt>` | Save output to a file |
| `--align=<type>` | Set text alignment (left/center/right/justify) |
| `--color=<color>` | Apply color to text |
| `--reverse=<fileName>` | Convert ASCII art back to text |

## Error Handling

The program includes proper error handling and usage messages for:
- Invalid flags or options
- File operations
- Color specifications
- Banner selection
- Terminal size limitations
