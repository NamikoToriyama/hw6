package app

import (
	"fmt"
	"net/http"
	"strings"
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
	fmt.Fprintf(w, "<input type=\"submit\" value=\"submit\"></input>\n")
	fmt.Fprintf(w, "</form>\n")
	mixWord := stringMix(r.FormValue("word1"), r.FormValue("word2"))
	fmt.Fprintf(w, "'%s'!!!</div>\n", mixWord) //返って来たのを表示

}
func stringMix(word1 string, word2 string) []string {

	//num := utf8.RuneCountInString(word1) + utf8.RuneCountInString(word2)
	var mix []string
	var number int
	//1文字ずつに分解
	splitWord1 := strings.Split(word1, "")
	splitWord2 := strings.Split(word2, "")
	if len(splitWord1) < len(splitWord2) {
		number = len(splitWord2)
	} else {
		number = len(splitWord1)
	}

	//新しい配列みたいなのに入れる　string型
	for i := 0; i < number; i++ {
		if i < len(splitWord1) {
			mix = append(mix, splitWord1[i])
		}
		if i < len(splitWord2) {
			mix = append(mix, splitWord2[i])
		}
	}

	//その新しい配列を返す
	return mix
}
