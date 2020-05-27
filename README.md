# Local GAE Support

GAE 2nd genでは、1st genのSDKに含まれていた `dev_appserver` が使えなくなってしまったため、
[app.yaml](https://cloud.google.com/appengine/docs/standard/go/config/appref)の
`static_dir` や `static_files` に設定したファイルを開発時にHTTPでアクセスするためには
開発用のハンドラを実装しなければなりません。

`dev_appserver` which provides static file access in development environoment
was included in GAE 1st gen SDK but is not included in GAE 2nd gen. So we have to
implement handlers to provide them.

`local_gae_support`は、以下のようにハンドラをラップすることで `static_dir` や
`static_files` の設定に基づいて、開発環境でのハンドラを定義します。

`local_gae_support` provides handlers for `static_dir` or `static_files` in
app.yaml like this:

```golang
import "github.com/akm/local_gae_support"
```

```golang
var err error
handler, err = localgaesupport.Static("app.yaml", handler)
if err != nil {
	panic(err)
}
```
