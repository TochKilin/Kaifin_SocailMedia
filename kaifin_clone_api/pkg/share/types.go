package share

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserFormCreate struct {
	FirstName   string
	LastName    string
	Nationality Nationality
}

type UserFormUpdate struct {
	Id            int
	FirstName     string
	LastName      string
	NationalityId int
	Nationality   Nationality
}

type Nationality struct {
	//
	Id     int
	Valvue string
}

// /////
type UserContext struct {
	UserID       int64     `json:"user_id"`
	UserName     string    `json:"user_name"`
	LoginSession string    `json:"login_session"`
	Exp          time.Time `json:"exp"`
	UserAgent    string    `json:"user_agent"`
	Ip           string    `json:"ip"`
	RoleID       int       `json:"role_id"`
}

type Paging struct {
	Page    int `json:"page" query:"page" validate:"required,min=1"`
	Perpage int `json:"per_page" query:"per_page" validate:"required,min=1"`
}

// sort product culomn price with a to asd asc dend from frontend
type Sort struct {
	Property  string `json:"property" validate:"required"`                 // column db table
	Direction string `json:"direction" validate:"required,oneof=asc desc"` // key  order by
}

// search or filter data
type Filter struct {
	Property string      `json:"property" validate:"required" query:"property"` // column db table
	Value    interface{} `json:"value" validate:"required" query:"value"`       // key  order by asc or desc
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// StatusData contains predefined status values.
var StatusData = []Status{
	{ID: 1, Name: "Active"},
	{ID: 2, Name: "Inactive"},
	{ID: 3, Name: "Suspended"},
	{ID: 4, Name: "Deleted"},
}

// GenerateToken បង្កើត JWT token សម្រាប់ user ដែល login បានជោគជ័យ។
func GenerateToken(userID float64, userName string, loginSession string, roleID int) (string, time.Time, error) {

	// កំណត់ពេលផុតកំណត់របស់ token ឲ្យមានសុពលភាព 24 ម៉ោងចាប់ពីពេលបង្កើត។
	expirationTime := time.Now().Add(24 * time.Hour) // from .env JWT_EXPIRE

	// រៀបចំទិន្នន័យដែលត្រូវដាក់ខាងក្នុង JWT token។
	claims := jwt.MapClaims{
		"user_id":       userID,
		"user_name":     userName,
		"login_session": loginSession,
		"exp":           expirationTime.Unix(),
		"role_id":       roleID,
	}

	// បង្កើត JWT token ដោយប្រើ signing method HS256។
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token ដោយប្រើ secret key ពី environment variable JWT_SECRET_KEY។
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		// បើ sign token មិនបាន ត្រឡប់ error ទៅ caller។
		return "", time.Time{}, err
	}

	// ត្រឡប់ token string, ពេលផុតកំណត់ និង nil error ពេលបង្កើតបានជោគជ័យ។
	return tokenString, expirationTime, nil
}
