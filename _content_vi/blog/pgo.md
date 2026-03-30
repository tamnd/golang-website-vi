---
title: Tối ưu hóa dựa trên profile trong Go 1.21
date: 2023-09-05
by:
- Michael Pratt
summary: Giới thiệu về tối ưu hóa dựa trên profile, ra mắt chính thức trong Go 1.21.
template: true
---

Đầu năm 2023, Go 1.20 [đã phát hành bản xem trước của tối ưu hóa dựa trên profile (PGO)](/blog/pgo-preview) để người dùng thử nghiệm.
Sau khi giải quyết các hạn chế đã biết trong bản xem trước, và với các cải tiến bổ sung nhờ phản hồi và đóng góp từ cộng đồng, hỗ trợ PGO trong Go 1.21 đã sẵn sàng để sử dụng trong production!
Xem [hướng dẫn người dùng tối ưu hóa dựa trên profile](/doc/pgo) để có tài liệu đầy đủ.

[Dưới đây](#example) chúng ta sẽ đi qua một ví dụ về việc sử dụng PGO để cải thiện hiệu năng của một ứng dụng.
Trước khi đến đó, "tối ưu hóa dựa trên profile" chính xác là gì?

Khi bạn build một binary Go, trình biên dịch Go thực hiện các tối ưu hóa để cố gắng tạo ra binary hoạt động tốt nhất có thể.
Ví dụ, lan truyền hằng số có thể đánh giá các biểu thức hằng số tại thời điểm biên dịch, tránh chi phí đánh giá tại runtime.
Phân tích escape tránh cấp phát heap cho các đối tượng có phạm vi cục bộ, tránh overhead của GC.
Inlining sao chép thân hàm của các hàm đơn giản vào nơi gọi, thường cho phép tối ưu hóa thêm tại nơi gọi (như lan truyền hằng số bổ sung hoặc phân tích escape tốt hơn).
Devirtualization chuyển đổi các lời gọi gián tiếp trên các giá trị interface có thể xác định kiểu tĩnh thành lời gọi trực tiếp đến phương thức cụ thể (thường cho phép inlining của lời gọi đó).

Go cải thiện các tối ưu hóa từ bản phát hành này sang bản phát hành khác, nhưng làm vậy không phải là nhiệm vụ dễ dàng.
Một số tối ưu hóa có thể điều chỉnh, nhưng trình biên dịch không thể chỉ "tăng hết mức" trên mọi tối ưu hóa vì các tối ưu hóa quá tích cực thực sự có thể gây hại cho hiệu năng hoặc gây ra thời gian build quá lâu.
Các tối ưu hóa khác yêu cầu trình biên dịch đưa ra phán đoán về đường dẫn "phổ biến" và "không phổ biến" trong một hàm là gì.
Trình biên dịch phải đưa ra phỏng đoán tốt nhất dựa trên các heuristic tĩnh vì nó không thể biết các trường hợp nào sẽ phổ biến tại thời điểm chạy.

Hay có thể biết được không?

Không có thông tin xác định về cách code được sử dụng trong môi trường production, trình biên dịch chỉ có thể hoạt động trên source code của các package.
Nhưng chúng ta có một công cụ để đánh giá hành vi production: [profiling](/doc/diagnostics#profiling).
Nếu chúng ta cung cấp một profile cho trình biên dịch, nó có thể đưa ra các quyết định sáng suốt hơn: tối ưu hóa tích cực hơn các hàm được sử dụng thường xuyên nhất, hoặc chọn các trường hợp phổ biến chính xác hơn.

Sử dụng các profile về hành vi ứng dụng để tối ưu hóa trình biên dịch được gọi là _Tối ưu hóa dựa trên Profile (PGO)_ (còn được gọi là Tối ưu hóa dựa trên Phản hồi (FDO)).

## Ví dụ

Hãy xây dựng một dịch vụ chuyển đổi Markdown sang HTML: người dùng tải Markdown nguồn lên `/render`, trả về kết quả chuyển đổi HTML.
Chúng ta có thể sử dụng [`gitlab.com/golang-commonmark/markdown`](https://pkg.go.dev/gitlab.com/golang-commonmark/markdown) để triển khai điều này một cách dễ dàng.

### Thiết lập

```
$ go mod init example.com/markdown
$ go get gitlab.com/golang-commonmark/markdown@bf3e522c626a
```

Trong `main.go`:

```
package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"

	"gitlab.com/golang-commonmark/markdown"
)

func render(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	src, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	md := markdown.New(
		markdown.XHTMLOutput(true),
		markdown.Typographer(true),
		markdown.Linkify(true),
		markdown.Tables(true),
	)

	var buf bytes.Buffer
	if err := md.Render(&buf, src); err != nil {
		log.Printf("error converting markdown: %v", err)
		http.Error(w, "Malformed markdown", http.StatusBadRequest)
		return
	}

	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/render", render)
	log.Printf("Serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Build và chạy server:

```
$ go build -o markdown.nopgo.exe
$ ./markdown.nopgo.exe
2023/08/23 03:55:51 Serving on port 8080...
```

Hãy thử gửi một số Markdown từ terminal khác.
Chúng ta có thể sử dụng `README.md` từ dự án Go làm tài liệu mẫu:

```
$ curl -o README.md -L "https://raw.githubusercontent.com/golang/go/c16c2c49e2fa98ae551fc6335215fadd62d33542/README.md"
$ curl --data-binary @README.md http://localhost:8080/render
<h1>The Go Programming Language</h1>
<p>Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.</p>
...
```

### Profiling

Bây giờ chúng ta có một dịch vụ hoạt động, hãy thu thập một profile và rebuild với PGO để xem liệu chúng ta có hiệu năng tốt hơn không.

Trong `main.go`, chúng ta đã import [net/http/pprof](https://pkg.go.dev/net/http/pprof) tự động thêm endpoint `/debug/pprof/profile` vào server để lấy CPU profile.

Thông thường bạn muốn thu thập profile từ môi trường production của mình để trình biên dịch nhận được cái nhìn đại diện về hành vi trong production.
Vì ví dụ này không có môi trường "production", tôi đã tạo một [chương trình đơn giản](https://github.com/prattmic/markdown-pgo/blob/main/load/main.go) để tạo tải trong khi chúng ta thu thập profile.
Tải và khởi động bộ tạo tải (đảm bảo server vẫn đang chạy!):

```
$ go run github.com/prattmic/markdown-pgo/load@latest
```

Trong khi đó đang chạy, tải xuống một profile từ server:

```
$ curl -o cpu.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
```

Khi hoàn thành, dừng bộ tạo tải và server.

### Sử dụng profile

Go toolchain sẽ tự động bật PGO khi tìm thấy một profile có tên `default.pgo` trong thư mục package chính.
Ngoài ra, cờ `-pgo` cho `go build` nhận đường dẫn đến profile để sử dụng cho PGO.

Chúng tôi khuyến nghị commit các tệp `default.pgo` vào repository của bạn.
Lưu trữ các profile cùng với source code đảm bảo người dùng tự động có quyền truy cập vào profile chỉ bằng cách lấy repository (qua hệ thống kiểm soát phiên bản, hoặc qua `go get`) và các build vẫn có thể tái tạo.

Hãy build:

```
$ mv cpu.pprof default.pgo
$ go build -o markdown.withpgo.exe
```

Chúng ta có thể kiểm tra xem PGO có được bật trong build không với `go version`:

```
$ go version -m markdown.withpgo.exe
./markdown.withpgo.exe: go1.21.0
...
        build   -pgo=/tmp/pgo121/default.pgo

```

### Đánh giá

Chúng ta sẽ sử dụng [phiên bản benchmark](https://github.com/prattmic/markdown-pgo/blob/main/load/bench_test.go) của bộ tạo tải để đánh giá tác động của PGO lên hiệu năng.

Đầu tiên, chúng ta sẽ benchmark server không có PGO.
Khởi động server đó:

```
$ ./markdown.nopgo.exe
```

Trong khi đó đang chạy, chạy một số lần lặp benchmark:

```
$ go get github.com/prattmic/markdown-pgo@latest
$ go test github.com/prattmic/markdown-pgo/load -bench=. -count=40 -source $(pwd)/README.md > nopgo.txt
```

Khi hoàn thành, dừng server gốc và khởi động phiên bản có PGO:

```
$ ./markdown.withpgo.exe
```

Trong khi đó đang chạy, chạy một số lần lặp benchmark:

```
$ go test github.com/prattmic/markdown-pgo/load -bench=. -count=40 -source $(pwd)/README.md > withpgo.txt
```

Khi hoàn thành, hãy so sánh kết quả:

```
$ go install golang.org/x/perf/cmd/benchstat@latest
$ benchstat nopgo.txt withpgo.txt
goos: linux
goarch: amd64
pkg: github.com/prattmic/markdown-pgo/load
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
        │  nopgo.txt  │            withpgo.txt             │
        │   sec/op    │   sec/op     vs base               │
Load-12   374.5µ ± 1%   360.2µ ± 0%  -3.83% (p=0.000 n=40)
```

Phiên bản mới nhanh hơn khoảng 3,8%!
Trong Go 1.21, các workload thường đạt được cải thiện mức sử dụng CPU từ 2% đến 7% khi bật PGO.
Các profile chứa nhiều thông tin về hành vi ứng dụng và Go 1.21 chỉ mới bắt đầu khai thác bề mặt bằng cách sử dụng thông tin này cho một tập hạn chế các tối ưu hóa.
Các bản phát hành tương lai sẽ tiếp tục cải thiện hiệu năng khi nhiều phần của trình biên dịch tận dụng PGO.

## Các bước tiếp theo

Trong ví dụ này, sau khi thu thập profile, chúng ta đã rebuild server sử dụng chính xác source code giống nhau được sử dụng trong build gốc.
Trong kịch bản thực tế, luôn có sự phát triển liên tục.
Vì vậy, chúng ta có thể thu thập profile từ production đang chạy code của tuần trước và sử dụng nó để build với source code ngày hôm nay.
Điều đó hoàn toàn ổn!
PGO trong Go có thể xử lý các thay đổi nhỏ đối với source code mà không gặp vấn đề gì.
Tất nhiên, theo thời gian source code sẽ ngày càng khác biệt hơn, vì vậy vẫn quan trọng là cập nhật profile đôi khi.

Để biết thêm nhiều thông tin về việc sử dụng PGO, các thực hành tốt nhất và những điều cần lưu ý, hãy xem [hướng dẫn người dùng tối ưu hóa dựa trên profile](/doc/pgo).
Nếu bạn tò mò về những gì đang xảy ra bên dưới, hãy tiếp tục đọc!

## Bên dưới mui xe

Để hiểu rõ hơn điều gì đã làm ứng dụng này nhanh hơn, hãy nhìn vào bên dưới để xem hiệu năng đã thay đổi như thế nào.
Chúng ta sẽ xem xét hai tối ưu hóa khác nhau được thúc đẩy bởi PGO.

### Inlining

Để quan sát các cải tiến inlining, hãy phân tích ứng dụng markdown này cả có và không có PGO.

Tôi sẽ so sánh điều này bằng một kỹ thuật gọi là differential profiling, nơi chúng ta thu thập hai profile (một có PGO và một không có) và so sánh chúng.
Để differential profiling, điều quan trọng là cả hai profile đại diện cho cùng một lượng **công việc**, không phải cùng một lượng thời gian, vì vậy tôi đã điều chỉnh server để tự động thu thập profile, và bộ tạo tải để gửi số lượng yêu cầu cố định rồi thoát khỏi server.

Các thay đổi tôi đã thực hiện với server cũng như các profile được thu thập có thể tìm thấy tại https://github.com/prattmic/markdown-pgo.
Bộ tạo tải được chạy với `-count=300000 -quit`.

Để kiểm tra nhanh sự nhất quán, hãy xem tổng thời gian CPU cần thiết để xử lý tất cả 300k yêu cầu:

```
$ go tool pprof -top cpu.nopgo.pprof | grep "Total samples"
Duration: 116.92s, Total samples = 118.73s (101.55%)
$ go tool pprof -top cpu.withpgo.pprof | grep "Total samples"
Duration: 113.91s, Total samples = 115.03s (100.99%)
```

Thời gian CPU giảm từ ~118s xuống ~115s, khoảng 3%.
Điều này phù hợp với kết quả benchmark của chúng ta, đây là dấu hiệu tốt rằng các profile này là đại diện.

Bây giờ chúng ta có thể mở một differential profile để tìm kiếm tiết kiệm:

```
$ go tool pprof -diff_base cpu.nopgo.pprof cpu.withpgo.pprof
File: markdown.profile.withpgo.exe
Type: cpu
Time: Aug 28, 2023 at 10:26pm (EDT)
Duration: 230.82s, Total samples = 118.73s (51.44%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top -cum
Showing nodes accounting for -0.10s, 0.084% of 118.73s total
Dropped 268 nodes (cum {{raw "<"}}= 0.59s)
Showing top 10 nodes out of 668
      flat  flat%   sum%        cum   cum%
    -0.03s 0.025% 0.025%     -2.56s  2.16%  gitlab.com/golang-commonmark/markdown.ruleLinkify
     0.04s 0.034% 0.0084%     -2.19s  1.84%  net/http.(*conn).serve
     0.02s 0.017% 0.025%     -1.82s  1.53%  gitlab.com/golang-commonmark/markdown.(*Markdown).Render
     0.02s 0.017% 0.042%     -1.80s  1.52%  gitlab.com/golang-commonmark/markdown.(*Markdown).Parse
    -0.03s 0.025% 0.017%     -1.71s  1.44%  runtime.mallocgc
    -0.07s 0.059% 0.042%     -1.62s  1.36%  net/http.(*ServeMux).ServeHTTP
     0.04s 0.034% 0.0084%     -1.58s  1.33%  net/http.serverHandler.ServeHTTP
    -0.01s 0.0084% 0.017%     -1.57s  1.32%  main.render
     0.01s 0.0084% 0.0084%     -1.56s  1.31%  net/http.HandlerFunc.ServeHTTP
    -0.09s 0.076% 0.084%     -1.25s  1.05%  runtime.newobject
(pprof) top
Showing nodes accounting for -1.41s, 1.19% of 118.73s total
Dropped 268 nodes (cum {{raw "<"}}= 0.59s)
Showing top 10 nodes out of 668
      flat  flat%   sum%        cum   cum%
    -0.46s  0.39%  0.39%     -0.91s  0.77%  runtime.scanobject
    -0.40s  0.34%  0.72%     -0.40s  0.34%  runtime.nextFreeFast (inline)
     0.36s   0.3%  0.42%      0.36s   0.3%  gitlab.com/golang-commonmark/markdown.performReplacements
    -0.35s  0.29%  0.72%     -0.37s  0.31%  runtime.writeHeapBits.flush
     0.32s  0.27%  0.45%      0.67s  0.56%  gitlab.com/golang-commonmark/markdown.ruleReplacements
    -0.31s  0.26%  0.71%     -0.29s  0.24%  runtime.writeHeapBits.write
    -0.30s  0.25%  0.96%     -0.37s  0.31%  runtime.deductAssistCredit
     0.29s  0.24%  0.72%      0.10s 0.084%  gitlab.com/golang-commonmark/markdown.ruleText
    -0.29s  0.24%  0.96%     -0.29s  0.24%  runtime.(*mspan).base (inline)
    -0.27s  0.23%  1.19%     -0.42s  0.35%  bytes.(*Buffer).WriteRune
```

Khi chỉ định `pprof -diff_base`, các giá trị hiển thị trong pprof là _sự chênh lệch_ giữa hai profile.
Vì vậy, ví dụ, `runtime.scanobject` sử dụng ít hơn 0,46s CPU với PGO so với không có.
Mặt khác, `gitlab.com/golang-commonmark/markdown.performReplacements` sử dụng nhiều hơn 0,36s CPU.
Trong differential profile, chúng ta thường muốn xem các giá trị tuyệt đối (cột `flat` và `cum`), vì các phần trăm không có ý nghĩa.

`top -cum` hiển thị các sự chênh lệch lớn nhất theo thay đổi tích lũy.
Tức là, sự chênh lệch CPU của một hàm và tất cả các callee bắc cầu từ hàm đó.
Điều này thường hiển thị các frame ngoài cùng trong call graph của chương trình, chẳng hạn như `main` hoặc một goroutine entry point khác.
Ở đây chúng ta có thể thấy hầu hết tiết kiệm đến từ phần `ruleLinkify` của việc xử lý các yêu cầu HTTP.

`top` hiển thị các sự chênh lệch lớn nhất chỉ giới hạn ở các thay đổi trong chính hàm đó.
Điều này thường hiển thị các frame bên trong trong call graph của chương trình, nơi hầu hết công việc thực sự đang xảy ra.
Ở đây chúng ta có thể thấy các tiết kiệm riêng lẻ đến chủ yếu từ các hàm `runtime`.

Đó là những gì? Hãy xem call stack để xem chúng đến từ đâu:

```
(pprof) peek scanobject$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.86s 94.51% |   runtime.gcDrain
                                            -0.09s  9.89% |   runtime.gcDrainN
                                             0.04s  4.40% |   runtime.markrootSpans
    -0.46s  0.39%  0.39%     -0.91s  0.77%                | runtime.scanobject
                                            -0.19s 20.88% |   runtime.greyobject
                                            -0.13s 14.29% |   runtime.heapBits.nextFast (inline)
                                            -0.08s  8.79% |   runtime.heapBits.next
                                            -0.08s  8.79% |   runtime.spanOfUnchecked (inline)
                                             0.04s  4.40% |   runtime.heapBitsForAddr
                                            -0.01s  1.10% |   runtime.findObject
----------------------------------------------------------+-------------
(pprof) peek gcDrain$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                               -1s   100% |   runtime.gcBgMarkWorker.func2
     0.15s  0.13%  0.13%        -1s  0.84%                | runtime.gcDrain
                                            -0.86s 86.00% |   runtime.scanobject
                                            -0.18s 18.00% |   runtime.(*gcWork).balance
                                            -0.11s 11.00% |   runtime.(*gcWork).tryGet
                                             0.09s  9.00% |   runtime.pollWork
                                            -0.03s  3.00% |   runtime.(*gcWork).tryGetFast (inline)
                                            -0.03s  3.00% |   runtime.markroot
                                            -0.02s  2.00% |   runtime.wbBufFlush
                                             0.01s  1.00% |   runtime/internal/atomic.(*Bool).Load (inline)
                                            -0.01s  1.00% |   runtime.gcFlushBgCredit
                                            -0.01s  1.00% |   runtime/internal/atomic.(*Int64).Add (inline)
----------------------------------------------------------+-------------
```

Vì vậy `runtime.scanobject` cuối cùng đến từ `runtime.gcBgMarkWorker`.
[Hướng dẫn GC Go](/doc/gc-guide#Identifying_costs) cho chúng ta biết rằng `runtime.gcBgMarkWorker` là một phần của garbage collector, vì vậy tiết kiệm từ `runtime.scanobject` phải là tiết kiệm GC.
Còn `nextFreeFast` và các hàm `runtime` khác thì sao?

```
(pprof) peek nextFreeFast$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.40s   100% |   runtime.mallocgc (inline)
    -0.40s  0.34%  0.34%     -0.40s  0.34%                | runtime.nextFreeFast
----------------------------------------------------------+-------------
(pprof) peek writeHeapBits
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.37s   100% |   runtime.heapBitsSetType
                                                 0     0% |   runtime.(*mspan).initHeapBits
    -0.35s  0.29%  0.29%     -0.37s  0.31%                | runtime.writeHeapBits.flush
                                            -0.02s  5.41% |   runtime.arenaIndex (inline)
----------------------------------------------------------+-------------
                                            -0.29s   100% |   runtime.heapBitsSetType
    -0.31s  0.26%  0.56%     -0.29s  0.24%                | runtime.writeHeapBits.write
                                             0.02s  6.90% |   runtime.arenaIndex (inline)
----------------------------------------------------------+-------------
(pprof) peek heapBitsSetType$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.82s   100% |   runtime.mallocgc
    -0.12s   0.1%   0.1%     -0.82s  0.69%                | runtime.heapBitsSetType
                                            -0.37s 45.12% |   runtime.writeHeapBits.flush
                                            -0.29s 35.37% |   runtime.writeHeapBits.write
                                            -0.03s  3.66% |   runtime.readUintptr (inline)
                                            -0.01s  1.22% |   runtime.writeHeapBitsForAddr (inline)
----------------------------------------------------------+-------------
(pprof) peek deductAssistCredit$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.37s   100% |   runtime.mallocgc
    -0.30s  0.25%  0.25%     -0.37s  0.31%                | runtime.deductAssistCredit
                                            -0.07s 18.92% |   runtime.gcAssistAlloc
----------------------------------------------------------+-------------
```

`nextFreeFast` và một số trong số 10 cái đứng đầu cuối cùng đến từ `runtime.mallocgc`, mà Hướng dẫn GC cho chúng ta biết là memory allocator.

Chi phí giảm trong GC và allocator cho thấy rằng chúng ta đang cấp phát ít hơn tổng thể.
Hãy xem các heap profile để có thêm thông tin:

```
$ go tool pprof -sample_index=alloc_objects -diff_base heap.nopgo.pprof heap.withpgo.pprof
File: markdown.profile.withpgo.exe
Type: alloc_objects
Time: Aug 28, 2023 at 10:28pm (EDT)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for -12044903, 8.29% of 145309950 total
Dropped 60 nodes (cum {{raw "<"}}= 726549)
Showing top 10 nodes out of 58
      flat  flat%   sum%        cum   cum%
  -4974135  3.42%  3.42%   -4974135  3.42%  gitlab.com/golang-commonmark/mdurl.Parse
  -4249044  2.92%  6.35%   -4249044  2.92%  gitlab.com/golang-commonmark/mdurl.(*URL).String
   -901135  0.62%  6.97%    -977596  0.67%  gitlab.com/golang-commonmark/puny.mapLabels
   -653998  0.45%  7.42%    -482491  0.33%  gitlab.com/golang-commonmark/markdown.(*StateInline).PushPending
   -557073  0.38%  7.80%    -557073  0.38%  gitlab.com/golang-commonmark/linkify.Links
   -557073  0.38%  8.18%    -557073  0.38%  strings.genSplit
   -436919   0.3%  8.48%    -232152  0.16%  gitlab.com/golang-commonmark/markdown.(*StateBlock).Lines
   -408617  0.28%  8.77%    -408617  0.28%  net/textproto.readMIMEHeader
    401432  0.28%  8.49%     499610  0.34%  bytes.(*Buffer).grow
    291659   0.2%  8.29%     291659   0.2%  bytes.(*Buffer).String (inline)
```

Tùy chọn `-sample_index=alloc_objects` hiển thị cho chúng ta số lần cấp phát, bất kể kích thước.
Điều này hữu ích vì chúng ta đang điều tra việc giảm mức sử dụng CPU, thường tương quan nhiều hơn với số lượng cấp phát hơn là kích thước.
Có khá nhiều giảm ở đây, nhưng hãy tập trung vào giảm lớn nhất, `mdurl.Parse`.

Để tham khảo, hãy xem tổng số lần cấp phát cho hàm này mà không có PGO:

```
$ go tool pprof -sample_index=alloc_objects -top heap.nopgo.pprof | grep mdurl.Parse
   4974135  3.42% 68.60%    4974135  3.42%  gitlab.com/golang-commonmark/mdurl.Parse
```

Tổng số trước là 4974135, nghĩa là `mdurl.Parse` đã loại bỏ 100% số lần cấp phát!

Quay lại differential profile, hãy thu thập thêm một chút bối cảnh:

```
(pprof) peek mdurl.Parse
Showing nodes accounting for -12257184, 8.44% of 145309950 total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                          -2956806 59.44% |   gitlab.com/golang-commonmark/markdown.normalizeLink
                                          -2017329 40.56% |   gitlab.com/golang-commonmark/markdown.normalizeLinkText
  -4974135  3.42%  3.42%   -4974135  3.42%                | gitlab.com/golang-commonmark/mdurl.Parse
----------------------------------------------------------+-------------
```

Các lời gọi đến `mdurl.Parse` đến từ `markdown.normalizeLink` và `markdown.normalizeLinkText`.

```
(pprof) list mdurl.Parse
Total: 145309950
ROUTINE ======================== gitlab.com/golang-commonmark/mdurl.Parse in /usr/local/google/home/mpratt/go/pkg/mod/gitlab.com/golang-commonmark/mdurl@v0.0.0-20191124015652-932350d1cb84/parse
.go
  -4974135   -4974135 (flat, cum)  3.42% of Total
         .          .     60:func Parse(rawurl string) (*URL, error) {
         .          .     61:   n, err := findScheme(rawurl)
         .          .     62:   if err != nil {
         .          .     63:           return nil, err
         .          .     64:   }
         .          .     65:
  -4974135   -4974135     66:   var url URL
         .          .     67:   rest := rawurl
         .          .     68:   hostless := false
         .          .     69:   if n > 0 {
         .          .     70:           url.RawScheme = rest[:n]
         .          .     71:           url.Scheme, rest = strings.ToLower(rest[:n]), rest[n+1:]
```

Source đầy đủ cho các hàm này và các caller có thể tìm thấy tại:

* [`mdurl.Parse`](https://gitlab.com/golang-commonmark/mdurl/-/blob/bd573caec3d827ead19e40b1f141a3802d956710/parse.go#L60)
* [`markdown.normalizeLink`](https://gitlab.com/golang-commonmark/markdown/-/blob/fd7971701a0cab12e9347109a4c889f5c0a1a479/util.go#L53)
* [`markdown.normalizeLinkText`](https://gitlab.com/golang-commonmark/markdown/-/blob/fd7971701a0cab12e9347109a4c889f5c0a1a479/util.go#L68)

Vậy điều gì đã xảy ra ở đây? Trong một build không có PGO, `mdurl.Parse` được coi là quá lớn để đủ điều kiện inlining.
Tuy nhiên, vì profile PGO của chúng ta chỉ ra rằng các lời gọi đến hàm này là hot, trình biên dịch đã inline chúng.
Chúng ta có thể thấy điều này từ chú thích "(inline)" trong các profile:

```
$ go tool pprof -top cpu.nopgo.pprof | grep mdurl.Parse
     0.36s   0.3% 63.76%      2.75s  2.32%  gitlab.com/golang-commonmark/mdurl.Parse
$ go tool pprof -top cpu.withpgo.pprof | grep mdurl.Parse
     0.55s  0.48% 58.12%      2.03s  1.76%  gitlab.com/golang-commonmark/mdurl.Parse (inline)
```

`mdurl.Parse` tạo một `URL` như là biến cục bộ trên dòng 66 (`var url URL`), và sau đó trả về một con trỏ đến biến đó trên dòng 145 (`return &url, nil`).
Thông thường điều này yêu cầu biến được cấp phát trên heap, vì tham chiếu đến nó sống qua lần trả về hàm.
Tuy nhiên, khi `mdurl.Parse` được inline vào `markdown.normalizeLink`, trình biên dịch có thể quan sát rằng biến không escape ra ngoài `normalizeLink`, điều này cho phép trình biên dịch cấp phát nó trên stack.
`markdown.normalizeLinkText` tương tự với `markdown.normalizeLink`.

Việc giảm lớn thứ hai được hiển thị trong profile, từ `mdurl.(*URL).String` là trường hợp tương tự của việc loại bỏ một escape sau khi inline.

Trong các trường hợp này, chúng ta có được hiệu năng cải thiện thông qua việc ít cấp phát heap hơn.
Một phần sức mạnh của PGO và các tối ưu hóa trình biên dịch nói chung là các tác động đến việc cấp phát không phải là một phần của triển khai PGO của trình biên dịch chút nào.
Thay đổi duy nhất mà PGO thực hiện là cho phép inlining các lời gọi hàm hot này.
Tất cả các tác động đến phân tích escape và cấp phát heap là các tối ưu hóa tiêu chuẩn áp dụng cho bất kỳ build nào.
Hành vi escape được cải thiện là một tác động downstream tuyệt vời của inlining, nhưng đó không phải là tác động duy nhất.
Nhiều tối ưu hóa có thể tận dụng inlining.
Ví dụ, lan truyền hằng số có thể đơn giản hóa code trong một hàm sau khi inline khi một số đầu vào là hằng số.

### Devirtualization

Ngoài inlining, mà chúng ta đã thấy trong ví dụ trên, PGO cũng có thể thúc đẩy devirtualization có điều kiện các lời gọi interface.

Trước khi đến PGO-driven devirtualization, hãy lùi lại và định nghĩa "devirtualization" nói chung.
Giả sử bạn có code trông như thế này:

```
f, _ := os.Open("foo.txt")
var r io.Reader = f
r.Read(b)
```

Ở đây chúng ta có một lời gọi đến phương thức interface `io.Reader` là `Read`.
Vì các interface có thể có nhiều triển khai, trình biên dịch tạo ra một lời gọi hàm _gián tiếp_, nghĩa là nó tra cứu phương thức đúng để gọi tại thời điểm chạy từ kiểu trong giá trị interface.
Các lời gọi gián tiếp có thêm một chi phí runtime nhỏ so với các lời gọi trực tiếp, nhưng quan trọng hơn chúng loại trừ một số tối ưu hóa trình biên dịch.
Ví dụ, trình biên dịch không thể thực hiện phân tích escape trên một lời gọi gián tiếp vì nó không biết triển khai phương thức cụ thể.

Nhưng trong ví dụ trên, chúng ta _biết_ triển khai phương thức cụ thể.
Nó phải là `os.(*File).Read`, vì `*os.File` là kiểu duy nhất có thể được gán cho `r`.
Trong trường hợp này, trình biên dịch sẽ thực hiện _devirtualization_, nơi nó thay thế lời gọi gián tiếp đến `io.Reader.Read` bằng lời gọi trực tiếp đến `os.(*File).Read`, từ đó cho phép các tối ưu hóa khác.

(Bạn có thể nghĩ "code đó vô dụng, tại sao ai lại viết như vậy?" Đây là điểm hợp lý, nhưng hãy lưu ý rằng code như trên có thể là kết quả của inlining.
Giả sử `f` được truyền vào một hàm nhận đối số `io.Reader`.
Khi hàm được inline, bây giờ `io.Reader` trở thành cụ thể.)

PGO-driven devirtualization mở rộng khái niệm này đến các tình huống mà kiểu cụ thể không được biết tĩnh, nhưng profiling có thể cho thấy, ví dụ, một lời gọi `io.Reader.Read` nhắm vào `os.(*File).Read` hầu hết thời gian.
Trong trường hợp này, PGO có thể thay thế `r.Read(b)` bằng thứ gì đó như:

```
if f, ok := r.(*os.File); ok {
    f.Read(b)
} else {
    r.Read(b)
}
```

Tức là, chúng ta thêm một kiểm tra runtime cho kiểu cụ thể có nhiều khả năng xuất hiện nhất, và nếu vậy sử dụng một lời gọi cụ thể, hoặc ngược lại rơi vào lời gọi gián tiếp tiêu chuẩn.
Lợi thế ở đây là đường dẫn phổ biến (sử dụng `*os.File`) có thể được inline và có các tối ưu hóa bổ sung được áp dụng, nhưng chúng ta vẫn duy trì một đường dẫn fallback vì một profile không phải là sự đảm bảo rằng điều này sẽ luôn là trường hợp như vậy.

Trong phân tích của chúng ta về markdown server, chúng ta không thấy PGO-driven devirtualization, nhưng chúng ta cũng chỉ xem xét các khu vực bị tác động nhiều nhất.
PGO (và hầu hết các tối ưu hóa trình biên dịch) thường mang lại lợi ích của chúng trong tổng hợp của các cải tiến rất nhỏ ở nhiều nơi khác nhau, vì vậy có thể có nhiều điều đang xảy ra hơn những gì chúng ta đã xem xét.

Inlining và devirtualization là hai tối ưu hóa được thúc đẩy bởi PGO có sẵn trong Go 1.21, nhưng như chúng ta đã thấy, chúng thường mở khóa các tối ưu hóa bổ sung.
Ngoài ra, các phiên bản tương lai của Go sẽ tiếp tục cải thiện PGO với các tối ưu hóa bổ sung.

## Lời cảm ơn

Thêm tối ưu hóa dựa trên profile vào Go là nỗ lực của cả nhóm, và tôi đặc biệt muốn nêu bật những đóng góp từ Raj Barik và Jin Lin tại Uber, và Cherry Mui và Austin Clements tại Google.
Loại hợp tác liên cộng đồng này là phần quan trọng để làm cho Go trở nên tuyệt vời.
