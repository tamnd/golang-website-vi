---
title: Binary nhỏ hơn trong Go 1.7
date: 2016-08-18
by:
- David Crawshaw
summary: Go 1.7 có một số cải tiến giảm kích thước binary quan trọng cho các thiết bị nhỏ.
template: true
---

## Giới thiệu

Go được thiết kế để viết server.
Đó là cách nó được sử dụng phổ biến nhất ngày nay, và kết quả là rất nhiều
công việc trên runtime và trình biên dịch tập trung vào các vấn đề quan trọng với server:
độ trễ, triển khai dễ dàng, thu gom rác chính xác,
thời gian khởi động nhanh, hiệu năng.

Khi Go được dùng cho nhiều loại chương trình hơn, có những vấn đề mới cần được xem xét.
Một trong số đó là kích thước binary.
Nó đã nằm trong danh sách theo dõi từ lâu
(issue [\#6853](/issue/6853) được tạo hơn hai
năm trước), nhưng sự quan tâm ngày càng tăng đến việc dùng Go để
triển khai binary trên các thiết bị nhỏ hơn như Raspberry Pi hay
thiết bị di động, có nghĩa là nó được chú ý cho bản phát hành Go 1.7.

## Công việc được thực hiện trong Go 1.7

Ba thay đổi quan trọng trong Go 1.7 ảnh hưởng đến kích thước binary.

Thứ nhất là backend SSA mới được bật cho AMD64 trong bản phát hành này.
Mặc dù động lực chính cho SSA là hiệu năng được cải thiện, code được tạo ra
tốt hơn cũng nhỏ hơn.
Backend SSA thu nhỏ binary Go khoảng 5%.
Chúng tôi kỳ vọng lợi ích lớn hơn cho các kiến trúc kiểu RISC hơn
như ARM và MIPS khi các backend đó được chuyển đổi sang SSA trong Go 1.8.

Thay đổi thứ hai là cắt tỉa phương thức.
Cho đến 1.6, tất cả các phương thức trên tất cả các kiểu được sử dụng đều được giữ lại, ngay cả khi một số
phương thức không bao giờ được gọi.
Điều này là vì chúng có thể được gọi qua interface, hoặc được gọi
động bằng gói reflect.
Nay trình biên dịch loại bỏ bất kỳ phương thức không xuất nào không khớp với
interface nào.
Tương tự trình liên kết có thể loại bỏ các phương thức xuất khác, những phương thức chỉ
có thể truy cập qua reflection, nếu các
[tính năng reflection tương ứng](/pkg/reflect/#Value.Call)
không được dùng ở bất kỳ đâu trong chương trình.
Thay đổi đó thu nhỏ binary 5-20%.

Thay đổi thứ ba là định dạng gọn hơn cho thông tin kiểu runtime
được dùng bởi gói reflect.
Định dạng mã hóa ban đầu được thiết kế để làm cho bộ giải mã trong
runtime và gói reflect đơn giản nhất có thể. Bằng cách làm code này
khó đọc hơn một chút, chúng tôi có thể nén định dạng mà không ảnh hưởng đến
hiệu năng runtime của các chương trình Go.
Định dạng mới thu nhỏ binary Go thêm 5-15%.
Các thư viện được build cho Android và archive được build cho iOS thu nhỏ thêm nữa
vì định dạng mới chứa ít con trỏ hơn, mỗi con trỏ đòi hỏi
dynamic relocation trong position independent code.

Ngoài ra, có nhiều cải tiến nhỏ như cải thiện bố cục dữ liệu interface,
bố cục dữ liệu tĩnh tốt hơn, và đơn giản hóa các dependency. Ví dụ, HTTP client
không còn liên kết toàn bộ HTTP server.
Danh sách đầy đủ các thay đổi có thể tìm thấy trong issue
[\#6853](/issue/6853).

## Kết quả

Các chương trình điển hình, từ các ví dụ nhỏ đến chương trình production lớn,
nhỏ hơn khoảng 30% khi được build với Go 1.7.

Chương trình hello world tiêu chuẩn giảm từ 2,3MB xuống 1,6MB:

	package main

	import "fmt"

	func main() {
		fmt.Println("Hello, World!")
	}

Khi biên dịch không có thông tin debug, binary liên kết tĩnh
nay dưới một megabyte.

{{image "go1.7-binary-size/graph.png"}}

Một chương trình production lớn được dùng để kiểm thử trong chu kỳ này, `jujud`, giảm từ 94MB
xuống 67MB.

Binary position-independent nhỏ hơn 50%.

Trong position-independent executable (PIE), con trỏ trong section dữ liệu chỉ đọc
đòi hỏi dynamic relocation.
Vì định dạng mới cho thông tin kiểu thay thế con trỏ bằng
section offset, nó tiết kiệm 28 byte cho mỗi con trỏ.

Position-independent executable với thông tin debug đã xóa
đặc biệt quan trọng với các nhà phát triển di động, vì đây là loại
chương trình được gửi đến điện thoại.
Tải xuống lớn tạo ra trải nghiệm người dùng kém, vì vậy việc giảm ở đây
là tin tốt.

## Công việc trong tương lai

Một số thay đổi đối với thông tin kiểu runtime đã quá muộn cho
đóng băng Go 1.7, nhưng hy vọng sẽ đưa vào 1.8, thu nhỏ thêm
các chương trình, đặc biệt là position-independent.

Tất cả các thay đổi này đều thận trọng, giảm kích thước binary mà không tăng
thời gian build, thời gian khởi động, thời gian thực thi tổng thể, hoặc sử dụng bộ nhớ.
Chúng ta có thể thực hiện các bước triệt để hơn để giảm kích thước binary: công cụ
[upx](http://upx.sourceforge.net/) để nén executable
thu nhỏ binary thêm 50% với chi phí tăng thời gian khởi động
và có thể tăng sử dụng bộ nhớ.
Đối với các hệ thống cực nhỏ (loại có thể tồn tại trên dây chìa khóa)
chúng ta có thể build phiên bản Go không có reflection, mặc dù
không rõ liệu ngôn ngữ bị hạn chế như vậy có đủ hữu ích không.
Đối với một số thuật toán trong runtime, chúng ta có thể dùng các triển khai chậm hơn nhưng gọn hơn
khi mỗi kilobyte có giá trị.
Tất cả những điều này đòi hỏi nhiều nghiên cứu hơn trong các chu kỳ phát triển sau.

Cảm ơn nhiều người đóng góp đã giúp làm cho binary Go 1.7 nhỏ hơn!
