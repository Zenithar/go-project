package helpers

import "github.com/dchest/uniuri"

// IDGeneratorFunc is the identifier generator
func IDGeneratorFunc() string {
	return uniuri.NewLen(32)
}
