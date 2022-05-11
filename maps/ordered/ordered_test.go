package ordered

const (
	ERROR    string = "experienced error: \"%s\"\n"
	EXPECTED string = "expected %s: (%v), got: (%v)\n"
)

func setup() (*Map[string, int], map[string]int) {
	ours := New[string, int]()
	theirs := make(map[string]int)
	keys := []string{"apple", "banana", "coffee", "donut"}
	vals := []int{1, 2, 3, 4}
	for i, key := range keys {
		ours.Ins(key, vals[i])
		theirs[key] = vals[i]
	}
	return ours, theirs
}
