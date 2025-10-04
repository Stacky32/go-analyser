package testdata

import (
	"bytes"
	"slices"
	"strings"
)

func Do3() {
	_ = bytes.Buffer{}
	slices.Sort([]int{})
	_ = strings.ToUpper("a")
}
