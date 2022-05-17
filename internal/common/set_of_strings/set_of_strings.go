package set_of_strings

type Type map[string]bool

func New(strings ...string) Type {
	result := make(Type, len(strings))
	for _, str := range strings {
		result.Add(str)
	}
	return result
}

func (s Type) Items() []string {
	keys := make([]string, len(s))

	i := 0
	for k := range s {
		keys[i] = k
		i++
	}

	return keys
}

func (s Type) Add(value string) {
	s[value] = true
}

func (s Type) Has(value string) bool {
	_, ok := s[value]

	return ok
}

func (s Type) Empty() bool {
	return len(s) == 0
}
