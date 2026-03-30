---
title: Ra mắt Module Mirror và Cơ sở dữ liệu Checksum
date: 2019-08-29
by:
- Katie Hockman
tags:
- tools
- versioning
summary: Module mirror và cơ sở dữ liệu checksum của Go cung cấp khả năng tải xuống nhanh hơn, được xác minh cho các dependency Go của bạn.
template: true
---


Chúng tôi vui mừng chia sẻ rằng [mirror](https://proxy.golang.org),
[index](https://index.golang.org), và
[cơ sở dữ liệu checksum](https://sum.golang.org) module của chúng tôi hiện đã sẵn sàng cho môi trường production! Lệnh `go`
sẽ mặc định sử dụng module mirror và cơ sở dữ liệu checksum cho
[người dùng module Go 1.13](/doc/go1.13#introduction). Xem
[proxy.golang.org/privacy](https://proxy.golang.org/privacy) để biết thông tin về quyền riêng tư
của các dịch vụ này và
[tài liệu lệnh go](/cmd/go/#hdr-Module_downloading_and_verification)
để biết chi tiết cấu hình, bao gồm cách tắt việc sử dụng các máy chủ này hoặc
sử dụng các máy chủ khác. Nếu bạn phụ thuộc vào các module không công khai, hãy xem
[tài liệu về cấu hình môi trường của bạn](/cmd/go/#hdr-Module_configuration_for_non_public_modules).

Bài viết này sẽ mô tả các dịch vụ này và lợi ích của việc sử dụng chúng, và
tóm tắt một số điểm từ bài nói
[Go Module Proxy: Life of a Query](https://youtu.be/KqTySYYhPUE) tại Gophercon 2019.
Xem [bản ghi](https://youtu.be/KqTySYYhPUE) nếu bạn quan tâm đến bài nói đầy đủ.

## Module Mirror

[Modules](/blog/versioning-proposal) là các tập hợp package Go
được đánh phiên bản cùng nhau, và nội dung của mỗi phiên bản là bất biến.
Tính bất biến đó cung cấp các cơ hội mới cho việc cache và xác thực.
Khi `go get` chạy trong chế độ module, nó phải tìm nạp module chứa
các package được yêu cầu, cũng như bất kỳ dependency mới nào được giới thiệu bởi module đó,
cập nhật các tệp
[go.mod](/cmd/go/#hdr-The_go_mod_file) và
[go.sum](/cmd/go/#hdr-Module_downloading_and_verification)
khi cần. Việc tìm nạp module từ hệ thống quản lý phiên bản có thể tốn kém về
độ trễ và lưu trữ trong hệ thống: lệnh `go` có thể bị buộc phải tải xuống
toàn bộ lịch sử commit của kho lưu trữ chứa dependency bắc cầu, ngay cả
cái không được build, chỉ để giải quyết phiên bản của nó.

Giải pháp là sử dụng module proxy, nói API phù hợp hơn với
nhu cầu của lệnh `go` (xem `go help goproxy`). Khi `go get` chạy trong
chế độ module với một proxy, nó sẽ hoạt động nhanh hơn bằng cách chỉ yêu cầu siêu dữ liệu module
hoặc mã nguồn cụ thể mà nó cần, và không lo lắng về phần còn lại. Dưới đây là
ví dụ về cách lệnh `go` có thể sử dụng proxy với `go get` bằng cách yêu cầu danh sách
phiên bản, sau đó tệp info, mod, và zip cho phiên bản tagged mới nhất.

{{image "module-mirror-launch/proxy-protocol.png" 800}}

Module mirror là một loại module proxy đặc biệt lưu trữ siêu dữ liệu và
mã nguồn trong hệ thống lưu trữ của riêng nó, cho phép mirror tiếp tục phục vụ
mã nguồn không còn có sẵn từ các vị trí gốc. Điều này có thể
tăng tốc tải xuống và bảo vệ bạn khỏi các dependency biến mất. Xem
[Go Modules trong năm 2019](/blog/modules2019) để biết thêm thông tin.

Nhóm Go duy trì một module mirror, được phục vụ tại
[proxy.golang.org](https://proxy.golang.org), mà lệnh `go` sẽ sử dụng theo
mặc định cho người dùng module kể từ Go 1.13. Nếu bạn đang chạy phiên bản trước đó của lệnh `go`,
bạn có thể sử dụng dịch vụ này bằng cách đặt
`GOPROXY=https://proxy.golang.org` trong môi trường cục bộ của bạn.

## Cơ sở dữ liệu Checksum

Modules đã giới thiệu tệp `go.sum`, là danh sách các hash SHA-256 của
mã nguồn và các tệp `go.mod` của mỗi dependency khi được tải xuống lần đầu.
Lệnh `go` có thể sử dụng các hash để phát hiện hành vi sai của máy chủ gốc hoặc
proxy cho bạn code khác nhau cho cùng phiên bản.

Hạn chế của tệp `go.sum` này là nó hoàn toàn hoạt động dựa trên tin tưởng vào _lần đầu tiên bạn dùng_. Khi bạn thêm một phiên bản của dependency mà bạn chưa từng thấy trước đó
vào module của bạn (có thể bằng cách nâng cấp một dependency hiện có), lệnh `go`
tìm nạp code và thêm các dòng vào tệp `go.sum` ngay lập tức. Vấn đề là
những dòng `go.sum` đó không được kiểm tra với bất kỳ ai khác: chúng có thể
khác với các dòng `go.sum` mà lệnh `go` vừa tạo ra cho
người khác, có thể vì proxy đã cố ý phục vụ code độc hại
nhắm vào bạn.

Giải pháp của Go là một nguồn toàn cầu của các dòng `go.sum`, được gọi là
[cơ sở dữ liệu checksum](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#checksum-database),
đảm bảo rằng lệnh `go` luôn thêm cùng các dòng vào tệp `go.sum` của mọi người.
Bất cứ khi nào lệnh `go` nhận mã nguồn mới, nó có thể xác minh hash
của code đó với cơ sở dữ liệu toàn cầu này để đảm bảo các hash khớp,
đảm bảo rằng mọi người đang sử dụng cùng code cho một phiên bản nhất định.

Cơ sở dữ liệu checksum được phục vụ bởi [sum.golang.org](https://sum.golang.org), và
được xây dựng trên [Transparent Log](https://research.swtch.com/tlog) (hay "Merkle
tree") của các hash được hỗ trợ bởi [Trillian](https://github.com/google/trillian). Ưu điểm
chính của Merkle tree là nó chống giả mạo và có các đặc tính
không cho phép hành vi sai phát hiện được, làm cho nó đáng tin cậy hơn
một cơ sở dữ liệu đơn giản. Lệnh `go` sử dụng cây này để kiểm tra
bằng chứng "inclusion" (rằng một bản ghi cụ thể tồn tại trong log) và bằng chứng "consistency"
(rằng cây chưa bị giả mạo) trước khi thêm các dòng `go.sum` mới
vào tệp `go.sum` của module. Dưới đây là một ví dụ về cây như vậy.

{{image "module-mirror-launch/tree.png" 800}}

Cơ sở dữ liệu checksum hỗ trợ
[một tập hợp các endpoint](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#checksum-database)
được sử dụng bởi lệnh `go` để yêu cầu và xác minh các dòng `go.sum`. Endpoint `/lookup`
cung cấp "signed tree head" (STH) và các dòng `go.sum` được yêu cầu. Endpoint
`/tile` cung cấp các chunk của cây được gọi là _tiles_ mà lệnh `go`
có thể sử dụng cho các bằng chứng. Dưới đây là ví dụ về cách lệnh `go` có thể
tương tác với cơ sở dữ liệu checksum bằng cách thực hiện `/lookup` một phiên bản module, rồi
yêu cầu các tile cần thiết cho các bằng chứng.

{{image "module-mirror-launch/sumdb-protocol.png" 800}}

Cơ sở dữ liệu checksum này cho phép lệnh `go` sử dụng an toàn một proxy
không tin cậy. Vì có một lớp bảo mật có thể kiểm tra nằm ở trên,
một proxy hoặc máy chủ gốc không thể cố ý, tùy tiện, hoặc vô tình
bắt đầu cung cấp cho bạn code sai mà không bị phát hiện. Ngay cả tác giả của một
module cũng không thể di chuyển thẻ của họ hoặc thay đổi các bit liên quan đến
một phiên bản cụ thể từ ngày này sang ngày khác mà không bị phát hiện.

Nếu bạn đang dùng Go 1.12 hoặc cũ hơn, bạn có thể kiểm tra thủ công tệp `go.sum`
với cơ sở dữ liệu checksum bằng
[gosumcheck](https://godoc.org/golang.org/x/mod/gosumcheck):

	$ go get golang.org/x/mod/gosumcheck
	$ gosumcheck /path/to/go.sum

Ngoài việc xác minh được thực hiện bởi lệnh `go`, các
kiểm toán viên bên thứ ba có thể giúp cơ sở dữ liệu checksum chịu trách nhiệm bằng cách lặp qua log
để tìm các mục xấu. Họ có thể làm việc cùng nhau và lan truyền thông tin về trạng thái
của cây khi nó phát triển để đảm bảo rằng nó vẫn không bị xâm phạm, và chúng tôi hy vọng
cộng đồng Go sẽ vận hành chúng.

## Module Index

Module index được phục vụ bởi [index.golang.org](https://index.golang.org), và
là nguồn cấp dữ liệu công khai của các phiên bản module mới có sẵn thông qua
[proxy.golang.org](https://proxy.golang.org). Điều này đặc biệt hữu ích cho
các nhà phát triển công cụ muốn duy trì cache của riêng họ về những gì có sẵn trong
[proxy.golang.org](https://proxy.golang.org), hoặc cập nhật về một số
module mới nhất mà mọi người đang sử dụng.

## Phản hồi hoặc lỗi

Chúng tôi hy vọng các dịch vụ này cải thiện trải nghiệm của bạn với modules, và khuyến khích bạn
[ghi lỗi](/issue/new?title=proxy.golang.org) nếu bạn gặp phải
vấn đề hoặc có phản hồi!
