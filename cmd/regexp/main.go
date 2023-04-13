package main

import (
	"fmt"

	"github.com/flytam/filenamify"
)

func main() {
	n := `<foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar><foo/bar>`
	output, err := filenamify.Filenamify(n, filenamify.Options{})
	fmt.Println(output, err) // => foo!bar,nil
	if output != n {
		fmt.Println("名称不合法1")
	}

	n = `大多数防守打法`
	output, err = filenamify.Filenamify(n, filenamify.Options{
		MaxLength: 100,
	})
	fmt.Println(output, err) // => foo!bar,nil
	if output != n {
		fmt.Println("名称不合法2")
	}
	//---
	n = `foo:"bar"`
	output, err = filenamify.Filenamify(n, filenamify.Options{
		Replacement: "🐴",
	})
	fmt.Println(output, err) // => foo🐴bar,nil
	if output != n {
		fmt.Println("名称不合法3")
	}
}
