package repositories

import "testing"

func TestHttpExampleRepository_GetPost(t *testing.T) {
	repo := HttpExampleRepository{}.GetPost()
	if repo != nil {
		for _, value := range repo {
			println("title : ", value.Title)
		}
	}
	println(repo)
}