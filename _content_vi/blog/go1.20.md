---
title: Go 1.20 đã được phát hành!
date: 2023-02-01
by:
- Robert Griesemer, on behalf of the Go team
summary: Go 1.20 mang đến PGO, build nhanh hơn, và nhiều cải tiến cho công cụ, ngôn ngữ và thư viện.
---

Hôm nay nhóm Go vô cùng hào hứng phát hành Go 1.20,
bạn có thể tải về từ [trang download](/dl/).

Go 1.20 được hưởng lợi từ giai đoạn phát triển kéo dài,
được thực hiện nhờ quá trình kiểm thử rộng rãi hơn và độ ổn định tổng thể được cải thiện
của codebase.

Chúng tôi đặc biệt hào hứng với việc ra mắt bản xem trước [tối ưu hóa dựa trên hồ sơ thực thi](/doc/pgo)
(PGO), cho phép trình biên dịch thực hiện các tối ưu hóa đặc thù theo ứng dụng và
workload dựa trên thông tin hồ sơ lúc chạy.
Cung cấp hồ sơ cho `go build` cho phép trình biên dịch tăng tốc các ứng dụng thông thường
khoảng 3-4%, và chúng tôi kỳ vọng các bản phát hành tương lai sẽ hưởng lợi nhiều hơn từ PGO.
Vì đây là bản xem trước của hỗ trợ PGO, chúng tôi khuyến khích mọi người thử,
nhưng vẫn còn một số điểm thô ráp có thể cản trở việc sử dụng trong production.

Go 1.20 cũng bao gồm một số thay đổi ngôn ngữ,
nhiều cải tiến cho công cụ và thư viện, và hiệu năng tổng thể tốt hơn.

## Thay đổi ngôn ngữ

- Ràng buộc [`comparable`](/ref/spec#Type_constraints) đã được khai báo trước nay cũng
[thỏa mãn](/ref/spec#Satisfying_a_type_constraint) bởi
các [kiểu có thể so sánh](/ref/spec#Comparison_operators) thông thường, chẳng hạn như interface,
điều này sẽ đơn giản hóa code generic.
- Các hàm `SliceData`, `String` và `StringData` đã được thêm vào
gói [`unsafe`](/ref/spec#Package_unsafe). Chúng hoàn thiện tập hàm
để thao tác slice và chuỗi độc lập với triển khai.
- Quy tắc chuyển đổi kiểu của Go đã được mở rộng để cho phép chuyển đổi trực tiếp
[từ slice sang mảng](/ref/spec#Conversions_from_slice_to_array_or_array_pointer).
- Đặc tả ngôn ngữ nay định nghĩa thứ tự chính xác trong đó các phần tử mảng
và trường struct được [so sánh](/ref/spec#Comparison_operators).
Điều này làm rõ điều gì xảy ra trong trường hợp panic trong quá trình so sánh.

## Cải tiến công cụ

- [Công cụ `cover`](/doc/build-cover) nay có thể thu thập hồ sơ coverage của toàn bộ chương trình,
không chỉ của unit test.
- [Công cụ `go`](/cmd/go) không còn dựa vào các file archive gói thư viện chuẩn đã biên dịch trước
trong thư mục `$GOROOT/pkg`, và chúng không còn được đi kèm với bản phân phối, giúp giảm kích thước tải về.
Thay vào đó, các gói trong thư viện chuẩn được build khi cần và được cache
trong build cache, giống như các gói khác.
- Triển khai của `go test -json` đã được cải thiện
để làm cho nó mạnh mẽ hơn khi có ghi thêm bất ngờ vào `stdout`.
- Các lệnh `go build`, `go install` và các lệnh liên quan đến build khác
nay chấp nhận cờ `-pgo` để bật tối ưu hóa dựa trên hồ sơ
cũng như cờ `-cover` cho phân tích coverage toàn chương trình.
- Lệnh `go` nay tắt `cgo` theo mặc định trên các hệ thống không có C toolchain.
Do đó, khi Go được cài đặt trên hệ thống không có trình biên dịch C, nó sẽ
dùng các bản build Go thuần cho các gói trong thư viện chuẩn có dùng cgo tùy chọn,
thay vì dùng các file archive gói đã phân phối sẵn (đã bị loại bỏ, như đã đề cập ở trên).
- [Công cụ `vet`](/cmd/vet) báo cáo thêm lỗi tham chiếu biến vòng lặp
có thể xảy ra trong các bài kiểm thử chạy song song.

## Bổ sung thư viện chuẩn

- Gói [`crypto/ecdh`](/pkg/crypto/ecdh) mới cung cấp hỗ trợ rõ ràng cho
trao đổi khóa Elliptic Curve Diffie-Hellman qua các đường cong NIST và Curve25519.
- Hàm mới [`errors.Join`](/pkg/errors#Join) trả về một lỗi bọc danh sách lỗi
có thể được lấy lại nếu kiểu lỗi triển khai phương thức `Unwrap() []error`.
- Kiểu [`http.ResponseController`](/pkg/net/http#ResponseController) mới
cung cấp quyền truy cập vào chức năng mở rộng theo từng request không được xử lý bởi
interface [`http.ResponseWriter`](/pkg/net/http#ResponseWriter).
- Proxy chuyển tiếp [`httputil.ReverseProxy`](/pkg/net/http/httputil#ReverseProxy)
bao gồm hàm hook `Rewrite` mới, thay thế hook `Director` trước đó.
- Hàm [`context.WithCancelCause`](/pkg/context#WithCancelCause) mới
cung cấp cách hủy một context với một lỗi cho trước.
Lỗi đó có thể được lấy lại bằng cách gọi hàm
[`context.Cause`](/pkg/context#Cause) mới.
- Các trường [`Cancel`](/pkg/os/exec#Cmd.Cancel)
và [`WaitDelay`](/pkg/os/exec#Cmd.WaitDelay) mới trong [`os/exec.Cmd`](/pkg/os/exec#Cmd) chỉ định hành vi của
`Cmd` khi `Context` liên quan bị hủy hoặc process của nó thoát.

## Cải thiện hiệu năng

- Cải tiến trình biên dịch và bộ gom rác đã giảm bộ nhớ sử dụng
và cải thiện hiệu năng CPU tổng thể lên đến 2%.
- Công việc nhắm cụ thể vào
thời gian biên dịch đã cải thiện tốc độ build lên đến 10%.
Điều này đưa tốc độ build trở lại mức của Go 1.17.

Khi [build bản phát hành Go từ nguồn](/doc/install/source),
Go 1.20 yêu cầu bản phát hành Go 1.17.13 hoặc mới hơn.
Trong tương lai, chúng tôi dự định tiến toolchain bootstrap lên xấp xỉ
một lần mỗi năm.
Ngoài ra, bắt đầu từ Go 1.21, một số hệ điều hành cũ hơn sẽ không còn được hỗ trợ:
bao gồm Windows 7, 8, Server 2008 và Server 2012,
macOS 10.13 High Sierra và 10.14 Mojave.
Mặt khác, Go 1.20 thêm hỗ trợ thử nghiệm cho FreeBSD trên RISC-V.

Để xem danh sách đầy đủ và chi tiết hơn về tất cả các thay đổi, xem [ghi chú phát hành đầy đủ](/doc/go1.20).

Cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code, báo cáo lỗi,
chia sẻ phản hồi, và kiểm thử các release candidate. Nỗ lực của bạn giúp
đảm bảo Go 1.20 ổn định nhất có thể.
Như thường lệ, nếu bạn phát hiện bất kỳ vấn đề nào, hãy [tạo issue](/issue/new).

Tận hưởng Go 1.20!
