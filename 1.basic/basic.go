package main

import "fmt"
import "math"
import "math/cmplx"
import "strconv"
import "reflect"
import "runtime"

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

func covertToBin(n int) string {
    result := ""
    for ; n > 0; n /=2 {
		// fmt.Println(n)
        lsb := n % 2
        result = strconv.Itoa(lsb) + result
	}
	fmt.Println(result)
	return result
}

func apply(callback func(int, int) int, a, b int) int {
    p := reflect.ValueOf(callback).Pointer()
    opName := runtime.FuncForPC(p).Name()
    fmt.Printf("calling function %s with args" + "(%d, %d)", opName, a, b)

    return callback(a, b)
}

func explainPoint()int {
	var a int = 2
	var pointA *int = &a
	*pointA = 3
	return a
}


func main() {
	fmt.Println("hello world")
	variableInitialValue()
	variableShorter()
	variableZeroValue()
	fmt.Println(aaa)
	fmt.Println(bbb)

	covertToBin(5)
	covertToBin(13)
	fmt.Println(7/3)
	fmt.Println(explainPoint())
}