package config

import "os"

var (
	api_env = os.Getenv("API_ENV")
	mysql_host = os.Getenv("MYSQL_HOST")
	mysql_protocol = os.Getenv("MYSQL_PROTOCOL")
    mysql_user = os.Getenv("MYSQL_USER")
    mysql_password = os.Getenv("MYSQL_PASSWORD")
    mysql_database_name = os.Getenv("MYSQL_DATABASE_NAME")
	redis_host = os.Getenv("REDIS_HOST")
	session_secret = os.Getenv("SESSION_SECRET")
)

func IsDevelopment() bool {
	return api_env == "development"
}

func IsProduction() bool {
	return api_env == "production"
}

func GetFrontURL() string {
	if IsDevelopment() {
		return "http://web.localhost.test"
	}
	return "https://board-web-service-2x4i4vgx5q-an.a.run.app"
}

func GetMySQLHost() string {
	return mysql_host
}

func GetMysqlProtocol() string {
	return mysql_protocol
}

func GetMySQLUserName() string{
	return mysql_user
}

func GetMySQLPassword() string {
	return mysql_password
}

func GetMySQLDatabaseName() string {
	return mysql_database_name
}

func GetRedisHost() string {
	return redis_host
}

func GetSessionSecret() string {
	return session_secret
}