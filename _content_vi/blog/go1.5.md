---
title: Go 1.5 đã được phát hành
date: 2015-08-19
by:
- Andrew Gerrand
summary: Go 1.5 bổ sung bộ gom rác mới nhanh hơn nhiều, nhiều song song hơn theo mặc định, go tool trace, và nhiều hơn nữa.
---


Hôm nay dự án Go tự hào phát hành Go 1.5,
bản phát hành ổn định lớn thứ sáu của Go.

Bản phát hành này bao gồm những thay đổi đáng kể đối với việc triển khai.
Toolchain của trình biên dịch đã được [dịch từ C sang Go](/doc/go1.5#c),
loại bỏ dấu vết cuối cùng của code C khỏi codebase Go.
Bộ gom rác đã được [thiết kế lại hoàn toàn](/doc/go1.5#gc),
mang lại [sự giảm đáng kể](/talks/2015/go-gc.pdf)
trong thời gian tạm dừng thu gom rác.
Các cải tiến liên quan đến bộ lập lịch cho phép chúng tôi thay đổi giá trị mặc định
[GOMAXPROCS](/pkg/runtime/#GOMAXPROCS)
(số goroutine thực thi đồng thời)
từ 1 thành số lượng CPU logic.
Các thay đổi đối với trình liên kết cho phép phân phối các gói Go như thư viện chia sẻ để
liên kết vào các chương trình Go, và build các gói Go thành archive hoặc thư viện chia sẻ
có thể được liên kết vào hoặc tải bởi các chương trình C
([tài liệu thiết kế](/s/execmodes)).

Bản phát hành cũng bao gồm [cải tiến cho các công cụ lập trình viên](/doc/go1.5#go_command).
Hỗ trợ cho [gói "internal"](/s/go14internal)
cho phép chia sẻ chi tiết triển khai giữa các gói.
[Hỗ trợ thử nghiệm](/s/go15vendor) cho "vendoring"
các dependency bên ngoài là một bước tiến tới một cơ chế chuẩn để quản lý
dependency trong các chương trình Go.
Lệnh "[go tool trace](/cmd/trace/)" mới cho phép
trực quan hóa các trace chương trình được tạo ra bởi cơ sở hạ tầng tracing mới trong
runtime.
Lệnh "[go doc](/cmd/go/#hdr-Show_documentation_for_package_or_symbol)" mới
cung cấp giao diện dòng lệnh cải tiến để xem tài liệu gói Go.

Cũng có thêm [cổng hệ điều hành và kiến trúc mới](/doc/go1.5#ports).
Các cổng mới trưởng thành hơn là darwin/arm,
darwin/arm64 (thiết bị iPhone và iPad của Apple),
và linux/arm64.
Ngoài ra còn có hỗ trợ thử nghiệm cho ppc64 và ppc64le
(IBM 64-bit PowerPC, big và little endian).

Cổng darwin/arm64 mới và tính năng external linking cung cấp nền tảng cho
[dự án Go mobile](https://godoc.org/golang.org/x/mobile), một thử nghiệm để
xem Go có thể được dùng như thế nào để build ứng dụng trên thiết bị Android và iOS.
(Bản thân công việc Go mobile không phải là một phần của bản phát hành này.)

Thay đổi ngôn ngữ duy nhất rất nhỏ,
[nới lỏng một hạn chế trong cú pháp map literal](/doc/go1.5#language)
để làm cho chúng súc tích và nhất quán hơn với slice literal.

Thư viện chuẩn cũng có nhiều bổ sung và cải tiến.
Gói flag nay hiển thị [thông báo sử dụng gọn hơn](/doc/go1.5#flag).
Gói math/big nay cung cấp kiểu [Float](/pkg/math/big/#Float)
để tính toán với số dấu phẩy động có độ chính xác tùy ý.
[Cải tiến](/doc/go1.5#net) đối với trình phân giải DNS trên
hệ thống Linux và BSD đã loại bỏ yêu cầu cgo cho các chương trình tra cứu tên.
Gói [go/types](/pkg/go/types/) đã được
[chuyển](/doc/go1.5#go_types) vào thư viện chuẩn từ
kho lưu trữ [golang.org/x/tools](https://godoc.org/golang.org/x/tools).
(Các gói [go/constant](/pkg/go/constant/) và
[go/importer](/pkg/go/importer/) mới cũng là kết quả
của việc chuyển này.)
Gói reflect đã thêm các hàm
[ArrayOf](/pkg/reflect/#ArrayOf) và
[FuncOf](/pkg/reflect/#FuncOf), tương tự như
hàm [SliceOf](/pkg/reflect/#SliceOf) hiện có.
Và tất nhiên, có danh sách thông thường về
[các sửa lỗi và cải tiến nhỏ hơn](/doc/go1.5#minor_library_changes).

Để xem toàn bộ câu chuyện, xem [ghi chú phát hành chi tiết](/doc/go1.5).
Hoặc nếu bạn không thể chờ để bắt đầu,
hãy đến [trang download](/dl/) để tải Go 1.5 ngay bây giờ.
