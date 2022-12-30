package models

type Spec struct {
    Paths   map[string]Path    `json:"paths"`
	Refs     map[string]interface{}	   `json:"refs"`
}


type Path struct {
    Method   string      `json:"method"`
    Request interface{} `json:"request"`
	Responses []interface{} `json:"responses"`
}
