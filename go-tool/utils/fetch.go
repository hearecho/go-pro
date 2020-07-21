package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserInfo struct {
	Name      string
	Mid       string
	VideoNum  float64
	Fans      float64
	Level     float64
	PlayNum   string
	Likes     string
}

func GetUserInfo(name string) []UserInfo {
	url := "https://api.bilibili.com/x/web-interface/search/type?context&search_type=bili_user&page=1&order" +
		"&keyword=" + name + "&category_id&user_type&order_sort&changing=mid&__refresh__=true&_extra&highlight=1&single_column=0&jsonp=jsonp"
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
		"Referer":    "https://search.bilibili.com/all?keyword=" + name,
	}
	var users []UserInfo
	body, err := fetch(url, headers)
	if err != nil {
		return users
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("解析json字符串出错")
		return users
	}
	data := res["data"].(map[string]interface{})
	pre_users := data["result"].([]interface{})
	for _, v := range pre_users {
		u := UserInfo{}
		user := v.(map[string]interface{})
		u.Name = user["uname"].(string)
		u.Mid = strconv.FormatFloat(user["mid"].(float64), 'f', -1, 64)
		u.VideoNum = user["videos"].(float64)
		u.Fans = user["fans"].(float64)
		u.Level = user["level"].(float64)
		users = append(users,u)
	}
	return users
}
func fetch(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if resp == nil {
		log.Printf("crawel error! url:%s\n", url)
		return nil, err
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("crawel error! url:%s\t status_code:%d\n", url, resp.StatusCode)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
