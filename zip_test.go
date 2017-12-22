package zip

import (
	"testing"
)

func TestFile(t *testing.T) {
	err := File("arquivos/arquivo1.txt", "arquivos")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDir(t *testing.T) {
	err := Dir("arquivos", "zip")
	if err != nil {
		t.Fatal(err)
	}
}
