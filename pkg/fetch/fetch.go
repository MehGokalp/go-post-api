package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"post-api/pkg/post"
	"post-api/pkg/user"
	"strconv"
)

//  ---------- Requester Begin  ----------
type Requester interface {
	Fetch(uri string) *http.Response
}

type JsonRequester struct {}

func (r *JsonRequester) Fetch(uri string) *http.Response {
	response, err := http.Get(uri)
	
	if err != nil {
		panic(err.Error())
	}
	
	if response.StatusCode != 200 {
		panic(fmt.Errorf("excepted status code 200 got %d in request: %s", response.StatusCode, response.Request.RequestURI))
	}
	
	return response
}
//  ---------- Requester End  ----------

//  ---------- Parser Begin ----------
type PostParser struct {}

func (p PostParser) Parse(data []byte) (parsedData []post.Post, err error) {
	err = json.Unmarshal(data, &parsedData)
	
	return parsedData, err
}

type UserParser struct {}

func (p UserParser) Parse(data []byte) (user.User, error)  {
	var parsedData []user.User
	
	err := json.Unmarshal(data, &parsedData)
	
	return parsedData[0], err
}
// ---------- Parser End ----------

//  ---------- Fetcher Begin ----------
type PostFetcher struct {}

func (s *PostFetcher) Fetch(id int) ([]post.Post, error) {
	requester := JsonRequester{}
	response := requester.Fetch("https://jsonplaceholder.typicode.com/posts?userId=" + strconv.Itoa(id))
	
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	
	parser := PostParser{}
	parsed, err := parser.Parse(body)
	
	if err != nil {
		panic(err.Error())
	}
	
	return parsed, err
}
//  ---------- Fetcher End ----------

type UserFetcher struct {}

func (s *UserFetcher) Fetch(id int) (user.User, error) {
	requester := JsonRequester{}
	response := requester.Fetch("https://jsonplaceholder.typicode.com/users?id=" + strconv.Itoa(id))
	
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	
	parser := UserParser{}
	parsed, err := parser.Parse(body)
	
	if err != nil {
		panic(err.Error())
	}
	
	return parsed, err
}