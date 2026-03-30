---
title: Go Modules trong năm 2019
date: 2018-12-19
by:
- Russ Cox
tags:
- tools
- versioning
summary: Những gì nhóm Go đang lên kế hoạch cho Go modules trong năm 2019.
template: true
---

## Thật là một năm tuyệt vời!

Năm 2018 là một năm tuyệt vời cho hệ sinh thái Go, với
quản lý package là một trong những trọng tâm chính của chúng tôi.
Vào tháng Hai, chúng tôi đã bắt đầu một cuộc thảo luận toàn cộng đồng về cách tích hợp
quản lý package trực tiếp vào bộ công cụ Go,
và vào tháng Tám chúng tôi đã cung cấp cài đặt thô đầu tiên của tính năng đó,
được gọi là Go modules, trong Go 1.11.
Việc chuyển đổi sang Go modules sẽ là thay đổi bao quát nhất
cho hệ sinh thái Go kể từ Go 1.
Chuyển đổi toàn bộ hệ sinh thái, code, người dùng, công cụ, v.v., từ
GOPATH sang modules sẽ đòi hỏi công việc trong nhiều lĩnh vực khác nhau.
Hệ thống module sẽ giúp chúng tôi cung cấp
xác thực tốt hơn và tốc độ build nhanh hơn cho hệ sinh thái Go.

Bài viết này là bản xem trước về những gì nhóm Go đang lên kế hoạch
liên quan đến modules trong năm 2019.

## Các bản phát hành

Go 1.11, được phát hành vào tháng Tám năm 2018, đã giới thiệu [hỗ trợ sơ bộ cho modules](/doc/go1.11#modules).
Hiện tại, hỗ trợ module được duy trì song song với
các cơ chế dựa trên GOPATH truyền thống.
Lệnh `go` mặc định chế độ module khi chạy
trong các cây thư mục bên ngoài GOPATH/src và
được đánh dấu bằng các tệp `go.mod` trong gốc của chúng.
Cài đặt này có thể được ghi đè bằng cách đặt biến môi trường chuyển tiếp
`$GO111MODULE` thành `on` hoặc `off`;
hành vi mặc định là chế độ `auto`.
Chúng tôi đã thấy sự áp dụng đáng kể của modules trong toàn bộ cộng đồng Go,
cùng với nhiều đề xuất hữu ích và báo cáo lỗi
giúp chúng tôi cải thiện modules.

Go 1.12, dự kiến vào tháng Hai năm 2019, sẽ tinh chỉnh hỗ trợ module
nhưng vẫn để nó ở chế độ `auto` theo mặc định.
Ngoài nhiều bản vá lỗi và các cải tiến nhỏ khác,
có lẽ thay đổi quan trọng nhất trong Go 1.12
là các lệnh như `go` `run` `x.go`
hoặc `go` `get` `rsc.io/2fa@v1.1.0`
hiện có thể hoạt động trong chế độ `GO111MODULE=on` mà không cần tệp `go.mod` rõ ràng.

Mục tiêu của chúng tôi là Go 1.13, dự kiến vào tháng Tám năm 2019, để bật chế độ module theo
mặc định (nghĩa là thay đổi mặc định từ `auto` sang `on`)
và không dùng nữa chế độ GOPATH.
Để làm điều đó, chúng tôi đã làm việc với hỗ trợ hệ thống công cụ tốt hơn
cùng với hỗ trợ tốt hơn cho hệ sinh thái module mã nguồn mở.

## Hệ thống công cụ & Tích hợp IDE

Trong tám năm chúng ta có GOPATH,
một lượng hệ thống công cụ đáng kinh ngạc đã được tạo ra
giả định rằng mã nguồn Go được lưu trữ trong GOPATH.
Việc chuyển sang modules đòi hỏi phải cập nhật tất cả code có
giả định đó.
Chúng tôi đã thiết kế một package mới,
[golang.org/x/tools/go/packages](https://godoc.org/golang.org/x/tools/go/packages),
trừu tượng hóa hoạt động tìm kiếm và tải thông tin
về mã nguồn Go cho một mục tiêu nhất định.
Package mới này tự động thích nghi với cả
chế độ GOPATH và modules và cũng có thể mở rộng
cho các bố cục code dành riêng cho công cụ, chẳng hạn như cái
được sử dụng bởi Bazel.
Chúng tôi đã làm việc với các tác giả công cụ trong toàn bộ cộng đồng Go
để giúp họ áp dụng golang.org/x/tools/go/packages trong các công cụ của họ.

Là một phần của nỗ lực này, chúng tôi cũng đã làm việc để
hợp nhất các công cụ truy vấn mã nguồn khác nhau
như gocode, godef, và go-outline
thành một công cụ duy nhất có thể được sử dụng từ
dòng lệnh và cũng hỗ trợ
[giao thức language server](https://langserver.org/)
được sử dụng bởi các IDE hiện đại.

Việc chuyển sang modules và những thay đổi trong tải package
cũng đã thúc đẩy một sự thay đổi đáng kể đối với phân tích chương trình Go.
Là một phần của việc sửa đổi `go` `vet` để hỗ trợ modules,
chúng tôi đã giới thiệu một framework tổng quát cho việc phân tích tăng dần
các chương trình Go,
trong đó một analyzer được gọi cho một package tại một thời điểm.
Trong framework này, việc phân tích một package có thể ghi ra các sự thật
có sẵn cho các phân tích của các package khác import gói đầu tiên.
Ví dụ, phân tích của `go` `vet` về [package log](/pkg/log/)
xác định và ghi lại sự thật rằng `log.Printf` là một wrapper của `fmt.Printf`.
Sau đó `go` `vet` có thể kiểm tra các chuỗi định dạng kiểu printf trong các package khác
gọi `log.Printf`.
Framework này nên cho phép nhiều công cụ phân tích chương trình mới, tinh vi
giúp các lập trình viên tìm lỗi sớm hơn
và hiểu code tốt hơn.

## Module Index

Một trong những phần quan trọng nhất của thiết kế ban đầu cho `go` `get`
là nó có tính _phi tập trung_:
chúng tôi tin khi đó, và chúng tôi vẫn tin ngày nay, rằng
bất kỳ ai cũng nên có thể xuất bản code của họ trên bất kỳ máy chủ nào,
trái ngược với các registry trung tâm
như CPAN của Perl, Maven của Java, hoặc NPM của Node.
Đặt tên miền ở đầu không gian import `go` `get`
đã tái sử dụng một hệ thống phi tập trung hiện có
và tránh phải giải quyết lại các vấn đề về
việc quyết định ai có thể sử dụng tên nào.
Nó cũng cho phép các công ty import code trên các máy chủ riêng
bên cạnh code từ các máy chủ công khai.
Điều quan trọng là phải duy trì tính phi tập trung này khi chúng ta chuyển sang Go modules.

Tính phi tập trung của các dependency Go đã có nhiều lợi ích,
nhưng nó cũng mang lại một vài nhược điểm đáng kể.
Nhược điểm đầu tiên là quá khó để tìm tất cả các package Go công khai.
Mỗi trang web muốn cung cấp thông tin về
các package phải
tự crawl, hoặc chờ cho đến khi người dùng hỏi
về một package cụ thể trước khi tìm nạp nó.

Chúng tôi đang làm việc trên một dịch vụ mới, Go Module Index,
sẽ cung cấp một log công khai của các package đang gia nhập hệ sinh thái Go.
Các trang web như godoc.org và goreportcard.com sẽ có thể theo dõi log này
để tìm các mục mới thay vì mỗi bên độc lập cài đặt code
để tìm các package mới.
Chúng tôi cũng muốn dịch vụ cho phép tra cứu các package
bằng các truy vấn đơn giản, để cho phép `goimports` thêm
các import cho các package chưa được tải xuống vào hệ thống cục bộ.

## Xác thực Module

Ngày nay, `go` `get` dựa vào xác thực ở cấp kết nối (HTTPS hoặc SSH)
để kiểm tra rằng nó đang nói chuyện với đúng máy chủ để tải xuống code.
Không có kiểm tra bổ sung nào về chính code,
để ngỏ khả năng tấn công man-in-the-middle
nếu cơ chế HTTPS hoặc SSH bị xâm phạm theo cách nào đó.
Tính phi tập trung có nghĩa là code cho một build được tìm nạp
từ nhiều máy chủ khác nhau, điều đó có nghĩa là build phụ thuộc vào
nhiều hệ thống để phục vụ code đúng.

Thiết kế Go modules cải thiện xác thực code bằng cách lưu trữ
tệp `go.sum` trong mỗi module;
tệp đó liệt kê hash mật mã
của cây tệp dự kiến cho mỗi dependency của module.
Khi sử dụng modules, lệnh `go` sử dụng `go.sum` để xác minh
rằng các dependency giống bit từng bit với các phiên bản dự kiến
trước khi sử dụng chúng trong một build.
Nhưng tệp `go.sum` chỉ liệt kê các hash cho các dependency cụ thể
được sử dụng bởi module đó.
Nếu bạn đang thêm một dependency mới
hoặc cập nhật các dependency với `go` `get` `-u`,
không có mục tương ứng trong `go.sum` và do đó
không có xác thực trực tiếp của các bit được tải xuống.

Đối với các module công khai, chúng tôi có ý định chạy một dịch vụ chúng tôi gọi là _notary_
theo dõi log module index,
tải xuống các module mới,
và ký mật mã các tuyên bố dạng
"module M tại phiên bản V có hash cây tệp H."
Dịch vụ notary sẽ xuất bản tất cả các hash đã được notarized
trong [log chống giả mạo](http://static.usenix.org/event/sec09/tech/full_papers/crosby.pdf) có thể truy vấn, kiểu
[Certificate Transparency](https://www.certificate-transparency.org/),
để bất kỳ ai cũng có thể xác minh rằng notary đang hoạt động đúng.
Log này sẽ đóng vai trò là tệp `go.sum` công khai, toàn cầu
mà `go` `get` có thể sử dụng để xác thực các module
khi thêm hoặc cập nhật các dependency.

Chúng tôi đang nhắm đến việc lệnh `go` kiểm tra các hash đã được notarized
cho các module công khai chưa có trong `go.sum`
bắt đầu từ Go 1.13.

## Module Mirrors

Vì `go` `get` phi tập trung tìm nạp code từ nhiều máy chủ gốc,
việc tìm nạp code chỉ nhanh và đáng tin cậy bằng máy chủ chậm nhất,
kém đáng tin cậy nhất.
Phương án bảo vệ duy nhất có sẵn trước modules là vendoring
các dependency vào các kho lưu trữ của bạn.
Mặc dù vendoring sẽ tiếp tục được hỗ trợ,
chúng tôi ưu tiên một giải pháp hoạt động cho tất cả modules, không chỉ những cái bạn đang sử dụng, và
không đòi hỏi phải sao chép dependency vào mỗi
kho lưu trữ sử dụng nó.

Thiết kế Go module giới thiệu ý tưởng về một module proxy,
là một máy chủ mà lệnh `go` yêu cầu modules,
thay vì các máy chủ gốc.
Một loại proxy quan trọng là _module mirror_,
trả lời các yêu cầu về modules bằng cách tìm nạp chúng
từ các máy chủ gốc rồi cache chúng để sử dụng trong
các yêu cầu tương lai.
Một mirror hoạt động tốt nên nhanh và đáng tin cậy
ngay cả khi một số máy chủ gốc đã ngừng hoạt động.
Chúng tôi đang lên kế hoạch ra mắt một dịch vụ mirror cho các module công khai vào năm 2019.
Các dự án khác, như GoCenter và Athens, cũng đang lên kế hoạch các dịch vụ mirror.
(Chúng tôi dự đoán rằng các công ty sẽ có nhiều lựa chọn để chạy
các mirror nội bộ của riêng họ, nhưng bài viết này đang tập trung vào các mirror công khai.)

Một vấn đề tiềm ẩn với mirrors là chúng
chính xác là các máy chủ man-in-the-middle,
làm cho chúng là mục tiêu tự nhiên cho các cuộc tấn công.
Các lập trình viên Go cần một số đảm bảo rằng các mirror
đang cung cấp cùng các bit mà các máy chủ gốc sẽ cung cấp.
Quá trình notary mà chúng tôi đã mô tả trong phần trước
giải quyết đúng mối lo ngại này, và nó
sẽ áp dụng cho các tải xuống sử dụng mirror
cũng như tải xuống sử dụng máy chủ gốc.
Bản thân các mirror không cần được tin tưởng.

Chúng tôi đang nhắm đến việc module mirror do Google vận hành
sẵn sàng được sử dụng theo mặc định trong lệnh `go` bắt đầu từ Go 1.13.
Sử dụng mirror thay thế, hoặc không sử dụng mirror nào, sẽ đơn giản
để cấu hình.

## Khám phá Module

Cuối cùng, chúng tôi đã đề cập trước đó rằng module index sẽ giúp dễ dàng hơn
xây dựng các trang như godoc.org.
Một phần công việc của chúng tôi trong năm 2019 sẽ là một sự sửa đổi lớn của godoc.org
để làm cho nó hữu ích hơn cho các lập trình viên cần
khám phá các module có sẵn
rồi quyết định có nên phụ thuộc vào một module nhất định hay không.

## Bức tranh toàn cảnh

Sơ đồ này cho thấy cách mã nguồn module
di chuyển qua thiết kế trong bài viết này.

{{image "modules2019/code.png" 374}}

Trước đây, tất cả người tiêu dùng mã nguồn Go, lệnh `go`
và bất kỳ trang web nào như godoc.org, đều tìm nạp code trực tiếp từ mỗi host code.
Bây giờ họ có thể tìm nạp code được cache từ một mirror nhanh, đáng tin cậy,
trong khi vẫn xác thực rằng các bit được tải xuống là chính xác.
Và dịch vụ index giúp các mirror, godoc.org,
và bất kỳ trang web tương tự nào khác dễ dàng theo kịp tất cả code tuyệt vời mới
được thêm vào hệ sinh thái Go mỗi ngày.

Chúng tôi hào hứng về tương lai của Go modules trong năm 2019,
và chúng tôi hy vọng bạn cũng vậy. Chúc mừng năm mới!
