package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	gopath := build.Default.GOPATH
	path := ""
	args := []string{}
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			if !strings.HasPrefix(v, "-") {
				path = filepath.Join(gopath, "src", v)
				args = append(args, path)
			} else {
				args = append(args, v)
			}
		}
	}
	b, err := exec.Command("ls", args...).CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(b))
}
