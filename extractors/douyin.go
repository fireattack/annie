package extractors

import (
	"encoding/json"

	"github.com/iawia002/annie/downloader"
	"github.com/iawia002/annie/request"
	"github.com/iawia002/annie/utils"
)

type douyinVideoURLData struct {
	URLList []string `json:"url_list"`
}

type douyinVideoData struct {
	PlayAddr     douyinVideoURLData `json:"play_addr"`
	RealPlayAddr string             `json:"real_play_addr"`
}

type douyinData struct {
	Video douyinVideoData `json:"video"`
	Desc  string          `json:"desc"`
}

// Douyin download function
func Douyin(url string) downloader.VideoData {
	html := request.Get(url)
	vData := utils.Match1(`var data = \[(.*?)\];`, html)[1]
	var dataDict douyinData
	json.Unmarshal([]byte(vData), &dataDict)

	size := request.Size(dataDict.Video.RealPlayAddr, url)
	urlData := downloader.URLData{
		URL:  dataDict.Video.RealPlayAddr,
		Size: size,
	}
	data := downloader.VideoData{
		Site:  "抖音 douyin.com",
		Title: utils.FileName(dataDict.Desc),
		Ext:   "mp4",
		URLs:  []downloader.URLData{urlData},
		Size:  size,
	}
	data.Download(url)
	return data
}