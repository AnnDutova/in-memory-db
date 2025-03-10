package engine

type (
	Engine interface {
		Set(key, value string) error
		Get(key string) (string, error)
		Delite(key string) error
	}
)
