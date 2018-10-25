package news_sample

type News struct {
	Id       int64    `json:"id"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	Tags     []string `json:"tags"`
	Date     int64    `json:"date"` // unix time
	IdAuthor int64    `json:"idAuthor"`
}

type NewsInput struct {
	Id       int64    `json:"id"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	Tags     []string `json:"tags"`
	Date     int64    `json:"date"` // unix time
	IdAuthor int64    `json:"idAuthor"`
}

type Comment struct {
	Id         int64  `json:"id"`
	CreateDate int64  `json:"createDate"`
	UpdateDate int64  `json:"updateDate"`
	IdAuthor   int64  `json:"idAuthor"`
	Text       string `json:"text"`
	Likes      int    `json:"likes"`
	Dislikes   int    `json:"dislikes"`
	IdNews     int64  `json:"idNews"`
}

type CommentInput struct {
	Id       int64  `json:"id"`
	IdAuthor int64  `json:"idAuthor"`
	Text     string `json:"text"`
	IdNews   int64  `json:"idNews"`
}

type Author struct {
	Id   int64
	Name string
}
