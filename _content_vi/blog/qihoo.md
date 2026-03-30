---
title: Qihoo 360 và Go
date: 2015-07-06
by:
- Yang Zhou
summary: Cách Qihoo 360 sử dụng Go.
template: true
---


_Bài đăng khách này được viết bởi Yang Zhou, Kỹ sư Phần mềm tại Qihoo 360._

[Qihoo 360](http://www.360safe.com/) là nhà cung cấp lớn các sản phẩm và
dịch vụ bảo mật Internet và di động tại Trung Quốc, và vận hành một nền tảng
phân phối di động dựa trên Android lớn. Vào cuối tháng 6 năm 2014, Qihoo có
khoảng 500 triệu người dùng Internet PC hoạt động hàng tháng và hơn 640 triệu
người dùng di động. Qihoo cũng vận hành một trong những trình duyệt Internet
PC và công cụ tìm kiếm PC phổ biến nhất của Trung Quốc.

Nhóm của tôi, Nhóm Dịch vụ Push, cung cấp các dịch vụ nhắn tin cơ bản cho
hơn 50 sản phẩm trong toàn công ty (cả PC lẫn di động), bao gồm hàng nghìn
ứng dụng trong nền tảng mở của chúng tôi.

"Chuyện tình" của chúng tôi với Go bắt đầu vào năm 2012 khi chúng tôi lần
đầu tiên cố gắng cung cấp dịch vụ push cho một trong các sản phẩm của Qihoo.
Phiên bản ban đầu được xây dựng với nginx + lua + redis, không đáp ứng được
yêu cầu về hiệu năng thời gian thực do tải quá cao. Trong hoàn cảnh đó, bản
phát hành Go 1.0.3 mới được chú ý đến. Chúng tôi đã hoàn thành một prototype
trong vài tuần, phần lớn nhờ vào các tính năng goroutine và channel mà nó cung
cấp.

Ban đầu, hệ thống dựa trên Go của chúng tôi chạy trên 20 máy chủ, với tổng
cộng 20 triệu kết nối thời gian thực. Hệ thống gửi 2 triệu tin nhắn mỗi ngày.
Hệ thống đó hiện chạy trên 400 máy chủ, hỗ trợ hơn 200 triệu kết nối thời
gian thực. Nó hiện gửi hơn 10 tỷ tin nhắn hàng ngày.

Với sự mở rộng kinh doanh nhanh chóng và nhu cầu ứng dụng ngày càng tăng cho
dịch vụ push của chúng tôi, hệ thống Go ban đầu nhanh chóng đạt đến điểm tắc
nghẽn: kích thước heap lên đến 69G, với thời gian dừng garbage collection (GC)
tối đa là 3-6 giây. Tệ hơn nữa, chúng tôi phải khởi động lại hệ thống mỗi
tuần để giải phóng bộ nhớ. Sẽ không trung thực nếu chúng tôi không thừa nhận
đã xem xét từ bỏ Go và thay vào đó viết lại toàn bộ thành phần cốt lõi bằng C.
Tuy nhiên, mọi thứ không diễn ra đúng như kế hoạch, chúng tôi gặp khó khăn
khi di chuyển code của Business Logic Layer. Kết quả là, nhân lực duy nhất
lúc đó (tôi) không thể vừa duy trì hệ thống Go vừa đảm bảo việc chuyển logic
sang framework dịch vụ C.

Vì vậy, tôi đã đưa ra quyết định ở lại với hệ thống Go (có lẽ là quyết định
khôn ngoan nhất tôi phải đưa ra), và chúng tôi đã đạt được nhiều tiến bộ
đáng kể.

Đây là một số điều chỉnh chúng tôi đã thực hiện và các bài học quan trọng:

  - Thay thế các kết nối ngắn bằng các kết nối persistent (sử dụng connection
    pool), để giảm việc tạo buffer và đối tượng trong quá trình giao tiếp.
  - Sử dụng Object Pool và Memory Pool một cách phù hợp, để giảm tải cho GC.

{{image "qihoo/image00.png"}}

  - Sử dụng Task Pool, một cơ chế với một nhóm goroutine chạy lâu dài tiêu
    thụ các hàng đợi task hoặc message toàn cục được gửi bởi các goroutine
    kết nối, để thay thế các goroutine ngắn hạn.

  - Giám sát và kiểm soát số lượng goroutine trong chương trình. Việc thiếu
    kiểm soát có thể gây ra gánh nặng không thể chịu đựng cho GC, do sự tăng
    đột biến của goroutine vì việc chấp nhận yêu cầu bên ngoài không bị hạn
    chế, vì các lời gọi RPC gửi đến các máy chủ bên trong có thể chặn các
    goroutine mới được tạo.

  - Nhớ thêm [deadline đọc và ghi](/pkg/net/#Conn)
    cho các kết nối khi ở mạng di động; nếu không, có thể dẫn đến chặn
    goroutine. Áp dụng nó đúng cách và thận trọng khi ở mạng LAN, nếu không
    hiệu quả giao tiếp RPC của bạn sẽ bị ảnh hưởng.

  - Sử dụng Pipeline (trong tính năng Full Duplex của TCP) để nâng cao hiệu
    quả giao tiếp của framework RPC.

Kết quả là, chúng tôi đã thành công triển khai ba lần lặp kiến trúc và hai
lần lặp framework RPC ngay cả với nguồn nhân lực hạn chế. Tất cả điều này đều
được quy về sự tiện lợi phát triển của Go. Bên dưới bạn có thể tìm thấy kiến
trúc hệ thống mới nhất:

{{image "qihoo/image01.png"}}

Hành trình cải tiến liên tục có thể được minh họa bằng một bảng:

{{image "qihoo/table.png"}}

Ngoài ra, không cần giải phóng bộ nhớ tạm thời hay khởi động lại hệ thống
sau các tối ưu hóa này.

Điều thú vị hơn nữa là chúng tôi đã phát triển một Nền tảng Hiển thị thời
gian thực trực tuyến để phân tích profiling các chương trình Go. Chúng tôi
hiện có thể dễ dàng truy cập và chẩn đoán trạng thái hệ thống, xác định bất
kỳ rủi ro tiềm ẩn nào. Đây là ảnh chụp màn hình của hệ thống đang hoạt động:

{{image "qihoo/image02.png"}}
{{image "qihoo/image03.png"}}

Điều tuyệt vời về nền tảng này là chúng tôi thực sự có thể mô phỏng kết nối
và hành vi của hàng triệu người dùng trực tuyến, bằng cách áp dụng Công cụ
Kiểm tra Căng thẳng Phân tán (cũng được xây dựng bằng Go), và quan sát tất
cả dữ liệu được trực quan hóa theo thời gian thực. Điều này cho phép chúng
tôi đánh giá hiệu quả của bất kỳ tối ưu hóa nào và ngăn chặn vấn đề bằng
cách xác định các điểm tắc nghẽn của hệ thống.

Hầu hết mọi tối ưu hóa hệ thống có thể đã được thực hành cho đến nay. Và
chúng tôi mong đợi thêm tin tốt từ nhóm GC để chúng tôi có thể được giảm
bớt thêm khỏi công việc phát triển nặng nề. Tôi đoán kinh nghiệm của chúng
tôi cũng có thể trở nên lỗi thời một ngày nào đó, khi Go tiếp tục phát triển.

Đây là lý do tôi muốn kết thúc phần chia sẻ của mình bằng cách bày tỏ lòng
biết ơn chân thành đến cơ hội được tham dự [Gopher China](http://gopherchina.org/).
Đó là một sự kiện hội tụ để chúng tôi học hỏi, chia sẻ và là cơ hội nhìn thấy
sự phổ biến và thịnh vượng của Go tại Trung Quốc. Nhiều nhóm khác trong Qihoo
đã biết đến Go, hoặc đã thử sử dụng Go.

Tôi tin chắc rằng nhiều công ty Internet Trung Quốc hơn sẽ tham gia cùng
chúng tôi trong việc tái tạo hệ thống của họ trong Go và những nỗ lực của
nhóm Go sẽ mang lại lợi ích cho nhiều nhà phát triển và doanh nghiệp hơn
trong tương lai có thể thấy trước.
