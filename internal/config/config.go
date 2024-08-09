package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"` //если мы парсим ямл файл, то у него будет такое название,по умолчанию будет локал, если пустой
	StoragePath string        `yaml:"storage_path" env-reguired:"./data"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-reguired:"TRUE"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}
type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config { //Must означает, что функция просто упадет и не вызовет ошибку и соотв вобще все упадет
	path := fetchConfigPath() //получили путь для конфиг файла
	if path == "" {
		panic("config file path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) { //stat типа содержимое + название возвращает
		panic("config file does not exist" + path)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

//Запустить конфиг файл можно через указание переменной окружения (CONFIG_PATH = ./что-то там), или можно написать sso -- config=./path...
