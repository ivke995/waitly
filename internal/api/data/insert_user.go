package data

func InsertUserQuery() string {
	return `INSERT INTO users (username, email) VALUES (?, ?)`
}
