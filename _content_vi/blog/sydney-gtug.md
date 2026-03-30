---
title: "Hai bài nói chuyện về Go: \"Lexical Scanning in Go\" và \"Cuddle: an App Engine Demo\""
date: 2011-09-01
by:
- Andrew Gerrand
tags:
- appengine
- lexer
- talk
- video
summary: "Hai bài nói chuyện về Go từ Sydney GTUG: Rob Pike giải thích về lexical scanning, và Andrew Gerrand xây dựng một ứng dụng chat thời gian thực đơn giản trên App Engine."
template: true
---


Tối thứ Ba, Rob Pike và Andrew Gerrand đã cùng trình bày tại [Sydney Google Technology User Group](http://www.sydney-gtug.org/).

Bài nói chuyện của Rob, "[Lexical Scanning in Go](http://www.youtube.com/watch?v=HxaD_trXwRE)",
thảo luận về thiết kế của một đoạn mã Go đặc biệt thú vị và mang tính đặc trưng,
đó là thành phần lexer của [gói template mới.](/pkg/exp/template/)

{{video "https://www.youtube.com/embed/HxaD_trXwRE"}}

Các slide [có tại đây](http://cuddle.googlecode.com/hg/talk/lex.html).
Gói template mới được cung cấp dưới dạng [exp/template](/pkg/exp/template/) trong bản phát hành Go r59.
Trong một bản phát hành tương lai, nó sẽ thay thế gói template cũ.

Bài nói chuyện của Andrew, "[Cuddle: an App Engine Demo](http://www.youtube.com/watch?v=HQtLRqqB-Kk)",
mô tả quá trình xây dựng một ứng dụng chat thời gian thực đơn giản sử dụng
các API [Datastore](http://code.google.com/appengine/docs/go/datastore/overview.html),
[Channel](http://code.google.com/appengine/docs/go/channel/overview.html),
và [Memcache](http://code.google.com/appengine/docs/go/datastore/memcache.html) của App Engine.
Bài nói chuyện cũng bao gồm phần hỏi đáp về [Go trên App Engine](http://code.google.com/appengine/docs/go/gettingstarted/)
và Go nói chung.

{{video "https://www.youtube.com/embed/HQtLRqqB-Kk"}}

Các slide [có tại đây](http://cuddle.googlecode.com/hg/talk/index.html).
Mã nguồn có tại [dự án cuddle trên Google Code](http://code.google.com/p/cuddle/).
