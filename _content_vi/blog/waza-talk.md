---
title: Đồng thời không phải là song song
date: 2013-01-16
by:
- Andrew Gerrand
tags:
- concurrency
- talk
- video
summary: Xem bài nói chuyện của Rob Pike, _Concurrency is not parallelism._
template: true
---


Nếu có một điều mà hầu hết mọi người biết về Go,
đó là nó được thiết kế cho tính đồng thời.
Không có phần giới thiệu nào về Go là đầy đủ nếu không có sự trình diễn về goroutine và channel.

Nhưng khi người ta nghe từ _đồng thời_ họ thường nghĩ đến _song song_,
một khái niệm liên quan nhưng khá khác biệt.
Trong lập trình, đồng thời là _sự kết hợp_ của các tiến trình thực thi độc lập,
trong khi song song là _thực thi đồng thời_ của các tính toán (có thể liên quan).
Đồng thời là về việc _xử lý_ nhiều thứ cùng một lúc.
Song song là về việc _làm_ nhiều thứ cùng một lúc.

Để làm rõ sự nhầm lẫn này, Rob Pike đã có bài nói chuyện tại hội nghị Waza của [Heroku](http://heroku.com/)
với tiêu đề
[_Concurrency is not parallelism_](https://blog.heroku.com/concurrency_is_not_parallelism),
và video ghi lại bài nói chuyện đã được phát hành vài tháng trước.

{{video "https://www.youtube.com/embed/oV9rvDllKEg" 500 281}}

Các slide có sẵn tại [go.dev/talks](/talks/2012/waza.slide)
(dùng phím mũi tên trái và phải để điều hướng).

Để tìm hiểu về các nguyên tố đồng thời của Go,
xem [Go concurrency patterns](http://www.youtube.com/watch?v=f6kdp27TYZs)
([slides](/talks/2012/concurrency.slide)).
