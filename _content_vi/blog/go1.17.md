---
title: Go 1.17 đã được phát hành
date: 2021-08-16
by:
- Matt Pearring
- Alex Rakoczy
summary: Go 1.17 bổ sung cải tiến hiệu năng, tối ưu module, arm64 trên Windows, và nhiều hơn nữa.
---


Hôm nay nhóm Go vô cùng hào hứng phát hành Go 1.17, bạn có thể tải về từ
[trang download](/dl/).

Bản phát hành này mang đến các cải tiến bổ sung cho trình biên dịch, cụ thể là
[cách mới truyền đối số và kết quả hàm](/doc/go1.17#compiler).
Thay đổi này cho thấy cải thiện hiệu năng khoảng 5% trong các chương trình Go và giảm kích thước binary
khoảng 2% cho nền tảng amd64. Hỗ trợ thêm cho các nền tảng khác sẽ có trong các bản phát hành tương lai.

Go 1.17 cũng thêm hỗ trợ cho
[kiến trúc ARM 64-bit trên Windows](/doc/go1.17#ports), cho phép các lập trình viên Go chạy
Go natively trên nhiều thiết bị hơn.

Chúng tôi cũng giới thiệu [đồ thị module được cắt tỉa (pruned module graphs)](/doc/go1.17#go-command) trong bản
phát hành này. Các module chỉ định `go 1.17` hoặc cao hơn trong file `go.mod` sẽ có đồ thị module
chỉ bao gồm các dependency trực tiếp của các module Go 1.17 khác, không phải toàn bộ dependency bắc cầu.
Điều này giúp tránh phải tải về hoặc đọc các file `go.mod` cho các dependency không liên quan,
tiết kiệm thời gian trong công việc hàng ngày.

Go 1.17 đi kèm với ba thay đổi nhỏ [đối với ngôn ngữ](/doc/go1.17#language).
Hai thay đổi đầu là các hàm mới trong gói `unsafe` giúp chương trình dễ tuân thủ
các quy tắc `unsafe.Pointer` hơn: `unsafe.Add` cho phép
[tính toán con trỏ an toàn hơn](/pkg/unsafe#Add), trong khi `unsafe.Slice` cho phép
[chuyển đổi con trỏ sang slice an toàn hơn](/pkg/unsafe#Slice). Thay đổi thứ ba là
mở rộng quy tắc chuyển đổi kiểu của ngôn ngữ để cho phép chuyển đổi từ
[slice sang array pointer](/ref/spec#Conversions_from_slice_to_array_pointer),
miễn là slice ít nhất bằng array về kích thước lúc runtime.

Ngoài ra còn khá nhiều cải tiến và sửa lỗi khác, bao gồm cải tiến xác minh
trong [crypto/x509](/doc/go1.17#crypto/x509) và thay đổi trong
[xử lý query URL](/doc/go1.17#semicolons). Để xem danh sách đầy đủ các thay đổi và
thông tin về các cải tiến trên, xem
[ghi chú phát hành đầy đủ](/doc/go1.17).

Cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code, báo cáo lỗi, chia sẻ phản hồi,
và kiểm thử bản beta và release candidate. Nỗ lực của bạn giúp đảm bảo Go 1.17 ổn định
nhất có thể. Như thường lệ, nếu bạn phát hiện bất kỳ vấn đề nào, hãy
[tạo issue](/issue/new).

Chúc bạn tận hưởng bản phát hành mới!
