package commands

import (
	"os"

	"github.com/Xunzhuo/bitbot/cmd/prowox/config"
	"github.com/Xunzhuo/bitbot/pkg/utils"
	"github.com/tetratelabs/multierror"
	"k8s.io/klog"
)

func Response(msg string) error {
	if msg != "" {
		klog.Info("send response: \n", msg)
		return comment(msg)
	}

	return nil
}

func Greeting() (err error) {
	if greeting, ok := os.LookupEnv("GREETING"); greeting != "" && ok {
		err = multierror.Append(err, comment(greeting))
	}

	return
}

func Help() (err error) {
	if help, ok := os.LookupEnv("HELP_INFO"); help != "" && ok {
		err = multierror.Append(err, comment(help))
	}

	return
}

func comment(msg string) error {
	return utils.ExecGitHubCmd(
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"comment",
		config.Get().ISSUE_NUMBER,
		"--body",
		msg)
}
