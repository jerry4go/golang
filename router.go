package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

// 首页
func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: HomePage")

}

// 返回所有文章
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// 返回单篇文章
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	//fmt.Fprintf(w, "Key: "+key)

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// 创建新的文章
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article

	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

// 删除指定文章
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {

		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			reqBody, _ := ioutil.ReadAll(r.Body)

			var article Article
			//_ = json.NewEncoder(r.Body).Decode(&article)
			json.Unmarshal(reqBody, &article)
			article.Id = vars["id"]
			Articles = append(Articles, article)
			json.NewEncoder(w).Encode(&article)
			return
		}
	}
	json.NewEncoder(w).Encode(Articles)
}

// 修改指定的文章
func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update article")
	//w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var updateArticle Article

	for index, article := range Articles {

		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)

			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &updateArticle)
			fmt.Println(updateArticle)
			//file, _ := os.open(r.Body)
			//json.Unmarshal(reqBody, &updateArticle)
			//json.NewEncoder(r.Body).Decode(&updateArticle)
			Articles = append(Articles, updateArticle)
			json.NewEncoder(w).Encode(updateArticle)
			return
		}
	}
}

// 路由
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8787", myRouter))
}

// 入口
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
