package changes

import (
	"../fileinfo"	
	"../change"
)

type Changes struct {
	Fileinfo fileinfo.Fileinfo
	Changes []change.Change
}
