---
title: Go 1.14 đã được phát hành
date: 2020-02-25
by:
- Alex Rakoczy
summary: Go 1.14 bổ sung hỗ trợ module sẵn sàng cho production, defer nhanh hơn, goroutine preemption tốt hơn, và nhiều hơn nữa.
---


Hôm nay nhóm Go rất vui mừng thông báo bản phát hành Go 1.14. Bạn có thể tải về từ [trang download](/dl).

Một số điểm nổi bật bao gồm:

  - Hỗ trợ module trong lệnh `go` đã sẵn sàng cho môi trường production. Chúng tôi khuyến khích tất cả người dùng [chuyển sang dùng `go` modules để quản lý dependency](/doc/go1.14#introduction).
  - [Nhúng interface với method set chồng nhau](/doc/go1.14#language)
  - [Cải thiện hiệu năng defer](/doc/go1.14#runtime)
  - [Goroutine hỗ trợ preemption không đồng bộ](/doc/go1.14#runtime)
  - [Bộ cấp phát trang hiệu quả hơn](/doc/go1.14#runtime)
  - [Timer nội bộ hiệu quả hơn](/doc/go1.14#runtime)

Để xem danh sách đầy đủ các thay đổi và thông tin về các cải tiến trên, xem [**ghi chú phát hành Go 1.14**](/doc/go1.14).

Chúng tôi muốn cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code, báo cáo lỗi, cung cấp phản hồi và/hoặc kiểm thử bản beta và release candidate.
Sự đóng góp và tận tâm của bạn giúp đảm bảo Go 1.14 ổn định nhất có thể.
Dù vậy, nếu bạn phát hiện bất kỳ vấn đề nào, hãy [tạo issue](/issue/new).

Chúc bạn tận hưởng bản phát hành mới!
