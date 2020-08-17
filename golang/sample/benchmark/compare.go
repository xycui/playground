package benchmark

import (
	"bytes"
	"fmt"
	"reflect"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/xycui/playground/infra/command"
)

const testComparePerfTrigger = "compare"

type compareBenchmarkExecutor struct {
}

func newCompareBenchmarkExecutor() command.Executor {
	return &compareBenchmarkExecutor{}
}

func (e *compareBenchmarkExecutor) Execute(args ...interface{}) error {
	e1 := newBasicEqualer()
	e2 := newFmtEqualer()
	e3 := newReflectEqualer()
	e4 := newMsgpEqualer()
	equalers := []equaler{e1, e2, e3, e4}
	fmt.Println("- Basic type compare")
	compareNPrint(equalers, 1, 1)
	fmt.Println("- Basic struct compare")
	compareNPrint(equalers, newSimpleStruct(), newSimpleStruct())
	fmt.Println("- Complex struct compare")
	compareNPrint(equalers, newComplexStruct(), newComplexStruct())

	return nil
}

func getExecutionTime(action func(interface{}, interface{}) bool, a, b interface{}, repeat int) time.Duration {
	start := time.Now()
	for i := 0; i < repeat; i++ {
		action(a, b)
	}
	end := time.Now()
	return end.Sub(start)
}

func compareNPrint(eqlrs []equaler, a interface{}, b interface{}) {
	header := fmt.Sprintf("%v\t%v\t%v\t%v\t%v", "EqualerName", "Equal", "TimeCost(100)", "TimeCost(1000)", "TimeCost(10000)")
	fmt.Println(header)
	for _, eqlr := range eqlrs {
		fmt.Printf("%v\t", eqlr.GetName())
		fmt.Printf("%v\t", eqlr.Equal(a, b))
		fmt.Printf("%-13v\t", getExecutionTime(eqlr.Equal, a, b, 100))
		fmt.Printf("%-14v\t", getExecutionTime(eqlr.Equal, a, b, 1000))
		fmt.Printf("%-15v\n", getExecutionTime(eqlr.Equal, a, b, 10000))
	}
	fmt.Println()
}

type basic struct {
	ID    int64
	Title string
	Intro string
}

type unit struct {
	Inner *basic
	Tags  []string
	Code  int64
}

type wrapper struct {
	Basic1    *basic
	Basic2    *basic
	Title     string
	Data      string
	Len       int64
	InnerUnit *unit
}

func newSimpleStruct() *basic {
	return &basic{
		ID:    123456789,
		Title: "Sample Title",
		Intro: "Complexity - Splitting features into microservices allows you to split code into smaller chunks. It harks back to the old unix adage of 'doing one thing well'. There's a tendency with monoliths to allow domains to become tightly coupled with one another, and concerns to become blurred. This leads to riskier, more complex updates, potentially more bugs and more difficult integrations.",
	}
}

func newComplexStruct() *wrapper {
	return &wrapper{
		Basic1: newSimpleStruct(),
		Basic2: newSimpleStruct(),
		Title:  "Wrapper title",
		Data:   "For example, this code always computes a positive elapsed time of approximately 20 milliseconds, even if the wall clock is changed during the operation being timed:",
		Len:    300,
		InnerUnit: &unit{
			Inner: newSimpleStruct(),
			Tags:  []string{"tag1", "tag2", "tag3"},
			Code:  200,
		},
	}
}

type equaler interface {
	GetName() string
	Equal(interface{}, interface{}) bool
}

type basicEqualer struct {
}

func newBasicEqualer() equaler {
	return &basicEqualer{}
}

func (e *basicEqualer) GetName() string {
	return "Basic Equaler"
}

func (e *basicEqualer) Equal(a interface{}, b interface{}) bool {
	return a == b
}

type reflectEqualer struct {
}

func newReflectEqualer() equaler {
	return &reflectEqualer{}
}

func (e *reflectEqualer) GetName() string {
	return "Reflect Equaler"
}
func (e *reflectEqualer) Equal(a interface{}, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

type fmtEqualer struct {
}

func newFmtEqualer() equaler {
	return &fmtEqualer{}
}

func (e *fmtEqualer) GetName() string {
	return "FMT Equaler"
}

func (e *fmtEqualer) Equal(a interface{}, b interface{}) bool {
	return fmt.Sprintf("%+v", a) == fmt.Sprintf("%+v", b)
}

type msgpEqualer struct {
}

func newMsgpEqualer() equaler {
	return &msgpEqualer{}
}

func (e *msgpEqualer) GetName() string {
	return "MessagePack"
}

func (e *msgpEqualer) Equal(a interface{}, b interface{}) bool {
	ba, err := msgpack.Marshal(a)
	if err != nil {
		panic(err)
	}
	bb, err := msgpack.Marshal(b)
	if err != nil {
		panic(err)
	}
	return bytes.Equal(ba, bb)
}
