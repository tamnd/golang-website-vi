---
title: Kết quả Khảo sát Go 2017
date: 2018-02-26
by:
- Steve Francia
tags:
- survey
- community
summary: Những gì chúng tôi thu được từ Khảo sát Người dùng Go tháng 12 năm 2017.
template: true
---

## Cảm ơn

Bài đăng này tóm tắt kết quả khảo sát người dùng năm 2017 cùng với
các nhận xét và phân tích. Bài đăng cũng so sánh các kết quả chính giữa khảo sát năm 2016 và
năm 2017.

Năm nay chúng tôi có 6.173 người tham gia khảo sát, nhiều hơn 70% so với 3.595 người trong
[Khảo sát Người dùng Go 2016](/blog/survey2016-results). Ngoài ra,
tỷ lệ hoàn thành cũng cao hơn một chút (84% lên 87%) và
tỷ lệ phản hồi với hầu hết các câu hỏi cũng cao hơn. Chúng tôi tin rằng độ dài khảo sát là
nguyên nhân chính của sự cải thiện này, vì khảo sát năm 2017 đã được rút ngắn dựa trên
phản hồi rằng khảo sát năm 2016 quá dài.

Chúng tôi biết ơn tất cả mọi người đã cung cấp phản hồi qua khảo sát để
giúp định hình tương lai của Go.

## Nền tảng lập trình

Lần đầu tiên, nhiều người tham gia khảo sát cho biết họ được trả lương để viết Go
hơn là viết Go ngoài giờ làm việc. Điều này cho thấy sự thay đổi đáng kể trong
cơ sở người dùng của Go và sự chấp nhận của các công ty đối với phát triển phần mềm chuyên nghiệp bằng Go.

Các lĩnh vực mà người tham gia khảo sát làm việc hầu hết nhất quán với
năm ngoái, tuy nhiên các ứng dụng di động và máy tính để bàn đã giảm đáng kể.

Một sự thay đổi quan trọng khác: ứng dụng số một của Go hiện nay là viết dịch vụ API/RPC (65%,
tăng 5% so với năm 2016), vượt qua vị trí đầu từ viết công cụ CLI bằng Go (63%).
Cả hai đều tận dụng tối đa các tính năng đặc trưng của Go và là những yếu tố then chốt của
điện toán đám mây hiện đại. Khi ngày càng nhiều công ty áp dụng Go, chúng tôi kỳ vọng hai ứng dụng
này của Go sẽ tiếp tục phát triển mạnh.

Hầu hết các số liệu đều tái khẳng định những điều chúng tôi đã tìm hiểu trong những năm trước. Các lập trình viên Go
vẫn ưa thích Go một cách áp đảo. Khi thời gian trôi qua, người dùng Go ngày càng đào sâu
kinh nghiệm của họ với Go. Mặc dù Go đã tăng khoảng cách dẫn đầu trong số các nhà phát triển Go,
thứ tự xếp hạng ngôn ngữ vẫn khá nhất quán với năm ngoái.

{{raw (file "survey2017/background.html")}}

## Sử dụng Go

Trong hầu hết mọi câu hỏi về việc sử dụng và nhận thức về Go, Go đã
cho thấy sự cải thiện so với khảo sát trước của chúng tôi. Người dùng hài lòng hơn khi dùng Go, và
tỷ lệ cao hơn thích dùng Go cho dự án tiếp theo của họ.

Khi được hỏi về những thách thức lớn nhất trong việc sử dụng Go cá nhân, người dùng
rõ ràng cho thấy rằng thiếu quản lý dependency và thiếu generics là hai vấn đề lớn nhất,
nhất quán với năm 2016. Trong năm 2017, chúng tôi đã đặt nền móng để
có thể giải quyết những vấn đề này. Chúng tôi đã cải thiện quy trình đề xuất và phát triển
với việc bổ sung
[Báo cáo Kinh nghiệm](/wiki/ExperienceReports), giúp
dự án thu thập và nhận được phản hồi quan trọng để thực hiện những
thay đổi đáng kể này. Chúng tôi cũng đã thực hiện
[những thay đổi đáng kể](/doc/go1.10#build) bên dưới về cách
Go lấy và xây dựng gói. Đây là công việc nền tảng thiết yếu để
giải quyết nhu cầu quản lý dependency của chúng tôi.

Hai vấn đề này sẽ tiếp tục là trọng tâm chính của dự án trong suốt năm 2018.

Trong phần này, chúng tôi đặt hai câu hỏi mới. Cả hai đều tập trung vào
những gì các nhà phát triển đang làm với Go theo cách chi tiết hơn so với những gì chúng tôi đã hỏi trước đây.
Chúng tôi hy vọng dữ liệu này sẽ cung cấp thông tin cho dự án Go và hệ sinh thái.

Kể từ năm ngoái, tỷ lệ người xác định "Go thiếu các tính năng quan trọng" là lý do họ không dùng Go nhiều hơn đã tăng,
và tỷ lệ người xác định "Go không phù hợp" đã giảm. Ngoài những thay đổi này,
danh sách vẫn nhất quán với năm ngoái.

{{raw (file "survey2017/usage.html")}}

## Phát triển và triển khai

Chúng tôi hỏi các lập trình viên họ phát triển Go trên hệ điều hành nào; tỷ lệ
phản hồi của họ vẫn nhất quán với năm ngoái. 64% người tham gia cho biết
họ dùng Linux, 49% dùng MacOS và 18% dùng Windows, với nhiều lựa chọn
được phép.

Tiếp tục tăng trưởng mạnh mẽ, VSCode hiện là trình soạn thảo phổ biến nhất trong số
các Gopher. IntelliJ/GoLand cũng có mức tăng sử dụng đáng kể. Những tăng trưởng này chủ yếu
đến từ Atom và Sublime Text, vốn đã giảm mức sử dụng tương đối.
Câu hỏi này có tỷ lệ phản hồi cao hơn 6% so với năm ngoái.

Người tham gia khảo sát cho thấy mức độ hài lòng cao hơn đáng kể với hỗ trợ Go
trong các trình soạn thảo so với năm 2016, với tỷ lệ hài lòng so với không hài lòng
tăng gấp đôi (9:1 lên 18:1). Cảm ơn tất cả mọi người đã làm việc về hỗ trợ trình soạn thảo Go
vì sự chăm chỉ của bạn.

Việc triển khai Go phân bổ tương đối đều nhau giữa các máy chủ do tư nhân quản lý và
các máy chủ đám mây được thuê ngoài. Đối với các ứng dụng Go, các dịch vụ Google Cloud đã tăng đáng kể
so với năm 2016. Đối với các ứng dụng không phải Go, AWS Lambda có mức tăng sử dụng lớn nhất.

{{raw (file "survey2017/dev.html")}}

## Làm việc Hiệu quả

Chúng tôi hỏi mức độ đồng ý hay không đồng ý của mọi người với các nhận định khác nhau về
Go. Tất cả các câu hỏi được lặp lại từ năm ngoái với việc bổ sung một câu hỏi mới
mà chúng tôi đưa ra để làm rõ thêm về cách người dùng có thể vừa tìm vừa **sử dụng** các thư viện Go.

Tất cả các phản hồi đều cho thấy một sự cải thiện nhỏ hoặc tương đương với năm 2016.

Như trong năm 2016, thư viện còn thiếu được yêu cầu nhiều nhất cho Go là một thư viện để
viết giao diện đồ họa (GUI), dù nhu cầu không nhiều như năm ngoái. Không có thư viện thiếu nào khác
nhận được số lượng phản hồi đáng kể.

Các nguồn chính để tìm câu trả lời cho các câu hỏi về Go là trang web Go,
Stack Overflow và đọc mã nguồn trực tiếp. Stack Overflow cho thấy mức tăng nhỏ
so với năm ngoái.

Các nguồn tin tức Go chính vẫn là blog Go, /r/golang trên Reddit và
Twitter; như năm ngoái, có thể có một số thiên lệch ở đây vì đây cũng là cách
khảo sát được thông báo.

{{raw (file "survey2017/effective.html")}}

## Dự án Go

59% người tham gia bày tỏ sự quan tâm đến việc đóng góp theo một cách nào đó cho cộng đồng và
các dự án Go, tăng từ 55% năm ngoái. Người tham gia cũng cho biết họ cảm thấy
được chào đón để đóng góp hơn nhiều so với năm 2016. Tiếc là,
người tham gia chỉ cho thấy sự cải thiện rất nhỏ trong việc hiểu cách
đóng góp. Chúng tôi sẽ tích cực làm việc với cộng đồng và các lãnh đạo của nó
để làm cho quy trình này dễ tiếp cận hơn.

Người tham gia cho thấy sự gia tăng đồng ý rằng họ tin tưởng vào
ban lãnh đạo của dự án Go (9:1 lên 11:1). Họ cũng cho thấy sự gia tăng nhỏ trong
việc đồng ý rằng ban lãnh đạo dự án hiểu nhu cầu của họ (2,6:1 lên 2,8:1)
và đồng ý rằng họ cảm thấy thoải mái khi tiếp cận ban lãnh đạo dự án với câu hỏi và phản hồi (2,2:1 lên 2,4:1). Dù đã có những cải tiến, đây
vẫn tiếp tục là lĩnh vực cần tập trung cho dự án và ban lãnh đạo của nó trong
tương lai. Chúng tôi sẽ tiếp tục cải thiện sự hiểu biết về nhu cầu người dùng và
khả năng tiếp cận của chúng tôi.

Chúng tôi đã thử [những cách mới](/blog/8years#TOC_1.3.) để tương tác
với người dùng trong năm 2017 và mặc dù đã có tiến bộ, chúng tôi vẫn đang làm việc để biến những
giải pháp này có thể mở rộng được cho cộng đồng ngày càng lớn mạnh của chúng tôi.

{{raw (file "survey2017/project.html")}}

## Cộng đồng

Ở cuối khảo sát, chúng tôi hỏi một số câu hỏi về nhân khẩu học.

Phân phối quốc gia của các phản hồi phần lớn tương tự như năm ngoái với những
biến động nhỏ. Như năm ngoái, phân phối quốc gia tương tự với
lượng truy cập vào golang.org, mặc dù một số quốc gia châu Á vẫn còn thiếu đại diện trong
khảo sát.

Có lẽ cải tiến đáng kể nhất so với năm 2016 đến từ câu hỏi hỏi người tham gia
đồng ý ở mức độ nào với nhận định "Tôi cảm thấy được chào đón trong cộng đồng Go".
Năm ngoái, tỷ lệ đồng ý so với không đồng ý là 15:1. Trong
năm 2017, tỷ lệ này gần gấp đôi lên 25:1.

Một phần quan trọng của cộng đồng là làm cho mọi người cảm thấy được chào đón, đặc biệt là
những người từ các nhóm ít được đại diện. Chúng tôi đặt một câu hỏi tùy chọn về
nhận dạng trong một vài nhóm thiểu số. Chúng tôi có tỷ lệ phản hồi tăng 4% so với
năm ngoái. Tỷ lệ của mỗi nhóm thiểu số đã tăng so với năm 2016, một số khá đáng kể.

Như năm ngoái, chúng tôi lấy kết quả của nhận định "Tôi cảm thấy được chào đón trong cộng đồng Go"
và phân tích theo các phản hồi về các nhóm thiểu số khác nhau. Như toàn bộ, hầu hết người tham gia xác định là
thiểu số cũng cảm thấy được chào đón trong cộng đồng Go nhiều hơn đáng kể so với
năm 2016. Người tham gia xác định là phụ nữ cho thấy sự cải thiện đáng kể nhất
với mức tăng hơn 400% trong tỷ lệ đồng ý so với không đồng ý với
nhận định này (3:1 lên 13:1). Người xác định là thiểu số về dân tộc hoặc chủng tộc
có mức tăng hơn 250% (7:1 lên 18:1). Như năm ngoái,
những người không xác định là thiểu số vẫn có tỷ lệ đồng ý cao hơn nhiều với nhận định này
so với những người xác định thuộc nhóm thiểu số.

Chúng tôi được khuyến khích bởi những tiến bộ này và hy vọng đà tiến tục tiếp diễn.

Câu hỏi cuối cùng của khảo sát chỉ là để vui: từ khóa Go yêu thích của bạn là gì?
Có lẽ không có gì đáng ngạc nhiên, câu trả lời phổ biến nhất là `go`, tiếp theo là
`defer`, `func`, `interface` và `select`, không thay đổi so với năm ngoái.

{{raw (file "survey2017/community.html")}}

Cuối cùng, thay mặt toàn bộ dự án Go, chúng tôi biết ơn tất cả mọi người đã
đóng góp cho dự án của chúng tôi, dù bằng cách là một phần của cộng đồng tuyệt vời của chúng tôi,
tham gia khảo sát này hay quan tâm đến Go.
