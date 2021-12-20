package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shelik/vk-go-back/models"
)

const api = `https://api.vk.com/method/`
const token = ""

type js struct {
	Error    vkError         `json:"error"`
	Response json.RawMessage `json:"response"`
}

type vkError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_msg"`
}

// type Api struct {
// 	Token string
// }

func (r *HttpRepo) request(method string, params map[string]string) (result json.RawMessage, e error) {

	request := api + method + "?"
	for key, value := range params {
		request += key + "=" + value + "&"
	}
	request += "access_token=" + r.Token

	fmt.Println(request)
	response, e := http.Get(request)
	if e != nil {
		return nil, e
	}
	defer response.Body.Close()

	result, e = ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, e
	}

	var j js
	if e := json.Unmarshal(result, &j); e != nil {
		return nil, e
	}

	if j.Error.Code != 0 {
		return nil, errors.New(fmt.Sprint("vk: ", j.Error.Code, ", \"", j.Error.Message, "\""))
	}

	return j.Response, nil
}

// RepoMock ...
type HttpRepo struct {
	Token string
}

// NewRepo ...
func NewRepo(token string) *HttpRepo {
	return &HttpRepo{
		Token: token,
	}
}

// Close ...
func (r *HttpRepo) Close() error {
	return nil
}

// Translate ...
func (r *HttpRepo) GetGalleries(ownerID string) []models.Gallery {
	galleries := []models.Gallery{}
	params := make(map[string]string)
	params["owner_id"] = ownerID
	params["v"] = "5.131"

	result, e := r.request("photos.getAlbums", params)
	fmt.Println(r.Token)
	if e != nil {
		log.Println(e)
	}

	// Remove [] from json
	result = result[1 : len(result)-1]

	json.Unmarshal(result, &galleries)

	fmt.Println(galleries)

	return galleries
}
