---
title: Kết quả Khảo sát Go 2018
date: 2019-03-28
by:
- Todd Kulesza, Steve Francia
tags:
- survey
- community
summary: Những gì chúng tôi thu được từ Khảo sát Người dùng Go tháng 12 năm 2018.
template: true
---

## Cảm ơn

<style>
    p.note {
        font-size: 0.80em;
        font-family: "Helvetica Neue", Arial, sans-serif; /* Helvetica on Mac aka sans-serif has broken U+2007 */
    }
</style>

Bài đăng này tóm tắt kết quả khảo sát người dùng năm 2018 của chúng tôi và so sánh
với kết quả các khảo sát trước của chúng tôi từ [năm 2016](/blog/survey2016-results)
và [năm 2017](/blog/survey2017-results).

Năm nay chúng tôi có 5.883 người tham gia khảo sát đến từ 103 quốc gia khác nhau.
Chúng tôi biết ơn tất cả mọi người đã cung cấp phản hồi qua khảo sát này
để giúp định hình tương lai của Go. Cảm ơn!

## Tóm tắt kết quả

  - Lần đầu tiên, **một nửa số người tham gia khảo sát hiện đang dùng Go như một phần trong thói quen hàng ngày**.
    Năm nay cũng chứng kiến sự gia tăng đáng kể về số người tham gia
    phát triển Go như một phần công việc của họ và sử dụng Go ngoài trách nhiệm công việc.
  - **Những ứng dụng phổ biến nhất của Go vẫn là dịch vụ API/RPC và công cụ CLI**.
    Các tác vụ tự động, dù không phổ biến như công cụ CLI và dịch vụ API,
    là lĩnh vực phát triển nhanh cho Go.
  - **Phát triển web vẫn là lĩnh vực phổ biến nhất** mà người tham gia khảo sát làm việc,
    nhưng **DevOps cho thấy mức tăng trưởng cao nhất theo năm** và hiện là lĩnh vực phổ biến thứ hai.
  - Đại đa số người tham gia khảo sát cho biết **Go là ngôn ngữ lập trình ưa thích nhất của họ**,
    dù nhìn chung cảm thấy ít thành thạo hơn so với ít nhất một ngôn ngữ khác.
  - **VS Code và GoLand đang tăng vọt về độ phổ biến** và hiện là các trình soạn thảo phổ biến nhất trong số người tham gia khảo sát.
  - Nổi bật lên tính di động của Go,
    **nhiều nhà phát triển Go sử dụng nhiều hệ điều hành chính** để phát triển.
    Linux và macOS đặc biệt phổ biến,
    với đại đa số người tham gia khảo sát sử dụng một hoặc cả hai hệ điều hành này để viết mã Go.
  - Người tham gia khảo sát có vẻ đang **dịch chuyển khỏi triển khai Go tại chỗ**
    và chuyển sang triển khai đám mây dạng container và serverless.
  - Đa số người tham gia cho biết họ cảm thấy được chào đón trong cộng đồng Go,
    và hầu hết các ý tưởng cải thiện cộng đồng Go tập trung vào **cải thiện trải nghiệm cho người mới**.

Đọc tiếp để xem toàn bộ chi tiết.

## Nền tảng lập trình

Kết quả năm nay cho thấy sự gia tăng đáng kể về số người tham gia khảo sát
được trả lương để viết Go như một phần công việc của họ (68% lên 72%),
tiếp tục xu hướng từng năm đã tăng trưởng kể từ khảo sát đầu tiên của chúng tôi vào năm 2016.
Chúng tôi cũng thấy sự gia tăng về số người tham gia lập trình Go ngoài
công việc (64% lên 70%).
Lần đầu tiên, số người tham gia khảo sát viết Go như một phần trong
thói quen hàng ngày đạt 50% (tăng từ 44% năm 2016).
Những phát hiện này gợi ý rằng các công ty đang tiếp tục đón nhận Go cho phát triển phần mềm
chuyên nghiệp ở một nhịp độ ổn định,
và rằng độ phổ biến chung của Go với các nhà phát triển vẫn mạnh.

{{image "survey2018/fig1.svg" 600}}
{{image "survey2018/fig2.svg" 600}}

Để hiểu rõ hơn nơi các nhà phát triển sử dụng Go,
chúng tôi chia phản hồi thành ba nhóm:
1) những người sử dụng Go cả trong và ngoài công việc,
2) những người dùng Go chuyên nghiệp nhưng không dùng ngoài công việc,
và 3) những người chỉ viết Go ngoài trách nhiệm công việc.
Gần một nửa (46%) người tham gia viết mã Go cả chuyên nghiệp lẫn thời gian tự do (tăng 10 điểm kể từ năm 2017),
trong khi những người tham gia còn lại được phân bổ gần đều giữa chỉ viết Go ở nơi làm việc
hoặc chỉ viết Go bên ngoài công việc.
Tỷ lệ lớn người tham gia vừa dùng Go ở nơi làm việc vừa chọn dùng nó
ngoài công việc cho thấy ngôn ngữ này hấp dẫn các nhà phát triển
không coi kỹ thuật phần mềm chỉ là công việc ban ngày:
họ cũng chọn code ngoài trách nhiệm công việc,
và (như được chứng minh bởi 85% người tham gia nói họ sẽ ưu tiên Go cho dự án tiếp theo,
xem phần _Thái độ đối với Go_ bên dưới) Go là ngôn ngữ hàng đầu họ muốn
sử dụng cho những dự án không liên quan đến công việc này.

{{image "survey2018/fig4.svg" 600}}

Khi được hỏi họ đã dùng Go bao lâu,
câu trả lời của người tham gia có xu hướng tăng lên rõ rệt theo thời gian,
với tỷ lệ phản hồi cao hơn trong các nhóm 2-4 năm và 4+ năm mỗi năm.
Điều này được kỳ vọng đối với một ngôn ngữ lập trình mới hơn,
và chúng tôi vui mừng nhận thấy rằng tỷ lệ người tham gia mới với Go
giảm chậm hơn tỷ lệ người tham gia đã dùng Go hơn 2 năm đang tăng,
vì điều này gợi ý rằng các nhà phát triển không bỏ hệ sinh thái sau
khi ban đầu học ngôn ngữ.

{{image "survey2018/fig5.svg" 600}}

Như những năm trước, Go xếp hạng đầu trong số các ngôn ngữ ưa thích và
ngôn ngữ có chuyên môn của người tham gia.
Đa số người tham gia (69%) cho biết có chuyên môn ở 5 ngôn ngữ khác nhau,
nổi bật lên rằng thái độ của họ đối với Go bị ảnh hưởng bởi kinh nghiệm
với các ngăn xếp lập trình khác.
Các biểu đồ dưới đây được sắp xếp theo số người tham gia xếp hạng mỗi
ngôn ngữ là ưa thích/hiểu nhất (thanh màu xanh đậm nhất),
điều này nổi bật lên ba điểm thú vị:

  - Trong khi khoảng 1/3 người tham gia coi Go là ngôn ngữ họ có chuyên môn nhất,
    gấp đôi số đó coi Go là ngôn ngữ lập trình ưa thích nhất của họ.
    Vì vậy, dù nhiều người tham gia cảm thấy họ chưa thành thạo Go bằng
    một số ngôn ngữ khác,
    họ vẫn thường ưu tiên phát triển bằng Go.
  - Ít người tham gia khảo sát xếp hạng Rust là ngôn ngữ họ có chuyên môn (6,8%),
    nhưng 19% xếp hạng nó là ngôn ngữ ưa thích hàng đầu,
    cho thấy mức độ quan tâm cao đối với Rust trong đối tượng này.
  - Chỉ có ba ngôn ngữ có nhiều người tham gia nói họ ưa thích ngôn ngữ đó
    hơn là nói họ có chuyên môn với nó:
    Rust (tỷ lệ ưa thích:chuyên môn là 2,41:1),
    Kotlin (1,95:1) và Go (1,02:1).
    Mức ưa thích cao hơn chuyên môn gợi ý sự quan tâm nhưng ít kinh nghiệm trực tiếp với một ngôn ngữ,
    trong khi mức ưa thích thấp hơn chuyên môn gợi ý có rào cản khi sử dụng thành thạo.
    Tỷ lệ gần 1,0 gợi ý rằng hầu hết các nhà phát triển đều có thể làm việc hiệu quả
    _và_ thú vị với một ngôn ngữ nhất định.
    Dữ liệu này được chứng minh bởi [khảo sát nhà phát triển 2018 của Stack Overflow](https://insights.stackoverflow.com/survey/2018/#most-loved-dreaded-and-wanted),
    cũng thấy Rust, Kotlin và Go là các ngôn ngữ lập trình được ưa thích nhất.

{{image "survey2018/fig6.svg" 600}}
{{image "survey2018/fig7.svg" 600}}

<p class="note">
    <i>Đọc dữ liệu</i>: Người tham gia có thể xếp hạng 5 ngôn ngữ hàng đầu. Mã màu bắt đầu bằng màu xanh đậm cho hạng nhất và nhạt dần cho mỗi hạng tiếp theo. Các biểu đồ này được sắp xếp theo tỷ lệ người tham gia xếp hạng mỗi ngôn ngữ là lựa chọn hàng đầu của họ.
</p>

## Lĩnh vực phát triển

Người tham gia khảo sát cho biết làm việc ở trung bình ba lĩnh vực khác nhau,
với đại đa số (72%) làm việc trong 2-5 lĩnh vực khác nhau.
Phát triển web phổ biến nhất với 65%,
và nó tăng cường sự thống trị là lĩnh vực chính mà người tham gia khảo sát làm việc
(tăng từ 61% năm ngoái):
phát triển web là lĩnh vực phổ biến nhất cho phát triển Go kể từ năm 2016.
Năm nay DevOps tăng lên đáng kể, từ 36% đến 41% người tham gia,
chiếm vị trí thứ hai từ tay Lập trình Hệ thống.
Chúng tôi không tìm thấy lĩnh vực nào có mức sử dụng thấp hơn năm 2018 so với năm 2017,
gợi ý rằng người tham gia đang áp dụng Go cho nhiều loại dự án hơn,
thay vì chuyển dịch sử dụng từ lĩnh vực này sang lĩnh vực khác.

{{image "survey2018/fig8.svg" 600}}

Kể từ năm 2016, hai ứng dụng hàng đầu của Go là viết dịch vụ API/RPC và
phát triển ứng dụng CLI.
Trong khi mức sử dụng CLI vẫn ổn định ở 63% trong ba năm,
mức sử dụng API/RPC đã tăng từ 60% năm 2016 lên 65% năm 2017 lên 73% hiện nay.
Những lĩnh vực này phát huy thế mạnh cốt lõi của Go và đều là trung tâm của
phát triển phần mềm cloud-native,
vì vậy chúng tôi kỳ vọng chúng sẽ vẫn là hai kịch bản chính cho các nhà phát triển Go trong tương lai.
Tỷ lệ người tham gia viết dịch vụ web trả về HTML trực tiếp đã giảm dần trong khi mức sử dụng API/RPC tăng,
gợi ý một số di chuyển sang mô hình API/RPC cho dịch vụ web.
Một xu hướng theo năm khác gợi ý rằng tự động hóa cũng là lĩnh vực phát triển cho Go,
với 38% người tham gia hiện dùng Go cho các tập lệnh và tác vụ tự động (tăng từ 31% năm 2016).

{{image "survey2018/fig9.svg" 600}}

Để hiểu rõ hơn bối cảnh mà các nhà phát triển đang sử dụng Go,
chúng tôi thêm câu hỏi về việc áp dụng Go trong các ngành khác nhau.
Có lẽ không có gì đáng ngạc nhiên đối với một ngôn ngữ còn tương đối mới,
hơn một nửa người tham gia khảo sát làm việc ở các công ty trong danh mục _Dịch vụ Internet/web_
và _Phần mềm_ (tức là các công ty công nghệ).
Các ngành duy nhất khác có >3% phản hồi là _Tài chính, ngân hàng hoặc bảo hiểm_
và _Truyền thông, quảng cáo, xuất bản hoặc giải trí_.
(Trong biểu đồ dưới đây, chúng tôi đã gộp tất cả các danh mục có tỷ lệ phản hồi
dưới 3% vào danh mục "Khác".) Chúng tôi sẽ tiếp tục theo dõi sự
áp dụng Go trong các ngành để hiểu rõ hơn nhu cầu nhà phát triển bên ngoài
các công ty công nghệ.

{{image "survey2018/fig10.svg" 600}}

## Thái độ đối với Go

Năm nay chúng tôi thêm một câu hỏi hỏi "Bạn có khả năng giới thiệu Go
cho bạn bè hoặc đồng nghiệp không?" để tính [Điểm Nhà Quảng bá Ròng (Net Promoter Score)](https://en.wikipedia.org/wiki/Net_Promoter).
Điểm này cố gắng đo lường có bao nhiêu "người quảng bá" hơn "người phản đối" đối với một sản phẩm
và dao động từ -100 đến 100;
một giá trị dương gợi ý hầu hết mọi người có khả năng giới thiệu dùng một sản phẩm,
trong khi các giá trị âm gợi ý hầu hết mọi người có khả năng khuyên không nên dùng nó.
Điểm năm 2018 của chúng tôi là 61 (68% người quảng bá - 7% người phản đối) và sẽ đóng vai trò là
đường cơ sở để giúp chúng tôi đánh giá tâm lý cộng đồng đối với hệ sinh thái Go theo thời gian.

{{image "survey2018/fig11.svg" 600}}

Ngoài NPS, chúng tôi đặt một số câu hỏi về mức độ hài lòng của nhà phát triển với Go.
Nhìn chung, người tham gia khảo sát cho thấy mức độ hài lòng cao,
nhất quán với những năm trước.
Đại đa số cho biết họ hài lòng với Go (89%),
muốn dùng Go cho dự án tiếp theo của họ (85%),
và cảm thấy nó đang hoạt động tốt cho nhóm của họ (66%),
trong khi đa số cảm thấy Go ít nhất có tầm quan trọng nhất định đối với sự thành công của công ty họ (44%).
Trong khi tất cả các số liệu này đều tăng trong năm 2017,
chúng hầu hết ổn định trong năm nay.
(Cách diễn đạt câu hỏi đầu tiên đã thay đổi trong năm 2018 từ "_Tôi sẽ giới thiệu dùng Go cho người khác_"
thành "_Nhìn chung, tôi hài lòng với Go_",
vì vậy những kết quả đó không thể so sánh trực tiếp.)

{{image "survey2018/fig12.svg" 600}}

Do xu hướng mạnh mẽ về việc ưu tiên Go cho phát triển trong tương lai,
chúng tôi muốn hiểu điều gì ngăn cản các nhà phát triển làm điều đó.
Những điều này hầu như không thay đổi so với năm ngoái:
khoảng 1/2 người tham gia khảo sát làm việc trên các dự án hiện có được viết bằng các ngôn ngữ khác,
và 1/3 làm việc trong một nhóm hoặc dự án ưa thích dùng ngôn ngữ khác.
Thiếu tính năng ngôn ngữ và thư viện là những lý do phổ biến nhất
người tham gia không dùng Go nhiều hơn.
Chúng tôi cũng hỏi về những thách thức lớn nhất mà nhà phát triển gặp phải khi dùng Go;
không giống như hầu hết các câu hỏi khảo sát của chúng tôi, người tham gia có thể nhập bất kỳ thứ gì
họ muốn để trả lời câu hỏi này.
Chúng tôi phân tích kết quả thông qua học máy để xác định các chủ đề phổ biến và
đếm số lượng phản hồi hỗ trợ mỗi chủ đề.
Ba thách thức chính hàng đầu chúng tôi xác định được là:

  - Quản lý gói (ví dụ: "Theo kịp việc vendoring",
    "quản lý dependency/packet [sic]/vendoring chưa thống nhất")
  - Khác biệt so với các ngôn ngữ lập trình quen thuộc hơn (ví dụ:
    "cú pháp gần với ngôn ngữ C nhưng ngữ nghĩa hơi khác khiến tôi tra cứu tài liệu tham khảo
    nhiều hơn mong muốn",
    "đồng nghiệp đến từ nền tảng không phải Go cố dùng Go như phiên bản ngôn ngữ trước của họ nhưng có channel và Goroutine")
  - Thiếu generics (ví dụ: "Thiếu generics khiến khó thuyết phục
    những người chưa thử Go rằng họ sẽ thấy nó hiệu quả.",
    "Khó xây dựng các trừu tượng phong phú hơn (muốn generics)")

{{image "survey2018/fig13.svg" 600}}
{{image "survey2018/fig14.svg" 600}}

Năm nay chúng tôi thêm một số câu hỏi về mức độ hài lòng của nhà phát triển với các khía cạnh khác nhau của Go.
Người tham gia khảo sát rất hài lòng với hiệu suất CPU của ứng dụng Go (46:1,
nghĩa là 46 người tham gia cho biết họ hài lòng cho mỗi 1 người tham gia
cho biết họ không hài lòng),
tốc độ build (37:1) và mức sử dụng bộ nhớ ứng dụng (32:1).
Tuy nhiên, phản hồi về khả năng debug ứng dụng (3,2:1) và kích thước binary (6,4:1)
gợi ý còn nhiều điều có thể cải thiện.

Sự không hài lòng về kích thước binary chủ yếu đến từ các nhà phát triển xây dựng CLI,
chỉ 30% trong số đó hài lòng với kích thước các binary được tạo ra bởi Go.
Tuy nhiên, đối với tất cả các loại ứng dụng khác,
mức độ hài lòng của nhà phát triển là > 50%, và kích thước binary nhất quán được xếp hạng
ở cuối danh sách các yếu tố quan trọng.

Khả năng debug, ngược lại, nổi bật khi chúng tôi xem xét cách người tham gia xếp hạng
tầm quan trọng của mỗi khía cạnh;
44% người tham gia xếp hạng khả năng debug là khía cạnh quan trọng nhất hoặc thứ hai của họ,
nhưng chỉ 36% hài lòng với trạng thái hiện tại của Go debugging.
Khả năng debug nhất quán được đánh giá quan trọng gần bằng mức sử dụng bộ nhớ
và tốc độ build nhưng với mức độ hài lòng thấp hơn đáng kể,
và mô hình này đúng bất kể loại phần mềm người tham gia đang xây dựng.
Hai bản phát hành Go gần đây nhất, Go 1.11 và 1.12,
đều chứa những cải tiến đáng kể về khả năng debug.
Chúng tôi dự định điều tra sâu hơn cách các nhà phát triển debug các ứng dụng Go trong năm nay,
với mục tiêu cải thiện trải nghiệm debug tổng thể cho các nhà phát triển Go.

{{image "survey2018/fig15.svg" 600}}
{{image "survey2018/fig29.svg" 600}}

## Môi trường phát triển

Chúng tôi hỏi người tham gia hệ điều hành nào họ chủ yếu dùng khi viết mã Go.
Đa số (65%) người tham gia cho biết họ dùng Linux,
50% dùng macOS và 18% dùng Windows, nhất quán với năm ngoái.
Năm nay chúng tôi cũng xem xét có bao nhiêu người tham gia phát triển trên nhiều hệ điều hành so với một hệ điều hành duy nhất.
Linux và macOS vẫn là những lựa chọn dẫn đầu rõ ràng,
với 81% người tham gia phát triển trên một số kết hợp của hai hệ thống này.
Chỉ 3% người tham gia phân bổ thời gian đồng đều giữa cả ba hệ điều hành.
Nhìn chung, 41% người tham gia sử dụng nhiều hệ điều hành để phát triển Go,
nổi bật lên tính đa nền tảng của Go.

{{image "survey2018/fig16.svg" 600}}

Năm ngoái, VS Code đã nhỉnh hơn Vim trở thành trình soạn thảo Go phổ biến nhất trong số người tham gia khảo sát.
Năm nay nó đã mở rộng đáng kể vị trí dẫn đầu để trở thành trình soạn thảo ưa thích
của hơn 1/3 người tham gia khảo sát (tăng từ 27% năm ngoái).
GoLand cũng trải qua tăng trưởng mạnh và hiện là trình soạn thảo ưa thích thứ hai với 22%,
đổi chỗ với Vim (giảm xuống 17%).
Sự tăng vọt phổ biến của VS Code và GoLand có vẻ đến từ
Sublime Text và Atom.
Vim cũng thấy số người tham gia xếp hạng nó là lựa chọn hàng đầu giảm,
nhưng nó vẫn là trình soạn thảo lựa chọn thứ hai phổ biến nhất với 14%.
Thú vị là, chúng tôi không tìm thấy sự khác biệt về mức độ hài lòng mà người tham gia
báo cáo đối với trình soạn thảo của họ.

Chúng tôi cũng hỏi người tham gia điều gì sẽ cải thiện hỗ trợ Go nhất trong trình soạn thảo ưa thích của họ.
Như câu hỏi "thách thức lớn nhất" ở trên,
người tham gia có thể tự viết câu trả lời thay vì chọn từ danh sách
nhiều lựa chọn.
Phân tích chủ đề trên các phản hồi cho thấy _cải thiện hỗ trợ debug_ (ví dụ:
"Debug trực tiếp", "Debug tích hợp",
"Debug tốt hơn nữa") là yêu cầu phổ biến nhất,
tiếp theo là _cải thiện hoàn thành mã_ (ví dụ:
"hiệu suất và chất lượng tự động hoàn thành", "tự động hoàn thành thông minh hơn").
Các yêu cầu khác bao gồm tích hợp tốt hơn với toolchain CLI của Go,
hỗ trợ tốt hơn cho module/gói và cải thiện hiệu suất chung.

{{image "survey2018/fig17.svg" 600}}
{{image "survey2018/fig18.svg" 600}}

Năm nay chúng tôi cũng thêm câu hỏi hỏi kiến trúc triển khai nào
quan trọng nhất đối với các nhà phát triển Go.
Không có gì đáng ngạc nhiên, người tham gia khảo sát xem x86/x86-64 là nền tảng triển khai hàng đầu của họ (76% người tham gia liệt kê nó là kiến trúc triển khai quan trọng nhất,
và 84% có nó trong top 3).
Tuy nhiên, thứ tự của các kiến trúc lựa chọn thứ hai và thứ ba
rất thú vị:
có sự quan tâm đáng kể đến ARM64 (45%),
WebAssembly (30%) và ARM (22%), nhưng rất ít quan tâm đến các nền tảng khác.

{{image "survey2018/fig19.svg" 600}}

## Triển khai và dịch vụ

Năm 2018, chúng tôi thấy sự tiếp nối xu hướng từ on-prem sang hosting đám mây
cho cả triển khai Go và không phải Go.
Tỷ lệ người tham gia khảo sát triển khai ứng dụng Go lên máy chủ on-prem giảm từ 43% xuống 32%,
phản ánh mức giảm 46% xuống 36% được báo cáo cho triển khai không phải Go.
Các dịch vụ đám mây có mức tăng trưởng cao nhất theo năm bao gồm AWS
Lambda (4% lên 11% cho Go,
10% lên 15% không phải Go) và Google Kubernetes Engine (8% lên 12% cho Go,
5% lên 10% không phải Go), gợi ý rằng serverless và container đang trở thành
các nền tảng triển khai ngày càng phổ biến.
Tuy nhiên, sự tăng trưởng dịch vụ này có vẻ được thúc đẩy bởi những người tham gia đã
áp dụng dịch vụ đám mây,
vì chúng tôi không tìm thấy mức tăng trưởng có ý nghĩa trong tỷ lệ người tham gia
triển khai lên ít nhất một dịch vụ đám mây năm nay (55% lên 56%).
Chúng tôi cũng thấy sự tăng trưởng ổn định trong triển khai Go lên GCP kể từ năm 2016,
tăng từ 12% lên 19% người tham gia.

{{image "survey2018/fig20.svg" 600}}

Có lẽ tương quan với sự giảm trong triển khai on-prem,
năm nay chúng tôi thấy lưu trữ đám mây trở thành dịch vụ được sử dụng nhiều thứ hai bởi người tham gia khảo sát,
tăng từ 32% lên 44%.
Các dịch vụ xác thực và liên kết cũng thấy mức tăng đáng kể (26% lên 33%).
Dịch vụ chính mà người tham gia khảo sát truy cập từ Go vẫn là
cơ sở dữ liệu quan hệ mã nguồn mở,
tăng nhẹ từ 61% lên 65% người tham gia.
Như biểu đồ dưới đây cho thấy, mức sử dụng dịch vụ tăng trên diện rộng.

{{image "survey2018/fig21.svg" 600}}

## Cộng đồng Go

Các nguồn cộng đồng hàng đầu để tìm câu trả lời cho câu hỏi về Go tiếp tục là
Stack Overflow (23% người tham gia đánh dấu đây là nguồn hàng đầu của họ),
các trang web Go (18% cho godoc.org, 14% cho golang.org),
và đọc mã nguồn (8% cho mã nguồn nói chung,
4% cho GitHub cụ thể).
Thứ tự hầu như nhất quán với những năm trước.
Các nguồn tin tức Go chính vẫn là blog Go,
r/golang của Reddit, Twitter và Hacker News.
Tuy nhiên, đây cũng là các phương thức phân phối chính cho khảo sát này,
vì vậy có thể có một số thiên lệch trong kết quả này.
Trong hai biểu đồ dưới đây, chúng tôi đã gộp các nguồn được sử dụng bởi ít hơn 5% người tham gia
vào danh mục "Khác".

{{image "survey2018/fig24.svg" 600}}
{{image "survey2018/fig25.svg" 600}}

Năm nay, 55% người tham gia khảo sát cho biết họ đã hoặc quan tâm đến việc
đóng góp cho cộng đồng Go,
giảm nhẹ từ 59% năm ngoái.
Vì hai lĩnh vực đóng góp phổ biến nhất (thư viện chuẩn
và các công cụ Go chính thức) đòi hỏi tương tác với nhóm Go cốt lõi,
chúng tôi nghi ngờ sự giảm này có thể liên quan đến việc giảm tỷ lệ người tham gia
đồng ý với các nhận định "Tôi cảm thấy thoải mái khi tiếp cận ban lãnh đạo dự án Go với câu hỏi và phản hồi" (30% xuống 25%) và "Tôi tin tưởng vào ban lãnh đạo của Go (54% xuống 46%).

{{image "survey2018/fig26.svg" 600}}
{{image "survey2018/fig27.svg" 600}}

Một khía cạnh quan trọng của cộng đồng là giúp mọi người cảm thấy được chào đón,
đặc biệt là những người từ các nhóm ít được đại diện theo truyền thống.
Để hiểu rõ hơn điều này, chúng tôi đặt câu hỏi tùy chọn về nhận dạng
trong một số nhóm ít được đại diện.
Trong năm 2017, chúng tôi thấy sự tăng trưởng theo năm trên diện rộng.
Đối với năm 2018, chúng tôi thấy tỷ lệ người tham gia tương tự (12%) xác định là thành viên của
nhóm thiểu số,
và điều này đi kèm với sự giảm đáng kể trong tỷ lệ người tham gia
**không** xác định là thành viên của nhóm thiểu số.
Trong năm 2017, với mỗi người xác định là thành viên của nhóm thiểu số,
có 3,5 người xác định không phải thành viên của nhóm thiểu số (tỷ lệ 3,5:1).
Năm 2018, tỷ lệ đó cải thiện xuống 3,08:1. Điều này gợi ý rằng cộng đồng Go
ít nhất đang duy trì tỷ lệ thành viên thiểu số tương tự,
và thậm chí có thể đang tăng lên.

{{image "survey2018/fig28.svg" 600}}

Duy trì một cộng đồng lành mạnh là cực kỳ quan trọng đối với dự án Go,
vì vậy trong ba năm qua, chúng tôi đã đo lường mức độ các nhà phát triển
cảm thấy được chào đón trong cộng đồng Go.
Năm nay chúng tôi thấy sự giảm trong tỷ lệ người tham gia khảo sát đồng ý
với nhận định "Tôi cảm thấy được chào đón trong cộng đồng Go", từ 66% xuống 59%.

Để hiểu rõ hơn về sự giảm này, chúng tôi xem xét kỹ hơn ai cho biết
cảm thấy kém được chào đón hơn.
Trong số các nhóm ít được đại diện theo truyền thống,
ít người hơn cho biết cảm thấy không được chào đón trong năm 2018,
gợi ý rằng việc tiếp cận trong lĩnh vực đó đã có ích.
Thay vào đó, chúng tôi tìm thấy mối quan hệ tuyến tính giữa thời gian ai đó đã dùng Go
và mức độ họ cảm thấy được chào đón:
các nhà phát triển Go mới cảm thấy kém được chào đón đáng kể (ở 50%) so với những nhà phát triển
có 1-2 năm kinh nghiệm (62%),
những người lại cảm thấy kém được chào đón hơn so với những nhà phát triển có vài năm kinh nghiệm (73%).
Cách giải thích dữ liệu này được hỗ trợ bởi các phản hồi cho câu hỏi
"Những thay đổi nào sẽ làm cho cộng đồng Go thân thiện hơn?".
Bình luận của người tham gia có thể được gộp lại thành bốn danh mục:

  - Giảm nhận thức về chủ nghĩa tinh hoa, đặc biệt là với người mới đến Go (ví dụ:
    "ít coi thường hơn", "Ít phòng thủ và kiêu ngạo hơn")
  - Tăng minh bạch ở cấp ban lãnh đạo (ví dụ:
    "Thảo luận về hướng đi và kế hoạch trong tương lai",
    "Ít lãnh đạo từ trên xuống hơn", "Dân chủ hơn")
  - Tăng tài nguyên giới thiệu (ví dụ: "Giới thiệu rõ ràng hơn cho người đóng góp",
    "Những thử thách thú vị để học thực hành tốt nhất")
  - Thêm sự kiện và buổi gặp mặt, tập trung vào phạm vi địa lý rộng hơn (ví dụ:
    "Thêm buổi gặp mặt và sự kiện xã hội", "Sự kiện ở nhiều thành phố hơn")

Phản hồi này rất hữu ích và cung cấp cho chúng tôi các lĩnh vực cụ thể có thể tập trung vào
để cải thiện trải nghiệm làm nhà phát triển Go.
Dù nó không đại diện cho một tỷ lệ lớn trong cơ sở người dùng của chúng tôi,
chúng tôi coi trọng phản hồi này rất nghiêm túc và đang làm việc để cải thiện từng lĩnh vực.

{{image "survey2018/fig22.svg" 600}}
{{image "survey2018/fig23.svg" 600}}

## Kết luận

Chúng tôi hy vọng bạn đã thích xem kết quả khảo sát nhà phát triển năm 2018 của chúng tôi.
Những kết quả này đang ảnh hưởng đến kế hoạch năm 2019 của chúng tôi,
và trong những tháng tới, chúng tôi sẽ chia sẻ một số ý tưởng với bạn để giải quyết các vấn đề
và nhu cầu cụ thể mà cộng đồng đã chỉ ra cho chúng tôi.
Một lần nữa, cảm ơn tất cả mọi người đã đóng góp vào cuộc khảo sát này!
