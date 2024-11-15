package database

type IDatabase interface {
	Write(data []byte)
	Read() []byte
}
