# KBC Blog  
#### 僕が通っている学校（河原電子ビジネス専門学校）専用の技術ブログサイトです。
### URL https://kbc-blog.com/  
（kbcblog.comが取れなかったのでしょうがなくハイフンを入れました。）

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

このアプリを作成するにあたって参考にさせてもらった記事
https://qiita.com/nasuB7373/items/0e507abad2017976c407


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
  * Vue.js 2.6.12
  * Vue/cli 4.5.6
  * Element UI 2.13.2
  * HTML / CSS / SCSS
* バックエンド
  * Golang 1.15.1
  * Echo 4.1.17
* インフラ
  * CircleCi
  * Docker/ docker-compose
  * nginx 1.19.6
  * mysql 5.8
  * AWS(EC2,ALB,ACM,S3,RDS,Route53,VPC,EIP,IAM)
* その他のツール
  * Visual Studio Code
  * Draw.io
 
 ## AWS構成図
 
 ![Image](https://github.com/DaikiKawaoka/KbcBlog/blob/master/document/kbcblogNetwork.jpg)
 
 ## 機能一覧

 * ユーザー登録関連
   * 新規登録、プロフィール編集機能
     * 名前、コメント、画像、好きな言語、githubリンク、Webページリンクなど
   * ログイン、ログアウト機能
   * ゲストユーザーログイン機能
     * 投稿やコメントはできない
 
 * 記事、質問関連
   * 新規作成、編集、削除機能
   * 記事へのコメント機能
   * 質問への回答機能
   
 * ランキング関連(1ヶ月毎と全て)
   * 投稿数が多いランキング（投稿数）
   * 投稿のいいね数が多いランキング（人気ユーザー）
   * 質問への回答数が多いランキング（回答数）
   * 質問数が多いランキング（やる気）  
   
   少しでもユーザーにアクションを起こしてもらえるようにランキングを作りました。  
   
 * 無限スクロール機能
   * 記事の一覧ページ
   * おすすめ一覧ページ
   * フォロー、フォロワーリスト
   
 * コメント機能
 * パスワード編集機能
 * 通知機能
 * タグ機能
 * フォロー機能
 * 画像アップロード機能
 * 記事や質問のmarkdown,src表示機能
   * コピぺしたら簡単に同じ記事を作成できる
 
 ※レスポンシブWebデザインはしてないです  
   時間がかかる上に学校のパソコンで見るのが大半だと判断したから


## DB設計
ER図下手すぎました。すみません。
 
  ![Image](https://github.com/DaikiKawaoka/KbcBlog/blob/master/document/kbcblogDB.jpg)
  
### 各テーブルについて

| テーブル名 | 説明        |
| ------ | ------------- |
| users    | ユーザー情報 |
| follows    | フォロー情報 |
| articles    | 技術記事情報 |
| article_comments    | 技術記事へのコメント情報 |
| article_likes    | 技術記事のいいね情報 |
| article_comment_likes    | 技術記事へのコメントのいいね情報 |
| questions    | 質問情報 |
| question_comments    | 質問への回答情報 |
| question_likes    | 質問のいいね情報 |
| question_comment_likes    | 質問への回答のいいね情報 |
| notifications    | 通知情報 |

## 最後に

今は僕と同じクラスの生徒約10名、後輩約10名、先生2名が使ってくれています。
徐々に範囲を広げていきたいですが、インフラの問題上、重くなってしまうのが問題点です。
本番環境で保守をするのは初めてで全然慣れないですがこれからも使ってもらえるように頑張っていきたいです。

