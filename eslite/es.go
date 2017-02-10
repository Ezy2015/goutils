package eslite

type ESLite interface {
	Open(host string, port int) error
	Close()
	Begin() error
	Write(index string, id string,
		typ string, v interface{}) error

	WriteDirect(index string, id string,
		typ string, v interface{}) error

	Commit() error
}
