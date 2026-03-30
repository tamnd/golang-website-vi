---
title: Giới thiệu bộ phát hiện race condition của Go
date: 2013-06-26
by:
- Dmitry Vyukov
- Andrew Gerrand
tags:
- concurrency
- technical
summary: Cách thức và lý do sử dụng bộ phát hiện race condition của Go để cải thiện chương trình.
template: true
---

## Giới thiệu

[Race condition](http://en.wikipedia.org/wiki/Race_condition) là một trong
những lỗi lập trình nguy hiểm và khó nắm bắt nhất. Chúng thường gây ra các
lỗi bất thường và khó hiểu, thường xảy ra lâu sau khi code đã được triển khai
lên môi trường production. Mặc dù các cơ chế concurrent của Go giúp dễ dàng
viết code concurrent sạch, nhưng chúng không ngăn chặn được race condition.
Cần có sự cẩn thận, tỉ mỉ và kiểm thử. Và các công cụ có thể giúp ích.

Chúng tôi vui mừng thông báo rằng Go 1.1 bao gồm một
[bộ phát hiện race condition](/doc/articles/race_detector.html),
một công cụ mới để tìm race condition trong code Go.
Hiện tại nó có sẵn cho các hệ thống Linux, OS X và Windows
với bộ xử lý x86 64-bit.

Bộ phát hiện race condition dựa trên
[thư viện runtime ThreadSanitizer](https://github.com/google/sanitizers) của C/C++,
vốn đã được sử dụng để phát hiện nhiều lỗi trong codebase nội bộ của Google và trong
[Chromium](http://www.chromium.org/).
Công nghệ này được tích hợp với Go vào tháng 9 năm 2012; kể từ đó nó đã phát hiện
[42 race condition](https://github.com/golang/go/issues?utf8=%E2%9C%93&q=ThreadSanitizer)
trong thư viện chuẩn. Nó hiện là một phần của quy trình build liên tục của chúng tôi,
nơi nó tiếp tục bắt được race condition khi chúng xuất hiện.

## Cách hoạt động

Bộ phát hiện race condition được tích hợp với chuỗi công cụ go. Khi cờ
dòng lệnh `-race` được đặt, trình biên dịch sẽ công cụ hóa tất cả các truy
cập bộ nhớ với code ghi lại thời điểm và cách thức bộ nhớ được truy cập, trong
khi thư viện runtime theo dõi các truy cập không đồng bộ vào các biến được
chia sẻ. Khi hành vi "race" như vậy được phát hiện, một cảnh báo được in ra.
(Xem [bài viết này](https://github.com/google/sanitizers/wiki/ThreadSanitizerAlgorithm)
để biết chi tiết về thuật toán.)

Do thiết kế của nó, bộ phát hiện race condition chỉ có thể phát hiện race
condition khi chúng thực sự được kích hoạt bởi code đang chạy, có nghĩa là
điều quan trọng là chạy các binary có bật race detector dưới các workload thực
tế. Tuy nhiên, các binary có bật race detector có thể sử dụng gấp mười lần
CPU và bộ nhớ, vì vậy việc bật race detector thường xuyên là không thực tế.
Một cách thoát khỏi tình huống này là chạy một số bài kiểm thử với race
detector được bật. Kiểm thử tải và kiểm thử tích hợp là các ứng cử viên tốt,
vì chúng có xu hướng thực thi các phần concurrent của code. Một cách khác sử
dụng workload production là triển khai một phiên bản duy nhất có bật race
detector trong một pool các máy chủ đang chạy.

## Sử dụng bộ phát hiện race condition

Bộ phát hiện race condition được tích hợp hoàn toàn với chuỗi công cụ Go.
Để build code với race detector được bật, chỉ cần thêm cờ `-race` vào dòng
lệnh:

	$ go test -race mypkg    // test the package
	$ go run -race mysrc.go  // compile and run the program
	$ go build -race mycmd   // build the command
	$ go install -race mypkg // install the package

Để tự thử bộ phát hiện race condition, hãy sao chép chương trình ví dụ này
vào `racy.go`:

{{raw `
	package main

	import "fmt"

	func main() {
		done := make(chan bool)
		m := make(map[string]string)
		m["name"] = "world"
		go func() {
			m["name"] = "data race"
			done <- true
		}()
		fmt.Println("Hello,", m["name"])
		<-done
	}
`}}

Sau đó chạy nó với race detector được bật:

	$ go run -race racy.go

## Ví dụ

Đây là hai ví dụ về các vấn đề thực tế được phát hiện bởi race detector.

### Ví dụ 1: Timer.Reset

Ví dụ đầu tiên là một phiên bản đơn giản hóa của lỗi thực tế được tìm thấy
bởi race detector. Nó sử dụng một timer để in một thông điệp sau một khoảng
thời gian ngẫu nhiên từ 0 đến 1 giây. Nó làm điều này liên tục trong năm giây.
Nó sử dụng [`time.AfterFunc`](/pkg/time/#AfterFunc) để tạo một
[`Timer`](/pkg/time/#Timer) cho thông điệp đầu tiên và sau đó
sử dụng phương thức [`Reset`](/pkg/time/#Timer.Reset) để
lên lịch thông điệp tiếp theo, tái sử dụng `Timer` mỗi lần.

{{play "race-detector/timer.go" `/func main/` `$` 0}}

Code này trông có vẻ hợp lý, nhưng trong một số trường hợp nó thất bại theo
một cách đáng ngạc nhiên:

	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xb code=0x1 addr=0x8 pc=0x41e38a]

	goroutine 4 [running]:
	time.stopTimer(0x8, 0x12fe6b35d9472d96)
		src/pkg/runtime/ztime_linux_amd64.c:35 +0x25
	time.(*Timer).Reset(0x0, 0x4e5904f, 0x1)
		src/pkg/time/sleep.go:81 +0x42
	main.func·001()
		race.go:14 +0xe3
	created by time.goFunc
		src/pkg/time/sleep.go:122 +0x48

Chuyện gì đang xảy ra ở đây? Chạy chương trình với race detector được bật sẽ
cho thấy nhiều hơn:

	==================
	WARNING: DATA RACE
	Read by goroutine 5:
	  main.func·001()
	     race.go:16 +0x169

	Previous write by goroutine 1:
	  main.main()
	      race.go:14 +0x174

	Goroutine 5 (running) created at:
	  time.goFunc()
	      src/pkg/time/sleep.go:122 +0x56
	  timerproc()
	     src/pkg/runtime/ztime_linux_amd64.c:181 +0x189
	==================

Race detector cho thấy vấn đề: việc đọc và ghi không được đồng bộ hóa của
biến `t` từ các goroutine khác nhau. Nếu thời gian chờ của timer ban đầu rất
nhỏ, hàm timer có thể kích hoạt trước khi goroutine chính đã gán một giá trị
cho `t` và do đó lời gọi đến `t.Reset` được thực hiện với `t` là nil.

Để sửa race condition, chúng ta thay đổi code để chỉ đọc và ghi biến `t` từ
goroutine chính:

{{play "race-detector/timer-fixed.go" `/func main/` `/^}/` 0}}

Ở đây goroutine chính hoàn toàn chịu trách nhiệm cài đặt và đặt lại `Timer`
`t`, và một channel reset mới truyền đạt nhu cầu đặt lại timer theo cách an
toàn với thread.

Một cách đơn giản hơn nhưng kém hiệu quả hơn là
[tránh tái sử dụng timer](/play/p/kuWTrY0pS4).

### Ví dụ 2: ioutil.Discard

Ví dụ thứ hai tinh tế hơn.

Đối tượng
[`Discard`](/pkg/io/ioutil/#Discard) của gói `ioutil` cài đặt
[`io.Writer`](/pkg/io/#Writer),
nhưng loại bỏ tất cả dữ liệu được ghi vào nó.
Hãy nghĩ nó như `/dev/null`: một nơi để gửi dữ liệu bạn cần đọc nhưng không
muốn lưu trữ.
Nó thường được sử dụng với [`io.Copy`](/pkg/io/#Copy)
để rút cạn một reader, như thế này:

	io.Copy(ioutil.Discard, reader)

Vào tháng 7 năm 2011, nhóm Go nhận thấy rằng việc sử dụng `Discard` theo
cách này là không hiệu quả: hàm `Copy` cấp phát một buffer nội bộ 32 kB mỗi
lần được gọi, nhưng khi được sử dụng với `Discard` thì buffer đó là không cần
thiết vì chúng ta chỉ vứt bỏ dữ liệu đã đọc đi. Chúng tôi nghĩ rằng cách sử
dụng thành ngữ của `Copy` và `Discard` không nên tốn kém như vậy.

Cách sửa rất đơn giản. Nếu `Writer` đã cho cài đặt phương thức `ReadFrom`,
một lời gọi `Copy` như thế này:

	io.Copy(writer, reader)

sẽ được ủy quyền cho lời gọi tiềm năng hiệu quả hơn này:

	writer.ReadFrom(reader)

Chúng tôi
[thêm một phương thức ReadFrom](/cl/4817041)
vào kiểu bên dưới của Discard, có một buffer nội bộ được chia sẻ giữa tất cả
người dùng của nó. Chúng tôi biết đây về lý thuyết là một race condition,
nhưng vì tất cả các lần ghi vào buffer đều bị vứt bỏ, chúng tôi không nghĩ
nó quan trọng.

Khi race detector được cài đặt, nó ngay lập tức
[đánh dấu code này](/issue/3970) là race condition.
Một lần nữa, chúng tôi xem xét rằng code có thể có vấn đề, nhưng quyết định
rằng race condition không phải là "có thật". Để tránh "dương tính giả" trong
build của chúng tôi, chúng tôi đã cài đặt
[một phiên bản không có race condition](/cl/6624059)
chỉ được bật khi race detector đang chạy.

Nhưng vài tháng sau, [Brad](https://bradfitz.com/) gặp phải một
[lỗi khó chịu và kỳ lạ](/issue/4589).
Sau vài ngày debug, ông đã thu hẹp nguyên nhân xuống còn một race condition
thực sự do `ioutil.Discard` gây ra.

Đây là code có race condition đã biết trong `io/ioutil`, nơi `Discard` là một
`devNull` chia sẻ một buffer duy nhất giữa tất cả người dùng của nó.

{{code "race-detector/blackhole.go" `/var blackHole/` `/^}/`}}

Chương trình của Brad bao gồm một kiểu `trackDigestReader`, bọc một `io.Reader`
và ghi lại hash digest của những gì nó đọc.

	type trackDigestReader struct {
		r io.Reader
		h hash.Hash
	}

	func (t trackDigestReader) Read(p []byte) (n int, err error) {
		n, err = t.r.Read(p)
		t.h.Write(p[:n])
		return
	}

Ví dụ, nó có thể được sử dụng để tính toán hash SHA-1 của một tệp khi đọc nó:

	tdr := trackDigestReader{r: file, h: sha1.New()}
	io.Copy(writer, tdr)
	fmt.Printf("File hash: %x", tdr.h.Sum(nil))

Trong một số trường hợp không có nơi nào để ghi dữ liệu nhưng vẫn cần hash
tệp, vì vậy `Discard` sẽ được sử dụng:

	io.Copy(ioutil.Discard, tdr)

Nhưng trong trường hợp này, buffer `blackHole` không chỉ là một lỗ đen; nó là
một nơi hợp pháp để lưu trữ dữ liệu giữa việc đọc từ nguồn `io.Reader` và
ghi vào `hash.Hash`. Với nhiều goroutine hash các tệp đồng thời, mỗi goroutine
chia sẻ cùng buffer `blackHole`, race condition biểu hiện bằng cách làm hỏng
dữ liệu giữa quá trình đọc và hash. Không có lỗi hoặc panic nào xảy ra, nhưng
các hash thì sai. Thật khó chịu!

	func (t trackDigestReader) Read(p []byte) (n int, err error) {
		// the buffer p is blackHole
		n, err = t.r.Read(p)
		// p may be corrupted by another goroutine here,
		// between the Read above and the Write below
		t.h.Write(p[:n])
		return
	}

Lỗi cuối cùng đã được
[sửa](/cl/7011047)
bằng cách cấp cho mỗi lần sử dụng `ioutil.Discard` một buffer riêng, loại bỏ
race condition trên buffer được chia sẻ.

## Kết luận

Race detector là một công cụ mạnh mẽ để kiểm tra tính đúng đắn của các chương
trình concurrent. Nó sẽ không phát ra dương tính giả, vì vậy hãy xem xét
nghiêm túc các cảnh báo của nó. Nhưng nó chỉ tốt khi bài kiểm thử của bạn
tốt; bạn phải đảm bảo rằng chúng thực thi kỹ lưỡng các thuộc tính concurrent
của code để race detector có thể làm tốt công việc của nó.

Còn chờ gì nữa? Hãy chạy `"go test -race"` trên code của bạn ngay hôm nay!
