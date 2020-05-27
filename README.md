# Local GAE Support

GAE 2nd genでは、1st genのSDKに含まれていた `dev_appserver` が使えなくなってしまったため、
[app.yaml](https://cloud.google.com/appengine/docs/standard/go/config/appref)の
`static_dir` や `static_files` に設定したファイルを開発時にHTTPでアクセスするためには
開発用のハンドラを実装しなければなりません。

`local_gae_support`は、以下のようにハンドラをラップすることで `static_dir` や
`static_files` の設定に基づいて、開発環境でのハンドラを定義します。

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
