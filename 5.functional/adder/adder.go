package main
import "fmt"

func adder () func(int) int {
	sum := 0;
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdd func(int) (int, iAdd)

func adderTwo(base int) iAdd{
	return func(v int) (int, iAdd) {
		return base + v, adderTwo(base + v)
	}
}

func main() {
	a := adder()

	for i := 0; i < 10 ; i++ {
		fmt.Println(a(i))
	}
}

