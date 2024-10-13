package repository

import (
	"testing"

	"github.com/KolesnikM8O/distributed-task-system/auth-service/internal/config/config"
)

// TestInitDB с использованием pgxmock для оригинальной структуры Repository
func TestInitDB(t *testing.T) {

	Config := &config.StorageConfig{
		Username: Username,
		Password: Password,
		Host:     Host,
		Port:     Port,
		Database: Database,
	}
	// repo := &Repository{}
	repo := New()
	db, err := repo.InitDB(Config)
	if err != nil {
		t.Errorf("InitDB failed: %v", err)
	}

	// Проверяем, что соединение с базой данных установлено
	if db == nil {
		t.Errorf("Database connection is not established")
	}
}
