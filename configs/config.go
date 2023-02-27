package configs

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var cfg *config = &config{}

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "user_todo")
	viper.SetDefault("database.pass", "1122")
	viper.SetDefault("database.database", "api_todo")
}

func Load() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	cfg = &config{}

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return nil
}

func GetDB() DBConfig {
	// fmt.Println("LER API. GetDB()", cfg.API)
	// fmt.Println("LER DB. GetDB()", cfg.DB)
	return cfg.DB
}

func GetServerPort() string {
	// fmt.Println("LER API. GetServerPort()", cfg.API)
	// fmt.Println("LER DB. GetServerPort()", cfg.DB)
	return cfg.API.Port
}
