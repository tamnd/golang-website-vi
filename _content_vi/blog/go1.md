---
title: Go phiên bản 1 đã được phát hành
date: 2012-03-28
by:
- Andrew Gerrand
tags:
- release
- go1
summary: "Một cột mốc quan trọng: thông báo Go 1, phiên bản ổn định đầu tiên của Go."
template: true
---


{{image "go1/gophermega.jpg"}}

Hôm nay đánh dấu một cột mốc quan trọng trong quá trình phát triển ngôn ngữ lập trình Go.
Chúng tôi thông báo Go phiên bản 1, hay viết tắt là Go 1,
định nghĩa một ngôn ngữ và một tập hợp thư viện lõi để cung cấp nền tảng ổn định
cho việc tạo ra các sản phẩm, dự án và ấn phẩm đáng tin cậy.

Go 1 là bản phát hành Go đầu tiên có sẵn dưới dạng bản phân phối nhị phân được hỗ trợ.
Chúng có sẵn cho Linux, FreeBSD, Mac OS X và,
chúng tôi vui mừng thông báo, Windows.

Động lực chính cho Go 1 là sự ổn định cho người dùng.
Những người viết chương trình Go 1 có thể tin tưởng rằng các chương trình đó sẽ
tiếp tục biên dịch và chạy mà không thay đổi,
trong nhiều môi trường, trong thời gian tính bằng năm.
Tương tự, các tác giả viết sách về Go 1 có thể yên tâm rằng các ví dụ
và giải thích của họ sẽ hữu ích cho người đọc hôm nay và trong tương lai.

Tương thích tiến là một phần của sự ổn định.
Code biên dịch được trong Go 1, với một vài ngoại lệ,
sẽ tiếp tục biên dịch và chạy trong suốt vòng đời của phiên bản đó,
ngay cả khi chúng tôi phát hành các bản cập nhật và sửa lỗi như Go phiên bản 1.1, 1.2 v.v.
[Tài liệu tương thích Go 1](/doc/go1compat.html)
giải thích chi tiết hơn về các hướng dẫn tương thích.

Go 1 là đại diện của Go được dùng ngày nay,
không phải là một thiết kế lại lớn.
Trong quá trình lên kế hoạch, chúng tôi tập trung vào việc dọn dẹp các vấn đề và sự không nhất quán
và cải thiện khả năng di động.
Đã từ lâu có nhiều thay đổi cho Go mà chúng tôi đã thiết kế và tạo prototype
nhưng không phát hành vì chúng không tương thích ngược.
Go 1 kết hợp các thay đổi này, cung cấp các cải tiến đáng kể
cho ngôn ngữ và thư viện nhưng đôi khi tạo ra sự không tương thích với các chương trình cũ.
May mắn thay, công cụ [go fix](/cmd/go/#Run_go_tool_fix_on_packages)
có thể tự động hóa phần lớn công việc cần thiết để đưa chương trình lên chuẩn Go 1.

Go 1 giới thiệu các thay đổi cho ngôn ngữ (chẳng hạn như các kiểu mới cho [ký tự Unicode](/doc/go1.html#rune)
và [lỗi](/doc/go1.html#errors)) và thư viện chuẩn (chẳng hạn như [gói time mới](/doc/go1.html#time)
và đổi tên trong [gói strconv](/doc/go1.html#strconv)).
Ngoài ra, cây gói đã được sắp xếp lại để nhóm các mục liên quan lại với nhau,
chẳng hạn như chuyển các tiện ích mạng,
ví dụ như [gói rpc](/pkg/net/rpc/),
vào các thư mục con của net.
Danh sách đầy đủ các thay đổi được ghi lại trong [ghi chú phát hành Go 1](/doc/go1.html).
Tài liệu đó là tài liệu tham khảo thiết yếu cho các lập trình viên di chuyển code từ
các phiên bản Go trước.

Chúng tôi cũng tái cấu trúc bộ công cụ Go xung quanh [lệnh go](/doc/go1.html#cmd_go) mới,
một chương trình để tải về, build, cài đặt và duy trì code Go.
Lệnh go loại bỏ nhu cầu dùng Makefile để viết code Go vì
nó dùng chính nguồn chương trình Go để suy ra các hướng dẫn build.
Không cần script build nữa!

Cuối cùng, việc phát hành Go 1 kích hoạt một bản phát hành mới của [Google App Engine SDK](https://developers.google.com/appengine/docs/go).
Một quy trình tương tự về sửa đổi và ổn định đã được áp dụng cho
các thư viện App Engine,
cung cấp nền tảng cho các lập trình viên build chương trình cho App Engine sẽ chạy trong nhiều năm.

Go 1 là kết quả của nỗ lực lớn từ nhóm Go lõi và những người đóng góp của chúng tôi
từ cộng đồng mã nguồn mở.
Chúng tôi cảm ơn tất cả những người đã giúp thực hiện điều này.

Chưa bao giờ có thời điểm tốt hơn để trở thành lập trình viên Go.
Mọi thứ bạn cần để bắt đầu đều có tại [golang.org](/).
