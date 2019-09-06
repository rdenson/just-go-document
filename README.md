# Document Templates
Patterns, examples and research for how to document golang projects.

## REST API Documentation
### Maker API
Simplistic api written plainly order to see usage of swagger (_openAPI_). There
are two instances of maker API: a vanilla implementation using go's `net/http`
package and a cleaner "2.0" version using gin-gonic. Both instances allow you to
serve up the swagger UI.

#### dependencies
before you begin, you'll need to execute the following
```sh
# install swagger (you could also go get go-swagger)
brew tap go-swagger/go-swagger
brew install go-swagger
# then grab the swagger-ui repo
git clone https://github.com/swagger-api/swagger-ui.git
# we need to copy the dist folder from this ↑ repo into ours
```
url in _dist/index.html_ needs to be changed to the spec file to be created later

#### commands
`swagger init spec` - initialize the specification

`swagger generate spec` - build the specification in json or yaml
```sh
# in project's root
#   option [m] includes models that were annotated with 'swagger:model'
#   option [o] writes to spec file
swagger generate spec -mo ./swagger.yml
```
