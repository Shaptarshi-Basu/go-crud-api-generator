package apitemplates

const MainTemplate = `
package main

import (

	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	{{- range $key, $value := . }}
	{{- range $path := $value}}
	router.HandleFunc("{{ $key }}", {{ ($path).FuncName }}).methods("{{ ($path).Method }}")
	{{- end }}
	{{- end }}
	http.ListenAndServe(":8080", router)
}
`
