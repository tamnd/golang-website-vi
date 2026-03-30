---
title: Xây dựng StatHat với Go
date: 2011-12-19
by:
- Patrick Crosby
tags:
- guest
summary: Cách StatHat sử dụng Go và lý do họ chọn nó.
template: true
---

## Giới thiệu

Tôi tên là Patrick Crosby và là người sáng lập một công ty tên là Numerotron.
Gần đây chúng tôi đã ra mắt [StatHat](http://www.stathat.com).
Bài đăng này nói về lý do chúng tôi chọn phát triển StatHat bằng [Go](/),
bao gồm các chi tiết về cách chúng tôi đang sử dụng Go.

[StatHat](http://www.stathat.com) là một công cụ để theo dõi thống kê và sự kiện trong mã của bạn.
Tất cả mọi người từ nhà thiết kế HTML đến kỹ sư backend đều có thể sử dụng StatHat một cách dễ dàng,
vì nó hỗ trợ gửi thống kê từ HTML, JavaScript,
Go, và mười hai ngôn ngữ khác.

Bạn gửi các con số của mình đến StatHat; nó tạo ra các biểu đồ đẹp,
có thể nhúng hoàn toàn về dữ liệu của bạn.
StatHat sẽ cảnh báo bạn khi các điều kiện đã chỉ định xảy ra,
gửi cho bạn báo cáo email hàng ngày, và nhiều hơn nữa.
Vì vậy, thay vì mất thời gian viết các công cụ theo dõi hoặc báo cáo cho ứng dụng của bạn,
bạn có thể tập trung vào mã.
Trong khi bạn thực hiện công việc thực sự, StatHat luôn cực kỳ cảnh giác,
như một con đại bàng trên tổ ở đỉnh núi, hoặc một người trông trẻ đầy nhiệt huyết.

Đây là ví dụ về biểu đồ StatHat về nhiệt độ ở NYC, Chicago và San Francisco:

{{image "stathat/weather.png"}}

## Tổng quan Kiến trúc

StatHat bao gồm hai dịch vụ chính: các lệnh gọi API nhận thống kê/sự kiện
và ứng dụng web để xem và phân tích thống kê.
Chúng tôi muốn giữ chúng tách biệt nhau nhất có thể để cô lập việc thu thập dữ liệu
với việc tương tác dữ liệu.
Chúng tôi làm điều này vì nhiều lý do, nhưng một lý do chính là chúng tôi dự đoán
sẽ xử lý một lượng lớn các yêu cầu HTTP API tự động đến và do đó sẽ có
các chiến lược tối ưu hóa khác nhau cho dịch vụ API so với một ứng dụng web
tương tác với con người.

{{image "stathat/architecture.png"}}

Dịch vụ ứng dụng web có nhiều tầng.
Máy chủ web xử lý tất cả các yêu cầu và gửi chúng đến tầng interactor.
Đối với các tác vụ đơn giản, interactor sẽ xử lý việc tạo ra bất kỳ dữ liệu cần thiết nào.
Đối với các tác vụ phức tạp, interactor dựa vào nhiều máy chủ ứng dụng
để xử lý các tác vụ như tạo biểu đồ hoặc phân tích tập dữ liệu.
Sau khi interactor hoàn thành, máy chủ web gửi kết quả đến một presenter.
Presenter phản hồi yêu cầu HTTP bằng HTML hoặc JSON.
Chúng tôi có thể mở rộng theo chiều ngang các máy chủ web, API, ứng dụng,
và cơ sở dữ liệu khi nhu cầu dịch vụ tăng và thay đổi theo thời gian.
Không có điểm lỗi đơn lẻ vì mỗi máy chủ ứng dụng có nhiều bản sao đang chạy.
Tầng interactor cho phép chúng tôi có các giao diện khác nhau với hệ thống:
http, dòng lệnh, kiểm tra tự động, API di động.
StatHat sử dụng MySQL để lưu trữ dữ liệu.

## Chọn Go

Khi thiết kế StatHat, chúng tôi có danh sách yêu cầu sau đây cho các công cụ phát triển:

  - cùng một ngôn ngữ lập trình cho hệ thống backend và frontend

  - hệ thống template HTML tốt, nhanh

  - khởi động nhanh, biên dịch lại, kiểm tra để thường xuyên thử nghiệm

  - nhiều kết nối trên một máy

  - các công cụ ngôn ngữ để xử lý đồng thời ở cấp độ ứng dụng

  - hiệu suất tốt

  - tầng RPC mạnh mẽ để liên lạc giữa các tầng

  - nhiều thư viện

  - mã nguồn mở

Chúng tôi đã đánh giá nhiều công nghệ web phổ biến và ít phổ biến hơn và
cuối cùng chọn phát triển bằng Go.

Khi Go được phát hành vào tháng 11 năm 2009, tôi đã cài đặt ngay lập tức và thích
thời gian biên dịch nhanh,
goroutines, channels, bộ gom rác,
và tất cả các gói có sẵn.
Tôi đặc biệt hài lòng với việc các ứng dụng của tôi sử dụng rất ít dòng mã.
Tôi sớm thử nghiệm tạo một ứng dụng web có tên [Langalot](http://langalot.com/)
đồng thời tìm kiếm qua năm từ điển ngôn ngữ nước ngoài khi
bạn gõ một truy vấn.
Nó cực kỳ nhanh. Tôi đưa nó lên mạng và nó đã chạy từ tháng 2 năm 2010.

Các phần sau đây trình bày chi tiết cách Go đáp ứng các yêu cầu của StatHat và
kinh nghiệm của chúng tôi khi sử dụng Go để giải quyết các vấn đề của mình.

## Runtime

Chúng tôi sử dụng [gói http](/pkg/http/) chuẩn của Go cho
máy chủ API và ứng dụng web của chúng tôi.
Tất cả các yêu cầu đầu tiên đi qua Nginx và mọi yêu cầu không phải tệp
được proxy đến các máy chủ http được cung cấp bởi Go.
Các máy chủ backend đều được viết bằng Go và sử dụng [gói rpc](/pkg/rpc/)
để giao tiếp với frontend.

## Templating

Chúng tôi đã xây dựng một hệ thống template sử dụng [gói template](/pkg/template/) chuẩn.
Hệ thống của chúng tôi bổ sung layouts, một số hàm định dạng phổ biến,
và khả năng biên dịch lại template ngay lập tức trong quá trình phát triển.
Chúng tôi rất hài lòng với hiệu suất và tính năng của các template Go.

## Thử nghiệm liên tục

Trong một công việc trước đây, tôi đã làm việc trên một trò chơi điện tử có tên
Throne of Darkness được viết bằng C++.
Chúng tôi có một vài file header mà khi được sửa đổi,
đòi hỏi phải xây dựng lại toàn bộ hệ thống, mất 20-30 phút.
Nếu ai đó thay đổi `Character.h`, họ sẽ phải chịu đựng sự tức giận của
mọi lập trình viên khác.
Ngoài sự khổ sở này, nó còn làm chậm đáng kể thời gian phát triển.

Kể từ đó, tôi luôn cố gắng chọn các công nghệ cho phép thử nghiệm nhanh, thường xuyên.
Với Go, thời gian biên dịch không còn là vấn đề.
Chúng tôi có thể biên dịch lại toàn bộ hệ thống trong vài giây, không phải vài phút.
Máy chủ web phát triển khởi động ngay lập tức,
các bài kiểm tra hoàn thành trong vài giây.
Như đã đề cập trước đây, các template được biên dịch lại khi chúng thay đổi.
Kết quả là hệ thống StatHat rất dễ làm việc,
và trình biên dịch không phải là nút cổ chai.

## RPC

Vì StatHat là hệ thống nhiều tầng, chúng tôi muốn có một tầng RPC để
tất cả giao tiếp đều được chuẩn hóa.
Với Go, chúng tôi đang sử dụng [gói rpc](/pkg/rpc/) và
[gói gob](/pkg/gob/) để mã hóa các đối tượng Go.
Trong Go, máy chủ RPC chỉ cần lấy bất kỳ đối tượng Go nào và đăng ký các phương thức
đã xuất của nó.
Không cần ngôn ngữ mô tả giao diện trung gian.
Chúng tôi thấy nó rất dễ sử dụng và nhiều máy chủ ứng dụng cốt lõi của chúng tôi
có dưới 300 dòng mã.

## Thư viện

Chúng tôi không muốn mất thời gian viết lại các thư viện cho những thứ như SSL,
driver cơ sở dữ liệu, trình phân tích cú pháp JSON/XML.
Mặc dù Go là một ngôn ngữ còn non trẻ, nó có rất nhiều gói hệ thống và
số lượng gói do người dùng đóng góp ngày càng tăng.
Chỉ với một vài ngoại lệ, chúng tôi đã tìm thấy các gói Go cho mọi thứ chúng tôi cần.

## Mã nguồn mở

Trong kinh nghiệm của chúng tôi, làm việc với các công cụ mã nguồn mở là vô cùng quý giá.
Nếu có điều gì đó không ổn, việc có thể kiểm tra
mã nguồn qua từng lớp mà không có bất kỳ hộp đen nào là cực kỳ hữu ích.
Có mã cho ngôn ngữ, máy chủ web,
các gói và công cụ cho phép chúng tôi hiểu cách mọi phần của hệ thống hoạt động.
Tất cả mọi thứ trong Go đều là mã nguồn mở. Trong codebase Go,
chúng tôi thường xuyên đọc các bài kiểm tra vì chúng thường cung cấp
các ví dụ tuyệt vời về cách sử dụng các gói và tính năng ngôn ngữ.

## Hiệu suất

Mọi người dựa vào StatHat để phân tích dữ liệu của họ đến từng phút và chúng tôi
cần hệ thống phản hồi nhanh nhất có thể.
Trong các bài kiểm tra của chúng tôi, hiệu suất của Go vượt xa hầu hết các đối thủ cạnh tranh.
Chúng tôi đã thử nghiệm nó với Rails, Sinatra, OpenResty và Node.
StatHat luôn tự giám sát bằng cách theo dõi tất cả các loại số liệu hiệu suất
về các yêu cầu,
thời lượng của các tác vụ nhất định, lượng bộ nhớ đang sử dụng.
Vì điều này, chúng tôi có thể dễ dàng đánh giá các công nghệ khác nhau.
Chúng tôi cũng đã tận dụng các tính năng kiểm tra hiệu suất benchmark
của gói testing trong Go.

## Đồng thời ở Cấp độ Ứng dụng

Trước đây, tôi từng là CTO tại OkCupid.
Kinh nghiệm của tôi ở đó khi sử dụng OKWS đã dạy tôi tầm quan trọng của lập trình
bất đồng bộ, đặc biệt khi liên quan đến các ứng dụng web động.
Không có lý do gì bạn nên thực hiện những việc như thế này theo tuần tự:
tải người dùng từ cơ sở dữ liệu, sau đó tìm thống kê của họ,
sau đó tìm cảnh báo của họ.
Những việc này nên được thực hiện đồng thời, nhưng thật ngạc nhiên,
nhiều framework phổ biến không có hỗ trợ bất đồng bộ.
Go hỗ trợ điều này ở cấp độ ngôn ngữ mà không cần mã callback rối rắm.
StatHat sử dụng goroutines rộng rãi để chạy nhiều hàm đồng thời
và channels để chia sẻ dữ liệu giữa các goroutines.

## Hosting và Triển khai

StatHat chạy trên các máy chủ EC2 của Amazon. Các máy chủ của chúng tôi được chia thành nhiều loại:

  - API

  - Web

  - Máy chủ ứng dụng

  - Cơ sở dữ liệu

Có ít nhất hai máy chủ của mỗi loại,
và chúng nằm ở các vùng khác nhau để đảm bảo tính khả dụng cao.
Thêm một máy chủ mới vào hệ thống chỉ mất vài phút.

Để triển khai, chúng tôi đầu tiên xây dựng toàn bộ hệ thống thành một thư mục
có dấu thời gian.
Kịch bản đóng gói của chúng tôi xây dựng các ứng dụng Go,
nén các tệp CSS và JS, và sao chép tất cả các kịch bản và tệp cấu hình.
Thư mục này sau đó được phân phối đến tất cả các máy chủ,
do đó tất cả đều có bản phân phối giống hệt nhau.
Một kịch bản trên mỗi máy chủ truy vấn các thẻ EC2 của nó và xác định những gì nó
chịu trách nhiệm chạy và khởi động/dừng/khởi động lại bất kỳ dịch vụ nào.
Chúng tôi thường chỉ triển khai cho một tập hợp con các máy chủ.

## Thêm thông tin

Để biết thêm thông tin về StatHat, vui lòng truy cập [stathat.com](http://www.stathat.com).
Chúng tôi đang phát hành một số mã Go chúng tôi đã viết.
Truy cập [www.stathat.com/src](http://www.stathat.com/src) để xem tất cả các
dự án StatHat mã nguồn mở.

Để tìm hiểu thêm về Go, hãy truy cập [golang.org](/).
