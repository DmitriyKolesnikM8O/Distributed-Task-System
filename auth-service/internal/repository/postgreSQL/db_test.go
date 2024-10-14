package repository

import (
	"testing"

	"github.com/KolesnikM8O/distributed-task-system/auth-service/internal/config/config"
)

func TestInitDB(t *testing.T) {
	Config := &config.StorageConfig{
		Username: Username,
		Password: Password,
		Host:     Host,
		Port:     Port,
		Database: Database,
	}
	repo := New()
	db, err := repo.InitDB(Config)
	if err != nil {
		t.Errorf("InitDB failed: %v", err)
	}

	if db == nil {
		t.Errorf("Database connection is not established")
	}
}
