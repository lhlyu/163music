package music163

import (
	"errors"
	"regexp"
)

const (
	key1 = "0CoJUm6Qyw8W8jud"
	key2 = "tDlm2QZLSRUIuwkj"
	enc  = "6eb1b01680f88137478197ed98ad62eb89a018b7a4bf6defc787529e6edc3ea096290a682fe60b08e2afd93d7eb040bf0cc43615b7d06faeef9fde42a2c136bc92842898f56b9059eb17babcbe1a7520714780a0ed666c7a83514973359e13eb3b33a9e05c7f2cd91bfc4dd0a224bae8a329434524415163301ca7d06e854622"
	iv   = "0102030405060708"
)

const (
	music_http_regexp  = "https:\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?"
	music_title_regexp = "<title>(.*)</title>"
)

const (
	favicon_format = "%s://%s/favicon.ico"
	// 歌曲详情参数
	song_detail_format = "{\"id\":\"%s\",\"c\":\"[{\\\"id\\\":\\\"%s\\\"}]\",\"csrf_token\":\"\"}"
	song_url_format    = "{\"ids\":\"[%s]\",\"level\":\"standard\",\"encodeType\":\"aac\",\"csrf_token\":\"\"}"
)

const (
	api_get_song_detail = "https://music.163.com/weapi/v3/song/detail?csrf_token="
	api_get_song_url    = "https://music.163.com/weapi/song/enhance/player/url/v1?csrf_token="

	referer_url = "https://music.163.com/"
)

var (
	musicShareLinkReg = regexp.MustCompile(music_http_regexp)
	musicTitleRegexp  = regexp.MustCompile(music_title_regexp)
)

var (
	ErrorNoFoundShareLink = errors.New("no found share link")
)
