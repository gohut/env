package env

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load(".env")
	if err != nil {
		log.Fatalf("Fail to parse '.env': %v", err)
	}
	LoadBase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	fmt.Println("Current Mode", RunMode)
}

func Load(key, section string) string {
	return Cfg.Section(section).Key(key).String()
}

func Msg() {
	fmt.Println("env.Msg")
}
