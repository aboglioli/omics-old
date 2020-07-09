package configuration

type Repository interface {
	Get(k string) string
}
