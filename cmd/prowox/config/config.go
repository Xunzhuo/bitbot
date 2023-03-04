package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

var globalConfig GlobalConfig

type GlobalConfig struct {
	TYPE          string
	LOGIN         string
	ISSUE_KIND    string
	ISSUE_NUMBER  string
	GH_REPOSITORY string
	GH_TOKEN      string
}

func Get() GlobalConfig {
	return globalConfig
}

func MustGet() error {
	var ok bool
	globalConfig.TYPE, ok = os.LookupEnv("TYPE")
	if !ok {
		return errors.New("no TYPE specified")
	}
	globalConfig.GH_REPOSITORY, ok = os.LookupEnv("GH_REPOSITORY")
	if !ok {
		return errors.New("no GH_REPOSITORY specified")
	}
	globalConfig.GH_TOKEN, ok = os.LookupEnv("GH_TOKEN")
	if !ok {
		return errors.New("no GH_TOKEN specified")
	}
	globalConfig.ISSUE_KIND, ok = os.LookupEnv("ISSUE_KIND")
	if !ok && globalConfig.TYPE != "schedule" {
		return errors.New("no ISSUE_KIND specified")
	}
	globalConfig.ISSUE_NUMBER, ok = os.LookupEnv("ISSUE_NUMBER")
	if !ok && globalConfig.TYPE != "schedule" {
		return errors.New("no ISSUE_NUMBER specified")
	}
	globalConfig.LOGIN, ok = os.LookupEnv("LOGIN")
	if !ok && globalConfig.TYPE != "schedule" {
		return errors.New("no LOGIN specified")
	}

	return nil
}

func InitConfig() error {
	if err := MustGet(); err != nil {
		return err
	}

	fmt.Println("config", spew.Sdump(globalConfig))
	return nil
}
