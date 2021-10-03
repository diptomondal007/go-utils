package variables

import (
	"strconv"
	"strings"
)

// ToInt32 returns parsed int32 from a given string
func ToInt32(value string) int32 {
	if s, err := strconv.ParseInt(strings.TrimSpace(value), 10, 32); err == nil {
		return int32(s)
	}
	return 0
}

// IntP returns pointer of a given int variable
func IntP(v int) *int {
	return &v
}

// Int32P returns pointer of a given int32 variable
func Int32P(v int32) *int32 {
	return &v
}

// Int64P returns pointer of a given int variable
func Int64P(v int64) *int64 {
	return &v
}

// StringP returns pointer of a given string variable
func StringP(v string) *string {
	return &v
}

// Float32P returns pointer of a given float32 variable
func Float32P(v float32) *float32 {
	return &v
}

// Float64P returns pointer of a given float64 variable
func Float64P(v float64) *float64 {
	return &v
}

// BoolP returns pointer of a given boolean variable
func BoolP(v bool) *bool {
	return &v
}

// MinInt returns minimum of two given integer value
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns minimum of two given integer value
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
