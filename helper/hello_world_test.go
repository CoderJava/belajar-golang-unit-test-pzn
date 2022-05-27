package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldYudi(t *testing.T) {
	result := HelloWorld("Yudi")
	if result != "Hello Yudi" {
		// error
		// t.Fail()
		t.Error("Result must be 'Hello Yudi'")
	}

	fmt.Println("TestHelloWorldYudi done")
}

func TestHelloWorldPablo(t *testing.T) {
	result := HelloWorld("Pablo")
	if result != "Hello Pablo" {
		// error
		// t.FailNow()
		t.Fatal("Result must be 'Hello Pablo'")
	}

	fmt.Println("TestHelloWorldPablo done")
}

func TestHelloWorldSetiawan(t *testing.T) {
	result := HelloWorld("Setiawan")
	if result != "Hello Setiawan" {
		// error
		// t.Fail()
		t.Error("Result must be 'Hello Setiawan'")
	}

	fmt.Println("TestHelloWorldSetiawan done")
}

// Cara 1 (nama function test-nya tidak ditampilkan)
// cd helper
// go test

// Cara 2 (nama function test-nya ditampilkan)
// cd helper
// go test -v

// Cara 3 (test function tertentu saja)
// Secara default unit test di golang menggunakan prefix di setiap nama function-nya.
// Jadi, kalau kita menggunakan `go test -run TestHelloWorld` maka, akan menjalankan test function `TestHelloWorldPablo` juga.
// cd helper
// go test -run TestHelloWorldPablo

// Cara 4 (menjalankan semua test-nya)
// go test ./...
