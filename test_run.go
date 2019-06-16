// +build ignore

package main

import (
	"fmt"
	"github.com/zetamatta/gowin-versioninfo"
)

func main() {
	v, err := versioninfo.Get(`C:\WINDOWS\system32\notepad.exe`)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("product=%d.%d.%d.%d\n",
		v.File[0], v.File[1], v.File[2], v.File[3])

	fmt.Printf("   file=%d.%d.%d.%d\n",
		v.Product[0], v.Product[1], v.Product[2], v.Product[3])
}
