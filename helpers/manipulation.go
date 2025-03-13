package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func NormalizeDate(date string) string {
	parts := strings.Split(date, ".")
	if len(parts) != 3 {
		return date // Invalid format, return as-is
	}

	day, _ := strconv.Atoi(parts[0])   // Convert to int to remove leading zeros
	month, _ := strconv.Atoi(parts[1]) // Convert to int to remove leading zeros
	year := parts[2]                   // Keep year as a string

	return fmt.Sprintf("%d.%d.%s", day, month, year)
}

func ShortenString(str string, length int) string {
	if len(str) > length {
		return str[:length-3] + "..."
	}
	return str
}

func SetDefaultInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		num = 0
	}
	return num
}
