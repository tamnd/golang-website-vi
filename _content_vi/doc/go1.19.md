---
path: /doc/go1.19
title: Ghi chú phát hành Go 1.19
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

## Giới thiệu về Go 1.19 {#introduction}

Bản phát hành Go mới nhất, phiên bản 1.19, ra đời năm tháng sau [Go 1.18](/doc/go1.18).
Phần lớn các thay đổi nằm ở phần triển khai toolchain, runtime và thư viện.
Như thường lệ, bản phát hành này duy trì [cam kết tương thích](/doc/go1compat) của Go 1.
Chúng tôi kỳ vọng hầu hết các chương trình Go sẽ tiếp tục biên dịch và chạy như trước đây.

## Thay đổi về ngôn ngữ {#language}

<!-- https://go.dev/issue/52038 -->
Chỉ có một thay đổi nhỏ về ngôn ngữ,
một [sửa chữa rất nhỏ](/issue/52038)
đối với [phạm vi của type parameter trong khai báo phương thức](/ref/spec#Declarations_and_scope).
Các chương trình hiện tại không bị ảnh hưởng.

## Mô hình bộ nhớ {#mem}

<!-- https://go.dev/issue/50859 -->
[Mô hình bộ nhớ Go](/ref/mem) đã được
[xem xét lại](https://research.swtch.com/gomm) để đồng bộ Go với
mô hình bộ nhớ được dùng bởi C, C++, Java, JavaScript, Rust và Swift.
Go chỉ cung cấp các atomic tuần tự nhất quán (sequentially consistent), không có bất kỳ dạng nới lỏng nào như trong các ngôn ngữ khác.
Cùng với việc cập nhật mô hình bộ nhớ,
Go 1.19 giới thiệu [các kiểu mới trong gói `sync/atomic`](#atomic_types)
giúp sử dụng các giá trị atomic dễ dàng hơn, chẳng hạn như
[atomic.Int64](/pkg/sync/atomic/#Int64)
và
[atomic.Pointer[T]](/pkg/sync/atomic/#Pointer).

## Nền tảng {#ports}

### LoongArch 64-bit {#loong64}

<!-- https://go.dev/issue/46229 -->
Go 1.19 thêm hỗ trợ cho kiến trúc 64-bit Loongson
[LoongArch](https://loongson.github.io/LoongArch-Documentation)
trên Linux (`GOOS=linux`, `GOARCH=loong64`).
ABI được triển khai là LP64D. Phiên bản kernel tối thiểu được hỗ trợ là 5.19.

Lưu ý rằng hầu hết các bản phân phối Linux thương mại hiện có cho LoongArch đi kèm
với kernel cũ hơn, với ABI lệnh gọi hệ thống không tương thích trong lịch sử.
Các binary đã biên dịch sẽ không chạy trên các hệ thống này, ngay cả khi được liên kết tĩnh.
Người dùng trên các hệ thống không được hỗ trợ như vậy chỉ có thể dùng
gói Go được cung cấp bởi bản phân phối.

### RISC-V {#riscv64}

<!-- CL 402374 -->
Cổng `riscv64` hiện hỗ trợ truyền đối số hàm
và kết quả bằng thanh ghi. Benchmark cho thấy cải thiện hiệu suất điển hình
từ 10% trở lên trên `riscv64`.

## Công cụ {#tools}

### Nhận xét tài liệu (Doc Comments) {#go-doc}

<!-- https://go.dev/issue/51082 -->
<!-- CL 384265, CL 397276, CL 397278, CL 397279, CL 397281, CL 397284 -->
Go 1.19 thêm hỗ trợ cho link, danh sách và tiêu đề rõ ràng hơn trong comment tài liệu.
Là một phần của thay đổi này, [`gofmt`](/cmd/gofmt)
hiện định dạng lại comment tài liệu để làm rõ hơn ý nghĩa được render.
Xem "[Go Doc Comments](/doc/comment)"
để biết chi tiết cú pháp và mô tả các lỗi phổ biến hiện được `gofmt` làm nổi bật.
Cũng là một phần của thay đổi này, gói mới [go/doc/comment](/pkg/go/doc/comment/)
cung cấp khả năng phân tích và định dạng lại comment tài liệu
cũng như hỗ trợ render chúng thành HTML, Markdown và text.

### Ràng buộc build `unix` mới {#go-unix}

<!-- CL 389934 -->
<!-- https://go.dev/issue/20322 -->
<!-- https://go.dev/issue/51572 -->
Ràng buộc build `unix` hiện được nhận diện
trong các dòng `//go:build`. Ràng buộc được thỏa mãn
nếu hệ điều hành đích, còn gọi là `GOOS`, là
một hệ thống Unix hoặc giống Unix. Với bản phát hành 1.19, nó được thỏa mãn
nếu `GOOS` là một trong
`aix`, `android`, `darwin`,
`dragonfly`, `freebsd`, `hurd`,
`illumos`, `ios`, `linux`,
`netbsd`, `openbsd`, hoặc `solaris`.
Trong các bản phát hành tương lai, ràng buộc `unix` có thể khớp
với các hệ điều hành được hỗ trợ mới.

### Lệnh Go {#go-command}

<!-- https://go.dev/issue/51461 -->

Cờ `-trimpath`, nếu được đặt, hiện được đưa vào cài đặt build
được đóng dấu vào các binary Go bởi `go` `build`, và có thể
được kiểm tra bằng cách dùng
[`go` `version` `-m`](https://pkg.go.dev/cmd/go#hdr-Print_Go_version)
hoặc [`debug.ReadBuildInfo`](https://pkg.go.dev/runtime/debug#ReadBuildInfo).

`go` `generate` hiện đặt biến môi trường `GOROOT`
một cách rõ ràng trong môi trường của generator, để
các generator có thể xác định `GOROOT` đúng ngay cả khi được build
với `-trimpath`.

<!-- CL 404134 -->
`go` `test` và `go` `generate` hiện đặt
`GOROOT/bin` ở đầu `PATH` được dùng cho
tiến trình con, để các kiểm thử và generator thực thi lệnh `go`
sẽ phân giải nó đến cùng `GOROOT`.

<!-- CL 398058 -->
`go` `env` hiện trích dẫn các mục chứa dấu cách trong
các biến `CGO_CFLAGS`, `CGO_CPPFLAGS`, `CGO_CXXFLAGS`, `CGO_FFLAGS`, `CGO_LDFLAGS`,
và `GOGCCFLAGS` mà nó báo cáo.

<!-- https://go.dev/issue/29666 -->
`go` `list` `-json` hiện chấp nhận một
danh sách các trường JSON được phân tách bằng dấu phẩy để điền vào. Nếu có danh sách,
output JSON sẽ chỉ bao gồm các trường đó, và
`go` `list` có thể tránh tính toán các trường không
được bao gồm. Trong một số trường hợp, điều này có thể ngăn chặn các lỗi mà
nếu không sẽ được báo cáo.

<!-- CL 410821 -->
Lệnh `go` hiện cache thông tin cần thiết để tải một số module,
điều này sẽ giúp tăng tốc một số lần gọi `go` `list`.

### Vet {#vet}

<!-- https://go.dev/issue/47528 -->
Trình kiểm tra vet "errorsas" hiện báo cáo khi
[`errors.As`](/pkg/errors/#As) được gọi
với đối số thứ hai kiểu `*error`,
một lỗi phổ biến.

## Runtime {#runtime}

<!-- https://go.dev/issue/48409 -->
<!-- CL 397018 -->
Runtime hiện bao gồm hỗ trợ cho giới hạn bộ nhớ mềm (soft memory limit). Giới hạn bộ nhớ này
bao gồm heap Go và tất cả bộ nhớ khác được runtime quản lý, và
loại trừ các nguồn bộ nhớ bên ngoài như mapping của chính binary,
bộ nhớ được quản lý bằng ngôn ngữ khác và bộ nhớ được hệ điều hành giữ
thay mặt chương trình Go. Giới hạn này có thể được quản lý qua
[`runtime/debug.SetMemoryLimit`](/pkg/runtime/debug/#SetMemoryLimit)
hoặc biến môi trường tương đương
[`GOMEMLIMIT`](/pkg/runtime/#hdr-Environment_Variables).
Giới hạn hoạt động kết hợp với
[`runtime/debug.SetGCPercent`](/pkg/runtime/debug/#SetGCPercent)
/ [`GOGC`](/pkg/runtime/#hdr-Environment_Variables),
và sẽ được tuân theo ngay cả khi `GOGC=off`, cho phép các chương trình Go
luôn sử dụng tối đa giới hạn bộ nhớ của chúng, cải thiện hiệu quả tài nguyên
trong một số trường hợp. Xem [hướng dẫn GC](/doc/gc-guide) để có
hướng dẫn chi tiết giải thích giới hạn bộ nhớ mềm, cũng như
nhiều trường hợp sử dụng và kịch bản phổ biến. Lưu ý rằng các giới hạn bộ nhớ nhỏ,
khoảng vài chục megabyte trở xuống, ít có khả năng được tuân theo
do các yếu tố độ trễ bên ngoài, chẳng hạn như lên lịch OS. Xem
[issue 52433](/issue/52433) để biết thêm chi tiết. Các
giới hạn bộ nhớ lớn hơn, khoảng hàng trăm megabyte trở lên, ổn định và
sẵn sàng cho môi trường sản xuất.

<!-- CL 353989 -->
Để hạn chế tác động của GC thrashing khi kích thước heap sống của chương trình
tiến gần đến giới hạn bộ nhớ mềm, Go runtime cũng cố gắng giới hạn
tổng mức sử dụng CPU của GC ở mức 50%, không tính thời gian idle, chọn dùng nhiều bộ nhớ hơn
để ngăn chặn ứng dụng bị chậm lại. Trong thực tế, chúng tôi kỳ vọng giới hạn này
chỉ đóng vai trò trong các trường hợp ngoại lệ, và
[runtime metric](/pkg/runtime/metrics/#hdr-Supported_metrics) mới
`/gc/limiter/last-enabled:gc-cycle` báo cáo khi điều này lần cuối
xảy ra.

<!-- https://go.dev/issue/44163 -->
Runtime hiện lên lịch ít goroutine worker GC hơn trên các luồng hệ điều hành
ở trạng thái idle khi ứng dụng đủ idle để buộc một chu kỳ GC định kỳ.

<!-- https://go.dev/issue/18138 -->
<!-- CL 345889 -->
Runtime hiện sẽ phân bổ stack goroutine ban đầu dựa trên
mức sử dụng stack trung bình lịch sử của goroutine. Điều này tránh một số
tăng trưởng stack sớm và sao chép cần thiết trong trường hợp trung bình, đổi lại
tối đa 2x không gian lãng phí trên các goroutine sử dụng dưới mức trung bình.

<!-- https://go.dev/issue/46279 -->
<!-- CL 393354 -->
<!-- CL 392415 -->
Trên các hệ điều hành Unix, các chương trình Go import gói
[os](/pkg/os/) hiện tự động tăng giới hạn file mở
(`RLIMIT_NOFILE`) lên giá trị tối đa cho phép;
tức là chúng thay đổi giới hạn mềm để khớp với giới hạn cứng.
Điều này sửa các giới hạn thấp giả tạo được đặt trên một số hệ thống để tương thích với các chương trình C rất cũ
sử dụng lệnh gọi hệ thống [_select_](https://en.wikipedia.org/wiki/Select_(Unix)).
Các chương trình Go không được hưởng lợi từ giới hạn đó, và thậm chí các chương trình đơn giản như `gofmt`
thường hết file descriptor trên các hệ thống như vậy khi xử lý nhiều tệp song song.
Một tác động của thay đổi này là các chương trình Go lần lượt thực thi các chương trình C rất cũ trong tiến trình con
có thể chạy các chương trình đó với giới hạn quá cao.
Điều này có thể được sửa bằng cách đặt giới hạn cứng trước khi gọi chương trình Go.

<!-- https://go.dev/issue/51485 -->
<!-- CL 390421 -->
Các lỗi fatal không thể khôi phục (như ghi map đồng thời, hoặc unlock mutex
chưa bị lock) hiện in traceback đơn giản hơn, loại trừ metadata runtime
(tương đương với fatal panic) trừ khi `GOTRACEBACK=system` hoặc
`crash`. Traceback lỗi fatal nội bộ runtime luôn bao gồm
đầy đủ metadata bất kể giá trị của `GOTRACEBACK`.

<!-- https://go.dev/issue/50614 -->
<!-- CL 395754 -->
Hỗ trợ cho các lệnh gọi hàm do debugger inject đã được thêm vào trên ARM64,
cho phép người dùng gọi các hàm từ binary của họ trong một phiên
gỡ lỗi tương tác khi sử dụng debugger được cập nhật để tận dụng
chức năng này.

<!-- https://go.dev/issue/44853 -->
[Hỗ trợ address sanitizer được thêm vào trong Go 1.18](/doc/go1.18#go-build-asan)
hiện xử lý đối số hàm và biến toàn cục chính xác hơn.

## Trình biên dịch {#compiler}

<!-- https://go.dev/issue/5496 -->
<!-- CL 357330, 395714, 403979 -->
Trình biên dịch hiện sử dụng
[bảng nhảy (jump table)](https://en.wikipedia.org/wiki/Branch_table) để triển khai các câu lệnh switch số nguyên và chuỗi lớn.
Cải thiện hiệu suất cho câu lệnh switch thay đổi nhưng có thể
nhanh hơn khoảng 20%.
(Chỉ `GOARCH=amd64` và `GOARCH=arm64`)

<!-- CL 391014 -->
Trình biên dịch Go hiện yêu cầu cờ `-p=importpath` để
build một object file có thể liên kết. Điều này đã được cung cấp bởi
lệnh `go` và Bazel. Bất kỳ hệ thống build nào khác
gọi trình biên dịch Go trực tiếp sẽ cần đảm bảo chúng
truyền cờ này.

<!-- CL 415235 -->
Trình biên dịch Go không còn chấp nhận cờ `-importmap`.
Các hệ thống build gọi trực tiếp trình biên dịch Go phải dùng
cờ `-importcfg` thay thế.

## Assembler {#assembler}

<!-- CL 404298 -->
Giống như trình biên dịch, assembler hiện yêu cầu
cờ `-p=importpath` để build một object file có thể liên kết.
Điều này đã được cung cấp bởi lệnh `go`. Bất kỳ hệ thống build nào khác
gọi trực tiếp assembler Go sẽ cần đảm bảo chúng
truyền cờ này.

## Linker {#linker}

<!-- https://go.dev/issue/50796, CL 380755 -->
Trên các nền tảng ELF, linker hiện phát ra các phần DWARF nén theo
định dạng gABI tiêu chuẩn (`SHF_COMPRESSED`), thay vì
định dạng kế thừa `.zdebug`.

## Thư viện chuẩn {#library}

### Các kiểu atomic mới {#atomic_types}

<!-- https://go.dev/issue/50860 -->
<!-- CL 381317 -->
Gói [`sync/atomic`](/pkg/sync/atomic/) định nghĩa các kiểu atomic mới
[`Bool`](/pkg/sync/atomic/#Bool),
[`Int32`](/pkg/sync/atomic/#Int32),
[`Int64`](/pkg/sync/atomic/#Int64),
[`Uint32`](/pkg/sync/atomic/#Uint32),
[`Uint64`](/pkg/sync/atomic/#Uint64),
[`Uintptr`](/pkg/sync/atomic/#Uintptr), và
[`Pointer`](/pkg/sync/atomic/#Pointer).
Các kiểu này ẩn các giá trị bên dưới để tất cả các truy cập buộc phải dùng
API atomic.
[`Pointer`](/pkg/sync/atomic/#Pointer) cũng tránh
nhu cầu chuyển đổi sang
[`unsafe.Pointer`](/pkg/unsafe/#Pointer) tại call site.
[`Int64`](/pkg/sync/atomic/#Int64) và
[`Uint64`](/pkg/sync/atomic/#Uint64) được
tự động căn chỉnh đến ranh giới 64-bit trong struct và dữ liệu được phân bổ,
ngay cả trên các hệ thống 32-bit.

### Tra cứu PATH {#os-exec-path}

<!-- https://go.dev/issue/43724 -->
<!-- CL 381374 -->
<!-- CL 403274 -->
[`Command`](/pkg/os/exec/#Command) và
[`LookPath`](/pkg/os/exec/#LookPath) không còn
cho phép kết quả từ tra cứu PATH được tìm thấy tương đối so với thư mục hiện tại.
Điều này loại bỏ một [nguồn vấn đề bảo mật phổ biến](/blog/path-security)
nhưng cũng có thể phá vỡ các chương trình hiện tại phụ thuộc vào việc dùng, chẳng hạn, `exec.Command("prog")`
để chạy một binary có tên `prog` (hoặc, trên Windows, `prog.exe`) trong thư mục hiện tại.
Xem tài liệu gói [`os/exec`](/pkg/os/exec/) để biết
thông tin về cách cập nhật tốt nhất các chương trình như vậy.

<!-- https://go.dev/issue/43947 -->
Trên Windows, `Command` và `LookPath` hiện tuân theo
biến môi trường [`NoDefaultCurrentDirectoryInExePath`](https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-needcurrentdirectoryforexepatha),
giúp có thể tắt
tìm kiếm ngầm mặc định của "`.`" trong tra cứu PATH trên hệ thống Windows.

### Các thay đổi nhỏ trong thư viện {#minor_library_changes}

Như thường lệ, có nhiều thay đổi và cập nhật nhỏ đối với thư viện,
được thực hiện với [cam kết tương thích](/doc/go1compat) của Go 1
trong tâm trí.
Cũng có nhiều cải tiến hiệu suất khác nhau, không được liệt kê ở đây.

#### [archive/zip](/pkg/archive/zip/)

<!-- CL 387976 -->
[`Reader`](/pkg/archive/zip/#Reader)
hiện bỏ qua dữ liệu không phải ZIP ở đầu tệp ZIP, khớp với hầu hết các triển khai khác.
Điều này cần thiết để đọc một số tệp Java JAR, trong số các mục đích sử dụng khác.

<!-- archive/zip -->

#### [crypto/elliptic](/pkg/crypto/elliptic/)

<!-- CL 382995 -->
Thao tác trên các điểm đường cong không hợp lệ (những điểm mà phương thức
`IsOnCurve` trả về false, và không bao giờ được trả về
bởi `Unmarshal` hoặc phương thức `Curve` thao tác trên một
điểm hợp lệ) luôn là hành vi không xác định và có thể dẫn đến
tấn công phục hồi khóa. Nếu một điểm không hợp lệ được cung cấp cho
[`Marshal`](/pkg/crypto/elliptic/#Marshal),
[`MarshalCompressed`](/pkg/crypto/elliptic/#MarshalCompressed),
[`Add`](/pkg/crypto/elliptic/#Curve.Add),
[`Double`](/pkg/crypto/elliptic/#Curve.Double), hoặc
[`ScalarMult`](/pkg/crypto/elliptic/#Curve.ScalarMult),
chúng hiện sẽ panic.

<!-- golang.org/issue/52182 -->
Các thao tác `ScalarBaseMult` trên các đường cong `P224`,
`P384` và `P521` hiện nhanh hơn đến ba
lần, dẫn đến cải thiện tương tự trong một số thao tác ECDSA. Triển khai
`P256` chung (không được tối ưu hóa cho nền tảng cụ thể) đã được
thay thế bằng một triển khai xuất phát từ mô hình được xác minh chính thức; điều này có thể
dẫn đến chậm lại đáng kể trên các nền tảng 32-bit.

<!-- crypto/elliptic -->

#### [crypto/rand](/pkg/crypto/rand/)

<!-- CL 370894 -->
<!-- CL 390038 -->
[`Read`](/pkg/crypto/rand/#Read) không còn buffer
dữ liệu ngẫu nhiên thu được từ hệ điều hành giữa các lần gọi. Các ứng dụng
thực hiện nhiều lần đọc nhỏ với tần suất cao có thể chọn bao bọc
[`Reader`](/pkg/crypto/rand/#Reader) trong một
[`bufio.Reader`](/pkg/bufio/#Reader) vì lý do hiệu suất,
cẩn thận dùng
[`io.ReadFull`](/pkg/io/#ReadFull)
để đảm bảo không có lần đọc một phần nào xảy ra.

<!-- CL 375215 -->
Trên Plan 9, `Read` đã được triển khai lại, thay thế thuật toán ANSI
X9.31 bằng bộ tạo key erasure nhanh.

<!-- CL 391554 -->
<!-- CL 387554 -->
Triển khai [`Prime`](/pkg/crypto/rand/#Prime)
đã được thay đổi để chỉ sử dụng rejection sampling,
giúp loại bỏ sai lệch khi tạo các số nguyên tố nhỏ trong ngữ cảnh không mã hóa,
loại bỏ một rò rỉ timing nhỏ có thể xảy ra,
và phù hợp hơn với hành vi của BoringSSL,
đồng thời đơn giản hóa việc triển khai.
Thay đổi này tạo ra các output khác nhau cho một stream nguồn ngẫu nhiên nhất định
so với triển khai trước đó,
điều này có thể phá vỡ các kiểm thử được viết để mong đợi các kết quả cụ thể từ
các nguồn ngẫu nhiên xác định cụ thể.
Để giúp ngăn chặn các vấn đề như vậy trong tương lai,
triển khai hiện có chủ ý không xác định (non-deterministic) đối với luồng đầu vào.

<!-- crypto/rand -->

#### [crypto/tls](/pkg/crypto/tls/)

<!-- CL 400974 -->
<!-- https://go.dev/issue/45428 -->
Tùy chọn `GODEBUG` `tls10default=1` đã được
loại bỏ. Vẫn có thể bật TLS 1.0 phía client bằng cách đặt
[`Config.MinVersion`](/pkg/crypto/tls/#Config.MinVersion).

<!-- CL 384894 -->
Server và client TLS hiện từ chối các extension trùng lặp trong
TLS handshake, theo yêu cầu của RFC 5246, Mục 7.4.1.4 và RFC 8446, Mục
4.2.

<!-- crypto/tls -->

#### [crypto/x509](/pkg/crypto/x509/)

<!-- CL 285872 -->
[`CreateCertificate`](/pkg/crypto/x509/#CreateCertificate)
không còn hỗ trợ tạo chứng chỉ với `SignatureAlgorithm`
được đặt thành `MD5WithRSA`.

<!-- CL 400494 -->
`CreateCertificate` không còn chấp nhận số serial âm.

<!-- CL 399827 -->
`CreateCertificate` sẽ không còn phát ra SEQUENCE rỗng
khi chứng chỉ được tạo ra không có extension.

<!-- CL 396774 -->
Việc loại bỏ tùy chọn `GODEBUG` `x509sha1=1`,
ban đầu được lên kế hoạch cho Go 1.19, đã được lên lịch lại cho bản phát hành tương lai.
Các ứng dụng đang sử dụng nó nên chuyển đổi. Các cuộc tấn công thực tế chống lại
SHA-1 đã được chứng minh từ năm 2017 và các Cơ quan Chứng nhận (Certificate Authorities) tin cậy công khai
đã không phát hành chứng chỉ SHA-1 kể từ năm 2015.

<!-- CL 383215 -->
[`ParseCertificate`](/pkg/crypto/x509/#ParseCertificate)
và [`ParseCertificateRequest`](/pkg/crypto/x509/#ParseCertificateRequest)
hiện từ chối các chứng chỉ và CSR chứa các extension trùng lặp.

<!-- https://go.dev/issue/46057 -->
<!-- https://go.dev/issue/35044 -->
<!-- CL 398237 -->
<!-- CL 400175 -->
<!-- CL 388915 -->
Các phương thức mới [`CertPool.Clone`](/pkg/crypto/x509/#CertPool.Clone)
và [`CertPool.Equal`](/pkg/crypto/x509/#CertPool.Equal)
cho phép clone một `CertPool` và kiểm tra sự tương đương của hai
`CertPool`.

<!-- https://go.dev/issue/50674 -->
<!-- CL 390834 -->
Hàm mới [`ParseRevocationList`](/pkg/crypto/x509/#ParseRevocationList)
cung cấp một bộ phân tích cú pháp CRL nhanh hơn, an toàn hơn để sử dụng, trả về một
[`RevocationList`](/pkg/crypto/x509/#RevocationList).
Phân tích cú pháp CRL cũng điền vào các trường `RevocationList` mới
`RawIssuer`, `Signature`,
`AuthorityKeyId` và `Extensions`, bị bỏ qua bởi
[`CreateRevocationList`](/pkg/crypto/x509/#CreateRevocationList).

Phương thức mới [`RevocationList.CheckSignatureFrom`](/pkg/crypto/x509/#RevocationList.CheckSignatureFrom)
kiểm tra rằng chữ ký trên CRL là một chữ ký hợp lệ từ một
[`Certificate`](/pkg/crypto/x509/#Certificate).

Các hàm [`ParseCRL`](/pkg/crypto/x509/#ParseCRL) và
[`ParseDERCRL`](/pkg/crypto/x509/#ParseDERCRL)
hiện không còn được khuyến nghị (deprecated) thay cho `ParseRevocationList`.
Phương thức [`Certificate.CheckCRLSignature`](/pkg/crypto/x509/#Certificate.CheckCRLSignature)
không còn được khuyến nghị thay cho `RevocationList.CheckSignatureFrom`.

<!-- CL 389555, CL 401115, CL 403554 -->
Bộ xây dựng đường dẫn của [`Certificate.Verify`](/pkg/crypto/x509/#Certificate.Verify)
đã được cải tổ và hiện nên tạo ra các chuỗi tốt hơn và/hoặc hiệu quả hơn trong các tình huống phức tạp.
Các ràng buộc tên (name constraints) hiện cũng được áp dụng trên các chứng chỉ không phải lá (non-leaf).

<!-- crypto/x509 -->

#### [crypto/x509/pkix](/pkg/crypto/x509/pkix/)

<!-- CL 390834 -->
Các kiểu [`CertificateList`](/pkg/crypto/x509/pkix/#CertificateList) và
[`TBSCertificateList`](/pkg/crypto/x509/pkix/#TBSCertificateList)
đã được đánh dấu không còn được khuyến nghị. Chức năng CRL mới của [`crypto/x509`](#crypto/x509)
nên được dùng thay thế.

<!-- crypto/x509/pkix -->

#### [debug/elf](/pkg/debug/elf/)

<!-- CL 396735 -->
Các hằng số mới `EM_LOONGARCH` và `R_LARCH_*`
hỗ trợ cổng loong64.

<!-- debug/elf -->

#### [debug/pe](/pkg/debug/pe/)

<!-- https://go.dev/issue/51868 -->
<!-- CL 394534 -->
Phương thức mới [`File.COFFSymbolReadSectionDefAux`](/pkg/debug/pe/#File.COFFSymbolReadSectionDefAux),
trả về một [`COFFSymbolAuxFormat5`](/pkg/debug/pe/#COFFSymbolAuxFormat5),
cung cấp quyền truy cập vào thông tin COMDAT trong các phần tệp PE.
Những thứ này được hỗ trợ bởi các hằng số `IMAGE_COMDAT_*` và `IMAGE_SCN_*` mới.

<!-- debug/pe -->

#### [encoding/binary](/pkg/encoding/binary/)

<!-- https://go.dev/issue/50601 -->
<!-- CL 386017 -->
<!-- CL 389636 -->
Interface mới
[`AppendByteOrder`](/pkg/encoding/binary/#AppendByteOrder)
cung cấp các phương thức hiệu quả để thêm (append) `uint16`, `uint32` hoặc `uint64`
vào một byte slice.
[`BigEndian`](/pkg/encoding/binary/#BigEndian) và
[`LittleEndian`](/pkg/encoding/binary/#LittleEndian) hiện triển khai interface này.

<!-- https://go.dev/issue/51644 -->
<!-- CL 400176 -->
Tương tự, các hàm mới
[`AppendUvarint`](/pkg/encoding/binary/#AppendUvarint) và
[`AppendVarint`](/pkg/encoding/binary/#AppendVarint)
là các phiên bản hiệu quả để thêm (append) của
[`PutUvarint`](/pkg/encoding/binary/#PutUvarint) và
[`PutVarint`](/pkg/encoding/binary/#PutVarint).

<!-- encoding/binary -->

#### [encoding/csv](/pkg/encoding/csv/)

<!-- https://go.dev/issue/43401 -->
<!-- CL 405675 -->
Phương thức mới
[`Reader.InputOffset`](/pkg/encoding/csv/#Reader.InputOffset)
báo cáo vị trí đầu vào hiện tại của reader dưới dạng byte offset,
tương tự như
[`Decoder.InputOffset`](/pkg/encoding/json/#Decoder.InputOffset) của `encoding/json`.

<!-- encoding/csv -->

#### [encoding/xml](/pkg/encoding/xml/)

<!-- https://go.dev/issue/45628 -->
<!-- CL 311270 -->
Phương thức mới
[`Decoder.InputPos`](/pkg/encoding/xml/#Decoder.InputPos)
báo cáo vị trí đầu vào hiện tại của reader dưới dạng dòng và cột,
tương tự như
[`Decoder.FieldPos`](/pkg/encoding/csv/#Decoder.FieldPos) của `encoding/csv`.

<!-- encoding/xml -->

#### [flag](/pkg/flag/)

<!-- https://go.dev/issue/45754 -->
<!-- CL 313329 -->
Hàm mới
[`TextVar`](/pkg/flag/#TextVar)
định nghĩa một cờ với giá trị triển khai
[`encoding.TextUnmarshaler`](/pkg/encoding/#TextUnmarshaler),
cho phép các biến cờ dòng lệnh có các kiểu như
[`big.Int`](/pkg/math/big/#Int),
[`netip.Addr`](/pkg/net/netip/#Addr), và
[`time.Time`](/pkg/time/#Time).

<!-- flag -->

#### [fmt](/pkg/fmt/)

<!-- https://go.dev/issue/47579 -->
<!-- CL 406177 -->
Các hàm mới
[`Append`](/pkg/fmt/#Append),
[`Appendf`](/pkg/fmt/#Appendf), và
[`Appendln`](/pkg/fmt/#Appendln)
thêm dữ liệu được định dạng vào byte slice.

<!-- fmt -->

#### [go/parser](/pkg/go/parser/)

<!-- CL 403696 -->
Parser hiện nhận diện `~x` như một biểu thức unary với toán tử
[token.TILDE](/pkg/go/token/#TILDE),
cho phép phục hồi lỗi tốt hơn khi một ràng buộc kiểu như `~int` được dùng trong ngữ cảnh không đúng.

<!-- go/parser -->

#### [go/types](/pkg/go/types/)

<!-- https://go.dev/issue/51682 -->
<!-- CL 395535 -->
Các phương thức mới [`Func.Origin`](/pkg/go/types/#Func.Origin)
và [`Var.Origin`](/pkg/go/types/#Var.Origin) trả về
[`Object`](/pkg/go/types/#Object) tương ứng của
kiểu generic cho các đối tượng [`Func`](/pkg/go/types/#Func)
và [`Var`](/pkg/go/types/#Var) tổng hợp được tạo trong quá trình
khởi tạo kiểu (type instantiation).

<!-- https://go.dev/issue/52728 -->
<!-- CL 404885 -->
Không còn có thể tạo ra vô số lần khởi tạo kiểu
[`Named`](/pkg/go/types/#Named) khác nhau nhưng giống hệt nhau qua
các lệnh gọi đệ quy đến
[`Named.Underlying`](/pkg/go/types/#Named.Underlying) hoặc
[`Named.Method`](/pkg/go/types/#Named.Method).

<!-- go/types -->

#### [hash/maphash](/pkg/hash/maphash/)

<!-- https://go.dev/issue/42710 -->
<!-- CL 392494 -->
Các hàm mới
[`Bytes`](/pkg/hash/maphash/#Bytes)
và
[`String`](/pkg/hash/maphash/#String)
cung cấp cách hiệu quả để hash một byte slice hoặc chuỗi đơn lẻ.
Chúng tương đương với việc dùng
[`Hash`](/pkg/hash/maphash/#Hash)
tổng quát hơn với một lần ghi duy nhất, nhưng chúng tránh chi phí khởi tạo cho các đầu vào nhỏ.

<!-- hash/maphash -->

#### [html/template](/pkg/html/template/)

<!-- https://go.dev/issue/46121 -->
<!-- CL 389156 -->
Kiểu [`FuncMap`](/pkg/html/template/#FuncMap)
hiện là alias cho
[`FuncMap`](/pkg/text/template/#FuncMap) của `text/template`
thay vì một kiểu có tên riêng.
Điều này cho phép viết code hoạt động trên `FuncMap` từ cả hai ngữ cảnh.

<!-- https://go.dev/issue/59153 -->
<!-- CL 481987 -->
Go 1.19.8 và mới hơn
[không cho phép các action trong ECMAScript 6 template literal.](/pkg/html/template#hdr-Security_Model)
Hành vi này có thể được đảo ngược bằng cài đặt `GODEBUG=jstmpllitinterp=1`.

<!-- html/template -->

#### [image/draw](/pkg/image/draw/)

<!-- CL 396795 -->
[`Draw`](/pkg/image/draw/#Draw) với toán tử
[`Src`](/pkg/image/draw/#Src) giữ nguyên
các màu alpha không nhân trước (non-premultiplied-alpha) khi cả ảnh đích và nguồn đều là
[`image.NRGBA`](/pkg/image/#NRGBA)
hoặc cả hai đều là [`image.NRGBA64`](/pkg/image/#NRGBA64).
Điều này khôi phục lại thay đổi hành vi vô tình được giới thiệu bởi một
tối ưu hóa thư viện Go 1.18; code hiện khớp với hành vi trong Go 1.17 và trước đó.

<!-- image/draw -->

#### [io](/pkg/io/)

<!-- https://go.dev/issue/51566 -->
<!-- CL 400236 -->
Kết quả của [`NopCloser`](/pkg/io/#NopCloser) hiện triển khai
[`WriterTo`](/pkg/io/#WriterTo)
khi đầu vào của nó triển khai.

<!-- https://go.dev/issue/50842 -->
Kết quả của [`MultiReader`](/pkg/io/#MultiReader) hiện triển khai
[`WriterTo`](/pkg/io/#WriterTo) vô điều kiện.
Nếu bất kỳ reader bên dưới nào không triển khai `WriterTo`,
nó sẽ được mô phỏng thích hợp.

<!-- io -->

#### [mime](/pkg/mime/)

<!-- CL 406894 -->
Chỉ trên Windows, gói mime hiện bỏ qua một mục registry
ghi rằng extension `.js` nên có kiểu MIME
`text/plain`. Đây là một
cấu hình sai vô tình phổ biến trên các hệ thống Windows. Hiệu ứng là
`.js` sẽ có kiểu MIME mặc định
`text/javascript; charset=utf-8`.
Các ứng dụng mong đợi `text/plain` trên Windows hiện phải
gọi rõ ràng
[`AddExtensionType`](/pkg/mime/#AddExtensionType).

<!-- mime -->

#### [mime/multipart](/pkg/mime/multipart)

<!-- https://go.dev/issue/59153 -->
<!-- CL 481985 -->
Trong Go 1.19.8 và mới hơn, gói này đặt giới hạn kích thước
dữ liệu MIME mà nó xử lý để bảo vệ chống lại các đầu vào độc hại.
`Reader.NextPart` và `Reader.NextRawPart` giới hạn
số header trong một phần tối đa 10000 và `Reader.ReadForm` giới hạn
tổng số header trong tất cả `FileHeader` tối đa 10000.
Các giới hạn này có thể được điều chỉnh với cài đặt `GODEBUG=multipartmaxheaders`.
`Reader.ReadForm` tiếp tục giới hạn số phần trong một form tối đa 1000.
Giới hạn này có thể được điều chỉnh với cài đặt `GODEBUG=multipartmaxparts`.

<!-- mime/multipart -->

#### [net](/pkg/net/)

<!-- CL 386016 -->
Resolver Go thuần túy hiện sẽ dùng EDNS(0) để bao gồm độ dài gói
trả lời tối đa được đề xuất, cho phép các gói trả lời chứa
tối đa 1232 byte (trước đây tối đa là 512).
Trong trường hợp không chắc xảy ra rằng điều này gây ra vấn đề với một DNS resolver
cục bộ, việc đặt biến môi trường
`GODEBUG=netdns=cgo` để sử dụng resolver dựa trên cgo
sẽ hoạt động.
Vui lòng báo cáo bất kỳ vấn đề nào như vậy trên [bộ theo dõi
issue](/issue/new).

<!-- https://go.dev/issue/51428 -->
<!-- CL 396877 -->
Khi một hàm hoặc phương thức trong gói net trả về lỗi "I/O timeout",
lỗi đó hiện sẽ thỏa mãn `errors.Is(err,
  context.DeadlineExceeded)`. Khi một hàm trong gói net
trả về lỗi "operation was canceled", lỗi đó hiện sẽ
thỏa mãn `errors.Is(err, context.Canceled)`.
Những thay đổi này nhằm giúp code dễ kiểm tra hơn
trong các trường hợp hủy context hoặc timeout khiến một hàm hoặc phương thức trong gói net
trả về lỗi, đồng thời giữ nguyên khả năng tương thích ngược cho các thông báo lỗi.

<!-- https://go.dev/issue/33097 -->
<!-- CL 400654 -->
[`Resolver.PreferGo`](/pkg/net/#Resolver.PreferGo)
hiện được triển khai trên Windows và Plan 9. Trước đây nó chỉ hoạt động trên các nền tảng Unix.
Kết hợp với
[`Dialer.Resolver`](/pkg/net/#Dialer.Resolver) và
[`Resolver.Dial`](/pkg/net/#Resolver.Dial), hiện
có thể viết các chương trình di động và kiểm soát tất cả các tra cứu tên DNS
khi dialing.

Gói `net` hiện có hỗ trợ ban đầu cho build tag `netgo`
trên Windows. Khi được dùng, gói sử dụng client DNS Go (như được dùng
bởi `Resolver.PreferGo`) thay vì hỏi Windows về
kết quả DNS. Server DNS ngược dòng mà nó tìm kiếm từ Windows
có thể chưa chính xác với cấu hình mạng hệ thống phức tạp.

<!-- net -->

#### [net/http](/pkg/net/http/)

<!-- CL 269997 -->
[`ResponseWriter.WriteHeader`](/pkg/net/http/#ResponseWriter)
hiện hỗ trợ gửi các header thông tin 1xx do người dùng định nghĩa.

<!-- CL 361397 -->
`io.ReadCloser` được trả về bởi
[`MaxBytesReader`](/pkg/net/http/#MaxBytesReader)
hiện sẽ trả về kiểu lỗi được định nghĩa
[`MaxBytesError`](/pkg/net/http/#MaxBytesError)
khi giới hạn đọc của nó bị vượt quá.

<!-- CL 375354 -->
HTTP client sẽ xử lý phản hồi 3xx không có
header `Location` bằng cách trả nó cho người gọi,
thay vì xử lý nó như một lỗi.

<!-- net/http -->

#### [net/url](/pkg/net/url/)

<!-- CL 374654 -->
Hàm mới
[`JoinPath`](/pkg/net/url/#JoinPath)
và phương thức
[`URL.JoinPath`](/pkg/net/url/#URL.JoinPath)
tạo một `URL` mới bằng cách nối một danh sách các phần tử đường dẫn.

<!-- https://go.dev/issue/46059 -->
Kiểu `URL` hiện phân biệt giữa các URL không có
authority và các URL có authority trống. Ví dụ:
`http:///path` có authority trống (host),
trong khi `http:/path` không có.

Trường mới [`URL`](/pkg/net/url/#URL)
`OmitHost` được đặt thành `true` khi một
`URL` có authority trống.

<!-- net/url -->

#### [os/exec](/pkg/os/exec/)

<!-- https://go.dev/issue/50599 -->
<!-- CL 401340 -->
Một [`Cmd`](/pkg/os/exec/#Cmd) với trường `Dir` không rỗng
và `Env` nil hiện ngầm đặt biến môi trường `PWD`
cho tiến trình con để khớp với `Dir`.

Phương thức mới [`Cmd.Environ`](/pkg/os/exec/#Cmd.Environ) báo cáo
môi trường được dùng để chạy lệnh, bao gồm
biến `PWD` được đặt ngầm.

<!-- os/exec -->

#### [reflect](/pkg/reflect/)

<!-- https://go.dev/issue/47066 -->
<!-- CL 357331 -->
Phương thức [`Value.Bytes`](/pkg/reflect/#Value.Bytes)
hiện chấp nhận các mảng có thể địa chỉ (addressable arrays) ngoài slice.

<!-- CL 400954 -->
Các phương thức [`Value.Len`](/pkg/reflect/#Value.Len)
và [`Value.Cap`](/pkg/reflect/#Value.Cap)
hiện hoạt động thành công trên con trỏ đến mảng và trả về độ dài của mảng đó,
để khớp với những gì [hàm built-in
`len` và `cap`](/ref/spec#Length_and_capacity) làm.

<!-- reflect -->

#### [regexp/syntax](/pkg/regexp/syntax/)

<!-- https://go.dev/issue/51684 -->
<!-- CL 401076 -->
Go 1.18 release candidate 1, Go 1.17.8 và Go 1.16.15 bao gồm một bản sửa lỗi bảo mật
cho parser biểu thức chính quy, khiến nó từ chối các biểu thức lồng nhau rất sâu.
Vì các bản vá phát hành Go không giới thiệu API mới,
parser trả về [`syntax.ErrInternalError`](/pkg/regexp/syntax/#ErrInternalError) trong trường hợp này.
Go 1.19 thêm lỗi cụ thể hơn, [`syntax.ErrNestingDepth`](/pkg/regexp/syntax/#ErrNestingDepth),
mà parser hiện trả về thay thế.

<!-- regexp -->

#### [runtime](/pkg/runtime/)

<!-- https://go.dev/issue/51461 -->
Hàm [`GOROOT`](/pkg/runtime/#GOROOT) hiện trả về chuỗi rỗng
(thay vì `"go"`) khi binary được build với
cờ `-trimpath` và biến `GOROOT`
không được đặt trong môi trường tiến trình.

<!-- runtime -->

#### [runtime/metrics](/pkg/runtime/metrics/)

<!-- https://go.dev/issue/47216 -->
<!-- CL 404305 -->
[Metric](/pkg/runtime/metrics/#hdr-Supported_metrics) mới `/sched/gomaxprocs:threads`
báo cáo giá trị
[`runtime.GOMAXPROCS`](/pkg/runtime/#GOMAXPROCS) hiện tại.

<!-- https://go.dev/issue/47216 -->
<!-- CL 404306 -->
[Metric](/pkg/runtime/metrics/#hdr-Supported_metrics) mới `/cgo/go-to-c-calls:calls`
báo cáo tổng số lệnh gọi từ Go sang C. Metric này
giống hệt với hàm
[`runtime.NumCgoCall`](/pkg/runtime/#NumCgoCall).

<!-- https://go.dev/issue/48409 -->
<!-- CL 403614 -->
[Metric](/pkg/runtime/metrics/#hdr-Supported_metrics) mới `/gc/limiter/last-enabled:gc-cycle`
báo cáo chu kỳ GC cuối cùng khi bộ giới hạn CPU GC được bật. Xem
[ghi chú runtime](#runtime) để biết chi tiết về bộ giới hạn CPU GC.

<!-- runtime/metrics -->

#### [runtime/pprof](/pkg/runtime/pprof/)

<!-- https://go.dev/issue/33250 -->
<!-- CL 387415 -->
Thời gian dừng thế giới (stop-the-world pause) đã được giảm đáng kể khi
thu thập goroutine profile, giảm tác động độ trễ tổng thể đến
ứng dụng.

<!-- CL 391434 -->
`MaxRSS` hiện được báo cáo trong heap profile cho tất cả các hệ điều hành Unix
(trước đây chỉ được báo cáo cho
`GOOS=android`, `darwin`, `ios` và
`linux`).

<!-- runtime/pprof -->

#### [runtime/race](/pkg/runtime/race/)

<!-- https://go.dev/issue/49761 -->
<!-- CL 333529 -->
Trình phát hiện race đã được nâng cấp lên phiên bản thread sanitizer
v3 trên tất cả các nền tảng được hỗ trợ
ngoại trừ `windows/amd64`
và `openbsd/amd64`, vẫn ở v2.
So với v2, nó thường nhanh hơn 1,5 đến 2 lần, sử dụng một nửa
bộ nhớ và hỗ trợ số lượng goroutine không giới hạn.
Trên Linux, trình phát hiện race hiện yêu cầu ít nhất glibc phiên bản
2.17 và GNU binutils 2.26.

<!-- CL 336549 -->
Trình phát hiện race hiện được hỗ trợ trên `GOARCH=s390x`.

<!-- https://go.dev/issue/52090 -->
Hỗ trợ trình phát hiện race cho `openbsd/amd64` đã được
loại bỏ khỏi upstream thread sanitizer, vì vậy khó có thể
được cập nhật từ v2.

<!-- runtime/race -->

#### [runtime/trace](/pkg/runtime/trace/)

<!-- CL 400795 -->
Khi tracing và
[CPU profiler](/pkg/runtime/pprof/#StartCPUProfile) được
bật đồng thời, execution trace bao gồm các mẫu CPU profile
như các sự kiện tức thời.

<!-- runtime/trace -->

#### [sort](/pkg/sort/)

<!-- CL 371574 -->
Thuật toán sắp xếp đã được viết lại để sử dụng
[pattern-defeating quicksort](https://arxiv.org/pdf/2106.05123.pdf), nhanh hơn
cho một số kịch bản phổ biến.

<!-- https://go.dev/issue/50340 -->
<!-- CL 396514 -->
Hàm mới
[`Find`](/pkg/sort/#Find)
giống như
[`Search`](/pkg/sort/#Search)
nhưng thường dễ sử dụng hơn: nó trả về thêm một boolean báo cáo liệu có tìm thấy giá trị bằng nhau không.

<!-- sort -->

#### [strconv](/pkg/strconv/)

<!-- CL 397255 -->
[`Quote`](/pkg/strconv/#Quote)
và các hàm liên quan hiện trích dẫn rune U+007F là `\x7f`,
không phải `\u007f`,
để nhất quán với các giá trị ASCII khác.

<!-- strconv -->

#### [syscall](/pkg/syscall/)

<!-- https://go.dev/issue/51192 -->
<!-- CL 385796 -->
Trên PowerPC (`GOARCH=ppc64`, `ppc64le`),
[`Syscall`](/pkg/syscall/#Syscall),
[`Syscall6`](/pkg/syscall/#Syscall6),
[`RawSyscall`](/pkg/syscall/#RawSyscall), và
[`RawSyscall6`](/pkg/syscall/#RawSyscall6)
hiện luôn trả về 0 cho giá trị trả về `r2` thay vì
một giá trị không xác định.

<!-- CL 391434 -->
Trên AIX và Solaris, [`Getrusage`](/pkg/syscall/#Getrusage) hiện được định nghĩa.

<!-- syscall -->

#### [time](/pkg/time/)

<!-- https://go.dev/issue/51414 -->
<!-- CL 393515 -->
Phương thức mới
[`Duration.Abs`](/pkg/time/#Duration.Abs)
cung cấp cách thuận tiện và an toàn để lấy giá trị tuyệt đối của một duration,
chuyển đổi -2^63 thành 2^63-1.
(Trường hợp ranh giới này có thể xảy ra do kết quả của việc trừ thời gian gần đây từ thời gian không.)

<!-- https://go.dev/issue/50062 -->
<!-- CL 405374 -->
Phương thức mới
[`Time.ZoneBounds`](/pkg/time/#Time.ZoneBounds)
trả về thời gian bắt đầu và kết thúc của múi giờ có hiệu lực tại một thời điểm nhất định.
Nó có thể được dùng trong một vòng lặp để liệt kê tất cả các chuyển đổi múi giờ đã biết tại một vị trí nhất định.

<!-- time -->

<!-- Silence these false positives from x/build/cmd/relnote: -->
<!-- CL 382460 -->
<!-- CL 384154 -->
<!-- CL 384554 -->
<!-- CL 392134 -->
<!-- CL 392414 -->
<!-- CL 396215 -->
<!-- CL 403058 -->
<!-- CL 410133 -->
<!-- https://go.dev/issue/27837 -->
<!-- https://go.dev/issue/38340 -->
<!-- https://go.dev/issue/42516 -->
<!-- https://go.dev/issue/45713 -->
<!-- https://go.dev/issue/46654 -->
<!-- https://go.dev/issue/48257 -->
<!-- https://go.dev/issue/50447 -->
<!-- https://go.dev/issue/50720 -->
<!-- https://go.dev/issue/50792 -->
<!-- https://go.dev/issue/51115 -->
<!-- https://go.dev/issue/51447 -->
