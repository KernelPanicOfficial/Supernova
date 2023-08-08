package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Structure
type FlagOptions struct {
	outFile     string
	inputFile   string
	language    string
	encryption  string
	obfuscation string
}

// global variables
var (
	__version__ = "1.0.0"
	__license__ = "MIT"
	__author__  = "@nickvourd"
	__github__  = "https://github.com/nickvourd/Supernova"
)

var __ascii__ = `

███████╗██╗   ██╗██████╗ ███████╗██████╗ ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗ 
██╔════╝██║   ██║██╔══██╗██╔════╝██╔══██╗████╗  ██║██╔═══██╗██║   ██║██╔══██╗
███████╗██║   ██║██████╔╝█████╗  ██████╔╝██╔██╗ ██║██║   ██║██║   ██║███████║
╚════██║██║   ██║██╔═══╝ ██╔══╝  ██╔══██╗██║╚██╗██║██║   ██║╚██╗ ██╔╝██╔══██║
███████║╚██████╔╝██║     ███████╗██║  ██║██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║
╚══════╝ ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝

Supernova v.%s - A real fucking shellcode encryptor.
Supernova is an open source tool licensed under %s.
Written with <3 by %s...
Please visit %s for more...

`

// ConvertShellcode2Template function
func ConvertShellcode2Template(shellcode string, language string) {
	switch language {
	case "nim":
		fmt.Println(language)
	case "rust":
		fmt.Println(language)
	case "c":
		fmt.Println(language)
	case "c#":
		fmt.Println(language)
	default:
		fmt.Println("[!] Unsupported programming language:", language)
	}
}

// ValidateArgument function
func ValidateArgument(argName string, argValue string, validValues []string) string {
	for _, valid := range validValues {
		if strings.ToLower(argValue) == strings.ToLower(valid) {
			valid = strings.ToLower(valid)
			return valid
		}
	}
	fmt.Printf("[!] Invalid value '%s' for argument '%s'. Valid values are: %v\n", argValue, argName, validValues)
	os.Exit(1)
	return ""
}

// ArgumentEmpty function
func ArgumentEmpty(statement string, option int) {
	if statement == "" {
		logger := log.New(os.Stderr, "[!] ", 0)
		switch option {
		case 1:
			logger.Fatal("Please provide a path to a file containing raw 64-bit shellcode.")
		case 2:
			logger.Fatal("Please provide a valid value for the programming language (e.g., C++, C#, Rust, Nim).")
		default:
			logger.Fatal("Invalid option specified for ArgumentEmpty function.")
		}
	}
}

// Options function
func Options() *FlagOptions {
	inputFile := flag.String("i", "", "Path to the raw 64-bit shellcode.")
	encryption := flag.String("enc", "", "Shellcode encryption (i.e., XOR, RC4, AES)")
	obfuscation := flag.String("obs", "", "Shellcode obfuscation")
	language := flag.String("lang", "", "Programming language to translate the shellcode (i.e., Nim, Rust, C, C#)")
	outFile := flag.String("o", "", "Name of output file")
	flag.Parse()

	return &FlagOptions{outFile: *outFile, inputFile: *inputFile, language: *language, encryption: *encryption, obfuscation: *obfuscation}
}

// main function
func main() {
	// Print ascii
	fmt.Printf(__ascii__, __version__, __license__, __author__, __github__)

	// Retrieve command-line options using the Options function
	options := Options()

	// if no arguments print help menu
	if len(os.Args) == 1 {
		fmt.Println("Usage of Suprenova.exe:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Call function named ArgumentEmpty
	ArgumentEmpty(options.inputFile, 1)

	// Call function named ArgumentEmpty
	ArgumentEmpty(options.language, 2)

	// Check for valid values of language argument
	foundLanguage := ValidateArgument("lang", options.language, []string{"Nim", "Rust", "C", "C#"})

	// Checks if either encryption or obfuscation options are provided.
	if options.encryption == "" && options.obfuscation == "" {
		logger := log.New(os.Stderr, "[!] ", 0)
		logger.Fatal("Please provide at least -enc or -obs option with a valid value...")
	}

	// Check for valid values of encryption argument
	if options.encryption != "" {
		ValidateArgument("enc", options.encryption, []string{"XOR", "RC4", "AES"})
	}

	// Check for valid values of obfuscation argument
	if options.obfuscation != "" {
		ValidateArgument("obs", options.obfuscation, []string{"IPv4", "IPv6", "MAC", "UUID"})
	}

	// Call function named ConvertShellcode2Template
	ConvertShellcode2Template(options.inputFile, foundLanguage)
}
