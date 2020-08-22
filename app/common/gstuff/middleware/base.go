package middleware

import (
	"app/common/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var cfg = config.GetConfig()

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// LogBody ..
func LogBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		req := c.Request()
		res := c.Response()

		//set requestID
		c.Set("reqID", res.Header().Get(echo.HeaderXRequestID))

		//request
		bodyRequest := []byte{}
		if req.Body != nil { // Read
			bodyRequest, _ = ioutil.ReadAll(req.Body)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyRequest)) // Reset

		//response
		bodyResponse := new(bytes.Buffer)
		mw := io.MultiWriter(res.Writer, bodyResponse)
		writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: res.Writer}
		res.Writer = writer

		start := time.Now()
		if err := next(c); err != nil {
			c.Error(err)
		}
		stop := time.Now()

		bodyRequestLimit := string(bodyRequest)
		if len(bodyRequestLimit) > 5000 {
			bodyRequestLimit = bodyRequestLimit[:5000]
		}

		bodyResponseLimit := string(bodyResponse.Bytes())
		if len(bodyResponseLimit) > 5000 {
			bodyResponseLimit = bodyResponseLimit[:5000]
		}
		data := map[string]interface{}{
			"method":        req.Method,
			"headers":       req.Header,
			"body-request":  bodyRequestLimit,
			"status":        res.Status,
			"body-response": bodyResponseLimit,
			"remote-ip":     c.RealIP(),
			"user-agent":    req.UserAgent(),
			"latency-human": stop.Sub(start).String(),
			"request-id":    res.Header().Get(echo.HeaderXRequestID),
			"uri":           req.RequestURI,
		}

		// send data to ElasticSearch (***)
		dataJSON, _ := json.MarshalIndent(data, "", "  ")
		fmt.Print(string(dataJSON))

		return nil
	}
}
