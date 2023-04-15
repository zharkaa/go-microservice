package main

import (
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

/*
We're storing the article list in memory
In a real application, this list will most likely be fetched
from a database or from static files
*/
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 Body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 Body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}
