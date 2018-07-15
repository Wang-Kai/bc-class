package middleware

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/apex/log"
	"github.com/urfave/negroni"

	"github.com/bc-class/utils"
)

// Logger log request body and relative infomation
func Logger() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// log http method、api、requset body
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.RespMsg(w, r, err)
		}
		reqBodyStr := string(reqBody[:])
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Body":   reqBodyStr,
			"API":    r.URL.Path,
		}).Info("LoggerMiddleware")

		// put request body into context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "reqBody", reqBody)

		next(w, r.WithContext(ctx))
	}
}
