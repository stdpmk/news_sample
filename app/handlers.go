package app

import (
	"encoding/json"
	"log"
	"net/http"

	api "github.com/stdpmk/news_sample"
	"github.com/stdpmk/news_sample/utils"
)

// News
func CreateNews(w http.ResponseWriter, r *http.Request) {

	var newsInput api.NewsInput

	if err := json.NewDecoder(r.Body).Decode(&newsInput); err != nil {
		log.Print("error, %v", err)
		responseJson(w, http.StatusBadRequest, ApiResponse{ErrInvalidJson, "Invalid input json", nil})
		return
	}

	id, err := app.db.CreateNews(newsInput)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Create news error", nil})
		return

	}
	responseJson(w, 201, ApiResponse{Ok, "", id})
}

func GetNewsList(w http.ResponseWriter, r *http.Request) {

	date := utils.QueryParamInt(r, "date")
	newsList, err := app.db.GetNewsList(date)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Get news list error", nil})
		return
	}

	responseJson(w, 200, ApiResponse{Ok, "", newsList})
}

func GetNews(w http.ResponseWriter, r *http.Request) {

	id := utils.PathParamInt(r, "id")
	news, err := app.db.GetNews(id)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Get news error", nil})
		return
	}

	responseJson(w, 200, ApiResponse{Ok, "", news})

}

func UpdateNews(w http.ResponseWriter, r *http.Request) {

	var newsInput api.NewsInput
	if err := json.NewDecoder(r.Body).Decode(&newsInput); err != nil {
		log.Print("error, %v", err)
		responseJson(w, http.StatusBadRequest, ApiResponse{ErrInvalidJson, "Invalid input json", nil})
		return
	}

	id, err := app.db.UpdateNews(newsInput)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Update news error", nil})
		return

	}
	responseJson(w, 200, ApiResponse{Ok, "", id})
}

// Commets
func CreateComment(w http.ResponseWriter, r *http.Request) {

	var commentInput api.CommentInput
	if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
		log.Print("error, %v", err)
		responseJson(w, http.StatusBadRequest, ApiResponse{ErrInvalidJson, "Invalid input json", nil})
		return
	}

	id, err := app.db.CreateComment(commentInput)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Create comment error", nil})
		return

	}
	responseJson(w, 201, ApiResponse{Ok, "", id})

}
func GetComment(w http.ResponseWriter, r *http.Request) {

	id := utils.PathParamInt(r, "id")
	comment, err := app.db.GetComment(id)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Get comment error", nil})
		return
	}

	responseJson(w, 200, ApiResponse{Ok, "", comment})

}
func GetCommentList(w http.ResponseWriter, r *http.Request) {

	idNews := utils.QueryParamInt(r, "idNews")
	commentList, err := app.db.GetCommentList(idNews)

	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Get comment list error", nil})
		return
	}

	responseJson(w, 200, ApiResponse{Ok, "", commentList})

}

func UpdateComment(w http.ResponseWriter, r *http.Request) {

	var commentInput api.CommentInput
	if err := json.NewDecoder(r.Body).Decode(&commentInput); err != nil {
		log.Print("error, %v", err)
		responseJson(w, http.StatusBadRequest, ApiResponse{ErrInvalidJson, "Invalid input json", nil})
		return
	}

	id, err := app.db.UpdateComment(commentInput)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Update comment error", nil})
		return

	}
	responseJson(w, 200, ApiResponse{Ok, "", id})

}

func DeleteComment(w http.ResponseWriter, r *http.Request) {

	id := utils.PathParamInt(r, "id")
	err := app.db.DeleteComment(id)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Delete comment error", nil})
		return
	}
	responseJson(w, 200, ApiResponse{Ok, "", nil})
}

// Author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var author api.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		log.Print("error, %v", err)
		responseJson(w, http.StatusBadRequest, ApiResponse{ErrInvalidJson, "Invalid input json", nil})
		return
	}

	id, err := app.db.CreateAuthor(author)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Create author error", nil})
		return

	}
	responseJson(w, 201, ApiResponse{Ok, "", id})

}

func GetAuthor(w http.ResponseWriter, r *http.Request) {

	id := utils.PathParamInt(r, "id")
	author, err := app.db.GetAuthor(id)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Get author error", nil})
		return
	}

	responseJson(w, 200, ApiResponse{Ok, "", author})

}

func GetAuthorList(w http.ResponseWriter, r *http.Request) {

	authorList, err := app.db.GetAuthorList()
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Get author list error", nil})
		return
	}

	responseJson(w, 200, ApiResponse{Ok, "", authorList})
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {

	var author api.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		log.Print("error, %v", err)
		responseJson(w, http.StatusBadRequest, ApiResponse{ErrInvalidJson, "Invalid input json", nil})
		return
	}

	id, err := app.db.UpdateAuthor(author)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Update author error", nil})
		return

	}
	responseJson(w, 200, ApiResponse{Ok, "", id})

}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {

	id := utils.PathParamInt(r, "id")
	err := app.db.DeleteAuthor(id)
	if err != nil {
		log.Printf("error: %v", err)
		responseJson(w, http.StatusInternalServerError, ApiResponse{ErrInternal, "Delete author error", nil})
		return
	}
	responseJson(w, 200, ApiResponse{Ok, "", nil})

}
