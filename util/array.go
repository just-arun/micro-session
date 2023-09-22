package util

type array struct{}

func Array() array {
	return array{}
}

func (st array) Includes(arr []string, cb func(item string, index int) bool) (includes bool) {
	for i, v := range arr {
		cond := cb(v, i)
		if cond {
			return cond
		}
	}
	return false
}
