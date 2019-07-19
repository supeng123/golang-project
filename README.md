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

#### If condition
~~~
func readFile() {
    const fileName = "abc.txt"
    contents, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }
}

//alternative way
func readFile() {
    const fileName = "abc.txt"
    if contents, err := ioutil.ReadFile(fileName); err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }
    //fmt.Printf("%s\n", contents), can not get contents becuase it's wraped in if condition
}

~~~

#### Switch
~~~
//without break;

func eval(a, b int, op string) int {
    var result int
    switch op {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*"
        result = a * b
    case "/":
        result = a / b
    default:
        panic("unsupported operator")
    }
    return result
}

//without condition after switch
func grade(score int) string {
    g := ''
    switch {
    case score < 60:
        g = "F"
    case score < 80:
        g = "D"
    case score < 100:
        g = "A"
    default:
        g = "unknow"
    }
    return g  
}

~~~

#### For condition
~~~
func covertToBin(n int) string {
    result := ""
    for ; n > 0; n /=2 {
        lsb := n % 2
        result = strconv.Itoa(lsb) + result
    }
    return result
}

//for uses as while
func printFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.println(scanner.Text())
    }
}
~~~

#### Functions
~~~
func div(a,b int)(q,r int) {
    return a/b, a%b
}

//alternative way
func div(a,b int)(q,r int) {
    q = a / b
    r = a % b
    return
}

q, r := div(13, 4)
//only use one output
q, _ := div(15, 8)

func eval(a,b int, op string)(int, error) {
    switch op {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*"
        return a * b, nil
    case "/":
        q, _ = div(a , b)
        return q, nil
    default:
        return 0, fmt.Errorf("unsupported operation: %s", op)
    }
}

func apply(callback func(int, int) int, a, b int) int {
    p := reflect.ValueOf(op).Pointer()
    opName := runtime.FuncForPC(p).Name()
    fmt.Printf("calling function %s with args" + "(%d, %d)", opName, a, b)

    return callback(a, b)
}

//rest arguments
func sum(numbers ...int) int {
    s := 0
    for i := range numbers {
        s += numbers[i]
    }
    return s
}

sum(1,3,4,6)
~~~

#### Pointer
~~~
// golang pass the value ,not the referene , so the value can not be changed
var a int = 2
var pointA *int = &a
*pointA = 3
fmt.Println(a)

func swap(a, b *int) {
    *b , *a = *a, *b
}

a, b := 3, 4
swap(&a, &b)
fmt.Println(a, b)

~~~
