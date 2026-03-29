---
title: Go 1.4 đã được phát hành
date: 2014-12-10
by:
- Andrew Gerrand
summary: Go 1.4 bổ sung hỗ trợ Android, go generate, tối ưu hóa, và nhiều hơn nữa.
---


Hôm nay chúng tôi thông báo Go 1.4, bản phát hành ổn định lớn thứ năm của Go, ra mắt sáu
tháng sau bản phát hành lớn trước của chúng tôi là [Go 1.3](/blog/go1.3).
Nó chứa một thay đổi ngôn ngữ nhỏ, hỗ trợ cho thêm hệ điều hành
và kiến trúc bộ xử lý, và cải tiến cho toolchain và thư viện.
Như thường lệ, Go 1.4 giữ lời hứa về tương thích, và hầu hết mọi thứ
sẽ tiếp tục biên dịch và chạy mà không thay đổi khi chuyển sang 1.4.
Để xem chi tiết đầy đủ, xem [ghi chú phát hành Go 1.4](/doc/go1.4).

Tính năng mới đáng chú ý nhất trong bản phát hành này là hỗ trợ chính thức cho Android.
Dùng hỗ trợ trong lõi và các thư viện trong
kho lưu trữ [golang.org/x/mobile](https://godoc.org/golang.org/x/mobile),
giờ đây có thể viết các ứng dụng Android đơn giản chỉ dùng code Go.
Ở giai đoạn này, các thư viện hỗ trợ vẫn còn non trẻ và đang trong quá trình phát triển mạnh.
Những người dùng sớm nên chuẩn bị cho trải nghiệm chưa mượt, nhưng chúng tôi hoan nghênh cộng đồng tham gia.

Thay đổi ngôn ngữ là một điều chỉnh nhỏ về cú pháp vòng lặp for-range.
Giờ bạn có thể viết "for range s {" để lặp qua từng mục trong s,
mà không cần gán giá trị, chỉ số vòng lặp hay khóa map.
Xem [ghi chú phát hành](/doc/go1.4#forrange) để biết chi tiết.

Lệnh go có một lệnh con mới, go generate, để tự động hóa việc chạy
các công cụ sinh code nguồn trước khi biên dịch.
Ví dụ, nó có thể dùng để tự động hóa việc sinh các phương thức String cho
hằng số có kiểu dùng
[công cụ stringer mới](https://godoc.org/golang.org/x/tools/cmd/stringer/).
Để biết thêm thông tin, xem [tài liệu thiết kế](/s/go1.4-generate).

Hầu hết các chương trình sẽ chạy với tốc độ tương đương hoặc nhanh hơn một chút trong 1.4 so với
1.3; một số sẽ chậm hơn một chút.
Có nhiều thay đổi, khiến việc đưa ra dự đoán chính xác là khó.
Xem [ghi chú phát hành](/doc/go1.4#performance) để thảo luận thêm.

Và tất nhiên, có nhiều cải tiến và sửa lỗi hơn nữa.

Nếu bạn chưa hay, vài tuần trước các kho lưu trữ con đã được chuyển đến vị trí mới.
Ví dụ, các gói go.tools nay được import từ "golang.org/x/tools".
Xem [bài thông báo](https://groups.google.com/d/msg/golang-announce/eD8dh3T9yyA/HDOEU_ZSmvAJ) để biết chi tiết.

Bản phát hành này cũng trùng với thời điểm dự án chuyển từ Mercurial sang Git (quản lý
mã nguồn), từ Rietveld sang Gerrit (đánh giá code), và từ Google Code sang
GitHub (theo dõi issue và wiki).
Việc chuyển đổi ảnh hưởng đến kho lưu trữ Go lõi và các kho lưu trữ con của nó.
Bạn có thể tìm thấy các kho lưu trữ Git chính thức tại
[go.googlesource.com](https://go.googlesource.com),
và issue tracker cùng wiki tại
[kho lưu trữ golang/go trên GitHub](https://github.com/golang/go).

Mặc dù quá trình phát triển đã chuyển sang cơ sở hạ tầng mới,
đối với bản phát hành 1.4, chúng tôi vẫn khuyến nghị người dùng
[cài đặt từ nguồn](/doc/install/source)
dùng các kho lưu trữ Mercurial.

Đối với người dùng App Engine, Go 1.4 nay có sẵn để kiểm thử beta.
Xem [thông báo](https://groups.google.com/d/msg/google-appengine-go/ndtQokV3oFo/25wV1W9JtywJ) để biết chi tiết.

Thay mặt tất cả mọi người trong nhóm Go, hãy tận hưởng Go 1.4, và chúc mừng mùa lễ!
