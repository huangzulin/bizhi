package main

import (
	"bizhi/model"
	"encoding/json"
	"github.com/redpois0n/wallpaper"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	resp, err := resty.R().Get("http://lcoc.top/bizhi/api.php?cid=9&start=0&count=200")
	if err!=nil {
	}

	wall := model.Wallpaper{}
	json.Unmarshal([]byte(resp.String()), &wall)

	length := len(wall.Data)
	rand.Seed(time.Now().UnixNano())
	n :=  random(0, length-1)
	url := wall.Data[n].URL
	println("设置的壁纸："+url)

	imgResp,err := resty.R().Get(url)
	path,err:=filepath.Abs("wallpaper.png")

	ioutil.WriteFile(path,imgResp.Body(),644)
	path = strings.Replace(path,"\\","/",-1)
	println(path)
	wallpaper.SetWallpaper(path)
	time.Sleep(3 * time.Second)
	os.Remove(path)

}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
