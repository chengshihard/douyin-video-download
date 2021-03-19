package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ResponseData struct {
	StatusCode int `json:"status_code""`
	AwemeList []struct{
		AwemeId string `json:"aweme_id"`
		Video struct{
			PlayAddr struct{
				UrlList []string `json:"url_list"`
			} `json:"play_addr"`
		} `json:"video"`
	} `json:"aweme_list"`
}

type VideoList struct {
	AwemeId string `json:"aweme_id""`
	Url string `json:"url"`
}

const path = ``

func main() {
	var startTime int64 = 1582387200000

	requestUrl := `https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=MS4wLjABAAAAKbhcZVC80794_3SNxHLzZEglkO5sDKRbZpN0TcbO2GA&count=21&max_cursor={$max_cursor}&aid=1128&_signature=SWImtgAAKTmhuATLVJI21EliJq&dytk=`

	videoList := make([]VideoList,0)

	for startTime < time.Now().UnixNano() / 1e6 {
		timeStr := strconv.FormatInt(startTime,10)
		url := strings.Replace(requestUrl,"{$max_cursor}",timeStr,1)
		respdata := requestClient(url)

		if len(respdata.AwemeList) > 0 {
			for _, v := range respdata.AwemeList {
				lastUrl :=  v.Video.PlayAddr.UrlList[len(v.Video.PlayAddr.UrlList) - 1]

				exist := false

				for _,vv := range videoList {
					if vv.Url == lastUrl{
						exist = true
					}
				}

				if !exist {
					var video VideoList
					video.AwemeId = v.AwemeId
					video.Url = v.Video.PlayAddr.UrlList[len(v.Video.PlayAddr.UrlList) - 1]
					videoList = append(videoList, video)
				}
			}
		}

		fmt.Printf("时间：%d，视频数：%d \n",startTime, len(videoList))

		startTime = startTime + 24 * 60 * 60 * 1000
	}


	for _, v := range videoList{
		videoDownload(v)
	}
}

func requestClient(url string) ResponseData {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var respData ResponseData
	err = json.Unmarshal(body, &respData)

	if err != nil {
		fmt.Println(err)
	}

	return respData
}

func videoDownload(video VideoList){
	res, err := http.Get(video.Url)
	if err != nil {
		panic(err)
	}

	fmt.Println(path + video.AwemeId + ".mp4")

	f, err := os.Create(path + video.AwemeId + ".mp4")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

