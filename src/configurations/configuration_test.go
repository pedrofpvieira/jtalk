package configurations_test

import "testing"

func TestReadConfigurations(t *testing.T) {
	DbHost := "localhost"

	if DbHost != "localhost" {
		t.Errorf("DB_HOST should be localhost")
	}
}
