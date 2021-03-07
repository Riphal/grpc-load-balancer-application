package storage

type Pinger interface {
	Ping() error
}
