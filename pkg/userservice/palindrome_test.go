package userservice

import "testing"

func FuzzIsPalindrome(f *testing.F) {

	f.Add("appa", true)
	f.Add("geagea", false)
	f.Add("potatop", true)
	f.Add("papa", false)
	f.Add("paap", true)
	f.Add("hoh", true)
	f.Add("hohoh", true)
	f.Add("a", true)
	f.Fuzz(func(t *testing.T, in string, expected bool) {
		out := isPalindrome(in)
		if out != expected {
			t.Errorf("Got %v, expected %v for %v", out, expected, in)
		}
	})
}
