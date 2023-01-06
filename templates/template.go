package apitemplates

const (
	MainTemplate = `
package main

import (
	"<PROJECT_NAME>/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	api := handler.APIHandler{}
	{{- range $key, $value := . }}
	{{- range $path := $value}}
	router.HandleFunc("{{ $key }}", api.{{ ($path).FuncName }}).Methods("{{ ($path).Method }}")
	{{- end }}
	{{- end }}
	http.ListenAndServe(":8080", router)
}
`
	HandlerFuncTemplate = `
	package handler

	import (

		"net/http"
	)

	type APIHandler struct {

	}
	{{- range $key, $value := . }}
	{{- range $path := $value}}
	func (api *APIHandler) {{ ($path).FuncName }}(w http.ResponseWriter, r *http.Request) {

	}
	{{- end }}
	{{- end }}
`
	ModeFileTemplate = `
module {{ . }}

go 1.13

require github.com/gorilla/mux v1.8.0
`
)
