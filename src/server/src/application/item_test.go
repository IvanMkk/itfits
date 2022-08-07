package application

import (
	"encoding/json"
	"io/ioutil"
	schema "main/src/schemas"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestAddItemHandler(t *testing.T) {
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
		app.DB.Exec(`DELETE FROM it_users WHERE email = '12124';`)

		return
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	bodyReader = strings.NewReader(`{"email": "12124", "password": "testinasg"}`)
	c.Request = httptest.NewRequest(http.MethodGet, "/v1/authentications", bodyReader)

	app.AuthenticationsHandler(c)

	res = w.Result()
	defer res.Body.Close()
	var data struct {
		Token schema.Token `json:"token"`
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("Can't register test data")
		app.DB.Exec(`DELETE FROM it_users WHERE email = '12124';`)

		return
	}

	if json.Unmarshal(body, &data) != nil {
		t.Error("Can't register test data")
		app.DB.Exec(`DELETE FROM it_users WHERE email = '12124';`)

		return
	}

	testUUID := uuid.New().String()

	addItemRequest := schema.AddItem{
		BrandName:      "test brand",
		GarmentId:      testUUID,
		SizeTypeId:     testUUID,
		SizeTypeItemId: testUUID,
		ItemState:      "active",
	}

	bodyJson, err := json.Marshal(addItemRequest)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	bodyReader = strings.NewReader(string(bodyJson))
	c.Request = httptest.NewRequest(http.MethodGet, "/v1/authentications", bodyReader)

	// add authorization header to the req
	var bearer = "Bearer " + data.Token.AuthToken
	c.Request.Header.Add("Authorization", bearer)

	app.AddItemHandler(c)

	res = w.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Can't add item")
	}

	app.DB.Exec(`DELETE FROM it_users WHERE email = '12124';`)
}
