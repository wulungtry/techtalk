package main

import (
	"fmt"
	"github.com/wulungtry/techtalk/infra"
	"github.com/wulungtry/techtalk/persist"
	"os"
)

func main() {
	persist.ViperInit()
	if err := infra.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
