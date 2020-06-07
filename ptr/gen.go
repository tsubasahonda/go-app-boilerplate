package ptr

import "time"

// NewString gen string pointer
func NewString(s string) *string {
	return &s
}

// NewInt64 gen int64 pointer
func NewInt64(i int64) *int64 {
	return &i
}

// NewTime gen time.Time pointer
func NewTime(t time.Time) *time.Time {
	return &t
}
