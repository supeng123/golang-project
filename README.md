 * [golang-project](#golang-project)
      * [1.Basic Syntax](#1basic-syntax)
         * [Variable type](#variable-type)
         * [Init variable](#init-variable)
         * [Build in variable type](#build-in-variable-type)
         * [Transformation of variable type](#transformation-of-variable-type)
         * [Constants](#constants)
         * [Enums](#enums)
            * [use brackets](#use-brackets)
            * [use iota](#use-iota)
         * [If condition](#if-condition)
            * [normal way to define If](#normal-way-to-define-if)
            * [alternative way to define If](#alternative-way-to-define-if)
         * [Switch](#switch)
            * [without break](#without-break)
            * [without condition after switch](#without-condition-after-switch)
         * [For condition](#for-condition)
            * [normal way use for](#normal-way-use-for)
            * [for uses as while](#for-uses-as-while)
         * [Functions](#functions)
            * [normal way declare function](#normal-way-declare-function)
            * [alternative way declare function](#alternative-way-declare-function)
            * [only use one output](#only-use-one-output)
            * [use function in function](#use-function-in-function)
            * [use rest arguments](#use-rest-arguments)
         * [Pointer](#pointer)
         * [Array](#array)
            * [traverse array](#traverse-array)
         * [Slice](#slice)
            * [change the value in slice](#change-the-value-in-slice)
            * [slice capacity and length](#slice-capacity-and-length)
            * [different ways to create slice](#different-ways-to-create-slice)
            * [copy slice](#copy-slice)
            * [remove slice from the middle of slice](#remove-slice-from-the-middle-of-slice)
            * [remove slice from the front of slice](#remove-slice-from-the-front-of-slice)
            * [remove slice from the end of slice](#remove-slice-from-the-end-of-slice)
         * [Map](#map)
            * [define map](#define-map)
            * [find value in map](#find-value-in-map)
            * [traverse chinese in map](#traverse-chinese-in-map)
         * [Struct](#struct)
            * [define a struct](#define-a-struct)
            * [extend struct TreeNode](#extend-struct-treenode)
            * [second example about struct](#second-example-about-struct)
            * [serialize struct using tag](#serialize-struct-using-tag)
            * [factory design pattern for go](#factory-design-pattern-for-go)
         * [Go Path](#go-path)
         * [Package](#package)
         * [Interface](#interface)
            * [other example for interface](#other-example-for-interface)
            * [two ways to check the interface type](#two-ways-to-check-the-interface-type)
            * [mixin interface](#mixin-interface)
         * [Functional programing](#functional-programing)
            * [normal way to declare closure](#normal-way-to-declare-closure)
            * [Alternative way to declare closure](#alternative-way-to-declare-closure)
         * [Defer](#defer)
            * [normal way to deal with errors using defer](#normal-way-to-deal-with-errors-using-defer)
            * [alternative way to deal with errors using defer](#alternative-way-to-deal-with-errors-using-defer)
         * [Goroutine](#goroutine)
         * [Channel](#channel)
            * [define send channel and receive channel](#define-send-channel-and-receive-channel)
            * [second example about send channel and receive channel](#second-example-about-send-channel-and-receive-channel)
            * [channel put all types of data](#channel-put-all-types-of-data)
            * [traverse channel](#traverse-channel)
            * [practice goroutine with channel](#practice-goroutine-with-channel)
            * [other practice goroutine with channel](#other-practice-goroutine-with-channel)
         * [Select](#select)
         * [Lock](#lock)
         * [Reflect](#reflect)
            * [declare reflect](#declare-reflect)
            * [other example for struct in reflect](#other-example-for-struct-in-reflect)
            * [use reflect changing the int pointer](#use-reflect-changing-the-int-pointer)
            * [prcatice about reflect](#prcatice-about-reflect)
         * [Golang File](#golang-file)
            * [normal way to get file](#normal-way-to-get-file)
            * [open file in buff](#open-file-in-buff)
            * [open small file at once](#open-small-file-at-once)
            * [write file](#write-file)
            * [change the old file](#change-the-old-file)
            * [append one file to the old file](#append-one-file-to-the-old-file)
            * [use ioutil to read and write file](#use-ioutil-to-read-and-write-file)
            * [check if file exist](#check-if-file-exist)
            * [copy file](#copy-file)
# golang-project
practices for golang,

## 1.Basic Syntax

### Variable type
~~~
var firstNumber int
var firstString string
fmt.Printf("%d %q\n", firstNumber, firstString)
~~~

### Init variable
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

### Build in variable type
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

### Transformation of variable type
~~~
func triangle() {
    var a, b int = 3,4
    var c int
    c = int(math.Sqrt(float64(a*a + b*b)))
    fmt.Println(c)
}
~~~

### Constants
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

### Enums
#### use brackets
~~~
func enums() {
    const (
        python = 1
        javascript = 2
        golang = 3
        dart = 4
    )
~~~
#### use iota
~~~
    const (
        python = iota
        javascript
        golang
        dart
    )
    fmt.Println(python, javascript, golang, dart)
}
~~~

### If condition

#### normal way to define If
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
~~~

#### alternative way to define If
~~~
func readFile() {
    const fileName = "abc.txt"
    if contents, err := ioutil.ReadFile(fileName); err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }
    fmt.Printf("%s\n", contents)
    //can not get contents becuase it's wraped in if condition
}
~~~

### Switch

#### without break
~~~
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
~~~
#### without condition after switch
~~~
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

### For condition

#### normal way use for
~~~
func covertToBin(n int) string {
    result := ""
    for ; n > 0; n /=2 {
        lsb := n % 2
        result = strconv.Itoa(lsb) + result
    }
    return result
}
~~~
#### for uses as while
~~~
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

### Functions

#### normal way declare function
~~~

func div(a,b int)(q,r int) {
    return a/b, a%b
}
~~~
#### alternative way declare function
~~~
func div(a,b int)(q,r int) {
    q = a / b
    r = a % b
    return
}

q, r := div(13, 4)
~~~
#### only use one output
~~~
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
~~~
#### use function in function
~~~
func apply(callback func(int, int) int, a, b int) int {
    p := reflect.ValueOf(op).Pointer()
    opName := runtime.FuncForPC(p).Name()
    fmt.Printf("calling function %s with args" + "(%d, %d)", opName, a, b)

    return callback(a, b)
}

~~~
#### use rest arguments
~~~
func sum(numbers ...int) int {
    s := 0
    for i := range numbers {
        s += numbers[i]
    }
    return s
}

sum(1,3,4,6)
~~~


### Pointer
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

### Array
~~~
// how to define array
var arr1 [5]int
arr2 := [3]int{1,2,33}
arr3 := [...]int {1,2,3,5,66}
var grid[4][5] bool

fmt.Println(arr1, arr2, arr3)
fmt.Println(grid)


~~~
#### traverse array
~~~

for index, value := range arr3 {
    fmt.Println(index, value)
}
//array is passed by value, not reference, not change to 100
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i,v)
	}
	arr[0] = 100
}

//if you want to change the value, pass it by reference
func printArray(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i,v)
	}
	arr[0] = 100
}
~~~

### Slice
~~~
//slice is passed by reference,not value, it's a view
func defineSlice() {
    oneSlice := make([]int, 10)
	arr := [...]int {0,2,356,42,4,5,6}
	s := arr[2:6]
	fmt.Println("arr[2:6] =", s)
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[:] =", arr[:])
}

func updateSlice(s []int){
	s[0] = 100
}

func main() {
    newArray := [5]int{0,6,4,7,8}
	newSlice := newArray[1:4]
	fmt.Println(newSlice)
	updateSlice(newSlice)
	fmt.Println(newSlice)
	fmt.Println(newArray)
}

~~~
#### change the value in slice
~~~

//conclusion, pass the slice instead of array to change the value in array
//without * and length of 5 here
func printArray(arr []int) {
	for i, v := range arr {
		fmt.Println(i,v)
	}
	arr[0] = 100
}

newArray := [5]int{0,6,4,7,8}
printArray(newArray[:])

~~~
#### slice capacity and length
~~~

//when the slice length is larger than the father container, it will look up forward to grandfather container if it's in the capacity, 
slice can be expanded
func expandedSlice() {
    newArray := [7]int{0,6,4,7,8,9,10}
    sliceOne := newArray[2:5]
    sliceTwo := sliceOne[3:6]
    fmt.Println(sliceOne)
    fmt.Println(sliceTwo)
}
//output [4 7 8] [9 10]

//newArray's length will not be changed because of it's defined capacity, but the last elements will be changed
by the appending operation through slice
func addElementToSlice() {
    newArray := [7]int{0,6,4,7,8,9,10}
    sliceOne := newArray[1:5]
    s2 = append(sliceOne, 11)
    s3 = append(s2, 12)
    s4 = append(s3, 13)

    fmt.Println(newArray)
}
//output [0 6 4 7 8 11 12]


~~~
#### different ways to create slice
~~~
func printSlice(s []int) {
    fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func defineDifferentSlice() {
     s1 := []int{2,4,6,8}
     s2 := make([]int,16)
     s3 := make([]int, 10, 32)

     printSlice(s1)
     printSlice(s2)
     printSlice(s3)
}

~~~
#### copy slice
~~~
func copySlice(){
    s2 := []int{2,4,6,8}
    s3 := make([]int, 10, 32)
    copy(s2, s3)
    fmt.Println(s2)
}

~~~
#### remove slice from the middle of slice
~~~

func removeSlice() {
	s2 := []int{2,4,6,8,9,9}
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
}

~~~
#### remove slice from the front of slice
~~~

func removeSliceFromFront() {
	s2 := []int{2,4,6,8,9,9}
	front = s2[0]
	s2 = s2[1:]
}

~~~
#### remove slice from the end of slice
~~~
func removeSliceFromEnd() {
	s2 := []int{2,4,6,8,9,9}
	tail = s2[len(s2) -1]
	s2 = s2[:len(s2)-1]
}
~~~

### Map

#### define map
~~~
func defineMap() {
	m := map[string]string {
		"name": "course",
		"site": "golang"
	}

	m2 := make(map[string]int) //m2== empty
	var m3 = make(map[string]int) //m3 == nil

}

~~~
#### find value in map
~~~
func iteratorMap() {
	m := map[string]string {
		"name": "course",
		"site": "golang"
	}

	if getName, ok := m["name"]; ok {
		fmt.Println(getName)
	} else {
		fmt.Println("key does not exist")
	}

	delete(m, "site")
	if site, ok := m["site"]; ok {
		fmt.Println("site has been deleted")
	} else {
		fmt.Println("site has not been deleted")
	}

	for key, value range m {
		fmt.Println(k,v)
	}
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxlength := 0

	for index, char := range []byte(s) {
		fmt.Println(index,char)
		lastI, ok := lastOccured[char]
		
		if ok && lastI >= start {
			fmt.Println("lastI",lastI)
			start = lastI + 1
		}
		if index - start + 1 > maxlength {
			
			maxlength = index - start + 1
			fmt.Println("maxlength:",maxlength)
		}
		lastOccured[char] = index
	}

	return maxlength
}

~~~
#### traverse chinese in map
~~~
func getChineseLength(s string) {
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)\n", i, ch)
	}
}
~~~

### Struct

#### define a struct
~~~
type treeNode struct {
	value int
	left, right *treeNode
}

//define the methods of struct, pass the value
func (node treeNode) print() {
	fmt.Println(node.value)
}

//define the methods of struct, pass the reference
func (node *treeNode) setValue(value int) {
	//need to pass the reference to change the value
	node.value = value
}

func createNode(value int) *treeNode {
	//return the address of the local area, save it on stack
	return &treeNode{value: value}
}

func defineTreeNodes() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{value: 4}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode {
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	fmt.Println(root)
	root.print()

	root.right.left.setValue(4)
	root.right.left.print()

	pRoot := &root
	pRoot.setValue(300)
	pRoot.print()
}


~~~
#### extend struct TreeNode
~~~
type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) posterOrder(){
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.posterOrder()
	right := myTreeNode{myNode.node.Right}
	right.posterOrder()
	myNode.node.Print()
}

~~~
#### second example about struct
~~~
type Student struct {
    Name string
    Age int
    Score int
}

func (s *Student) ShowInfo() {
    fmt.Println(s.Name)
}

func (s *Student) SetScore(score int) {
    s.Score = score
}

type Graduate struct {
    //only need to embed the class you want to extend
    Student
}

type UnderGraduate struct {
    Student
    Name: string
}

func (* UnderGraduate) SetScore(score int){
    s.Score = score+ 5
}

var newGraduate = &Graduate{}
newGraduate.Student.Name = 'slogan'
newGraduate.Student.Age = 11
newGraduate.Student.SetScore(44)

//or it could be simplified, becuase newGraduate does not have Name,Age,SetScore, so it will get all the property from its parent
newGraduate.Name = 'slogan'
newGraduate.Age = 11
newGraduate.SetScore(44)

var newUnderGraduate = &UnderGraduate{}
newUnderGraduate.Name = 'sunminjuan'
newUnderGraduate.Age = 11
newUnderGraduate.SetScore(44)

// will not get 'sunminjuan' because it will get the closest Name in struct Student, the outcome will be empty
newUnderGraduate.ShowInfo()


~~~
#### serialize struct using tag
~~~
import 'encoding/json'

type Monster struct {
    Name string `json: "name"`
    Age string `json: "age"`
    Skill string `json: "skill"`
}

jsonString, err = json.Marshal(Monster)
if err != nil {
    fmt.Println(string(jsonString))

    fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", calculator.getsum()))
}

~~~
#### factory design pattern for go
~~~
type student struct {
    Name: string,
    score: float64,
}

func NewStudent (n string, s float64) *student {
    return &student{
        Name: n,
        score: s,
    }
}
func (s *student) GetScore() float64 {
    return s.score
}
~~~

### Go Path
~~~
use "~/.bash_profile" to check the GOPATH
go get github.com/gpmgo/gopm
~~~

### Package
~~~
public if the first letter is capital
~~~

### Interface
~~~
//define interface and it's function
type Retriever interface {
	Get(urr string)string
}

func download(r Retriever) string {
	return r.Get("http://www.google.com")
}

//define another class to implement this interface
type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}

//check the class
var r2 Retriever
r2 = real.Retriever{
    UserAgent: "Chrome",
    TimeOut: time.Minute,
}
fmt.Printf("%T %v\n", r2, r2)
~~~

#### other example for interface
~~~
type Usb interface {
    Start()
    Stop()
}

type Phone struct {

}

func (p *Phone) Start() {
    fmt.Println("Phone is starting to work")
}

func (p *Phone) Stop() {
    fmt.Println("Phone is Stopping to work")
}

type Computer struct {

}

func (c *Computer) Working(usb Usb) {
    usb.Start()
    usb.Stop()
}

var computer = Computer{}
var phone = &Phone{}
computer.Working(phone)
~~~

#### two ways to check the interface type
~~~

**r.(type)
**r.(int|class|string), r needs to be interface
func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("conent",v.Content)
	case *real.Retriever:
		fmt.Println("useragent", v.UserAgent)
	}
}
//check the type
if mockRetriever, ok:= r.(mock.Retriever); ok {
    fmt.Println(mockRetriever.Content)
}

~~~

#### mixin interface
~~~
type Retriever interface {
	Get(urr string)string
}
type Poster interface {
	Post(url string, 
		form map[string]string) string
}

**should not have repeat Post or Get in previous interfaces
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) {
	s.Get("http://sina.com")
	s.Post("sup", map[string]string{
		"name": "slogan",
		"course": "golang",
	})
}

~~~

### Functional programing

#### normal way to declare closure
~~~
func adder () func(int) int {
	sum := 0;
	return func(v int) int {
		sum += v
		return sum
	}
}

~~~
#### Alternative way to declare closure
~~~
type iAdd func(int) (int, iAdd)

func adderTwo(base int) iAdd{
	return func(v int) (int, iAdd) {
		return base + v, adderTwo(base + v)
	}
}

type Node struct {
	Value int
	Left, Right *Node
}

func (node *Node) TraverseFunc(f func(*Node)) {
    if node == nil {
        return
    }
    node.Left.TraverseFunc(f)
    f(node)
    node.Right.TraverseFunc(f)
})

func (node *Node) Traverse() {
    node.TraverseFunc(func(n * Node){
        n.Print()
    })
    fmt.Println()
}

func main() {
	a := adder()

	for i := 0; i < 10 ; i++ {
		fmt.Println(a(i))
	}
}
~~~

### Defer
#### normal way to deal with errors using defer
~~~
//defer put data in a stack, use defer will not interrupted by panic and other errors
func tryDefer() {
    defer fmt.Println(1)
    defer fmt.Println(2)
    fmt.Println(2)
}

func writeFile(filename string){
    file, err := os.Create(filename)
    if err != nil{
        <!-- panic(err) -->
        <!-- err = errors.New("this is customized error") -->
        <! -- fmt.Println('file exist', err.Error()) -->
        if pathError, ok := err.(*os.PathError); !ok {
            panic(err)
        } else {
            fmt.Println(pathError.Op, pathError.Path, pathError.Err)
        }
        
        return
    }
    defer file.Close()
    writer := bufio.NewWriter(file)
    defer writer.Flush()
}
~~~
#### alternative way to deal with errors using defer
~~~

if  err != nil {
        defer func() {
            r := recover()
            if err, ok := r.(error); ok {
                fmt.Println("Error occured:", err)
            } else {
                panic(r)
            }
        }()
    }
~~~

### Goroutine
~~~
spontaneously invoking the anonymous function
func goRoutine() {
    for i := 0; i < 10 ; i ++ {
        go func(i int) {
            for {
                fmt.Println("Hello from"+ "%d\n", i)
                // handover
                runtime.Gosched()
            }
        }(i)
    }
    time.Sleep(time.Millisecond)
}

//main go routine have multiple small routine, it's light weight threading
import "runtime"
func useMultipleCPU() {
    cupNumber := runtime.NumCPU()
    runtime.GOMAXPROCS(number - 1)
}

// when the main string ended? how the multiple sub threading using the same source
~~~

### Channel

#### define send channel and receive channel
~~~
ch <- v    // channel reieve value 
v := <-ch // channel send value

chan<- //use to send
<-chan //use to recieve

//<-chan used to recieve
// chan<- used to send data
func createSendChannel(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println("work %d recieve channel %c\n", id, <-c)
		}
	}()
	return c
}

func sendChan() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createSendChannel(i)
	}

	for i := 0; i < 10; i++ {
		//channel send 'a'+i
		channels[i] <-'a' + i
	}
}

func buffChannel() {
	c := make(chan int, 3)
	go work(c)
	c <- 1
	c <- 2
	c <- 3
	close(c)
}
~~~
#### second example about send channel and receive channel
~~~
func channelDefined() {
    var intChan chan int
    intChan = make(chan int, 3)
    //channel write int 10
    intChan <- 10
    intChan <- 20

    //get the int from the channel
    var num2 int
    num2 = <-intChan
    fmt.Println(num2)// 10
    num3 := <-intChan // 20

}

mapChan := make(chan map[string]string, 10)
m1 := make(map[string]string, 20)
m1["city1"] = "beijijng"
m1["city2"] = "chongqin"

mapchan <- m1

~~~
#### channel put all types of data
~~~
var allChan chan interface{}
allChan = make(chan interface{}, 10)

cat1 := Cat{Name: "tom", Age: 18}
allChan <- cat1;
allChan <- 10
allChan <- "Jack"

cat11 := <-allChan
//type interface is empty with no methods
fmt.Println(cat11.Name) //error
//use type assertion
a := newCat.(Cat)
fmt.Println(a.Name)

//close the channel, can't write any thing,but it can be read
Close(allChan)

~~~
#### traverse channel
~~~
intChan2 := make(chan int, 10)
for i := 0; i < 100; i++ {
    intChan2 <- i*2
}
// need to close channel otherwise it will be deadlocked
close(intChan2)
for v := range intChan2 {
    fmt.Println("v=", v)
}

~~~
#### practice goroutine with channel
~~~
intChan = make(chan int, 50)
exitChan = make(chan int, 1)

go writeData()
go readData()

for {
    _, ok := <-exitChan
    if !ok {
        break
    }
}

func writeData(intChan chan int) {
    for i := 0; i < 50; i++ {
        intChan <- i
    }
    close(intChan)
}

func readChan(intChan chan int, exitChan chan bool){
    for {
        v, ok := <-intChan
        if !ok {
            break
        }
        fmt.Println("read Data =%V", v)
    }
    exitChan <- true
    close(exitChan)
}

~~~
#### other practice goroutine with channel
~~~
func putNum(intChan chan int){
    for i := 0; i < 8000; i++ {
        intChan <- i
    }
    close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
    var flag bool
    for {
        num, ok := <-intChan
        if !ok {
            break
        }

        flag = true
        //check if it's a prime
        for i := 2; i < num; i++ {
            if num%i == 0 {
                flag = false
                break
            }
        }

        if flag {
            primeChan <- num
        }
    }

    exitChan <- true
}

func main() {
    intChan := make(chan int, 1000)
    primeChan := make(chan int, 2000)
    //flag chan
    exitChan := make(chan bool, 4)

    go putNum(intChan)

    for i := 0; i < 4; i++ {
        go primeNum(intChan, primeChan, exitChan)
    }

    go func() {
        for i := 0; i < 4; i++ {
                <-exitChan
            }

        close(primeChan)
    }()

    for {
        res, ok := <- primeChan
        if !ok {
            break
        }
        fmt.Println("primeNum:%v\n", res)
    }
}
    


~~~
### Select
~~~
func generator() chan int{
    out := make(chan int)

    go func() {
        for {
            time.Sleep(time.Duration(rand.Intn(5000)*time.Millisecond))
            out <- i
            i ++
        }
    }()
    return out
}

func useSelect(){
    var c1, c2 = generator(), generator()

    for {
        select {
            case n:= <-c1:
            fmt.Println("received from c1:" n)
            case n:= <-c2:
            fmt.Println("received from c1:" n)
            case <-time.After(880 * time.Millisecond):
            fmt.Println("timeout")
            default:
            fmt.Println("no value received")
            return
        }
    }
    
}

~~~

### Lock
~~~
type atomicInt struct {
    value int
    lock sync.Mutex
}

func (a * atomicInt) increment(){

    func() {
        a.lock.Lock()
        defer a.lock.Unlock()
        a.value ++
    }()
    
}

func (a * atomicInt) get () int {
    a.lock.Lock()
    defer a.lock.Unlock()
    return a.value
}

func main(){
    var a atomicInt
    bb := atomicInt{}
    a.increment()
    go func(){
        a.increment()
    }()
    fmt.Println(a.get())
}
~~~

### Reflect
#### declare reflect
~~~
reflect.TypeOf()
reflect.ValueOf()

func test (b interface{}) {
    //get the type of the argument
    rType := reflect.TypeOf(b)

    //get the value of the argument
    rVal := reflect.ValueOf(b)
    //get the true value if the type is Int
    trueValue := rVal.Int()

    //change the Int argument back
    iVal := rVal.Interface()
    //put the parameter int for int value 10
    num2 := iVal.(int)
}
~~~
#### other example for struct in reflect
~~~
type Student struct {
    name String
    Age int
}

func testStruct (b interface{}) {
    rType := reflect.TypeOf(b)

    rVal := reflect.ValueOf(b)

    iVal := rVal.Interface()

    //get the kind
    kind1 := rVal.Kind()

    //need type assert before output the value
    assertType := TypeJudge(iVal)
    //or
    stu. ok := iVal.(Student)
    if ok {
        fmt.Printf("stu.Name=%v\n", stu.Name)
    }

}

func TypeJudge(items ...interface{}) {
    for i, x := range items {
        switch x.(type) {
            case bool:
                fmt.Print("param %d is a bool , value is %v\n", i, x)
            case float64:
                fmt.Print("param %d is a float64 , value is %v\n", i, x)
            case int, int64:
                fmt.Print("param %d is a int64 , value is %v\n", i, x)
            case nil:
                fmt.Print("param %d is a nil , value is %v\n", i, x)
            case string:
                fmt.Print("param %d is a string , value is %v\n", i, x)
            default:
                fmt.Print("param %d is a unknown , value is %v\n", i, x)
        }
    }
}

~~~
#### use reflect changing the int pointer
~~~
a  := 10
rVal := reflect.Value(&a)
rVal.Elem().setInt(20)

~~~
#### prcatice about reflect
~~~

type Monster struct {
    Name strig `json: "name"`
    Age int `json:"monster_age"`
    Score float32
    Sex string
}

func (s Monster) Print () {
    fmt.Println(s)
}

func (s Monster) GetSum (n1, n2 int) int {
    return n1 + n2
}

func (s Monster) Set(name string, age int,score flat32, sex string) {
    s.Name = name
    s.Age = age
    s.Score = score
    s.Sex = sex
}

func TestStruct (a interface{}) {
    typ := reflect.TypeOf(a)
    val := reflect.ValueOf(a)
    kd := val.Kind()
    if kd != reflect.Struct {
        fmt.Println("expect struct")
        return
    }

    //get the number of properties
    num := val.NumField()

    for i := 0; i < num ; i++ {
        fmt.Println("Field %d: value is %v\n", i, val.Field(i))

        //get the tag
        tagVal := typ.Field(i).Tag.Get("json")
        if tagVal != "" {
            fmt.Println("Field %d: tage is = %v\n", i, tagVal)
        }
    }

    //get the number of methods
    numOfMethod := val.NumMethod()

    val.Method(1).Call(nil)

    //invoke the first method inside the struct
    //the index based on the first letter in order
    var params []reflect.Value
    params = append(params, reflect.ValueOf(10))
    params = append(params, reflect.ValueOf(20))

    res := val.Method(0).Call(params)
    fmt.Println("res=", res[0].Int())

}
~~~

### Golang File
#### normal way to get file
~~~
import "os"

func openFile(distName){
    if file, err := os.Open("d:/text.txt"); err != nil {
        fmt.Println(err)
    }
    fmt.Printf("file=%v", file)

    err = file.Close()
    if err != nil {
        fmt.Println(err)
    }
}

~~~
#### open file in buff
~~~
import "bufio"
import "io"
func openFile(distName){
    if file, err := os.Open("d:/text.txt"); err != nil {
        fmt.Println(err)
    }

    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        str, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }

    }
}

~~~
#### open small file at once
~~~
import "io/ioutil"
func openFile(distName){
    content, err := ioutil.ReadFile("d:/text.txt")
    if err != nil {
        fmt.Println("err")
    }
    fmt.Printf("%v", string(content))
}

~~~
#### write file
~~~
import "os"
import "bufio"

func writeFile(fileDistName){
    file, err := os.OpenFile(fileDistName, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    str := "hello ,slogan\r\n"
    writer := bufio.NewWriter(file)
    for i:=0; i< 5 ; i++ {
        writer.WriteString(str)
    }
    writer.Flush()
}

~~~
#### change the old file
~~~
func writeFile(fileDistName){
    file, err := os.OpenFile(fileDistName, os.O_WRONLY | os.TRUNC, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    str := "hello ,slogan\r\n"
    writer := bufio.NewWriter(file)
    for i:=0; i< 5 ; i++ {
        writer.WriteString(str)
    }
    writer.Flush()
}

~~~
#### append one file to the old file
~~~
func writeFile(fileDistName){
    file, err := os.OpenFile(fileDistName, os.O_RDONLY | os.APPEND, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    str := "hello ,slogan\r\n"
    writer := bufio.NewWriter(file)
    for i:=0; i< 5 ; i++ {
        writer.WriteString(str)
    }
    writer.Flush()
}

~~~
#### use ioutil to read and write file
~~~
import "os"
import "io/ioutil"

func writeFile () {
    file1Path := "d:/abc.txt"
    file2Path := "e:/edf.txt"

    data, err := ioutil.ReadFile(file1path)
    if err != nil {
        fmt.Println(err)
        return
    }
    err := ioutil.WriteFile(file2Path, data, 0666)
    if (err != nil) {
        fmt.Println(err)
    }
}
~~~
#### check if file exist
~~~
func PathExist(path string)(bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err){
        return false, nil
    }
    return false, err
}

~~~
#### copy file
~~~
func CopeFile(dstFilename string, srcFilename string)(written int64, err error) {
    srcFile, err := os.Open(srcFilename)
    if err != nil{
        fmt.Println(err)
    }

    defer srcFile.Close()

    reader := bufio.NewReader(srcFile)
    dstfile, err := os.OpenFile(dstFilename, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer dstFile.Close()

    writer := bufio.NewWriter(dstfile)
    return io.Copy(writer, reader)
}

**transform to Chinese
str = []rune(str)

**os.Args get the command line arguments
**flag package to parse the command line arguments
~~~