package connection

import "github.tools.sap/distribution-store/moovies/pkg/env"

type Postgres struct {
	Host     string
	Port     int
	Password string
	User     string
	DBName   string
}

func LoadConfig() (Postgres, error) {
	var config Postgres
	err := env.ReadEnvExec(
		env.ReadString("DB_HOST", &config.Host),
		env.ReadInt("DB_PORT", &config.Port),
		env.ReadString("DB_PASS", &config.Password),
		env.ReadString("DB_USER", &config.User),
		env.ReadString("DB_NAME", &config.DBName),
	)
	if err != nil {
		return Postgres{}, err
	}

	return config, nil
}
