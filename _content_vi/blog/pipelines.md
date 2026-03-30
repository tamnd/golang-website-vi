---
title: "Các mẫu đồng thời trong Go: Pipeline và hủy bỏ"
date: 2014-03-13
by:
- Sameer Ajmani
tags:
- concurrency
- pipelines
- cancellation
summary: Cách sử dụng tính đồng thời của Go để xây dựng pipeline xử lý dữ liệu.
template: true
---

## Giới thiệu

Các nguyên tố đồng thời của Go giúp dễ dàng xây dựng các pipeline dữ liệu phát trực tuyến
sử dụng I/O và nhiều CPU một cách hiệu quả. Bài viết này trình bày
các ví dụ về các pipeline như vậy, làm nổi bật những điểm tinh tế phát sinh khi các thao tác
thất bại, và giới thiệu các kỹ thuật để xử lý các thất bại một cách gọn gàng.

## Pipeline là gì?

Không có định nghĩa chính thức về pipeline trong Go; đó chỉ là một trong nhiều loại
chương trình đồng thời. Một cách không chính thức, pipeline là một loạt các _giai đoạn_ được kết nối
bởi các channel, trong đó mỗi giai đoạn là một nhóm goroutine chạy cùng một
hàm. Trong mỗi giai đoạn, các goroutine

  - nhận các giá trị từ _upstream_ qua các channel _inbound_
  - thực hiện một hàm nào đó trên dữ liệu đó, thường tạo ra các giá trị mới
  - gửi các giá trị _downstream_ qua các channel _outbound_

Mỗi giai đoạn có bất kỳ số lượng channel inbound và outbound nào, ngoại trừ
giai đoạn đầu tiên và cuối cùng, chỉ có channel outbound hoặc inbound,
tương ứng. Giai đoạn đầu tiên đôi khi được gọi là _source_ hoặc
_producer_; giai đoạn cuối cùng là _sink_ hoặc _consumer_.

Chúng ta sẽ bắt đầu với một ví dụ pipeline đơn giản để giải thích các ý tưởng và kỹ thuật.
Sau đó, chúng ta sẽ trình bày một ví dụ thực tế hơn.

## Bình phương các số

Xem xét một pipeline với ba giai đoạn.

Giai đoạn đầu tiên, `gen`, là một hàm chuyển đổi một danh sách các số nguyên thành
channel phát ra các số nguyên trong danh sách đó. Hàm `gen` khởi động một
goroutine gửi các số nguyên trên channel và đóng channel khi tất cả
các giá trị đã được gửi:

{{code "pipelines/square.go" `/func gen/` `/^}/`}}

Giai đoạn thứ hai, `sq`, nhận các số nguyên từ một channel và trả về
channel phát ra bình phương của mỗi số nguyên nhận được. Sau khi
channel inbound được đóng và giai đoạn này đã gửi tất cả các giá trị
downstream, nó đóng channel outbound:

{{code "pipelines/square.go" `/func sq/` `/^}/`}}

Hàm `main` thiết lập pipeline và chạy giai đoạn cuối cùng: nó nhận
các giá trị từ giai đoạn thứ hai và in từng giá trị, cho đến khi channel bị đóng:

{{code "pipelines/square.go" `/func main/` `/^}/`}}

Vì `sq` có cùng kiểu cho channel inbound và outbound của nó, chúng ta
có thể kết hợp nó bất kỳ số lần nào. Chúng ta cũng có thể viết lại `main` dưới dạng
vòng lặp range, giống như các giai đoạn khác:

{{code "pipelines/square2.go" `/func main/` `/^}/`}}

## Fan-out, fan-in

Nhiều hàm có thể đọc từ cùng một channel cho đến khi channel đó bị đóng;
điều này được gọi là _fan-out_. Điều này cung cấp một cách để phân phối công việc trong một nhóm
worker để song song hóa việc sử dụng CPU và I/O.

Một hàm có thể đọc từ nhiều đầu vào và tiến hành cho đến khi tất cả đều đóng bằng cách
ghép kênh các channel đầu vào vào một channel duy nhất được đóng khi tất cả các
đầu vào bị đóng. Điều này được gọi là _fan-in_.

Chúng ta có thể thay đổi pipeline để chạy hai phiên bản `sq`, mỗi phiên bản đọc từ cùng
một channel đầu vào. Chúng ta giới thiệu một hàm mới, _merge_, để fan-in các
kết quả:

{{code "pipelines/sqfan.go" `/func main/` `/^}/`}}

Hàm `merge` chuyển đổi một danh sách các channel thành một channel duy nhất bằng cách khởi động
một goroutine cho mỗi channel inbound sao chép các giá trị vào channel outbound duy nhất.
Khi tất cả các goroutine `output` đã được khởi động, `merge` khởi động thêm một
goroutine nữa để đóng channel outbound sau khi tất cả các lần gửi trên channel đó được
hoàn thành.

Gửi trên một channel đã đóng sẽ panic, vì vậy điều quan trọng là đảm bảo tất cả các lần gửi
đã xong trước khi gọi close. Kiểu
[`sync.WaitGroup`](/pkg/sync/#WaitGroup)
cung cấp một cách đơn giản để sắp xếp đồng bộ hóa này:

{{code "pipelines/sqfan.go" `/func merge/` `/^}/`}}

## Dừng sớm

Có một mẫu trong các hàm pipeline của chúng ta:

  - các giai đoạn đóng channel outbound khi tất cả các thao tác gửi đã hoàn thành.
  - các giai đoạn tiếp tục nhận giá trị từ các channel inbound cho đến khi các channel đó bị đóng.

Mẫu này cho phép mỗi giai đoạn nhận được viết dưới dạng vòng lặp `range` và
đảm bảo tất cả các goroutine thoát khi tất cả các giá trị đã được gửi thành công
downstream.

Nhưng trong các pipeline thực tế, các giai đoạn không phải lúc nào cũng nhận tất cả các giá trị
inbound. Đôi khi điều này là theo thiết kế: receiver chỉ có thể cần một
tập hợp con các giá trị để tiếp tục. Thường xuyên hơn, một giai đoạn thoát sớm
vì một giá trị inbound đại diện cho một lỗi ở giai đoạn trước đó. Trong
cả hai trường hợp, receiver không nên phải chờ các giá trị còn lại đến,
và chúng ta muốn các giai đoạn trước đó dừng tạo ra các giá trị mà các giai đoạn sau
không cần.

Trong pipeline ví dụ của chúng ta, nếu một giai đoạn không tiêu thụ hết tất cả các giá trị inbound,
các goroutine cố gắng gửi các giá trị đó sẽ bị block vô thời hạn:

{{code "pipelines/sqleak.go" `/first value/` `/^}/`}}

Đây là rò rỉ tài nguyên: các goroutine tiêu thụ bộ nhớ và tài nguyên runtime, và
các tham chiếu heap trong các stack goroutine ngăn dữ liệu bị thu gom rác.
Các goroutine không được thu gom rác; chúng phải tự thoát.

Chúng ta cần sắp xếp để các giai đoạn upstream của pipeline thoát ngay cả khi
các giai đoạn downstream không nhận tất cả các giá trị inbound. Một cách để làm điều này là
thay đổi các channel outbound để có buffer. Buffer có thể giữ một số lượng cố định
các giá trị; các thao tác gửi hoàn thành ngay lập tức nếu có chỗ trong
buffer:

{{raw `
	c := make(chan int, 2) // buffer size 2
	c <- 1  // succeeds immediately
	c <- 2  // succeeds immediately
	c <- 3  // blocks until another goroutine does <-c and receives 1
`}}

Khi số lượng giá trị sẽ được gửi được biết tại thời điểm tạo channel, buffer
có thể đơn giản hóa code. Ví dụ, chúng ta có thể viết lại `gen` để sao chép danh sách
các số nguyên vào một channel có buffer và tránh tạo một goroutine mới:

{{code "pipelines/sqbuffer.go" `/func gen/` `/^}/`}}

Quay lại các goroutine bị block trong pipeline của chúng ta, chúng ta có thể xem xét thêm
buffer vào channel outbound được trả về bởi `merge`:

{{code "pipelines/sqbuffer.go" `/func merge/` `/unchanged/`}}

Mặc dù điều này sửa goroutine bị block trong chương trình này, đây là code xấu. Việc
chọn kích thước buffer là 1 ở đây phụ thuộc vào việc biết số lượng giá trị mà `merge`
sẽ nhận và số lượng giá trị mà các giai đoạn downstream sẽ tiêu thụ. Điều này
mong manh: nếu chúng ta truyền một giá trị bổ sung cho `gen`, hoặc nếu giai đoạn downstream
đọc ít giá trị hơn, chúng ta sẽ lại có các goroutine bị block.

Thay vào đó, chúng ta cần cung cấp một cách để các giai đoạn downstream chỉ ra cho
các sender rằng họ sẽ ngừng chấp nhận đầu vào.

## Hủy bỏ tường minh

Khi `main` quyết định thoát mà không nhận tất cả các giá trị từ
`out`, nó phải nói với các goroutine ở các giai đoạn upstream để từ bỏ
các giá trị chúng đang cố gửi. Nó làm điều này bằng cách gửi các giá trị trên một
channel gọi là `done`. Nó gửi hai giá trị vì có
hai sender tiềm năng bị block:

{{code "pipelines/sqdone1.go" `/func main/` `/^}/`}}

Các goroutine gửi thay thế thao tác gửi của chúng bằng câu lệnh `select`
tiến hành khi lần gửi trên `out` xảy ra hoặc khi chúng nhận được một giá trị
từ `done`. Kiểu giá trị của `done` là struct rỗng vì giá trị
không quan trọng: sự kiện nhận là thứ chỉ ra rằng lần gửi trên `out` nên
được từ bỏ. Các goroutine `output` tiếp tục lặp trên channel inbound
của chúng, `c`, vì vậy các giai đoạn upstream không bị block. (Chúng ta sẽ thảo luận trong giây lát
về cách cho phép vòng lặp này trả về sớm.)

{{code "pipelines/sqdone1.go" `/func merge/` `/unchanged/`}}

Cách tiếp cận này có một vấn đề: _mỗi_ receiver downstream cần biết số lượng
sender upstream tiềm năng bị block và sắp xếp để báo hiệu cho những sender đó khi
trả về sớm. Việc theo dõi các số đếm này rất tẻ nhạt và dễ xảy ra lỗi.

Chúng ta cần một cách để nói với số lượng goroutine không biết và không giới hạn để
ngừng gửi các giá trị của chúng downstream. Trong Go, chúng ta có thể làm điều này bằng cách
đóng một channel, vì
[thao tác nhận trên một channel đã đóng luôn có thể tiến hành ngay lập tức, trả về giá trị zero của kiểu phần tử.](/ref/spec#Receive_operator)

Điều này có nghĩa là `main` có thể bỏ block tất cả các sender chỉ bằng cách đóng
channel `done`. Việc đóng này thực sự là một tín hiệu phát sóng cho
các sender. Chúng ta mở rộng _mỗi_ hàm pipeline của mình để chấp nhận
`done` như một tham số và sắp xếp cho việc đóng xảy ra qua câu lệnh
`defer`, để tất cả các đường trả về từ `main` sẽ báo hiệu
các giai đoạn pipeline thoát.

{{code "pipelines/sqdone3.go" `/func main/` `/^}/`}}

Mỗi giai đoạn pipeline của chúng ta bây giờ có thể trả về ngay khi `done` bị đóng.
Routine `output` trong `merge` có thể trả về mà không cần thoát hết channel inbound của nó,
vì nó biết sender upstream, `sq`, sẽ ngừng cố gắng gửi khi
`done` bị đóng. `output` đảm bảo `wg.Done` được gọi trên tất cả các đường trả về qua
câu lệnh `defer`:

{{code "pipelines/sqdone3.go" `/func merge/` `/unchanged/`}}

Tương tự, `sq` có thể trả về ngay khi `done` bị đóng. `sq` đảm bảo channel `out`
của nó bị đóng trên tất cả các đường trả về qua câu lệnh `defer`:

{{code "pipelines/sqdone3.go" `/func sq/` `/^}/`}}

Đây là các hướng dẫn cho việc xây dựng pipeline:

  - các giai đoạn đóng channel outbound khi tất cả các thao tác gửi đã hoàn thành.
  - các giai đoạn tiếp tục nhận giá trị từ các channel inbound cho đến khi các channel đó bị đóng hoặc các sender được bỏ block.

Các pipeline bỏ block các sender bằng cách đảm bảo có đủ buffer cho tất cả
các giá trị được gửi hoặc bằng cách báo hiệu tường minh cho các sender khi receiver có thể
từ bỏ channel.

## Tính toán digest của một cây

Hãy xem xét một pipeline thực tế hơn.

MD5 là thuật toán message-digest hữu ích như một file checksum. Tiện ích dòng lệnh
`md5sum` in các giá trị digest cho một danh sách các tệp.

	% md5sum *.go
	d47c2bbc28298ca9befdfbc5d3aa4e65  bounded.go
	ee869afd31f83cbb2d10ee81b2b831dc  parallel.go
	b88175e65fdcbc01ac08aaf1fd9b5e96  serial.go

Chương trình ví dụ của chúng ta giống như `md5sum` nhưng thay vào đó nhận một thư mục duy nhất làm
đối số và in các giá trị digest cho mỗi tệp thông thường trong thư mục đó,
được sắp xếp theo tên đường dẫn.

	% go run serial.go .
	d47c2bbc28298ca9befdfbc5d3aa4e65  bounded.go
	ee869afd31f83cbb2d10ee81b2b831dc  parallel.go
	b88175e65fdcbc01ac08aaf1fd9b5e96  serial.go

Hàm main của chương trình gọi một hàm helper `MD5All`, trả về một
map từ tên đường dẫn đến giá trị digest, sau đó sắp xếp và in kết quả:

{{code "pipelines/serial.go" `/func main/` `/^}/`}}

Hàm `MD5All` là tâm điểm của cuộc thảo luận. Trong
[serial.go](pipelines/serial.go), triển khai không sử dụng tính đồng thời và
chỉ đơn giản là đọc và tính tổng mỗi tệp khi nó duyệt cây.

{{code "pipelines/serial.go" `/MD5All/` `/^}/`}}

## Tính toán song song

Trong [parallel.go](pipelines/parallel.go), chúng ta chia `MD5All` thành pipeline hai giai đoạn.
Giai đoạn đầu tiên, `sumFiles`, duyệt cây, tính digest của mỗi tệp trong
một goroutine mới, và gửi kết quả trên một channel với kiểu giá trị `result`:

{{code "pipelines/parallel.go" `/type result/` `/}/` "HLresult"}}

`sumFiles` trả về hai channel: một cho `results` và một cho lỗi
được trả về bởi `filepath.Walk`. Hàm walk khởi động một goroutine mới để
xử lý mỗi tệp thông thường, sau đó kiểm tra `done`. Nếu `done` bị đóng, walk
dừng ngay lập tức:

{{code "pipelines/parallel.go" `/func sumFiles/` `/^}/`}}

`MD5All` nhận các giá trị digest từ `c`. `MD5All` trả về sớm khi có lỗi,
đóng `done` qua `defer`:

{{code "pipelines/parallel.go" `/func MD5All/` `/^}/` "HLdone"}}

## Song song có giới hạn

Triển khai `MD5All` trong [parallel.go](pipelines/parallel.go)
khởi động một goroutine mới cho mỗi tệp. Trong một thư mục với nhiều tệp lớn,
điều này có thể cấp phát nhiều bộ nhớ hơn mức có sẵn trên máy.

Chúng ta có thể giới hạn các cấp phát này bằng cách giới hạn số lượng tệp được đọc song song.
Trong [bounded.go](pipelines/bounded.go), chúng ta làm điều này bằng cách
tạo một số lượng cố định các goroutine để đọc các tệp. Pipeline của chúng ta
bây giờ có ba giai đoạn: duyệt cây, đọc và tính digest các tệp, và
thu thập các digest.

Giai đoạn đầu tiên, `walkFiles`, phát ra các đường dẫn của các tệp thông thường trong cây:

{{code "pipelines/bounded.go" `/func walkFiles/` `/^}/`}}

Giai đoạn giữa khởi động một số lượng cố định các goroutine `digester` nhận
tên tệp từ `paths` và gửi `results` trên channel `c`:

{{code "pipelines/bounded.go" `/func digester/` `/^}/` "HLpaths"}}

Không giống như các ví dụ trước của chúng ta, `digester` không đóng channel đầu ra của nó, vì
nhiều goroutine đang gửi trên một channel chia sẻ. Thay vào đó, code trong `MD5All`
sắp xếp để channel bị đóng khi tất cả các `digester` đã xong:

{{code "pipelines/bounded.go" `/fixed number/` `/End of pipeline/` "HLc"}}

Thay vào đó chúng ta có thể có mỗi digester tạo và trả về channel đầu ra riêng của nó,
nhưng sau đó chúng ta sẽ cần các goroutine bổ sung để fan-in các
kết quả.

Giai đoạn cuối cùng nhận tất cả `results` từ `c` rồi kiểm tra lỗi
từ `errc`. Kiểm tra này không thể xảy ra sớm hơn, vì trước
điểm này, `walkFiles` có thể bị block khi gửi các giá trị downstream:

{{code "pipelines/bounded.go" `/m := make/` `/^}/` "HLerrc"}}

## Kết luận

Bài viết này đã trình bày các kỹ thuật để xây dựng các pipeline dữ liệu phát trực tuyến
trong Go. Xử lý các thất bại trong các pipeline như vậy là khó khăn, vì mỗi giai đoạn trong
pipeline có thể bị block khi cố gắng gửi các giá trị downstream, và các giai đoạn downstream có thể
không còn quan tâm đến dữ liệu đến. Chúng ta đã chỉ ra cách đóng một channel
có thể phát sóng tín hiệu "done" đến tất cả các goroutine được khởi động bởi một
pipeline và xác định các hướng dẫn để xây dựng pipeline đúng cách.

Đọc thêm:

  - [Go Concurrency Patterns](/talks/2012/concurrency.slide#1)
    ([video](https://www.youtube.com/watch?v=f6kdp27TYZs)) trình bày các khái niệm cơ bản
    về các nguyên tố đồng thời của Go và một số cách áp dụng chúng.
  - [Advanced Go Concurrency Patterns](/blog/advanced-go-concurrency-patterns)
    ([video](http://www.youtube.com/watch?v=QDDwwePbDtw)) đề cập đến các cách sử dụng phức tạp hơn
    của các nguyên tố Go,
    đặc biệt là `select`.
  - Bài báo của Douglas McIlroy [Squinting at Power Series](https://swtch.com/~rsc/thread/squint.pdf)
    chỉ ra cách tính đồng thời giống Go cung cấp hỗ trợ thanh lịch cho các tính toán phức tạp.
