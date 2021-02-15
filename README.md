# KBC Blog  
#### 僕が通っている学校（河原電子ビジネス専門学校）専用の技術ブログサイトです。
### URL https://kbc-blog.com/  

## アプリ概要
### こちらのアプリのコンセプトは、以下の4点です。

* 学校全体（先生、生徒）の技術力の向上
* 学校の情報など世の中のブログにはない情報を見れる
* ブログの作成経験を得る
* 他学年他クラスとの情報交換

#### 基本的にはQiitaのような記事と質問の投稿、コメント、いいね、通知、フォロー機能があるブログサイトです。
* 記事と質問の作成、編集、削除ができる
* 記事、質問にタグやカテゴリーを付け検索機能
* 投稿数、いいね数、回答数、のランキング機能（1ヶ月ごとに集計）
* 様々なアクションの後に通知機能


## 開発した背景

僕はこのアプリを作成する前に`Ruby on Rails`でInstagramに似たアイテム管理アプリを作成しました。  
初めて作ったアプリでかなり力を入れて作ったのですが誰にも使われませんでした。  
この時いくら力を入れ、自分が作りたいアプリを作っても使われなかったら意味がないことに気づきました。  
この反省点を活かし、次こそは絶対に使ってもらえる需要のあるアプリを作ろうと決めました。   
#### 必ず使ってもらう為に抑えたことは以下の2点です。
* 使用ユーザーの範囲を狭める(全員に使ってもらえるアプリを作るのは難しく、既に素晴らしいものが作られているから)
* 自分があったら便利だと思うものを作る

#### 対策
* 使用ユーザーの範囲を狭める → 学校内のみ
* 便利だと思うもの → 学校内のニッチな情報を素早く取得でき、情報交換できるアプリ

これらを踏まえて作成したのがこのKBC-Blogです。

## 使用画面のイメージ

![Image](https://github.com/DaikiKawaoka/KbcBlog/blob/master/document/image.jpg)

## 使用技術
* フロントエンド
  * vue/cli 4.5.6
  * HTML / CSS / SCSS
*バックエンド
  * Golang 1.15.1
  * echo 4.1.17
* インフラ
  * CircleCi
  * Docker/ docker-compose
  * nginx 1.17.10
  * mysql 5.8
  * AWS(EC2,ALB,ACM,S3,RDS,Route53,VPC,EIP,IAM)
* その他のツール
  * Visual Studio Code
  * Draw.io
 
 ## AWS構成図
 
 ![Image](https://github.com/DaikiKawaoka/KbcBlog/blob/master/document/kbcblogNetwork.jpg)

