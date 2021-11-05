package config

type Config struct {
	ApplicationKey string
	DatabaseHost string
	DatabaseUser string
	DatabaseName string
	DatabasePort int
	DatabasePassword string
	MailFrom string
	MailerAccess string

}



var globalConfig Config

func GetConfig()  Config {
	return globalConfig
}

func SetConfig(config Config)  {
	globalConfig = config
}