package models

type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status bool `json:"status"`
} 

const (
	Tablename = "tasks"
	CreateTableSQL  = `CREATE TABLE IF NOT EXISTS ` + Tablename + ` (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status BOOLEAN NOT NULL DEFAULT FALSE
	);`
)