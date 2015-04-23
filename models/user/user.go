package user
import (
	"github.com/yydzero/cherry/models"
	"github.com/ninedata/goat"
	"database/sql"
	"log"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
}

var db *sqlx.DB

func init() {
	db = models.GetDB()
}

func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func HasUser(u *User) (bool, error) {
	u1, err := Find(u)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return u1 != nil, nil
}

func Find(u *User) (*User, error) {
	var u1 User
	row := db.QueryRowx(db.Rebind("SELECT id, email, name, password FROM users WHERE email = ?"), u.Email)
	err := row.StructScan(&u1)

	if err != nil {
		return nil, err
	}

	return &u1, nil
}

func FindByEmailPassword(u *User) (*User, error) {
	var u1 User
	row := db.QueryRowx(db.Rebind("SELECT id, name, email  FROM users WHERE email = ? and password = ?"), u.Email, u.Password)
	err := row.StructScan(&u1)

	if err != nil {
		return nil, err
	}

	return &u1, nil
}

func (u *User) Save() error {
	var id int

	err := db.QueryRow(db.Rebind("SELECT id FROM users WHERE email = ?"), u.Email).Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	log.Printf("INSERT INTO users (name, email, password) VALUES ('%s', '%s', '%s')", u.Name, u.Email, u.Password)

	r, err := db.Exec(db.Rebind(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`), u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	goat.PrintVarInJson(r)
	return nil
}
