package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json: "title"`
	Content string `json:"content"`
}

/* 
	We're storing the article list in memory
	In a real application, this list will most likely be fetched
	from a database or from static files
*/
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 Body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 Body"},
}
// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}
