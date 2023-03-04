package commands

import (
	"errors"
	"fmt"

	"github.com/tetratelabs/multierror"
	"github.com/xunzhuo/prowox/cmd/prowox/config"
	"github.com/xunzhuo/prowox/pkg/utils"
)

func init() {
	registerCommand(kindCommandName, kindCommandFunc)
	registerCommand(removeKindCommandName, removeKindCommandFunc)
	registerCommand(areaCommandName, areaCommandFunc)
	registerCommand(removeAreaCommandName, removeAreaCommandFunc)
	registerCommand(addApproveCommandName, addApproveCommandFunc)
	registerCommand(addLGTMCommandName, addLGTMCommandFunc)
	registerCommand(addHoldCommandName, addHoldCommandFunc)
	registerCommand(addPriorityCommandName, addPriorityCommandFunc)
	registerCommand(removePriorityCommandName, removePriorityCommandFunc)
	registerCommand(addHelpCommandName, addHelpCommandFunc)
	registerCommand(addGoodCommandName, addGoodCommandFunc)
}

var kindCommandFunc = addKind
var kindCommandName CommandName = "kind"

func addKind(args ...string) error {
	var errs error
	for _, l := range args {
		if err := label(fmt.Sprintf("kind/%s", l)); err != nil {
			errs = multierror.Append(errs, fmt.Errorf("run kind command failed: %s. Make sure label has been created firstly", err.Error()))
		}
	}
	return errs
}

var removeKindCommandFunc = removeKind
var removeKindCommandName CommandName = "remove-kind"

func removeKind(args ...string) error {
	var errs error
	for _, l := range args {
		if err := unlabel(fmt.Sprintf("kind/%s", l)); err != nil {
			errs = multierror.Append(errs, fmt.Errorf("run remove-kind command failed: %s. Make sure label has been created firstly", err.Error()))
		}
	}
	return errs
}

var areaCommandFunc = addArea
var areaCommandName CommandName = "area"

func addArea(args ...string) error {
	var errs error
	for _, l := range args {
		if err := label(fmt.Sprintf("area/%s", l)); err != nil {
			errs = multierror.Append(errs, fmt.Errorf("run area command failed: %s. Make sure label has been created firstly", err.Error()))
		}
	}
	return errs
}

var removeAreaCommandFunc = removeArea
var removeAreaCommandName CommandName = "remove-area"

func removeArea(args ...string) error {
	var errs error
	for _, l := range args {
		if err := unlabel(fmt.Sprintf("area/%s", l)); err != nil {
			errs = multierror.Append(errs, fmt.Errorf("run remove-area command failed: %s. Make sure label has been created firstly", err.Error()))
		}
	}
	return errs
}

var addApproveCommandFunc = addApprove
var addApproveCommandName CommandName = "approve"

func approveNotifier() string {
	return fmt.Sprintf(`[APPROVALNOTIFIER] This PR is **APPROVED**

This pull-request has been approved by: @%s
	
The full list of commands accepted by this bot can be found [here](https://github.com/Xunzhuo/prowox/blob/main/COMMAND.md).
	
The pull request process is described [here](https://github.com/Xunzhuo/prowox/blob/main/PROCESS.md).

<details>
    <summary>Details</summary>

Needs approval from an approver in each of these files:
+ [OWNERS](https://github.com/%s/blob/main/OWNERS)

Approvers can indicate their approval by writing %s in a comment
Approvers can cancel approval by writing %s in a comment

</details>
	`, config.Get().LOGIN, config.Get().GH_REPOSITORY, "`/approve`", "`/approve cancel`")
}

func addApprove(args ...string) error {
	if config.Get().ISSUE_KIND != "pr" {
		return errors.New("you can only approve PRs")
	}

	if isYouSelf() {
		return errors.New("you cannot approve your own PRs")
	}

	if len(args) != 1 {
		comment(approveNotifier())
		if err := label("approved"); err != nil {
			return fmt.Errorf("run approve command failed: %s. Make sure approved label has been created firstly", err.Error())
		}
	} else {
		if args[0] == "cancel" {
			if err := unlabel("approved"); err != nil {
				return fmt.Errorf("run approve cancel command failed: %s", err.Error())
			}
		}
	}

	return nil
}

var addLGTMCommandFunc = addLGTM
var addLGTMCommandName CommandName = "lgtm"

func lgtmNotifier() string {
	return fmt.Sprintf(`[LGTMNOTIFIER] This PR is **LGTM**

This pull-request has been lgtm by: @%s
	
The full list of commands accepted by this bot can be found [here](https://github.com/Xunzhuo/prowox/blob/main/COMMAND.md).
	
The pull request process is described [here](https://github.com/Xunzhuo/prowox/blob/main/PROCESS.md).

<details>
    <summary>Details</summary>

Needs reviewers from an reviewer in each of these files:
+ [OWNERS](https://github.com/%s/blob/main/OWNERS)

Reviewers can indicate their approval by writing %s in a comment
Reviewers can cancel approval by writing %s in a comment

</details>
	`, config.Get().LOGIN, config.Get().GH_REPOSITORY, "`/lgtm`", "`/lgtm cancel`")
}

func addLGTM(args ...string) error {
	if config.Get().ISSUE_KIND != "pr" {
		return errors.New("you can only lgtm PRs")
	}

	if isYouSelf() {
		return errors.New("you cannot lgtm your own PRs")
	}

	if len(args) != 1 {
		// we usually assign the reviewer who has lgtm the PR.
		addAssignee([]string{config.Get().LOGIN})

		comment(lgtmNotifier())

		if err := label("lgtm"); err != nil {
			return fmt.Errorf("run lgtm command failed: %s. Make sure lgtm label has been created firstly", err.Error())
		}
	} else {
		if args[0] == "cancel" {
			if err := unlabel("lgtm"); err != nil {
				return fmt.Errorf("run lgtm cancel command failed: %s", err.Error())
			}
		}
	}

	return nil
}

var addHoldCommandFunc = addHold
var addHoldCommandName CommandName = "hold"

func addHold(args ...string) error {
	if len(args) != 1 {
		if err := label("do-not-merge"); err != nil {
			return fmt.Errorf("run hold command failed: %s. Make sure do-not-merge label has been created firstly", err.Error())
		}
	} else {
		if args[0] == "cancel" {
			if err := unlabel("do-not-merge"); err != nil {
				return fmt.Errorf("run hold cancel command failed: %s", err.Error())
			}
		}
	}

	return nil
}

var addHelpCommandFunc = addHelp
var addHelpCommandName CommandName = "help-wanted"

func addHelp(args ...string) error {
	if len(args) != 1 {
		if err := label("help wanted"); err != nil {
			return fmt.Errorf("run help-wanted command failed: %s. Make sure help wanted label has been created firstly", err.Error())
		}
	} else {
		if args[0] == "cancel" {
			if err := unlabel("help wanted"); err != nil {
				return fmt.Errorf("run help-wanted cancel command failed: %s", err.Error())
			}
		}
	}

	return nil
}

var addGoodCommandFunc = addGood
var addGoodCommandName CommandName = "good-first-issue"

func addGood(args ...string) error {
	if len(args) != 1 {
		if err := label("good first issue"); err != nil {
			return fmt.Errorf("run good-first-issue command failed: %s. Make sure good first issue label has been created firstly", err.Error())
		}
	} else {
		if args[0] == "cancel" {
			if err := unlabel("good first issue"); err != nil {
				return fmt.Errorf("run good-first-issue cancel command failed: %s", err.Error())
			}
		}
	}

	return nil
}

var addPriorityCommandFunc = addPriority
var addPriorityCommandName CommandName = "priority"

func addPriority(args ...string) error {
	if len(args) == 1 {
		if err := label(fmt.Sprintf("priority/%s", args[0])); err != nil {
			return fmt.Errorf("run priority cancel command failed: %s, make sure priority/%s exist", err.Error(), args[0])
		}
	} else {
		return fmt.Errorf("run priority cancel command failed, you could only set one priority in one time")
	}

	return nil
}

var removePriorityCommandFunc = removePriority
var removePriorityCommandName CommandName = "remove-priority"

func removePriority(args ...string) error {
	if len(args) == 1 {
		if err := unlabel(fmt.Sprintf("priority/%s", args[0])); err != nil {
			return fmt.Errorf("run remove-priority cancel command failed: %s, make sure priority/%s exist", err.Error(), args[0])
		}
	} else {
		return fmt.Errorf("run remove-priority cancel command failed, you could only remove one priority in one time")
	}

	return nil
}

func label(label string) error {
	return utils.ExecGitHubCmd(
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"edit",
		config.Get().ISSUE_NUMBER,
		"--add-label",
		label)
}

func unlabel(label string) error {
	return utils.ExecGitHubCmd(
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"edit",
		config.Get().ISSUE_NUMBER,
		"--remove-label",
		label)
}

func HasLabel(label string) bool {
	pattern := fmt.Sprintf(".labels[] | select(.name == \"%s\")", label)
	output, err := utils.ExecGitHubCmdWithOutput(
		config.Get().ISSUE_KIND,
		"-R",
		config.Get().GH_REPOSITORY,
		"view",
		config.Get().ISSUE_NUMBER,
		"--json",
		"labels",
		"--jq",
		pattern)
	if err != nil {
		return false
	}

	if output != "" {
		return true
	}
	return false
}
