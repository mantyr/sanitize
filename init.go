package sanitize

import (
    "github.com/mantyr/goquery"
    "os"
    "bytes"
)

func New() (s *Sani) {
    s = new(Sani)
    s.is_revalue = false
    return
}

func (sani *Sani) LoadFile(address string) error {
    var file *os.File
    file, sani.Error = os.Open(address)
    if sani.Error != nil {
        return sani.Error
    }
    defer file.Close()

    var doc *goquery.Document
    doc, sani.Error = goquery.NewDocumentFromReader(file)
    sani.Dom = doc.Clone()
    return sani.Error
}

func (sani *Sani) SetSelection(sel *goquery.Selection) *Sani {
    sani.Dom = sel
    return sani
}

func (sani *Sani) LoadString(text string) error {
    reader := bytes.NewReader([]byte(text))

    var doc *goquery.Document
    doc, sani.Error = goquery.NewDocumentFromReader(reader)
    sani.Dom = doc.Clone()
    return sani.Error
}

// сработал ли хотя бы один из фильтров
func (sani *Sani) IsEdit() bool {
    return sani.is_revalue
}

const (
    flash_mime_type = "application/x-shockwave-flash"
)

var (
    Remove_Tags = "script, style, input, select, textarea, form, frame, base, basefont, applet, map, area, aside"
    Remove_Attr = "id class style srcdoc seamless onclick onblur onchange ondblclick onfocus onkeydown onkeypress onkeyup onload onmousedown onmousemove onmouseout onmouseover onmouseup onreset onselect onsubmit onunload"
    Remove_Empty_Tags = "p, i, a, s, b, div, span, sup, pub"
)

