package internal

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"math"
	"net/http"
)

// Fields represents JSON
type Fields map[string]interface{}

// Response is written to http.ResponseWriter
type Response struct {
	Code    int
	Payload interface{}
}

// Make creates a http handler from a request handler func
func Make(
	f func(req *http.Request) Response,
) func(w http.ResponseWriter, req *http.Request) {
	handler := func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		if (*req).Method == "OPTIONS" {
			req.Body.Close()
			return
		}
		w.Header().Set("Content-Type", "application/json")
		res := f(req)
		JSON, err := json.Marshal(res.Payload)
		if err != nil {
			logrus.WithError(err).Fatal("json marshal failed")
		}

		w.WriteHeader(res.Code)
		w.Write(JSON)
		req.Body.Close()

	}

	return handler
}

// deadcode
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetBytesFromReq(reqBody io.ReadCloser) (rawBody []byte) {
	rawBody, err := ioutil.ReadAll(reqBody)
	if err != nil {
		logrus.WithError(err).Warn("could not read bytes from request")
	}
	return
}

func BadRequest(msg string) Response {
	return Response{
		http.StatusBadRequest,
		Fields{"error": msg, "code": http.StatusBadRequest},
	}
}

func InternalServerError(msg string) Response {
	return Response{
		http.StatusInternalServerError,
		Fields{"error": msg, "code": http.StatusInternalServerError},
	}
}

func IsSquare(x int) (bool, int) {
	sq := int(math.Sqrt(float64(x)))
	return sq * sq == x, sq
}
