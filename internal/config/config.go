//Package config store config information in user service
package config

//Config structure represents config structure in user service
type Config struct {
	PostgresURL     string `env:"POSTGRESURL" envDefault:"postgresql://postgres:passwd@localhost:5432/test"`
	UserServicePort string `env:"USERSERVICEPORT" envDefault:"localhost:8087"`
}
