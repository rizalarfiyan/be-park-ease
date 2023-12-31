package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

func RandSuffix(long int) string {
	b := make([]byte, long)
	if _, err := rand.Read(b); err != nil {
		return strings.Repeat("A", long)
	}
	return fmt.Sprintf("%X", b)
}

func GenerateEntryHistoryId() string {
	saveLong := 16
	idx := fmt.Sprintf("H%d%s", time.Now().Nanosecond(), RandSuffix(5))
	if len(idx) > saveLong {
		idx = idx[:saveLong]
	}
	return idx
}
