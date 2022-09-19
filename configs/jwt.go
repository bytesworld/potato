package configs

type Jwt struct {
	Secret string `yaml:"secret"`
	JwtTtl int `yaml:"jwt_ttl"`
}
