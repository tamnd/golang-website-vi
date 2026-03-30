---
title: Pkg.go.dev có diện mạo mới!
date: 2020-11-10T12:00:00Z
by:
- Julie Qiu
summary: Thông báo về trải nghiệm người dùng mới trên pkg.go.dev.
template: true
---


Kể từ khi ra mắt [pkg.go.dev](https://pkg.go.dev), chúng tôi đã nhận được nhiều phản hồi tuyệt vời về thiết kế và khả năng sử dụng.
Đặc biệt, rõ ràng là cách tổ chức thông tin đã gây nhầm lẫn cho người dùng khi điều hướng trên trang web.

Hôm nay chúng tôi hào hứng chia sẻ pkg.go.dev được thiết kế lại,
với hy vọng sẽ rõ ràng và hữu ích hơn.
Bài đăng này trình bày những điểm nổi bật. Để biết chi tiết,
xem [Go issue 41585](/issue/41585).

## Trang đích nhất quán cho mọi đường dẫn

Thay đổi chính là trang pkg.go.dev/\<path> đã được tổ chức lại
xung quanh khái niệm về một đường dẫn.
Một đường dẫn đại diện cho một thư mục trong một phiên bản cụ thể của module.
Bây giờ, bất kể nội dung trong thư mục đó là gì,
mọi trang đường dẫn sẽ có cùng bố cục,
với mục tiêu làm cho trải nghiệm nhất quán, hữu ích và dễ đoán.

<figure class="image">
  <img src="pkgsite-redesign/path.png" width="800" alt="Landing page for cloud.google.com/go/storage" style="border: 1px solid black;">
  <figcaption>
    Hình 1. Trang đích cho
    <a href="https://pkg.go.dev/cloud.google.com/go/storage">https://pkg.go.dev/cloud.google.com/go/storage</a>.
  </figcaption>
</figure>

Trang đường dẫn sẽ hiển thị README tại đường dẫn đó nếu có.
Trước đây, tab tổng quan chỉ hiển thị README nếu có tại thư mục gốc của module.
Đây là một trong nhiều thay đổi chúng tôi đang thực hiện để đưa thông tin quan trọng nhất lên đầu trang.

## Điều hướng tài liệu

Phần tài liệu hiện hiển thị một mục lục cùng với thanh điều hướng bên cạnh.
Điều này cho phép xem toàn bộ API của package,
trong khi vẫn có ngữ cảnh khi điều hướng qua phần tài liệu.
Ngoài ra còn có một ô nhập liệu "Jump To" mới ở thanh điều hướng bên trái,
để tìm kiếm các định danh.


<figure class="image">
  <img src="pkgsite-redesign/nav.png" width="800" alt="Jump To feature navigating net/http" style="border: 1px solid black;">
  <figcaption>
    Hình 2. Tính năng Jump To trên
    <a href="https://pkg.go.dev/net/http">https://pkg.go.dev/net/http</a>.
  </figcaption>
</figure>

Xem [Go issue 41587](/issue/41587) để biết chi tiết về các thay đổi trong phần tài liệu.

## Siêu dữ liệu trên trang chính

Thanh tiêu đề trên mỗi trang hiện hiển thị thêm siêu dữ liệu,
chẳng hạn như số lượng "imports" và "imported by" của mỗi package.
Các biểu ngữ cũng hiển thị thông tin về các phiên bản minor và major mới nhất của module.
Xem [Go issue 41588](/issue/41588) để biết chi tiết.

<figure class="image">
  <img src="pkgsite-redesign/meta.png" width="800" alt="Header metadata for github.com/russross/blackfriday" style="border: 1px solid black;">
  <figcaption>
    Hình 3. Siêu dữ liệu tiêu đề cho
    <a href="https://pkg.go.dev/github.com/russross/blackfriday">https://pkg.go.dev/github.com/russross/blackfriday</a>.
  </figcaption>
</figure>

## Video hướng dẫn

Tuần trước tại [Google Open Source Live](https://opensourcelive.withgoogle.com/events/go),
chúng tôi đã trình bày hướng dẫn trải nghiệm trang web mới trong bài nói chuyện,
[Level Up: Go Package Discovery and Editor Tooling](https://www.youtube.com/watch?v=n7ayE29b7QA&feature=emb_logo).

{{video "https://www.youtube.com/embed/n7ayE29b7QA" 650 400}}

## Phản hồi

Chúng tôi rất vui khi chia sẻ bản thiết kế cập nhật này với bạn.
Như thường lệ, hãy cho chúng tôi biết suy nghĩ của bạn qua các liên kết "Share Feedback"
và "Report an Issue" ở cuối mỗi trang của trang web.

Và nếu bạn quan tâm đến việc đóng góp cho dự án này, pkg.go.dev là mã nguồn mở! Hãy xem
[hướng dẫn đóng góp](https://go.googlesource.com/pkgsite/+/refs/heads/master/CONTRIBUTING.md)
để tìm hiểu thêm.
