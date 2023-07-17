package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Arch-4ng3l/TaskManage/types"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	connStr := "user=moritz dbname=taskmanage password=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil
	}

	return &Postgres{
		db,
	}

}

func (psql *Postgres) Init() {
	query := `CREATE TABLE IF NOT EXISTS users(
		id serial PRIMARY KEY,
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT,
		created_at TIMESTAMP
	)`
	query2 := `CREATE TABLE IF NOT EXISTS tasks(
		id serial PRIMARY KEY,
		creator TEXT,
		task_name TEXT, 
		task_content TEXT,
		created_at TIMESTAMP
	)`

	_, err := psql.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	_, err = psql.db.Query(query2)

	if err != nil {
		log.Fatal(err)
	}

}

func (psql *Postgres) AddNewAccount(acc *types.Account) error {

	query := `INSERT INTO users (username, email, password, created_at) VALUES($1, $2, $3, $4)`

	_, err := psql.db.Exec(query, acc.Username, acc.Email, acc.Password, acc.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (psql *Postgres) RemoveAccount(acc *types.Account) error {
	query := `DELETE FROM users WHERE username=$1 AND email=$2`
	_, err := psql.db.Exec(query, acc.Username, acc.Email)
	return err
}
func (psql *Postgres) GetAccountByName(name string) (*types.Account, error) {
	query := `SELECT * FROM users WHERE username=$1`
	rows, err := psql.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	acc := &types.Account{}
	id := 0
	for rows.Next() {
		if err := rows.Scan(&id, &acc.Username, &acc.Email, &acc.Password, &acc.CreatedAt); err != nil {
			return nil, err
		}

	}

	return acc, nil
}

func (psql *Postgres) GetAccountByEmail(email string) (*types.Account, error) {
	query := `SELECT * FROM users WHERE email=$1`
	rows, err := psql.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	acc := &types.Account{}
	id := 0
	for rows.Next() {
		if err := rows.Scan(&id, &acc.Username, &acc.Email, &acc.Password, &acc.CreatedAt); err != nil {
			return nil, err
		}

	}

	return acc, nil
}

func (psql *Postgres) AddNewTask(task *types.Task) error {

	query := `INSERT INTO tasks (creator, task_name, task_content, created_at) 
			  SELECT $1, $2, $3, $4
			  WHERE NOT EXISTS(
				  SELECT id FROM tasks WHERE creator = $1 AND task_name = $2
			  )`

	res, err := psql.db.Exec(query, task.Name, task.TaskName, task.TaskContent, task.CreatedAt)
	if i, _ := res.RowsAffected(); i < 1 {
		return fmt.Errorf("Task Already Exists")
	}
	return err

}
func (psql *Postgres) RemoveTask(task *types.Task) error {

	query := `DELETE FROM tasks WHERE creator=$1 AND task_name=$2`

	_, err := psql.db.Exec(query, task.Name, task.TaskName)

	return err
}
func (psql *Postgres) TaskFromUser(username, taskname string) (*types.Task, error) {

	query := `SELECT * FROM tasks WHERE creator=$1 AND task_name=$2`
	rows, err := psql.db.Query(query, username, taskname)
	if err != nil {
		return nil, err
	}
	task := &types.Task{}
	id := 0
	for rows.Next() {
		err := rows.Scan(&id, &task.Name, &task.TaskName, &task.TaskContent, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return task, nil
}

func (psql *Postgres) AllTasksFromUser(username string) ([]*types.Task, error) {
	query := `SELECT * FROM tasks WHERE creator=$1`
	rows, err := psql.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	var tasks []*types.Task
	id := 0
	for rows.Next() {

		task := types.Task{}

		err := rows.Scan(&id, &task.Name, &task.TaskName, &task.TaskContent, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}
