package dotenv
// cannot have two packages in the same folder..?

import (
	"testing"
	"os"
)

func TestDotenv(t *testing.T) {
	LoadEnv("test.env")
	t1 := os.Getenv("testi")
	t2 := os.Getenv("Test2")

	if t1 != "Hello, World!" {
		t.Error("TEST env should be \"Hello, World!\"")
	}

	if t2 != "Holla Amigos" {
		t.Error(`Test2 env should be "Holla Amigos"`)
	}
}