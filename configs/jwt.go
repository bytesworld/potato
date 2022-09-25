package configs

type Jwt struct {
	Secret string `yaml:"secret"`
	JwtTtl int64 `yaml:"jwt_ttl"`
}
