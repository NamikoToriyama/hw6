package app

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handlePata)
	http.HandleFunc("/test", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<div style=\"text-align:center;\">")
	fmt.Fprintf(w, "<div>Hello STEP!</div>")
	fmt.Fprintf(w, "<form method=\"get\">\n")
	fmt.Fprintf(w, "<div>word1\n")
	fmt.Fprintf(w, "<input type=\"text\" name=\"word1\"></input></div>\n")
	fmt.Fprintf(w, "<div>word2\n")
	fmt.Fprintf(w, "<input type=\"text\" name=\"word2\"></input></div>\n")
	fmt.Fprintf(w, "<input type=\"submit\" value=\"submit\"></input></div>\n")
	fmt.Fprintf(w, "</form>\n")
	fmt.Fprintf(w, "foo is '%s'!\n", r.FormValue("word1")) //default: "GET /?word1=rrr&word2=www HTTP/1.1" 200 277
	fmt.Fprintf(w, "foo is '%s'!\n", r.FormValue("word2")) //default: "GET /?word1=aaa&word2=ccc HTTP/1.1" 200 277
	//GETで送られていることがわかる
}
