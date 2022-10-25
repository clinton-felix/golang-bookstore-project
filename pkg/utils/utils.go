package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// the utils file helps to unmarshall the marshalled json requests from users
// and marshal the unmarshalled responses into JSON

/* creating a function to help parse the body JSON requests from users
calling a POST method; eg while creating a Book in our example instance */
func ParseBody(r *http.Request, x interface{})  {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}