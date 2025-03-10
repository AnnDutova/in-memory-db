package db

type (
	Storage interface {
		Set(key, value string)
		Get(key string) (string, bool)
		Delite(key string)

		Length() int
	}
)
