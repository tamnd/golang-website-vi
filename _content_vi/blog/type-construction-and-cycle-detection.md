---
title: Xây dựng kiểu và phát hiện chu trình
date: 2026-03-24
by:
- Mark Freeman
summary: Go 1.26 đơn giản hóa quá trình xây dựng kiểu và nâng cao khả năng phát hiện chu trình cho một số loại kiểu đệ quy nhất định.
template: true
---

Kiểu tĩnh của Go là một phần quan trọng lý giải tại sao Go phù hợp cho các hệ thống
production cần sự mạnh mẽ và đáng tin cậy. Khi một package Go được biên dịch, nó
trước tiên được phân tích cú pháp, tức là mã nguồn Go trong package đó được chuyển đổi
thành cây cú pháp trừu tượng (AST). AST này sau đó được truyền vào
*trình kiểm tra kiểu* (type checker) của Go.

Trong bài viết này, chúng ta sẽ đi sâu vào một phần của trình kiểm tra kiểu mà chúng tôi đã
cải thiện đáng kể trong Go 1.26. Điều này thay đổi gì từ góc nhìn của người dùng Go?
Trừ khi bạn thích các định nghĩa kiểu phức tạp, không có thay đổi nào có thể quan sát được ở đây.
Cải tiến này nhằm giảm các trường hợp đặc biệt, chuẩn bị cho các cải tiến trong tương lai của Go.
Ngoài ra, đây là một góc nhìn thú vị về điều mà có vẻ khá bình thường với các lập trình viên Go,
nhưng lại chứa đựng những sự tinh tế thực sự bên trong.

Nhưng trước tiên, *kiểm tra kiểu* là gì? Đây là bước trong trình biên dịch Go
loại bỏ toàn bộ các lớp lỗi tại thời gian biên dịch. Cụ thể, trình kiểm tra kiểu Go xác minh rằng:

1. Các kiểu xuất hiện trong AST là hợp lệ (ví dụ, kiểu khóa của một map phải
   `comparable`).
2. Các phép toán liên quan đến các kiểu đó (hoặc giá trị của chúng) là hợp lệ (ví dụ,
   không thể cộng một `int` và một `string`).

Để thực hiện điều này, trình kiểm tra kiểu xây dựng một biểu diễn nội tại cho
mỗi kiểu nó gặp trong khi duyệt AST, một quá trình được gọi không chính thức là
*xây dựng kiểu* (type construction).

Như chúng ta sẽ thấy, mặc dù Go được biết đến với hệ thống kiểu đơn giản, xây dựng kiểu
có thể phức tạp đến mức đánh lừa được trong một số góc của ngôn ngữ.

## Xây dựng kiểu

Hãy bắt đầu bằng cách xem xét một cặp khai báo kiểu đơn giản:

```go
type T []U
type U *int
```

Khi trình kiểm tra kiểu được gọi, nó đầu tiên gặp khai báo kiểu cho
`T`. Ở đây, AST ghi lại một định nghĩa kiểu của tên kiểu `T` và một
*biểu thức kiểu* `[]U`. `T` là một [kiểu đã định nghĩa](/ref/spec#Types); để biểu diễn
cấu trúc dữ liệu thực tế mà trình kiểm tra kiểu sử dụng khi xây dựng các kiểu đã định nghĩa,
chúng ta sẽ sử dụng một struct `Defined`.

Struct `Defined` chứa một con trỏ đến kiểu của biểu thức kiểu bên phải tên kiểu.
Trường `underlying` này có liên quan đến việc xác định
[kiểu cơ sở](/ref/spec#Underlying_types) của kiểu. Để minh họa trạng thái
của trình kiểm tra kiểu, hãy xem cách duyệt AST điền vào các cấu trúc
dữ liệu, bắt đầu với:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/01.svg" />

Tại thời điểm này, `T` đang *trong quá trình xây dựng*, được chỉ thị bởi màu vàng. Vì
chúng ta chưa đánh giá biểu thức kiểu `[]U` (nó vẫn màu đen), `underlying`
trỏ đến `nil`, được chỉ thị bởi một mũi tên rỗng.

Khi chúng ta đánh giá `[]U`, trình kiểm tra kiểu xây dựng một struct `Slice`, là
cấu trúc dữ liệu nội tại dùng để biểu diễn các kiểu slice. Tương tự như `Defined`,
nó chứa một con trỏ đến kiểu phần tử cho slice. Chúng ta chưa biết tên
`U` chỉ đến cái gì, mặc dù chúng ta kỳ vọng nó chỉ đến một kiểu. Vì vậy, lại một lần nữa,
con trỏ này là `nil`. Chúng ta kết thúc với:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/02.svg" />

Đến đây bạn có thể đã nắm được ý chính rồi, vì vậy chúng ta sẽ tăng tốc một chút.

Để chuyển đổi tên kiểu `U` thành một kiểu, trước tiên chúng ta xác định vị trí khai báo của nó.
Nhận thấy nó đại diện cho một kiểu đã định nghĩa khác, chúng ta xây dựng một `Defined` riêng
cho `U` tương ứng. Kiểm tra bên phải của `U`, chúng ta thấy biểu thức kiểu
`*int`, được đánh giá thành một struct `Pointer`, với kiểu cơ sở của con trỏ là biểu thức kiểu `int`.

Khi chúng ta đánh giá `int`, có điều gì đó đặc biệt xảy ra: chúng ta nhận lại một kiểu
được khai báo trước. Các kiểu được khai báo trước được xây dựng *trước* khi trình kiểm tra kiểu
bắt đầu duyệt AST. Vì kiểu cho `int` đã được xây dựng, không có gì cho chúng ta làm
ngoài việc trỏ đến kiểu đó.

Bây giờ chúng ta có:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/03.svg" />

Lưu ý rằng kiểu `Pointer` là *hoàn chỉnh* (complete) tại thời điểm này, được chỉ thị bởi màu
xanh lá. Tính hoàn chỉnh có nghĩa là cấu trúc dữ liệu nội tại của kiểu có tất cả các
trường được điền và các kiểu được trỏ đến bởi các trường đó là hoàn chỉnh.
Tính hoàn chỉnh là một thuộc tính quan trọng của kiểu vì nó đảm bảo rằng
việc truy cập vào nội bộ (deconstruction) của kiểu đó là hợp lệ: chúng ta có tất cả
thông tin mô tả kiểu.

Trong hình trên, struct `Pointer` chỉ chứa trường `base`, trỏ đến `int`.
Vì `int` không có trường nào để điền, nó "rỗng hoàn chỉnh" (vacuously complete),
làm cho kiểu cho `*int` là hoàn chỉnh.

Từ đây, trình kiểm tra kiểu bắt đầu tháo gỡ stack. Vì kiểu cho
`*int` là hoàn chỉnh, chúng ta có thể hoàn thành kiểu cho `U`, nghĩa là chúng ta có thể hoàn thành
kiểu cho `[]U`, và tương tự cho `T`. Khi quá trình này kết thúc, chúng ta kết thúc với
chỉ các kiểu hoàn chỉnh, như được hiển thị bên dưới:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/04.svg" />

Đánh số ở trên cho thấy thứ tự các kiểu được hoàn thành (sau `Pointer`). Lưu ý rằng
kiểu ở dưới cùng được hoàn thành trước. Xây dựng kiểu về cơ bản là một quá trình ưu tiên chiều sâu,
vì việc hoàn thành một kiểu đòi hỏi các phụ thuộc của nó phải được hoàn thành trước.

## Các kiểu đệ quy

Với ví dụ đơn giản này đã qua, hãy thêm một chút phức tạp. Hệ thống kiểu của Go
cũng cho phép chúng ta diễn tả các kiểu đệ quy. Một ví dụ điển hình là:

```go
type Node struct {
  next *Node
}
```

Nếu chúng ta xem xét lại ví dụ trên, chúng ta có thể thêm một chút đệ quy bằng cách
thay `*int` bằng `*T` như sau:

```go
type T []U
type U *T
```

Bây giờ hãy theo dõi: hãy bắt đầu lại với `T`, nhưng bỏ qua một số bước để minh họa
hiệu quả của thay đổi này. Như người ta có thể nghi ngờ từ ví dụ trước,
trình kiểm tra kiểu sẽ tiếp cận việc đánh giá `*T` với trạng thái dưới đây:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/05.svg" />

Câu hỏi là phải làm gì với kiểu cơ sở cho `*T`. Chúng ta có ý tưởng về
`T` là gì (một `Defined`), nhưng nó hiện đang được xây dựng (trường `underlying` của nó
vẫn là `nil`).

Chúng ta đơn giản trỏ kiểu cơ sở cho `*T` đến `T`, mặc dù `T` chưa hoàn chỉnh:

<img id="example" style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/06.svg" />

Chúng ta làm điều này với giả định rằng `T` sẽ hoàn chỉnh khi nó kết thúc quá trình xây dựng
*trong tương lai* (bằng cách trỏ đến một kiểu hoàn chỉnh). Khi điều đó xảy ra, `base` sẽ
trỏ đến một kiểu hoàn chỉnh, do đó làm cho `*T` trở nên hoàn chỉnh.

Trong thời gian đó, chúng ta sẽ bắt đầu đi ngược lại stack:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/07.svg" />

Khi chúng ta quay lại đỉnh và kết thúc việc xây dựng `T`, "vòng lặp" của các kiểu
sẽ đóng lại, hoàn thành mỗi kiểu trong vòng lặp đồng thời:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/08.svg" />

Trước khi chúng ta xem xét các kiểu đệ quy, việc đánh giá một biểu thức kiểu luôn
trả về một kiểu hoàn chỉnh. Đó là thuộc tính tiện lợi vì nó có nghĩa là trình kiểm tra kiểu
có thể luôn phân tách (look inside) một kiểu được trả về từ đánh giá.

Nhưng trong [ví dụ trên](#example), đánh giá `T` trả về một kiểu *chưa hoàn chỉnh*,
có nghĩa là việc phân tách `T` là không hợp lệ cho đến khi nó hoàn chỉnh. Nói chung
nói, các kiểu đệ quy có nghĩa là trình kiểm tra kiểu không còn có thể giả định rằng
các kiểu được trả về từ đánh giá sẽ là hoàn chỉnh.

Tuy nhiên, kiểm tra kiểu bao gồm nhiều kiểm tra đòi hỏi phải phân tách một kiểu.
Một ví dụ điển hình là xác nhận rằng khóa của map là `comparable`, đòi hỏi
kiểm tra trường `underlying`. Làm sao chúng ta có thể tương tác an toàn với các kiểu
chưa hoàn chỉnh như `T`?

Nhớ lại rằng tính hoàn chỉnh của kiểu là điều kiện tiên quyết để phân tách một kiểu.
Trong trường hợp này, xây dựng kiểu không bao giờ phân tách một kiểu, nó chỉ tham chiếu đến
các kiểu. Nói cách khác, tính hoàn chỉnh của kiểu *không* cản trở xây dựng kiểu ở đây.

Vì xây dựng kiểu không bị cản trở, trình kiểm tra kiểu có thể đơn giản trì hoãn các
kiểm tra như vậy cho đến khi kết thúc quá trình kiểm tra kiểu, khi tất cả các kiểu đều hoàn chỉnh
(lưu ý rằng bản thân các kiểm tra cũng không cản trở xây dựng kiểu). Nếu một kiểu tiết lộ
một lỗi kiểu, không có gì khác biệt khi lỗi đó được báo cáo trong quá trình kiểm tra
kiểu, chỉ cần nó được báo cáo cuối cùng.

Với kiến thức này, hãy xem xét một ví dụ phức tạp hơn liên quan đến
các giá trị của kiểu chưa hoàn chỉnh.

## Kiểu đệ quy và giá trị

Hãy dừng lại một chút và xem xét
[kiểu mảng](/ref/spec#Array_types) của Go. Quan trọng là, kiểu mảng có kích thước,
đây là [hằng số](/ref/spec#Array_types) là một phần của kiểu. Một số
phép toán, như các hàm tích hợp `unsafe.Sizeof` và `len` có thể trả về
các hằng số khi được áp dụng cho [một số giá trị nhất định](/ref/spec#Package_unsafe) hoặc
[biểu thức](/ref/spec#Length_and_capacity), có nghĩa là chúng có thể xuất hiện như
kích thước mảng. Quan trọng là, các giá trị được truyền cho các hàm đó có thể thuộc bất kỳ kiểu nào,
ngay cả một kiểu chưa hoàn chỉnh. Chúng ta gọi những giá trị này là *giá trị chưa hoàn chỉnh*.

Hãy xem xét ví dụ này:

```go
type T [unsafe.Sizeof(T{})]int
```

Theo cách tương tự như trước, chúng ta sẽ đạt đến trạng thái như bên dưới:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/09.svg" />

Để xây dựng `Array`, chúng ta phải tính kích thước của nó. Từ biểu thức giá trị
`unsafe.Sizeof(T{})`, đó là kích thước của `T`. Đối với các kiểu mảng (như `T`),
việc tính kích thước đòi hỏi phân tách: chúng ta cần nhìn vào bên trong kiểu để
xác định độ dài của mảng và kích thước của từng phần tử.

Nói cách khác, xây dựng kiểu cho `Array` *thực sự* phân tách `T`,
có nghĩa là `Array` không thể kết thúc quá trình xây dựng (chứ chưa nói đến hoàn chỉnh) trước khi
`T` hoàn chỉnh. Mẹo "vòng lặp" mà chúng ta đã sử dụng trước đó, trong đó một vòng lặp các kiểu
đồng thời hoàn chỉnh khi kiểu bắt đầu vòng lặp kết thúc quá trình xây dựng, không hoạt động ở đây.

Điều này để chúng ta trong một tình huống khó:

* `T` không thể hoàn chỉnh cho đến khi `Array` hoàn chỉnh.
* `Array` không thể hoàn chỉnh cho đến khi `T` hoàn chỉnh.
* Chúng *không thể* hoàn chỉnh đồng thời (không giống như trước).

Rõ ràng, điều này là không thể thỏa mãn. Trình kiểm tra kiểu phải làm gì?

### Phát hiện chu trình

Về cơ bản, mã như thế này là không hợp lệ vì kích thước của `T` không thể
được xác định mà không biết kích thước của `T`, bất kể trình kiểm tra kiểu
hoạt động như thế nào. Trường hợp cụ thể này, định nghĩa kích thước theo chu trình, là một phần của
một lớp lỗi gọi là *lỗi chu trình*, thường liên quan đến định nghĩa chu trình của
các cấu trúc Go. Như một ví dụ khác, xem xét `type T T`, cũng thuộc
lớp này, nhưng vì lý do khác. Quá trình tìm và báo cáo lỗi chu trình
trong quá trình kiểm tra kiểu được gọi là *phát hiện chu trình*.

Bây giờ, phát hiện chu trình hoạt động như thế nào cho `type T [unsafe.Sizeof(T{})]int`?
Để trả lời điều này, hãy xem xét `T{}` bên trong. Vì `T{}` là một biểu thức composite literal,
trình kiểm tra kiểu biết rằng giá trị kết quả của nó có kiểu `T`.
Vì `T` chưa hoàn chỉnh, chúng ta gọi giá trị `T{}` là *giá trị chưa hoàn chỉnh*.

Chúng ta phải thận trọng, vì thao tác trên một giá trị chưa hoàn chỉnh chỉ hợp lệ nếu
nó không phân tách kiểu của giá trị. Ví dụ, `type T [unsafe.Sizeof(new(T))]int`
*là* hợp lệ, vì giá trị `new(T)` (kiểu `*T`) không bao giờ bị phân tách,
tất cả các con trỏ đều có cùng kích thước. Để nhấn mạnh lại, việc xác định kích thước
của một giá trị chưa hoàn chỉnh kiểu `*T` là hợp lệ, nhưng không phải kiểu `T`.

Điều này bởi vì "tính con trỏ" của `*T` cung cấp đủ thông tin kiểu cho
`unsafe.Sizeof`, trong khi chỉ `T` thì không. Trên thực tế, không bao giờ hợp lệ để thao tác
trên một giá trị chưa hoàn chỉnh *mà kiểu của nó là một kiểu đã định nghĩa*, vì chỉ một tên kiểu
không truyền tải bất kỳ thông tin kiểu (cơ sở) nào cả.

#### Nơi thực hiện phát hiện

Cho đến nay chúng ta đã tập trung vào `unsafe.Sizeof` trực tiếp thao tác trên
các giá trị có thể chưa hoàn chỉnh. Trong `type T [unsafe.Sizeof(T{})]int`, lời gọi đến `unsafe.Sizeof`
chỉ là "gốc" của biểu thức độ dài mảng. Chúng ta có thể dễ dàng hình dung
giá trị chưa hoàn chỉnh `T{}` như một toán hạng trong biểu thức giá trị nào đó khác.

Ví dụ, nó có thể được truyền cho một hàm (tức là
`type T [unsafe.Sizeof(f(T{}))]int`), được slice (tức là
`type T [unsafe.Sizeof(T{}[:])]int`), được lập chỉ mục (tức là
`type T [unsafe.Sizeof(T{}[0])]int`), v.v. Tất cả những cái này không hợp lệ vì chúng
đòi hỏi phân tách `T`. Ví dụ, việc lập chỉ mục `T`
[đòi hỏi kiểm tra](/ref/spec#Index_expressions) kiểu cơ sở của `T`.
Vì các biểu thức này "tiêu thụ" các giá trị có thể chưa hoàn chỉnh, hãy gọi chúng là
*downstream*. Có nhiều ví dụ hơn về toán tử downstream, một số trong đó không rõ ràng về mặt cú pháp.

Tương tự, `T{}` chỉ là một ví dụ về biểu thức "tạo ra" một giá trị có thể chưa hoàn chỉnh,
hãy gọi các loại biểu thức này là *upstream*:

<img style="background-color:#f3f3f3" src="type-construction-and-cycle-detection/10.svg" />

So sánh, có ít hơn và các biểu thức giá trị rõ ràng hơn về mặt cú pháp
có thể dẫn đến các giá trị chưa hoàn chỉnh. Ngoài ra, việc liệt kê
các trường hợp này bằng cách kiểm tra định nghĩa cú pháp của Go khá đơn giản. Vì những lý do này,
sẽ đơn giản hơn để triển khai logic phát hiện chu trình của chúng ta thông qua các upstream, nơi
các giá trị có thể chưa hoàn chỉnh bắt nguồn. Dưới đây là một số ví dụ về chúng:

```go

type T [unsafe.Sizeof(T(42))]int                // chuyển đổi

func f() T
type T [unsafe.Sizeof(f())]int                  // gọi hàm

var i interface{}
type T [unsafe.Sizeof(i.(T))]int                // type assertion

type T [unsafe.Sizeof({{raw "<"}}-(make({{raw "<"}}-chan T)))]int   // nhận channel

type T [unsafe.Sizeof(make(map[int]T)[42])]int  // truy cập map

type T [unsafe.Sizeof(*new(T))]int              // dereference

// ... và một vài trường hợp khác
```

Đối với mỗi trường hợp này, trình kiểm tra kiểu có logic bổ sung nơi loại biểu thức giá trị
cụ thể đó được đánh giá. Ngay khi chúng ta biết kiểu của
giá trị kết quả, chúng ta chèn một bài kiểm tra đơn giản xác nhận rằng kiểu đó là hoàn chỉnh.

Ví dụ, trong ví dụ chuyển đổi `type T [unsafe.Sizeof(T(42))]int`,
có một đoạn trong trình kiểm tra kiểu giống như:

```go
func callExpr(call *syntax.CallExpr) operand {
  x := typeOrValue(call.Fun)
  switch x.mode() {
  // ... các trường hợp khác
  case typeExpr:
    // T(), nghĩa là đây là một chuyển đổi
    T := x.typ()
    // ... xử lý chuyển đổi, T *không an toàn* để phân tách
  }
}
```

Ngay khi chúng ta quan sát thấy `CallExpr` là một chuyển đổi sang `T`, chúng ta biết rằng
kiểu kết quả sẽ là `T` (giả sử không có lỗi trước đó). Trước khi chúng ta trả
lại một giá trị (ở đây là một `operand`) của kiểu `T` cho phần còn lại của trình kiểm tra kiểu,
chúng ta cần kiểm tra tính hoàn chỉnh của `T`:

```go
func callExpr(call *syntax.CallExpr) operand {
  x := typeOrValue(call.Fun)
  switch x.mode() {
  // ... các trường hợp khác
  case typeExpr:
    // T(), nghĩa là đây là một chuyển đổi
    T := x.typ()
+   if !isComplete(T) {
+     reportCycleErr(T)
+     return invalid
+   }
    // ... xử lý chuyển đổi, T *an toàn* để phân tách
  }
}
```

Thay vì trả về một giá trị chưa hoàn chỉnh, chúng ta trả về một toán hạng `invalid` đặc biệt,
báo hiệu rằng biểu thức gọi hàm không thể được đánh giá. Phần còn lại của
trình kiểm tra kiểu có xử lý đặc biệt cho các toán hạng không hợp lệ. Bằng cách thêm điều này,
chúng ta đã ngăn các giá trị chưa hoàn chỉnh "thoát" xuống downstream, cả vào phần còn lại của
logic chuyển đổi kiểu lẫn các toán tử downstream, và thay vào đó đã báo cáo lỗi chu trình
mô tả vấn đề với `T`.

Một mô hình mã tương tự được sử dụng trong tất cả các trường hợp khác, thực hiện phát hiện chu trình
cho các giá trị chưa hoàn chỉnh.

## Kết luận

Phát hiện chu trình có hệ thống liên quan đến các giá trị chưa hoàn chỉnh là một bổ sung mới vào
trình kiểm tra kiểu. Trước Go 1.26, chúng tôi đã sử dụng một thuật toán xây dựng kiểu phức tạp hơn,
bao gồm phát hiện chu trình đặc thù hơn và không phải lúc nào cũng hoạt động. Cách tiếp cận
mới, đơn giản hơn của chúng tôi đã giải quyết một số sự cố trình biên dịch (thực sự là bí truyền)
(các vấn đề [\#75918](/issue/75918), [\#76383](/issue/76383),
[\#76384](/issue/76384), [\#76478](/issue/76478), và nhiều hơn nữa), dẫn đến một
trình biên dịch ổn định hơn.

Là các lập trình viên, chúng ta đã quen với các tính năng như định nghĩa kiểu đệ quy
và kiểu mảng có kích thước đến mức có thể bỏ qua sự tinh tế của
độ phức tạp cơ bản của chúng. Mặc dù bài viết này có bỏ qua một số chi tiết tinh tế hơn,
hy vọng chúng ta đã truyền đạt được sự hiểu biết sâu sắc hơn (và có lẽ là sự đánh giá cao hơn)
về các vấn đề xung quanh kiểm tra kiểu trong Go.
