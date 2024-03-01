package url

import (
	"database/sql"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	AddURL(longURL string) (string, error)
	FindLongURL(shortURL string) (string, error)
	IncrementClicks(shortURL string) error
	GetTotalLinks() (int, error)
	GetClicksByID(id string) (int, error)
}

type mysqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *mysqlRepository {
	return &mysqlRepository{db: db}
}

func generateShortURL(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (r *mysqlRepository) AddURL(longURL string) (string, error) {
	shortURL := generateShortURL(6)

	var exists string
	err := r.db.QueryRow("SELECT short_url FROM urls WHERE short_url = ?", shortURL).Scan(&exists)

	for err == nil {
		shortURL = generateShortURL(6)
		err = r.db.QueryRow("SELECT short_url FROM urls WHERE short_url = ?", shortURL).Scan(&exists)
	}

	if err != sql.ErrNoRows {
		return "", err
	}

	statement, err := r.db.Prepare("INSERT INTO urls (short_url, long_url) VALUES (?, ?)")
	if err != nil {
		return "", err
	}
	defer statement.Close()

	_, err = statement.Exec(shortURL, longURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (r *mysqlRepository) FindLongURL(shortURL string) (string, error) {
	var longURL string
	err := r.db.QueryRow("SELECT long_url FROM urls WHERE short_url = ?", shortURL).Scan(&longURL)
	if err != nil {
		return "", err
	}

	return longURL, nil
}

func (r *mysqlRepository) IncrementClicks(shortURL string) error {
	_, err := r.db.Exec("UPDATE urls SET clics = clics + 1 WHERE short_url = ?", shortURL)
	return err
}

func (r *mysqlRepository) GetTotalLinks() (int, error) {
	var total int
	err := r.db.QueryRow("SELECT COUNT(*) FROM urls").Scan(&total)
	return total, err
}

func (r *mysqlRepository) GetClicksByID(id string) (int, error) {
	var clicks int
	err := r.db.QueryRow("SELECT clics FROM urls WHERE id = ?", id).Scan(&clicks)
	return clicks, err
}
