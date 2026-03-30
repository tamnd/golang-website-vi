---
title: "Trải nghiệm tìm kiếm mới trên pkg.go.dev"
date: 2021-11-09
by:
- Julie Qiu
summary: Tìm kiếm package trên pkg.go.dev đã được cập nhật, và bạn hiện có thể tìm kiếm các symbol!
template: true
---

Chúng tôi vui mừng ra mắt trải nghiệm tìm kiếm mới trên [pkg.go.dev](https://pkg.go.dev/).

Những thay đổi này được thúc đẩy bởi [phản hồi chúng tôi đã nhận được](/issue/47321) về trang tìm kiếm, và chúng tôi hy vọng bạn sẽ thích chúng. Bài đăng này cung cấp tổng quan về những gì bạn có thể mong đợi thấy trên trang web.

## Nhóm các kết quả tìm kiếm package liên quan

Các kết quả tìm kiếm cho các package trong cùng một module hiện được nhóm lại với nhau. Package phù hợp nhất với yêu cầu tìm kiếm được tô sáng. Thay đổi này được thực hiện để giảm nhiễu khi nhiều package trong cùng một module có thể phù hợp với một tìm kiếm. Ví dụ, tìm kiếm "markdown" hiển thị một hàng liệt kê "Other packages in module" cho một số kết quả.

{{image "pkgsite-search-redesign/markdown.png" 850}}

Các kết quả cho các phiên bản major khác nhau của cùng một module giờ đây cũng được nhóm lại với nhau. Phiên bản major cao nhất chứa một bản phát hành được gắn thẻ sẽ được tô sáng. Ví dụ, tìm kiếm "github" hiển thị module v39, với các phiên bản cũ hơn được liệt kê là "Other major versions".

{{image "pkgsite-search-redesign/github.png" 850}}

Cuối cùng, chúng tôi đã tổ chức lại thông tin liên quan đến imports, phiên bản và giấy phép. Chúng tôi cũng đã thêm các liên kết đến các tab này trực tiếp từ trang kết quả tìm kiếm.

## Giới thiệu tìm kiếm symbol

Trong năm qua, chúng tôi đã giới thiệu thêm thông tin về các symbol trên pkg.go.dev và làm việc để cải thiện cách trình bày thông tin đó. Chúng tôi đã ra mắt khả năng xem lịch sử API của bất kỳ package nào. Chúng tôi cũng đánh dấu các symbol bị deprecated trong mục lục tài liệu và ẩn chúng theo mặc định trong tài liệu package.

Với bản cập nhật tìm kiếm này, pkg.go.dev hiện cũng hỗ trợ tìm kiếm các symbol trong các package Go. Khi người dùng nhập một symbol vào thanh tìm kiếm, họ sẽ được đưa đến một tab tìm kiếm mới dành cho kết quả tìm kiếm symbol. Có một vài cách khác nhau mà pkg.go.dev xác định người dùng đang tìm kiếm một symbol. Chúng tôi đã thêm các ví dụ vào trang chủ pkg.go.dev, và hướng dẫn chi tiết vào [trang trợ giúp tìm kiếm](https://pkg.go.dev/search-help).

{{image "pkgsite-search-redesign/httpclient.png" 850}}

## Phản hồi

Chúng tôi rất vui khi chia sẻ trải nghiệm tìm kiếm mới này với bạn, và chúng tôi rất muốn nghe phản hồi của bạn!

Như thường lệ, hãy sử dụng nút "Report an Issue" ở cuối mỗi trang trên trang web để chia sẻ ý kiến của bạn.

Nếu bạn quan tâm đến việc đóng góp cho dự án này, pkg.go.dev là mã nguồn mở! Hãy xem [hướng dẫn đóng góp](https://go.googlesource.com/pkgsite/+/refs/heads/master/CONTRIBUTING.md) để tìm hiểu thêm.
