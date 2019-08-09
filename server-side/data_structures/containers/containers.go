package containers

import (
	"../fileinfo"
	"../change"
	"../changes"
)

type OneChangeContainer struct {
	Fileinfo fileinfo.Fileinfo
	Change change.Change
}

type MultipleChangesContainer struct {
	Fileinfo fileinfo.Fileinfo
	Changes changes.Changes
}
