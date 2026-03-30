---
title: Tại sao cần Generics?
date: 2019-07-31
by:
- Ian Lance Taylor
tags:
- go2
- proposals
- generics
summary: Tại sao chúng ta nên thêm generics vào Go, và chúng có thể trông như thế nào?
template: true
---

## Giới thiệu

Đây là phiên bản bài đăng blog của bài nói chuyện của tôi tuần trước tại Gophercon 2019.

{{video "https://www.youtube.com/embed/WzgLqE-3IhY?rel=0"}}

Bài viết này bàn về ý nghĩa của việc thêm generics vào Go, và
lý do tôi nghĩ chúng ta nên làm điều đó.
Tôi cũng sẽ đề cập đến một bản cập nhật về thiết kế khả thi cho
việc thêm generics vào Go.

Go được phát hành vào ngày 10 tháng 11 năm 2009.
Chưa đến 24 giờ sau, chúng tôi đã thấy
[bình luận đầu tiên về generics](https://groups.google.com/d/msg/golang-nuts/70-pdwUUrbI/onMsQspcljcJ).
(Bình luận đó cũng đề cập đến ngoại lệ (exceptions), thứ chúng tôi đã thêm vào
ngôn ngữ, dưới dạng `panic` và `recover`, vào đầu năm 2010.)

Trong ba năm khảo sát Go, sự thiếu vắng generics luôn được liệt kê
là một trong ba vấn đề hàng đầu cần khắc phục trong ngôn ngữ.

## Tại sao cần generics?

Nhưng thêm generics có nghĩa là gì, và tại sao chúng ta muốn điều đó?

Để diễn giải lại
[Jazayeri và cộng sự](https://www.dagstuhl.de/en/program/calendar/semhp/?semnr=98171):
lập trình generic cho phép biểu diễn hàm và cấu trúc dữ liệu
ở dạng tổng quát, với kiểu dữ liệu được tách ra ngoài.

Điều đó có nghĩa là gì?

Xét một ví dụ đơn giản, hãy giả sử chúng ta muốn đảo ngược các phần tử trong
một slice.
Đây không phải là điều nhiều chương trình cần làm, nhưng cũng không quá hiếm.

Giả sử đó là một slice of int.

{{raw `
	func ReverseInts(s []int) {
		first := 0
		last := len(s)
		for first < last {
			s[first], s[last] = s[last], s[first]
			first++
			last--
		}
	}
`}}

Khá đơn giản, nhưng ngay cả với một hàm đơn giản như vậy, bạn cũng muốn
viết một vài test case.
Thực ra, khi tôi làm vậy, tôi tìm thấy một lỗi.
Tôi chắc nhiều độc giả đã nhận ra nó.

{{raw `
	func ReverseInts(s []int) {
		first := 0
		last := len(s) - 1
		for first < last {
			s[first], s[last] = s[last], s[first]
			first++
			last--
		}
	}
`}}

Chúng ta cần trừ 1 khi thiết lập biến last.

Bây giờ hãy đảo ngược một slice of string.

{{raw `
	func ReverseStrings(s []string) {
		first := 0
		last := len(s) - 1
		for first < last {
			s[first], s[last] = s[last], s[first]
			first++
			last--
		}
	}
`}}

Nếu bạn so sánh `ReverseInts` và `ReverseStrings`, bạn sẽ thấy rằng
hai hàm hoàn toàn giống nhau, ngoại trừ kiểu của tham số.
Tôi không nghĩ độc giả nào ngạc nhiên về điều đó.

Điều mà một số người mới học Go thấy ngạc nhiên là không có cách nào để
viết một hàm `Reverse` đơn giản hoạt động cho một slice của bất kỳ kiểu nào.

Hầu hết các ngôn ngữ khác cho phép bạn viết loại hàm đó.

Trong một ngôn ngữ kiểu động như Python hay JavaScript, bạn có thể
đơn giản viết hàm đó mà không cần chỉ định kiểu phần tử. Điều này không hoạt động trong Go vì Go là ngôn ngữ kiểu tĩnh, và
yêu cầu bạn viết rõ kiểu của slice và kiểu của
các phần tử slice.

Hầu hết các ngôn ngữ kiểu tĩnh khác, như C++, Java, Rust hay
Swift, đều hỗ trợ generics để giải quyết chính xác loại vấn đề này.

## Lập trình generic trong Go ngày nay

Vậy mọi người viết loại mã này trong Go như thế nào?

Trong Go bạn có thể viết một hàm duy nhất hoạt động cho các kiểu slice khác nhau
bằng cách sử dụng kiểu interface, và định nghĩa một phương thức trên các kiểu slice
bạn muốn truyền vào.
Đó là cách hàm `sort.Sort` trong thư viện chuẩn hoạt động.

Nói cách khác, kiểu interface trong Go là một dạng lập trình generic.
Chúng cho phép chúng ta nắm bắt các khía cạnh chung của các kiểu khác nhau và biểu diễn
chúng dưới dạng phương thức.
Sau đó chúng ta có thể viết các hàm sử dụng các kiểu interface đó, và những
hàm đó sẽ hoạt động với bất kỳ kiểu nào triển khai các phương thức đó.

Nhưng cách tiếp cận này vẫn thiếu so với những gì chúng ta muốn.
Với interface, bạn phải tự viết các phương thức.
Thật phiền khi phải định nghĩa một kiểu được đặt tên với một vài phương thức
chỉ để đảo ngược một slice.
Và các phương thức bạn viết hoàn toàn giống nhau cho mỗi kiểu slice, vì vậy
về một nghĩa nào đó chúng ta chỉ di chuyển và cô đọng mã trùng lặp, chứ không loại bỏ nó.
Mặc dù interface là một dạng generics, chúng không cung cấp cho chúng ta
mọi thứ chúng ta muốn từ generics.

Một cách khác để sử dụng interface cho generics, có thể tránh được
việc phải tự viết các phương thức, là để ngôn ngữ
định nghĩa phương thức cho một số loại kiểu.
Đây không phải là điều ngôn ngữ hỗ trợ ngày nay, nhưng, ví dụ,
ngôn ngữ có thể định nghĩa rằng mọi kiểu slice đều có phương thức Index
trả về một phần tử.
Nhưng để sử dụng phương thức đó trong thực tế, nó sẽ phải trả về
kiểu interface rỗng, và sau đó chúng ta mất tất cả lợi ích của kiểu tĩnh.
Tế nhị hơn, sẽ không có cách nào để định nghĩa một hàm generic nhận
hai slice khác nhau với cùng kiểu phần tử, hoặc nhận một map của
một kiểu phần tử và trả về một slice của cùng kiểu phần tử đó.
Go là ngôn ngữ kiểu tĩnh vì điều đó giúp dễ dàng hơn
viết các chương trình lớn; chúng ta không muốn mất lợi ích của kiểu tĩnh
để đổi lấy lợi ích của generics.

Một cách tiếp cận khác là viết hàm `Reverse` generic bằng cách sử dụng
gói reflect, nhưng điều đó quá phiền khi viết và chậm khi chạy
đến mức ít người làm vậy.
Cách tiếp cận đó cũng yêu cầu kiểm tra kiểu tường minh và không có kiểm tra kiểu tĩnh.

Hoặc, bạn có thể viết một trình tạo mã nhận một kiểu và tạo ra một
hàm `Reverse` cho slice của kiểu đó.
Có một số trình tạo mã làm chính xác điều đó.
Nhưng điều này thêm một bước vào mọi gói cần `Reverse`,
nó làm phức tạp quá trình build vì tất cả các bản sao khác nhau phải được
biên dịch, và việc sửa lỗi trong mã nguồn chính đòi hỏi phải tạo lại
tất cả các instance, một số trong số đó có thể thuộc các dự án hoàn toàn khác.

Tất cả các cách tiếp cận này đều đủ phiền phức để tôi nghĩ hầu hết
những người phải đảo ngược một slice trong Go chỉ viết hàm cho
kiểu slice cụ thể mà họ cần.
Sau đó họ sẽ cần viết test case cho hàm, để đảm bảo
họ không mắc lỗi đơn giản như tôi đã mắc ban đầu.
Và họ sẽ cần chạy các bài test đó thường xuyên.

Dù chúng ta làm thế nào, điều đó có nghĩa là rất nhiều công việc thêm chỉ cho một hàm
trông hoàn toàn giống nhau ngoại trừ kiểu phần tử.
Không phải là không thể làm được.
Rõ ràng là có thể làm được, và các lập trình viên Go đang làm điều đó.
Chỉ là nên có một cách tốt hơn.

Đối với ngôn ngữ kiểu tĩnh như Go, cách tốt hơn đó là generics.
Như tôi đã viết trước đó, lập trình generic cho phép
biểu diễn hàm và cấu trúc dữ liệu ở dạng tổng quát,
với kiểu dữ liệu được tách ra ngoài.
Đó chính xác là những gì chúng ta muốn ở đây.

## Những gì generics có thể mang lại cho Go

Điều đầu tiên và quan trọng nhất chúng ta muốn từ generics trong Go là
có thể viết các hàm như `Reverse` mà không quan tâm đến
kiểu phần tử của slice.
Chúng ta muốn tách kiểu phần tử đó ra.
Sau đó chúng ta có thể viết hàm một lần, viết các bài test một lần, đặt chúng trong
một gói có thể dùng go get, và gọi chúng bất cứ khi nào chúng ta muốn.

Còn tốt hơn, vì đây là thế giới mã nguồn mở, người khác có thể
viết `Reverse` một lần, và chúng ta có thể sử dụng triển khai của họ.

Tại đây tôi nên nói rằng "generics" có thể có nhiều ý nghĩa khác nhau.
Trong bài viết này, ý tôi về "generics" là những gì tôi vừa mô tả.
Cụ thể, tôi không có ý nói đến template như trong ngôn ngữ C++,
cái hỗ trợ khá nhiều hơn những gì tôi đã viết ở đây.

Tôi đã trình bày chi tiết về `Reverse`, nhưng có nhiều hàm khác
mà chúng ta có thể viết theo cách generic, chẳng hạn như:

  - Tìm phần tử nhỏ nhất/lớn nhất trong slice
  - Tính trung bình/độ lệch chuẩn của slice
  - Tính hợp/giao của các map
  - Tìm đường đi ngắn nhất trong đồ thị nút/cạnh
  - Áp dụng hàm biến đổi vào slice/map, trả về slice/map mới

Những ví dụ này có sẵn trong hầu hết các ngôn ngữ khác.
Thực ra, tôi đã viết danh sách này bằng cách lướt qua thư viện template chuẩn của C++.

Cũng có những ví dụ đặc thù cho Go với sự hỗ trợ mạnh mẽ về concurrency.

  - Đọc từ channel với timeout
  - Kết hợp hai channel thành một channel duy nhất
  - Gọi một danh sách hàm song song, trả về một slice kết quả
  - Gọi một danh sách hàm, sử dụng Context, trả về kết quả của hàm đầu tiên hoàn thành, hủy và dọn dẹp các goroutine dư thừa

Tôi đã thấy tất cả những hàm này được viết ra nhiều lần với các kiểu khác nhau.
Không khó để viết chúng trong Go.
Nhưng sẽ tốt hơn nếu có thể tái sử dụng một triển khai hiệu quả và đã được kiểm tra lỗi
hoạt động với bất kỳ kiểu giá trị nào.

Để rõ ràng hơn, đây chỉ là các ví dụ.
Có nhiều hàm mục đích chung hơn có thể được viết
dễ dàng và an toàn hơn khi sử dụng generics.

Ngoài ra, như tôi đã viết trước đó, không chỉ là hàm.
Mà còn là cấu trúc dữ liệu.

Go có hai cấu trúc dữ liệu generic mục đích chung được tích hợp vào
ngôn ngữ: slice và map.
Slice và map có thể chứa các giá trị của bất kỳ kiểu dữ liệu nào, với kiểm tra kiểu tĩnh
cho các giá trị được lưu trữ và truy xuất.
Các giá trị được lưu trữ dưới dạng chính chúng, không phải dưới dạng kiểu interface.
Tức là, khi tôi có `[]int`, slice chứa trực tiếp các int, không phải
các int được chuyển đổi sang kiểu interface.

Slice và map là các cấu trúc dữ liệu generic hữu ích nhất, nhưng chúng
không phải là loại duy nhất.
Đây là một số ví dụ khác.

  - Set
  - Cây tự cân bằng, với khả năng chèn và duyệt hiệu quả theo thứ tự đã sắp xếp
  - Multimap, với nhiều instance của một khóa
  - Map băm đồng thời, hỗ trợ chèn và tra cứu song song mà không cần khóa đơn

Nếu chúng ta có thể viết kiểu generic, chúng ta có thể định nghĩa các cấu trúc dữ liệu mới, như
những cái này, có cùng ưu điểm kiểm tra kiểu như slice và map:
trình biên dịch có thể kiểm tra kiểu tĩnh các kiểu giá trị mà
chúng chứa, và các giá trị có thể được lưu trữ dưới dạng chính chúng, không phải dưới dạng
kiểu interface.

Cũng nên có khả năng lấy các thuật toán như những cái đã đề cập
trước đó và áp dụng chúng cho các cấu trúc dữ liệu generic.

Những ví dụ này đều nên giống như `Reverse`: hàm generic
và cấu trúc dữ liệu được viết một lần, trong một gói, và tái sử dụng bất cứ khi nào
chúng cần.
Chúng nên hoạt động như slice và map, ở chỗ chúng không nên lưu trữ
các giá trị kiểu interface rỗng, mà nên lưu trữ các kiểu cụ thể, và
các kiểu đó nên được kiểm tra tại thời điểm biên dịch.

Vậy đó là những gì Go có thể đạt được từ generics.
Generics có thể cung cấp cho chúng ta các khối xây dựng mạnh mẽ giúp chúng ta chia sẻ mã
và xây dựng chương trình dễ dàng hơn.

Tôi hy vọng tôi đã giải thích lý do tại sao điều này đáng để khám phá.

## Lợi ích và chi phí

Nhưng generics không đến từ
[Vùng đất kẹo ngọt](https://mainlynorfolk.info/folk/songs/bigrockcandymountain.html),
vùng đất nơi mặt trời chiếu sáng mỗi ngày trên những
[suối nước chanh](http://www.lat-long.com/Latitude-Longitude-773297-Montana-Lemonade_Springs.html).
Mọi thay đổi ngôn ngữ đều có chi phí.
Không còn nghi ngờ gì nữa rằng việc thêm generics vào Go sẽ làm cho ngôn ngữ
phức tạp hơn.
Như với bất kỳ thay đổi nào đối với ngôn ngữ, chúng ta cần nói về việc tối đa hóa
lợi ích và giảm thiểu chi phí.

Trong Go, chúng tôi đã nhắm đến việc giảm độ phức tạp thông qua các tính năng ngôn ngữ
độc lập, trực giao có thể kết hợp tự do.
Chúng tôi giảm độ phức tạp bằng cách làm cho các tính năng riêng lẻ đơn giản, và chúng tôi
tối đa hóa lợi ích của các tính năng bằng cách cho phép kết hợp tự do chúng.
Chúng tôi muốn làm điều tương tự với generics.

Để làm rõ hơn, tôi sẽ liệt kê một vài hướng dẫn mà
chúng ta nên tuân theo.

### Giảm thiểu khái niệm mới

Chúng ta nên thêm càng ít khái niệm mới vào ngôn ngữ càng tốt.
Điều đó có nghĩa là tối thiểu cú pháp mới và tối thiểu từ khóa mới
và các tên khác.

### Độ phức tạp rơi vào người viết mã generic, không phải người dùng

Càng nhiều càng tốt, độ phức tạp nên rơi vào lập trình viên
đang viết gói generic.
Chúng ta không muốn người dùng gói phải lo lắng về generics.
Điều này có nghĩa là nên có thể gọi hàm generic theo cách
tự nhiên, và bất kỳ lỗi nào khi sử dụng gói generic nên được báo cáo theo cách
dễ hiểu và dễ sửa.
Cũng nên dễ dàng gỡ lỗi các lời gọi vào mã generic.

### Người viết và người dùng có thể làm việc độc lập

Tương tự, chúng ta nên dễ dàng tách biệt các mối quan tâm của
người viết mã generic và người dùng, để họ có thể phát triển
mã của mình một cách độc lập.
Họ không nên phải lo lắng về những gì người kia đang làm, hơn gì
người viết và người gọi một hàm bình thường trong các gói khác nhau
phải lo lắng.
Điều này nghe có vẻ hiển nhiên, nhưng không đúng với generics trong mọi
ngôn ngữ lập trình khác.

### Thời gian build ngắn, thời gian thực thi nhanh

Đương nhiên, càng nhiều càng tốt, chúng ta muốn giữ thời gian build ngắn
và thời gian thực thi nhanh mà Go cung cấp cho chúng ta ngày nay.
Generics có xu hướng tạo ra sự đánh đổi giữa build nhanh và thực thi nhanh.
Càng nhiều càng tốt, chúng ta muốn cả hai.

### Giữ nguyên sự rõ ràng và đơn giản của Go

Quan trọng nhất, Go ngày nay là một ngôn ngữ đơn giản.
Các chương trình Go thường rõ ràng và dễ hiểu.
Một phần lớn trong quá trình dài khám phá không gian này của chúng tôi là
cố gắng hiểu cách thêm generics trong khi vẫn giữ nguyên sự rõ ràng
và đơn giản đó.
Chúng ta cần tìm các cơ chế phù hợp tốt với ngôn ngữ hiện có,
mà không biến nó thành thứ gì đó hoàn toàn khác.

Những hướng dẫn này nên áp dụng cho bất kỳ triển khai generics nào trong Go.
Đó là thông điệp quan trọng nhất tôi muốn để lại cho bạn hôm nay:
**generics có thể mang lại lợi ích đáng kể cho ngôn ngữ, nhưng chúng chỉ đáng làm nếu Go vẫn cảm thấy như Go**.

## Bản thiết kế nháp

May mắn thay, tôi nghĩ điều này có thể thực hiện được.
Để kết thúc bài viết này, tôi sẽ chuyển từ việc thảo luận tại sao
chúng ta muốn generics, và các yêu cầu về chúng là gì, sang thảo luận ngắn gọn
về một thiết kế về cách chúng ta nghĩ có thể thêm chúng vào ngôn ngữ.

Ghi chú thêm tháng 1 năm 2022: Bài đăng blog này được viết năm 2019 và
không mô tả phiên bản generics cuối cùng đã được chấp nhận.
Để biết thông tin cập nhật, vui lòng xem mô tả về tham số kiểu trong
[đặc tả ngôn ngữ](/ref/spec) và
[tài liệu thiết kế generics](/design/43651-type-parameters).

Tại Gophercon năm nay, Robert Griesemer và tôi đã công bố
[một bản thiết kế nháp](https://github.com/golang/proposal/blob/master/design/go2draft-contracts.md)
cho việc thêm generics vào Go.
Xem bản nháp để biết chi tiết đầy đủ.
Tôi sẽ trình bày một số điểm chính ở đây.

Đây là hàm Reverse generic trong thiết kế này.

{{raw `
	func Reverse (type Element) (s []Element) {
		first := 0
		last := len(s) - 1
		for first < last {
			s[first], s[last] = s[last], s[first]
			first++
			last--
		}
	}
`}}

Bạn sẽ nhận thấy rằng phần thân của hàm hoàn toàn giống nhau.
Chỉ có chữ ký thay đổi.

Kiểu phần tử của slice đã được tách ra.
Nó giờ được đặt tên là `Element` và đã trở thành cái chúng ta gọi là
_tham số kiểu_.
Thay vì là một phần của kiểu tham số slice, nó giờ là một
tham số kiểu riêng biệt, bổ sung.

Để gọi một hàm với tham số kiểu, trong trường hợp tổng quát bạn truyền
một đối số kiểu, cái giống như bất kỳ đối số nào khác ngoại trừ nó là một kiểu.

	func ReverseAndPrint(s []int) {
		Reverse(int)(s)
		fmt.Println(s)
	}

Đó là `(int)` thấy sau `Reverse` trong ví dụ này.

May mắn thay, trong hầu hết các trường hợp, bao gồm trường hợp này, trình biên dịch có thể
suy ra đối số kiểu từ các kiểu của đối số thông thường, và
bạn không cần đề cập đến đối số kiểu chút nào.

Gọi một hàm generic trông giống như gọi bất kỳ hàm nào khác.

	func ReverseAndPrint(s []int) {
		Reverse(s)
		fmt.Println(s)
	}

Nói cách khác, mặc dù hàm generic `Reverse` phức tạp hơn một chút
so với `ReverseInts` và `ReverseStrings`, độ phức tạp đó
rơi vào người viết hàm, không phải người gọi.

### Contract

Vì Go là ngôn ngữ kiểu tĩnh, chúng ta phải nói về kiểu
của một tham số kiểu.
_Kiểu meta_ này cho trình biên dịch biết những loại đối số kiểu nào được
phép khi gọi một hàm generic, và những loại thao tác nào
hàm generic có thể thực hiện với các giá trị của tham số kiểu.

Hàm `Reverse` có thể hoạt động với slice của bất kỳ kiểu nào.
Điều duy nhất nó làm với các giá trị của kiểu `Element` là gán,
cái hoạt động với bất kỳ kiểu nào trong Go.
Đối với loại hàm generic này, cái là một trường hợp rất phổ biến, chúng ta
không cần nói gì đặc biệt về tham số kiểu.

Hãy xem nhanh một hàm khác.

{{raw `
	func IndexByte (type T Sequence) (s T, b byte) int {
		for i := 0; i < len(s); i++ {
			if s[i] == b {
				return i
			}
		}
		return -1
	}
`}}

Hiện tại cả gói bytes và gói strings trong
thư viện chuẩn đều có hàm `IndexByte`.
Hàm này trả về chỉ số của `b` trong chuỗi `s`, nơi `s`
là `string` hoặc `[]byte`.
Chúng ta có thể sử dụng hàm generic đơn lẻ này để thay thế hai hàm
trong các gói bytes và strings.
Trong thực tế, chúng ta có thể không bận tâm làm điều đó, nhưng đây là một ví dụ đơn giản hữu ích.

Ở đây chúng ta cần biết rằng tham số kiểu `T` hoạt động như `string`
hoặc `[]byte`.
Chúng ta có thể gọi `len` trên nó, có thể lập chỉ số vào nó, và có thể so sánh
kết quả của thao tác lập chỉ số với một giá trị byte.

Để biên dịch được điều này, bản thân tham số kiểu `T` cần có một kiểu.
Đó là kiểu meta, nhưng vì đôi khi chúng ta cần mô tả nhiều
kiểu liên quan, và vì nó mô tả mối quan hệ giữa
triển khai hàm generic và người gọi nó, chúng ta thực sự
gọi kiểu của `T` là contract.
Ở đây contract được đặt tên là `Sequence`.
Nó xuất hiện sau danh sách tham số kiểu.

Đây là cách contract Sequence được định nghĩa cho ví dụ này.

	contract Sequence(T) {
		T string, []byte
	}

Khá đơn giản, vì đây là một ví dụ đơn giản: tham số kiểu
`T` có thể là `string` hoặc `[]byte`.
Ở đây `contract` có thể là một từ khóa mới, hoặc một định danh đặc biệt
được nhận ra trong phạm vi gói; xem bản thiết kế nháp để biết chi tiết.

Bất kỳ ai nhớ [thiết kế chúng tôi đã trình bày tại Gophercon 2018](https://github.com/golang/proposal/blob/4a530dae40977758e47b78fae349d8e5f86a6c0a/design/go2draft-contracts.md)
sẽ thấy rằng cách viết contract này đơn giản hơn rất nhiều.
Chúng tôi đã nhận được nhiều phản hồi về thiết kế trước đó rằng contract
quá phức tạp, và chúng tôi đã cố gắng tính đến điều đó.
Các contract mới đơn giản hơn nhiều để viết, để đọc và để hiểu.

Chúng cho phép bạn chỉ định kiểu cơ sở của một tham số kiểu, và/hoặc
liệt kê các phương thức của tham số kiểu.
Chúng cũng cho phép bạn mô tả mối quan hệ giữa các tham số kiểu khác nhau.

### Contract với phương thức

Đây là một ví dụ đơn giản khác, về một hàm sử dụng phương thức String
để trả về `[]string` là biểu diễn chuỗi của tất cả các
phần tử trong `s`.

	func ToStrings (type E Stringer) (s []E) []string {
		r := make([]string, len(s))
		for i, v := range s {
			r[i] = v.String()
		}
		return r
	}

Khá rõ ràng: duyệt qua slice, gọi phương thức `String`
trên mỗi phần tử, và trả về một slice các chuỗi kết quả.

Hàm này yêu cầu kiểu phần tử triển khai phương thức `String`.
Contract Stringer đảm bảo điều đó.

	contract Stringer(T) {
		T String() string
	}

Contract chỉ đơn giản nói rằng `T` phải triển khai phương thức `String`.

Bạn có thể nhận thấy rằng contract này trông giống như interface `fmt.Stringer`,
vì vậy đáng chú ý rằng đối số của hàm
`ToStrings` không phải là một slice của `fmt.Stringer`.
Đó là một slice của một kiểu phần tử nào đó, nơi kiểu phần tử triển khai
`fmt.Stringer`.
Biểu diễn bộ nhớ của một slice kiểu phần tử và một slice
`fmt.Stringer` thường khác nhau, và Go không hỗ trợ
chuyển đổi trực tiếp giữa chúng.
Vì vậy, điều này đáng viết, mặc dù `fmt.Stringer` tồn tại.

### Contract với nhiều kiểu

Đây là ví dụ về contract với nhiều tham số kiểu.

	type Graph (type Node, Edge G) struct { ... }

	contract G(Node, Edge) {
		Node Edges() []Edge
		Edge Nodes() (from Node, to Node)
	}

	func New (type Node, Edge G) (nodes []Node) *Graph(Node, Edge) {
		...
	}

	func (g *Graph(Node, Edge)) ShortestPath(from, to Node) []Edge {
		...
	}

Ở đây chúng ta đang mô tả một đồ thị, được xây dựng từ các nút và cạnh.
Chúng ta không yêu cầu một cấu trúc dữ liệu cụ thể cho đồ thị.
Thay vào đó, chúng ta nói rằng kiểu `Node` phải có phương thức `Edges`
trả về danh sách các cạnh kết nối với `Node`.
Và kiểu `Edge` phải có phương thức `Nodes` trả về hai
`Node` mà `Edge` kết nối.

Tôi đã bỏ qua phần triển khai, nhưng điều này cho thấy chữ ký của một
hàm `New` trả về `Graph`, và chữ ký của phương thức `ShortestPath` trên `Graph`.

Điểm quan trọng cần ghi nhớ ở đây là contract không chỉ về
một kiểu duy nhất. Nó có thể mô tả mối quan hệ giữa hai hoặc nhiều
kiểu.

### Kiểu có thứ tự

Một phàn nàn đáng ngạc nhiên phổ biến về Go là nó không có hàm `Min`.
Hay, tương tự, một hàm `Max`.
Điều đó là vì một hàm `Min` hữu ích nên hoạt động cho bất kỳ kiểu có thứ tự nào,
nghĩa là nó phải là generic.

Trong khi `Min` khá tầm thường để tự viết, bất kỳ triển khai generics hữu ích nào
cũng nên cho phép chúng ta thêm nó vào thư viện chuẩn.
Đây là hình dạng của nó với thiết kế của chúng ta.

{{raw `
	func Min (type T Ordered) (a, b T) T {
		if a < b {
			return a
		}
		return b
	}
`}}

Contract `Ordered` nói rằng kiểu T phải là kiểu có thứ tự,
nghĩa là nó hỗ trợ các toán tử như nhỏ hơn, lớn hơn,
và tương tự.

	contract Ordered(T) {
		T int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr,
			float32, float64,
			string
	}

Contract `Ordered` chỉ là danh sách tất cả các kiểu có thứ tự
được định nghĩa bởi ngôn ngữ.
Contract này chấp nhận bất kỳ kiểu nào trong danh sách, hoặc bất kỳ kiểu được đặt tên nào có
kiểu cơ sở là một trong những kiểu đó.
Về cơ bản, bất kỳ kiểu nào bạn có thể sử dụng với toán tử nhỏ hơn.

Hóa ra việc đơn giản liệt kê các kiểu hỗ trợ toán tử nhỏ hơn
dễ hơn nhiều so với việc phát minh một ký hiệu mới
hoạt động cho tất cả các toán tử.
Xét cho cùng, trong Go, chỉ có các kiểu tích hợp hỗ trợ toán tử.

Cách tiếp cận tương tự có thể được sử dụng cho bất kỳ toán tử nào, hay tổng quát hơn
để viết contract cho bất kỳ hàm generic nào được thiết kế để làm việc với
các kiểu tích hợp.
Nó cho phép người viết hàm generic chỉ định rõ ràng tập hợp
các kiểu mà hàm được mong đợi sử dụng.
Nó cho phép người gọi hàm generic thấy rõ ràng liệu
hàm có áp dụng được cho các kiểu đang sử dụng hay không.

Trong thực tế, contract này có thể sẽ được đưa vào thư viện chuẩn,
và vì vậy thực sự hàm `Min` (cái có thể cũng sẽ có trong
thư viện chuẩn ở đâu đó) sẽ trông như thế này.
Ở đây chúng ta chỉ tham chiếu đến contract `Ordered` được định nghĩa trong gói
contracts.

{{raw `
	func Min (type T contracts.Ordered) (a, b T) T {
		if a < b {
			return a
		}
		return b
	}
`}}

### Cấu trúc dữ liệu generic

Cuối cùng, hãy xem xét một cấu trúc dữ liệu generic đơn giản, một cây nhị phân.
Trong ví dụ này, cây có hàm so sánh, vì vậy không có
yêu cầu gì về kiểu phần tử.

	type Tree (type E) struct {
		root    *node(E)
		compare func(E, E) int
	}

	type node (type E) struct {
		val         E
		left, right *node(E)
	}

Đây là cách tạo một cây nhị phân mới.
Hàm so sánh được truyền vào hàm `New`.

	func New (type E) (cmp func(E, E) int) *Tree(E) {
		return &Tree(E){compare: cmp}
	}

Một phương thức không xuất khẩu trả về con trỏ đến ô chứa v,
hoặc đến vị trí trong cây nơi nó sẽ được đặt vào.

{{raw `
	func (t *Tree(E)) find(v E) **node(E) {
		pn := &t.root
		for *pn != nil {
			switch cmp := t.compare(v, (*pn).val); {
			case cmp < 0:
				pn = &(*pn).left
			case cmp > 0:
				pn = &(*pn).right
			default:
				return pn
			}
		}
		return pn
	}
`}}

Các chi tiết ở đây không thực sự quan trọng, đặc biệt vì tôi chưa
kiểm thử mã này.
Tôi chỉ cố gắng cho thấy giao diện của việc viết một cấu trúc dữ liệu generic đơn giản.

Đây là mã để kiểm tra xem cây có chứa một giá trị không.

	func (t *Tree(E)) Contains(v E) bool {
		return *t.find(e) != nil
	}

Đây là mã để chèn một giá trị mới.

	func (t *Tree(E)) Insert(v E) bool {
		pn := t.find(v)
		if *pn != nil {
			return false
		}
		*pn = &node(E){val: v}
		return true
	}

Lưu ý rằng kiểu `node` có đối số kiểu `E`.
Đây là giao diện của việc viết một cấu trúc dữ liệu generic.
Như bạn có thể thấy, nó trông giống như viết mã Go thông thường, ngoại trừ
một số đối số kiểu được rải rác đây đó.

Sử dụng cây này khá đơn giản.

	var intTree = tree.New(func(a, b int) int { return a - b })

	func InsertAndCheck(v int) {
		intTree.Insert(v)
		if !intTree.Contains(v) {
			log.Fatalf("%d not found after insertion", v)
		}
	}

Đúng như mong đợi.
Viết một cấu trúc dữ liệu generic khó hơn một chút, vì bạn thường
phải viết tường minh các đối số kiểu cho các kiểu hỗ trợ, nhưng
càng nhiều càng tốt, sử dụng nó không khác gì sử dụng một cấu trúc dữ liệu
không generic thông thường.

### Các bước tiếp theo

Chúng tôi đang làm việc trên các triển khai thực tế để cho phép chúng tôi thử nghiệm
với thiết kế này.
Điều quan trọng là có thể thử thiết kế trong thực tế, để đảm bảo
rằng chúng ta có thể viết các loại chương trình mà chúng ta muốn viết.
Nó chưa tiến nhanh như chúng tôi hy vọng, nhưng chúng tôi sẽ gửi thêm chi tiết
về các triển khai này khi chúng có sẵn.

Robert Griesemer đã viết một
[CL sơ bộ](/cl/187317)
sửa đổi gói go/types.
Điều này cho phép kiểm tra xem mã sử dụng generics và contract có thể
kiểm tra kiểu không.
Nó chưa hoàn chỉnh ngay bây giờ, nhưng nó hầu như hoạt động cho một gói duy nhất,
và chúng tôi sẽ tiếp tục làm việc trên nó.

Những gì chúng tôi muốn mọi người làm với triển khai này và các triển khai tương lai là
thử viết và sử dụng mã generic và xem điều gì xảy ra.
Chúng tôi muốn đảm bảo rằng mọi người có thể viết mã họ cần, và
rằng họ có thể sử dụng nó như mong đợi.
Dĩ nhiên không phải mọi thứ sẽ hoạt động ngay từ đầu, và khi chúng ta khám phá
không gian này, chúng ta có thể phải thay đổi mọi thứ.
Và, để rõ ràng, chúng tôi quan tâm đến phản hồi về ngữ nghĩa nhiều hơn
là về chi tiết cú pháp.

Tôi muốn cảm ơn tất cả những người đã bình luận về thiết kế trước đó, và
tất cả những người đã thảo luận về generics trong Go có thể trông như thế nào.
Chúng tôi đã đọc tất cả các bình luận, và chúng tôi đánh giá cao công việc
mà mọi người đã bỏ ra.
Chúng tôi sẽ không ở đây hôm nay nếu không có công việc đó.

Mục tiêu của chúng tôi là đạt được một thiết kế giúp có thể viết
các loại mã generic tôi đã thảo luận hôm nay, mà không làm cho
ngôn ngữ quá phức tạp để sử dụng hoặc làm cho nó không còn cảm giác như Go nữa.
Chúng tôi hy vọng rằng thiết kế này là một bước hướng tới mục tiêu đó, và chúng tôi mong đợi
tiếp tục điều chỉnh nó khi chúng ta học, từ kinh nghiệm của chúng tôi và của bạn,
điều gì hiệu quả và điều gì không.
Nếu chúng ta đạt được mục tiêu đó, thì chúng ta sẽ có thứ gì đó mà chúng ta có thể
đề xuất cho các phiên bản tương lai của Go.
