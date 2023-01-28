package config

import (
    "log"
    "os"
	"github.com/joho/godotenv"
)

// load .env
func LoadEnv() {
	var envflag=os.Getenv("ENV_FLAG")
	if (envflag=="1"){
		log.Printf("Load from environment")
	}else{
		err := godotenv.Load(".env")
		if (err !=nil){
			log.Printf("load .env error",err)
		}else{
			log.Printf("Load from .env file")
		}
		}
	}