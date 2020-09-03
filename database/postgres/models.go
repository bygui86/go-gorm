package postgres

type internalConfig struct {
	host     string
	port     int
	username string
	password string
	params   string
}
