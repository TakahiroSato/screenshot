# screenshot

メインウィンドウの指定矩形のスクリーンショットを撮る CLI ツール

## 使い方

```
$ screenshot.exe [flags] {矩形左上X座標} {矩形左上Y座標} {矩形右下X座標} {矩形右下Y座標} [{出力パス}]
```

出力パスに PNG 形式画像が出力されます。

出力パス未指定の場合は実行時のディレクトリに`screenshot.png`が出力されます。

## フラグ

- -binary={true | false} : 二値化出力するかどうか(default は false)
- -th={0~255} : 二値化閾値(default は 127)
- -reverse={true | false} : 二値化時の白黒反転(default は false 画素が閾値以上の時白=255 になり、閾値未満は黒=0)
