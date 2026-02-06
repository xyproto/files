package files

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	if err := Run("ls"); err != nil {
		t.Error(err)
	}

	output, err := Shell("echo SHELL")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(output)

	output, err = Bash("echo BASH")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(output)

	//output, err = Fish("echo FISH")
	//if err != nil {
	//    t.Error(err)
	//}
	//fmt.Println(output)
}
