---
title: Tổng kết GopherCon 2015
date: 2015-07-28
by:
- Andrew Gerrand
tags:
- conference
- report
- gopher
summary: Báo cáo từ GopherCon 2015.
template: true
---


Vài tuần trước, các lập trình viên Go từ khắp nơi trên thế giới đã đổ về Denver,
Colorado cho GopherCon 2015. Hội nghị một track kéo dài hai ngày thu hút
hơn 1.250 người tham dự - gần gấp đôi số lượng năm ngoái - và có 22
bài nói chuyện được trình bày bởi các thành viên cộng đồng Go.

{{image "gophercon2015/cowboy.jpg" 550}}

<p>
<small>Cowboy Gopher (một món đồ chơi được tặng cho mỗi người tham dự) canh giữ trang trại.<br>
<i>Ảnh chụp bởi <a href="https://twitter.com/nathany/status/619861336399351808">Nathan Youngman</a>. Gopher bởi Renee French.</i>
</small>
</p>

Hôm nay ban tổ chức đã đăng các video trực tuyến để bạn có thể thưởng thức
hội nghị từ xa:

[Ngày 1](http://gophercon.com/schedule/8july/):

  - Go, Open Source, Community - Russ Cox ([video](https://www.youtube.com/watch?v=XvZOdpd_9tc))
    ([bản văn](/blog/open-source))
  - Go kit: A Standard Library for Distributed Programming - Peter Bourgon
    ([video](https://www.youtube.com/watch?v=1AjaZi4QuGo)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Go%20kit/go-kit.pdf))
  - Delve Into Go - Derek Parker ([video](https://www.youtube.com/watch?v=InG72scKPd4))
    ([slides](http://go-talks.appspot.com/github.com/derekparker/talks/gophercon-2015/delve-into-go.slide))
  - How a complete beginner learned Go as her first backend language in 5
    weeks - Audrey Lim ([video](https://www.youtube.com/watch?v=fZh8uCInEfw))
    ([slides](https://github.com/gophercon/2015-talks/blob/master/Audrey%20Lim%20-%20How%20a%20Complete%20Beginner%20Picked%20Up%20Go%20as%20Her%20First%20Backend%20Language%20in%205%20weeks/audreylim_slides.pdf))
  - A Practical Guide to Preventing Deadlocks and Leaks in Go - Richard
    Fliam ([video](https://www.youtube.com/watch?v=3EW1hZ8DVyw))
  - Go GC: Solving the Latency Problem - Rick Hudson ([video](https://www.youtube.com/watch?v=aiv1JOfMjm0))
    ([slides](/talks/2015/go-gc.pdf))
  - Simplicity and Go - Katherine Cox-Buday ([video](https://www.youtube.com/watch?v=S6mEo_FHZ5Y))
    ([slides](https://github.com/gophercon/2015-talks/blob/master/Katherine%20Cox-Buday:%20Simplicity%20%26%20Go/Simplicity%20%26%20Go.pdf))
  - Rebuilding Parse.com in Go - an opinionated rewrite - Abhishek Kona
    ([video](https://www.youtube.com/watch?v=_f9LS-OWfeA)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Abhishek%20Kona%20Rewriting%20Parse%20in%20GO/myslides.pdf))
  - Prometheus: Designing and Implementing a Modern Monitoring Solution in
    Go - Björn Rabenstein ([video](https://www.youtube.com/watch?v=1V7eJ0jN8-E))
    ([slides](https://github.com/gophercon/2015-talks/blob/master/Bj%C3%B6rn%20Rabenstein%20-%20Prometheus/slides.pdf))
  - What Could Go Wrong? - Kevin Cantwell ([video](https://www.youtube.com/watch?v=VC3QXZ-x5yI))
  - The Roots of Go - Baishampayan Ghose ([video](https://www.youtube.com/watch?v=0hPOopcJ8-E))
    ([slides](https://speakerdeck.com/bg/the-roots-of-go))

[Ngày 2](http://gophercon.com/schedule/9july/):

  - The Evolution of Go - Robert Griesemer ([video](https://www.youtube.com/watch?v=0ReKdcpNyQg))
    ([slides](/talks/2015/gophercon-goevolution.slide))
  - Static Code Analysis Using SSA - Ben Johnson ([video](https://www.youtube.com/watch?v=D2-gaMvWfQY))
    ([slides](https://speakerdeck.com/benbjohnson/static-code-analysis-using-ssa))
  - Go on Mobile - Hana Kim ([video](https://www.youtube.com/watch?v=sQ6-HyPxHKg))
    ([slides](/talks/2015/gophercon-go-on-mobile.slide))
  - Go Dynamic Tools - Dmitry Vyukov ([video](https://www.youtube.com/watch?v=a9xrxRsIbSU))
    ([slides](/talks/2015/dynamic-tools.slide))
  - Embrace the Interface - Tomás Senart ([video](https://www.youtube.com/watch?v=xyDkyFjzFVc))
    ([slides](https://github.com/gophercon/2015-talks/blob/master/Tom%C3%A1s%20Senart%20-%20Embrace%20the%20Interface/ETI.pdf))
  - Uptime: Building Resilient Services with Go - Blake Caldwell ([video](https://www.youtube.com/watch?v=PyBJQA4clfc))
    ([slides](https://github.com/gophercon/2015-talks/blob/master/Blake%20Caldwell%20-%20Uptime:%20Building%20Resilient%20Services%20with%20Go/2015-GopherCon-Talk-Uptime.pdf))
  - Cayley: Building a Graph Database - Barak Michener ([video](https://www.youtube.com/watch?v=-9kWbPmSyCI))
    ([slides](https://github.com/gophercon/2015-talks/blob/master/Barak%20Michener%20-%20Cayley:%20Building%20a%20Graph%20Database/Cayley%20-%20Building%20a%20Graph%20Database.pdf))
  - Code Generation For The Sake Of Consistency - Sarah Adams ([video](https://www.youtube.com/watch?v=kGAgHwfjg1s))
  - The Many Faces of Struct Tags - Sam Helman and Kyle Erf ([video](https://www.youtube.com/watch?v=_SCRvMunkdA))
    ([slides](https://github.com/gophercon/2015-talks/blob/master/Sam%20Helman%20%26%20Kyle%20Erf%20-%20The%20Many%20Faces%20of%20Struct%20Tags/StructTags.pdf))
  - Betting the Company on Go and Winning - Kelsey Hightower ([video](https://www.youtube.com/watch?v=wqVbLlHqAeY))
  - How Go Was Made - Andrew Gerrand ([video](https://www.youtube.com/watch?v=0ht89TxZZnk))
    ([slides](/talks/2015/how-go-was-made.slide))

[Ngày hack](http://gophercon.com/schedule/10july/) cũng rất vui,
với nhiều giờ [bài nói chớp nhoáng](https://www.youtube.com/playlist?list=PL2ntRZ1ySWBeHqlHM8DmvS8axgbrpvF9b)
và nhiều hoạt động từ lập trình robot
đến giải đấu Magic: the Gathering.

Cảm ơn rất nhiều đến ban tổ chức sự kiện Brian Ketelsen và Eric St. Martin và
nhóm sản xuất của họ, các nhà tài trợ, các diễn giả, và những người tham dự đã tạo nên
một hội nghị vui vẻ và sôi động như vậy. Hy vọng gặp bạn ở đó năm tới!
