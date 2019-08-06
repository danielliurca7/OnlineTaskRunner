package change

type Key struct {
	Row int
	Column int
}

type Change struct {
	Key Key
	Value byte
}
