# Oraxia - Deterministic Password Generator

Oraxia is a command-line password generator written in Go that creates secure, deterministic passwords based on a single input word. Inspired by Warframe's mysterious and powerful themes, Oraxia ensures that the same input word always generates the same password, making it perfect for creating memorable yet strong passwords.

## Features

- **Deterministic Generation**: Same words always produce the same password
- **Secure Hashing**: Uses HMAC-SHA256 with a secret token for cryptographic strength
- **Command-Line Interface**: Simple to use from the terminal
- **Base64 Output**: Generates printable passwords with alphanumeric and special characters
- **Secret Token Protection**: Requires a secret file for added security

## How It Works

Oraxia takes your input word and computes an HMAC-SHA256 hash using your secret token. It uses the first 128 bits (16 bytes) of the hash, base64-encoded to create a strong, 22-character password.

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/oraxia.git
   cd oraxia
   ```

2. Build the binary:

   ```bash
   go build -o oraxia main.go
   ```

3. (Optional) Install globally for command-line access:

   ```bash
   # Option 1: Move to system bin (requires sudo)
   sudo mv oraxia /usr/local/bin/

   # Option 2: Add to user bin (create ~/bin if it doesn't exist)
   mkdir -p ~/bin
   mv oraxia ~/bin/
   # Add ~/bin to PATH in your ~/.bashrc or ~/.zshrc: export PATH="$HOME/bin:$PATH"
   ```

4. Set up your secret token (see Security section below)

## Usage

After building and optionally installing the binary, use it from the command line:

```bash
# If not installed globally, run from the project directory
./oraxia <word>

# If installed globally (see Installation step 3), use anywhere
oraxia <word>
```

### Examples

```bash
$ oraxia github
wzfJ26DDJZPlbMi6gzbJQA==

$ oraxia github
wzfJ26DDJZPlbMi6gzbJQA==
```

Notice that the same word always gives the same output.

## Security

Oraxia generates 22-character passwords with 128-bit entropy using HMAC-SHA256. The security relies entirely on keeping your `secret.txt` file private. Never commit it to version control.

### Generating a Secure Secret

Use one of these methods to generate a strong secret:

```bash
# Using OpenSSL (recommended)
openssl rand -base64 32 > secret.txt

# Or using /dev/urandom
head -c 32 /dev/urandom | base64 > secret.txt

# Or using Go
go run -c 'package main; import ("crypto/rand"; "encoding/base64"; "fmt"); func main() { b := make([]byte, 32); rand.Read(b); fmt.Println(base64.StdEncoding.EncodeToString(b)) }' > secret.txt
```

### Security Considerations

- The secret token should be at least 32 bytes (256 bits) long
- Keep `secret.txt` secure and private
- Use different secrets for different purposes if needed
- The generated passwords are cryptographically strong but remember that deterministic passwords have different security properties than random ones

## Requirements

- Go 1.16 or later


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This tool is for educational and personal use. While it uses strong cryptography, the deterministic nature means it's not suitable for all security scenarios. Always use unique, strong passwords and consider using a password manager for critical accounts.
