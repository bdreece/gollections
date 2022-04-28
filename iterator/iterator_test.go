/**
 * MIT License
 *
 * Copyright (c) 2022 Brian Reece
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package iterator

import (
	"github.com/bdreece/gollections/vector"
	"testing"
)

const (
	EXPECTED string = "expected (%d), got (%d)\n"
	ERROR    string = "experienced error: \"%s\"\n"
)

func setup(t *testing.T) (*vector.Vector[int], []int) {
	t.Log("Setting up test")
	vec := vector.New[int]()
	t.Log("Created vector")
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		vec.PushBack(number)
		t.Logf("Pushed (%d) to back of vector\n", number)
	}
	return vec, numbers
}

func TestForEach(t *testing.T) {
	vec, numbers := setup(t)
	iter := vec.Iterator()
	i := 0
	if err := ForEach[int](&iter, func(item *int) {
		if *item != numbers[i] {
			t.Errorf(EXPECTED, numbers[i], item)
			i++
		}
	}); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}
