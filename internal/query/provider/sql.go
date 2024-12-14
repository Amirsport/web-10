package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectQuery() (string, error) {
	var msg string

	err := p.conn.QueryRow("SELECT name FROM users WHERE name = $1 LIMIT 1").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return msg, nil
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
	_, err := p.conn.Exec("INSERT INTO users (naem) VALUES ($1)", msg)
	if err != nil {
		return err
	}

	return nil
}
