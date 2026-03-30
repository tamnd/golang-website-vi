---
title: Chuẩn hóa văn bản trong Go
date: 2013-11-26
by:
- Marcel van Lohuizen
tags:
- strings
- bytes
- runes
- characters
summary: Cách và lý do chuẩn hóa văn bản UTF-8 trong Go.
template: true
---

## Giới thiệu

Một [bài trước](/blog/strings) đã nói về chuỗi, byte
và ký tự trong Go. Tôi đã làm việc trên nhiều gói xử lý văn bản đa ngôn ngữ
cho kho go.text. Một số gói trong số này xứng đáng có một
bài đăng blog riêng, nhưng hôm nay tôi muốn tập trung vào
[go.text/unicode/norm](https://pkg.go.dev/golang.org/x/text/unicode/norm),
xử lý chuẩn hóa, một chủ đề được đề cập trong
[bài viết về chuỗi](/blog/strings) và là chủ đề của bài đăng này.
Chuẩn hóa hoạt động ở cấp độ trừu tượng cao hơn so với byte thô.

Để tìm hiểu hầu hết mọi thứ bạn muốn biết về chuẩn hóa
(và thậm chí nhiều hơn thế), [Phụ lục 15 của Tiêu chuẩn Unicode](http://unicode.org/reports/tr15/)
là tài liệu đáng đọc. Một bài viết dễ tiếp cận hơn là
[trang Wikipedia tương ứng](http://en.wikipedia.org/wiki/Unicode_equivalence). Ở đây chúng ta
tập trung vào cách chuẩn hóa liên quan đến Go.

## Chuẩn hóa là gì?

Thường có nhiều cách để biểu diễn cùng một chuỗi. Ví dụ, chữ é
(e-acute) có thể được biểu diễn trong một chuỗi dưới dạng một rune đơn ("\u00e9") hoặc chữ 'e'
theo sau bởi dấu mũ ("\u0301"). Theo tiêu chuẩn Unicode,
hai cái này là "tương đương theo chuẩn" và nên được coi là bằng nhau.

Sử dụng so sánh byte-to-byte để xác định sự bằng nhau rõ ràng sẽ không cho
kết quả đúng cho hai chuỗi này. Unicode định nghĩa một tập hợp các dạng chuẩn
sao cho nếu hai chuỗi tương đương theo chuẩn và được chuẩn hóa về
cùng một dạng chuẩn, thì biểu diễn byte của chúng là giống nhau.

Unicode cũng định nghĩa "tương đương tương thích" để coi các ký tự
biểu diễn cùng ký tự là bằng nhau, nhưng có thể có giao diện hình ảnh khác nhau. Ví dụ,
chữ số mũ '⁹' và chữ số thông thường '9' tương đương trong
dạng này.

Cho mỗi một trong hai dạng tương đương này, Unicode định nghĩa một dạng tổng hợp và
dạng phân tích. Dạng trước thay thế các rune có thể kết hợp thành một rune đơn
bằng rune đơn đó. Dạng sau phân tách các rune thành các thành phần của chúng.
Bảng này hiển thị các tên, tất cả bắt đầu bằng NF, theo đó
Unicode Consortium xác định các dạng này:

{{raw (file "normalization/table1.html")}}

## Cách tiếp cận của Go với chuẩn hóa

Như đã đề cập trong bài đăng blog về chuỗi, Go không đảm bảo rằng các ký tự trong
một chuỗi được chuẩn hóa. Tuy nhiên, các gói go.text có thể bù đắp cho điều đó. Ví dụ, gói
[collate](https://pkg.go.dev/golang.org/x/text/collate), có thể
sắp xếp chuỗi theo cách đặc thù của ngôn ngữ, hoạt động chính xác ngay cả với
các chuỗi không được chuẩn hóa. Các gói trong go.text không phải lúc nào cũng yêu cầu đầu vào đã được chuẩn hóa,
nhưng nhìn chung chuẩn hóa có thể cần thiết để có kết quả nhất quán.

Chuẩn hóa không miễn phí nhưng nhanh, đặc biệt đối với collation và
tìm kiếm hoặc nếu một chuỗi đã ở dạng NFD hoặc NFC và có thể được chuyển đổi sang NFD
bằng cách phân tích mà không cần sắp xếp lại byte. Trong thực tế,
[99,98%](http://www.macchiato.com/unicode/nfc-faq#TOC-How-much-text-is-already-NFC-) nội dung
trang HTML trên web ở dạng NFC (không tính đánh dấu, trong trường hợp đó
sẽ còn nhiều hơn). Cho đến nay hầu hết NFC có thể được phân tích thành NFD mà không cần
sắp xếp lại (vốn yêu cầu phân bổ). Ngoài ra, việc phát hiện
khi nào cần sắp xếp lại là hiệu quả, vì vậy chúng ta có thể tiết kiệm thời gian bằng cách làm điều đó chỉ cho các phân đoạn hiếm gặp
cần nó.

Để làm cho mọi thứ tốt hơn nữa, gói collation thường không sử dụng
gói norm trực tiếp, mà thay vào đó sử dụng gói norm để xen kẽ
thông tin chuẩn hóa với các bảng của riêng nó. Xen kẽ hai vấn đề
cho phép sắp xếp lại và chuẩn hóa ngay lúc chạy với hầu như không ảnh hưởng đến
hiệu năng. Chi phí của việc chuẩn hóa ngay lúc chạy được bù đắp bằng cách không phải
chuẩn hóa văn bản trước và đảm bảo rằng dạng chuẩn được duy trì
khi có chỉnh sửa. Điều sau có thể phức tạp. Ví dụ, kết quả của việc ghép
hai chuỗi được chuẩn hóa NFC không được đảm bảo là ở dạng NFC.

Tất nhiên, chúng ta cũng có thể tránh hoàn toàn phần chi phí nếu chúng ta biết trước rằng một
chuỗi đã được chuẩn hóa, đây thường là trường hợp.

## Tại sao lại lo lắng?

Sau tất cả cuộc thảo luận này về việc tránh chuẩn hóa, bạn có thể hỏi tại sao
đáng lo lắng về điều đó. Lý do là có những trường hợp
chuẩn hóa là bắt buộc và điều quan trọng là phải hiểu những trường hợp đó là gì, và
theo đó cách làm đúng.

Trước khi thảo luận về những điều đó, trước tiên chúng ta phải làm rõ khái niệm 'ký tự'.

## Ký tự là gì?

Như đã đề cập trong bài đăng blog về chuỗi, các ký tự có thể trải dài qua nhiều rune.
Ví dụ, chữ 'e' và '◌́' (dấu mũ "\u0301") có thể kết hợp để tạo thành 'é' ("e\u0301"
trong NFD). Cùng nhau, hai rune này là một ký tự. Định nghĩa về
ký tự có thể thay đổi tùy thuộc vào ứng dụng. Đối với chuẩn hóa, chúng ta sẽ
định nghĩa nó là một dãy rune bắt đầu bằng một starter, một rune không
sửa đổi hoặc kết hợp ngược với bất kỳ rune nào khác, theo sau là một
dãy có thể rỗng các non-starter, tức là, các rune có (thường là dấu phụ). Thuật toán
chuẩn hóa xử lý một ký tự tại một thời điểm.

Về mặt lý thuyết, không có giới hạn về số lượng rune có thể tạo nên một
ký tự Unicode. Thực tế, không có hạn chế nào về số lượng bộ sửa đổi có thể theo sau một ký tự
và một bộ sửa đổi có thể được lặp lại, hoặc xếp chồng. Bạn đã từng thấy chữ 'e' với ba dấu mũ chưa?
Đây là nó: 'é́́'. Đó là một ký tự 4-rune hoàn toàn hợp lệ theo tiêu chuẩn.

Do đó, ngay cả ở cấp độ thấp nhất, văn bản cần được xử lý theo
các đoạn có kích thước không giới hạn. Điều này đặc biệt khó xử với phương pháp
xử lý văn bản theo luồng, như được sử dụng bởi các interface Reader và
Writer chuẩn của Go, vì mô hình đó có khả năng yêu cầu bất kỳ bộ đệm trung gian nào
cũng phải có kích thước không giới hạn. Ngoài ra, một triển khai đơn giản của
chuẩn hóa sẽ có thời gian chạy O(n²).

Thực sự không có cách hiểu có ý nghĩa nào cho các dãy bộ sửa đổi lớn như vậy đối với các ứng dụng thực tế.
Unicode định nghĩa định dạng Văn bản An toàn Theo Luồng, cho phép giới hạn số lượng bộ sửa đổi (non-starter) tối đa
là 30, đủ cho bất kỳ mục đích thực tế nào. Các bộ sửa đổi tiếp theo sẽ được
đặt sau một Combining Grapheme Joiner (CGJ hoặc U+034F) được chèn vào. Go
áp dụng cách tiếp cận này cho tất cả các thuật toán chuẩn hóa. Quyết định này từ bỏ một
chút tuân thủ nhưng đạt được một chút an toàn.

## Viết ở dạng chuẩn

Ngay cả khi bạn không cần chuẩn hóa văn bản trong code Go của mình, bạn có thể vẫn
muốn làm như vậy khi giao tiếp với thế giới bên ngoài. Ví dụ, chuẩn hóa
sang NFC có thể nén văn bản của bạn, làm cho việc gửi đi rẻ hơn. Đối với một số
ngôn ngữ, như tiếng Hàn, tiết kiệm có thể đáng kể. Ngoài ra, một số API bên ngoài
có thể kỳ vọng văn bản ở một dạng chuẩn nhất định. Hoặc bạn chỉ có thể muốn hòa mình
và xuất văn bản của mình dưới dạng NFC như phần còn lại của thế giới.

Để viết văn bản của bạn dưới dạng NFC, sử dụng gói
[unicode/norm](https://pkg.go.dev/golang.org/x/text/unicode/norm)
để bọc `io.Writer` bạn chọn:

	wc := norm.NFC.Writer(w)
	defer wc.Close()
	// viết như trước...

Nếu bạn có một chuỗi nhỏ và muốn chuyển đổi nhanh, bạn có thể sử dụng biểu mẫu đơn giản hơn này:

	norm.NFC.Bytes(b)

Gói norm cung cấp nhiều phương thức khác để chuẩn hóa văn bản.
Hãy chọn phương thức phù hợp nhất với nhu cầu của bạn.

## Phát hiện các ký tự trông giống nhau

Bạn có thể phân biệt được sự khác biệt giữa 'K' ("\u004B") và 'K' (dấu hiệu Kelvin
"\u212A") hoặc 'Ω' ("\u03a9") và 'Ω' (dấu hiệu Ohm "\u2126") không? Thật dễ bỏ qua
sự khác biệt đôi khi nhỏ giữa các biến thể của cùng một ký tự cơ bản.
Nhìn chung, đây là ý tưởng tốt để không cho phép các biến thể như vậy trong các định danh
hoặc bất cứ thứ gì mà việc lừa dối người dùng bằng các ký tự trông giống như vậy có thể gây ra rủi ro bảo mật.

Các dạng chuẩn tương thích, NFKC và NFKD, sẽ ánh xạ nhiều dạng trông gần như
giống nhau về mặt hình ảnh thành một giá trị duy nhất. Lưu ý rằng nó sẽ không làm như vậy khi hai ký hiệu
trông giống nhau, nhưng thực sự từ hai bảng chữ cái khác nhau. Ví dụ, chữ Latinh
'o', chữ Hy Lạp 'ο' và chữ Cyrillic 'о' vẫn là các ký tự khác nhau theo các dạng này.

## Sửa đổi văn bản đúng cách

Gói norm cũng có thể giúp ích khi cần sửa đổi văn bản.
Hãy xem xét một trường hợp bạn muốn tìm kiếm và thay thế từ "cafe" bằng dạng số nhiều "cafes".
Một đoạn code có thể trông như thế này.

	s := "We went to eat at multiple cafe"
	cafe := "cafe"
	if p := strings.Index(s, cafe); p != -1 {
		p += len(cafe)
		s = s[:p] + "s" + s[p:]
	}
	fmt.Println(s)

Điều này in ra "We went to eat at multiple cafes" như mong muốn và kỳ vọng. Bây giờ
hãy xem xét văn bản của chúng ta chứa chính tả tiếng Pháp "café" ở dạng NFD:

	s := "We went to eat at multiple cafe\u0301"

Sử dụng cùng code từ trên, "s" số nhiều vẫn sẽ được chèn sau
chữ 'e', nhưng trước dấu mũ, dẫn đến "We went to eat at multiple
cafeś". Hành vi này là không mong muốn.

Vấn đề là code không tôn trọng ranh giới giữa các ký tự nhiều rune và chèn một rune vào giữa một ký tự. Sử dụng gói norm,
chúng ta có thể viết lại đoạn code này như sau:

	s := "We went to eat at multiple cafe\u0301"
	cafe := "cafe"
	if p := strings.Index(s, cafe); p != -1 {
		p += len(cafe)
		if bp := norm.FirstBoundary(s[p:]); bp > 0 {
			p += bp
		}
		s = s[:p] + "s" + s[p:]
	}
	fmt.Println(s)

Đây có thể là một ví dụ được tạo ra, nhưng ý nghĩa nên rõ ràng. Hãy chú ý đến
thực tế là các ký tự có thể trải dài qua nhiều rune. Nhìn chung, các vấn đề loại
này có thể tránh được bằng cách sử dụng chức năng tìm kiếm tôn trọng ranh giới ký tự
(chẳng hạn như gói go.text/search đã được lên kế hoạch).

## Lặp lại

Một công cụ khác được cung cấp bởi gói norm có thể giúp xử lý ranh giới ký tự
là iterator của nó,
[`norm.Iter`](https://pkg.go.dev/golang.org/x/text/unicode/norm#Iter).
Nó lặp qua các ký tự từng cái một trong dạng chuẩn được chọn.

## Thực hiện điều kỳ diệu

Như đã đề cập trước đó, hầu hết văn bản ở dạng NFC, nơi các ký tự cơ bản và
bộ sửa đổi được kết hợp thành một rune duy nhất khi có thể. Với mục đích
phân tích ký tự, thường dễ hơn khi xử lý các rune sau khi phân tích
thành các thành phần nhỏ nhất của chúng. Đây là nơi dạng NFD phát huy tác dụng. Ví dụ,
đoạn code sau tạo ra một `transform.Transformer` phân tích
văn bản thành các phần nhỏ nhất của nó, xóa tất cả các dấu phụ, và sau đó
tổng hợp lại văn bản thành NFC:

	import (
		"unicode"

		"golang.org/x/text/transform"
		"golang.org/x/text/unicode/norm"
	)

	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

`Transformer` kết quả có thể được sử dụng để xóa dấu phụ từ một `io.Reader`
bất kỳ như sau:

	r = transform.NewReader(r, t)
	// đọc như trước ...

Ví dụ, điều này sẽ chuyển đổi bất kỳ đề cập nào về "cafés" trong văn bản thành "cafes",
bất kể dạng chuẩn mà văn bản gốc được mã hóa.

## Thông tin chuẩn hóa

Như đã đề cập trước đó, một số gói tính toán trước các chuẩn hóa vào các bảng của chúng
để giảm thiểu nhu cầu chuẩn hóa tại thời điểm chạy. Kiểu `norm.Properties`
cung cấp quyền truy cập vào thông tin per-rune cần thiết bởi các gói này, đáng chú ý nhất là
Lớp Kết hợp Chuẩn và thông tin phân tích. Đọc
[tài liệu](https://pkg.go.dev/golang.org/x/text/unicode/norm#Properties)
cho kiểu này nếu bạn muốn tìm hiểu sâu hơn.

## Hiệu năng

Để hình dung hiệu năng của chuẩn hóa, chúng ta so sánh nó với
hiệu năng của strings.ToLower. Mẫu trong hàng đầu tiên vừa viết thường vừa là NFC và có thể
được trả về nguyên vẹn trong mọi trường hợp. Mẫu thứ hai thì không và
yêu cầu viết một phiên bản mới.

{{raw (file "normalization/table2.html")}}

Cột với kết quả cho iterator hiển thị cả phép đo với
và không có khởi tạo iterator, chứa các bộ đệm không cần
khởi tạo lại khi tái sử dụng.

Như bạn có thể thấy, việc phát hiện xem một chuỗi có được chuẩn hóa không có thể khá
hiệu quả. Nhiều chi phí của việc chuẩn hóa trong hàng thứ hai là do
khởi tạo bộ đệm, chi phí của nó được phân bổ khi xử lý các chuỗi lớn hơn.
Hóa ra, các bộ đệm này hiếm khi cần đến, vì vậy chúng ta có thể thay đổi triển khai vào
một thời điểm nào đó để tăng tốc trường hợp phổ biến cho các chuỗi nhỏ thêm nữa.

## Kết luận

Nếu bạn đang xử lý văn bản trong Go, nhìn chung bạn không cần sử dụng gói
unicode/norm để chuẩn hóa văn bản của mình. Gói này vẫn có thể hữu ích
cho những thứ như đảm bảo rằng các chuỗi được chuẩn hóa trước khi gửi chúng đi hoặc
để thực hiện các thao tác văn bản nâng cao.

Bài viết này đã đề cập ngắn gọn về sự tồn tại của các gói go.text khác cũng như
xử lý văn bản đa ngôn ngữ và có thể đã nêu ra nhiều câu hỏi hơn là
câu trả lời. Tuy nhiên, việc thảo luận về các chủ đề này sẽ phải
chờ đến một ngày khác.
