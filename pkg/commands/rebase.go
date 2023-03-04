package commands

import (
	"fmt"

	"github.com/xunzhuo/prowox/cmd/prowox/config"
	"github.com/xunzhuo/prowox/pkg/utils"
)

func init() {
	registerCommand(rebaseCommandName, rebaseCommandFunc)
}

var rebaseCommandFunc = rebase
var rebaseCommandName CommandName = "rebase"

func rebase(args ...string) error {
	return utils.ExecGitHubCmd("api",
		"--method",
		"PUT",
		"-H",
		"Accept: application/vnd.github+json",
		fmt.Sprintf("/repos/%s/pulls/%s/update-branch", config.Get().GH_REPOSITORY, config.Get().ISSUE_NUMBER),
		"-f",
		"update_method=rebase",
	)
}
