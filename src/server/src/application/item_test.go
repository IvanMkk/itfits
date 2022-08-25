package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	schema "main/src/schemas"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestAddDeleteListItemHandler(t *testing.T) {
	token, err := CreateTestUser()
	if err != nil {
		t.Error(err)
		return
	}
	app := New()
	defer app.DB.Exec(`DELETE FROM it_users WHERE email = '12124';`)

	testUUID := uuid.New().String()

	addItemRequest := schema.AddItem{
		BrandName:      "test brand",
		GarmentId:      testUUID,
		SizeTypeId:     testUUID,
		SizeTypeItemId: testUUID,
	}

	bodyJson, err := json.Marshal(addItemRequest)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	bodyReader := strings.NewReader(string(bodyJson))
	c.Request = httptest.NewRequest(http.MethodGet, "/v1/item", bodyReader)

	// add authorization header to the req
	var bearer = "Bearer " + token.AuthToken
	c.Request.Header.Add("Authorization", bearer)

	app.AddItemHandler(c)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Can't add item")
		return
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("response read error: %v", err)
		return
	}

	itemId := struct {
		Id string `json:"id"`
	}{}

	if err = json.Unmarshal(response, &itemId); err != nil {
		t.Errorf("unexpected output after adding an item")
		return
	}

	if itemId.Id == "" {
		t.Error("unexpected item id")
		return
	}
	t.Logf("itemId.Id: %v\n", itemId.Id)

	// List items
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	bodyReader = strings.NewReader("")
	c.Request = httptest.NewRequest(http.MethodGet, "/v1/item", bodyReader)
	c.Request.Header.Add("Authorization", bearer)

	app.ListItemsHandler(c)
	res = w.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Can't list items")
		return
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/item/%s", itemId.Id), nil)
	c.Request.Header.Add("Authorization", bearer)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: itemId.Id,
	})

	app.GetItemHandler(c)

	res = w.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Can't get item")
		return
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/item/%s", itemId.Id), nil)
	c.Request.Header.Add("Authorization", bearer)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: itemId.Id,
	})

	app.DeleteItemHandler(c)

	res = w.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Can't remove item")
		return
	}

}

func TestEditItemHandler(t *testing.T) {
	token, err := CreateTestUser()
	if err != nil {
		t.Error(err)
		return
	}
	app := New()
	defer app.DB.Exec(`DELETE FROM it_users WHERE email = '12124';`)

	testUUID := uuid.New().String()

	addItemRequest := schema.AddItem{
		BrandName:      "test brand",
		GarmentId:      testUUID,
		SizeTypeId:     testUUID,
		SizeTypeItemId: testUUID,
	}

	bodyJson, err := json.Marshal(addItemRequest)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	bodyReader := strings.NewReader(string(bodyJson))
	c.Request = httptest.NewRequest(http.MethodGet, "/v1/item", bodyReader)

	// add authorization header to the req
	var bearer = "Bearer " + token.AuthToken
	c.Request.Header.Add("Authorization", bearer)

	app.AddItemHandler(c)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Can't add item")
		return
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("response read error: %v", err)
		return
	}

	itemId := struct {
		Id string `json:"id"`
	}{}

	if err = json.Unmarshal(response, &itemId); err != nil {
		t.Errorf("unexpected output after adding an item")
		return
	}

	if itemId.Id == "" {
		t.Error("unexpected item id")
		return
	}
	t.Logf("itemId.Id: %v\n", itemId.Id)

	// Edit item
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	bodyReader = strings.NewReader(string(bodyJson))
	c.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/item/%s", itemId.Id), bodyReader)
	c.Request.Header.Add("Authorization", bearer)
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: itemId.Id,
	})

	app.GetItemHandler(c)

	res = w.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Can't edit item")
	}
}
