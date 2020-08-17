package container

import (
	"github.com/xycui/playground/infra/command"
	"github.com/xycui/playground/sample/benchmark"
	"github.com/xycui/playground/sample/db"
	"github.com/xycui/playground/sample/pattern"
)

var delMap = map[string]command.NewCmdDel{
	pattern.TestInheritTriggerStr: pattern.NewTestInheritCommand,
	db.TestGormTriggerStr:         db.NewTestGormCommand,
	benchmark.BenchMarkTriggerStr: benchmark.NewBenchMarkCommand,
}

// Level0Builder for build first layer command
var Level0Builder = command.NewCmdBuilder(delMap)
