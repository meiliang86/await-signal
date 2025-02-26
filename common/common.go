package common

import (
	"fmt"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
	"log"
	"os"
)

const TaskQueue = "replay-2025"

func GetTemporalClient() client.Client {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("could not construct bootstrap logger: %v", err))
	}

	hostPort := getEnvOrDefaultString(logger, "FRONTEND_ADDRESS", client.DefaultHostPort)

	c, err := getTemporalClient(logger, hostPort)
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	return c
}

func getTemporalClient(logger *zap.Logger, hostPort string) (client.Client, error) {
	namespace := getEnvOrDefaultString(logger, "NAMESPACE", client.DefaultNamespace)

	return client.Dial(client.Options{
		Namespace: namespace,
		HostPort:  hostPort,
	})
}

func GetEitherTemporalClient() client.Client {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("could not construct bootstrap logger: %v", err))
	}

	localHostPort := getEnvOrDefaultString(logger, "FRONTEND_ADDRESS", client.DefaultHostPort)
	cloudHostPort := getEnvOrDefaultString(logger, "CLOUD_FRONTEND_ADDRESS", client.DefaultHostPort)

	c, err := getTemporalClient(logger, localHostPort)
	if err == nil {
		return c
	}

	c, err = getTemporalClient(logger, cloudHostPort)
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	return c
}

func getEnvOrDefaultString(logger *zap.Logger, envVarName string, defaultValue string) string {
	value := os.Getenv(envVarName)
	if value == "" {
		logger.Info(fmt.Sprintf("'%s' env variable not set, defaulting to '%s'", envVarName, defaultValue))
		value = defaultValue
	} else {
		logger.Info(fmt.Sprintf("'%s' env variable read as '%s'", envVarName, value))
	}
	return value
}
