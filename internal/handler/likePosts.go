package handler

import (
	"log"
	"net/http"

	"forum/internal/render"
)

func (h *Handler) likePostsGET(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likeposts" {
		log.Printf("likePostsGET:StatusNotFound:%s\n", r.URL.Path)
		h.renderError(w, http.StatusNotFound) // 404
		return
	}
	if r.Method != http.MethodGet {
		log.Printf("likePostsGET:StatusMethodNotAllowed:%s\n", r.Method)
		h.renderError(w, http.StatusMethodNotAllowed) // 405
		return
	}

	user := h.getUserFromContext(r)

	posts, err := h.service.GetPostsByLike(user.Id)
	if err != nil {
		log.Printf("likePostsGET:GetPostsByLike:%s\n", err.Error())
		h.renderError(w, http.StatusInternalServerError) // 500
		return
	}
	h.renderPage(w, "home.html", &render.Data{
		User:  user,
		Posts: posts,
	})
}
