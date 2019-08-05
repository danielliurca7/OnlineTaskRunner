package changes

import (
	"./change"
	"./file"
)

type Changes struct {
	file File
	changes []Change
}
