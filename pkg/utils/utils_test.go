package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecGithubCmdHelp(t *testing.T) {
	err := ExecGitHubCmd("help")
	require.NoError(t, err)
}
