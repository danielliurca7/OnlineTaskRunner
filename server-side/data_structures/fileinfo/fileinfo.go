package fileinfo

type Fileinfo struct {
	Path  []string
	IsDir bool
}

func (fi *Fileinfo) Equals(other *Fileinfo) bool {
	pathEq := len(fi.Path) == len(other.Path)

	if pathEq {
		for index, item := range fi.Path {
			if item != other.Path[index] {
				pathEq = false
			}
		}
	}
	isDirEq := fi.IsDir == other.IsDir
	return pathEq && isDirEq
}
