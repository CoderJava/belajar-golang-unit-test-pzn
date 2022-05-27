package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

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

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Yudi Setiawan")
	assert.Equal(t, "Hello Yudi Setiawan", result, "Result must be 'Hello Yudi Setiawan'")
	fmt.Println("TestHelloWorld with Assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dita")
	require.Equal(t, "Hello Dita", result, "Result must be 'Hello Dita'")
	fmt.Println("TestHelloWorldRequire with Require done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		assert.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawan'")
	})
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
