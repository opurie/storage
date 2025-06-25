package config

type Config struct {
	Host string
	Port int
	Username string
	Password string
	Name string
	Charset string
}


func GetConfig() *Config{
	return &Config{
		Host: "127.0.0.1",
		Port: 3306,
		Username: "Opurie",
		Password: "asdasdbfd",
		Name: "myapp",
		Charset: "utf8",
	}
}