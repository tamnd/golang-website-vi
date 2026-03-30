---
title: Go và Google Cloud Platform
date: 2013-06-12
by:
- Andrew Gerrand
summary: Hai bài nói chuyện về việc sử dụng Go với Google Cloud Platform, từ Google I/O 2013.
template: true
---

## Giới thiệu

Năm 2011, chúng tôi công bố [Go runtime cho App Engine](https://developers.google.com/appengine/docs/go/overview).
Kể từ đó, chúng tôi đã tiếp tục cải thiện trải nghiệm Go trên App Engine,
và nói chung đã cải thiện hỗ trợ Go cho Google Cloud Platform.
Chẳng hạn, [google-api-go-client](http://code.google.com/p/google-api-go-client) cung cấp
giao diện Go cho một loạt các API công khai của Google,
bao gồm Compute Engine, Cloud Storage, BigQuery,
Drive, và nhiều thứ khác.

Tìm hiểu thêm bằng cách xem các bài nói chuyện này từ Google I/O năm nay:

## Ứng dụng hiệu năng cao với Go trên App Engine

_Go runtime cho App Engine là một engine hiệu năng cao để_
_chạy các ứng dụng web. Nó cho thời gian phản hồi nhanh,_
_khởi động instance trong một phần giây, tận dụng tối đa_
_giờ instance, và cho phép ứng dụng của bạn thực hiện xử lý nghiêm túc_
_ở tốc độ máy đầy đủ._
_Hãy đến để nghe cách khai thác toàn bộ sức mạnh của Go trên App_
_Engine và làm cho ứng dụng web của bạn tốt nhất có thể._

{{video "https://www.youtube.com/embed/fc25ihfXhbg"}}

## Tất cả các tàu trên thế giới

Trực quan hóa dữ liệu với Google Cloud và Maps

_Hàng chục nghìn con tàu báo cáo vị trí của chúng ít nhất một lần_
_mỗi 5 phút, 24 giờ một ngày._
_Trực quan hóa lượng dữ liệu đó và phục vụ nó cho số lượng lớn_
_người dùng đòi hỏi nhiều sức mạnh cả ở trình duyệt lẫn máy chủ._
_Phiên này sẽ khám phá việc sử dụng Maps,_
_App Engine, Go, Compute Engine, BigQuery, Cloud Storage,_
_và WebGL để thực hiện trực quan hóa dữ liệu quy mô lớn._

{{video "https://www.youtube.com/embed/MT7cd4M9vzs"}}
