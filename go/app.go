package app

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<div style=\"text-align:center;\">")
	fmt.Fprintf(w, "<div>Hello STEP!</div>")
	fmt.Fprintf(w, "<div>word1\n")
	fmt.Fprintf(w, "<input type=\"text\"></input></div>\n")
	fmt.Fprintf(w, "<div>word2\n")
	fmt.Fprintf(w, "<input type=\"text\"></input></div>\n")
	fmt.Fprintf(w, "<input type=\"button\" value=\"submit\"></input></div>\n")
	fmt.Fprintf(w, "</div>")
}
func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "foo is '%s'!", r.FormValue("foo"))
}
