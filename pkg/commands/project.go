package commands

import (
	"fmt"

	"github.com/Xunzhuo/bitbot/cmd/prowox/config"
	"github.com/Xunzhuo/bitbot/pkg/utils"
)

func init() {
	registerCommand(addProjectCommandName, addProjectCommandFunc)
	registerCommand(removeProjectCommandName, removeProjectCommandFunc)
}

var addProjectCommandFunc = addProject
var addProjectCommandName CommandName = "project"

func addProject(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("run project command failed. Cannot put one more project in one time")
	}
	if err := addPrj(args[0]); err != nil {
		return fmt.Errorf("run project command failed: %s. Make sure project has been created firstly", err.Error())
	}
	return nil
}

var removeProjectCommandFunc = removeProject
var removeProjectCommandName CommandName = "remove-project"

func removeProject(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("run remove-project command failed. Cannot put one more project in one time")
	}
	if err := removePrj(args[0]); err != nil {
		return fmt.Errorf("run remove-project command failed: %s. Make sure project has been created firstly", err.Error())
	}
	return nil
}

func addPrj(project string) error {
	return utils.ExecGitHubCmd(
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"edit",
		config.Get().ISSUE_NUMBER,
		"--add-project",
		project)
}

func removePrj(project string) error {
	return utils.ExecGitHubCmd(
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"edit",
		config.Get().ISSUE_NUMBER,
		"--remove-project",
		project)
}
