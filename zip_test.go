package zip

import (
	"testing"
)

func TestFile(t *testing.T) {
	err := File("arquivos/arquivo1.txt")
	if err != nil {
		t.Fatal(err)
	}
}
