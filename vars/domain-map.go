package vars

import (
	"fmt"
	"reflect"
)

func AddNameSpace(name string, class interface{}) {
	fmt.Println("AddNameSpace", name, reflect.TypeOf(class))
}
