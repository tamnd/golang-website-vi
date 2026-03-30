---
title: Các mẫu concurrency Go nâng cao
date: 2013-05-23
by:
- Andrew Gerrand
tags:
- talk
- video
- concurrency
summary: Xem bài nói chuyện của Sameer Ajmani, "Advanced Go Concurrency Patterns," từ Google I/O 2013.
template: true
---


Tại Google I/O một năm trước, Rob Pike đã trình bày [_Go Concurrency Patterns_](/talks/2012/concurrency.slide),
một phần giới thiệu về mô hình concurrency của Go.
Tuần trước, tại I/O 2013, thành viên nhóm Go Sameer Ajmani tiếp tục câu chuyện
với [_Advanced Go Concurrency Patterns_](/talks/2013/advconc.slide),
một cái nhìn sâu vào một bài toán lập trình đồng thời thực tế.
Bài nói chuyện cho thấy cách phát hiện và tránh deadlock và race condition,
và minh họa triển khai deadline,
hủy bỏ, và nhiều thứ khác.
Với những ai muốn đưa lập trình Go của mình lên tầm cao mới, đây là nội dung bắt buộc phải xem.

{{video "https://www.youtube.com/embed/QDDwwePbDtw?rel=0"}}

Các slide [có thể tải tại đây](/talks/2013/advconc.slide)
(dùng mũi tên trái và phải để điều hướng).

Các slide được tạo bằng [công cụ present](https://pkg.go.dev/golang.org/x/tools/present),
và các đoạn code có thể chạy được hỗ trợ bởi [Go Playground](/play/).
Mã nguồn của bài nói chuyện này nằm trong [sub-repository go.talks](https://github.com/golang/talks/tree/master/content/2013/advconc).
