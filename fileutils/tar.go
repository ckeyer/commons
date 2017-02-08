package fileutils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type tarContext struct {
	src string
	dst string
	gz  bool
}

func Tar(src, dst string) error {
	return newTarContext(src, dst, false).tar()
}

func Tgz(src, dst string) error {
	return newTarContext(src, dst, true).tar()
}

func UnTar(targetFile, to string) error {
	return newTarContext(to, targetFile, false).untar()
}

func UnTgz(targetFile, to string) error {
	return newTarContext(to, targetFile, true).untar()
}

func newTarContext(src, dst string, gz bool) *tarContext {
	return &tarContext{
		src: src,
		dst: dst,
		gz:  gz,
	}
}

func (tc *tarContext) tar() error {
	f, err := tc.targetFile()
	if err != nil {
		return err
	}
	defer f.Close()

	var fw io.WriteCloser
	if tc.gz {
		fw = gzip.NewWriter(f)
		defer fw.Close()
	} else {
		fw = f
	}

	tw := tar.NewWriter(fw)
	defer tw.Close()

	return tc.writeTar(tw)
}

func (tc *tarContext) targetFile() (*os.File, error) {
	_, err := os.Stat(tc.dst)
	if os.IsNotExist(err) {
		_, err := os.Stat(filepath.Dir(tc.dst))
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(tc.dst), 0755)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err = os.Remove(tc.dst)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(tc.dst, os.O_CREATE|os.O_RDWR, 0644)
}

func (tc *tarContext) writeTar(tw *tar.Writer) error {
	return filepath.Walk(tc.src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path == tc.src {
			return nil
		}

		hd, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		hd.Name = trimDirPrefix(path, tc.src)
		tw.WriteHeader(hd)

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		_, err = io.Copy(tw, file)
		if err != nil {
			return err

		}
		return nil
	})
}

func (tc *tarContext) untar() error {
	f, err := os.Open(tc.dst)
	if err != nil {
		return err
	}
	defer f.Close()

	var tr *tar.Reader
	if tc.gz {
		gr, err := gzip.NewReader(f)
		if err != nil {
			return err
		}
		defer gr.Close()

		tr = tar.NewReader(gr)
	} else {
		tr = tar.NewReader(f)
	}

	for {
		hd, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		localName := filepath.Join(tc.src, hd.Name)
		if hd.FileInfo().IsDir() {
			os.MkdirAll(localName, 0755)
			continue
		}

		fw, err := os.OpenFile(localName, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return err
		}

		_, err = io.CopyN(fw, tr, hd.Size)
		if err != nil {
			return err
		}
	}

	return nil
}

func trimDirPrefix(path, dir string) string {
	return strings.TrimPrefix(path, strings.TrimSuffix(dir, "/")+"/")
}
