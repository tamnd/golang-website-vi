---
title: "Mảng, slice (và chuỗi): Cơ chế của 'append'"
date: 2013-09-26
by:
- Rob Pike
tags:
- array
- slice
- string
- copy
- append
summary: Cách mảng và slice trong Go hoạt động, và cách sử dụng copy cùng append.
template: true
---

## Giới thiệu

Một trong những tính năng phổ biến nhất của các ngôn ngữ lập trình thủ tục là
khái niệm về mảng.
Mảng có vẻ đơn giản nhưng có nhiều câu hỏi phải được
giải đáp khi thêm chúng vào một ngôn ngữ, chẳng hạn như:

  - kích thước cố định hay thay đổi được?
  - kích thước có phải là một phần của kiểu không?
  - mảng nhiều chiều trông như thế nào?
  - mảng rỗng có ý nghĩa gì không?

Câu trả lời cho những câu hỏi này ảnh hưởng đến việc mảng chỉ là
một tính năng của ngôn ngữ hay là một phần cốt lõi trong thiết kế của nó.

Trong giai đoạn phát triển sớm của Go, phải mất khoảng một năm để quyết định câu trả lời
cho những câu hỏi này trước khi thiết kế cảm thấy đúng.
Bước ngoặt quan trọng là sự ra đời của _slice_, được xây dựng trên nền các _mảng_ có kích thước cố định
để tạo ra một cấu trúc dữ liệu linh hoạt và có thể mở rộng.
Tuy nhiên cho đến ngày nay, những lập trình viên mới làm quen với Go vẫn thường vấp váp về cách slice
hoạt động, có lẽ vì kinh nghiệm từ các ngôn ngữ khác đã ảnh hưởng đến cách suy nghĩ của họ.

Trong bài viết này chúng ta sẽ cố gắng làm sáng tỏ sự nhầm lẫn đó.
Chúng ta sẽ làm điều này bằng cách lắp ghép các mảnh để giải thích cách hàm tích hợp `append`
hoạt động và tại sao nó hoạt động theo cách đó.

## Mảng

Mảng là một khối xây dựng quan trọng trong Go, nhưng giống như nền móng của một tòa nhà,
chúng thường bị ẩn bên dưới các thành phần nổi bật hơn.
Chúng ta phải thảo luận ngắn gọn về chúng trước khi chuyển sang ý tưởng thú vị hơn,
mạnh mẽ hơn và nổi bật hơn là slice.

Mảng không thường xuất hiện trong các chương trình Go vì
kích thước của mảng là một phần kiểu của nó, điều này giới hạn khả năng biểu đạt.

Khai báo

{{code "slices/prog010.go" `/var buffer/`}}

khai báo biến `buffer`, giữ 256 byte.
Kiểu của `buffer` bao gồm kích thước của nó, `[256]byte`.
Một mảng với 512 byte sẽ có kiểu khác biệt là `[512]byte`.

Dữ liệu liên kết với một mảng chỉ đơn giản là: một mảng các phần tử.
Nhìn về mặt sơ đồ, buffer của chúng ta trông như thế này trong bộ nhớ:

	buffer: byte byte byte ... 256 lần ... byte byte byte

Nghĩa là, biến giữ 256 byte dữ liệu và không có gì khác. Chúng ta có thể
truy cập các phần tử bằng cú pháp chỉ số quen thuộc, `buffer[0]`, `buffer[1]`,
và tiếp tục đến `buffer[255]`. (Phạm vi chỉ số từ 0 đến 255 bao gồm
256 phần tử.) Cố gắng truy cập `buffer` với một giá trị ngoài phạm vi này
sẽ làm chương trình sụp đổ.

Có một hàm tích hợp gọi là `len` trả về số phần tử
của một mảng hoặc slice và cũng của một số kiểu dữ liệu khác.
Đối với mảng, `len` trả về điều rõ ràng.
Trong ví dụ của chúng ta, `len(buffer)` trả về giá trị cố định 256.

Mảng có chỗ dùng của chúng, chẳng hạn chúng là biểu diễn tốt cho ma trận biến đổi,
nhưng mục đích phổ biến nhất của chúng trong Go là cung cấp vùng lưu trữ cho một slice.

## Slice: Slice header

Slice mới là phần hành động thực sự, nhưng để dùng chúng tốt ta phải hiểu
chính xác chúng là gì và chúng làm gì.

Slice là một cấu trúc dữ liệu mô tả một đoạn liên tiếp của mảng
được lưu trữ riêng biệt với chính biến slice.
_Một slice không phải là một mảng_.
Một slice _mô tả_ một phần của mảng.

Với biến mảng `buffer` từ phần trước, chúng ta có thể tạo
một slice mô tả các phần tử từ 100 đến 150 (chính xác là từ 100 đến 149,
bao gồm cả hai đầu) bằng cách _cắt_ mảng:

{{code "slices/prog010.go" `/var slice/`}}

Trong đoạn đó, chúng ta đã dùng khai báo biến đầy đủ để rõ ràng.
Biến `slice` có kiểu `[]byte`, đọc là "slice of bytes",
và được khởi tạo từ mảng tên `buffer`
bằng cách lấy các phần tử từ 100 (bao gồm) đến 150 (không bao gồm).
Cú pháp thông dụng hơn sẽ bỏ kiểu đi, vì nó được suy ra từ biểu thức khởi tạo:

	var slice = buffer[100:150]

Trong một hàm, chúng ta có thể sử dụng dạng khai báo ngắn gọn:

	slice := buffer[100:150]

Chính xác thì biến slice này là gì?
Chưa phải toàn bộ câu chuyện, nhưng bây giờ hãy nghĩ về
slice như một cấu trúc dữ liệu nhỏ với hai thành phần: một độ dài và một con trỏ tới một phần tử
của mảng.
Bạn có thể nghĩ về nó như được xây dựng như thế này đằng sau hậu trường:

	type sliceHeader struct {
		Length        int
		ZerothElement *byte
	}

	slice := sliceHeader{
		Length:        50,
		ZerothElement: &buffer[100],
	}

Tất nhiên, đây chỉ là minh họa.
Mặc dù đoạn này nói rằng struct `sliceHeader` không thực sự hiển thị
với lập trình viên, và kiểu
của con trỏ phần tử phụ thuộc vào kiểu của các phần tử,
nhưng điều này đưa ra ý tưởng chung về cơ chế.

Cho đến nay chúng ta đã dùng thao tác slice trên một mảng, nhưng chúng ta cũng có thể cắt một slice, như thế này:

	slice2 := slice[5:10]

Cũng như trước, thao tác này tạo ra một slice mới, trong trường hợp này với các phần tử
từ 5 đến 9 (bao gồm cả hai đầu) của slice gốc, tức là các phần tử
từ 105 đến 109 của mảng gốc.
Struct `sliceHeader` bên dưới của biến `slice2` trông như sau:

	slice2 := sliceHeader{
		Length:        5,
		ZerothElement: &buffer[105],
	}

Lưu ý rằng header này vẫn trỏ đến cùng mảng bên dưới, được lưu trong
biến `buffer`.

Chúng ta cũng có thể _reslice_, nghĩa là cắt một slice và lưu kết quả trở lại
vào cấu trúc slice gốc. Sau khi thực hiện

	slice = slice[5:10]

cấu trúc `sliceHeader` của biến `slice` trông giống hệt như đối với biến `slice2`.
Bạn sẽ thấy reslice được sử dụng thường xuyên, ví dụ để cắt ngắn một slice. Câu lệnh này loại bỏ
phần tử đầu và cuối của slice:

	slice = slice[1:len(slice)-1]

[Bài tập: Viết ra cấu trúc `sliceHeader` trông như thế nào sau khi gán này.]

Bạn sẽ thường nghe các lập trình viên Go có kinh nghiệm nói về "slice header"
vì đó thực sự là thứ được lưu trong một biến slice.
Ví dụ, khi bạn gọi một hàm nhận slice làm đối số, chẳng hạn
[bytes.IndexRune](/pkg/bytes/#IndexRune), đó là
header được truyền vào hàm.
Trong lệnh gọi này:

	slashPos := bytes.IndexRune(slice, '/')

đối số `slice` được truyền vào hàm `IndexRune` thực chất là
một "slice header".

Còn một mục dữ liệu nữa trong slice header mà chúng ta sẽ nói về dưới đây,
nhưng trước tiên hãy xem sự tồn tại của slice header có ý nghĩa gì khi bạn
lập trình với slice.

## Truyền slice vào hàm

Điều quan trọng cần hiểu là mặc dù một slice chứa một con trỏ,
bản thân nó là một giá trị.
Phía dưới, nó là một giá trị struct chứa một con trỏ và một độ dài.
Nó _không_ phải là con trỏ tới một struct.

Điều này quan trọng.

Khi chúng ta gọi `IndexRune` trong ví dụ trước,
một _bản sao_ của slice header được truyền vào.
Hành vi đó có những hệ quả quan trọng.

Hãy xem xét hàm đơn giản này:

{{code "slices/prog010.go" `/^func/` `/^}/`}}

Hàm này làm đúng như tên của nó, duyệt qua các chỉ số của một slice
(sử dụng vòng lặp `for` `range`), tăng giá trị các phần tử.

Thử nó:

{{play "slices/prog010.go" `/^func main/` `/^}/`}}

(Bạn có thể chỉnh sửa và thực thi lại các đoạn code có thể chạy này nếu muốn khám phá.)

Mặc dù slice _header_ được truyền theo giá trị, header bao gồm
một con trỏ tới các phần tử của mảng, vì vậy cả slice header gốc
lẫn bản sao của header được truyền vào hàm đều mô tả
cùng một mảng.
Do đó, khi hàm trả về, các phần tử đã được sửa đổi có thể
được nhìn thấy thông qua biến slice gốc.

Đối số của hàm thực sự là một bản sao, như ví dụ này cho thấy:

{{play "slices/prog020.go" `/^func/` `$`}}

Ở đây chúng ta thấy rằng _nội dung_ của một đối số slice có thể bị sửa đổi bởi một hàm,
nhưng _header_ của nó thì không.
Độ dài được lưu trong biến `slice` không bị sửa đổi bởi lời gọi hàm,
vì hàm nhận một bản sao của slice header, không phải bản gốc.
Do đó, nếu muốn viết một hàm sửa đổi header, chúng ta phải trả nó về như một tham số kết quả,
đúng như chúng ta đã làm ở đây.
Biến `slice` không thay đổi nhưng giá trị được trả về có độ dài mới,
sau đó được lưu vào `newSlice`.

## Con trỏ tới slice: Bộ nhận method

Một cách khác để hàm sửa đổi slice header là truyền con trỏ tới nó.
Đây là biến thể của ví dụ trước thực hiện điều đó:

{{play "slices/prog030.go" `/^func/` `$`}}

Điều đó trông có vẻ cồng kềnh trong ví dụ, đặc biệt là khi xử lý mức độ gián tiếp bổ sung
(một biến tạm sẽ giúp ích),
nhưng có một trường hợp phổ biến khi bạn thấy con trỏ tới slice.
Việc sử dụng bộ nhận con trỏ cho một method sửa đổi slice là thông dụng.

Giả sử chúng ta muốn có một method trên một slice để cắt ngắn nó tại dấu gạch chéo cuối cùng.
Chúng ta có thể viết nó như thế này:

{{play "slices/prog040.go" `/^type/` `$`}}

Nếu bạn chạy ví dụ này bạn sẽ thấy nó hoạt động đúng, cập nhật slice trong hàm gọi.

[Bài tập: Thay đổi kiểu của bộ nhận thành giá trị thay vì
con trỏ và chạy lại. Giải thích điều gì xảy ra.]

Mặt khác, nếu muốn viết một method cho `path` chuyển các chữ cái ASCII trong đường dẫn thành chữ hoa
(bỏ qua các tên không phải tiếng Anh), method có thể
dùng bộ nhận giá trị vì bộ nhận giá trị vẫn sẽ trỏ tới cùng mảng bên dưới.

{{play "slices/prog050.go" `/^type/` `$`}}

Ở đây method `ToUpper` sử dụng hai biến trong cấu trúc `for` `range`
để nắm bắt chỉ số và phần tử slice.
Dạng vòng lặp này tránh việc viết `p[i]` nhiều lần trong thân vòng lặp.

[Bài tập: Chuyển đổi method `ToUpper` để sử dụng bộ nhận con trỏ và xem hành vi có thay đổi không.]

[Bài tập nâng cao: Chuyển đổi method `ToUpper` để xử lý các chữ cái Unicode, không chỉ ASCII.]

## Dung lượng

Hãy xem hàm sau mở rộng slice đối số của `int` thêm một phần tử:

{{code "slices/prog060.go" `/^func Extend/` `/^}/`}}

(Tại sao nó cần trả về slice đã sửa đổi?) Bây giờ chạy nó:

{{play "slices/prog060.go" `/^func main/` `/^}/`}}

Hãy xem slice tăng lên cho đến khi... nó không tăng được nữa.

Đã đến lúc nói về thành phần thứ ba của slice header: _dung lượng_ của nó.
Ngoài con trỏ mảng và độ dài, slice header cũng lưu dung lượng của nó:

	type sliceHeader struct {
		Length        int
		Capacity      int
		ZerothElement *byte
	}

Trường `Capacity` ghi lại lượng không gian mà mảng bên dưới thực sự có; đó là giá trị
tối đa mà `Length` có thể đạt đến.
Cố gắng mở rộng slice vượt quá dung lượng của nó sẽ vượt quá giới hạn của mảng và kích hoạt một panic.

Sau khi slice ví dụ của chúng ta được tạo bởi

	slice := iBuffer[0:0]

header của nó trông như thế này:

	slice := sliceHeader{
		Length:        0,
		Capacity:      10,
		ZerothElement: &iBuffer[0],
	}

Trường `Capacity` bằng độ dài của mảng bên dưới,
trừ đi chỉ số trong mảng của phần tử đầu tiên của slice (không trong trường hợp này).
Nếu bạn muốn hỏi dung lượng của một slice là bao nhiêu, hãy sử dụng hàm tích hợp `cap`:

	if cap(slice) == len(slice) {
		fmt.Println("slice is full!")
	}

## Make

Điều gì xảy ra nếu chúng ta muốn mở rộng slice vượt quá dung lượng của nó?
Không thể làm điều đó!
Theo định nghĩa, dung lượng là giới hạn cho sự tăng trưởng.
Nhưng bạn có thể đạt được kết quả tương đương bằng cách cấp phát một mảng mới, sao chép dữ liệu sang, và sửa đổi
slice để mô tả mảng mới.

Hãy bắt đầu với việc cấp phát.
Chúng ta có thể sử dụng hàm tích hợp `new` để cấp phát một mảng lớn hơn
và sau đó cắt kết quả,
nhưng sẽ đơn giản hơn khi dùng hàm tích hợp `make`.
Nó cấp phát một mảng mới và
tạo một slice header để mô tả nó, tất cả cùng một lúc.
Hàm `make` nhận ba đối số: kiểu của slice, độ dài ban đầu của nó, và dung lượng, là
độ dài của mảng mà `make` cấp phát để giữ dữ liệu slice.
Lệnh gọi này tạo ra một slice có độ dài 10 với chỗ cho thêm 5 (15-10), như bạn có thể thấy khi chạy:

{{play "slices/prog070.go" `/slice/` `/fmt/`}}

Đoạn này nhân đôi dung lượng của slice `int` nhưng giữ nguyên độ dài:

{{play "slices/prog080.go" `/slice/` `/OMIT/`}}

Sau khi chạy đoạn code này, slice có nhiều không gian hơn để tăng trưởng trước khi cần cấp phát lại.

Khi tạo slice, thường thì độ dài và dung lượng sẽ giống nhau.
Hàm tích hợp `make` có cách viết tắt cho trường hợp phổ biến này.
Đối số độ dài mặc định bằng dung lượng, vì vậy bạn có thể bỏ nó đi
để đặt cả hai bằng cùng giá trị.
Sau khi thực hiện

	gophers := make([]Gopher, 10)

slice `gophers` có cả độ dài lẫn dung lượng đều bằng 10.

## Copy

Khi chúng ta nhân đôi dung lượng của slice trong phần trước,
chúng ta đã viết một vòng lặp để sao chép dữ liệu cũ sang slice mới.
Go có hàm tích hợp `copy` để làm điều này dễ dàng hơn.
Đối số của nó là hai slice, và nó sao chép dữ liệu từ đối số bên phải sang đối số bên trái.
Đây là ví dụ của chúng ta được viết lại để dùng `copy`:

{{play "slices/prog090.go" `/newSlice/` `/newSlice/`}}

Hàm `copy` khá thông minh.
Nó chỉ sao chép những gì có thể, chú ý đến độ dài của cả hai đối số.
Nói cách khác, số phần tử nó sao chép là giá trị nhỏ hơn trong hai độ dài slice.
Điều này có thể tiết kiệm một chút công đặt sách.
Ngoài ra, `copy` trả về một giá trị nguyên, số phần tử đã sao chép, mặc dù không phải lúc nào cũng cần kiểm tra.

Hàm `copy` cũng xử lý đúng khi nguồn và đích chồng chéo nhau, nghĩa là nó có thể được sử dụng để dịch chuyển
các mục trong một slice duy nhất.
Đây là cách sử dụng `copy` để chèn một giá trị vào giữa một slice.

{{code "slices/prog100.go" `/Insert/` `/^}/`}}

Có một vài điều cần lưu ý trong hàm này.
Thứ nhất, tất nhiên, nó phải trả về slice đã cập nhật vì độ dài của nó đã thay đổi.
Thứ hai, nó sử dụng một cách viết tắt tiện lợi.
Biểu thức

	slice[i:]

có nghĩa chính xác như

	slice[i:len(slice)]

Ngoài ra, mặc dù chúng ta chưa sử dụng thủ thuật đó, chúng ta cũng có thể bỏ phần tử đầu tiên của biểu thức slice;
nó mặc định là không. Như vậy

	slice[:]

chỉ có nghĩa là bản thân slice, điều này hữu ích khi cắt một mảng.
Biểu thức này là cách ngắn nhất để nói "một slice mô tả tất cả các phần tử của mảng":

	array[:]

Đã xong phần đó, hãy chạy hàm `Insert` của chúng ta.

{{play "slices/prog100.go" `/make/` `/OMIT/`}}

## Append: Một ví dụ

Vài phần trước, chúng ta đã viết hàm `Extend` để mở rộng một slice thêm một phần tử.
Tuy nhiên nó có lỗi, vì nếu dung lượng của slice quá nhỏ, hàm sẽ
bị crash.
(Ví dụ `Insert` của chúng ta cũng có vấn đề tương tự.)
Bây giờ chúng ta đã có các mảnh để khắc phục điều đó, vì vậy hãy viết một triển khai robust của
`Extend` cho các slice số nguyên.

{{code "slices/prog110.go" `/func Extend/` `/^}/`}}

Trong trường hợp này, việc trả về slice đặc biệt quan trọng, vì khi cấp phát lại
slice kết quả mô tả một mảng hoàn toàn khác.
Đây là một đoạn nhỏ để minh họa điều gì xảy ra khi slice đầy:

{{play "slices/prog110.go" `/START/` `/END/`}}

Lưu ý quá trình cấp phát lại khi mảng ban đầu kích thước 5 bị đầy.
Cả dung lượng lẫn địa chỉ của phần tử thứ không đều thay đổi khi mảng mới được cấp phát.

Với hàm `Extend` robust làm hướng dẫn, chúng ta có thể viết một hàm hay hơn cho phép
mở rộng slice bằng nhiều phần tử.
Để làm điều này, chúng ta sử dụng khả năng của Go để biến danh sách đối số hàm thành một slice khi hàm
được gọi.
Nghĩa là, chúng ta sử dụng tính năng hàm biến tham số (variadic) của Go.

Hãy đặt tên hàm là `Append`.
Đối với phiên bản đầu tiên, chúng ta có thể chỉ cần gọi `Extend` lặp đi lặp lại để cơ chế của hàm variadic được rõ ràng.
Chữ ký của `Append` là:

	func Append(slice []int, items ...int) []int

Điều đó có nghĩa là `Append` nhận một đối số, một slice, theo sau bởi không hoặc nhiều
đối số `int`.
Những đối số đó chính xác là một slice của `int` theo quan điểm của việc triển khai
`Append`, như bạn có thể thấy:

{{code "slices/prog120.go" `/Append/` `/^}/`}}

Chú ý vòng lặp `for` `range` duyệt qua các phần tử của đối số `items`, có kiểu ngầm định là `[]int`.
Cũng chú ý việc sử dụng định danh trống `_` để loại bỏ chỉ số trong vòng lặp, mà chúng ta không cần trong trường hợp này.

Thử nó:

{{play "slices/prog120.go" `/START/` `/END/`}}

Một kỹ thuật mới khác trong ví dụ này là chúng ta khởi tạo slice bằng cách viết một composite literal,
gồm kiểu của slice theo sau bởi các phần tử của nó trong dấu ngoặc nhọn:

{{code "slices/prog120.go" `/slice := /`}}

Hàm `Append` thú vị vì một lý do khác.
Không chỉ có thể nối các phần tử, chúng ta có thể nối toàn bộ một slice thứ hai
bằng cách "bung" slice thành các đối số sử dụng ký hiệu `...` tại điểm gọi:

{{play "slices/prog130.go" `/START/` `/END/`}}

Tất nhiên, chúng ta có thể làm `Append` hiệu quả hơn bằng cách chỉ cấp phát không quá một lần,
xây dựng trên nền của `Extend`:

{{code "slices/prog140.go" `/Append/` `/^}/`}}

Ở đây, hãy chú ý cách chúng ta sử dụng `copy` hai lần, một lần để di chuyển dữ liệu slice sang bộ nhớ mới được cấp phát,
và sau đó để sao chép các mục được nối vào cuối dữ liệu cũ.

Thử nó; hành vi giống như trước:

{{play "slices/prog140.go" `/START/` `/END/`}}

## Append: Hàm tích hợp

Và như vậy chúng ta đến động lực cho thiết kế của hàm tích hợp `append`.
Nó làm chính xác những gì ví dụ `Append` của chúng ta làm, với hiệu quả tương đương, nhưng nó
hoạt động cho bất kỳ kiểu slice nào.

Một điểm yếu của Go là bất kỳ thao tác kiểu generic nào cũng phải được cung cấp bởi
runtime. Có lẽ một ngày nào đó điều đó sẽ thay đổi, nhưng bây giờ, để làm việc với slice
dễ dàng hơn, Go cung cấp hàm `append` generic tích hợp.
Nó hoạt động giống như phiên bản slice `int` của chúng ta, nhưng cho _bất kỳ_ kiểu slice nào.

Hãy nhớ rằng, vì slice header luôn được cập nhật bởi lời gọi `append`, bạn cần
lưu slice được trả về sau lời gọi.
Thực ra, trình biên dịch sẽ không cho phép bạn gọi append mà không lưu kết quả.

Đây là một số câu lệnh một dòng xen kẽ với các câu lệnh in. Thử chúng, chỉnh sửa và khám phá:

{{play "slices/prog150.go" `/START/` `/END/`}}

Đáng dành một chút thời gian để suy nghĩ chi tiết về câu lệnh một dòng cuối cùng của ví dụ đó để hiểu
cách thiết kế của slice làm cho lời gọi đơn giản này hoạt động đúng.

Có nhiều ví dụ hơn về `append`, `copy`, và các cách khác để sử dụng slice
trên trang Wiki
["Slice Tricks"](/wiki/SliceTricks) do cộng đồng xây dựng.

## Nil

Như một lưu ý bên lề, với kiến thức mới có được, chúng ta có thể thấy biểu diễn của một slice `nil` là gì.
Tự nhiên, đó là giá trị không của slice header:

	sliceHeader{
		Length:        0,
		Capacity:      0,
		ZerothElement: nil,
	}

hoặc chỉ đơn giản là

	sliceHeader{}

Chi tiết quan trọng là con trỏ phần tử cũng là `nil`. Slice được tạo bởi

	array[0:0]

có độ dài bằng không (và thậm chí có thể dung lượng bằng không) nhưng con trỏ của nó không phải là `nil`, vì vậy
nó không phải là một slice nil.

Như đã rõ ràng, một slice rỗng có thể tăng trưởng (giả sử nó có dung lượng khác không), nhưng một slice `nil`
không có mảng để đặt giá trị vào và không bao giờ có thể tăng trưởng để giữ dù chỉ một phần tử.

Mặt khác, một slice `nil` về mặt chức năng tương đương với một slice có độ dài bằng không, dù nó không trỏ tới đâu.
Nó có độ dài bằng không và có thể được nối vào, với việc cấp phát.
Ví dụ, hãy xem câu lệnh một dòng ở trên sao chép một slice bằng cách nối vào
một slice `nil`.

## Chuỗi

Bây giờ một phần ngắn về chuỗi trong Go trong bối cảnh của slice.

Chuỗi thực ra rất đơn giản: chúng chỉ là các slice byte chỉ đọc với một chút
hỗ trợ cú pháp thêm từ ngôn ngữ.

Vì chúng chỉ đọc, không cần dung lượng (bạn không thể mở rộng chúng),
nhưng ngoài ra với hầu hết các mục đích bạn có thể coi chúng giống như các slice byte chỉ đọc.

Để bắt đầu, chúng ta có thể truy cập theo chỉ số để lấy từng byte:

	slash := "/usr/ken"[0] // trả về giá trị byte '/'.

Chúng ta có thể cắt một chuỗi để lấy chuỗi con:

	usr := "/usr/ken"[0:4] // trả về chuỗi "/usr"

Bây giờ đã rõ điều gì xảy ra đằng sau khi chúng ta cắt một chuỗi.

Chúng ta cũng có thể lấy một slice byte thông thường và tạo một chuỗi từ nó bằng chuyển đổi đơn giản:

	str := string(slice)

và đi theo hướng ngược lại cũng được:

	slice := []byte(usr)

Mảng bên dưới một chuỗi bị ẩn khỏi tầm nhìn; không có cách nào truy cập nội dung của nó
ngoài việc thông qua chuỗi. Điều đó có nghĩa là khi chúng ta thực hiện bất kỳ chuyển đổi nào trong số này,
một bản sao của mảng phải được tạo ra.
Tất nhiên Go xử lý việc này, vì vậy bạn không phải làm.
Sau một trong hai chuyển đổi này, những sửa đổi đối với
mảng bên dưới slice byte không ảnh hưởng đến chuỗi tương ứng.

Một hệ quả quan trọng của thiết kế giống slice này cho chuỗi là
việc tạo chuỗi con rất hiệu quả.
Tất cả những gì cần xảy ra
là việc tạo một string header gồm hai từ. Vì chuỗi là chỉ đọc, chuỗi gốc
và chuỗi thu được từ thao tác cắt có thể chia sẻ cùng mảng một cách an toàn.

Một ghi chú lịch sử: Những triển khai sớm nhất của chuỗi luôn cấp phát, nhưng khi slice
được thêm vào ngôn ngữ, chúng cung cấp một mô hình để xử lý chuỗi hiệu quả. Một số
benchmark đã thấy tốc độ tăng lên đáng kể nhờ điều đó.

Còn nhiều điều nữa về chuỗi, và một
[bài viết riêng](/blog/strings) đề cập đến chúng sâu hơn.

## Kết luận

Để hiểu cách slice hoạt động, sẽ hữu ích khi hiểu cách chúng được triển khai.
Có một cấu trúc dữ liệu nhỏ, slice header, là mục được liên kết với biến slice,
và header đó mô tả một đoạn của mảng được cấp phát riêng biệt.
Khi chúng ta truyền các giá trị slice, header được sao chép nhưng mảng mà nó trỏ tới
luôn được chia sẻ.

Khi bạn hiểu cách chúng hoạt động, slice không chỉ trở nên dễ sử dụng mà còn
mạnh mẽ và biểu cảm, đặc biệt với sự hỗ trợ của các hàm tích hợp `copy` và `append`.

## Đọc thêm

Có rất nhiều tài liệu về slice trong Go trên internet.
Như đã đề cập trước đó,
trang Wiki ["Slice Tricks"](/wiki/SliceTricks)
có nhiều ví dụ.
Bài blog [Go Slices](/blog/go-slices-usage-and-internals)
mô tả chi tiết bố cục bộ nhớ với các sơ đồ rõ ràng.
Bài viết [Go Data Structures](https://research.swtch.com/godata) của Russ Cox bao gồm
thảo luận về slice cùng với một số cấu trúc dữ liệu nội bộ khác của Go.

Còn nhiều tài liệu hơn, nhưng cách tốt nhất để tìm hiểu về slice là sử dụng chúng.
