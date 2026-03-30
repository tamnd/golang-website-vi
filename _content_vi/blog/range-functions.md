---
title: Range Qua Các Kiểu Hàm
date: 2024-08-20
by:
- Ian Lance Taylor
tags:
- iterators
summary: Mô tả về range qua các kiểu hàm, một tính năng mới trong Go 1.23.
template: true
---

## Giới thiệu

Đây là phiên bản bài viết của bài nói chuyện của tôi tại GopherCon 2024.

{{video "https://www.youtube.com/embed/i9zwUT9dlVc"}}

Range qua các kiểu hàm là một tính năng ngôn ngữ mới trong bản phát hành Go 1.23.
Bài viết này sẽ giải thích tại sao chúng tôi thêm tính năng mới này, nó
chính xác là gì, và cách sử dụng nó.

## Tại sao?

Kể từ Go 1.18, chúng ta đã có khả năng viết các kiểu container generic mới
trong Go.
Ví dụ, hãy xem xét kiểu `Set` rất đơn giản này, một kiểu generic
được xây dựng trên nền tảng của một map.

```
// Set holds a set of elements.
type Set[E comparable] struct {
	m map[E]struct{}
}

// New returns a new [Set].
func New[E comparable]() *Set[E] {
	return &Set[E]{m: make(map[E]struct{})}
}
```

Tự nhiên một kiểu set có cách để thêm các phần tử và kiểm tra
xem các phần tử có tồn tại không. Chi tiết ở đây không quan trọng.

```
// Add adds an element to a set.
func (s *Set[E]) Add(v E) {
	s.m[v] = struct{}{}
}

// Contains reports whether an element is in a set.
func (s *Set[E]) Contains(v E) bool {
	_, ok := s.m[v]
	return ok
}
```

Và trong số những thứ khác, chúng ta sẽ muốn có một hàm trả về hợp của
hai set.

```
// Union returns the union of two sets.
func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
	r := New[E]()
	// Note for/range over internal Set field m.
	// We are looping over the maps in s1 and s2.
	for v := range s1.m {
		r.Add(v)
	}
	for v := range s2.m {
		r.Add(v)
	}
	return r
}
```

Hãy xem cài đặt của hàm `Union` trong một phút.
Để tính hợp của hai set, chúng ta cần một cách để lấy tất cả
các phần tử trong mỗi set.
Trong mã này, chúng ta sử dụng câu lệnh for/range trên một trường không xuất khẩu của
kiểu set.
Điều đó chỉ hoạt động nếu hàm `Union` được định nghĩa trong gói set.

Nhưng có nhiều lý do khiến ai đó có thể muốn lặp qua tất cả
các phần tử trong một set.
Gói set này phải cung cấp một số cách để người dùng làm điều đó.

Làm thế nào điều đó nên hoạt động?

### Đẩy các phần tử Set

Một cách tiếp cận là cung cấp một phương thức `Set` nhận một hàm, và
gọi hàm đó với mọi phần tử trong Set.
Chúng ta sẽ gọi điều này là `Push`, vì `Set` đẩy mọi giá trị cho
hàm.
Ở đây, nếu hàm trả về false, chúng ta ngừng gọi nó.

```
func (s *Set[E]) Push(f func(E) bool) {
	for v := range s.m {
		if !f(v) {
			return
		}
	}
}
```

Trong thư viện chuẩn Go, chúng ta thấy mẫu chung này được sử dụng cho các trường hợp
như phương thức [`sync.Map.Range`](https://pkg.go.dev/sync#Map.Range),
hàm [`flag.Visit`](https://pkg.go.dev/flag#Visit), và
hàm [`filepath.Walk`](https://pkg.go.dev/path/filepath#Walk).
Đây là một mẫu chung, không phải mẫu chính xác; trong thực tế, không có
ba ví dụ nào trong số đó hoạt động hoàn toàn theo cách giống nhau.

Đây là cách sử dụng phương thức `Push` để in tất cả các
phần tử của một set trông như thế nào: bạn gọi `Push` với một hàm thực hiện những gì bạn
muốn với phần tử.

```
func PrintAllElementsPush[E comparable](s *Set[E]) {
	s.Push(func(v E) bool {
		fmt.Println(v)
		return true
	})
}
```

### Kéo các phần tử Set

Một cách tiếp cận khác để lặp qua các phần tử của một `Set` là trả về
một hàm.
Mỗi lần hàm được gọi, nó sẽ trả về một giá trị từ `Set`,
cùng với một boolean cho biết giá trị đó có hợp lệ không.
Kết quả boolean sẽ là false khi vòng lặp đã đi qua tất cả
các phần tử.
Trong trường hợp này, chúng ta cũng cần một hàm dừng có thể được gọi khi không
cần thêm giá trị nữa.

Cài đặt này sử dụng một cặp channel, một cái cho các giá trị trong
set và một cái để dừng trả về giá trị.
Chúng ta sử dụng một goroutine để gửi giá trị trên channel.
Hàm `next` trả về một phần tử từ set bằng cách đọc từ
channel phần tử, và hàm `stop` báo cho goroutine
thoát bằng cách đóng channel dừng.
Chúng ta cần hàm `stop` để đảm bảo goroutine thoát khi
không cần thêm giá trị nữa.

{{raw `
	// Pull returns a next function that returns each
	// element of s with a bool for whether the value
	// is valid. The stop function should be called
	// when finished calling the next function.
	func (s *Set[E]) Pull() (func() (E, bool), func()) {
		ch := make(chan E)
		stopCh := make(chan bool)

		go func() {
			defer close(ch)
			for v := range s.m {
				select {
				case ch <- v:
				case <-stopCh:
					return
				}
			}
		}()

		next := func() (E, bool) {
			v, ok := <-ch
			return v, ok
		}

		stop := func() {
			close(stopCh)
		}

		return next, stop
	}
`}}

Không có gì trong thư viện chuẩn hoạt động chính xác theo cách này. Cả
[`runtime.CallersFrames`](https://pkg.go.dev/runtime#CallersFrames)
và
[`reflect.Value.MapRange`](https://pkg.go.dev/reflect#Value.MapRange)
đều tương tự, mặc dù chúng trả về các giá trị với các phương thức thay vì
trả về các hàm trực tiếp.

Đây là cách sử dụng phương thức `Pull` để in tất cả các
phần tử của một `Set` trông như thế nào.
Bạn gọi `Pull` để lấy một hàm, và bạn gọi hàm đó lặp lại trong một vòng lặp for.

```
func PrintAllElementsPull[E comparable](s *Set[E]) {
	next, stop := s.Pull()
	defer stop()
	for v, ok := next(); ok; v, ok = next() {
		fmt.Println(v)
	}
}
```

## Chuẩn hóa cách tiếp cận

Chúng ta đã thấy hai cách tiếp cận khác nhau để lặp qua tất cả các
phần tử của một set.
Các gói Go khác nhau sử dụng các cách tiếp cận này và một số cách khác.
Điều đó có nghĩa là khi bạn bắt đầu sử dụng một gói container Go mới, bạn
có thể phải học một cơ chế lặp mới.
Điều đó cũng có nghĩa là chúng ta không thể viết một hàm hoạt động với nhiều
loại container khác nhau, vì các kiểu container sẽ xử lý
việc lặp theo những cách khác nhau.

Chúng ta muốn cải thiện hệ sinh thái Go bằng cách phát triển các cách tiếp cận chuẩn
để lặp qua các container.

### Iterators

Tất nhiên, đây là một vấn đề phát sinh trong nhiều ngôn ngữ lập trình.

[Cuốn sách Design Patterns](https://en.wikipedia.org/wiki/Design_Patterns) phổ biến, được xuất bản lần đầu
vào năm 1994, mô tả điều này là mẫu iterator.
Bạn sử dụng một iterator để "cung cấp một cách để truy cập các phần tử của
một đối tượng tổng hợp một cách tuần tự mà không cần phơi bày cách biểu diễn bên dưới của nó."
Điều mà câu trích dẫn này gọi là đối tượng tổng hợp là những gì tôi gọi là
container.
Một đối tượng tổng hợp, hay container, chỉ là một giá trị chứa các giá trị khác,
như kiểu `Set` chúng ta đã thảo luận.

Giống như nhiều ý tưởng trong lập trình, các iterator có nguồn gốc từ
[ngôn ngữ CLU](https://en.wikipedia.org/wiki/CLU_(programming_language)) của Barbara Liskov,
được phát triển vào những năm 1970.

Ngày nay nhiều ngôn ngữ phổ biến cung cấp các iterator theo cách này hay cách khác,
bao gồm, trong số những ngôn ngữ khác, C++, Java, Javascript, Python và Rust.

Tuy nhiên, Go trước phiên bản 1.23 thì không.

### For/range

Như chúng ta đều biết, Go có các kiểu container được tích hợp vào
ngôn ngữ: slice, mảng và map.
Và nó có cách để truy cập các phần tử của những giá trị đó mà không cần
phơi bày cách biểu diễn bên dưới: câu lệnh for/range.
Câu lệnh for/range hoạt động cho các kiểu container tích hợp của Go (và
cũng cho chuỗi, channel, và kể từ Go 1.22, int).

Câu lệnh for/range là lặp, nhưng nó không phải là các iterator như chúng
xuất hiện trong các ngôn ngữ phổ biến ngày nay.
Tuy nhiên, sẽ rất tốt nếu có thể sử dụng for/range để lặp qua
một container do người dùng định nghĩa như kiểu `Set`.

Tuy nhiên, Go trước phiên bản 1.23 không hỗ trợ điều này.

### Cải tiến trong bản phát hành này

Đối với Go 1.23, chúng tôi đã quyết định hỗ trợ cả for/range trên
các kiểu container do người dùng định nghĩa, và một dạng iterator chuẩn hóa.

Chúng tôi đã mở rộng câu lệnh for/range để hỗ trợ range qua các kiểu
hàm.
Chúng ta sẽ thấy bên dưới điều này giúp lặp qua các container do người dùng định nghĩa như thế nào.

Chúng tôi cũng đã thêm các kiểu và hàm thư viện chuẩn để hỗ trợ sử dụng
các kiểu hàm như các iterator.
Một định nghĩa chuẩn về iterator cho phép chúng ta viết các hàm hoạt động
trơn tru với các kiểu container khác nhau.

### Range qua (một số) kiểu hàm

Câu lệnh for/range được cải tiến không hỗ trợ các kiểu hàm tùy ý.
Kể từ Go 1.23, nó hỗ trợ range qua các hàm nhận một
đối số duy nhất.
Đối số duy nhất phải là một hàm nhận từ không đến hai
đối số và trả về bool; theo quy ước, chúng ta gọi nó là hàm yield.

```
func(yield func() bool)

func(yield func(V) bool)

func(yield func(K, V) bool)
```

Khi chúng ta nói về một iterator trong Go, chúng ta có nghĩa là một hàm có một trong
ba kiểu này.
Như chúng ta sẽ thảo luận bên dưới, có một loại iterator khác trong
thư viện chuẩn: pull iterator.
Khi cần phân biệt giữa các iterator chuẩn và
pull iterator, chúng ta gọi các iterator chuẩn là push iterator.
Đó là vì, như chúng ta sẽ thấy, chúng đẩy ra một chuỗi giá trị bằng cách
gọi hàm yield.

### Iterator chuẩn (push)

Để làm cho các iterator dễ sử dụng hơn, gói iter trong thư viện chuẩn mới
định nghĩa hai kiểu: `Seq` và `Seq2`.
Đây là tên cho các kiểu hàm iterator, các kiểu có thể được
sử dụng với câu lệnh for/range.
Tên `Seq` là viết tắt của sequence (chuỗi), vì các iterator lặp qua một
chuỗi giá trị.

```
package iter

type Seq[V any] func(yield func(V) bool)

type Seq2[K, V any] func(yield func(K, V) bool)

// for now, no Seq0
```

Sự khác biệt giữa `Seq` và `Seq2` chỉ là `Seq2` là một
chuỗi các cặp, chẳng hạn như key và value từ một map.
Trong bài viết này, chúng ta sẽ tập trung vào `Seq` cho đơn giản, nhưng hầu hết những gì chúng ta
nói cũng áp dụng cho `Seq2`.

Cách dễ nhất để giải thích cách iterator hoạt động là thông qua một ví dụ.
Ở đây, phương thức `All` của `Set` trả về một hàm.
Kiểu trả về của `All` là `iter.Seq[E]`, vì vậy chúng ta biết nó trả về
một iterator.

```
// All is an iterator over the elements of s.
func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}
```

Bản thân hàm iterator nhận một hàm khác, hàm yield,
làm đối số.
Iterator gọi hàm yield với mọi giá trị trong set.
Trong trường hợp này, iterator, hàm được trả về bởi `Set.All`, giống
như hàm `Set.Push` chúng ta đã thấy trước đó.

Điều này cho thấy cách các iterator hoạt động: đối với một số chuỗi giá trị, chúng gọi
một hàm yield với mỗi giá trị trong chuỗi.
Nếu hàm yield trả về false, không cần thêm giá trị nữa, và
iterator có thể trả về, thực hiện bất kỳ việc dọn dẹp nào có thể được yêu cầu.
Nếu hàm yield không bao giờ trả về false, iterator có thể trả về sau khi gọi yield với tất cả các giá trị trong chuỗi.

Đó là cách chúng hoạt động, nhưng hãy thừa nhận rằng lần đầu tiên bạn
nhìn thấy một trong những thứ này, phản ứng đầu tiên của bạn có thể là "có rất nhiều
hàm bay xung quanh ở đây."
Bạn không sai về điều đó.
Hãy tập trung vào hai điều.

Điều đầu tiên là một khi bạn đi qua dòng đầu tiên của mã hàm này,
cài đặt thực sự của iterator khá đơn giản: gọi
yield với mọi phần tử của set, dừng nếu yield trả về false.

```
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
```

Điều thứ hai là việc sử dụng điều này thực sự rất dễ dàng.
Bạn gọi `s.All` để lấy một iterator, rồi sử dụng for/range để
lặp qua tất cả các phần tử trong `s`.
Câu lệnh for/range hỗ trợ bất kỳ iterator nào, và điều này cho thấy cách dễ dàng
sử dụng như thế nào.

```
func PrintAllElements[E comparable](s *Set[E]) {
	for v := range s.All() {
		fmt.Println(v)
	}
}
```

Trong loại mã này, `s.All` là một phương thức trả về một hàm.
Chúng ta đang gọi `s.All`, rồi sử dụng for/range để range qua
hàm mà nó trả về.
Trong trường hợp này, chúng ta có thể đã làm cho `Set.All` là một hàm iterator
chính nó, thay vì để nó trả về một hàm iterator.
Tuy nhiên, trong một số trường hợp điều đó sẽ không hoạt động, chẳng hạn như nếu hàm
trả về iterator cần nhận một đối số, hoặc cần thực hiện một số
công việc cài đặt.
Về mặt quy ước, chúng tôi khuyến khích tất cả các kiểu container cung cấp
một phương thức `All` trả về một iterator, để các lập trình viên không
phải nhớ liệu có nên range trực tiếp trên `All` hay gọi `All` để lấy một giá trị để range.
Họ luôn có thể làm theo cách sau.

Nếu bạn nghĩ về nó, bạn sẽ thấy rằng trình biên dịch phải điều chỉnh
vòng lặp để tạo một hàm yield để truyền cho iterator được trả về
bởi `s.All`.
Có khá nhiều sự phức tạp trong trình biên dịch và runtime Go để
làm điều này hiệu quả, và để xử lý đúng những thứ như `break` hoặc
`panic` trong vòng lặp.
Chúng ta sẽ không đề cập đến bất kỳ điều nào trong số đó trong bài viết này.
May mắn thay, các chi tiết cài đặt không quan trọng khi thực sự sử dụng tính năng này.

### Pull iterator

Chúng ta đã thấy cách sử dụng các iterator trong vòng lặp for/range.
Nhưng một vòng lặp đơn giản không phải là cách duy nhất để sử dụng một iterator.
Ví dụ, đôi khi chúng ta cần lặp qua hai container song song.
Làm thế nào để làm điều đó?

Câu trả lời là chúng ta sử dụng một loại iterator khác: pull
iterator.
Chúng ta đã thấy rằng một iterator chuẩn, còn được gọi là push iterator, là
một hàm nhận hàm yield làm đối số và đẩy mỗi
giá trị trong chuỗi bằng cách gọi hàm yield.

Một pull iterator hoạt động theo hướng ngược lại: nó là một hàm được
viết sao cho mỗi lần bạn gọi nó, nó trả về giá trị tiếp theo trong
chuỗi.

Chúng ta sẽ nhắc lại sự khác biệt giữa hai loại iterator để giúp
bạn nhớ:
- Một push iterator đẩy mỗi giá trị trong chuỗi đến một hàm yield.
  Push iterator là các iterator chuẩn trong thư viện chuẩn Go,
  và được hỗ trợ trực tiếp bởi câu lệnh for/range.
- Một pull iterator hoạt động theo hướng ngược lại.
  Mỗi lần bạn gọi một pull iterator, nó kéo một giá trị khác từ một
  chuỗi và trả về nó.
  Pull iterator _không_ được hỗ trợ trực tiếp bởi câu lệnh for/range;
  tuy nhiên, khá đơn giản để viết một câu lệnh for thông thường
  lặp qua một pull iterator.
  Thực ra, chúng ta đã thấy một ví dụ trước đó khi chúng ta xem xét sử dụng
  phương thức `Set.Pull`.

Bạn có thể tự viết một pull iterator, nhưng thường thì bạn không cần phải làm vậy.
Hàm thư viện chuẩn mới
[`iter.Pull`](https://pkg.go.dev/iter#Pull) nhận một iterator chuẩn,
tức là một hàm là push iterator, và trả về một cặp hàm.
Hàm đầu tiên là pull iterator: một hàm trả về giá trị tiếp theo
trong chuỗi mỗi lần nó được gọi.
Hàm thứ hai là hàm dừng nên được gọi khi chúng ta xong
với pull iterator.
Điều này giống như phương thức `Set.Pull` chúng ta đã thấy trước đó.

Hàm đầu tiên được trả về bởi `iter.Pull`, pull iterator, trả về
một giá trị và một boolean cho biết liệu giá trị đó có hợp lệ không.
Boolean sẽ là false ở cuối chuỗi.

`iter.Pull` trả về một hàm dừng trong trường hợp chúng ta không đọc hết
chuỗi đến cuối.
Trong trường hợp tổng quát, push iterator, đối số cho `iter.Pull`,
có thể khởi động các goroutine, hoặc xây dựng các cấu trúc dữ liệu mới cần được
dọn dẹp khi lặp hoàn thành.
Push iterator sẽ thực hiện bất kỳ việc dọn dẹp nào khi hàm yield trả về
false, có nghĩa là không cần thêm giá trị nào.
Khi được sử dụng với câu lệnh for/range, câu lệnh for/range sẽ
đảm bảo rằng nếu vòng lặp thoát sớm, thông qua câu lệnh `break` hoặc
vì bất kỳ lý do nào khác, thì hàm yield sẽ trả về false.
Với pull iterator, mặt khác, không có cách nào để buộc
hàm yield trả về false, vì vậy hàm dừng là cần thiết.

Một cách khác để nói điều này là gọi hàm dừng sẽ khiến
hàm yield trả về false khi nó được gọi bởi push
iterator.

Về mặt kỹ thuật, bạn không cần gọi hàm dừng nếu pull
iterator trả về false để cho biết nó đã đến cuối
chuỗi, nhưng thường đơn giản hơn khi luôn gọi nó.

Đây là một ví dụ sử dụng pull iterator để đi qua hai
chuỗi song song.
Hàm này báo cáo liệu hai chuỗi tùy ý có chứa cùng
các phần tử theo cùng thứ tự không.

```
// EqSeq reports whether two iterators contain the same
// elements in the same order.
func EqSeq[E comparable](s1, s2 iter.Seq[E]) bool {
	next1, stop1 := iter.Pull(s1)
	defer stop1()
	next2, stop2 := iter.Pull(s2)
	defer stop2()
	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if !ok1 {
			return !ok2
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
}
```

Hàm sử dụng `iter.Pull` để chuyển đổi hai push iterator, `s1`
và `s2`, thành pull iterator.
Nó sử dụng các câu lệnh `defer` để đảm bảo rằng các pull iterator được
dừng khi chúng ta xong với chúng.

Sau đó mã lặp, gọi các pull iterator để lấy giá trị.
Nếu chuỗi đầu tiên đã xong, nó trả về true nếu chuỗi thứ hai
cũng xong, hoặc false nếu không.
Nếu các giá trị khác nhau, nó trả về false.
Sau đó nó lặp để kéo hai giá trị tiếp theo.

Như với push iterator, có một số phức tạp trong runtime Go để
làm cho pull iterator hiệu quả, nhưng điều này không ảnh hưởng đến mã thực sự
sử dụng hàm `iter.Pull`.

## Lặp trên các iterator

Bây giờ bạn biết tất cả mọi thứ về range qua các kiểu hàm
và về các iterator.
Hy vọng bạn thích sử dụng chúng!

Tuy nhiên, có thêm một vài điều đáng đề cập.

### Adapter

Một ưu điểm của định nghĩa chuẩn về iterator là khả năng
viết các hàm adapter chuẩn sử dụng chúng.

Ví dụ, đây là một hàm lọc một chuỗi giá trị,
trả về một chuỗi mới.
Hàm `Filter` này nhận một iterator làm đối số và trả về
một iterator mới.
Đối số kia là một hàm lọc quyết định giá trị nào
nên có trong iterator mới mà `Filter` trả về.

```
// Filter returns a sequence that contains the elements
// of s for which f returns true.
func Filter[V any](f func(V) bool, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
```

Như với ví dụ trước đó, các chữ ký hàm trông phức tạp
khi bạn nhìn thấy chúng lần đầu.
Một khi bạn đi qua các chữ ký, cài đặt là
đơn giản.

```
		for v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
```

Mã range qua iterator đầu vào, kiểm tra hàm lọc,
và gọi yield với các giá trị nên đi vào output
iterator.

Chúng ta sẽ cho thấy một ví dụ sử dụng `Filter` bên dưới.

(Hiện tại không có phiên bản `Filter` nào trong thư viện chuẩn Go, nhưng
có thể được thêm vào trong các bản phát hành tương lai.)

### Cây nhị phân

Như một ví dụ về cách thuận tiện push iterator có thể là để lặp qua một
kiểu container, hãy xem xét kiểu cây nhị phân đơn giản này.

```
// Tree is a binary tree.
type Tree[E any] struct {
	val         E
	left, right *Tree[E]
}
```

Chúng ta sẽ không hiển thị mã để chèn các giá trị vào cây, nhưng tự nhiên
nên có cách nào đó để range qua tất cả các giá trị trong cây.

Hóa ra mã iterator dễ viết hơn nếu nó trả về
bool.
Vì các kiểu hàm được hỗ trợ bởi for/range không trả về gì,
phương thức `All` ở đây trả về một hàm literal nhỏ gọi
bản thân iterator, ở đây được gọi là `push`, và bỏ qua kết quả bool.

```
// All returns an iterator over the values in t.
func (t *Tree[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		t.push(yield)
	}
}

// push pushes all elements to the yield function.
func (t *Tree[E]) push(yield func(E) bool) bool {
	if t == nil {
		return true
	}
	return t.left.push(yield) &&
		yield(t.val) &&
		t.right.push(yield)
}
```

Phương thức `push` sử dụng đệ quy để đi qua toàn bộ cây, gọi
yield trên mỗi phần tử.
Nếu hàm yield trả về false, phương thức trả về false tất cả cách
lên stack.
Ngược lại, nó chỉ trả về sau khi việc lặp hoàn thành.

Điều này cho thấy cách tiếp cận iterator này dễ dàng như thế nào để
lặp qua ngay cả các cấu trúc dữ liệu phức tạp.
Không cần duy trì một stack riêng để ghi lại vị trí
trong cây; chúng ta chỉ có thể sử dụng call stack của goroutine để làm điều đó
cho chúng ta.

### Các hàm iterator mới.

Cũng mới trong Go 1.23 là các hàm trong các gói slices và maps
hoạt động với các iterator.

Đây là các hàm mới trong gói slices.
`All` và `Values` là các hàm trả về iterator qua các
phần tử của một slice.
`Collect` lấy các giá trị từ một iterator và trả về một slice
chứa những giá trị đó.
Xem docs cho những cái còn lại.

- [`All([]E) iter.Seq2[int, E]`](https://pkg.go.dev/slices#All)
- [`Values([]E) iter.Seq[E]`](https://pkg.go.dev/slices#Values)
- [`Collect(iter.Seq[E]) []E`](https://pkg.go.dev/slices#Collect)
- [`AppendSeq([]E, iter.Seq[E]) []E`](https://pkg.go.dev/slices#AppendSeq)
- [`Backward([]E) iter.Seq2[int, E]`](https://pkg.go.dev/slices#Backward)
- [`Sorted(iter.Seq[E]) []E`](https://pkg.go.dev/slices#Sorted)
- [`SortedFunc(iter.Seq[E], func(E, E) int) []E`](https://pkg.go.dev/slices#SortedFunc)
- [`SortedStableFunc(iter.Seq[E], func(E, E) int) []E`](https://pkg.go.dev/slices#SortedStableFunc)
- [`Repeat([]E, int) []E`](https://pkg.go.dev/slices#Repeat)
- [`Chunk([]E, int) iter.Seq([]E)`](https://pkg.go.dev/slices#Chunk)

Đây là các hàm mới trong gói maps.
`All`, `Keys` và `Values` trả về iterator qua nội dung của map.
`Collect` lấy các key và value từ một iterator và trả về một
map mới.

- [`All(map[K]V) iter.Seq2[K, V]`](https://pkg.go.dev/maps#All)
- [`Keys(map[K]V) iter.Seq[K]`](https://pkg.go.dev/maps#Keys)
- [`Values(map[K]V) iter.Seq[V]`](https://pkg.go.dev/maps#Values)
- [`Collect(iter.Seq2[K, V]) map[K, V]`](https://pkg.go.dev/maps#Collect)
- [`Insert(map[K, V], iter.Seq2[K, V])`](https://pkg.go.dev/maps#Insert)

### Ví dụ iterator thư viện chuẩn

Đây là một ví dụ về cách bạn có thể sử dụng các hàm mới này cùng với
hàm `Filter` chúng ta đã thấy trước đó.
Hàm này nhận một map từ int sang string và trả về một slice
chỉ chứa các giá trị trong map dài hơn một số đối số `n` nào đó.

```
// LongStrings returns a slice of just the values
// in m whose length is n or more.
func LongStrings(m map[int]string, n int) []string {
	isLong := func(s string) bool {
		return len(s) >= n
	}
	return slices.Collect(Filter(isLong, maps.Values(m)))
}
```

Hàm `maps.Values` trả về một iterator qua các giá trị trong `m`.
`Filter` đọc iterator đó và trả về một iterator mới chỉ
chứa các chuỗi dài.
`slices.Collect` đọc từ iterator đó vào một slice mới.

Tất nhiên, bạn có thể viết một vòng lặp để làm điều này đủ dễ dàng, và trong
nhiều trường hợp một vòng lặp sẽ rõ ràng hơn.
Chúng tôi không muốn khuyến khích mọi người viết mã theo phong cách này mọi lúc.
Tuy nhiên, ưu điểm của việc sử dụng iterator là loại
hàm này hoạt động theo cùng một cách với bất kỳ chuỗi nào.
Trong ví dụ này, hãy chú ý cách Filter đang sử dụng map làm đầu vào và một
slice làm đầu ra, mà không phải thay đổi mã trong Filter
chút nào.

### Lặp qua các dòng trong một tệp

Mặc dù hầu hết các ví dụ chúng ta đã thấy liên quan đến container,
các iterator rất linh hoạt.

Hãy xem xét mã đơn giản này, không sử dụng iterator, để lặp qua
các dòng trong một byte slice.
Cách này dễ viết và khá hiệu quả.

```
	nl := []byte{'\n'}
	// Trim a trailing newline to avoid a final empty blank line.
	for _, line := range bytes.Split(bytes.TrimSuffix(data, nl), nl) {
		handleLine(line)
	}
```

Tuy nhiên, `bytes.Split` có phân bổ và trả về một slice của byte slice
để chứa các dòng.
Bộ gom rác sẽ phải làm một chút công việc để cuối cùng giải phóng
slice đó.

Đây là một hàm trả về một iterator qua các dòng của một byte slice nào đó.
Sau các chữ ký iterator thông thường, hàm khá đơn giản.
Chúng ta tiếp tục chọn các dòng ra từ dữ liệu cho đến khi không còn gì, và chúng ta
truyền mỗi dòng cho hàm yield.

```
// Lines returns an iterator over lines in data.
func Lines(data []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for len(data) > 0 {
			line, rest, _ := bytes.Cut(data, []byte{'\n'})
			if !yield(line) {
				return
			}
			data = rest
		}
	}
}
```

Bây giờ mã của chúng ta để lặp qua các dòng của một byte slice trông như thế này.

```
	for line := range Lines(data) {
		handleLine(line)
	}
```

Cách này cũng dễ viết như mã trước đó, và hiệu quả hơn một chút
vì nó không phải phân bổ một slice của các dòng.

### Truyền hàm cho push iterator

Với ví dụ cuối cùng của chúng ta, chúng ta sẽ thấy rằng bạn không cần phải sử dụng push
iterator trong câu lệnh range.

Trước đó chúng ta đã thấy hàm `PrintAllElements` in ra mỗi
phần tử của một set.
Đây là một cách khác để in tất cả các phần tử của một set: gọi `s.All`
để lấy một iterator, rồi truyền vào một hàm yield được viết tay.
Hàm yield này chỉ in một giá trị và trả về true.
Lưu ý rằng có hai lời gọi hàm ở đây: chúng ta gọi `s.All` để lấy
một iterator là chính nó một hàm, và chúng ta gọi hàm đó với
hàm yield được viết tay của chúng ta.


```
func PrintAllElements[E comparable](s *Set[E]) {
	s.All()(func(v E) bool {
		fmt.Println(v)
		return true
	})
}
```

Không có lý do cụ thể nào để viết mã theo cách này.
Đây chỉ là một ví dụ để cho thấy rằng hàm yield không có gì kỳ diệu.
Nó có thể là bất kỳ hàm nào bạn thích.

## Cập nhật go.mod

Ghi chú cuối cùng: mỗi module Go chỉ định phiên bản ngôn ngữ mà nó
sử dụng.
Điều đó có nghĩa là để sử dụng các tính năng ngôn ngữ mới trong một module hiện có, bạn
có thể cần cập nhật phiên bản đó.
Điều này đúng với tất cả các tính năng ngôn ngữ mới; đây không phải là điều
cụ thể đối với range qua các kiểu hàm.
Vì range qua các kiểu hàm là mới trong bản phát hành Go 1.23, sử dụng nó
đòi hỏi chỉ định ít nhất phiên bản ngôn ngữ Go 1.23.

Có (ít nhất) bốn cách để đặt phiên bản ngôn ngữ:
- Trên dòng lệnh, chạy `go get go@1.23` (hoặc `go mod edit -go=1.23`
  để chỉ chỉnh sửa chỉ thị `go`).
- Chỉnh sửa thủ công tệp `go.mod` và thay đổi dòng `go`.
- Giữ phiên bản ngôn ngữ cũ cho toàn bộ module, nhưng sử dụng
  build tag `//go:build go1.23` để cho phép sử dụng range qua các kiểu hàm
  trong một tệp cụ thể.
