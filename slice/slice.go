package slice

/*
 * IsContains 元素是否包含
 */
func IsContains(item string, sl []string) bool {
	set := make(map[string]struct{}, len(sl))
	for _, s := range sl {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

/*
 * IsContainsInterface 元素是否包含（interface）
 */
func IsContainsInterface(item interface{}, sl []interface{}) bool {
	set := make(map[interface{}]struct{}, len(sl))
	for _, s := range sl {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

/*
 * DeleteValueByIndex 删除slice中的元素
 */
func DeleteValueByIndex(obj []interface{}, idx int) []interface{} {
	if idx >= len(obj) {
		return obj
	}
	copy(obj[idx:], obj[idx+1:])
	obj = obj[:len(obj)-1]
	return obj
}
