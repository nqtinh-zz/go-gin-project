package setting

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations
var (
	Version = "1.0"
	Config  appConfig
)

// appConfig structure for configuration
type appConfig struct {
	// PostgresDB config
	DSN                string        `mapstructure:"dsn" validate:"required"`
	DbMaxOpenConns     int           `mapstructure:"db_max_open_conns"`
	DbConnExecTimeout  int           `mapstructure:"db_conn_exec_timeout"`
	DbMaxIdleConnsRate float64       `mapstructure:"db_max_idle_conns_rate"`
	DbConnMaxLifetime  time.Duration `mapstructure:"db_conn_max_lifetime"`

	// Redis config
	RedisClientPort int    `mapstructure:"redis_client_port" validate:"required"`
	RedisClientHost string `mapstructure:"redis_client_host" validate:"required"`

	// Server port. Defaults to 8765
	ServerPort     int    `mapstructure:"server_port"`
	Address        string `mapstructure:"address"`
	AccessTokenTTL int    `mapstructure:"access_token_TTL"`

	// the signing method for JWT. Defaults to "HS256"
	JWTSigningMethod string `mapstructure:"jwt_signing_method"`
	// JWT signing key. required.
	JWTSigningKey string `mapstructure:"jwt_signing_key"`
	// JWT verification key. required.
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
}

// Validate to check appConfig struct
func (config appConfig) Validate() error {
	return validator.New().Struct(&config)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
func LoadConfig(configPaths ...string) error {
	v := viper.New()

	// Postgreslq config
	v.SetDefault("db_max_open_conns", 5)
	v.SetDefault("db_conn_exec_timeout", 3000000)
	v.SetDefault("db_max_idle_conns_rate", 0.5)
	v.SetDefault("db_conn_max_lifetime", time.Minute*30)

	// Server port
	v.SetDefault("server_port", 9393)
	v.SetDefault("address", "http://localhost:9393")
	v.SetDefault("access_token_TTL", 8)
	v.SetDefault("jwt_signing_method", "HS256")

	logrus.Infof("Load configuration in DEV mode")
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	return Config.Validate()
}

// BindEnvs use to bind env vars
func BindEnvs(iface interface{}, vp *viper.Viper, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), vp, append(parts, tv)...)
		default:
			vp.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
