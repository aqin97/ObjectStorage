package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	path := os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2]
	fmt.Println(path)
	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(file, r.Body)
}
