package repository

import (
	"testing"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
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

func TestInitDBErrorUsername(t *testing.T) {
	Config := &config.StorageConfig{
		Username: Fake,
		Password: Password,
		Host:     Host,
		Port:     Port,
		Database: Database,
	}
	repo := New()
	db, err := repo.InitDB(Config)
	if err == nil {
		t.Errorf("InitDB should return an error")
	}

	if db != nil {
		t.Errorf("Database connection should not be established")
	}
}

func TestInitDBErrorPasswords(t *testing.T) {
	Config := &config.StorageConfig{
		Username: Username,
		Password: Fake,
		Host:     Host,
		Port:     Port,
		Database: Database,
	}
	repo := New()
	db, err := repo.InitDB(Config)
	if err == nil {
		t.Errorf("InitDB should return an error")
	}

	if db != nil {
		t.Errorf("Database connection should not be established")
	}
}

func TestInitDBErrorHost(t *testing.T) {
	Config := &config.StorageConfig{
		Username: Username,
		Password: Password,
		Host:     Fake,
		Port:     Port,
		Database: Database,
	}
	repo := New()
	db, err := repo.InitDB(Config)
	if err == nil {
		t.Errorf("InitDB should return an error")
	}

	if db != nil {
		t.Errorf("Database connection should not be established")
	}
}

func TestInitDBErrorPort(t *testing.T) {
	Config := &config.StorageConfig{
		Username: Username,
		Password: Password,
		Host:     Host,
		Port:     Fake,
		Database: Database,
	}
	repo := New()
	db, err := repo.InitDB(Config)
	if err == nil {
		t.Errorf("InitDB should return an error")
	}

	if db != nil {
		t.Errorf("Database connection should not be established")
	}
}

func TestInitDBErrorDatabase(t *testing.T) {
	Config := &config.StorageConfig{
		Username: Username,
		Password: Password,
		Host:     Host,
		Port:     Port,
		Database: Fake,
	}
	repo := New()
	db, err := repo.InitDB(Config)
	if err == nil {
		t.Errorf("InitDB should return an error")
	}

	if db != nil {
		t.Errorf("Database connection should not be established")
	}
}
