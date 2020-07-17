package strdemo3

import (
	"fmt"
	"testing"
)

func TestTrim(t *testing.T) {
	fmt.Println(len(trim(" . Hello   World .  ")))
	fmt.Println(len(". Hello   World ."))
	fmt.Println(len(trim(" . Hello   World .  "))==len(". Hello   World ."))
}
