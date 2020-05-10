package datastructures

// Course models the course tabel in the database
type Course struct {
	Name         string `json:"name"`
	Year         int    `json:"year"`
	Series       string `json:"series"`
	Abbreviation string `json:"abbreviation"`
}

// Equals checks the equality between courses
func (c1 Course) Equals(c2 Course) bool {
	return c1.Name == c2.Name && c1.Year == c2.Year && c1.Series == c2.Series
}

// Student models the student tabel in the database
type Student struct {
	Year    int      `json:"year"`
	Group   int      `json:"group"`
	Series  string   `json:"series"`
	Courses []Course `json:"courses"`
}

// Assistant models the assistant tabel in the database
type Assistant struct {
	Courses []Course `json:"courses"`
}

// Professor models the professor tabel in the database
type Professor struct {
	Courses []Course `json:"courses"`
}

// UserData is the data structure that holds all of the data required by a user
type UserData struct {
	Name      string    `json:"name"`
	Student   Student   `json:"student"`
	Assistant Assistant `json:"assistant"`
	Professor Professor `json:"professor"`
}
