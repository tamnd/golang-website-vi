---
title: "Các câu hỏi thường gặp (FAQ)"
sidebar: "faq"
breadcrumb: true
template: true
---

## Nguồn gốc {#Origins}

### Mục đích của dự án là gì? {#What_is_the_purpose_of_the_project}

Vào thời điểm Go ra đời năm 2007, thế giới lập trình rất khác so với ngày nay.
Phần mềm cho môi trường production thường được viết bằng C++ hoặc Java,
GitHub chưa tồn tại, hầu hết máy tính chưa phải máy đa nhân,
và ngoài Visual Studio cùng Eclipse, hầu như không có IDE hay công cụ cấp cao nào,
chứ chưa nói đến miễn phí trên Internet.

Trong khi đó, chúng tôi ngày càng thấy bực bội với sự phức tạp không cần thiết khi
xây dựng các dự án phần mềm lớn bằng các ngôn ngữ mình đang dùng và hệ thống build đi kèm.
Máy tính đã nhanh hơn rất nhiều kể từ khi C, C++ và Java ra đời,
nhưng bản thân hoạt động lập trình lại không tiến bộ tương xứng.
Ngoài ra, rõ ràng là bộ xử lý đa nhân đang trở nên phổ biến,
nhưng hầu hết ngôn ngữ không cung cấp nhiều hỗ trợ để lập trình trên chúng
một cách hiệu quả và an toàn.

Chúng tôi quyết định lùi lại và suy nghĩ về những vấn đề lớn sẽ chi phối
kỹ thuật phần mềm trong những năm tới khi công nghệ phát triển,
và một ngôn ngữ mới có thể giúp giải quyết chúng như thế nào.
Ví dụ, sự trỗi dậy của CPU đa nhân cho thấy ngôn ngữ cần hỗ trợ ở cấp độ
đầu tiên một dạng concurrency hoặc song song nào đó.
Và để quản lý tài nguyên khả thi trong một chương trình concurrent lớn,
bộ gom rác, hoặc ít nhất là một dạng quản lý bộ nhớ tự động an toàn, là cần thiết.

Những cân nhắc đó dẫn đến
[một loạt cuộc thảo luận](https://commandcenter.blogspot.com/2017/09/go-ten-years-and-climbing.html)
từ đó Go ra đời, ban đầu là một tập hợp ý tưởng và mong muốn,
rồi trở thành một ngôn ngữ.
Mục tiêu bao trùm là Go làm được nhiều hơn để hỗ trợ lập trình viên
bằng cách cải thiện hệ thống công cụ, tự động hóa các tác vụ nhàm chán như
định dạng mã, và loại bỏ các rào cản khi làm việc với các codebase lớn.

Một mô tả chi tiết hơn về mục tiêu của Go và cách chúng được đáp ứng,
hoặc ít nhất là tiệm cận, có trong bài viết
[Go tại Google: Thiết kế ngôn ngữ phục vụ kỹ thuật phần mềm](/talks/2012/splash.article).

### Lịch sử của dự án ra sao? {#history}

Robert Griesemer, Rob Pike và Ken Thompson bắt đầu phác thảo
các mục tiêu cho một ngôn ngữ mới trên bảng trắng vào ngày 21 tháng 9 năm 2007.
Trong vài ngày, các mục tiêu đã hình thành thành một kế hoạch cụ thể
và một ý tưởng khá rõ về hướng đi. Thiết kế tiếp tục bán thời gian
song song với các công việc khác. Đến tháng 1 năm 2008, Ken đã bắt đầu
xây dựng một trình biên dịch để khám phá các ý tưởng; nó tạo ra mã C làm đầu ra.
Đến giữa năm, ngôn ngữ đã trở thành một dự án toàn thời gian và đã đủ ổn định
để thử xây dựng trình biên dịch cho môi trường production. Tháng 5 năm 2008,
Ian Taylor độc lập bắt đầu xây dựng GCC front end cho Go dựa trên bản đặc tả nháp.
Russ Cox gia nhập cuối năm 2008 và giúp đưa ngôn ngữ và thư viện từ prototype sang thực tế.

Go trở thành dự án mã nguồn mở công khai vào ngày 10 tháng 11 năm 2009.
Vô số người từ cộng đồng đã đóng góp ý tưởng, thảo luận và mã nguồn.

Hiện nay có hàng triệu lập trình viên Go, hay còn gọi là gopher, trên toàn thế giới,
và con số đó tăng lên mỗi ngày.
Sự thành công của Go đã vượt xa kỳ vọng của chúng tôi.

### Nguồn gốc của linh vật gopher là gì? {#gopher}

Linh vật và logo được thiết kế bởi
[Renée French](https://reneefrench.blogspot.com), người cũng thiết kế
[Glenda](https://9p.io/plan9/glenda.html),
thỏ của Plan 9.
Một [bài đăng trên blog](/blog/gopher)
về gopher giải thích cách nó được phát triển
từ hình ảnh cô ấy đã dùng cho thiết kế áo phông của [WFMU](https://wfmu.org/)
vài năm trước.
Logo và linh vật được bảo vệ bởi giấy phép
[Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/).

Gopher có một
[tờ mô hình](/doc/gopher/modelsheet.jpg)
minh họa các đặc điểm của nó và cách biểu diễn chúng đúng cách.
Tờ mô hình này lần đầu được trình bày trong một
[bài nói chuyện](https://www.youtube.com/watch?v=4rw_B4yY69k)
của Renée tại Gophercon năm 2016.
Nó có những đặc điểm riêng; đây là *Go gopher*, không phải gopher thông thường nào đó.

### Ngôn ngữ được gọi là Go hay Golang? {#go_or_golang}

Ngôn ngữ có tên là Go.
Biệt danh "golang" xuất hiện vì trang web ban đầu là *golang.org*.
(Lúc đó chưa có tên miền *.dev*.)
Nhiều người vẫn dùng tên golang và nó tiện lợi như một nhãn hiệu.
Ví dụ, thẻ mạng xã hội của ngôn ngữ là "#golang".
Tuy nhiên, tên ngôn ngữ vẫn chỉ là Go.

Một lưu ý nhỏ: mặc dù
[logo chính thức](/blog/go-brand)
có hai chữ in hoa, tên ngôn ngữ được viết là Go, không phải GO.

### Tại sao lại tạo ra một ngôn ngữ mới? {#creating_a_new_language}

Go ra đời từ sự bực bội với các ngôn ngữ và môi trường hiện có
trong công việc chúng tôi đang làm tại Google.
Lập trình đã trở nên quá khó và lựa chọn ngôn ngữ có phần trách nhiệm trong đó.
Người ta phải chọn giữa biên dịch hiệu quả, thực thi hiệu quả, hoặc dễ lập trình;
cả ba không cùng có trong một ngôn ngữ mainstream.
Lập trình viên có khả năng đang ưu tiên sự dễ dàng hơn an toàn và hiệu quả
bằng cách chuyển sang các ngôn ngữ kiểu động như Python và JavaScript
thay vì C++ hoặc Java.

Chúng tôi không đơn độc trong mối lo này.
Sau nhiều năm yên tĩnh trong lĩnh vực ngôn ngữ lập trình,
Go là một trong những ngôn ngữ mới đầu tiên -- Rust, Elixir, Swift và nhiều hơn nữa --
đã biến phát triển ngôn ngữ lập trình thành một lĩnh vực sôi động, gần như chính thống trở lại.

Go giải quyết những vấn đề này bằng cách cố gắng kết hợp sự dễ lập trình
của một ngôn ngữ thông dịch, kiểu động với hiệu quả và an toàn của một
ngôn ngữ kiểu tĩnh, biên dịch. Nó cũng nhằm thích nghi tốt hơn với phần cứng hiện tại,
với hỗ trợ cho máy tính mạng và đa nhân.
Cuối cùng, làm việc với Go được thiết kế để *nhanh*: chỉ cần vài giây
để build một tệp thực thi lớn trên một máy đơn.
Để đạt được các mục tiêu này, chúng tôi phải xem xét lại một số cách tiếp cận lập trình,
dẫn đến: hệ thống kiểu theo tổ hợp thay vì phân cấp;
hỗ trợ concurrency và bộ gom rác; đặc tả dependency chặt chẽ; và nhiều thứ khác.
Những điều này không thể xử lý tốt bằng thư viện hoặc công cụ;
cần phải có một ngôn ngữ mới.

Bài viết [Go tại Google](/talks/2012/splash.article)
thảo luận về bối cảnh và động lực đằng sau thiết kế của ngôn ngữ Go,
cũng như cung cấp thêm chi tiết về nhiều câu trả lời trong FAQ này.

### Tổ tiên của Go là gì? {#ancestors}

Go chủ yếu thuộc họ C (cú pháp cơ bản),
với đóng góp đáng kể từ họ Pascal/Modula/Oberon
(khai báo, package),
cộng thêm một số ý tưởng từ các ngôn ngữ
lấy cảm hứng từ CSP của Tony Hoare,
như Newsqueak và Limbo (concurrency).
Tuy nhiên, đây là một ngôn ngữ mới hoàn toàn.
Ở mọi khía cạnh, ngôn ngữ được thiết kế bằng cách suy nghĩ
về những gì lập trình viên làm và cách làm cho lập trình,
ít nhất là loại lập trình chúng tôi làm, hiệu quả hơn, tức là thú vị hơn.

### Các nguyên tắc hướng dẫn trong thiết kế là gì? {#principles}

Khi Go được thiết kế, Java và C++ là hai ngôn ngữ được dùng phổ biến nhất
để viết server, ít nhất là tại Google.
Chúng tôi cảm thấy các ngôn ngữ đó đòi hỏi
quá nhiều thao tác hành chính và lặp lại.
Một số lập trình viên phản ứng bằng cách chuyển sang các ngôn ngữ
động hơn, linh hoạt hơn như Python, đánh đổi hiệu quả và an toàn kiểu.
Chúng tôi cho rằng hoàn toàn có thể có cả hiệu quả,
an toàn và sự linh hoạt trong một ngôn ngữ duy nhất.

Go cố gắng giảm lượng gõ phím theo cả nghĩa đen và nghĩa bóng.
Trong suốt quá trình thiết kế, chúng tôi đã cố gắng giảm sự lộn xộn và
phức tạp. Không có khai báo tiền tố, không có tệp header;
mọi thứ được khai báo đúng một lần. Khởi tạo rõ ràng, tự động và dễ dùng.
Cú pháp gọn gàng, ít từ khóa. Sự lặp lại (`foo.Foo* myFoo = new(foo.Foo)`)
được rút gọn bằng suy luận kiểu đơn giản qua cú pháp khai báo và khởi tạo `:=`.
Và có lẽ triệt để nhất, không có phân cấp kiểu: các kiểu chỉ *tồn tại*,
chúng không cần phải công bố mối quan hệ của mình.
Những đơn giản hóa này cho phép Go biểu đạt tốt mà vẫn dễ hiểu
mà không mất đi năng suất.

Một nguyên tắc quan trọng khác là giữ các khái niệm trực giao.
Phương thức có thể được triển khai cho bất kỳ kiểu nào;
struct đại diện cho dữ liệu trong khi interface đại diện cho trừu tượng; và tương tự.
Tính trực giao giúp dễ hiểu điều gì xảy ra khi các thứ kết hợp lại.

## Sử dụng {#Usage}

### Google có dùng Go nội bộ không? {#internal_usage}

Có. Go được dùng rộng rãi trong môi trường production bên trong Google.
Một ví dụ là máy chủ tải xuống của Google, `dl.google.com`,
phân phối các tệp nhị phân Chrome và các gói cài đặt lớn khác như gói `apt-get`.

Go không phải ngôn ngữ duy nhất được dùng tại Google, nhưng đây là ngôn ngữ
quan trọng cho nhiều lĩnh vực bao gồm
[kỹ thuật độ tin cậy trang web (SRE)](/talks/2013/go-sreops.slide)
và xử lý dữ liệu quy mô lớn.
Nó cũng là một phần cốt lõi của phần mềm chạy Google Cloud.

### Các công ty nào khác dùng Go? {#external_usage}

Việc sử dụng Go đang tăng trưởng trên toàn thế giới, đặc biệt nhưng không chỉ
trong không gian điện toán đám mây.
Một số dự án hạ tầng đám mây lớn viết bằng Go là Docker và Kubernetes,
nhưng còn nhiều hơn nữa.

Không chỉ là đám mây, như bạn có thể thấy từ danh sách các công ty
trên [trang web go.dev](/)
cùng với một số
[câu chuyện thành công](/solutions/case-studies).
Ngoài ra, Go Wiki có một
[trang](/wiki/GoUsers),
được cập nhật thường xuyên, liệt kê một số trong nhiều công ty đang dùng Go.

Wiki cũng có một trang với các liên kết đến thêm
[câu chuyện thành công](/wiki/SuccessStories)
về các công ty và dự án đang sử dụng ngôn ngữ này.

### Chương trình Go có liên kết với chương trình C/C++ không? {#Do_Go_programs_link_with_Cpp_programs}

Có thể dùng C và Go cùng nhau trong cùng một không gian địa chỉ,
nhưng đây không phải là sự kết hợp tự nhiên và có thể cần phần mềm giao tiếp đặc biệt.
Ngoài ra, liên kết C với mã Go từ bỏ các tính năng an toàn bộ nhớ
và quản lý stack mà Go cung cấp.
Đôi khi hoàn toàn cần thiết phải dùng thư viện C để giải quyết một vấn đề,
nhưng làm vậy luôn mang lại một yếu tố rủi ro không có trong mã Go thuần túy,
vì vậy hãy cẩn thận.

Nếu bạn cần dùng C với Go, cách tiến hành phụ thuộc vào cách triển khai trình biên dịch Go.
Trình biên dịch "chuẩn", một phần của Go toolchain được nhóm Go tại Google hỗ trợ,
được gọi là `gc`.
Ngoài ra, còn có trình biên dịch dựa trên GCC (`gccgo`) và
trình biên dịch dựa trên LLVM (`gollvm`),
cũng như danh sách ngày càng tăng của các trình biên dịch đặc biệt phục vụ các mục đích khác nhau,
đôi khi triển khai các tập con ngôn ngữ,
chẳng hạn như [TinyGo](https://tinygo.org/).

`Gc` dùng quy ước gọi hàm và linker khác với C và
do đó không thể được gọi trực tiếp từ chương trình C, hoặc ngược lại.
Chương trình [`cgo`](/cmd/cgo/) cung cấp cơ chế cho
"foreign function interface" để cho phép gọi an toàn
các thư viện C từ mã Go.
SWIG mở rộng khả năng này sang các thư viện C++.

Bạn cũng có thể dùng `cgo` và SWIG với `gccgo` và `gollvm`.
Vì chúng dùng ABI truyền thống, cũng có thể, với sự cẩn thận cao,
liên kết mã từ các trình biên dịch này trực tiếp với các chương trình C hoặc C++ được biên dịch bởi GCC/LLVM.
Tuy nhiên, làm vậy một cách an toàn đòi hỏi hiểu biết về các quy ước gọi hàm
cho tất cả các ngôn ngữ liên quan, cũng như quan tâm đến giới hạn stack khi gọi C hoặc C++
từ Go.

### Go hỗ trợ những IDE nào? {#ide}

Dự án Go không bao gồm một IDE tùy chỉnh, nhưng ngôn ngữ và
thư viện được thiết kế để dễ phân tích mã nguồn.
Do đó, hầu hết các trình soạn thảo và IDE nổi tiếng đều hỗ trợ Go tốt,
hoặc trực tiếp hoặc thông qua plugin.

Nhóm Go cũng hỗ trợ một Go language server cho giao thức LSP, gọi là
[`gopls`](https://pkg.go.dev/golang.org/x/tools/gopls#section-readme).
Các công cụ hỗ trợ LSP có thể dùng `gopls` để tích hợp hỗ trợ theo ngôn ngữ cụ thể.

Danh sách các IDE và trình soạn thảo nổi tiếng cung cấp hỗ trợ Go tốt
bao gồm Emacs, Vim, VSCode, Atom, Eclipse, Sublime, IntelliJ
(qua một biến thể tùy chỉnh gọi là GoLand) và nhiều hơn nữa.
Khả năng cao môi trường yêu thích của bạn cũng là một môi trường
năng suất để lập trình bằng Go.

### Go có hỗ trợ protocol buffers của Google không? {#protocol_buffers}

Một dự án mã nguồn mở riêng cung cấp plugin trình biên dịch và thư viện cần thiết.
Nó có tại
[github.com/golang/protobuf/](https://github.com/golang/protobuf).

## Thiết kế {#Design}

### Go có runtime không? {#runtime}

Go có một thư viện runtime phong phú, thường chỉ được gọi là *runtime*,
là một phần của mọi chương trình Go.
Thư viện này triển khai bộ gom rác, concurrency,
quản lý stack và các tính năng quan trọng khác của ngôn ngữ Go.
Mặc dù nó có vai trò trung tâm hơn với ngôn ngữ, runtime của Go tương tự
như `libc`, thư viện C.

Tuy nhiên, điều quan trọng cần hiểu là runtime của Go không
bao gồm máy ảo, như được cung cấp bởi Java runtime.
Các chương trình Go được biên dịch trước thời gian thực thi sang mã máy gốc
(hoặc JavaScript hoặc WebAssembly, cho một số triển khai biến thể).
Do đó, mặc dù thuật ngữ thường được dùng để mô tả môi trường ảo
nơi chương trình chạy, trong Go từ "runtime"
chỉ là tên được đặt cho thư viện cung cấp các dịch vụ ngôn ngữ quan trọng.

### Sao lại có định danh Unicode? {#unicode_identifiers}

Khi thiết kế Go, chúng tôi muốn đảm bảo rằng nó không
quá lấy ASCII làm trung tâm,
điều đó có nghĩa là mở rộng không gian định danh ra ngoài
phạm vi 7-bit ASCII.
Quy tắc của Go -- ký tự định danh phải là
chữ cái hoặc chữ số theo định nghĩa của Unicode -- đơn giản để hiểu
và triển khai nhưng có những hạn chế.
Ký tự kết hợp bị loại trừ theo thiết kế,
chẳng hạn, và điều đó loại trừ một số ngôn ngữ như Devanagari.

Quy tắc này có một hậu quả đáng tiếc khác.
Vì định danh được xuất phải bắt đầu bằng chữ hoa,
các định danh được tạo từ ký tự trong một số ngôn ngữ
theo định nghĩa không thể được xuất.
Hiện tại giải pháp duy nhất là dùng thứ gì đó như `X日本語`,
rõ ràng là không thỏa đáng.

Kể từ phiên bản đầu tiên của ngôn ngữ, đã có nhiều suy nghĩ
về cách tốt nhất để mở rộng không gian định danh để phù hợp với
lập trình viên sử dụng các ngôn ngữ tự nhiên khác.
Chính xác phải làm gì vẫn là chủ đề thảo luận tích cực,
và một phiên bản tương lai của ngôn ngữ có thể sẽ thoáng hơn
trong định nghĩa về định danh.
Ví dụ, nó có thể áp dụng một số ý tưởng từ
[khuyến nghị](http://unicode.org/reports/tr31/)
của tổ chức Unicode về định danh.
Dù điều gì xảy ra, nó phải được thực hiện một cách tương thích trong khi vẫn giữ
(hoặc có thể mở rộng) cách chữ hoa/thường xác định khả năng hiển thị của
định danh, vẫn là một trong những tính năng yêu thích của chúng tôi trong Go.

Hiện tại, chúng tôi có một quy tắc đơn giản có thể mở rộng sau này
mà không làm hỏng chương trình, một quy tắc tránh các lỗi chắc chắn sẽ phát sinh
từ một quy tắc chấp nhận các định danh mơ hồ.

### Tại sao Go không có tính năng X? {#Why_doesnt_Go_have_feature_X}

Mọi ngôn ngữ đều có các tính năng mới lạ và bỏ qua tính năng yêu thích của ai đó.
Go được thiết kế với cái nhìn về sự dễ chịu khi lập trình, tốc độ biên dịch,
tính trực giao của các khái niệm, và nhu cầu hỗ trợ các tính năng
như concurrency và bộ gom rác. Tính năng yêu thích của bạn có thể bị thiếu
vì nó không phù hợp, vì nó ảnh hưởng đến tốc độ biên dịch hoặc sự rõ ràng của thiết kế,
hoặc vì nó sẽ làm mô hình hệ thống cơ bản quá khó.

Nếu bạn khó chịu vì Go thiếu tính năng *X*,
hãy thứ lỗi cho chúng tôi và khám phá các tính năng mà Go có.
Bạn có thể thấy chúng bù đắp theo những cách thú vị cho sự thiếu vắng của *X*.

### Go có generic types từ khi nào? {#generics}

Bản phát hành Go 1.18 đã thêm tham số kiểu vào ngôn ngữ.
Điều này cho phép một dạng lập trình đa hình hoặc generic.
Xem [đặc tả ngôn ngữ](/ref/spec) và
[đề xuất](/design/43651-type-parameters) để biết chi tiết.

### Tại sao Go ban đầu được phát hành mà không có generic types? {#beginning_generics}

Go được thiết kế như một ngôn ngữ để viết các chương trình server
dễ bảo trì theo thời gian.
(Xem [bài viết này](/talks/2012/splash.article) để biết thêm bối cảnh.)
Thiết kế tập trung vào những thứ như khả năng mở rộng, khả năng đọc và concurrency.
Lập trình đa hình có vẻ không thiết yếu với các mục tiêu của ngôn ngữ vào thời điểm đó,
và do đó ban đầu bị bỏ qua để giữ sự đơn giản.

Generics rất tiện lợi nhưng chúng có chi phí về độ phức tạp trong
hệ thống kiểu và thời gian chạy.
Mất một thời gian để phát triển một thiết kế mà chúng tôi tin là mang lại giá trị
tương xứng với độ phức tạp.

### Tại sao Go không có exception? {#exceptions}

Chúng tôi cho rằng việc gắn exception vào cấu trúc điều khiển,
như trong mô hình `try-catch-finally`, dẫn đến mã phức tạp.
Nó cũng có xu hướng khuyến khích lập trình viên gán nhãn
quá nhiều lỗi thông thường, chẳng hạn như không mở được tệp,
là ngoại lệ.

Go có cách tiếp cận khác. Để xử lý lỗi đơn giản, giá trị trả về đa giá trị của Go
giúp dễ dàng báo cáo lỗi mà không làm quá tải giá trị trả về.
[Một kiểu lỗi chuẩn, kết hợp với các tính năng khác của Go](/doc/articles/error_handling.html),
làm cho xử lý lỗi dễ chịu nhưng khá khác biệt so với các ngôn ngữ khác.

Go cũng có một vài hàm tích hợp để báo hiệu và khôi phục từ
các điều kiện thực sự bất thường. Cơ chế khôi phục chỉ được thực thi
như một phần của trạng thái hàm đang bị hủy sau khi xảy ra lỗi,
đủ để xử lý thảm họa nhưng không cần cấu trúc điều khiển bổ sung và,
khi dùng tốt, có thể dẫn đến mã xử lý lỗi gọn gàng.

Xem bài viết [Defer, Panic, và Recover](/doc/articles/defer_panic_recover.html) để biết chi tiết.
Ngoài ra, bài đăng trên blog [Lỗi là giá trị](/blog/errors-are-values)
mô tả một cách tiếp cận xử lý lỗi gọn gàng trong Go bằng cách chứng minh rằng,
vì lỗi chỉ là các giá trị, toàn bộ sức mạnh của Go có thể được áp dụng trong xử lý lỗi.

### Tại sao Go không có assertions? {#assertions}

Go không cung cấp assertions. Chúng chắc chắn thuận tiện, nhưng kinh nghiệm của chúng tôi
là lập trình viên dùng chúng như nạng để tránh suy nghĩ về xử lý và báo cáo lỗi đúng cách.
Xử lý lỗi đúng cách có nghĩa là server tiếp tục hoạt động thay vì sụp đổ
sau một lỗi không nghiêm trọng.
Báo cáo lỗi đúng cách có nghĩa là lỗi trực tiếp và rõ ràng,
giúp lập trình viên không phải giải mã một stack trace sụp đổ lớn.
Lỗi chính xác đặc biệt quan trọng khi lập trình viên nhìn thấy lỗi
không quen với mã nguồn.

Chúng tôi hiểu đây là điểm gây tranh cãi. Có nhiều thứ trong
ngôn ngữ và thư viện Go khác với các thực hành hiện đại, đơn giản là
vì đôi khi chúng tôi cảm thấy đáng để thử một cách tiếp cận khác.

### Tại sao xây dựng concurrency dựa trên ý tưởng của CSP? {#csp}

Concurrency và lập trình đa luồng theo thời gian đã phát triển
tiếng xấu vì sự khó khăn. Chúng tôi tin điều này một phần do các thiết kế phức tạp
như [pthreads](https://en.wikipedia.org/wiki/POSIX_Threads)
và một phần do quá nhấn mạnh vào các chi tiết cấp thấp
như mutex, biến điều kiện và memory barrier.
Các giao diện cấp cao hơn cho phép mã đơn giản hơn nhiều, dù vẫn còn
mutex và các thứ tương tự bên dưới.

Một trong những mô hình thành công nhất để cung cấp hỗ trợ ngôn ngữ cấp cao
cho concurrency đến từ Communicating Sequential Processes của Hoare, hay CSP.
Occam và Erlang là hai ngôn ngữ nổi tiếng xuất phát từ CSP.
Các nguyên thủy concurrency của Go xuất phát từ một nhánh khác của gia đình
với đóng góp chính là khái niệm mạnh mẽ về channel như các đối tượng hạng nhất.
Kinh nghiệm với một số ngôn ngữ trước đó đã chứng tỏ mô hình CSP
phù hợp tốt với framework ngôn ngữ thủ tục.

### Tại sao là goroutine thay vì thread? {#goroutines}

Goroutine là một phần làm cho concurrency dễ dùng. Ý tưởng, đã tồn tại một thời gian,
là ghép các hàm thực thi độc lập -- coroutine -- lên một tập hợp các thread.
Khi một coroutine bị block, chẳng hạn bằng cách gọi một lời gọi hệ thống blocking,
runtime tự động chuyển các coroutine khác trên cùng thread hệ điều hành
sang một thread khác, có thể chạy để chúng không bị block.
Lập trình viên không thấy điều này, đó là điểm mấu chốt.
Kết quả, mà chúng tôi gọi là goroutine, có thể rất rẻ: chúng ít overhead
hơn ngoài bộ nhớ cho stack, chỉ vài kilobyte.

Để giữ các stack nhỏ, runtime của Go dùng các stack có thể thay đổi kích thước, có giới hạn.
Một goroutine mới được tạo ra nhận vài kilobyte, gần như luôn đủ.
Khi không đủ, runtime tự động tăng (và thu nhỏ) bộ nhớ để lưu
stack, cho phép nhiều goroutine tồn tại trong một lượng bộ nhớ vừa phải.
Overhead CPU trung bình khoảng ba lệnh rẻ tiền mỗi lần gọi hàm.
Hoàn toàn khả thi để tạo hàng trăm nghìn goroutine trong cùng
một không gian địa chỉ.
Nếu goroutine chỉ là thread, tài nguyên hệ thống sẽ cạn kiệt
ở một số lượng nhỏ hơn nhiều.

### Tại sao các thao tác map không được định nghĩa là atomic? {#atomic_maps}

Sau khi thảo luận dài, quyết định được đưa ra là cách dùng map thông thường không cần
truy cập an toàn từ nhiều goroutine, và trong những trường hợp cần thiết,
map thường là một phần của cấu trúc dữ liệu hoặc tính toán lớn hơn
đã được đồng bộ hóa. Do đó, yêu cầu tất cả các thao tác map phải giữ mutex
sẽ làm chậm hầu hết các chương trình và thêm an toàn cho ít trường hợp.
Tuy nhiên, đây không phải là quyết định dễ dàng,
vì nó có nghĩa là truy cập map không được kiểm soát có thể làm chương trình sụp đổ.

Ngôn ngữ không ngăn cập nhật map atomic. Khi cần,
chẳng hạn khi lưu trú một chương trình không đáng tin cậy, triển khai có thể khóa truy cập map.

Truy cập map không an toàn chỉ khi có cập nhật đang xảy ra.
Miễn là tất cả goroutine chỉ đọc -- tra cứu các phần tử trong map,
bao gồm lặp qua nó bằng vòng lặp `for` `range` -- và không thay đổi map
bằng cách gán cho các phần tử hoặc xóa,
chúng an toàn để truy cập map đồng thời mà không cần đồng bộ hóa.

Để hỗ trợ sử dụng map đúng cách, một số triển khai ngôn ngữ
chứa kiểm tra đặc biệt tự động báo cáo lúc chạy khi map bị sửa đổi
không an toàn bởi thực thi đồng thời.
Ngoài ra, có một kiểu trong thư viện sync gọi là
[`sync.Map`](https://pkg.go.dev/sync#Map) hoạt động
tốt cho một số mẫu sử dụng như cache tĩnh, mặc dù nó không
phù hợp như một thay thế tổng quát cho kiểu map tích hợp.

### Bạn có chấp nhận thay đổi ngôn ngữ của tôi không? {#language_changes}

Mọi người thường đề xuất cải tiến cho ngôn ngữ --
[mailing list](https://groups.google.com/group/golang-nuts)
chứa lịch sử phong phú của các cuộc thảo luận như vậy -- nhưng rất ít thay đổi này
được chấp nhận.

Mặc dù Go là dự án mã nguồn mở, ngôn ngữ và thư viện được bảo vệ
bởi [cam kết tương thích](/doc/go1compat.html) ngăn chặn
các thay đổi phá vỡ chương trình hiện có, ít nhất là ở cấp độ mã nguồn
(các chương trình đôi khi cần được biên dịch lại để cập nhật).
Nếu đề xuất của bạn vi phạm đặc tả Go 1, chúng tôi thậm chí không thể xem xét
ý tưởng, bất kể giá trị của nó.
Một bản phát hành chính của Go trong tương lai có thể không tương thích với Go 1,
nhưng các cuộc thảo luận về chủ đề đó mới chỉ bắt đầu và một điều chắc chắn:
sẽ có rất ít sự không tương thích như vậy được đưa vào trong quá trình.
Hơn nữa, cam kết tương thích khuyến khích chúng tôi cung cấp một đường dẫn tự động
để các chương trình cũ thích nghi nếu tình huống đó xảy ra.

Ngay cả khi đề xuất của bạn tương thích với đặc tả Go 1,
nó có thể không phù hợp với tinh thần mục tiêu thiết kế của Go.
Bài viết *[Go tại Google: Thiết kế ngôn ngữ phục vụ kỹ thuật phần mềm](/talks/2012/splash.article)*
giải thích nguồn gốc của Go và động lực đằng sau thiết kế của nó.

## Kiểu dữ liệu {#types}

### Go có phải là ngôn ngữ hướng đối tượng không? {#Is_Go_an_object-oriented_language}

Có và không. Mặc dù Go có kiểu và phương thức và cho phép
phong cách lập trình hướng đối tượng, nhưng không có phân cấp kiểu.
Khái niệm "interface" trong Go cung cấp một cách tiếp cận khác mà
chúng tôi tin là dễ dùng và theo một số cách tổng quát hơn. Cũng có
những cách để nhúng kiểu vào các kiểu khác để cung cấp thứ gì đó
tương tự -- nhưng không giống hệt -- với kế thừa phụ lớp.
Hơn nữa, phương thức trong Go tổng quát hơn trong C++ hoặc Java:
chúng có thể được định nghĩa cho bất kỳ loại dữ liệu nào, kể cả các kiểu tích hợp
như số nguyên thuần, "không đóng hộp".
Chúng không bị giới hạn ở struct (class).

Ngoài ra, việc thiếu phân cấp kiểu làm cho "đối tượng" trong Go
cảm thấy nhẹ nhàng hơn nhiều so với các ngôn ngữ như C++ hoặc Java.

### Làm thế nào để có dispatch phương thức động? {#How_do_I_get_dynamic_dispatch_of_methods}

Cách duy nhất để có phương thức được dispatch động là thông qua một interface.
Các phương thức trên struct hoặc bất kỳ kiểu cụ thể nào khác luôn được giải quyết tĩnh.

### Tại sao không có kế thừa kiểu? {#inheritance}

Lập trình hướng đối tượng, ít nhất trong các ngôn ngữ nổi tiếng nhất,
liên quan đến quá nhiều thảo luận về mối quan hệ giữa các kiểu,
các mối quan hệ thường có thể được suy ra tự động. Go có cách tiếp cận khác.

Thay vì yêu cầu lập trình viên khai báo trước hai kiểu có liên quan,
trong Go một kiểu tự động thỏa mãn bất kỳ interface nào
chỉ định một tập con phương thức của nó. Ngoài việc giảm thao tác hành chính,
cách tiếp cận này có những lợi thế thực sự. Các kiểu có thể thỏa mãn
nhiều interface cùng một lúc, không có sự phức tạp của kế thừa đa dạng truyền thống.
Interface có thể rất nhẹ -- một interface với
một hoặc thậm chí không có phương thức nào có thể biểu đạt một khái niệm hữu ích.
Interface có thể được thêm vào sau thực tế nếu có ý tưởng mới
hoặc để kiểm thử -- mà không cần chú thích các kiểu gốc.
Vì không có mối quan hệ rõ ràng giữa các kiểu và interface,
không có phân cấp kiểu để quản lý hoặc thảo luận.

Có thể dùng các ý tưởng này để xây dựng thứ gì đó tương tự
với Unix pipes an toàn về kiểu. Ví dụ, hãy xem cách `fmt.Fprintf`
cho phép in định dạng vào bất kỳ đầu ra nào, không chỉ tệp, hoặc cách package
`bufio` có thể hoàn toàn tách biệt khỏi I/O tệp,
hoặc cách các package `image` tạo ra các tệp hình ảnh nén.
Tất cả những ý tưởng này xuất phát từ một interface duy nhất
(`io.Writer`) đại diện cho một phương thức duy nhất
(`Write`). Và đó chỉ là bề mặt.
Interface của Go có ảnh hưởng sâu sắc đến cách các chương trình được cấu trúc.

Cần thời gian để quen nhưng phong cách phụ thuộc kiểu ngầm định này
là một trong những điều năng suất nhất về Go.

### Tại sao `len` là hàm chứ không phải phương thức? {#methods_on_basics}

Chúng tôi đã tranh luận vấn đề này nhưng quyết định
triển khai `len` và các hàm tương tự là các hàm thông thường trong thực tế
và không làm phức tạp các câu hỏi về interface (theo nghĩa kiểu Go)
của các kiểu cơ bản.

### Tại sao Go không hỗ trợ overloading phương thức và toán tử? {#overloading}

Dispatch phương thức được đơn giản hóa nếu nó không cần phải khớp kiểu.
Kinh nghiệm với các ngôn ngữ khác cho thấy rằng có nhiều
phương thức cùng tên nhưng chữ ký khác nhau đôi khi hữu ích
nhưng trong thực tế cũng có thể gây nhầm lẫn và dễ vỡ. Chỉ khớp theo tên
và yêu cầu nhất quán về kiểu là một quyết định đơn giản hóa lớn
trong hệ thống kiểu của Go.

Về overloading toán tử, có vẻ đó là sự tiện lợi hơn là yêu cầu tuyệt đối.
Một lần nữa, mọi thứ đơn giản hơn khi không có nó.

### Tại sao Go không có khai báo "implements"? {#implements_interface}

Một kiểu Go triển khai một interface bằng cách triển khai các phương thức của interface đó,
không cần gì thêm. Thuộc tính này cho phép interface được định nghĩa và dùng
mà không cần sửa đổi mã hiện có. Nó tạo ra một dạng
[structural typing](https://en.wikipedia.org/wiki/Structural_type_system) thúc đẩy
phân tách mối quan tâm và cải thiện tái sử dụng mã, và giúp dễ dàng hơn
để xây dựng trên các mẫu xuất hiện khi mã phát triển.
Ngữ nghĩa của interface là một trong những lý do chính cho cảm giác
nhanh nhẹn, nhẹ nhàng của Go.

Xem [câu hỏi về kế thừa kiểu](#inheritance) để biết thêm chi tiết.

### Làm thế nào để đảm bảo kiểu của tôi thỏa mãn một interface? {#guarantee_satisfies_interface}

Bạn có thể yêu cầu trình biên dịch kiểm tra rằng kiểu `T` triển khai
interface `I` bằng cách thử một phép gán dùng giá trị zero cho
`T` hoặc con trỏ tới `T`, tùy theo trường hợp:

```
type T struct{}
var _ I = T{}       // Xác nhận T triển khai I.
var _ I = (*T)(nil) // Xác nhận *T triển khai I.
```

Nếu `T` (hoặc `*T`, tương ứng) không triển khai
`I`, lỗi sẽ bị phát hiện tại thời điểm biên dịch.

Nếu bạn muốn người dùng của một interface khai báo rõ ràng rằng họ triển khai nó,
bạn có thể thêm một phương thức với tên mô tả vào tập phương thức của interface.
Ví dụ:

```
type Fooer interface {
    Foo()
    ImplementsFooer()
}
```

Một kiểu sau đó phải triển khai phương thức `ImplementsFooer` để là một
`Fooer`, tài liệu hóa rõ ràng thực tế và thông báo nó trong
đầu ra của [go doc](/cmd/go/#hdr-Show_documentation_for_package_or_symbol).

```
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

Hầu hết mã không dùng các ràng buộc như vậy, vì chúng hạn chế tiện ích
của ý tưởng interface. Tuy nhiên, đôi khi chúng cần thiết để giải quyết sự mơ hồ
giữa các interface tương tự.

### Tại sao kiểu T không thỏa mãn interface Equal? {#t_and_equal_interface}

Hãy xem xét interface đơn giản này để đại diện cho một đối tượng có thể so sánh
chính nó với giá trị khác:

```
type Equaler interface {
    Equal(Equaler) bool
}
```

và kiểu này, `T`:

```
type T int
func (t T) Equal(u T) bool { return t == u } // không thỏa mãn Equaler
```

Không giống tình huống tương tự trong một số hệ thống kiểu đa hình,
`T` không triển khai `Equaler`.
Kiểu đối số của `T.Equal` là `T`,
không phải kiểu bắt buộc `Equaler`.

Trong Go, hệ thống kiểu không tự động nâng cấp đối số của
`Equal`; đó là trách nhiệm của lập trình viên, như
minh họa bởi kiểu `T2`, thực sự triển khai
`Equaler`:

```
type T2 int
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }  // thỏa mãn Equaler
```

Ngay cả điều này cũng không giống các hệ thống kiểu khác, vì trong Go *bất kỳ*
kiểu nào thỏa mãn `Equaler` đều có thể được truyền như
đối số cho `T2.Equal`, và khi chạy chúng ta phải
kiểm tra rằng đối số có kiểu `T2`.
Một số ngôn ngữ sắp xếp để đảm bảo điều đó tại thời điểm biên dịch.

Một ví dụ liên quan đi theo chiều ngược lại:

```
type Opener interface {
   Open() Reader
}

func (t T3) Open() *os.File
```

Trong Go, `T3` không thỏa mãn `Opener`,
mặc dù nó có thể trong ngôn ngữ khác.

Mặc dù đúng là hệ thống kiểu của Go làm ít hơn cho lập trình viên
trong những trường hợp như vậy, nhưng việc thiếu subtyping làm cho các quy tắc về
sự thỏa mãn interface rất dễ phát biểu: tên và chữ ký của hàm có
khớp chính xác với interface không?
Quy tắc của Go cũng dễ triển khai hiệu quả.
Chúng tôi cảm thấy những lợi ích này bù đắp cho việc thiếu
tự động thăng cấp kiểu.

### Tôi có thể chuyển đổi []T thành []interface{} không? {#convert_slice_of_interface}

Không trực tiếp.
Điều này bị ngôn ngữ cấm vì hai kiểu
không có cùng biểu diễn trong bộ nhớ.
Cần phải sao chép các phần tử riêng lẻ sang slice đích.
Ví dụ này chuyển đổi một slice của `int` thành slice của
`interface{}`:

```
t := []int{1, 2, 3, 4}
s := make([]interface{}, len(t))
for i, v := range t {
    s[i] = v
}
```

### Tôi có thể chuyển đổi []T1 thành []T2 nếu T1 và T2 có cùng kiểu cơ bản không? {#convert_slice_with_same_underlying_type}

Dòng cuối của đoạn mã này không biên dịch được.

```
type T1 int
type T2 int
var t1 T1
var x = T2(t1) // OK
var st1 []T1
var sx = ([]T2)(st1) // NOT OK
```

Trong Go, các kiểu gắn chặt với phương thức, trong đó mọi kiểu có tên
có một tập phương thức (có thể rỗng).
Quy tắc chung là bạn có thể thay đổi tên của kiểu đang được
chuyển đổi (và do đó có thể thay đổi tập phương thức của nó) nhưng bạn không thể
thay đổi tên (và tập phương thức) của các phần tử của kiểu composite.
Go yêu cầu bạn phải rõ ràng về chuyển đổi kiểu.

### Tại sao giá trị nil error của tôi không bằng nil? {#nil_error}

Bên dưới, interface được triển khai như hai phần tử, kiểu `T`
và giá trị `V`.
`V` là một giá trị cụ thể như `int`,
`struct` hoặc con trỏ, không bao giờ là bản thân interface, và có
kiểu `T`.
Ví dụ, nếu chúng ta lưu giá trị `int` 3 vào một interface,
giá trị interface kết quả có, về mặt giản đồ,
(`T=int`, `V=3`).
Giá trị `V` còn được gọi là giá trị *động* của interface,
vì một biến interface nhất định có thể chứa các giá trị `V` khác nhau
(và các kiểu `T` tương ứng)
trong suốt quá trình thực thi chương trình.

Một giá trị interface là `nil` chỉ khi cả `V` và `T`
đều chưa được đặt, (`T=nil`, `V` chưa được đặt).
Đặc biệt, một interface `nil` sẽ luôn chứa kiểu `nil`.
Nếu chúng ta lưu một con trỏ `nil` có kiểu `*int` vào
một giá trị interface, kiểu bên trong sẽ là `*int` bất kể giá trị của con trỏ:
(`T=*int`, `V=nil`).
Một giá trị interface như vậy sẽ không phải `nil`
*ngay cả khi giá trị con trỏ `V` bên trong là* `nil`.

Tình huống này có thể gây nhầm lẫn, và phát sinh khi giá trị `nil` được
lưu bên trong giá trị interface như một trả về `error`:

```
func returnsError() error {
    var p *MyError = nil
    if bad() {
        p = ErrBad
    }
    return p // Luôn trả về lỗi không nil.
}
```

Nếu mọi thứ đều ổn, hàm trả về một `p` `nil`,
vì vậy giá trị trả về là một giá trị interface `error`
chứa (`T=*MyError`, `V=nil`).
Điều này có nghĩa là nếu người gọi so sánh lỗi trả về với `nil`,
nó sẽ luôn trông như thể có lỗi ngay cả khi không có gì xấu xảy ra.
Để trả về `nil` `error` đúng nghĩa cho người gọi,
hàm phải trả về một `nil` rõ ràng:

```
func returnsError() error {
    if bad() {
        return ErrBad
    }
    return nil
}
```

Thực hành tốt là các hàm
trả về lỗi luôn dùng kiểu `error` trong
chữ ký của chúng (như chúng ta đã làm ở trên) thay vì một kiểu cụ thể
như `*MyError`, để giúp đảm bảo lỗi được
tạo đúng cách. Ví dụ,
[`os.Open`](/pkg/os/#Open)
trả về một `error` mặc dù, nếu không phải `nil`,
nó luôn có kiểu cụ thể
[`*os.PathError`](/pkg/os/#PathError).

Các tình huống tương tự như được mô tả ở đây có thể phát sinh bất cứ khi nào interface được dùng.
Chỉ cần nhớ rằng nếu bất kỳ giá trị cụ thể nào
đã được lưu trong interface, interface sẽ không phải `nil`.
Để biết thêm thông tin, xem
[Luật phản chiếu](/doc/articles/laws_of_reflection.html).

### Tại sao các kiểu kích thước zero lại hoạt động kỳ lạ? {#zero_size_types}

Go hỗ trợ các kiểu kích thước zero, như struct không có trường
(`struct{}`) hoặc mảng không có phần tử (`[0]byte`).
Không có gì bạn có thể lưu vào một kiểu kích thước zero, nhưng các kiểu này
đôi khi hữu ích khi không cần giá trị, như trong
`map[int]struct{}` hoặc một kiểu có phương thức nhưng không có giá trị.

Các biến khác nhau có kiểu kích thước zero có thể được đặt tại cùng
vị trí trong bộ nhớ.
Điều này an toàn vì không có giá trị nào có thể được lưu trong các biến đó.

Hơn nữa, ngôn ngữ không đảm bảo liệu
con trỏ tới hai biến kích thước zero khác nhau sẽ bằng nhau hay không.
Các so sánh đó thậm chí có thể trả về `true` tại một điểm trong chương trình
và sau đó trả về `false` tại một điểm khác, tùy thuộc vào cách chính xác
chương trình được biên dịch và thực thi.

Một vấn đề riêng với các kiểu kích thước zero là con trỏ tới một trường struct kích thước zero
không được trùng với con trỏ tới một đối tượng khác trong bộ nhớ.
Điều đó có thể gây nhầm lẫn cho bộ gom rác.
Điều này có nghĩa là nếu trường cuối cùng trong một struct có kích thước zero, struct
sẽ được đệm để đảm bảo con trỏ tới trường cuối không
trùng với bộ nhớ ngay sau struct.
Do đó, chương trình này:

```
func main() {
	type S struct {
		f1 byte
		f2 struct{}
	}
	fmt.Println(unsafe.Sizeof(S{}))
}
```

sẽ in `2`, không phải `1`, trong hầu hết các triển khai Go.

### Tại sao không có union không gắn thẻ, như trong C? {#unions}

Union không gắn thẻ sẽ vi phạm đảm bảo an toàn bộ nhớ của Go.

### Tại sao Go không có kiểu biến thể? {#variant_types}

Kiểu biến thể, còn gọi là kiểu đại số, cung cấp cách chỉ định
rằng một giá trị có thể nhận một trong một tập hợp các kiểu khác, nhưng chỉ những kiểu đó.
Một ví dụ phổ biến trong lập trình hệ thống sẽ chỉ định rằng một
lỗi là, ví dụ, lỗi mạng, lỗi bảo mật hoặc lỗi ứng dụng
và cho phép người gọi phân biệt nguồn gốc của vấn đề
bằng cách kiểm tra kiểu của lỗi. Một ví dụ khác là cây cú pháp
trong đó mỗi nút có thể là kiểu khác nhau: khai báo, câu lệnh,
phép gán và tương tự.

Chúng tôi đã xem xét thêm kiểu biến thể vào Go, nhưng sau khi thảo luận
quyết định bỏ qua chúng vì chúng chồng chéo theo những cách gây nhầm lẫn
với interface. Điều gì sẽ xảy ra nếu các phần tử của kiểu biến thể
tự thân là interface?

Ngoài ra, một phần của những gì kiểu biến thể giải quyết đã được bao phủ bởi
ngôn ngữ. Ví dụ lỗi dễ biểu đạt bằng cách dùng giá trị interface
để giữ lỗi và type switch để phân biệt các trường hợp. Ví dụ
cây cú pháp cũng có thể thực hiện được, mặc dù không thanh lịch bằng.

### Tại sao Go không có kiểu kết quả covariant? {#covariant_types}

Kiểu kết quả covariant có nghĩa là một interface như

```
type Copyable interface {
    Copy() interface{}
}
```

sẽ được thỏa mãn bởi phương thức

```
func (v Value) Copy() Value
```

vì `Value` triển khai interface rỗng.
Trong Go, các kiểu phương thức phải khớp chính xác, vì vậy `Value` không
triển khai `Copyable`.
Go tách biệt khái niệm những gì một
kiểu làm -- các phương thức của nó -- khỏi triển khai của kiểu.
Nếu hai phương thức trả về các kiểu khác nhau, chúng không làm cùng một việc.
Lập trình viên muốn kiểu kết quả covariant thường đang cố gắng
biểu đạt phân cấp kiểu thông qua interface.
Trong Go, tự nhiên hơn là có sự tách biệt rõ ràng giữa interface
và triển khai.

## Giá trị {#Values}

### Tại sao Go không cung cấp chuyển đổi số ngầm định? {#conversions}

Sự tiện lợi của chuyển đổi tự động giữa các kiểu số trong C
bị lấn át bởi sự nhầm lẫn mà nó gây ra. Khi nào biểu thức là unsigned?
Giá trị lớn bao nhiêu? Có tràn không? Kết quả có portable không, độc lập
với máy thực thi?
Nó cũng làm phức tạp trình biên dịch; "chuyển đổi số học thông thường" của C
không dễ triển khai và không nhất quán giữa các kiến trúc.
Vì lý do portable, chúng tôi quyết định làm mọi thứ rõ ràng và đơn giản
với chi phí là một số chuyển đổi rõ ràng trong mã.
Định nghĩa về hằng số trong Go -- các giá trị độ chính xác tùy ý không có
chú thích về dấu và kích thước -- cải thiện vấn đề đáng kể, tuy nhiên.

Một chi tiết liên quan là, không giống C, `int` và `int64`
là các kiểu khác nhau ngay cả khi `int` là kiểu 64-bit. Kiểu `int`
là kiểu tổng quát; nếu bạn quan tâm đến số bit một số nguyên chứa, Go
khuyến khích bạn phải rõ ràng.

### Hằng số hoạt động như thế nào trong Go? {#constants}

Mặc dù Go nghiêm ngặt về chuyển đổi giữa các biến kiểu số khác nhau,
hằng số trong ngôn ngữ linh hoạt hơn nhiều.
Hằng số chữ như `23`, `3.14159`
và [`math.Pi`](/pkg/math/#pkg-constants)
chiếm một không gian số lý tưởng, với độ chính xác tùy ý và
không có tràn hay thiếu.
Ví dụ, giá trị của `math.Pi` được chỉ định đến 63 chữ số thập phân
trong mã nguồn, và các biểu thức hằng số liên quan đến giá trị giữ
độ chính xác vượt quá những gì `float64` có thể chứa.
Chỉ khi hằng số hoặc biểu thức hằng số được gán cho một
biến -- một vị trí bộ nhớ trong chương trình -- nó mới trở thành một số "máy tính" với
các thuộc tính và độ chính xác dấu phẩy động thông thường.

Ngoài ra,
vì chúng chỉ là các số, không phải giá trị có kiểu, hằng số trong Go có thể được
dùng tự do hơn các biến, do đó làm mềm đi một phần sự khó chịu
xung quanh các quy tắc chuyển đổi chặt chẽ.
Người ta có thể viết các biểu thức như

```
sqrt2 := math.Sqrt(2)
```

mà không bị trình biên dịch phàn nàn vì số lý tưởng `2`
có thể được chuyển đổi an toàn và chính xác
thành `float64` cho lời gọi `math.Sqrt`.

Bài đăng trên blog có tiêu đề [Hằng số](/blog/constants)
khám phá chủ đề này chi tiết hơn.

### Tại sao map được tích hợp sẵn? {#builtin_maps}

Cùng lý do với string: chúng là cấu trúc dữ liệu mạnh mẽ và quan trọng đến mức
cung cấp một triển khai xuất sắc với hỗ trợ cú pháp
làm cho lập trình dễ chịu hơn. Chúng tôi tin rằng triển khai map của Go
đủ mạnh để phục vụ phần lớn các trường hợp sử dụng.
Nếu một ứng dụng cụ thể có thể hưởng lợi từ một triển khai tùy chỉnh, hoàn toàn có thể
viết một cái nhưng sẽ không tiện về mặt cú pháp; đây là sự đánh đổi hợp lý.

### Tại sao map không cho phép slice làm key? {#map_keys}

Tra cứu map đòi hỏi toán tử bằng nhau, mà slice không triển khai.
Chúng không triển khai bằng nhau vì bằng nhau không được xác định rõ ràng trên những kiểu đó;
có nhiều cân nhắc liên quan đến so sánh nông vs. sâu, so sánh con trỏ vs.
giá trị, cách xử lý các kiểu đệ quy, và tương tự.
Chúng tôi có thể xem xét lại vấn đề này -- và triển khai bằng nhau cho slice
sẽ không làm mất hiệu lực bất kỳ chương trình hiện có nào -- nhưng thiếu ý tưởng rõ ràng
về bằng nhau của slice là gì, đơn giản hơn là bỏ qua nó hiện tại.

Bằng nhau được định nghĩa cho struct và mảng, vì vậy chúng có thể được dùng làm map key.

### Tại sao map, slice và channel là reference trong khi mảng là giá trị? {#references}

Có nhiều lịch sử về chủ đề đó. Lúc đầu, map và channel
là các con trỏ về mặt cú pháp và không thể khai báo hoặc dùng một
instance không phải con trỏ. Ngoài ra, chúng tôi đã vật lộn với cách mảng nên hoạt động.
Cuối cùng chúng tôi quyết định rằng sự tách biệt nghiêm ngặt giữa con trỏ và
giá trị làm ngôn ngữ khó dùng hơn. Thay đổi các kiểu này
để hoạt động như tham chiếu tới cấu trúc dữ liệu chia sẻ liên quan đã giải quyết
những vấn đề này. Thay đổi này thêm vào một chút phức tạp đáng tiếc cho
ngôn ngữ nhưng có tác động lớn đến khả năng sử dụng: Go trở nên một ngôn ngữ
năng suất và thoải mái hơn khi nó được giới thiệu.

## Viết Mã {#Writing_Code}

### Thư viện được tài liệu hóa như thế nào? {#How_are_libraries_documented}

Để truy cập tài liệu từ dòng lệnh, công cụ
[go](/pkg/cmd/go/) có lệnh con
[doc](/pkg/cmd/go/#hdr-Show_documentation_for_package_or_symbol)
cung cấp giao diện văn bản cho tài liệu
về khai báo, tệp, package và tương tự.

Trang khám phá package toàn cầu
[pkg.go.dev/pkg/](/pkg/).
chạy một máy chủ trích xuất tài liệu package từ mã nguồn Go
ở bất cứ đâu trên web
và phục vụ nó dưới dạng HTML với các liên kết đến khai báo và các phần tử liên quan.
Đây là cách dễ nhất để tìm hiểu về các thư viện Go hiện có.

Trong những ngày đầu của dự án, có một chương trình tương tự, `godoc`,
cũng có thể chạy để trích xuất tài liệu cho các tệp trên máy cục bộ;
[pkg.go.dev/pkg/](/pkg/) về cơ bản là hậu duệ của nó.
Hậu duệ khác là lệnh
[`pkgsite`](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite)
có thể, giống như `godoc`, chạy cục bộ, mặc dù
nó chưa được tích hợp vào
kết quả hiển thị bởi `go` `doc`.

### Có hướng dẫn phong cách lập trình Go không? {#Is_there_a_Go_programming_style_guide}

Không có hướng dẫn phong cách rõ ràng, mặc dù chắc chắn
có "phong cách Go" có thể nhận ra.

Go đã thiết lập các quy ước để hướng dẫn các quyết định về
đặt tên, bố cục và tổ chức tệp.
Tài liệu [Effective Go](effective_go.html)
chứa một số lời khuyên về các chủ đề này.
Trực tiếp hơn, chương trình `gofmt` là một pretty-printer
có mục đích áp đặt các quy tắc bố cục; nó thay thế
tập hợp thông thường của các điều nên và không nên làm để giải thích.
Tất cả mã Go trong kho lưu trữ, và phần lớn trong thế giới mã nguồn mở, đã được chạy qua `gofmt`.

Tài liệu có tên
[Nhận xét đánh giá mã Go](/s/comments)
là tập hợp các bài luận ngắn về chi tiết của Go idiom thường bị
lập trình viên bỏ qua.
Đây là tài liệu tham khảo hữu ích cho những người đang đánh giá mã cho các dự án Go.

### Làm thế nào để tôi gửi patch cho thư viện Go? {#How_do_I_submit_patches_to_the_Go_libraries}

Mã nguồn thư viện nằm trong thư mục `src` của kho lưu trữ.
Nếu bạn muốn thực hiện một thay đổi đáng kể, vui lòng thảo luận trên mailing list trước khi bắt đầu.

Xem tài liệu
[Đóng góp cho dự án Go](contribute.html)
để biết thêm thông tin về cách tiến hành.

### Tại sao "go get" dùng HTTPS khi clone kho lưu trữ? {#git_https}

Các công ty thường chỉ cho phép lưu lượng ra trên các cổng TCP chuẩn 80 (HTTP)
và 443 (HTTPS), chặn lưu lượng ra trên các cổng khác, bao gồm cổng TCP 9418
(git) và cổng TCP 22 (SSH).
Khi dùng HTTPS thay vì HTTP, `git` áp đặt xác thực chứng chỉ theo
mặc định, cung cấp bảo vệ chống lại các cuộc tấn công man-in-the-middle, nghe lén và giả mạo.
Do đó, lệnh `go get` dùng HTTPS để an toàn.

`Git` có thể được cấu hình để xác thực qua HTTPS hoặc để dùng SSH thay cho HTTPS.
Để xác thực qua HTTPS, bạn có thể thêm một dòng
vào tệp `$HOME/.netrc` mà git tham chiếu:

```
machine github.com login *USERNAME* password *APIKEY*
```

Đối với tài khoản GitHub, password có thể là một
[personal access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/).

`Git` cũng có thể được cấu hình để dùng SSH thay cho HTTPS cho các URL khớp với tiền tố nhất định.
Ví dụ, để dùng SSH cho tất cả truy cập GitHub,
thêm các dòng này vào `~/.gitconfig` của bạn:

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

Khi làm việc với các module riêng tư nhưng dùng proxy module công khai cho dependency, bạn có thể cần đặt `GOPRIVATE`.
Xem [module riêng tư](/ref/mod#private-modules) để biết chi tiết và các cài đặt bổ sung.

### Làm thế nào để quản lý phiên bản package bằng "go get"? {#get_version}

Go toolchain có một hệ thống tích hợp để quản lý các tập package có phiên bản liên quan, được gọi là *module*.
Module được giới thiệu trong [Go 1.11](/doc/go1.11#modules) và đã sẵn sàng cho môi trường production kể từ [1.14](/doc/go1.14#introduction).

Để tạo một dự án dùng module, chạy [`go mod init`](/ref/mod#go-mod-init).
Lệnh này tạo tệp `go.mod` theo dõi các phiên bản dependency.

```
go mod init example/project
```

Để thêm, nâng cấp hoặc hạ cấp một dependency, chạy [`go get`](/ref/mod#go-get):

```
go get golang.org/x/text@v0.3.5
```

Xem [Hướng dẫn: Tạo module](/doc/tutorial/create-module.html) để biết thêm thông tin về bắt đầu.

Xem [Phát triển module](/doc/#developing-modules) để biết các hướng dẫn quản lý dependency với module.

Các package trong module nên duy trì tương thích ngược khi phát triển, theo dõi [quy tắc tương thích import](https://research.swtch.com/vgo-import):

> Nếu một package cũ và một package mới có cùng đường dẫn import,\
> package mới phải tương thích ngược với package cũ.

[Hướng dẫn tương thích Go 1](/doc/go1compat.html) là tài liệu tham khảo tốt ở đây:
không xóa tên đã xuất, khuyến khích literal composite có thẻ, và tương tự.

Nếu cần chức năng khác, hãy thêm tên mới thay vì thay đổi tên cũ.

Module mã hóa điều này với [semantic versioning](https://semver.org/) và semantic import versioning.
Nếu cần phá vỡ tương thích, hãy phát hành module ở phiên bản chính mới.
Module ở phiên bản chính 2 trở lên yêu cầu [hậu tố phiên bản chính](/ref/mod#major-version-suffixes) như một phần đường dẫn của chúng (như `/v2`).
Điều này bảo toàn quy tắc tương thích import: các package trong các phiên bản chính khác nhau của một module có đường dẫn khác nhau.

## Con trỏ và Phân bổ bộ nhớ {#Pointers}

### Khi nào tham số hàm được truyền theo giá trị? {#pass_by_value}

Như trong tất cả các ngôn ngữ thuộc họ C, mọi thứ trong Go đều được truyền theo giá trị.
Tức là, một hàm luôn nhận một bản sao của
thứ được truyền, như thể có một câu lệnh gán giá trị
cho tham số. Ví dụ, truyền một giá trị `int`
cho một hàm tạo ra một bản sao của `int`, và truyền một
giá trị con trỏ tạo ra một bản sao của con trỏ, nhưng không phải dữ liệu mà nó trỏ tới.
(Xem [phần sau](/doc/faq#methods_on_values_or_pointers)
để thảo luận về cách điều này ảnh hưởng đến receiver phương thức.)

Giá trị map và slice hoạt động như con trỏ: chúng là các mô tả chứa
con trỏ tới dữ liệu map hoặc slice cơ bản. Sao chép một map hoặc
giá trị slice không sao chép dữ liệu mà nó trỏ tới. Sao chép giá trị interface
tạo một bản sao của thứ được lưu trong giá trị interface. Nếu giá trị interface
giữ một struct, sao chép giá trị interface tạo một bản sao của struct đó.
Nếu giá trị interface giữ một con trỏ, sao chép giá trị interface
tạo một bản sao của con trỏ, nhưng một lần nữa không phải dữ liệu mà nó trỏ tới.

Lưu ý rằng thảo luận này là về ngữ nghĩa của các thao tác.
Các triển khai thực tế có thể áp dụng tối ưu hóa để tránh sao chép
miễn là các tối ưu hóa không thay đổi ngữ nghĩa.

### Khi nào tôi nên dùng con trỏ tới một interface? {#pointer_to_interface}

Hầu như không bao giờ. Con trỏ tới giá trị interface chỉ phát sinh trong những tình huống
hiếm gặp, phức tạp liên quan đến việc che giấu kiểu của giá trị interface để đánh giá trễ.

Đây là một lỗi phổ biến khi truyền con trỏ tới giá trị interface
cho một hàm mong đợi một interface. Trình biên dịch sẽ phàn nàn về
lỗi này nhưng tình huống vẫn có thể gây nhầm lẫn, vì đôi khi một
[con trỏ là cần thiết để thỏa mãn một interface](#different_method_sets).
Điểm mấu chốt là mặc dù con trỏ tới một kiểu cụ thể có thể thỏa mãn
một interface, với một ngoại lệ *một con trỏ tới interface không bao giờ có thể thỏa mãn interface*.

Xem xét khai báo biến,

```
var w io.Writer
```

Hàm in `fmt.Fprintf` nhận như đối số đầu tiên
một giá trị thỏa mãn `io.Writer` -- thứ gì đó triển khai
phương thức `Write` chuẩn. Vì vậy chúng ta có thể viết

```
fmt.Fprintf(w, "hello, world\n")
```

Tuy nhiên nếu chúng ta truyền địa chỉ của `w`, chương trình sẽ không biên dịch được.

```
fmt.Fprintf(&w, "hello, world\n") // Lỗi lúc biên dịch.
```

Ngoại lệ duy nhất là bất kỳ giá trị nào, kể cả con trỏ tới interface, có thể được gán cho
một biến của kiểu interface rỗng (`interface{}`).
Dù vậy, gần như chắc chắn là lỗi nếu giá trị là con trỏ tới interface;
kết quả có thể gây nhầm lẫn.

### Tôi nên định nghĩa phương thức trên giá trị hay con trỏ? {#methods_on_values_or_pointers}

```
func (s *MyStruct) pointerMethod() { } // phương thức trên con trỏ
func (s MyStruct)  valueMethod()   { } // phương thức trên giá trị
```

Đối với lập trình viên chưa quen với con trỏ, sự khác biệt giữa hai
ví dụ này có thể gây nhầm lẫn, nhưng tình huống thực sự rất đơn giản.
Khi định nghĩa một phương thức trên một kiểu, receiver (`s` trong các
ví dụ trên) hoạt động giống như thể nó là một đối số cho phương thức.
Liệu có nên định nghĩa receiver là giá trị hay con trỏ là câu hỏi tương tự,
như liệu đối số hàm nên là giá trị hay con trỏ.
Có một số cân nhắc.

Đầu tiên, và quan trọng nhất, phương thức có cần sửa đổi
receiver không?
Nếu có, receiver *phải* là con trỏ.
(Slice và map hoạt động như reference, vì vậy câu chuyện của chúng phức tạp hơn một chút,
nhưng ví dụ để thay đổi độ dài của một slice
trong một phương thức, receiver vẫn phải là con trỏ.)
Trong các ví dụ trên, nếu `pointerMethod` sửa đổi
các trường của `s`,
người gọi sẽ thấy những thay đổi đó, nhưng `valueMethod`
được gọi với một bản sao của đối số của người gọi (đó là định nghĩa
của truyền theo giá trị), vì vậy các thay đổi nó thực hiện sẽ không thấy được với người gọi.

Nhân tiện, trong Java receiver phương thức luôn là con trỏ,
mặc dù bản chất con trỏ của chúng bị che khuất phần nào
(và các phát triển gần đây đang mang lại value receiver cho Java).
Đó là value receiver trong Go mới là điều bất thường.

Thứ hai là cân nhắc về hiệu quả. Nếu receiver lớn,
chẳng hạn một `struct` lớn, có thể rẻ hơn khi
dùng pointer receiver.

Tiếp theo là tính nhất quán. Nếu một số phương thức của kiểu phải có
pointer receiver, các phương thức còn lại cũng nên có, để tập phương thức nhất quán
bất kể kiểu được dùng như thế nào.
Xem phần về [tập phương thức](#different_method_sets)
để biết chi tiết.

Đối với các kiểu như kiểu cơ bản, slice và `struct` nhỏ,
value receiver rất rẻ, vì vậy trừ khi ngữ nghĩa của phương thức
yêu cầu con trỏ, value receiver là hiệu quả và rõ ràng.

### Sự khác biệt giữa new và make là gì? {#new_and_make}

Tóm lại: `new` cấp phát bộ nhớ, trong khi `make` khởi tạo
các kiểu slice, map và channel.

Xem [phần liên quan
của Effective Go](/doc/effective_go.html#allocation_new) để biết thêm chi tiết.

### Kích thước của `int` trên máy 64 bit là bao nhiêu? {#q_int_sizes}

Kích thước của `int` và `uint` phụ thuộc vào triển khai
nhưng giống nhau trên một nền tảng nhất định.
Để đảm bảo portable, mã phụ thuộc vào kích thước giá trị cụ thể nên dùng kiểu có kích thước rõ ràng, như `int64`.
Trên máy 32-bit, trình biên dịch dùng số nguyên 32-bit theo mặc định,
trong khi trên máy 64-bit, số nguyên có 64 bit.
(Về mặt lịch sử, điều này không phải lúc nào cũng đúng.)

Mặt khác, các scalar dấu phẩy động và kiểu phức tạp
luôn có kích thước cụ thể (không có kiểu cơ bản `float` hoặc `complex`),
vì lập trình viên nên biết về độ chính xác khi dùng số dấu phẩy động.
Kiểu mặc định được dùng cho hằng số dấu phẩy động (không có kiểu) là `float64`.
Do đó `foo` `:=` `3.0` khai báo biến `foo`
có kiểu `float64`.
Đối với biến `float32` được khởi tạo bởi hằng số (không có kiểu), kiểu biến
phải được chỉ định rõ ràng trong khai báo biến:

```
var foo float32 = 3.0
```

Ngoài ra, hằng số phải được đặt kiểu bằng chuyển đổi như
`foo := float32(3.0)`.

### Làm thế nào để tôi biết biến được cấp phát trên heap hay stack? {#stack_or_heap}

Từ quan điểm đúng đắn, bạn không cần biết.
Mỗi biến trong Go tồn tại miễn là có tham chiếu tới nó.
Vị trí lưu trữ được trình biên dịch chọn không liên quan đến
ngữ nghĩa của ngôn ngữ.

Vị trí lưu trữ có ảnh hưởng đến việc viết chương trình hiệu quả.
Khi có thể, trình biên dịch Go sẽ cấp phát các biến cục bộ
cho một hàm trong stack frame của hàm đó. Tuy nhiên, nếu
trình biên dịch không thể chứng minh rằng biến không được tham chiếu sau khi
hàm trả về, thì trình biên dịch phải cấp phát biến trên
heap được gom rác để tránh lỗi dangling pointer.
Ngoài ra, nếu một biến cục bộ rất lớn, có thể hợp lý hơn
để lưu nó trên heap thay vì stack.

Trong các trình biên dịch hiện tại, nếu một biến có địa chỉ được lấy, biến đó
là ứng viên để cấp phát trên heap. Tuy nhiên, một phân tích *escape*
cơ bản nhận ra một số trường hợp khi những biến như vậy sẽ không
sống qua lần trả về của hàm và có thể nằm trên stack.

### Tại sao process Go của tôi dùng nhiều bộ nhớ ảo như vậy? {#Why_does_my_Go_process_use_so_much_virtual_memory}

Bộ cấp phát bộ nhớ Go dành riêng một vùng lớn bộ nhớ ảo như một arena
để cấp phát. Bộ nhớ ảo này là cục bộ cho process Go cụ thể;
sự đặt chỗ không lấy đi bộ nhớ của các process khác.

Để tìm lượng bộ nhớ thực sự được cấp phát cho một process Go, hãy dùng lệnh Unix
`top` và tham khảo các cột `RES` (Linux) hoặc
`RSIZE` (macOS).
<!-- TODO(adg): find out how this works on Windows -->

## Concurrency {#Concurrency}

### Thao tác nào là atomic? Còn mutex thì sao? {#What_operations_are_atomic_What_about_mutexes}

Mô tả về tính atomic của các thao tác trong Go có thể tìm thấy trong
tài liệu [Mô hình bộ nhớ Go](/ref/mem).

Đồng bộ hóa cấp thấp và các nguyên thủy atomic có sẵn trong các package
[sync](/pkg/sync) và
[sync/atomic](/pkg/sync/atomic).
Các package này phù hợp cho các tác vụ đơn giản như tăng
số đếm tham chiếu hoặc đảm bảo loại trừ lẫn nhau ở quy mô nhỏ.

Đối với các thao tác cấp cao hơn, chẳng hạn phối hợp giữa
các server concurrent, các kỹ thuật cấp cao hơn có thể dẫn đến
các chương trình đẹp hơn, và Go hỗ trợ cách tiếp cận này thông qua
goroutine và channel.
Ví dụ, bạn có thể cấu trúc chương trình để chỉ một
goroutine tại một thời điểm chịu trách nhiệm về một dữ liệu cụ thể.
Cách tiếp cận đó được tóm tắt bởi
[Go proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c) gốc,

Đừng giao tiếp bằng cách chia sẻ bộ nhớ. Thay vào đó, chia sẻ bộ nhớ bằng cách giao tiếp.

Xem [Chia sẻ bộ nhớ bằng cách giao tiếp](/doc/codewalk/sharemem/) code walk
và [bài viết liên quan](/blog/share-memory-by-communicating)
để thảo luận chi tiết về khái niệm này.

Các chương trình concurrent lớn có thể vay mượn từ cả hai bộ công cụ này.

### Tại sao chương trình của tôi không chạy nhanh hơn với nhiều CPU? {#parallel_slow}

Liệu một chương trình có chạy nhanh hơn với nhiều CPU phụ thuộc vào vấn đề
nó đang giải quyết.
Ngôn ngữ Go cung cấp các nguyên thủy concurrency, như goroutine
và channel, nhưng concurrency chỉ cho phép song song
khi vấn đề cơ bản vốn dĩ là song song.
Các vấn đề vốn dĩ tuần tự không thể tăng tốc bằng cách thêm
nhiều CPU, trong khi những vấn đề có thể được chia thành các phần
thực thi song song có thể được tăng tốc, đôi khi đáng kể.

Đôi khi thêm nhiều CPU có thể làm chậm chương trình.
Trong thực tế, các chương trình dành nhiều thời gian hơn
đồng bộ hóa hoặc giao tiếp hơn là thực hiện tính toán hữu ích
có thể trải qua suy giảm hiệu suất khi dùng
nhiều OS thread.
Điều này là do truyền dữ liệu giữa các thread liên quan đến chuyển
ngữ cảnh, có chi phí đáng kể, và chi phí đó có thể tăng
với nhiều CPU hơn.
Ví dụ, [ví dụ sàng số nguyên tố](/ref/spec#An_example_package)
từ đặc tả Go không có song song đáng kể mặc dù nó khởi chạy nhiều
goroutine; tăng số thread (CPU) có nhiều khả năng làm chậm hơn là
tăng tốc.

Để biết thêm chi tiết về chủ đề này, xem bài nói chuyện có tiêu đề
[Concurrency không phải là Parallelism](/blog/concurrency-is-not-parallelism).

### Làm thế nào để kiểm soát số lượng CPU? {#number_cpus}

Số lượng CPU có sẵn đồng thời để thực thi goroutine được
kiểm soát bởi biến môi trường shell `GOMAXPROCS`,
giá trị mặc định của nó là số lõi CPU có sẵn.
Các chương trình có khả năng thực thi song song do đó
sẽ đạt được điều đó theo mặc định trên máy nhiều CPU.
Để thay đổi số CPU song song cần dùng,
đặt biến môi trường hoặc dùng
[hàm](/pkg/runtime/#GOMAXPROCS)
cùng tên của package runtime để cấu hình
hỗ trợ run-time để sử dụng số lượng thread khác.
Đặt nó thành 1 loại bỏ khả năng song song thực sự,
buộc các goroutine độc lập phải thay nhau thực thi.

Runtime có thể cấp phát nhiều thread hơn giá trị
của `GOMAXPROCS` để phục vụ nhiều yêu cầu
I/O đồng thời.
`GOMAXPROCS` chỉ ảnh hưởng đến bao nhiêu goroutine
có thể thực sự thực thi cùng một lúc; nhiều hơn tùy ý có thể bị chặn
trong các lời gọi hệ thống.

Bộ lập lịch goroutine của Go hoạt động tốt trong việc cân bằng goroutine
và thread, và thậm chí có thể ngắt thực thi của một goroutine
để đảm bảo các goroutine khác trên cùng thread không bị đói.
Tuy nhiên, nó không hoàn hảo.
Nếu bạn thấy vấn đề hiệu suất,
đặt `GOMAXPROCS` trên cơ sở mỗi ứng dụng có thể giúp ích.

### Tại sao không có goroutine ID? {#no_goroutine_id}

Goroutine không có tên; chúng chỉ là các worker ẩn danh.
Chúng không expose định danh duy nhất, tên hoặc cấu trúc dữ liệu cho lập trình viên.
Một số người ngạc nhiên về điều này, mong đợi câu lệnh `go`
trả về một mục nào đó có thể được dùng để truy cập và kiểm soát
goroutine sau này.

Lý do cơ bản goroutine ẩn danh là để
toàn bộ ngôn ngữ Go có sẵn khi lập trình mã concurrent.
Ngược lại, các mẫu sử dụng phát triển khi thread và goroutine
được đặt tên có thể hạn chế những gì một thư viện dùng chúng có thể làm.

Đây là một minh họa về những khó khăn.
Một khi ai đó đặt tên cho một goroutine và xây dựng một mô hình xung quanh
nó, nó trở nên đặc biệt, và người ta bị cám dỗ để liên kết tất cả tính toán
với goroutine đó, bỏ qua khả năng
dùng nhiều goroutine, có thể chia sẻ để xử lý.
Nếu package `net/http` liên kết trạng thái mỗi-yêu cầu
với một goroutine,
client sẽ không thể dùng nhiều goroutine hơn
khi phục vụ một yêu cầu.

Hơn nữa, kinh nghiệm với các thư viện như những thư viện cho hệ thống đồ họa
yêu cầu tất cả xử lý xảy ra trên "main thread"
đã cho thấy cách tiếp cận có thể bất tiện và hạn chế như thế nào khi
triển khai trong một ngôn ngữ concurrent.
Sự tồn tại của một thread hoặc goroutine đặc biệt
buộc lập trình viên phải biến dạng chương trình để tránh các sự cố
và các vấn đề khác gây ra bởi việc vô tình hoạt động
trên thread sai.

Đối với những trường hợp mà một goroutine cụ thể thực sự đặc biệt,
ngôn ngữ cung cấp các tính năng như channel có thể được
dùng theo những cách linh hoạt để tương tác với nó.

## Hàm và Phương thức {#Functions_methods}

### Tại sao T và *T có tập phương thức khác nhau? {#different_method_sets}

Như [đặc tả Go](/ref/spec#Types) nói,
tập phương thức của kiểu `T` bao gồm tất cả các phương thức
với kiểu receiver `T`,
trong khi kiểu con trỏ tương ứng
`*T` bao gồm tất cả các phương thức có receiver `*T` hoặc
`T`.
Điều đó có nghĩa là tập phương thức của `*T`
bao gồm của `T`,
nhưng không ngược lại.

Sự phân biệt này phát sinh vì
nếu một giá trị interface chứa một con trỏ `*T`,
một lời gọi phương thức có thể lấy giá trị bằng cách dereference con trỏ,
nhưng nếu một giá trị interface chứa một giá trị `T`,
không có cách an toàn nào để một lời gọi phương thức lấy con trỏ.
(Làm như vậy sẽ cho phép một phương thức sửa đổi nội dung của
giá trị bên trong interface, điều này không được phép bởi
đặc tả ngôn ngữ.)

Ngay cả trong các trường hợp mà trình biên dịch có thể lấy địa chỉ của một giá trị
để truyền cho phương thức, nếu phương thức sửa đổi giá trị thì các thay đổi
sẽ bị mất ở người gọi.

Ví dụ, nếu mã dưới đây hợp lệ:

```
var buf bytes.Buffer
io.Copy(buf, os.Stdin)
```

nó sẽ sao chép standard input vào một *bản sao* của `buf`,
không phải vào chính `buf`.
Đây gần như không bao giờ là hành vi mong muốn và do đó bị ngôn ngữ cấm.

### Điều gì xảy ra với closure chạy như goroutine? {#closures_and_goroutines}

Do cách biến vòng lặp hoạt động, trước Go phiên bản 1.22 (xem
cuối phần này để biết cập nhật),
một số nhầm lẫn có thể phát sinh khi dùng closure với concurrency.
Xem xét chương trình sau:

{{raw `
<pre>
func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done &lt;- true
        }()
    }

    // wait for all goroutines to complete before exiting
    for _ = range values {
        &lt;-done
    }
}
</pre>
`}}

Người ta có thể nhầm mong đợi thấy `a, b, c` là đầu ra.
Thực tế bạn sẽ thấy là `c, c, c`. Điều này là vì
mỗi lần lặp vòng lặp dùng cùng một instance của biến `v`, vì vậy
mỗi closure chia sẻ biến đơn đó. Khi closure chạy, nó in
giá trị của `v` vào thời điểm `fmt.Println` được thực thi,
nhưng `v` có thể đã được sửa đổi kể từ khi goroutine được khởi chạy.
Để giúp phát hiện điều này và các vấn đề khác trước khi chúng xảy ra, hãy chạy
[`go vet`](/cmd/go/#hdr-Run_go_tool_vet_on_packages).

Để ràng buộc giá trị hiện tại của `v` với mỗi closure khi nó được khởi chạy,
người ta phải sửa đổi vòng lặp bên trong để tạo một biến mới cho mỗi lần lặp.
Một cách là truyền biến như đối số cho closure:

{{raw `
<pre>
    for _, v := range values {
        go func(<b>u</b> string) {
            fmt.Println(<b>u</b>)
            done &lt;- true
        }(<b>v</b>)
    }
</pre>
`}}

Trong ví dụ này, giá trị của `v` được truyền như đối số cho
hàm ẩn danh. Giá trị đó sau đó có thể truy cập được bên trong hàm như
biến `u`.

Thậm chí dễ hơn là chỉ tạo một biến mới, dùng phong cách khai báo có thể
trông kỳ lạ nhưng hoạt động tốt trong Go:

{{raw `
<pre>
    for _, v := range values {
        <b>v := v</b> // tạo 'v' mới.
        go func() {
            fmt.Println(<b>v</b>)
            done &lt;- true
        }()
    }
</pre>
`}}

Hành vi này của ngôn ngữ, không định nghĩa biến mới cho
mỗi lần lặp, được coi là lỗi nhìn lại,
và đã được giải quyết trong [Go 1.22](/wiki/LoopvarExperiment), phiên bản này
thực sự tạo một biến mới cho mỗi lần lặp, loại bỏ vấn đề này.

## Luồng điều khiển {#Control_flow}

### Tại sao Go không có toán tử `?:`? {#Does_Go_have_a_ternary_form}

Không có thao tác kiểm tra ternary trong Go.
Bạn có thể dùng cách sau để đạt được kết quả tương tự:

```
if expr {
    n = trueVal
} else {
    n = falseVal
}
```

Lý do `?:` vắng mặt khỏi Go là người thiết kế ngôn ngữ
đã thấy thao tác này được dùng quá thường xuyên để tạo ra các biểu thức phức tạp đến khó hiểu.
Dạng `if-else`, mặc dù dài hơn,
rõ ràng hơn không nghi ngờ gì.
Một ngôn ngữ chỉ cần một cấu trúc điều khiển luồng điều kiện.

## Tham số kiểu {#Type_Parameters}

### Tại sao Go có tham số kiểu? {#why_generics}

Tham số kiểu cho phép những gì được gọi là lập trình generic, trong đó
các hàm và cấu trúc dữ liệu được định nghĩa theo các kiểu được
chỉ định sau, khi những hàm và cấu trúc dữ liệu đó được dùng.
Ví dụ, chúng giúp có thể viết một hàm trả về
giá trị nhỏ nhất của hai giá trị thuộc bất kỳ kiểu có thứ tự nào, mà không cần viết
một phiên bản riêng biệt cho mỗi kiểu có thể.
Để giải thích sâu hơn với các ví dụ, xem bài đăng trên blog
[Tại sao có Generics?](/blog/why-generics).

### Generics được triển khai như thế nào trong Go? {#generics_implementation}

Trình biên dịch có thể chọn liệu có biên dịch mỗi instantiation
riêng biệt hay biên dịch các instantiation tương tự như
một triển khai duy nhất.
Cách tiếp cận triển khai đơn giống như một hàm với một
tham số interface.
Các trình biên dịch khác nhau sẽ đưa ra lựa chọn khác nhau cho các trường hợp khác nhau.
Trình biên dịch Go chuẩn thường emit một instantiation duy nhất
cho mọi đối số kiểu có cùng hình dạng, trong đó hình dạng được
xác định bởi các thuộc tính của kiểu như kích thước và vị trí
của các con trỏ mà nó chứa.
Các bản phát hành tương lai có thể thử nghiệm với sự đánh đổi giữa thời gian biên dịch,
hiệu quả run-time và kích thước mã.

### Generics trong Go so sánh với generics trong các ngôn ngữ khác như thế nào? {#generics_comparison}

Chức năng cơ bản trong tất cả các ngôn ngữ là tương tự: có thể
viết các kiểu và hàm dùng các kiểu được chỉ định sau.
Tuy nhiên, có một số khác biệt.

* Java

	Trong Java, trình biên dịch kiểm tra các kiểu generic tại thời điểm biên dịch nhưng xóa
	các kiểu khi chạy.
	Điều này được gọi là
	[type erasure](https://en.wikipedia.org/wiki/Generics_in_Java#Problems_with_type_erasure).
	Ví dụ, một kiểu Java được gọi là `List<Integer>` tại
	thời điểm biên dịch sẽ trở thành kiểu không generic `List` khi
	chạy.
	Điều này có nghĩa là, ví dụ, khi dùng dạng Java của type
	reflection, không thể phân biệt một giá trị của
	kiểu `List<Integer>` khỏi giá trị của
	kiểu `List<Float>`.
	Trong Go, thông tin reflection cho một kiểu generic bao gồm đầy đủ
	thông tin kiểu tại thời điểm biên dịch.

	Java dùng wildcard kiểu như {{raw `<code>List&lt;? extends Number&gt;</code>`}}
	hoặc {{raw `<code>List<? super Number></code>`}} để triển khai generic
	covariance và contravariance.
	Go không có các khái niệm này, làm cho các kiểu generic trong Go
	đơn giản hơn nhiều.

* C++

	Theo truyền thống, C++ template không áp đặt bất kỳ ràng buộc nào lên
	đối số kiểu, mặc dù C++20 hỗ trợ các ràng buộc tùy chọn thông qua
	[concepts](https://en.wikipedia.org/wiki/Concepts_(C%2B%2B)).
	Trong Go, các ràng buộc là bắt buộc cho tất cả tham số kiểu.
	C++20 concepts được biểu đạt như các đoạn mã nhỏ phải biên dịch
	với các đối số kiểu.
	Go constraints là các kiểu interface định nghĩa tập hợp tất cả
	các đối số kiểu được phép.

	C++ hỗ trợ template metaprogramming; Go thì không.
	Trong thực tế, tất cả trình biên dịch C++ biên dịch mỗi template tại điểm
	nó được instantiate; như đã lưu ý trên, Go có thể và sử dụng
	các cách tiếp cận khác nhau cho các instantiation khác nhau.

* Rust

	Phiên bản Rust của constraints được gọi là trait bounds.
	Trong Rust, mối liên kết giữa một trait bound và một kiểu phải được
	định nghĩa rõ ràng, hoặc trong crate định nghĩa trait bound
	hoặc crate định nghĩa kiểu.
	Trong Go, các đối số kiểu ngầm định thỏa mãn constraints, giống như các kiểu Go
	ngầm định triển khai các kiểu interface.
	Thư viện chuẩn Rust định nghĩa các trait chuẩn cho các thao tác như
	so sánh hoặc cộng; thư viện chuẩn Go không, vì những thứ này có thể
	được biểu đạt trong mã người dùng thông qua các kiểu interface. Ngoại lệ duy nhất
	là interface được khai báo trước `comparable` của Go, nắm bắt
	một thuộc tính không thể biểu đạt trong hệ thống kiểu.

* Python

	Python không phải là ngôn ngữ kiểu tĩnh, vì vậy người ta có thể nói hợp lý
	rằng tất cả các hàm Python luôn là generic theo mặc định: chúng luôn có thể
	được gọi với các giá trị của bất kỳ kiểu nào, và bất kỳ lỗi kiểu nào được
	phát hiện khi chạy.

### Tại sao Go dùng dấu ngoặc vuông cho danh sách tham số kiểu? {#generic_brackets}

Java và C++ dùng dấu ngoặc nhọn cho danh sách tham số kiểu, như
Java `List<Integer>` và C++
`std::vector<int>`.
Tuy nhiên, tùy chọn đó không có sẵn cho Go, vì nó dẫn đến
một vấn đề cú pháp: khi phân tích mã trong một hàm, chẳng hạn
như `v := F<T>`, tại điểm nhìn thấy
{{raw `<code>&lt;</code>`}}, không rõ liệu chúng ta đang thấy một
instantiation hay một biểu thức dùng toán tử {{raw `<code>&lt;</code>`}}.
Điều này rất khó giải quyết mà không có thông tin kiểu.

Ví dụ, hãy xem xét một câu lệnh như

{{raw `
<pre>
    a, b = w &lt; x, y &gt; (z)
</pre>
`}}

Không có thông tin kiểu, không thể quyết định liệu vế phải
của phép gán là một cặp biểu thức ({{raw `<code>w &lt; x</code>`}}
và {{raw `<code>y &gt; z</code>`}}), hay liệu nó là một generic function
instantiation và gọi trả về hai kết quả
({{raw `<code>(w&lt;x, y&gt;)(z)</code>`}}).

Đây là quyết định thiết kế then chốt của Go rằng việc phân tích phải có thể thực hiện
mà không cần thông tin kiểu, điều có vẻ không thể khi dùng dấu ngoặc nhọn cho
generics.

Go không phải là duy nhất hay nguyên bản khi dùng dấu ngoặc vuông; có các ngôn ngữ khác
như Scala cũng dùng dấu ngoặc vuông cho mã generic.

### Tại sao Go không hỗ trợ phương thức với tham số kiểu? {#generic_methods}

Go cho phép một kiểu generic có phương thức, nhưng, ngoài
receiver, các đối số cho những phương thức đó không thể dùng các kiểu được tham số hóa.
Chúng tôi không dự đoán rằng Go sẽ bao giờ thêm generic methods.

Vấn đề là cách triển khai chúng.
Cụ thể, hãy xem xét việc kiểm tra xem một giá trị trong một
interface có triển khai một interface khác với các phương thức bổ sung không.
Ví dụ, hãy xem xét kiểu này, một struct rỗng với một
phương thức generic `Nop` trả về đối số của nó, cho bất kỳ kiểu nào có thể:

```
type Empty struct{}

func (Empty) Nop[T any](x T) T {
	return x
}
```

Bây giờ giả sử một giá trị `Empty` được lưu trong một `any` và truyền
cho mã khác kiểm tra nó có thể làm gì:

```
func TryNops(x any) {
	if x, ok := x.(interface{ Nop(string) string }); ok {
		fmt.Printf("string %s\n", x.Nop("hello"))
	}
	if x, ok := x.(interface{ Nop(int) int }); ok {
		fmt.Printf("int %d\n", x.Nop(42))
	}
	if x, ok := x.(interface{ Nop(io.Reader) io.Reader }); ok {
		data, err := io.ReadAll(x.Nop(strings.NewReader("hello world")))
		fmt.Printf("reader %q %v\n", data, err)
	}
}
```

Mã đó hoạt động như thế nào nếu `x` là một `Empty`?
Có vẻ như `x` phải thỏa mãn cả ba kiểm tra,
cùng với bất kỳ dạng nào khác với bất kỳ kiểu nào khác.

Mã nào chạy khi những phương thức đó được gọi?
Đối với các phương thức không generic, trình biên dịch tạo mã
cho tất cả các triển khai phương thức và liên kết chúng vào chương trình cuối cùng.
Nhưng đối với các phương thức generic, có thể có số lượng triển khai phương thức vô hạn,
vì vậy cần một chiến lược khác.

Có bốn lựa chọn:

 1. Tại thời điểm link, tạo danh sách tất cả các kiểm tra interface động có thể,
    rồi tìm kiếm các kiểu thỏa mãn chúng nhưng thiếu
    các phương thức đã biên dịch, và sau đó gọi lại trình biên dịch để thêm những phương thức đó.

    Điều này sẽ làm chậm quá trình build đáng kể, do cần phải dừng sau khi
    link và lặp lại một số biên dịch. Nó sẽ đặc biệt làm chậm các incremental build.
    Tệ hơn, có thể mã phương thức được biên dịch mới chính nó
    có các kiểm tra interface động mới, và quá trình
    phải được lặp lại. Các ví dụ có thể được xây dựng nơi
    quá trình thậm chí không bao giờ kết thúc.

 2. Triển khai một dạng JIT, biên dịch mã phương thức cần thiết khi chạy.

    Go hưởng lợi rất nhiều từ sự đơn giản và hiệu suất có thể dự đoán
    của việc biên dịch hoàn toàn ahead-of-time.
    Chúng tôi không muốn đảm nhận sự phức tạp của JIT chỉ để triển khai
    một tính năng ngôn ngữ.

 3. Sắp xếp để phát ra một fallback chậm cho mỗi phương thức generic dùng
    một bảng các hàm cho mỗi thao tác ngôn ngữ có thể trên tham số kiểu,
    và sau đó dùng triển khai fallback đó cho các kiểm tra động.

    Cách tiếp cận này sẽ làm cho một phương thức generic được tham số hóa bởi
    kiểu không mong đợi chậm hơn nhiều so với phương thức tương tự
    được tham số hóa bởi kiểu được quan sát tại thời điểm biên dịch.
    Điều này sẽ làm cho hiệu suất khó dự đoán hơn nhiều.

 4. Định nghĩa rằng các phương thức generic không thể được dùng để thỏa mãn interface.

    Interface là một phần thiết yếu của lập trình trong Go.
    Không cho phép các phương thức generic thỏa mãn interface là không thể chấp nhận
    từ quan điểm thiết kế.

Không có lựa chọn nào trong số này là tốt, vì vậy chúng tôi chọn "không có cái nào ở trên."

Thay vì các phương thức với tham số kiểu, hãy dùng hàm top-level với
tham số kiểu, hoặc thêm tham số kiểu vào kiểu receiver.

Để biết thêm chi tiết, bao gồm thêm ví dụ, xem
[đề xuất](/design/43651-type-parameters#no-parameterized-methods).

### Tại sao tôi không thể dùng kiểu cụ thể hơn cho receiver của kiểu được tham số hóa? {#types_in_method_declaration}

Các khai báo phương thức của kiểu generic được viết với một receiver
bao gồm tên tham số kiểu.
Có lẽ vì sự tương đồng về cú pháp để chỉ định các kiểu
tại một call site,
một số người đã nghĩ điều này cung cấp cơ chế để tạo ra
một phương thức được tùy chỉnh cho một số đối số kiểu nhất định bằng cách đặt tên
một kiểu cụ thể trong receiver, chẳng hạn như `string`:

```
type S[T any] struct { f T }

func (s S[string]) Add(t string) string {
    return s.f + t
}
```

Điều này thất bại vì từ `string` được
trình biên dịch hiểu là tên của đối số kiểu trong phương thức.
Thông báo lỗi trình biên dịch sẽ có dạng như "`operator + not defined on
s.f (variable of type string)`".
Điều này có thể gây nhầm lẫn vì toán tử `+`
hoạt động tốt trên kiểu khai báo trước `string`,
nhưng khai báo đã ghi đè, cho phương thức này, định nghĩa của `string`,
và toán tử không hoạt động trên phiên bản không liên quan đó của `string`.
Ghi đè một tên được khai báo trước như thế này là hợp lệ, nhưng là điều kỳ lạ để làm và
thường là một lỗi.

### Tại sao trình biên dịch không thể suy luận đối số kiểu trong chương trình của tôi? {#type_inference}

Có nhiều trường hợp mà lập trình viên có thể dễ dàng thấy
đối số kiểu cho một kiểu hoặc hàm generic phải là gì, nhưng ngôn ngữ không
cho phép trình biên dịch suy luận nó.
Suy luận kiểu được giới hạn có chủ đích để đảm bảo không bao giờ
có sự nhầm lẫn nào về kiểu nào được suy luận.
Kinh nghiệm với các ngôn ngữ khác cho thấy suy luận kiểu không mong đợi
có thể dẫn đến nhầm lẫn đáng kể khi đọc và
gỡ lỗi một chương trình.
Luôn luôn có thể chỉ định đối số kiểu rõ ràng được dùng
trong lời gọi.
Trong tương lai, các dạng suy luận mới có thể được hỗ trợ, miễn là
các quy tắc vẫn đơn giản và rõ ràng.

## Package và Kiểm thử {#Packages_Testing}

### Làm thế nào để tôi tạo một package đa tệp? {#How_do_I_create_a_multifile_package}

Đặt tất cả các tệp nguồn cho package vào một thư mục riêng.
Các tệp nguồn có thể tham chiếu đến các mục từ các tệp khác tùy ý; không
cần khai báo tiền tố hoặc tệp header.

Ngoài việc được chia thành nhiều tệp, package sẽ biên dịch và kiểm thử
giống như một package đơn tệp.

### Làm thế nào để tôi viết unit test? {#How_do_I_write_a_unit_test}

Tạo một tệp mới kết thúc bằng `_test.go` trong cùng thư mục
với mã nguồn package. Bên trong tệp đó, `import "testing"`
và viết các hàm có dạng

```
func TestFoo(t *testing.T) {
    ...
}
```

Chạy `go test` trong thư mục đó.
Script đó tìm các hàm `Test`,
build tệp nhị phân kiểm thử và chạy nó.

Xem tài liệu [Cách viết mã Go](/doc/code.html),
package [`testing`](/pkg/testing/)
và lệnh con [`go test`](/cmd/go/#hdr-Test_packages) để biết thêm chi tiết.

### Hàm helper kiểm thử yêu thích của tôi ở đâu? {#testing_framework}

Package [`testing`](/pkg/testing/) chuẩn của Go giúp dễ dàng viết unit test, nhưng nó thiếu
các tính năng được cung cấp trong các framework kiểm thử của ngôn ngữ khác như các hàm assertion.
Một [phần trước đó](#assertions) của tài liệu này giải thích tại sao Go
không có assertions, và
các lập luận tương tự áp dụng cho việc dùng `assert` trong kiểm thử.
Xử lý lỗi đúng cách có nghĩa là để các kiểm thử khác chạy sau khi một kiểm thử thất bại,
để người gỡ lỗi lỗi có được bức tranh đầy đủ về những gì
sai. Hữu ích hơn nếu một kiểm thử báo cáo rằng
`isPrime` cho kết quả sai cho 2, 3, 5 và 7 (hoặc cho
2, 4, 8 và 16) hơn là báo cáo rằng `isPrime` cho kết quả sai
cho 2 và do đó không có kiểm thử nào khác được chạy. Lập trình viên
kích hoạt lỗi kiểm thử có thể không quen với mã thất bại.
Thời gian đầu tư để viết thông báo lỗi tốt sẽ được đền đáp sau khi
kiểm thử bị hỏng.

Một điểm liên quan là các framework kiểm thử có xu hướng phát triển thành các ngôn ngữ mini
của riêng chúng, với các điều kiện và điều khiển và cơ chế in ấn,
nhưng Go đã có tất cả những khả năng đó; tại sao phải tạo lại chúng?
Chúng tôi muốn viết kiểm thử bằng Go; đó là một ngôn ngữ ít hơn cần học và cách
tiếp cận giữ cho các kiểm thử đơn giản và dễ hiểu.

Nếu lượng mã bổ sung cần thiết để viết
lỗi tốt có vẻ lặp đi lặp lại và quá nhiều, kiểm thử có thể hoạt động tốt hơn nếu
theo bảng, lặp qua danh sách các đầu vào và đầu ra được định nghĩa
trong cấu trúc dữ liệu (Go có hỗ trợ xuất sắc cho literal cấu trúc dữ liệu).
Công việc để viết một kiểm thử tốt và thông báo lỗi tốt sau đó sẽ được khấu hao trên nhiều
trường hợp kiểm thử. Thư viện Go chuẩn đầy ắp các ví dụ minh họa, chẳng hạn như trong
[các kiểm thử định dạng cho package `fmt`](/src/fmt/fmt_test.go).

### Tại sao *X* không có trong thư viện chuẩn? {#x_in_std}

Mục đích của thư viện chuẩn là hỗ trợ thư viện runtime, kết nối với
hệ điều hành, và cung cấp chức năng chính mà nhiều chương trình Go
yêu cầu, như I/O định dạng và mạng.
Nó cũng chứa các phần tử quan trọng cho lập trình web, bao gồm
mật mã và hỗ trợ cho các chuẩn như HTTP, JSON và XML.

Không có tiêu chí rõ ràng định nghĩa những gì được bao gồm vì trong
một thời gian dài, đây là thư viện Go *duy nhất*.
Tuy nhiên có các tiêu chí định nghĩa những gì được thêm ngày nay.

Các bổ sung mới vào thư viện chuẩn rất hiếm và rào cản để
đưa vào cao.
Mã được bao gồm trong thư viện chuẩn chịu chi phí bảo trì liên tục lớn
(thường được gánh chịu bởi những người khác ngoài tác giả gốc),
phải tuân theo [cam kết tương thích Go 1](/doc/go1compat.html)
(chặn các sửa chữa bất kỳ lỗi nào trong API),
và phải tuân theo
[lịch phát hành](/s/releasesched) của Go,
ngăn các sửa lỗi đến tay người dùng nhanh chóng.

Hầu hết mã mới nên sống bên ngoài thư viện chuẩn và có thể truy cập
qua lệnh `go get` của [công cụ `go`](/cmd/go/).
Mã đó có thể có người bảo trì riêng, chu kỳ phát hành,
và đảm bảo tương thích.
Người dùng có thể tìm các package và đọc tài liệu của chúng tại
[pkg.go.dev](https://pkg.go.dev/).

Mặc dù có những phần trong thư viện chuẩn không thực sự thuộc về đó,
chẳng hạn như `log/syslog`, chúng tôi tiếp tục duy trì mọi thứ trong
thư viện vì cam kết tương thích Go 1.
Nhưng chúng tôi khuyến khích hầu hết mã mới sống ở nơi khác.

## Triển khai {#Implementation}

### Công nghệ trình biên dịch nào được dùng để xây dựng các trình biên dịch? {#What_compiler_technology_is_used_to_build_the_compilers}

Có một số trình biên dịch cho môi trường production dành cho Go, và một số trình biên dịch khác
đang phát triển cho các nền tảng khác nhau.

Trình biên dịch mặc định, `gc`, được bao gồm trong
bản phân phối Go như một phần hỗ trợ cho lệnh `go`.
`Gc` ban đầu được viết bằng C
do khó khăn của bootstrapping -- bạn cần một trình biên dịch Go để
thiết lập môi trường Go.
Nhưng mọi thứ đã tiến bộ và kể từ bản phát hành Go 1.5, trình biên dịch là
một chương trình Go.
Trình biên dịch đã được chuyển đổi từ C sang Go bằng các công cụ dịch tự động, như
được mô tả trong [tài liệu thiết kế](/s/go13compiler) này
và [bài nói chuyện](/talks/2015/gogo.slide#1) này.
Do đó trình biên dịch bây giờ là "self-hosting", có nghĩa là chúng tôi cần phải đối mặt
với vấn đề bootstrapping.
Giải pháp là có một cài đặt Go hoạt động sẵn có,
giống như người ta thường có với một cài đặt C hoạt động.
Câu chuyện về cách khởi động một môi trường Go mới từ mã nguồn
được mô tả [ở đây](/s/go15bootstrap) và
[ở đây](/doc/install/source).

`Gc` được viết bằng Go với bộ phân tích recursive descent
và dùng loader tùy chỉnh, cũng được viết bằng Go nhưng
dựa trên Plan 9 loader, để tạo các tệp nhị phân ELF/Mach-O/PE.

Trình biên dịch `Gccgo` là front end được viết bằng C++
với bộ phân tích recursive descent kết hợp với
GCC back end chuẩn. Một
[LLVM back end](https://go.googlesource.com/gollvm/) thử nghiệm đang
dùng cùng front end.

Vào đầu dự án, chúng tôi đã xem xét dùng LLVM cho
`gc` nhưng quyết định nó quá lớn và chậm để đáp ứng
các mục tiêu hiệu suất của chúng tôi.
Quan trọng hơn nhìn lại, bắt đầu với LLVM sẽ làm khó hơn
để giới thiệu một số thay đổi ABI và liên quan, như
quản lý stack, mà Go yêu cầu nhưng không phải là một phần của thiết lập C chuẩn.

Go hóa ra là ngôn ngữ tốt để triển khai trình biên dịch Go,
mặc dù đó không phải là mục tiêu ban đầu của nó.
Việc không phải self-hosting từ đầu cho phép thiết kế của Go
tập trung vào trường hợp sử dụng ban đầu của nó, là các server mạng.
Nếu chúng tôi quyết định Go nên tự biên dịch sớm, chúng tôi có thể đã
kết thúc với một ngôn ngữ hướng nhiều hơn cho việc xây dựng trình biên dịch,
đây là mục tiêu xứng đáng nhưng không phải mục tiêu chúng tôi ban đầu có.

Mặc dù `gc` có triển khai riêng của nó, lexer và
parser gốc có sẵn trong package [`go/parser`](/pkg/go/parser/) và
cũng có [type checker](/pkg/go/types) gốc.
Trình biên dịch `gc` dùng các biến thể của các thư viện này.

### Hỗ trợ run-time được triển khai như thế nào? {#How_is_the_run_time_support_implemented}

Một lần nữa do các vấn đề bootstrapping, mã run-time ban đầu được viết hầu hết bằng C (với
một chút assembler) nhưng kể từ đó nó đã được dịch sang Go
(ngoại trừ một số bit assembler).
Hỗ trợ run-time của `Gccgo` dùng `glibc`.
Trình biên dịch `gccgo` triển khai goroutine bằng cách dùng
một kỹ thuật gọi là segmented stacks,
được hỗ trợ bởi các sửa đổi gần đây cho gold linker.
`Gollvm` tương tự được xây dựng trên cơ sở hạ tầng LLVM tương ứng.

### Tại sao chương trình nhỏ của tôi lại là tệp nhị phân lớn? {#Why_is_my_trivial_program_such_a_large_binary}

Linker trong `gc` toolchain
tạo các tệp nhị phân liên kết tĩnh theo mặc định.
Do đó, tất cả tệp nhị phân Go bao gồm Go
runtime, cùng với thông tin kiểu run-time cần thiết để hỗ trợ các kiểm tra kiểu động,
reflection và thậm chí stack trace khi panic.

Một chương trình C "hello, world" đơn giản được biên dịch và liên kết tĩnh bằng
gcc trên Linux khoảng 750 kB, bao gồm một triển khai của
`printf`.
Một chương trình Go tương đương dùng
`fmt.Printf` nặng vài megabyte, nhưng điều đó bao gồm
hỗ trợ run-time mạnh mẽ hơn và thông tin kiểu và gỡ lỗi.

Một chương trình Go được biên dịch với `gc` có thể được liên kết với
flag `-ldflags=-w` để vô hiệu hóa tạo DWARF,
loại bỏ thông tin gỡ lỗi khỏi tệp nhị phân nhưng không mất
chức năng khác.
Điều này có thể giảm đáng kể kích thước tệp nhị phân.

### Tôi có thể tắt những phàn nàn về biến/import chưa dùng không? {#unused_variables_and_imports}

Sự hiện diện của biến chưa dùng có thể chỉ ra lỗi, trong khi
import chưa dùng chỉ làm chậm biên dịch,
một tác động có thể trở nên đáng kể khi chương trình tích lũy
mã và lập trình viên theo thời gian.
Vì những lý do này, Go từ chối biên dịch các chương trình có biến chưa dùng
hoặc import,
đánh đổi sự tiện lợi ngắn hạn để có tốc độ build dài hạn và
sự rõ ràng của chương trình.

Tuy nhiên, khi đang phát triển mã, việc tạo ra những tình huống này
tạm thời là phổ biến và có thể khó chịu khi phải chỉnh sửa chúng trước khi
chương trình sẽ biên dịch.

Một số người đã yêu cầu một tùy chọn trình biên dịch để tắt những kiểm tra đó
hoặc ít nhất giảm chúng thành cảnh báo.
Tùy chọn như vậy chưa được thêm vào, tuy nhiên,
vì tùy chọn trình biên dịch không nên ảnh hưởng đến ngữ nghĩa của
ngôn ngữ và vì trình biên dịch Go không báo cáo cảnh báo, chỉ
lỗi ngăn biên dịch.

Có hai lý do để không có cảnh báo. Đầu tiên, nếu đáng phàn nàn,
đáng sửa trong mã. (Ngược lại, nếu không đáng sửa, không đáng đề cập.)
Thứ hai, việc trình biên dịch tạo cảnh báo khuyến khích triển khai cảnh báo về
các trường hợp yếu có thể làm biên dịch ồn ào, che giấu lỗi thực
*nên* được sửa.

Tuy nhiên, dễ dàng giải quyết tình huống này. Dùng blank identifier
để để những thứ chưa dùng tồn tại trong khi bạn đang phát triển.

```
import "unused"

// Khai báo này đánh dấu import là được dùng bằng cách tham chiếu một
// mục từ package.
var _ = unused.Item  // TODO: Xóa trước khi commit!

func main() {
    debugData := debug.Profile()
    _ = debugData // Chỉ dùng trong quá trình gỡ lỗi.
    ....
}
```

Hiện nay, hầu hết lập trình viên Go dùng một công cụ,
[goimports](https://godoc.org/golang.org/x/tools/cmd/goimports),
tự động viết lại tệp nguồn Go để có các import đúng,
loại bỏ vấn đề import chưa dùng trong thực tế.
Chương trình này dễ dàng kết nối với hầu hết các trình soạn thảo và IDE để chạy tự động khi tệp nguồn Go được ghi.
Chức năng này cũng được tích hợp vào `gopls`, như
[đã thảo luận ở trên](/doc/faq#ide).

### Tại sao phần mềm quét virus của tôi nghĩ bản phân phối Go hoặc tệp nhị phân đã biên dịch của tôi bị nhiễm? {#virus}

Đây là hiện tượng phổ biến, đặc biệt trên máy Windows, và hầu như luôn là dương tính giả.
Các chương trình quét virus thương mại thường bị nhầm lẫn bởi cấu trúc của tệp nhị phân Go,
mà chúng không thấy thường xuyên như những tệp được biên dịch từ các ngôn ngữ khác.

Nếu bạn vừa cài đặt bản phân phối Go và hệ thống báo cáo nó bị nhiễm, đó chắc chắn là lỗi.
Để thực sự kỹ lưỡng, bạn có thể xác minh tải xuống bằng cách so sánh checksum với những checksum trên
[trang tải xuống](/dl/).

Trong mọi trường hợp, nếu bạn tin báo cáo là sai, vui lòng báo cáo lỗi cho nhà cung cấp máy quét virus của bạn.
Có lẽ theo thời gian, máy quét virus có thể học cách hiểu các chương trình Go.

## Hiệu năng {#Performance}

### Tại sao Go thực hiện kém trên benchmark X? {#Why_does_Go_perform_badly_on_benchmark_x}

Một trong các mục tiêu thiết kế của Go là tiếp cận hiệu suất của C cho
các chương trình tương đương, nhưng trên một số benchmark nó thực hiện khá kém, bao gồm một số
trong [golang.org/x/exp/shootout](https://go.googlesource.com/exp/+/master/shootout/).
Những cái chậm nhất phụ thuộc vào các thư viện mà phiên bản có hiệu suất tương đương
không có trong Go.
Ví dụ, [pidigits.go](https://go.googlesource.com/exp/+/master/shootout/pidigits.go)
phụ thuộc vào package toán học đa độ chính xác, và các phiên bản C,
không giống Go, dùng [GMP](https://gmplib.org/) (được
viết bằng assembler tối ưu).
Các benchmark phụ thuộc vào regular expression
([regex-dna.go](https://go.googlesource.com/exp/+/master/shootout/regex-dna.go),
ví dụ) về cơ bản so sánh [package regexp](/pkg/regexp) gốc của Go với
các thư viện regular expression trưởng thành, được tối ưu hóa cao như PCRE.

Các trò chơi benchmark được giành chiến thắng bằng sự điều chỉnh rộng rãi và các phiên bản Go của hầu hết
các benchmark cần được chú ý. Nếu bạn đo lường các chương trình C
và Go thực sự tương đương
([reverse-complement.go](https://go.googlesource.com/exp/+/master/shootout/reverse-complement.go)
là một ví dụ), bạn sẽ thấy hai ngôn ngữ gần nhau hơn nhiều trong hiệu suất thô
so với bộ này cho thấy.

Tuy nhiên, vẫn còn chỗ để cải thiện. Các trình biên dịch tốt nhưng có thể
tốt hơn, nhiều thư viện cần công việc hiệu suất lớn, và bộ gom rác
chưa đủ nhanh. (Ngay cả khi nó như vậy, cẩn thận không tạo ra
rác không cần thiết có thể có tác động rất lớn.)

Trong mọi trường hợp, Go thường có thể rất cạnh tranh.
Đã có cải thiện đáng kể trong hiệu suất của nhiều chương trình
khi ngôn ngữ và công cụ phát triển.
Xem bài đăng trên blog về
[profiling
các chương trình Go](/blog/profiling-go-programs) để biết ví dụ thông tin.
Nó khá cũ nhưng vẫn chứa thông tin hữu ích.

## Thay đổi từ C {#change_from_c}

### Tại sao cú pháp lại khác C đến vậy? {#different_syntax}

Ngoài cú pháp khai báo, các khác biệt không lớn và xuất phát từ
hai mong muốn. Đầu tiên, cú pháp nên cảm thấy nhẹ nhàng, không có
quá nhiều từ khóa bắt buộc, lặp lại hoặc khó hiểu. Thứ hai, ngôn ngữ
được thiết kế để dễ phân tích
và có thể được phân tích mà không cần bảng ký hiệu. Điều này làm cho
việc xây dựng các công cụ như debugger, trình phân tích dependency, trình trích xuất tài liệu tự động,
plugin IDE và tương tự dễ dàng hơn nhiều. C và
các hậu duệ của nó nổi tiếng là khó khăn trong vấn đề này.

### Tại sao các khai báo lại ngược? {#declarations_backwards}

Chúng chỉ ngược nếu bạn đã quen với C. Trong C, khái niệm là một
biến được khai báo giống như biểu thức biểu thị kiểu của nó, đây là một
ý tưởng hay, nhưng ngữ pháp kiểu và biểu thức không kết hợp tốt và
kết quả có thể gây nhầm lẫn; hãy xem xét các con trỏ hàm. Go chủ yếu
tách biệt cú pháp biểu thức và kiểu và điều đó đơn giản hóa mọi thứ (dùng
tiền tố `*` cho con trỏ là ngoại lệ chứng minh quy tắc). Trong C,
khai báo

```
    int* a, b;
```

khai báo `a` là con trỏ nhưng không phải `b`; trong Go

```
    var a, b *int
```

khai báo cả hai là con trỏ. Điều này rõ ràng và đồng đều hơn.
Ngoài ra, dạng khai báo ngắn `:=` cho thấy rằng khai báo biến đầy đủ
nên trình bày cùng thứ tự như `:=` vì vậy

```
    var a uint64 = 1
```

có tác dụng tương tự như

```
    a := uint64(1)
```

Phân tích cú pháp cũng được đơn giản hóa bằng cách có ngữ pháp riêng biệt cho các kiểu
không chỉ là ngữ pháp biểu thức; các từ khóa như `func`
và `chan` giữ mọi thứ rõ ràng.

Xem bài viết về
[Cú pháp khai báo của Go](/doc/articles/gos_declaration_syntax.html)
để biết thêm chi tiết.

### Tại sao không có phép tính số học con trỏ? {#no_pointer_arithmetic}

An toàn. Không có phép tính số học con trỏ, có thể tạo ra một
ngôn ngữ không bao giờ có thể suy ra một địa chỉ bất hợp pháp thành công
một cách sai lầm. Công nghệ trình biên dịch và phần cứng đã tiến bộ đến
mức một vòng lặp dùng chỉ số mảng có thể hiệu quả như vòng lặp
dùng phép tính số học con trỏ. Ngoài ra, việc thiếu phép tính số học con trỏ có thể
đơn giản hóa triển khai của bộ gom rác.

### Tại sao `++` và `--` là câu lệnh chứ không phải biểu thức? Và tại sao là postfix, không phải prefix? {#inc_dec}

Không có phép tính số học con trỏ, giá trị tiện lợi của các toán tử tăng và giảm pre- và postfix giảm đi.
Bằng cách loại bỏ chúng khỏi phân cấp biểu thức hoàn toàn, cú pháp biểu thức được đơn giản hóa và các vấn đề lộn xộn
xung quanh thứ tự đánh giá của `++` và `--`
(hãy xem xét `f(i++)` và `p[i] = q[++i]`)
cũng được loại bỏ. Sự đơn giản hóa là
đáng kể. Còn về postfix vs. prefix, cả hai sẽ hoạt động tốt nhưng
phiên bản postfix là truyền thống hơn; sự khăng khăng về prefix phát sinh
với STL, một thư viện cho ngôn ngữ có tên chứa, mỉa mai thay, một
postfix increment.

### Tại sao có dấu ngoặc nhọn nhưng không có dấu chấm phẩy? Và tại sao tôi không thể đặt dấu ngoặc mở trên dòng tiếp theo? {#semicolons}

Go dùng dấu ngoặc nhọn để nhóm câu lệnh, một cú pháp quen thuộc với
lập trình viên đã làm việc với bất kỳ ngôn ngữ nào trong họ C.
Tuy nhiên, dấu chấm phẩy là cho parser, không phải cho người dùng, và chúng tôi muốn
loại bỏ chúng càng nhiều càng tốt. Để đạt mục tiêu này, Go vay mượn
một mẹo từ BCPL: các dấu chấm phẩy phân tách câu lệnh có trong
ngữ pháp hình thức nhưng được chèn tự động, không cần lookahead, bởi
lexer ở cuối bất kỳ dòng nào có thể là cuối câu lệnh.
Điều này hoạt động rất tốt trong thực tế nhưng có tác dụng là bắt buộc phong cách dấu ngoặc.
Ví dụ, dấu ngoặc mở của một hàm không thể
xuất hiện trên một dòng riêng.

Một số người đã lập luận rằng lexer nên lookahead để cho phép
dấu ngoặc sống trên dòng tiếp theo. Chúng tôi không đồng ý. Vì mã Go được định dạng tự động bởi
[`gofmt`](/cmd/gofmt/),
*một số* phong cách phải được chọn. Phong cách đó có thể khác với những gì
bạn đã dùng trong C hoặc Java, nhưng Go là ngôn ngữ khác và
phong cách của `gofmt` tốt như bất kỳ phong cách nào khác. Quan trọng hơn
-- quan trọng hơn nhiều -- lợi thế của định dạng duy nhất, được quy định bởi chương trình
cho tất cả các chương trình Go vượt xa bất kỳ bất lợi nhận thức nào của phong cách cụ thể.
Cũng lưu ý rằng phong cách của Go có nghĩa là một triển khai tương tác của
Go có thể dùng cú pháp chuẩn một dòng tại một thời điểm mà không cần quy tắc đặc biệt.

### Tại sao dùng bộ gom rác? Nó có đắt tiền quá không? {#garbage_collection}

Một trong những nguồn bookkeeping lớn nhất trong các chương trình hệ thống là
quản lý vòng đời của các đối tượng được phân bổ.
Trong các ngôn ngữ như C nơi nó được thực hiện thủ công,
nó có thể tiêu tốn lượng đáng kể thời gian lập trình viên và thường
là nguyên nhân của các lỗi nguy hiểm.
Ngay cả trong các ngôn ngữ như C++ hoặc Rust cung cấp các cơ chế
để hỗ trợ, những cơ chế đó có thể có tác động đáng kể lên
thiết kế của phần mềm, thường thêm overhead lập trình
của riêng nó.
Chúng tôi cảm thấy điều quan trọng là loại bỏ những
overhead như vậy cho lập trình viên, và những tiến bộ trong công nghệ bộ gom rác
trong những năm gần đây cho chúng tôi sự tự tin rằng nó
có thể được triển khai đủ rẻ, và với độ trễ đủ thấp,
rằng nó có thể là một cách tiếp cận khả thi cho các
hệ thống mạng.

Phần lớn khó khăn của lập trình concurrent
có gốc rễ trong vấn đề vòng đời đối tượng:
khi các đối tượng được truyền giữa các thread, việc đảm bảo
chúng được giải phóng an toàn trở nên cồng kềnh.
Bộ gom rác tự động làm cho mã concurrent dễ viết hơn nhiều.
Tất nhiên, triển khai bộ gom rác trong môi trường concurrent tự nó
là một thách thức, nhưng gặp nó một lần thay vì trong mỗi
chương trình giúp ích cho mọi người.

Cuối cùng, ngoài concurrency, bộ gom rác làm cho các interface
đơn giản hơn vì chúng không cần chỉ định cách bộ nhớ được quản lý qua chúng.

Điều này không có nghĩa là công việc gần đây trong các ngôn ngữ
như Rust mang lại các ý tưởng mới cho vấn đề quản lý
tài nguyên là sai lầm; chúng tôi khuyến khích công việc này và hào hứng khi thấy
nó phát triển như thế nào.
Nhưng Go có cách tiếp cận truyền thống hơn bằng cách giải quyết
vòng đời đối tượng thông qua
bộ gom rác, và chỉ bộ gom rác mà thôi.

Triển khai hiện tại là bộ thu mark-and-sweep.
Nếu máy là bộ xử lý đa nhân, bộ thu chạy trên một lõi CPU riêng
song song với chương trình chính.
Công việc lớn trên bộ thu trong những năm gần đây đã giảm thời gian tạm dừng
thường xuống phạm vi dưới millisecond, ngay cả đối với các heap lớn,
gần như loại bỏ một trong những phản đối lớn đối với bộ gom rác
trong các server mạng.
Công việc tiếp tục để tinh chỉnh thuật toán, giảm thêm overhead và
độ trễ, và để khám phá các cách tiếp cận mới.
[Keynote ISMM](/blog/ismmkeynote) năm 2018
của Rick Hudson từ nhóm Go
mô tả tiến độ cho đến nay và đề xuất một số cách tiếp cận trong tương lai.

Về chủ đề hiệu suất, hãy lưu ý rằng Go cho lập trình viên
khả năng kiểm soát đáng kể về bố cục bộ nhớ và phân bổ, nhiều hơn
so với thông thường trong các ngôn ngữ được gom rác. Một lập trình viên cẩn thận có thể giảm
overhead bộ gom rác đáng kể bằng cách dùng ngôn ngữ tốt;
xem bài viết về
[profiling các chương trình Go](/blog/profiling-go-programs) để biết ví dụ có hướng dẫn,
bao gồm một bản trình diễn về các công cụ profiling của Go.
