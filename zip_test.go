package zip

import (
	"testing"
	"fmt"
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

func Test(t *testing.T) {
	fmt.Println("oi")
}
