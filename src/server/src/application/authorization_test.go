package application

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthenticationsHandler(t *testing.T) {
	bodyReader := strings.NewReader(`{"email": "12124", "password": "testinasg"}`)
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
		t.Error("Can't register test data")
		return
	}

	// w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	bodyReader = strings.NewReader(`{"email": "12124", "password": "testinasg"}`)
	c.Request = httptest.NewRequest(http.MethodGet, "/v1/authentications", bodyReader)

	app.AuthenticationsHandler(c)

	res = w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("Can't register test data")
		return
	}
	app.DB.Exec(`DELETE FROM it_users WHERE email = '12124';`)
}
