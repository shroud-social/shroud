package postgresql

import "database/sql"

func init() {
	db, err := sql.Open("postgres", "")
}
