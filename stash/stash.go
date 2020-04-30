package stash

import (
	"database/sql"
)

type LogSender interface {
	Send(requestId, log string) error
}

type Stash struct {
	repository *sql.DB // хранилище логов
	table      string  // имя таблицы
}

func NewStash(db *sql.DB) *Stash {
	return &Stash{repository: db}
}

func (s *Stash) Send(requestId, log string) error {
	result, err := s.repository.Exec("INSERT INTO $1 (request_id, log) VALUES ($2, $3)", s.table, requestId, log)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
