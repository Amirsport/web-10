package provider

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func (p *Provider) SelectQuery(w http.ResponseWriter, r *http.Request) (string, error) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Parameter 'name' is required", http.StatusBadRequest)
		return
	}

	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM public.users WHERE name = $1)`, name).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, fmt.Sprintf("User '%s' not found", name), http.StatusNotFound)
		return
	}

	response := fmt.Sprintf("Hello, %s!", name)
	w.Write([]byte(response))
}

func (p *Provider) CheckQueryExitByMsg(msg string) (bool, error) {
	// Получаем одно сообщение из таблицы hello
	err := p.conn.QueryRow("SELECT name FROM users WHERE name = $1 LIMIT 1", msg).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *Provider) InsertQuery(msg string) error {
	_, err := p.conn.Exec("INSERT INTO users (name) VALUES ($1)", msg)
	if err != nil {
		return err
	}

	return nil
}
