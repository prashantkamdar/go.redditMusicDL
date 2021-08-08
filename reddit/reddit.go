package reddit

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetToken() int {
	data := url.Values{}
	data.Set("grant_type", "")
	data.Set("username", "")
	data.Set("password", "")

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Authorization", "Basic ")
	req.Header.Set("User-Agent", "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	sb := string(body)
	log.Print(sb)
	return 0
}
