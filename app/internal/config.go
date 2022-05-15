package internal

type Config struct {
	IsProd   bool   `env:"IS_PROD" env-default:"false"`
	LogLevel string `env:"LOG_LEVEL" env-default:"info"`

	Listen struct {
		Host string `env:"HOST" env-default:"127.0.0.1"`
		Port string `env:"PORT" env-default:"9990"`
	}
	Postgres struct{}
}
