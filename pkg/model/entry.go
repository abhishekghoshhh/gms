package model

type Entry struct {
	key   string
	value string
}

func NewEntry(key, val string) *Entry {
	return &Entry{
		key,
		val,
	}
}
