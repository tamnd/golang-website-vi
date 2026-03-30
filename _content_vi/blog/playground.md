---
title: Bên trong Go Playground
date: 2013-12-12
by:
- Andrew Gerrand
tags:
- playground
summary: Cách Go playground hoạt động.
template: true
---

## Giới thiệu

_LƯU Ý: Bài viết này không mô tả phiên bản hiện tại của Go Playground._

Vào tháng 9 năm 2010, chúng tôi đã [giới thiệu Go Playground](/blog/introducing-go-playground),
một dịch vụ web biên dịch và thực thi code Go tùy ý và trả về
kết quả đầu ra của chương trình.

Nếu bạn là lập trình viên Go, bạn có thể đã sử dụng playground
bằng cách truy cập trực tiếp [Go Playground](/play/),
tham gia [Go Tour](/tour/),
hoặc chạy [các ví dụ có thể thực thi](/pkg/strings/#pkg-examples)
từ tài liệu Go.

Bạn cũng có thể đã sử dụng nó bằng cách nhấp vào một trong các nút "Run" trong bản trình chiếu
trên [go.dev/talks](/talks/) hoặc một bài đăng trên blog này
(chẳng hạn như [bài viết gần đây về Strings](/blog/strings)).

Trong bài viết này, chúng ta sẽ xem xét cách playground được triển khai
và tích hợp với các dịch vụ này.
Việc triển khai liên quan đến môi trường hệ điều hành và runtime biến thể,
và mô tả ở đây giả định bạn đã có một số kiến thức về lập trình hệ thống với Go.

## Tổng quan

{{image "playground/overview.png"}}

Dịch vụ playground có ba phần:

  - Một back end chạy trên các máy chủ của Google.
    Nó nhận các yêu cầu RPC, biên dịch chương trình người dùng bằng chuỗi công cụ gc,
    thực thi chương trình người dùng, và trả về kết quả đầu ra của chương trình (hoặc
    lỗi biên dịch) dưới dạng phản hồi RPC.
  - Một front end chạy trên [Google App Engine](https://cloud.google.com/appengine/docs/go/).
    Nó nhận các yêu cầu HTTP từ client và thực hiện các yêu cầu RPC tương ứng đến back end.
    Nó cũng thực hiện một số caching.
  - Một client JavaScript triển khai giao diện người dùng và thực hiện các yêu cầu HTTP đến front end.

## Back end

Bản thân chương trình back end rất đơn giản, nên chúng ta sẽ không thảo luận về việc triển khai
của nó ở đây. Phần thú vị là cách chúng ta thực thi code người dùng tùy ý một cách an toàn
trong một môi trường bảo mật trong khi vẫn cung cấp các chức năng cốt lõi như thời gian, mạng
và hệ thống tệp.

Để cô lập các chương trình người dùng khỏi cơ sở hạ tầng của Google, back end chạy chúng
dưới [Native Client](https://developers.google.com/native-client/)
(hay "NaCl"), một công nghệ do Google phát triển để cho phép thực thi an toàn
các chương trình x86 bên trong trình duyệt web. Back end sử dụng một phiên bản đặc biệt của chuỗi
công cụ gc để tạo ra các tệp thực thi NaCl.

(Chuỗi công cụ đặc biệt này đã được hợp nhất vào Go 1.3.
Để tìm hiểu thêm, hãy đọc [tài liệu thiết kế](/s/go13nacl).)

NaCl giới hạn lượng CPU và RAM mà một chương trình có thể sử dụng, và nó ngăn chặn
các chương trình truy cập mạng hoặc hệ thống tệp.
Tuy nhiên, điều này tạo ra một vấn đề.
Hỗ trợ concurrency và mạng của Go là một trong những điểm mạnh chính của nó,
và quyền truy cập vào hệ thống tệp là điều cần thiết cho nhiều chương trình.
Để thể hiện concurrency hiệu quả, chúng ta cần thời gian, và để thể hiện
mạng và hệ thống tệp, chúng ta rõ ràng cần mạng và hệ thống tệp.

Mặc dù tất cả những thứ này đều được hỗ trợ ngày nay, phiên bản đầu tiên của
playground, ra mắt vào năm 2010, không có gì trong số chúng.
Thời gian hiện tại được cố định ở ngày 10 tháng 11 năm 2009, `time.Sleep` không có hiệu lực,
và hầu hết các hàm của các package `os` và `net` đều bị thay thế bằng stub để
trả về lỗi `EINVALID`.

Một năm trước, chúng tôi đã
[triển khai thời gian giả](https://groups.google.com/d/msg/golang-nuts/JBsCrDEVyVE/30MaQsiQcWoJ)
trong playground, để các chương trình có sleep hoạt động đúng.
Một bản cập nhật gần đây hơn cho playground đã giới thiệu một ngăn xếp mạng giả
và một hệ thống tệp giả, làm cho chuỗi công cụ của playground tương tự với một
chuỗi công cụ Go bình thường.
Các cơ sở này được mô tả trong các phần sau.

### Giả lập thời gian

Các chương trình playground bị giới hạn về lượng thời gian CPU và bộ nhớ chúng có thể
sử dụng, nhưng chúng cũng bị hạn chế về lượng thời gian thực mà chúng có thể sử dụng.
Điều này là do mỗi chương trình đang chạy tiêu thụ tài nguyên trên back end
và bất kỳ cơ sở hạ tầng stateful nào giữa nó và client.
Giới hạn thời gian chạy của mỗi chương trình playground làm cho dịch vụ của chúng tôi
dễ dự đoán hơn và bảo vệ chúng tôi khỏi các cuộc tấn công từ chối dịch vụ.

Nhưng những hạn chế này trở nên ngột ngạt khi chạy code sử dụng thời gian.
Bài nói chuyện [Go Concurrency Patterns](/talks/2012/concurrency.slide)
thể hiện concurrency với các ví dụ sử dụng các hàm thời gian như
[`time.Sleep`](/pkg/time/#Sleep) và
[`time.After`](/pkg/time/#After).
Khi chạy dưới các phiên bản đầu tiên của playground, các lệnh sleep của những chương trình này
sẽ không có hiệu lực và hành vi của chúng sẽ trở nên kỳ lạ (và đôi khi sai).

Bằng cách sử dụng một thủ thuật thông minh, chúng ta có thể khiến một chương trình Go _nghĩ_ rằng nó
đang ngủ, khi thực tế các lần ngủ không mất thời gian nào cả.
Để giải thích thủ thuật, trước tiên chúng ta cần hiểu cách scheduler quản lý
các goroutine đang ngủ.

Khi một goroutine gọi `time.Sleep` (hoặc tương tự), scheduler thêm một bộ hẹn giờ vào
một heap các bộ hẹn giờ đang chờ xử lý và đưa goroutine đi ngủ.
Trong khi đó, một goroutine timer đặc biệt quản lý heap đó.
Khi goroutine timer khởi động, nó yêu cầu scheduler đánh thức nó
khi bộ hẹn giờ đang chờ tiếp theo sắp kích hoạt và sau đó đi ngủ.
Khi nó thức dậy, nó kiểm tra những bộ hẹn giờ nào đã hết hạn, đánh thức
các goroutine phù hợp và tiếp tục ngủ.

Thủ thuật là thay đổi điều kiện đánh thức goroutine timer.
Thay vì đánh thức nó sau một khoảng thời gian cụ thể, chúng ta sửa đổi scheduler để
chờ một deadlock; trạng thái mà tất cả các goroutine đều bị chặn.

Phiên bản playground của runtime duy trì đồng hồ nội bộ riêng của nó. Khi
scheduler đã sửa đổi phát hiện một deadlock, nó kiểm tra xem có bộ hẹn giờ nào đang chờ không.
Nếu có, nó tiến đồng hồ nội bộ đến thời điểm kích hoạt của bộ hẹn giờ sớm nhất
và sau đó đánh thức goroutine timer. Thực thi tiếp tục và
chương trình tin rằng thời gian đã trôi qua, khi thực tế lần ngủ gần như là tức thời.

Những thay đổi này đối với scheduler có thể được tìm thấy trong [`proc.c`](/cl/73110043)
và [`time.goc`](/cl/73110043).

Thời gian giả khắc phục vấn đề cạn kiệt tài nguyên trên back end, nhưng còn đầu ra
chương trình thì sao? Sẽ thật kỳ lạ khi thấy một chương trình có sleep chạy đến
hoàn thành một cách đúng đắn mà không mất thời gian nào.

Chương trình sau đây in thời gian hiện tại mỗi giây và sau đó thoát sau
ba giây. Hãy thử chạy nó.

{{play "playground/time.go" `/^func main/` `$`}}

Điều này hoạt động như thế nào? Đó là sự cộng tác giữa back end, front end và client.

Chúng ta ghi lại thời điểm của mỗi lần ghi vào đầu ra chuẩn và lỗi chuẩn và
cung cấp cho client. Sau đó client có thể "phát lại" các lần ghi với
thời gian chính xác, để đầu ra xuất hiện như thể chương trình đang chạy
cục bộ.

Package `runtime` của playground cung cấp một
[hàm `write` đặc biệt](https://github.com/golang/go/blob/go1.3/src/pkg/runtime/sys_nacl_amd64p32.s#L54)
bao gồm một "playback header" nhỏ trước mỗi lần ghi.
Playback header bao gồm một chuỗi magic, thời gian hiện tại và
độ dài của dữ liệu ghi. Một lần ghi với playback header có cấu trúc như sau:

{{raw `
	0 0 P B <8-byte time> <4-byte data length> <data>
`}}

Đầu ra thô của chương trình trên trông như thế này:

	\x00\x00PB\x11\x74\xef\xed\xe6\xb3\x2a\x00\x00\x00\x00\x1e2009-11-10 23:00:01 +0000 UTC
	\x00\x00PB\x11\x74\xef\xee\x22\x4d\xf4\x00\x00\x00\x00\x1e2009-11-10 23:00:02 +0000 UTC
	\x00\x00PB\x11\x74\xef\xee\x5d\xe8\xbe\x00\x00\x00\x00\x1e2009-11-10 23:00:03 +0000 UTC

Front end phân tích đầu ra này thành một loạt các sự kiện
và trả về danh sách các sự kiện cho client dưới dạng đối tượng JSON:

	{
		"Errors": "",
		"Events": [
			{
				"Delay": 1000000000,
				"Message": "2009-11-10 23:00:01 +0000 UTC\n"
			},
			{
				"Delay": 1000000000,
				"Message": "2009-11-10 23:00:02 +0000 UTC\n"
			},
			{
				"Delay": 1000000000,
				"Message": "2009-11-10 23:00:03 +0000 UTC\n"
			}
		]
	}

Client JavaScript (chạy trong trình duyệt web của người dùng) sau đó phát lại
các sự kiện bằng cách sử dụng các khoảng trễ đã cung cấp.
Đối với người dùng, có vẻ như chương trình đang chạy trong thời gian thực.

### Giả lập hệ thống tệp

Các chương trình được xây dựng với chuỗi công cụ NaCl của Go không thể truy cập hệ thống tệp
của máy cục bộ. Thay vào đó, các hàm liên quan đến tệp của package `syscall`
(`Open`, `Read`, `Write`, v.v.) hoạt động trên một hệ thống tệp trong bộ nhớ
được triển khai bởi chính package `syscall`.
Vì package `syscall` là giao diện giữa code Go và kernel hệ điều hành,
các chương trình người dùng nhìn thấy hệ thống tệp hoàn toàn giống như
một hệ thống tệp thực.

Chương trình ví dụ sau ghi dữ liệu vào một tệp, sau đó sao chép
nội dung của nó vào đầu ra chuẩn. Hãy thử chạy nó. (Bạn cũng có thể chỉnh sửa nó!)

{{play "playground/os.go" `/^func main/` `$`}}

Khi một tiến trình khởi động, hệ thống tệp được điền với một số thiết bị dưới
`/dev` và một thư mục `/tmp` trống. Chương trình có thể thao tác với hệ thống tệp như bình thường,
nhưng khi tiến trình thoát, mọi thay đổi đối với hệ thống tệp đều bị mất.

Cũng có một cơ chế để tải một tệp zip vào hệ thống tệp trong quá trình khởi tạo
(xem [`unzip_nacl.go`](https://github.com/golang/go/blob/go1.3/src/pkg/syscall/unzip_nacl.go)).
Cho đến nay, chúng tôi chỉ sử dụng tiện ích unzip để cung cấp các tệp dữ liệu cần thiết
để chạy các bài kiểm tra thư viện chuẩn, nhưng chúng tôi có kế hoạch cung cấp cho các chương trình
playground một bộ tệp có thể được sử dụng trong các ví dụ tài liệu, bài đăng blog
và Go Tour.

Việc triển khai có thể được tìm thấy trong các tệp
[`fs_nacl.go`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fs_nacl.go) và
[`fd_nacl.go`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fd_nacl.go)
(những tệp này, nhờ hậu tố `_nacl`, chỉ được tích hợp vào package `syscall`
khi `GOOS` được đặt thành `nacl`).

Bản thân hệ thống tệp được biểu diễn bởi
[`fsys` struct](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fs_nacl.go#L26),
trong đó một instance toàn cục (có tên `fs`) được tạo trong quá trình khởi tạo.
Các hàm liên quan đến tệp khác nhau sau đó hoạt động trên `fs` thay vì thực hiện
lời gọi hệ thống thực sự.
Ví dụ, đây là hàm [`syscall.Open`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fs_nacl.go#L473):

	func Open(path string, openmode int, perm uint32) (fd int, err error) {
		fs.mu.Lock()
		defer fs.mu.Unlock()
		f, err := fs.open(path, openmode, perm&0777|S_IFREG)
		if err != nil {
			return -1, err
		}
		return newFD(f), nil
	}

Các file descriptor được theo dõi bởi một slice toàn cục có tên
[`files`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fd_nacl.go#L17).
Mỗi file descriptor tương ứng với một [`file`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fd_nacl.go#L23)
và mỗi `file` cung cấp một giá trị triển khai interface [`fileImpl`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fd_nacl.go#L30).
Có một số triển khai của interface:

  - các tệp thông thường và thiết bị (như `/dev/random`) được biểu diễn bởi [`fsysFile`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fs_nacl.go#L58),
  - đầu vào chuẩn, đầu ra chuẩn và lỗi chuẩn là các instance của [`naclFile`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/fd_nacl.go#L216),
    sử dụng các lời gọi hệ thống để tương tác với các tệp thực sự (đây là
    cách duy nhất để chương trình playground tương tác với thế giới bên ngoài),
  - các network socket có triển khai riêng của chúng, được thảo luận trong phần tiếp theo.

### Giả lập mạng

Giống như hệ thống tệp, ngăn xếp mạng của playground là một giả lập trong tiến trình
được triển khai bởi package `syscall`. Nó cho phép các dự án playground sử dụng
giao diện loopback (`127.0.0.1`). Các yêu cầu đến các host khác sẽ thất bại.

Để xem ví dụ có thể thực thi, hãy chạy chương trình sau. Nó lắng nghe trên một cổng TCP,
chờ một kết nối đến, sao chép dữ liệu từ kết nối đó vào
đầu ra chuẩn và thoát. Trong một goroutine khác, nó thực hiện một kết nối đến cổng
đang lắng nghe, ghi một chuỗi vào kết nối và đóng nó.

{{play "playground/net.go" `/^func main/` `$`}}

Giao diện với mạng phức tạp hơn giao diện với tệp, vì vậy
việc triển khai mạng giả lớn hơn và phức tạp hơn hệ thống tệp giả. Nó phải
mô phỏng thời gian chờ đọc và ghi, các loại địa chỉ và giao thức khác nhau, v.v.

Việc triển khai có thể được tìm thấy trong [`net_nacl.go`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/net_nacl.go).
Một điểm tốt để bắt đầu đọc là [`netFile`](https://github.com/golang/go/blob/2197321db1dd997165c0091ba2bcb3b6be7633d0/src/syscall/net_nacl.go#L461),
việc triển khai network socket của interface `fileImpl`.

## Front end

Front end của playground là một chương trình đơn giản khác (ngắn hơn 100 dòng).
Nó nhận các yêu cầu HTTP từ client, thực hiện các yêu cầu RPC đến back end
và thực hiện một số caching.

Front end phục vụ một HTTP handler tại `https://golang.org/compile`.
Handler mong đợi một yêu cầu POST với trường `body`
(chương trình Go cần chạy) và một trường `version` tùy chọn
(đối với hầu hết các client, đây nên là `"2"`).

Khi front end nhận được một yêu cầu biên dịch, đầu tiên nó kiểm tra
[memcache](https://developers.google.com/appengine/docs/memcache/)
để xem liệu nó có cached kết quả của một lần biên dịch trước đó của source code đó không.
Nếu tìm thấy, nó trả về phản hồi đã được cached.
Cache ngăn chặn các chương trình phổ biến như những chương trình trên
[Go home page](/) khỏi việc làm quá tải các back end.
Nếu không có phản hồi đã được cached, front end thực hiện một yêu cầu RPC đến back
end, lưu trữ phản hồi trong memcache, phân tích các sự kiện phát lại và trả về
một đối tượng JSON cho client dưới dạng phản hồi HTTP (như mô tả ở trên).

## Client

Các trang web khác nhau sử dụng playground đều chia sẻ một số code JavaScript
chung để thiết lập giao diện người dùng (các hộp code và đầu ra, nút run
và v.v.) và giao tiếp với front end của playground.

Việc triển khai này nằm trong tệp
[`playground.js`](https://github.com/golang/tools/blob/f8e922be8efeabd06a510065ca5836b62fa10b9a/godoc/static/playground.js)
trong kho lưu trữ `go.tools`, có thể được import từ
package [`golang.org/x/tools/godoc/static`](https://godoc.org/golang.org/x/tools/godoc/static).
Một phần của nó rõ ràng và một phần hơi lộn xộn, vì nó là kết quả của
việc hợp nhất một số triển khai khác nhau của code client.

Hàm [`playground`](https://github.com/golang/tools/blob/f8e922be8efeabd06a510065ca5836b62fa10b9a/godoc/static/playground.js#L227)
nhận một số phần tử HTML và biến chúng thành một
widget playground tương tác. Bạn nên sử dụng hàm này nếu muốn đặt
playground trên trang web của riêng bạn (xem 'Các client khác' bên dưới).

Interface [`Transport`](https://github.com/golang/tools/blob/f8e922be8efeabd06a510065ca5836b62fa10b9a/godoc/static/playground.js#L6)
(không được định nghĩa chính thức, đây là JavaScript)
trừu tượng hóa giao diện người dùng khỏi phương tiện giao tiếp với web front end.
[`HTTPTransport`](https://github.com/golang/tools/blob/f8e922be8efeabd06a510065ca5836b62fa10b9a/godoc/static/playground.js#L43)
là một triển khai của `Transport` sử dụng giao thức dựa trên HTTP
được mô tả trước đó.
[`SocketTransport`](https://github.com/golang/tools/blob/f8e922be8efeabd06a510065ca5836b62fa10b9a/godoc/static/playground.js#L115)
là một triển khai khác sử dụng WebSocket (xem 'Sử dụng offline' bên dưới).

Để tuân thủ [chính sách same-origin](https://en.wikipedia.org/wiki/Same-origin_policy),
các máy chủ web khác nhau (godoc chẳng hạn) proxy các yêu cầu đến
`/compile` qua dịch vụ playground tại `https://golang.org/compile`.
Package chung [`golang.org/x/tools/playground`](https://godoc.org/golang.org/x/tools/playground)
thực hiện proxy này.

## Sử dụng offline

Cả [Go Tour](/tour/) và
[Present Tool](https://godoc.org/golang.org/x/tools/present) đều có thể được
chạy offline. Điều này rất tốt cho những người có kết nối internet hạn chế
hoặc những người thuyết trình tại hội nghị không thể (và _không nên_) dựa vào
kết nối internet đang hoạt động.

Để chạy offline, các công cụ chạy phiên bản back end playground của riêng chúng trên
máy cục bộ. Back end sử dụng một chuỗi công cụ Go thông thường mà không có
những sửa đổi đã đề cập và sử dụng WebSocket để giao tiếp với
client.

Việc triển khai back end WebSocket có thể được tìm thấy trong
package [`golang.org/x/tools/playground/socket`](https://godoc.org/golang.org/x/tools/playground/socket).
Bài nói chuyện [Inside Present](/talks/2012/insidepresent.slide#1) thảo luận chi tiết về code này.

## Các client khác

Dịch vụ playground được sử dụng bởi nhiều hơn chỉ dự án Go chính thức
([Go by Example](https://gobyexample.com/) là một ví dụ khác)
và chúng tôi rất vui khi bạn sử dụng nó trên trang web của riêng bạn. Tất cả những gì chúng tôi yêu cầu là
bạn [liên hệ với chúng tôi trước](mailto:golang-dev@googlegroups.com),
sử dụng một user agent duy nhất trong các yêu cầu của bạn (để chúng tôi có thể nhận dạng bạn), và rằng
dịch vụ của bạn mang lại lợi ích cho cộng đồng Go.

## Kết luận

Từ godoc đến tour đến chính blog này, playground đã trở thành
một phần thiết yếu trong câu chuyện tài liệu Go của chúng tôi. Với những bổ sung gần đây
của hệ thống tệp giả và ngăn xếp mạng, chúng tôi hào hứng mở rộng
các tài liệu học tập của mình để bao gồm những lĩnh vực đó.

Nhưng, cuối cùng, playground chỉ là bề mặt của tảng băng chìm.
Với hỗ trợ Native Client được lên kế hoạch cho Go 1.3,
chúng tôi mong chờ xem những gì cộng đồng có thể làm với nó.

_Bài viết này là phần 12 của_
[Go Advent Calendar](https://blog.gopheracademy.com/go-advent-2013),
_một loạt bài đăng blog hàng ngày trong suốt tháng 12._
