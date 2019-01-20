package modules

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello from the api!")

	renderTemplate(w, "main")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// fmt.Println("INFO: Current Directory: ", dir)

	tempalteDir := "cmd/api/html/"
	fmt.Println("INFO: Current templates directory: ", dir+"/"+tempalteDir)

	t, err := template.ParseFiles(tempalteDir+tmpl+".html", tempalteDir+"header.html", tempalteDir+"footer.html")

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	} else {
		t.Execute(w, "")
	}

}
