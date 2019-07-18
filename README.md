# golang-project
practices for golang,

## 1.Basic Syntax

#### Variable type
~~~
var firstNumber int
var firstString string

fmt.Printf("%d %q\n", firstNumber, firstString)
~~~

#### Init variable
~~~
var a, b int = 3, 4
var s string = "abc"
fmt.Println(a, b, s)

//use bracket to init variables
var (
    aaa = 3
    bbb = true
    ccc = "rueer"
)

//or use ":=" to init variables without var for prefix but it can only be used in functions
a, b, c, s := 3, 5, true, "def"
fmt.Println(a,b,c,s)

~~~

#### Build in variable type
~~~
bool, int , string, byte, rune

(u)int, (u)int8, (u)int16 ..., uintptr
float32, float64, complex64, complex124

//plural example

func euler() {
    var pluralExample =cmplx.Exp(math.E, 1i*Math.Pi) + 1

    fmt.Println(pluraExample)
}

~~~

#### Transformation of variable type
~~~
func triangle() {
    var a, b int = 3,4
    var c int
    c = int(math.Sqrt(float64(a*a + b*b)))
    fmt.Println(c)
}

~~~

#### Constants
~~~
//do not need to use float64 to tranform the data type
func consts() {
    const (
        a, b = 3,4
        filename = "abc.txt"
    )
    var c int
    c = int(math.Sqrt(a*a + b*b))
    fmt.Println(c)
}

~~~

#### Enums
~~~
func enums() {
    const (
        python = 1
        javascript = 2
        golang = 3
        dart = 4
    )

    or 

    const (
        python = iota
        javascript
        golang
        dart
    )
    fmt.Println(python, javascript, golang, dart)
}

~~~

