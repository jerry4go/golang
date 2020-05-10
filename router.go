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

	fmt.Fprintf(w, "Welcome to the HomePage!") // 根据格式说明来将字符串拼接格式化并写到w中
	fmt.Println("Endpoint Hit: HomePage")      // 输出一行

}

// 返回所有文章
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) // 返回一个json格式的response
}

// 返回单篇文章
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) 会捕捉到请求中所解析出的所有参数(map[string]string)
	vars := mux.Vars(r)
	key := vars["id"]
	//fmt.Fprintf(w, "Key: "+key)

	// for 循环遍历出匹配的id ，并且以json形式返回结果
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// 创建新的文章
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// 从 http.Request.Body 或 http.Response.Body 中读取post数据
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article
	// 对body进行json解码,并且添加到数组里面
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	// 对请求的参数进行编码，并返回给用户
	json.NewEncoder(w).Encode(article)
}

// 删除指定文章
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) 会捕捉到请求中所解析出的所有参数(map[string]string)
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {

		if article.Id == id {
			// 删除ID所匹配的index所在的值
			Articles = append(Articles[:index], Articles[index+1:]...)

		}
	}
	json.NewEncoder(w).Encode(Articles)
}

// 修改指定的文章
func updateArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var updateArticle Article

	for index, article := range Articles {

		if article.Id == id {
			// 删除ID所匹配的index所在的值
			Articles = append(Articles[:index], Articles[index+1:]...)
			// 获取请求中新的post参数，进行json解码
			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &updateArticle)
			// 把解码的数据添加到数组中，并进行json编码返回
			Articles = append(Articles, updateArticle)
			json.NewEncoder(w).Encode(updateArticle)
			return

		}
	}
}

// 路由
func handleRequests() {
	// 新建一个路由的实例
	// 并且设置URL末尾的斜线模糊匹配，即 /path/ 可以访问到 /path, /path 也可以访问到/path/
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	// 监听8787端口，并把hanlder实例化函数myRouter绑定
	log.Fatal(http.ListenAndServe(":8787", myRouter))
}

// 主入口
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
