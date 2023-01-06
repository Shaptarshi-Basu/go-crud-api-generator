## Go API generator
 Generate basic REST API tempate with handlers for the methods with corresponing request and response for with go from specfication file which could be in yaml or json.

### Basic spec file structure

```json
{
	"paths": {
		"/path": [{
			"method": "post",
			"request": "refObject",
			"responses": [{
				"200": "someObj"
			}]
		}]
	},
	"refs": {
		"SomeObj": {
			"sometrribute1": ["ref.SomeObj3"],
			"sometrribute2": "string",
			"sometrribute3": 1.4,
			"sometrribute4": true,
			"sometrribute5": "ref.SomeObj2"
		},
		"SomeObj2": {
			"attr1": [1],
			"attr2": "string"
		},
		"SomeObj3": {
			"attr4": [true],
			"attr5": "string",
			"attr6": {
				"string": true
			}
		}
	}
}
```
