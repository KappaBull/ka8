---
title: "Gotanada.unity #10 に行ってきた"
date: 2019-01-24T06:00:23+06:00
hero: /images/posts/study-group/IMG_2644.PNG
description: Gotanda.unityの参加レポート
menu:
  sidebar:
    name: Gotanada.unity#10
    identifier: gotandaunity
    parent: study-group
    weight: 600
---
  
Gotanda.unityの勉強会に行ってきたのでそのまとめ記事です。  
今回は株式会社アカツキさんのところで行われたGotanda.unity #10へ参加してきました。
  
<iframe class="embed-card embed-webcard" style="display: block; width: 100%; height: 155px; max-width: 500px; margin: 10px 0px;" title="Gotanda.unity #10" src="https://hatenablog-parts.com/embed?url=https%3A%2F%2Fmeetup.unity3d.jp%2Fjp%2Fevents%2F1032" frameborder="0" scrolling="no"></iframe>
土足厳禁の芝生という良い会場でした。弊社にもちょっと欲しい。（もしかして何処かにある？  
タイムスケジュール順に上から資料、その際のツイート、コメントという形で紹介していきます。  
  
### LT #1 会場/アカツキ枠 Norihiro Funakoshi : スクリプトの処理を効率化してゲームを軽くしよう
(資料が公開されて無いので公開されたら追加)  
{{< twitter_simple 1088024511834800129 >}}
{{< twitter_simple 1088025505943568384 >}}
{{< twitter_simple 1088025673891905536 >}}
Linqの実行タイミングをずらして先にキャッシュさせて使うみたいなところは結構忘れがちなところもあるので気を付けたいところ  
あと、MonoBehaviourの空メソッドは許してはならない。  
  
### LT #2 Hiroto Imoto : どこから始めるUnityテスト
{{< twitter_simple 1088024511834800129 >}}
{{< twitter_simple 1088025505943568384 >}}
{{< twitter_simple 1088025673891905536 >}}
  
<iframe id="talk_frame_486760" style="border: 0; padding: 0; margin: 0; background: transparent;" src="//speakerdeck.com/player/5922ddb8deb7440e8ab9100edd2546ec" width="710" height="532" frameborder="0" allowfullscreen="allowfullscreen"></iframe>
{{< twitter_simple 1088028806051516417 >}}
{{< twitter_simple 1088027924291145728 >}}
{{< twitter_simple 1088028923085021185 >}}
DI、疎結合をしっかりする事でテストがしやすくなる。リファクタリングに近いからテストは未来の開発短縮に繋がるし設計見直しも出来るよという事なので、密結合部分はテスト出来るのかどうかという点もふまえて見直せると良いのかも。  
  
### LT #3 mao : ECSベースのVRMSpringBoneを実装してみた話
<iframe class="embed-card embed-webcard" style="display: block; width: 100%; height: 155px; max-width: 500px; margin: 10px 0px;" title="ECSベースのVRMSpringBoneを実装してみた話" src="https://hatenablog-parts.com/embed?url=https%3A%2F%2Ft.co%2FmeUjgTbiYW" frameborder="0" scrolling="no"></iframe>
{{< twitter_simple 1088029441345781762 >}}
{{< twitter_simple 1088030917103280128 >}}
{{< twitter_simple 1088031172712693761 >}}
UnsafeUtilityでC++感溢れるDirectXを触ってた頃を思い出すような動きが出来るっぽいです。ポインタあんまり触りたくないですが、突き詰めると触る事になりそう…  
あと、ECSをまだ触ってないのでpreview外れる前に一度見ておきたい  
  
### LT #4 Nozomi Tanaka : Unity(C#)メモリ小話
<iframe id="talk_frame_486807" style="border: 0; padding: 0; margin: 0; background: transparent;" src="//speakerdeck.com/player/1bc568b2c004493f838dd8f235831520" width="710" height="399" frameborder="0" allowfullscreen="allowfullscreen"></iframe>
{{< twitter_simple 1088031918598250496 >}}
{{< twitter_simple 1088031801677729792 >}}
{{< twitter_simple 1088033408855433216 >}}
スタックとヒープを意識してGCを抑えようというお話でした。  
ラムダにメソッドを食わせると大抵はゴミが出るので、インゲームでは気をつけていきたい。  
  
### LT #5 小林拓 : VRMと暗号化
<iframe class="embed-card embed-webcard" style="display: block; width: 100%; height: 155px; max-width: 500px; margin: 10px 0px;" title="Gotanda.unity #10資料(VRMと暗号)" src="https://hatenablog-parts.com/embed?url=https%3A%2F%2Fdocs.google.com%2Fpresentation%2Fd%2F1VRLE5DqObO8VNN3IrXgbkkZmVDHnPoK1INxu7we0zEs%2Fedit" frameborder="0" scrolling="no"></iframe>
<iframe class="embed-card embed-webcard" style="display: block; width: 100%; height: 155px; max-width: 500px; margin: 10px 0px;" title="dwango/UniVRM" src="https://hatenablog-parts.com/embed?url=https%3A%2F%2Fgithub.com%2Fdwango%2FUniVRM" frameborder="0" scrolling="no"></iframe>
{{< twitter_simple 1088034977625735168 >}}
{{< twitter_simple 1088034440197025792 >}}
{{< twitter_simple 1088035362339905536 >}}
実際にVRoid SDK内で使用した暗号化のお話でした。1/28(月)に出るらしい。  
時間切れで最後まで聞けなかったですが、Saltと初期ベクターを生成の度にちゃんと変更するセキュアな暗号化が出来る奴でした。  
  
  
{{< twitter_simple 1088035969259864064 >}}
  
  
### LT #6 Tomohiro Okamura : PUN2を使おう！
<iframe class="embed-card embed-webcard" style="display: block; width: 100%; height: 155px; max-width: 500px; margin: 10px 0px;" title="PUN2を使おう！.pdf" src="https://hatenablog-parts.com/embed?url=https%3A%2F%2Fdrive.google.com%2Ffile%2Fd%2F1Ge3PVoxJjagdMHQghPEQrX312YH8LaZR%2Fview" frameborder="0" scrolling="no"></iframe>
{{< twitter_simple 1088038827816407040 >}}
{{< twitter_simple 1088039549802934272 >}}
お行儀(フォルダ、Namespace構成)と動作が軽くなって戻ってきた感じだった。
だが100MBを超えるファイルが含まれるらしいのでGitLFSとか使わないといけなさそう。  
  
### LT #7 Naoya Mizota : VRM対応スマホアプリのすすめ
<iframe style="border: 1px solid #CCC; border-width: 1px; margin-bottom: 5px; max-width: 100%;" src="https://www.slideshare.net/slideshow/embed_code/key/DDDGSNGb12pUZX" width="427" height="356" frameborder="0" marginwidth="0" marginheight="0" scrolling="no" allowfullscreen=""> </iframe>
{{< twitter_simple 1088040739940913153 >}}
{{< twitter_simple 1088041849380458496 >}}
懇親会でお話を聞いた感じだと、VRMは名前だけとか特定の情報だけロードするのが標準では無理っぽいので、今は起動時に全ロードさせているらしい。  
CreateFromFileみたいなロード欲しいね…。  
  
### LT #8 joshuaarthurelcaro : もう「ぼくが持ってるUnityと違う…」ではない！
<iframe class="embed-card embed-webcard" style="display: block; width: 100%; height: 155px; max-width: 500px; margin: 10px 0px;" title="もう「ぼくが持ってるUnityと違う…」ではない！HDRPで劇的にルックを変えよう！～部屋とかキャラのセットアップをしてみる。～ - さやちゃんぐbotスクラップス" src="https://hatenablog-parts.com/embed?url=https%3A%2F%2Fscrapbox.io%2Fsayachang%2F%25E3%2582%2582%25E3%2581%2586%25E3%2580%258C%25E3%2581%25BC%25E3%2581%258F%25E3%2581%258C%25E6%258C%2581%25E3%2581%25A3%25E3%2581%25A6%25E3%2582%258BUnity%25E3%2581%25A8%25E9%2581%2595%25E3%2581%2586%25E2%2580%25A6%25E3%2580%258D%25E3%2581%25A7%25E3%2581%25AF%25E3%2581%25AA%25E3%2581%2584%25EF%25BC%2581HDRP%25E3%2581%25A7%25E5%258A%2587%25E7%259A%2584%25E3%2581%25AB%25E3%2583%25AB%25E3%2583%2583%25E3%2582%25AF%25E3%2582%2592%25E5%25A4%2589%25E3%2581%2588%25E3%2582%2588%25E3%2581%2586%25EF%25BC%2581%25EF%25BD%259E%25E9%2583%25A8%25E5%25B1%258B%25E3%2581%25A8%25E3%2581%258B%25E3%2582%25AD%25E3%2583%25A3%25E3%2583%25A9%25E3%2581%25AE%25E3%2582%25BB%25E3%2583%2583%25E3%2583%2588%25E3%2582%25A2%25E3%2583%2583%25E3%2583%2597%25E3%2582%2592%25E3%2581%2597%25E3%2581%25A6%25E3%2581%25BF%25E3%2582%258B%25E3%2580%2582%25EF%25BD%259E" frameborder="0" scrolling="no"></iframe>
{{< twitter_simple 1088042336662183938 >}}
{{< twitter_simple 1088043429446135808 >}}
アンリアルみたいな良い感じの絵が出せる奴。鏡とかワンポチで作れる。  
いずれガンに効く。  
  
### LT #9 aoiaioi : 簡単な弾幕をつくる
<iframe class="embed-card embed-webcard" style="display: block; width: 100%; height: 155px; max-width: 500px; margin: 10px 0px;" title="簡単な弾幕をつくる" src="https://hatenablog-parts.com/embed?url=https%3A%2F%2Fdocs.google.com%2Fpresentation%2Fd%2F1nSYcTwCT0vCo6oQxnwFxgGlrNHM5HZHTcW-vIRa9DDc%2Fedit" frameborder="0" scrolling="no"></iframe>
{{< twitter_simple 1088045605698465792 >}}
龍神録プログラミングの館を思い出した。Vector3は実装した後にあったんかい案件に遭遇した事もあるので見てちゃんと活用していきたい。
  
  
今回、Gotanda.unityは初参加で別シリーズ(？)の完全に理解したシリーズは行ったりしてましたが、ガチ勢から初心者向けまで幅広くLTが行われており、とても良い勉強会でした。</p>
<p>こういうのを企画してくれるUnityAmbassadorのお二人には足を向けて寝れませんね。


LTの内容はVRMとメモリ周りが多かったイメージですが、  
VRMは今度、勉強会が開催されるのでこちらも枠が空いたらこちらも是非参加したいです。  
<iframe class="embed-card embed-webcard" style="display: block; width: 100%; height: 155px; max-width: 500px; margin: 10px 0px;" title="第1回 VRM勉強会 (2019/02/19 19:00〜)" src="https://hatenablog-parts.com/embed?url=https%3A%2F%2Fconnpass.com%2Fevent%2F116985%2F" frameborder="0" scrolling="no"></iframe>