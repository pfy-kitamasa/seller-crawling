package crawler

type DBConfig struct {
	Host     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	Port     int    `envconfig:"DB_PORT" default:"5432"`
	Name     string `envconfig:"DB_NAME" default:"seller-finding"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"password"`
}

type Config struct {
	Env string `default:"local"`
	DB  DBConfig
}
