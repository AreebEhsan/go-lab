package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/atotto/clipboard"
)

func generateOTP(length int) (string, error) {
	otp := ""
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += fmt.Sprintf("%d", n)
	}
	return otp, nil
}

func main() {
	const otpLength = 6
	const timeout = 30 * time.Second

	otp, err := generateOTP(otpLength)
	if err != nil {
		log.Fatalf("Failed to generate OTP: %v", err)
	}

	err = clipboard.WriteAll(otp)
	if err != nil {
		log.Fatalf("Failed to copy to clipboard: %v", err)
	}

	fmt.Printf("OTP '%s' copied to clipboard. It will be cleared in %d seconds.\n", otp, int(timeout.Seconds()))

	time.AfterFunc(timeout, func() {
		err := clipboard.WriteAll("")
		if err != nil {
			log.Println("Failed to clear clipboard:", err)
		} else {
			fmt.Println("OTP cleared from clipboard.")
		}
	})

	time.Sleep(timeout + 1*time.Second)
}
