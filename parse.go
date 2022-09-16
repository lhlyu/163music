package music163

import (
	"fmt"
	"net/url"
	"strings"
)

// Parse 解析分享链接
func Parse(shareUrl string) (*musicInfo, error) {
	link := extractedLink(shareUrl)
	if link == "" {
		return nil, ErrorNoFoundShareLink
	}
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	logo := fmt.Sprintf(favicon_format, u.Scheme, u.Host)

	data := &musicInfo{
		Url:  link,
		Logo: logo,
	}

	err = getMusicTitle(data)

	if err != nil {
		return data, err
	}

	musicId := u.Query().Get("id")

	if musicId == "" {
		return data, nil
	}

	err = getSongDetail(data, musicId)
	if err != nil {
		return data, err
	}
	err = getSongUrl(data, musicId)
	if err != nil {
		return data, err
	}

	return data, nil
}

// 提取分享链接
func extractedLink(shareUrl string) string {
	link := musicShareLinkReg.FindString(shareUrl)
	if link == "" {
		return ""
	}
	return link
}

// 获取标题
func getMusicTitle(data *musicInfo) error {
	val, err := doGet(data.Url)
	if err != nil {
		return err
	}
	titles := musicTitleRegexp.FindStringSubmatch(val)
	if len(titles) > 1 {
		data.Title = titles[1]
	}
	return nil
}

// 获取歌曲详情
func getSongDetail(data *musicInfo, musicId string) error {
	d := fmt.Sprintf(song_detail_format, musicId, musicId)
	uv := url.Values{}
	uv.Set("params", aesEncrypt(aesEncrypt(d, key1), key2))
	uv.Set("encSecKey", enc)
	r, err := doPost(api_get_song_detail, data.Url, strings.NewReader(uv.Encode()))
	if err != nil {
		return err
	}
	if r == nil {
		return nil
	}
	song := r.Get("songs.0")
	data.Name = song.Get("name").String()
	data.Album = song.Get("al.name").String()
	data.Duration = song.Get("dt").Float()
	data.Cover = song.Get("al.picUrl").String() + "?param=200y200"
	artists := make([]string, 0)
	for _, result := range song.Get("ar").Array() {
		artists = append(artists, result.Get("name").String())
	}
	data.Artist = strings.Join(artists, "/")
	return nil
}

func getSongUrl(data *musicInfo, musicId string) error {
	d := fmt.Sprintf(song_url_format, musicId)
	uv := url.Values{}
	uv.Set("params", aesEncrypt(aesEncrypt(d, key1), key2))
	uv.Set("encSecKey", enc)
	r, err := doPost(api_get_song_url, referer_url, strings.NewReader(uv.Encode()))
	if err != nil {
		return err
	}
	if r == nil {
		return nil
	}
	dt := r.Get("data.0")
	data.Music = dt.Get("url").String()
	return nil
}
