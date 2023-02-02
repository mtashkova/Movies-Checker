package env

import "github.tools.sap/distribution-store/moovies/pkg/env"

type BasicAuth struct {
	Username string
	Password string
}

func BasicAuthConfiguration() (BasicAuth, error) {
	var config BasicAuth
	err := env.ReadEnvExec(
		env.ReadString("APP_USERNAME", &config.Username),
		env.ReadString("APP_PASSWORD", &config.Password),
	)
	if err != nil {
		return BasicAuth{}, err
	}

	return config, nil
}
