package datastructures

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Course struct {
	Name         string `json:"name"`
	Year         int    `json:"year"`
	Series       string `json:"series"`
	Abbreviation string `json:"abbreviation"`
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
