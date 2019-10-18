package containers

import (
	"../change"
	"../file"
	"../fileinfo"
	"../workspace"
)

type ChangeContainer struct {
	Workspace workspace.Workspace
	Fileinfo  fileinfo.Fileinfo
	Change    change.Change
}

type OneFileContainer struct {
	Workspace workspace.Workspace
	Fileinfo  fileinfo.Fileinfo
	Data      string
}

type TwoFileinfoContainer struct {
	Workspace workspace.Workspace
	Fileinfo  [2]fileinfo.Fileinfo
}

type WorkspaceContainer struct {
	Workspace workspace.Workspace
	Files     []file.File
}
