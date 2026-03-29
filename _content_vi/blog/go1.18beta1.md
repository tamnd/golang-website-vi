---
title: Go 1.18 Beta 1 đã có mặt, kèm generics
date: 2021-12-14
by:
- Russ Cox, for the Go team
summary: Go 1.18 Beta 1 là bản xem trước đầu tiên của Go 1.18. Hãy thử và cho chúng tôi biết nếu bạn gặp vấn đề.
---

Chúng tôi vừa phát hành Go 1.18 Beta 1,
bạn có thể tải về từ [trang download](/dl/#go1.18beta1).

Bản phát hành chính thức Go 1.18 sẽ chưa có trong vài tháng nữa.
Đây là bản xem trước đầu tiên của Go 1.18, để bạn thử nghiệm,
trải nghiệm và cho chúng tôi biết bạn gặp phải vấn đề gì.
Go 1.18 Beta 1 là kết quả của một khối lượng công việc khổng lồ
từ toàn bộ nhóm Go tại Google và những người đóng góp Go trên khắp thế giới,
và chúng tôi rất háo hức được nghe ý kiến của bạn.

Go 1.18 Beta 1 là bản xem trước đầu tiên chứa
hỗ trợ mới của Go cho [code generic dùng kiểu được tham số hóa](/blog/why-generics).
Generics là thay đổi quan trọng nhất đối với Go kể từ khi phát hành Go 1,
và chắc chắn là thay đổi ngôn ngữ đơn lẻ lớn nhất chúng tôi từng thực hiện.
Với bất kỳ tính năng lớn mới nào, người dùng mới thường phát hiện ra bug mới,
và chúng tôi không kỳ vọng generics là ngoại lệ;
hãy tiếp cận chúng với sự cẩn thận phù hợp.
Ngoài ra, một số trường hợp tinh tế, chẳng hạn như các loại generic đệ quy cụ thể,
đã được hoãn lại cho các bản phát hành tương lai.
Dù vậy, chúng tôi biết những người dùng sớm khá hài lòng,
và nếu bạn có trường hợp sử dụng mà bạn cho là phù hợp đặc biệt với generics,
chúng tôi hy vọng bạn sẽ thử.
Chúng tôi đã xuất bản
[hướng dẫn ngắn về cách bắt đầu với generics](/doc/tutorial/generics)
và đã trình bày
[tại GopherCon tuần trước](https://www.youtube.com/watch?v=35eIxI_n5ZM&t=1755s).
Bạn thậm chí có thể thử ngay trên
[Go playground ở chế độ Go dev branch](/play/?v=gotip).

Go 1.18 Beta 1 bổ sung hỗ trợ sẵn có để viết
[kiểm thử dựa trên fuzzing](/blog/fuzz-beta),
tự động tìm đầu vào khiến chương trình crash hoặc trả về kết quả không hợp lệ.

Go 1.18 Beta 1 thêm "[Go workspace mode](/design/45713-workspace)" mới,
cho phép bạn làm việc với nhiều module Go cùng lúc,
một trường hợp sử dụng quan trọng cho các dự án lớn.

Go 1.18 Beta 1 chứa lệnh `go version -m` được mở rộng,
nay ghi lại các chi tiết build như cờ compiler.
Một chương trình có thể truy vấn thông tin build của chính nó dùng
[debug.ReadBuildInfo](https://pkg.go.dev/runtime/debug@master#BuildInfo),
và giờ có thể đọc chi tiết build từ các binary khác dùng gói
[debug/buildinfo](https://pkg.go.dev/debug/buildinfo@master) mới.
Chức năng này được thiết kế làm nền tảng
cho bất kỳ công cụ nào cần tạo danh sách nguyên liệu phần mềm (SBOM) cho các binary Go.

Đầu năm nay, Go 1.17 đã thêm quy ước gọi hàm dựa trên register
để tăng tốc code Go trên hệ thống x86-64.
Go 1.18 Beta 1 mở rộng tính năng đó sang ARM64 và PPC64,
mang lại tốc độ nhanh hơn tới 20%.

Cảm ơn tất cả những người đã đóng góp cho bản beta này,
và đặc biệt cho nhóm tại Google đã
làm việc không ngừng nghỉ trong nhiều năm để biến generics thành hiện thực.
Đó là một chặng đường dài, chúng tôi rất hài lòng với kết quả,
và chúng tôi hy vọng bạn cũng thích.

Xem đầy đủ [ghi chú phát hành dự thảo cho Go 1.18](https://tip.golang.org/doc/go1.18) để biết thêm chi tiết.

Như thường lệ, đặc biệt với bản beta, nếu bạn phát hiện bất kỳ vấn đề nào,
hãy [tạo issue](/issue/new).

Chúng tôi hy vọng bạn thích kiểm thử bản beta,
và chúc tất cả mọi người có phần còn lại của năm 2021 thư thái.
Chúc mừng lễ!
