package execute

import (
	"github.com/robertkrimen/otto"
)

var vm *otto.Otto

func ExecJs(funcName string, param []interface{}) {
	vm.Call(funcName, nil, 1, 2)
}

func GetScript(funcName string) string {

	return ""
}
