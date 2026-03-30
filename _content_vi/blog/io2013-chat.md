---
title: Cuộc trò chuyện với nhóm Go
date: 2013-06-06
summary: Tại Google I/O 2013, một số thành viên của nhóm Go đã tổ chức một buổi "Fireside chat."
template: true
---


Tại Google I/O 2013, một số thành viên của nhóm Go đã tổ chức một buổi "Fireside chat."
Robert Griesemer, Rob Pike, David Symonds, Andrew Gerrand, Ian Lance Taylor,
Sameer Ajmani, Brad Fitzpatrick và Nigel Tao đã nhận câu hỏi từ khán giả
và mọi người trên khắp thế giới về các khía cạnh khác nhau của dự án Go.

{{video "https://www.youtube.com/embed/p9VUCp98ay4"}}

Chúng tôi cũng đã tổ chức một phiên tương tự tại I/O năm ngoái:
[_Meet the Go team_](http://www.youtube.com/watch?v=sln-gJaURzk).

Có nhiều câu hỏi hơn từ Google Moderator so với những gì chúng tôi có thể
trả lời trong phiên ngắn 40 phút.
Dưới đây chúng tôi trả lời một số câu hỏi mà chúng tôi đã bỏ qua trong phiên trực tiếp.

_Tốc độ liên kết (và mức sử dụng bộ nhớ) cho chuỗi công cụ gc là một vấn đề đã biết._
_Có kế hoạch nào để giải quyết vấn đề này trong chu kỳ 1.2 không?_

**Rob:** Có. Chúng tôi luôn suy nghĩ về cách cải thiện hiệu năng của
các công cụ cũng như ngôn ngữ và thư viện.

_Tôi rất vui khi thấy Go dường như đang nhanh chóng được chấp nhận._
_Bạn có thể chia sẻ về những phản ứng bạn đã trải qua khi làm việc với_
_các nhà phát triển khác trong và ngoài Google không? Còn điểm khó khăn lớn nào không?_

**Robert:** Nhiều nhà phát triển đã thực sự thử Go đều rất hài lòng với nó.
Nhiều người báo cáo một codebase nhỏ hơn, dễ đọc hơn và do đó dễ bảo trì hơn:
Giảm 50% kích thước code hoặc hơn khi chuyển từ C++ là điều thường gặp.
Các nhà phát triển chuyển sang Go từ Python đều đồng thuận hài lòng với
mức cải thiện hiệu năng. Những phàn nàn thường gặp là về các điểm không nhất quán nhỏ trong
ngôn ngữ (một số trong đó chúng tôi có thể giải quyết vào một lúc nào đó). Điều ngạc nhiên với tôi là
hầu như không ai phàn nàn về việc thiếu generics.

_Go sẽ trở thành ngôn ngữ bậc nhất cho phát triển Android khi nào?_

**Andrew:** Điều này sẽ rất tuyệt, nhưng chúng tôi không có gì để thông báo.

_Có lộ trình phát triển cho phiên bản Go tiếp theo không?_

**Andrew:** Chúng tôi không có lộ trình tính năng như vậy. Những người đóng góp thường làm việc trên
những thứ họ quan tâm. Các lĩnh vực đang được phát triển tích cực bao gồm trình biên dịch gc và gccgo,
bộ gom rác và runtime, và nhiều thứ khác. Chúng tôi kỳ vọng phần lớn những bổ sung thú vị mới
sẽ ở dạng cải tiến cho các công cụ của chúng tôi. Bạn có thể tìm thấy thảo luận thiết kế và code review trên
[danh sách gửi thư golang-dev](http://groups.google.com/group/golang-dev).

Về mốc thời gian, chúng tôi có
[kế hoạch cụ thể](https://docs.google.com/document/d/106hMEZj58L9nq9N9p7Zll_WKfo-oyZHFyI6MttuZmBU/edit?usp=sharing):
chúng tôi dự kiến phát hành Go 1.2 vào ngày 1 tháng 12 năm 2013.

_Các bạn muốn thấy Go được sử dụng ở đâu bên ngoài?_
_Điều gì sẽ được coi là một chiến thắng lớn cho Go khi được chấp nhận bên ngoài Google?_
_Go có tiềm năng tạo ra tác động đáng kể ở lĩnh vực nào?_

**Rob:** Nơi Go được triển khai là do người dùng quyết định, không phải chúng tôi. Chúng tôi vui mừng thấy
nó được chấp nhận ở bất kỳ nơi nào nó giúp ích. Nó được thiết kế với phần mềm phía máy chủ
trong tâm trí, và đang cho thấy tiềm năng ở đó, nhưng cũng đã thể hiện thế mạnh trong nhiều
lĩnh vực khác và câu chuyện thực sự chỉ mới bắt đầu. Còn nhiều bất ngờ phía trước.

**Ian:** Các startup dễ dùng Go hơn, vì họ không có một codebase đã ăn sâu mà họ cần làm việc cùng. Vì vậy tôi thấy hai chiến thắng lớn trong tương lai
cho Go. Một là Go được một công ty phần mềm lớn hiện có ngoài Google sử dụng đáng kể. Một khác là một IPO hoặc thương vụ mua lại đáng kể
của một startup chủ yếu sử dụng Go. Cả hai đều gián tiếp: rõ ràng
việc lựa chọn ngôn ngữ lập trình là một yếu tố rất nhỏ trong sự thành công của một công ty.
Nhưng đó sẽ là một cách khác để chứng minh rằng Go có thể là một phần của hệ thống phần mềm thành công.

_Bạn có nghĩ (thêm) về tiềm năng của việc tải động_
_các gói hoặc đối tượng Go và cách nó có thể hoạt động trong Go không?_
_Tôi nghĩ điều này có thể cho phép một số cấu trúc thực sự thú vị và biểu đạt,_
_đặc biệt kết hợp với các interface._

**Rob:** Đây là chủ đề đang được thảo luận tích cực. Chúng tôi hiểu khái niệm này có thể mạnh mẽ đến mức nào và hy vọng có thể tìm ra cách triển khai nó trong thời gian không quá lâu.
Có những thách thức nghiêm túc trong cách tiếp cận thiết kế và nhu cầu làm cho nó hoạt động
đa nền tảng.

_Đã có một thảo luận không lâu trước đây về việc tập hợp một số_
driver `database/sql` _tốt nhất vào một nơi trung tâm hơn._
_Một số người có quan điểm trái chiều mạnh mẽ._
_`database/sql` _và các driver của nó sẽ đi về đâu trong năm tới?_

**Brad:** Dù chúng tôi có thể tạo một subrepo chính thức ("go.db") cho các
driver cơ sở dữ liệu, chúng tôi lo ngại điều đó sẽ vô tình đặc cách cho một số driver. Ở thời điểm này chúng tôi
vẫn muốn thấy sự cạnh tranh lành mạnh giữa các driver khác nhau. Trang
[wiki SQLDrivers](/wiki/SQLDrivers)
liệt kê một số driver tốt.

Gói `database/sql` không được chú ý nhiều trong một thời gian, do thiếu
driver. Nay đã có driver, việc sử dụng gói đang tăng lên và các lỗi
về tính đúng đắn và hiệu năng đang được báo cáo (và sửa). Các sửa lỗi sẽ
tiếp tục, nhưng không có thay đổi lớn nào cho giao diện của `database/sql` được lên kế hoạch.
Có thể có các mở rộng nhỏ đây đó khi cần thiết cho hiệu năng hoặc để
hỗ trợ một số driver.

_Trạng thái của versioning là gì?_
_Việc import code từ GitHub có phải là best practice được nhóm Go khuyến nghị không?_
_Điều gì xảy ra khi chúng tôi xuất bản code phụ thuộc vào một repo GitHub và_
_API của bên phụ thuộc thay đổi?_

**Ian:** Điều này thường xuyên được thảo luận trên danh sách gửi thư. Những gì chúng tôi làm nội bộ
là chụp snapshot của code được import, và cập nhật snapshot đó theo thời gian. Theo cách đó, codebase
của chúng tôi sẽ không bị hỏng bất ngờ nếu API thay đổi.
Nhưng chúng tôi hiểu rằng cách tiếp cận đó không hoạt động tốt lắm với những người
tự thân đang cung cấp một thư viện. Chúng tôi chào đón các gợi ý tốt trong lĩnh vực này.
Hãy nhớ rằng đây là một khía cạnh của các công cụ xung quanh ngôn ngữ chứ không phải
bản thân ngôn ngữ; nơi cần sửa là ở các công cụ, không phải ngôn ngữ.

_Go và Giao diện người dùng đồ họa thì sao?_

**Rob:** Đây là chủ đề gần gũi với trái tim tôi. Newsqueak, một ngôn ngữ tiền thân rất sớm,
được thiết kế đặc biệt để viết các chương trình đồ họa (đó là thứ chúng tôi từng gọi là app). Bức tranh đã thay đổi nhiều nhưng tôi nghĩ mô hình concurrency của Go có nhiều điều hữu ích để cung cấp trong lĩnh vực đồ họa tương tác.

**Andrew:** Có nhiều
[binding cho các thư viện đồ họa hiện có](/wiki/Projects#Graphics_and_Audio)
ngoài kia, và một số dự án dành riêng cho Go. Một trong những cái hứa hẹn hơn là
[go.uik](https://github.com/skelterjohn/go.uik), nhưng nó vẫn còn trong giai đoạn đầu.
Tôi nghĩ có rất nhiều tiềm năng cho một bộ công cụ UI tuyệt vời dành riêng cho Go để
viết ứng dụng native (hãy nghĩ đến việc xử lý sự kiện người dùng bằng cách nhận từ một
channel), nhưng phát triển một gói chất lượng production là một công việc đáng kể. Tôi không nghi ngờ sẽ có một cái vào thời điểm thích hợp.

Trong thời gian đó, web là nền tảng rộng rãi nhất cho giao diện người dùng.
Go cung cấp hỗ trợ tuyệt vời để xây dựng ứng dụng web, dù chỉ ở phần backend.

_Trong danh sách gửi thư, Adam Langley đã tuyên bố rằng code TLS chưa được_
_các nhóm bên ngoài đánh giá, và do đó không nên sử dụng trong production._
_Có kế hoạch nào để xem xét code không?_
_Một triển khai TLS đồng thời an toàn tốt sẽ rất tuyệt._

**Adam**: Mật mã học nổi tiếng dễ mắc lỗi theo những cách tinh tế và bất ngờ
và tôi chỉ là người thường. Tôi không cảm thấy có thể đảm bảo rằng code TLS của Go hoàn toàn không có lỗi và tôi không muốn đánh giá sai về nó.

Có một vài nơi trong code có vấn đề về side-channel đã biết:
code RSA được che khuất nhưng không phải constant time, các đường cong elliptic khác
ngoài P-224 không phải constant time và cuộc tấn công Lucky13 có thể hoạt động. Tôi hy vọng giải quyết hai vấn đề sau trong khung thời gian Go 1.2 với việc triển khai P-256 constant-time
và AES-GCM.

Tuy nhiên, chưa có ai bước ra để thực hiện đánh giá TLS stack và tôi chưa
điều tra xem có thể nhờ Matasano hay tương tự không. Điều đó phụ thuộc vào
việc Google có muốn tài trợ hay không.

_Bạn nghĩ gì về_ [_GopherCon 2014_](http://www.gophercon.com/)_?_
_Có ai trong nhóm dự định tham dự không?_

**Andrew:** Thật thú vị. Tôi chắc chắn một số chúng tôi sẽ ở đó.
