---
title: Go 1.8 đã được phát hành
date: 2017-02-16
by:
- Chris Broadfoot
summary: Go 1.8 bổ sung code non-x86 được biên dịch nhanh hơn, thời gian tạm dừng GC dưới mili giây, HTTP/2 push, và nhiều hơn nữa.
template: true
---


Hôm nay nhóm Go vui mừng thông báo phát hành Go 1.8.
Bạn có thể tải về từ [trang download](/dl/).
Có những cải tiến hiệu năng đáng kể và thay đổi trên thư viện chuẩn.

Backend trình biên dịch được giới thiệu trong [Go 1.7](/blog/go1.7) cho x86 64-bit nay được dùng
trên tất cả kiến trúc, và các kiến trúc đó nên thấy [cải tiến hiệu năng đáng kể](/doc/go1.8#compiler).
Ví dụ, thời gian CPU yêu cầu bởi các chương trình benchmark của chúng tôi đã giảm 20-30% trên hệ thống ARM 32-bit.
Cũng có một số cải tiến hiệu năng khiêm tốn trong bản phát hành này cho hệ thống x86 64-bit.
Trình biên dịch và trình liên kết đã được làm nhanh hơn.
Thời gian biên dịch nên được cải thiện khoảng 15% so với Go 1.7.
Vẫn còn nhiều công việc cần làm trong lĩnh vực này: kỳ vọng tốc độ biên dịch nhanh hơn trong các bản phát hành tương lai.

Thời gian tạm dừng thu gom rác nên [ngắn hơn đáng kể](/doc/go1.8#gc),
thường dưới 100 micro giây và thường thấp tới 10 micro giây.

HTTP server thêm hỗ trợ [HTTP/2 Push](/doc/go1.8#h2push),
cho phép server gửi trước phản hồi đến client.
Điều này hữu ích để giảm thiểu độ trễ mạng bằng cách loại bỏ các vòng đi lại.
HTTP server cũng thêm hỗ trợ [tắt graceful](/doc/go1.8#http_shutdown),
cho phép server giảm thiểu thời gian ngừng hoạt động bằng cách chỉ tắt sau khi phục vụ tất cả các request đang thực hiện.

[Context](/pkg/context/) (được thêm vào thư viện chuẩn trong Go 1.7)
cung cấp cơ chế hủy bỏ và timeout.
Go 1.8 [thêm](/doc/go1.8#more_context) hỗ trợ context trong nhiều phần hơn của thư viện chuẩn,
bao gồm gói [`database/sql`](/pkg/database/sql) và [`net`](/pkg/net)
và [`Server.Shutdown`](http://beta.golang.org/pkg/net/http/#Server.Shutdown) trong gói `net/http`.

Bây giờ đơn giản hơn nhiều để sắp xếp slice dùng hàm [`Slice`](/pkg/sort/#Slice) mới được thêm vào
trong gói `sort`. Ví dụ, để sắp xếp một slice struct theo trường `Name`:

{{raw `
	sort.Slice(s, func(i, j int) bool { return s[i].Name < s[j].Name })
`}}

Go 1.8 bao gồm nhiều bổ sung, cải tiến và sửa lỗi hơn nữa.
Tìm tập thay đổi đầy đủ, và thêm thông tin về các cải tiến được liệt kê ở trên, trong
[ghi chú phát hành Go 1.8](/doc/go1.8.html).

Để kỷ niệm bản phát hành, Nhóm người dùng Go trên toàn thế giới đang tổ chức [tiệc phát hành](/wiki/Go-1.8-release-party) tuần này.
Tiệc phát hành đã trở thành truyền thống trong cộng đồng Go, vì vậy nếu bạn lỡ lần này, hãy theo dõi khi 1.9 đến gần.

Cảm ơn hơn 200 người đóng góp đã giúp đỡ trong bản phát hành này.
