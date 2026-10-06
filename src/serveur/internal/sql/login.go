package sqlite

import (
	"database/sql"

	"github.com/Chefmine8/OpenCroupier/internal"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByUsername(db *sql.DB, name string, password string) (bool, internal.LoginSQL) {
	var result internal.LoginSQL
	err := db.QueryRow("SELECT password, uid FROM user WHERE userName = ?", name).Scan(&result.Password, &result.Uid)
	if err == sql.ErrNoRows {
		return false, internal.LoginSQL{}
	} else if err != nil {
		return false, internal.LoginSQL{}
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	return err == nil, result
}
