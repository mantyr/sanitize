# Golang HTML Sanitize - alfa-alfa version

[![Build Status](https://travis-ci.org/mantyr/sanitize.svg?branch=master)](https://travis-ci.org/mantyr/sanitize) [![GoDoc](https://godoc.org/github.com/mantyr/sanitize?status.png)](http://godoc.org/github.com/mantyr/sanitize) [![Software License](https://img.shields.io/badge/license-The%20Not%20Free%20License,%20Commercial%20License-brightgreen.svg)](LICENSE.md)

This don't stable version

## Installation

    $ go get github.com/mantyr/sanitize
    $ go get github.com/mantyr/goquery
    $ go get github.com/mantyr/runner

## Example

```GO
package main

import (
    "github.com/mantyr/sanitize"
)

func main() {
    sani := sanitize.New()
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

    sani.RemoveEmptyTags()

    html, err := sani.Dom.Find("body").Html()
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
