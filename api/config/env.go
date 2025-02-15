package config

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

var Env Environment

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	_, err = env.UnmarshalFromEnviron(&Env)
	if err != nil {
		panic(err)
	}
}
