// package deptestsmall is a tiny test lib with no dependencies.

package deptestsmall

type SillyType struct{}

const sillystr = "I'm a silly test type with no package deps"

func (s SillyType) String() string {
	return sillystr
}
