---
title: Go 1.16 đã được phát hành
date: 2021-02-16
by:
- Matt Pearring
- Dmitri Shuralyov
summary: Go 1.16 bổ sung nhúng file, hỗ trợ Apple Silicon, và nhiều hơn nữa.
---


Hôm nay nhóm Go rất vui mừng thông báo bản phát hành Go 1.16.
Bạn có thể tải về từ [trang download](/dl/).

[Gói embed](/doc/go1.16#library-embed) mới
cho phép truy cập các file được nhúng vào lúc biên dịch dùng chỉ thị `//go:embed` mới.
Giờ đây việc đóng gói các file dữ liệu hỗ trợ vào chương trình Go trở nên dễ dàng,
giúp quá trình phát triển với Go mượt mà hơn nữa.
Bạn có thể bắt đầu bằng cách xem
[tài liệu gói embed](https://pkg.go.dev/embed).
Carl Johnson cũng đã viết một hướng dẫn hay,
"[How to use Go embed](https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/)".

Go 1.16 cũng bổ sung
[hỗ trợ macOS ARM64](/doc/go1.16#darwin)
(còn gọi là Apple silicon).
Kể từ khi Apple thông báo kiến trúc arm64 mới, chúng tôi đã làm việc chặt chẽ với họ để đảm bảo Go được hỗ trợ đầy đủ; xem bài đăng blog của chúng tôi
"[Go trên ARM và xa hơn](/blog/ports)"
để biết thêm.

Lưu ý rằng Go 1.16
[yêu cầu dùng Go modules theo mặc định](/doc/go1.16#modules),
bởi theo khảo sát lập trình viên Go 2020 của chúng tôi,
96% lập trình viên Go đã chuyển sang dùng modules.
Gần đây chúng tôi đã thêm tài liệu chính thức về [phát triển và xuất bản module](/doc/modules/developing).

Cuối cùng, có nhiều cải tiến và sửa lỗi khác,
bao gồm quá trình build nhanh hơn tới 25% và dùng ít bộ nhớ hơn tới 15%.
Để xem danh sách đầy đủ các thay đổi và thông tin về các cải tiến trên,
xem
[ghi chú phát hành Go 1.16](/doc/go1.16).

Chúng tôi muốn cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code,
báo cáo lỗi, cung cấp phản hồi, và kiểm thử bản beta và release candidate.

Sự đóng góp và tận tâm của bạn giúp đảm bảo Go 1.16 ổn định nhất có thể.
Dù vậy, nếu bạn phát hiện bất kỳ vấn đề nào, hãy
[tạo issue](/issue/new).

Chúc bạn tận hưởng bản phát hành mới!
