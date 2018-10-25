package db

import (
	"database/sql"
	"fmt"
	"time"

	api "github.com/stdpmk/news_sample"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

type Tx struct {
	*sqlx.Tx
}

type ConnOpts struct {
	User         string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

type queryRowInterface interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

type dbInterface interface {
	queryRowInterface
	sqlx.Preparer
	sqlx.Execer
	sqlx.Queryer
}

func openDbConnection(openParams string) (opDb *sqlx.DB) {
	var err error

	opDb, err = sqlx.Open("postgres", openParams)
	opDb.SetMaxOpenConns(30)
	opDb.SetMaxIdleConns(8)
	if err != nil {
		panic(err)
	}
	err = opDb.Ping()
	if err != nil {
		panic(err)
	}
	return
}

func NewDatabase(opts *ConnOpts) *DB {
	connectionParams := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		opts.User, opts.Password, opts.Host, opts.Port, opts.DatabaseName)

	d := openDbConnection(connectionParams)
	return &DB{DB: d}
}

func (db *DB) Now() (time.Time, error) {
	var t pq.NullTime
	err := db.QueryRow(`SELECT now()`).Scan(&t)
	if err != nil {
		return time.Time{}, err
	}
	return t.Time, nil
}

// News

func (db *DB) GetNews(id int64) (api.News, error) {

	var news api.News
	var tagsArray pq.StringArray
	var date time.Time

	err := db.QueryRow(
		`SELECT n.id, n.title, n.text, n.tags, n.date, n.id_author
		FROM news n WHERE id = $1`, id).Scan(&news.Id, &news.Title, &news.Text, &tagsArray, &date, &news.IdAuthor)

	news.Tags = tagsArray
	news.Date = date.Unix()

	return news, err
}

func (db *DB) GetNewsList(date int64) ([]api.News, error) {

	dateUTC := time.Unix(date, 0).UTC()

	rows, err := db.Query(
		`SELECT n.id, n.title, n.text, n.tags, n.date, n.id_author
		FROM news n WHERE date >= $1`, dateUTC)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList = []api.News{}
	var news api.News
	var tagsArray pq.StringArray
	var ddate time.Time

	for rows.Next() {
		err = rows.Scan(&news.Id, &news.Title, &news.Text, &tagsArray, &ddate, &news.IdAuthor)
		if err != nil {
			return nil, err
		}
		news.Tags = tagsArray
		news.Date = ddate.Unix()

		newsList = append(newsList, news)
	}

	return newsList, nil
}

func (db *DB) CreateNews(news api.NewsInput) (int64, error) {

	var id int64
	err := db.QueryRow(
		`INSERT INTO news(title, text, tags, date, id_author) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		news.Title, news.Text, pq.StringArray(news.Tags), time.Unix(news.Date, 0).UTC(), news.IdAuthor,
	).Scan(&id)

	return id, err

}

func (db *DB) UpdateNews(news api.NewsInput) (int64, error) {
	var id int64
	err := db.QueryRow(
		`UPDATE news SET (title, text, tags, date, id_author) = ($2, $3, $4, $5, $6) WHERE id = $1 RETURNING id`,
		news.Id, news.Title, news.Text, pq.StringArray(news.Tags), time.Unix(news.Date, 0).UTC(), news.IdAuthor,
	).Scan(&id)

	return id, err
}

// Comments
func (db *DB) GetComment(id int64) (api.Comment, error) {

	var comment api.Comment
	var createDate time.Time
	var updDate time.Time

	err := db.QueryRow(
		`SELECT c.id, c.createdate, c.updatedate, c.id_author, c.text, c.likes, c.dislikes, c.id_news
		FROM comment c WHERE id = $1`, id).Scan(&comment.Id, &createDate, &updDate, &comment.IdAuthor, &comment.Text, &comment.Likes, &comment.Dislikes, &comment.IdNews)

	comment.CreateDate = createDate.Unix()
	comment.UpdateDate = updDate.Unix()

	return comment, err

}

func (db *DB) GetCommentList(idNews int64) ([]api.Comment, error) {

	var commentList = []api.Comment{}
	var comment api.Comment
	var createDate time.Time
	var updDate time.Time

	rows, err := db.Query(
		`SELECT c.id, c.createdate, c.updatedate, c.id_author, c.text, c.likes, c.dislikes, c.id_news
		FROM comment c WHERE id_news = $1`, idNews)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&comment.Id, &createDate, &updDate, &comment.IdAuthor, &comment.Text, &comment.Likes, &comment.Dislikes, &comment.IdNews)
		if err != nil {
			return nil, err
		}
		comment.CreateDate = createDate.Unix()
		comment.UpdateDate = updDate.Unix()

		commentList = append(commentList, comment)
	}

	return commentList, err
}

func (db *DB) CreateComment(comment api.CommentInput) (int64, error) {
	var id int64
	err := db.QueryRow(
		`INSERT INTO comment(id_author, text, id_news) VALUES ($1, $2, $3) RETURNING id`,
		comment.IdAuthor, comment.Text, comment.IdNews,
	).Scan(&id)

	return id, err
}

func (db *DB) UpdateComment(comment api.CommentInput) (int64, error) {

	var id int64
	err := db.QueryRow(
		`UPDATE comment SET (id_author, text, id_news) = ($2, $3, $4) WHERE id = $1 RETURNING id`,
		comment.Id, comment.IdAuthor, comment.Text, comment.IdNews,
	).Scan(&id)

	return id, err

}

func (db *DB) DeleteComment(id int64) error {

	err := db.QueryRow(`DELETE FROM comment WHERE id = $1 RETURNING id`, id).Scan(&id)
	return err
}

// Authors
func (db *DB) GetAuthor(id int64) (api.Author, error) {

	var author api.Author
	err := db.QueryRow(
		`SELECT id, name
		FROM author WHERE id = $1`, id).Scan(&author.Id, &author.Name)
	return author, err

}

func (db *DB) GetAuthorList() ([]api.Author, error) {

	var author api.Author
	var authorList = []api.Author{}

	rows, err := db.Query(
		`SELECT id, name FROM author`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&author.Id, &author.Name)
		if err != nil {
			return nil, err
		}
		authorList = append(authorList, author)
	}

	return authorList, nil

}

func (db *DB) CreateAuthor(author api.Author) (int64, error) {

	var id int64
	err := db.QueryRow(
		`INSERT INTO author(name) VALUES ($1) RETURNING id`,
		author.Name,
	).Scan(&id)

	return id, err
}

func (db *DB) UpdateAuthor(author api.Author) (int64, error) {

	var id int64
	err := db.QueryRow(
		`UPDATE author SET (name) = ($2) WHERE id = $1 RETURNING id`,
		author.Id, author.Name,
	).Scan(&id)

	return id, err
}

func (db *DB) DeleteAuthor(id int64) error {
	err := db.QueryRow(`DELETE FROM author WHERE id = $1 RETURNING id`, id).Scan(&id)
	return err
}
