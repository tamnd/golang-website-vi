---
title: Go 1.2 đã được phát hành
date: 2013-12-01
by:
- Andrew Gerrand
tags:
- release
summary: Go 1.2 bổ sung kết quả code coverage khi kiểm thử, goroutine preemption, và nhiều hơn nữa.
---


Chúng tôi vui mừng thông báo bản phát hành Go 1.2, phiên bản ổn định mới nhất của
Ngôn ngữ lập trình Go.

Bản phân phối nhị phân có thể được tải về từ
[nơi thông thường](/doc/install) hoặc nếu bạn muốn
[biên dịch từ nguồn](/doc/install/source) thì nên dùng
tag `release` hoặc `go1.2`.

Bản phát hành mới này ra đời gần bảy tháng sau khi phát hành Go 1.1 vào tháng 5,
một khoảng thời gian ngắn hơn nhiều so với 14 tháng giữa 1.1 và 1.0.
Chúng tôi dự kiến khoảng thời gian tương tự giữa các bản phát hành lớn trong tương lai.

[Go 1.2](/doc/go1.2) bao gồm một vài thay đổi ngôn ngữ nhỏ, một số cải tiến cho việc triển khai ngôn ngữ và
công cụ, một số cải tiến hiệu năng, và nhiều bổ sung và
thay đổi (tương thích ngược) cho thư viện chuẩn.

Hãy đọc [ghi chú phát hành](/doc/go1.2) để biết
toàn bộ chi tiết, vì một số thay đổi có thể ảnh hưởng đến hành vi của các chương trình hiện có (có lỗi).
Dưới đây là những điểm nổi bật của bản phát hành.

[Cú pháp slice ba chỉ số](/doc/go1.2#three_index) mới
thêm khả năng chỉ định cả dung lượng lẫn độ dài. Điều này cho phép
lập trình viên truyền một giá trị slice chỉ có thể truy cập một phần giới hạn của
mảng bên dưới, một kỹ thuật trước đây đòi hỏi phải dùng gói unsafe.

Một tính năng mới lớn của toolchain là khả năng tính toán và hiển thị
[kết quả test coverage](/doc/go1.2#cover).
Xem tài liệu [`go test`](/cmd/go/#hdr-Description_of_testing_flags)
và [công cụ cover](https://golang.org/x/tools/cmd/cover)
để biết chi tiết. Cuối tuần này chúng tôi sẽ xuất bản một bài viết thảo luận chi tiết về tính năng mới này.

Goroutine nay được [lên lịch theo kiểu preemptive](/doc/go1.2#preemption),
trong đó bộ lập lịch được gọi thỉnh thoảng khi vào một hàm.
Điều này có thể ngăn các goroutine bận rộn chiếm dụng tài nguyên của các goroutine khác trên cùng thread.

Tăng kích thước stack goroutine mặc định sẽ cải thiện hiệu năng của một số chương trình.
(Kích thước cũ có xu hướng gây ra việc chuyển đổi segment stack tốn kém trong các phần quan trọng về hiệu năng.)
Ở chiều ngược lại, các giới hạn mới về
[kích thước stack](/doc/go1.2#stack_size) và
[số lượng thread hệ điều hành](/doc/go1.2#thread_limit)
nên ngăn các chương trình hoạt động sai tiêu thụ hết tài nguyên của máy.
(Các giới hạn này có thể được điều chỉnh bằng các hàm mới trong
[gói `runtime/debug`](/pkg/runtime/debug).)

Cuối cùng, trong số [nhiều thay đổi cho thư viện chuẩn](/doc/go1.2#library),
các thay đổi quan trọng bao gồm
[gói `encoding` mới](/doc/go1.2#encoding),
[đối số có chỉ số](/doc/go1.2#fmt_indexed_arguments) trong chuỗi định dạng `Printf` và
một số [bổ sung tiện lợi](/doc/go1.2#text_template) cho các gói template.

Là một phần của bản phát hành, [Go Playground](/play/) đã được
cập nhật lên Go 1.2. Điều này cũng ảnh hưởng đến các dịch vụ dùng Playground, chẳng hạn như
[Go Tour](/tour/) và blog này.
Cập nhật cũng thêm khả năng dùng thread và các gói `os`, `net` và
`unsafe` bên trong sandbox, làm cho nó giống môi trường Go thực hơn.

Với tất cả những người đã giúp làm nên bản phát hành này, từ nhiều người dùng đã
gửi báo cáo lỗi đến 116 (!) người đóng góp đã commit hơn 1600
thay đổi vào lõi: Sự giúp đỡ của bạn là vô giá đối với dự án. Cảm ơn!

_Bài đăng blog này là bài đầu tiên trong_
[Go Advent Calendar](http://blog.gopheracademy.com/day-01-go-1.2),
_một loạt bài viết hàng ngày được trình bày bởi_
[Gopher Academy](http://gopheracademy.com/) _từ ngày 1 đến 25 tháng 12._
