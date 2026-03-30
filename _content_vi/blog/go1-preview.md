---
title: Xem trước Go phiên bản 1
date: 2011-10-05
by:
- Russ Cox
tags:
- go1
- release
summary: Những gì nhóm Go đang lên kế hoạch cho Go phiên bản 1.
---


Chúng tôi muốn có thể cung cấp một nền tảng ổn định cho người dùng Go.
Mọi người nên có thể viết các chương trình Go và kỳ vọng rằng chúng sẽ tiếp tục
biên dịch và chạy được mà không thay đổi,
trong thời gian tính bằng năm.
Tương tự, mọi người nên có thể viết sách về Go,
có thể nói phiên bản Go nào cuốn sách đang mô tả,
và số phiên bản đó vẫn còn ý nghĩa về sau.
Không có thuộc tính nào trong số này đúng với Go hiện tại.

Chúng tôi đề xuất phát hành một bản Go vào đầu năm tới có tên "Go phiên bản 1",
viết tắt là Go 1, đây sẽ là bản phát hành Go đầu tiên ổn định theo cách này.
Code biên dịch được trong Go phiên bản 1 sẽ,
với một vài ngoại lệ, tiếp tục biên dịch được trong suốt vòng đời của phiên bản đó,
khi chúng tôi phát hành các bản cập nhật và sửa lỗi như Go phiên bản 1.1, 1.2, v.v.
Nó cũng sẽ được duy trì với các bản sửa lỗi và vá bảo mật ngay cả khi
các phiên bản khác có thể tiếp tục phát triển.
Ngoài ra, các môi trường production như Google App Engine sẽ hỗ trợ nó
trong một thời gian dài.

Go phiên bản 1 sẽ là ngôn ngữ ổn định với thư viện ổn định.
Ngoài các bản sửa lỗi nghiêm trọng, các thay đổi đối với thư viện và gói trong phiên bản 1.1,
1.2 v.v. có thể bổ sung chức năng nhưng sẽ không làm hỏng các chương trình Go phiên bản 1 hiện có.

Mục tiêu của chúng tôi là Go 1 là phiên bản ổn định của Go hiện tại,
không phải tái thiết kế toàn bộ ngôn ngữ.
Đặc biệt, chúng tôi đang rõ ràng phản đối bất kỳ nỗ lực nào thiết kế các tính năng ngôn ngữ mới
"theo ủy ban".

Tuy nhiên, có nhiều thay đổi khác nhau đối với ngôn ngữ và gói Go mà
chúng tôi đã dự định từ lâu và đã tạo prototype nhưng chưa triển khai,
chủ yếu vì chúng quan trọng và không tương thích ngược.
Nếu Go 1 tồn tại lâu dài, điều quan trọng là chúng tôi phải lên kế hoạch,
thông báo, triển khai và kiểm thử các thay đổi này như một phần của quá trình chuẩn bị Go 1,
thay vì trì hoãn cho đến sau khi phát hành và do đó tạo ra
sự phân kỳ mâu thuẫn với mục tiêu của chúng tôi.

Hôm nay, chúng tôi đang công bố [kế hoạch sơ bộ cho Go 1](https://docs.google.com/document/pub?id=1ny8uI-_BHrDCZv_zNBSthNKAMX_fR_0dc6epA6lztRE)
để lấy phản hồi từ cộng đồng Go.
Nếu bạn có phản hồi, hãy trả lời vào [chủ đề trên danh sách thư golang-nuts](http://groups.google.com/group/golang-nuts/browse_thread/thread/badc4f323431a4f6).
