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

**UI(Presentation)層 : /presenter**  
ユーザーからの入力を受け付け入力をサニタイジング、ApplicationService層に渡す役割を持つ  
また、ApplicationService層からの出力を受け取り、外界とのやり取りに適した形式に変換を行う  
  
**Infrastructure層 : /infrastructure**  
データの永続化,メッセージ送信などの技術的機能をここで定義  
/domain/service/xxx_repository.goに記載するinterfaceの実装をここで行う  

**ApplicationService層 : /usecase**    
アプリケーション固有のロジックを定義し、処理順序やドメイン層とのデータの流れをここで管理  
  
**DomainService層 : /domain/service**  
エンティティや値オブジェクトの責務ではないドメインモデルのロジックをここで管理  
/domain/repositoryにはDIPのためのinterfaceを定義  
  
**DomainModel層 : /domain/model**  
エンティティ、値オブジェクト、集約といった「ドメインオブジェクト」をここで管理  
  
**/interactor**  
DIコンテナの役割  
最終的な依存の解決はinteractorが行う  

__参考__　　
![onion](https://user-images.githubusercontent.com/28241735/120328338-b0a8f080-c325-11eb-965b-4c355a03e983.jpeg)  
[https://qiita.com/nanamen/items/f37d1047368929e377fd](https://qiita.com/nanamen/items/f37d1047368929e377fd)

  

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
  
  