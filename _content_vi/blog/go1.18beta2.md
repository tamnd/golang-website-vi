---
title: Thông báo Go 1.18 Beta 2
date: 2022-01-31
by:
- Jeremy Faller and Steve Francia, for the Go team
summary: Go 1.18 Beta 2 là bản xem trước thứ hai của Go 1.18. Hãy thử và cho chúng tôi biết nếu bạn gặp vấn đề.
---

Chúng tôi rất vui trước sự hào hứng xung quanh bản phát hành Go 1.18 sắp tới,
bổ sung hỗ trợ cho
[generics](/blog/why-generics),
[fuzzing](/blog/fuzz-beta) và
[Go workspace mode](/design/45713-workspace) mới.

Chúng tôi đã phát hành Go 1.18 beta 1 hai tháng trước,
và đây hiện là bản beta Go được tải về nhiều nhất từ trước đến nay,
với số lượt tải gấp đôi bất kỳ bản phát hành nào trước đó.
Beta 1 cũng đã chứng tỏ rất đáng tin cậy;
trên thực tế, chúng tôi đã chạy nó trong môi trường production tại Google.

Phản hồi của bạn về Beta 1 đã giúp chúng tôi xác định các lỗi ít gặp
trong hỗ trợ mới cho generics và đảm bảo bản phát hành cuối cùng ổn định hơn.
Chúng tôi đã giải quyết các vấn đề này trong bản phát hành Go 1.18 Beta 2 hôm nay,
và chúng tôi khuyến khích mọi người thử.
Cách dễ nhất để cài đặt song song với Go toolchain hiện có là chạy:

	go install golang.org/dl/go1.18beta2@latest
	go1.18beta2 download

Sau đó, bạn có thể dùng `go1.18beta2` thay thế trực tiếp cho `go`.
Để xem thêm tùy chọn tải về, truy cập [go.dev/dl/#go1.18beta2](/dl/#go1.18beta2).

Vì chúng tôi dành thời gian phát hành thêm một bản beta,
chúng tôi nay kỳ vọng release candidate của Go 1.18 sẽ được phát hành vào tháng 2,
với bản phát hành Go 1.18 chính thức vào tháng 3.

Language server Go `gopls` và extension VS Code Go
nay hỗ trợ generics.
Để cài đặt `gopls` có hỗ trợ generics, xem
[tài liệu này](https://github.com/golang/tools/blob/master/gopls/doc/advanced.md#working-with-generic-code),
và để cấu hình extension VS Code Go, làm theo [hướng dẫn này](https://github.com/golang/vscode-go/blob/master/docs/advanced.md#using-go118).

Như thường lệ, đặc biệt với bản beta,
nếu bạn phát hiện bất kỳ vấn đề nào, hãy [tạo issue](/issue/new).
