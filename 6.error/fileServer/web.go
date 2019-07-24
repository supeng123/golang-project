package main

import "net/http"
import "os"
import "io/ioutil"
import "fmt"

func handleFileList() {
	http.HandleFunc("/list/", 
		func(write http.ResponseWriter,
			request *http.Request){
				path := request.URL.Path[len("/list/"):]
				file, err := os.Open(path)

				if  err != nil {
					http.Error(write, err.Error(), http.StatusInternalServerError)
					return
				}
				defer file.Close()
				all, err := ioutil.ReadAll(file)
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
				write.Write(all) 

			})
	err := http.ListenAndServe(":8888", nil)
	if  err != nil {

	}
}

func main() {
	
} 