---
title: Báo cáo chuyến đi GopherChina
date: 2015-07-01
by:
- Robert Griesemer
tags:
- community
- china
summary: Báo cáo từ GopherChina 2015, hội nghị Go đầu tiên tại Trung Quốc.
template: true
---


Chúng tôi đã biết trong một thời gian rằng Go phổ biến ở Trung Quốc hơn bất kỳ quốc gia nào khác.
Theo Google Trends, hầu hết [các tìm kiếm từ khóa "golang"](https://www.google.com/trends/explore#q=golang) đến từ Cộng hòa Nhân dân hơn bất kỳ nơi nào khác.
[Những người khác](http://herman.asia/why-is-go-popular-in-china) cũng đã suy đoán về
cùng quan sát này, tuy nhiên cho đến nay chúng tôi có
[thông tin cụ thể ít ỏi](https://news.ycombinator.com/item?id=8872400)
về hiện tượng này.

Hội nghị Go đầu tiên ở Trung Quốc, [GopherChina](http://gopherchina.org/),
có vẻ là cơ hội tuyệt vời để khám phá tình hình bằng cách đưa một số Gopher phương Tây đặt chân lên đất Trung Quốc. Lời mời thực sự đã biến điều đó thành hiện thực và tôi
quyết định chấp nhận và thuyết trình về tác động của gofmt đến phát triển phần mềm.

{{image "gopherchina/image04.jpg"}}

_Xin chào, Thượng Hải!_

Hội nghị diễn ra trong một cuối tuần tháng 4 ở Thượng Hải, trong
[Tòa nhà Puruan](https://www.google.com/maps/place/Puruan+Bldg,+Pudong,+Shanghai,+China)
của Khu Công nghiệp Phần mềm Pudong Thượng Hải, dễ dàng tiếp cận bằng tàu điện ngầm trong vòng một giờ
hoặc ít hơn từ các khu trung tâm hơn của Thượng Hải.
Được mô phỏng theo [GopherCon](http://www.gophercon.com), hội nghị là
một track duy nhất, với tất cả các bài thuyết trình được trình bày trong một phòng hội nghị chứa khoảng 400
người tham dự.
Nó được tổ chức bởi các tình nguyện viên, dẫn đầu bởi [Asta Xie](https://github.com/astaxie),
và với sự bảo trợ mạnh mẽ từ các tên tuổi công nghiệp lớn. Theo
các ban tổ chức, nhiều người hơn muốn tham dự nhưng không thể được tiếp nhận
do hạn chế về không gian.

{{image "gopherchina/image01.jpg"}}

_Ủy ban chào đón với Asta Xie (thứ 2 từ trái), người tổ chức chính._

Mỗi người tham dự nhận được một túi đầy áo phông GopherChina bắt buộc,
các tờ thông tin quảng cáo liên quan đến nhà tài trợ, nhãn dán và đôi khi
có "gì đó" nhồi bông (không có Gopher mềm, tuy nhiên). Ít nhất một nhà cung cấp bên thứ ba
đang quảng cáo sách kỹ thuật, bao gồm một số sách Go gốc (không dịch
từ tiếng Anh).

{{image "gopherchina/image05.jpg"}}

_Sách Go!_

Ấn tượng đầu tiên, người tham dự trung bình có vẻ khá trẻ, tạo ra một đám đông nhiệt tình, và sự kiện có vẻ được tổ chức tốt.

Ngoại trừ bài thuyết trình của tôi, tất cả các bài thuyết trình đều được thực hiện bằng tiếng Quan Thoại và
do đó không thể hiểu được với tôi. Asta Xie, người tổ chức chính, đã hỗ trợ
với một số bản dịch đồng thời thì thầm vào tai tôi, và các slide tiếng Anh thỉnh thoảng cung cấp thêm manh mối: "69GB" nổi bật ngay cả khi không có kiến thức tiếng Quan Thoại nào (thêm về điều đó bên dưới). Do đó, cuối cùng tôi chỉ nghe một số ít bài thuyết trình, và thay vào đó dành nhiều thời gian nói chuyện với
những người tham dự bên ngoài phòng hội nghị chính. Tuy nhiên, dựa trên các slide, chất lượng của hầu hết các bài thuyết trình có vẻ cao, so sánh được với kinh nghiệm của chúng tôi tại
GopherCon ở Denver năm ngoái. Mỗi bài nói có một khung thời gian một giờ cho phép
nhiều chi tiết kỹ thuật, và nhiều (hàng chục) câu hỏi từ khán giả nhiệt tình.

Như dự kiến, nhiều bài thuyết trình là về dịch vụ web, backend cho
ứng dụng di động, và v.v. Một số hệ thống có vẻ rất lớn theo bất kỳ tiêu chí nào.
Ví dụ, một bài nói của [Yang Zhou](http://gopherchina.org/user/zhouyang)
mô tả hệ thống nhắn tin nội bộ quy mô lớn, được sử dụng bởi
[Qihoo 360](http://www.360.cn/), một công ty phần mềm lớn của Trung Quốc, tất cả được viết
bằng Go. Bài thuyết trình thảo luận về cách nhóm của anh ấy xoay sở để giảm kích thước heap ban đầu 69GB (!) và các khoảng dừng GC lâu dài từ 3-6 giây xuống các con số dễ quản lý hơn, và cách họ chạy hàng triệu goroutine trên mỗi máy, trên một đội hàng nghìn máy. Một bài đăng blog khách trong tương lai được lên kế hoạch mô tả
hệ thống này chi tiết hơn.

{{image "gopherchina/image03.jpg"}}

_Phòng hội nghị đông người vào thứ Bảy._

Trong một bài thuyết trình khác, [Feng Guo](http://gopherchina.org/user/guofeng) từ
[DaoCloud](https://www.daocloud.io/) đã nói về cách họ sử dụng Go trong
công ty của họ cho cái mà họ gọi là "continuous delivery" (phân phối liên tục) của ứng dụng. DaoCloud
đảm nhận việc tự động chuyển phần mềm được lưu trữ trên GitHub (và các
tương đương của Trung Quốc) lên đám mây. Một nhà phát triển phần mềm chỉ cần push phiên bản mới lên
GitHub và DaoCloud lo phần còn lại: chạy kiểm thử,
[Dockerize](https://www.docker.com/) nó, và ship nó bằng
nhà cung cấp dịch vụ đám mây ưa thích của bạn.

Một số diễn giả đến từ các công ty phần mềm lớn được công nhận rộng rãi (tôi đã cho chương trình hội nghị cho những người không làm kỹ thuật và họ dễ dàng nhận ra tên của một số công ty). Nhiều hơn ở Mỹ, có vẻ Go không chỉ
cực kỳ phổ biến với những người mới và startup, mà đã tìm đường vào các tổ chức lớn hơn và được áp dụng ở quy mô mà chúng tôi chỉ bắt đầu
thấy ở nơi khác.

Không phải là chuyên gia về dịch vụ web, trong bài thuyết trình của mình tôi đã lệch khỏi
chủ đề chung của hội nghị một chút bằng cách nói về
[gofmt](/cmd/gofmt/) và cách việc sử dụng rộng rãi của nó đã bắt đầu
định hình kỳ vọng không chỉ cho Go mà còn cho các ngôn ngữ khác.
Tôi thuyết trình bằng tiếng Anh nhưng đã dịch các slide của mình sang tiếng Quan Thoại trước đó. Do
rào cản ngôn ngữ đáng kể, tôi không kỳ vọng nhiều câu hỏi về bài nói của mình.
Thay vào đó tôi quyết định giữ ngắn gọn và để lại nhiều thời gian cho các câu hỏi chung
về Go, điều mà khán giả đánh giá cao.

{{image "gopherchina/image06.jpg"}}

_Không có sự kiện xã hội nào ở Trung Quốc là hoàn chỉnh nếu thiếu đồ ăn tuyệt vời._

Vài ngày sau hội nghị, tôi đã thăm công ty khởi nghiệp 4 tuổi
[Qiniu](http://www.qiniu.com/) ("Bảy Con Bò"), theo lời mời của
[CEO](http://gopherchina.org/user/xushiwei) Wei Hsu, được tổ chức và
dịch với sự giúp đỡ của Asta Xie. Qiniu là nhà cung cấp lưu trữ đám mây
cho ứng dụng di động; Wei Hsu đã thuyết trình tại hội nghị và cũng là
tác giả của một trong những cuốn sách tiếng Trung đầu tiên về Go (cái ngoài cùng bên trái trong
ảnh trên).

{{image "gopherchina/image02.jpg"}}
{{image "gopherchina/image00.jpg"}}

_Sảnh Qiniu, kỹ thuật._

Qiniu là một cửa hàng toàn Go cực kỳ thành công, với khoảng 160 nhân viên, phục vụ
hơn 150.000 công ty và nhà phát triển, lưu trữ hơn 50 tỷ file, và
tăng trưởng hơn 500 triệu file mỗi ngày. Khi được hỏi về lý do
thành công của Go ở Trung Quốc, Wei Hsu nhanh chóng trả lời: PHP cực kỳ phổ biến ở
Trung Quốc, nhưng tương đối chậm và không phù hợp cho các hệ thống lớn. Như ở
Mỹ, các trường đại học dạy C++ và Java là ngôn ngữ chính, nhưng đối với nhiều
ứng dụng C++ là công cụ quá phức tạp và Java quá cồng kềnh. Theo ý kiến của anh,
Go giờ đóng vai trò mà PHP truyền thống giữ, nhưng Go chạy nhanh hơn nhiều,
an toàn kiểu, và mở rộng dễ hơn. Anh thích thực tế rằng Go đơn giản và
ứng dụng dễ triển khai. Anh cho rằng ngôn ngữ là "hoàn hảo" với
họ và yêu cầu chính của anh là một gói được khuyến nghị hoặc thậm chí chuẩn hóa để dễ dàng truy cập các hệ thống cơ sở dữ liệu. Anh có đề cập rằng họ đã gặp vấn đề GC trong
quá khứ nhưng có thể khắc phục được. Hy vọng bản phát hành 1.5 sắp tới của chúng tôi
sẽ giải quyết vấn đề này. Với Qiniu, Go xuất hiện đúng lúc và đúng
nơi (mã nguồn mở).

Theo Asta Xie, Qiniu chỉ là một trong nhiều cửa hàng Go ở Trung Quốc. Các công ty lớn
như Alibaba, Baidu, Tencent, và Weibo, giờ đều đang sử dụng Go theo
hình thức này hay hình thức khác. Anh chỉ ra rằng trong khi Thượng Hải và các thành phố lân cận
như [Tô Châu](https://www.google.com/maps/place/Suzhou,+Jiangsu,+China) là
các trung tâm công nghệ cao, thậm chí nhiều nhà phát triển phần mềm hơn được tìm thấy ở khu vực Bắc Kinh.
Cho năm 2016, Asta hy vọng tổ chức một hội nghị tiếp nối lớn hơn (1000, có thể 1500 người) ở Bắc Kinh.

Có vẻ chúng tôi đã tìm thấy người dùng Go ở Trung Quốc: Họ ở khắp nơi!

_Một số tài liệu GopherChina, bao gồm các video, giờ có sẵn cùng với các khóa học Go trên một_ [_trang bên thứ ba_](http://www.imooc.com/view/407).
