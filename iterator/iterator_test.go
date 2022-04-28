package iter

import (
	"github.com/bdreece/gollections/vector"
	"testing"
)

func TestForEach(t *testing.T) {
	vec := vector.New[int]()
	vec.Fill(1, 2, 3, 4, 5)
}
