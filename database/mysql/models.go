package mysql

type internalConfig struct {
	host     string
	port     int
	dbName   string
	username string
	password string
	params   string
}
