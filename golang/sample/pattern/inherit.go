package pattern

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"

	"github.com/xycui/playground/infra/command"
)

// TestInheritTriggerStr is the entry command for inherit case
const TestInheritTriggerStr = "inherit"

var testInheritExecutor command.Executor

func init() {
	testInheritExecutor = newTestInheritExecutor()
}

// ISample is the test interface
type ISample interface {
	PublicFunc() error
	privateFunc() error
}

// SampleA is the implement of ISample
type SampleA struct {
	ISample
	PublicField  string
	privateField string
}

// NewSampleA is the constructor
func NewSampleA(pub, pri string) *SampleA {
	return &SampleA{
		PublicField:  pub,
		privateField: pri,
	}
}

// PublicFunc is the public func of Sample A
func (s *SampleA) PublicFunc() error {
	fmt.Println("In Public Func")
	d, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("Marshal fail")
	} else {
		fmt.Println(string(d))
	}
	fmt.Println("----------------")
	return nil
}

func (s *SampleA) privateFunc() error {
	fmt.Println("In private Func")
	return nil
}

// SampleB is another implement of ISample
type SampleB struct {
	*SampleA
	newPrivateField string
}

// NewSampleB is the constructor
func NewSampleB(pub, pri string, override bool) *SampleB {
	s := &SampleB{
		newPrivateField: pri,
	}
	if override {
		s.SampleA = NewSampleA(pub, pri)
	}
	return s
}

// SampleOverride is the override test implement of ISample
type SampleOverride struct {
	*SampleA
}

// NewSampleOverride is the constructor
func NewSampleOverride(pub, pri string) *SampleOverride {
	return &SampleOverride{
		SampleA: NewSampleA(pub, pri),
	}
}

// PublicFunc is the override method
func (s *SampleOverride) PublicFunc() error {
	if s.SampleA != nil {
		fmt.Println("Call parent:")
		s.SampleA.PublicFunc()
		fmt.Println("-------------------")
	}
	fmt.Println("Call override:")
	fmt.Println("-------------------")
	return nil
}

// TestInheritArgs is for parsing string args
type TestInheritArgs struct {
	Override bool
	PubStr   string
	PriStr   string
}

// TestInheritExecutor will be used as singlaton
type TestInheritExecutor struct {
	command.Executor
}

func newTestInheritExecutor() *TestInheritExecutor {
	return &TestInheritExecutor{}
}

// Execute for test
func (e *TestInheritExecutor) Execute(args ...interface{}) error {
	arg, err := e.parseParam(args...)
	if err != nil {
		return err
	}

	var sampleA ISample = NewSampleA("Base Public string", "Base private string")
	var sampleB ISample = NewSampleB(arg.PubStr, arg.PriStr, arg.Override)
	var sampleOverride ISample = NewSampleOverride(arg.PubStr, arg.PriStr)

	sampleA.PublicFunc()
	sampleA.privateFunc()
	sampleB.PublicFunc()
	sampleB.privateFunc()
	sampleOverride.PublicFunc()
	sampleOverride.privateFunc()

	return nil
}

func (e *TestInheritExecutor) parseParam(args ...interface{}) (*TestInheritArgs, error) {
	for _, item := range args {
		if arg, ok := item.(*TestInheritArgs); ok {
			return arg, nil
		}
	}

	return nil, errors.New("Args not valid")
}

// TestInheritCommand is the test command
type TestInheritCommand struct {
	command.ICommand
	args    *TestInheritArgs
	flagSet *flag.FlagSet
	strArgs []string
}

// NewTestInheritCommand is the constructor
func NewTestInheritCommand(strArgs []string) command.ICommand {
	c := &TestInheritCommand{
		flagSet: flag.NewFlagSet(TestInheritTriggerStr, flag.ContinueOnError),
		args:    &TestInheritArgs{},
		strArgs: strArgs,
	}

	c.flagSet.BoolVar(&c.args.Override, "o", false, "override parent value")
	c.flagSet.StringVar(&c.args.PriStr, "pri", "private string", "private string value")
	c.flagSet.StringVar(&c.args.PubStr, "pub", "public string", "public string value")

	return c
}

// Run for command run
func (c *TestInheritCommand) Run() error {
	c.flagSet.Parse(c.strArgs)
	return testInheritExecutor.Execute(c.args)
}

// PrintHelp to print help
func (c *TestInheritCommand) PrintHelp() {
	fmt.Printf("Command '%v' with following args:\n", TestInheritTriggerStr)
	c.flagSet.PrintDefaults()
	fmt.Println()
}
