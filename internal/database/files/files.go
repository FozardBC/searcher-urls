package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	fPath      = filepath.Join(basepath, "db.txt")
)

type FilesDB struct {
	f *os.File
}

//*
//func (df *FilesDB) Write(data []byte) (n int, err error) {

//n, err = df.f.WriteString(string(data))
//if err != nil {
//return nil, err
///	}
//}

func New() *FilesDB {

	file, err := fileNew(fPath)
	if err != nil {
		e := fmt.Errorf("can't make the file:%w", err)
		panic(e)
	}

	return &FilesDB{
		f: file,
	}

}

func fileNew(fPath string) (*os.File, error) {

	if _, err := os.Stat(fPath); errors.Is(err, os.ErrNotExist) {

		f, err := os.Create(fPath)
		if err != nil {
			return nil, err
		}

		return f, nil

	} else {
		f, err := os.OpenFile(fPath, os.O_RDWR, 0666)

		if err != nil {
			return nil, err
		}

		return f, nil
	}
}

func (fDb *FilesDB) Clear() error {

	err := fDb.f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = fDb.f.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}

func (fDb *FilesDB) Read(p []byte) (n int, e error) {

	n, e = fDb.f.Read(p)
	if e != nil {
		return 0, e
	}

	return n, nil
}

func (fDb *FilesDB) Write(p []byte) (n int, e error) {

	n, e = fDb.f.Write(p)
	if e != nil {
		return 0, e
	}

	return n, nil
}
