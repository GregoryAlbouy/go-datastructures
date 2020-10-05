package compare

// A Func compares two elements a and b of the same type.
// It should return -1 if a < b, or 1 if a > b, or 0 if a == b.
type Func func(a, b interface{}) int

// Sup returns true if a > b, given a compare.Func f.
func Sup(a, b interface{}, f Func) bool {
	return f(a, b) == 1
}

// Inf returns true if a < b, given a compare.Func f.
func Inf(a, b interface{}, f Func) bool {
	return f(a, b) == -1
}

// Eq returns true if a == b, given a compare.Func f.
func Eq(a, b interface{}, f Func) bool {
	return f(a, b) == 0
}

// SupOrEq returns true if a >= b, given a compare.Func f.
func SupOrEq(a, b interface{}, f Func) bool {
	return Sup(a, b, f) || Eq(a, b, f)
}

// InfOrEq returns true if a <= b, given a compare.Func f.
func InfOrEq(a, b interface{}, f Func) bool {
	return Inf(a, b, f) || Eq(a, b, f)
}
