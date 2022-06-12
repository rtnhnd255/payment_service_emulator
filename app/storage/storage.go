package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	config *Config
	db     *pgxpool.Pool
}

func NewStorage(cfg *Config) *Storage {
	return &Storage{
		config: cfg,
	}
}

func (s *Storage) Open() error {
	db, err := pgxpool.Connect(context.Background(), s.config.DBURL)
	if err != nil {
		log.Print("Trouble with opening pgx connection")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		log.Println("Trouble pinging pgx conn")
		return err
	}

	s.db = db
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) WriteTransaction(t *Transaction) error {
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		log.Println("Error with db connection")
		return err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"INSERT INTO transactions (userid, usermail, sum, currency, dtcreated, dtlastchanged, status)",
		t.UserID, t.UserEmail, t.Sum, t.Currency, t.DTCreated, t.DTLastChanged, t.Status)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		log.Println("Error while inserting transaction")
		return err
	}
	t.ID = id
	return nil
}

/*
func (s *Storage) ReadTransaction(query string) (t *Transaction, err error) {
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		log.Println("Error with db connection")
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx context.Context, sql string, args ...interface{})

	return &Transaction{}, nil
}*/

func (s *Storage) UpdateTransaction() error {
	return nil
}

func (s *Storage) DeleteTransaction() error {
	return nil
}
