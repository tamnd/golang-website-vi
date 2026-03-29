---
title: Go 1.7 đã được phát hành
date: 2016-08-15
by:
- Chris Broadfoot
summary: Go 1.7 bổ sung code x86 được biên dịch nhanh hơn, context trong thư viện chuẩn, và nhiều hơn nữa.
---


Hôm nay chúng tôi vui mừng thông báo phát hành Go 1.7.
Bạn có thể tải về từ [trang download](/dl/).
Có một số thay đổi quan trọng trong bản phát hành này: cổng cho
[Linux trên IBM z Systems](https://en.wikipedia.org/wiki/IBM_System_z) (s390x),
cải tiến trình biên dịch, bổ sung gói [context](/pkg/context/),
và hỗ trợ [kiểm thử và benchmark phân cấp](/pkg/testing/#hdr-Subtests_and_Sub_benchmarks).

Một backend trình biên dịch mới, dựa trên dạng [static single-assignment](https://en.wikipedia.org/wiki/Static_single_assignment_form) (SSA),
đã được phát triển trong năm qua.
Bằng cách biểu diễn một chương trình dưới dạng SSA, trình biên dịch có thể thực hiện các tối ưu hóa nâng cao dễ dàng hơn.
Backend mới này tạo ra code gọn hơn, hiệu quả hơn bao gồm các tối ưu hóa như
[loại bỏ kiểm tra biên giới](https://en.wikipedia.org/wiki/Bounds-checking_elimination) và
[loại bỏ biểu thức con chung](https://en.wikipedia.org/wiki/Common_subexpression_elimination).
Chúng tôi quan sát thấy tốc độ tăng từ 5-35% trên các benchmark.
Hiện tại, backend mới chỉ có sẵn cho nền tảng x86 64-bit ("amd64"),
nhưng chúng tôi đang lên kế hoạch chuyển đổi thêm backend kiến trúc sang SSA trong các bản phát hành tương lai.

Frontend trình biên dịch dùng định dạng dữ liệu export mới, gọn hơn và
xử lý khai báo import hiệu quả hơn.
Mặc dù những [thay đổi trong toolchain trình biên dịch](/doc/go1.7#compiler) này hầu hết không thấy được,
người dùng đã [quan sát](http://dave.cheney.net/2016/04/02/go-1-7-toolchain-improvements)
tốc độ biên dịch tăng đáng kể và kích thước binary giảm tới 20-30%.

Các chương trình nên chạy nhanh hơn một chút nhờ tăng tốc trong bộ gom rác và tối ưu hóa trong thư viện chuẩn.
Các chương trình có nhiều goroutine nhàn rỗi sẽ trải qua thời gian tạm dừng thu gom rác ngắn hơn nhiều so với Go 1.6.

Trong vài năm qua, gói
[golang.org/x/net/context](https://godoc.org/golang.org/x/net/context/)
đã chứng tỏ là thiết yếu với nhiều ứng dụng Go.
Context được dùng hiệu quả trong các ứng dụng liên quan đến mạng, hạ tầng và microservice
(chẳng hạn như [Kubernetes](http://kubernetes.io/) và [Docker](https://www.docker.com/)).
Chúng giúp dễ dàng bật hủy bỏ, timeout và truyền dữ liệu theo phạm vi request.
Để sử dụng context trong thư viện chuẩn và khuyến khích sử dụng rộng rãi hơn,
gói đã được chuyển từ kho lưu trữ [x/net](https://godoc.org/golang.org/x/net/context/)
vào thư viện chuẩn như gói [context](/pkg/context/).
Hỗ trợ context đã được thêm vào các gói
[net](/pkg/net/),
[net/http](/pkg/net/http/) và
[os/exec](/pkg/os/exec/).
Để biết thêm về context, xem [tài liệu gói](/pkg/context)
và bài đăng blog Go [_Go Concurrency Patterns: Context_](/blog/context).

Go 1.5 đã giới thiệu hỗ trợ thử nghiệm cho [thư mục "vendor"](/cmd/go/#hdr-Vendor_Directories),
được bật bằng biến môi trường `GO15VENDOREXPERIMENT`.
Go 1.6 đã bật hành vi này theo mặc định, và trong Go 1.7, switch này đã bị loại bỏ và hành vi "vendor" luôn được bật.

Go 1.7 bao gồm nhiều bổ sung, cải tiến và sửa lỗi hơn nữa.
Tìm tập thay đổi đầy đủ, và chi tiết của các điểm trên, trong
[ghi chú phát hành Go 1.7](/doc/go1.7.html).

Cuối cùng, nhóm Go muốn cảm ơn tất cả những người đã đóng góp cho bản phát hành.
170 người đã đóng góp cho bản phát hành này, trong đó 140 người từ cộng đồng Go.
Các đóng góp này từ thay đổi cho trình biên dịch và trình liên kết, đến thư viện chuẩn, đến tài liệu và đánh giá code.
Chúng tôi hoan nghênh đóng góp; nếu bạn muốn tham gia, xem
[hướng dẫn đóng góp](/doc/contribute.html).
