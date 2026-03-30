---
title: Hỗ trợ đo độ phủ code cho kiểm thử tích hợp
layout: article
template: true
---

Mục lục:

 [Tổng quan](#overview)\
 [Xây dựng binary để đo độ phủ](#building)\
 [Chạy binary đã được thêm công cụ đo độ phủ](#running)\
 [Làm việc với tệp dữ liệu độ phủ](#working)\
 [Câu hỏi thường gặp](#FAQ)\
 [Tài nguyên](#resources)\
 [Thuật ngữ](#glossary)


Bắt đầu từ Go 1.20, Go hỗ trợ thu thập hồ sơ độ phủ từ ứng dụng và từ kiểm thử tích hợp, loại kiểm thử lớn hơn và phức tạp hơn cho các chương trình Go.

# Tổng quan {#overview}

Go cung cấp hỗ trợ dễ dùng để thu thập hồ sơ độ phủ ở cấp kiểm thử đơn vị package thông qua lệnh "`go test -coverprofile=... <pkg_target>`".
Bắt đầu từ Go 1.20, người dùng có thể thu thập hồ sơ độ phủ cho [kiểm thử tích hợp](#glos-integration-test) quy mô lớn hơn: những kiểm thử nặng và phức tạp hơn, thực hiện nhiều lần chạy một binary ứng dụng nhất định.

Đối với kiểm thử đơn vị, thu thập hồ sơ độ phủ và tạo báo cáo cần hai bước: một lần chạy `go test -coverprofile=...`, sau đó gọi `go tool cover {-func,-html}` để tạo báo cáo.

Đối với kiểm thử tích hợp, cần ba bước: bước [xây dựng](#building), bước [chạy](#running) (có thể gồm nhiều lần gọi binary từ bước xây dựng), và cuối cùng là bước [báo cáo](#reporting), như mô tả dưới đây.

# Xây dựng binary để đo độ phủ {#building}

Để xây dựng ứng dụng nhằm thu thập hồ sơ độ phủ, truyền cờ `-cover` khi gọi `go build` trên binary ứng dụng mục tiêu. Xem phần [dưới đây](#packageselection) để biết ví dụ lệnh `go build -cover`.
Binary kết quả sau đó có thể được chạy bằng cách thiết lập biến môi trường để thu thập hồ sơ độ phủ (xem phần tiếp theo về [chạy](#running)).

## Cách chọn gói để thêm công cụ đo {#packageselection}

Trong một lần gọi "`go build -cover`" nhất định, lệnh Go sẽ chọn các gói trong module chính để đo độ phủ; các gói khác tham gia vào quá trình xây dựng (dependency được liệt kê trong go.mod, hoặc các gói là một phần của thư viện chuẩn Go) sẽ không được bao gồm theo mặc định.

Ví dụ, đây là một chương trình đơn giản chứa gói main, một gói module-chính cục bộ `greetings` và một tập hợp các gói được import từ bên ngoài module, bao gồm (trong số các gói khác) `rsc.io/quote` và `fmt` ([liên kết đến chương trình đầy đủ](/play/p/VSQJN8xkkf-?v=gotip)).

```
$ cat go.mod
module mydomain.com

go 1.20

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

$ cat myprogram.go
package main

import (
	"fmt"
	"mydomain.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Printf("I say %q and %q\n", quote.Hello(), greetings.Goodbye())
}
$ cat greetings/greetings.go
package greetings

func Goodbye() string {
	return "see ya"
}
$ go build -cover -o myprogram.exe .
$
```

Nếu bạn xây dựng chương trình này với cờ "`-cover`" và chạy nó, chính xác hai gói sẽ được đưa vào hồ sơ: `main` và `mydomain.com/greetings`; các gói phụ thuộc khác sẽ bị loại trừ.

Người dùng muốn kiểm soát chặt hơn các gói nào được đưa vào đo độ phủ có thể xây dựng với cờ "`-coverpkg`". Ví dụ:

```
$ go build -cover -o myprogramMorePkgs.exe -coverpkg=io,mydomain.com,rsc.io/quote .
$
```

Trong lệnh xây dựng trên, gói main từ `mydomain.com` cũng như các gói `rsc.io/quote` và `io` được chọn để đo; vì `mydomain.com/greetings` không được liệt kê cụ thể, nó sẽ bị loại khỏi hồ sơ, dù nó nằm trong module chính.

# Chạy binary đã được thêm công cụ đo độ phủ {#running}

Các binary được xây dựng với "`-cover`" sẽ ghi ra tệp dữ liệu hồ sơ vào cuối quá trình thực thi vào một thư mục được chỉ định qua biến môi trường `GOCOVERDIR`. Ví dụ:

```
$ go build -cover -o myprogram.exe myprogram.go
$ mkdir somedata
$ GOCOVERDIR=somedata ./myprogram.exe
I say "Hello, world." and "see ya"
$ ls somedata
covcounters.c6de772f99010ef5925877a7b05db4cc.2424989.1670252383678349347
covmeta.c6de772f99010ef5925877a7b05db4cc
$
```

Lưu ý hai tệp được ghi vào thư mục `somedata`: các tệp (nhị phân) này chứa kết quả độ phủ. Xem phần tiếp theo về [báo cáo](#reporting) để biết thêm cách tạo kết quả dễ đọc từ các tệp dữ liệu này.

Nếu biến môi trường `GOCOVERDIR` không được thiết lập, binary đã được thêm công cụ đo độ phủ vẫn sẽ thực thi đúng, nhưng sẽ đưa ra cảnh báo.
Ví dụ:

```
$ ./myprogram.exe
warning: GOCOVERDIR not set, no coverage data emitted
I say "Hello, world." and "see ya"
$
```

## Kiểm thử gồm nhiều lần chạy

Kiểm thử tích hợp trong nhiều trường hợp có thể gồm nhiều lần chạy chương trình; khi chương trình được xây dựng với "`-cover`", mỗi lần chạy sẽ tạo ra một tệp dữ liệu mới. Ví dụ:

```
$ mkdir somedata2
$ GOCOVERDIR=somedata2 ./myprogram.exe          // lần chạy thứ nhất
I say "Hello, world." and "see ya"
$ GOCOVERDIR=somedata2 ./myprogram.exe -flag    // lần chạy thứ hai
I say "Hello, world." and "see ya"
$ ls somedata2
covcounters.890814fca98ac3a4d41b9bd2a7ec9f7f.2456041.1670259309405583534
covcounters.890814fca98ac3a4d41b9bd2a7ec9f7f.2456047.1670259309410891043
covmeta.890814fca98ac3a4d41b9bd2a7ec9f7f
$
```

Các tệp đầu ra dữ liệu độ phủ có hai dạng: tệp siêu dữ liệu (chứa các mục không thay đổi giữa các lần chạy, như tên tệp nguồn và tên hàm), và tệp dữ liệu bộ đếm (ghi lại các phần của chương trình đã được thực thi).

Trong ví dụ trên, lần chạy đầu tiên tạo ra hai tệp (bộ đếm và siêu dữ liệu), trong khi lần chạy thứ hai chỉ tạo ra một tệp dữ liệu bộ đếm: vì siêu dữ liệu không thay đổi giữa các lần chạy, nó chỉ cần được ghi một lần.

# Làm việc với tệp dữ liệu độ phủ {#working}

Go 1.20 giới thiệu một công cụ mới, '`covdata`', có thể được dùng để đọc và xử lý tệp dữ liệu độ phủ từ thư mục `GOCOVERDIR`.

Công cụ `covdata` của Go chạy ở nhiều chế độ khác nhau. Dạng chung của lệnh gọi công cụ `covdata` là:

```
$ go tool covdata <mode> -i=<dir1,dir2,...> ...flags...
```

trong đó cờ "`-i`" cung cấp danh sách các thư mục cần đọc, mỗi thư mục được tạo ra từ một lần thực thi binary đã được thêm công cụ đo độ phủ (thông qua `GOCOVERDIR`).

## Tạo báo cáo hồ sơ độ phủ {#reporting}

Phần này thảo luận cách dùng "`go tool covdata`" để tạo báo cáo dễ đọc từ tệp dữ liệu độ phủ.

### Báo cáo phần trăm câu lệnh được bao phủ

Để báo cáo chỉ số "phần trăm câu lệnh được bao phủ" cho mỗi gói được thêm công cụ đo, dùng lệnh "`go tool covdata percent -i=<directory>`".
Sử dụng ví dụ từ phần [chạy](#running) ở trên:

```
$ ls somedata
covcounters.c6de772f99010ef5925877a7b05db4cc.2424989.1670252383678349347
covmeta.c6de772f99010ef5925877a7b05db4cc
$ go tool covdata percent -i=somedata
	main	coverage: 100.0% of statements
	mydomain.com/greetings	coverage: 100.0% of statements
$
```

Các phần trăm "câu lệnh được bao phủ" ở đây tương ứng trực tiếp với những gì được báo cáo bởi `go test -cover`.

## Chuyển đổi sang định dạng văn bản cũ

Bạn có thể chuyển đổi các tệp dữ liệu độ phủ nhị phân sang định dạng văn bản cũ được tạo bởi "`go test -coverprofile=<outfile>`" bằng cách dùng trình chọn `textfmt` của covdata. Tệp văn bản kết quả sau đó có thể được dùng với "`go tool cover -func`" hoặc "`go tool cover -html`" để tạo thêm báo cáo. Ví dụ:

```
$ ls somedata
covcounters.c6de772f99010ef5925877a7b05db4cc.2424989.1670252383678349347
covmeta.c6de772f99010ef5925877a7b05db4cc
$ go tool covdata textfmt -i=somedata -o profile.txt
$ cat profile.txt
mode: set
mydomain.com/myprogram.go:10.13,12.2 1 1
mydomain.com/greetings/greetings.go:3.23,5.2 1 1
$ go tool cover -func=profile.txt
mydomain.com/greetings/greetings.go:3:	Goodbye		100.0%
mydomain.com/myprogram.go:10:		main		100.0%
total:					(statements)	100.0%
$
```

## Hợp nhất

Lệnh con `merge` của "`go tool covdata`" có thể được dùng để hợp nhất các hồ sơ từ nhiều thư mục dữ liệu.

Ví dụ, hãy xem xét một chương trình chạy trên cả macOS và Windows.
Tác giả của chương trình này có thể muốn kết hợp các hồ sơ độ phủ từ các lần chạy riêng biệt trên mỗi hệ điều hành thành một kho hồ sơ duy nhất, để tạo ra một bản tóm tắt độ phủ đa nền tảng.
Ví dụ:

```
$ ls windows_datadir
covcounters.f3833f80c91d8229544b25a855285890.1025623.1667481441036838252
covcounters.f3833f80c91d8229544b25a855285890.1025628.1667481441042785007
covmeta.f3833f80c91d8229544b25a855285890
$ ls macos_datadir
covcounters.b245ad845b5068d116a4e25033b429fb.1025358.1667481440551734165
covcounters.b245ad845b5068d116a4e25033b429fb.1025364.1667481440557770197
covmeta.b245ad845b5068d116a4e25033b429fb
$ ls macos_datadir
$ mkdir merged
$ go tool covdata merge -i=windows_datadir,macos_datadir -o merged
$
```

Thao tác hợp nhất trên sẽ kết hợp dữ liệu từ các thư mục đầu vào được chỉ định và ghi một tập hợp tệp dữ liệu đã hợp nhất mới vào thư mục "merged".

## Chọn gói

Hầu hết các lệnh "`go tool covdata`" đều hỗ trợ cờ "`-pkg`" để thực hiện chọn gói như một phần của thao tác; đối số cho "`-pkg`" có cùng dạng như đối số được dùng bởi cờ "`-coverpkg`" của lệnh Go.
Ví dụ:

```

$ ls somedata
covcounters.c6de772f99010ef5925877a7b05db4cc.2424989.1670252383678349347
covmeta.c6de772f99010ef5925877a7b05db4cc
$ go tool covdata percent -i=somedata -pkg=mydomain.com/greetings
	mydomain.com/greetings	coverage: 100.0% of statements
$ go tool covdata percent -i=somedata -pkg=nonexistentpackage
$
```

Cờ "`-pkg`" có thể được dùng để chọn tập hợp con cụ thể các gói quan tâm cho một báo cáo nhất định.

#

## Câu hỏi thường gặp {#FAQ}

1. [Làm thế nào để yêu cầu thêm công cụ đo độ phủ cho tất cả các gói được import trong tệp `go.mod`](#gomodselect)
2. [Tôi có thể dùng `go build -cover` ở chế độ GOPATH/GO111MODULE=off không?](#gopathmode)
3. [Nếu chương trình của tôi bị panic, dữ liệu độ phủ có được ghi không?](#panicprof)
4. [`-coverpkg=main` có chọn gói main của tôi để đo không?](#mainpkg)


#### Làm thế nào để yêu cầu thêm công cụ đo độ phủ cho tất cả các gói được import trong tệp `go.mod` {#gomodselect}

Theo mặc định, `go build -cover` sẽ thêm công cụ đo cho tất cả các gói module chính để đo độ phủ, nhưng sẽ không thêm công cụ đo cho các import bên ngoài module chính (ví dụ: gói thư viện chuẩn hoặc các import được liệt kê trong `go.mod`).
Một cách để yêu cầu thêm công cụ đo cho tất cả các dependency không phải thư viện chuẩn là đưa đầu ra của `go list` vào `-coverpkg`.
Đây là một ví dụ, một lần nữa sử dụng [chương trình ví dụ](/play/p/VSQJN8xkkf-?v=gotip) được đề cập ở trên:

```
$ go list -f '{{"{{if not .Standard}}{{.ImportPath}}{{end}}"}}' -deps . | paste -sd "," > pkgs.txt
$ go build -o myprogram.exe -coverpkg=`cat pkgs.txt` .
$ mkdir somedata
$ GOCOVERDIR=somedata ./myprogram.exe
$ go tool covdata percent -i=somedata
	golang.org/x/text/internal/tag	coverage: 78.4% of statements
	golang.org/x/text/language	coverage: 35.5% of statements
	mydomain.com	coverage: 100.0% of statements
	mydomain.com/greetings	coverage: 100.0% of statements
	rsc.io/quote	coverage: 25.0% of statements
	rsc.io/sampler	coverage: 86.7% of statements
$
```

#### Tôi có thể dùng `go build -cover` ở chế độ GO111MODULE=off không? {#gopathmode}

Có, `go build -cover` hoạt động với `GO111MODULE=off`.
Khi xây dựng chương trình ở chế độ GO111MODULE=off, chỉ có gói được đặt tên cụ thể làm mục tiêu trên dòng lệnh mới được thêm công cụ đo để đo. Dùng cờ `-coverpkg` để bao gồm thêm các gói vào hồ sơ.

#### Nếu chương trình của tôi bị panic, dữ liệu độ phủ có được ghi không? {#panicprof}

Các chương trình được xây dựng với `go build -cover` chỉ ghi ra dữ liệu hồ sơ đầy đủ vào cuối quá trình thực thi nếu chương trình gọi `os.Exit()` hoặc trả về bình thường từ `main.main`.
Nếu chương trình kết thúc do panic không được khôi phục, hoặc nếu chương trình gặp ngoại lệ nghiêm trọng (như vi phạm phân đoạn bộ nhớ, chia cho không, v.v.), dữ liệu hồ sơ từ các câu lệnh đã thực thi trong lần chạy đó sẽ bị mất.

#### `-coverpkg=main` có chọn gói main của tôi để đo không? {#mainpkg}

Cờ `-coverpkg` nhận danh sách các đường dẫn import, không phải danh sách tên gói. Nếu bạn muốn chọn gói `main` của mình để thêm công cụ đo độ phủ, hãy xác định nó bằng đường dẫn import, không phải bằng tên. Ví dụ (sử dụng [chương trình ví dụ này](/play/p/VSQJN8xkkf-?v=gotip)):

```
$ go list -m
mydomain.com
$ go build -coverpkg=main -o oops.exe .
warning: no packages being built depend on matches for pattern main
$ go build -coverpkg=mydomain.com -o myprogram.exe .
$ mkdir somedata
$ GOCOVERDIR=somedata ./myprogram.exe
I say "Hello, world." and "see ya"
$ go tool covdata percent -i=somedata
	mydomain.com	coverage: 100.0% of statements
$
```

## Tài nguyên {#resources}

- **Bài viết blog giới thiệu đo độ phủ kiểm thử đơn vị trong Go 1.2**:
  - Đo độ phủ cho kiểm thử đơn vị được giới thiệu trong bản phát hành Go 1.2; xem [bài viết blog này](/blog/cover) để biết chi tiết.
- **Tài liệu**:
  - Tài liệu gói [`cmd/go`](https://pkg.go.dev/cmd/go) mô tả các cờ xây dựng và kiểm thử liên quan đến độ phủ.
- **Chi tiết kỹ thuật**:
  - [Bản thiết kế nháp](/design/51430-revamp-code-coverage)
  - [Đề xuất](/issue/51430)

## Thuật ngữ {#glossary}

<a id="glos-unit-test"></a>
**kiểm thử đơn vị:** Các kiểm thử trong tệp `*_test.go` gắn với một gói Go cụ thể, sử dụng gói `testing` của Go.

<a id="glos-integration-test"></a>
**kiểm thử tích hợp:** Kiểm thử toàn diện hơn, nặng hơn cho một ứng dụng hoặc binary nhất định. Kiểm thử tích hợp thường bao gồm việc xây dựng một hoặc nhiều chương trình, sau đó thực hiện một loạt lần chạy chương trình với nhiều đầu vào và kịch bản khác nhau, dưới sự kiểm soát của một khung kiểm thử có thể dựa trên hoặc không dựa trên gói `testing` của Go.

