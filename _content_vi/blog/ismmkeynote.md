---
title: "Hành trình đến với Go: Câu chuyện của bộ gom rác Go"
date: 2018-07-12
by:
- Rick Hudson
summary: Bài nói kỹ thuật về cấu trúc và chi tiết của bộ gom rác Go mới với độ trễ thấp.
template: true
---


Đây là bản ghi lại bài keynote tôi trình bày tại Hội nghị chuyên đề Quốc tế
về Quản lý Bộ nhớ (ISMM) ngày 18 tháng 6 năm 2018.
Trong 25 năm qua, ISMM là diễn đàn hàng đầu để công bố các bài báo về quản lý
bộ nhớ và thu gom rác, và tôi cảm thấy vinh dự khi được mời làm diễn giả keynote.

## Tóm tắt

Các tính năng, mục tiêu và trường hợp sử dụng của ngôn ngữ Go đã buộc chúng tôi phải
suy nghĩ lại toàn bộ tầng xử lý bộ gom rác và dẫn chúng tôi đến một nơi bất ngờ.
Hành trình đó thực sự đầy hứng khởi. Bài nói này mô tả hành trình của chúng tôi.
Đó là hành trình được thúc đẩy bởi mã nguồn mở và yêu cầu của môi trường production tại Google.
Bao gồm cả những chặng đường lạc vào các ngõ cụt nơi con số đã dẫn chúng tôi trở về.
Bài nói này sẽ cung cấp cái nhìn sâu sắc về cách chúng tôi làm và lý do tại sao,
vị trí của chúng tôi vào năm 2018, và sự chuẩn bị của Go cho phần tiếp theo của hành trình.

## Tiểu sử

Richard L. Hudson (Rick) được biết đến nhiều nhất qua công trình nghiên cứu về quản lý bộ nhớ,
bao gồm việc phát minh các thuật toán Train, Sapphire và Mississippi Delta,
cũng như GC stack maps cho phép thu gom rác trong các ngôn ngữ có kiểu tĩnh
như Modula-3, Java, C# và Go.
Rick hiện là thành viên của nhóm Go tại Google, nơi ông đang làm việc về
thu gom rác và các vấn đề runtime của Go.

Liên hệ: rlh@golang.org

Bình luận: Xem [thảo luận trên golang-dev](https://groups.google.com/forum/#!topic/golang-dev/UuDv7W1Hsns).

## Bản ghi

{{image "ismmkeynote/image63.png"}}

Rick Hudson đây.

Đây là bài nói về runtime của Go, cụ thể là bộ gom rác.
Tôi có khoảng 45 đến 50 phút tài liệu chuẩn bị và sau đó chúng ta sẽ
có thời gian thảo luận, và tôi sẽ ở đây nên đừng ngại đến gặp tôi sau buổi nói chuyện.

{{image "ismmkeynote/image24.png"}}

Trước khi bắt đầu, tôi muốn ghi nhận một số người.

Rất nhiều điều hay trong bài nói này là công sức của Austin Clements.
Những người khác trong nhóm Go Cambridge, Russ, Than, Cherry và David
là một nhóm luôn cuốn hút, thú vị và vui vẻ khi làm việc cùng.

Chúng tôi cũng muốn cảm ơn 1,6 triệu người dùng Go trên toàn thế giới vì đã
đặt ra cho chúng tôi những bài toán thú vị.
Nếu không có họ, nhiều vấn đề trong số này sẽ không bao giờ được phát hiện.

Và cuối cùng, tôi muốn ghi nhận công lao của Renee French vì tất cả những chú Gopher
đáng yêu mà cô ấy đã tạo ra trong nhiều năm qua.
Bạn sẽ thấy một số chú Gopher đó xuất hiện trong bài nói.

{{image "ismmkeynote/image38.png"}}

Trước khi đi sâu vào nội dung, chúng ta thực sự cần phải cho thấy cách nhìn của GC đối với Go trông như thế nào.

{{image "ismmkeynote/image32.png"}}

Trước tiên, chương trình Go có hàng trăm nghìn stack.
Chúng được quản lý bởi bộ lên lịch (scheduler) của Go và luôn bị preempt tại các điểm GC safepoint.
Bộ lên lịch Go ghép nhiều goroutine lên các luồng OS, lý tưởng là một luồng OS mỗi luồng phần cứng.
Chúng tôi quản lý các stack và kích thước của chúng bằng cách sao chép và cập nhật con trỏ trong stack.
Đây là thao tác cục bộ nên nó mở rộng khá tốt.

{{image "ismmkeynote/image22.png"}}

Điều quan trọng tiếp theo là Go là một ngôn ngữ định hướng giá trị (value-oriented)
theo truyền thống của các ngôn ngữ hệ thống dạng C, chứ không phải định hướng tham chiếu (reference-oriented)
như hầu hết các ngôn ngữ runtime được quản lý.
Ví dụ, đây là cách một kiểu trong package tar được bố cục trong bộ nhớ.
Tất cả các trường đều được nhúng trực tiếp vào giá trị Reader.
Điều này cho phép lập trình viên kiểm soát nhiều hơn về bố cục bộ nhớ khi cần thiết.
Người ta có thể đặt gần nhau các trường có giá trị liên quan, điều này giúp cải thiện hiệu suất bộ nhớ đệm (cache).

Định hướng giá trị cũng hỗ trợ giao tiếp với hàm ngoại vi (foreign function interface).
Chúng tôi có FFI nhanh với C và C++. Rõ ràng Google có rất nhiều cơ sở hạ tầng được xây dựng bằng C++.
Go không thể chờ đợi để triển khai lại tất cả những thứ đó bằng Go, vì vậy Go phải có
khả năng truy cập các hệ thống đó thông qua giao tiếp hàm ngoại vi.

Một quyết định thiết kế này đã dẫn đến một số điều thú vị nhất phải xảy ra với runtime.
Đây có lẽ là điều quan trọng nhất phân biệt Go với các ngôn ngữ GC khác.

{{image "ismmkeynote/image60.png"}}

Tất nhiên Go có thể có con trỏ, và thực tế chúng có thể có con trỏ nội tại (interior pointer).
Các con trỏ như vậy giữ cho toàn bộ giá trị sống và chúng khá phổ biến.

{{image "ismmkeynote/image29.png"}}

Chúng tôi cũng có hệ thống biên dịch trước (ahead of time compilation) nên binary chứa toàn bộ runtime.

Không có JIT recompilation. Điều này có ưu và nhược điểm.
Trước tiên, khả năng tái tạo lại quá trình thực thi chương trình dễ dàng hơn rất nhiều,
giúp việc cải thiện trình biên dịch nhanh hơn.

Mặt tiêu cực là chúng tôi không có cơ hội để thực hiện các tối ưu hóa dựa trên phản hồi như với JIT.

Vì vậy, có cả ưu và nhược điểm.

{{image "ismmkeynote/image13.png"}}

Go đi kèm với hai núm điều chỉnh để kiểm soát GC.
Núm đầu tiên là GCPercent. Về cơ bản đây là núm điều chỉnh để bạn muốn sử dụng
bao nhiêu CPU và bao nhiêu bộ nhớ.
Giá trị mặc định là 100, nghĩa là một nửa heap dành cho bộ nhớ đang sống
và nửa còn lại dành cho phân bổ.
Bạn có thể điều chỉnh theo cả hai hướng.

MaxHeap, hiện chưa được phát hành nhưng đang được sử dụng và đánh giá nội bộ,
cho phép lập trình viên đặt kích thước heap tối đa.
Hết bộ nhớ (OOM) rất khó chịu với Go; các đột biến tăng bộ nhớ tạm thời nên được
xử lý bằng cách tăng chi phí CPU, không phải bằng cách hủy bỏ.
Về cơ bản, nếu GC nhận thấy áp lực bộ nhớ, nó sẽ thông báo cho ứng dụng rằng
ứng dụng nên giảm tải.
Khi mọi thứ trở lại bình thường, GC thông báo cho ứng dụng rằng nó có thể
quay lại tải thông thường.
MaxHeap cũng cung cấp nhiều tính linh hoạt hơn trong việc lên lịch.
Thay vì luôn lo lắng về lượng bộ nhớ còn lại, runtime có thể tăng kích thước
heap lên đến MaxHeap.

Điều này kết thúc phần thảo luận của chúng ta về các phần của Go quan trọng với bộ gom rác.

{{image "ismmkeynote/image3.png"}}

Bây giờ hãy nói về runtime của Go và cách chúng ta đã đến đây, đã đến được nơi chúng ta đang đứng.

{{image "ismmkeynote/image59.png"}}

Đó là năm 2014. Nếu Go không giải quyết được vấn đề độ trễ GC này bằng cách nào đó
thì Go sẽ không thành công. Điều đó đã rõ ràng.

Các ngôn ngữ mới khác cũng đang đối mặt với vấn đề tương tự.
Các ngôn ngữ như Rust đã đi theo hướng khác, nhưng chúng ta sẽ nói về con đường mà Go đã chọn.

Tại sao độ trễ lại quan trọng đến vậy?

{{image "ismmkeynote/image7.png"}}

Toán học hoàn toàn không thể tha thứ về điều này.

Mục tiêu mức độ dịch vụ (SLO) về độ trễ GC cô lập ở phân vị 99 (99%ile),
ví dụ như 99% thời gian một chu kỳ GC mất dưới 10ms, đơn giản là không mở rộng được.
Điều quan trọng là độ trễ trong toàn bộ phiên làm việc hoặc trong quá trình sử dụng ứng dụng
nhiều lần trong ngày.
Giả sử một phiên duyệt một số trang web kết thúc với 100 yêu cầu đến máy chủ
trong phiên, hoặc tạo 20 yêu cầu và bạn có 5 phiên trong ngày.
Trong trường hợp đó, chỉ 37% người dùng sẽ có trải nghiệm dưới 10ms một cách nhất quán
trong toàn bộ phiên làm việc.

Nếu bạn muốn 99% người dùng có trải nghiệm dưới 10ms,
như chúng tôi đề xuất, toán học nói rằng bạn thực sự cần nhắm đến 4 số 9, tức là phân vị 99,99%.

Đó là năm 2014 và Jeff Dean vừa công bố bài báo có tên "The Tail at Scale"
đào sâu hơn vào vấn đề này.
Nó đang được đọc rộng rãi ở Google vì nó có hậu quả nghiêm trọng
cho việc mở rộng quy mô của Google trong tương lai.

Chúng tôi gọi vấn đề này là sự chuyên quyền của những con số 9.

{{image "ismmkeynote/image36.png"}}

Vậy làm thế nào để chống lại sự chuyên quyền của những con số 9?

Nhiều việc đã được thực hiện vào năm 2014.

Nếu bạn muốn 10 câu trả lời, hãy yêu cầu nhiều hơn và lấy 10 cái đầu tiên để đưa lên trang tìm kiếm.
Nếu yêu cầu vượt phân vị 50, hãy phát lại hoặc chuyển tiếp yêu cầu đó đến máy chủ khác.
Nếu GC sắp chạy, hãy từ chối các yêu cầu mới hoặc chuyển tiếp đến máy chủ khác cho đến khi GC hoàn thành.
Và vân vân.

Tất cả những cách làm này đến từ những người rất thông minh với những vấn đề rất thực,
nhưng chúng không giải quyết gốc rễ của vấn đề độ trễ GC.
Ở quy mô Google, chúng tôi phải giải quyết vấn đề gốc rễ. Tại sao?

{{image "ismmkeynote/image48.png"}}

Dư thừa sẽ không mở rộng được, dư thừa rất tốn kém. Nó đòi hỏi các trang trại máy chủ mới.

Chúng tôi hy vọng có thể giải quyết vấn đề này và coi đó là cơ hội để cải thiện
hệ sinh thái máy chủ, và trong quá trình đó, cứu một số cánh đồng ngô và
cho từng hạt ngô có cơ hội đạt tầm gối vào ngày 4 tháng 7 và phát huy hết tiềm năng của nó.

{{image "ismmkeynote/image56.png"}}

Đây là SLO năm 2014. Đúng là tôi đang đặt mục tiêu thấp,
tôi mới gia nhập nhóm, đó là một quy trình mới với tôi,
và tôi không muốn hứa hẹn quá nhiều.

Hơn nữa, các bài thuyết trình về độ trễ GC trong các ngôn ngữ khác khá đáng sợ.

{{image "ismmkeynote/image67.png"}}

Kế hoạch ban đầu là thực hiện GC sao chép đồng thời không cần read barrier.
Đó là kế hoạch dài hạn. Có rất nhiều sự không chắc chắn về chi phí của read barrier
nên Go muốn tránh chúng.

Nhưng ngắn hạn vào năm 2014, chúng tôi phải thu xếp lại trật tự.
Chúng tôi phải chuyển đổi toàn bộ runtime và trình biên dịch sang Go.
Chúng được viết bằng C vào thời điểm đó. Không còn C nữa,
không còn đuôi dài của các lỗi do lập trình viên C không hiểu GC nhưng có ý tưởng hay
về cách sao chép chuỗi.
Chúng tôi cũng cần một cái gì đó nhanh chóng và tập trung vào độ trễ nhưng
mức độ ảnh hưởng đến hiệu suất phải nhỏ hơn mức tăng tốc mà trình biên dịch cung cấp.
Vì vậy, chúng tôi bị giới hạn. Về cơ bản, chúng tôi chỉ có một năm cải thiện hiệu suất trình biên dịch
mà chúng tôi có thể tiêu thụ để làm GC đồng thời.
Nhưng chỉ vậy thôi. Chúng tôi không thể làm chậm chương trình Go.
Điều đó sẽ không thể chấp nhận được vào năm 2014.

{{image "ismmkeynote/image28.png"}}

Vì vậy, chúng tôi lùi lại một chút. Chúng tôi sẽ không thực hiện phần sao chép.

Quyết định là thực hiện thuật toán đồng thời ba màu (tri-color concurrent algorithm).
Trước đó trong sự nghiệp của tôi, Eliot Moss và tôi đã thực hiện các chứng minh về tạp chí
cho thấy thuật toán của Dijkstra hoạt động với nhiều luồng ứng dụng.
Chúng tôi cũng chỉ ra rằng chúng tôi có thể loại bỏ các vấn đề STW (stop-the-world),
và chúng tôi có bằng chứng rằng điều đó có thể thực hiện được.

Chúng tôi cũng lo ngại về tốc độ trình biên dịch,
tức là mã mà trình biên dịch tạo ra.
Nếu chúng tôi giữ write barrier tắt hầu hết thời gian, các tối ưu hóa trình biên dịch
sẽ bị ảnh hưởng tối thiểu và nhóm trình biên dịch có thể tiến bộ nhanh.
Go cũng rất cần thành công ngắn hạn vào năm 2015.

{{image "ismmkeynote/image55.png"}}

Hãy xem xét một số điều chúng tôi đã làm.

Chúng tôi dùng span phân tách theo kích thước (size segregated span). Con trỏ nội tại là một vấn đề.

Bộ gom rác cần tìm điểm bắt đầu của đối tượng một cách hiệu quả.
Nếu nó biết kích thước của các đối tượng trong một span, nó chỉ cần làm tròn xuống theo kích thước đó
và đó sẽ là điểm bắt đầu của đối tượng.

Tất nhiên, span phân tách theo kích thước có một số ưu điểm khác.

Phân mảnh thấp: Kinh nghiệm với C, ngoài TCMalloc và Hoard của Google,
tôi đã gắn bó chặt chẽ với Scalable Malloc của Intel và công trình đó đã
cho chúng tôi sự tự tin rằng phân mảnh sẽ không phải là vấn đề với bộ cấp phát không di chuyển.

Cấu trúc nội tại: Chúng tôi hiểu đầy đủ và có kinh nghiệm với chúng.
Chúng tôi biết cách thực hiện span phân tách theo kích thước,
chúng tôi biết cách thực hiện các đường phân bổ với ít hoặc không có tranh chấp.

Tốc độ: Không sao chép không làm chúng tôi lo ngại, việc phân bổ có thể chậm hơn một chút
nhưng vẫn ở mức độ của C.
Có thể không nhanh như bump pointer nhưng điều đó ổn.

Chúng tôi cũng có vấn đề về giao tiếp hàm ngoại vi.
Nếu chúng tôi không di chuyển các đối tượng, chúng tôi không phải đối phó với
cái đuôi dài của các lỗi mà bạn có thể gặp phải nếu có bộ thu gom di chuyển
khi bạn cố gắng ghim các đối tượng và đặt các mức độ gián tiếp giữa C và đối tượng Go mà bạn đang làm việc.

{{image "ismmkeynote/image5.png"}}

Lựa chọn thiết kế tiếp theo là nơi đặt metadata của đối tượng.
Chúng tôi cần có một số thông tin về các đối tượng vì chúng tôi không có header.
Các bit đánh dấu được giữ ở phần bên cạnh và được sử dụng cho cả đánh dấu lẫn phân bổ.
Mỗi từ có 2 bit liên kết để cho biết nó là scalar hay con trỏ bên trong từ đó.
Nó cũng mã hóa liệu có thêm con trỏ trong đối tượng để chúng tôi có thể
dừng quét sớm hơn.
Chúng tôi cũng có thêm một bit mã hóa mà chúng tôi có thể sử dụng như một bit đánh dấu bổ sung
hoặc để gỡ lỗi.
Điều này thực sự có giá trị trong việc chạy được mã và tìm ra các lỗi.

{{image "ismmkeynote/image19.png"}}

Vậy còn write barrier? Write barrier chỉ bật trong quá trình GC.
Lúc khác, mã biên dịch tải một biến toàn cục và kiểm tra nó.
Vì GC thường tắt, phần cứng dự đoán chính xác để bỏ qua write barrier.
Khi chúng ta ở trong GC, biến đó khác,
và write barrier chịu trách nhiệm đảm bảo không có đối tượng nào có thể đạt đến bị mất
trong quá trình thao tác ba màu.

{{image "ismmkeynote/image50.png"}}

Phần còn lại của mã là GC Pacer.
Đây là một trong những công trình xuất sắc mà Austin đã làm.
Về cơ bản nó dựa trên vòng lặp phản hồi xác định thời điểm tốt nhất để bắt đầu chu kỳ GC.
Nếu hệ thống ở trạng thái ổn định và không có thay đổi pha,
việc đánh dấu sẽ kết thúc đúng lúc bộ nhớ sắp hết.

Điều đó có thể không xảy ra vì vậy Pacer cũng phải giám sát tiến trình đánh dấu
và đảm bảo việc phân bổ không vượt quá đánh dấu đồng thời.

Khi cần, Pacer làm chậm phân bổ trong khi tăng tốc đánh dấu.
Ở cấp độ cao, Pacer dừng Goroutine,
đang thực hiện nhiều phân bổ, và đặt nó vào công việc đánh dấu.
Lượng công việc tỷ lệ thuận với lượng phân bổ của Goroutine.
Điều này tăng tốc bộ gom rác trong khi làm chậm mutator.

Khi tất cả điều này hoàn thành, Pacer lấy những gì nó đã học được từ chu kỳ GC này
cũng như các chu kỳ trước và dự đoán khi nào nên bắt đầu GC tiếp theo.

Nó làm nhiều hơn thế nhưng đó là cách tiếp cận cơ bản.

Toán học hoàn toàn hấp dẫn, hãy liên hệ với tôi để xem tài liệu thiết kế.
Nếu bạn đang thực hiện GC đồng thời, bạn thực sự nên tìm hiểu toán học này
và xem nó có giống với toán học của bạn không.
Nếu bạn có bất kỳ đề xuất nào, hãy cho chúng tôi biết.

[\*Định thời gian bộ gom rác đồng thời Go 1.5](/s/go15gcpacing)
và [Đề xuất: Tách biệt mục tiêu kích thước heap mềm và cứng](https://github.com/golang/proposal/blob/master/design/14951-soft-heap-limit.md)

{{image "ismmkeynote/image40.png"}}

Đúng vậy, chúng tôi đã có nhiều thành công. Một Rick trẻ hơn và liều lĩnh hơn đã
muốn xăm một số đồ thị này lên vai, tôi tự hào về chúng đến vậy.

{{image "ismmkeynote/image20.png"}}

Đây là một loạt đồ thị được thực hiện cho một máy chủ production tại Twitter.
Tất nhiên chúng tôi không liên quan gì đến máy chủ production đó.
Brian Hatfield đã thực hiện các phép đo này và thú vị thay, đã tweet về chúng.

Trục Y là độ trễ GC tính bằng mili giây.
Trục X là thời gian. Mỗi điểm là thời gian dừng thế giới (stop the world pause)
trong quá trình GC đó.

Trong bản phát hành đầu tiên của chúng tôi vào tháng 8 năm 2015,
chúng tôi thấy sự giảm từ khoảng 300 đến 400 mili giây xuống còn 30 hoặc 40 mili giây.
Điều này tốt, tốt theo bậc độ lớn.

Chúng ta sẽ thay đổi trục Y một cách mạnh mẽ từ 0 đến 400 mili giây xuống còn 0 đến 50 mili giây.

{{image "ismmkeynote/image54.png"}}

Đây là 6 tháng sau. Sự cải thiện chủ yếu là do việc loại bỏ có hệ thống tất cả
những thứ O(heap) mà chúng tôi đã làm trong thời gian dừng thế giới.
Đây là cải thiện bậc độ lớn thứ hai khi chúng tôi giảm từ 40 mili giây xuống còn 4 hoặc 5.

{{image "ismmkeynote/image1.png"}}

Có một số lỗi mà chúng tôi phải sửa và chúng tôi đã làm điều này trong
bản phát hành nhỏ 1.6.3.
Điều này đã giảm độ trễ xuống dưới 10 mili giây, đúng với SLO của chúng tôi.

Chúng ta sắp thay đổi trục Y một lần nữa, lần này xuống còn 0 đến 5 mili giây.

{{image "ismmkeynote/image68.png"}}

Đây rồi, đây là tháng 8 năm 2016, một năm sau bản phát hành đầu tiên.
Một lần nữa chúng tôi tiếp tục loại bỏ các quá trình dừng thế giới O(kích thước heap).
Chúng ta đang nói về heap 18GB ở đây.
Chúng tôi có những heap lớn hơn nhiều và khi chúng tôi loại bỏ các dừng thế giới O(kích thước heap) này,
kích thước heap có thể rõ ràng tăng đáng kể mà không ảnh hưởng đến độ trễ.
Vì vậy, đây là một sự hỗ trợ nhỏ trong 1.7.

{{image "ismmkeynote/image58.png"}}

Bản phát hành tiếp theo vào tháng 3 năm 2017. Chúng tôi đã có lần cuối cùng giảm độ trễ lớn
do tìm ra cách tránh dừng thế giới để quét stack ở cuối chu kỳ GC.
Điều đó đưa chúng tôi vào phạm vi dưới mili giây.
Trục Y sắp thay đổi thành 1,5 mili giây và chúng ta thấy sự cải thiện bậc độ lớn thứ ba.

{{image "ismmkeynote/image45.png"}}

Bản phát hành tháng 8 năm 2017 có ít cải thiện.
Chúng tôi biết những gì đang gây ra các dừng còn lại.
Con số thì thầm SLO ở đây khoảng 100-200 micro giây và chúng tôi sẽ hướng tới đó.
Nếu bạn thấy bất kỳ điều gì trên vài trăm micro giây, chúng tôi thực sự muốn
nói chuyện với bạn và tìm hiểu xem nó có phù hợp với những gì chúng tôi biết
hay đó là điều gì đó mới mà chúng tôi chưa xem xét.
Dù sao, có vẻ như ít có yêu cầu về độ trễ thấp hơn.
Điều quan trọng cần lưu ý là các mức độ trễ này có thể xảy ra vì nhiều lý do
không liên quan đến GC và như câu nói "Bạn không cần nhanh hơn con gấu,
bạn chỉ cần nhanh hơn người đứng cạnh bạn."

Không có thay đổi đáng kể nào trong bản phát hành 1.10 tháng 2 năm 2018,
chỉ là một số dọn dẹp và xử lý các trường hợp biên.

{{image "ismmkeynote/image6.png"}}

Một năm mới và một SLO mới. Đây là SLO năm 2018 của chúng tôi.

Chúng tôi đã giảm tổng CPU xuống CPU sử dụng trong chu kỳ GC.

Heap vẫn ở mức 2x.

Chúng tôi hiện có mục tiêu 500 micro giây dừng thế giới mỗi chu kỳ GC. Có lẽ hơi thấp một chút.

Phân bổ sẽ tiếp tục tỷ lệ thuận với các GC assist.

Pacer đã tốt hơn nhiều nên chúng tôi mong muốn thấy các GC assist tối thiểu ở trạng thái ổn định.

Chúng tôi khá hài lòng với điều này. Một lần nữa, đây không phải là SLA mà là SLO,
vì vậy đó là mục tiêu, không phải thỏa thuận, vì chúng tôi không thể kiểm soát những thứ như OS.

{{image "ismmkeynote/image64.png"}}

Đó là phần tốt đẹp. Hãy chuyển sang nói về những thất bại của chúng tôi.
Đây là những vết sẹo của chúng tôi; chúng giống như hình xăm và mọi người đều có.
Dù sao, chúng đi kèm với những câu chuyện hay hơn, vì vậy hãy cùng kể một số câu chuyện đó.

{{image "ismmkeynote/image46.png"}}

Nỗ lực đầu tiên của chúng tôi là thực hiện một thứ gọi là request oriented collector (ROC). Giả thuyết có thể thấy ở đây.

{{image "ismmkeynote/image34.png"}}

Điều này có nghĩa là gì?

Goroutine là các luồng nhẹ trông giống như Gopher,
vì vậy ở đây chúng ta có hai Goroutine.
Chúng chia sẻ một số thứ như hai đối tượng màu xanh ở giữa.
Chúng có stack riêng và bộ đối tượng riêng của mình.
Giả sử bạn bên trái muốn chia sẻ đối tượng màu xanh lá.

{{image "ismmkeynote/image9.png"}}

Goroutine đặt nó vào vùng chia sẻ để Goroutine kia có thể truy cập nó.
Họ có thể gắn nó vào một thứ gì đó trong heap chia sẻ hoặc gán nó vào một biến toàn cục
và Goroutine kia có thể thấy nó.

{{image "ismmkeynote/image26.png"}}

Cuối cùng, Goroutine bên trái đến giường chết, nó sắp chết, buồn quá.

{{image "ismmkeynote/image14.png"}}

Như bạn biết, bạn không thể mang theo đối tượng của mình khi chết.
Bạn cũng không thể mang theo stack. Stack thực ra trống lúc này
và các đối tượng không thể tiếp cận được, vì vậy bạn có thể đơn giản thu hồi chúng.

{{image "ismmkeynote/image2.png"}}

Điều quan trọng ở đây là tất cả các hành động đều là cục bộ và không đòi hỏi
bất kỳ sự đồng bộ hóa toàn cục nào.
Điều này về cơ bản khác so với các cách tiếp cận như GC thế hệ,
và hy vọng là sự mở rộng chúng tôi sẽ đạt được từ việc không phải thực hiện đồng bộ hóa đó
sẽ đủ để chúng tôi có lợi thế.

{{image "ismmkeynote/image27.png"}}

Vấn đề khác đang xảy ra với hệ thống này là write barrier luôn bật.
Bất cứ khi nào có một lần ghi, chúng tôi phải xem liệu nó có đang ghi con trỏ
đến một đối tượng riêng tư vào một đối tượng công khai không.
Nếu vậy, chúng tôi phải công khai đối tượng tham chiếu và sau đó thực hiện
một lần duyệt bắc cầu của các đối tượng có thể tiếp cận để đảm bảo chúng cũng công khai.
Đó là một write barrier khá đắt đỏ có thể gây ra nhiều lần trượt bộ nhớ đệm.

{{image "ismmkeynote/image30.png"}}

Dù vậy, chúng tôi đã có một số thành công khá tốt.

Đây là một benchmark RPC đầu cuối. Trục Y bị gán nhãn sai đi từ 0 đến
5 mili giây (thấp hơn là tốt hơn),
dù sao thì đó chỉ là vậy.
Trục X về cơ bản là ballast hay kích thước cơ sở dữ liệu trong bộ nhớ.

Như bạn có thể thấy, nếu bạn bật ROC và không có nhiều chia sẻ,
mọi thứ thực sự mở rộng khá tốt.
Nếu bạn không bật ROC thì không tốt bằng.

{{image "ismmkeynote/image35.png"}}

Nhưng điều đó chưa đủ tốt, chúng tôi cũng phải đảm bảo ROC không làm chậm
các phần khác của hệ thống.
Lúc đó có nhiều lo ngại về trình biên dịch và chúng tôi không thể làm chậm trình biên dịch.
Thật không may, trình biên dịch chính xác là những chương trình mà ROC không làm tốt.
Chúng tôi thấy 30, 40, 50% và hơn nữa là chậm lại và điều đó không thể chấp nhận được.
Go tự hào về việc trình biên dịch của nó nhanh như thế nào, vì vậy chúng tôi không thể
làm chậm trình biên dịch, chắc chắn không nhiều đến vậy.

{{image "ismmkeynote/image61.png"}}

Sau đó chúng tôi xem xét một số chương trình khác.
Đây là các benchmark hiệu suất của chúng tôi. Chúng tôi có kho khoảng 200 hoặc 300 benchmark
và đây là những cái mà nhóm trình biên dịch đã quyết định là quan trọng
để họ làm việc và cải thiện.
Những cái này không được chọn bởi nhóm GC chút nào.
Các con số đồng loạt xấu và ROC sẽ không trở thành người chiến thắng.

{{image "ismmkeynote/image44.png"}}

Đúng là chúng tôi đã mở rộng nhưng chúng tôi chỉ có hệ thống 4 đến 12 luồng phần cứng
nên chúng tôi không thể vượt qua thuế write barrier.
Có lẽ trong tương lai khi chúng ta có hệ thống 128 lõi và Go tận dụng chúng,
các thuộc tính mở rộng của ROC có thể là một chiến thắng.
Khi điều đó xảy ra, chúng tôi có thể quay lại và xem xét lại điều này,
nhưng hiện tại ROC là một đề xuất thua lỗ.

{{image "ismmkeynote/image66.png"}}

Vậy chúng tôi sẽ làm gì tiếp theo? Hãy thử GC thế hệ (generational GC).
Đây là phương pháp cũ nhưng tốt. ROC không hiệu quả nên hãy quay lại những thứ
chúng tôi có nhiều kinh nghiệm hơn.

{{image "ismmkeynote/image41.png"}}

Chúng tôi sẽ không từ bỏ độ trễ, chúng tôi sẽ không từ bỏ việc không di chuyển.
Vì vậy, chúng tôi cần một GC thế hệ không di chuyển.

{{image "ismmkeynote/image27.png"}}

Chúng ta có thể làm điều này không? Có, nhưng với GC thế hệ,
write barrier luôn bật.
Khi chu kỳ GC đang chạy, chúng tôi sử dụng cùng write barrier như hiện nay,
nhưng khi GC tắt, chúng tôi sử dụng write barrier GC nhanh đệm các con trỏ
và sau đó xả bộ đệm vào bảng card mark khi tràn.

{{image "ismmkeynote/image4.png"}}

Vậy điều này sẽ hoạt động như thế nào trong tình huống không di chuyển? Đây là bản đồ đánh dấu/phân bổ.
Về cơ bản bạn duy trì một con trỏ hiện tại.
Khi bạn phân bổ, bạn tìm số không tiếp theo và khi bạn tìm thấy số không đó
bạn phân bổ một đối tượng trong khoảng trống đó.

{{image "ismmkeynote/image51.png"}}

Bạn sau đó cập nhật con trỏ hiện tại đến số 0 tiếp theo.

{{image "ismmkeynote/image17.png"}}

Bạn tiếp tục cho đến khi đến lúc thực hiện generation GC.
Bạn sẽ nhận thấy rằng nếu có số 1 trong vector đánh dấu/phân bổ thì
đối tượng đó đã sống từ lần GC cuối nên nó đã trưởng thành.
Nếu nó là số 0 và bạn tiếp cận được nó thì bạn biết nó còn trẻ.

{{image "ismmkeynote/image53.png"}}

Vậy bạn thực hiện thăng cấp như thế nào. Nếu bạn tìm thấy thứ gì đó được đánh dấu bằng 1
trỏ đến thứ gì đó được đánh dấu bằng 0 thì bạn thăng cấp đối tượng tham chiếu
đơn giản bằng cách đặt số 0 đó thành 1.

{{image "ismmkeynote/image49.png"}}

Bạn phải thực hiện một lần duyệt bắc cầu để đảm bảo tất cả các đối tượng có thể tiếp cận đều được thăng cấp.

{{image "ismmkeynote/image69.png"}}

Khi tất cả các đối tượng có thể tiếp cận đã được thăng cấp, GC nhỏ kết thúc.

{{image "ismmkeynote/image62.png"}}

Cuối cùng, để hoàn thành chu kỳ GC thế hệ, bạn chỉ cần đặt lại
con trỏ hiện tại về đầu vector và bạn có thể tiếp tục.
Tất cả các số không không được tiếp cận trong chu kỳ GC đó đều tự do và có thể được tái sử dụng.
Như nhiều người trong các bạn biết, điều này được gọi là 'sticky bits' và được phát minh bởi Hans
Boehm và các đồng nghiệp của ông.

{{image "ismmkeynote/image21.png"}}

Vậy hiệu suất trông như thế nào? Không tệ với các heap lớn.
Đây là các benchmark mà GC nên làm tốt. Tất cả đều ổn.

{{image "ismmkeynote/image65.png"}}

Sau đó chúng tôi chạy nó trên các benchmark hiệu suất và mọi thứ không tiến triển tốt. Vậy điều gì đang xảy ra?

{{image "ismmkeynote/image43.png"}}

Write barrier nhanh nhưng đơn giản là không đủ nhanh.
Hơn nữa, việc tối ưu hóa nó khó khăn. Ví dụ,
write barrier elision có thể xảy ra nếu có một lần ghi khởi tạo giữa
lúc đối tượng được phân bổ và safepoint tiếp theo.
Nhưng chúng tôi phải chuyển sang hệ thống có GC safepoint tại mỗi lệnh
nên thực sự không có write barrier nào chúng tôi có thể loại bỏ trong tương lai.

{{image "ismmkeynote/image47.png"}}

Chúng tôi cũng có escape analysis và nó ngày càng tốt hơn.
Hãy nhớ những thứ định hướng giá trị mà chúng ta đã nói? Thay vì truyền
con trỏ đến một hàm, chúng ta sẽ truyền giá trị thực tế.
Vì chúng ta đang truyền giá trị, escape analysis chỉ phải thực hiện
phân tích thoát trong nội bộ hàm (intraprocedural), không phải liên thủ tục (interprocedural).

Tất nhiên trong trường hợp con trỏ đến đối tượng cục bộ thoát ra, đối tượng sẽ được phân bổ trên heap.

Không phải là giả thuyết thế hệ không đúng với Go,
mà là các đối tượng trẻ sống và chết trẻ trên stack.
Kết quả là việc thu gom thế hệ kém hiệu quả hơn nhiều so với những gì
bạn có thể tìm thấy trong các ngôn ngữ runtime được quản lý khác.

{{image "ismmkeynote/image10.png"}}

Vì vậy, những lực lượng chống lại write barrier này bắt đầu tập hợp.
Ngày nay, trình biên dịch của chúng tôi tốt hơn nhiều so với năm 2014.
Escape analysis đang nhặt nhiều đối tượng đó và đặt chúng trên
stack, những đối tượng mà bộ thu thế hệ sẽ giúp ích.
Chúng tôi bắt đầu tạo ra các công cụ giúp người dùng tìm các đối tượng đã thoát và
nếu nó nhỏ, họ có thể thay đổi mã và giúp trình biên dịch phân bổ trên stack.

Người dùng ngày càng thông minh hơn trong việc áp dụng các phương pháp định hướng giá trị
và số lượng con trỏ đang giảm.
Mảng và map chứa các giá trị chứ không phải con trỏ đến struct. Mọi thứ đều tốt.

Nhưng đó không phải là lý do chính thuyết phục tại sao write barrier trong Go có cuộc chiến khó khăn phía trước.

{{image "ismmkeynote/image8.png"}}

Hãy nhìn vào đồ thị này. Đây chỉ là đồ thị phân tích về chi phí đánh dấu.
Mỗi đường biểu diễn một ứng dụng khác nhau có thể có chi phí đánh dấu.
Giả sử chi phí đánh dấu của bạn là 20%, khá cao nhưng có thể.
Đường đỏ là 10%, vẫn cao.
Đường dưới là 5%, khoảng chi phí của một write barrier hiện nay.
Vậy điều gì xảy ra nếu bạn tăng gấp đôi kích thước heap? Đó là điểm bên phải.
Chi phí tích lũy của giai đoạn đánh dấu giảm đáng kể vì các chu kỳ GC ít thường xuyên hơn.
Chi phí write barrier là hằng số nên chi phí của việc tăng kích thước heap
sẽ đưa chi phí đánh dấu đó xuống dưới chi phí write barrier.

{{image "ismmkeynote/image39.png"}}

Đây là chi phí phổ biến hơn cho write barrier,
là 4%, và chúng tôi thấy rằng ngay cả với điều đó, chúng tôi có thể đưa chi phí
đánh dấu xuống dưới chi phí write barrier chỉ bằng cách tăng kích thước heap.

Giá trị thực sự của GC thế hệ là,
khi nhìn vào thời gian GC, chi phí write barrier bị bỏ qua vì chúng
được phết đều lên mutator.
Đây là lợi thế lớn của GC thế hệ,
nó giảm đáng kể các lần STW dài trong các chu kỳ GC đầy đủ nhưng không nhất thiết cải thiện thông lượng.
Go không có vấn đề dừng thế giới này nên nó phải nhìn kỹ hơn
vào các vấn đề thông lượng và đó là những gì chúng tôi đã làm.

{{image "ismmkeynote/image23.png"}}

Đó là rất nhiều thất bại và với những thất bại như vậy đến cả thức ăn và bữa trưa.
Tôi đang than vãn như thường lệ "Ồ sẽ tuyệt vời biết bao nếu không có write barrier."

Trong khi đó Austin vừa dành một giờ nói chuyện với một số chuyên gia HW GC
tại Google và anh ấy đang nói chúng tôi nên nói chuyện với họ và cố gắng tìm hiểu
cách nhận được hỗ trợ HW GC có thể giúp ích.
Sau đó tôi bắt đầu kể những câu chuyện chiến trận về dòng cache lấp đầy không (zero-fill cache lines),
các chuỗi nguyên tử có thể khởi động lại, và những thứ khác không thành công khi tôi
làm việc cho một công ty phần cứng lớn.
Đúng là chúng tôi đã đưa được một số thứ vào con chip gọi là Itanium,
nhưng chúng tôi không thể đưa chúng vào các chip phổ biến hơn hiện nay.
Vì vậy, bài học đạo đức đơn giản là chỉ cần sử dụng phần cứng chúng ta có.

Dù sao điều đó đã khiến chúng tôi nói chuyện, còn điều gì điên rồ không?

{{image "ismmkeynote/image25.png"}}

Điều gì về việc đánh dấu card mà không cần write barrier? Hóa ra Austin
có những tệp này và anh ấy viết vào những tệp này tất cả những ý tưởng điên rồ của mình
mà vì lý do nào đó anh ấy không nói với tôi.
Tôi cho rằng đó là một loại liệu pháp tâm lý.
Tôi đã từng làm điều tương tự với Eliot. Các ý tưởng mới dễ bị đập vỡ và
người ta cần bảo vệ chúng và làm cho chúng mạnh mẽ hơn trước khi để chúng ra thế giới.
Dù sao anh ấy đã lấy ý tưởng này ra.

Ý tưởng là bạn duy trì một hash của các con trỏ trưởng thành trong mỗi card.
Nếu các con trỏ được ghi vào card, hash sẽ thay đổi và card sẽ
được coi là đã đánh dấu.
Điều này sẽ đánh đổi chi phí write barrier để lấy chi phí hashing.

{{image "ismmkeynote/image31.png"}}

Nhưng quan trọng hơn, nó phù hợp với phần cứng.

Các kiến trúc hiện đại ngày nay có lệnh AES (Advanced Encryption Standard).
Một trong những lệnh đó có thể thực hiện hashing cấp độ mã hóa và với hashing cấp độ mã hóa
chúng ta không phải lo lắng về va chạm nếu chúng ta cũng tuân theo các chính sách mã hóa tiêu chuẩn.
Vì vậy, hashing sẽ không tốn kém nhiều nhưng chúng ta phải tải những gì chúng ta sẽ hash.
May mắn thay, chúng ta đang duyệt qua bộ nhớ tuần tự nên chúng ta đạt được
hiệu suất bộ nhớ và bộ nhớ đệm rất tốt.
Nếu bạn có DIMM và bạn truy cập các địa chỉ tuần tự,
thì đó là một chiến thắng vì chúng sẽ nhanh hơn so với truy cập các địa chỉ ngẫu nhiên.
Các bộ tiền nạp phần cứng sẽ hoạt động và điều đó cũng sẽ giúp ích.
Dù sao chúng ta có 50 năm, 60 năm thiết kế phần cứng để chạy Fortran,
để chạy C và để chạy các benchmark SPECint.
Không có gì ngạc nhiên khi kết quả là phần cứng chạy loại công việc này nhanh.

{{image "ismmkeynote/image12.png"}}

Chúng tôi đã thực hiện phép đo. Khá tốt. Đây là bộ benchmark cho heap lớn, đáng lẽ phải tốt.

{{image "ismmkeynote/image18.png"}}

Sau đó chúng tôi nói điều đó trông như thế nào với benchmark hiệu suất? Không tốt lắm,
có một số ngoại lệ.
Nhưng bây giờ chúng tôi đã chuyển write barrier từ luôn bật trong mutator
sang chạy như một phần của chu kỳ GC.
Bây giờ việc quyết định xem chúng tôi có thực hiện GC thế hệ hay không
bị trì hoãn đến đầu chu kỳ GC.
Chúng tôi có nhiều quyền kiểm soát hơn ở đó vì chúng tôi đã cục bộ hóa công việc card.
Bây giờ chúng tôi có các công cụ, chúng tôi có thể giao nó cho Pacer,
và nó có thể làm tốt việc cắt xén động các chương trình rơi sang phải
và không được hưởng lợi từ GC thế hệ.
Nhưng điều này có chiến thắng trong tương lai không? Chúng tôi cần biết hoặc ít nhất nghĩ đến
phần cứng sẽ trông như thế nào trong tương lai.

{{image "ismmkeynote/image52.png"}}

Bộ nhớ của tương lai là gì?

{{image "ismmkeynote/image11.png"}}

Hãy nhìn vào đồ thị này. Đây là đồ thị định luật Moore cổ điển của bạn.
Bạn có thang log trên trục Y cho thấy số lượng transistor trong một chip.
Trục X là các năm từ 1971 đến 2016.
Tôi sẽ lưu ý rằng đây là những năm mà ai đó ở đâu đó đã dự đoán rằng
định luật Moore đã chết.

Định luật Dennard đã kết thúc sự cải thiện tần số khoảng mười năm trước.
Các quy trình mới đang mất nhiều thời gian hơn để tăng tốc. Vì vậy, thay vì 2 năm, chúng
bây giờ là 4 năm hoặc hơn.
Vì vậy, khá rõ ràng rằng chúng ta đang bước vào kỷ nguyên mà định luật Moore chậm lại.

Hãy chỉ nhìn vào các chip trong vòng tròn đỏ. Đây là những chip tốt nhất trong việc duy trì định luật Moore.

Chúng là những chip trong đó logic ngày càng đơn giản và được sao chép nhiều lần.
Nhiều lõi giống nhau, nhiều bộ điều khiển bộ nhớ và bộ nhớ đệm,
GPU, TPU, và vân vân.

Khi chúng ta tiếp tục đơn giản hóa và tăng sự sao chép, chúng ta tiệm cận
kết thúc với một vài dây,
một transistor và một tụ điện.
Nói cách khác là một ô nhớ DRAM.

Nói theo cách khác, chúng tôi nghĩ rằng việc tăng gấp đôi bộ nhớ sẽ là giá trị tốt hơn so với tăng gấp đôi lõi.

[Đồ thị gốc](http://www.kurzweilai.net/ask-ray-the-future-of-moores-law)
tại www.kurzweilai.net/ask-ray-the-future-of-moores-law.

{{image "ismmkeynote/image57.png"}}

Hãy xem một đồ thị khác tập trung vào DRAM.
Đây là các con số từ một luận án tiến sĩ gần đây từ CMU.
Nếu chúng ta nhìn vào điều này, chúng ta thấy rằng định luật Moore là đường màu xanh.
Đường màu đỏ là dung lượng và có vẻ như đang theo định luật Moore.
Thú vị thay, tôi đã thấy một đồ thị quay ngược đến năm 1939 khi chúng ta đang
sử dụng bộ nhớ trống và dung lượng đó cùng định luật Moore đang tiến lên cùng nhau,
vì vậy đồ thị này đã tiếp tục trong một thời gian dài,
chắc chắn lâu hơn có lẽ bất kỳ ai trong phòng này đã sống.

Nếu chúng ta so sánh đồ thị này với tần số CPU hoặc các đồ thị Moore's-law-is-dead khác nhau,
chúng ta được dẫn đến kết luận rằng bộ nhớ,
hoặc ít nhất là dung lượng chip, sẽ theo định luật Moore lâu hơn CPU.
Băng thông, đường màu vàng, liên quan không chỉ đến tần số
bộ nhớ mà còn đến số lượng chân có thể đưa ra khỏi chip nên nó
không theo kịp tốt nhưng không làm tệ.

Độ trễ, đường màu xanh lá, đang làm rất kém,
mặc dù tôi sẽ lưu ý rằng độ trễ cho các lần truy cập tuần tự tốt hơn
độ trễ cho truy cập ngẫu nhiên.

(Dữ liệu từ "Understanding and Improving the Latency of DRAM-Based Memory
Systems Submitted in partial fulfillment of the requirements for the degree
of Doctor of Philosophy in Electrical and Computer Engineering Kevin K.
Chang M.S., Electrical & Computer Engineering,
Carnegie Mellon University B.S., Electrical & Computer Engineering,
Carnegie Mellon University Carnegie Mellon University Pittsburgh, PA May, 2017".
Xem [luận án của Kevin K. Chang.](http://repository.cmu.edu/cgi/viewcontent.cgi?article%3D1946%26context%3Ddissertations&amp;sa=D&amp;ust=1531164842660000)
Đồ thị gốc trong phần giới thiệu không ở dạng tôi có thể vẽ
đường định luật Moore lên nó dễ dàng nên tôi đã thay đổi trục X để đồng đều hơn.)

{{image "ismmkeynote/image15.png"}}

Hãy đến nơi cao su gặp mặt đường.
Đây là giá DRAM thực tế và nó nhìn chung đã giảm từ 2005 đến 2016.
Tôi chọn 2005 vì đó là khoảng thời gian mà định luật Dennard kết thúc và
cùng với nó là sự cải thiện tần số.

Nếu bạn nhìn vào vòng tròn đỏ, về cơ bản là khoảng thời gian công việc giảm
độ trễ GC của Go đang diễn ra,
chúng ta thấy rằng trong vài năm đầu giá đã ổn.
Gần đây không tốt như vậy, khi nhu cầu vượt cung dẫn đến tăng giá trong hai năm qua.
Tất nhiên, transistor chưa trở nên lớn hơn và trong một số trường hợp dung lượng chip
đã tăng nên điều này được thúc đẩy bởi lực lượng thị trường.
RAMBUS và các nhà sản xuất chip khác nói rằng tiến lên phía trước, chúng ta sẽ thấy
quá trình thu nhỏ tiếp theo trong khung thời gian 2019-2020.

Tôi sẽ kiềm chế không suy đoán về các lực lượng thị trường toàn cầu trong ngành công nghiệp bộ nhớ
ngoài việc lưu ý rằng giá cả có chu kỳ và về lâu dài nguồn cung có xu hướng đáp ứng nhu cầu.

Về lâu dài, chúng tôi tin rằng giá bộ nhớ sẽ giảm với tốc độ nhanh hơn nhiều so với giá CPU.

(Nguồn [https://hblok.net/blog/](https://hblok.net/blog/) và [https://hblok.net/storage\_data/storage\_memory\_prices\_2005-2017-12.png](https://hblok.net/storage_data/storage_memory_prices_2005-2017-12.png))

{{image "ismmkeynote/image37.png"}}

Hãy nhìn vào đường khác này. Ồ sẽ tốt biết bao nếu chúng ta ở trên đường này.
Đây là đường SSD. Nó đang làm tốt hơn trong việc giữ giá thấp.
Vật lý vật liệu của các chip này phức tạp hơn nhiều so với DRAM.
Logic phức tạp hơn, thay vì một transistor mỗi ô, có khoảng nửa tá.

Nhìn về phía trước có một đường giữa DRAM và SSD nơi NVRAM như Intel
3D XPoint và Phase Change Memory (PCM) sẽ tồn tại.
Trong thập kỷ tới, sự tăng cường cung cấp của loại bộ nhớ này có khả năng
trở nên phổ biến hơn và điều này sẽ chỉ củng cố ý tưởng rằng thêm
bộ nhớ là cách rẻ tiền để tăng thêm giá trị cho máy chủ của chúng ta.

Quan trọng hơn, chúng ta có thể mong đợi thấy các giải pháp thay thế cạnh tranh khác với DRAM.
Tôi sẽ không giả vờ biết cái nào sẽ được ưa thích trong năm hay mười năm nữa nhưng
sự cạnh tranh sẽ gay gắt và bộ nhớ heap sẽ tiến gần hơn đến đường SSD màu xanh nổi bật ở đây.

Tất cả điều này củng cố quyết định của chúng tôi về việc tránh các barrier luôn bật để ủng hộ việc tăng bộ nhớ.

{{image "ismmkeynote/image16.png"}}

Vậy tất cả điều này có nghĩa gì với Go trong tương lai?

{{image "ismmkeynote/image42.png"}}

Chúng tôi có ý định làm cho runtime linh hoạt và mạnh mẽ hơn khi chúng tôi xem xét
các trường hợp biên đến từ người dùng.
Hy vọng là thắt chặt bộ lên lịch và có được tính xác định và công bằng tốt hơn
nhưng chúng tôi không muốn hy sinh bất kỳ hiệu suất nào.

Chúng tôi cũng không có ý định tăng bề mặt API GC.
Chúng tôi đã có gần một thập kỷ và chúng tôi có hai núm điều chỉnh và điều đó cảm thấy phù hợp.
Không có ứng dụng nào quan trọng đến mức chúng tôi cần thêm một cờ mới.

Chúng tôi cũng sẽ xem xét cách cải thiện escape analysis vốn đã khá tốt của chúng tôi
và tối ưu hóa cho lập trình định hướng giá trị của Go.
Không chỉ trong lập trình mà còn trong các công cụ chúng tôi cung cấp cho người dùng.

Về mặt thuật toán, chúng tôi sẽ tập trung vào các phần của không gian thiết kế
giảm thiểu việc sử dụng barrier,
đặc biệt là những barrier luôn bật.

Cuối cùng, và quan trọng nhất, chúng tôi hy vọng tận dụng xu hướng của định luật Moore
ủng hộ RAM hơn CPU chắc chắn trong 5 năm tới và hy vọng trong thập kỷ tới.

Vậy là hết. Cảm ơn.

{{image "ismmkeynote/image33.png"}}

P.S. Nhóm Go đang tìm kiếm các kỹ sư để giúp phát triển và duy trì bộ công cụ runtime và trình biên dịch của Go.

Quan tâm? Hãy xem các [vị trí tuyển dụng](https://go-jobs-at-goog.firebaseapp.com) của chúng tôi.
