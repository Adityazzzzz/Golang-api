package students

import "net/http"

func New() http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request){
		
		res.Write([]byte("helllo"))
	}
}