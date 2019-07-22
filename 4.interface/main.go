package main

import "fmt"
import "./mock"
import "./real"
import "time"

type Retriever interface {
	Get(urr string)string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("conent",v.Content)
	case *real.Retriever:
		fmt.Println("useragent", v.UserAgent)
	}
}



func post(post Poster){
	post.Post("http://baidu.com", map[string]string{
		"name": "slogan",
		"course": "golang",
	})
}

type Poster interface {
	Post(url string, 
		form map[string]string) string
}
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

func main(){
	var r Retriever
	r = mock.Retriever{"this is fake imooc it"}
	fmt.Printf("%T %v\n", r, r)
	// fmt.Println(download(r))
	reality := real.Retriever{}
	fmt.Printf("%T %v\n", reality, reality)
	// fmt.Println(download(reality))

	var r2 Retriever
	r2 = real.Retriever{
		UserAgent: "Chrome",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v\n", r2, r2)
	//check the type
	if mockRetriever, ok:= r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Content)
	}
	
	
}