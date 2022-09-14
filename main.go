package function

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"	
)

func init() {
	functions.HTTP("FnDeployCI", fnDeployCI)
}

// helloWorld writes "Hello, World!" to the HTTP response.
func fnDeployCI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "function deployed using circle-ci")
}
