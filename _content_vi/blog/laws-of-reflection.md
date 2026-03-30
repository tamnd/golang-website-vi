---
title: Các quy luật của reflection
date: 2011-09-06
by:
- Rob Pike
tags:
- interface
- reflect
- type
- technical
summary: Reflection trong Go hoạt động như thế nào, cách suy nghĩ về nó, và cách sử dụng nó.
template: true
---

## Giới thiệu

Reflection trong máy tính là khả năng của một chương trình kiểm tra cấu trúc của chính nó,
đặc biệt thông qua các kiểu;
đó là một dạng lập trình meta.
Nó cũng là nguồn gốc lớn của sự nhầm lẫn.

Trong bài viết này, chúng ta cố gắng làm rõ mọi thứ bằng cách giải thích cách reflection hoạt động trong Go.
Mô hình reflection của mỗi ngôn ngữ là khác nhau (và nhiều ngôn ngữ không hỗ trợ nó chút nào),
nhưng bài viết này nói về Go, vì vậy trong phần còn lại, từ "reflection"
nên được hiểu là "reflection trong Go".

Ghi chú thêm tháng 1 năm 2022: Bài viết blog này được viết vào năm 2011 và trước
tính đa hình tham số (hay còn gọi là generics) trong Go.
Mặc dù không có điều gì quan trọng trong bài viết trở nên sai sót do sự phát triển đó trong ngôn ngữ,
nó đã được chỉnh sửa ở một vài chỗ để tránh
gây nhầm lẫn cho người quen thuộc với Go hiện đại.

## Kiểu và interface

Vì reflection xây dựng trên hệ thống kiểu, hãy bắt đầu với ôn lại về kiểu trong Go.

Go là ngôn ngữ kiểu tĩnh. Mỗi biến có một kiểu tĩnh,
tức là, chính xác một kiểu được biết và cố định tại thời điểm biên dịch:
`int`, `float32`, `*MyType`, `[]byte`, và vân vân. Nếu chúng ta khai báo

	type MyInt int

	var i int
	var j MyInt

thì `i` có kiểu `int` và `j` có kiểu `MyInt`.
Các biến `i` và `j` có kiểu tĩnh khác nhau và,
mặc dù chúng có cùng kiểu cơ bản,
chúng không thể gán cho nhau mà không có chuyển đổi.

Một danh mục quan trọng của kiểu là kiểu interface,
biểu diễn các tập hợp phương thức cố định.
(Khi thảo luận về reflection, chúng ta có thể bỏ qua việc sử dụng
định nghĩa interface như ràng buộc trong mã đa hình.)
Một biến interface có thể lưu trữ bất kỳ giá trị concrete (không phải interface) nào miễn là
giá trị đó triển khai các phương thức của interface.
Một cặp ví dụ nổi tiếng là `io.Reader` và `io.Writer`,
các kiểu `Reader` và `Writer` từ [package io](/pkg/io/):

	// Reader là interface bao bọc phương thức Read cơ bản.
	type Reader interface {
	    Read(p []byte) (n int, err error)
	}

	// Writer là interface bao bọc phương thức Write cơ bản.
	type Writer interface {
	    Write(p []byte) (n int, err error)
	}

Bất kỳ kiểu nào triển khai phương thức `Read` (hoặc `Write`) với chữ ký này
được cho là triển khai `io.Reader` (hoặc `io.Writer`).
Để phục vụ cho cuộc thảo luận này, điều đó có nghĩa là biến kiểu
`io.Reader` có thể chứa bất kỳ giá trị nào có kiểu với phương thức `Read`:

	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	// và vân vân

Điều quan trọng cần làm rõ là dù giá trị concrete nào `r` có thể chứa,
kiểu của `r` luôn là `io.Reader`:
Go là kiểu tĩnh và kiểu tĩnh của `r` là `io.Reader`.

Một ví dụ cực kỳ quan trọng của kiểu interface là interface rỗng:

	interface{}

hoặc alias tương đương của nó,

	any

Nó biểu diễn tập hợp phương thức rỗng và được thỏa mãn bởi bất kỳ giá trị nào, vì mọi giá trị đều có không hoặc nhiều phương thức.

Một số người nói rằng interface của Go được kiểu động,
nhưng điều đó gây hiểu nhầm.
Chúng có kiểu tĩnh: một biến kiểu interface luôn có cùng kiểu tĩnh,
và mặc dù tại thời điểm chạy, giá trị được lưu trong biến interface có thể thay đổi kiểu,
giá trị đó luôn thỏa mãn interface.

Chúng ta cần chính xác về tất cả những điều này vì reflection và interface có liên quan chặt chẽ.

## Biểu diễn của interface

Russ Cox đã viết một [bài blog chi tiết](https://research.swtch.com/2009/12/go-data-structures-interfaces.html)
về biểu diễn của các giá trị interface trong Go.
Không cần thiết phải lặp lại toàn bộ câu chuyện ở đây,
nhưng một tóm tắt đơn giản là cần thiết.

Một biến kiểu interface lưu một cặp:
giá trị concrete được gán cho biến,
và mô tả kiểu của giá trị đó.
Cụ thể hơn, giá trị là dữ liệu concrete cơ bản triển khai interface và kiểu mô tả kiểu đầy đủ của mục đó. Ví dụ, sau

	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
	    return nil, err
	}
	r = tty

`r` chứa, về mặt sơ đồ, cặp (giá trị, kiểu),
(`tty`, `*os.File`).
Lưu ý rằng kiểu `*os.File` triển khai các phương thức khác ngoài `Read`;
mặc dù giá trị interface chỉ cung cấp quyền truy cập vào phương thức `Read`,
giá trị bên trong mang tất cả thông tin kiểu về giá trị đó.
Đó là lý do tại sao chúng ta có thể làm những thứ như thế này:

	var w io.Writer
	w = r.(io.Writer)

Biểu thức trong phép gán này là một type assertion;
những gì nó khẳng định là mục bên trong `r` cũng triển khai `io.Writer`,
và do đó chúng ta có thể gán nó cho `w`.
Sau phép gán, `w` sẽ chứa cặp (`tty`, `*os.File`).
Đó là cùng cặp với `r`. Kiểu tĩnh của interface
xác định phương thức nào có thể được gọi với biến interface,
mặc dù giá trị concrete bên trong có thể có tập phương thức lớn hơn.

Tiếp tục, chúng ta có thể làm điều này:

	var empty interface{}
	empty = w

và giá trị interface rỗng của chúng ta `empty` sẽ một lần nữa chứa cùng cặp,
(`tty`, `*os.File`).
Điều đó tiện lợi: một interface rỗng có thể chứa bất kỳ giá trị nào và chứa tất cả
thông tin chúng ta có thể cần về giá trị đó.

(Chúng ta không cần type assertion ở đây vì được biết tĩnh rằng
`w` thỏa mãn interface rỗng.
Trong ví dụ chúng ta chuyển giá trị từ `Reader` sang `Writer`,
chúng ta cần rõ ràng và sử dụng type assertion vì các phương thức của `Writer`
không phải là tập con của `Reader`.)

Một chi tiết quan trọng là cặp bên trong biến interface luôn có dạng (giá trị,
kiểu concrete) và không thể có dạng (giá trị, kiểu interface).
Interface không chứa giá trị interface.

Bây giờ chúng ta sẵn sàng để phản chiếu.

## Quy luật đầu tiên của reflection

## 1. Reflection đi từ giá trị interface đến đối tượng reflection.

Ở cấp độ cơ bản, reflection chỉ là cơ chế để kiểm tra cặp kiểu và
giá trị được lưu bên trong biến interface.
Để bắt đầu, có hai kiểu chúng ta cần biết trong [package reflect](/pkg/reflect/):
[Type](/pkg/reflect/#Type) và [Value](/pkg/reflect/#Value).
Hai kiểu đó cung cấp quyền truy cập vào nội dung của biến interface,
và hai hàm đơn giản, gọi là `reflect.TypeOf` và `reflect.ValueOf`,
truy xuất các phần `reflect.Type` và `reflect.Value` từ một giá trị interface.
(Cũng từ một `reflect.Value`, thật dễ dàng để lấy `reflect.Type` tương ứng,
nhưng hãy tạm thời giữ các khái niệm `Value` và `Type` tách biệt.)

Hãy bắt đầu với `TypeOf`:

	package main

	import (
	    "fmt"
	    "reflect"
	)

	func main() {
	    var x float64 = 3.4
	    fmt.Println("type:", reflect.TypeOf(x))
	}

Chương trình này in ra

	type: float64

Bạn có thể thắc mắc interface ở đây là gì,
vì chương trình trông như thể đang truyền biến `float64` `x`,
không phải giá trị interface, cho `reflect.TypeOf`.
Nhưng nó có ở đó; như [godoc báo cáo](/pkg/reflect/#TypeOf),
chữ ký của `reflect.TypeOf` bao gồm một interface rỗng:

	// TypeOf trả về kiểu reflection của giá trị trong interface{}.
	func TypeOf(i interface{}) Type

Khi chúng ta gọi `reflect.TypeOf(x)`, `x` đầu tiên được lưu trong một interface rỗng,
sau đó được truyền như đối số;
`reflect.TypeOf` giải nén interface rỗng đó để lấy lại thông tin kiểu.

Hàm `reflect.ValueOf`, tất nhiên,
lấy lại giá trị (từ đây chúng ta sẽ bỏ qua mã boilerplate và chỉ tập trung
vào mã thực thi):

	var x float64 = 3.4
	fmt.Println("value:", reflect.ValueOf(x).String())

in ra

{{raw `
	value: <float64 Value>
`}}

(Chúng ta gọi phương thức `String` một cách rõ ràng vì theo mặc định package `fmt`
đào sâu vào một `reflect.Value` để hiển thị giá trị concrete bên trong.
Phương thức `String` thì không.)

Cả `reflect.Type` và `reflect.Value` đều có nhiều phương thức để chúng ta
kiểm tra và thao tác với chúng.
Một ví dụ quan trọng là `Value` có phương thức `Type` trả về
`Type` của một `reflect.Value`.
Một ví dụ khác là cả `Type` và `Value` đều có phương thức `Kind` trả về
một hằng số chỉ ra loại mục nào được lưu:
`Uint`, `Float64`, `Slice`, và vân vân.
Cũng các phương thức trên `Value` với tên như `Int` và `Float` cho phép chúng ta lấy giá trị
(như `int64` và `float64`) được lưu bên trong:

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

in ra

	type: float64
	kind is float64: true
	value: 3.4

Cũng có các phương thức như `SetInt` và `SetFloat` nhưng để sử dụng chúng, chúng ta cần
hiểu khả năng thiết lập (settability),
chủ đề của quy luật thứ ba của reflection, được thảo luận bên dưới.

Thư viện reflection có một vài thuộc tính đáng chú ý.
Thứ nhất, để API đơn giản, các phương thức "getter" và "setter" của `Value`
hoạt động trên kiểu lớn nhất có thể chứa giá trị:
`int64` cho tất cả các số nguyên có dấu, chẳng hạn.
Tức là, phương thức `Int` của `Value` trả về `int64` và giá trị `SetInt`
nhận `int64`;
có thể cần chuyển đổi sang kiểu thực tế liên quan:

	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	x = uint8(v.Uint())                                       // v.Uint trả về uint64.

Thuộc tính thứ hai là `Kind` của đối tượng reflection mô tả
kiểu cơ bản,
không phải kiểu tĩnh.
Nếu đối tượng reflection chứa giá trị của kiểu số nguyên do người dùng định nghĩa, như trong

	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)

`Kind` của `v` vẫn là `reflect.Int`,
mặc dù kiểu tĩnh của `x` là `MyInt`, không phải `int`.
Nói cách khác, `Kind` không thể phân biệt `int` với `MyInt` mặc dù
`Type` có thể.

## Quy luật thứ hai của reflection

## 2. Reflection đi từ đối tượng reflection đến giá trị interface.

Giống như phản chiếu vật lý, reflection trong Go tạo ra nghịch đảo của chính nó.

Với một `reflect.Value`, chúng ta có thể lấy lại giá trị interface bằng phương thức `Interface`;
thực tế phương thức đóng gói lại thông tin kiểu và giá trị vào biểu diễn interface
và trả về kết quả:

	// Interface trả về giá trị của v như một interface{}.
	func (v Value) Interface() interface{}

Kết quả là chúng ta có thể nói

	y := v.Interface().(float64) // y sẽ có kiểu float64.
	fmt.Println(y)

để in giá trị `float64` được biểu diễn bởi đối tượng reflection `v`.

Chúng ta còn có thể làm tốt hơn, thực ra. Các đối số cho `fmt.Println`,
`fmt.Printf` và vân vân đều được truyền như giá trị interface rỗng,
sau đó được giải nén bởi package `fmt` nội bộ giống như những gì chúng ta
đã làm trong các ví dụ trước.
Do đó tất cả những gì cần để in nội dung của `reflect.Value` đúng cách
là truyền kết quả của phương thức `Interface` cho quy trình in được định dạng:

	fmt.Println(v.Interface())

(Kể từ khi bài viết này được viết lần đầu, đã có thay đổi với package `fmt`
để nó tự động giải nén `reflect.Value` như vậy, vì vậy
chúng ta chỉ cần nói

	fmt.Println(v)

để có kết quả giống nhau, nhưng để rõ ràng, chúng ta sẽ giữ các lệnh gọi `.Interface()`
ở đây.)

Vì giá trị của chúng ta là `float64`,
chúng ta thậm chí có thể sử dụng định dạng dấu phẩy động nếu muốn:

	fmt.Printf("value is %7.1e\n", v.Interface())

và thu được trong trường hợp này

	3.4e+00

Một lần nữa, không cần type-assert kết quả của `v.Interface()` sang `float64`;
giá trị interface rỗng có thông tin kiểu của giá trị concrete bên trong
và `Printf` sẽ lấy nó ra.

Tóm lại, phương thức `Interface` là nghịch đảo của hàm `ValueOf`,
ngoại trừ kết quả của nó luôn có kiểu tĩnh là `interface{}`.

Nhắc lại: Reflection đi từ giá trị interface đến đối tượng reflection và trở lại.

## Quy luật thứ ba của reflection

## 3. Để sửa đổi đối tượng reflection, giá trị phải có thể thiết lập.

Quy luật thứ ba là tinh tế và gây nhầm lẫn nhất, nhưng nó đủ dễ hiểu nếu chúng ta bắt đầu từ các nguyên tắc đầu tiên.

Đây là một số mã không hoạt động, nhưng đáng nghiên cứu.

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1) // Lỗi: sẽ panic.

Nếu bạn chạy mã này, nó sẽ panic với thông báo khó hiểu

	panic: reflect.Value.SetFloat using unaddressable value

Vấn đề không phải là giá trị `7.1` không có địa chỉ;
mà là `v` không thể thiết lập.
Khả năng thiết lập là thuộc tính của `Value` reflection,
và không phải tất cả `Value` reflection đều có nó.

Phương thức `CanSet` của `Value` báo cáo khả năng thiết lập của một `Value`; trong trường hợp của chúng ta,

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())

in ra

	settability of v: false

Việc gọi phương thức `Set` trên `Value` không thể thiết lập là lỗi. Nhưng khả năng thiết lập là gì?

Khả năng thiết lập giống khả năng đánh địa chỉ, nhưng nghiêm ngặt hơn.
Đó là thuộc tính mà đối tượng reflection có thể sửa đổi storage thực sự
đã được sử dụng để tạo đối tượng reflection.
Khả năng thiết lập được xác định bởi liệu đối tượng reflection có giữ mục gốc hay không. Khi chúng ta nói

	var x float64 = 3.4
	v := reflect.ValueOf(x)

chúng ta truyền bản sao của `x` cho `reflect.ValueOf`,
vì vậy giá trị interface được tạo như đối số cho `reflect.ValueOf` là
bản sao của `x`, không phải `x` chính nó.
Do đó, nếu câu lệnh

	v.SetFloat(7.1)

được phép thành công, nó sẽ không cập nhật `x`,
mặc dù `v` trông như được tạo từ `x`.
Thay vào đó, nó sẽ cập nhật bản sao của `x` được lưu bên trong giá trị reflection
và `x` chính nó sẽ không bị ảnh hưởng.
Điều đó sẽ gây nhầm lẫn và vô dụng, vì vậy nó bất hợp pháp,
và khả năng thiết lập là thuộc tính được sử dụng để tránh vấn đề này.

Nếu điều này có vẻ lạ, thực ra không phải. Đây thực sự là một tình huống quen thuộc trong bộ quần áo khác thường.
Hãy nghĩ về việc truyền `x` cho một hàm:

	f(x)

Chúng ta sẽ không mong đợi `f` có thể sửa đổi `x` vì chúng ta đã truyền bản sao
của giá trị `x`, không phải `x` chính nó.
Nếu muốn `f` sửa đổi `x` trực tiếp, chúng ta phải truyền cho hàm địa chỉ
của `x` (tức là, con trỏ đến `x`):

	f(&x)

Điều này đơn giản và quen thuộc, và reflection hoạt động theo cách tương tự.
Nếu muốn sửa đổi `x` bằng reflection, chúng ta phải cung cấp cho thư viện reflection
một con trỏ đến giá trị chúng ta muốn sửa đổi.

Hãy làm điều đó. Đầu tiên chúng ta khởi tạo `x` như thường lệ rồi tạo một giá trị reflection trỏ đến nó, gọi là `p`.

	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Lưu ý: lấy địa chỉ của x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

Đầu ra cho đến nay là

	type of p: *float64
	settability of p: false

Đối tượng reflection `p` không thể thiết lập,
nhưng không phải `p` chúng ta muốn thiết lập, mà là (thực tế) `*p`.
Để đến được những gì `p` trỏ đến, chúng ta gọi phương thức `Elem` của `Value`,
điều hướng qua con trỏ, và lưu kết quả trong một `Value` reflection gọi là `v`:

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

Bây giờ `v` là một đối tượng reflection có thể thiết lập, như đầu ra minh chứng,

	settability of v: true

và vì nó biểu diễn `x`, cuối cùng chúng ta có thể sử dụng `v.SetFloat` để sửa đổi giá trị của `x`:

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)

Đầu ra, như mong đợi, là

	7.1
	7.1

Reflection có thể khó hiểu nhưng nó đang làm chính xác những gì ngôn ngữ làm,
mặc dù qua `Types` và `Values` reflection có thể che giấu những gì đang xảy ra.
Chỉ cần nhớ rằng các `Value` reflection cần địa chỉ của một thứ gì đó
để sửa đổi những gì chúng biểu diễn.

## Struct

Trong ví dụ trước, `v` không phải là con trỏ chính nó,
nó chỉ được lấy từ một con trỏ.
Một cách phổ biến để tình huống này xuất hiện là khi sử dụng reflection để sửa đổi
các trường của cấu trúc.
Miễn là chúng ta có địa chỉ của cấu trúc,
chúng ta có thể sửa đổi các trường của nó.

Đây là một ví dụ đơn giản phân tích giá trị struct `t`.
Chúng ta tạo đối tượng reflection với địa chỉ của struct vì chúng ta sẽ
muốn sửa đổi nó sau.
Sau đó chúng ta đặt `typeOfT` theo kiểu của nó và lặp qua các trường bằng cách gọi phương thức đơn giản
(xem [package reflect](/pkg/reflect/) để biết chi tiết).
Lưu ý rằng chúng ta trích xuất tên của các trường từ kiểu struct,
nhưng các trường chính chúng là các đối tượng `reflect.Value` thông thường.

{{raw `
	type T struct {
	    A int
	    B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
	    f := s.Field(i)
	    fmt.Printf("%d: %s %s = %v\n", i,
	        typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
`}}

Đầu ra của chương trình này là

	0: A int = 23
	1: B string = skidoo

Có thêm một điểm về khả năng thiết lập được giới thiệu ở đây:
tên trường của `T` là chữ hoa (exported) vì chỉ các trường exported
của struct mới có thể thiết lập.

Vì `s` chứa đối tượng reflection có thể thiết lập, chúng ta có thể sửa đổi các trường của cấu trúc.

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)

Và đây là kết quả:

	t is now {77 Sunset Strip}

Nếu chúng ta sửa đổi chương trình để `s` được tạo từ `t`,
không phải `&t`, các lệnh gọi đến `SetInt` và `SetString` sẽ thất bại vì các trường
của `t` sẽ không thể thiết lập.

## Kết luận

Đây lại là các quy luật của reflection:

  - Reflection đi từ giá trị interface đến đối tượng reflection.

  - Reflection đi từ đối tượng reflection đến giá trị interface.

  - Để sửa đổi đối tượng reflection, giá trị phải có thể thiết lập.

Khi bạn hiểu những quy luật này, reflection trong Go trở nên dễ sử dụng hơn nhiều,
mặc dù nó vẫn tinh tế.
Đây là công cụ mạnh mẽ nên được sử dụng cẩn thận và tránh trừ khi thực sự cần thiết.

Còn nhiều điều về reflection mà chúng ta chưa đề cập, gửi và
nhận trên channel,
cấp phát bộ nhớ, sử dụng slice và map,
gọi phương thức và hàm, nhưng bài viết này đã đủ dài.
Chúng ta sẽ đề cập một số chủ đề đó trong bài viết sau.
