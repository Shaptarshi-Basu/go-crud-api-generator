package apitemplates

const (
	ModelTemplate = `
type StructName struct {
		{{- range $key, $value := . }}
		{{ $key }} {{ $value }}
		{{- end }}
	}
`
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
	http.ListenAndServe(":<PORT>", router)
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
module {{ .ProjectName }}

go {{ .GoVersion }}

require github.com/gorilla/mux v1.8.0
`
)
