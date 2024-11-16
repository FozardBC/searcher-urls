package files

import (
	"errors"
	"os"
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

func New() (*FilesDB, error) {

	if _, err := os.Stat("/db.txt"); errors.Is(err, os.ErrNotExist) {

		f, err := os.Create("db.txt")
		if err != nil {
			return nil, err
		}
		defer f.Close()

		fDb := FilesDB{
			f: f,
		}

		return &fDb, nil

	} else {
		file, err := os.Open("/db.txt")
		if err != nil {
			return nil, err
		}

		fDb := FilesDB{
			f: file,
		}
		return &fDb, nil
	}

}
