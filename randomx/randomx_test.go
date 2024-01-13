package randomx

import (
	"strings"
	"testing"
)

func TestProvider_GenerateRandomString(t *testing.T) {
	rx := New().(*provider)

	tests := []struct {
		length  int
		charSet string
	}{
		{10, LowercaseLetters + UppercaseLetters + Digits + Symbols},
		{5, LowercaseLetters},
		{8, Digits},
	}

	for _, test := range tests {
		randomStr := rx.GenerateRandomString(test.length, test.charSet)
		if len(randomStr) != test.length {
			t.Errorf("Generated string length does not match expected length for charSet %s", test.charSet)
		}

		for _, char := range randomStr {
			if !strings.Contains(test.charSet, string(char)) {
				t.Errorf("Generated string contains unexpected character: %s", string(char))
			}
		}
	}
}

func TestProvider_GenerateRandomStringFromBuiltin(t *testing.T) {
	rx := New().(*provider)

	defaultOptions := GenerateOptions{
		UseLowercase: true,
		UseUppercase: true,
		UseDigits:    true,
		UseSymbols:   true,
	}
	randomDefault := rx.GenerateRandomStringFromBuiltin(15, defaultOptions)
	for _, char := range randomDefault {
		if !strings.Contains(LowercaseLetters+UppercaseLetters+Digits+Symbols, string(char)) {
			t.Errorf("Generated string contains unexpected character: %s", string(char))
		}
	}

	customOptions := GenerateOptions{
		UseLowercase: true,
		UseUppercase: true,
		UseDigits:    false,
		UseSymbols:   true,
	}
	randomCustom := rx.GenerateRandomStringFromBuiltin(12, customOptions)
	for _, char := range randomCustom {
		if !strings.Contains(LowercaseLetters+UppercaseLetters+Symbols, string(char)) {
			t.Errorf("Generated string contains unexpected character: %s", string(char))
		}
	}
}

func TestProvider_GenerateRandomStringFromBuiltin_PredefinedOptions(t *testing.T) {
	rx := New().(*provider)

	tests := []struct {
		length  int
		options GenerateOptions
	}{
		{15, OptionAllChars},
		{10, OptionLetters},
		{8, OptionDigitsOnly},
	}

	for _, test := range tests {
		randomStr := rx.GenerateRandomStringFromBuiltin(test.length, test.options)

		expectedCharSet := ""
		if test.options.UseLowercase {
			expectedCharSet += LowercaseLetters
		}
		if test.options.UseUppercase {
			expectedCharSet += UppercaseLetters
		}
		if test.options.UseDigits {
			expectedCharSet += Digits
		}
		if test.options.UseSymbols {
			expectedCharSet += Symbols
		}

		if len(randomStr) != test.length {
			t.Errorf("Generated string length does not match expected length for options %+v", test.options)
		}

		for _, char := range randomStr {
			if !strings.Contains(expectedCharSet, string(char)) {
				t.Errorf("Generated string contains unexpected character: %s", string(char))
			}
		}
	}
}
