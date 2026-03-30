---
title: Độ phủ mã cho kiểm thử tích hợp Go
date: 2023-03-08
by:
- Than McIntosh
summary: Độ phủ mã cho kiểm thử tích hợp, có sẵn trong Go 1.20.
template: true
---

Các công cụ độ phủ mã giúp nhà phát triển xác định tỷ lệ mã nguồn được thực thi (được bao phủ) khi một bộ kiểm thử nhất định chạy.

Go từ lâu đã cung cấp hỗ trợ ([được giới thiệu](/blog/cover) trong bản phát hành Go 1.2) để đo độ phủ mã ở cấp độ gói, sử dụng cờ **"-cover"** của lệnh "go test".

Hệ thống công cụ này hoạt động tốt trong hầu hết các trường hợp, nhưng có một số điểm yếu đối với các ứng dụng Go lớn hơn.
Với các ứng dụng như vậy, các nhà phát triển thường viết các kiểm thử "tích hợp" để xác minh hành vi của toàn bộ chương trình (ngoài các kiểm thử đơn vị cấp gói).

Loại kiểm thử này thường liên quan đến việc xây dựng một tệp nhị phân ứng dụng hoàn chỉnh, sau đó chạy tệp nhị phân đó trên một tập các đầu vào đại diện (hoặc dưới tải production nếu là server) để đảm bảo rằng tất cả các gói thành phần đang hoạt động đúng cùng nhau, thay vì kiểm thử từng gói riêng lẻ.

Vì các tệp nhị phân kiểm thử tích hợp được xây dựng bằng "go build" chứ không phải "go test", hệ thống công cụ của Go không cung cấp cách dễ dàng nào để thu thập profile độ phủ cho các kiểm thử này, cho đến tận bây giờ.

Với Go 1.20, bạn có thể xây dựng các chương trình được đo lường độ phủ bằng "go build -cover", sau đó đưa các tệp nhị phân được đo lường này vào một kiểm thử tích hợp để mở rộng phạm vi kiểm thử độ phủ.

Trong bài viết blog này, chúng tôi sẽ đưa ra ví dụ về cách các tính năng mới này hoạt động, và phác thảo một số trường hợp sử dụng và quy trình làm việc để thu thập profile độ phủ từ các kiểm thử tích hợp.

## Ví dụ

Hãy lấy một chương trình mẫu rất nhỏ, viết một kiểm thử tích hợp đơn giản cho nó, rồi thu thập một profile độ phủ từ kiểm thử tích hợp.

Cho bài tập này, chúng tôi sẽ dùng công cụ xử lý markdown "mdtool" từ [`gitlab.com/golang-commonmark/mdtool`](https://pkg.go.dev/gitlab.com/golang-commonmark/mdtool).
Đây là chương trình demo được thiết kế để chỉ cách các client có thể dùng gói [`gitlab.com/golang-commonmark/markdown`](https://pkg.go.dev/gitlab.com/golang-commonmark/markdown), một thư viện chuyển đổi markdown sang HTML.

## Thiết lập cho mdtool

Đầu tiên hãy tải về một bản sao của "mdtool" (chúng tôi chọn một phiên bản cụ thể để các bước này có thể tái tạo được):

```
$ git clone https://gitlab.com/golang-commonmark/mdtool.git
...
$ cd mdtool
$ git tag example e210a4502a825ef7205691395804eefce536a02f
$ git checkout example
...
$
```

## Một kiểm thử tích hợp đơn giản

Bây giờ chúng tôi sẽ viết một kiểm thử tích hợp đơn giản cho "mdtool"; kiểm thử của chúng tôi sẽ xây dựng tệp nhị phân "mdtool", rồi chạy nó trên một tập các tệp markdown đầu vào.
Script rất đơn giản này chạy tệp nhị phân "mdtool" trên mỗi tệp từ thư mục dữ liệu kiểm thử, kiểm tra để đảm bảo rằng nó tạo ra một số đầu ra và không bị crash.

```
$ cat integration_test.sh
#!/bin/sh
BUILDARGS="$*"
#
# Terminate the test if any command below does not complete successfully.
#
set -e
#
# Download some test inputs (the 'website' repo contains various *.md files).
#
if [ ! -d testdata ]; then
  git clone https://go.googlesource.com/website testdata
  git -C testdata tag example 8bb4a56901ae3b427039d490207a99b48245de2c
  git -C testdata checkout example
fi
#
# Build mdtool binary for testing purposes.
#
rm -f mdtool.exe
go build $BUILDARGS -o mdtool.exe .
#
# Run the tool on a set of input files from 'testdata'.
#
FILES=$(find testdata -name "*.md" -print)
N=$(echo $FILES | wc -w)
for F in $FILES
do
  ./mdtool.exe +x +a $F > /dev/null
done
echo "finished processing $N files, no crashes"
$
```

Đây là ví dụ chạy kiểm thử của chúng tôi:

```
$ /bin/sh integration_test.sh
...
finished processing 380 files, no crashes
$
```

Thành công: chúng tôi đã xác minh rằng tệp nhị phân "mdtool" xử lý thành công một tập các tệp đầu vào... nhưng chúng tôi thực sự đã thực thi bao nhiêu mã nguồn của công cụ?
Trong phần tiếp theo, chúng tôi sẽ thu thập một profile độ phủ để tìm hiểu.


## Sử dụng kiểm thử tích hợp để thu thập dữ liệu độ phủ

Hãy viết một script wrapper khác gọi script trước, nhưng xây dựng công cụ với đo lường độ phủ và sau đó xử lý hậu kỳ các profile kết quả:

```
$ cat wrap_test_for_coverage.sh
#!/bin/sh
set -e
PKGARGS="$*"
#
# Setup
#
rm -rf covdatafiles
mkdir covdatafiles
#
# Pass in "-cover" to the script to build for coverage, then
# run with GOCOVERDIR set.
#
GOCOVERDIR=covdatafiles \
  /bin/sh integration_test.sh -cover $PKGARGS
#
# Post-process the resulting profiles.
#
go tool covdata percent -i=covdatafiles
$
```

Một số điểm quan trọng cần lưu ý về wrapper trên:

  * nó truyền vào cờ "-cover" khi chạy `integration_test.sh`, điều này cho chúng ta một tệp nhị phân "mdtool.exe" được đo lường độ phủ
  * nó đặt biến môi trường GOCOVERDIR thành một thư mục nơi các tệp dữ liệu độ phủ sẽ được ghi vào
  * khi kiểm thử hoàn tất, nó chạy "go tool covdata percent" để tạo ra báo cáo về tỷ lệ phần trăm câu lệnh được bao phủ

Đây là đầu ra khi chúng ta chạy script wrapper mới này:

```
$ /bin/sh wrap_test_for_coverage.sh
...
	gitlab.com/golang-commonmark/mdtool	coverage: 48.1% of statements
$
# Note: covdatafiles now contains 381 files.
```

Voila!
Bây giờ chúng tôi có một số ý tưởng về mức độ tốt của các kiểm thử tích hợp trong việc thực thi mã nguồn của ứng dụng "mdtool".

Nếu chúng tôi thực hiện các thay đổi để cải thiện test harness, rồi chạy thu thập độ phủ lần thứ hai, chúng tôi sẽ thấy các thay đổi được phản ánh trong báo cáo độ phủ.
Ví dụ, giả sử chúng tôi cải thiện kiểm thử bằng cách thêm hai dòng bổ sung này vào `integration_test.sh`:

{{raw `
	./mdtool.exe +ty testdata/README.md  > /dev/null
	./mdtool.exe +ta < testdata/README.md  > /dev/null
`}}

Chạy wrapper kiểm thử độ phủ lại:

```
$ /bin/sh wrap_test_for_coverage.sh
finished processing 380 files, no crashes
	gitlab.com/golang-commonmark/mdtool	coverage: 54.6% of statements
$
```

Chúng ta có thể thấy hiệu ứng của thay đổi: độ phủ câu lệnh đã tăng từ 48% lên 54%.

## Chọn các gói cần đo độ phủ

Theo mặc định, "go build -cover" sẽ chỉ đo lường các gói thuộc Go module đang được xây dựng, trong trường hợp này là gói `gitlab.com/golang-commonmark/mdtool`.
Tuy nhiên, trong một số trường hợp sẽ hữu ích khi mở rộng đo lường độ phủ sang các gói khác; điều này có thể thực hiện bằng cách truyền "-coverpkg" vào "go build -cover".

Đối với chương trình mẫu của chúng tôi, "mdtool" thực ra chủ yếu chỉ là một wrapper xung quanh gói `gitlab.com/golang-commonmark/markdown`, vì vậy sẽ thú vị khi đưa `markdown` vào tập các gói được đo lường.

Đây là tệp `go.mod` của "mdtool":

```
$ head go.mod
module gitlab.com/golang-commonmark/mdtool

go 1.17

require (
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	gitlab.com/golang-commonmark/markdown v0.0.0-20211110145824-bf3e522c626a
)
```

Chúng ta có thể dùng cờ "-coverpkg" để kiểm soát các gói được chọn để đưa vào phân tích độ phủ bao gồm một trong các dependency trên.
Đây là ví dụ:

```
$ /bin/sh wrap_test_for_coverage.sh -coverpkg=gitlab.com/golang-commonmark/markdown,gitlab.com/golang-commonmark/mdtool
...
	gitlab.com/golang-commonmark/markdown	coverage: 70.6% of statements
	gitlab.com/golang-commonmark/mdtool	coverage: 54.6% of statements
$
```

# Làm việc với các tệp dữ liệu độ phủ

Khi một kiểm thử tích hợp độ phủ hoàn tất và ghi ra một tập các tệp dữ liệu thô (trong trường hợp của chúng tôi, nội dung của thư mục `covdatafiles`), chúng ta có thể xử lý hậu kỳ các tệp này theo nhiều cách.

## Chuyển đổi profile sang định dạng văn bản '-coverprofile'

Khi làm việc với kiểm thử đơn vị, bạn có thể chạy `go test -coverprofile=abc.txt` để ghi một profile độ phủ định dạng văn bản cho một lần chạy kiểm thử độ phủ nhất định.

Với các tệp nhị phân được xây dựng bằng `go build -cover`, bạn có thể tạo một profile định dạng văn bản sau đó bằng cách chạy `go tool covdata textfmt` trên các tệp được phát ra vào thư mục GOCOVERDIR.

Khi bước này hoàn tất, bạn có thể dùng `go tool cover -func=<file>` hoặc `go tool cover -html=<file>` để diễn giải/trực quan hóa dữ liệu, giống như bạn làm với `go test -coverprofile`.

Ví dụ:

```
$ /bin/sh wrap_test_for_coverage.sh
...
$ go tool covdata textfmt -i=covdatafiles -o=cov.txt
$ go tool cover -func=cov.txt
gitlab.com/golang-commonmark/mdtool/main.go:40:		readFromStdin	100.0%
gitlab.com/golang-commonmark/mdtool/main.go:44:		readFromFile	80.0%
gitlab.com/golang-commonmark/mdtool/main.go:54:		readFromWeb	0.0%
gitlab.com/golang-commonmark/mdtool/main.go:64:		readInput	80.0%
gitlab.com/golang-commonmark/mdtool/main.go:74:		extractText	100.0%
gitlab.com/golang-commonmark/mdtool/main.go:88:		writePreamble	100.0%
gitlab.com/golang-commonmark/mdtool/main.go:111:	writePostamble	100.0%
gitlab.com/golang-commonmark/mdtool/main.go:118:	handler		0.0%
gitlab.com/golang-commonmark/mdtool/main.go:139:	main		51.6%
total:							(statements)	54.6%
$
```

## Hợp nhất các profile thô với 'go tool covdata merge'

Mỗi lần thực thi một ứng dụng được xây dựng với "-cover" sẽ ghi ra một hoặc nhiều tệp dữ liệu vào thư mục được chỉ định trong biến môi trường GOCOVERDIR.
Nếu một kiểm thử tích hợp thực hiện N lần chạy chương trình, bạn sẽ có O(N) tệp trong thư mục đầu ra.
Thường có nhiều nội dung trùng lặp trong các tệp dữ liệu, vì vậy để thu gọn dữ liệu và/hoặc kết hợp các tập dữ liệu từ các lần chạy kiểm thử tích hợp khác nhau, bạn có thể dùng lệnh `go tool covdata merge` để hợp nhất các profile lại.
Ví dụ:

```
$ /bin/sh wrap_test_for_coverage.sh
finished processing 380 files, no crashes
	gitlab.com/golang-commonmark/mdtool	coverage: 54.6% of statements
$ ls covdatafiles
covcounters.13326b42c2a107249da22f6e0d35b638.772307.1677775306041466651
covcounters.13326b42c2a107249da22f6e0d35b638.772314.1677775306053066987
...
covcounters.13326b42c2a107249da22f6e0d35b638.774973.1677775310032569308
covmeta.13326b42c2a107249da22f6e0d35b638
$ ls covdatafiles | wc
    381     381   27401
$ rm -rf merged ; mkdir merged ; go tool covdata merge -i=covdatafiles -o=merged
$ ls merged
covcounters.13326b42c2a107249da22f6e0d35b638.0.1677775331350024014
covmeta.13326b42c2a107249da22f6e0d35b638
$
```

Thao tác `go tool covdata merge` cũng chấp nhận cờ `-pkg` có thể dùng để chọn một gói hoặc tập hợp gói cụ thể, nếu cần.

Khả năng hợp nhất này cũng hữu ích để kết hợp kết quả từ các loại lần chạy kiểm thử khác nhau, bao gồm cả các lần chạy được tạo bởi các test harness khác.


# Tổng kết

Vậy là đã đủ: với bản phát hành 1.20, hệ thống công cụ độ phủ của Go không còn bị giới hạn trong các kiểm thử gói, mà hỗ trợ thu thập profile từ các kiểm thử tích hợp lớn hơn.
Chúng tôi hy vọng bạn sẽ tận dụng tốt các tính năng mới này để hiểu rõ hơn các kiểm thử lớn và phức tạp hơn của bạn đang hoạt động như thế nào, và chúng đang thực thi những phần nào trong mã nguồn của bạn.

Hãy thử các tính năng mới này, và như thường lệ nếu bạn gặp vấn đề, hãy đăng issue trên [GitHub issue tracker](https://github.com/golang/go/issues) của chúng tôi.
Cảm ơn.
