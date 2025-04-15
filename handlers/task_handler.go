package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cleisondev/apirestcrud/models"
)

type TaskHandler struct {
	DB *sql.DB
}

//Construtor do TaskHandler
func NewTaskHandler(db *sql.DB) *TaskHandler{
	return &TaskHandler{DB:db}
}

// GET
func (taskHandler *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := taskHandler.fetchTasks()
	if err != nil {
		http.Error(w, "Erro ao buscar tarefas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (handler *TaskHandler) fetchTasks() ([]models.Task, error) {
	rows, err := handler.DB.Query("SELECT id, title, description, status FROM Tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// POST
func (taskHandler *TaskHandler) WriteTask(write http.ResponseWriter, request *http.Request){
	var task models.Task

	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil{
		http.Error(write, "Dados inválidos", http.StatusBadRequest)
	}

	query := "INSERT INTO Tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id"
	err = taskHandler.DB.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		http.Error(write, "Erro ao criar task", http.StatusInternalServerError)
		log.Println("Erro no INSERT:", err)
		return
	}


	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(write).Encode(task)

}


