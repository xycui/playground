package command

// NewCmdDel is the delegate
type NewCmdDel func([]string) ICommand

// ICommand is the interface for the command pattern
type ICommand interface {
	Run() error
	PrintHelp()
}

// Executor is the interface for execute pattern
type Executor interface {
	Execute(args ...interface{}) error
}

// CommandBuilder for create command
type CommandBuilder interface {
	Build(args []string) ICommand
}

// BaseCmdBuilder create command with delegate mapping
type BaseCmdBuilder struct {
	CommandBuilder
	delMap map[string]NewCmdDel
}

func NewCmdBuilder(delMap map[string]NewCmdDel) CommandBuilder {
	return &BaseCmdBuilder{
		delMap: delMap,
	}
}

// Build will create sub commands
func (b *BaseCmdBuilder) Build(args []string) ICommand {
	if len(args) != 0 {
		del, ok := b.delMap[args[0]]
		subSlice := args[1:]
		if ok {
			return del(subSlice)
		}
	}

	var commands []ICommand
	for _, v := range b.delMap {
		commands = append(commands, v(args))
	}

	return &FallbackCmd{
		commands: commands,
	}
}
