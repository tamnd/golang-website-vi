---
title: Khi nào nên dùng Generics
date: 2022-04-12
by:
- Ian Lance Taylor
tags:
- go2
- generics
summary: Khi nào nên dùng generics khi viết code Go, và khi nào không nên dùng.
template: true
---

## Giới thiệu

Đây là phiên bản bài blog của các bài nói của tôi tại Google Open Source Live:

{{video "https://www.youtube.com/embed/nr8EpUO9jhw"}}

và GopherCon 2021:

{{video "https://www.youtube.com/embed/Pa_e9EeCdy8?start=1244"}}

Bản phát hành Go 1.18 bổ sung một tính năng ngôn ngữ lớn: hỗ trợ lập trình generic. Trong bài viết này tôi sẽ không mô tả generics là gì hay cách sử dụng chúng. Bài viết này nói về khi nào nên dùng generics trong code Go, và khi nào không nên dùng.

Để rõ ràng, tôi sẽ cung cấp các hướng dẫn chung, không phải quy tắc cứng nhắc. Hãy dùng phán đoán của chính bạn. Nhưng nếu bạn không chắc, tôi khuyến nghị sử dụng các hướng dẫn được thảo luận ở đây.

## Viết code

Hãy bắt đầu với một hướng dẫn chung cho lập trình Go: viết chương trình Go bằng cách viết code, không phải bằng cách định nghĩa kiểu. Khi nói đến generics, nếu bạn bắt đầu viết chương trình bằng cách định nghĩa các ràng buộc tham số kiểu, bạn có thể đang đi sai hướng. Hãy bắt đầu bằng cách viết hàm. Việc thêm tham số kiểu sau này khi rõ ràng chúng sẽ hữu ích là điều dễ dàng.

## Khi nào tham số kiểu hữu ích?

Với điều đó đã nói, hãy xem xét các trường hợp tham số kiểu có thể hữu ích.

### Khi sử dụng các kiểu container do ngôn ngữ định nghĩa

Một trường hợp là khi viết các hàm hoạt động trên các kiểu container đặc biệt do ngôn ngữ định nghĩa: slice, map và channel. Nếu một hàm có các tham số với các kiểu đó, và code hàm không đưa ra bất kỳ giả định cụ thể nào về các kiểu phần tử, thì có thể hữu ích khi dùng tham số kiểu.

Ví dụ, đây là một hàm trả về một slice của tất cả các khóa trong một map bất kỳ kiểu:

{{raw `
	// MapKeys trả về một slice của tất cả các khóa trong m.
	// Các khóa không được trả về theo thứ tự cụ thể nào.
	func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
		s := make([]Key, 0, len(m))
		for k := range m {
			s = append(s, k)
		}
		return s
	}
`}}

Code này không giả định gì về kiểu khóa map, và nó không sử dụng kiểu giá trị map chút nào. Nó hoạt động với bất kỳ kiểu map nào. Điều đó làm cho nó trở thành ứng viên tốt cho việc sử dụng tham số kiểu.

Thay thế cho tham số kiểu cho loại hàm này thường là sử dụng reflection, nhưng đó là mô hình lập trình kém thoải mái hơn, không được kiểm tra kiểu tĩnh tại thời điểm build, và thường chậm hơn lúc runtime.

### Cấu trúc dữ liệu đa dụng

Một trường hợp khác tham số kiểu có thể hữu ích là cho các cấu trúc dữ liệu đa dụng. Cấu trúc dữ liệu đa dụng là thứ như slice hoặc map, nhưng không được tích hợp vào ngôn ngữ, chẳng hạn như danh sách liên kết, hoặc cây nhị phân.

Ngày nay, các chương trình cần những cấu trúc dữ liệu như vậy thường làm một trong hai: viết chúng với kiểu phần tử cụ thể, hoặc sử dụng kiểu interface. Thay thế kiểu phần tử cụ thể bằng tham số kiểu có thể tạo ra cấu trúc dữ liệu tổng quát hơn có thể được sử dụng ở các phần khác của chương trình, hoặc bởi các chương trình khác. Thay thế kiểu interface bằng tham số kiểu có thể cho phép dữ liệu được lưu trữ hiệu quả hơn, tiết kiệm tài nguyên bộ nhớ; nó cũng có thể cho phép code tránh các type assertion và được kiểm tra kiểu hoàn toàn tại thời điểm build.

Ví dụ, đây là một phần về cấu trúc dữ liệu cây nhị phân có thể trông như thế nào khi sử dụng tham số kiểu:

{{raw `
	// Tree là một cây nhị phân.
	type Tree[T any] struct {
		cmp  func(T, T) int
		root *node[T]
	}

	// Một nút trong Tree.
	type node[T any] struct {
		left, right  *node[T]
		val          T
	}

	// find trả về con trỏ đến nút chứa val,
	// hoặc, nếu val không có mặt, con trỏ đến nơi
	// nó sẽ được đặt nếu thêm vào.
	func (bt *Tree[T]) find(val T) **node[T] {
		pl := &bt.root
		for *pl != nil {
			switch cmp := bt.cmp(val, (*pl).val); {
			case cmp < 0:
				pl = &(*pl).left
		   	case cmp > 0:
				pl = &(*pl).right
			default:
				return pl
			}
		}
		return pl
	}

	// Insert chèn val vào bt nếu chưa có,
	// và báo cáo liệu nó có được chèn hay không.
	func (bt *Tree[T]) Insert(val T) bool {
		pl := bt.find(val)
		if *pl != nil {
			return false
		}
		*pl = &node[T]{val: val}
		return true
	}
`}}

Mỗi nút trong cây chứa một giá trị của tham số kiểu `T`. Khi cây được khởi tạo với một đối số kiểu cụ thể, các giá trị của kiểu đó sẽ được lưu trực tiếp trong các nút. Chúng sẽ không được lưu như các kiểu interface.

Đây là cách sử dụng tham số kiểu hợp lý vì cấu trúc dữ liệu `Tree`, bao gồm code trong các method, phần lớn độc lập với kiểu phần tử `T`.

Cấu trúc dữ liệu `Tree` cần biết cách so sánh các giá trị của kiểu phần tử `T`; nó sử dụng hàm so sánh được truyền vào. Bạn có thể thấy điều này ở dòng thứ tư của method `find`, trong lời gọi `bt.cmp`. Ngoài điều đó, tham số kiểu không quan trọng chút nào.

### Với tham số kiểu, ưu tiên hàm hơn method

Ví dụ `Tree` minh họa một hướng dẫn chung khác: khi bạn cần thứ gì đó như hàm so sánh, ưu tiên hàm hơn method.

Chúng ta có thể đã định nghĩa kiểu `Tree` sao cho kiểu phần tử phải có method `Compare` hoặc `Less`. Điều này sẽ được thực hiện bằng cách viết một ràng buộc yêu cầu method đó, có nghĩa là bất kỳ đối số kiểu nào được dùng để khởi tạo kiểu `Tree` sẽ cần có method đó.

Hệ quả là bất kỳ ai muốn dùng `Tree` với kiểu dữ liệu đơn giản như `int` sẽ phải định nghĩa kiểu số nguyên riêng của họ và viết method so sánh riêng. Nếu chúng ta định nghĩa `Tree` để nhận hàm so sánh, như trong code trình bày ở trên, thì việc truyền hàm mong muốn vào rất dễ dàng. Việc viết hàm so sánh đó cũng dễ dàng như viết method.

Nếu kiểu phần tử `Tree` tình cờ đã có method `Compare`, thì chúng ta có thể đơn giản sử dụng biểu thức method như `ElementType.Compare` làm hàm so sánh.

Nói cách khác, việc biến một method thành hàm đơn giản hơn nhiều so với việc thêm một method vào kiểu. Vì vậy, với các kiểu dữ liệu đa dụng, ưu tiên hàm hơn là viết ràng buộc yêu cầu method.

### Cài đặt một method chung

Một trường hợp khác tham số kiểu có thể hữu ích là khi các kiểu khác nhau cần cài đặt một method chung, và các cài đặt cho các kiểu khác nhau trông giống nhau.

Ví dụ, hãy xem xét `sort.Interface` của thư viện chuẩn. Nó yêu cầu một kiểu cài đặt ba method: `Len`, `Swap` và `Less`.

Đây là ví dụ về kiểu generic `SliceFn` cài đặt `sort.Interface` cho bất kỳ kiểu slice nào:

{{raw `
	// SliceFn cài đặt sort.Interface cho một slice của T.
	type SliceFn[T any] struct {
		s    []T
		less func(T, T) bool
	}

	func (s SliceFn[T]) Len() int {
		return len(s.s)
	}
	func (s SliceFn[T]) Swap(i, j int) {
		s.s[i], s.s[j] = s.s[j], s.s[i]
	}
	func (s SliceFn[T]) Less(i, j int) bool {
		return s.less(s.s[i], s.s[j])
	}
`}}

Với bất kỳ kiểu slice nào, các method `Len` và `Swap` hoàn toàn giống nhau. Method `Less` yêu cầu một phép so sánh, đó là phần `Fn` của tên `SliceFn`. Như trong ví dụ `Tree` trước đó, chúng ta sẽ truyền vào một hàm khi tạo `SliceFn`.

Đây là cách sử dụng `SliceFn` để sắp xếp bất kỳ slice nào bằng hàm so sánh:

{{raw `
	// SortFn sắp xếp s tại chỗ sử dụng hàm so sánh.
	func SortFn[T any](s []T, less func(T, T) bool) {
		sort.Sort(SliceFn[T]{s, less})
	}
`}}

Điều này tương tự với hàm `sort.Slice` của thư viện chuẩn, nhưng hàm so sánh được viết sử dụng các giá trị thay vì các chỉ số slice.

Sử dụng tham số kiểu cho loại code này phù hợp vì các method trông hoàn toàn giống nhau với tất cả các kiểu slice.

(Tôi cần đề cập rằng Go 1.19, không phải 1.18, rất có thể sẽ bao gồm một hàm generic để sắp xếp một slice sử dụng hàm so sánh, và hàm generic đó rất có thể sẽ không sử dụng `sort.Interface`. Xem [đề xuất #47619](/issue/47619). Nhưng điểm chung vẫn đúng ngay cả khi ví dụ cụ thể này rất có thể sẽ không hữu ích: hợp lý khi sử dụng tham số kiểu khi bạn cần cài đặt các method trông giống nhau với tất cả các kiểu liên quan.)

## Khi nào tham số kiểu không hữu ích?

Bây giờ hãy nói về mặt kia của câu hỏi: khi nào không nên dùng tham số kiểu.

### Đừng thay thế kiểu interface bằng tham số kiểu

Như chúng ta đều biết, Go có kiểu interface. Kiểu interface cho phép một loại lập trình generic.

Ví dụ, interface `io.Reader` được sử dụng rộng rãi cung cấp cơ chế generic để đọc dữ liệu từ bất kỳ giá trị nào chứa thông tin (ví dụ, tệp) hoặc tạo ra thông tin (ví dụ, bộ tạo số ngẫu nhiên). Nếu tất cả những gì bạn cần làm với một giá trị của một kiểu nào đó là gọi một method trên giá trị đó, hãy dùng kiểu interface, không phải tham số kiểu. `io.Reader` dễ đọc, hiệu quả và hiệu quả. Không cần dùng tham số kiểu để đọc dữ liệu từ một giá trị bằng cách gọi method `Read`.

Ví dụ, có thể hấp dẫn khi thay đổi chữ ký hàm đầu tiên ở đây, chỉ dùng kiểu interface, thành phiên bản thứ hai dùng tham số kiểu.

{{raw `
	func ReadSome(r io.Reader) ([]byte, error)

	func ReadSome[T io.Reader](r T) ([]byte, error)
`}}

Đừng thực hiện loại thay đổi đó. Bỏ tham số kiểu làm cho hàm dễ viết hơn, dễ đọc hơn, và thời gian thực thi sẽ có thể như nhau.

Đáng nhấn mạnh điểm cuối cùng. Mặc dù có thể cài đặt generics theo nhiều cách khác nhau, và các cài đặt sẽ thay đổi và cải thiện theo thời gian, cài đặt được dùng trong Go 1.18 sẽ trong nhiều trường hợp xử lý các giá trị có kiểu là tham số kiểu giống như các giá trị có kiểu là kiểu interface. Điều này có nghĩa là việc sử dụng tham số kiểu thường không nhanh hơn sử dụng kiểu interface. Vì vậy, đừng chuyển từ kiểu interface sang tham số kiểu chỉ vì tốc độ, vì nó có thể không chạy nhanh hơn.

### Đừng dùng tham số kiểu nếu các cài đặt method khác nhau

Khi quyết định sử dụng tham số kiểu hay kiểu interface, hãy xem xét việc cài đặt các method. Trước đó chúng ta đã nói rằng nếu cài đặt của một method giống nhau với tất cả các kiểu, hãy dùng tham số kiểu. Ngược lại, nếu cài đặt khác nhau với từng kiểu, thì hãy dùng kiểu interface và viết các cài đặt method khác nhau, đừng dùng tham số kiểu.

Ví dụ, cài đặt `Read` từ một tệp hoàn toàn khác với cài đặt `Read` từ bộ tạo số ngẫu nhiên. Điều đó có nghĩa là chúng ta nên viết hai method `Read` khác nhau, và dùng kiểu interface như `io.Reader`.

### Dùng reflection khi thích hợp

Go có [reflection lúc runtime](https://pkg.go.dev/reflect). Reflection cho phép một loại lập trình generic, ở chỗ nó cho phép bạn viết code hoạt động với bất kỳ kiểu nào.

Nếu một thao tác phải hỗ trợ ngay cả các kiểu không có method (vì vậy kiểu interface không giúp ích), và nếu thao tác khác nhau với từng kiểu (vì vậy tham số kiểu không phù hợp), hãy dùng reflection.

Một ví dụ là gói [encoding/json](https://pkg.go.dev/encoding/json). Chúng ta không muốn yêu cầu mọi kiểu chúng ta encode phải có method `MarshalJSON`, vì vậy chúng ta không thể dùng kiểu interface. Nhưng việc encode kiểu interface hoàn toàn khác với encode kiểu struct, vì vậy chúng ta không nên dùng tham số kiểu. Thay vào đó, gói sử dụng reflection. Code không đơn giản, nhưng nó hoạt động. Để biết chi tiết, hãy xem [source code](/src/encoding/json/encode.go).

## Một hướng dẫn đơn giản

Tóm lại, cuộc thảo luận về khi nào nên dùng generics có thể được rút gọn thành một hướng dẫn đơn giản.

Nếu bạn thấy mình viết đúng cùng một đoạn code nhiều lần, nơi sự khác biệt duy nhất giữa các bản sao là code sử dụng các kiểu khác nhau, hãy cân nhắc xem bạn có thể dùng tham số kiểu không.

Nói cách khác, bạn nên tránh dùng tham số kiểu cho đến khi bạn nhận thấy rằng bạn sắp viết đúng cùng một đoạn code nhiều lần.
