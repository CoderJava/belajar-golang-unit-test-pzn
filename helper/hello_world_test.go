package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yudi")
	if result != "Hello Yudi" {
		// error
		panic("Result is not 'Hello Yudi'")
	}
}

func TestHelloWorldPablo(t *testing.T) {
	result := HelloWorld("Pablo")
	if result != "Hello Pablo" {
		// error
		panic("Result is not 'Hello Pablo'")
	}
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
