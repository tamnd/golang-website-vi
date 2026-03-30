---
title: Kết quả Khảo sát Developer Go 2019
date: 2020-04-20
by:
- Todd Kulesza
tags:
- survey
- community
summary: Phân tích kết quả từ Khảo sát Developer Go 2019.
template: true
---

## Thật đáng kinh ngạc!

Tôi muốn bắt đầu với một lời **cảm ơn chân thành** gửi đến hàng nghìn developer Go
đã tham gia khảo sát năm nay.
Năm 2019, chúng tôi nhận được 10.975 phản hồi, gần [gấp đôi so với năm ngoái](/blog/survey2018-results)!
Thay mặt toàn bộ nhóm, tôi không thể diễn đạt hết mức độ trân trọng của chúng tôi
khi bạn dành thời gian và công sức để kể cho chúng tôi nghe về trải nghiệm với Go. Xin cảm ơn!

## Lưu ý về các năm trước

Những độc giả tinh ý có thể nhận thấy rằng các so sánh theo từng năm của chúng tôi không
hoàn toàn khớp với những con số chúng tôi đã chia sẻ trong quá khứ.
Lý do là từ năm 2016 đến 2018, chúng tôi tính tỷ lệ phần trăm cho mỗi
câu hỏi bằng cách dùng tổng số người đã bắt đầu khảo sát làm mẫu số.
Mặc dù cách này nhất quán, nhưng nó bỏ qua thực tế là không phải ai cũng
hoàn thành khảo sát, với tới 40% người tham gia dừng lại trước khi đến trang cuối,
điều này khiến các câu hỏi xuất hiện muộn hơn trong khảo sát có vẻ kém hơn chỉ vì chúng đứng sau.
Vì vậy, năm nay chúng tôi đã tính lại toàn bộ kết quả (bao gồm cả các phản hồi từ 2016 đến 2018
được hiển thị trong bài này) bằng cách dùng số người đã trả lời một câu hỏi cụ thể làm mẫu số cho câu hỏi đó.
Chúng tôi đã đưa vào số lượng phản hồi năm 2019 cho mỗi biểu đồ, dưới dạng
"n=[số người trả lời]" trên trục x hoặc trong chú thích biểu đồ, để
giúp độc giả hiểu rõ hơn về mức độ tin cậy của bằng chứng đằng sau mỗi phát hiện.

Tương tự, chúng tôi nhận ra rằng trong các khảo sát trước, các lựa chọn xuất hiện sớm hơn
trong danh sách câu trả lời có tỷ lệ được chọn cao hơn một cách không tương xứng.
Để khắc phục điều này, chúng tôi đã thêm yếu tố ngẫu nhiên vào khảo sát.
Một số câu hỏi nhiều lựa chọn của chúng tôi có danh sách lựa chọn không có thứ tự logic,
chẳng hạn như "Tôi viết những loại ứng dụng sau trong Go:
[danh sách các loại ứng dụng]".
Trước đây, các lựa chọn này được sắp xếp theo bảng chữ cái,
nhưng với năm 2019, chúng được hiển thị theo thứ tự ngẫu nhiên cho mỗi người tham gia.
Điều này có nghĩa là so sánh theo từng năm đối với một số câu hỏi không có giá trị cho giai đoạn 2018 đến 2019,
nhưng xu hướng từ 2016 đến 2018 vẫn còn hiệu lực.
Bạn có thể nghĩ đây là việc thiết lập một đường cơ sở chính xác hơn cho năm 2019.
Chúng tôi vẫn giữ nguyên thứ tự bảng chữ cái trong các trường hợp mà người trả lời có khả năng
tìm kiếm một tên cụ thể, chẳng hạn như trình soạn thảo ưa thích của họ.
Chúng tôi sẽ chỉ rõ câu hỏi nào áp dụng điều này ở phần bên dưới.

Thay đổi lớn thứ ba là cải thiện cách phân tích các câu hỏi có câu trả lời văn bản tự do, mở.
Năm ngoái, chúng tôi sử dụng machine learning để phân loại các câu trả lời này một cách sơ bộ nhưng nhanh chóng.
Năm nay, hai nhà nghiên cứu đã phân tích và phân loại thủ công các câu trả lời này,
cho phép phân tích chi tiết hơn nhưng không thể so sánh hợp lệ với
số liệu của năm ngoái.
Giống như việc ngẫu nhiên hóa đã thảo luận ở trên, mục đích của sự thay đổi này là
cung cấp cho chúng tôi một đường cơ sở đáng tin cậy cho năm 2019 trở đi.

## Không nói thêm nữa...

Đây là một bài viết dài. Dưới đây là tóm tắt các phát hiện chính:

- Thông tin nhân khẩu học của người trả lời tương tự như người trả lời khảo sát của Stack Overflow,
  điều này tăng thêm sự tự tin rằng những kết quả này đại diện cho
  cộng đồng developer Go rộng lớn hơn.
- Đa số người trả lời sử dụng Go hàng ngày, và con số này đang có xu hướng tăng dần qua các năm.
- Việc sử dụng Go vẫn tập trung chủ yếu ở các công ty công nghệ,
  nhưng Go ngày càng được tìm thấy trong nhiều ngành công nghiệp hơn,
  chẳng hạn như tài chính và truyền thông.
- Những thay đổi về phương pháp cho thấy hầu hết các chỉ số theo năm
  đều ổn định và cao hơn chúng tôi đã nhận ra trước đây.
- Người trả lời đang dùng Go để giải quyết các vấn đề tương tự,
  đặc biệt là xây dựng dịch vụ API/RPC và CLI,
  bất kể quy mô tổ chức họ làm việc.
- Hầu hết các nhóm cố gắng cập nhật lên bản phát hành Go mới nhất một cách nhanh chóng;
  khi các nhà cung cấp bên thứ ba chậm hỗ trợ bản phát hành Go hiện tại,
  điều này tạo ra rào cản triển khai cho developer.
- Hầu như mọi người trong hệ sinh thái Go hiện đều đang sử dụng module, nhưng vẫn còn một số nhầm lẫn xung quanh việc quản lý package.
- Các lĩnh vực ưu tiên cao cần cải thiện bao gồm cải thiện trải nghiệm developer khi gỡ lỗi,
  làm việc với module và làm việc với dịch vụ cloud.
- VS Code và GoLand tiếp tục thấy mức độ sử dụng ngày càng tăng; hiện chúng được ưa thích bởi 3 trong 4 người trả lời.

## Chúng tôi đã nghe từ ai?

Năm nay chúng tôi đặt ra một số câu hỏi nhân khẩu học mới để giúp chúng tôi hiểu rõ hơn
về những người đã trả lời khảo sát này.
Cụ thể, chúng tôi hỏi về thời gian kinh nghiệm lập trình chuyên nghiệp
và quy mô của các tổ chức nơi mọi người làm việc.
Những câu hỏi này được mô phỏng theo các câu hỏi mà StackOverflow đặt ra trong khảo sát thường niên của họ,
và phân bố các câu trả lời chúng tôi thấy rất gần với kết quả năm 2019 của StackOverflow.
Kết luận của chúng tôi là người trả lời khảo sát này có kinh nghiệm chuyên nghiệp tương đương
và tỷ lệ đại diện tương xứng của các quy mô tổ chức khác nhau
như đối tượng khảo sát StackOverflow (với sự khác biệt rõ ràng là chúng tôi chủ yếu
nghe từ các developer làm việc với Go).
Điều này tăng thêm sự tự tin khi tổng quát hóa các phát hiện này cho khoảng
1 triệu developer Go trên toàn thế giới.
Các câu hỏi nhân khẩu học này cũng sẽ giúp chúng tôi trong tương lai để xác định
những thay đổi theo từng năm nào có thể là kết quả của sự thay đổi trong đối tượng trả lời khảo sát,
thay vì những thay đổi trong tình cảm hay hành vi.

{{image "survey2019/fig1.svg" 700}}
{{image "survey2019/fig2.svg" 700}}

Nhìn vào kinh nghiệm Go, chúng tôi thấy rằng đa số người trả lời (56%) còn
khá mới với Go, đã sử dụng nó chưa đến hai năm.
Đa số cũng cho biết họ sử dụng Go tại nơi làm việc (72%) và ngoài giờ làm việc (62%).
Tỷ lệ người trả lời sử dụng Go trong công việc có vẻ đang tăng dần qua các năm.

Như bạn có thể thấy trong biểu đồ dưới đây, năm 2018 chúng tôi thấy một sự tăng đột biến trong các số liệu này,
nhưng mức tăng đó biến mất vào năm nay.
Đây là một trong nhiều tín hiệu cho thấy đối tượng đã trả lời khảo sát năm 2018
khác biệt đáng kể so với ba năm còn lại.
Trong trường hợp này, họ có nhiều khả năng hơn đáng kể là đang sử dụng Go ngoài
giờ làm việc và dùng một ngôn ngữ khác trong khi làm việc,
nhưng chúng tôi thấy những điểm ngoại lệ tương tự trên nhiều câu hỏi khảo sát.

{{image "survey2019/fig3.svg" 700}}
{{image "survey2019/fig4.svg" 700}}

Những người trả lời đã sử dụng Go lâu nhất có nền tảng khác
so với các developer Go mới hơn.
Những Gopher kỳ cựu này có nhiều khả năng hơn để tuyên bố chuyên môn trong C/C++ và ít
có khả năng hơn để tuyên bố chuyên môn trong JavaScript,
TypeScript và PHP.
Một lưu ý là đây là chuyên môn "tự báo cáo";
có thể hữu ích hơn khi nghĩ về nó như là "sự quen thuộc".
Python có vẻ là ngôn ngữ (ngoài Go) mà hầu hết người trả lời quen thuộc,
bất kể họ đã làm việc với Go bao lâu.

{{image "survey2019/fig5.svg" 750}}

Năm ngoái chúng tôi đã hỏi về ngành công nghiệp mà người trả lời làm việc,
và phát hiện ra rằng đa số cho biết họ làm việc trong các công ty phần mềm,
internet hoặc dịch vụ web.
Năm nay có vẻ như người trả lời đại diện cho nhiều ngành công nghiệp hơn.
Tuy nhiên, chúng tôi cũng đã đơn giản hóa danh sách các ngành công nghiệp để giảm nhầm lẫn từ
các danh mục có thể chồng lấp (ví dụ,
các danh mục riêng biệt cho "Phần mềm" và "Internet / dịch vụ web" từ
năm 2018 được gộp thành "Công nghệ" cho năm 2019).
Vì vậy, đây không phải là một so sánh hoàn toàn tương đương.
Ví dụ, có thể một tác động của việc đơn giản hóa danh sách danh mục
là giảm việc sử dụng danh mục "Phần mềm" như một danh mục tổng hợp cho những người trả lời
viết phần mềm Go cho một ngành không có trong danh sách.

{{image "survey2019/fig6.svg" 700}}

Go là một dự án mã nguồn mở thành công, nhưng điều đó không có nghĩa là các developer
làm việc với nó cũng đang viết phần mềm miễn phí hoặc mã nguồn mở.
Giống như các năm trước, chúng tôi thấy rằng hầu hết người trả lời không thường xuyên đóng góp
cho các dự án mã nguồn mở Go,
với 75% cho biết họ làm như vậy "không thường xuyên" hoặc "không bao giờ".
Khi cộng đồng Go mở rộng, chúng tôi thấy tỷ lệ người trả lời chưa bao giờ
đóng góp cho các dự án mã nguồn mở Go đang tăng dần lên.

{{image "survey2019/fig7.svg" 700}}

## Công cụ developer

Giống như các năm trước, đại đa số người trả lời khảo sát cho biết họ làm việc
với Go trên các hệ thống Linux và macOS.
Đây là một lĩnh vực có sự khác biệt rõ rệt giữa người trả lời của chúng tôi và kết quả năm 2019 của StackOverflow:
trong khảo sát của chúng tôi, chỉ có 20% người trả lời sử dụng Windows làm nền tảng phát triển chính,
trong khi ở StackOverflow con số đó là 45%.
Linux được dùng bởi 66% và macOS bởi 53%, cả hai đều cao hơn nhiều so với đối tượng StackOverflow,
vốn báo cáo 25% và 30%, tương ứng.

{{image "survey2019/fig8.svg" 700}}
{{image "survey2019/fig9.svg" 700}}

Xu hướng hợp nhất trình soạn thảo đã tiếp tục trong năm nay.
GoLand có mức tăng sử dụng mạnh nhất trong năm nay,
tăng từ 24% lên 34%.
Tốc độ tăng trưởng của VS Code chậm lại, nhưng nó vẫn là trình soạn thảo phổ biến nhất trong số người trả lời với 41%.
Kết hợp lại, hai trình soạn thảo này hiện được ưa thích bởi 3 trong 4 người trả lời.

Mọi trình soạn thảo khác đều có mức giảm nhỏ. Điều này không có nghĩa là các trình soạn thảo đó
không được sử dụng,
nhưng chúng không phải là những gì người trả lời cho biết họ _muốn_ sử dụng để viết mã Go.

{{image "survey2019/fig10.svg" 700}}

Năm nay chúng tôi đã thêm một câu hỏi về hệ thống công cụ tài liệu Go nội bộ,
chẳng hạn như [gddo](https://github.com/golang/gddo).
Một thiểu số nhỏ người trả lời (6%) cho biết tổ chức của họ đang chạy
máy chủ tài liệu Go riêng,
mặc dù tỷ lệ này gần như tăng gấp đôi (lên 11%) khi chúng tôi nhìn vào người trả lời
tại các tổ chức lớn (những tổ chức có ít nhất 5.000 nhân viên).
Câu hỏi tiếp theo được hỏi những người trả lời cho biết tổ chức của họ đã ngừng
chạy máy chủ tài liệu riêng cho thấy rằng lý do hàng đầu để loại bỏ
máy chủ là sự kết hợp giữa lợi ích cảm nhận thấp (23%) so với
lượng công sức cần thiết để thiết lập ban đầu và duy trì nó (38%).

{{image "survey2019/fig11.svg" 700}}

## Cảm nhận đối với Go

Đa số lớn người trả lời đồng ý rằng Go đang hoạt động tốt cho
nhóm của họ (86%) và họ muốn sử dụng nó cho dự án tiếp theo (89%).
Chúng tôi cũng thấy rằng hơn một nửa số người trả lời (59%) tin rằng Go rất quan trọng
đối với sự thành công của công ty họ.
Tất cả các chỉ số này đã giữ ổn định kể từ năm 2016.

Việc chuẩn hóa kết quả đã thay đổi hầu hết các con số này cho các năm trước.
Ví dụ, tỷ lệ người trả lời đồng ý với câu "Go đang hoạt động tốt cho nhóm của tôi"
trước đây nằm trong khoảng 50 và 60
vì sự bỏ rơi của người tham gia;
khi chúng tôi loại bỏ những người tham gia chưa bao giờ thấy câu hỏi,
chúng tôi thấy nó đã khá ổn định kể từ năm 2016.

{{image "survey2019/fig12.svg" 700}}

Nhìn vào cảm nhận đối với việc giải quyết vấn đề trong hệ sinh thái Go,
chúng tôi thấy kết quả tương tự.
Tỷ lệ lớn người trả lời đồng ý với từng câu (82% đến 88%),
và những tỷ lệ này đã khá ổn định trong bốn năm qua.

{{image "survey2019/fig13.svg" 700}}

Năm nay chúng tôi đã xem xét sâu hơn về sự hài lòng theo ngành công nghiệp
để thiết lập một đường cơ sở.
Nhìn chung, người trả lời có thái độ tích cực về việc sử dụng Go tại nơi làm việc,
bất kể lĩnh vực ngành công nghiệp.
Chúng tôi thấy có sự biến động nhỏ về sự không hài lòng trong một số lĩnh vực,
đáng chú ý nhất là sản xuất, mà chúng tôi dự định điều tra bằng nghiên cứu tiếp theo.
Tương tự, chúng tôi đã hỏi về sự hài lòng và tầm quan trọng của các
khía cạnh khác nhau của phát triển Go.
Ghép các thước đo này lại với nhau đã làm nổi bật ba chủ đề được chú ý đặc biệt:
gỡ lỗi (bao gồm gỡ lỗi concurrency),
sử dụng module và sử dụng dịch vụ cloud.
Mỗi chủ đề này được đánh giá là "rất" hoặc "cực kỳ" quan trọng bởi đa số
người trả lời nhưng có điểm hài lòng thấp hơn đáng kể so với các chủ đề khác.

{{image "survey2019/fig14.svg" 800}}
{{image "survey2019/fig15.svg" 750}}

Khi nói đến cảm nhận đối với cộng đồng Go,
chúng tôi thấy có sự khác biệt so với các năm trước.
Trước tiên, có sự sụt giảm trong tỷ lệ người trả lời đồng ý với
câu "Tôi cảm thấy được chào đón trong cộng đồng Go", từ 82% xuống 75%.
Đào sâu hơn cho thấy tỷ lệ người trả lời "hơi đồng ý"
hoặc "đồng ý vừa phải" giảm xuống,
trong khi tỷ lệ "không đồng ý cũng không phản đối" và "hoàn toàn đồng ý"
cả hai đều tăng (lên 5 và 7 điểm, tương ứng).
Sự phân cực này gợi ý rằng có hai hoặc nhiều nhóm có trải nghiệm trong
cộng đồng Go đang có sự phân kỳ,
và vì vậy đây là một lĩnh vực khác mà chúng tôi dự định điều tra thêm.

Những khác biệt lớn khác là xu hướng tăng rõ ràng trong các câu trả lời cho câu
"Tôi cảm thấy được chào đón để đóng góp cho dự án Go" và một mức tăng lớn theo năm
trong tỷ lệ người trả lời cảm thấy ban lãnh đạo dự án Go
hiểu nhu cầu của họ.

Tất cả những kết quả này cho thấy mô hình đồng ý cao hơn tương quan với
kinh nghiệm Go ngày càng nhiều,
bắt đầu từ khoảng hai năm.
Nói cách khác, người trả lời đã sử dụng Go càng lâu,
thì họ càng có khả năng đồng ý với mỗi câu trong số những câu đó.

{{image "survey2019/fig16.svg" 700}}

Điều này có lẽ không gây ngạc nhiên, nhưng những người đã trả lời Khảo sát Developer Go
có xu hướng thích Go.
Tuy nhiên, chúng tôi cũng muốn hiểu xem người trả lời thích làm việc với _ngôn ngữ nào khác_.
Hầu hết các con số này không thay đổi đáng kể so với các năm trước,
với hai ngoại lệ:
TypeScript (đã tăng 10 điểm),
và Rust (tăng 7 điểm).
Khi chúng tôi phân tích các kết quả này theo thời gian sử dụng Go,
chúng tôi thấy cùng một mô hình như chúng tôi đã tìm thấy đối với chuyên môn ngôn ngữ.
Cụ thể, Python là ngôn ngữ và hệ sinh thái mà các developer Go
có nhiều khả năng cũng thích xây dựng.

{{image "survey2019/fig17.svg" 700}}

Năm 2018, lần đầu tiên chúng tôi hỏi câu hỏi "Bạn có giới thiệu không..." [Net Promoter Score](https://en.wikipedia.org/wiki/Net_Promoter) (NPS),
cho điểm số là 61.
Năm nay kết quả NPS của chúng tôi là 60, về mặt thống kê không thay đổi (67% "người quảng bá"
trừ đi 7% "người phản đối").

{{image "survey2019/fig18.svg" 700}}

## Làm việc với Go

Xây dựng dịch vụ API/RPC (71%) và CLI (62%) vẫn là những cách sử dụng Go phổ biến nhất.
Biểu đồ dưới đây có vẻ cho thấy những thay đổi lớn từ năm 2018,
nhưng đây rất có khả năng là kết quả của việc ngẫu nhiên hóa thứ tự các lựa chọn,
vốn trước đây được liệt kê theo thứ tự bảng chữ cái:
3 trong 4 lựa chọn bắt đầu bằng chữ 'A' giảm xuống,
trong khi tất cả những cái khác vẫn ổn định hoặc tăng lên.
Vì vậy, biểu đồ này được hiểu rõ nhất là một đường cơ sở chính xác hơn cho năm 2019
với xu hướng từ 2016 đến 2018.
Ví dụ, chúng tôi tin rằng tỷ lệ người trả lời xây dựng dịch vụ web
trả về HTML đã giảm kể từ năm 2016 nhưng có khả năng bị đếm thiếu vì câu trả lời này luôn nằm ở cuối danh sách dài các lựa chọn.
Chúng tôi cũng đã phân tích điều này theo quy mô tổ chức và ngành công nghiệp nhưng không tìm thấy sự khác biệt đáng kể:
có vẻ như người trả lời sử dụng Go theo những cách tương tự nhau dù họ làm việc
tại một công ty khởi nghiệp công nghệ nhỏ hay một doanh nghiệp bán lẻ lớn.

Một câu hỏi liên quan đã hỏi về các lĩnh vực rộng hơn mà người trả lời làm việc với Go.
Lĩnh vực phổ biến nhất cho đến nay là phát triển web (66%),
nhưng các lĩnh vực phổ biến khác bao gồm cơ sở dữ liệu (45%),
lập trình mạng (42%), lập trình hệ thống (38%),
và các tác vụ DevOps (37%).

{{image "survey2019/fig19.svg" 700}}
{{image "survey2019/fig20.svg" 700}}

Ngoài những gì người trả lời đang xây dựng,
chúng tôi cũng hỏi về một số kỹ thuật phát triển họ sử dụng.
Đa số lớn người trả lời cho biết họ dựa vào text log để gỡ lỗi (88%),
và các câu trả lời văn bản tự do của họ cho thấy điều này là vì các công cụ thay thế
rất khó sử dụng hiệu quả.
Tuy nhiên, gỡ lỗi từng bước cục bộ (ví dụ: với Delve),
profiling và kiểm tra với race detector cũng không hiếm gặp,
với khoảng 50% người trả lời dựa vào ít nhất một trong những kỹ thuật này.

{{image "survey2019/fig21.svg" 700}}

Về quản lý package, chúng tôi thấy rằng đại đa số người trả lời
đã áp dụng module cho Go (89%).
Đây là một sự thay đổi lớn đối với các developer,
và gần như toàn bộ cộng đồng dường như đang trải qua nó cùng một lúc.

{{image "survey2019/fig22.svg" 700}}

Chúng tôi cũng thấy rằng 75% người trả lời đánh giá bản phát hành Go hiện tại để sử dụng trong production,
với thêm 12% chờ đợi một chu kỳ phát hành.
Điều này gợi ý rằng đa số lớn các developer Go đang sử dụng (hoặc ít nhất là
cố gắng sử dụng) bản phát hành ổn định hiện tại hoặc trước đó,
làm nổi bật tầm quan trọng của các nhà cung cấp platform-as-a-service trong việc nhanh chóng
hỗ trợ các bản phát hành ổn định mới của Go.

{{image "survey2019/fig23.svg" 700}}

## Go trong đám mây

Go được thiết kế với tính toán phân tán hiện đại trong đầu,
và chúng tôi muốn tiếp tục cải thiện trải nghiệm developer khi xây dựng
dịch vụ đám mây với Go.
Năm nay chúng tôi đã mở rộng các câu hỏi về phát triển đám mây để
hiểu rõ hơn cách người trả lời đang làm việc với các nhà cung cấp đám mây,
những gì họ thích về trải nghiệm developer hiện tại,
và những gì có thể được cải thiện.
Như đã đề cập trước đó, một số kết quả năm 2018 có vẻ là điểm ngoại lệ,
chẳng hạn như kết quả bất ngờ thấp cho các máy chủ tự sở hữu,
và kết quả bất ngờ cao cho triển khai GCP.

Chúng tôi thấy hai xu hướng rõ ràng:

1. Ba nhà cung cấp đám mây toàn cầu lớn nhất (Amazon Web Services,
Google Cloud Platform và Microsoft Azure) đều có vẻ đang có xu hướng tăng
về việc sử dụng trong số người trả lời khảo sát,
trong khi hầu hết các nhà cung cấp khác được sử dụng bởi tỷ lệ nhỏ hơn của người trả lời mỗi năm.
2. Triển khai on-prem lên các máy chủ tự sở hữu hoặc thuộc sở hữu công ty tiếp tục
giảm và hiện tại đang kết đặc với AWS (44% so với
42%) là các mục tiêu triển khai phổ biến nhất.

Nhìn vào loại nền tảng đám mây mà người trả lời đang sử dụng,
chúng tôi thấy sự khác biệt giữa các nhà cung cấp lớn.
Người trả lời triển khai lên AWS và Azure có nhiều khả năng nhất là đang sử dụng VMs
trực tiếp (65% và 51%,
tương ứng), trong khi những người triển khai lên GCP có khả năng gần gấp đôi
để sử dụng nền tảng Kubernetes được quản lý (GKE,
64%) so với VMs (35%).
Chúng tôi cũng thấy rằng người trả lời triển khai lên AWS có khả năng bằng nhau
để sử dụng nền tảng Kubernetes được quản lý (32%) hay nền tảng serverless được quản lý
(AWS Lambda, 33%).
Cả GCP (17%) và Azure (7%) đều có tỷ lệ thấp hơn của người trả lời sử dụng
nền tảng serverless,
và các câu trả lời văn bản tự do gợi ý lý do chính là sự chậm trễ trong hỗ trợ
runtime Go mới nhất trên các nền tảng này.

Nhìn chung, đa số người trả lời đều hài lòng với việc sử dụng Go trên cả ba
nhà cung cấp đám mây lớn.
Người trả lời báo cáo mức độ hài lòng tương tự khi phát triển Go cho
AWS (80% hài lòng) và GCP (78%).
Azure nhận được điểm hài lòng thấp hơn (57% hài lòng),
và các câu trả lời văn bản tự do cho thấy yếu tố chính là nhận thức rằng
Go thiếu hỗ trợ hạng nhất trên nền tảng này (25% câu trả lời văn bản tự do).
Ở đây, "hỗ trợ hạng nhất" đề cập đến việc luôn cập nhật bản phát hành Go mới nhất,
và đảm bảo các tính năng mới có sẵn cho developer Go vào thời điểm ra mắt.
Đây cũng là vấn đề đau đớn hàng đầu được báo cáo bởi người trả lời sử dụng GCP (14%),
và đặc biệt tập trung vào hỗ trợ runtime Go mới nhất trong các triển khai serverless.
Người trả lời triển khai lên AWS, ngược lại,
có nhiều khả năng nhất để nói rằng SDK cần được cải thiện,
chẳng hạn như trở nên idiomatic hơn (21%).
Cải thiện SDK cũng là yêu cầu phổ biến thứ hai cho cả developer GCP (9%)
và Azure (18%).

{{image "survey2019/fig24.svg" 700}}
{{image "survey2019/fig25.svg" 700}}
{{image "survey2019/fig26.svg" 700}}

## Các điểm đau

Những lý do hàng đầu khiến người trả lời không thể sử dụng Go nhiều hơn vẫn là làm việc
trên một dự án bằng ngôn ngữ khác (56%),
làm việc trong một nhóm thích sử dụng ngôn ngữ khác (37%),
và sự thiếu hụt một tính năng quan trọng trong chính Go (25%).

Đây là một trong những câu hỏi mà chúng tôi đã ngẫu nhiên hóa danh sách lựa chọn,
vì vậy các so sánh theo từng năm không có giá trị,
mặc dù xu hướng 2016 đến 2018 thì có.
Ví dụ, chúng tôi tự tin rằng số lượng developer không thể sử dụng
Go thường xuyên hơn vì nhóm của họ thích ngôn ngữ khác đang giảm dần qua các năm,
nhưng chúng tôi không biết liệu sự giảm đó có tăng tốc mạnh trong năm nay,
hay luôn thấp hơn một chút so với ước tính từ 2016 đến 2018 của chúng tôi.

{{image "survey2019/fig27.svg" 700}}

Hai rào cản áp dụng hàng đầu (làm việc trên dự án Go không phải mới và
làm việc trong nhóm thích ngôn ngữ khác) không có giải pháp kỹ thuật trực tiếp,
nhưng những rào cản còn lại thì có thể.
Vì vậy, năm nay chúng tôi đã yêu cầu thêm chi tiết,
để hiểu rõ hơn cách chúng tôi có thể giúp developer tăng cường việc sử dụng Go.
Các biểu đồ trong phần còn lại của mục này dựa trên các câu trả lời văn bản tự do
được phân loại thủ công,
vì vậy chúng có _đuôi rất dài_;
các danh mục chiếm ít hơn 3% tổng số câu trả lời đã được nhóm
vào danh mục "Khác" cho mỗi biểu đồ.
Một câu trả lời có thể đề cập đến nhiều chủ đề,
vì vậy các biểu đồ không cộng lại thành 100%.

Trong số 25% người trả lời cho biết Go thiếu các tính năng ngôn ngữ mà họ cần,
79% đề đến generics như một tính năng còn thiếu quan trọng.
Cải tiến liên tục cho xử lý lỗi (ngoài các thay đổi Go 1.13) được trích dẫn bởi 22%,
trong khi 13% yêu cầu thêm tính năng lập trình hàm,
đặc biệt là chức năng map/filter/reduce tích hợp sẵn.
Để rõ ràng, những con số này là từ tập hợp con người trả lời cho biết họ
có thể sử dụng Go nhiều hơn nếu nó không thiếu một hoặc nhiều tính năng quan trọng mà họ cần,
không phải từ toàn bộ đối tượng người trả lời khảo sát.

{{image "survey2019/fig28.svg" 700}}

Những người trả lời cho biết Go "không phải là ngôn ngữ phù hợp" cho những gì họ làm việc
có nhiều lý do và trường hợp sử dụng khác nhau.
Phổ biến nhất là họ làm việc trên một số hình thức phát triển front-end (22%),
chẳng hạn như GUI cho web, desktop hoặc di động.
Một câu trả lời phổ biến khác là người trả lời nói rằng họ làm việc trong một lĩnh vực
với một ngôn ngữ đã chiếm ưu thế (9%),
khiến việc sử dụng thứ gì đó khác trở nên thử thách.
Một số người trả lời cũng cho chúng tôi biết họ đang đề cập đến lĩnh vực nào (hoặc đơn giản
đề cập đến lĩnh vực mà không đề cập đến ngôn ngữ khác phổ biến hơn),
chúng tôi hiển thị thông qua các hàng "Tôi làm việc trên [lĩnh vực]" bên dưới.
Một lý do hàng đầu khác được người trả lời trích dẫn là cần hiệu suất tốt hơn (9%),
đặc biệt là cho tính toán thời gian thực.

{{image "survey2019/fig29.svg" 700}}

Những thách thức lớn nhất mà người trả lời báo cáo vẫn phần lớn nhất quán với năm ngoái.
Việc thiếu generics và quản lý module/package trong Go vẫn đứng đầu danh sách
(15% và 12% số câu trả lời,
tương ứng), và tỷ lệ người trả lời nêu bật các vấn đề về hệ thống công cụ đã tăng lên.
Những con số này khác với các biểu đồ trên vì câu hỏi này
được hỏi _tất cả_ người trả lời,
bất kể họ nói rào cản áp dụng Go lớn nhất của họ là gì.
Tất cả ba lĩnh vực này đều là trọng tâm của nhóm Go trong năm nay,
và chúng tôi hy vọng sẽ cải thiện đáng kể trải nghiệm developer,
đặc biệt là xung quanh module, hệ thống công cụ và trải nghiệm bắt đầu,
trong những tháng tới.

{{image "survey2019/fig30.svg" 700}}

Chẩn đoán lỗi và vấn đề hiệu suất có thể là thử thách trong bất kỳ ngôn ngữ nào.
Người trả lời cho chúng tôi biết thách thức hàng đầu của họ cho cả hai vấn đề này không phải là điều gì đó
đặc thù cho việc triển khai hoặc hệ thống công cụ của Go,
mà là một vấn đề cơ bản hơn:
sự thiếu hụt kiến thức, kinh nghiệm hoặc các thực hành tốt nhất được tự báo cáo.
Chúng tôi hy vọng sẽ giúp giải quyết những khoảng trống kiến thức này thông qua tài liệu và các
tài liệu giáo dục khác trong năm nay.
Các vấn đề lớn khác liên quan đến hệ thống công cụ,
cụ thể là nhận thức về tỷ lệ chi phí/lợi ích không thuận lợi khi học/sử dụng
hệ thống công cụ gỡ lỗi và profiling của Go,
và những thách thức trong việc khiến hệ thống công cụ hoạt động trong các môi trường khác nhau (ví dụ:
gỡ lỗi trong container hoặc lấy hồ sơ hiệu suất từ hệ thống production).

{{image "survey2019/fig31.svg" 700}}
{{image "survey2019/fig32.svg" 700}}

Cuối cùng, khi chúng tôi hỏi điều gì sẽ cải thiện hỗ trợ Go nhất trong môi trường
soạn thảo của người trả lời,
câu trả lời phổ biến nhất là cải tiến tổng thể hoặc hỗ trợ tốt hơn
cho language server (gopls, 19%).
Điều này được mong đợi, vì gopls thay thế khoảng 80 công cụ hiện có và vẫn đang trong giai đoạn beta.
Khi người trả lời cụ thể hơn về những gì họ muốn thấy được cải thiện,
họ có nhiều khả năng nhất để báo cáo trải nghiệm gỡ lỗi (14%) và hoàn thành mã
nhanh hơn hoặc đáng tin cậy hơn (13%).
Một số người tham gia cũng đề cập rõ ràng đến nhu cầu thường xuyên
khởi động lại VS Code khi sử dụng gopls (8%);
trong thời gian kể từ khi khảo sát này được thực hiện (cuối tháng 11 đến đầu tháng 12 năm 2019),
nhiều cải tiến gopls này đã được đưa vào,
và đây tiếp tục là một lĩnh vực ưu tiên cao cho nhóm.

{{image "survey2019/fig33.svg" 700}}

## Cộng đồng Go

Khoảng hai phần ba người trả lời đã dùng Stack Overflow để trả lời các câu hỏi liên quan đến Go (64%).
Các nguồn câu trả lời hàng đầu khác là godoc.org (47%),
đọc trực tiếp mã nguồn (42%) và golang.org (33%).

{{image "survey2019/fig34.svg" 700}}

Đuôi dài trong biểu đồ trước làm nổi bật sự đa dạng lớn của các
nguồn khác nhau (hầu hết tất cả đều do cộng đồng điều khiển) và các phương thức mà người trả lời
dựa vào để vượt qua các thách thức khi phát triển với Go.
Thực ra, đối với nhiều Gopher, đây có thể là một trong những điểm tương tác chính
với cộng đồng rộng lớn hơn:
khi cộng đồng của chúng tôi mở rộng, chúng tôi thấy tỷ lệ ngày càng cao hơn của người trả lời
không tham dự bất kỳ sự kiện nào liên quan đến Go.
Vào năm 2019, tỷ lệ đó gần như đạt đến hai phần ba người trả lời (62%).

{{image "survey2019/fig35.svg" 700}}

Do cập nhật các hướng dẫn quyền riêng tư trên toàn Google,
chúng tôi không còn có thể hỏi về quốc gia nào mà người trả lời sống.
Thay vào đó, chúng tôi hỏi về ngôn ngữ nói/viết được ưa thích như một đại diện rất thô
cho việc sử dụng Go trên toàn thế giới,
với lợi ích là cung cấp dữ liệu cho các nỗ lực bản địa hóa tiềm năng.

Vì khảo sát này bằng tiếng Anh, có thể có sự thiên lệch mạnh về
người nói tiếng Anh và người từ các khu vực nơi tiếng Anh là ngôn ngữ thứ hai hoặc thứ ba phổ biến.
Vì vậy, các con số không phải tiếng Anh nên được hiểu là mức tối thiểu có thể
thay vì ước tính đối tượng toàn cầu của Go.

{{image "survey2019/fig36.svg" 700}}

Chúng tôi thấy rằng 12% người trả lời xác định với một nhóm truyền thống thiểu đại diện (ví dụ:
sắc tộc, bản sắc giới tính, v.v.) và 3% tự xác định là nữ.
(Câu hỏi này nên đã nói "phụ nữ" thay vì "nữ".
Lỗi này đã được sửa trong bản nháp khảo sát của chúng tôi cho năm 2020,
và chúng tôi xin lỗi vì điều đó.)
Chúng tôi nghi ngờ mạnh mẽ rằng 3% này đang đếm thiếu phụ nữ trong cộng đồng Go.
Ví dụ, chúng tôi biết rằng các nhà phát triển phần mềm nữ ở Mỹ trả lời Khảo sát Developer StackOverflow ở [khoảng nửa tỷ lệ chúng tôi kỳ vọng dựa trên số liệu việc làm ở Mỹ](https://insights.stackoverflow.com/survey/2019#developer-profile-_-developer-type) (11% so với 20%).
Vì chúng tôi không biết tỷ lệ phản hồi ở Mỹ,
chúng tôi không thể an toàn suy luận từ những con số này ngoài việc nói rằng tỷ lệ thực tế
có thể cao hơn 3%.
Hơn nữa, GDPR yêu cầu chúng tôi thay đổi cách chúng tôi hỏi về thông tin nhạy cảm,
bao gồm giới tính và các nhóm truyền thống thiểu đại diện.
Thật không may, những thay đổi này ngăn chúng tôi có thể thực hiện so sánh hợp lệ
của những con số này với các năm trước.

Những người trả lời xác định với các nhóm thiểu đại diện hoặc thích không
trả lời câu hỏi này cho thấy tỷ lệ không đồng ý cao hơn với câu
"Tôi cảm thấy được chào đón trong cộng đồng Go" (8% so với
4%) so với những người không xác định với nhóm thiểu đại diện,
làm nổi bật tầm quan trọng của các nỗ lực tiếp cận liên tục của chúng tôi.

{{image "survey2019/fig37.svg" 700}}
{{image "survey2019/fig38.svg" 700}}
{{image "survey2019/fig39.svg" 800}}

## Kết luận

Chúng tôi hy vọng bạn đã thích xem kết quả khảo sát developer năm 2019 của chúng tôi.
Hiểu được kinh nghiệm và thách thức của developer giúp chúng tôi lên kế hoạch và ưu tiên công việc cho năm 2020.
Một lần nữa, xin cảm ơn rất nhiều đến tất cả những người đã đóng góp cho khảo sát này,
phản hồi của bạn đang giúp định hướng hướng đi của Go trong năm tới và hơn thế nữa.
