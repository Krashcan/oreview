package main

import(
	"log"
	"os"
	"golang.org/x/net/html"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"net/url"
	"fmt"
	"sync"
	"strconv"
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

var movie struct{
	Name string;
	Year string;
}
var root *node
var Movies []fileInfo

func main(){
	router := httprouter.New()

	router.GET("/",GetFileNames)
	router.POST("/",ProcessFileNames)
	router.ServeFiles("/static/*filepath",http.Dir(os.Getenv("GOPATH")+"/src/github.com/krashcan/oreview/static/"))
	port := os.Getenv("PORT")
	if port == "" {
	  	port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port,router))
}

func GetFileNames(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	t,err := template.ParseFiles(os.Getenv("GOPATH")+"/src/github.com/krashcan/oreview/template/index.html")
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
		go GetTitleAndYear("https://opensubtitles.co/search?q=" + url.QueryEscape(queryNames[i]),wg)
    	wg.Add(1)
    }
    wg.Wait()
    t,err := template.ParseFiles(os.Getenv("GOPATH")+"/src/github.com/krashcan/oreview/template/movies.tpl") 
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
	resp,err := http.Get(url)
	if err!=nil{
		fmt.Println(err)
		wg.Add(1)
		GetTitleAndYear(url,wg)
		return
	}
	defer resp.Body.Close()
	var movieData string
	if resp.StatusCode != 200 {
		fmt.Println("Error statuscode: ",resp.StatusCode)
		wg.Add(1)
		GetTitleAndYear(url,wg)
		return	
	}
	z := html.NewTokenizer(resp.Body)
	for{
		tt := z.Next()

		if tt == html.ErrorToken{
			return
		}else if tt==html.StartTagToken{
			t:= z.Token()
			if t.Data=="h4"{
				tt = z.Next()
				tt = z.Next()
				tt = z.Next()
				t = z.Token()
				movieData = strings.TrimSpace(t.Data)				
				break
			}
		}
	}

	movie.Name = movieData[:len(movieData)-6]
	movie.Year = movieData[len(movieData)-5:len(movieData)-1]
	movie.Name = strings.Replace(movie.Name, " ", "+", -1)
	uri := "http://www.omdbapi.com/?t=" + movie.Name + "&y=" + movie.Year + "&plot=short&r=json"  
	resp,err = http.Get(uri)

	if err!=nil{
		log.Fatal(err)
	}
	x := fileInfo{}
	jsonParser := json.NewDecoder(resp.Body)
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

