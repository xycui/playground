package benchmark

import (
	"errors"
	"fmt"

	"github.com/xycui/playground/infra/command"
)

// BenchMarkTriggerStr is used for benchmark playground triggering
const BenchMarkTriggerStr = "benchmark"

var executorMap = map[string]command.Executor{
	testComparePerfTrigger: newCompareBenchmarkExecutor(),
}

type benchMarkCommand struct {
	strArgs []string
}

func NewBenchMarkCommand(strArgs []string) command.ICommand {
	return &benchMarkCommand{
		strArgs: strArgs,
	}
}

func (c *benchMarkCommand) Run() error {
	if len(c.strArgs) == 0 {
		return errors.New("input command invalid")
	}
	executor, ok := executorMap[c.strArgs[0]]
	if !ok {
		return errors.New("executor not found")
	}

	return executor.Execute()
}

func (c *benchMarkCommand) PrintHelp() {
	fmt.Println("Support the following benchmark executor:")
	for k := range executorMap {
		fmt.Printf("\t%v\n", k)
	}
	fmt.Println()
}
