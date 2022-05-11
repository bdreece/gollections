package unordered

import (
	"testing"
)

func TestGet(t *testing.T) {
	ours, theirs := setup()

	for key := range theirs {
		val, err := ours.Get(key)
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != theirs[key] {
			t.Errorf(EXPECTED, "val", theirs[key], *val)
		}
	}
}

func TestSet(t *testing.T) {
	ours, theirs := setup()

	for key, val := range theirs {
		if err := ours.Set(key, val); err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if value, err := ours.Get(key); err != nil {
			t.Errorf(ERROR, err.Error())
		} else if *value != val {
			t.Errorf(EXPECTED, "val", val, *value)
		}
	}

	if err := ours.Set("giraffe", 6); err == nil {
		t.Errorf(EXPECTED, "err", "errors.NotFound", "nil")
	}
}

func TestIns(t *testing.T) {
	ours, theirs := setup()

	for key := range theirs {
		if err := ours.Ins(key, 0); err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if val, err := ours.Get(key); err != nil {
			t.Errorf(ERROR, err.Error())
		} else if *val != 0 {
			t.Errorf(EXPECTED, "val", 0, *val)
		}
	}

	if err := ours.Ins("giraffe", 5); err != nil {
		t.Errorf(ERROR, err.Error())
	}
}

func TestDel(t *testing.T) {
	ours, theirs := setup()

	for key := range theirs {
		val, err := ours.Del(key)
		if err != nil {
			t.Errorf(ERROR, err.Error())
		}
		if *val != theirs[key] {
			t.Errorf(EXPECTED, "val", theirs[key], *val)
		}
		if val, err := ours.Get(key); err == nil {
			t.Errorf(EXPECTED, "error", "errors.NotFound", *val)
		}
	}
}
