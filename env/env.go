package env

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

var isProduction bool
var isStage bool
var isDev bool

var devEnv = "dev.env"
var stageEnv = "stage.env"
var productionEnv = "prod.env"

func init() {
	isProduction = os.Getenv("ENV") == "production"
	isStage = os.Getenv("ENV") == "stage"
	isDev = !isProduction && !isStage

	if isProduction {
		log.Info("Loading environment for production")
		err := godotenv.Load(productionEnv)
		if err != nil {
			log.Warnf("Could not load production .env file \"%s\"", productionEnv)
		}
	} else if isStage {
		log.Info("Loading environment for stage")
		err := godotenv.Load(stageEnv)
		if err != nil {
			log.Warnf("Could not load stage .env file \"%s\"", stageEnv)
		}
	} else {
		log.Info("Loading environment for dev")
		err := godotenv.Load(devEnv)
		if err != nil {
			log.Warnf("Could not load dev .env file \"%s\"", devEnv)
		}
	}
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}

func IsDev() bool { return isDev }
