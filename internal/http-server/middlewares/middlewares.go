package middlewares

import (
	"bytes"
	"encoding/json"
	"factorial_calculate/internal/http-server/requests"
	"factorial_calculate/internal/lib/rp"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CheckParamsMiddleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		var reqBody requests.RequestBody

		body, _ := io.ReadAll(r.Body)

		if err := json.Unmarshal(body, &reqBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(rp.Error("Bad JSON input"))
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		if reqBody.A <= 0 || reqBody.B <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(rp.Error("Incorrect input"))
			return
		}

		n(w, r, ps)
	}
}
