package config

import (
	"github.com/spf13/viper"
)

var configuration Configuration

type Configuration struct {
	IsDevelopment bool
	Server        ServerConfig
	Database      DbConfig
}

type ServerConfig struct {
	Port         int
	JwtSecret    string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	LogPath      string
	LogLevel     string
	UploadDir    string
}

type DbConfig struct {
	Name        string
	Username    string
	Password    string
	Host        string
	Port        string
	LogMode     bool
	MaxPoolSize int
	MaxIdleConn int
}

func setDefaults() {
	viper.SetDefault("isDevelopment", false)
	viper.SetDefault("server.port", 3000)
	viper.SetDefault("server.logPath", "./log")
	viper.SetDefault("server.logLevel", "debug")
	viper.SetDefault("server.secret", "gizli-ve-guvenli-secret")
	viper.SetDefault("server.readTimeout", 5)
	viper.SetDefault("server.writeTimeout", 10)
	viper.SetDefault("server.IdleTimeout", 120)
	viper.SetDefault("server.uploadDir", "./upload")

	viper.SetDefault("database.name", "ozanpay")
	viper.SetDefault("database.username", "ozanpay")
	viper.SetDefault("database.password", "ozanpay2024!!!")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.logmode", false)
	viper.SetDefault("database.maxPoolSize", 10)
	viper.SetDefault("database.maxIdleConn", 1)
}

// os environment'ından okumak için. şimdilik çokda lazım değil..
func bindEnvs() {
	viper.BindEnv("isDevelopment", "IS_DEVELOPMENT")
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("server.logPath", "SERVER_LOG_PATH")
	viper.BindEnv("server.logLevel", "SERVER_LOG_LEVEL")
	viper.BindEnv("server.secret", "SERVER_SECRET")
	viper.BindEnv("server.timeout", "SERVER_TIMEOUT")
	viper.BindEnv("server.readTimeout", "SERVER_READ_TIMEOUT")
	viper.BindEnv("server.writeTimeout", "SERVER_WRITE_TIMEOUT")
	viper.BindEnv("server.IdleTimeout", "SERVER_IDLE_TIMEOUT")
	viper.BindEnv("server.uploadDir", "SERVER_UPLOAD_DIR")

	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.username", "DB_USERNAME")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.logmode", "DB_LOG_MODE")
	viper.BindEnv("database.maxPoolSize", "DB_MAX_POOL_SIZE")
	viper.BindEnv("database.maxIdleConn", "DB_MAX_IDLE_CONN")
}

func Setup() (Configuration, error) {
	setDefaults()
	bindEnvs()
	// Auto read env variables
	viper.AutomaticEnv()

	// Unmarshal config file to struct
	if err := viper.Unmarshal(&configuration); err != nil {
		return configuration, err
	}

	return configuration, nil
}

func Get() Configuration {
	return configuration
}
