package strutil

// Equals return true if the strings are equals
func Equals(f string, s string) (bool, error) {
	if len([]byte(f)) == len([]byte(s)) {
		for i := range []byte(f) {
			if f[i] != s[i] {
				return false, nil
			}
		}
		return true, nil
	}
	return false, nil
}
