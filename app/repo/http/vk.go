package http

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

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

type resultGalleries struct {
	Count     int              `json:"count"`
	Galleries []models.Gallery `json:"items"`
}

type resultPhotos struct {
	Count  int            `json:"count"`
	Photos []models.Photo `json:"items"`
}

func getInsecureHTTPClient() *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{
		Transport: tr,
		Timeout:   120 * time.Second,
	}

	// var netTransport = &http.Transport{
	// 	Dial: (&net.Dialer{
	// 		Timeout: 5 * time.Second,
	// 	}).Dial,
	// 	TLSHandshakeTimeout: 5 * time.Second,
	// }
	// netTransport.TLSClientConfig = &tls.Config{}

	// var netClient = &http.Client{
	// 	Timeout:   10 * time.Second,
	// 	Transport: netTransport,
	// }

	// caCert, err := ioutil.ReadFile("cert.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{
	// 			RootCAs: caCertPool,
	// 		},
	// 	},
	// }
	return client
}

func (r *HttpRepo) request(method string, params map[string]string) (result json.RawMessage, e error) {

	// request := api + method + "?"
	// for key, value := range params {
	// 	request += key + "=" + value + "&"
	// }
	// request += "access_token=" + r.Token

	paramValues := url.Values{}

	for key, value := range params {
		paramValues.Add(key, value)
	}
	// paramValues.Add("access_token", r.Token)
	paramValues.Add("lang", "ru")
	paramValues.Add("v", "5.131")

	url := fmt.Sprintf("%s%s?%s", api, method, paramValues.Encode())
	// url := fmt.Sprintf("%s%s", api, method)

	client := getInsecureHTTPClient()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)

		return
	}

	req.Close = true

	response, e := client.Do(req)
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

// HttpRepo ...
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

// GetGalleries ...
func (r *HttpRepo) GetGalleries(ownerID, token string) []models.Gallery {
	res := resultGalleries{}
	params := make(map[string]string)
	params["owner_id"] = ownerID
	params["access_token"] = token

	result, e := r.request("photos.getAlbums", params)
	if e != nil {
		log.Println(e)
	}

	json.Unmarshal(result, &res)

	return res.Galleries
}

// Translate ...
func (r *HttpRepo) GetPhotos(ownerID, token string, galleryIDs []string, count int) []models.Photo {
	var allPhotos []models.Photo

	res := resultPhotos{}
	for _, gID := range galleryIDs {
		params := make(map[string]string)
		params["album_id"] = gID
		params["owner_id"] = ownerID
		params["access_token"] = token
		params["count"] = strconv.Itoa(count)

		result, e := r.request("photos.get", params)
		if e != nil {
			log.Println(e)
		}

		json.Unmarshal(result, &res)

		allPhotos = append(allPhotos, res.Photos...)
	}

	return allPhotos
}
