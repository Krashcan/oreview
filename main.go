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
	var Movies []fileInfo
	Movies = nil

	wg := new(sync.WaitGroup)
	rawMovies := r.FormValue("files")
	queryNames := strings.Split(rawMovies,"$`&")
	
	for i:=0;i<len(queryNames)-1;i++{
		go GetTitleAndYear("https://www.opensubtitles.org/libs/suggest.php?format=json3&MovieName=" + url.QueryEscape(queryNames[i]),wg,&Movies)
    	wg.Add(1)
    }
    wg.Wait()	
	fmt.Println("Done")
	
	MergeSort(&Movies,0,len(Movies)-1)
    t,err := template.ParseFiles("template/movies.tpl") 
	if(err!=nil){
		log.Fatal(err)
	}
	t.Execute(w,Movies)	
}


func GetTitleAndYear(url string,wg *sync.WaitGroup,Movies *([]fileInfo)){
	defer wg.Done()
	var movie struct{
		Id  string `json:"pic"`
	}
	resp,err := http.Get(url)
	if err!=nil{
		fmt.Println(err)
		wg.Add(1)
		GetTitleAndYear(url,wg,Movies)
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
	jsonParser := json.NewDecoder(strings.NewReader(data))
	if err := jsonParser.Decode(&movie); err!=nil{
		fmt.Println("Parsing config file: ",err)
	}
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
    if x.Title =="Sample This"{
    	return
    }
    *(Movies) = append(*(Movies),x)
    fmt.Println(x.Title,x.Year)
 }

func MergeSort(Movies *([]fileInfo),l int,r int){
	if(l<r){
		mid := l + (r-l)/2
		MergeSort(Movies,l,mid)
		MergeSort(Movies,mid+1,r)
		Merge(Movies,l,mid,r)
	}
}

func Merge(Movies *([]fileInfo),l int,mid int,r int){
	n1 :=  mid - l + 1 
	n2 := r - mid 

	var L []fileInfo
	var R []fileInfo

	for i:=0;i<n1;i++{
		L = append(L,(*Movies)[l+i])
	}
	for i:=0;i<n2;i++{
		R = append(R,(*Movies)[mid+1+i])
	}
	i:=0
	j:=0
	k:=l

	for i<n1 && j<n2{
		a,_ := strconv.ParseFloat(L[i].Rating,64)
		b,_ := strconv.ParseFloat(R[j].Rating,64)
		if a>=b{
			(*Movies)[k]=L[i]
			i++
		}else{
			(*Movies)[k]=R[j]
			j++
		}
		k++
	}
	for i<n1{
		(*Movies)[k] = L[i]
		i++
		k++
	}
	for j<n2{
		(*Movies)[k] = R[j]
		j++
		k++
	}
}