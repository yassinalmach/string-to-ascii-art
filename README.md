# String To Ascii Art  - A Go Learning Journey

A Go program that converts strings into ASCII art with various styling options and features.

## Description

This ASCII Art Generator transforms text input into graphical ASCII art representations. It supports multiple banner styles, alignment options, color customization, and can even reverse ASCII art back to text.


## Project Purpose

This project serves as an immersive learning experience for mastering Go programming fundamentals through practical application. By building an ASCII art generator, we explore core Go concepts while creating something fun and useful.

## Learning Outcomes

### Go Programming Fundamentals
- 📦 **Package Management**: Understanding Go modules and project structure
- 🔄 **Control Structures**: Mastering loops, conditionals, and program flow
- 📝 **String Manipulation**: Deep dive into string handling and rune operations
- 🔧 **Error Handling**: Implementing robust error checking and user feedback
- 📁 **File Operations**: Working with the Go file system API
- 🎯 **Flag Parsing**: Command-line argument handling in Go
- 🎨 **ANSI Color Implementation**: Terminal color manipulation
- 📏 **Text Alignment Algorithms**: String positioning and spacing calculations
- 🖥️ **Terminal Manipulation**: Working with terminal dimensions and output
- 🔄 **Data Transformation**: Converting between text and ASCII art representations

## Technical Skills I Developed

### 1. Go Language Features
- Slices and arrays manipulation
- String and rune handling
- Interface implementation
- Error handling patterns

### 2. File System Operations
- Reading banner files
- Writing output to files
- File existence checking
- Error handling in file operations

### 3. Data Structures
- Multi-dimensional arrays for character representation
- Maps for character lookup
- Custom types and structs

### 4. Algorithms
- Text alignment algorithms
- Color parsing and application
- ASCII art conversion logic
- Terminal width adaptation

### Project Structure

## Testing Your Knowledge

1. Try implementing new banner styles
2. Add support for new color formats
3. Implement custom alignment algorithms
4. Add animation features
5. Optimize the program for large inputs

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Go By Example](https://gobyexample.com/)
- [ASCII Art References](https://en.wikipedia.org/wiki/ASCII_art)
- [ANSI Color Codes](https://en.wikipedia.org/wiki/ANSI_escape_code)


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
