package commands

import (
	"strings"

	"github.com/Xunzhuo/bitbot/cmd/prowox/config"
	"github.com/Xunzhuo/bitbot/pkg/utils"
)

func init() {
	registerCommand(retitleCommandName, retitleCommandFunc)
}

var retitleCommandFunc = retitle
var retitleCommandName CommandName = "retitle"

func retitle(args ...string) error {
	return title(strings.Join(args, " "))
}

func title(newtitle string) error {
	return utils.ExecGitHubCmd(
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"edit",
		config.Get().ISSUE_NUMBER,
		"--title",
		newtitle)
}
