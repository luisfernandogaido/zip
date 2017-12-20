package zip

import (
	"os"
	"archive/zip"
	"path/filepath"
	"io/ioutil"
)

func File(file string) error {
	//fixme tudo errado.
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	w := zip.NewWriter(in)
	f, err := w.Create(filepath.Base(file))
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}
