package sqlite

const (
	fileStorage     storageType = "file"
	inMemoryStorage storageType = "in-memory"
)

type storageType string

func (s storageType) string() string {
	return string(s)
}

type internalConfig struct {
	storageType storageType
	filename    string
}
