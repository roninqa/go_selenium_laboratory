package xxMoreExercises

import (
	"log"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1

	t.Run("1+", func(t *testing.T) {
		t.Parallel()
		if a+2 != 2 {
			t.Fatal("fail!")
		}
	})

	t.Run("2+", func(t *testing.T) {
		t.Parallel()
		if a+2 != 3 {
			t.Fatal("fail!")
		}
	})
}

func TestSum(t *testing.T) {
	a := 1

	t.Run("1+", func(t *testing.T) {
		t.Parallel()
		if a+2 != 2 {
			t.Fatal("fail!")
		}
	})

	t.Run("2+", func(t *testing.T) {
		t.Parallel()
		if a+2 != 3 {
			t.Fatal("fail!")
		}
	})
}

func TestMain(m *testing.M) {
	Setup()
	code := m.Run()
	Teardown()
	os.Exit(code)
}

func Setup() {
	log.Println("In Setup...")
}

func Teardown() {
	log.Println("In Teardown...")
}
