package commands

func init() {
	registerCommand(retestCommandName, retestCommandFunc)
}

// TODO
var retestCommandFunc = retest
var retestCommandName CommandName = "retest"

func retest(args ...string) error {

	return nil
}
