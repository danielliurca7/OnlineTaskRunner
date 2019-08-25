package changes

import (
	"../change"
	"../fileinfo"
)

type Changes struct {
	Fileinfo fileinfo.Fileinfo
	Changes  []change.Change
}
