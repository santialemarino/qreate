package config

import (
	"os"
	"strconv"
)

func lookupEnvString(key string) (string, bool) {
	v, ok := os.LookupEnv(key)
	return v, ok && v != ""
}

func lookupEnvInt(key string) (int, bool) {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return 0, false
	}
	iv, err := strconv.Atoi(v)
	if err != nil {
		return 0, false
	}
	return iv, true
}

func overrideFromEnv(s *Settings) {
	// App
	if v, ok := lookupEnvString("APP_NAME"); ok {
		s.App.Name = v
	}
	if v, ok := lookupEnvString("APP_ENV"); ok {
		s.App.Env = v
	}
	if v, ok := lookupEnvInt("APP_PORT"); ok {
		s.App.Port = v
	}

	// Environment
	if v, ok := lookupEnvString("ENVIRONMENT"); ok {
		s.Environment = Environment(v)
	}

	// Database
	if v, ok := lookupEnvString("DATABASE_DRIVER"); ok {
		s.Database.Driver = v
	}
	if v, ok := lookupEnvString("DATABASE_HOST"); ok {
		s.Database.Host = v
	}
	if v, ok := lookupEnvInt("DATABASE_PORT"); ok {
		s.Database.Port = v
	}
	if v, ok := lookupEnvString("DATABASE_NAME"); ok {
		s.Database.Name = v
	}
	if v, ok := lookupEnvString("DATABASE_USER"); ok {
		s.Database.User = v
	}
	if v, ok := lookupEnvString("DATABASE_PASSWORD"); ok {
		s.Database.Password = v
	}
	if v, ok := lookupEnvString("DATABASE_SSLMODE"); ok {
		s.Database.SSLMode = v
	}

	// Redis
	if v, ok := lookupEnvString("REDIS_HOST"); ok {
		s.Redis.Host = v
	}
	if v, ok := lookupEnvInt("REDIS_PORT"); ok {
		s.Redis.Port = v
	}
	if v, ok := lookupEnvString("REDIS_PASSWORD"); ok {
		s.Redis.Password = v
	}
	if v, ok := lookupEnvInt("REDIS_DB"); ok {
		s.Redis.DB = v
	}

	// JWT
	if v, ok := lookupEnvString("JWT_SECRET"); ok {
		s.JWT.Secret = v
	}
	if v, ok := lookupEnvInt("JWT_ACCESS_TOKEN_TTL_SECONDS"); ok {
		s.JWT.AccessTokenTTLSeconds = v
	}
	if v, ok := lookupEnvInt("JWT_REFRESH_TOKEN_TTL_SECONDS"); ok {
		s.JWT.RefreshTokenTTLSeconds = v
	}

	// QR
	if v, ok := lookupEnvString("QR_BASE_URL"); ok {
		s.QR.BaseURL = v
	}
	if v, ok := lookupEnvInt("QR_SHORT_CODE_LENGTH"); ok {
		s.QR.ShortCodeLength = v
	}
}
