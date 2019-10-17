package gen

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestMakeIndex(t *testing.T) {
	var r1, r2 rune
	index1 := make([]string, 2)
	index2 := make([]string, 2)
	index1 = MakeIndex(1, 2)
	index2 = MakeIndex(2, 1)

	if (index1 != nil) && (index1[0] != index1[1]) {
		r1, _ = utf8.DecodeRuneInString(index1[0])
		r2, _ = utf8.DecodeRuneInString(index1[0])
		fmt.Println("Condition 1 has passed.")
	} else {
		t.Errorf("Condition 1 has failed.")
	}
	fmt.Println(r1)
	fmt.Println(r2)

	if (index2 != nil) && (index2[0] != index2[1]) {
		r1, _ = utf8.DecodeLastRuneInString(index2[0])
		r2, _ = utf8.DecodeLastRuneInString(index2[0])
		fmt.Println("Condition 2 has passed.")
	} else {
		t.Errorf("Condition 2 has failed.")
	}
	fmt.Println(r1)
	fmt.Println(r2)
}
