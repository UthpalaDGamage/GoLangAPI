package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

)

const (
	connHost = "localhost"
	connPort = "9001"
	connType = "tcp"
)

type Article struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

//global article array to populate data
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home Endpoint Hit")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){ //--   /articles
	fmt.Println("Endpoint Hit : returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){  //--   /article/{id}
	vars := mux.Vars(r) // reading vars sending from request uri
	key := vars["id"] //get if from vars 

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

	fmt.Fprintf(w, "Key: "+ key)
}

func createNewArticle(w http.ResponseWriter, r *http.Request){
    // get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(r.Body)
   // fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) // reading vars sending from request uri
	id := vars["id"] //get if from vars 

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests(){
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles",returnAllArticles)
	myRouter.HandleFunc("/article",createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}",deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}",returnSingleArticle)
	log.Fatal(http.ListenAndServe(":9000",myRouter))
}

func main(){
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
        Article{Id:"1", Title: "Hello 1", Desc: "Article Description sd ", Content: "Article Content1"},
        Article{Id:"2", Title: "Hello 2", Desc: "Article Description sad ", Content: "Article Content2"},
    }
	handleRequests()
}