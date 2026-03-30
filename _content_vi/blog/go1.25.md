---
title: Go 1.25 đã được phát hành
date: 2025-08-12
by:
- Dmitri Shuralyov, on behalf of the Go team
summary: Go 1.25 bổ sung GOMAXPROCS theo container, gói testing/synctest, GC thử nghiệm, encoding/json/v2 thử nghiệm, và nhiều hơn nữa.
---

Hôm nay nhóm Go vui mừng phát hành Go 1.25.
Bạn có thể tìm các file nhị phân và trình cài đặt trên [trang download](/dl/).

Go 1.25 có nhiều cải tiến so với Go 1.24 trong
[các công cụ](/doc/go1.25#tools),
[runtime](/doc/go1.25#runtime),
[trình biên dịch](/doc/go1.25#compiler),
[trình liên kết](/doc/go1.25#linker),
và [thư viện chuẩn](/doc/go1.25#library),
bao gồm việc bổ sung một [gói mới](/doc/go1.25#new-testingsynctest-package).
Có các thay đổi [theo nền tảng cụ thể](/doc/go1.25#ports)
và cập nhật [cài đặt `GODEBUG`](/doc/godebug#go-125).

Một số tính năng trong Go 1.25 đang ở giai đoạn thử nghiệm
và chỉ hiển thị khi bạn bật tùy chọn một cách rõ ràng.
Đáng chú ý là [bộ gom rác thử nghiệm mới](/doc/go1.25#new-experimental-garbage-collector),
và [gói `encoding/json/v2` thử nghiệm mới](/doc/go1.25#json_v2)
có sẵn để bạn thử trước và cung cấp phản hồi.
Sẽ rất hữu ích nếu bạn làm được điều đó!

Hãy tham khảo [Ghi chú phát hành Go 1.25](/doc/go1.25) để xem danh sách đầy đủ
các bổ sung, thay đổi và cải tiến trong Go 1.25.

Trong vài tuần tới, các bài đăng blog tiếp theo sẽ đề cập một số chủ đề
liên quan đến Go 1.25 một cách chi tiết hơn. Hãy quay lại sau để đọc các bài đó.

Cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code, báo cáo lỗi,
thử các tính năng thử nghiệm, chia sẻ phản hồi, và kiểm thử các release candidate.
Nỗ lực của bạn giúp Go 1.25 ổn định nhất có thể.
Như thường lệ, nếu bạn phát hiện bất kỳ vấn đề nào, hãy [tạo issue](/issue/new).

Chúc bạn tận hưởng bản phát hành mới!
