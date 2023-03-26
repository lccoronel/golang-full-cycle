package configs

type conf struct {
	DBDriver     string
	DBHost       string
	DBUser       string
	DBPort       string
	DBPassword   string
	DBName       string
	EbServerPort string
	JWTSecret    string
	JWTExpiresIn int
}

var config *conf

func LoadConfig(path string) (*conf, error) {

}
