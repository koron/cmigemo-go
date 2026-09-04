# koron-go/cmigemo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/cmigemo)](https://pkg.go.dev/github.com/koron-go/cmigemo)
[![Actions/Go](https://github.com/koron-go/cmigemo/actions/workflows/go.yml/badge.svg)](https://github.com/koron-go/cmigemo/actions/workflows/go.yml)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/koron-go/cmigemo)

[C/Migemo][cmigemo] wrapper for Go using [PureGo][purego]

[cmigemo]:https://github.com/koron/cmigemo
[purego]:https://github.com/ebitengine/purego

## Example code

``go
import "github.com/koron/cmigemo-go"

mo, err := cmigemo.Open("/usr/local/share/cmigemo/utf-8/migemo-dict)
if err != nil {
	panic(err.Error())
}

pat := mo.Query("aka")
println(pat)

rx, err := mo.Regexp("aka")
if err != nil {
	panic(err.Error())
}
m := rx.FindAllStringIndex("...")
``
