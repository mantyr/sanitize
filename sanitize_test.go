package sanitize

import (
    "testing"
    "os"
    "io/ioutil"
    "fmt"
)

func TestClearAll(t *testing.T) {
    sani := New()
    sani.LoadFile("./testdata/test1.html")
    if sani.Error != nil {
        t.Errorf("Error open file, %q", "./testdata/test1.html")
    }
    sani.SetBaseHost("http://example.com/")
    sani.SetAudioPreload("none")

    sani.RemoveTags()
    sani.RemoveAttr()

    sani.RemoveParam()
    sani.FilterIframe()
    sani.FilterObject()
    sani.FilterEmbed()

    sani.FilterA()
    sani.FilterImg()

    sani.RemoveDubleBr()    // убираем часто повторяющиеся <br>
    sani.RemoveEmptyTags()  // убираем пустые теги в самом конце так как они могли появиться во время фильтрации
    sani.RemoveDubleBr()    // убираем часто повторяющиеся <br>

    var err error
    var html_1, html_2 string

    html_1, err = sani.Dom.Find("body").Html()
    if err != nil {
        t.Errorf("Error generate HTML, %q", err)
    }

    html_2, err = GetFile("./testdata/test1_result.html")
    if err != nil {
        t.Errorf("Error open HTML ./testdata/test1_result.html, %q", err)
    }

    if html_1 == "" {
        t.Errorf("Error generate HTML, zero value")
    }
    if html_1 != html_2 {
        t.Errorf("Error generate, HTML != test1_result.html") // переписать, добавить вывод diff файлов по строчно
        fmt.Println("======== generate file =======")
        fmt.Println(html_1)
        fmt.Println("======== expected file =======")
        fmt.Println(html_2)
    }

}
/*
func TestClearAll2(t *testing.T) {
t.Errorf("test")

    sani := New()
    sani.LoadFile("./testdata/test2.html")
    if sani.Error != nil {
        t.Errorf("Error open file, %q", "./testdata/test1.html")
    }
    sani.SetBaseHost("http://example.com/")
    sani.SetAudioPreload("none")

    sani.RemoveTags()
    sani.RemoveAttr()

    sani.RemoveParam()
    sani.FilterIframe()
    sani.FilterObject()
    sani.FilterEmbed()

    sani.FilterA()
    sani.FilterImg()

    sani.RemoveDubleBr()    // убираем часто повторяющиеся <br>
    sani.RemoveEmptyTags()  // убираем пустые теги в самом конце так как они могли появиться во время фильтрации
    sani.RemoveDubleBr()    // убираем часто повторяющиеся <br>
    sani.RemoveEmptyTags()  // убираем пустые теги в самом конце так как они могли появиться во время фильтрации
    var err error
    var html_1, html_2 string

    html_1, err = sani.Dom.Find("body").Html()
    if err != nil {
        t.Errorf("Error generate HTML, %q", err)
    }

    html_2, err = GetFile("./testdata/test1_result.html")
    if err != nil {
        t.Errorf("Error open HTML ./testdata/test1_result.html, %q", err)
    }

    if html_1 == "" {
        t.Errorf("Error generate HTML, zero value")
    }
    if html_1 != html_2 {
        t.Errorf("Error generate, HTML != test1_result.html") // переписать, добавить вывод diff файлов по строчно
        fmt.Println("======== generate file =======")
        fmt.Println(html_1)
//        fmt.Println("======== expected file =======")
//        fmt.Println(html_2)
    }
}
*/

func GetFile(address string) (st string, err error) {
    var file *os.File

    file, err = os.Open(address)
    if err != nil {
        return
    }
    st2, err := ioutil.ReadAll(file)
    st = string(st2)
    return
}