package chain_test

import (
	"fmt"
	"testing"

	"github.com/bdreece/gollections/pkg/iterator"
	"github.com/bdreece/gollections/pkg/iterator/chain"
	"github.com/bdreece/gollections/pkg/slice"
)

func TestChainOne(t *testing.T) {
	iter := slice.Marshal([]int{11, 22, 33, 44, 55}).Iter()
	oddStrings := chain.From[int, string](iter).
		Filter(func(i int) bool {
			return i%2 != 0
		}).
		Map(func(i int) string {
			return fmt.Sprint(i)
		}).
		Value()

	char, err := chain.From[string, rune](oddStrings).
		Take(3).
		FlatMap(func(s string) iterator.IntoIterator[rune] {
			return slice.Marshal([]rune(s))
		}).
		Find(func(r rune) bool {
			return r == rune(51)
		})

	if err != nil {
		t.Error(err)
	}

	if char == nil {
		t.Errorf("t is nil")
		return
	}

	if *char != '3' {
		t.Errorf("t == %c\n", *char)
	}
}
