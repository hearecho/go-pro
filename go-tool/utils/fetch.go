package utils

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type UserInfo struct {
	Name      string
	Mid       string
	VideoNum  float64
	Following float64
	Follower  float64
	PlayNum   string
	Likes     string
}

func GetUserInfo(uid string) UserInfo {
	wg := &sync.WaitGroup{}
	u := &UserInfo{}
	wg.Add(4)
	url1 := "https://api.bilibili.com/x/space/acc/info?mid=" + uid + "&jsonp=jsonp"
	url2 := "https://api.bilibili.com/x/relation/stat?vmid=" + uid
	url3 := "https://api.bilibili.com/x/space/upstat?mid=" + uid
	url4 := "https://api.bilibili.com/x/space/navnum?mid=" + uid
	go fetch(url1, wg, u)
	go fetch(url2, wg, u)
	go fetch(url3, wg, u)
	go fetch(url4, wg, u)
	if err != nil {
		log.Println("程序出错！")
	}
	wg.Wait()
	return *u
}

func fetch(url string, wg *sync.WaitGroup, u *UserInfo) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		log.Println(err)
	}
	for k, v := range viper.GetStringMap("headers") {
		req.Header.Set(k, v.(string))
	}
	resp, err := client.Do(req)
	if resp == nil {
		log.Printf("crawel error! url:%s\n", url)
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("crawel error! url:%s\t status_code:%d\n", url, resp.StatusCode)
	}
	defer resp.Body.Close()
	defer wg.Done()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("解析json字符串出错")
	}
	//存储数据
	data := res["data"].(map[string]interface{})
	if _, ok := data["video"]; ok {
		u.VideoNum = data["video"].(float64)
	}
	if _, ok := data["archive"]; ok {
		archive := data["archive"].(map[string]interface{})
		u.PlayNum = strconv.FormatFloat(archive["view"].(float64),'f',-1,64)
	}
	if _, ok := data["likes"]; ok {
		u.Likes = strconv.FormatFloat(data["likes"].(float64),'f',-1,64)
	}
	if _, ok := data["following"]; ok {
		u.Following = data["following"].(float64)
	}
	if _, ok := data["follower"]; ok {
		u.Follower = data["follower"].(float64)
	}
	if _, ok := data["mid"]; ok {
		u.Mid = strconv.FormatFloat(data["mid"].(float64), 'f', -1, 64)
	}
	if _, ok := data["name"]; ok {
		u.Name = data["name"].(string)
	}

}
