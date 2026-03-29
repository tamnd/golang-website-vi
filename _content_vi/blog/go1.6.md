---
title: Go 1.6 đã được phát hành
date: 2016-02-17
by:
- Andrew Gerrand
summary: Go 1.6 bổ sung HTTP/2, template block, và nhiều hơn nữa.
template: true
---


Hôm nay chúng tôi phát hành [Go phiên bản 1.6](/doc/go1.6),
bản phát hành ổn định lớn thứ bảy của Go.
Bạn có thể tải ngay từ [trang download](/dl/).
Mặc dù [bản phát hành Go 1.5](/blog/go1.5) sáu tháng trước
chứa những thay đổi triển khai đáng kể,
bản phát hành này mang tính chất tăng dần hơn.

Thay đổi quan trọng nhất là hỗ trợ [HTTP/2](https://http2.github.io/)
trong [gói net/http](/pkg/net/http/).
HTTP/2 là một giao thức mới, tiếp nối của HTTP và đã được
các nhà sản xuất trình duyệt và các trang web lớn áp dụng rộng rãi.
Trong Go 1.6, hỗ trợ HTTP/2 được [bật theo mặc định](/doc/go1.6#http2)
cho cả máy chủ lẫn máy khách khi dùng HTTPS,
mang lại [các lợi ích](https://http2.github.io/faq/) của giao thức mới
cho nhiều dự án Go,
chẳng hạn như [web server Caddy](https://caddyserver.com/download) nổi tiếng.

Các gói template đã học một số thủ thuật mới,
với hỗ trợ [cắt khoảng trắng xung quanh template action](/pkg/text/template/#hdr-Text_and_spaces)
để tạo ra đầu ra template gọn hơn,
và việc giới thiệu {{raw "[`{{block}}` action]"}}(/pkg/text/template/#hdr-Actions)
có thể dùng để tạo template kế thừa từ template khác.
Một [chương trình mẫu template mới](https://cs.opensource.google/go/x/example/+/master:template) minh họa các tính năng mới này.

Go 1.5 đã giới thiệu [hỗ trợ thử nghiệm](/s/go15vendor)
cho thư mục "vendor" được bật bằng biến môi trường.
Trong Go 1.6, tính năng này nay được [bật theo mặc định](/doc/go1.6#go_command).
Các cây mã nguồn chứa thư mục tên "vendor" không được dùng theo cách đúng với tính năng mới
sẽ cần thay đổi để tránh build bị lỗi (cách sửa đơn giản nhất là đổi tên thư mục).

Runtime đã thêm tính năng phát hiện nhẹ, cố gắng hết sức về việc dùng map đồng thời không đúng cách.
Như thường lệ, nếu một goroutine đang ghi vào map, không goroutine nào khác nên đọc hoặc ghi vào map đó đồng thời.
Nếu runtime phát hiện điều kiện này, nó in một chẩn đoán và crash chương trình.
Cách tốt nhất để tìm hiểu thêm về vấn đề là chạy nó dưới
[bộ phát hiện race](/blog/race-detector),
sẽ xác định race đáng tin cậy hơn và cung cấp thêm chi tiết.

Runtime cũng đã thay đổi cách in panic kết thúc chương trình.
Nó nay chỉ in stack của goroutine bị panic, thay vì tất cả goroutine đang tồn tại.
Hành vi này có thể được cấu hình dùng biến môi trường
[GOTRACEBACK](/pkg/runtime/#hdr-Environment_Variables)
hoặc bằng cách gọi hàm [debug.SetTraceback](/pkg/runtime/debug/#SetTraceback).

Người dùng cgo nên biết về những thay đổi lớn đối với các quy tắc chia sẻ con trỏ giữa code Go và C.
Các quy tắc được thiết kế để đảm bảo rằng code C đó có thể cùng tồn tại với bộ gom rác của Go
và được kiểm tra trong quá trình thực thi chương trình, vì vậy code có thể cần thay đổi để tránh crash.
Xem [ghi chú phát hành](/doc/go1.6#cgo) và
[tài liệu cgo](/cmd/cgo/#hdr-Passing_pointers) để biết chi tiết.

Trình biên dịch, trình liên kết và lệnh go có cờ `-msan` mới
tương tự như `-race` và chỉ có sẵn trên linux/amd64,
cho phép tương tác với
[Clang MemorySanitizer](http://clang.llvm.org/docs/MemorySanitizer.html).
Điều này hữu ích để kiểm thử chương trình chứa code C hoặc C++ đáng ngờ.
Bạn có thể muốn thử trong khi kiểm thử code cgo của mình với các quy tắc con trỏ mới.

Hiệu năng của các chương trình Go được build với Go 1.6 vẫn tương tự như những chương trình được build với Go 1.5.
Thời gian tạm dừng thu gom rác thậm chí còn thấp hơn so với Go 1.5,
nhưng điều này đặc biệt đáng chú ý đối với các chương trình dùng lượng lớn bộ nhớ.
Về hiệu năng của toolchain trình biên dịch,
thời gian build nên tương tự như Go 1.5.

Thuật toán bên trong [sort.Sort](/pkg/sort/#Sort)
đã được cải thiện để chạy nhanh hơn khoảng 10%,
nhưng thay đổi có thể làm hỏng các chương trình kỳ vọng một thứ tự cụ thể
của các phần tử bằng nhau nhưng có thể phân biệt được.
Các chương trình như vậy nên tinh chỉnh phương thức `Less` để chỉ thứ tự mong muốn
hoặc dùng [sort.Stable](/pkg/sort/#Stable)
để bảo toàn thứ tự đầu vào cho các giá trị bằng nhau.

Và tất nhiên, có nhiều bổ sung, cải tiến và sửa lỗi hơn nữa.
Bạn có thể tìm thấy tất cả trong [ghi chú phát hành](/doc/go1.6) toàn diện.

Để kỷ niệm bản phát hành,
[Nhóm người dùng Go trên toàn thế giới](/wiki/Go-1.6-release-party)
đang tổ chức tiệc phát hành vào ngày 17 tháng 2.
Trực tuyến, những người đóng góp Go đang tổ chức phiên hỏi đáp
trên [golang subreddit](https://reddit.com/r/golang) trong 24 giờ tới.
Nếu bạn có câu hỏi về dự án, bản phát hành, hay chỉ về Go nói chung,
hãy [tham gia thảo luận](https://www.reddit.com/r/golang/comments/46bd5h/ama_we_are_the_go_contributors_ask_us_anything/).

Cảm ơn tất cả những người đã đóng góp cho bản phát hành.
Chúc lập trình vui vẻ.
