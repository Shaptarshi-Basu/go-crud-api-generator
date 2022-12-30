## Go API generator
 Generate basic REST API tempate with handlers for the methods with corresponing request and response for with go from specfication file which could be in yaml or json.

### Basic spec file structure

```json
{
	"Paths": {
		"/path": {
			"method": "post"
			"request": "refObject"
			"responses": ["code" : "refObject"]
		}
	}
	"refObject" : "interface {}"
	"databases": [] --> optional -- start with mariadb
}
```
