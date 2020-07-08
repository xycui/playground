package db

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/xycui/playground/infra/command"
	"github.com/xycui/playground/sample/db/model"
)

const TestGormTriggerStr = "gorm"

type TestGormArgs struct {
	MysqlConnStr string
}

type TestGormExecutor struct {
	command.Executor

	args *TestGormArgs
}

func newTestGormExecutor(args *TestGormArgs) *TestGormExecutor {
	return &TestGormExecutor{
		args: args,
	}
}

func (e *TestGormExecutor) Execute(args ...interface{}) error {
	db, err := gorm.Open("mysql", e.args.MysqlConnStr)
	if err != nil {
		return err
	}
	defer db.Close()
	db.AutoMigrate(&model.DataItem{})
	dataModel := &model.DataItem{
		Data: "test data",
	}
	if err := db.Create(dataModel).Error; err != nil {
		return err
	}

	jsonStrBytes, err := json.MarshalIndent(dataModel, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonStrBytes))

	return nil
}

func (e *TestGormExecutor) parseParam(args ...interface{}) (*TestGormArgs, error) {
	for _, item := range args {
		if arg, ok := item.(*TestGormArgs); ok {
			return arg, nil
		}
	}

	return nil, errors.New("Args not valid")
}

// TestGormCommand is the test command
type TestGormCommand struct {
	command.ICommand
	args    *TestGormArgs
	flagSet *flag.FlagSet
	strArgs []string
}

// NewTestGormCommand is the constructor
func NewTestGormCommand(strArgs []string) command.ICommand {
	c := &TestGormCommand{
		flagSet: flag.NewFlagSet(TestGormTriggerStr, flag.ContinueOnError),
		args:    &TestGormArgs{},
		strArgs: strArgs,
	}

	c.flagSet.StringVar(&c.args.MysqlConnStr, "mysql-conn", "", "mysql connection string")

	return c
}

// Run for command run
func (c *TestGormCommand) Run() error {
	c.flagSet.Parse(c.strArgs)
	return newTestGormExecutor(c.args).Execute()
}

// PrintHelp to print help
func (c *TestGormCommand) PrintHelp() {
	fmt.Printf("Command '%v' with following args:\n", TestGormTriggerStr)
	c.flagSet.PrintDefaults()
	fmt.Println()
}
