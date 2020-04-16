package utils

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"

	datastructures "../data_structures"
)

type claims struct {
	Payload datastructures.UserData `json:"payload"`
	jwt.StandardClaims
}

// BuildToken build a token based on the user data and the SECRET_KEY environment
func BuildToken(userdata datastructures.UserData) ([]byte, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		Payload: userdata,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return []byte(tokenString), err
}

// VerifyToken verifies a token based on the string token and the SECRET_KEY environment
func VerifyToken(tokenString string) (datastructures.UserData, error) {
	claims := &claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return datastructures.UserData{}, err
	} else if !token.Valid {
		return datastructures.UserData{}, errors.New("Invalid token")
	}

	return claims.Payload, nil
}

// VerifyAuthorization verifies if the user has the permision to access a resourse
func VerifyAuthorization(workspace datastructures.Workspace, userdata datastructures.UserData) bool {
	course := datastructures.Course{
		Name:   workspace.Course,
		Year:   workspace.Year,
		Series: workspace.Series,
	}

	for _, studentCourse := range userdata.Student.Courses {
		if course.Equals(studentCourse) && workspace.Owner == userdata.Name {
			return true
		}
	}

	for _, assistantCourse := range userdata.Assistant.Courses {
		if course.Equals(assistantCourse) {
			return true
		}
	}

	for _, professorCourse := range userdata.Professor.Courses {
		if course.Equals(professorCourse) {
			return true
		}
	}

	return false
}
