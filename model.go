package music163

type musicInfo struct {
	// 分享原始链接
	Url string `json:"url,omitempty"`
	// 网易云音乐logo
	Logo string `json:"logo,omitempty"`
	// 标题
	Title string `json:"title,omitempty"`
	// 音乐下载地址
	Music string `json:"music,omitempty"`
	// 名字
	Name string `json:"name,omitempty"`
	// 作者
	Artist string `json:"artist,omitempty"`
	// 专辑
	Album string `json:"album,omitempty"`
	// 封面
	Cover string `json:"cover,omitempty"`
	// 时长(毫秒)
	Duration float64 `json:"duration,omitempty"`
}
