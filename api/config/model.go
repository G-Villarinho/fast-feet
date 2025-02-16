package config

type Environment struct {
	Postgres Postgres
	Redis    Redis
	API      API
	Session  Session
	SMTP     SMTP
}

type Postgres struct {
	Host        string `env:"POSTGRES_HOST"`
	Port        int    `env:"POSTGRES_PORT"`
	User        string `env:"POSTGRES_USER"`
	Password    string `env:"POSTGRES_PASSWORD"`
	DBName      string `env:"POSTGRES_NAME"`
	DBSSLMode   string `env:"POSTGRES_SSL_MODE"`
	MaxConn     int    `env:"POSTGRES_MAX_CONN"`
	MaxIdle     int    `env:"POSTGRES_MAX_IDLE"`
	MaxLifeTime int    `env:"POSTGRES_MAX_LIFE_TIME"`
	Timeout     int    `env:"POSTGRES_TIMEOUT"`
}

type Redis struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB"`
	Timeout  int    `env:"REDIS_TIMEOUT"`
}

type API struct {
	Port int `env:"API_PORT,default=8080"`
}

type Session struct {
	TokenExp   int    `env:"TOKEN_EXP,default=6"`
	JWTSecret  string `env:"JWT_SECRET"`
	CookieName string `env:"COOKIE_NAME"`
}

type SMTP struct {
	Host     string `env:"SMTP_HOST"`
	Port     int    `env:"SMTP_PORT"`
	User     string `env:"SMTP_USER"`
	Password string `env:"SMTP_PASSWORD"`
}
