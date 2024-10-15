package server

import (
    "database/sql"
    "net/http"

    "github.com/jmoiron/sqlx"
    _ "github.com/mattn/go-sqlite3"
)

type Server struct {
    db *sqlx.DB
}

func NewServer(db *sqlx.DB) *Server {
    return &Server{db: db}
}

func InitDB(dbFile string) (*sqlx.DB, error) {
    db, err := sqlx.Connect("sqlite3", dbFile)
    if err != nil {
        return nil, err
    }
    
    return db, nil
}

func (s *Server) HandleTasks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        s.handleGetTasks(w, r)
    case http.MethodPost:
        s.handleCreateTask(w, r)
    case http.MethodPut:
        s.handleUpdateTask(w, r)
    case http.MethodDelete:
        s.handleDeleteTask(w, r)
    default:
        http.Error(w, "Invalid request method", http.StatusBadRequest)
    }
}

func (s *Server) handleGetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := s.getTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(tasks)
}

func (s *Server) handleCreateTask(w http.ResponseWriter, r *http.Request) {
    var task Task
    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    err = s.createTask(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
}

func (s *Server) handleUpdateTask(w http.ResponseWriter, r *http.Request) {
    var task Task
    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    err = s.updateTask(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
    var task Task
    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    err = s.deleteTask(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
}

func (s *Server) getTasks() ([]Task, error) {
    tasks := []Task{}
    err := s.db.Select(&tasks, "SELECT * FROM tasks")
    if err != nil {
        return nil, err
    }
    
    return tasks, nil
}

func (s *Server) createTask(task Task) error {
    _, err := s.db.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", task.Title, task.Description)
    if err != nil {
        return err
    }
    
    return nil
}

func (s *Server) updateTask(task Task) error {
    _, err := s.db.Exec("UPDATE tasks SET title = ?, description = ? WHERE id = ?", task.Title, task.Description, task.ID)
    if err != nil {
        return err
    }
    
    return nil
}

func (s *Server) deleteTask(task Task) error {
    _, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", task.ID)
    if err != nil {
        return err
    }
    
    return nil
}

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}
    db, err := sqlx.Connect("sqlite3", dbFile)
    if err != nil {
        return nil, err
    }
    
    return db, nil


func (s *Server) HandleTasks(w http.ResponseWriter, r *http.Request) {

	}
