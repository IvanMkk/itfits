package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	schema "main/src/schemas"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateTestUser() (*schema.Token, error) {
	bodyReader := strings.NewReader(`{"email": "12124", "password": "testing"}`)
	req := httptest.NewRequest(http.MethodPost, "/v1/registration", bodyReader)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	os.Chdir("../..") // to read test.env

	// set the request to the gin context
	c.Request = req

	app := New()
	app.RegistrationHandler(c)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Can't register test data")
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	bodyReader = strings.NewReader(`{"email": "12124", "password": "testing"}`)
	c.Request = httptest.NewRequest(http.MethodGet, "/v1/authentications", bodyReader)

	app.AuthenticationsHandler(c)

	res = w.Result()
	defer res.Body.Close()
	var data struct {
		Token schema.Token `json:"token"`
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Can't register test data")
	}

	if json.Unmarshal(body, &data) != nil {
		return nil, fmt.Errorf("Can't register test data")
	}
	return &data.Token, nil
}
