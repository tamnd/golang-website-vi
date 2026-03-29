---
title: Đề xuất cho Go 1.15
date: 2020-01-28
by:
- Robert Griesemer, for the Go team
tags:
- go1.15
- proposals
- community
- language
- vet
summary: Đối với Go 1.15, chúng tôi đề xuất ba thay đổi dọn dẹp ngôn ngữ nhỏ.
---

## Tình trạng

Chúng ta đang gần với bản phát hành Go 1.14, dự kiến vào tháng 2 nếu mọi thứ diễn ra
suôn sẻ, với RC1 gần sẵn sàng. Theo quy trình được phác thảo trong bài đăng blog
[Go 2, here we come!](/blog/go2-here-we-come),
đây lại là thời điểm trong chu kỳ phát triển và phát hành của chúng ta để xem xét liệu
và những thay đổi ngôn ngữ hoặc thư viện nào chúng ta có thể muốn đưa vào bản phát hành tiếp theo,
Go 1.15, dự kiến vào tháng 8 năm nay.

Các mục tiêu chính của Go vẫn là quản lý gói và phiên bản, hỗ trợ xử lý lỗi tốt hơn,
và generics. Hỗ trợ module đang trong tình trạng tốt và ngày càng được cải thiện,
và chúng tôi cũng đang tiến triển về mặt generics (sẽ có thêm về điều đó sau trong năm nay).
Nỗ lực bảy tháng trước của chúng tôi trong việc cung cấp cơ chế xử lý lỗi tốt hơn,
[đề xuất `try`](/issue/32437), nhận được sự ủng hộ tốt
nhưng cũng có sự phản đối mạnh mẽ và chúng tôi quyết định từ bỏ nó. Sau đó có
nhiều đề xuất tiếp theo, nhưng không cái nào trong số họ có vẻ đủ thuyết phục,
rõ ràng tốt hơn đề xuất `try`, hoặc ít có khả năng gây tranh cãi tương tự.
Do đó, hiện tại chúng tôi không tiếp tục theo đuổi các thay đổi đối với xử lý lỗi.
Có lẽ một cái nhìn sâu sắc trong tương lai sẽ giúp chúng ta cải thiện hiện trạng.

## Đề xuất

Với modules và generics đang được tích cực phát triển, và với các thay đổi xử lý lỗi
tạm thời được gác lại, chúng ta nên theo đuổi những thay đổi nào khác, nếu có? Có một số
yêu cầu thường xuyên như enum và kiểu bất biến, nhưng không có ý tưởng nào trong số đó
được phát triển đủ, cũng không đủ cấp bách để được nhóm Go chú ý nhiều,
đặc biệt khi cũng xem xét chi phí của việc thay đổi ngôn ngữ.

Sau khi xem xét tất cả các đề xuất có thể khả thi, và quan trọng hơn, vì
chúng tôi không muốn thêm tính năng mới theo từng bước mà không có kế hoạch dài hạn, chúng tôi
kết luận rằng tốt hơn là nên chờ đợi với các thay đổi lớn lần này. Thay vào đó
chúng tôi tập trung vào một vài kiểm tra `vet` mới và điều chỉnh nhỏ đối với
ngôn ngữ. Chúng tôi đã chọn ba đề xuất sau:

[\#32479](/issue/32479).
Chẩn đoán chuyển đổi `string(int)` trong `go vet`.

Chúng tôi đã lên kế hoạch hoàn thành điều này cho bản phát hành Go 1.14 sắp tới nhưng chưa
kịp làm, vì vậy đây là lần nữa. Chuyển đổi `string(int)` được giới thiệu
sớm trong Go để thuận tiện, nhưng nó gây nhầm lẫn cho người mới (`string(10)` là
`"\n"` không phải `"10"`) và không còn hợp lý nữa khi chuyển đổi có sẵn
trong gói `unicode/utf8`.
Vì [loại bỏ chuyển đổi này](/issue/3939) không phải là thay đổi tương thích ngược,
chúng tôi đề xuất bắt đầu bằng lỗi `vet` thay thế.

[\#4483](/issue/4483).
Chẩn đoán type assertion interface-interface không thể xảy ra trong `go vet`.

Hiện tại, Go cho phép bất kỳ type assertion `x.(T)` nào (và trường hợp type switch tương ứng)
khi kiểu của `x` và `T` đều là interface. Tuy nhiên, nếu cả `x` và `T` đều có phương thức
cùng tên nhưng chữ ký khác nhau thì không thể có giá trị nào được gán
cho `x` cũng triển khai `T`; các type assertion như vậy sẽ luôn thất bại tại runtime
(panic hoặc trả về `false`). Vì chúng ta biết điều này tại compile time, trình biên dịch
hoàn toàn có thể báo lỗi. Tuy nhiên báo lỗi trình biên dịch trong trường hợp này không phải là
thay đổi tương thích ngược, vì vậy chúng tôi cũng đề xuất bắt đầu bằng lỗi `vet` thay thế.

[\#28591](/issue/28591).
Tính toán hằng số các biểu thức chỉ số và slice với chuỗi và chỉ số hằng số.

Hiện tại, việc lập chỉ số hoặc slice một chuỗi hằng số với chỉ số hằng số
tạo ra giá trị `byte` hoặc `string` không phải hằng số, tương ứng. Nhưng nếu tất cả các toán hạng
là hằng số, trình biên dịch có thể tính toán hằng số các biểu thức đó và tạo ra
kết quả hằng số (có thể không được định kiểu). Đây là thay đổi hoàn toàn tương thích ngược
và chúng tôi đề xuất thực hiện các điều chỉnh cần thiết cho đặc tả và trình biên dịch.

(Đính chính: Chúng tôi phát hiện ra sau khi đăng rằng thay đổi này không tương thích ngược;
xem [bình luận](/issue/28591#issuecomment-579993684) để biết chi tiết.)

## Lịch trình

Chúng tôi tin rằng không có đề xuất nào trong ba đề xuất này gây tranh cãi, nhưng luôn có
khả năng chúng tôi đã bỏ sót điều gì đó quan trọng. Vì lý do đó, chúng tôi lên kế hoạch
triển khai các đề xuất vào đầu chu kỳ phát hành Go 1.15
(vào hoặc ngay sau khi phát hành Go 1.14) để có đủ thời gian để
thu thập kinh nghiệm và cung cấp phản hồi. Theo
[quy trình đánh giá đề xuất](/blog/go2-here-we-come),
quyết định cuối cùng sẽ được đưa ra vào cuối chu kỳ phát triển, vào
đầu tháng 5 năm 2020.

## Và thêm một điều nữa...

Chúng tôi nhận được nhiều đề xuất thay đổi ngôn ngữ hơn
([các issue được gắn nhãn LanguageChange](https://github.com/golang/go/labels/LanguageChange))
so với những gì chúng tôi có thể xem xét kỹ lưỡng. Ví dụ, chỉ riêng về xử lý lỗi,
có 57 issue, trong đó năm issue hiện vẫn còn mở. Vì chi phí
thay đổi ngôn ngữ, dù nhỏ đến đâu, là cao và lợi ích
thường không rõ ràng, chúng tôi phải thận trọng về phía cẩn thận. Do đó, hầu hết
các đề xuất thay đổi ngôn ngữ bị từ chối sớm hay muộn, đôi khi với phản hồi tối thiểu.
Điều này không làm hài lòng tất cả các bên liên quan. Nếu bạn đã dành
nhiều thời gian và công sức để phác thảo ý tưởng của mình một cách chi tiết, sẽ là tốt nếu
nó không bị từ chối ngay lập tức. Mặt khác, vì
[quy trình đề xuất](https://github.com/golang/proposal/blob/master/README.md) chung
cố tình đơn giản, rất dễ tạo ra các đề xuất thay đổi ngôn ngữ
chỉ được khám phá ở mức độ tối thiểu, gây ra cho ủy ban đánh giá một khối lượng công việc đáng kể.
Để cải thiện trải nghiệm này cho mọi người, chúng tôi đang thêm
[bảng câu hỏi](https://github.com/golang/proposal/blob/master/go2-language-changes.md) mới
cho các thay đổi ngôn ngữ: điền vào mẫu đó sẽ giúp người đánh giá đánh giá
các đề xuất hiệu quả hơn vì họ không cần cố gắng tự trả lời những câu hỏi đó.
Và hy vọng nó cũng cung cấp hướng dẫn tốt hơn cho người đề xuất bằng cách đặt kỳ vọng
ngay từ đầu. Đây là một thử nghiệm mà chúng tôi sẽ tinh chỉnh theo thời gian khi cần.

Cảm ơn bạn đã giúp chúng tôi cải thiện trải nghiệm Go!
