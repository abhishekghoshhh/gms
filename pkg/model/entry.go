package model

type entry struct {
	key   string
	value string
}

func Entry(key, val string) *entry {
	return &entry{
		key,
		val,
	}
}
