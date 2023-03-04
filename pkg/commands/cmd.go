package commands

type CommandName string
type commandFunc func(...string) error

var commandGroups = map[CommandName]commandFunc{}

func registerCommand(name CommandName, command commandFunc) {
	commandGroups[name] = command
}

func GetCommand(name CommandName) (commandFunc, bool) {
	cfunc, found := commandGroups[name]
	return cfunc, found
}
