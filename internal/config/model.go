package config

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Timezone string
}

type Config struct {
	Server   Server
	Database Database
}
