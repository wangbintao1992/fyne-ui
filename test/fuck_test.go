package test

import (
	"fmt"
	"fucker/frame/comp"
	"fyne.io/fyne/v2/widget"
	"github.com/robertkrimen/otto"
	"reflect"
	"testing"
)

func TestConfigFile(t *testing.T) {
	p := "/Users/james/Desktop/config.json"

	comp.StartFromConfigFile(p)
}
func TestLabel(t *testing.T) {

	label := widget.NewLabel("1")

	of := reflect.TypeOf(label)

	fmt.Println(of)
}

func TestJs(t *testing.T) {
	vm := otto.New()

	jsCode := `
        function add(a, b){
            return a + b;
        };
    `

	if value, err := vm.Run(jsCode); err == nil {
		fmt.Printf("Result: %s\n", value.String())
	} else {
		fmt.Println(err)
	}

	call, _ := vm.Call("add", nil, 1, 2)
	fmt.Println(call)
}
