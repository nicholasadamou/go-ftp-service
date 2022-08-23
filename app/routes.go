package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
)

func SFTP(c *gin.Context) {
	r := c.Request
	w := c.Writer

	var fileItem FileItem

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 30000000))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &fileItem); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	file := fileItem
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	Upload(file)

	if err := json.NewEncoder(w).Encode(file); err != nil {
		panic(err)
	}
}

func health(c *gin.Context) {
	w := c.Writer

	w.WriteHeader(http.StatusOK)
}
