package queries

const (
	QuerySelectUserByEmail = `SELECT * FROM users WHERE email=$1;`
	QueryInsertUser        = `INSERT INTO users (email, password) VALUES ($1, $2);`
)
