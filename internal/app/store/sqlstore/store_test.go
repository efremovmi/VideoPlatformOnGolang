package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// Функция, которая позволяет инициализировать databaseURL вне зависи  мости от платформы.

// TestMain ...
func TestMain(m *testing.M) {
	//databaseURL = "postgres://postgres:1@localhost:5432/grpc_dev?sslmode=disable"
	databaseURL = os.Getenv("DATABASE_URL") + "?sslmode=" + os.Getenv("DATABASE_SSL_MODE")
	os.Exit(m.Run())
}
