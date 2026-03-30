---
title: Hướng tới Go 2
date: 2017-07-13
by:
- Russ Cox
tags:
- community
summary: Cách tất cả chúng ta sẽ cùng nhau hướng tới Go 2.
template: true
---

## Giới thiệu

[Đây là nội dung bài nói chuyện của tôi hôm nay tại](https://www.youtube.com/watch?v=0Zbh_vmAKvk)
Gophercon 2017, kêu gọi toàn bộ cộng đồng Go cùng tham gia
thảo luận và lên kế hoạch cho Go 2.]

Vào ngày 25 tháng 9 năm 2007, sau khi Rob Pike, Robert Griesemer và Ken
Thompson đã thảo luận về một ngôn ngữ lập trình mới được vài
ngày, Rob đề xuất tên "Go."

<div style="margin-left: 2em;">
{{image "toward-go2/mail.png" 446}}
</div>

Năm tiếp theo, Ian Lance Taylor và tôi gia nhập nhóm, và cùng
năm người chúng tôi đã xây dựng hai trình biên dịch và một thư viện chuẩn, dẫn đến
[bản phát hành mã nguồn mở](https://opensource.googleblog.com/2009/11/hey-ho-lets-go.html) vào ngày 10 tháng 11 năm 2009.

<div style="margin-left: 2em;">
{{image "toward-go2/tweet.png" 467}}
</div>

Trong hai năm tiếp theo, với sự giúp đỡ của cộng đồng Go mã nguồn mở mới,
chúng tôi đã thử nghiệm các thay đổi lớn và nhỏ, tinh chỉnh Go
và dẫn đến [kế hoạch cho Go 1](/blog/preview-of-go-version-1), được đề xuất vào ngày 5 tháng 10 năm 2011.

<div style="margin-left: 2em;">
{{image "toward-go2/go1-preview.png" 560}}
</div>

Với sự giúp đỡ thêm từ cộng đồng Go, chúng tôi đã sửa đổi và triển khai kế hoạch đó,
cuối cùng [phát hành Go 1](/blog/go1) vào ngày 28 tháng 3 năm 2012.

<div style="margin-left: 2em;">
{{image "toward-go2/go1-release.png" 556}}
</div>

Bản phát hành Go 1 đánh dấu đỉnh cao của gần năm năm
nỗ lực sáng tạo, sôi nổi đưa chúng tôi từ một cái tên và danh sách ý tưởng
đến một ngôn ngữ ổn định, sẵn sàng cho môi trường production. Nó cũng đánh dấu sự chuyển đổi rõ ràng
từ thay đổi và biến động sang ổn định.

Trong những năm dẫn đến Go 1, chúng tôi đã thay đổi Go và phá vỡ chương trình Go của mọi người
gần như mỗi tuần. Chúng tôi hiểu rằng điều này ngăn Go
được sử dụng trong môi trường production, nơi các chương trình không thể được viết lại
hàng tuần để theo kịp các thay đổi ngôn ngữ.
Như [bài đăng blog thông báo Go 1](/blog/go1) nói, động lực chính là cung cấp một nền tảng ổn định
để tạo ra các sản phẩm, dự án và ấn phẩm đáng tin cậy (blog,
hướng dẫn, bài nói hội nghị và sách), để người dùng tự tin rằng
các chương trình của họ sẽ tiếp tục biên dịch và chạy không thay đổi
trong nhiều năm tới.

Sau khi Go 1 được phát hành, chúng tôi biết rằng chúng tôi cần dành thời gian sử dụng Go
trong các môi trường production mà nó được thiết kế cho. Chúng tôi chuyển hướng
rõ ràng khỏi việc thực hiện thay đổi ngôn ngữ sang sử dụng Go trong
các dự án của chính mình và cải thiện việc triển khai: chúng tôi đã chuyển Go sang nhiều
hệ thống mới, chúng tôi đã viết lại gần như mọi phần quan trọng về hiệu suất để
làm cho Go chạy hiệu quả hơn, và chúng tôi đã thêm các công cụ chính như
[bộ phát hiện race condition](/blog/race-detector).

Bây giờ chúng tôi có năm năm kinh nghiệm sử dụng Go để xây dựng các hệ thống
lớn, chất lượng production. Chúng tôi đã phát triển khả năng nhận biết điều gì hoạt động
và điều gì không. Bây giờ là lúc bắt đầu bước tiếp theo trong sự tiến hóa và phát triển của Go,
lên kế hoạch cho tương lai của Go. Tôi ở đây hôm nay để xin
tất cả các bạn trong cộng đồng Go, dù bạn đang ở khán phòng tại
GopherCon hay xem qua video hay đọc blog Go sau đó hôm nay,
hãy cùng làm việc với chúng tôi khi chúng tôi lên kế hoạch và triển khai Go 2.

Trong phần còn lại của bài nói này, tôi sẽ giải thích các mục tiêu của chúng tôi cho Go 2;
các ràng buộc và hạn chế; quy trình tổng thể; tầm quan trọng của
việc viết về trải nghiệm sử dụng Go của chúng ta, đặc biệt khi chúng liên quan đến
các vấn đề mà chúng tôi có thể cố gắng giải quyết; các loại giải pháp có thể; cách
chúng tôi sẽ phát hành Go 2; và cách tất cả các bạn có thể giúp đỡ.

## Mục tiêu

Các mục tiêu chúng tôi có cho Go hôm nay giống như năm 2007. Chúng tôi muốn
làm cho các lập trình viên hiệu quả hơn trong việc quản lý hai loại quy mô:
quy mô production, đặc biệt là các hệ thống đồng thời tương tác với nhiều
máy chủ khác, được thể hiện ngày nay bởi phần mềm đám mây; và quy mô phát triển,
đặc biệt là các codebase lớn được nhiều kỹ sư làm việc
với sự phối hợp lỏng lẻo, được thể hiện ngày nay bởi sự phát triển
mã nguồn mở hiện đại.

Các loại quy mô này xuất hiện ở các công ty thuộc mọi quy mô. Ngay cả một startup
năm người cũng có thể sử dụng các dịch vụ API đám mây lớn được cung cấp bởi
các công ty khác và sử dụng nhiều phần mềm mã nguồn mở hơn phần mềm họ
tự viết. Quy mô production và quy mô phát triển cũng quan trọng đối với
startup đó như đối với Google.

Mục tiêu của chúng tôi cho Go 2 là sửa các cách quan trọng nhất mà Go không
đáp ứng được yêu cầu về quy mô.

(Để biết thêm về các mục tiêu này, xem
bài viết năm 2012 của Rob Pike "[Go at Google: Language Design in the Service of Software Engineering](/talks/2012/splash.article)"
và bài nói GopherCon 2015 của tôi "[Go, Open Source, Community](/blog/open-source).")

## Các ràng buộc

Các mục tiêu cho Go chưa thay đổi kể từ đầu, nhưng
các ràng buộc đối với Go chắc chắn đã thay đổi. Ràng buộc quan trọng nhất là
việc sử dụng Go hiện tại. Chúng tôi ước tính có ít nhất
[nửa triệu lập trình viên Go trên toàn thế giới](https://research.swtch.com/gophercount),
có nghĩa là có hàng triệu tệp nguồn Go và ít nhất
một tỷ dòng mã Go. Những lập trình viên đó và mã nguồn đó
đại diện cho thành công của Go, nhưng họ cũng là ràng buộc chính
đối với Go 2.

Go 2 phải đưa theo tất cả những lập trình viên đó. Chúng tôi phải yêu cầu họ
quên đi thói quen cũ và học thói quen mới chỉ khi phần thưởng đủ lớn.
Ví dụ, trước Go 1, phương thức được triển khai bởi các kiểu lỗi được
đặt tên là `String`. Trong Go 1, chúng tôi đổi tên thành `Error`, để phân biệt kiểu lỗi
với các kiểu khác có thể định dạng chính chúng. Một hôm tôi đang
triển khai một kiểu lỗi, và không suy nghĩ tôi đặt tên phương thức của nó
là `String` thay vì `Error`, điều này tất nhiên không biên dịch được. Sau năm
năm tôi vẫn chưa hoàn toàn quên được cách cũ. Loại đổi tên làm rõ nghĩa đó
là một thay đổi quan trọng cần thực hiện trong Go 1 nhưng sẽ
quá gây rối cho Go 2 nếu không có lý do rất chính đáng.

Go 2 cũng phải đưa theo tất cả mã nguồn Go 1 hiện có. Chúng tôi không được
chia rẽ hệ sinh thái Go. Các chương trình hỗn hợp, trong đó các package được viết
bằng Go 2 import các package được viết bằng Go 1 và ngược lại, phải hoạt động
liền mạch trong thời gian chuyển đổi nhiều năm. Chúng tôi sẽ phải
tìm ra chính xác cách làm điều đó; các công cụ tự động như go fix
chắc chắn sẽ đóng một phần.

Để giảm thiểu sự gián đoạn, mỗi thay đổi sẽ cần được suy nghĩ cẩn thận,
lên kế hoạch và tạo công cụ hỗ trợ, điều này lại giới hạn số lượng thay đổi chúng tôi
có thể thực hiện. Có thể chúng tôi có thể làm hai hoặc ba, chắc chắn không hơn năm.

Tôi không tính các thay đổi nhỏ như có thể cho phép các định danh
bằng nhiều ngôn ngữ hơn hoặc thêm hằng số nguyên dạng nhị phân. Các thay đổi nhỏ
như thế này cũng quan trọng, nhưng dễ thực hiện đúng hơn. Tôi đang tập trung hôm nay
vào các thay đổi lớn có thể, chẳng hạn như
hỗ trợ bổ sung để xử lý lỗi, hoặc giới thiệu các giá trị bất biến hoặc
chỉ đọc, hoặc thêm một dạng generics, hoặc các chủ đề quan trọng khác
chưa được đề xuất. Chúng tôi chỉ có thể thực hiện một số trong những thay đổi lớn đó.
Chúng tôi sẽ phải chọn lựa cẩn thận.

## Quy trình

Điều đó đặt ra một câu hỏi quan trọng. Quy trình phát triển
Go là gì?

Trong những ngày đầu của Go, khi chỉ có năm người chúng tôi, chúng tôi làm việc trong
một cặp văn phòng chia sẻ liền kề được ngăn cách bởi một bức tường kính. Rất dễ dàng để kéo mọi người
vào một văn phòng để thảo luận về một vấn đề nào đó và sau đó
quay lại bàn làm việc để triển khai giải pháp. Khi có vấn đề nào đó xuất hiện
trong quá trình triển khai, cũng dễ dàng tập hợp mọi người lại. Văn phòng của Rob
và Robert có một chiếc ghế sofa nhỏ và một bảng trắng, vì vậy thường
một trong chúng tôi bước vào và bắt đầu viết ví dụ lên bảng. Thường thì
khi ví dụ đã lên bảng, mọi người khác đã đạt được điểm dừng tốt trong công việc của mình
và sẵn sàng ngồi xuống thảo luận.
Sự tự nhiên đó rõ ràng không thể mở rộng cho cộng đồng Go toàn cầu ngày nay.

Một phần công việc kể từ khi Go được phát hành mã nguồn mở là chuyển
quy trình không chính thức của chúng tôi vào thế giới chính thức hơn của danh sách thư và
hệ thống theo dõi vấn đề với nửa triệu người dùng, nhưng tôi không nghĩ chúng tôi đã từng
mô tả rõ ràng quy trình tổng thể của mình. Có thể chúng tôi chưa bao giờ
nghĩ về nó một cách có ý thức. Nhìn lại, tuy nhiên, tôi nghĩ đây là
phác thảo cơ bản về công việc của chúng tôi với Go, quy trình chúng tôi đã tuân theo
kể từ khi nguyên mẫu đầu tiên chạy được.

<div style="margin-left: 2em;">
{{image "toward-go2/process.png" 410}}
</div>

Bước 1 là sử dụng Go, tích lũy kinh nghiệm với nó.

Bước 2 là xác định một vấn đề với Go có thể cần giải quyết và
trình bày rõ ràng nó, giải thích cho người khác, viết nó ra.

Bước 3 là đề xuất giải pháp cho vấn đề đó, thảo luận với
người khác và sửa đổi giải pháp dựa trên cuộc thảo luận đó.

Bước 4 là triển khai giải pháp, đánh giá nó và tinh chỉnh nó dựa trên
đánh giá đó.

Cuối cùng, bước 5 là phát hành giải pháp, thêm nó vào ngôn ngữ, hoặc
thư viện, hoặc tập hợp các công cụ mà mọi người sử dụng hàng ngày.

Không nhất thiết một người phải thực hiện tất cả các bước này cho một
thay đổi cụ thể. Trên thực tế, thường nhiều người cộng tác ở bất kỳ bước nào,
và nhiều giải pháp có thể được đề xuất cho một vấn đề duy nhất. Ngoài ra, ở bất kỳ
điểm nào chúng tôi có thể nhận ra rằng chúng tôi không muốn tiến xa hơn với một ý tưởng cụ thể
và quay lại bước trước đó.

Mặc dù tôi không tin rằng chúng tôi đã bao giờ nói về quy trình này như một
tổng thể, chúng tôi đã giải thích các phần của nó. Năm 2012, khi chúng tôi phát hành Go 1
và nói rằng bây giờ là lúc sử dụng Go và ngừng thay đổi nó, chúng tôi đang
giải thích bước 1. Năm 2015, khi chúng tôi giới thiệu quy trình đề xuất thay đổi Go,
chúng tôi đang giải thích các bước 3, 4 và 5. Nhưng chúng tôi chưa bao giờ
giải thích bước 2 một cách chi tiết, vì vậy tôi muốn làm điều đó bây giờ.

(Để biết thêm về sự phát triển của Go 1 và sự chuyển đổi khỏi
các thay đổi ngôn ngữ, xem bài nói OSCON 2012 của Rob Pike và Andrew Gerrand
"[The Path to Go 1](/blog/the-path-to-go-1)."
Để biết thêm về quy trình đề xuất, xem
bài nói GopherCon 2015 của Andrew Gerrand "[How Go was Made](https://www.youtube.com/watch?v=0ht89TxZZnk)" và
[tài liệu về quy trình đề xuất](/s/proposal).)

## Giải thích các vấn đề

<div style="margin-left: 2em;">
{{image "toward-go2/process2.png" 410}}
</div>

Có hai phần để giải thích một vấn đề. Phần đầu tiên, phần dễ hơn,
là trình bày chính xác vấn đề là gì. Chúng ta là những nhà phát triển
khá giỏi ở điều này. Xét cho cùng, mọi bài kiểm tra chúng ta viết đều là một phát biểu
về vấn đề cần giải quyết, bằng ngôn ngữ chính xác đến mức ngay cả máy tính
cũng có thể hiểu. Phần thứ hai, phần khó hơn, là mô tả tầm quan trọng
của vấn đề đủ rõ để mọi người hiểu tại sao chúng ta nên dành thời gian giải quyết nó
và duy trì một giải pháp. Trái ngược với việc trình bày vấn đề một cách chính xác,
chúng ta không cần mô tả tầm quan trọng của vấn đề thường xuyên,
và chúng ta không giỏi điều đó bằng. Máy tính chưa bao giờ hỏi chúng ta
"tại sao test case này quan trọng? Bạn có chắc đây là vấn đề bạn cần giải quyết không?
Giải quyết vấn đề này có phải là điều quan trọng nhất bạn có thể làm không?" Có thể
họ sẽ hỏi vào một ngày nào đó, nhưng không phải hôm nay.

Hãy xem xét một ví dụ cũ từ năm 2011. Đây là những gì tôi đã viết về
việc đổi tên os.Error thành error.Value trong khi chúng tôi đang lên kế hoạch Go 1.

<div style="margin-left: 2em;">
{{image "toward-go2/error.png" 495}}
</div>

Nó bắt đầu với một phát biểu chính xác, một dòng về vấn đề: trong các
thư viện cấp rất thấp mọi thứ đều import "os" vì os.Error. Sau đó có
năm dòng, mà tôi đã gạch chân ở đây, dành để mô tả tầm quan trọng của vấn đề:
các package mà "os" sử dụng không thể tự trình bày lỗi trong API của chúng,
và các package khác phụ thuộc vào "os" vì những lý do không liên quan đến dịch vụ hệ điều hành.

Năm dòng này có thuyết phục _bạn_ rằng vấn đề này quan trọng không?
Điều đó phụ thuộc vào mức độ bạn có thể điền vào ngữ cảnh tôi đã bỏ qua:
được hiểu đòi hỏi phải dự đoán những gì người khác cần biết. Đối với
khán giả của tôi vào thời điểm đó, mười người khác trong nhóm Go tại Google
đang đọc tài liệu đó, năm mươi từ đó là đủ. Để
trình bày vấn đề tương tự cho khán giả tại GothamGo vào mùa thu năm ngoái,
một khán giả với nền tảng và lĩnh vực chuyên môn đa dạng hơn nhiều, tôi
cần cung cấp nhiều ngữ cảnh hơn, và tôi đã sử dụng khoảng hai trăm từ,
cùng với các ví dụ mã thực tế và sơ đồ. Đây là thực tế của cộng đồng Go
toàn cầu ngày nay: mô tả tầm quan trọng của bất kỳ vấn đề nào
đòi hỏi phải thêm ngữ cảnh, đặc biệt được minh họa bằng các ví dụ cụ thể,
mà bạn sẽ bỏ qua khi nói chuyện với đồng nghiệp.

Thuyết phục người khác rằng một vấn đề quan trọng là một bước thiết yếu.
Khi một vấn đề có vẻ không đáng kể, hầu hết mọi giải pháp sẽ có vẻ
quá tốn kém. Nhưng đối với một vấn đề quan trọng, thường có nhiều
giải pháp với chi phí hợp lý. Khi chúng ta không đồng ý về việc có nên
áp dụng một giải pháp cụ thể hay không, chúng ta thường thực sự đang bất đồng về
tầm quan trọng của vấn đề đang được giải quyết. Điều này rất quan trọng đến mức tôi
muốn xem xét hai ví dụ gần đây thể hiện điều này rõ ràng, ít nhất là nhìn lại.

### Ví dụ: giây nhuận

Ví dụ đầu tiên của tôi là về thời gian.

Giả sử bạn muốn đo thời gian một sự kiện kéo dài bao lâu. Bạn ghi thời gian
bắt đầu, chạy sự kiện, ghi thời gian kết thúc, sau đó trừ đi
thời gian bắt đầu khỏi thời gian kết thúc. Nếu sự kiện mất mười mili giây,
phép trừ cho ra kết quả là mười mili giây, có thể cộng hoặc trừ một sai số đo lường nhỏ.

	start := time.Now()       // 3:04:05.000
	event()
	end := time.Now()         // 3:04:05.010

	elapsed := end.Sub(start) // 10 ms

Quy trình rõ ràng này có thể thất bại trong một [giây nhuận](https://en.wikipedia.org/wiki/Leap_second).
Khi đồng hồ của chúng ta không đồng bộ hoàn toàn với vòng quay hàng ngày của Trái Đất,
một giây nhuận, chính thức là 11:59 tối và 60 giây, được chèn vào ngay trước
nửa đêm. Không giống như năm nhuận, giây nhuận không theo một
mô hình có thể dự đoán được, điều này làm cho chúng khó phù hợp với các chương trình và API.
Thay vì cố gắng biểu diễn phút 61 giây đôi khi xảy ra, các hệ điều hành
thường triển khai giây nhuận bằng cách quay đồng hồ trở lại
một giây ngay trước những gì đáng lẽ là nửa đêm, để 11:59 tối
và 59 giây xảy ra hai lần. Việc đặt lại đồng hồ này làm thời gian có vẻ
đi ngược lại, do đó sự kiện mười mili giây của chúng ta có thể được tính là
mất âm 990 mili giây.

	start := time.Now()       // 11:59:59.995
	event()
	end := time.Now()         // 11:59:59.005 (really 11:59:60.005)

	elapsed := end.Sub(start) // –990 ms

Vì đồng hồ giờ trong ngày không chính xác để tính thời gian các sự kiện qua
các lần đặt lại đồng hồ như thế này, các hệ điều hành bây giờ cung cấp một đồng hồ thứ hai,
đồng hồ đơn điệu, không có nghĩa tuyệt đối nhưng đếm giây
và không bao giờ được đặt lại.

Ngoại trừ trong các lần đặt lại đồng hồ lẻ, đồng hồ đơn điệu không tốt hơn
đồng hồ giờ trong ngày, và đồng hồ giờ trong ngày có lợi thêm là
hữu ích để cho biết thời gian, vì vậy để đơn giản API thời gian của Go 1
chỉ hiển thị đồng hồ giờ trong ngày.

Vào tháng 10 năm 2015, một [báo cáo lỗi](/issue/12914) lưu ý rằng các chương trình Go không thể tính thời gian
sự kiện một cách chính xác qua các lần đặt lại đồng hồ, đặc biệt là giây nhuận điển hình.
Cách sửa được đề xuất cũng là tiêu đề vấn đề ban đầu: "thêm API mới để truy cập nguồn
đồng hồ đơn điệu." Tôi lập luận rằng vấn đề này không
đủ quan trọng để biện minh cho API mới. Vài tháng trước đó, cho
giây nhuận giữa năm 2015, Akamai, Amazon và Google đã làm chậm đồng hồ của họ
một chút trong cả ngày, hấp thụ giây thêm
mà không quay đồng hồ ngược lại. Có vẻ như cuối cùng
việc áp dụng rộng rãi cách tiếp cận "[điều chỉnh giây nhuận](https://developers.google.com/time/smear)" này sẽ loại bỏ
các lần đặt lại đồng hồ giây nhuận như một vấn đề trên các hệ thống production. Ngược lại,
việc thêm API mới vào Go sẽ thêm các vấn đề mới: chúng tôi sẽ phải
giải thích hai loại đồng hồ, hướng dẫn người dùng khi nào nên sử dụng
mỗi loại, và chuyển đổi nhiều dòng mã hiện có, tất cả cho một vấn đề
hiếm khi xảy ra và có thể hợp lý sẽ tự biến mất.

Chúng tôi đã làm những gì chúng tôi luôn làm khi có một vấn đề mà không có giải pháp rõ ràng:
chúng tôi đã chờ đợi. Chờ đợi cho chúng tôi thêm thời gian để tích lũy kinh nghiệm và
hiểu biết về vấn đề và cũng thêm thời gian để tìm một giải pháp tốt.
Trong trường hợp này, việc chờ đợi đã thêm vào sự hiểu biết của chúng tôi về
tầm quan trọng của vấn đề, dưới dạng một sự cố may mắn là nhỏ tại
[Cloudflare](https://www.theregister.co.uk/2017/01/04/cloudflare_trips_over_leap_second/).
Mã Go của họ đã tính thời gian các yêu cầu DNS trong giây nhuận cuối năm 2016
khoảng âm 990 mili giây, điều này đã gây ra các panic đồng thời trên các máy chủ của họ,
phá vỡ 0,2% truy vấn DNS ở mức cao điểm.

Cloudflare là chính xác loại hệ thống đám mây mà Go được thiết kế cho,
và họ đã có sự cố production dựa trên Go không thể tính thời gian
sự kiện một cách chính xác. Sau đó, và đây là điểm mấu chốt, Cloudflare đã báo cáo
trải nghiệm của họ trong một bài đăng blog của John Graham-Cumming có tiêu đề
"[How and why the leap second affected Cloudflare DNS](https://blog.cloudflare.com/how-and-why-the-leap-second-affected-cloudflare-dns/)." Bằng cách chia sẻ
chi tiết cụ thể về trải nghiệm của họ với Go trong production, John và Cloudflare đã giúp chúng tôi
hiểu rằng vấn đề về tính chính xác thời gian qua các lần đặt lại đồng hồ giây nhuận
quá quan trọng để không được sửa. Hai tháng sau
khi bài viết đó được xuất bản, chúng tôi đã thiết kế và triển khai một giải pháp
sẽ [ra mắt trong Go 1.9](https://beta.golang.org/doc/go1.9#monotonic-time)
(và thực tế chúng tôi đã làm với [không có API mới](/design/12914-monotonic)).

### Ví dụ: khai báo alias

Ví dụ thứ hai của tôi là hỗ trợ khai báo alias trong Go.

Trong vài năm qua, Google đã thành lập một nhóm tập trung vào
các thay đổi mã nguồn quy mô lớn, có nghĩa là di chuyển API và sửa lỗi được áp dụng trên toàn bộ
[codebase của hàng triệu tệp nguồn và hàng tỷ dòng mã](http://cacm.acm.org/magazines/2016/7/204032-why-google-stores-billions-of-lines-of-code-in-a-single-repository/pdf)
được viết bằng C++, Go, Java, Python và các ngôn ngữ khác. Một
điều tôi đã học được từ công việc của nhóm đó là tầm quan trọng, khi
thay đổi API từ việc sử dụng một tên sang tên khác, của việc có thể
cập nhật mã khách hàng theo nhiều bước, không phải tất cả cùng một lúc. Để làm điều này,
phải có thể viết một khai báo chuyển tiếp việc sử dụng tên cũ
sang tên mới. C++ có #define, typedef và khai báo using
để cho phép việc chuyển tiếp này, nhưng Go không có gì. Tất nhiên, một trong những
mục tiêu của Go là mở rộng tốt cho các codebase lớn, và khi lượng mã Go
tại Google tăng lên, rõ ràng là chúng tôi cần một loại cơ chế chuyển tiếp nào đó
và các dự án và công ty khác cũng sẽ
gặp phải vấn đề này khi codebase Go của họ phát triển.

Vào tháng 3 năm 2016, tôi bắt đầu nói chuyện với Robert Griesemer và Rob Pike
về cách Go có thể xử lý việc cập nhật codebase dần dần, và chúng tôi đã đi đến
khai báo alias, đây chính xác là cơ chế chuyển tiếp cần thiết.
Vào thời điểm này, tôi cảm thấy rất tốt về cách Go đang phát triển. Chúng tôi đã
nói về alias từ những ngày đầu của Go, trên thực tế, bản nháp đặc tả đầu tiên có
[một ví dụ sử dụng khai báo alias](https://go.googlesource.com/go/+/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec#1182), nhưng mỗi khi chúng tôi
thảo luận về alias, và sau đó là type alias, chúng tôi không có trường hợp sử dụng rõ ràng
cho chúng, vì vậy chúng tôi đã bỏ qua. Bây giờ chúng tôi đề xuất thêm alias
không phải vì chúng là một khái niệm thanh lịch mà vì chúng giải quyết
một vấn đề thực tiễn quan trọng với Go đáp ứng mục tiêu của nó là phát triển phần mềm có thể mở rộng.
Tôi hy vọng điều này sẽ là mô hình cho các thay đổi trong tương lai của Go.

Vào cuối mùa xuân, Robert và Rob đã viết [một đề xuất](/design/16339-alias-decls),
và Robert đã trình bày nó trong một [bài nói chớp nhoáng GopherCon 2016](https://www.youtube.com/watch?v=t-w6MyI2qlU). Vài tháng tiếp theo
không diễn ra suôn sẻ, và chúng chắc chắn không phải là mô hình cho các thay đổi trong tương lai của Go.
Một trong những bài học chúng tôi đã học được là tầm quan trọng
của việc mô tả tầm quan trọng của vấn đề.

Vài phút trước, tôi đã giải thích vấn đề cho bạn, đưa ra một số nền tảng
về cách nó có thể phát sinh và tại sao, nhưng không có ví dụ cụ thể nào
có thể giúp bạn đánh giá liệu vấn đề có thể ảnh hưởng đến bạn vào một lúc nào đó không.
Đề xuất của mùa hè năm ngoái và bài nói chớp nhoáng đã đưa ra một
ví dụ trừu tượng, liên quan đến các package C, L, L1 và C1 đến Cn, nhưng không có
ví dụ cụ thể nào mà các nhà phát triển có thể liên hệ. Kết quả là, hầu hết
phản hồi từ cộng đồng dựa trên ý tưởng rằng alias
chỉ giải quyết vấn đề cho Google, không phải cho mọi người khác.

Cũng giống như chúng tôi tại Google lúc đầu không hiểu tầm quan trọng của
việc xử lý đúng các lần đặt lại đồng hồ giây nhuận, chúng tôi đã không
truyền đạt hiệu quả cho cộng đồng Go rộng lớn hơn tầm quan trọng của việc xử lý
di chuyển và sửa chữa mã dần dần trong các thay đổi quy mô lớn.

Vào mùa thu chúng tôi bắt đầu lại. Tôi đã có một [bài nói](https://www.youtube.com/watch?v=h6Cw9iCDVcU) và viết
[một bài viết trình bày vấn đề](/talks/2016/refactor.article)
sử dụng nhiều ví dụ cụ thể được lấy từ
các codebase mã nguồn mở, cho thấy vấn đề này phát sinh ở khắp mọi nơi, không chỉ
bên trong Google. Bây giờ nhiều người hiểu vấn đề hơn và
có thể thấy tầm quan trọng của nó, chúng tôi đã có một [cuộc thảo luận hiệu quả](/issue/18130) về loại
giải pháp nào sẽ tốt nhất. Kết quả là [type alias](/design/18130-type-alias) sẽ
[được đưa vào Go 1.9](https://beta.golang.org/doc/go1.9#language) và sẽ giúp Go mở rộng cho các codebase ngày càng lớn hơn.

### Báo cáo kinh nghiệm

Bài học ở đây là khó nhưng cần thiết để mô tả tầm quan trọng
của vấn đề theo cách mà người làm việc trong một môi trường khác có thể hiểu.
Để thảo luận các thay đổi lớn trong Go như một cộng đồng, chúng tôi sẽ cần chú ý đặc biệt
đến việc mô tả tầm quan trọng của bất kỳ vấn đề nào chúng tôi muốn giải quyết.
Cách rõ ràng nhất để làm điều đó là bằng cách cho thấy cách vấn đề ảnh hưởng đến
các chương trình thực tế và hệ thống production thực tế, như trong
[bài đăng blog của Cloudflare](https://blog.cloudflare.com/how-and-why-the-leap-second-affected-cloudflare-dns/) và trong
[bài viết tái cấu trúc của tôi](/talks/2016/refactor.article).

Các báo cáo kinh nghiệm như thế này biến một vấn đề trừu tượng thành vấn đề cụ thể
và giúp chúng ta hiểu tầm quan trọng của nó. Chúng cũng đóng vai trò là các bài kiểm tra:
bất kỳ giải pháp được đề xuất nào cũng có thể được đánh giá bằng cách xem xét ảnh hưởng của nó
đối với các vấn đề thực tế mà các báo cáo mô tả.

Ví dụ, tôi đã xem xét generics gần đây, nhưng tôi không có trong đầu
một hình ảnh rõ ràng về các vấn đề chi tiết, cụ thể mà người dùng Go
cần generics để giải quyết. Kết quả là, tôi không thể trả lời một câu hỏi thiết kế
như có nên hỗ trợ các phương thức generic không, tức là các
phương thức được tham số hóa riêng biệt so với receiver. Nếu chúng tôi có
một tập hợp lớn các trường hợp sử dụng thực tế, chúng tôi có thể bắt đầu trả lời
câu hỏi như thế này bằng cách xem xét các trường hợp quan trọng.

Như một ví dụ khác, tôi đã thấy các đề xuất để mở rộng interface lỗi
theo nhiều cách khác nhau, nhưng tôi chưa thấy bất kỳ báo cáo kinh nghiệm nào cho thấy cách
các chương trình Go lớn cố gắng hiểu và xử lý lỗi như thế nào, chứ đừng nói đến
việc cho thấy cách interface lỗi hiện tại cản trở những nỗ lực đó.
Các báo cáo này sẽ giúp tất cả chúng ta hiểu rõ hơn các chi tiết và
tầm quan trọng của vấn đề, điều mà chúng ta phải làm trước khi giải quyết nó.

Tôi có thể tiếp tục. Mọi thay đổi lớn tiềm năng đối với Go nên được thúc đẩy
bởi một hoặc nhiều báo cáo kinh nghiệm ghi lại cách mọi người sử dụng Go ngày nay
và tại sao điều đó không hoạt động đủ tốt. Đối với các thay đổi lớn rõ ràng
chúng tôi có thể xem xét cho Go, tôi không biết nhiều báo cáo như vậy,
đặc biệt là không phải các báo cáo được minh họa bằng ví dụ thực tế.

Các báo cáo này là nguyên liệu thô cho quy trình đề xuất Go 2, và
chúng tôi cần tất cả các bạn viết chúng, để giúp chúng tôi hiểu trải nghiệm của bạn
với Go. Có nửa triệu bạn, làm việc trong
một loạt các môi trường rộng lớn, và không có nhiều chúng tôi.
Viết một bài đăng trên blog của riêng bạn,
hoặc viết một bài đăng [Medium](https://www.medium.com/),
hoặc viết một [GitHub Gist](https://gist.github.com/) (thêm phần mở rộng tệp `.md` cho Markdown),
hoặc viết một [tài liệu Google](https://docs.google.com/),
hoặc sử dụng bất kỳ cơ chế xuất bản nào bạn thích.
Sau khi đã đăng, hãy thêm bài đăng vào trang wiki mới của chúng tôi,
[golang.org/wiki/ExperienceReports](/wiki/ExperienceReports).

## Giải pháp

<div style="margin-left: 2em;">
{{image "toward-go2/process34.png" 410}}
</div>

Bây giờ chúng ta đã biết cách chúng ta sẽ xác định và giải thích các vấn đề cần
giải quyết, tôi muốn lưu ý ngắn gọn rằng không phải tất cả các vấn đề đều được
giải quyết tốt nhất bằng cách thay đổi ngôn ngữ, và điều đó là ổn.

Một vấn đề chúng tôi có thể muốn giải quyết là máy tính thường có thể tính toán
các kết quả bổ sung trong các phép toán số học cơ bản, nhưng Go không
cung cấp quyền truy cập trực tiếp vào các kết quả đó. Năm 2013, Robert đề xuất rằng
chúng tôi có thể mở rộng ý tưởng của các biểu thức hai kết quả ("comma-ok") sang
số học cơ bản. Ví dụ, nếu x và y là, giả sử, các giá trị uint32,
`lo, hi = x * y`
sẽ trả về không chỉ 32 bit thấp thông thường mà còn cả 32 bit cao
của tích. Vấn đề này không có vẻ đặc biệt quan trọng, vì vậy
chúng tôi [ghi lại giải pháp tiềm năng](/issue/6815) nhưng không triển khai nó. Chúng tôi đã chờ đợi.

Gần đây hơn, chúng tôi đã thiết kế cho Go 1.9 một [package math/bits](https://beta.golang.org/doc/go1.9#math-bits) chứa
nhiều hàm thao tác bit khác nhau:

	package bits // import "math/bits"

	func LeadingZeros32(x uint32) int
	func Len32(x uint32) int
	func OnesCount32(x uint32) int
	func Reverse32(x uint32) uint32
	func ReverseBytes32(x uint32) uint32
	func RotateLeft32(x uint32, k int) uint32
	func TrailingZeros32(x uint32) int
	...

Package này có các triển khai Go tốt cho mỗi hàm, nhưng các trình biên dịch cũng thay thế
bằng các lệnh phần cứng đặc biệt khi có sẵn. Dựa trên kinh nghiệm này
với math/bits, cả Robert và tôi bây giờ đều tin rằng việc làm cho
các kết quả số học bổ sung có sẵn bằng cách thay đổi ngôn ngữ là
không khôn ngoan, và thay vào đó chúng ta nên định nghĩa các hàm phù hợp trong một
package như math/bits. Đây giải pháp tốt nhất là thay đổi thư viện,
không phải thay đổi ngôn ngữ.

Một vấn đề khác chúng tôi có thể muốn giải quyết, sau Go 1.0, là
thực tế rằng goroutine và bộ nhớ chia sẻ làm cho quá dễ dàng để
đưa race condition vào các chương trình Go, gây ra sự cố và hành vi sai
trong production. Giải pháp dựa trên ngôn ngữ sẽ là tìm cách
không cho phép data race, để làm cho việc viết hoặc ít nhất là biên dịch
một chương trình có data race là không thể. Cách phù hợp với ngôn ngữ như Go
vẫn là một câu hỏi chưa có lời giải trong thế giới ngôn ngữ lập trình. Thay vào đó
chúng tôi đã thêm một công cụ vào bản phân phối chính và làm cho nó dễ sử dụng:
công cụ đó, [bộ phát hiện race condition](/blog/race-detector), đã trở thành
một phần không thể thiếu của trải nghiệm Go. Đây giải pháp tốt nhất là thay đổi runtime và công cụ,
không phải thay đổi ngôn ngữ.

Sẽ có các thay đổi ngôn ngữ cũng vậy, tất nhiên, nhưng không phải tất cả
các vấn đề đều được giải quyết tốt nhất trong ngôn ngữ.

## Phát hành Go 2

<div style="margin-left: 2em;">
{{image "toward-go2/process5.png" 410}}
</div>

Cuối cùng, chúng ta sẽ phát hành và phân phối Go 2 như thế nào?

Tôi nghĩ kế hoạch tốt nhất sẽ là phát hành [các phần tương thích ngược](/doc/go1compat)
của Go 2 từng bước, tính năng theo tính năng, như một phần của chuỗi phát hành Go 1.
Điều này có một số thuộc tính quan trọng. Thứ nhất, nó giữ các bản phát hành Go 1
theo [lịch trình thông thường](/wiki/Go-Release-Cycle), để tiếp tục sửa lỗi kịp thời và
các cải tiến mà người dùng hiện đang phụ thuộc. Thứ hai, nó tránh chia rẽ
nỗ lực phát triển giữa Go 1 và Go 2. Thứ ba, nó tránh phân kỳ
giữa Go 1 và Go 2, để dễ dàng di chuyển cho mọi người. Thứ tư,
nó cho phép chúng tôi tập trung và phát hành một thay đổi mỗi lần, điều này
nên giúp duy trì chất lượng. Thứ năm, nó sẽ khuyến khích chúng tôi thiết kế
các tính năng tương thích ngược.

Chúng tôi sẽ cần thời gian để thảo luận và lên kế hoạch trước khi bất kỳ thay đổi nào bắt đầu
đến trong các bản phát hành Go 1, nhưng có vẻ hợp lý với tôi rằng chúng tôi có thể bắt đầu thấy
các thay đổi nhỏ khoảng một năm nữa, cho Go 1.12 hoặc tương tự. Điều đó cũng
cho chúng tôi thời gian để đưa hỗ trợ quản lý package vào trước.

Khi tất cả công việc tương thích ngược đã hoàn thành, giả sử trong Go 1.20, thì
chúng tôi có thể thực hiện các thay đổi không tương thích ngược trong Go 2.0. Nếu
không có thay đổi không tương thích ngược, có lẽ chúng ta chỉ
tuyên bố rằng Go 1.20 _là_ Go 2.0. Dù sao đi nữa, tại điểm đó chúng tôi sẽ
chuyển từ làm việc trên chuỗi phát hành Go 1.X sang làm việc trên
chuỗi Go 2.X, có thể với cửa sổ hỗ trợ mở rộng cho
bản phát hành Go 1.X cuối cùng.

Tất cả điều này có phần suy đoán, và các số phát hành cụ thể
tôi vừa đề cập là các ước tính tạm thời,
nhưng tôi muốn làm rõ rằng chúng tôi không
từ bỏ Go 1, và trên thực tế chúng tôi sẽ đưa Go 1 theo
trong phạm vi tối đa có thể.

## Cần Sự Giúp Đỡ

**Chúng tôi cần sự giúp đỡ của bạn.**

Cuộc trò chuyện cho Go 2 bắt đầu hôm nay, và đó là cuộc trò chuyện sẽ xảy ra
một cách công khai, trong các diễn đàn công cộng như danh sách thư và
hệ thống theo dõi vấn đề. Hãy giúp chúng tôi ở mỗi bước trên đường.

Hôm nay, điều chúng tôi cần nhất là các báo cáo kinh nghiệm. Hãy cho chúng tôi biết cách Go
đang hoạt động cho bạn, và quan trọng hơn là không hoạt động cho bạn. Viết một
bài đăng blog, bao gồm các ví dụ thực tế, chi tiết cụ thể và
kinh nghiệm thực tế. Và liên kết nó trên [trang wiki](/wiki/ExperienceReports) của chúng tôi.
Đó là cách chúng ta sẽ bắt đầu nói về những gì chúng ta, cộng đồng Go,
có thể muốn thay đổi về Go.

Cảm ơn.
