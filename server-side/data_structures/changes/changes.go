package changes

import (
	"./change"
	"./file"
)

type Changes struct {
	File File
	Changes []Change
}
