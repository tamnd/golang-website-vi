---
title: "Frequently Asked Questions (FAQ)"
sidebar: "faq"
breadcrumb: true
template: true
---

## Nguồn gốc {#Origins}

### Mục đích của dự án là gì? {#What_is_the_purpose_of_the_project}

Vào thời điểm Go ra đời năm 2007, thế giới lập trình còn rất khác so với ngày nay.
Phần mềm production thường được viết bằng C++ hoặc Java,
GitHub chưa tồn tại, hầu hết máy tính chưa có bộ vi xử lý đa nhân,
và ngoài Visual Studio và Eclipse, hầu như không có IDE hay công cụ cấp cao nào khác,
chứ chưa nói đến miễn phí trên Internet.

Trong khi đó, chúng tôi ngày càng thất vọng với sự phức tạp thái quá cần thiết
để xây dựng các dự án phần mềm lớn bằng những ngôn ngữ mình đang dùng và hệ thống build đi kèm.
Máy tính đã trở nên nhanh hơn rất nhiều kể từ khi C, C++ và Java lần đầu được phát triển,
nhưng bản thân việc lập trình lại chưa tiến bộ được bao nhiêu.
Ngoài ra, rõ ràng bộ vi xử lý đa nhân đang dần trở nên phổ biến,
nhưng hầu hết các ngôn ngữ lại cung cấp rất ít hỗ trợ để lập trình hiệu quả và an toàn trên chúng.

Chúng tôi quyết định lùi lại và suy nghĩ xem những vấn đề lớn nào sẽ
chi phối kỹ thuật phần mềm trong những năm tới khi công nghệ phát triển,
và một ngôn ngữ mới có thể giúp giải quyết chúng như thế nào.
Chẳng hạn, sự trỗi dậy của CPU đa nhân cho thấy một ngôn ngữ nên
cung cấp hỗ trợ bậc nhất cho một dạng tính toán song song hay concurrency nào đó.
Và để quản lý tài nguyên hợp lý trong một chương trình concurrent lớn,
cần có bộ gom rác, hoặc ít nhất là một dạng quản lý bộ nhớ tự động an toàn.

Những cân nhắc này dẫn đến
[một loạt cuộc thảo luận](https://commandcenter.blogspot.com/2017/09/go-ten-years-and-climbing.html)
từ đó Go ra đời, ban đầu là một tập hợp ý tưởng và mong muốn,
sau đó trở thành một ngôn ngữ.
Mục tiêu bao trùm là Go làm được nhiều hơn để hỗ trợ lập trình viên trong công việc hàng ngày
bằng cách tạo điều kiện cho hệ thống công cụ, tự động hóa các tác vụ nhàm chán như định dạng mã nguồn,
và loại bỏ các trở ngại khi làm việc với cơ sở mã lớn.

Một mô tả chi tiết hơn về các mục tiêu của Go và cách chúng được hiện thực hóa,
hoặc ít nhất là tiếp cận, có trong bài viết
[Go at Google: Language Design in the Service of Software Engineering](/talks/2012/splash.article).

### Lịch sử của dự án như thế nào? {#history}

Robert Griesemer, Rob Pike và Ken Thompson bắt đầu phác thảo
các mục tiêu cho một ngôn ngữ mới trên bảng trắng vào ngày 21 tháng 9 năm 2007.
Chỉ vài ngày sau, các mục tiêu đã định hình thành một kế hoạch để làm điều gì đó
và một ý tưởng rõ ràng về hướng đi. Quá trình thiết kế tiếp tục theo kiểu bán thời gian
song song với công việc khác. Đến tháng 1 năm 2008, Ken đã bắt đầu viết một trình biên dịch
để khám phá các ý tưởng; nó tạo ra mã C làm đầu ra. Đến giữa năm, ngôn ngữ đã trở thành
dự án toàn thời gian và đã đủ ổn định để thử xây dựng trình biên dịch production.
Vào tháng 5 năm 2008, Ian Taylor độc lập bắt đầu xây dựng một GCC front end cho Go
dựa trên bản đặc tả nháp. Russ Cox gia nhập vào cuối năm 2008 và giúp đưa ngôn ngữ
và các thư viện từ giai đoạn prototype thành hiện thực.

Go trở thành dự án mã nguồn mở vào ngày 10 tháng 11 năm 2009.
Vô số người từ cộng đồng đã đóng góp ý tưởng, thảo luận và mã nguồn.

Hiện nay có hàng triệu lập trình viên Go, hay gopher, trên toàn thế giới,
và con số này ngày càng tăng.
Thành công của Go đã vượt xa kỳ vọng của chúng tôi.

### Nguồn gốc của linh vật gopher là gì? {#gopher}

Linh vật và logo được thiết kế bởi
[Renée French](https://reneefrench.blogspot.com), người cũng thiết kế
[Glenda](https://9p.io/plan9/glenda.html),
chú thỏ của Plan 9.
Một [bài đăng trên blog](/blog/gopher)
về gopher giải thích cách nó được lấy cảm hứng từ một hình ảnh cô ấy đã dùng
cho thiết kế áo phông [WFMU](https://wfmu.org/) vài năm trước.
Logo và linh vật được bảo vệ bởi giấy phép
[Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/).

Gopher có một
[bảng mẫu nhân vật](/doc/gopher/modelsheet.jpg)
minh họa các đặc điểm của nó và cách thể hiện chúng đúng cách.
Bảng mẫu lần đầu được trình bày trong một
[bài nói chuyện](https://www.youtube.com/watch?v=4rw_B4yY69k)
của Renée tại Gophercon năm 2016.
Nó có những đặc điểm riêng biệt; đây là *Go gopher*, không phải gopher tùy tiện nào cũng được.

### Ngôn ngữ được gọi là Go hay Golang? {#go_or_golang}

Ngôn ngữ được gọi là Go.
Biệt danh "golang" xuất hiện vì trang web ban đầu là
*golang.org*.
(Khi đó chưa có tên miền *.dev*.)
Nhiều người vẫn dùng tên golang và nó cũng tiện dùng
làm nhãn hiệu.
Chẳng hạn, hashtag mạng xã hội cho ngôn ngữ là "#golang".
Dù vậy, tên chính thức của ngôn ngữ vẫn chỉ là Go.

Chú thích bên lề: Mặc dù
[logo chính thức](/blog/go-brand)
có hai chữ cái viết hoa, tên ngôn ngữ được viết là Go, không phải GO.

### Tại sao các bạn tạo ra một ngôn ngữ mới? {#creating_a_new_language}

Go ra đời từ sự thất vọng với các ngôn ngữ và môi trường hiện có
cho công việc chúng tôi đang làm tại Google.
Việc lập trình đã trở nên quá khó khăn và sự lựa chọn ngôn ngữ một phần là nguyên nhân.
Người ta phải chọn giữa biên dịch hiệu quả, thực thi hiệu quả, hoặc dễ lập trình;
cả ba không tồn tại trong cùng một ngôn ngữ mainstream. Những lập trình viên có thể lựa chọn
đang ưu tiên sự tiện lợi hơn an toàn và hiệu suất bằng cách chuyển sang các ngôn ngữ
kiểu động như Python và JavaScript thay vì C++ hay, ở mức độ thấp hơn, Java.

Chúng tôi không đơn độc trong những lo ngại này.
Sau nhiều năm trong một bức tranh khá tĩnh lặng về ngôn ngữ lập trình,
Go là một trong những ngôn ngữ mới đầu tiên, cùng với Rust, Elixir, Swift và nhiều hơn nữa,
đã biến việc phát triển ngôn ngữ lập trình trở thành một lĩnh vực năng động, gần như là xu hướng chủ đạo.

Go giải quyết những vấn đề này bằng cách cố gắng kết hợp sự dễ lập trình của một ngôn ngữ
thông dịch, kiểu động với hiệu suất và tính an toàn của ngôn ngữ kiểu tĩnh, biên dịch sẵn.
Nó cũng nhắm đến việc thích ứng tốt hơn với phần cứng hiện đại, với hỗ trợ cho tính toán
mạng và đa nhân. Cuối cùng, làm việc với Go được thiết kế để *nhanh*: việc build một
file thực thi lớn trên một máy tính đơn lẻ chỉ nên mất tối đa vài giây.
Đáp ứng những mục tiêu này đã khiến chúng tôi phải suy nghĩ lại một số cách tiếp cận
lập trình từ các ngôn ngữ hiện có, dẫn đến:
hệ thống kiểu dựa trên tổ hợp thay vì phân cấp;
hỗ trợ concurrency và bộ gom rác; đặc tả dependency chặt chẽ;
và nhiều hơn nữa.
Những điều này không thể giải quyết tốt bằng thư viện hay công cụ; một ngôn ngữ mới là cần thiết.

Bài viết [Go at Google](/talks/2012/splash.article)
thảo luận về bối cảnh và động lực đằng sau thiết kế của ngôn ngữ Go,
đồng thời cung cấp thêm chi tiết về nhiều câu trả lời trong FAQ này.

### Tổ tiên của Go là gì? {#ancestors}

Go chủ yếu thuộc họ C (cú pháp cơ bản),
với ảnh hưởng đáng kể từ họ Pascal/Modula/Oberon
(khai báo, package),
cộng với một số ý tưởng từ các ngôn ngữ
lấy cảm hứng từ CSP của Tony Hoare,
như Newsqueak và Limbo (concurrency).
Tuy nhiên, nó là một ngôn ngữ hoàn toàn mới xét ở mọi khía cạnh.
Ngôn ngữ được thiết kế bằng cách suy nghĩ về
những gì lập trình viên làm và cách làm cho việc lập trình, ít nhất là
loại lập trình chúng tôi làm, hiệu quả hơn, tức là thú vị hơn.

### Những nguyên tắc hướng dẫn trong thiết kế là gì? {#principles}

Khi Go được thiết kế, Java và C++ là các ngôn ngữ được dùng phổ biến nhất
để viết máy chủ, ít nhất là tại Google.
Chúng tôi cảm thấy những ngôn ngữ này đòi hỏi quá nhiều thao tác ghi chép và lặp lại.
Một số lập trình viên đã phản ứng bằng cách chuyển sang các ngôn ngữ
năng động và linh hoạt hơn như Python, nhưng phải đánh đổi hiệu suất và an toàn kiểu.
Chúng tôi cho rằng hoàn toàn có thể có được cả hiệu suất, an toàn và sự linh hoạt trong cùng một ngôn ngữ.

Go cố gắng giảm thiểu lượng việc phải gõ theo cả hai nghĩa của từ đó.
Xuyên suốt quá trình thiết kế, chúng tôi đã cố gắng giảm bớt rắc rối và
sự phức tạp. Không có khai báo trước, không có file header;
mọi thứ đều được khai báo đúng một lần. Khởi tạo biểu cảm, tự động và dễ dùng.
Cú pháp gọn nhẹ với ít từ khóa. Sự lặp lại (`foo.Foo* myFoo = new(foo.Foo)`) được giảm bớt
bằng suy luận kiểu đơn giản qua cú pháp khai báo và khởi tạo `:=`.
Và có lẽ triệt để nhất, không có phân cấp kiểu: kiểu dữ liệu chỉ *tồn tại*, không cần
tuyên bố mối quan hệ của chúng. Những đơn giản hóa này cho phép Go vừa biểu cảm vừa dễ hiểu
mà không đánh đổi năng suất.

Một nguyên tắc quan trọng khác là giữ các khái niệm trực giao với nhau.
Phương thức có thể được cài đặt cho bất kỳ kiểu nào; struct biểu diễn dữ liệu trong khi
interface biểu diễn trừu tượng; và cứ thế. Tính trực giao giúp dễ hiểu hơn
điều gì xảy ra khi các thứ kết hợp với nhau.

## Sử dụng {#Usage}

### Google có dùng Go nội bộ không? {#internal_usage}

Có. Go được dùng rộng rãi trong môi trường production bên trong Google.
Một ví dụ là máy chủ tải xuống của Google, `dl.google.com`,
phân phối các file nhị phân của Chrome và các gói cài đặt lớn khác như `apt-get`.

Go không phải ngôn ngữ duy nhất được dùng tại Google, nhưng nó là ngôn ngữ chủ chốt
cho một số lĩnh vực bao gồm
[site reliability engineering (SRE)](/talks/2013/go-sreops.slide)
và xử lý dữ liệu quy mô lớn.
Nó cũng là một phần quan trọng của phần mềm vận hành Google Cloud.

### Những công ty nào khác dùng Go? {#external_usage}

Mức độ sử dụng Go đang tăng trưởng trên toàn thế giới, đặc biệt nhưng không chỉ giới hạn
trong lĩnh vực điện toán đám mây.
Một vài dự án hạ tầng đám mây lớn được viết bằng Go là
Docker và Kubernetes, nhưng còn rất nhiều hơn thế.

Không chỉ có đám mây, như bạn có thể thấy từ danh sách
các công ty trên
[trang web go.dev](/)
cùng với một số
[câu chuyện thành công](/solutions/case-studies).
Ngoài ra, Go Wiki có một
[trang](/wiki/GoUsers),
được cập nhật thường xuyên, liệt kê một số trong nhiều công ty đang dùng Go.

Wiki cũng có một trang với các liên kết đến thêm nhiều
[câu chuyện thành công](/wiki/SuccessStories)
về các công ty và dự án đang dùng ngôn ngữ này.

### Các chương trình Go có thể liên kết với các chương trình C/C++ không? {#Do_Go_programs_link_with_Cpp_programs}

Có thể dùng C và Go cùng nhau trong cùng một không gian địa chỉ,
nhưng đây không phải sự kết hợp tự nhiên và có thể đòi hỏi phần mềm giao tiếp đặc biệt.
Ngoài ra, việc liên kết C với mã Go làm mất đi các thuộc tính
an toàn bộ nhớ và quản lý stack mà Go cung cấp.
Đôi khi bắt buộc phải dùng thư viện C để giải quyết vấn đề,
nhưng làm vậy luôn đưa vào một yếu tố rủi ro không có trong
mã Go thuần túy, vì vậy hãy cẩn thận.

Nếu bạn cần dùng C với Go, cách tiến hành phụ thuộc vào
cài đặt trình biên dịch Go.
Trình biên dịch "chuẩn", là một phần của chuỗi công cụ Go được hỗ trợ bởi
nhóm Go tại Google, được gọi là `gc`.
Ngoài ra còn có trình biên dịch dựa trên GCC (`gccgo`) và
trình biên dịch dựa trên LLVM (`gollvm`),
cũng như danh sách ngày càng tăng các trình biên dịch đặc biệt phục vụ các mục đích khác nhau,
đôi khi cài đặt một tập con của ngôn ngữ,
như [TinyGo](https://tinygo.org/).

`Gc` dùng quy ước gọi hàm và linker khác với C và
do đó không thể được gọi trực tiếp từ các chương trình C, hay ngược lại.
Chương trình [`cgo`](/cmd/cgo/) cung cấp cơ chế cho
"foreign function interface" để cho phép gọi an toàn các thư viện C từ mã Go.
SWIG mở rộng khả năng này sang các thư viện C++.

Bạn cũng có thể dùng `cgo` và SWIG với `gccgo` và `gollvm`.
Vì chúng dùng ABI truyền thống, nên cũng có thể, với rất nhiều cẩn thận,
liên kết mã từ các trình biên dịch này trực tiếp với các chương trình C hoặc C++ được biên dịch bởi GCC/LLVM.
Tuy nhiên, để làm điều đó an toàn cần phải hiểu rõ các quy ước gọi hàm của
tất cả các ngôn ngữ liên quan, cũng như phải chú ý đến giới hạn stack khi gọi C hoặc C++
từ Go.

### Go hỗ trợ những IDE nào? {#ide}

Dự án Go không đi kèm một IDE tùy chỉnh, nhưng ngôn ngữ và
các thư viện được thiết kế để dễ phân tích mã nguồn.
Do đó, hầu hết các trình soạn thảo và IDE phổ biến đều hỗ trợ Go tốt,
trực tiếp hoặc thông qua plugin.

Nhóm Go cũng hỗ trợ một language server Go cho giao thức LSP, được gọi là
[`gopls`](https://pkg.go.dev/golang.org/x/tools/gopls#section-readme).
Các công cụ hỗ trợ LSP có thể dùng `gopls` để tích hợp hỗ trợ theo ngôn ngữ cụ thể.

Danh sách các IDE và trình soạn thảo phổ biến có hỗ trợ Go tốt
bao gồm Emacs, Vim, VSCode, Atom, Eclipse, Sublime, IntelliJ
(thông qua một biến thể tùy chỉnh gọi là GoLand), và nhiều hơn nữa.
Nhiều khả năng môi trường yêu thích của bạn cũng là một môi trường năng suất để
lập trình bằng Go.

### Go có hỗ trợ protocol buffers của Google không? {#protocol_buffers}

Một dự án mã nguồn mở riêng biệt cung cấp plugin trình biên dịch và thư viện cần thiết.
Nó có sẵn tại
[github.com/golang/protobuf/](https://github.com/golang/protobuf).

## Thiết kế {#Design}

### Go có runtime không? {#runtime}

Go có một thư viện runtime mở rộng, thường chỉ được gọi là *runtime*,
là một phần của mọi chương trình Go.
Thư viện này cài đặt bộ gom rác, concurrency,
quản lý stack, và các tính năng quan trọng khác của ngôn ngữ Go.
Mặc dù nó trung tâm hơn với ngôn ngữ, runtime của Go tương tự như
`libc`, thư viện C.

Tuy nhiên, điều quan trọng cần hiểu là runtime của Go không
bao gồm một máy ảo, như được cung cấp bởi Java runtime.
Các chương trình Go được biên dịch trước thành mã máy gốc
(hoặc JavaScript hay WebAssembly, cho một số cài đặt biến thể).
Do đó, mặc dù thuật ngữ thường được dùng để mô tả môi trường ảo
mà chương trình chạy trong đó, trong Go từ "runtime"
chỉ là tên được đặt cho thư viện cung cấp các dịch vụ ngôn ngữ quan trọng.

### Chuyện gì với các định danh Unicode? {#unicode_identifiers}

Khi thiết kế Go, chúng tôi muốn đảm bảo rằng nó không
quá thiên về ASCII, điều đó có nghĩa là mở rộng không gian định danh
từ phạm vi 7-bit ASCII.
Quy tắc của Go&mdash;các ký tự định danh phải là
chữ cái hoặc chữ số theo định nghĩa của Unicode&mdash;đơn giản để hiểu
và cài đặt, nhưng có những hạn chế.
Các ký tự kết hợp bị loại trừ theo thiết kế,
điều này loại bỏ một số ngôn ngữ như Devanagari.

Quy tắc này có một hệ quả đáng tiếc khác.
Vì một định danh được xuất phải bắt đầu bằng
chữ cái viết hoa, các định danh được tạo ra từ các ký tự
trong một số ngôn ngữ, theo định nghĩa, không thể được xuất.
Hiện tại,
giải pháp duy nhất là dùng gì đó như `X日本語`,
điều này rõ ràng là không thỏa đáng.

Kể từ phiên bản đầu tiên của ngôn ngữ, đã có rất nhiều
suy nghĩ về cách tốt nhất để mở rộng không gian định danh nhằm đáp ứng
nhu cầu của lập trình viên dùng các ngôn ngữ mẹ đẻ khác.
Chính xác phải làm gì vẫn là chủ đề thảo luận đang diễn ra, và một phiên bản
tương lai của ngôn ngữ có thể tự do hơn trong định nghĩa về định danh.
Chẳng hạn, nó có thể áp dụng một số ý tưởng từ
[khuyến nghị](http://unicode.org/reports/tr31/)
của tổ chức Unicode về định danh.
Dù thế nào đi nữa, việc đó phải được thực hiện tương thích trong khi
bảo toàn (hoặc có thể mở rộng) cách chữ hoa xác định khả năng hiển thị của
định danh, vẫn là một trong những tính năng yêu thích của chúng tôi trong Go.

Hiện tại, chúng tôi có một quy tắc đơn giản có thể được mở rộng sau này
mà không phá vỡ chương trình, một quy tắc tránh các lỗi chắc chắn sẽ phát sinh
từ một quy tắc cho phép định danh mơ hồ.

### Tại sao Go không có tính năng X? {#Why_doesnt_Go_have_feature_X}

Mọi ngôn ngữ đều có những tính năng mới lạ và bỏ qua tính năng yêu thích của ai đó.
Go được thiết kế với trọng tâm là sự thuận tiện khi lập trình, tốc độ biên dịch,
tính trực giao của các khái niệm, và nhu cầu hỗ trợ các tính năng
như concurrency và bộ gom rác. Tính năng yêu thích của bạn có thể bị thiếu
vì nó không phù hợp, vì nó ảnh hưởng đến tốc độ biên dịch hoặc sự rõ ràng của thiết kế,
hoặc vì nó sẽ khiến mô hình hệ thống cơ bản quá khó hiểu.

Nếu bạn thấy khó chịu vì Go thiếu tính năng *X*,
xin hãy tha thứ cho chúng tôi và khám phá những tính năng mà Go có. Bạn có thể thấy rằng
chúng bù đắp cho sự thiếu vắng của *X* theo những cách thú vị.

### Khi nào Go có kiểu generic? {#generics}

Bản phát hành Go 1.18 đã thêm các tham số kiểu vào ngôn ngữ.
Điều này cho phép một dạng lập trình đa hình hay generic.
Xem [đặc tả ngôn ngữ](/ref/spec) và
[đề xuất](/design/43651-type-parameters) để biết thêm chi tiết.

### Tại sao Go ban đầu được phát hành mà không có kiểu generic? {#beginning_generics}

Go được dự định là một ngôn ngữ để viết các chương trình máy chủ
dễ bảo trì theo thời gian.
(Xem [bài viết này](/talks/2012/splash.article) để biết thêm bối cảnh.)
Thiết kế tập trung vào các yếu tố như khả năng mở rộng, khả năng đọc và concurrency.
Lập trình đa hình không có vẻ cần thiết cho các mục tiêu của ngôn ngữ tại thời điểm đó,
và vì vậy ban đầu bị bỏ qua để giữ sự đơn giản.

Generics tiện lợi nhưng chúng có chi phí về sự phức tạp trong
hệ thống kiểu và runtime.
Phải mất một thời gian để phát triển một thiết kế mà chúng tôi tin là mang lại giá trị
xứng đáng với sự phức tạp.

### Tại sao Go không có exception? {#exceptions}

Chúng tôi tin rằng việc gắn exception với một cấu trúc điều khiển,
như trong thành ngữ `try-catch-finally`, dẫn đến mã phức tạp rối rắm.
Nó cũng có xu hướng khuyến khích lập trình viên gắn nhãn
quá nhiều lỗi thông thường, chẳng hạn như không mở được file, là
ngoại lệ.

Go tiếp cận theo cách khác. Đối với xử lý lỗi thông thường, khả năng trả về nhiều giá trị của Go
giúp báo cáo lỗi dễ dàng mà không làm quá tải giá trị trả về.
[Một kiểu lỗi chuẩn, kết hợp với các tính năng khác của Go](/doc/articles/error_handling.html),
làm cho việc xử lý lỗi dễ chịu nhưng khá khác biệt
so với các ngôn ngữ khác.

Go cũng có một vài
hàm tích hợp để báo hiệu và phục hồi từ các điều kiện thực sự ngoại lệ.
Cơ chế phục hồi chỉ được thực thi như một phần của trạng thái của một hàm
đang bị dọn dẹp sau lỗi, đủ để xử lý thảm họa nhưng không cần cấu trúc điều khiển bổ sung và,
khi được dùng tốt, có thể dẫn đến mã xử lý lỗi gọn gàng.

Xem bài viết [Defer, Panic, and Recover](/doc/articles/defer_panic_recover.html) để biết chi tiết.
Ngoài ra, bài đăng blog [Errors are values](/blog/errors-are-values)
mô tả một cách tiếp cận để xử lý lỗi gọn gàng trong Go bằng cách chứng minh rằng,
vì lỗi chỉ là các giá trị, toàn bộ sức mạnh của Go có thể được vận dụng trong xử lý lỗi.

### Tại sao Go không có assertion? {#assertions}

Go không cung cấp assertion. Chúng thực sự tiện lợi, nhưng kinh nghiệm của chúng tôi
là lập trình viên dùng chúng như một cái nạng để tránh phải suy nghĩ
về xử lý và báo cáo lỗi đúng cách. Xử lý lỗi đúng cách nghĩa là
các máy chủ tiếp tục hoạt động thay vì crash sau một lỗi không nghiêm trọng.
Báo cáo lỗi đúng cách nghĩa là các lỗi trực tiếp và đi thẳng vào vấn đề,
giúp lập trình viên không phải giải mã một stack trace crash dài.
Lỗi chính xác đặc biệt quan trọng khi lập trình viên đọc lỗi
không quen thuộc với mã nguồn.

Chúng tôi hiểu đây là điểm gây tranh cãi. Có nhiều thứ trong
ngôn ngữ và thư viện Go khác với thực tiễn hiện đại, đơn giản vì
chúng tôi cảm thấy đôi khi đáng thử một cách tiếp cận khác.

### Tại sao xây dựng concurrency dựa trên các ý tưởng của CSP? {#csp}

Concurrency và lập trình đa luồng theo thời gian đã
tạo ra danh tiếng khó khăn. Chúng tôi tin điều này một phần là do các thiết kế phức tạp như
[pthreads](https://en.wikipedia.org/wiki/POSIX_Threads)
và một phần là do quá chú trọng vào các chi tiết cấp thấp
như mutex, condition variable và memory barrier.
Các giao diện cấp cao hơn cho phép mã đơn giản hơn nhiều, ngay cả khi vẫn còn
mutex và những thứ tương tự bên dưới.

Một trong những mô hình thành công nhất để cung cấp hỗ trợ ngôn ngữ cấp cao
cho concurrency đến từ Communicating Sequential Processes (CSP) của Hoare.
Occam và Erlang là hai ngôn ngữ nổi tiếng bắt nguồn từ CSP.
Các primitives concurrency của Go bắt nguồn từ một nhánh khác trong cây gia đình đó,
với đóng góp chính là khái niệm mạnh mẽ về channel như các đối tượng hạng nhất.
Kinh nghiệm với một số ngôn ngữ trước đây đã chỉ ra rằng mô hình CSP
phù hợp tốt với framework ngôn ngữ thủ tục.

### Tại sao dùng goroutine thay vì thread? {#goroutines}

Goroutine là một phần để làm cho concurrency dễ dùng. Ý tưởng, đã tồn tại từ lâu,
là ghép kênh các hàm thực thi độc lập&mdash;coroutine&mdash;lên một tập hợp các thread.
Khi một coroutine bị chặn, chẳng hạn bởi một system call chặn,
runtime tự động di chuyển các coroutine khác trên cùng một thread hệ điều hành
sang một thread có thể chạy khác để chúng không bị chặn.
Lập trình viên không thấy điều này, đó chính là mục đích.
Kết quả là, cái chúng tôi gọi là goroutine, có thể rất rẻ: chúng có rất ít
chi phí ngoài bộ nhớ cho stack, chỉ vài kilobyte.

Để giữ stack nhỏ, runtime của Go dùng stack có thể thay đổi kích thước, có giới hạn.
Một goroutine mới được cấp vài kilobyte, thường là đủ.
Khi không đủ, runtime tự động tăng (và giảm) bộ nhớ lưu trữ
cho stack, cho phép nhiều goroutine tồn tại trong một lượng bộ nhớ vừa phải.
Chi phí CPU trung bình khoảng ba lệnh rẻ cho mỗi lần gọi hàm.
Hoàn toàn thực tế khi tạo hàng trăm nghìn goroutine trong cùng
một không gian địa chỉ.
Nếu goroutine chỉ là các thread, tài nguyên hệ thống sẽ
cạn kiệt ở số lượng nhỏ hơn nhiều.

### Tại sao các thao tác map không được định nghĩa là nguyên tử? {#atomic_maps}

Sau nhiều thảo luận, đã quyết định rằng việc dùng map thông thường không đòi hỏi
truy cập an toàn từ nhiều goroutine, và trong những trường hợp có yêu cầu đó, map
có lẽ là một phần của một cấu trúc dữ liệu hoặc tính toán lớn hơn đã được đồng bộ hóa.
Do đó, yêu cầu tất cả các thao tác map phải giữ một mutex sẽ làm chậm
hầu hết các chương trình và thêm an toàn vào rất ít chương trình. Tuy nhiên đây không phải quyết định dễ dàng,
vì nó có nghĩa là truy cập map không kiểm soát có thể crash chương trình.

Ngôn ngữ không cấm các cập nhật map nguyên tử. Khi cần,
chẳng hạn như khi lưu trữ một chương trình không đáng tin cậy, cài đặt có thể khóa
truy cập map.

Truy cập map chỉ không an toàn khi có cập nhật đang xảy ra.
Miễn là tất cả goroutine chỉ đọc, tức là tra cứu các phần tử trong map,
bao gồm lặp qua nó bằng vòng lặp `for` `range`, và không thay đổi map
bằng cách gán cho các phần tử hay xóa,
thì việc chúng truy cập map đồng thời mà không đồng bộ hóa là an toàn.

Để hỗ trợ dùng map đúng cách, một số cài đặt của ngôn ngữ
chứa một kiểm tra đặc biệt tự động báo cáo tại runtime khi một map bị sửa đổi
không an toàn bởi thực thi đồng thời.
Ngoài ra, trong thư viện sync có một kiểu được gọi là
[`sync.Map`](https://pkg.go.dev/sync#Map) hoạt động
tốt cho một số pattern sử dụng như cache tĩnh, mặc dù nó không
phù hợp như một thay thế chung cho kiểu map tích hợp sẵn.

### Các bạn có chấp nhận thay đổi ngôn ngữ của tôi không? {#language_changes}

Mọi người thường đề xuất cải tiến ngôn ngữ, và
[danh sách thư](https://groups.google.com/group/golang-nuts)
có lịch sử phong phú về những thảo luận như vậy, nhưng rất ít trong số này được chấp nhận.

Mặc dù Go là một dự án mã nguồn mở, ngôn ngữ và thư viện được bảo vệ
bởi một [cam kết tương thích](/doc/go1compat.html) ngăn chặn
các thay đổi làm hỏng các chương trình hiện có, ít nhất ở cấp mã nguồn
(các chương trình đôi khi cần được biên dịch lại để cập nhật).
Nếu đề xuất của bạn vi phạm đặc tả Go 1, chúng tôi thậm chí không thể xem xét
ý tưởng đó, bất kể giá trị của nó.
Một bản phát hành lớn tương lai của Go có thể không tương thích với Go 1, nhưng các thảo luận
về chủ đề đó mới chỉ bắt đầu và một điều chắc chắn là:
sẽ có rất ít sự không tương thích như vậy được đưa vào trong quá trình này.
Hơn nữa, cam kết tương thích khuyến khích chúng tôi cung cấp một con đường tự động
để các chương trình cũ thích ứng nếu tình huống đó xảy ra.

Ngay cả khi đề xuất của bạn tương thích với đặc tả Go 1, nó có thể
không phù hợp với tinh thần của các mục tiêu thiết kế của Go.
Bài viết *[Go
at Google: Language Design in the Service of Software Engineering](/talks/2012/splash.article)*
giải thích nguồn gốc của Go và động lực đằng sau thiết kế của nó.

## Kiểu dữ liệu {#types}

### Go có phải là ngôn ngữ hướng đối tượng không? {#Is_Go_an_object-oriented_language}

Vừa có vừa không. Mặc dù Go có kiểu và phương thức và cho phép
phong cách lập trình hướng đối tượng, nhưng không có phân cấp kiểu.
Khái niệm "interface" trong Go cung cấp một cách tiếp cận khác mà
chúng tôi tin là dễ dùng và theo một số cách tổng quát hơn. Cũng có
các cách để nhúng kiểu vào kiểu khác để cung cấp điều gì đó
tương tự&mdash;nhưng không giống hệt&mdash;với phân lớp.
Hơn nữa, các phương thức trong Go tổng quát hơn trong C++ hoặc Java:
chúng có thể được định nghĩa cho bất kỳ loại dữ liệu nào, kể cả các kiểu tích hợp như
các số nguyên thuần túy "unboxed".
Chúng không bị giới hạn chỉ trong struct (class).

Ngoài ra, sự thiếu vắng của phân cấp kiểu làm cho "đối tượng" trong Go cảm thấy
nhẹ nhàng hơn nhiều so với các ngôn ngữ như C++ hoặc Java.

### Làm thế nào để có dynamic dispatch của phương thức? {#How_do_I_get_dynamic_dispatch_of_methods}

Cách duy nhất để có các phương thức được dispatch động là thông qua
interface. Các phương thức trên một struct hoặc bất kỳ kiểu cụ thể nào khác luôn được giải quyết tĩnh.

### Tại sao không có kế thừa kiểu? {#inheritance}

Lập trình hướng đối tượng, ít nhất là trong các ngôn ngữ nổi tiếng nhất,
liên quan đến quá nhiều thảo luận về mối quan hệ giữa các kiểu,
những mối quan hệ thường có thể được suy ra tự động. Go tiếp cận
theo cách khác.

Thay vì yêu cầu lập trình viên khai báo trước rằng hai
kiểu có liên quan, trong Go một kiểu tự động thỏa mãn bất kỳ interface nào
chỉ định một tập con các phương thức của nó. Ngoài việc giảm công ghi chép,
cách tiếp cận này có những ưu điểm thực tế. Các kiểu có thể thỏa mãn
nhiều interface cùng một lúc, mà không có sự phức tạp của đa kế thừa truyền thống.
Interface có thể rất nhẹ&mdash;một interface với
một hoặc thậm chí không có phương thức nào có thể biểu đạt một khái niệm hữu ích.
Interface có thể được thêm vào sau nếu một ý tưởng mới xuất hiện hoặc để kiểm thử,
mà không cần chú thích các kiểu gốc.
Vì không có mối quan hệ rõ ràng nào giữa các kiểu
và interface, không có phân cấp kiểu để quản lý hay thảo luận.

Có thể dùng những ý tưởng này để xây dựng thứ gì đó tương tự như
Unix pipe an toàn kiểu. Ví dụ, xem cách `fmt.Fprintf`
cho phép in định dạng ra bất kỳ đầu ra nào, không chỉ file, hoặc cách gói
`bufio` có thể hoàn toàn tách biệt với I/O file,
hoặc cách các gói `image` tạo ra các file ảnh nén.
Tất cả những ý tưởng này bắt nguồn từ một interface duy nhất
(`io.Writer`) biểu diễn một phương thức duy nhất
(`Write`). Và đó chỉ là bề nổi.
Interface của Go có ảnh hưởng sâu sắc đến cách các chương trình được cấu trúc.

Cần một thời gian để quen, nhưng phong cách phụ thuộc kiểu ngầm định này
là một trong những điều hiệu quả nhất về Go.

### Tại sao `len` là hàm chứ không phải phương thức? {#methods_on_basics}

Chúng tôi đã tranh luận vấn đề này nhưng quyết định
cài đặt `len` và các hàm tương tự như hàm là ổn trong thực tế và
không làm phức tạp thêm câu hỏi về interface (theo nghĩa kiểu Go)
của các kiểu cơ bản.

### Tại sao Go không hỗ trợ overloading phương thức và toán tử? {#overloading}

Dispatch phương thức được đơn giản hóa nếu nó không cần phải khớp kiểu cùng lúc.
Kinh nghiệm với các ngôn ngữ khác cho chúng tôi biết rằng có nhiều
phương thức cùng tên nhưng chữ ký khác nhau đôi khi hữu ích
nhưng cũng có thể gây nhầm lẫn và dễ vỡ trong thực tế. Chỉ khớp theo tên
và yêu cầu nhất quán về kiểu là một quyết định đơn giản hóa lớn
trong hệ thống kiểu của Go.

Về operator overloading, nó có vẻ tiện lợi hơn là yêu cầu tuyệt đối.
Lại như vậy, mọi thứ đơn giản hơn khi không có nó.

### Tại sao Go không có khai báo "implements"? {#implements_interface}

Một kiểu Go cài đặt một interface bằng cách cài đặt các phương thức của interface đó,
không hơn không kém. Thuộc tính này cho phép interface được định nghĩa và dùng mà không cần
sửa đổi mã hiện có. Nó cho phép một dạng
[structural typing](https://en.wikipedia.org/wiki/Structural_type_system) thúc đẩy
sự tách biệt mối quan tâm và cải thiện tái sử dụng mã, đồng thời dễ dàng hơn
để xây dựng trên các pattern xuất hiện khi mã phát triển.
Ngữ nghĩa của interface là một trong những lý do chính cho cảm giác
linh hoạt, nhẹ nhàng của Go.

Xem [câu hỏi về kế thừa kiểu](#inheritance) để biết thêm chi tiết.

### Làm thế nào để đảm bảo kiểu của tôi thỏa mãn một interface? {#guarantee_satisfies_interface}

Bạn có thể yêu cầu trình biên dịch kiểm tra rằng kiểu `T` cài đặt
interface `I` bằng cách thử gán dùng giá trị zero cho
`T` hoặc con trỏ đến `T`, tùy trường hợp:

```
type T struct{}
var _ I = T{}       // Verify that T implements I.
var _ I = (*T)(nil) // Verify that *T implements I.
```

Nếu `T` (hoặc `*T`, tương ứng) không cài đặt
`I`, lỗi sẽ bị phát hiện tại thời điểm biên dịch.

Nếu bạn muốn người dùng của một interface tường minh khai báo rằng họ cài đặt nó,
bạn có thể thêm một phương thức có tên mô tả vào tập phương thức của interface.
Ví dụ:

```
type Fooer interface {
    Foo()
    ImplementsFooer()
}
```

Khi đó, một kiểu phải cài đặt phương thức `ImplementsFooer` để là một
`Fooer`, tài liệu hóa rõ ràng thực tế đó và thông báo trong
đầu ra của [go doc](/cmd/go/#hdr-Show_documentation_for_package_or_symbol).

```
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

Hầu hết mã không dùng các ràng buộc như vậy, vì chúng giới hạn tính hữu dụng của
ý tưởng interface. Tuy nhiên đôi khi chúng cần thiết để giải quyết sự mơ hồ
giữa các interface tương tự nhau.

### Tại sao kiểu T không thỏa mãn interface Equal? {#t_and_equal_interface}

Xem xét interface đơn giản này để biểu diễn một đối tượng có thể so sánh
bản thân với một giá trị khác:

```
type Equaler interface {
    Equal(Equaler) bool
}
```

và kiểu này, `T`:

```
type T int
func (t T) Equal(u T) bool { return t == u } // does not satisfy Equaler
```

Không giống như tình huống tương tự trong một số hệ thống kiểu đa hình,
`T` không cài đặt `Equaler`.
Kiểu đối số của `T.Equal` là `T`,
không phải kiểu bắt buộc `Equaler`.

Trong Go, hệ thống kiểu không tự động nâng cấp đối số của
`Equal`; đó là trách nhiệm của lập trình viên, như
minh họa bởi kiểu `T2`, thực sự cài đặt
`Equaler`:

```
type T2 int
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }  // satisfies Equaler
```

Tuy nhiên, điều này cũng không giống các hệ thống kiểu khác, vì trong Go *bất kỳ*
kiểu nào thỏa mãn `Equaler` có thể được truyền như
đối số đến `T2.Equal`, và tại runtime chúng ta phải
kiểm tra rằng đối số có kiểu `T2`.
Một số ngôn ngữ đảm bảo điều đó tại thời điểm biên dịch.

Một ví dụ liên quan đi theo chiều ngược lại:

```
type Opener interface {
   Open() Reader
}

func (t T3) Open() *os.File
```

Trong Go, `T3` không thỏa mãn `Opener`,
mặc dù có thể trong một ngôn ngữ khác.

Trong khi đúng là hệ thống kiểu của Go làm ít hơn cho lập trình viên
trong những trường hợp như vậy, sự thiếu vắng của subtyping làm cho các quy tắc về
thỏa mãn interface rất dễ phát biểu: tên và chữ ký của hàm
có chính xác là của interface không?
Quy tắc của Go cũng dễ cài đặt hiệu quả.
Chúng tôi cảm thấy những lợi ích này bù đắp cho việc thiếu
tự động nâng cấp kiểu.

### Tôi có thể chuyển đổi []T sang []interface{} không? {#convert_slice_of_interface}

Không được trực tiếp.
Điều này bị cấm bởi đặc tả ngôn ngữ vì hai kiểu
không có cùng biểu diễn trong bộ nhớ.
Cần phải sao chép các phần tử riêng lẻ sang slice đích. Ví dụ này chuyển đổi một slice của `int` thành slice của
`interface{}`:

```
t := []int{1, 2, 3, 4}
s := make([]interface{}, len(t))
for i, v := range t {
    s[i] = v
}
```

### Tôi có thể chuyển đổi []T1 sang []T2 nếu T1 và T2 có cùng kiểu cơ bản không? {#convert_slice_with_same_underlying_type}

Dòng cuối cùng của đoạn mã mẫu này không biên dịch được.

```
type T1 int
type T2 int
var t1 T1
var x = T2(t1) // OK
var st1 []T1
var sx = ([]T2)(st1) // NOT OK
```

Trong Go, các kiểu gắn chặt với phương thức, trong đó mỗi kiểu được đặt tên có
một tập phương thức (có thể rỗng).
Quy tắc chung là bạn có thể thay đổi tên của kiểu đang được
chuyển đổi (và do đó có thể thay đổi tập phương thức của nó) nhưng bạn không thể
thay đổi tên (và tập phương thức) của các phần tử của kiểu kết hợp.
Go yêu cầu bạn phải tường minh về chuyển đổi kiểu.

### Tại sao giá trị nil error của tôi lại không bằng nil? {#nil_error}

Bên dưới, interface được cài đặt như hai phần tử, một kiểu `T`
và một giá trị `V`.
`V` là một giá trị cụ thể như `int`,
`struct` hoặc con trỏ, không bao giờ là bản thân interface, và có
kiểu `T`.
Chẳng hạn, nếu chúng ta lưu giá trị `int` 3 vào một interface,
giá trị interface kết quả có, theo sơ đồ,
(`T=int`, `V=3`).
Giá trị `V` còn được gọi là giá trị *động* của interface,
vì một biến interface cụ thể có thể giữ các giá trị `V` khác nhau
(và các kiểu `T` tương ứng)
trong quá trình thực thi chương trình.

Một giá trị interface là `nil` chỉ khi cả `V` và `T`
đều không được đặt, (`T=nil`, `V` không được đặt).
Đặc biệt, một interface `nil` sẽ luôn giữ một kiểu `nil`.
Nếu chúng ta lưu một con trỏ `nil` kiểu `*int` vào bên trong
một giá trị interface, kiểu bên trong sẽ là `*int` bất kể giá trị của con trỏ:
(`T=*int`, `V=nil`).
Vì vậy, giá trị interface đó sẽ là không phải `nil`
*ngay cả khi giá trị con trỏ `V` bên trong là* `nil`.

Tình huống này có thể gây nhầm lẫn, và phát sinh khi một giá trị `nil`
được lưu bên trong một giá trị interface như return `error`:

```
func returnsError() error {
    var p *MyError = nil
    if bad() {
        p = ErrBad
    }
    return p // Will always return a non-nil error.
}
```

Nếu mọi thứ đều ổn, hàm trả về `nil` `p`,
vì vậy giá trị trả về là một giá trị interface `error`
giữ (`T=*MyError`, `V=nil`).
Điều này có nghĩa là nếu người gọi so sánh lỗi trả về với `nil`,
nó sẽ luôn có vẻ như có lỗi ngay cả khi không có gì xảy ra.
Để trả về một `nil` `error` đúng nghĩa cho người gọi,
hàm phải trả về một `nil` tường minh:

```
func returnsError() error {
    if bad() {
        return ErrBad
    }
    return nil
}
```

Đối với các hàm trả về lỗi, nên luôn dùng kiểu `error` trong
chữ ký của chúng (như chúng tôi đã làm ở trên) thay vì một kiểu cụ thể như
`*MyError`, để giúp đảm bảo lỗi được
tạo ra đúng cách. Ví dụ,
[`os.Open`](/pkg/os/#Open)
trả về một `error` mặc dù, nếu không phải `nil`,
nó luôn có kiểu cụ thể là
[`*os.PathError`](/pkg/os/#PathError).

Các tình huống tương tự như những gì được mô tả ở đây có thể phát sinh bất cứ khi nào interface được dùng.
Chỉ cần ghi nhớ rằng nếu bất kỳ giá trị cụ thể nào
đã được lưu trong interface, interface sẽ không phải `nil`.
Để biết thêm thông tin, xem
[The Laws of Reflection](/doc/articles/laws_of_reflection.html).

### Tại sao các kiểu kích thước không có lại hoạt động kỳ lạ? {#zero_size_types}

Go hỗ trợ các kiểu kích thước không, chẳng hạn như struct không có trường
(`struct{}`) hoặc array không có phần tử (`[0]byte`).
Không có gì bạn có thể lưu trong kiểu kích thước không, nhưng những kiểu này
đôi khi hữu ích khi không cần giá trị, như trong
`map[int]struct{}` hoặc một kiểu có phương thức nhưng không có giá trị.

Các biến khác nhau có kiểu kích thước không có thể được đặt tại cùng
vị trí trong bộ nhớ.
Điều này an toàn vì không có giá trị nào có thể được lưu trong những biến đó.

Hơn nữa, ngôn ngữ không đảm bảo rằng các con trỏ đến hai biến kích thước không
khác nhau sẽ so sánh bằng nhau hay không.
Những so sánh như vậy thậm chí có thể trả về `true` tại một điểm trong chương trình
và sau đó trả về `false` tại một điểm khác, tùy thuộc vào cách
chương trình được biên dịch và thực thi.

Một vấn đề riêng biệt với các kiểu kích thước không là một con trỏ đến một
trường struct kích thước không không được chồng lên con trỏ đến một đối tượng khác trong
bộ nhớ.
Điều đó có thể gây nhầm lẫn cho bộ gom rác.
Điều này có nghĩa là nếu trường cuối cùng trong một struct có kích thước không, struct
sẽ được đệm để đảm bảo rằng con trỏ đến trường cuối cùng không
chồng lên bộ nhớ ngay sau struct.
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

sẽ in ra `2`, không phải `1`, trong hầu hết các cài đặt Go.

### Tại sao không có untagged union như trong C? {#unions}

Untagged union sẽ vi phạm các đảm bảo an toàn bộ nhớ của Go.

### Tại sao Go không có kiểu variant? {#variant_types}

Kiểu variant, còn được gọi là kiểu đại số, cung cấp một cách để chỉ định
rằng một giá trị có thể nhận một trong một tập các kiểu khác, nhưng chỉ những kiểu đó thôi.
Một ví dụ phổ biến trong lập trình hệ thống sẽ chỉ định rằng một
lỗi là lỗi mạng, lỗi bảo mật hoặc lỗi ứng dụng và cho phép người gọi phân biệt nguồn gốc của vấn đề
bằng cách kiểm tra kiểu của lỗi. Một ví dụ khác là cây cú pháp
trong đó mỗi nút có thể là một kiểu khác nhau: khai báo, câu lệnh,
phép gán và cứ thế.

Chúng tôi đã xem xét việc thêm kiểu variant vào Go, nhưng sau khi thảo luận
đã quyết định bỏ qua vì chúng chồng chéo theo những cách gây nhầm lẫn
với interface. Điều gì sẽ xảy ra nếu các phần tử của kiểu variant
bản thân là interface?

Ngoài ra, một số những gì kiểu variant giải quyết đã được bao phủ bởi
ngôn ngữ. Ví dụ lỗi dễ biểu đạt bằng cách dùng một giá trị interface
để giữ lỗi và một type switch để phân biệt trường hợp. Ví dụ
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

vì `Value` cài đặt interface rỗng.
Trong Go, kiểu phương thức phải khớp chính xác, vì vậy `Value` không
cài đặt `Copyable`.
Go tách biệt khái niệm về những gì một
kiểu làm&mdash;các phương thức của nó&mdash;khỏi cài đặt của kiểu.
Nếu hai phương thức trả về các kiểu khác nhau, chúng không làm cùng một thứ.
Các lập trình viên muốn kiểu kết quả covariant thường đang cố gắng
biểu đạt một phân cấp kiểu thông qua interface.
Trong Go, tự nhiên hơn khi có sự tách biệt rõ ràng giữa interface
và cài đặt.

## Giá trị {#Values}

### Tại sao Go không cung cấp chuyển đổi số ẩn? {#conversions}

Sự tiện lợi của chuyển đổi tự động giữa các kiểu số trong C
bị lấn át bởi sự nhầm lẫn nó gây ra. Khi nào một biểu thức không có dấu?
Giá trị lớn đến đâu? Có tràn số không? Kết quả có khả chuyển không, độc lập
với máy thực thi không?
Nó cũng làm phức tạp trình biên dịch; "usual arithmetic conversions" của C
không dễ cài đặt và không nhất quán trên các kiến trúc.
Vì lý do khả chuyển, chúng tôi đã quyết định làm cho mọi thứ rõ ràng và thẳng thắn
với chi phí là một số chuyển đổi tường minh trong mã.
Tuy nhiên, định nghĩa về hằng số trong Go&mdash;các giá trị độ chính xác tùy ý không có
chú thích về dấu và kích thước&mdash;cải thiện đáng kể tình trạng này.

Một chi tiết liên quan là, không giống như C, `int` và `int64`
là các kiểu riêng biệt ngay cả khi `int` là kiểu 64-bit. Kiểu `int`
là tổng quát; nếu bạn quan tâm đến số bit mà một số nguyên giữ, Go
khuyến khích bạn phải tường minh.

### Hằng số hoạt động như thế nào trong Go? {#constants}

Mặc dù Go nghiêm ngặt về chuyển đổi giữa các biến có kiểu số khác nhau,
hằng số trong ngôn ngữ linh hoạt hơn nhiều.
Các hằng số nghĩa đen như `23`, `3.14159`
và [`math.Pi`](/pkg/math/#pkg-constants)
chiếm một không gian số lý tưởng, với độ chính xác tùy ý và
không có tràn số hay thiếu số.
Ví dụ, giá trị của `math.Pi` được chỉ định đến 63 chữ số thập phân
trong mã nguồn, và các biểu thức hằng số liên quan đến giá trị này giữ
độ chính xác vượt quá những gì `float64` có thể giữ.
Chỉ khi hằng số hoặc biểu thức hằng số được gán cho một
biến&mdash;một vị trí bộ nhớ trong chương trình&mdash;thì
nó mới trở thành một số "máy tính" với
các thuộc tính và độ chính xác dấu phẩy động thông thường.

Ngoài ra,
vì chúng chỉ là số, không phải giá trị có kiểu, hằng số trong Go có thể được
dùng tự do hơn biến, do đó làm mềm đi một số bất tiện
xung quanh các quy tắc chuyển đổi nghiêm ngặt.
Người ta có thể viết các biểu thức như

```
sqrt2 := math.Sqrt(2)
```

mà không bị trình biên dịch phàn nàn vì số lý tưởng `2`
có thể được chuyển đổi an toàn và chính xác
thành `float64` cho lời gọi `math.Sqrt`.

Một bài đăng blog có tiêu đề [Constants](/blog/constants)
khám phá chủ đề này chi tiết hơn.

### Tại sao map được tích hợp sẵn? {#builtin_maps}

Lý do tương tự như string: chúng là cấu trúc dữ liệu mạnh mẽ và quan trọng đến mức
cung cấp một cài đặt xuất sắc với hỗ trợ cú pháp
làm cho lập trình dễ chịu hơn. Chúng tôi tin rằng cài đặt map của Go
đủ mạnh để phục vụ đại đa số các trường hợp dùng.
Nếu một ứng dụng cụ thể có thể được lợi từ một cài đặt tùy chỉnh, hoàn toàn có thể
viết một cái nhưng nó sẽ không tiện lợi về mặt cú pháp; đây có vẻ là một sự đánh đổi hợp lý.

### Tại sao map không cho phép slice làm khóa? {#map_keys}

Tra cứu map cần một toán tử bằng nhau, mà slice không cài đặt.
Chúng không cài đặt bằng nhau vì bằng nhau không được định nghĩa rõ ràng trên những kiểu như vậy;
có nhiều cân nhắc liên quan đến so sánh nông và sâu, con trỏ so với
giá trị, cách xử lý các kiểu đệ quy, và cứ thế.
Chúng tôi có thể xem lại vấn đề này&mdash;và cài đặt bằng nhau cho slice
sẽ không làm mất hiệu lực bất kỳ chương trình nào hiện có&mdash;nhưng nếu không có ý tưởng rõ ràng về
ý nghĩa của bằng nhau của slice, đơn giản hơn là bỏ qua hiện tại.

Bằng nhau được định nghĩa cho struct và array, vì vậy chúng có thể được dùng làm khóa map.

### Tại sao map, slice và channel là tham chiếu trong khi array là giá trị? {#references}

Có rất nhiều lịch sử về chủ đề này. Trong giai đoạn đầu, map và channel
về mặt cú pháp là con trỏ và không thể khai báo hoặc dùng một
instance không phải con trỏ. Ngoài ra, chúng tôi gặp khó khăn với cách array nên hoạt động.
Cuối cùng chúng tôi quyết định rằng sự tách biệt nghiêm ngặt giữa con trỏ và
giá trị làm cho ngôn ngữ khó dùng hơn. Thay đổi những
kiểu này để hoạt động như tham chiếu đến các cấu trúc dữ liệu chia sẻ liên quan đã giải quyết
những vấn đề này. Thay đổi này thêm một số phức tạp đáng tiếc vào
ngôn ngữ nhưng có ảnh hưởng lớn đến khả năng dùng: Go trở nên năng suất hơn, thoải mái hơn khi được giới thiệu.

## Viết mã {#Writing_Code}

### Thư viện được tài liệu hóa như thế nào? {#How_are_libraries_documented}

Để truy cập tài liệu từ dòng lệnh, công cụ
[go](/pkg/cmd/go/) có lệnh con
[doc](/pkg/cmd/go/#hdr-Show_documentation_for_package_or_symbol)
cung cấp giao diện văn bản đến tài liệu
cho các khai báo, file, package và cứ thế.

Trang khám phá package toàn cầu
[pkg.go.dev/pkg/](/pkg/).
chạy một máy chủ trích xuất tài liệu package từ mã nguồn Go
ở bất cứ đâu trên web
và phục vụ nó dưới dạng HTML với các liên kết đến khai báo và các phần tử liên quan.
Đây là cách dễ nhất để tìm hiểu về các thư viện Go hiện có.

Trong những ngày đầu của dự án, có một chương trình tương tự, `godoc`,
cũng có thể được chạy để trích xuất tài liệu cho các file trên máy cục bộ;
[pkg.go.dev/pkg/](/pkg/) về cơ bản là hậu duệ của nó.
Một hậu duệ khác là lệnh
[`pkgsite`](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite)
có thể, như `godoc`, được chạy cục bộ, mặc dù
nó chưa được tích hợp vào
kết quả được hiển thị bởi `go` `doc`.

### Có hướng dẫn phong cách lập trình Go không? {#Is_there_a_Go_programming_style_guide}

Không có hướng dẫn phong cách rõ ràng, mặc dù chắc chắn có
một "phong cách Go" dễ nhận biết.

Go đã thiết lập các quy ước để hướng dẫn các quyết định về
đặt tên, bố cục và tổ chức file.
Tài liệu [Effective Go](effective_go.html)
chứa một số lời khuyên về những chủ đề này.
Trực tiếp hơn, chương trình `gofmt` là một pretty-printer
có mục đích là thực thi các quy tắc bố cục; nó thay thế danh sách
các điều nên và không nên làm thông thường cho phép giải thích.
Tất cả mã Go trong kho lưu trữ, và đại đa số trong
thế giới mã nguồn mở, đã được chạy qua `gofmt`.

Tài liệu có tiêu đề
[Go Code Review Comments](/s/comments)
là một tập hợp các bài luận ngắn về các chi tiết thành ngữ Go thường
bị lập trình viên bỏ qua.
Đây là tài liệu tham khảo hữu ích cho những người thực hiện code review cho các dự án Go.

### Làm thế nào để tôi gửi bản vá cho các thư viện Go? {#How_do_I_submit_patches_to_the_Go_libraries}

Các nguồn thư viện nằm trong thư mục `src` của kho lưu trữ.
Nếu bạn muốn thực hiện một thay đổi đáng kể, hãy thảo luận trên danh sách thư trước khi bắt đầu.

Xem tài liệu
[Contributing to the Go project](contribute.html)
để biết thêm thông tin về cách tiến hành.

### Tại sao "go get" dùng HTTPS khi nhân bản kho lưu trữ? {#git_https}

Các công ty thường chỉ cho phép lưu lượng đầu ra trên các cổng TCP tiêu chuẩn 80 (HTTP)
và 443 (HTTPS), chặn lưu lượng đầu ra trên các cổng khác, bao gồm TCP cổng 9418
(git) và TCP cổng 22 (SSH).
Khi dùng HTTPS thay vì HTTP, `git` thực thi xác thực chứng chỉ theo
mặc định, cung cấp bảo vệ chống lại các cuộc tấn công man-in-the-middle, nghe lén và giả mạo.
Do đó lệnh `go get` dùng HTTPS để an toàn.

`Git` có thể được cấu hình để xác thực qua HTTPS hoặc dùng SSH thay cho HTTPS.
Để xác thực qua HTTPS, bạn có thể thêm một dòng
vào file `$HOME/.netrc` mà git tham khảo:

```
machine github.com login *USERNAME* password *APIKEY*
```

Đối với tài khoản GitHub, mật khẩu có thể là một
[personal access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/).

`Git` cũng có thể được cấu hình để dùng SSH thay cho HTTPS cho các URL khớp với một tiền tố nhất định.
Chẳng hạn, để dùng SSH cho tất cả truy cập GitHub,
thêm những dòng này vào `~/.gitconfig` của bạn:

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

Khi làm việc với các module riêng tư, nhưng dùng module proxy công khai cho dependency, bạn có thể cần đặt `GOPRIVATE`.
Xem [private modules](/ref/mod#private-modules) để biết chi tiết và các cài đặt bổ sung.

### Làm thế nào để quản lý phiên bản package bằng "go get"? {#get_version}

Chuỗi công cụ Go có một hệ thống tích hợp để quản lý các tập package liên quan có phiên bản, được gọi là *module*.
Module được giới thiệu trong [Go 1.11](/doc/go1.11#modules) và đã sẵn sàng cho môi trường production từ [1.14](/doc/go1.14#introduction).

Để tạo một dự án dùng module, chạy [`go mod init`](/ref/mod#go-mod-init).
Lệnh này tạo ra một file `go.mod` theo dõi các phiên bản dependency.

```
go mod init example/project
```

Để thêm, nâng cấp hoặc hạ cấp một dependency, chạy [`go get`](/ref/mod#go-get):

```
go get golang.org/x/text@v0.3.5
```

Xem [Tutorial: Create a module](/doc/tutorial/create-module.html) để biết thêm thông tin về cách bắt đầu.

Xem [Developing modules](/doc/#developing-modules) để có hướng dẫn về quản lý dependency với module.

Các package trong module nên duy trì khả năng tương thích ngược khi phát triển, tuân theo [quy tắc tương thích import](https://research.swtch.com/vgo-import):

> Nếu một package cũ và một package mới có cùng đường dẫn import,\
> package mới phải tương thích ngược với package cũ.

[Hướng dẫn tương thích Go 1](/doc/go1compat.html) là một tài liệu tham khảo tốt ở đây:
không xóa các tên được xuất, khuyến khích composite literal có tag, và cứ thế.
Nếu cần chức năng khác, hãy thêm một tên mới thay vì thay đổi tên cũ.

Module mã hóa điều này bằng [semantic versioning](https://semver.org/) và semantic import versioning.
Nếu cần phá vỡ tương thích, hãy phát hành một module ở phiên bản chính mới.
Các module ở phiên bản chính 2 trở lên yêu cầu một [hậu tố phiên bản chính](/ref/mod#major-version-suffixes) như một phần của đường dẫn của chúng (như `/v2`).
Điều này bảo toàn quy tắc tương thích import: các package trong các phiên bản chính khác nhau của một module có các đường dẫn riêng biệt.

## Con trỏ và Phân bổ {#Pointers}

### Khi nào các tham số hàm được truyền theo giá trị? {#pass_by_value}

Như trong tất cả các ngôn ngữ thuộc họ C, mọi thứ trong Go đều được truyền theo giá trị.
Tức là, một hàm luôn nhận được một bản sao của
thứ được truyền, như thể có một câu lệnh gán giá trị
cho tham số. Chẳng hạn, truyền một giá trị `int`
cho một hàm tạo ra một bản sao của `int`, và truyền một giá trị con trỏ
tạo ra một bản sao của con trỏ, nhưng không sao chép dữ liệu mà nó trỏ tới.
(Xem [phần sau](/doc/faq#methods_on_values_or_pointers)
để thảo luận về cách điều này ảnh hưởng đến receiver phương thức.)

Các giá trị map và slice hoạt động như con trỏ: chúng là các descriptor
chứa con trỏ đến dữ liệu map hoặc slice cơ bản. Sao chép một giá trị map hoặc
slice không sao chép dữ liệu mà nó trỏ tới. Sao chép một giá trị interface
tạo ra một bản sao của thứ được lưu trong giá trị interface. Nếu giá trị interface
giữ một struct, sao chép giá trị interface tạo ra một bản sao của
struct. Nếu giá trị interface giữ một con trỏ, sao chép giá trị interface
tạo ra một bản sao của con trỏ, nhưng lại không sao chép dữ liệu mà nó trỏ tới.

Lưu ý rằng thảo luận này là về ngữ nghĩa của các thao tác.
Các cài đặt thực tế có thể áp dụng tối ưu hóa để tránh sao chép
miễn là các tối ưu hóa không thay đổi ngữ nghĩa.

### Khi nào nên dùng con trỏ tới interface? {#pointer_to_interface}

Hầu như không bao giờ. Con trỏ tới giá trị interface chỉ xuất hiện trong những tình huống hiếm gặp và phức tạp liên quan đến việc che giấu kiểu của giá trị interface để trì hoãn việc đánh giá.

Một lỗi phổ biến là truyền con trỏ tới giá trị interface vào hàm đang chờ nhận một interface. Trình biên dịch sẽ báo lỗi này, nhưng tình huống vẫn có thể gây nhầm lẫn vì đôi khi
[con trỏ là cần thiết để thỏa mãn một interface](#different_method_sets).
Điểm mấu chốt là: mặc dù con trỏ tới một kiểu cụ thể có thể thỏa mãn một interface, thì *con trỏ tới một interface không bao giờ có thể thỏa mãn một interface*, trừ một trường hợp ngoại lệ.

Xét khai báo biến:

```
var w io.Writer
```

Hàm in `fmt.Fprintf` nhận tham số đầu tiên là một giá trị thỏa mãn `io.Writer` -- tức là thứ gì đó có cài đặt phương thức `Write` chuẩn. Vì vậy ta có thể viết:

```
fmt.Fprintf(w, "hello, world\n")
```

Nhưng nếu truyền địa chỉ của `w`, chương trình sẽ không biên dịch được.

```
fmt.Fprintf(&w, "hello, world\n") // Lỗi tại thời điểm biên dịch.
```

Trường hợp ngoại lệ duy nhất là bất kỳ giá trị nào, kể cả con trỏ tới interface, đều có thể gán cho biến kiểu interface rỗng (`interface{}`). Tuy vậy, nếu giá trị là con trỏ tới interface thì gần như chắc chắn đó là lỗi; kết quả có thể gây nhầm lẫn.

### Nên định nghĩa phương thức trên giá trị hay con trỏ? {#methods_on_values_or_pointers}

```
func (s *MyStruct) pointerMethod() { } // phương thức trên con trỏ
func (s MyStruct)  valueMethod()   { } // phương thức trên giá trị
```

Với những lập trình viên chưa quen con trỏ, sự khác biệt giữa hai ví dụ này có thể gây nhầm lẫn, nhưng thực ra tình huống khá đơn giản. Khi định nghĩa phương thức trên một kiểu, receiver (`s` trong ví dụ trên) hoạt động y như là một đối số của phương thức đó. Vì vậy, câu hỏi nên khai báo receiver là giá trị hay con trỏ cũng giống như câu hỏi đối số hàm nên là giá trị hay con trỏ. Có một số yếu tố cần xem xét.

Thứ nhất và quan trọng nhất: phương thức có cần sửa đổi receiver không? Nếu có, receiver *phải* là con trỏ. (Slice và map hoạt động như tham chiếu, nên câu chuyện của chúng phức tạp hơn một chút, nhưng ví dụ để thay đổi độ dài của slice trong một phương thức thì receiver vẫn phải là con trỏ.) Trong ví dụ trên, nếu `pointerMethod` sửa đổi các trường của `s`, người gọi sẽ thấy những thay đổi đó, còn `valueMethod` được gọi với một bản sao của đối số từ người gọi (đó là định nghĩa của truyền giá trị), nên những thay đổi nó thực hiện sẽ không hiển thị với người gọi.

Nhân tiện, trong Java receiver phương thức luôn là con trỏ, mặc dù bản chất con trỏ của chúng bị che giấu phần nào (và những phát triển gần đây đang đưa value receiver vào Java). Chính value receiver trong Go mới là điều bất thường.

Thứ hai là vấn đề hiệu quả. Nếu receiver lớn, chẳng hạn là một `struct` to, thì dùng pointer receiver có thể tiết kiệm hơn.

Tiếp theo là sự nhất quán. Nếu một số phương thức của kiểu phải có pointer receiver, phần còn lại cũng nên như vậy, để tập phương thức nhất quán bất kể kiểu được dùng thế nào. Xem phần về [tập phương thức](#different_method_sets) để biết thêm chi tiết.

Với các kiểu như kiểu cơ bản, slice, và `struct` nhỏ, value receiver rất rẻ, vì vậy trừ khi ngữ nghĩa của phương thức yêu cầu con trỏ, value receiver là lựa chọn hiệu quả và rõ ràng.

### Sự khác biệt giữa new và make là gì? {#new_and_make}

Nói ngắn gọn: `new` cấp phát bộ nhớ, còn `make` khởi tạo các kiểu slice, map và channel.

Xem [phần liên quan trong Effective Go](/doc/effective_go.html#allocation_new) để biết thêm chi tiết.

### Kích thước của `int` trên máy 64 bit là bao nhiêu? {#q_int_sizes}

Kích thước của `int` và `uint` phụ thuộc vào cài đặt cụ thể nhưng bằng nhau trên cùng một nền tảng. Để đảm bảo tính di động, mã phụ thuộc vào kích thước cụ thể của giá trị nên dùng kiểu có kích thước tường minh, như `int64`. Trên máy 32-bit, trình biên dịch dùng số nguyên 32-bit theo mặc định, còn trên máy 64-bit, số nguyên có 64 bit. (Trong lịch sử, điều này không phải lúc nào cũng đúng.)

Mặt khác, số thực và kiểu số phức luôn có kích thước xác định (không có kiểu cơ bản `float` hay `complex`), vì lập trình viên cần ý thức về độ chính xác khi dùng số thực. Kiểu mặc định cho hằng số thực (không có kiểu) là `float64`. Vì vậy `foo` `:=` `3.0` khai báo biến `foo` kiểu `float64`. Với biến `float32` được khởi tạo bởi hằng số (không có kiểu), kiểu biến phải được chỉ định tường minh trong khai báo:

```
var foo float32 = 3.0
```

Hoặc hằng số phải được gán kiểu bằng một phép chuyển đổi như `foo := float32(3.0)`.

### Làm sao biết một biến được cấp phát trên heap hay stack? {#stack_or_heap}

Về mặt tính đúng đắn, bạn không cần biết. Mỗi biến trong Go tồn tại miễn là còn tham chiếu tới nó. Vị trí lưu trữ mà cài đặt chọn không ảnh hưởng đến ngữ nghĩa của ngôn ngữ.

Tuy nhiên, vị trí lưu trữ có ảnh hưởng đến việc viết chương trình hiệu quả. Khi có thể, trình biên dịch Go sẽ cấp phát các biến cục bộ của hàm trong stack frame của hàm đó. Nhưng nếu trình biên dịch không thể chứng minh rằng biến không được tham chiếu sau khi hàm trả về, trình biên dịch phải cấp phát biến trên heap được gom rác để tránh lỗi con trỏ treo. Ngoài ra, nếu một biến cục bộ rất lớn, có thể hợp lý hơn khi lưu nó trên heap thay vì stack.

Trong các trình biên dịch hiện tại, nếu một biến được lấy địa chỉ, biến đó là ứng viên để cấp phát trên heap. Tuy nhiên, phân tích *escape* cơ bản nhận ra một số trường hợp khi các biến đó sẽ không tồn tại sau khi hàm trả về và có thể nằm trên stack.

### Tại sao tiến trình Go của tôi dùng nhiều bộ nhớ ảo thế? {#Why_does_my_Go_process_use_so_much_virtual_memory}

Bộ cấp phát bộ nhớ Go đặt trước một vùng bộ nhớ ảo lớn làm arena để cấp phát. Bộ nhớ ảo này là riêng của tiến trình Go cụ thể đó; việc đặt trước này không lấy đi bộ nhớ của các tiến trình khác.

Để biết lượng bộ nhớ thực sự được cấp phát cho một tiến trình Go, hãy dùng lệnh `top` trên Unix và xem cột `RES` (Linux) hoặc `RSIZE` (macOS).
<!-- TODO(adg): find out how this works on Windows -->

## Đồng thời {#Concurrency}

### Những thao tác nào là nguyên tử? Còn mutex thì sao? {#What_operations_are_atomic_What_about_mutexes}

Mô tả về tính nguyên tử của các thao tác trong Go có thể tìm thấy trong tài liệu [Go Memory Model](/ref/mem).

Các primitive đồng bộ hóa cấp thấp và nguyên tử có trong các gói [sync](/pkg/sync) và [sync/atomic](/pkg/sync/atomic). Các gói này phù hợp với các tác vụ đơn giản như tăng bộ đếm tham chiếu hay đảm bảo loại trừ lẫn nhau ở quy mô nhỏ.

Với các thao tác cấp cao hơn, chẳng hạn như phối hợp giữa các server đồng thời, các kỹ thuật cấp cao hơn có thể tạo ra chương trình tốt hơn, và Go hỗ trợ cách tiếp cận này thông qua goroutine và channel. Ví dụ, bạn có thể cấu trúc chương trình sao cho chỉ một goroutine tại một thời điểm chịu trách nhiệm với một phần dữ liệu cụ thể. Cách tiếp cận đó được tóm tắt bởi câu châm ngôn Go gốc:

Đừng giao tiếp bằng cách chia sẻ bộ nhớ. Thay vào đó, hãy chia sẻ bộ nhớ bằng cách giao tiếp.

Xem bài đi bộ mã [Share Memory By Communicating](/doc/codewalk/sharemem/) và [bài viết liên quan](/blog/share-memory-by-communicating) để thảo luận chi tiết về khái niệm này.

Các chương trình đồng thời lớn có thể sẽ vay mượn từ cả hai bộ công cụ này.

### Tại sao chương trình của tôi không chạy nhanh hơn khi có nhiều CPU? {#parallel_slow}

Liệu một chương trình có chạy nhanh hơn với nhiều CPU hay không phụ thuộc vào bài toán nó đang giải. Ngôn ngữ Go cung cấp các primitive đồng thời như goroutine và channel, nhưng đồng thời chỉ cho phép song song khi bài toán cơ bản về bản chất là có thể song song. Các bài toán về bản chất tuần tự không thể tăng tốc bằng cách thêm CPU, trong khi những bài toán có thể chia thành các phần thực thi song song thì có thể tăng tốc, đôi khi đáng kể.

Đôi khi thêm CPU có thể làm chương trình chậm lại. Trên thực tế, các chương trình dành nhiều thời gian đồng bộ hóa hoặc giao tiếp hơn là thực hiện tính toán hữu ích có thể bị giảm hiệu suất khi dùng nhiều luồng OS. Điều này là do truyền dữ liệu giữa các luồng liên quan đến việc chuyển đổi ngữ cảnh, tốn chi phí đáng kể, và chi phí đó có thể tăng với nhiều CPU hơn. Ví dụ, [ví dụ sàng số nguyên tố](/ref/spec#An_example_package) từ đặc tả Go không có tính song song đáng kể mặc dù nó khởi chạy nhiều goroutine; tăng số luồng (CPU) có khả năng làm chậm hơn là tăng tốc.

Để biết thêm chi tiết về chủ đề này, xem bài nói chuyện [Concurrency is not Parallelism](/blog/concurrency-is-not-parallelism).

### Làm sao kiểm soát số lượng CPU? {#number_cpus}

Số lượng CPU có sẵn đồng thời cho các goroutine đang thực thi được kiểm soát bởi biến môi trường shell `GOMAXPROCS`, giá trị mặc định là số lõi CPU có sẵn. Các chương trình có tiềm năng thực thi song song do đó sẽ đạt được điều đó theo mặc định trên máy nhiều CPU. Để thay đổi số lượng CPU song song, hãy đặt biến môi trường hoặc dùng [hàm cùng tên](/pkg/runtime/#GOMAXPROCS) của gói runtime để cấu hình hỗ trợ run-time sử dụng số luồng khác. Đặt thành 1 loại bỏ khả năng song song thực sự, buộc các goroutine độc lập thay phiên nhau thực thi.

Runtime có thể cấp phát nhiều luồng hơn giá trị của `GOMAXPROCS` để phục vụ nhiều yêu cầu I/O đang chờ. `GOMAXPROCS` chỉ ảnh hưởng đến số lượng goroutine có thể thực sự thực thi cùng lúc; nhiều goroutine tùy ý có thể bị chặn trong các lời gọi hệ thống.

Bộ lập lịch goroutine của Go thực hiện tốt việc cân bằng goroutine và luồng, và thậm chí có thể ưu tiên thực thi của goroutine để đảm bảo những goroutine khác trên cùng luồng không bị bỏ đói. Tuy nhiên, nó không hoàn hảo. Nếu bạn thấy vấn đề về hiệu suất, đặt `GOMAXPROCS` cho từng ứng dụng có thể giúp ích.

### Tại sao không có goroutine ID? {#no_goroutine_id}

Goroutine không có tên; chúng chỉ là những worker ẩn danh. Chúng không cung cấp định danh duy nhất, tên, hay cấu trúc dữ liệu nào cho lập trình viên. Một số người ngạc nhiên về điều này, mong đợi câu lệnh `go` trả về một thứ gì đó có thể dùng để truy cập và kiểm soát goroutine sau này.

Lý do cơ bản khiến goroutine ẩn danh là để toàn bộ ngôn ngữ Go có sẵn khi lập trình mã đồng thời. Ngược lại, các mẫu sử dụng hình thành khi luồng và goroutine được đặt tên có thể hạn chế những gì một thư viện dùng chúng có thể làm.

Đây là một minh họa về những khó khăn. Khi đặt tên cho một goroutine và xây dựng mô hình xung quanh nó, nó trở nên đặc biệt, và người ta bị cám dỗ liên kết tất cả tính toán với goroutine đó, bỏ qua khả năng sử dụng nhiều goroutine, có thể dùng chung, cho việc xử lý. Nếu gói `net/http` liên kết trạng thái per-request với một goroutine, client sẽ không thể dùng thêm goroutine khi phục vụ một request.

Ngoài ra, kinh nghiệm với các thư viện như những thư viện cho hệ thống đồ họa yêu cầu tất cả xử lý xảy ra trên "luồng chính" đã cho thấy cách tiếp cận này phức tạp và hạn chế như thế nào khi triển khai trong ngôn ngữ đồng thời. Sự tồn tại của một luồng hay goroutine đặc biệt buộc lập trình viên phải bóp méo chương trình để tránh crash và các vấn đề khác gây ra bởi việc vô tình thao tác trên sai luồng.

Với những trường hợp mà một goroutine cụ thể thực sự đặc biệt, ngôn ngữ cung cấp các tính năng như channel có thể dùng theo những cách linh hoạt để tương tác với nó.

## Hàm và Phương thức {#Functions_methods}

### Tại sao T và *T có tập phương thức khác nhau? {#different_method_sets}

Như [đặc tả Go](/ref/spec#Types) nói, tập phương thức của kiểu `T` bao gồm tất cả các phương thức có receiver kiểu `T`, trong khi kiểu con trỏ tương ứng `*T` bao gồm tất cả các phương thức có receiver `*T` hoặc `T`. Điều đó có nghĩa là tập phương thức của `*T` bao gồm tập phương thức của `T`, nhưng không có chiều ngược lại.

Sự phân biệt này xuất phát từ việc nếu một giá trị interface chứa một con trỏ `*T`, một lời gọi phương thức có thể lấy được giá trị bằng cách dereference con trỏ, nhưng nếu một giá trị interface chứa một giá trị `T`, không có cách an toàn nào để lời gọi phương thức lấy được con trỏ. (Làm vậy sẽ cho phép phương thức sửa đổi nội dung của giá trị bên trong interface, điều này không được phép bởi đặc tả ngôn ngữ.)

Ngay cả trong những trường hợp trình biên dịch có thể lấy địa chỉ của giá trị để truyền cho phương thức, nếu phương thức sửa đổi giá trị thì các thay đổi sẽ bị mất ở người gọi.

Ví dụ, nếu đoạn mã dưới đây hợp lệ:

```
var buf bytes.Buffer
io.Copy(buf, os.Stdin)
```

nó sẽ sao chép đầu vào chuẩn vào một *bản sao* của `buf`, không phải vào `buf` thực sự. Đây hầu như không bao giờ là hành vi mong muốn và do đó bị ngôn ngữ cấm.

### Điều gì xảy ra với closure chạy như goroutine? {#closures_and_goroutines}

Do cách biến vòng lặp hoạt động, trước Go phiên bản 1.22 (xem phần cuối mục này để cập nhật), một số nhầm lẫn có thể nảy sinh khi dùng closure với đồng thời. Xét chương trình sau:

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

Người ta có thể nhầm mong đợi thấy `a, b, c` là đầu ra. Thay vào đó, bạn có thể sẽ thấy `c, c, c`. Điều này là vì mỗi lần lặp của vòng lặp dùng cùng một instance của biến `v`, nên mỗi closure chia sẻ biến đơn đó. Khi closure chạy, nó in giá trị của `v` tại thời điểm `fmt.Println` thực thi, nhưng `v` có thể đã bị thay đổi kể từ khi goroutine được khởi chạy. Để giúp phát hiện vấn đề này và các vấn đề khác trước khi chúng xảy ra, hãy chạy [`go vet`](/cmd/go/#hdr-Run_go_tool_vet_on_packages).

Để gắn giá trị hiện tại của `v` với mỗi closure khi nó được khởi chạy, cần sửa vòng lặp bên trong để tạo một biến mới ở mỗi lần lặp. Một cách là truyền biến làm đối số cho closure:

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

Trong ví dụ này, giá trị của `v` được truyền làm đối số cho hàm ẩn danh. Giá trị đó sau đó có thể truy cập bên trong hàm qua biến `u`.

Còn đơn giản hơn là chỉ cần tạo một biến mới, dùng kiểu khai báo có vẻ lạ nhưng hoạt động tốt trong Go:

{{raw `
<pre>
    for _, v := range values {
        <b>v := v</b> // create a new 'v'.
        go func() {
            fmt.Println(<b>v</b>)
            done &lt;- true
        }()
    }
</pre>
`}}

Hành vi này của ngôn ngữ, không định nghĩa biến mới cho mỗi lần lặp, về sau được coi là một sai lầm, và đã được khắc phục trong [Go 1.22](/wiki/LoopvarExperiment), phiên bản này thực sự tạo một biến mới cho mỗi lần lặp, loại bỏ vấn đề này.

## Luồng điều khiển {#Control_flow}

### Tại sao Go không có toán tử `?:`? {#Does_Go_have_a_ternary_form}

Không có thao tác kiểm tra ba ngôi trong Go. Bạn có thể dùng cách sau để đạt kết quả tương tự:

```
if expr {
    n = trueVal
} else {
    n = falseVal
}
```

Lý do `?:` vắng mặt trong Go là vì những người thiết kế ngôn ngữ đã thấy toán tử này được dùng quá thường xuyên để tạo ra các biểu thức phức tạp đến mức không thể hiểu được. Dạng `if-else`, mặc dù dài hơn, nhưng rõ ràng hơn không thể phủ nhận. Một ngôn ngữ chỉ cần một cấu trúc điều khiển luồng điều kiện duy nhất.

## Tham số Kiểu {#Type_Parameters}

### Tại sao Go có tham số kiểu? {#why_generics}

Tham số kiểu cho phép lập trình tổng quát, trong đó các hàm và cấu trúc dữ liệu được định nghĩa theo các kiểu được chỉ định sau, khi những hàm và cấu trúc dữ liệu đó được sử dụng. Ví dụ, chúng giúp viết một hàm trả về giá trị nhỏ nhất trong hai giá trị của bất kỳ kiểu có thứ tự nào, mà không cần viết phiên bản riêng cho mỗi kiểu có thể. Để giải thích sâu hơn kèm ví dụ, xem bài đăng blog [Why Generics?](/blog/why-generics).

### Generics được cài đặt thế nào trong Go? {#generics_implementation}

Trình biên dịch có thể chọn biên dịch mỗi instantiation riêng lẻ hay biên dịch các instantiation tương tự thành một cài đặt duy nhất. Cách tiếp cận cài đặt đơn tương tự như hàm với tham số interface. Các trình biên dịch khác nhau sẽ đưa ra lựa chọn khác nhau cho các trường hợp khác nhau. Trình biên dịch Go chuẩn thường phát ra một instantiation duy nhất cho mỗi đối số kiểu có cùng hình dạng, trong đó hình dạng được xác định bởi các thuộc tính của kiểu như kích thước và vị trí của các con trỏ mà nó chứa. Các bản phát hành trong tương lai có thể thử nghiệm sự đánh đổi giữa thời gian biên dịch, hiệu quả run-time và kích thước mã.

### Generics trong Go so sánh với generics trong các ngôn ngữ khác thế nào? {#generics_comparison}

Chức năng cơ bản trong mọi ngôn ngữ là tương tự: có thể viết các kiểu và hàm sử dụng các kiểu được chỉ định sau. Tuy nhiên, có một số khác biệt.

* Java

	Trong Java, trình biên dịch kiểm tra kiểu generic tại thời điểm biên dịch nhưng loại bỏ kiểu đó tại run time. Điều này được gọi là
	[type erasure](https://en.wikipedia.org/wiki/Generics_in_Java#Problems_with_type_erasure).
	Ví dụ, một kiểu Java gọi là `List<Integer>` tại thời điểm biên dịch sẽ trở thành kiểu không generic `List` tại run time. Điều này có nghĩa là, khi dùng dạng reflection của Java, không thể phân biệt một giá trị kiểu `List<Integer>` với một giá trị kiểu `List<Float>`. Trong Go, thông tin reflection cho một kiểu generic bao gồm đầy đủ thông tin kiểu tại thời điểm biên dịch.

	Java dùng type wildcard như {{raw `<code>List&lt;? extends Number&gt;</code>`}} hoặc {{raw `<code>List<? super Number></code>`}} để cài đặt covariance và contravariance generic. Go không có những khái niệm này, điều này làm cho kiểu generic trong Go đơn giản hơn nhiều.

* C++

	Theo truyền thống, template C++ không áp đặt bất kỳ ràng buộc nào lên các đối số kiểu, mặc dù C++20 hỗ trợ các ràng buộc tùy chọn thông qua [concepts](https://en.wikipedia.org/wiki/Concepts_(C%2B%2B)). Trong Go, ràng buộc là bắt buộc với tất cả tham số kiểu. C++20 concepts được biểu diễn bằng các đoạn mã nhỏ phải biên dịch được với các đối số kiểu. Ràng buộc Go là các kiểu interface xác định tập hợp tất cả các đối số kiểu được phép.

	C++ hỗ trợ template metaprogramming; Go thì không. Trên thực tế, tất cả trình biên dịch C++ biên dịch mỗi template tại điểm nó được khởi tạo; như đã đề cập ở trên, Go có thể và thực sự dùng các cách tiếp cận khác nhau cho các instantiation khác nhau.

* Rust

	Phiên bản ràng buộc của Rust được gọi là trait bound. Trong Rust, liên kết giữa trait bound và một kiểu phải được định nghĩa tường minh, trong crate định nghĩa trait bound hoặc crate định nghĩa kiểu. Trong Go, các đối số kiểu ngầm thỏa mãn ràng buộc, giống như các kiểu Go ngầm cài đặt kiểu interface. Thư viện chuẩn Rust định nghĩa các trait chuẩn cho các thao tác như so sánh hay cộng; thư viện chuẩn Go không làm vậy, vì những thao tác này có thể được biểu diễn trong mã người dùng qua kiểu interface. Ngoại lệ duy nhất là interface được định nghĩa sẵn `comparable` của Go, nắm bắt một thuộc tính không thể biểu diễn trong hệ thống kiểu.

* Python

	Python không phải là ngôn ngữ kiểu tĩnh, vì vậy người ta có thể nói một cách hợp lý rằng tất cả các hàm Python luôn là generic theo mặc định: chúng luôn có thể được gọi với giá trị của bất kỳ kiểu nào, và mọi lỗi kiểu đều được phát hiện tại run time.

### Tại sao Go dùng dấu ngoặc vuông cho danh sách tham số kiểu? {#generic_brackets}

Java và C++ dùng dấu ngoặc nhọn cho danh sách tham số kiểu, như `List<Integer>` trong Java và `std::vector<int>` trong C++. Tuy nhiên, tùy chọn đó không có sẵn cho Go, vì nó dẫn đến vấn đề cú pháp: khi phân tích mã bên trong hàm, chẳng hạn như `v := F<T>`, tại điểm thấy {{raw `<code>&lt;</code>`}} thì mơ hồ không biết đang thấy một instantiation hay một biểu thức dùng toán tử {{raw `<code>&lt;</code>`}}. Điều này rất khó giải quyết mà không có thông tin kiểu.

Ví dụ, xét câu lệnh như:

{{raw `
<pre>
    a, b = w &lt; x, y &gt; (z)
</pre>
`}}

Không có thông tin kiểu, không thể quyết định liệu vế phải của phép gán là một cặp biểu thức ({{raw `<code>w &lt; x</code>`}} và {{raw `<code>y &gt; z</code>`}}), hay là một instantiation hàm generic và lời gọi trả về hai giá trị kết quả ({{raw `<code>(w&lt;x, y&gt;)(z)</code>`}}).

Đây là quyết định thiết kế quan trọng của Go là việc phân tích cú pháp có thể thực hiện mà không cần thông tin kiểu, điều này có vẻ không thể khi dùng dấu ngoặc nhọn cho generics.

Go không phải là ngôn ngữ duy nhất hay đầu tiên dùng dấu ngoặc vuông; có các ngôn ngữ khác như Scala cũng dùng dấu ngoặc vuông cho mã generic.

### Tại sao Go không hỗ trợ phương thức với tham số kiểu? {#generic_methods}

Go cho phép kiểu generic có phương thức, nhưng ngoài receiver, các đối số của những phương thức đó không thể dùng kiểu được tham số hóa. Chúng ta không dự đoán rằng Go sẽ thêm phương thức generic.

Vấn đề là cách cài đặt chúng. Cụ thể, xét việc kiểm tra liệu một giá trị trong interface có cài đặt một interface khác với các phương thức bổ sung hay không. Ví dụ, xét kiểu này, một struct rỗng với phương thức `Nop` generic trả về đối số của nó, cho bất kỳ kiểu nào có thể:

```
type Empty struct{}

func (Empty) Nop[T any](x T) T {
	return x
}
```

Giả sử một giá trị `Empty` được lưu trong `any` và truyền cho mã khác kiểm tra xem nó có thể làm gì:

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

Đoạn mã đó hoạt động thế nào nếu `x` là một `Empty`? Có vẻ như `x` phải thỏa mãn cả ba kiểm tra, cùng với bất kỳ hình thức nào khác với bất kỳ kiểu nào khác.

Mã nào chạy khi những phương thức đó được gọi? Với các phương thức không generic, trình biên dịch tạo mã cho tất cả các cài đặt phương thức và liên kết chúng vào chương trình cuối cùng. Nhưng với các phương thức generic, có thể có số cài đặt phương thức vô hạn, nên cần một chiến lược khác.

Có bốn lựa chọn:

 1. Tại thời điểm liên kết, lập danh sách tất cả các kiểm tra interface động có thể có, sau đó tìm kiếm các kiểu thỏa mãn chúng nhưng thiếu phương thức đã biên dịch, và sau đó kích hoạt lại trình biên dịch để thêm các phương thức đó.

    Điều này sẽ làm chậm đáng kể việc build, vì cần dừng lại sau khi liên kết và lặp lại một số lần biên dịch. Nó sẽ đặc biệt làm chậm các build gia tăng. Tệ hơn, có thể mã phương thức mới biên dịch sẽ chính nó có các kiểm tra interface động mới, và quá trình sẽ phải lặp lại. Có thể xây dựng các ví dụ mà quá trình không bao giờ kết thúc.

 2. Cài đặt một loại JIT, biên dịch mã phương thức cần thiết tại runtime.

    Go được hưởng lợi rất nhiều từ sự đơn giản và hiệu suất có thể dự đoán được của việc biên dịch ahead-of-time thuần túy. Chúng ta không muốn gánh thêm sự phức tạp của JIT chỉ để cài đặt một tính năng ngôn ngữ.

 3. Sắp xếp để phát ra một fallback chậm cho mỗi phương thức generic sử dụng bảng hàm cho mọi thao tác ngôn ngữ có thể có trên tham số kiểu, và sau đó dùng cài đặt fallback đó cho các kiểm tra động.

    Cách tiếp cận này sẽ làm cho phương thức generic được tham số hóa bởi một kiểu bất ngờ chậm hơn nhiều so với cùng phương thức được tham số hóa bởi một kiểu quan sát được tại thời điểm biên dịch. Điều này sẽ làm hiệu suất kém có thể dự đoán hơn nhiều.

 4. Định nghĩa rằng các phương thức generic không thể dùng để thỏa mãn interface.

    Interface là phần thiết yếu của lập trình trong Go. Không cho phép phương thức generic thỏa mãn interface là không thể chấp nhận từ góc độ thiết kế.

Không có lựa chọn nào trong số này là tốt, vì vậy chúng ta đã chọn "không có lựa chọn nào ở trên."

Thay vì phương thức với tham số kiểu, hãy dùng hàm top-level với tham số kiểu, hoặc thêm tham số kiểu vào kiểu receiver.

Để biết thêm chi tiết, bao gồm thêm ví dụ, xem [đề xuất](/design/43651-type-parameters#no-parameterized-methods).

### Tại sao tôi không thể dùng kiểu cụ thể hơn cho receiver của kiểu được tham số hóa? {#types_in_method_declaration}

Khai báo phương thức của kiểu generic được viết với receiver bao gồm tên tham số kiểu. Có lẽ do sự tương đồng của cú pháp để chỉ định kiểu tại nơi gọi, một số người đã nghĩ rằng điều này cung cấp cơ chế để tạo ra phương thức được tùy chỉnh cho các đối số kiểu nhất định bằng cách đặt tên kiểu cụ thể trong receiver, chẳng hạn như `string`:

```
type S[T any] struct { f T }

func (s S[string]) Add(t string) string {
    return s.f + t
}
```

Điều này thất bại vì từ `string` được trình biên dịch hiểu là tên của đối số kiểu trong phương thức. Thông báo lỗi của trình biên dịch sẽ có dạng như "`operator + not defined on s.f (variable of type string)`". Điều này có thể gây nhầm lẫn vì toán tử `+` hoạt động tốt trên kiểu định sẵn `string`, nhưng khai báo đã ghi đè, cho phương thức này, định nghĩa của `string`, và toán tử không hoạt động trên phiên bản không liên quan của `string` đó. Việc ghi đè một tên được định sẵn như vậy là hợp lệ, nhưng là điều kỳ lạ để làm và thường là lỗi.

### Tại sao trình biên dịch không thể suy luận đối số kiểu trong chương trình của tôi? {#type_inference}

Có nhiều trường hợp mà lập trình viên có thể dễ dàng thấy đối số kiểu cho một kiểu hoặc hàm generic phải là gì, nhưng ngôn ngữ không cho phép trình biên dịch suy luận ra nó. Suy luận kiểu được giới hạn có chủ ý để đảm bảo không bao giờ có sự nhầm lẫn về kiểu nào được suy luận. Kinh nghiệm với các ngôn ngữ khác cho thấy suy luận kiểu bất ngờ có thể dẫn đến nhầm lẫn đáng kể khi đọc và debug chương trình. Luôn luôn có thể chỉ định tường minh đối số kiểu để dùng trong lời gọi. Trong tương lai, các hình thức suy luận mới có thể được hỗ trợ, miễn là các quy tắc vẫn đơn giản và rõ ràng.

## Gói và Kiểm thử {#Packages_Testing}

### Làm sao tạo một gói nhiều tệp? {#How_do_I_create_a_multifile_package}

Đặt tất cả các tệp nguồn cho gói vào một thư mục riêng. Các tệp nguồn có thể tham chiếu đến các mục từ các tệp khác theo ý muốn; không cần khai báo trước hay tệp header.

Ngoài việc được chia thành nhiều tệp, gói sẽ biên dịch và kiểm thử giống hệt như một gói một tệp.

### Làm sao viết unit test? {#How_do_I_write_a_unit_test}

Tạo một tệp mới kết thúc bằng `_test.go` trong cùng thư mục với các tệp nguồn của gói. Bên trong tệp đó, `import "testing"` và viết các hàm có dạng:

```
func TestFoo(t *testing.T) {
    ...
}
```

Chạy `go test` trong thư mục đó. Script đó tìm các hàm `Test`, xây dựng binary kiểm thử và chạy nó.

Xem tài liệu [How to Write Go Code](/doc/code.html), gói [`testing`](/pkg/testing/) và subcommand [`go test`](/cmd/go/#hdr-Test_packages) để biết thêm chi tiết.

### Hàm trợ giúp yêu thích của tôi cho kiểm thử ở đâu? {#testing_framework}

Gói [`testing`](/pkg/testing/) chuẩn của Go giúp dễ dàng viết unit test, nhưng nó thiếu các tính năng được cung cấp trong các framework kiểm thử của ngôn ngữ khác như hàm assertion. Một [phần trước](#assertions) của tài liệu này giải thích tại sao Go không có assertion, và các lập luận tương tự áp dụng cho việc dùng `assert` trong test. Xử lý lỗi đúng cách có nghĩa là để các test khác chạy sau khi một test thất bại, để người debug thất bại có được bức tranh đầy đủ về những gì sai. Sẽ hữu ích hơn khi test báo cáo rằng `isPrime` đưa ra câu trả lời sai cho 2, 3, 5 và 7 (hoặc cho 2, 4, 8 và 16) hơn là báo cáo rằng `isPrime` đưa ra câu trả lời sai cho 2 và do đó không có thêm test nào được chạy. Lập trình viên kích hoạt thất bại test có thể không quen với mã thất bại. Thời gian đầu tư viết thông báo lỗi tốt sẽ được đền đáp sau khi test bị hỏng.

Một điểm liên quan là các framework kiểm thử có xu hướng phát triển thành các ngôn ngữ mini của riêng chúng, với các điều kiện, điều khiển và cơ chế in ấn, nhưng Go đã có tất cả những khả năng đó; tại sao phải tạo lại chúng? Chúng ta muốn viết test bằng Go; đó là một ngôn ngữ ít hơn phải học và cách tiếp cận giữ cho test thẳng thắn và dễ hiểu.

Nếu lượng mã phụ thêm cần thiết để viết các lỗi tốt có vẻ lặp đi lặp lại và quá nhiều, test có thể hoạt động tốt hơn nếu dạng bảng, lặp qua danh sách đầu vào và đầu ra được định nghĩa trong một cấu trúc dữ liệu (Go có hỗ trợ tuyệt vời cho literal cấu trúc dữ liệu). Công sức để viết test tốt và thông báo lỗi tốt khi đó sẽ được phân bổ qua nhiều trường hợp kiểm thử. Thư viện chuẩn Go chứa đầy các ví dụ minh họa, chẳng hạn như trong [các kiểm thử định dạng cho gói `fmt`](/src/fmt/fmt_test.go).

### Tại sao *X* không có trong thư viện chuẩn? {#x_in_std}

Mục đích của thư viện chuẩn là hỗ trợ thư viện runtime, kết nối với hệ điều hành, và cung cấp chức năng chính mà nhiều chương trình Go yêu cầu, chẳng hạn như I/O định dạng và mạng. Nó cũng chứa các thành phần quan trọng cho lập trình web, bao gồm mật mã học và hỗ trợ cho các tiêu chuẩn như HTTP, JSON và XML.

Không có tiêu chí rõ ràng nào định nghĩa những gì được đưa vào vì trong thời gian dài, đây là thư viện Go *duy nhất*. Tuy nhiên, hiện nay có các tiêu chí xác định những gì được thêm vào.

Các bổ sung mới vào thư viện chuẩn rất hiếm và rào cản để được đưa vào là cao. Mã được đưa vào thư viện chuẩn chịu chi phí bảo trì liên tục lớn (thường do những người khác ngoài tác giả gốc chịu), tuân theo [cam kết tương thích Go 1](/doc/go1compat.html) (chặn các bản sửa lỗi cho bất kỳ sai sót nào trong API), và tuân theo [lịch phát hành](/s/releasesched) của Go, ngăn các bản sửa lỗi được cung cấp cho người dùng nhanh chóng.

Hầu hết mã mới nên tồn tại ngoài thư viện chuẩn và có thể truy cập qua lệnh `go get` của [`go` tool](/cmd/go/). Mã đó có thể có người bảo trì riêng, chu kỳ phát hành và đảm bảo tương thích riêng. Người dùng có thể tìm các gói và đọc tài liệu của chúng tại [pkg.go.dev](https://pkg.go.dev/).

Mặc dù có một số mảnh trong thư viện chuẩn không thực sự thuộc về đó, chẳng hạn như `log/syslog`, chúng ta vẫn tiếp tục duy trì mọi thứ trong thư viện vì cam kết tương thích Go 1. Nhưng chúng ta khuyến khích hầu hết mã mới tồn tại ở nơi khác.

## Cài đặt {#Implementation}

### Công nghệ trình biên dịch nào được dùng để xây dựng trình biên dịch? {#What_compiler_technology_is_used_to_build_the_compilers}

Có một số trình biên dịch production cho Go, và một số trình biên dịch khác đang được phát triển cho các nền tảng khác nhau.

Trình biên dịch mặc định, `gc`, được bao gồm trong bản phân phối Go như một phần hỗ trợ cho lệnh `go`. `Gc` ban đầu được viết bằng C vì những khó khăn của bootstrapping -- bạn cần một trình biên dịch Go để thiết lập môi trường Go. Nhưng mọi thứ đã tiến bộ và kể từ bản phát hành Go 1.5, trình biên dịch đã là một chương trình Go. Trình biên dịch đã được chuyển đổi từ C sang Go bằng các công cụ dịch tự động, như được mô tả trong [tài liệu thiết kế](/s/go13compiler) và [bài nói chuyện](/talks/2015/gogo.slide#1) này. Vì vậy, trình biên dịch hiện nay là "self-hosting", nghĩa là chúng ta cần đối mặt với bài toán bootstrapping. Giải pháp là có một bản cài đặt Go đang hoạt động sẵn, giống như người ta thường có với bản cài đặt C đang hoạt động. Câu chuyện về cách khởi động một môi trường Go mới từ nguồn được mô tả [ở đây](/s/go15bootstrap) và [ở đây](/doc/install/source).

`Gc` được viết bằng Go với bộ phân tích cú pháp đệ quy xuống và dùng một trình tải tùy chỉnh, cũng được viết bằng Go nhưng dựa trên trình tải Plan 9, để tạo các binary ELF/Mach-O/PE.

Trình biên dịch `Gccgo` là một front end được viết bằng C++ với bộ phân tích cú pháp đệ quy xuống kết hợp với back end GCC chuẩn. Một [back end LLVM](https://go.googlesource.com/gollvm/) thử nghiệm đang dùng cùng front end.

Vào đầu dự án, chúng ta đã xem xét việc dùng LLVM cho `gc` nhưng quyết định rằng nó quá lớn và chậm để đáp ứng các mục tiêu hiệu suất của chúng ta. Quan trọng hơn khi nhìn lại, bắt đầu với LLVM sẽ khiến việc giới thiệu một số thay đổi ABI và liên quan, chẳng hạn như quản lý stack mà Go yêu cầu nhưng không phải là một phần của thiết lập C chuẩn, trở nên khó khăn hơn.

Go hóa ra là ngôn ngữ tốt để cài đặt trình biên dịch Go, mặc dù đó không phải là mục tiêu ban đầu của nó. Không là self-hosting từ đầu cho phép thiết kế của Go tập trung vào trường hợp sử dụng ban đầu của nó, là các server có kết nối mạng. Nếu chúng ta quyết định Go nên biên dịch chính nó sớm, chúng ta có thể đã kết thúc với một ngôn ngữ nhắm mục tiêu nhiều hơn cho việc xây dựng trình biên dịch, đó là mục tiêu đáng khen nhưng không phải là mục tiêu chúng ta ban đầu có.

Mặc dù `gc` có cài đặt riêng, một lexer và parser native có sẵn trong gói [`go/parser`](/pkg/go/parser/) và cũng có một [type checker](/pkg/go/types) native. Trình biên dịch `gc` dùng các biến thể của những thư viện này.

### Hỗ trợ run-time được cài đặt thế nào? {#How_is_the_run_time_support_implemented}

Cũng do các vấn đề bootstrapping, mã run-time ban đầu được viết chủ yếu bằng C (với một chút assembler) nhưng từ đó đã được dịch sang Go (ngoại trừ một số bit assembler). Hỗ trợ run-time của `Gccgo` dùng `glibc`. Trình biên dịch `gccgo` cài đặt goroutine bằng kỹ thuật gọi là segmented stacks, được hỗ trợ bởi các sửa đổi gần đây cho trình liên kết gold. `Gollvm` tương tự được xây dựng trên cơ sở hạ tầng LLVM tương ứng.

### Tại sao chương trình đơn giản của tôi lại tạo ra binary lớn? {#Why_is_my_trivial_program_such_a_large_binary}

Trình liên kết trong toolchain `gc` tạo các binary liên kết tĩnh theo mặc định. Do đó, tất cả các binary Go bao gồm runtime Go, cùng với thông tin kiểu run-time cần thiết để hỗ trợ kiểm tra kiểu động, reflection, và thậm chí cả stack trace tại thời điểm panic.

Một chương trình C "hello, world" đơn giản được biên dịch và liên kết tĩnh bằng gcc trên Linux vào khoảng 750 kB, bao gồm một cài đặt của `printf`. Một chương trình Go tương đương dùng `fmt.Printf` nặng vài megabyte, nhưng bao gồm hỗ trợ run-time mạnh mẽ hơn và thông tin kiểu và debug.

Một chương trình Go biên dịch với `gc` có thể được liên kết với cờ `-ldflags=-w` để tắt tạo DWARF, loại bỏ thông tin debug khỏi binary mà không mất chức năng nào khác. Điều này có thể giảm đáng kể kích thước binary.

### Tôi có thể tắt các cảnh báo về biến/import không được sử dụng không? {#unused_variables_and_imports}

Sự có mặt của biến không được dùng có thể chỉ ra lỗi, trong khi import không được dùng chỉ làm chậm quá trình biên dịch, hiệu ứng này có thể trở nên đáng kể khi chương trình tích lũy mã và lập trình viên theo thời gian. Vì những lý do này, Go từ chối biên dịch các chương trình có biến hoặc import không được dùng, đánh đổi sự tiện lợi ngắn hạn để có tốc độ build lâu dài và sự rõ ràng của chương trình.

Tuy nhiên, khi phát triển mã, thường xuyên tạo ra những tình huống này tạm thời và có thể khó chịu khi phải chỉnh sửa chúng trước khi chương trình có thể biên dịch.

Một số người đã yêu cầu tùy chọn trình biên dịch để tắt các kiểm tra đó hoặc ít nhất giảm chúng thành cảnh báo. Tùy chọn đó không được thêm vào, vì các tùy chọn trình biên dịch không nên ảnh hưởng đến ngữ nghĩa của ngôn ngữ và vì trình biên dịch Go không báo cảnh báo, chỉ báo lỗi ngăn biên dịch.

Có hai lý do để không có cảnh báo. Thứ nhất, nếu đáng phàn nàn về nó thì đáng sửa trong mã. (Ngược lại, nếu không đáng sửa, thì không đáng đề cập.) Thứ hai, việc trình biên dịch tạo cảnh báo khuyến khích cài đặt cảnh báo về các trường hợp yếu có thể làm quá trình biên dịch ồn ào, che khuất các lỗi thực sự *nên* được sửa.

Tuy nhiên, dễ dàng giải quyết tình huống này. Dùng định danh trống để cho phép những thứ không được dùng tồn tại trong khi bạn đang phát triển.

```
import "unused"

// This declaration marks the import as used by referencing an
// item from the package.
var _ = unused.Item  // TODO: Delete before committing!

func main() {
    debugData := debug.Profile()
    _ = debugData // Used only during debugging.
    ....
}
```

Ngày nay, hầu hết lập trình viên Go dùng một công cụ, [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports), tự động viết lại tệp nguồn Go để có các import đúng, loại bỏ vấn đề import không được dùng trong thực tế. Chương trình này dễ dàng kết nối với hầu hết các trình soạn thảo và IDE để chạy tự động khi tệp nguồn Go được ghi. Chức năng này cũng được tích hợp vào `gopls`, như [đã thảo luận ở trên](/doc/faq#ide).

### Tại sao phần mềm quét virus của tôi cho rằng bản phân phối Go hoặc binary được biên dịch của tôi bị nhiễm? {#virus}

Đây là trường hợp phổ biến, đặc biệt trên máy Windows, và hầu như luôn là dương tính giả. Các chương trình quét virus thương mại thường bị nhầm lẫn bởi cấu trúc của các binary Go, mà chúng không thấy thường xuyên như những binary được biên dịch từ các ngôn ngữ khác.

Nếu bạn vừa cài đặt bản phân phối Go và hệ thống báo cáo nó bị nhiễm, đó chắc chắn là lỗi. Để thực sự kỹ lưỡng, bạn có thể xác minh bản tải xuống bằng cách so sánh checksum với những checksum trên [trang tải xuống](/dl/).

Trong mọi trường hợp, nếu bạn tin rằng báo cáo là sai, hãy báo cáo lỗi cho nhà cung cấp phần mềm quét virus của bạn. Có thể theo thời gian, phần mềm quét virus có thể học cách hiểu các chương trình Go.

## Hiệu suất {#Performance}

### Tại sao Go hoạt động kém trên benchmark X? {#Why_does_Go_perform_badly_on_benchmark_x}

Một trong các mục tiêu thiết kế của Go là tiếp cận hiệu suất của C cho các chương trình tương đương, nhưng trên một số benchmark nó hoạt động khá kém, bao gồm một số trong [golang.org/x/exp/shootout](https://go.googlesource.com/exp/+/master/shootout/). Những trường hợp chậm nhất phụ thuộc vào các thư viện mà các phiên bản có hiệu suất tương đương không có sẵn trong Go. Ví dụ, [pidigits.go](https://go.googlesource.com/exp/+/master/shootout/pidigits.go) phụ thuộc vào gói toán học đa độ chính xác, và các phiên bản C, không giống Go, dùng [GMP](https://gmplib.org/) (được viết bằng assembler được tối ưu hóa). Các benchmark phụ thuộc vào biểu thức chính quy (chẳng hạn [regex-dna.go](https://go.googlesource.com/exp/+/master/shootout/regex-dna.go)) về cơ bản là so sánh [gói regexp](/pkg/regexp) native của Go với các thư viện biểu thức chính quy trưởng thành, được tối ưu hóa cao như PCRE.

Các trò chơi benchmark được thắng bằng cách điều chỉnh rộng rãi và các phiên bản Go của hầu hết các benchmark cần được chú ý. Nếu bạn đo lường các chương trình C và Go thực sự tương đương ([reverse-complement.go](https://go.googlesource.com/exp/+/master/shootout/reverse-complement.go) là một ví dụ), bạn sẽ thấy hai ngôn ngữ gần nhau hơn nhiều về hiệu suất thô so với bộ benchmark này cho thấy.

Tuy vậy, vẫn còn chỗ để cải thiện. Các trình biên dịch tốt nhưng có thể tốt hơn, nhiều thư viện cần cải thiện hiệu suất lớn, và bộ gom rác chưa đủ nhanh. (Ngay cả khi nó đủ nhanh, việc cẩn thận không tạo ra rác không cần thiết vẫn có thể có tác động lớn.)

Trong mọi trường hợp, Go thường có thể rất cạnh tranh. Đã có cải thiện đáng kể về hiệu suất của nhiều chương trình khi ngôn ngữ và công cụ phát triển. Xem bài đăng blog về [profiling các chương trình Go](/blog/profiling-go-programs) để có ví dụ thông tin. Nó khá cũ nhưng vẫn chứa thông tin hữu ích.

## Thay đổi từ C {#change_from_c}

### Tại sao cú pháp lại khác C đến vậy? {#different_syntax}

Ngoài cú pháp khai báo, các khác biệt không lớn và xuất phát từ hai mong muốn. Thứ nhất, cú pháp phải cảm thấy nhẹ nhàng, không có quá nhiều từ khóa bắt buộc, lặp lại hay điều khó hiểu. Thứ hai, ngôn ngữ đã được thiết kế để dễ phân tích và có thể được phân tích cú pháp mà không cần bảng ký hiệu. Điều này làm cho việc xây dựng các công cụ như debugger, bộ phân tích dependency, bộ trích xuất tài liệu tự động, plugin IDE, v.v. dễ dàng hơn nhiều. C và các ngôn ngữ kế thừa của nó nổi tiếng khó khăn trong vấn đề này.

### Tại sao khai báo lại ngược chiều? {#declarations_backwards}

Chúng chỉ ngược chiều nếu bạn đã quen với C. Trong C, khái niệm là một biến được khai báo giống như một biểu thức biểu thị kiểu của nó, đó là ý tưởng hay, nhưng ngữ pháp kiểu và biểu thức không kết hợp tốt với nhau và kết quả có thể gây nhầm lẫn; hãy xem xét con trỏ hàm. Go chủ yếu tách biệt cú pháp biểu thức và kiểu và điều đó đơn giản hóa mọi thứ (dùng tiền tố `*` cho con trỏ là ngoại lệ chứng minh quy tắc). Trong C, khai báo:

```
    int* a, b;
```

khai báo `a` là con trỏ nhưng không phải `b`; trong Go:

```
    var a, b *int
```

khai báo cả hai đều là con trỏ. Điều này rõ ràng và đều đặn hơn. Ngoài ra, dạng khai báo ngắn `:=` lập luận rằng khai báo biến đầy đủ nên trình bày cùng thứ tự như `:=` vì vậy:

```
    var a uint64 = 1
```

có hiệu ứng tương tự như:

```
    a := uint64(1)
```

Phân tích cú pháp cũng được đơn giản hóa bằng cách có ngữ pháp riêng biệt cho kiểu không chỉ là ngữ pháp biểu thức; các từ khóa như `func` và `chan` giữ mọi thứ rõ ràng.

Xem bài viết về [Cú pháp khai báo của Go](/doc/articles/gos_declaration_syntax.html) để biết thêm chi tiết.

### Tại sao không có phép toán con trỏ? {#no_pointer_arithmetic}

An toàn. Không có phép toán con trỏ, có thể tạo ra một ngôn ngữ không bao giờ có thể dẫn xuất địa chỉ bất hợp pháp thành công không đúng cách. Công nghệ trình biên dịch và phần cứng đã tiến bộ đến mức một vòng lặp dùng chỉ số mảng có thể hiệu quả như vòng lặp dùng phép toán con trỏ. Ngoài ra, việc thiếu phép toán con trỏ có thể đơn giản hóa việc cài đặt bộ gom rác.

### Tại sao `++` và `--` là câu lệnh chứ không phải biểu thức? Và tại sao là hậu tố, không phải tiền tố? {#inc_dec}

Không có phép toán con trỏ, giá trị tiện lợi của các toán tử tăng/giảm tiền tố và hậu tố giảm đi. Bằng cách loại bỏ chúng hoàn toàn khỏi phân cấp biểu thức, cú pháp biểu thức được đơn giản hóa và các vấn đề lộn xộn xung quanh thứ tự đánh giá của `++` và `--` (xét `f(i++)` và `p[i] = q[++i]`) cũng được loại bỏ. Sự đơn giản hóa là đáng kể. Còn về hậu tố so với tiền tố, cả hai đều hoạt động tốt nhưng phiên bản hậu tố truyền thống hơn; việc nhấn mạnh vào tiền tố xuất phát từ STL, một thư viện cho ngôn ngữ có tên chứa, trớ trêu thay, một phép tăng hậu tố.

### Tại sao có dấu ngoặc nhọn mà không có dấu chấm phẩy? Và tại sao tôi không thể đặt dấu ngoặc nhọn mở đầu trên dòng tiếp theo? {#semicolons}

Go dùng dấu ngoặc nhọn để nhóm câu lệnh, một cú pháp quen thuộc với các lập trình viên đã làm việc với bất kỳ ngôn ngữ nào trong họ C. Tuy nhiên, dấu chấm phẩy là cho trình phân tích cú pháp, không phải cho con người, và chúng ta muốn loại bỏ chúng càng nhiều càng tốt. Để đạt mục tiêu này, Go mượn một kỹ thuật từ BCPL: các dấu chấm phẩy phân tách câu lệnh có trong ngữ pháp hình thức nhưng được tự động chèn vào, không cần nhìn trước, bởi lexer ở cuối bất kỳ dòng nào có thể là kết thúc của một câu lệnh. Điều này hoạt động rất tốt trong thực tế nhưng có tác dụng là buộc một phong cách dấu ngoặc nhọn. Ví dụ, dấu ngoặc nhọn mở đầu của một hàm không thể xuất hiện trên dòng riêng của nó.

Một số người đã lập luận rằng lexer nên nhìn trước để cho phép dấu ngoặc nhọn sống trên dòng tiếp theo. Chúng ta không đồng ý. Vì mã Go có nghĩa là được định dạng tự động bởi [`gofmt`](/cmd/gofmt/), *một số* phong cách phải được chọn. Phong cách đó có thể khác với những gì bạn đã dùng trong C hay Java, nhưng Go là một ngôn ngữ khác và phong cách của `gofmt` cũng tốt như bất kỳ phong cách nào khác. Quan trọng hơn -- quan trọng hơn nhiều -- lợi thế của một định dạng duy nhất, được quy định theo chương trình cho tất cả các chương trình Go, vượt xa bất kỳ bất lợi được nhận thức nào của phong cách cụ thể. Cũng lưu ý rằng phong cách của Go có nghĩa là một cài đặt tương tác của Go có thể dùng cú pháp chuẩn từng dòng một mà không cần quy tắc đặc biệt.

### Tại sao dùng bộ gom rác? Nó có không quá tốn kém không? {#garbage_collection}

Một trong những nguồn quản lý sổ sách lớn nhất trong các chương trình hệ thống là quản lý vòng đời của các đối tượng được cấp phát. Trong các ngôn ngữ như C, nơi nó được thực hiện thủ công, có thể tiêu tốn đáng kể thời gian của lập trình viên và thường là nguyên nhân của các lỗi nguy hiểm. Ngay cả trong các ngôn ngữ như C++ hay Rust cung cấp các cơ chế hỗ trợ, những cơ chế đó có thể có ảnh hưởng đáng kể đến thiết kế của phần mềm, thường thêm overhead lập trình riêng của nó. Chúng ta cảm thấy quan trọng phải loại bỏ những overhead lập trình như vậy, và những tiến bộ trong công nghệ bộ gom rác trong những năm gần đây đã cho chúng ta sự tự tin rằng nó có thể được cài đặt đủ rẻ, với độ trễ đủ thấp, để có thể là cách tiếp cận khả thi cho các hệ thống mạng.

Phần lớn khó khăn của lập trình đồng thời có nguồn gốc từ vấn đề vòng đời đối tượng: khi các đối tượng được truyền giữa các luồng, việc đảm bảo chúng được giải phóng an toàn trở nên cồng kềnh. Thu gom rác tự động làm cho mã đồng thời dễ viết hơn nhiều. Tất nhiên, việc cài đặt bộ gom rác trong môi trường đồng thời bản thân nó là một thách thức, nhưng giải quyết nó một lần thay vì trong mọi chương trình giúp ích cho mọi người.

Cuối cùng, ngoài đồng thời, bộ gom rác làm cho các interface đơn giản hơn vì chúng không cần chỉ định cách bộ nhớ được quản lý qua chúng.

Điều này không có nghĩa là công việc gần đây trong các ngôn ngữ như Rust đưa ra các ý tưởng mới cho vấn đề quản lý tài nguyên là sai lầm; chúng ta khuyến khích công việc này và hào hứng chờ xem nó phát triển thế nào. Nhưng Go theo cách tiếp cận truyền thống hơn bằng cách giải quyết vòng đời đối tượng thông qua bộ gom rác, và chỉ bộ gom rác mà thôi.

Cài đặt hiện tại là bộ gom rác mark-and-sweep. Nếu máy là bộ đa xử lý, bộ gom rác chạy trên lõi CPU riêng song song với chương trình chính. Công việc lớn trên bộ gom rác trong những năm gần đây đã giảm thời gian tạm dừng thường xuống phạm vi dưới mili giây, ngay cả với các heap lớn, gần như loại bỏ một trong những phản đối lớn về bộ gom rác trong các server mạng. Công việc tiếp tục để tinh chỉnh thuật toán, giảm overhead và độ trễ hơn nữa, và khám phá các cách tiếp cận mới. [Bài phát biểu chính ISMM 2018](/blog/ismmkeynote) của Rick Hudson từ nhóm Go mô tả tiến độ cho đến nay và gợi ý một số cách tiếp cận trong tương lai.

Về chủ đề hiệu suất, hãy nhớ rằng Go cung cấp cho lập trình viên quyền kiểm soát đáng kể đối với bố cục bộ nhớ và cấp phát, nhiều hơn so với điển hình trong các ngôn ngữ được gom rác. Một lập trình viên cẩn thận có thể giảm đáng kể overhead bộ gom rác bằng cách dùng ngôn ngữ tốt; xem bài viết về [profiling các chương trình Go](/blog/profiling-go-programs) để có ví dụ thực tiễn, bao gồm minh họa về các công cụ profiling của Go.
