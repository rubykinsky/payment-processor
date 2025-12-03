package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrInvalidAmount   = errors.New("invalid amount")
	ErrInvalidCurrency = errors.New("invalid currency")
)

func GenerateTransactionID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate transaction ID: %w", err)
	}
	return hex.EncodeToString(b), nil
}

func ValidateAmount(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	return nil
}

func ValidateCurrency(currency string) error {
	currency = strings.ToUpper(currency)
	switch currency {
	case "USD", "EUR", "GBP", "JPY":
		return nil
	default:
		return ErrInvalidCurrency
	}
}

func FormatTimestamp(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

func MaskCardNumber(cardNumber string) string {
	if len(cardNumber) < 4 {
		return cardNumber
	}
	return "****-****-****-" + cardNumber[len(cardNumber)-4:]
}