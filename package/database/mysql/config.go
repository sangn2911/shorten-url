package mysql

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_PORT"`
	Database string `env:"MYSQL_DATABASE"`
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
}
