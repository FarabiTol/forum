package handler

import (
	"errors"
	"fmt"
	"forum/internal/models"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func (h *Handler) getUserFromContext(r *http.Request) *models.User {
	user, ok := r.Context().Value(keyUser).(*models.User)
	if !ok {
		return nil
	}
	return user
}

func (h *Handler) getPostIdFromURL(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		return 0, errors.New("incorrect path")
	}
	rx := regexp.MustCompile(`^[^0,+,-]{1,}\d*$`)
	if !rx.MatchString(parts[2]) {
		return 0, fmt.Errorf("incorrect request vote = %s", parts[2])
	}

	postId, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	return postId, nil
}

func (h *Handler) getVote(voteStr string) (int, error) {
	rx := regexp.MustCompile(`^[^0,+]{1,}\d*$`)
	if !rx.MatchString(voteStr) {
		return 0, fmt.Errorf("incorrect request vote = %s", voteStr)
	}
	vote, err := strconv.Atoi(voteStr)
	if err != nil {
		return 0, err
	}
	if vote != 1 && vote != -1 {
		return 0, fmt.Errorf("incorrect request vote = %d", vote)
	}
	return vote, nil
}

func (h *Handler) getIntFromForm(r *http.Request, key string) (int, error) {
	value := r.Form.Get(key)

	rx := regexp.MustCompile(`^[^0,+,-]{1,}\d*$`)
	if !rx.MatchString(value) {
		return 0, fmt.Errorf("incorrect request vote = %s", value)
	}

	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return id, nil
}
