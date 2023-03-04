package commands

import (
	"os"
	"strings"

	"github.com/xunzhuo/prowox/cmd/prowox/config"
	"github.com/xunzhuo/prowox/pkg/utils"
)

func init() {
	registerCommand(assignCommandName, assignCommandFunc)
	registerCommand(unassignCommandName, unassignCommand)
}

var assignCommandFunc = assign
var assignCommandName CommandName = "assign"

func assign(args ...string) error {
	var assignee []string
	if len(args) == 0 {
		assignee = []string{config.Get().LOGIN}
	} else {
		assignee = formatUserIDs(args)
	}

	return addAssignee(assignee)
}

var unassignCommand = unassign
var unassignCommandName CommandName = "unassign"

func unassign(args ...string) error {
	var assignee []string
	if len(args) == 0 {
		assignee = []string{config.Get().LOGIN}
	} else {
		assignee = formatUserIDs(args)
	}

	return removeAssignee(assignee)
}

func addAssignee(IDs []string) error {
	ids := []string{
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"edit",
		config.Get().ISSUE_NUMBER,
	}
	for _, id := range IDs {
		ids = append(ids, "--add-assignee", id)
	}

	return utils.ExecGitHubCmd(ids...)
}

func removeAssignee(IDs []string) error {
	ids := []string{
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"edit",
		config.Get().ISSUE_NUMBER,
	}
	for _, id := range IDs {
		ids = append(ids, "--remove-assignee", id)
	}

	return utils.ExecGitHubCmd(ids...)
}

func formatUserIDs(names []string) []string {
	formatedIDs := []string{}
	for _, name := range names {
		formatedIDs = append(formatedIDs, strings.TrimPrefix(name, "@"))
	}
	return formatedIDs
}

func isYouSelf() bool {
	return strings.TrimSpace(os.Getenv("AUTHOR")) == config.Get().LOGIN
}
