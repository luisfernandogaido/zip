package zip

import (
	"path/filepath"
	"strings"
	"io/ioutil"
	"archive/zip"
	"os"
	"fmt"
)

//Cria um zip com o arquivo file na pasta destino.
// O nome do arquivo criado é o mesmo que o de origem, com a extensão trocada para zip.
func File(arquivoOrigem, pastaDestino string) error {
	base := filepath.Base(arquivoOrigem)
	indice := strings.Index(base, filepath.Ext(base))
	nome := base[:indice]
	out, err := os.Create(filepath.Join(pastaDestino, nome+".zip"))
	if err != nil {
		return err
	}
	w := zip.NewWriter(out)
	f, err := w.Create(base)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(arquivoOrigem)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return w.Close()
}

//Cria um arquivo com o conteúdo da pasta de origem na pasta de destino.
// O nome do arquivo é o mesmo que o da pasta de origem.
func Dir(pastaOrigem, pastaDestino string) error {
	base := filepath.Base(pastaOrigem)
	out, err := os.Create(filepath.Join(pastaDestino, base+".zip"))
	if err != nil {
		return err
	}
	w := zip.NewWriter(out)
	err = filepath.Walk(pastaOrigem, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		return err
	})
	return w.Close()
}
