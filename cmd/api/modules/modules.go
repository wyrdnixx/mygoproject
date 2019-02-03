package modules

import (
	"fmt"
)

func HelloWorld() {
	fmt.Println("Hello-World")
	fmt.Println("hello-world() Info: ", AppConfig.Info)
}
