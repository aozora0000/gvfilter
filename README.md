# gvfilter
go言語製、セマンティックバージョンフィルター

標準入力からパイプさせて利用します。

```shell
cat versions.txt | gvfilter "~5.3" | gvfilter "latest"
```

https://github.com/Masterminds/semver　の比較に準拠します。

比較記法、ハイフン記法、ワイルドカード記法、チルダ記法、キャレット記法が使えます。

また```latest```を指定する事で、最新のバージョンが取得出来ます。