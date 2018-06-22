package uweb

import (
	"testing"
)

func TestExistingMethods(t *testing.T) {
	t.Run("should_check_existing_methods", func(t *testing.T) {
		if GET != "GET" {
			t.Errorf("GET is incorrect: %s", GET)
		}

		if POST != "POST" {
			t.Errorf("POST is incorrect: %s", POST)
		}

		if PUT != "PUT" {
			t.Errorf("PUT is incorrect: %s", PUT)
		}

		if HEAD != "HEAD" {
			t.Errorf("HEAD is incorrect: %s", HEAD)
		}

		if DELETE != "DELETE" {
			t.Errorf("DELETE is incorrect: %s", DELETE)
		}

		if TRACE != "TRACE" {
			t.Errorf("TRACE is incorrect: %s", TRACE)
		}

		if CONNECT != "CONNECT" {
			t.Errorf("CONNECT is incorrect: %s", CONNECT)
		}

		if OPTIONS != "OPTIONS" {
			t.Errorf("OPTIONS is incorrect: %s", OPTIONS)
		}
	})
}
