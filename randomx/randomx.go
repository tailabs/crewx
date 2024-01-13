package randomx

import (
	"math/rand"
	"time"
)

// Character sets
const (
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits           = "0123456789"
	Symbols          = "!@#$%^&*()-_=+[]{}|;:,.<>?/~"
)

// Predefined GenerateOptions for commonly used character sets
var (
	OptionAllChars = GenerateOptions{
		UseLowercase: true,
		UseUppercase: true,
		UseDigits:    true,
		UseSymbols:   true,
	}

	OptionLetters = GenerateOptions{
		UseLowercase: true,
		UseUppercase: true,
		UseDigits:    false,
		UseSymbols:   false,
	}

	OptionDigitsOnly = GenerateOptions{
		UseLowercase: false,
		UseUppercase: false,
		UseDigits:    true,
		UseSymbols:   false,
	}
)

// Interface defines the methods for generating random strings.
type Interface interface {
	GenerateRandomString(length int, charSet string) string
	GenerateRandomStringFromBuiltin(length int, options GenerateOptions) string
}

// provider is the random string generator.
type provider struct {
	randSrc *rand.Rand
}

// New creates a new random string provider.
func New() Interface {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	return &provider{
		randSrc: rand.New(source),
	}
}

// GenerateRandomString generates a random string of the specified length using the given character set.
func (p *provider) GenerateRandomString(length int, charSet string) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charSet[p.randSrc.Intn(len(charSet))]
	}
	return string(result)
}

// GenerateOptions is the parameter options for GenerateRandomStringFromBuiltin.
type GenerateOptions struct {
	UseLowercase bool
	UseUppercase bool
	UseDigits    bool
	UseSymbols   bool
}

// GenerateRandomStringFromBuiltin generates a random string of the specified length using predefined or customized character sets.
func (p *provider) GenerateRandomStringFromBuiltin(length int, options GenerateOptions) string {
	charSetBuilder := ""
	if options.UseLowercase {
		charSetBuilder += LowercaseLetters
	}
	if options.UseUppercase {
		charSetBuilder += UppercaseLetters
	}
	if options.UseDigits {
		charSetBuilder += Digits
	}
	if options.UseSymbols {
		charSetBuilder += Symbols
	}

	return p.GenerateRandomString(length, charSetBuilder)
}
