package main

import(
	"log"
	"os"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"net/url"
	"fmt"
	"sync"
	"strconv"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
)

type fileInfo struct{
	Title string `json:"Title"`;
	Year string `json:"Year"`;
	Runtime string `json:"Runtime"`;
	Genre string `json:"Genre"` ;
	Rating string `json:"imdbRating"`;
	Description string `json:"Plot"`;
	Image string `json:"Poster"`;
	Awards string `json:"Awards"`;
}

type node struct{
	Movie fileInfo 
	Left *(node) 
	Right *(node) 
}


var root *node
var Movies []fileInfo

func main(){
	router := httprouter.New()

	router.GET("/",GetFileNames)
	router.POST("/",ProcessFileNames)
	router.ServeFiles("/static/*filepath",http.Dir("static/"))
	port := os.Getenv("PORT")
	if port == "" {
	  	port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port,router))
}

func GetFileNames(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	t,err := template.ParseFiles("template/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	t.Execute(w,nil)
}

func ProcessFileNames(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	MakeGlobalNil()
	wg := new(sync.WaitGroup)
	rawMovies := r.FormValue("files")
	queryNames := strings.Split(rawMovies,"$`&")

	for i:=0;i<len(queryNames)-1;i++{
		go GetTitleAndYear("https://www.opensubtitles.org/libs/suggest.php?format=json3&MovieName=" + url.QueryEscape(queryNames[i]),wg)
    	wg.Add(1)
    }
    wg.Wait()
    t,err := template.ParseFiles("template/movies.tpl") 
	if(err!=nil){
		log.Fatal(err)
	}
	t.Execute(w,Movies)	
}

func MakeGlobalNil(){
	Movies = nil
	root = nil
}

func GetTitleAndYear(url string,wg *sync.WaitGroup){
	defer wg.Done()
	var movie struct{
		Id  string `json:"pic"`
	}
	resp,err := http.Get(url)
	if err!=nil{
		fmt.Println(err)
		wg.Add(1)
		GetTitleAndYear(url,wg)
		return
	}
	defer resp.Body.Close()
	movie.Id = ""
	fmt.Println(url)
	body,_:= ioutil.ReadAll(resp.Body)
	data := string(body)
	y:= strings.Index(data,"},{")
	if y!=-1{
		data = data[1:y+1]
	}else{
		data = data[1:len(data)-1]
			
	}
	fmt.Println(data)
	jsonParser := json.NewDecoder(strings.NewReader(data))
	if err := jsonParser.Decode(&movie); err!=nil{
		fmt.Println("Parsing config file: ",err)
	}
	fmt.Println(movie.Id)
	if movie.Id == ""{
		return
	}
	rep := 7 - len(movie.Id)
	url = "http://www.omdbapi.com/?i=tt" + strings.Repeat("0",rep) + movie.Id + "&plot=short&r=json"  
	
	resp,err = http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}
	x := fileInfo{}
	jsonParser = json.NewDecoder(resp.Body)
    if err := jsonParser.Decode(&x); err != nil {
        log.Fatal("parsing config file", err)
    }
    if x == (fileInfo{}){
     	return
    }
    root = InsertTree(root,x)
    Movies = nil
    
    InorderTraversal(root)
    fmt.Println(x.Title,x.Year)
 }

func InsertTree(leaf *node,x fileInfo) *node{
	a,_ := strconv.ParseFloat(x.Rating,32)
	
	if leaf == nil{
		return &node{x,nil,nil}
	}else if b,_ := strconv.ParseFloat(leaf.Movie.Rating,32); a>b{
		leaf.Left = InsertTree(leaf.Left,x)
		return leaf
	}
	leaf.Right = InsertTree(leaf.Right,x)
	return leaf
	
}

func InorderTraversal(leaf *node){
	if leaf == nil{
		return
	}
	InorderTraversal(leaf.Left)
	Movies = append(Movies,leaf.Movie)
	InorderTraversal(leaf.Right)
}

