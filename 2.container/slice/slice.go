package main
import "fmt"

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

func expandedSlice() {
    newArray := [7]int{0,6,4,7,8,9,10}
    sliceOne := newArray[2:5]
    sliceTwo := sliceOne[3:5]
	fmt.Println(sliceOne)
    fmt.Println(sliceTwo)
}
func addElementToSlice() {
    newArray := [7]int{0,6,4,7,8,9,10}
    sliceOne := newArray[1:5]
    s2 := append(sliceOne, 11)
    s3 := append(s2, 12)
    s4 := append(s3, 13)
	fmt.Println(s4)
    fmt.Println(newArray)
}


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

func copySlice(){
    s2 := []int{2,4,6,8}
	s3 := make([]int, 10, 32)
	fmt.Println(s2)
    copy(s3, s2)
    fmt.Println(s2)
}

func removeSlice() {
	s2 := []int{2,4,6,8,9,9}
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
}

func main() {
	// defineSlice()

	// newArray := [5]int{0,6,4,7,8}
	// newSlice := newArray[1:4]
	// fmt.Println(newSlice)
	// updateSlice(newSlice)
	// fmt.Println(newSlice)
	// fmt.Println(newArray)
	// expandedSlice()
	// addElementToSlice()
	// defineDifferentSlice()
	copySlice()
	removeSlice()
}