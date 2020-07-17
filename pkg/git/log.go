package git

import (
	"strings"
)

// Log of git
type Log struct {
	ShortCode    string
	Message      string
	CoAuthoredBy string
}

// CreateLog to create git log from raw message
func CreateLog(raw string) *Log {
	if len(raw) < 7 {
		return nil
	}
	raw = strings.TrimSpace(raw)
	message := raw[7:]
	coAuthoredBy := ""
	if i := strings.Index(message, "Co-Authored-By:"); i >= 0 {
		coAuthoredBy = message[i+len("Co-Authored-By:"):]
		message = message[:i]

	}
	return &Log{
		ShortCode:    strings.TrimSpace(raw[:7]),
		Message:      strings.TrimSpace(message),
		CoAuthoredBy: strings.TrimSpace(coAuthoredBy),
	}
}
