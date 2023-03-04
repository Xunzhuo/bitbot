package commands

import (
	"fmt"

	"github.com/xunzhuo/prowox/cmd/prowox/config"
	"github.com/xunzhuo/prowox/pkg/utils"
)

func init() {
	registerCommand(addMilestoneCommandName, addMilestoneCommandFunc)
	registerCommand(removeMilestoneCommandName, removeMilestoneCommandFunc)
}

var addMilestoneCommandFunc = addMilestone
var addMilestoneCommandName CommandName = "milestone"

func addMilestone(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("run milestone command failed. Cannot put one more milestone in one time")
	}
	if err := milestone(args[0]); err != nil {
		return fmt.Errorf("run milestone command failed: %s. Make sure milestone has been created firstly", err.Error())
	}
	return nil
}

var removeMilestoneCommandFunc = removeMilestone
var removeMilestoneCommandName CommandName = "remove-milestone"

func removeMilestone(args ...string) error {
	if err := milestone(""); err != nil {
		return fmt.Errorf("run remove-milestone command failed: %s", err.Error())
	}
	return nil
}

func milestone(milestone string) error {
	return utils.ExecGitHubCmd(
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"edit",
		config.Get().ISSUE_NUMBER,
		"--milestone",
		milestone)
}
