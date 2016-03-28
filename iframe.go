package sanitize

import (
    "github.com/mantyr/goquery"
    "github.com/mantyr/runner"
    "net/url"
)

// удалем отдельно стоящие param
func (sani *Sani) RemoveParam() *Sani {
    sani.Dom.Find("param").Each(func(i int, s *goquery.Selection) {
        if s.Closest("object").Length() == 0 {
            s.Remove()
            sani.is_revalue = true
        }
    })
    return sani
}

func (sani *Sani) FilterIframe(params ...string) *Sani {
    access_hosts := runner.Access_iframe_hosts
    if len(params) > 0 {
        access_hosts = params
    }
    sani.Dom.Find("iframe").Each(func(i int, s *goquery.Selection) {
        s_src := s.GetObjectSrc()

        if len(s_src) == 0 {
            s.Remove() // удаляем пустой iframe
            sani.is_revalue = true
            return
        }

        s_url, err := url.Parse(s_src)
        if err != nil {
            s.Remove() // удаляем iframe если внутри непонятный url
            sani.is_revalue = true
            return
        }

        if s_url.Scheme != "http" && s_url.Scheme != "https" && s_url.Scheme != ""  {
            s.Remove()
            sani.is_revalue = true
            return
        }

        if !runner.AccessHost(s_url.Host, access_hosts) {
            s.Remove()
            sani.is_revalue = true
            return
        }
    })

    return sani
}

func (sani *Sani) FilterObject(params ...string) *Sani {
    access_hosts := runner.Access_iframe_hosts
    if len(params) > 0 {
        access_hosts = params
    }
    sani.Dom.Find("object").Each(func(i int, s *goquery.Selection) {
        sani.is_revalue = true

        s_src  := s.GetObjectSrc()
        s_type := s.GetMimeType()

        if len(s_src) == 0 || s_type != flash_mime_type {
            s.Remove() // удаляем пустой object или object в котором неизвестный mime-type
            return
        }

        s_url, err := url.Parse(s_src)
        if err != nil || s_url.Host == "" {
            s.Remove() // удаляем object и все вложенные элементы если внутри непонятный url
            return
        }

        if s_url.Scheme != "http" && s_url.Scheme != "https" && s_url.Scheme != "" {
            s.Remove() // удаляем если неизвестная схема, если пустая - оставляем
            return
        }

        if !runner.AccessHost(s_url.Host, access_hosts) {
            s.Remove() // удаляем если домен не входит в белый список
            return
        }

        // если object всё таки оставили - патчим вложенные param и embed
        s.Find("param[name=movie]").SetAttr("value", s_src)
        s.Find("embed").SetAttr("src", s_src)
        s.SetAttr("data", s_src)                              // так как мы не знаем где взяли адрес патчим и object
    })
    return sani
}

// ищем отдельно стоящие embed
func (sani *Sani) FilterEmbed(params ...string) *Sani {
    access_hosts := runner.Access_iframe_hosts
    if len(params) > 0 {
        access_hosts = params
    }
    sani.Dom.Find("embed").Each(func(i int, s *goquery.Selection) {
        if s.Closest("object").Length() == 0 {
            s_src  := s.GetObjectSrc()
            s_type := s.GetMimeType()

            if len(s_src) == 0 || s_type != flash_mime_type {
                s.Remove() // удаляем пустой embed или embed в котором неизвестный mime-type
                sani.is_revalue = true
                return
            }

            s_url, err := url.Parse(s_src)
            if err != nil || s_url.Host == "" {
                s.Remove() // удаляем embed и все вложенные элементы если внутри непонятный url
                sani.is_revalue = true
                return
            }

            if s_url.Scheme != "http" && s_url.Scheme != "https" && s_url.Scheme != "" {
                s.Remove() // удаляем если неизвестная схема, если пустая - оставляем
                sani.is_revalue = true
                return
            }

            if !runner.AccessHost(s_url.Host, access_hosts) {
                s.Remove() // удаляем если домен не входит в белый список
                sani.is_revalue = true
                return
            }

            // если object всё таки оставили - патчим вложенные param
            if s.Find("param[name=movie]").SetAttr("value", s_src).Length() > 0 {
                sani.is_revalue = true
            }
        }
    })
    return sani
}