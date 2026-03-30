---
title: Kết quả Khảo sát Go 2016
date: 2017-03-06
by:
- Steve Francia, for the Go team
tags:
- survey
- community
summary: Những gì chúng tôi thu được từ Khảo sát Người dùng Go tháng 12 năm 2016.
template: true
---

## Cảm ơn

Bài đăng này tóm tắt kết quả khảo sát người dùng tháng 12 năm 2016 cùng với
các nhận xét và phân tích của chúng tôi.
Chúng tôi biết ơn tất cả mọi người đã cung cấp phản hồi qua khảo sát để giúp
định hình tương lai của Go.

## Nền tảng lập trình

Trong số 3.595 người tham gia khảo sát, 89% cho biết họ lập trình bằng Go tại nơi làm việc hoặc
bên ngoài công việc, trong đó 39% dùng Go cả ở nhà lẫn ở nơi làm việc, 27% chỉ dùng Go ở nhà,
và 23% chỉ dùng Go ở nơi làm việc.

Chúng tôi hỏi về các lĩnh vực mà mọi người làm việc.
63% cho biết họ làm trong phát triển web, nhưng chỉ 9% liệt kê phát triển web là lĩnh vực duy nhất.
Trên thực tế, 77% chọn hai hoặc nhiều lĩnh vực hơn, và 53% chọn ba lĩnh vực trở lên.

Chúng tôi cũng hỏi về loại chương trình mọi người viết bằng Go.
63% người tham gia viết chương trình dòng lệnh, 60% viết dịch vụ API hoặc RPC, và 52% viết dịch vụ web.
Giống như câu hỏi trước, hầu hết đều chọn nhiều lựa chọn,
với 85% chọn hai hoặc nhiều hơn và 72% chọn ba hoặc nhiều hơn.

Chúng tôi hỏi về mức độ thành thạo và sở thích của mọi người trong các ngôn ngữ lập trình.
Không có gì đáng ngạc nhiên, Go xếp hạng cao nhất trong lựa chọn đầu tiên của người tham gia ở cả
mức độ thành thạo (26%) và sở thích (62%).
Nếu loại trừ Go, năm lựa chọn đầu tiên về mức độ thành thạo ngôn ngữ là
Python (18%), Java (17%), JavaScript (13%), C (11%) và PHP (8%);
và năm lựa chọn đầu tiên về ngôn ngữ ưa thích là
Python (22%), JavaScript (10%), C (9%), Java (9%) và Ruby (7%).
Rõ ràng Go đang thu hút nhiều lập trình viên từ các ngôn ngữ động.

{{raw (file "survey2016/background.html")}}

## Sử dụng Go

Người dùng hài lòng với Go một cách áp đảo:
họ đồng ý rằng họ sẽ giới thiệu Go cho người khác theo tỷ lệ 19:1,
rằng họ muốn dùng Go cho dự án tiếp theo (14:1),
và rằng Go đang hoạt động tốt cho nhóm của họ (18:1).
Ít người dùng hơn đồng ý rằng Go là thiết yếu cho sự thành công của công ty họ (2,5:1).

Khi được hỏi điều gì họ thích nhất về Go, người dùng thường đề cập nhiều nhất đến
sự đơn giản, dễ sử dụng, tính năng đồng thời và hiệu suất của Go.
Khi được hỏi những thay đổi nào sẽ cải thiện Go nhiều nhất,
người dùng thường đề cập nhiều nhất đến generics, phiên bản gói và quản lý dependency.
Các phản hồi phổ biến khác bao gồm giao diện đồ họa (GUI), gỡ lỗi và xử lý lỗi.

Khi được hỏi về những thách thức lớn nhất trong việc sử dụng Go cá nhân,
người dùng đề cập đến nhiều thay đổi kỹ thuật được gợi ý trong câu hỏi trước.
Các chủ đề phổ biến nhất trong các thách thức phi kỹ thuật là thuyết phục người khác dùng Go
và truyền đạt giá trị của Go cho người khác, bao gồm cả cấp quản lý.
Một chủ đề phổ biến khác là học Go hoặc giúp người khác học,
bao gồm tìm kiếm tài liệu như hướng dẫn bắt đầu,
các bài hướng dẫn, ví dụ và thực hành tốt nhất.

Một số phản hồi tiêu biểu, đã được diễn đạt lại để đảm bảo bảo mật:

{{raw (file "survey2016/quotes.html")}}

Chúng tôi đánh giá cao phản hồi được cung cấp để xác định những thách thức mà người dùng và cộng đồng của chúng tôi đang đối mặt.
Trong năm 2017, chúng tôi tập trung vào việc giải quyết những vấn đề này và hy vọng sẽ thực hiện được nhiều cải tiến đáng kể nhất có thể.
Chúng tôi chào đón các đề xuất và đóng góp từ cộng đồng để biến những thách thức thành điểm mạnh cho Go.

{{raw (file "survey2016/usage.html")}}

## Phát triển và triển khai

Khi được hỏi hệ điều hành nào họ dùng để phát triển Go,
63% người tham gia cho biết họ sử dụng Linux, 44% dùng MacOS và 19% dùng Windows,
với nhiều lựa chọn được phép và 49% người tham gia phát triển trên nhiều hệ thống.
51% phản hồi chỉ chọn một hệ thống, phân bổ thành
29% trên Linux, 17% trên MacOS, 5% trên Windows và 0,2% trên các hệ thống khác.

Triển khai Go phân bổ tương đối đều nhau giữa các máy chủ do tư nhân quản lý và
các máy chủ đám mây được thuê ngoài.

{{raw (file "survey2016/dev.html")}}

## Làm việc Hiệu quả

Chúng tôi hỏi mức độ đồng ý hay không đồng ý của mọi người với các nhận định khác nhau về Go.
Người dùng đồng ý nhiều nhất rằng hiệu suất của Go đáp ứng nhu cầu của họ (tỷ lệ đồng ý so với không đồng ý là 57:1),
rằng họ có thể nhanh chóng tìm được câu trả lời cho câu hỏi của mình (20:1),
và rằng họ có thể sử dụng hiệu quả các tính năng đồng thời của Go (14:1).
Mặt khác, người dùng ít đồng ý nhất rằng họ có thể gỡ lỗi hiệu quả
các tính năng đồng thời của Go (2,7:1).

Người dùng hầu hết đồng ý rằng họ có thể nhanh chóng tìm thấy thư viện cần thiết (7,5:1).
Khi được hỏi còn thiếu thư viện nào, yêu cầu phổ biến nhất là thư viện để viết giao diện đồ họa (GUI).
Một chủ đề phổ biến khác là các yêu cầu xung quanh xử lý dữ liệu, phân tích và tính toán số học và khoa học.

Trong số 30% người dùng đề xuất cách cải thiện tài liệu của Go,
đề xuất phổ biến nhất là có thêm nhiều ví dụ hơn.

Các nguồn tin tức Go chính là blog Go,
/r/golang trên Reddit và Twitter;
có thể có một số thiên lệch ở đây vì đây cũng là cách khảo sát được thông báo.

Các nguồn chính để tìm câu trả lời cho các câu hỏi về Go là trang web Go,
Stack Overflow và đọc mã nguồn trực tiếp.

{{raw (file "survey2016/effective.html")}}

## Dự án Go

55% người tham gia bày tỏ sự quan tâm đến việc đóng góp theo một cách nào đó cho cộng đồng và các dự án Go.
Tiếc là, tương đối ít người đồng ý rằng họ cảm thấy được chào đón khi làm điều đó (3,3:1)
và thậm chí ít hơn cho rằng quy trình đó rõ ràng (1,3:1).
Trong năm 2017, chúng tôi dự định tập trung vào việc cải thiện quy trình đóng góp và tiếp tục
làm cho tất cả người đóng góp cảm thấy được chào đón.

Người tham gia đồng ý rằng họ tin tưởng vào ban lãnh đạo của dự án Go (9:1),
nhưng họ đồng ý ít hơn nhiều rằng ban lãnh đạo dự án hiểu nhu cầu của họ (2,6:1),
và họ đồng ý thậm chí ít hơn rằng họ cảm thấy thoải mái khi tiếp cận ban lãnh đạo dự án với câu hỏi và phản hồi (2,2:1).
Trên thực tế, đây là những câu hỏi duy nhất trong khảo sát mà hơn một nửa người tham gia
không đánh dấu "đồng ý một phần", "đồng ý" hoặc "hoàn toàn đồng ý" (nhiều người trung lập hoặc không trả lời).

Chúng tôi hy vọng rằng cuộc khảo sát và bài đăng blog này sẽ truyền đạt đến những ai
không thoải mái khi tiếp cận rằng ban lãnh đạo dự án Go đang lắng nghe.
Trong suốt năm 2017, chúng tôi sẽ khám phá các cách mới để tương tác với người dùng nhằm hiểu rõ hơn nhu cầu của họ.

{{raw (file "survey2016/project.html")}}

## Cộng đồng

Ở cuối khảo sát, chúng tôi hỏi một số câu hỏi về nhân khẩu học.
Phân phối quốc gia của các phản hồi gần tương ứng với phân phối quốc gia của lượng truy cập vào golang.org,
nhưng các phản hồi thiếu đại diện từ một số quốc gia châu Á.
Cụ thể, Ấn Độ, Trung Quốc và Nhật Bản mỗi nước chiếm khoảng 5% lượng truy cập vào golang.org trong năm 2016
nhưng chỉ chiếm 3%, 2% và 1% phản hồi khảo sát.

Một phần quan trọng của cộng đồng là làm cho mọi người cảm thấy được chào đón,
đặc biệt là những người từ các nhóm ít được đại diện.
Chúng tôi đặt một câu hỏi tùy chọn về nhận dạng trong một vài nhóm đa dạng.
37% người tham gia để trống câu hỏi và 12% chọn "Tôi không muốn trả lời",
vì vậy chúng tôi không thể rút ra nhiều kết luận rộng từ dữ liệu.
Tuy nhiên, một so sánh nổi bật: 9% người xác nhận họ thuộc nhóm thiểu số đồng ý
với nhận định "Tôi cảm thấy được chào đón trong cộng đồng Go" theo tỷ lệ 7,5:1,
so với 15:1 trong toàn bộ khảo sát.
Chúng tôi muốn làm cho cộng đồng Go trở nên thân thiện hơn nữa.
Chúng tôi ủng hộ và được khích lệ bởi những nỗ lực của các tổ chức như GoBridge và Women Who Go.

Câu hỏi cuối cùng của khảo sát chỉ là để vui: từ khóa Go yêu thích của bạn là gì?
Có lẽ không có gì đáng ngạc nhiên, câu trả lời phổ biến nhất là `go`, tiếp theo là `defer`, `func`, `interface` và `select`.

{{raw (file "survey2016/community.html")}}
