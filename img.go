package sanitize

import (
    "github.com/mantyr/goquery"
    "github.com/mantyr/runner"
    "strings"
    "net/url"
)

func (sani *Sani) SetBaseHost(base_host string) *Sani {
    sani.base_host, _ = url.Parse(base_host)
    return sani
}

// чистим картинки, при надобности закачиваем
func (sani *Sani) FilterImg(params ...func(f_sani *Sani, s *goquery.Selection, url *url.URL)) *Sani {
    sani.Dom.Find("img").Each(func(i int, s *goquery.Selection) {
        sani.is_revalue = true

        s_href := runner.Trim(s.AttrOr("src", ""))
        if len(s_href) ==  0 {
            s.Remove()
            return
        }

        s_url, err := url.Parse(s_href)
        if err != nil {
            s.Remove()
            return
        }

        if !runner.InSlice(s_url.Scheme, []string{"http", "https", ""}) {
            s.Remove()
            return
        }

        if s_url.Scheme == "" {
            s_url.Scheme = sani.base_host.Scheme
        }

        if s_url.Host == "" {
            s_url.Host = sani.base_host.Host
            s_url.Path = "/"+strings.TrimLeft(s_url.Path, "./")
        }

        if len(runner.Trim(s_url.Path)) == 0 {
            s.Remove()
            return
        }
        s.SetAttr("src", s_url.String())

        if len(params) > 0 {
            for _, f := range params {
                f(sani, s, s_url)
            }
        }
    })
    return sani
}