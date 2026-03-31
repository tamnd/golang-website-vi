---
title: Ghi chú phát hành Go 1.12
---

<!--
NOTE: In this document and others in this directory, the convention is to
set fixed-width phrases with non-fixed-width spaces, as in
`hello` `world`.
Do not send CLs removing the interior tags from such phrases.
-->

<style>
  main ul li { margin: 0.5em 0; }
</style>

## Giới thiệu về Go 1.12 {#introduction}

Bản phát hành Go mới nhất, phiên bản 1.12, ra đời sáu tháng sau [Go 1.11](go1.11).
Phần lớn các thay đổi nằm ở phần triển khai toolchain, runtime và thư viện.
Như thường lệ, bản phát hành này duy trì [cam kết tương thích](/doc/go1compat) của Go 1.
Chúng tôi kỳ vọng hầu hết các chương trình Go sẽ tiếp tục biên dịch và chạy như trước đây.

## Thay đổi về ngôn ngữ {#language}

Không có thay đổi nào trong đặc tả ngôn ngữ.

## Nền tảng {#ports}

<!-- CL 138675 -->
Trình phát hiện race (race detector) hiện đã được hỗ trợ trên `linux/arm64`.

Go 1.12 là bản phát hành cuối cùng được hỗ trợ trên FreeBSD 10.x, vốn đã hết
vòng đời (end-of-life). Go 1.13 sẽ yêu cầu FreeBSD 11.2+ hoặc FreeBSD 12.0+.
FreeBSD 12.0+ yêu cầu kernel có tùy chọn COMPAT\_FREEBSD11 được bật (đây là mặc định).

<!-- CL 146898 -->
cgo hiện đã được hỗ trợ trên `linux/ppc64`.

<!-- CL 146023 -->
`hurd` hiện là một giá trị được nhận diện cho `GOOS`, dành riêng
cho hệ thống GNU/Hurd khi sử dụng với `gccgo`.

### Windows {#windows}

Cổng `windows/arm` mới của Go hỗ trợ chạy Go trên Windows 10
IoT Core trên các chip ARM 32-bit như Raspberry Pi 3.

### AIX {#aix}

Go hiện hỗ trợ AIX 7.2 trở lên trên kiến trúc POWER8 (`aix/ppc64`). Liên kết ngoài (external linking), cgo, pprof và trình phát hiện race chưa được hỗ trợ.

### Darwin {#darwin}

Go 1.12 là bản phát hành cuối cùng chạy được trên macOS 10.10 Yosemite.
Go 1.13 sẽ yêu cầu macOS 10.11 El Capitan trở lên.

<!-- CL 141639 -->
`libSystem` hiện được sử dụng khi thực hiện syscall trên Darwin,
đảm bảo khả năng tương thích tiến với các phiên bản macOS và iOS trong tương lai. <!-- CL 153338 -->
Việc chuyển sang `libSystem` đã kích hoạt thêm các kiểm tra App Store
đối với việc sử dụng API riêng tư. Vì được coi là riêng tư,
`syscall.Getdirentries` hiện luôn thất bại với
`ENOSYS` trên iOS.
Ngoài ra, [`syscall.Setrlimit`](/pkg/syscall/#Setrlimit)
trả về `invalid` `argument` ở những nơi mà trước đây thành công.
Những hậu quả này không đặc thù riêng của Go và người dùng nên kỳ vọng
hành vi tương đương với triển khai của `libSystem` trong tương lai.

## Công cụ {#tools}

### `go tool vet` không còn được hỗ trợ {#vet}

Lệnh `go vet` đã được viết lại để làm nền tảng cho nhiều
công cụ phân tích mã nguồn khác nhau. Xem
gói [golang.org/x/tools/go/analysis](https://godoc.org/golang.org/x/tools/go/analysis)
để biết thêm chi tiết. Một tác dụng phụ là `go tool vet`
không còn được hỗ trợ nữa. Các công cụ bên ngoài sử dụng `go tool
  vet` phải được cập nhật để dùng `go
  vet`. Việc dùng `go vet` thay cho `go tool
  vet` sẽ hoạt động với tất cả các phiên bản Go được hỗ trợ.

Trong khuôn khổ thay đổi này, tùy chọn thử nghiệm `-shadow`
không còn khả dụng với `go vet`. Việc kiểm tra
biến shadowing có thể thực hiện qua:

	go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
	go vet -vettool=$(which shadow)


### Tour {#tour}

<!-- CL 152657 -->
Go tour không còn được bao gồm trong gói phân phối nhị phân chính. Để
chạy tour cục bộ, thay vì dùng `go` `tool` `tour`,
hãy cài đặt thủ công:

	go get -u golang.org/x/tour
	tour


### Yêu cầu build cache {#gocache}

[Build cache](/cmd/go/#hdr-Build_and_test_caching) hiện là
bắt buộc như một bước hướng tới việc loại bỏ
`$GOPATH/pkg`. Đặt biến môi trường
`GOCACHE=off` sẽ khiến các lệnh `go` ghi vào
cache bị lỗi.

### Gói chỉ nhị phân (binary-only packages) {#binary-only}

Go 1.12 là bản phát hành cuối cùng hỗ trợ các gói chỉ nhị phân.

### Cgo {#cgo}

Go 1.12 sẽ dịch kiểu C `EGLDisplay` sang kiểu Go `uintptr`.
Thay đổi này tương tự cách Go 1.10 và mới hơn xử lý các kiểu
CoreFoundation của Darwin và JNI của Java. Xem
[tài liệu cgo](/cmd/cgo/#hdr-Special_cases)
để biết thêm thông tin.

<!-- CL 152657 -->
Các tên C bị xáo trộn (mangled C names) không còn được chấp nhận trong các gói sử dụng Cgo. Hãy dùng
tên Cgo thay thế. Ví dụ: dùng tên cgo được ghi chép là `C.char`
thay vì tên bị xáo trộn `_Ctype_char` do cgo sinh ra.

### Module {#modules}

<!-- CL 148517 -->
Khi `GO111MODULE` được đặt thành `on`, lệnh `go`
hiện hỗ trợ các thao tác nhận biết module bên ngoài thư mục module,
miễn là các thao tác đó không cần phân giải đường dẫn import tương đối so với
thư mục hiện tại hoặc chỉnh sửa trực tiếp tệp `go.mod`.
Các lệnh như `go`&nbsp;`get`,
`go`&nbsp;`list`, và
`go`&nbsp;`mod`&nbsp;`download` hoạt động như thể trong một
module với yêu cầu ban đầu trống.
Trong chế độ này, `go`&nbsp;`env`&nbsp;`GOMOD` trả về
thiết bị null của hệ thống (`/dev/null` hoặc `NUL`).

<!-- CL 146382 -->
Các lệnh `go` tải xuống và giải nén module hiện an toàn để
gọi đồng thời.
Cache module (`GOPATH/pkg/mod`) phải nằm trên filesystem hỗ trợ
file locking.

<!-- CL 147282, 147281 -->
Chỉ thị `go` trong tệp `go.mod` hiện cho biết
phiên bản ngôn ngữ được sử dụng bởi các tệp trong module đó.
Nó sẽ được đặt thành bản phát hành hiện tại
(`go`&nbsp;`1.12`) nếu chưa có phiên bản nào.
Nếu chỉ thị `go` của một module chỉ định một phiên bản _mới hơn_
toolchain đang dùng, lệnh `go` sẽ thử build các gói bất kể,
và chỉ ghi nhận sự không khớp nếu quá trình build thất bại.

<!-- CL 147282, 147281 -->
Sự thay đổi này trong cách dùng chỉ thị `go` có nghĩa là nếu
bạn dùng Go 1.12 để build một module, qua đó ghi `go 1.12`
vào tệp `go.mod`, bạn sẽ gặp lỗi khi
cố gắng build module đó với Go 1.11 đến Go 1.11.3.
Go 1.11.4 trở lên sẽ hoạt động bình thường, cũng như các bản phát hành cũ hơn Go 1.11.
Nếu bạn phải dùng Go 1.11 đến 1.11.3, bạn có thể tránh vấn đề bằng cách
đặt phiên bản ngôn ngữ thành 1.11, dùng công cụ go của Go 1.12,
qua `go mod edit -go=1.11`.

<!-- CL 152739 -->
Khi không thể phân giải một import bằng các module đang hoạt động,
lệnh `go` sẽ thử dùng các module được đề cập trong các chỉ thị `replace`
của module chính trước khi tham khảo cache module
và các nguồn mạng thông thường.
Nếu tìm thấy một bản thay thế phù hợp nhưng chỉ thị `replace`
không chỉ định phiên bản, lệnh `go` sẽ dùng một pseudo-version
dẫn xuất từ `time.Time` bằng không (ví dụ
như `v0.0.0-00010101000000-000000000000`).

### Toolchain trình biên dịch {#compiler}

<!-- CL 134155, 134156 -->
Phân tích biến sống (live variable analysis) của trình biên dịch đã được cải thiện. Điều này có thể có nghĩa là
các finalizer sẽ được thực thi sớm hơn trong bản phát hành này so với các bản trước.
Nếu điều đó gây ra vấn đề, hãy xem xét việc thêm lệnh gọi
[`runtime.KeepAlive`](/pkg/runtime/#KeepAlive) phù hợp.

<!-- CL 147361 -->
Nhiều hàm hơn hiện đủ điều kiện được inline theo mặc định, bao gồm
các hàm chỉ đơn giản gọi một hàm khác.
Việc inline thêm này làm cho việc dùng
[`runtime.CallersFrames`](/pkg/runtime/#CallersFrames)
thay vì lặp trực tiếp qua kết quả của
[`runtime.Callers`](/pkg/runtime/#Callers) trở nên quan trọng hơn.

```
// Old code which no longer works correctly (it will miss inlined call frames).
var pcs [10]uintptr
n := runtime.Callers(1, pcs[:])
for _, pc := range pcs[:n] {
	f := runtime.FuncForPC(pc)
	if f != nil {
		fmt.Println(f.Name())
	}
}
```

```
// New code which will work correctly.
var pcs [10]uintptr
n := runtime.Callers(1, pcs[:])
frames := runtime.CallersFrames(pcs[:n])
for {
	frame, more := frames.Next()
	fmt.Println(frame.Function)
	if !more {
		break
	}
}
```

<!-- CL 153477 -->
Các wrapper do trình biên dịch tạo ra để triển khai method expression
không còn được báo cáo bởi [`runtime.CallersFrames`](/pkg/runtime/#CallersFrames)
và [`runtime.Stack`](/pkg/runtime/#Stack). Chúng
cũng không được in trong stack trace khi panic.
Thay đổi này giúp toolchain `gc` phù hợp với
toolchain `gccgo`, vốn đã bỏ qua các wrapper như vậy
khỏi stack trace.
Các client của các API này có thể cần điều chỉnh để thích nghi với các frame bị thiếu.
Đối với code cần hoạt động liên thông giữa các bản phát hành 1.11 và 1.12,
bạn có thể thay thế method expression `x.M`
bằng hàm literal ` func (...) { x.M(...) }  `.

<!-- CL 144340 -->
Trình biên dịch hiện chấp nhận cờ `-lang` để đặt phiên bản ngôn ngữ Go
cần sử dụng. Ví dụ: `-lang=go1.8` khiến trình biên dịch
phát ra lỗi nếu chương trình dùng type alias, được thêm vào trong Go 1.9.
Các thay đổi ngôn ngữ thực hiện trước Go 1.12 không được kiểm tra nhất quán.

<!-- CL 147160 -->
Toolchain trình biên dịch hiện sử dụng các quy ước khác nhau để gọi các hàm
Go và hàm assembly. Điều này sẽ vô hình với người dùng,
ngoại trừ các lệnh gọi đồng thời vượt qua biên giới Go-assembly
_và_ vượt qua ranh giới gói. Nếu quá trình liên kết dẫn đến
lỗi như "relocation target not defined for ABIInternal (but
is defined for ABI0)", hãy tham khảo
[phần tương thích](https://github.com/golang/proposal/blob/master/design/27539-internal-abi.md#compatibility)
của tài liệu thiết kế ABI.

<!-- CL 145179 -->
Đã có nhiều cải tiến về thông tin gỡ lỗi DWARF
do trình biên dịch tạo ra, bao gồm cải tiến về in tham số
và thông tin vị trí biến.

<!-- CL 61511 -->
Các chương trình Go hiện cũng duy trì stack frame pointer trên `linux/arm64`
để hỗ trợ các công cụ profiling như `perf`. Việc duy trì frame pointer
có chi phí runtime nhỏ, biến động nhưng trung bình khoảng 3%.
Để build toolchain không dùng frame pointer, đặt
`GOEXPERIMENT=noframepointer` khi chạy `make.bash`.

<!-- CL 142717 -->
Chế độ trình biên dịch "safe" đã lỗi thời (bật bằng gcflag `-u`) đã bị loại bỏ.

### `godoc` và `go` `doc` {#godoc}

Trong Go 1.12, `godoc` không còn có giao diện dòng lệnh và
chỉ là web server. Người dùng nên dùng `go` `doc`
để xuất trợ giúp dòng lệnh thay thế. Go 1.12 là bản phát hành cuối cùng bao gồm
web server `godoc`; trong Go 1.13 nó sẽ có sẵn
qua `go` `get`.

<!-- CL 141977 -->
`go` `doc` hiện hỗ trợ cờ `-all`,
khiến nó in ra tất cả các API được xuất khẩu và tài liệu của chúng,
như lệnh `godoc` trước đây từng làm.

<!-- CL 140959 -->
`go` `doc` cũng có thêm cờ `-src`,
sẽ hiển thị mã nguồn của đích.

### Trace {#trace}

<!-- CL 60790 -->
Công cụ trace hiện hỗ trợ vẽ đồ thị mutator utilization curves,
bao gồm tham chiếu chéo đến execution trace. Chúng hữu ích
để phân tích tác động của bộ gom rác đến độ trễ và
thông lượng của ứng dụng.

### Assembler {#assembler}

<!-- CL 147218 -->
Trên `arm64`, thanh ghi nền tảng được đổi tên từ
`R18` thành `R18_PLATFORM` để ngăn chặn việc sử dụng nhầm,
vì hệ điều hành có thể chọn dành riêng thanh ghi này.

## Runtime {#runtime}

<!-- CL 138959 -->
Go 1.12 cải thiện đáng kể hiệu suất của quá trình sweeping khi
một phần lớn heap vẫn còn sống. Điều này giảm độ trễ phân bổ
ngay sau một lần garbage collection.

<!-- CL 139719 -->
Go runtime hiện trả lại bộ nhớ cho hệ điều hành tích cực hơn,
đặc biệt để phản hồi các phân bổ lớn không thể tái sử dụng không gian heap hiện có.

<!-- CL 146342, CL 146340, CL 146345, CL 146339, CL 146343, CL 146337, CL 146341, CL 146338 -->
Mã timer và deadline của Go runtime nhanh hơn và mở rộng tốt hơn
với số lượng CPU cao hơn. Điều này cải thiện
hiệu suất của việc thao tác deadline kết nối mạng.

<!-- CL 135395 -->
Trên Linux, runtime hiện dùng `MADV_FREE` để giải phóng
bộ nhớ chưa sử dụng. Điều này hiệu quả hơn nhưng có thể dẫn đến RSS
được báo cáo cao hơn. Kernel sẽ thu hồi dữ liệu chưa sử dụng khi cần.
Để quay lại hành vi của Go 1.11 (`MADV_DONTNEED`), đặt
biến môi trường `GODEBUG=madvdontneed=1`.

<!-- CL 149578 -->
Thêm cpu._extension_=off vào biến môi trường
[GODEBUG](/doc/diagnostics.html#godebug)
hiện vô hiệu hóa việc sử dụng các phần mở rộng tập lệnh CPU tùy chọn
trong thư viện chuẩn và runtime. Điều này chưa được hỗ trợ trên Windows.

<!-- CL 158337 -->
Go 1.12 cải thiện độ chính xác của memory profile bằng cách sửa
lỗi đếm quá mức (overcounting) của các phân bổ heap lớn.

<!-- CL 159717 -->
Traceback, `runtime.Caller`,
và `runtime.Callers` không còn bao gồm
các hàm khởi tạo do trình biên dịch sinh ra. Thực hiện traceback
trong quá trình khởi tạo một biến toàn cục sẽ hiển thị
một hàm có tên `PKG.init.ializers`.

## Thư viện chuẩn {#library}

### TLS 1.3 {#tls_1_3}

Go 1.12 bổ sung hỗ trợ opt-in cho TLS 1.3 trong gói `crypto/tls` theo
đặc tả của [RFC 8446](https://www.rfc-editor.org/info/rfc8446). Tính năng này có thể
được bật bằng cách thêm giá trị `tls13=1` vào biến môi trường `GODEBUG`.
Nó sẽ được bật mặc định trong Go 1.13.

Để đàm phán TLS 1.3, hãy đảm bảo bạn không đặt `MaxVersion` rõ ràng trong
[`Config`](/pkg/crypto/tls/#Config) và chạy chương trình với
biến môi trường `GODEBUG=tls13=1` được đặt.

Tất cả các tính năng TLS 1.2 ngoại trừ `TLSUnique` trong
[`ConnectionState`](/pkg/crypto/tls/#ConnectionState)
và renegotiation đều có sẵn trong TLS 1.3 và cung cấp bảo mật cùng
hiệu suất tương đương hoặc tốt hơn. Lưu ý rằng dù TLS 1.3 tương thích ngược
với các phiên bản trước, một số hệ thống kế thừa nhất định có thể không hoạt động
đúng khi cố gắng đàm phán nó. Các khóa chứng chỉ RSA quá nhỏ
để an toàn (bao gồm khóa 512-bit) sẽ không hoạt động với TLS 1.3.

Các cipher suite của TLS 1.3 không thể cấu hình. Tất cả các cipher suite được hỗ trợ đều
an toàn, và nếu `PreferServerCipherSuites` được đặt trong
[`Config`](/pkg/crypto/tls/#Config), thứ tự ưu tiên
dựa trên phần cứng khả dụng.

Early data (còn gọi là "0-RTT mode") hiện không được hỗ trợ như là
client hoặc server. Ngoài ra, server Go 1.12 không hỗ trợ bỏ qua
early data bất ngờ nếu client gửi nó. Vì chế độ TLS 1.3 0-RTT
liên quan đến client giữ trạng thái về server nào hỗ trợ 0-RTT,
server Go 1.12 không thể là một phần của nhóm cân bằng tải (load-balancing pool) nơi một số server khác
hỗ trợ 0-RTT. Nếu chuyển một domain từ server hỗ trợ 0-RTT sang
server Go 1.12, 0-RTT phải được tắt trong ít nhất thời gian
tồn tại của session ticket đã cấp trước khi chuyển đổi để đảm bảo
hoạt động không bị gián đoạn.

Trong TLS 1.3, client là người nói cuối cùng trong quá trình handshake, vì vậy nếu nó gây ra
lỗi trên server, lỗi đó sẽ được trả về cho client qua lần gọi
[`Read`](/pkg/crypto/tls/#Conn.Read) đầu tiên, không phải qua
[`Handshake`](/pkg/crypto/tls/#Conn.Handshake). Ví dụ:
đây sẽ là trường hợp nếu server từ chối chứng chỉ client.
Tương tự, session ticket hiện là tin nhắn post-handshake, vì vậy chỉ
được client nhận tại lần gọi
[`Read`](/pkg/crypto/tls/#Conn.Read) đầu tiên của nó.

### Các thay đổi nhỏ trong thư viện {#minor_library_changes}

Như thường lệ, có nhiều thay đổi và cập nhật nhỏ đối với thư viện,
được thực hiện với [cam kết tương thích](/doc/go1compat) của Go 1
trong tâm trí.

<!-- TODO: CL 115677: https://golang.org/cl/115677: cmd/vet: check embedded field tags too -->

#### [bufio](/pkg/bufio/)

<!-- CL 149297 -->
Các phương thức [`UnreadRune`](/pkg/bufio/#Reader.UnreadRune) và
[`UnreadByte`](/pkg/bufio/#Reader.UnreadByte) của `Reader` hiện sẽ trả về lỗi
nếu chúng được gọi sau [`Peek`](/pkg/bufio/#Reader.Peek).

<!-- bufio -->

#### [bytes](/pkg/bytes/)

<!-- CL 137855 -->
Hàm mới [`ReplaceAll`](/pkg/bytes/#ReplaceAll) trả về bản sao của
một byte slice với tất cả các thể hiện không chồng lấp của một giá trị được thay thế bằng giá trị khác.

<!-- CL 145098 -->
Con trỏ đến [`Reader`](/pkg/bytes/#Reader) với giá trị bằng không hiện
tương đương về chức năng với [`NewReader`](/pkg/bytes/#NewReader)`(nil)`.
Trước Go 1.12, cái trước không thể dùng thay thế cho cái sau trong mọi trường hợp.

<!-- bytes -->

#### [crypto/rand](/pkg/crypto/rand/)

<!-- CL 139419 -->
Một cảnh báo sẽ được in ra standard error lần đầu tiên
`Reader.Read` bị chặn hơn 60 giây khi chờ
đọc entropy từ kernel.

<!-- CL 120055 -->
Trên FreeBSD, `Reader` hiện dùng lệnh gọi hệ thống `getrandom`
nếu có, ngược lại dùng `/dev/urandom`.

<!-- crypto/rand -->

#### [crypto/rc4](/pkg/crypto/rc4/)

<!-- CL 130397 -->
Bản phát hành này loại bỏ các triển khai assembly, chỉ giữ lại
phiên bản Go thuần túy. Trình biên dịch Go tạo ra code
tốt hơn hoặc tệ hơn một chút, tùy thuộc vào CPU cụ thể.
RC4 không an toàn và chỉ nên dùng để tương thích
với các hệ thống kế thừa.

<!-- crypto/rc4 -->

#### [crypto/tls](/pkg/crypto/tls/)

<!-- CL 143177 -->
Nếu client gửi tin nhắn ban đầu không trông giống TLS, server
sẽ không còn trả lời với một cảnh báo, và nó sẽ để lộ
`net.Conn` bên dưới trong trường mới `Conn` của
[`RecordHeaderError`](/pkg/crypto/tls/#RecordHeaderError).

<!-- crypto/tls -->

#### [database/sql](/pkg/database/sql/)

<!-- CL 145738 -->
Con trỏ truy vấn (query cursor) hiện có thể thu được bằng cách truyền một giá trị
[`*Rows`](/pkg/database/sql/#Rows)
vào phương thức [`Row.Scan`](/pkg/database/sql/#Row.Scan).

<!-- database/sql -->

#### [expvar](/pkg/expvar/)

<!-- CL 139537 -->
Phương thức mới [`Delete`](/pkg/expvar/#Map.Delete) cho phép
xóa các cặp key/value khỏi [`Map`](/pkg/expvar/#Map).

<!-- expvar -->

#### [fmt](/pkg/fmt/)

<!-- CL 142737 -->
Map hiện được in theo thứ tự sắp xếp key để dễ kiểm thử. Các quy tắc sắp xếp là:

  - Khi áp dụng, nil so sánh thấp hơn
  - int, float và string sắp xếp theo <
  - NaN so sánh nhỏ hơn float không phải NaN
  - bool so sánh false trước true
  - Complex so sánh phần thực, sau đó phần ảo
  - Con trỏ so sánh theo địa chỉ máy
  - Giá trị channel so sánh theo địa chỉ máy
  - Struct so sánh từng trường lần lượt
  - Array so sánh từng phần tử lần lượt
  - Giá trị interface so sánh trước tiên theo `reflect.Type` mô tả kiểu cụ thể
    và sau đó theo giá trị cụ thể như được mô tả trong các quy tắc trước.


<!-- CL 129777 -->
Khi in map, các giá trị key không phản xạ (non-reflexive) như `NaN` trước đây
được hiển thị là `<nil>`. Từ bản phát hành này, các giá trị đúng được in.

<!-- fmt -->

#### [go/doc](/pkg/go/doc/)

<!-- CL 140958 -->
Để giải quyết một số vấn đề tồn đọng trong [`cmd/doc`](/cmd/doc/),
gói này có bit [`Mode`](/pkg/go/doc/#Mode) mới,
`PreserveAST`, kiểm soát liệu dữ liệu AST có bị xóa không.

<!-- go/doc -->

#### [go/token](/pkg/go/token/)

<!-- CL 134075 -->
Kiểu [`File`](/pkg/go/token#File) có trường mới
[`LineStart`](/pkg/go/token#File.LineStart),
trả về vị trí bắt đầu của một dòng cho trước. Điều này đặc biệt hữu ích
trong các chương trình đôi khi xử lý các tệp không phải Go, như assembly, nhưng muốn dùng
cơ chế `token.Pos` để xác định vị trí tệp.

<!-- go/token -->

#### [image](/pkg/image/)

<!-- CL 118755 -->
Hàm [`RegisterFormat`](/pkg/image/#RegisterFormat) hiện an toàn để sử dụng đồng thời.

<!-- image -->

#### [image/png](/pkg/image/png/)

<!-- CL 134235 -->
Ảnh Paletted với ít hơn 16 màu hiện mã hóa thành output nhỏ hơn.

<!-- image/png -->

#### [io](/pkg/io/)

<!-- CL 139457 -->
Interface mới [`StringWriter`](/pkg/io#StringWriter) bao bọc
hàm [`WriteString`](/pkg/io/#WriteString).

<!-- io -->

#### [math](/pkg/math/)

<!-- CL 153059 -->
Các hàm
[`Sin`](/pkg/math/#Sin),
[`Cos`](/pkg/math/#Cos),
[`Tan`](/pkg/math/#Tan),
và [`Sincos`](/pkg/math/#Sincos) hiện
áp dụng Payne-Hanek range reduction cho các đối số rất lớn. Điều này
tạo ra kết quả chính xác hơn, nhưng chúng sẽ không giống bit-cho-bit
với các kết quả trong các bản phát hành trước.

<!-- math -->

#### [math/bits](/pkg/math/bits/)

<!-- CL 123157 -->
Các phép toán độ chính xác mở rộng mới [`Add`](/pkg/math/bits/#Add), [`Sub`](/pkg/math/bits/#Sub), [`Mul`](/pkg/math/bits/#Mul), và [`Div`](/pkg/math/bits/#Div) có sẵn trong các phiên bản `uint`, `uint32` và `uint64`.

<!-- math/bits -->

#### [net](/pkg/net/)

<!-- CL 146659 -->
Cài đặt
[`Dialer.DualStack`](/pkg/net/#Dialer.DualStack) hiện bị bỏ qua và không còn được dùng (deprecated);
RFC 6555 Fast Fallback ("Happy Eyeballs") hiện được bật mặc định. Để tắt, đặt
[`Dialer.FallbackDelay`](/pkg/net/#Dialer.FallbackDelay) thành một giá trị âm.

<!-- CL 107196 -->
Tương tự, TCP keep-alive hiện được bật mặc định nếu
[`Dialer.KeepAlive`](/pkg/net/#Dialer.KeepAlive) bằng không.
Để tắt, đặt nó thành một giá trị âm.

<!-- CL 113997 -->
Trên Linux, [system call `splice`](https://man7.org/linux/man-pages/man2/splice.2.html) hiện được sử dụng khi sao chép từ
[`UnixConn`](/pkg/net/#UnixConn) sang
[`TCPConn`](/pkg/net/#TCPConn).

<!-- net -->

#### [net/http](/pkg/net/http/)

<!-- CL 143177 -->
HTTP server hiện từ chối các yêu cầu HTTP bị định hướng sai (misdirected) đến server HTTPS với phản hồi plaintext "400 Bad Request".

<!-- CL 130115 -->
Phương thức mới [`Client.CloseIdleConnections`](/pkg/net/http/#Client.CloseIdleConnections)
gọi `CloseIdleConnections` của `Transport` bên dưới của `Client`
nếu nó có.

<!-- CL 145398 -->
[`Transport`](/pkg/net/http/#Transport) không còn từ chối các phản hồi HTTP khai báo
HTTP Trailer nhưng không dùng chunked encoding. Thay vào đó, các trailer được khai báo hiện chỉ đơn giản bị bỏ qua.

<!-- CL 152080 -->
<!-- CL 151857 -->
[`Transport`](/pkg/net/http/#Transport) không còn xử lý các giá trị `MAX_CONCURRENT_STREAMS`
được quảng bá từ server HTTP/2 một cách nghiêm ngặt như trong Go 1.10 và Go 1.11. Hành vi mặc định hiện trở lại
như trong Go 1.9: mỗi kết nối đến một server có thể có tối đa `MAX_CONCURRENT_STREAMS` yêu cầu
đang hoạt động và sau đó các kết nối TCP mới được tạo khi cần. Trong Go 1.10 và Go 1.11, gói `http2`
sẽ chặn và chờ các yêu cầu hoàn thành thay vì tạo kết nối mới.
Để lấy lại hành vi nghiêm ngặt hơn, hãy import trực tiếp gói
[`golang.org/x/net/http2`](https://godoc.org/golang.org/x/net/http2)
và đặt
[`Transport.StrictMaxConcurrentStreams`](https://godoc.org/golang.org/x/net/http2#Transport.StrictMaxConcurrentStreams) thành
`true`.

<!-- net/http -->

#### [net/url](/pkg/net/url/)

<!-- CL 159157, CL 160178 -->
[`Parse`](/pkg/net/url/#Parse),
[`ParseRequestURI`](/pkg/net/url/#ParseRequestURI),
và
[`URL.Parse`](/pkg/net/url/#URL.Parse)
hiện trả về lỗi cho các URL chứa ký tự điều khiển ASCII, bao gồm NULL,
tab và xuống dòng.

<!-- net/url -->

#### [net/http/httputil](/pkg/net/http/httputil/)

<!-- CL 146437 -->
[`ReverseProxy`](/pkg/net/http/httputil/#ReverseProxy) hiện tự động
proxy các yêu cầu WebSocket.

<!-- net/http/httputil -->

#### [os](/pkg/os/)

<!-- CL 125443 -->
Phương thức mới [`ProcessState.ExitCode`](/pkg/os/#ProcessState.ExitCode)
trả về mã thoát của tiến trình.

<!-- CL 135075 -->
`ModeCharDevice` đã được thêm vào bitmask `ModeType`, cho phép
`ModeDevice | ModeCharDevice` được khôi phục khi che (masking) một
[`FileMode`](/pkg/os/#FileMode) với `ModeType`.

<!-- CL 139418 -->
Hàm mới [`UserHomeDir`](/pkg/os/#UserHomeDir) trả về
thư mục home của người dùng hiện tại.

<!-- CL 146020 -->
[`RemoveAll`](/pkg/os/#RemoveAll) hiện hỗ trợ các đường dẫn dài hơn 4096 ký tự
trên hầu hết các hệ thống Unix.

<!-- CL 130676 -->
[`File.Sync`](/pkg/os/#File.Sync) hiện sử dụng `F_FULLFSYNC` trên macOS
để xả đúng nội dung tệp vào bộ nhớ vĩnh cửu.
Điều này có thể khiến phương thức chạy chậm hơn so với các bản phát hành trước.

<!--CL 155517 -->
[`File`](/pkg/os/#File) hiện hỗ trợ
phương thức [`SyscallConn`](/pkg/os/#File.SyscallConn)
trả về giá trị interface
[`syscall.RawConn`](/pkg/syscall/#RawConn).
Điều này có thể được dùng để gọi các thao tác đặc thù hệ thống
trên file descriptor bên dưới.

<!-- os -->

#### [path/filepath](/pkg/path/filepath/)

<!-- CL 145220 -->
Hàm [`IsAbs`](/pkg/path/filepath/#IsAbs) hiện trả về true khi được truyền
một tên tệp dành riêng trên Windows như `NUL`.
[Danh sách các tên dành riêng.](https://docs.microsoft.com/en-us/windows/desktop/fileio/naming-a-file#naming-conventions)

<!-- path/filepath -->

#### [reflect](/pkg/reflect/)

<!-- CL 33572 -->
Kiểu [`MapIter`](/pkg/reflect#MapIter) mới là
một iterator để duyệt qua một map. Kiểu này được truy cập qua
phương thức mới [`MapRange`](/pkg/reflect#Value.MapRange) của kiểu
[`Value`](/pkg/reflect#Value).
Nó tuân theo cùng ngữ nghĩa lặp như một câu lệnh range, với `Next`
để tiến iterator, và `Key`/`Value` để truy cập từng mục.

<!-- reflect -->

#### [regexp](/pkg/regexp/)

<!-- CL 139784 -->
[`Copy`](/pkg/regexp/#Regexp.Copy) không còn cần thiết
để tránh tranh chấp khóa (lock contention), vì vậy nó đã được đánh dấu không còn được khuyến nghị một phần.
[`Copy`](/pkg/regexp/#Regexp.Copy)
vẫn có thể phù hợp nếu lý do sử dụng là để tạo hai bản sao với
cài đặt [`Longest`](/pkg/regexp/#Regexp.Longest) khác nhau.

<!-- regexp -->

#### [runtime/debug](/pkg/runtime/debug/)

<!-- CL 144220 -->
Kiểu [`BuildInfo`](/pkg/runtime/debug/#BuildInfo) mới
để lộ thông tin build đọc từ binary đang chạy, chỉ có sẵn trong
các binary được build với hỗ trợ module. Điều này bao gồm đường dẫn gói chính,
thông tin module chính và các phụ thuộc module. Kiểu này được cung cấp qua
hàm [`ReadBuildInfo`](/pkg/runtime/debug/#ReadBuildInfo)
trên [`BuildInfo`](/pkg/runtime/debug/#BuildInfo).

<!-- runtime/debug -->

#### [strings](/pkg/strings/)

<!-- CL 137855 -->
Hàm mới [`ReplaceAll`](/pkg/strings/#ReplaceAll) trả về bản sao của
một chuỗi với tất cả các thể hiện không chồng lấp của một giá trị được thay thế bằng giá trị khác.

<!-- CL 145098 -->
Con trỏ đến [`Reader`](/pkg/strings/#Reader) với giá trị bằng không hiện
tương đương về chức năng với [`NewReader`](/pkg/strings/#NewReader)`(nil)`.
Trước Go 1.12, cái trước không thể dùng thay thế cho cái sau trong mọi trường hợp.

<!-- CL 122835 -->
Phương thức mới [`Builder.Cap`](/pkg/strings/#Builder.Cap) trả về dung lượng của byte slice bên dưới builder.

<!-- CL 131495 -->
Các hàm ánh xạ ký tự [`Map`](/pkg/strings/#Map),
[`Title`](/pkg/strings/#Title),
[`ToLower`](/pkg/strings/#ToLower),
[`ToLowerSpecial`](/pkg/strings/#ToLowerSpecial),
[`ToTitle`](/pkg/strings/#ToTitle),
[`ToTitleSpecial`](/pkg/strings/#ToTitleSpecial),
[`ToUpper`](/pkg/strings/#ToUpper), và
[`ToUpperSpecial`](/pkg/strings/#ToUpperSpecial)
hiện luôn đảm bảo trả về UTF-8 hợp lệ. Trong các bản phát hành trước, nếu đầu vào là UTF-8 không hợp lệ nhưng không cần thay thế ký tự nào,
các hàm này trả về UTF-8 không hợp lệ đó mà không thay đổi.

<!-- strings -->

#### [syscall](/pkg/syscall/)

<!-- CL 138595 -->
Inode 64-bit hiện được hỗ trợ trên FreeBSD 12. Một số kiểu đã được điều chỉnh theo.

<!-- CL 125456 -->
Họ địa chỉ Unix socket
([`AF_UNIX`](https://blogs.msdn.microsoft.com/commandline/2017/12/19/af_unix-comes-to-windows/))
hiện được hỗ trợ cho các phiên bản Windows tương thích.

<!-- CL 147117 -->
Hàm mới [`Syscall18`](/pkg/syscall/?GOOS=windows&GOARCH=amd64#Syscall18)
đã được giới thiệu cho Windows, cho phép gọi với tối đa 18 đối số.

<!-- syscall -->

#### [syscall/js](/pkg/syscall/js/)

<!-- CL 153559 -->

Kiểu `Callback` và hàm `NewCallback` đã được đổi tên;
chúng hiện được gọi là
[`Func`](/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Func) và
[`FuncOf`](/pkg/syscall/js/?GOOS=js&GOARCH=wasm#FuncOf).
Đây là một thay đổi phá vỡ, nhưng hỗ trợ WebAssembly vẫn là thử nghiệm
và chưa tuân theo
[cam kết tương thích Go 1](/doc/go1compat). Bất kỳ code nào dùng
tên cũ cần được cập nhật.

<!-- CL 141644 -->
Nếu một kiểu triển khai interface mới
[`Wrapper`](/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Wrapper),
[`ValueOf`](/pkg/syscall/js/?GOOS=js&GOARCH=wasm#ValueOf)
sẽ dùng nó để trả về giá trị JavaScript cho kiểu đó.

<!-- CL 143137 -->
Ý nghĩa của
[`Value`](/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Value) bằng không
đã thay đổi. Nó hiện đại diện cho giá trị JavaScript `undefined`
thay vì số không.
Đây là một thay đổi phá vỡ, nhưng hỗ trợ WebAssembly vẫn là thử nghiệm
và chưa tuân theo
[cam kết tương thích Go 1](/doc/go1compat). Bất kỳ code nào phụ thuộc vào
[`Value`](/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Value) bằng không
để có nghĩa là số không cần được cập nhật.

<!-- CL 144384 -->
Phương thức mới
[`Value.Truthy`](/pkg/syscall/js/?GOOS=js&GOARCH=wasm#Value.Truthy)
báo cáo
["tính truthy" JavaScript](https://developer.mozilla.org/en-US/docs/Glossary/Truthy)
của một giá trị cho trước.

<!-- syscall/js -->

#### [testing](/pkg/testing/)

<!-- CL 139258 -->
Cờ [`-benchtime`](/cmd/go/#hdr-Testing_flags) hiện hỗ trợ đặt số lần lặp rõ ràng thay vì thời gian khi giá trị kết thúc bằng "`x`". Ví dụ: `-benchtime=100x` chạy benchmark 100 lần.

<!-- testing -->

#### [text/template](/pkg/text/template/)

<!-- CL 142217 -->
Khi thực thi template, các giá trị context dài không còn bị cắt ngắn trong thông báo lỗi.

`executing "tmpl" at <.very.deep.context.v...>: map has no entry for key "notpresent"`

hiện là

`executing "tmpl" at <.very.deep.context.value.notpresent>: map has no entry for key "notpresent"`

<!-- CL 143097 -->
Nếu một hàm do người dùng định nghĩa được gọi bởi template bị panic, panic đó
hiện được bắt và trả về dưới dạng lỗi bởi
phương thức `Execute` hoặc `ExecuteTemplate`.

<!-- text/template -->

#### [time](/pkg/time/)

<!-- CL 151299 -->
Cơ sở dữ liệu múi giờ trong `$GOROOT/lib/time/zoneinfo.zip`
đã được cập nhật lên phiên bản 2018i. Lưu ý rằng tệp ZIP này chỉ
được dùng nếu hệ điều hành không cung cấp cơ sở dữ liệu múi giờ.

<!-- time -->

#### [unsafe](/pkg/unsafe/)

<!-- CL 146058 -->
Việc chuyển đổi một `unsafe.Pointer` nil sang `uintptr` và ngược lại với số học là không hợp lệ.
(Điều này đã không hợp lệ trước đó, nhưng hiện sẽ khiến trình biên dịch hoạt động sai.)

<!-- unsafe -->
