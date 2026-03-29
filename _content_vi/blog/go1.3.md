---
title: Go 1.3 đã được phát hành
date: 2014-06-18
by:
- Andrew Gerrand
summary: Go 1.3 mang lại hiệu năng tốt hơn, phân tích tĩnh trong godoc, và nhiều hơn nữa.
---


Hôm nay chúng tôi vui mừng thông báo phát hành [Go 1.3](/doc/go1.3).
Bản phát hành này ra đời sáu tháng sau bản phát hành lớn trước, cung cấp hiệu năng tốt hơn,
công cụ được cải thiện, hỗ trợ chạy Go trong môi trường mới, và nhiều hơn nữa.
Tất cả người dùng Go nên nâng cấp lên Go 1.3.
Bạn có thể lấy bản phát hành từ [trang download](/dl/) và
tìm danh sách đầy đủ các cải tiến và sửa lỗi trong
[ghi chú phát hành](/doc/go1.3).
Dưới đây là một số điểm nổi bật.

[Godoc](https://godoc.org/code.google.com/p/go.tools/cmd/godoc),
máy chủ tài liệu Go, nay thực hiện phân tích tĩnh.
Khi được bật với cờ -analysis, kết quả phân tích được trình bày
trong cả chế độ xem nguồn lẫn tài liệu gói, giúp việc điều hướng và hiểu
chương trình Go dễ dàng hơn bao giờ hết.

Toolchain gc nay hỗ trợ sandbox thực thi Native Client (NaCl) trên
kiến trúc Intel 32-bit và 64-bit.
Điều này cho phép thực thi an toàn code không đáng tin cậy, hữu ích trong các môi trường như
[Playground](/blog/playground).
Để thiết lập NaCl trên hệ thống của bạn, xem [trang wiki NativeClient](/wiki/NativeClient).

Bản phát hành này cũng bao gồm hỗ trợ thử nghiệm cho các hệ điều hành DragonFly BSD,
Plan 9 và Solaris. Để dùng Go trên các hệ thống này, bạn phải
[cài đặt từ nguồn](/doc/install/source).

Các thay đổi trong runtime đã cải thiện
[hiệu năng](/doc/go1.3#performance) của các binary Go,
với bộ gom rác được cải thiện, [chiến lược quản lý stack goroutine "liên tục" mới](/s/contigstacks),
bộ phát hiện race nhanh hơn, và cải tiến cho engine biểu thức chính quy.

Là một phần của việc [tái cấu trúc tổng thể](/s/go13linker) trình liên kết Go,
các trình biên dịch và trình liên kết đã được tái cấu trúc. Giai đoạn chọn lệnh
vốn là một phần của trình liên kết đã được chuyển vào trình biên dịch.
Điều này có thể tăng tốc build tăng dần cho các dự án lớn.

[Bộ gom rác](/doc/go1.3#garbage_collector) nay
chính xác khi kiểm tra stack (việc thu thập heap đã chính xác kể từ Go
1.1), có nghĩa là giá trị không phải con trỏ như số nguyên sẽ không bao giờ
bị nhầm là con trỏ và ngăn bộ nhớ không dùng được thu hồi. Thay đổi này
ảnh hưởng đến code dùng gói unsafe; nếu bạn có code unsafe, hãy đọc kỹ
[ghi chú phát hành](/doc/go1.3#garbage_collector)
để xem code của bạn có cần cập nhật không.

Chúng tôi muốn cảm ơn nhiều người đã đóng góp cho bản phát hành này;
nó sẽ không thể thực hiện được nếu không có sự giúp đỡ của bạn.

Vậy, bạn còn chờ gì nữa?
Hãy đến [trang download](/dl/) và bắt đầu code thôi.
