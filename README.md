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

### If condition
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

### Switch
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

### For condition
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

### Functions
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

//iterate array

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

//different ways to create slice
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

//slice copy
func copySlice(){
    s2 := []int{2,4,6,8}
    s3 := make([]int, 10, 32)
    copy(s2, s3)
    fmt.Println(s2)
}

//remove slice from the middle of slice
func removeSlice() {
	s2 := []int{2,4,6,8,9,9}
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
}

//remove slice from the front of slice
func removeSliceFromFront() {
	s2 := []int{2,4,6,8,9,9}
	front = s2[0]
	s2 = s2[1:]
}

//remove slice from the end of slice
func removeSliceFromEnd() {
	s2 := []int{2,4,6,8,9,9}
	tail = s2[len(s2) -1]
	s2 = s2[:len(s2)-1]
}

~~~

### Map
~~~
// how to define map
func defineMap() {
	m := map[string]string {
		"name": "course",
		"site": "golang"
	}

	m2 := make(map[string]int) //m2== empty
	var m3 = make(map[string]int) //m3 == nil

}

// how to find value in map
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

//iterate chinese 
func getChineseLength(s string) {
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)\n", i, ch)
	}
}
~~~

### Struct
~~~
//define a struct
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

entend struct TreeNode
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

### Go Path
~~~
use "~/.bash_profile" to check the GOPATH

go get github.com/gpmgo/gopm

~~~

### Package
~~~
public if the first letter is capital

~~~