package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Helper function to generate a random verification code
func GenerateVerificationCode() string {
	// Инициализация генератора случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%04d", r.Intn(10000))
}
