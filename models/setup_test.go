package models

import "testing"

func TestConnectDB(t *testing.T) {
	err := ConnectDB()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := DB.Exec("SELECT 1").Error; err != nil {
		t.Fatalf("Database connection failed: %v", err)
	}
}
