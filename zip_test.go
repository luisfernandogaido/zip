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
	err := Dir("arquivos", "")
	if err != nil {
		t.Fatal(err)
	}
}

func Test(t *testing.T) {
	fmt.Println(len("136d740d8906b8202c56cc93235bf44bf6791b921b608241fee9bb8c3410fe07"))
}

