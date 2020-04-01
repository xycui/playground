package container

import (
	"github.com/xycui/playground/infra/command"
	"github.com/xycui/playground/sample/pattern"
)

var delMap = map[string]command.NewCmdDel{
	pattern.TestInheritTriggerStr: pattern.NewTestInheritCommand,
}

// Level0Builder for build first layer command
var Level0Builder = command.NewCmdBuilder(delMap)
