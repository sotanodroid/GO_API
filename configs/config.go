package configs

//Config main config struct to check type of env variable
type Config struct {
	DbSource string `env:"DB_URL,required"`
	Port     string `env:"PORT,required"`
}
