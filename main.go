package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Body  string `json:"body"`
	Title string `json:"title"`
}
type Comment struct {
	ID     string `json:"id"`
	Text   string `json:"Text"`
	PostID string `json:"PostID"`
}

var posts = []Post{
	{ID: "1", Title: "Post1", Body: "Post1"},
	{ID: "2", Title: "Post2", Body: "Post2"},
	{ID: "3", Title: "Post3", Body: "Post3"},
}
var Comments = []Comment{
	{ID: "1", Text: "comm1", PostID: "1"},
	{ID: "2", Text: "comm2", PostID: "1"},
	{ID: "3", Text: "comm3", PostID: "2"},
}

func getAllposts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func Createpost(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	newPost.ID = strconv.Itoa(rand.Intn(1000))
	posts = append(posts, newPost)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

func Updatepost(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, post := range posts {
		if post.ID == params["id"] {
			p := &posts[index]
			p.Body = newPost.Body
			//p.Title = newPost.Title
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	json.NewEncoder(w).Encode(newPost)
}

func Deletepost(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, post := range posts {
		if post.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	json.NewEncoder(w).Encode(newPost)
}

func GetCommentsOfPost(w http.ResponseWriter, r *http.Request) {
	var Comment []Comment
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)

	for _, obj := range Comments {
		if obj.PostID == param["id"] {
			Comment = append(Comment, obj)
		}
	}
	json.NewEncoder(w).Encode(Comment)
}

//task-66926
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var newComment Comment
	json.NewDecoder(r.Body).Decode(&newComment)
	newComment.ID = strconv.Itoa(rand.Intn(1000))
	Comments = append(Comments, newComment)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(newComment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var newComment Comment
	json.NewDecoder(r.Body).Decode(&newComment)
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, comment := range Comments {
		if comment.ID == params["id"] {
			fmt.Println(comment.ID)
			c := &Comments[index]
			c.Text = newComment.Text
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	json.NewEncoder(w).Encode(newComment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	var newComment Comment
	json.NewDecoder(r.Body).Decode(&newComment)
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, comment := range Comments {
		if comment.ID == params["id"] {
			Comments = append(Comments[:index], Comments[index+1:]...)
			json.NewEncoder(w).Encode(comment)
			return
		}
	}
	json.NewEncoder(w).Encode(newComment)
}

func newRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/Post/getAllposts", getAllposts).Methods("GET")
	r.HandleFunc("/Post/Createpost", Createpost).Methods("POST")
	r.HandleFunc("/Post/Updatepost/{id}", Updatepost).Methods("PUT")
	r.HandleFunc("/Post/DeletePost/{id}", Deletepost).Methods("DELETE")
	r.HandleFunc("/Comment/getComments/{id}", GetCommentsOfPost).Methods("GET")
	r.HandleFunc("/Comment/CreateComment", CreateComment).Methods("POST")
	r.HandleFunc("/Comment/UpdateComment/{id}", UpdateComment).Methods("PUT")
	r.HandleFunc("/Comment/DeleteComment/{id}", DeleteComment).Methods("DELETE")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}
func main() {

	r := newRouter()
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, r)
}
