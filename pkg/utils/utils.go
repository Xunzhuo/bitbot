package utils

import (
	"os/exec"
	"strings"

	"github.com/Xunzhuo/bitbot/cmd/prowox/config"
	"k8s.io/klog"
)

const GitHubCMD = "gh"

func ExecGitHubCommonCmd(args ...string) error {
	options := append([]string{config.Get().ISSUE_KIND, "-R", config.Get().GH_REPOSITORY}, args...)
	cmd := exec.Command(GitHubCMD, options...)
	cmdOutput, err := cmd.CombinedOutput()
	klog.Info("command: ", "gh ", strings.Join(options, " "), "\n", string(cmdOutput), "\n")
	if err != nil {
		klog.Error(err, "\n")
		return err
	}

	klog.Info(string(cmdOutput))
	return nil
}

func ExecGitHubCmd(args ...string) error {
	cmd := exec.Command(GitHubCMD, args...)
	cmdOutput, err := cmd.CombinedOutput()
	klog.Info("command: ", "gh ", strings.Join(args, " "), "\n", string(cmdOutput), "\n")
	if err != nil {
		klog.Error(err, "\n")
		return err
	}

	klog.Info(string(cmdOutput))
	return nil
}

func ExecGitHubCmdWithOutput(args ...string) (string, error) {
	cmd := exec.Command(GitHubCMD, args...)
	cmdOutput, err := cmd.Output()
	klog.Info("command: ", "gh ", strings.Join(args, " "), "\n", string(cmdOutput), "\n")
	if err != nil {
		klog.Error(err, "\n")
		return "", err
	}

	return string(cmdOutput), nil
}
