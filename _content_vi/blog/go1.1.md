---
title: Go 1.1 đã được phát hành
date: 2013-05-13
by:
- Andrew Gerrand
tags:
- release
summary: Go 1.1 nhanh hơn, bớt khắt khe với câu lệnh return, và bổ sung method expression.
template: true
---


Chúng tôi vô cùng vui mừng thông báo bản phát hành Go 1.1.

{{image "go1.1/gopherbiplane5.jpg"}}

Tháng 3 năm ngoái chúng tôi đã phát hành Go 1.0, và kể từ đó chúng tôi đã phát hành thêm ba
bản cập nhật nhỏ "point release".
Các point release chỉ được tạo ra để sửa các vấn đề nghiêm trọng,
vì vậy Go 1.0.3 bạn đang dùng ngày nay về bản chất vẫn là
Go 1.0 chúng tôi phát hành vào tháng 3 năm 2012.

Go 1.1 bao gồm nhiều cải tiến so với 1.0.

Các cải tiến quan trọng nhất liên quan đến hiệu năng.
Chúng tôi đã tối ưu hóa trình biên dịch và trình liên kết,
bộ gom rác, bộ lập lịch goroutine, triển khai map,
và một số phần của thư viện chuẩn.
Có khả năng code Go của bạn sẽ chạy nhanh hơn đáng kể khi được biên dịch với Go 1.1.

Có một số thay đổi nhỏ đối với bản thân ngôn ngữ,
trong đó hai điểm đáng được nêu ra ở đây:
[các thay đổi về yêu cầu câu lệnh return](/doc/go1.1#return) sẽ
dẫn đến các chương trình súc tích và đúng đắn hơn,
và việc giới thiệu [giá trị method](/doc/go1.1#method_values) cung cấp
một cách diễn đạt để gắn một method với receiver của nó như một giá trị hàm.

Lập trình đồng thời an toàn hơn trong Go 1.1 với việc bổ sung
bộ phát hiện race condition để tìm lỗi đồng bộ hóa bộ nhớ trong các chương trình.
Chúng tôi sẽ thảo luận về bộ phát hiện race condition nhiều hơn trong một bài viết sắp tới,
nhưng hiện tại [hướng dẫn sử dụng](/doc/articles/race_detector.html) là
nơi tuyệt vời để bắt đầu.

Các công cụ và thư viện chuẩn đã được cải thiện và mở rộng.
Bạn có thể đọc toàn bộ thông tin trong [ghi chú phát hành](/doc/go1.1).

Theo [hướng dẫn tương thích](/doc/go1compat.html) của chúng tôi,
Go 1.1 vẫn tương thích với Go 1.0 và chúng tôi khuyến nghị tất cả người dùng Go nâng cấp lên bản phát hành mới.

Tất cả những điều này sẽ không thể thực hiện được nếu không có sự giúp đỡ của những người đóng góp từ
cộng đồng mã nguồn mở.
Kể từ Go 1.0, lõi nhận được hơn 2600 commit từ 161 người bên ngoài Google.
Cảm ơn mọi người vì thời gian và công sức.
Đặc biệt, chúng tôi muốn cảm ơn Shenghou Ma,
Rémy Oudompheng, Dave Cheney, Mikio Hara,
Alex Brainman, Jan Ziak và Daniel Morsing vì những đóng góp xuất sắc.

Để tải bản phát hành mới, hãy làm theo [hướng dẫn cài đặt](/doc/install) thông thường. Chúc lập trình vui vẻ!

_Cảm ơn Renée French vì chú Gopher!_
