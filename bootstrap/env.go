package bootstrap

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func ConvertToInt(stringNumber string) int {

	if stringNumber == "" {
		return 2
	}

	i, err := strconv.Atoi(stringNumber)
	if err != nil {
		log.Fatal(err)
		return 2
	}

	return i
}

func NewEnv() *Env {

	if os.Getenv("IS_DOCKER") != "TRUE" {

		env := Env{}
		viper.SetConfigFile(".env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Can't find the file .env : ", err)
		}

		err = viper.Unmarshal(&env)
		if err != nil {
			log.Fatal("Environment can't be loaded: ", err)
		}

		if env.AppEnv == "development" {
			log.Println("The App is running in development env")
		}

		return &env
	}

	return &Env{
		AppEnv:                 os.Getenv("APP_ENV"),
		ServerAddress:          os.Getenv("SERVER_ADDRESS"),
		ContextTimeout:         ConvertToInt(os.Getenv("CONTEXT_TIMEOUT")),
		DBHost:                 os.Getenv("DB_HOST"),
		DBPort:                 os.Getenv("DB_PORT"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPass:                 os.Getenv("DB_PASS"),
		DBName:                 os.Getenv("DB_NAME"),
		AccessTokenExpiryHour:  ConvertToInt(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR")),
		RefreshTokenExpiryHour: ConvertToInt(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR")),
		AccessTokenSecret:      os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:     os.Getenv("REFRESH_TOKEN_SECRET"),
	}
}
