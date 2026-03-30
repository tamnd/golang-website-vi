---
title: Go 1.26 đã được phát hành
date: 2026-02-10
by:
- Carlos Amedee, on behalf of the Go team
summary: Go 1.26 bổ sung bộ gom rác mới, giảm chi phí cgo, gói simd/archsimd thử nghiệm, gói runtime/secret thử nghiệm, và nhiều hơn nữa.
---

Hôm nay nhóm Go vui mừng phát hành Go 1.26.
Bạn có thể tìm các file nhị phân và trình cài đặt trên [trang download](/dl/).

## Thay đổi ngôn ngữ

Go 1.26 giới thiệu hai tinh chỉnh quan trọng cho
[cú pháp và hệ thống kiểu](/doc/go1.26#language) của ngôn ngữ.

Thứ nhất, hàm built-in `new`, dùng để tạo biến mới, nay cho phép toán hạng của nó là một
biểu thức, chỉ định giá trị ban đầu của biến.

Một ví dụ đơn giản về thay đổi này có nghĩa là code như sau:

```go
x := int64(300)
ptr := &x
```

Có thể được đơn giản hóa thành:

```go
ptr := new(int64(300))
```

Thứ hai, kiểu generic nay có thể tự tham chiếu trong danh sách tham số kiểu của chính nó. Thay đổi này
đơn giản hóa việc triển khai các cấu trúc dữ liệu và interface phức tạp.

## Cải thiện hiệu năng

[Bộ gom rác Green Tea](/doc/go1.26#new-garbage-collector) trước đây ở giai đoạn thử nghiệm
nay được bật mặc định.

[Chi phí cgo cơ bản đã được giảm](/doc/go1.26#faster-cgo-calls)
khoảng 30%.

Trình biên dịch nay có thể [cấp phát bộ nhớ backing store](/doc/go1.26#compiler) cho
slice trên stack trong nhiều tình huống hơn, cải thiện hiệu năng.

## Cải tiến công cụ

Lệnh `go fix` đã được viết lại hoàn toàn để dùng
[khung phân tích Go](/pkg/golang.org/x/tools/go/analysis), và nay bao gồm
vài chục "[modernizer](/pkg/golang.org/x/tools/go/analysis/passes/modernize)", các trình phân tích
đề xuất các sửa đổi an toàn giúp code của bạn tận dụng các tính năng mới hơn của ngôn ngữ
và thư viện chuẩn. Nó cũng bao gồm
[trình phân tích `inline`](/pkg/golang.org/x/tools/go/analysis/passes/inline#hdr-Analyzer_inline), cố gắng
nội tuyến tất cả các lời gọi đến từng hàm được chú thích với chỉ thị `//go:fix inline`.
Hai bài đăng blog sắp tới sẽ đề cập các tính năng này chi tiết hơn.

## Thêm cải tiến và thay đổi

Go 1.26 có nhiều cải tiến so với Go 1.25 trong
[các công cụ](/doc/go1.26#tools), [runtime](/doc/go1.26#runtime),
[trình biên dịch](/doc/go1.26#compiler), [trình liên kết](/doc/go1.26#linker),
và [thư viện chuẩn](/doc/go1.26#library).
Bao gồm việc bổ sung ba gói mới: [`crypto/hpke`](/doc/go1.26#new-cryptohpke-package),
[`crypto/mlkem/mlkemtest`](/doc/go1.26#cryptomlkempkgcryptomlkem), và
[`testing/cryptotest`](/doc/go1.26#testingcryptotestpkgtestingcryptotest).
Có các thay đổi [theo nền tảng cụ thể](/doc/go1.26#ports)
và cập nhật [cài đặt `GODEBUG`](/doc/godebug#go-126).

Một số tính năng trong Go 1.26 đang ở giai đoạn thử nghiệm
và chỉ hiển thị khi bạn bật tùy chọn một cách rõ ràng. Đáng chú ý:

- [Gói `simd/archsimd` thử nghiệm](/doc/go1.26#simd) cung cấp quyền truy cập vào các phép tính
"single instruction, multiple data" (SIMD).

- [Gói `runtime/secret` thử nghiệm](/doc/go1.26#new-experimental-runtimesecret-package) cung cấp
cơ chế để xóa an toàn các biến tạm dùng trong code thao tác thông tin bí mật,
thường mang tính chất mật mã.

- [Hồ sơ `goroutineleak` thử nghiệm](/doc/go1.26#goroutineleak-profiles)
trong gói `runtime/pprof` báo cáo các goroutine bị rò rỉ.

Tất cả các thử nghiệm này dự kiến sẽ chính thức sẵn dùng trong
phiên bản Go tương lai. Chúng tôi khuyến khích bạn thử trước.
Phản hồi của bạn rất có giá trị với chúng tôi!

Hãy tham khảo [Ghi chú phát hành Go 1.26](/doc/go1.26) để xem danh sách đầy đủ
các bổ sung, thay đổi và cải tiến trong Go 1.26.

Trong vài tuần tới, các bài đăng blog tiếp theo sẽ đề cập một số chủ đề
liên quan đến Go 1.26 chi tiết hơn. Hãy quay lại sau để đọc các bài đó.

Cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code, báo cáo lỗi,
thử các tính năng thử nghiệm, chia sẻ phản hồi, và kiểm thử các release candidate.
Nỗ lực của bạn giúp Go 1.26 ổn định nhất có thể.
Như thường lệ, nếu bạn phát hiện bất kỳ vấn đề nào, hãy [tạo issue](/issue/new).

Chúc bạn tận hưởng bản phát hành mới!
