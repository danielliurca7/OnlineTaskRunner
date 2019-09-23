package containers

import (
	"../change"
	"../file"
	"../fileinfo"
	"../workspace"
)

type OneChangeContainer struct {
	Fileinfo fileinfo.Fileinfo
	Change   change.Change
}

type WorkspaceContainer struct {
	Workspace workspace.Workspace
	Files     []file.File
}
