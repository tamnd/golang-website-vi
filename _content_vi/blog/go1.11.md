---
title: Go 1.11 đã được phát hành
date: 2018-08-24
by:
- Andrew Bonventre
summary: Go 1.11 bổ sung hỗ trợ sơ bộ cho Go modules, WebAssembly, và nhiều hơn nữa.
---


Ai bảo phát hành vào thứ Sáu là ý tưởng tệ?

Hôm nay nhóm Go vui mừng thông báo bản phát hành Go 1.11.
Bạn có thể tải về từ [trang download](/dl/).

Có nhiều thay đổi và cải tiến đối với toolchain,
runtime và thư viện, nhưng hai tính năng nổi bật đặc biệt thú vị:
modules và hỗ trợ WebAssembly.

Bản phát hành này thêm hỗ trợ sơ bộ cho [một khái niệm mới gọi là "modules,"](/doc/go1.11#modules)
một phương án thay thế cho GOPATH với hỗ trợ tích hợp cho việc phiên bản hóa và phân phối gói.
Hỗ trợ module được coi là thử nghiệm,
và vẫn còn một vài điểm thô ráp cần được làm mịn,
vì vậy hãy thoải mái dùng [issue tracker](/issue/new).

Go 1.11 cũng thêm một cổng thử nghiệm cho [WebAssembly](/doc/go1.11#wasm) (`js/wasm`).
Điều này cho phép lập trình viên biên dịch các chương trình Go thành định dạng nhị phân tương thích với bốn trình duyệt web chính.
Bạn có thể đọc thêm về WebAssembly (viết tắt là "Wasm") tại [webassembly.org](https://webassembly.org/)
và xem [trang wiki này](/wiki/WebAssembly) để biết cách
bắt đầu dùng Wasm với Go.
Đặc biệt cảm ơn [Richard Musiol](https://github.com/neelance) đã đóng góp cổng WebAssembly!

Chúng tôi cũng muốn cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code,
báo cáo lỗi, cung cấp phản hồi và/hoặc kiểm thử các bản beta và release candidate.
Sự đóng góp và tận tâm của bạn giúp đảm bảo Go 1.11 ít lỗi nhất có thể.
Dù vậy, nếu bạn phát hiện bất kỳ vấn đề nào, hãy [tạo issue](/issues/new).

Để biết thêm chi tiết về các thay đổi trong Go 1.11, xem [ghi chú phát hành](/doc/go1.11).

Chúc cuối tuần tuyệt vời và tận hưởng bản phát hành!
