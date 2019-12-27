package types

// Int return the value of `v`.
func Int(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// IntPtr return a pointer of `v`.
func IntPtr(v int) *int {
	return &v
}

// Int64 return the value of `v`.
func Int64(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Ptr return a pointer of `v`.
func Int64Ptr(v int64) *int64 {
	return &v
}

// Float64 return the value of `v`.
func Float64(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Ptr return a pointer of `v`.
func Float64Ptr(v float64) *float64 {
	return &v
}

// String return the value of `v`.
func String(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringPtr return the value of `v`.
func StringPtr(v string) *string {
	return &v
}
