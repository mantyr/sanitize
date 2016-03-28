package sanitize

import (
    "github.com/mantyr/goquery"
    "github.com/mantyr/runner"
    "strings"
    "net/url"
)

// <audio preload="none | metadata | auto"> - выставить в none
func (sani *Sani) SetAudioPreload(params ...string) *Sani {
    status := "none"
    if len(params) > 0 {
        status = params[0]
    }
    if sani.Dom.Find("audio").SetAttr("preload", status).Length() > 0 {
        sani.is_revalue = true
    }
    return sani
}

// чистим ссылки
func (sani *Sani) FilterA(params ...func(sani *Sani, s *goquery.Selection, url *url.URL)) *Sani {
    sani.Dom.Find("a").Each(func(i int, s *goquery.Selection) {
        sani.is_revalue = true

        s_href := runner.Trim(s.AttrOr("href", ""))
        if len(s_href) ==  0 {
            if sani.ignore_a_empty_href {
                return
            } else {
                s.Remove()
                sani.is_revalue = true
                return
            }
        }

        s_url, err := url.Parse(s_href)
        if err != nil {
            if sani.ignore_a_error_href {
                s.SetAttr("href", "") // удалем адрес ссылки если не можем его распознать
            } else {
                s.Remove()            // удаляем ссылку если не можем её распознать
            }
            return
        }

        if !runner.InSlice(s_url.Scheme, []string{"http", "https", "ftp", "sftp", ""}) {
            if sani.ignore_a_error_scheme {
                s.SetAttr("href", "") // удалем адрес ссылки если схема не известна
            } else {
                s.Remove()            // удаляем ссылку если не можем её распознать
            }
            return
        }

        if s_url.Host == "" {
            s_url.Host = sani.base_host
            s_url.Path = "/"+strings.TrimLeft(s_url.Path, "./")
            s.SetAttr("href", s_url.String())
        }

        if len(params) > 0 {
            for _, f := range params {
                f(sani, s, s_url)
            }
        }

    })
    return sani
}
