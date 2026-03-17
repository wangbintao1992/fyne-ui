package execute

import (
	"github.com/robertkrimen/otto"
	"log"
)

var vm *otto.Otto

func ExecJs(funcName string, param []interface{}) {
	vm.Call(funcName, nil, 1, 2)
}

func GetScript(funcName string) string {

	return ""
}

func init() {
	p := ""
	if _, err := vm.Run(GetScript(p)); err == nil {
		log.Println("loan js func script success")
	} else {
		panic(err)
	}
}
