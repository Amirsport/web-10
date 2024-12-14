package provider

import (
	"database/sql"
	"errors"
	"net/http"
)

func (p *Provider) SelectQuery(w http.ResponseWriter, r *http.Request) (string, error) {
	name := r.URL.Query().Get("name")
	err := p.conn.QueryRow("SELECT message FROM hello ORDER BY RANDOM() LIMIT 1").Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	name = "Hell0" + name
	return name, nil

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
