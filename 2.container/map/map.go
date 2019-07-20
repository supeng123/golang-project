package main

import "fmt"

// func defineMap() {
// 	m := map[string]string {
// 		"name": "course",
// 		"site": "golang"
// 	}

// 	m2 := make(map[string]int) //m2== empty
// 	var m3 = make(map[string]int) //m3 == nil

// }

// func iteratorMap() {
// 	m := map[string]string {
// 		"name": "course",
// 		"site": "golang"
// 	}

// 	if getName, ok := m["name"]; ok {
// 		fmt.Println(getName)
// 	} else {
// 		fmt.Println("key does not exist")
// 	}

// 	delete(m, "site")
// 	if site, ok := m["site"]; ok {
// 		fmt.Println("site has been deleted")
// 	} else {
// 		fmt.Println("site has not been deleted")
// 	}

// 	for key, value range m {
// 		fmt.Println(k,v)
// 	}
// }

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

func getChineseLength(s string) {
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)\n", i, ch)
	}
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcdcbababcd"))
	bbb := "i love 中古"
	getChineseLength(bbb)
}