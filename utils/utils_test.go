package utils_test

import (
	"fmt"
	"testing"

	"github.com/harrycmfan/golang-leetcode/utils"
)

func TestCapToLowerCase(t *testing.T) {
	fmt.Println(utils.CapToLowerCase("Kth Largest Element in an Array"))
}

// Interpreted string literals are character sequences between double quotes "" using the (possibly multi-byte) UTF-8 encoding of individual characters.
// In UTF-8, ASCII characters are single-byte corresponding to the first 128 Unicode characters.
func TestTypeOfSubstring(t *testing.T) {
	defer func() {}()
	s := "abc"
	fmt.Printf("%v, %T",s[0], s[0])

	fmt.Println(map[uint8]int{s[0]:1})

	fmt.Println(("2"[0]-48) * ("5"[0]-48))
}