---
title: Go 1.19 đã được phát hành!
date: 2022-08-02
by:
- The Go Team
summary: Go 1.19 bổ sung comment tài liệu phong phú hơn, cải tiến hiệu năng, và nhiều hơn nữa.
---

Hôm nay nhóm Go vô cùng hào hứng phát hành Go 1.19,
bạn có thể tải về từ [trang download](/dl/).

Go 1.19 tinh chỉnh và cải thiện bản [Go 1.18](/blog/go1.18) lớn của chúng tôi hồi đầu năm.
Chúng tôi tập trung phát triển generics trong Go 1.19 vào việc giải quyết các vấn đề tinh tế
và các trường hợp biên mà cộng đồng đã báo cáo,
cũng như các cải tiến hiệu năng quan trọng (lên đến 20% với một số chương trình generic).

Doc comment nay hỗ trợ [liên kết, danh sách và cú pháp heading rõ ràng hơn](/doc/comment).
Thay đổi này giúp người dùng viết doc comment rõ ràng và dễ điều hướng hơn,
đặc biệt trong các gói có API lớn.
Là một phần của thay đổi này, `gofmt` nay định dạng lại doc comment để áp dụng
định dạng chuẩn cho cách dùng các tính năng này.
Xem "[Go Doc Comments](/doc/comment)" để biết toàn bộ chi tiết.

[Mô hình bộ nhớ của Go](/ref/mem) nay định nghĩa rõ ràng
hành vi của [gói sync/atomic](/pkg/sync/atomic/).
Định nghĩa chính thức về quan hệ happens-before đã được sửa lại
để phù hợp với các mô hình bộ nhớ dùng bởi C, C++, Java, JavaScript, Rust và Swift.
Các chương trình hiện có không bị ảnh hưởng.
Cùng với cập nhật mô hình bộ nhớ, có thêm
[các kiểu mới trong gói sync/atomic](/doc/go1.19#atomic_types),
chẳng hạn như [atomic.Int64](/pkg/sync/atomic/#Int64) và [atomic.Pointer[T]](/pkg/sync/atomic/#Pointer),
để dùng giá trị atomic dễ dàng hơn.

Vì [lý do bảo mật](/blog/path-security), gói os/exec
không còn xem xét các đường dẫn tương đối trong tìm kiếm PATH.
Xem [tài liệu gói](/pkg/os/exec/#hdr-Executables_in_the_current_directory)
để biết chi tiết.
Các cách dùng hiện có của [golang.org/x/sys/execabs](https://pkg.go.dev/golang.org/x/sys/execabs)
có thể chuyển lại về os/exec trong các chương trình chỉ build dùng Go 1.19 hoặc mới hơn.

Bộ gom rác đã thêm hỗ trợ cho giới hạn bộ nhớ mềm,
được thảo luận chi tiết trong [hướng dẫn thu gom rác mới](/doc/gc-guide#Memory_limit).
Giới hạn này đặc biệt hữu ích để tối ưu hóa các chương trình Go
chạy hiệu quả nhất có thể trong container với lượng bộ nhớ cố định.

Ràng buộc build mới `unix` được thỏa mãn khi hệ điều hành đích (`GOOS`)
là bất kỳ hệ thống giống Unix nào.
Hiện tại, giống Unix có nghĩa là tất cả hệ điều hành đích của Go trừ `js`, `plan9`, `windows` và `zos`.

Cuối cùng, Go 1.19 bao gồm nhiều cải tiến hiệu năng và triển khai, bao gồm
định kích thước động cho stack goroutine ban đầu để giảm sao chép stack,
tự động sử dụng thêm file descriptor trên hầu hết các hệ thống Unix,
jump table cho lệnh switch lớn trên x86-64 và ARM64,
hỗ trợ gọi hàm inject bởi debugger trên ARM64,
hỗ trợ register ABI trên RISC-V,
và hỗ trợ thử nghiệm cho
Linux chạy trên kiến trúc 64-bit Loongson LoongArch (`GOARCH=loong64`).

Cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code, báo cáo lỗi, chia sẻ phản hồi,
và kiểm thử bản beta và release candidate.
Nỗ lực của bạn giúp đảm bảo Go 1.19 ổn định nhất có thể.
Như thường lệ, nếu bạn phát hiện bất kỳ vấn đề nào, hãy [tạo issue](/issue/new).

Tận hưởng Go 1.19!
