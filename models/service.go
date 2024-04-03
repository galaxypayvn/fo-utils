package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"bitbucket.org/finesys/finesys-utility/constants"
	"bitbucket.org/finesys/finesys-utility/libs/serror"
	"bitbucket.org/finesys/finesys-utility/utils/utstring"
)

type (
	Service struct {
		Key       string `json:"key" valid:"required"`
		Name      string `json:"name" valid:"required"`
		Version   string `json:"version" valid:"required,semver"`
		Host      string `json:"host" valid:"required,host"`
		Port      int    `json:"port" valid:"required,port"`
		RestPort  int    `json:"restPort" valid:"-"`
		Namespace string `json:"namespace" valid:"-"`
	}
)

func (ox Service) Validate() (errx serror.SError) {
	if err := validator.New().Struct(ox); err != nil {
		errx = serror.NewFromErrorc(err, "Failed to validate service config")
	}

	return errx
}

func (ox Service) UserAgent() string {
	return fmt.Sprintf("%s@%s", ox.Key, ox.Version)
}

func IsDebug() bool {
	return utstring.Env(constants.AppDebug) == "TRUE"
}

func Environment() string {
	return utstring.Env(constants.AppEnv, constants.EnvProduction)
}
