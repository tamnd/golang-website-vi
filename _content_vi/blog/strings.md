---
title: Strings, bytes, runes và characters trong Go
date: 2013-10-23
by:
- Rob Pike
tags:
- strings
- bytes
- runes
- characters
summary: Cách strings hoạt động trong Go và cách sử dụng chúng.
template: true
---

## Giới thiệu

[Bài đăng blog trước](/blog/slices) đã giải thích cách slices
hoạt động trong Go, sử dụng nhiều ví dụ để minh họa cơ chế đằng sau
quá trình triển khai của chúng.
Dựa trên nền tảng đó, bài đăng này thảo luận về strings trong Go.
Thoạt đầu, strings có vẻ là một chủ đề quá đơn giản cho một bài đăng blog, nhưng để
sử dụng chúng tốt đòi hỏi phải hiểu không chỉ cách chúng hoạt động,
mà còn sự khác biệt giữa byte, character và rune,
sự khác biệt giữa Unicode và UTF-8,
sự khác biệt giữa string và string literal,
và nhiều phân biệt tinh tế hơn khác.

Một cách tiếp cận chủ đề này là nghĩ nó như câu trả lời cho câu hỏi thường được hỏi,
"Khi tôi truy cập vào vị trí _n_ trong một Go string, tại sao tôi không lấy được
ký tự thứ _n_?"
Như bạn sẽ thấy, câu hỏi này dẫn chúng ta đến nhiều chi tiết về cách văn bản hoạt động
trong thế giới hiện đại.

Một giới thiệu xuất sắc về một số vấn đề này, độc lập với Go,
là bài đăng blog nổi tiếng của Joel Spolsky,
[The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](http://www.joelonsoftware.com/articles/Unicode.html).
Nhiều điểm ông đề cập sẽ được nhắc lại ở đây.

## String là gì?

Hãy bắt đầu với một số kiến thức cơ bản.

Trong Go, một string về cơ bản là một slice byte chỉ đọc.
Nếu bạn còn không chắc chắn về slice of bytes là gì hoặc nó hoạt động như thế nào,
vui lòng đọc [bài đăng blog trước](/blog/slices);
chúng tôi sẽ giả sử ở đây rằng bạn đã đọc rồi.

Điều quan trọng cần nêu ngay từ đầu là một string chứa các byte _tùy ý_.
Nó không yêu cầu phải chứa văn bản Unicode, văn bản UTF-8, hoặc bất kỳ định dạng xác định nào khác.
Về mặt nội dung của một string, nó hoàn toàn tương đương với một
slice of bytes.

Đây là một string literal (sẽ nói thêm về điều này sau) sử dụng ký hiệu
`\xNN` để định nghĩa một hằng string chứa một số giá trị byte đặc biệt.
(Dĩ nhiên, các byte nằm trong khoảng giá trị thập lục phân từ 00 đến FF, bao gồm cả hai đầu.)

{{code "strings/basic.go" `/const sample/`}}

## In strings

Vì một số byte trong chuỗi mẫu của chúng ta không phải là ASCII hợp lệ, thậm chí
không phải UTF-8 hợp lệ, việc in string trực tiếp sẽ tạo ra đầu ra xấu.
Câu lệnh print đơn giản

{{code "strings/basic.go" `/println/` `/println/`}}

tạo ra kết quả lộn xộn này (diện mạo chính xác tùy thuộc vào môi trường):

	��=� ⌘

Để tìm hiểu string thực sự chứa gì, chúng ta cần tách nó ra và kiểm tra các phần.
Có một số cách để làm điều này.
Cách rõ ràng nhất là vòng lặp qua nội dung của nó và lấy ra từng byte
một, như trong vòng lặp `for` này:

{{code "strings/basic.go" `/byte loop/` `/byte loop/`}}

Như đã ngụ ý từ đầu, việc truy cập vào một string truy cập từng byte, không phải
characters. Chúng ta sẽ quay lại chủ đề đó chi tiết hơn bên dưới. Bây giờ, hãy
chỉ tập trung vào các byte.
Đây là đầu ra từ vòng lặp byte theo byte:

	bd b2 3d bc 20 e2 8c 98

Chú ý cách từng byte khớp với
các ký tự thoát thập lục phân đã định nghĩa string.

Một cách ngắn hơn để tạo ra đầu ra dễ đọc cho một string lộn xộn
là sử dụng format verb `%x` (thập lục phân) của `fmt.Printf`.
Nó chỉ đơn giản đổ ra các byte tuần tự của string dưới dạng
chữ số thập lục phân, hai chữ số cho mỗi byte.

{{code "strings/basic.go" `/percent x/` `/percent x/`}}

So sánh đầu ra của nó với phần trên:

	bdb23dbc20e28c98

Một mẹo hay là sử dụng cờ "space" trong format đó, đặt
khoảng trắng giữa `%` và `x`. So sánh chuỗi format
được sử dụng ở đây với cái trên,

{{code "strings/basic.go" `/percent space x/` `/percent space x/`}}

và chú ý cách các byte xuất hiện
với khoảng trắng ở giữa, làm cho kết quả ít đáng sợ hơn một chút:

	bd b2 3d bc 20 e2 8c 98

Còn nhiều hơn nữa. Verb `%q` (quoted) sẽ thoát bất kỳ chuỗi byte không in được
trong một string để đầu ra rõ ràng.

{{code "strings/basic.go" `/percent q/` `/percent q/`}}

Kỹ thuật này hữu ích khi phần lớn string
có thể hiểu được dưới dạng văn bản nhưng có những điểm kỳ lạ cần tìm ra; nó tạo ra:

	"\xbd\xb2=\xbc ⌘"

Nếu nhìn kỹ vào đó, chúng ta có thể thấy rằng ẩn trong nhiễu là một dấu bằng ASCII,
cùng với một khoảng trắng thông thường, và ở cuối xuất hiện ký hiệu "Place of Interest"
nổi tiếng của Thụy Điển.
Ký hiệu đó có giá trị Unicode là U+2318, được mã hóa dưới dạng UTF-8 bởi các byte
sau khoảng trắng (giá trị thập lục phân `20`): `e2` `8c` `98`.

Nếu chúng ta không quen hoặc bị nhầm lẫn bởi các giá trị lạ trong string,
chúng ta có thể sử dụng cờ "plus" cho verb `%q`. Cờ này khiến đầu ra thoát
không chỉ các chuỗi không in được, mà còn bất kỳ byte không phải ASCII nào, tất cả
trong khi giải thích UTF-8.
Kết quả là nó làm lộ ra các giá trị Unicode của UTF-8 được định dạng đúng
đại diện cho dữ liệu không phải ASCII trong string:

{{code "strings/basic.go" `/percent plus q/` `/percent plus q/`}}

Với format đó, giá trị Unicode của ký hiệu Thụy Điển xuất hiện dưới dạng
ký tự thoát `\u`:

	"\xbd\xb2=\xbc \u2318"

Các kỹ thuật in này rất hữu ích khi gỡ lỗi
nội dung của strings, và sẽ hữu ích trong phần thảo luận tiếp theo.
Cũng đáng chú ý rằng tất cả các phương pháp này hoạt động hoàn toàn
giống nhau cho các byte slice cũng như cho strings.

Đây là bộ đầy đủ các tùy chọn in mà chúng tôi đã liệt kê, được trình bày dưới dạng
một chương trình hoàn chỉnh bạn có thể chạy (và chỉnh sửa) ngay trong trình duyệt:

{{play "strings/basic.go" `/package/` `/^}/`}}

[Bài tập: Sửa đổi các ví dụ trên để sử dụng một slice of bytes
thay vì string. Gợi ý: Sử dụng chuyển đổi để tạo slice.]

[Bài tập: Lặp qua string sử dụng format `%q` trên từng byte.
Đầu ra cho bạn biết điều gì?]

## UTF-8 và string literals

Như chúng ta đã thấy, việc truy cập vào một string tạo ra các byte của nó, không phải characters: một string chỉ là
một đống bytes.
Điều đó có nghĩa là khi chúng ta lưu trữ một giá trị character trong một string,
chúng ta lưu trữ biểu diễn từng byte của nó.
Hãy xem một ví dụ được kiểm soát hơn để xem điều đó xảy ra như thế nào.

Đây là một chương trình đơn giản in một hằng string với một ký tự duy nhất
theo ba cách khác nhau, một lần dưới dạng string thông thường, một lần dưới dạng string
chỉ ASCII được trích dẫn, và một lần dưới dạng từng byte riêng lẻ dưới dạng thập lục phân.
Để tránh nhầm lẫn, chúng ta tạo một "raw string", được bao quanh bởi dấu backtick,
để nó chỉ có thể chứa văn bản literal. (Các string thông thường, được bao quanh bởi dấu
ngoặc kép, có thể chứa các chuỗi thoát như chúng ta đã thấy ở trên.)

{{play "strings/utf8.go" `/^func/` `/^}/`}}

Đầu ra là:

	plain string: ⌘
	quoted string: "\u2318"
	hex bytes: e2 8c 98

nhắc nhở chúng ta rằng giá trị ký tự Unicode U+2318, ký hiệu "Place
of Interest" ⌘, được biểu diễn bởi các byte `e2` `8c` `98`, và
những byte đó là mã hóa UTF-8 của giá trị thập lục phân 2318.

Điều này có thể rõ ràng hoặc có thể tinh tế, tùy thuộc vào sự quen thuộc của bạn với
UTF-8, nhưng đáng để dành một chút thời gian để giải thích cách biểu diễn UTF-8
của string được tạo ra.
Thực tế đơn giản là: nó được tạo ra khi mã nguồn được viết.

Mã nguồn trong Go _được định nghĩa_ là văn bản UTF-8; không có biểu diễn nào khác được
cho phép. Điều đó ngụ ý rằng khi, trong mã nguồn, chúng ta viết văn bản

	`⌘`

trình soạn thảo văn bản được sử dụng để tạo chương trình đặt mã hóa UTF-8
của ký hiệu ⌘ vào văn bản nguồn.
Khi chúng ta in ra các byte thập lục phân, chúng ta chỉ đổ ra
dữ liệu mà trình soạn thảo đã đặt trong tệp.

Nói tóm lại, mã nguồn Go là UTF-8, vì vậy
_mã nguồn cho string literal là văn bản UTF-8_.
Nếu string literal đó không chứa chuỗi thoát nào, điều mà một raw
string không thể, thì string được tạo ra sẽ chứa chính xác
văn bản nguồn giữa các dấu trích dẫn.
Do đó theo định nghĩa và
theo cách tạo dựng, raw string sẽ luôn chứa một biểu diễn UTF-8 hợp lệ
cho nội dung của nó.
Tương tự, trừ khi nó chứa các ký tự thoát phá vỡ UTF-8 như những cái
từ phần trước, một string literal thông thường cũng sẽ luôn chứa UTF-8 hợp lệ.

Một số người nghĩ rằng Go strings luôn là UTF-8, nhưng chúng
không phải: chỉ có string literals mới là UTF-8.
Như chúng ta đã thấy trong phần trước, các _giá trị_ string có thể chứa các byte tùy ý;
như chúng ta đã thấy trong phần này, các _literals_ string luôn chứa văn bản UTF-8
miễn là chúng không có các ký tự thoát ở cấp độ byte.

Tóm lại, strings có thể chứa các byte tùy ý, nhưng khi được tạo
từ string literals, những byte đó (gần như luôn luôn) là UTF-8.

## Code points, characters và runes

Chúng ta đã rất cẩn thận cho đến nay trong cách chúng ta sử dụng các từ "byte" và "character".
Một phần là vì strings chứa bytes, và một phần vì ý niệm về "character"
hơi khó định nghĩa.
Tiêu chuẩn Unicode sử dụng thuật ngữ "code point" để chỉ mục được biểu diễn
bởi một giá trị duy nhất.
Code point U+2318, với giá trị thập lục phân 2318, đại diện cho ký hiệu ⌘.
(Để biết thêm thông tin về code point đó, hãy xem
[trang Unicode của nó](http://unicode.org/cldr/utility/character.jsp?a=2318).)

Để chọn một ví dụ thực tế hơn, code point Unicode U+0061 là chữ cái
Latin viết thường 'A': a.

Nhưng còn chữ cái viết thường có dấu huyền 'A', à thì sao?
Đó là một character, và nó cũng là một code point (U+00E0), nhưng nó có các biểu diễn khác.
Ví dụ chúng ta có thể sử dụng code point dấu huyền "combining", U+0300,
và gắn nó vào chữ cái viết thường a, U+0061, để tạo ra cùng character à.
Nhìn chung, một character có thể được biểu diễn bằng một số chuỗi code points khác nhau,
và do đó bằng các chuỗi bytes UTF-8 khác nhau.

Do đó khái niệm character trong điện toán là mơ hồ, hoặc ít nhất là
khó hiểu, vì vậy chúng ta sử dụng nó một cách cẩn thận.
Để làm cho mọi thứ đáng tin cậy, có các kỹ thuật _chuẩn hóa_ đảm bảo rằng
một character nhất định luôn được biểu diễn bởi cùng một code points, nhưng chủ đề đó
đưa chúng ta đi quá xa khỏi chủ đề chính hiện tại.
Một bài đăng blog sau sẽ giải thích cách các thư viện Go xử lý chuẩn hóa.

"Code point" hơi dài, vì vậy Go giới thiệu một thuật ngữ ngắn hơn cho
khái niệm này: _rune_.
Thuật ngữ này xuất hiện trong các thư viện và mã nguồn, và có nghĩa chính xác
giống như "code point", với một bổ sung thú vị.

Ngôn ngữ Go định nghĩa từ `rune` là bí danh của kiểu `int32`, vì vậy
các chương trình có thể rõ ràng khi một giá trị số nguyên đại diện cho một code point.
Hơn nữa, những gì bạn có thể nghĩ là hằng character được gọi là
_rune constant_ trong Go.
Kiểu và giá trị của biểu thức

	'⌘'

là `rune` với giá trị số nguyên `0x2318`.

Tóm lại, đây là những điểm chính:

  - Mã nguồn Go luôn là UTF-8.
  - Một string chứa các byte tùy ý.
  - Một string literal, không có các ký tự thoát ở cấp độ byte, luôn chứa các chuỗi UTF-8 hợp lệ.
  - Những chuỗi đó đại diện cho các code points Unicode, được gọi là runes.
  - Không có đảm bảo nào được đưa ra trong Go rằng các characters trong strings đã được chuẩn hóa.

## Vòng lặp range

Ngoài chi tiết hiển nhiên rằng mã nguồn Go là UTF-8,
thực sự chỉ có một cách mà Go xử lý UTF-8 đặc biệt, đó là khi sử dụng
vòng lặp `for` `range` trên một string.

Chúng ta đã thấy điều gì xảy ra với một vòng lặp `for` thông thường.
Ngược lại, vòng lặp `for` `range` giải mã một rune được mã hóa UTF-8 trong mỗi
lần lặp.
Mỗi lần vòng lặp, chỉ số của vòng lặp là vị trí bắt đầu của
rune hiện tại, được đo bằng bytes, và code point là giá trị của nó.
Đây là một ví dụ sử dụng format `Printf` tiện dụng khác, `%#U`, hiển thị
giá trị Unicode của code point và biểu diễn in của nó:

{{play "strings/range.go" `/const/` `/}/`}}

Đầu ra cho thấy cách mỗi code point chiếm nhiều byte:

	U+65E5 '日' starts at byte position 0
	U+672C '本' starts at byte position 3
	U+8A9E '語' starts at byte position 6

[Bài tập: Đặt một chuỗi byte UTF-8 không hợp lệ vào string. (Làm như thế nào?)
Điều gì xảy ra với các lần lặp của vòng lặp?]

## Thư viện

Thư viện chuẩn của Go cung cấp hỗ trợ mạnh mẽ cho việc diễn giải văn bản UTF-8.
Nếu vòng lặp `for` `range` không đủ cho mục đích của bạn,
rất có thể tính năng bạn cần được cung cấp bởi một gói trong thư viện.

Gói quan trọng nhất như vậy là
[`unicode/utf8`](/pkg/unicode/utf8/),
chứa
các hàm trợ giúp để xác thực, tháo rời và lắp ráp lại các strings UTF-8.
Đây là một chương trình tương đương với ví dụ `for` `range` ở trên,
nhưng sử dụng hàm `DecodeRuneInString` từ gói đó để
thực hiện công việc.
Các giá trị trả về từ hàm là rune và độ rộng của nó trong
các byte được mã hóa UTF-8.

{{play "strings/encoding.go" `/const/` `/}/`}}

Chạy nó để thấy rằng nó thực hiện cùng một kết quả.
Vòng lặp `for` `range` và `DecodeRuneInString` được định nghĩa để tạo ra
chính xác cùng một chuỗi lặp.

Xem
[tài liệu](/pkg/unicode/utf8/)
cho gói `unicode/utf8` để thấy
các tính năng khác mà nó cung cấp.

## Kết luận

Để trả lời câu hỏi được đặt ra ở đầu: Strings được xây dựng từ bytes
vì vậy việc truy cập vào chúng tạo ra bytes, không phải characters.
Một string thậm chí có thể không chứa characters.
Trên thực tế, định nghĩa về "character" là mơ hồ và sẽ là
một sai lầm khi cố gắng giải quyết sự mơ hồ bằng cách định nghĩa rằng strings được tạo
thành từ characters.

Còn rất nhiều điều để nói về Unicode, UTF-8 và thế giới xử lý
văn bản đa ngôn ngữ, nhưng nó có thể đợi đến một bài đăng khác.
Hiện tại, chúng tôi hy vọng bạn có hiểu biết tốt hơn về cách Go strings hoạt động
và mặc dù chúng có thể chứa các byte tùy ý, UTF-8 là một phần trung tâm
trong thiết kế của chúng.
