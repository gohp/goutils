package slice

func IsContains(item string, sl []string) bool {
	set := make(map[string]struct{}, len(sl))
	for _, s := range sl {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func IsContainsInterface(item interface{}, sl []interface{}) bool {
	set := make(map[interface{}]struct{}, len(sl))
	for _, s := range sl {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
