package sanitize

import (
    "github.com/mantyr/goquery"
    "github.com/mantyr/runner"
    "strings"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
)

// удаляем лишние теги, такие как скрипты, стили, формы и прочее
func (sani *Sani) RemoveTags(params ...string) *Sani {
    tags := Remove_Tags
    if len(params) > 0 {
        tags = strings.Join(params, ", ")
    }

    if sani.Dom.Find(tags).Remove().Length() > 0 {
        sani.is_revalue = true
    }

    return sani
}

// убираем атрибуты которые могут вызвать скрипты, а так же стили, классы и идентификаторы, seamless - выставляемт расширения для iframe, srcdoc - содержит html
func (sani *Sani) RemoveAttr(params ...string) *Sani {
    attr := Remove_Attr
    if len(params) > 0 {
        attr = strings.Join(params, ", ")
    }

    if sani.Dom.FindRemoveAttr(attr).Length() > 0 {
        sani.is_revalue = true
    }
    return sani
}

// удаляем пустые теги
func (sani *Sani) RemoveEmptyTags(params ...string) *Sani {
    tags := Remove_Empty_Tags
    if len(params) > 0 {
        tags = strings.Join(params, ", ")
    }
    sani.Dom.Find(tags).Each(func(i int, s *goquery.Selection) {
        line, err := s.Html()
        line = runner.Trim(line)

        if err != nil || line == "" || line == "<br/>" || line == "<br>" {
            s.Remove()
            sani.is_revalue = true
        }
    })

    return sani
}

func (sani *Sani) RemoveDubleBr() *Sani {
    sani.Dom.Find("br").Each(func(i int, s *goquery.Selection){
        if len(s.Nodes) == 0 {
            return
        }

        next := s.Nodes[0].NextSibling
        for {
            if next == nil {
                break
            }
            if next.Type == html.TextNode && runner.Trim(next.Data) != "" {  // если текст
                break
            }
            if next.Type == html.ElementNode && next.DataAtom != atom.Br {   // если дальше идёт тег не <br>
                break
            }
            if next.Type == html.ElementNode && next.DataAtom == atom.Br {   // если дальше идёт тег <br> то текущий убираем, следующий проверим при следующей итерации
                s.Remove()
                break
            }
            next = next.NextSibling
        }
    })
    return sani
}