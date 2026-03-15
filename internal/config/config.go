package config

type Config struct {
	DBPath   string
	HTTPPort string
	GRPCPort string
	Provider string
}

func Load() Config {
	return Config{
		DBPath:   getEnv("MINIKMS_DB_PATH", "./data/minikms.db"),
		HTTPPort: getEnv("MINIKMS_REST_PORT", "8080"),
		GRPCPort: getEnv("MINIKMS_GRPC_PORT", "9090"),
		Provider: getEnv("MINIKMS_PROVIDER", "local"),
	}
}
