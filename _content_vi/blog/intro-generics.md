---
title: Giới thiệu về Generics
date: 2022-03-22
by:
- Robert Griesemer
- Ian Lance Taylor
tags:
- go2
- generics
summary: Giới thiệu về generics trong Go.
template: true
---

## Giới thiệu

Bài viết blog này dựa trên bài nói chuyện của chúng tôi tại GopherCon 2021:

{{video "https://www.youtube.com/embed/Pa_e9EeCdy8"}}

Bản phát hành Go 1.18 bổ sung hỗ trợ cho generics.
Generics là thay đổi lớn nhất chúng tôi thực hiện với Go kể từ bản phát hành mã nguồn mở đầu tiên.
Trong bài viết này, chúng tôi sẽ giới thiệu các tính năng ngôn ngữ mới.
Chúng tôi sẽ không cố gắng bao phủ tất cả các chi tiết, nhưng sẽ đề cập đến tất cả những điểm quan trọng.
Để có mô tả chi tiết và đầy đủ hơn, kèm theo nhiều ví dụ, hãy xem [tài liệu đề xuất](https://go.googlesource.com/proposal/+/HEAD/design/43651-type-parameters.md).
Để có mô tả chính xác hơn về các thay đổi ngôn ngữ, hãy xem [spec ngôn ngữ đã được cập nhật](/ref/spec).
(Lưu ý rằng triển khai thực tế của 1.18 áp đặt một số hạn chế so với những gì tài liệu đề xuất cho phép; spec sẽ là chính xác.
Các bản phát hành tương lai có thể gỡ bỏ một số hạn chế đó.)

Generics là cách viết mã độc lập với các kiểu cụ thể đang được sử dụng.
Các hàm và kiểu dữ liệu giờ đây có thể được viết để sử dụng với bất kỳ kiểu nào trong một tập hợp các kiểu.

Generics bổ sung ba điều lớn mới vào ngôn ngữ:
1. Tham số kiểu cho hàm và kiểu dữ liệu.
2. Định nghĩa các kiểu interface là tập hợp các kiểu, bao gồm các kiểu không có phương thức.
3. Suy luận kiểu, cho phép bỏ qua các đối số kiểu trong nhiều trường hợp khi gọi hàm.

## Tham số kiểu

Các hàm và kiểu giờ đây được phép có tham số kiểu.
Danh sách tham số kiểu trông giống như danh sách tham số thông thường, ngoại trừ việc nó dùng dấu ngoặc vuông thay vì dấu ngoặc đơn.

Để minh họa cách hoạt động, hãy bắt đầu với hàm `Min` cơ bản không sử dụng generics cho các giá trị dấu phẩy động:

{{raw `
	func Min(x, y float64) float64 {
		if x < y {
			return x
		}
		return y
	}
`}}

Chúng ta có thể biến hàm này thành generic, tức là làm cho nó hoạt động với các kiểu khác nhau, bằng cách thêm một danh sách tham số kiểu.
Trong ví dụ này, chúng ta thêm một danh sách tham số kiểu với một tham số kiểu duy nhất `T`, và thay thế các lần dùng `float64` bằng `T`.

{{raw `
	import "golang.org/x/exp/constraints"

	func GMin[T constraints.Ordered](x, y T) T {
		if x < y {
			return x
		}
		return y
	}
`}}

Bây giờ có thể gọi hàm này với một đối số kiểu bằng cách viết một lời gọi như

{{raw `
	x := GMin[int](2, 3)
`}}

Việc cung cấp đối số kiểu cho `GMin`, trong trường hợp này là `int`, được gọi là _khởi tạo_.
Khởi tạo xảy ra theo hai bước.
Đầu tiên, trình biên dịch thay thế tất cả các đối số kiểu cho các tham số kiểu tương ứng trong toàn bộ hàm hoặc kiểu generic.
Thứ hai, trình biên dịch kiểm tra rằng mỗi đối số kiểu thỏa mãn ràng buộc tương ứng.
Chúng ta sẽ sớm tìm hiểu điều đó có nghĩa là gì, nhưng nếu bước thứ hai thất bại, quá trình khởi tạo thất bại và chương trình không hợp lệ.

Sau khi khởi tạo thành công, chúng ta có một hàm không generic có thể được gọi như bất kỳ hàm nào khác.
Ví dụ, trong mã như

{{raw `
	fmin := GMin[float64]
	m := fmin(2.71, 3.14)
`}}

quá trình khởi tạo `GMin[float64]` tạo ra thực chất là hàm `Min` kiểu dấu phẩy động gốc của chúng ta, và chúng ta có thể dùng nó trong một lời gọi hàm.

Tham số kiểu cũng có thể được dùng với các kiểu dữ liệu.

{{raw `
	type Tree[T interface{}] struct {
		left, right *Tree[T]
		value       T
	}

	func (t *Tree[T]) Lookup(x T) *Tree[T] { ... }

	var stringTree Tree[string]
`}}

Ở đây kiểu generic `Tree` lưu trữ các giá trị của tham số kiểu `T`.
Các kiểu generic có thể có phương thức, như `Lookup` trong ví dụ này.
Để sử dụng một kiểu generic, nó phải được khởi tạo; `Tree[string]` là ví dụ về khởi tạo `Tree` với đối số kiểu `string`.

## Tập hợp kiểu

Hãy xem xét sâu hơn các đối số kiểu có thể được dùng để khởi tạo một tham số kiểu.

Một hàm thông thường có một kiểu cho mỗi tham số giá trị; kiểu đó định nghĩa một tập hợp các giá trị.
Ví dụ, nếu chúng ta có kiểu `float64` như trong hàm `Min` không generic ở trên, tập hợp các giá trị đối số được phép là tập hợp các giá trị dấu phẩy động có thể được biểu diễn bởi kiểu `float64`.

Tương tự, các danh sách tham số kiểu có một kiểu cho mỗi tham số kiểu.
Vì một tham số kiểu bản thân là một kiểu, các kiểu của tham số kiểu định nghĩa các tập hợp các kiểu.
Siêu kiểu này được gọi là _ràng buộc kiểu_.

Trong `GMin` generic, ràng buộc kiểu được import từ [gói constraints](https://golang.org/x/exp/constraints).
Ràng buộc `Ordered` mô tả tập hợp tất cả các kiểu có các giá trị có thể được sắp xếp, hay nói cách khác, so sánh được với toán tử {{" < "}} (hoặc {{" <= "}}, {{" > "}}, v.v.).
Ràng buộc đảm bảo rằng chỉ các kiểu có giá trị có thể sắp xếp mới có thể được truyền vào `GMin`.
Nó cũng có nghĩa là trong thân hàm `GMin`, các giá trị của tham số kiểu đó có thể được dùng trong so sánh với toán tử {{" < "}}.

Trong Go, các ràng buộc kiểu phải là interface.
Tức là, một kiểu interface có thể được dùng như một kiểu giá trị, và nó cũng có thể được dùng như một siêu kiểu.
Các interface định nghĩa các phương thức, vì vậy rõ ràng chúng ta có thể biểu đạt các ràng buộc kiểu yêu cầu một số phương thức nhất định phải có mặt.
Nhưng `constraints.Ordered` cũng là một kiểu interface, và toán tử {{" < "}} không phải là một phương thức.

Để làm cho điều này hoạt động, chúng ta nhìn nhận interface theo cách mới.

Cho đến gần đây, spec Go cho biết một interface định nghĩa một tập phương thức, đó là tập hợp các phương thức được liệt kê trong interface.
Bất kỳ kiểu nào triển khai tất cả các phương thức đó đều triển khai interface đó.

{{image "intro-generics/method-sets.png"}}

Nhưng cách nhìn khác là nói rằng interface định nghĩa một tập hợp các kiểu, cụ thể là các kiểu triển khai các phương thức đó.
Từ góc nhìn này, bất kỳ kiểu nào là phần tử của tập hợp kiểu của interface đều triển khai interface đó.

{{image "intro-generics/type-sets.png"}}

Hai quan điểm dẫn đến cùng kết quả: Với mỗi tập phương thức, chúng ta có thể hình dung tập hợp kiểu tương ứng triển khai các phương thức đó, và đó là tập hợp kiểu được định nghĩa bởi interface.

Tuy nhiên, với mục đích của chúng ta, quan điểm tập hợp kiểu có lợi thế so với quan điểm tập phương thức: chúng ta có thể thêm trực tiếp các kiểu vào tập hợp, và do đó kiểm soát tập hợp kiểu theo những cách mới.

Chúng ta đã mở rộng cú pháp cho các kiểu interface để thực hiện điều này.
Ví dụ, `interface{ int|string|bool }` định nghĩa tập hợp kiểu chứa các kiểu `int`, `string` và `bool`.

{{image "intro-generics/type-sets-2.png"}}

Cách nói khác là interface này chỉ được thỏa mãn bởi `int`, `string` hoặc `bool`.

Bây giờ hãy xem định nghĩa thực tế của `constraints.Ordered`:

{{raw `
	type Ordered interface {
		Integer|Float|~string
	}
`}}

Khai báo này nói rằng interface `Ordered` là tập hợp tất cả các kiểu số nguyên, dấu phẩy động và chuỗi.
Thanh dọc biểu diễn hợp của các kiểu (hoặc tập hợp các kiểu trong trường hợp này).
`Integer` và `Float` là các kiểu interface được định nghĩa tương tự trong gói `constraints`.
Lưu ý rằng không có phương thức nào được định nghĩa bởi interface `Ordered`.

Với các ràng buộc kiểu, thường chúng ta không quan tâm đến một kiểu cụ thể, chẳng hạn như `string`; chúng ta quan tâm đến tất cả các kiểu chuỗi.
Đó là mục đích của token `~`.
Biểu thức `~string` có nghĩa là tập hợp tất cả các kiểu có kiểu cơ bản là `string`.
Điều này bao gồm kiểu `string` bản thân cũng như tất cả các kiểu được khai báo với các định nghĩa như `type MyString string`.

Tất nhiên chúng ta vẫn muốn chỉ định phương thức trong interface, và chúng ta muốn tương thích ngược.
Trong Go 1.18, một interface có thể chứa các phương thức và interface nhúng như trước, nhưng nó cũng có thể nhúng các kiểu không phải interface, các union và các tập hợp của các kiểu cơ bản.

Khi được dùng làm ràng buộc kiểu, tập hợp kiểu được định nghĩa bởi một interface xác định chính xác các kiểu được phép làm đối số kiểu cho tham số kiểu tương ứng.
Trong thân hàm generic, nếu kiểu của một toán hạng là tham số kiểu `P` với ràng buộc `C`, các phép toán được phép nếu chúng được phép bởi tất cả các kiểu trong tập hợp kiểu của `C` (hiện có một số hạn chế triển khai ở đây, nhưng mã thông thường khó gặp phải chúng).

Các interface được dùng làm ràng buộc có thể được đặt tên (như `Ordered`), hoặc chúng có thể là các interface literal nội tuyến trong một danh sách tham số kiểu.
Ví dụ:

{{raw `
	[S interface{~[]E}, E interface{}]
`}}

Ở đây `S` phải là kiểu slice với kiểu phần tử có thể là bất kỳ kiểu nào.

Vì đây là trường hợp phổ biến, `interface{}` bao ngoài có thể bỏ qua với các interface ở vị trí ràng buộc, và chúng ta có thể viết đơn giản:

{{raw `
	[S ~[]E, E interface{}]
`}}

Vì interface rỗng phổ biến trong danh sách tham số kiểu, và trong mã Go thông thường nói chung, Go 1.18 giới thiệu một định danh được khai báo trước mới `any` như một alias cho kiểu interface rỗng.
Với điều đó, chúng ta đến với đoạn mã thành ngữ này:

{{raw `
	[S ~[]E, E any]
`}}

Interface như tập hợp kiểu là một cơ chế mạnh mẽ mới và là chìa khóa để làm cho các ràng buộc kiểu hoạt động trong Go.
Hiện tại, các interface sử dụng các dạng cú pháp mới chỉ có thể được dùng làm ràng buộc.
Nhưng không khó để tưởng tượng cách các interface bị ràng buộc kiểu tường minh có thể hữu ích một cách tổng quát.

## Suy luận kiểu

Tính năng ngôn ngữ lớn mới cuối cùng là suy luận kiểu.
Theo một nghĩa nào đó, đây là thay đổi phức tạp nhất trong ngôn ngữ, nhưng nó quan trọng vì nó cho phép mọi người dùng phong cách tự nhiên khi viết mã gọi các hàm generic.

### Suy luận kiểu từ đối số hàm

Với tham số kiểu, cần truyền các đối số kiểu, điều này có thể làm cho mã trở nên dài dòng.
Trở lại với hàm generic `GMin` của chúng ta:

{{raw `
	func GMin[T constraints.Ordered](x, y T) T { ... }
`}}

tham số kiểu `T` được dùng để chỉ định các kiểu của các đối số không phải kiểu thông thường `x` và `y`.
Như chúng ta đã thấy trước đó, có thể gọi hàm này với một đối số kiểu tường minh

{{raw `
	var a, b, m float64

	m = GMin[float64](a, b) // đối số kiểu tường minh
`}}

Trong nhiều trường hợp, trình biên dịch có thể suy ra đối số kiểu cho `T` từ các đối số thông thường.
Điều này làm cho mã ngắn hơn trong khi vẫn rõ ràng.

{{raw `
	var a, b, m float64

	m = GMin(a, b) // không có đối số kiểu
`}}

Điều này hoạt động bằng cách so khớp các kiểu của đối số `a` và `b` với các kiểu của tham số `x` và `y`.

Loại suy luận này, suy ra các đối số kiểu từ các kiểu của các đối số truyền vào hàm, được gọi là _suy luận kiểu từ đối số hàm_.

Suy luận kiểu từ đối số hàm chỉ hoạt động cho các tham số kiểu được dùng trong các tham số hàm, không dành cho các tham số kiểu chỉ được dùng trong kết quả hàm hoặc chỉ trong thân hàm.
Ví dụ, nó không áp dụng cho các hàm như `MakeT[T any]() T`, chỉ dùng `T` cho kết quả.

### Suy luận kiểu từ ràng buộc

Ngôn ngữ hỗ trợ một loại suy luận kiểu khác, _suy luận kiểu từ ràng buộc_.
Để mô tả điều này, hãy bắt đầu với ví dụ về nhân các phần tử của một slice số nguyên:

{{raw `
	// Scale trả về bản sao của s với mỗi phần tử được nhân với c.
	// Triển khai này có một vấn đề, như chúng ta sẽ thấy.
	func Scale[E constraints.Integer](s []E, c E) []E {
		r := make([]E, len(s))
		for i, v := range s {
			r[i] = v * c
		}
		return r
	}
`}}

Đây là một hàm generic hoạt động với slice của bất kỳ kiểu số nguyên nào.

Bây giờ giả sử chúng ta có kiểu `Point` nhiều chiều, trong đó mỗi `Point` đơn giản là một danh sách số nguyên biểu diễn tọa độ của điểm.
Tự nhiên kiểu này sẽ có một số phương thức.

{{raw `
	type Point []int32

	func (p Point) String() string {
		// Chi tiết không quan trọng.
	}
`}}

Đôi khi chúng ta muốn nhân (scale) một `Point`.
Vì một `Point` chỉ là một slice số nguyên, chúng ta có thể dùng hàm `Scale` đã viết trước đó:

{{raw `
	// ScaleAndPrint nhân đôi một Point và in ra.
	func ScaleAndPrint(p Point) {
		r := Scale(p, 2)
		fmt.Println(r.String()) // KHÔNG BIÊN DỊCH ĐƯỢC
	}
`}}

Thật không may, điều này không biên dịch được, thất bại với lỗi như `r.String undefined (type []int32 has no field or method String)`.

Vấn đề là hàm `Scale` trả về một giá trị kiểu `[]E` trong đó `E` là kiểu phần tử của slice đối số.
Khi chúng ta gọi `Scale` với giá trị kiểu `Point`, có kiểu cơ bản là `[]int32`, chúng ta nhận lại giá trị kiểu `[]int32`, không phải kiểu `Point`.
Điều này xuất phát từ cách mã generic được viết, nhưng đó không phải là điều chúng ta muốn.

Để sửa điều này, chúng ta phải thay đổi hàm `Scale` để dùng một tham số kiểu cho kiểu slice.

{{raw `
	// Scale trả về bản sao của s với mỗi phần tử được nhân với c.
	func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
		r := make(S, len(s))
		for i, v := range s {
			r[i] = v * c
		}
		return r
	}
`}}

Chúng ta đã giới thiệu một tham số kiểu mới `S` là kiểu của đối số slice.
Chúng ta đã ràng buộc nó sao cho kiểu cơ bản là `S` chứ không phải `[]E`, và kiểu kết quả giờ là `S`.
Vì `E` được ràng buộc là số nguyên, hiệu ứng là giống như trước: đối số đầu tiên phải là một slice của một kiểu số nguyên nào đó.
Thay đổi duy nhất trong thân hàm là bây giờ chúng ta truyền `S`, thay vì `[]E`, khi gọi `make`.

Hàm mới hoạt động giống như trước nếu chúng ta gọi với một slice thuần túy, nhưng nếu chúng ta gọi với kiểu `Point`, bây giờ chúng ta nhận lại giá trị kiểu `Point`.
Đó là điều chúng ta muốn.
Với phiên bản này của `Scale`, hàm `ScaleAndPrint` trước đó sẽ biên dịch và chạy như mong đợi.

Nhưng có thể hỏi: tại sao có thể viết lời gọi tới `Scale` mà không truyền đối số kiểu tường minh?
Tức là, tại sao chúng ta có thể viết `Scale(p, 2)`, không có đối số kiểu, thay vì phải viết `Scale[Point, int32](p, 2)`?
Hàm `Scale` mới của chúng ta có hai tham số kiểu, `S` và `E`.
Trong một lời gọi tới `Scale` không truyền bất kỳ đối số kiểu nào, suy luận kiểu từ đối số hàm, được mô tả ở trên, cho phép trình biên dịch suy ra rằng đối số kiểu cho `S` là `Point`.
Nhưng hàm cũng có tham số kiểu `E` là kiểu của hệ số nhân `c`.
Đối số hàm tương ứng là `2`, và vì `2` là một hằng số _không có kiểu_, suy luận kiểu từ đối số hàm không thể suy ra kiểu chính xác cho `E` (nhiều nhất nó có thể suy ra kiểu mặc định cho `2` là `int`, điều này sẽ không đúng).
Thay vào đó, quá trình mà trình biên dịch suy ra rằng đối số kiểu cho `E` là kiểu phần tử của slice được gọi là _suy luận kiểu từ ràng buộc_.

Suy luận kiểu từ ràng buộc suy ra các đối số kiểu từ các ràng buộc tham số kiểu.
Nó được dùng khi một tham số kiểu có ràng buộc được định nghĩa theo tham số kiểu khác.
Khi đối số kiểu của một trong các tham số kiểu đó đã biết, ràng buộc được dùng để suy ra đối số kiểu của tham số còn lại.

Trường hợp thông thường áp dụng điều này là khi một ràng buộc dùng dạng `~`_`type`_ cho một số kiểu nào đó, trong đó kiểu đó được viết bằng các tham số kiểu khác.
Chúng ta thấy điều này trong ví dụ `Scale`.
`S` là `~[]E`, tức là `~` theo sau là kiểu `[]E` được viết theo tham số kiểu khác.
Nếu chúng ta biết đối số kiểu cho `S`, chúng ta có thể suy ra đối số kiểu cho `E`.
`S` là một kiểu slice, và `E` là kiểu phần tử của slice đó.

Đây chỉ là phần giới thiệu về suy luận kiểu từ ràng buộc.
Để biết chi tiết đầy đủ, hãy xem [tài liệu đề xuất](https://go.googlesource.com/proposal/+/HEAD/design/43651-type-parameters.md) hoặc [spec ngôn ngữ](/ref/spec).

### Suy luận kiểu trong thực tế

Chi tiết chính xác về cách suy luận kiểu hoạt động khá phức tạp, nhưng sử dụng nó thì không: suy luận kiểu hoặc thành công hoặc thất bại.
Nếu thành công, các đối số kiểu có thể bỏ qua, và việc gọi các hàm generic trông không khác gì gọi các hàm thông thường.
Nếu suy luận kiểu thất bại, trình biên dịch sẽ báo lỗi, và trong những trường hợp đó chúng ta chỉ cần cung cấp các đối số kiểu cần thiết.

Trong việc thêm suy luận kiểu vào ngôn ngữ, chúng tôi đã cố gắng cân bằng giữa sức mạnh suy luận và độ phức tạp.
Chúng tôi muốn đảm bảo rằng khi trình biên dịch suy ra kiểu, những kiểu đó không bao giờ gây ngạc nhiên.
Chúng tôi đã cố gắng thận trọng để nghiêng về phía thất bại trong việc suy ra một kiểu hơn là nghiêng về phía suy ra kiểu sai.
Chúng tôi có lẽ chưa hoàn toàn đúng, và chúng tôi có thể tiếp tục cải thiện nó trong các bản phát hành tương lai.
Hiệu ứng là ngày càng nhiều chương trình có thể được viết mà không cần đối số kiểu tường minh.
Các chương trình không cần đối số kiểu hôm nay sẽ cũng không cần chúng vào ngày mai.

## Kết luận

Generics là một tính năng ngôn ngữ mới lớn trong 1.18.
Các thay đổi ngôn ngữ mới này đòi hỏi một lượng lớn mã mới chưa được kiểm thử đáng kể trong môi trường production.
Điều đó chỉ xảy ra khi ngày càng nhiều người viết và sử dụng mã generic.
Chúng tôi tin rằng tính năng này được triển khai tốt và có chất lượng cao.
Tuy nhiên, không giống như hầu hết các khía cạnh của Go, chúng tôi không thể hỗ trợ niềm tin đó bằng kinh nghiệm thực tế.
Do đó, trong khi chúng tôi khuyến khích sử dụng generics khi hợp lý, hãy thận trọng khi triển khai mã generic trong môi trường production.

Dù vậy, chúng tôi rất vui khi có generics, và chúng tôi hy vọng chúng sẽ giúp các lập trình viên Go làm việc hiệu quả hơn.
