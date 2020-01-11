package main

import (
	"fmt"
	"github.com/go-post-api/pkg/fetch"
	"github.com/go-post-api/pkg/post"
	"github.com/go-post-api/pkg/user"
	"os"
	"runtime"
	"strconv"
	"sync"
)

var postService = fetch.PostFetcher{}
var userService = fetch.UserFetcher{}
var wg = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(4)
	userIds := os.Args[1:]

	for _, id := range userIds {
		wg.Add(1)
		//To run asynchronously
		fmt.Println("Fetching: " + id)
		id, _ := strconv.Atoi(id)
		go func() {
			fetchedPost := fetchPost(id)

			fmt.Println(fetchedPost)
		}()
	}

	wg.Wait()
	fmt.Println("Done")
}

func fetchPost(userId int) []post.Post {
	defer wg.Done()
	postData, err := postService.Fetch(userId)

	if err != nil {
		panic(err.Error())
	}

	fetchedUser := fetchUser(userId)

	for i := range postData {
		postData[i].User = &fetchedUser
	}

	return postData
}

func fetchUser(id int) user.User {
	fetched, err := userService.Fetch(id)

	if err != nil {
		panic(err.Error())
	}

	return fetched
}
