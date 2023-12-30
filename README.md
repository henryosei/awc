# Go-based Word Counter CLI

## Overview

The **Go-based Word Counter CLI** is a command-line tool written in the Go programming language that allows users to analyze text files and count the number of words.

## Features

- **Word Counting:** Quickly and accurately counts the number of words in a given text file.
- **Command-Line Interface:** Easy-to-use command-line interface for seamless interaction.
- **Fast and Efficient:** Written in Go, the CLI is designed for optimal performance.

## Usage

To use the Word Counter CLI, follow these steps:

1. **Download the Binary:**
    - Releases

2. **Run the CLI:**
   ```bash
   ./awc -f path/to/your/file.txt



## Usage

The Word Counter CLI supports the following flags:

- `-c`: Display the character count.
- `-l`: Display the line count.
- `-w`: Display the word count.
- `-m`: Display the total character count, excluding spaces.

### Examples

1. Count characters and words in a file:

```bash
./awc -c -w input.txt
```

2. Count lines and characters in a string:

```bash
./awc -l -c -m "Hello, this is a sample text."
```

3. Count words in a file and exclude spaces from the character count:

```bash
./awc -w -m file.txt
```

## License

This Word Counter CLI is distributed under the [MIT License](LICENSE).
