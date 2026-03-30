---
title: "Go Doc Comments"
layout: article
date: 2022-06-01T00:00:00Z
template: true
---

Mục lục:

 [Gói](#package)\
 [Lệnh](#cmd)\
 [Kiểu](#type)\
 [Hàm](#func)\
 [Hằng số](#const)\
 [Biến](#var)\
 [Cú pháp](#syntax)\
 [Lỗi phổ biến và cạm bẫy](#mistakes)

"Doc comment" là các comment xuất hiện ngay trước khai báo package, const, func, type và var ở cấp cao nhất mà không có dòng trống nào xen vào.
Mọi tên được xuất (viết hoa) đều nên có doc comment.

Các gói [go/doc](/pkg/go/doc) và [go/doc/comment](/pkg/go/doc/comment)
cung cấp khả năng trích xuất tài liệu từ mã nguồn Go,
và nhiều công cụ sử dụng chức năng này.
Lệnh [`go` `doc`](/cmd/go#hdr-Show_documentation_for_package_or_symbol)
tra cứu và in doc comment cho một gói hoặc symbol nhất định.
(Symbol là const, func, type hoặc var ở cấp cao nhất.)
Máy chủ web [pkg.go.dev](https://pkg.go.dev/) hiển thị tài liệu
cho các gói Go công khai (khi giấy phép của chúng cho phép điều đó).
Chương trình phục vụ trang đó là
[golang.org/x/pkgsite/cmd/pkgsite](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite),
cũng có thể được chạy cục bộ để xem tài liệu cho các module riêng tư
hoặc không có kết nối internet.
Máy chủ ngôn ngữ [gopls](https://pkg.go.dev/golang.org/x/tools/gopls)
cung cấp tài liệu khi chỉnh sửa tệp nguồn Go trong các IDE.

Phần còn lại của trang này ghi lại cách viết Go doc comment.

## Gói {#package}

Mọi gói đều nên có package comment giới thiệu gói.
Nó cung cấp thông tin liên quan đến toàn bộ gói
và thường đặt ra kỳ vọng cho gói.
Đặc biệt trong các gói lớn, package comment có thể hữu ích
khi cung cấp một tổng quan ngắn gọn về các phần quan trọng nhất của API,
liên kết đến các doc comment khác khi cần.

Nếu gói đơn giản, package comment có thể ngắn gọn.
Ví dụ:

	// Package path implements utility routines for manipulating slash-separated
	// paths.
	//
	// The path package should only be used for paths separated by forward
	// slashes, such as the paths in URLs. This package does not deal with
	// Windows paths with drive letters or backslashes; to manipulate
	// operating system paths, use the [path/filepath] package.
	package path

Dấu ngoặc vuông trong `[path/filepath]` tạo ra [liên kết tài liệu](#links).

Như có thể thấy trong ví dụ này, Go doc comment sử dụng câu hoàn chỉnh.
Đối với package comment, điều đó có nghĩa là [câu đầu tiên](/pkg/go/doc/#Package.Synopsis)
bắt đầu bằng "Package <name>".

Đối với các gói có nhiều tệp, package comment chỉ nên xuất hiện trong một tệp nguồn.
Nếu nhiều tệp có package comment, chúng được nối với nhau để tạo thành một
comment lớn duy nhất cho toàn bộ gói.

## Lệnh {#cmd}

Package comment cho một lệnh tương tự, nhưng nó mô tả hành vi
của chương trình thay vì các symbol Go trong gói.
Câu đầu tiên theo quy ước bắt đầu bằng tên của chương trình,
viết hoa vì nó đứng đầu câu.
Ví dụ, đây là phiên bản rút gọn của package comment cho [gofmt](/cmd/gofmt):

	/*
	Gofmt formats Go programs.
	It uses tabs for indentation and blanks for alignment.
	Alignment assumes that an editor is using a fixed-width font.

	Without an explicit path, it processes the standard input. Given a file,
	it operates on that file; given a directory, it operates on all .go files in
	that directory, recursively. (Files starting with a period are ignored.)
	By default, gofmt prints the reformatted sources to standard output.

	Usage:

		gofmt [flags] [path ...]

	The flags are:

		-d
			Do not print reformatted sources to standard output.
			If a file's formatting is different than gofmt's, print diffs
			to standard output.
		-w
			Do not print reformatted sources to standard output.
			If a file's formatting is different from gofmt's, overwrite it
			with gofmt's version. If an error occurred during overwriting,
			the original file is restored from an automatic backup.

	When gofmt reads from standard input, it accepts either a full Go program
	or a program fragment. A program fragment must be a syntactically
	valid declaration list, statement list, or expression. When formatting
	such a fragment, gofmt preserves leading indentation as well as leading
	and trailing spaces, so that individual sections of a Go program can be
	formatted by piping them through gofmt.
	*/
	package main

Phần đầu comment được viết sử dụng
[semantic linefeeds](https://rhodesmill.org/brandon/2012/one-sentence-per-line/),
trong đó mỗi câu mới hoặc cụm từ dài được đặt trên một dòng riêng,
có thể giúp các diff dễ đọc hơn khi mã và comment phát triển.
Các đoạn sau đó không tuân theo quy ước này
và đã được ngắt dòng thủ công.
Điều gì tốt nhất cho codebase của bạn là ổn.
Dù thế nào, `go` `doc` và `pkgsite` tự động ngắt dòng văn bản doc comment khi in.
Ví dụ:

	$ go doc gofmt
	Gofmt formats Go programs. It uses tabs for indentation and blanks for
	alignment. Alignment assumes that an editor is using a fixed-width font.

	Without an explicit path, it processes the standard input. Given a file, it
	operates on that file; given a directory, it operates on all .go files in that
	directory, recursively. (Files starting with a period are ignored.) By default,
	gofmt prints the reformatted sources to standard output.

	Usage:

		gofmt [flags] [path ...]

	The flags are:

		-d
			Do not print reformatted sources to standard output.
			If a file's formatting is different than gofmt's, print diffs
			to standard output.
	...

Các dòng thụt lề được coi là văn bản đã định dạng:
chúng không được ngắt dòng và được in bằng font code
trong các bản trình bày HTML và Markdown.
(Phần [Cú pháp](#syntax) bên dưới cung cấp chi tiết.)

## Kiểu {#type}

Doc comment của một kiểu nên giải thích mỗi instance của kiểu đó đại diện hoặc cung cấp gì.
Nếu API đơn giản, doc comment có thể rất ngắn.
Ví dụ:

	package zip

	// A Reader serves content from a ZIP archive.
	type Reader struct {
		...
	}

Theo mặc định, lập trình viên nên kỳ vọng rằng một kiểu chỉ an toàn khi sử dụng bởi
một goroutine tại một thời điểm.
Nếu một kiểu cung cấp các đảm bảo mạnh hơn, doc comment nên nêu chúng.
Ví dụ:

	package regexp

	// Regexp is the representation of a compiled regular expression.
	// A Regexp is safe for concurrent use by multiple goroutines,
	// except for configuration methods, such as Longest.
	type Regexp struct {
		...
	}

Các kiểu Go cũng nên hướng đến việc làm cho giá trị zero có ý nghĩa hữu ích.
Nếu điều đó không rõ ràng, ý nghĩa đó nên được ghi lại. Ví dụ:

	package bytes

	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	// The zero value for Buffer is an empty buffer ready to use.
	type Buffer struct {
		...
	}

Đối với struct có các trường được xuất, doc comment của kiểu hoặc comment từng trường
nên giải thích ý nghĩa của mỗi trường được xuất.
Ví dụ, doc comment của kiểu này giải thích các trường:

{{raw `
	package io

	// A LimitedReader reads from R but limits the amount of
	// data returned to just N bytes. Each call to Read
	// updates N to reflect the new amount remaining.
	// Read returns EOF when N <= 0.
	type LimitedReader struct {
		R   Reader // underlying reader
		N   int64  // max bytes remaining
	}
`}}

Ngược lại, doc comment của kiểu này để lại phần giải thích cho các comment từng trường:

{{raw `
	package comment

	// A Printer is a doc comment printer.
	// The fields in the struct can be filled in before calling
	// any of the printing methods
	// in order to customize the details of the printing process.
	type Printer struct {
		// HeadingLevel is the nesting level used for
		// HTML and Markdown headings.
		// If HeadingLevel is zero, it defaults to level 3,
		// meaning to use <h3> and ###.
		HeadingLevel int
		...
	}
`}}

Như với các gói (ở trên) và hàm (bên dưới), doc comment cho kiểu
bắt đầu bằng câu hoàn chỉnh đặt tên symbol được khai báo.
Chủ thể rõ ràng thường làm cho cách diễn đạt rõ ràng hơn,
và làm cho văn bản dễ tìm kiếm hơn, dù trên trang web
hay dòng lệnh.
Ví dụ:

	$ go doc -all regexp | grep pairs
	pairs within the input string: result[2*n:2*n+2] identifies the indexes
	    FindReaderSubmatchIndex returns a slice holding the index pairs identifying
	    FindStringSubmatchIndex returns a slice holding the index pairs identifying
	    FindSubmatchIndex returns a slice holding the index pairs identifying the
	$

## Hàm {#func}

Doc comment của một hàm nên giải thích hàm trả về gì
hoặc, đối với các hàm được gọi vì tác dụng phụ, nó làm gì.
Các tham số và kết quả có tên có thể được tham chiếu trực tiếp trong
comment, không cần cú pháp đặc biệt như backtick.
(Hệ quả của quy ước này là các tên như `a`,
có thể bị nhầm là từ thông thường, thường được tránh.)
Ví dụ:

	package strconv

	// Quote returns a double-quoted Go string literal representing s.
	// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100)
	// for control characters and non-printable characters as defined by IsPrint.
	func Quote(s string) string {
		...
	}

Và:

	package os

	// Exit causes the current program to exit with the given status code.
	// Conventionally, code zero indicates success, non-zero an error.
	// The program terminates immediately; deferred functions are not run.
	//
	// For portability, the status code should be in the range [0, 125].
	func Exit(code int) {
		...
	}

Doc comment thường dùng cụm từ "reports whether"
để mô tả các hàm trả về boolean.
Cụm từ "or not" là không cần thiết.
Ví dụ:

	package strings

	// HasPrefix reports whether the string s begins with prefix.
	func HasPrefix(s, prefix string) bool

Nếu doc comment cần giải thích nhiều kết quả,
việc đặt tên cho kết quả có thể giúp doc comment dễ hiểu hơn,
ngay cả khi các tên không được sử dụng trong thân hàm.
Ví dụ:

	package io

	// Copy copies from src to dst until either EOF is reached
	// on src or an error occurs. It returns the total number of bytes
	// written and the first error encountered while copying, if any.
	//
	// A successful Copy returns err == nil, not err == EOF.
	// Because Copy is defined to read from src until EOF, it does
	// not treat an EOF from Read as an error to be reported.
	func Copy(dst Writer, src Reader) (n int64, err error) {
		...
	}

Ngược lại, khi các kết quả không cần được đặt tên trong doc comment,
chúng thường cũng bị bỏ qua trong mã, như trong ví dụ `Quote` ở trên,
để tránh làm lộn xộn phần trình bày.

Các quy tắc này đều áp dụng cho cả hàm thông thường và phương thức.
Đối với phương thức, sử dụng cùng tên receiver tránh sự biến đổi không cần thiết
khi liệt kê tất cả phương thức của một kiểu:

	$ go doc bytes.Buffer
	package bytes // import "bytes"

	type Buffer struct {
		// Has unexported fields.
	}
	    A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	    The zero value for Buffer is an empty buffer ready to use.

	func NewBuffer(buf []byte) *Buffer
	func NewBufferString(s string) *Buffer
	func (b *Buffer) Bytes() []byte
	func (b *Buffer) Cap() int
	func (b *Buffer) Grow(n int)
	func (b *Buffer) Len() int
	func (b *Buffer) Next(n int) []byte
	func (b *Buffer) Read(p []byte) (n int, err error)
	func (b *Buffer) ReadByte() (byte, error)
	...

Ví dụ này cũng cho thấy rằng các hàm cấp cao nhất trả về kiểu `T` hoặc con trỏ `*T`,
có thể kèm thêm một kết quả lỗi,
được hiển thị cùng với kiểu `T` và các phương thức của nó,
dựa trên giả định rằng chúng là các constructor của `T`.

Theo mặc định, lập trình viên có thể giả định rằng hàm cấp cao nhất
an toàn để gọi từ nhiều goroutine;
thực tế này không cần phải nêu rõ.

Mặt khác, như đã lưu ý trong phần trước,
việc sử dụng một instance của một kiểu theo bất kỳ cách nào,
bao gồm cả việc gọi một phương thức, thường được giả định là
bị hạn chế cho một goroutine duy nhất tại một thời điểm.
Nếu các phương thức an toàn để dùng đồng thời
không được ghi lại trong doc comment của kiểu,
chúng nên được ghi lại trong các comment từng phương thức.
Ví dụ:

	package sql

	// Close returns the connection to the connection pool.
	// All operations after a Close will return with ErrConnDone.
	// Close is safe to call concurrently with other operations and will
	// block until all other operations finish. It may be useful to first
	// cancel any used context and then call Close directly after.
	func (c *Conn) Close() error {
		...
	}

Lưu ý rằng doc comment của hàm và phương thức tập trung vào
những gì thao tác trả về hoặc làm,
chi tiết những gì người gọi cần biết.
Các trường hợp đặc biệt có thể đặc biệt quan trọng khi ghi lại.
Ví dụ:

{{raw `
	package math

	// Sqrt returns the square root of x.
	//
	// Special cases are:
	//
	//	Sqrt(+Inf) = +Inf
	//	Sqrt(±0) = ±0
	//	Sqrt(x < 0) = NaN
	//	Sqrt(NaN) = NaN
	func Sqrt(x float64) float64 {
		...
	}
`}}

Doc comment không nên giải thích các chi tiết nội bộ
như thuật toán được sử dụng trong triển khai hiện tại.
Những chi tiết đó tốt nhất để lại cho các comment trong thân hàm.
Có thể thích hợp để đưa ra các giới hạn thời gian hoặc không gian tiệm cận
khi chi tiết đó đặc biệt quan trọng với người gọi.
Ví dụ:

	package sort

	// Sort sorts data in ascending order as determined by the Less method.
	// It makes one call to data.Len to determine n and O(n*log(n)) calls to
	// data.Less and data.Swap. The sort is not guaranteed to be stable.
	func Sort(data Interface) {
		...
	}

Vì doc comment này không đề cập đến thuật toán sắp xếp nào được sử dụng,
việc thay đổi triển khai để sử dụng thuật toán khác trong tương lai sẽ dễ dàng hơn.

## Hằng số {#const}

Cú pháp khai báo của Go cho phép nhóm các khai báo,
trong trường hợp đó một doc comment duy nhất có thể giới thiệu một nhóm các hằng số liên quan,
với các hằng số riêng lẻ chỉ được ghi lại bởi các comment ngắn ở cuối dòng.
Ví dụ:

	package scanner // import "text/scanner"

	// The result of Scan is one of these tokens or a Unicode character.
	const (
		EOF = -(iota + 1)
		Ident
		Int
		Float
		Char
		...
	)

Đôi khi nhóm không cần doc comment nào cả. Ví dụ:

	package unicode // import "unicode"

	const (
		MaxRune         = '\U0010FFFF' // maximum valid Unicode code point.
		ReplacementChar = '\uFFFD'     // represents invalid code points.
		MaxASCII        = '\u007F'     // maximum ASCII value.
		MaxLatin1       = '\u00FF'     // maximum Latin-1 value.
	)

Mặt khác, các hằng số không được nhóm thường cần một
doc comment đầy đủ bắt đầu bằng câu hoàn chỉnh. Ví dụ:

	package unicode

	// Version is the Unicode edition from which the tables are derived.
	const Version = "13.0.0"

Các hằng số có kiểu được hiển thị cạnh khai báo kiểu của chúng
và do đó thường bỏ qua doc comment của nhóm const để ưu tiên
doc comment của kiểu.
Ví dụ:

	package syntax

	// An Op is a single regular expression operator.
	type Op uint8

	const (
		OpNoMatch        Op = 1 + iota // matches no strings
		OpEmptyMatch                   // matches empty string
		OpLiteral                      // matches Runes sequence
		OpCharClass                    // matches Runes interpreted as range pair list
		OpAnyCharNotNL                 // matches any character except newline
		...
	)

(Xem [pkg.go.dev/regexp/syntax#Op](https://pkg.go.dev/regexp/syntax#Op) để biết phần trình bày HTML.)

## Biến {#var}

Các quy ước cho biến giống với các quy ước cho hằng số.
Ví dụ, đây là một tập hợp biến được nhóm:

	package fs

	// Generic file system errors.
	// Errors returned by file systems can be tested against these errors
	// using errors.Is.
	var (
		ErrInvalid    = errInvalid()    // "invalid argument"
		ErrPermission = errPermission() // "permission denied"
		ErrExist      = errExist()      // "file already exists"
		ErrNotExist   = errNotExist()   // "file does not exist"
		ErrClosed     = errClosed()     // "file already closed"
	)

Và một biến đơn:

	package unicode

	// Scripts is the set of Unicode script tables.
	var Scripts = map[string]*RangeTable{
		"Adlam":                  Adlam,
		"Ahom":                   Ahom,
		"Anatolian_Hieroglyphs":  Anatolian_Hieroglyphs,
		"Arabic":                 Arabic,
		"Armenian":               Armenian,
		...
	}

## Cú pháp {#syntax}

Go doc comment được viết theo cú pháp đơn giản hỗ trợ
đoạn văn, tiêu đề, liên kết, danh sách và khối mã đã định dạng.
Để giữ cho comment nhẹ và dễ đọc trong tệp nguồn,
không có hỗ trợ cho các tính năng phức tạp như thay đổi font hoặc HTML thô.
Những người quen với Markdown có thể coi cú pháp này như một tập con đơn giản hóa của Markdown.

Công cụ định dạng chuẩn [gofmt](/cmd/gofmt) định dạng lại doc comment
để sử dụng định dạng chuẩn cho mỗi tính năng này.
Gofmt nhắm đến tính dễ đọc và quyền kiểm soát của người dùng đối với cách comment
được viết trong mã nguồn nhưng sẽ điều chỉnh phần trình bày để làm cho
ý nghĩa ngữ nghĩa của một comment cụ thể rõ ràng hơn,
tương tự như việc định dạng lại `1+2 * 3` thành `1 + 2*3` trong mã nguồn thông thường.

Gofmt loại bỏ các dòng trống ở đầu và cuối trong doc comment.
Nếu tất cả các dòng trong doc comment bắt đầu bằng cùng một chuỗi
khoảng trắng và tab, gofmt loại bỏ tiền tố đó.

### Đoạn văn {#paragraphs}

Đoạn văn là một span các dòng không thụt lề, không trống.
Chúng ta đã thấy nhiều ví dụ về đoạn văn.

Một cặp backtick liên tiếp (\` U+0060)
được hiểu là dấu ngoặc kép trái Unicode (" U+201C),
và một cặp dấu nháy đơn liên tiếp (\' U+0027)
được hiểu là dấu ngoặc kép phải Unicode (" U+201D).

Gofmt giữ nguyên ngắt dòng trong văn bản đoạn: nó không tự động ngắt dòng.
Điều này cho phép sử dụng [semantic linefeeds](https://rhodesmill.org/brandon/2012/one-sentence-per-line/),
như đã thấy trước đó.
Gofmt thay thế các dòng trống trùng lặp giữa các đoạn
bằng một dòng trống duy nhất.
Gofmt cũng định dạng lại các backtick hoặc dấu nháy đơn liên tiếp
thành các giải thích Unicode của chúng.

#### Ghi chú {#notes}

Ghi chú là các comment đặc biệt có dạng `MARKER(uid): body`.
MARKER nên bao gồm 2 hoặc nhiều chữ cái in hoa `[A-Z]`,
xác định loại ghi chú, trong khi uid là ít nhất 1 ký tự,
thường là tên người dùng của người có thể cung cấp thêm thông tin.
`:` theo sau uid là tùy chọn.

Các ghi chú được thu thập và hiển thị trong phần riêng của chúng trên pkg.go.dev.

Ví dụ:

	// TODO(user1): refactor to use standard library context
	// BUG(user2): not cleaned up
	var ctx context.Context

#### Ngừng hỗ trợ {#deprecations}

Các đoạn bắt đầu bằng `Deprecated: ` được coi là thông báo ngừng hỗ trợ.
Một số công cụ sẽ cảnh báo khi các định danh đã ngừng hỗ trợ được sử dụng.
[pkg.go.dev](https://pkg.go.dev) sẽ ẩn tài liệu của chúng theo mặc định.

Thông báo ngừng hỗ trợ được theo sau bởi một số thông tin về việc ngừng hỗ trợ,
và khuyến nghị về những gì nên dùng thay thế, nếu có.
Đoạn không nhất thiết phải là đoạn cuối cùng trong doc comment.

Ví dụ:

	// Package rc4 implements the RC4 stream cipher.
	//
	// Deprecated: RC4 is cryptographically broken and should not be used
	// except for compatibility with legacy systems.
	//
	// This package is frozen and no new functionality will be added.
	package rc4

	// Reset zeros the key data and makes the Cipher unusable.
	//
	// Deprecated: Reset can't guarantee that the key will be entirely removed from
	// the process's memory.
	func (c *Cipher) Reset()

### Tiêu đề {#headings}

Tiêu đề là một dòng bắt đầu bằng ký hiệu số (U+0023) rồi một khoảng trắng và văn bản tiêu đề.
Để được nhận dạng là tiêu đề, dòng phải không thụt lề và được tách biệt khỏi văn bản đoạn liền kề
bằng các dòng trống.

Ví dụ:

	// Package strconv implements conversions to and from string representations
	// of basic data types.
	//
	// # Numeric Conversions
	//
	// The most common numeric conversions are [Atoi] (string to int) and [Itoa] (int to string).
	...
	package strconv

Mặt khác:

	// #This is not a heading, because there is no space.
	//
	// # This is not a heading,
	// # because it is multiple lines.
	//
	// # This is not a heading,
	// because it is also multiple lines.
	//
	// The next paragraph is not a heading, because there is no additional text:
	//
	// #
	//
	// In the middle of a span of non-blank lines,
	// # this is not a heading either.
	//
	//     # This is not a heading, because it is indented.

Cú pháp # được thêm vào trong Go 1.19.
Trước Go 1.19, tiêu đề được xác định ngầm định bởi các đoạn một dòng
thỏa mãn một số điều kiện nhất định, đáng chú ý nhất là thiếu bất kỳ dấu câu kết thúc nào.

Gofmt định dạng lại [các dòng được coi là tiêu đề ngầm định](https://github.com/golang/proposal/blob/master/design/51082-godocfmt.md#headings)
bởi các phiên bản Go cũ hơn để sử dụng tiêu đề # thay thế.
Nếu việc định dạng lại không phù hợp, tức là nếu dòng đó không có ý định là tiêu đề, cách dễ nhất
để biến nó thành đoạn là thêm dấu câu kết thúc
như dấu chấm hoặc dấu hai chấm, hoặc chia thành hai dòng.

### Liên kết {#links}

Một span các dòng không thụt lề, không trống xác định các mục tiêu liên kết
khi mỗi dòng có dạng "[Text]: URL".
Trong các văn bản khác trong cùng doc comment,
"[Text]" đại diện cho liên kết đến URL sử dụng văn bản đó, trong HTML là
\<a href="URL">Text\</a>.
Ví dụ:

	// Package json implements encoding and decoding of JSON as defined in
	// [RFC 7159]. The mapping between JSON and Go values is described
	// in the documentation for the Marshal and Unmarshal functions.
	//
	// For an introduction to this package, see the article
	// "[JSON and Go]."
	//
	// [RFC 7159]: https://tools.ietf.org/html/rfc7159
	// [JSON and Go]: https://golang.org/doc/articles/json_and_go.html
	package json

Bằng cách giữ URL trong một phần riêng,
định dạng này chỉ làm gián đoạn tối thiểu luồng của văn bản thực tế.
Nó cũng gần giống với định dạng
[shortcut reference link](https://spec.commonmark.org/0.30/#shortcut-reference-link) của Markdown,
không có văn bản tiêu đề tùy chọn.

Nếu không có khai báo URL tương ứng,
thì (ngoại trừ các doc link, được mô tả trong phần tiếp theo)
"[Text]" không phải là siêu liên kết, và dấu ngoặc vuông được giữ nguyên
khi hiển thị.
Mỗi doc comment được xem xét độc lập:
các định nghĩa mục tiêu liên kết trong một comment không ảnh hưởng đến các comment khác.

Mặc dù các khối định nghĩa mục tiêu liên kết có thể được xen kẽ với
các đoạn thông thường, gofmt di chuyển tất cả định nghĩa mục tiêu liên kết đến
cuối doc comment,
trong tối đa hai khối: đầu tiên là khối chứa tất cả các mục tiêu liên kết
được tham chiếu trong comment, rồi là khối chứa tất cả các mục tiêu _không_ được tham chiếu trong comment.
Khối riêng biệt giúp các mục tiêu không được dùng dễ
nhận thấy và sửa (trong trường hợp liên kết hoặc định nghĩa có lỗi đánh máy)
hoặc xóa (trong trường hợp các định nghĩa không còn cần thiết).

Văn bản thuần túy được nhận dạng là URL sẽ tự động được liên kết trong các bản hiển thị HTML.

### Doc link {#doclinks}

Doc link là các liên kết có dạng "[Name1]" hoặc "[Name1.Name2]" để tham chiếu
đến các định danh được xuất trong gói hiện tại, hoặc "[pkg]",
"[pkg.Name1]", hoặc "[pkg.Name1.Name2]" để tham chiếu đến các định danh trong các gói khác.

Ví dụ:

	package bytes

	// ReadFrom reads data from r until EOF and appends it to the buffer, growing
	// the buffer as needed. The return value n is the number of bytes read. Any
	// error except [io.EOF] encountered during the read is also returned. If the
	// buffer becomes too large, ReadFrom will panic with [ErrTooLarge].
	func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
		...
	}

Văn bản trong ngoặc vuông cho liên kết symbol
có thể bao gồm một dấu sao đứng đầu tùy chọn, giúp dễ dàng tham chiếu đến
các kiểu con trỏ, chẳng hạn như \[\*bytes.Buffer\].

Khi tham chiếu đến các gói khác, "pkg" có thể là đường dẫn import đầy đủ
hoặc tên gói giả định của một import hiện có. Tên gói giả định là
định danh trong import được đổi tên hoặc ngược lại là
[tên được giả định bởi
goimports](https://pkg.go.dev/golang.org/x/tools/internal/imports#ImportPathToAssumedName).
(Goimports chèn việc đổi tên khi giả định đó không đúng, vì vậy
quy tắc này sẽ hoạt động cho hầu như tất cả mã Go.)
Ví dụ, nếu gói hiện tại import encoding/json,
thì "[json.Decoder]" có thể được viết thay cho "[encoding/json.Decoder]"
để liên kết đến tài liệu cho Decoder của encoding/json.
Nếu các tệp nguồn khác nhau trong một gói import các gói khác nhau sử dụng cùng tên,
thì cú pháp ngắn gọn đó là mơ hồ và không thể sử dụng.

"pkg" chỉ được
giả định là đường dẫn import đầy đủ nếu nó bắt đầu bằng tên miền (một
phần tử đường dẫn có dấu chấm) hoặc là một trong các gói từ thư viện chuẩn
("[os]", "[encoding/json]", v.v.).
Ví dụ, `[os.File]` và `[example.com/sys.File]` là doc link tài liệu
(cái sau sẽ là liên kết hỏng),
nhưng `[os/sys.File]` thì không, vì không có gói os/sys trong thư viện chuẩn.

Để tránh các vấn đề với
map, generics và kiểu mảng, doc link phải được đứng trước và
theo sau bởi dấu câu, khoảng trắng, tab, hoặc đầu hoặc cuối của một dòng.
Ví dụ, văn bản "map[ast.Expr]TypeAndValue" không chứa
doc link.

### Danh sách {#lists}

Danh sách là một span các dòng thụt lề hoặc trống
(nếu không sẽ là khối mã,
như được mô tả trong phần tiếp theo)
trong đó dòng thụt lề đầu tiên bắt đầu bằng
dấu danh sách bullet hoặc dấu danh sách đánh số.

Dấu danh sách bullet là dấu sao, dấu cộng, gạch ngang hoặc bullet Unicode
(*, +, -, •; U+002A, U+002B, U+002D, U+2022)
theo sau là khoảng trắng hoặc tab rồi văn bản.
Trong danh sách bullet, mỗi dòng bắt đầu bằng dấu danh sách bullet
bắt đầu một mục danh sách mới.

Ví dụ:

	package url

	// PublicSuffixList provides the public suffix of a domain. For example:
	//   - the public suffix of "example.com" is "com",
	//   - the public suffix of "foo1.foo2.foo3.co.uk" is "co.uk", and
	//   - the public suffix of "bar.pvt.k12.ma.us" is "pvt.k12.ma.us".
	//
	// Implementations of PublicSuffixList must be safe for concurrent use by
	// multiple goroutines.
	//
	// An implementation that always returns "" is valid and may be useful for
	// testing but it is not secure: it means that the HTTP server for foo.com can
	// set a cookie for bar.com.
	//
	// A public suffix list implementation is in the package
	// golang.org/x/net/publicsuffix.
	type PublicSuffixList interface {
		...
	}

Dấu danh sách đánh số là một số thập phân bất kỳ độ dài
theo sau là dấu chấm hoặc dấu ngoặc đơn phải, rồi khoảng trắng hoặc tab, và văn bản.
Trong danh sách đánh số, mỗi dòng bắt đầu bằng dấu danh sách số bắt đầu một mục danh sách mới.
Số mục được giữ nguyên, không bao giờ được đánh lại.

Ví dụ:

	package path

	// Clean returns the shortest path name equivalent to path
	// by purely lexical processing. It applies the following rules
	// iteratively until no further processing can be done:
	//
	//  1. Replace multiple slashes with a single slash.
	//  2. Eliminate each . path name element (the current directory).
	//  3. Eliminate each inner .. path name element (the parent directory)
	//     along with the non-.. element that precedes it.
	//  4. Eliminate .. elements that begin a rooted path:
	//     that is, replace "/.." by "/" at the beginning of a path.
	//
	// The returned path ends in a slash only if it is the root "/".
	//
	// If the result of this process is an empty string, Clean
	// returns the string ".".
	//
	// See also Rob Pike, "[Lexical File Names in Plan 9]."
	//
	// [Lexical File Names in Plan 9]: https://9p.io/sys/doc/lexnames.html
	func Clean(path string) string {
		...
	}

Các mục danh sách chỉ chứa đoạn văn, không chứa khối mã hoặc danh sách lồng nhau.
Điều này tránh bất kỳ sự phức tạp về đếm khoảng trắng cũng như các câu hỏi về
bao nhiêu khoảng trắng một tab tính trong thụt lề không nhất quán.

Gofmt định dạng lại danh sách bullet để sử dụng gạch ngang làm dấu bullet,
hai khoảng trắng thụt lề trước gạch ngang,
và bốn khoảng trắng thụt lề cho các dòng tiếp theo.

Gofmt định dạng lại danh sách đánh số để sử dụng một khoảng trắng trước số,
dấu chấm sau số, và một lần nữa
bốn khoảng trắng thụt lề cho các dòng tiếp theo.

Gofmt giữ nguyên nhưng không yêu cầu một dòng trống giữa danh sách và đoạn trước đó.
Nó chèn một dòng trống giữa danh sách và đoạn hoặc tiêu đề tiếp theo.

### Khối mã {#code}

Khối mã là một span các dòng thụt lề hoặc trống
không bắt đầu bằng dấu danh sách bullet hoặc dấu danh sách đánh số.
Nó được hiển thị dưới dạng văn bản đã định dạng (khối \<pre> trong HTML).

Khối mã thường chứa mã Go. Ví dụ:

{{raw `
	package sort

	// Search uses binary search...
	//
	// As a more whimsical example, this program guesses your number:
	//
	//	func GuessingGame() {
	//		var s string
	//		fmt.Printf("Pick an integer from 0 to 100.\n")
	//		answer := sort.Search(100, func(i int) bool {
	//			fmt.Printf("Is your number <= %d? ", i)
	//			fmt.Scanf("%s", &s)
	//			return s != "" && s[0] == 'y'
	//		})
	//		fmt.Printf("Your number is %d.\n", answer)
	//	}
	func Search(n int, f func(int) bool) int {
		...
	}
`}}

Tất nhiên, khối mã cũng thường chứa văn bản đã định dạng ngoài mã. Ví dụ:

{{raw `
	package path

	// Match reports whether name matches the shell pattern.
	// The pattern syntax is:
	//
	//	pattern:
	//		{ term }
	//	term:
	//		'*'         matches any sequence of non-/ characters
	//		'?'         matches any single non-/ character
	//		'[' [ '^' ] { character-range } ']'
	//		            character class (must be non-empty)
	//		c           matches character c (c != '*', '?', '\\', '[')
	//		'\\' c      matches character c
	//
	//	character-range:
	//		c           matches character c (c != '\\', '-', ']')
	//		'\\' c      matches character c
	//		lo '-' hi   matches character c for lo <= c <= hi
	//
	// Match requires pattern to match all of name, not just a substring.
	// The only possible returned error is [ErrBadPattern], when pattern
	// is malformed.
	func Match(pattern, name string) (matched bool, err error) {
		...
	}
`}}

Gofmt thụt lề tất cả các dòng trong khối mã bằng một tab duy nhất,
thay thế bất kỳ thụt lề nào khác mà các dòng không trống có điểm chung.
Gofmt cũng chèn một dòng trống trước và sau mỗi khối mã,
phân biệt rõ ràng khối mã với văn bản đoạn xung quanh.

### Chỉ thị {#directives}

Các comment chỉ thị như `//go:generate` không được
coi là một phần của doc comment và bị bỏ qua khỏi
tài liệu đã hiển thị.
Gofmt di chuyển các comment chỉ thị đến cuối doc comment,
đứng trước bởi một dòng trống.
Ví dụ:

	package regexp

	// An Op is a single regular expression operator.
	//
	//go:generate stringer -type Op -trimprefix Op
	type Op uint8

Comment chỉ thị là một dòng bắt đầu bằng biểu thức chính quy
`//(line |extern |export |[a-z0-9]+:[a-z0-9])`.

Các công cụ có thể xác định các comment chỉ thị của riêng chúng sử dụng dạng
`//toolname:directive arguments`.
Các chỉ thị công cụ khớp với biểu thức chính quy
`//([a-z0-9]+):([a-z0-9]\PZ*)($|\pZ+)(.*)`, trong đó nhóm đầu tiên
là tên công cụ và nhóm thứ hai là tên chỉ thị.
Các đối số tùy chọn được tách biệt khỏi tên chỉ thị bởi
một hoặc nhiều ký tự khoảng trắng Unicode.
Mỗi công cụ có thể xác định cú pháp đối số riêng, nhưng một quy ước phổ biến là một
chuỗi các đối số phân tách bởi khoảng trắng, trong đó đối số có thể là
một từ thuần túy, hoặc chuỗi Go được trích dẫn kép hoặc trích dẫn backtick.
Tên công cụ `go` được dành riêng để sử dụng bởi Go toolchain.

Hàm [`go/ast.ParseDirective`](/pkg/go/ast#ParseDirective) và các
kiểu liên quan của nó phân tích cú pháp chỉ thị công cụ.

## Lỗi phổ biến và cạm bẫy {#mistakes}

Quy tắc rằng bất kỳ span dòng thụt lề hoặc trống nào
trong doc comment đều được hiển thị dưới dạng khối mã
có từ những ngày đầu của Go.
Thật không may, việc thiếu hỗ trợ cho doc comment trong gofmt
đã dẫn đến nhiều comment hiện có sử dụng thụt lề
mà không có ý định tạo khối mã.

Ví dụ, danh sách không thụt lề này luôn được godoc hiểu là đoạn ba dòng theo sau là khối mã một dòng:

	package http

	// cancelTimerBody is an io.ReadCloser that wraps rc with two features:
	// 1) On Read error or close, the stop func is called.
	// 2) On Read failure, if reqDidTimeout is true, the error is wrapped and
	//    marked as net.Error that hit its timeout.
	type cancelTimerBody struct {
		...
	}

Điều này luôn hiển thị trong `go` `doc` là:

	cancelTimerBody is an io.ReadCloser that wraps rc with two features:
	1) On Read error or close, the stop func is called. 2) On Read failure,
	if reqDidTimeout is true, the error is wrapped and

	    marked as net.Error that hit its timeout.

Tương tự, lệnh trong comment này là đoạn một dòng
theo sau là khối mã một dòng:

	package smtp

	// localhostCert is a PEM-encoded TLS cert generated from src/crypto/tls:
	//
	// go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \
	//     --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
	var localhostCert = []byte(`...`)

Điều này hiển thị trong `go` `doc` là:

	localhostCert is a PEM-encoded TLS cert generated from src/crypto/tls:

	go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \

	    --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h

Và comment này là đoạn hai dòng (dòng thứ hai là "{"),
theo sau là khối mã sáu dòng thụt lề và đoạn một dòng ("}").

	// On the wire, the JSON will look something like this:
	// {
	//	"kind":"MyAPIObject",
	//	"apiVersion":"v1",
	//	"myPlugin": {
	//		"kind":"PluginA",
	//		"aOption":"foo",
	//	},
	// }

Và điều này hiển thị trong `go` `doc` là:

	On the wire, the JSON will look something like this: {

	    "kind":"MyAPIObject",
	    "apiVersion":"v1",
	    "myPlugin": {
	    	"kind":"PluginA",
	    	"aOption":"foo",
	    },

	}

Một lỗi phổ biến khác là định nghĩa hàm Go hoặc câu lệnh khối không thụt lề,
tương tự được bao bởi "{" và "}".

Việc giới thiệu định dạng lại doc comment trong gofmt của Go 1.19 làm cho các lỗi
như thế này dễ nhận thấy hơn bằng cách thêm các dòng trống xung quanh các khối mã.

Phân tích năm 2022 phát hiện rằng chỉ có 3% doc comment trong các module Go công khai
được định dạng lại hoàn toàn bởi bản nháp gofmt Go 1.19.
Giới hạn ở các comment đó, khoảng 87% việc định dạng lại của gofmt
đã bảo tồn cấu trúc mà một người sẽ suy ra khi đọc comment;
khoảng 6% bị vấp phải bởi các loại danh sách không thụt lề,
các lệnh shell nhiều dòng không thụt lề, và các khối mã phân tách bởi ngoặc nhọn không thụt lề.

Dựa trên phân tích này, gofmt Go 1.19 áp dụng một vài heuristic để hợp nhất
các dòng không thụt lề vào danh sách hoặc khối mã thụt lề liền kề.
Với những điều chỉnh đó, gofmt Go 1.19 định dạng lại các ví dụ trên thành:

	// cancelTimerBody is an io.ReadCloser that wraps rc with two features:
	//  1. On Read error or close, the stop func is called.
	//  2. On Read failure, if reqDidTimeout is true, the error is wrapped and
	//     marked as net.Error that hit its timeout.

	// localhostCert is a PEM-encoded TLS cert generated from src/crypto/tls:
	//
	//	go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \
	//	    --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h

	// On the wire, the JSON will look something like this:
	//
	//	{
	//		"kind":"MyAPIObject",
	//		"apiVersion":"v1",
	//		"myPlugin": {
	//			"kind":"PluginA",
	//			"aOption":"foo",
	//		},
	//	}

Việc định dạng lại này làm cho ý nghĩa rõ ràng hơn cũng như làm cho các doc comment
hiển thị đúng trong các phiên bản Go cũ hơn.
Nếu heuristic đôi khi đưa ra quyết định sai, nó có thể được ghi đè bằng cách chèn
một dòng trống để tách rõ ràng văn bản đoạn khỏi văn bản không phải đoạn.

Ngay cả với những heuristic này, các comment hiện có khác vẫn sẽ cần điều chỉnh thủ công
để sửa cách hiển thị của chúng.
Lỗi phổ biến nhất là thụt lề một dòng văn bản được ngắt không thụt lề.
Ví dụ:

	// TODO Revisit this design. It may make sense to walk those nodes
	//      only once.

	// According to the document:
	// "The alignment factor (in bytes) that is used to align the raw data of sections in
	//  the image file. The value should be a power of 2 between 512 and 64 K, inclusive."

Trong cả hai trường hợp này, dòng cuối cùng bị thụt lề, khiến nó trở thành khối mã.
Cách sửa là bỏ thụt lề các dòng.

Một lỗi phổ biến khác là không thụt lề dòng tiếp theo được ngắt của một danh sách hoặc khối mã thụt lề.
Ví dụ:

	// Uses of this error model include:
	//
	//   - Partial errors. If a service needs to return partial errors to the
	// client,
	//     it may embed the `Status` in the normal response to indicate the
	// partial
	//     errors.
	//
	//   - Workflow errors. A typical workflow has multiple steps. Each step
	// may
	//     have a `Status` message for error reporting.

Cách sửa là thụt lề các dòng tiếp theo.

Go doc comment không hỗ trợ danh sách lồng nhau, vì vậy gofmt định dạng lại

	// Here is a list:
	//
	//  - Item 1.
	//    * Subitem 1.
	//    * Subitem 2.
	//  - Item 2.
	//  - Item 3.

thành

	// Here is a list:
	//
	//  - Item 1.
	//  - Subitem 1.
	//  - Subitem 2.
	//  - Item 2.
	//  - Item 3.
