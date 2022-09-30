package config

type Server struct {
	Port string `mapstructure:"port"`
}

type DB struct {
	Username string `mapstructure:"username"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
