// MIT License
// Copyright (c) 2022 Brian Reece

package iterator

import "fmt"

type ZipError struct {
	a error
	b error
}

func (z ZipError) Error() string {
	if z.a != nil && z.b != nil {
		return fmt.Sprintf("iter a: (%s), iter b: (%s)", z.a.Error(), z.b.Error())
	} else if z.a != nil {
		return fmt.Sprintf("iter a: (%s)", z.a.Error())
	} else if z.b != nil {
		return fmt.Sprintf("iter b: (%s)", z.b.Error())
	} else {
		return "unreachable! bug!"
	}
}

type ZipItem[T, U any] struct {
	A *T
	B *U
}

type Zip[T, U any] struct {
	a Iterator[T]
	b Iterator[U]
}

func NewZip[T any, U any](a Iterator[T], b Iterator[U]) *Zip[T, U] {
	return &Zip[T, U]{a, b}
}

func (z *Zip[T, U]) Next() (ZipItem[T, U], error) {
	a, err_a := z.a.Next()
	b, err_b := z.b.Next()

	item := ZipItem[T, U]{a, b}

	if err_a != nil || err_b != nil {
		return item, ZipError{
			a: err_a,
			b: err_b,
		}
	} else {
		return item, nil
	}
}
