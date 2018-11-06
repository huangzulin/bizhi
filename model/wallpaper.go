package model

type Wallpaper struct {
	Errno   string `json:"errno"`
	Errmsg  string `json:"errmsg"`
	Consume string `json:"consume"`
	Total   string `json:"total"`
	Data    []struct {
		ID            string        `json:"id"`
		ClassID       string        `json:"class_id"`
		Resolution    string        `json:"resolution"`
		URLMobile     string        `json:"url_mobile"`
		URL           string        `json:"url"`
		URLThumb      string        `json:"url_thumb"`
		URLMid        string        `json:"url_mid"`
		DownloadTimes string        `json:"download_times"`
		Imgcut        string        `json:"imgcut"`
		Tag           string        `json:"tag"`
		CreateTime    string        `json:"create_time"`
		UpdateTime    string        `json:"update_time"`
		AdID          string        `json:"ad_id"`
		AdImg         string        `json:"ad_img"`
		AdPos         string        `json:"ad_pos"`
		AdURL         string        `json:"ad_url"`
		Ext1          string        `json:"ext_1"`
		Ext2          string        `json:"ext_2"`
		Utag          string        `json:"utag"`
		Tempdata      string        `json:"tempdata"`
		Rdata         []interface{} `json:"rdata"`
		Img1600900    string        `json:"img_1600_900"`
		Img1440900    string        `json:"img_1440_900"`
		Img1366768    string        `json:"img_1366_768"`
		Img1280800    string        `json:"img_1280_800"`
		Img12801024   string        `json:"img_1280_1024"`
		Img1024768    string        `json:"img_1024_768"`
	} `json:"data"`
}
