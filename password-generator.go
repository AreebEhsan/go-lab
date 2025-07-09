package main

import (
    "crypto/rand"
    "flag"
    "fmt"
    "math/big"
    "os"
)

const (
    lowerChars  = "abcdefghijklmnopqrstuvwxyz"
    upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    numberChars = "0123456789"
    symbolChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func generatePassword(length int, useSymbols, useNumbers, useUpper bool) (string, error) {
    charset := lowerChars
    if useUpper {
        charset += upperChars
    }
    if useNumbers {
        charset += numberChars
    }
    if useSymbols {
        charset += symbolChars
    }

    if len(charset) == 0 {
        return "", fmt.Errorf("no character set selected")
    }

    password := make([]byte, length)
    for i := range password {
        num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            return "", err
        }
        password[i] = charset[num.Int64()]
    }
    return string(password), nil
}

func main() {
    length := flag.Int("len", 12, "Length of the password")
    useSymbols := flag.Bool("symbols", false, "Include symbols")
    useNumbers := flag.Bool("numbers", false, "Include numbers")
    useUpper := flag.Bool("upper", false, "Include uppercase letters")

    flag.Parse()

    pwd, err := generatePassword(*length, *useSymbols, *useNumbers, *useUpper)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }

    fmt.Println("Generated Password:", pwd)
}