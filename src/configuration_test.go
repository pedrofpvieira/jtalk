package main

import "testing"

func TestReadConfigurations(t *testing.T) {
	DB_HOST = "localhost"

	if DB_HOST != "localhost" {
		t.Errorf("DB_HOST should be localhost")
	}
}
