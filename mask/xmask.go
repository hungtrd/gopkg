package xmask

import "strings"

const (
	maskChar   = '*'
	maxMaskLen = 5
)

// String masks the middle of s, keeping unmaskedPrefix runes at start and
// unmaskedSuffix runes at end. If s is too short to mask, returns s unchanged.
// The mask portion is capped at maxMaskLen characters regardless of actual length.
func String(s string, unmaskedPrefix int, unmaskedSuffix int) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return ""
	}
	if unmaskedPrefix+unmaskedSuffix >= n {
		return s
	}
	maskLen := min(n-unmaskedPrefix-unmaskedSuffix, maxMaskLen)
	var b strings.Builder
	b.WriteString(string(runes[:unmaskedPrefix]))
	b.WriteString(strings.Repeat(string(maskChar), maskLen))
	if unmaskedSuffix > 0 {
		b.WriteString(string(runes[n-unmaskedSuffix:]))
	}
	return b.String()
}

// Prefix shows the first n runes and masks the rest.
// e.g. Prefix("abc123xyz", 3) → "abc******"
func Prefix(s string, n int) string {
	return String(s, n, 0)
}

// Suffix shows the last n runes and masks the rest.
// e.g. Suffix("0901234567", 4) → "******4567"
func Suffix(s string, n int) string {
	return String(s, 0, n)
}

// Email masks the local part of an email, keeping the first 2 characters.
// e.g. "hungtran@domain.com" → "hu*****@domain.com"
func Email(email string) string {
	at := strings.LastIndex(email, "@")
	if at < 0 {
		return String(email, 2, 2)
	}
	return String(email[:at], 2, 2) + email[at:]
}

// PhoneNumber masks a phone number, keeping the last 4 characters.
// e.g. "0901234567" → "******4567"
func PhoneNumber(phone string) string {
	return Suffix(phone, 4)
}
