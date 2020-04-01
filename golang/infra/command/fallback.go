package command

// FallbackCmd is for catch all 
type FallbackCmd struct{
	ICommand
	commands []ICommand
}

// Run will print help
func (c *FallbackCmd) Run() error {
	c.PrintHelp()
	return nil
}

// PrintHelp will print all command help
func (c *FallbackCmd) PrintHelp() {
	for _, item := range c.commands{
		item.PrintHelp()
	}
}

