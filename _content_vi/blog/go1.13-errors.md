---
title: Làm việc với lỗi trong Go 1.13
date: 2019-10-17
by:
- Damien Neil and Jonathan Amsterdam
tags:
- errors
- technical
summary: Cách dùng các interface và hàm xử lý lỗi mới trong Go 1.13.
---

## Giới thiệu

Cách Go xử lý [lỗi như giá trị](/blog/errors-are-values)
đã phục vụ chúng ta tốt trong thập kỷ qua. Mặc dù hỗ trợ của thư viện chuẩn
cho lỗi khá tối giản, chỉ có hai hàm `errors.New` và `fmt.Errorf`
tạo ra lỗi chỉ chứa thông báo, interface `error` built-in
cho phép lập trình viên Go thêm bất kỳ thông tin nào họ muốn. Tất cả những gì cần là
một kiểu triển khai phương thức `Error`:

	type QueryError struct {
		Query string
		Err   error
	}

	func (e *QueryError) Error() string { return e.Query + ": " + e.Err.Error() }

Các kiểu lỗi như thế này rất phổ biến, và thông tin chúng lưu trữ rất đa dạng,
từ timestamp đến tên file đến địa chỉ server. Thường thì thông tin đó
bao gồm một lỗi khác ở mức thấp hơn để cung cấp thêm ngữ cảnh.

Mẫu một lỗi chứa lỗi khác rất phổ biến trong code Go đến mức,
sau [thảo luận kéo dài](/issue/29934), Go 1.13 đã thêm
hỗ trợ rõ ràng cho nó. Bài viết này mô tả các bổ sung cho thư viện chuẩn
cung cấp hỗ trợ đó: ba hàm mới trong gói `errors`,
và một động từ định dạng mới cho `fmt.Errorf`.

Trước khi mô tả chi tiết các thay đổi, hãy xem lại cách lỗi được kiểm tra
và xây dựng trong các phiên bản ngôn ngữ trước đó.

## Lỗi trước Go 1.13

### Kiểm tra lỗi

Lỗi Go là giá trị. Chương trình đưa ra quyết định dựa trên các giá trị đó theo một vài
cách. Phổ biến nhất là so sánh lỗi với `nil` để xem một thao tác có
thất bại không.

	if err != nil {
		// có gì đó không ổn
	}

Đôi khi chúng ta so sánh lỗi với một giá trị _sentinel_ đã biết, để xem một lỗi cụ thể có xảy ra không.

	var ErrNotFound = errors.New("not found")

	if err == ErrNotFound {
		// không tìm thấy gì đó
	}

Một giá trị lỗi có thể thuộc bất kỳ kiểu nào thỏa mãn interface `error` được định nghĩa trong ngôn ngữ.
Chương trình có thể dùng type assertion hoặc type switch để xem giá trị lỗi như một kiểu cụ thể hơn.

	type NotFoundError struct {
		Name string
	}

	func (e *NotFoundError) Error() string { return e.Name + ": not found" }

	if e, ok := err.(*NotFoundError); ok {
		// e.Name không tìm thấy
	}

### Thêm thông tin

Thường thì một hàm chuyển lỗi lên call stack trong khi thêm thông tin
vào nó, như mô tả ngắn gọn về những gì đang xảy ra khi lỗi xảy ra. Một
cách đơn giản là xây dựng lỗi mới bao gồm văn bản của lỗi trước:

	if err != nil {
		return fmt.Errorf("decompress %v: %v", name, err)
	}

Tạo lỗi mới với `fmt.Errorf` bỏ đi mọi thứ từ lỗi gốc
ngoại trừ văn bản. Như chúng ta đã thấy ở trên với `QueryError`, đôi khi chúng ta muốn
định nghĩa một kiểu lỗi mới chứa lỗi bên dưới, giữ nguyên nó để
code kiểm tra. Đây là `QueryError` một lần nữa:

	type QueryError struct {
		Query string
		Err   error
	}

Chương trình có thể nhìn bên trong giá trị `*QueryError` để đưa ra quyết định dựa trên
lỗi bên dưới. Bạn đôi khi thấy điều này được gọi là "unwrap" lỗi.

	if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {
		// query thất bại vì vấn đề quyền truy cập
	}

Kiểu `os.PathError` trong thư viện chuẩn là một ví dụ khác về lỗi chứa lỗi khác.

## Lỗi trong Go 1.13

### Phương thức Unwrap

Go 1.13 giới thiệu các tính năng mới cho các gói thư viện chuẩn `errors` và `fmt`
để đơn giản hóa việc làm việc với lỗi chứa lỗi khác. Quan trọng nhất trong số này
là một quy ước hơn là một thay đổi: lỗi chứa lỗi khác có thể triển khai phương thức `Unwrap`
trả về lỗi bên dưới. Nếu `e1.Unwrap()` trả về `e2`, thì chúng ta nói rằng `e1` _bọc_ `e2`, và
rằng bạn có thể _unwrap_ `e1` để lấy `e2`.

Theo quy ước này, chúng ta có thể cung cấp cho kiểu `QueryError` ở trên một phương thức `Unwrap`
trả về lỗi mà nó chứa:

	func (e *QueryError) Unwrap() error { return e.Err }

Kết quả của việc unwrap một lỗi bản thân nó có thể có phương thức `Unwrap`; chúng ta gọi
chuỗi lỗi được tạo ra bởi việc unwrap lặp lại là _chuỗi lỗi_.

### Kiểm tra lỗi với Is và As

Gói `errors` trong Go 1.13 bao gồm hai hàm mới để kiểm tra lỗi: `Is` và `As`.

Hàm `errors.Is` so sánh lỗi với một giá trị.

	// Tương tự:
	//   if err == ErrNotFound { … }
	if errors.Is(err, ErrNotFound) {
		// không tìm thấy gì đó
	}

Hàm `As` kiểm tra xem lỗi có phải là một kiểu cụ thể không.

	// Tương tự:
	//   if e, ok := err.(*QueryError); ok { … }
	var e *QueryError
	// Lưu ý: *QueryError là kiểu của lỗi.
	if errors.As(err, &e) {
		// err là *QueryError, và e được gán giá trị của lỗi
	}

Trong trường hợp đơn giản nhất, hàm `errors.Is` hoạt động như một so sánh với
lỗi sentinel, và hàm `errors.As` hoạt động như một type assertion. Tuy nhiên khi
làm việc với các lỗi được bọc, các hàm này xem xét tất cả lỗi trong
chuỗi. Hãy xem lại ví dụ ở trên về việc unwrap `QueryError`
để kiểm tra lỗi bên dưới:

	if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {
		// query thất bại vì vấn đề quyền truy cập
	}

Dùng hàm `errors.Is`, chúng ta có thể viết:

	if errors.Is(err, ErrPermission) {
		// err, hoặc lỗi nào đó mà nó bọc, là vấn đề quyền truy cập
	}

Gói `errors` cũng bao gồm hàm `Unwrap` mới trả về
kết quả của việc gọi phương thức `Unwrap` của lỗi, hoặc `nil` khi lỗi không có
phương thức `Unwrap`. Tuy nhiên, thường tốt hơn là dùng `errors.Is` hoặc `errors.As`,
vì các hàm này sẽ kiểm tra toàn bộ chuỗi chỉ trong một lần gọi.

Lưu ý: mặc dù có vẻ lạ khi lấy con trỏ đến con trỏ, trong trường hợp này
điều đó là đúng. Hãy nghĩ về nó như là lấy con trỏ đến giá trị của kiểu lỗi;
tình cờ trong trường hợp này lỗi được trả về là một kiểu con trỏ.

### Bọc lỗi với %w

Như đã đề cập trước đó, việc dùng hàm `fmt.Errorf` để thêm thông tin bổ sung vào lỗi là phổ biến.

	if err != nil {
		return fmt.Errorf("decompress %v: %v", name, err)
	}

Trong Go 1.13, hàm `fmt.Errorf` hỗ trợ động từ mới `%w`. Khi động từ này
có mặt, lỗi được trả về bởi `fmt.Errorf` sẽ có phương thức `Unwrap`
trả về đối số của `%w`, đối số này phải là lỗi. Ở mọi khía cạnh khác, `%w`
giống hệt `%v`.

	if err != nil {
		// Trả về lỗi có thể unwrap về err.
		return fmt.Errorf("decompress %v: %w", name, err)
	}

Bọc lỗi với `%w` làm cho nó có thể dùng được với `errors.Is` và `errors.As`:

	err := fmt.Errorf("access denied: %w", ErrPermission)
	...
	if errors.Is(err, ErrPermission) ...

### Khi nào nên bọc

Khi thêm ngữ cảnh bổ sung vào lỗi, dù dùng `fmt.Errorf` hay bằng cách
triển khai kiểu tùy chỉnh, bạn cần quyết định xem lỗi mới có nên bọc
lỗi gốc hay không. Không có câu trả lời duy nhất; nó phụ thuộc vào
ngữ cảnh tạo ra lỗi mới. Bọc lỗi để để lộ nó cho
người gọi. Không bọc lỗi khi làm vậy sẽ lộ chi tiết triển khai.

Ví dụ, hãy tưởng tượng hàm `Parse` đọc cấu trúc dữ liệu phức tạp
từ `io.Reader`. Nếu lỗi xảy ra, chúng ta muốn báo cáo số dòng và cột
mà lỗi xảy ra. Nếu lỗi xảy ra khi đọc từ
`io.Reader`, chúng ta sẽ muốn bọc lỗi đó để cho phép kiểm tra
vấn đề bên dưới. Vì người gọi cung cấp `io.Reader` cho hàm,
sẽ hợp lý khi để lộ lỗi do nó tạo ra.

Ngược lại, hàm thực hiện nhiều lần gọi đến cơ sở dữ liệu có lẽ không nên
trả về lỗi có thể unwrap về kết quả của một trong các lần gọi đó. Nếu
cơ sở dữ liệu được dùng bởi hàm là chi tiết triển khai, thì việc để lộ các
lỗi này là vi phạm trừu tượng. Ví dụ, nếu hàm `LookupUser` của
gói `pkg` của bạn dùng gói `database/sql` của Go, thì nó có thể gặp lỗi
`sql.ErrNoRows`. Nếu bạn trả về lỗi đó với
`fmt.Errorf("accessing DB: %v", err)`
thì người gọi không thể nhìn bên trong để tìm `sql.ErrNoRows`. Nhưng nếu
hàm thay vào đó trả về `fmt.Errorf("accessing DB: %w", err)`, thì người gọi
có thể hợp lý viết

	err := pkg.LookupUser(...)
	if errors.Is(err, sql.ErrNoRows) …

Lúc đó, hàm phải luôn trả về `sql.ErrNoRows` nếu bạn không muốn
phá vỡ các client, ngay cả khi bạn chuyển sang gói cơ sở dữ liệu khác. Nói cách khác,
bọc lỗi làm cho lỗi đó trở thành một phần của API của bạn. Nếu bạn không
muốn cam kết hỗ trợ lỗi đó như một phần của API trong tương lai, bạn
không nên bọc lỗi.

Điều quan trọng cần nhớ là dù bạn có bọc hay không, văn bản lỗi vẫn sẽ là
như nhau. Một _người_ cố gắng hiểu lỗi sẽ có cùng thông tin
dù sao; quyết định bọc là về việc có cung cấp cho _chương trình_ thêm
thông tin để chúng có thể đưa ra quyết định sáng suốt hơn, hay giữ lại thông tin đó
để bảo toàn tầng trừu tượng.

## Tùy chỉnh kiểm tra lỗi với phương thức Is và As

Hàm `errors.Is` kiểm tra từng lỗi trong chuỗi để tìm khớp với
giá trị đích. Theo mặc định, lỗi khớp với đích nếu cả hai
[bằng nhau](/ref/spec#Comparison_operators). Ngoài ra, một
lỗi trong chuỗi có thể khai báo rằng nó khớp với đích bằng cách triển khai phương thức `Is`.

Ví dụ, hãy xem lỗi này lấy cảm hứng từ
[gói lỗi Upspin](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html)
so sánh lỗi với template, chỉ xem xét các trường khác không trong template:

	type Error struct {
		Path string
		User string
	}

	func (e *Error) Is(target error) bool {
		t, ok := target.(*Error)
		if !ok {
			return false
		}
		return (e.Path == t.Path || t.Path == "") &&
			   (e.User == t.User || t.User == "")
	}

	if errors.Is(err, &Error{User: "someuser"}) {
		// trường User của err là "someuser".
	}

Hàm `errors.As` cũng tham khảo phương thức `As` khi có.

## Lỗi và API gói

Một gói trả về lỗi (và hầu hết đều trả về) nên mô tả những thuộc tính nào
của các lỗi đó mà lập trình viên có thể dựa vào. Gói được thiết kế tốt cũng sẽ tránh
trả về lỗi với các thuộc tính không nên được dựa vào.

Đặc tả đơn giản nhất là nói rằng các thao tác thành công hoặc thất bại,
trả về giá trị lỗi nil hoặc non-nil tương ứng. Trong nhiều trường hợp, không cần
thêm thông tin.

Nếu chúng ta muốn một hàm trả về điều kiện lỗi có thể nhận dạng, chẳng hạn như "item
không tìm thấy," chúng ta có thể trả về lỗi bọc một sentinel.

	var ErrNotFound = errors.New("not found")

	// FetchItem trả về item có tên được chỉ định.
	//
	// Nếu không tồn tại item nào với tên đó, FetchItem trả về lỗi
	// bọc ErrNotFound.
	func FetchItem(name string) (*Item, error) {
		if itemNotFound(name) {
			return nil, fmt.Errorf("%q: %w", name, ErrNotFound)
		}
		// ...
	}

Có các mẫu hiện có khác để cung cấp lỗi có thể được người gọi kiểm tra về mặt ngữ nghĩa,
chẳng hạn như trực tiếp trả về giá trị sentinel, kiểu cụ thể,
hoặc giá trị có thể được kiểm tra với hàm predicate.

Trong mọi trường hợp, cần thận trọng không để lộ chi tiết nội bộ cho người dùng.
Như chúng ta đã đề cập trong "Khi nào nên bọc" ở trên, khi bạn trả về
lỗi từ gói khác, bạn nên chuyển đổi lỗi sang dạng không
để lộ lỗi bên dưới, trừ khi bạn sẵn sàng cam kết trả về
lỗi cụ thể đó trong tương lai.

	f, err := os.Open(filename)
	if err != nil {
		// *os.PathError được trả về bởi os.Open là chi tiết nội bộ.
		// Để tránh để lộ nó cho người gọi, đóng gói lại nó như một
		// lỗi mới với cùng văn bản. Chúng ta dùng động từ định dạng %v, vì
		// %w sẽ cho phép người gọi unwrap *os.PathError gốc.
		return fmt.Errorf("%v", err)
	}

Nếu một hàm được định nghĩa để trả về lỗi bọc một sentinel hoặc kiểu nào đó,
không trả về lỗi bên dưới trực tiếp.

	var ErrPermission = errors.New("permission denied")

	// DoSomething trả về lỗi bọc ErrPermission nếu người dùng
	// không có quyền thực hiện gì đó.
	func DoSomething() error {
		if !userHasPermission() {
			// Nếu chúng ta trả về ErrPermission trực tiếp, người gọi có thể
			// bắt đầu phụ thuộc vào giá trị lỗi chính xác, viết code như sau:
			//
			//     if err := pkg.DoSomething(); err == pkg.ErrPermission { … }
			//
			// Điều này sẽ gây ra vấn đề nếu chúng ta muốn thêm thêm
			// ngữ cảnh vào lỗi trong tương lai. Để tránh điều này, chúng ta
			// trả về lỗi bọc sentinel để người dùng phải
			// luôn unwrap nó:
			//
			//     if err := pkg.DoSomething(); errors.Is(err, pkg.ErrPermission) { ... }
			return fmt.Errorf("%w", ErrPermission)
		}
		// ...
	}

## Kết luận

Mặc dù các thay đổi chúng ta đã thảo luận chỉ là ba hàm và một
động từ định dạng, chúng tôi hy vọng chúng sẽ giúp cải thiện đáng kể cách xử lý lỗi
trong các chương trình Go. Chúng tôi kỳ vọng rằng việc bọc để cung cấp thêm ngữ cảnh
sẽ trở nên phổ biến, giúp chương trình đưa ra quyết định tốt hơn và giúp
lập trình viên tìm ra bug nhanh hơn.

Như Russ Cox đã nói trong [keynote GopherCon 2019](/blog/experiment),
trên con đường đến Go 2 chúng ta thử nghiệm, đơn giản hóa và phát hành. Bây giờ chúng ta đã
phát hành những thay đổi này, chúng tôi mong đợi các thử nghiệm tiếp theo.
