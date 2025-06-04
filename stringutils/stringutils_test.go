package stringutils

import (
	"testing"
	"unicode/utf8"
)

// Fuzz test for UnSafeReverseString function
func FuzzUnSafeReverseString(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		rev := UnSafeReverseString(orig)
		doubleRev := UnSafeReverseString(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q",
				rev)
		}
	})
}

// Fuzz test for ReverseString function
func FuzzReverseString(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := ReverseString(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := ReverseString(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
