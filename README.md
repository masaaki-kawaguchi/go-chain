# go-chain

## サーバー起動

```shell
go run main.go
```
もしくは  
```shell
docker-compose up
```

## ディレクトリ構成とオニオンアーキテクチャの関係性

```shell
UI(Presentation)層 : /presenter
Infrastructure層 : /infrastructure
ApplicationService層 : /usecase
DomainService層 : /domain/service, /domain/repository
DomainModel層 : /domain/model
```
   
上から下へ  

参考　　
![onion](https://user-images.githubusercontent.com/28241735/120328338-b0a8f080-c325-11eb-965b-4c355a03e983.jpeg)

  

## チェーンの内容
  
- ブロックナンバー  
- ナンス  
- ハッシュ
- 一つ前のハッシュ  
- トランザクション  
  
## トランザクションの内容  
  
- 送り手のアドレス  
- 受け手のアドレス  
- 額 
- メモ
- 署名


## 想定される機能  
  
- マイニング
- アドレス(公開鍵)の妥当性判定

## 実装したいもの

- SHAの二重がけ(出来ればオリジナルでやりたい、最優先ではない)  
- ナンスのランダム探索  
- Mysql接続
- Mysql保存
- 公開鍵と秘密鍵のペア妥当性判定
- 一意の公開鍵の作成(おそらくtimestampを用いてみては)
- 最長チェーン判定
- テスト

## 技術上でやりたいこと

- テスト駆動開発
- ドメイン駆動開発
- クリーンアーキテクチャorオニオンアーキテクチャ
- docker

## TODOリスト  
  
  