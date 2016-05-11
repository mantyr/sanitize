package sanitize

import (
    "github.com/mantyr/goquery"
    "net/url"
)

type Sani struct {
    Dom        *goquery.Selection
    Error      error
    is_revalue bool

    base_host  *url.URL           // адрес сайта для подстановки в адреса у которых относительные ссылки, тоже и для картинок

    ignore_a_empty_href   bool   // нужно ли удалять пустые ссылки, по умолчанию false - не нужно
    ignore_a_error_href   bool   // нужно ли удалять битые ссылки, по умолчанию false - не нужно
    ignore_a_error_scheme bool   // нужно ли удалять ссылки с неизвестными Scheme, по умолчанию false - не нужно. Известные Scheme это http, https, ftp, sftp
}


