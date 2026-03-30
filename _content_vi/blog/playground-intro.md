---
title: Giới thiệu Go Playground
date: 2010-09-15
by:
- Andrew Gerrand
tags:
- playground
summary: "Thông báo về Go Playground, https://play.golang.org/."
template: true
---


Nếu bạn truy cập [golang.org](/) hôm nay, bạn sẽ thấy giao diện mới của chúng tôi.
Chúng tôi đã khoác lên trang web một lớp sơn mới và tổ chức lại nội dung
để dễ tìm kiếm hơn.
Những thay đổi này cũng được phản ánh trong giao diện web của [godoc](/cmd/godoc/),
công cụ tài liệu Go.
Nhưng tin thực sự đáng chú ý là một tính năng mới nổi bật: [Go Playground](/).

{{image "playground-intro/screenshot.png"}}

Playground cho phép bất kỳ ai có trình duyệt web đều có thể viết code Go mà chúng tôi
ngay lập tức biên dịch, liên kết và chạy trên các máy chủ của mình.
Có một vài chương trình ví dụ để giúp bạn bắt đầu (xem menu thả xuống "Examples").
Chúng tôi hy vọng rằng điều này sẽ cho các lập trình viên tò mò cơ hội thử ngôn ngữ
trước khi [cài đặt nó](/doc/install.html),
và cung cấp cho người dùng Go có kinh nghiệm một nơi thuận tiện để thử nghiệm.
Ngoài trang chủ, tính năng này có tiềm năng làm cho các tài liệu tham khảo và hướng dẫn
của chúng tôi trở nên sinh động hơn.
Chúng tôi hy vọng sẽ mở rộng việc sử dụng nó trong tương lai gần.

Tất nhiên, có một số hạn chế đối với các loại chương trình bạn có thể chạy trong Playground.
Chúng tôi không thể đơn giản chấp nhận code tùy ý và chạy nó trên máy chủ của mình mà không có hạn chế.
Các chương trình được xây dựng và chạy trong một sandbox với thư viện chuẩn được thu gọn;
giao tiếp duy nhất của chương trình với thế giới bên ngoài là qua đầu ra chuẩn,
và có giới hạn về việc sử dụng CPU và bộ nhớ.
Do vậy, hãy coi đây chỉ là một hương vị của thế giới Go tuyệt vời;
để có trải nghiệm đầy đủ, bạn cần [tự tải về](/doc/install.html).
Nếu bạn đã có ý định thử Go nhưng chưa có cơ hội,
tại sao không ghé thăm [golang.org](/) để thử ngay bây giờ?
