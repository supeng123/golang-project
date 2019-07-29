package main

import "net/http"
// import "net/http/httputil"
import "io/ioutil"
import "fmt"

func main(){
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error: statue code", resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s\n", all)
}