package deptestsmall

import "testing"

func TestSillyTypeString(t *testing.T) {
	v := SillyType{}
	if s := v.String(); s != sillystr {
		t.Errorf("Bad output: got %s, want %s", s, sillystr)
	}
}
