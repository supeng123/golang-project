package main

import "fmt"
import "math"
import "math/cmplx"

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n",a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, c, s := 3, 5, true, "def"
	fmt.Println(a,b,c,s)
}

var (
    aaa = 3
    bbb = true
    ccc = "rueer"
)

func euler() {
    fmt.Printf("%.3f\n",
        cmplx.Exp(1i*math.Pi)+1)
}

func consts() {
    const (
        a, b = 3,4
        filename = "abc.txt"
    )
    var c int
    c = int(math.Sqrt(a*a + b*b))
    fmt.Println(c)
}

const (
	python = iota
	javascript
	golang
	dart
)


func main() {
	fmt.Println("hello world")
	variableInitialValue()
	variableShorter()
	variableZeroValue()
	fmt.Println(aaa)
	fmt.Println(bbb)
}