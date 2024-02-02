package config

type Config struct {
	GRPCServer Server
	HTTPServer Server
}

type Server struct {
	Port int64
}

func NewConfig() *Config {
	return &Config{
		GRPCServer: Server{
			Port: 9090,
		},
		HTTPServer: Server{
			Port: 8080,
		},
	}
}
