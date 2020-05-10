package utils

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	datastructures "../data_structures"
)

var postgres *sql.DB
var mongodb *mongo.Database
var mongoctx context.Context

func init() {
	var err error

	if postgres, err = sql.Open("postgres", "postgres://root:root@postgres:5432/university?sslmode=disable"); err != nil {
		panic(err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@mongo:27017"))
	if err != nil {
		panic(err)
	}

	mongoctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := client.Connect(mongoctx); err != nil {
		panic(err)
	}

	mongodb = client.Database("otr_data")
}

// VerifyCredentials receives a credentials datastructures
// and query the database to see if they are valid
func VerifyCredentials(credentials datastructures.Credentials) (bool, error) {
	sum := sha256.Sum256([]byte(credentials.Password))

	username := credentials.Username
	password := fmt.Sprintf("%x", sum)

	var exists bool
	row := postgres.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE user_name=$1 AND password_hash=$2);", username, password)
	if err := row.Scan(&exists); err != nil {
		return false, err
	} else if !exists {
		return false, nil
	}

	return true, nil
}

// GetUserData returns the data neccesary to build a token for a user
func GetUserData(username string) (datastructures.UserData, error) {
	var userdata datastructures.UserData
	userdata.Name = username

	var year, group int
	var course, series, abbreviation string

	row := postgres.QueryRow(`
		SELECT school_year, group_no, series.name FROM students
		JOIN groups ON students.group_id = groups.id
		JOIN series ON groups.series_id = series.id
		WHERE students.name=$1;
	`, username)

	if err := row.Scan(&year, &group, &series); err != nil {
		return datastructures.UserData{}, err
	}

	userdata.Student.Year = year
	userdata.Student.Group = group
	userdata.Student.Series = series

	rows, err := postgres.Query(`
		SELECT courses.name, school_year, series.name, abbreviation FROM students_courses
		JOIN courses ON students_courses.course_id = courses.id
		JOIN series ON courses.id = series.id
		WHERE student_name=$1;
	`, username)

	if err != nil {
		return datastructures.UserData{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&course, &year, &series, &abbreviation); err != nil {
			return datastructures.UserData{}, err
		}

		userdata.Student.Courses = append(userdata.Student.Courses, datastructures.Course{
			Name:         course,
			Year:         year,
			Series:       series,
			Abbreviation: abbreviation,
		})
	}

	rows, err = postgres.Query(`
		SELECT courses.name, school_year, series.name, abbreviation FROM assistants_courses
		JOIN courses ON assistants_courses.course_id = courses.id
		JOIN series ON courses.id = series.id
		WHERE assistant_name=$1;
	`, username)

	if err != nil {
		return datastructures.UserData{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&course, &year, &series, &abbreviation); err != nil {
			return datastructures.UserData{}, err
		}

		userdata.Assistant.Courses = append(userdata.Assistant.Courses, datastructures.Course{
			Name:         course,
			Year:         year,
			Series:       series,
			Abbreviation: abbreviation,
		})
	}

	rows, err = postgres.Query(`
		SELECT courses.name, school_year, series.name, abbreviation FROM courses
		JOIN series ON courses.id = series.id
		WHERE professor=$1;
	`, username)

	if err != nil {
		return datastructures.UserData{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&course, &year, &series, &abbreviation); err != nil {
			return datastructures.UserData{}, err
		}

		userdata.Professor.Courses = append(userdata.Professor.Courses, datastructures.Course{
			Name:         course,
			Year:         year,
			Series:       series,
			Abbreviation: abbreviation,
		})
	}

	return userdata, nil
}

// GetCourses returns courses for an username and a collection
func GetCourses(username string, collection string) ([]byte, error) {
	col := mongodb.Collection(collection)

	var course bson.M

	if err := col.FindOne(
		mongoctx,
		bson.M{"name": username},
		options.FindOne().SetProjection(bson.M{"_id": 0, "courses": 1}),
	).Decode(&course); err != nil && err.Error() != "mongo: no documents in result" {
		return nil, err
	}

	b, err := json.Marshal(course)
	if err != nil {
		return nil, err
	}

	return b, nil
}
