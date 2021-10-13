package crx

import (
	"archive/zip"
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	crxVersion uint32 = 2
)

var (
	crxMagicNumber = []byte{'C', 'r', '2', '4'}
)

func Zip(folder string) (*zip.Writer, error) {
	z := zip.NewWriter(new(bytes.Buffer))
	defer z.Close()

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		filename, _ := filepath.Rel(folder, path)
		if info.IsDir() {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = filename
		dst, err := z.CreateHeader(header)
		if err != nil {
			return err
		}

		src, err := os.Open(path)
		if err != nil {
			return err
		}

		bytes, err := io.Copy(dst, src)
		if err != nil {
			return err
		}

		fmt.Printf("Included %q %d\n", filename, bytes)

		return nil
	})

	if err != nil {
		return z, err
	}

	return z, nil
}

func LoadKey(file string) (*rsa.PrivateKey, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("key not found")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func Write(w io.Writer, key *rsa.PrivateKey) error {
	if _, err := w.Write(crxMagicNumber); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, crxVersion); err != nil {
		return err
	}

	// if _, err := io.Copy(w, w.); err != nil {
	// 	return err
	// }

	return nil
}
