package datastructures

type Course struct {
	Name         string `json:"name"`
	Year         int    `json:"year"`
	Series       string `json:"series"`
	Abbreviation string `json:"abbreviation"`
}

func (c1 Course) Equals(c2 Course) bool {
	return c1.Name == c2.Name && c1.Year == c2.Year && c1.Series == c2.Series
}

type Student struct {
	Year    int      `json:"year"`
	Group   int      `json:"group"`
	Series  string   `json:"series"`
	Courses []Course `json:"courses"`
}

type Assistant struct {
	Courses []Course `json:"courses"`
}

type Professor struct {
	Courses []Course `json:"courses"`
}

type UserData struct {
	Name      string    `json:"name"`
	Student   Student   `json:"student"`
	Assistant Assistant `json:"assistant"`
	Professor Professor `json:"professor"`
}
