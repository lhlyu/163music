package music163

import (
	"encoding/json"
	"testing"
)

func TestParse1(t *testing.T) {
	s := `分享乐小桃的单曲《音你闪耀（王者荣耀孙尚香皮肤同名主题曲）》: https://y.music.163.com/m/song?id=1969809433&uct=JITdEtoueT1leXz61yr0AQ%3D%3D&dlt=0846&app_version=8.8.31&sc=wmv&tn= (来自@网易云音乐)`
	t.Log(extractedLink(s))
	data, err := Parse(s)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(toPrettierJson(data))
}

func TestParse2(t *testing.T) {
	s := `分享Seto的单曲《愛？》https://y.music.163.com/m/song?app_version=8.8.31&id=1965902593&uct=xXUJLXWiMjHXhRV8QVZH8w%3D%3D&dlt=0846 (@网易云音乐)`
	t.Log(extractedLink(s))
	data, err := Parse(s)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(toPrettierJson(data))
}

func TestParse3(t *testing.T) {
	s := `https://music.163.com/song?id=454698657&userid=258842242`
	t.Log(extractedLink(s))
	data, err := Parse(s)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(toPrettierJson(data))
}

// 新版本
func TestParse4(t *testing.T) {
	s := `分享Lil Peep/KirbLaGoop的单曲《red drop shawty》: http://163cn.tv/MrEle0 (来自@网易云音乐)`
	t.Log(extractedLink(s))
	data, err := Parse(s)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(toPrettierJson(data))
}

func toPrettierJson(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
