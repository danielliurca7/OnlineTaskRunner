package fileinfo

import "../workspace"

type Fileinfo struct {
	workspace.Workspace
	Path []string
}

func (fi *Fileinfo) Equals(other *Fileinfo) bool {
	ownerEq := fi.Owner == other.Owner
	subjectEq := fi.Subject == other.Subject
	assignmentNameEq := fi.AssignmentName == other.AssignmentName
	yearEq := fi.Year == other.Year
	pathEq := len(fi.Path) == len(other.Path)
	if pathEq {
		for index, item := range fi.Path {
			if item != other.Path[index] {
				pathEq = false
			}
		}
	}
	return ownerEq && subjectEq && assignmentNameEq && yearEq && pathEq
}
