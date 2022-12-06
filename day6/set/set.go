package set

type Set map[string]bool

func New(str string) Set {
	set := Set{}

	for _, s := range str {
		if set[string(s)] {
			continue
		}

		set[string(s)] = true
	}

	return set
}
