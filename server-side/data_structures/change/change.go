package change

type Key struct {
	row int
	column int
}

type Change struct {
	key Key
	value byte
}
