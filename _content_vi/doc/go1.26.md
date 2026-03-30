---
title: Ghi chú phát hành Go 1.26
---

<style>
  main ul li { margin: 0.5em 0; }
</style>

## Giới thiệu về Go 1.26 {#introduction}

Bản phát hành Go mới nhất, phiên bản 1.26, ra đời vào [tháng 2 năm 2026](/doc/devel/release#go1.26.0), sáu tháng sau [Go 1.25](/doc/go1.25).
Phần lớn các thay đổi nằm ở phần triển khai toolchain, runtime và thư viện.
Như thường lệ, bản phát hành này duy trì cam kết tương thích của Go 1.
Chúng tôi kỳ vọng hầu hết các chương trình Go sẽ tiếp tục biên dịch và chạy như trước đây.

## Thay đổi về ngôn ngữ {#language}

<!-- https://go.dev/issue/45624 --->

Hàm built-in `new`, dùng để tạo một biến mới, hiện cho phép
toán hạng của nó là một biểu thức, chỉ định giá trị ban đầu của
biến.

Tính năng này đặc biệt hữu ích khi làm việc với các gói serialization
như `encoding/json` hoặc protocol buffer sử dụng con trỏ để biểu diễn giá trị tùy chọn, vì nó cho phép
điền một trường tùy chọn trong một biểu thức đơn giản, ví dụ:

```go
import "encoding/json"

type Person struct {
	Name string   `json:"name"`
	Age  *int     `json:"age"` // age if known; nil otherwise
}

func personJSON(name string, born time.Time) ([]byte, error) {
	return json.Marshal(Person{
		Name: name,
		Age:  new(yearsSince(born)),
	})
}

func yearsSince(t time.Time) int {
	return int(time.Since(t).Hours() / (365.25 * 24)) // approximately
}
```

<!-- https://go.dev/issue/75883 --->

Hạn chế rằng một kiểu generic không được tham chiếu chính nó trong danh sách tham số kiểu
đã được dỡ bỏ.
Hiện có thể chỉ định các ràng buộc kiểu (type constraint) tham chiếu đến kiểu generic đang được
ràng buộc.
Ví dụ, một kiểu generic `Adder` có thể yêu cầu nó được khởi tạo với
một kiểu giống chính nó:

```go
type Adder[A Adder[A]] interface {
	Add(A) A
}

func algo[A Adder[A]](x, y A) A {
	return x.Add(y)
}
```

Trước đây, tham chiếu tự đến `Adder` ở dòng đầu tiên không được cho phép.
Ngoài việc làm cho các ràng buộc kiểu mạnh hơn, thay đổi này cũng đơn giản hóa các quy tắc spec
cho tham số kiểu đôi chút.

## Công cụ {#tools}

### Lệnh Go {#go-command}

<!-- go.dev/issue/75432 -->
Lệnh `go fix` truyền thống đã được đại tu hoàn toàn và hiện là nơi chứa
các *modernizer* của Go. Nó cung cấp một cách đáng tin cậy để cập nhật
code base Go lên các idiom và API thư viện lõi mới nhất. Bộ modernizer ban đầu bao gồm
hàng chục fixer để tận dụng các tính năng hiện đại của ngôn ngữ
và thư viện Go, cũng như một inliner ở cấp độ nguồn cho phép người dùng
tự động hóa các lần chuyển đổi API của riêng họ bằng cách sử dụng
[chỉ thị `//go:fix inline`](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/inline#hdr-Analyzer_inline).
Các fixer này không nên thay đổi hành vi của chương trình, vì vậy nếu bạn gặp
bất kỳ vấn đề nào với một sửa chữa được thực hiện bởi `go fix`, vui lòng [báo cáo](/issue/new).

Lệnh `go fix` được viết lại dựa trên chính [framework phân tích
Go](https://pkg.go.dev/golang.org/x/tools/go/analysis) như `go vet`.
Điều này có nghĩa là các analyzer cung cấp chẩn đoán trong `go vet`
có thể được dùng để đề xuất và áp dụng các sửa chữa trong `go fix`.
Các fixer lịch sử của lệnh `go fix`, tất cả đều lỗi thời,
đã được loại bỏ.

Hai bài đăng blog Go sắp tới sẽ đi vào chi tiết hơn về modernizer, inliner,
và cách tận dụng tốt nhất `go fix`.
<!-- TODO(adonovan): link to blog posts once live. -->

<!-- go.dev/issue/74748 -->
`go mod init` hiện mặc định phiên bản `go` thấp hơn trong các tệp `go.mod` mới.
Chạy `go mod init`
bằng toolchain phiên bản `1.N.X` sẽ tạo tệp `go.mod`
chỉ định phiên bản Go `go 1.(N-1).0`. Các phiên bản pre-release của `1.N` sẽ
tạo các tệp `go.mod` chỉ định `go 1.(N-2).0`. Ví dụ, các
release candidate Go 1.26 sẽ tạo các tệp `go.mod` với `go 1.24.0`, và Go 1.26
cùng các bản phát hành nhỏ của nó sẽ tạo các tệp `go.mod` với `go 1.25.0`. Điều này nhằm
khuyến khích tạo các module tương thích với các phiên bản Go hiện đang được hỗ trợ. Để kiểm soát thêm về phiên bản `go` trong các module mới,
`go mod init` có thể được tiếp theo bởi `go get go@version`.

<!-- go.dev/issue/74667 -->
`cmd/doc` và `go tool doc` đã bị xóa. `go doc` có thể được dùng như
một thay thế cho `go tool doc`: nó nhận cùng cờ và đối số và
có cùng hành vi.

### Pprof {#pprof}

<!-- go.dev/issue/74774 -->
Giao diện web của công cụ `pprof`, được bật bằng cờ `-http`, hiện mặc định sử dụng chế độ xem flame graph.
Chế độ xem đồ thị trước đây vẫn có sẵn trong menu "View -> Graph", hoặc qua `/ui/graph`.

## Runtime {#runtime}

### Garbage collector mới

Garbage collector Green Tea, trước đây có sẵn như một thử nghiệm trong
Go 1.25, hiện được bật mặc định sau khi tiếp nhận phản hồi.

Thiết kế của garbage collector này cải thiện hiệu suất của việc marking và
scanning các đối tượng nhỏ thông qua cải thiện locality và khả năng mở rộng CPU.
Kết quả benchmark thay đổi, nhưng chúng tôi kỳ vọng giảm từ 10 đến 40% chi phí
garbage collection trong các chương trình thực tế sử dụng nhiều garbage collector.
Cải thiện thêm, khoảng 10% chi phí garbage collection,
được kỳ vọng khi chạy trên các nền tảng CPU dựa trên amd64 mới hơn (Intel Ice
Lake hoặc AMD Zen 4 trở lên), vì garbage collector hiện tận dụng
các lệnh vector để scanning các đối tượng nhỏ khi có thể.

Garbage collector mới có thể bị tắt bằng cách đặt
`GOEXPERIMENT=nogreenteagc` tại thời điểm build.
Cài đặt opt-out này dự kiến sẽ bị loại bỏ trong Go 1.27.
Nếu bạn tắt garbage collector mới vì bất kỳ lý do nào liên quan đến
hiệu suất hoặc hành vi của nó, vui lòng [gửi issue](/issue/new).

### Lệnh gọi cgo nhanh hơn

<!-- CL 646198 -->

Chi phí runtime cơ bản của các lệnh gọi cgo đã giảm khoảng 30%.

### Ngẫu nhiên hóa địa chỉ cơ sở heap

<!-- CL 674835 -->

Trên các nền tảng 64-bit, runtime hiện ngẫu nhiên hóa địa chỉ cơ sở heap
khi khởi động.
Đây là một cải tiến bảo mật giúp kẻ tấn công khó dự đoán địa chỉ bộ nhớ
và khai thác lỗ hổng khi sử dụng cgo hơn.
Tính năng này có thể bị tắt bằng cách đặt
`GOEXPERIMENT=norandomizedheapbase64` tại thời điểm build.
Cài đặt opt-out này dự kiến sẽ bị loại bỏ trong một bản phát hành Go tương lai.

### Profile rò rỉ goroutine thử nghiệm {#goroutineleak-profiles}

<!-- CL 688335 -->

Một kiểu profile mới báo cáo goroutine bị rò rỉ hiện có sẵn như một
thử nghiệm.
Kiểu profile mới, có tên `goroutineleak` trong gói
[`runtime/pprof`](/pkg/runtime/pprof), có thể được bật bằng cách đặt
`GOEXPERIMENT=goroutineleakprofile` tại thời điểm build.
Bật thử nghiệm cũng làm cho profile có sẵn như một
endpoint của [`net/http/pprof`](/pkg/net/http/pprof),
`/debug/pprof/goroutineleak`.

Một goroutine *bị rò rỉ* là goroutine bị chặn trên một nguyên thủy đồng thời
(channel, [`sync.Mutex`](/pkg/sync#Mutex), [`sync.Cond`](/pkg/sync#Cond), v.v.) mà
không thể bị bỏ chặn.
Runtime phát hiện goroutine bị rò rỉ bằng garbage collector: nếu
goroutine G bị chặn trên nguyên thủy đồng thời P, và P không thể được truy cập từ
bất kỳ goroutine có thể chạy nào hoặc bất kỳ goroutine nào mà *những goroutine đó* có thể bỏ chặn, thì P
không thể bị bỏ chặn, vì vậy goroutine G không bao giờ có thể thức dậy.
Mặc dù không thể phát hiện tất cả các goroutine bị chặn vĩnh viễn trong mọi trường hợp,
phương pháp này phát hiện một lớp lớn các rò rỉ như vậy.

Ví dụ sau đây trình bày một rò rỉ goroutine trong thực tế
có thể được tiết lộ bởi profile mới:

```go
type result struct {
	res workResult
	err error
}

func processWorkItems(ws []workItem) ([]workResult, error) {
	// Process work items in parallel, aggregating results in ch.
	ch := make(chan result)
	for _, w := range ws {
		go func() {
			res, err := processWorkItem(w)
			ch <- result{res, err}
		}()
	}

	// Collect the results from ch, or return an error if one is found.
	var results []workResult
	for range len(ws) {
		r := <-ch
		if r.err != nil {
			// This early return may cause goroutine leaks.
			return nil, r.err
		}
		results = append(results, r.res)
	}
	return results, nil
}
```

Vì `ch` không có buffer, nếu `processWorkItems` trả về sớm do
lỗi, tất cả các goroutine `processWorkItem` còn lại sẽ bị rò rỉ.
Ngay sau khi điều này xảy ra, `ch` sẽ trở nên không thể truy cập đối với tất cả các
goroutine khác không tham gia vào rò rỉ, cho phép runtime phát hiện
và báo cáo các goroutine bị rò rỉ.

Vì kỹ thuật này dựa trên khả năng truy cập, runtime có thể không xác định được
các rò rỉ gây ra bởi việc chặn trên các nguyên thủy đồng thời có thể truy cập qua các biến
toàn cục hoặc các biến cục bộ của các goroutine có thể chạy.

Đặc biệt cảm ơn Vlad Saioc tại Uber đã đóng góp công việc này.
Lý thuyết nền tảng được trình bày chi tiết trong [một bài đăng của
Saioc et al](https://dl.acm.org/doi/pdf/10.1145/3676641.3715990).

Việc triển khai đã sẵn sàng cho môi trường sản xuất, và chỉ được coi là
thử nghiệm để thu thập phản hồi về API,
cụ thể là lựa chọn biến nó thành một profile mới.
Tính năng cũng được thiết kế để không phát sinh bất kỳ chi phí runtime bổ sung nào
trừ khi nó đang được sử dụng chủ động.

Chúng tôi khuyến khích người dùng thử nghiệm tính năng mới trong [Go
playground](/play/p/3C71z4Dpav-?v=gotip),
trong các kiểm thử, trong tích hợp liên tục và trong môi trường sản xuất.
Chúng tôi hoan nghênh phản hồi bổ sung trên [issue
đề xuất](/issue/74609).

Chúng tôi mục tiêu bật goroutine leak profile theo mặc định trong Go 1.27.

## Trình biên dịch {#compiler}

<!-- CLs 707755, 722440 -->

Trình biên dịch hiện có thể phân bổ bộ nhớ hỗ trợ cho slice trên stack trong nhiều
tình huống hơn, giúp cải thiện hiệu suất. Nếu thay đổi này gây ra sự cố, công cụ
[bisect](https://pkg.go.dev/golang.org/x/tools/cmd/bisect) có thể được dùng để
tìm phân bổ gây ra sự cố bằng cờ `-compile=variablemake`. Tất cả
các phân bổ stack mới như vậy cũng có thể được tắt bằng cách dùng
`-gcflags=all=-d=variablemakehash=n`.
Nếu bạn gặp vấn đề với tối ưu hóa này, vui lòng [gửi issue](/issue/new).

## Linker {#linker}

Trên Windows dựa trên ARM 64-bit (cổng `windows/arm64`), linker hiện hỗ trợ chế độ
liên kết nội bộ (internal linking mode) của các chương trình cgo, có thể được yêu cầu bằng cờ
`-ldflags=-linkmode=internal`.

Có một số thay đổi nhỏ đối với các tệp thực thi. Những thay đổi này
không ảnh hưởng đến việc chạy các chương trình Go. Chúng có thể ảnh hưởng đến các chương trình
phân tích binary Go, và có thể ảnh hưởng đến những người dùng chế độ liên kết ngoài
với script linker tùy chỉnh.

 - Cấu trúc `moduledata` hiện nằm trong phần riêng của nó, có tên
   `.go.module`.
 - Trường `cutab` của `moduledata`, là một slice, hiện có
   độ dài đúng; trước đây độ dài bị gấp bốn lần.
 - `pcHeader` được tìm thấy ở đầu phần `.gopclntab`
   không còn ghi lại điểm bắt đầu của phần text. Trường đó hiện
   luôn là không.
 - Thay đổi `pcHeader` đó được thực hiện để phần `.gopclntab`
   không còn chứa bất kỳ relocation nào. Trên các nền tảng hỗ trợ
   relro, phần đã chuyển từ segment relro sang segment
   rodata.
 - Các ký hiệu funcdata và findfunctab đã chuyển từ phần
   `.rodata` sang phần `.gopclntab`.
 - Phần `.gosymtab` đã bị xóa. Trước đây nó luôn
   có mặt nhưng rỗng.
 - Khi sử dụng liên kết nội bộ, các phần ELF hiện xuất hiện trong
   danh sách section header được sắp xếp theo địa chỉ. Thứ tự trước đây
   khá khó đoán.

Các tham chiếu đến tên phần ở đây dùng tên ELF như được thấy trên
Linux và các hệ thống khác. Tên Mach-O như được thấy trên Darwin bắt đầu bằng
dấu gạch dưới kép và không chứa bất kỳ dấu chấm nào.

## Bootstrap {#bootstrap}

<!-- go.dev/issue/69315 -->
Như đã đề cập trong [ghi chú phát hành Go 1.24](/doc/go1.24#bootstrap), Go 1.26 hiện yêu cầu
Go 1.24.6 trở lên để bootstrap.
Chúng tôi kỳ vọng rằng Go 1.28 sẽ yêu cầu một bản phát hành nhỏ của Go 1.26 trở lên để bootstrap.

## Thư viện chuẩn {#library}

### Gói crypto/hpke mới

Gói mới [`crypto/hpke`](/pkg/crypto/hpke) triển khai Hybrid Public Key Encryption
(HPKE) theo đặc tả trong [RFC 9180](https://rfc-editor.org/rfc/rfc9180.html), bao gồm hỗ trợ cho các
hybrid KEM hậu lượng tử (post-quantum hybrid KEM).

### Gói simd/archsimd thử nghiệm mới {#simd}

Go 1.26 giới thiệu gói thử nghiệm mới [`simd/archsimd`](/pkg/simd/archsimd/),
có thể được bật bằng cách đặt biến môi trường
`GOEXPERIMENT=simd` tại thời điểm build.
Gói này cung cấp quyền truy cập vào các thao tác SIMD đặc thù kiến trúc.
Hiện có sẵn trên kiến trúc amd64 và hỗ trợ
các kiểu vector 128-bit, 256-bit và 512-bit, chẳng hạn như
[`Int8x16`](/pkg/simd/archsimd#Int8x16) và
[`Float64x8`](/pkg/simd/archsimd#Float64x8), với các thao tác như
[`Int8x16.Add`](/pkg/simd/archsimd#Int8x16.Add).
API chưa được coi là ổn định.

Chúng tôi có kế hoạch hỗ trợ các kiến trúc khác trong các phiên bản tương lai, nhưng
API có chủ ý đặc thù kiến trúc và do đó không di động.
Ngoài ra, chúng tôi có kế hoạch phát triển một gói SIMD di động cấp cao
trong tương lai.

Xem [tài liệu gói](/pkg/simd/archsimd) và [issue đề xuất](/issue/73787) để biết thêm chi tiết.

### Gói runtime/secret thử nghiệm mới

<!-- https://go.dev/issue/21865 --->

Gói mới [`runtime/secret`](/pkg/runtime/secret) có sẵn như một thử nghiệm,
có thể được bật bằng cách đặt biến môi trường
`GOEXPERIMENT=runtimesecret` tại thời điểm build.
Nó cung cấp phương tiện để xóa an toàn các biến tạm thời được dùng trong
code thao tác thông tin bí mật (thường là mã hóa), chẳng hạn như
thanh ghi, stack, phân bổ heap mới.
Gói này nhằm giúp dễ dàng đảm bảo [bảo mật tiếp theo
(forward secrecy)](https://en.wikipedia.org/wiki/Forward_secrecy).
Hiện hỗ trợ kiến trúc amd64 và arm64 trên Linux.

### Các thay đổi nhỏ trong thư viện {#minor_library_changes}

#### [`bytes`](/pkg/bytes/)

Phương thức mới [`Buffer.Peek`](/pkg/bytes#Buffer.Peek) trả về n byte tiếp theo từ buffer mà không
tiến buffer.

#### [`crypto`](/pkg/crypto/)

Các interface mới [`Encapsulator`](/pkg/crypto#Encapsulator) và [`Decapsulator`](/pkg/crypto#Decapsulator) cho phép chấp nhận
khóa đóng gói (encapsulation) hoặc giải đóng gói (decapsulation) KEM trừu tượng.

#### [`crypto/dsa`](/pkg/crypto/dsa/)

Tham số random trong [`GenerateKey`](/pkg/crypto/dsa#GenerateKey) hiện bị bỏ qua.
Thay vào đó, nó hiện luôn dùng nguồn byte ngẫu nhiên bảo mật về mặt mã hóa.
Để kiểm thử xác định, dùng hàm mới [`testing/cryptotest.SetGlobalRandom`](/pkg/testing/cryptotest#SetGlobalRandom).
Cài đặt GODEBUG mới `cryptocustomrand=1` tạm thời khôi phục hành vi cũ.

#### [`crypto/ecdh`](/pkg/crypto/ecdh/)

Tham số random trong [`Curve.GenerateKey`](/pkg/crypto/ecdh#Curve.GenerateKey) hiện bị bỏ qua.
Thay vào đó, nó hiện luôn dùng nguồn byte ngẫu nhiên bảo mật về mặt mã hóa.
Để kiểm thử xác định, dùng hàm mới [`testing/cryptotest.SetGlobalRandom`](/pkg/testing/cryptotest#SetGlobalRandom).
Cài đặt GODEBUG mới `cryptocustomrand=1` tạm thời khôi phục hành vi cũ.

Interface mới [`KeyExchanger`](/pkg/crypto/ecdh#KeyExchanger), được triển khai bởi [`PrivateKey`](/pkg/crypto/ecdh#PrivateKey), giúp có thể
chấp nhận các khóa riêng ECDH trừu tượng, ví dụ: các khóa được triển khai trong phần cứng.

#### [`crypto/ecdsa`](/pkg/crypto/ecdsa/)

Các trường `big.Int` của [`PublicKey`](/pkg/crypto/ecdsa#PublicKey) và [`PrivateKey`](/pkg/crypto/ecdsa#PrivateKey) hiện không còn được khuyến nghị (deprecated).

Tham số random trong [`GenerateKey`](/pkg/crypto/ecdsa#GenerateKey), [`SignASN1`](/pkg/crypto/ecdsa#SignASN1), [`Sign`](/pkg/crypto/ecdsa#Sign), và [`PrivateKey.Sign`](/pkg/crypto/ecdsa#PrivateKey.Sign) hiện bị bỏ qua.
Thay vào đó, chúng hiện luôn dùng nguồn byte ngẫu nhiên bảo mật về mặt mã hóa.
Để kiểm thử xác định, dùng hàm mới [`testing/cryptotest.SetGlobalRandom`](/pkg/testing/cryptotest#SetGlobalRandom).
Cài đặt GODEBUG mới `cryptocustomrand=1` tạm thời khôi phục hành vi cũ.

#### [`crypto/ed25519`](/pkg/crypto/ed25519/)

Nếu tham số random trong [`GenerateKey`](/pkg/crypto/ed25519#GenerateKey) là nil, GenerateKey hiện luôn dùng
nguồn byte ngẫu nhiên bảo mật về mặt mã hóa, thay vì [`crypto/rand.Reader`](/pkg/crypto/rand#Reader)
(có thể đã bị ghi đè). Cài đặt GODEBUG mới `cryptocustomrand=1`
tạm thời khôi phục hành vi cũ.

#### [`crypto/fips140`](/pkg/crypto/fips140/)

[Module Mã hóa Go FIPS 140-3](/doc/security/fips140) v1.26.0 bao gồm các thay đổi được thực hiện đối với các gói `crypto/internal/fips140/...` đến bản phát hành này, và hiện có thể được chọn bằng GOFIPS140.

Các hàm mới [`WithoutEnforcement`](/pkg/crypto/fips140#WithoutEnforcement) và [`Enforced`](/pkg/crypto/fips140#Enforced) hiện cho phép chạy
trong chế độ `GODEBUG=fips140=only` đồng thời vô hiệu hóa có chọn lọc các kiểm tra FIPS 140-3 nghiêm ngặt.

[`Version`](/pkg/crypto/fips140#Version) trả về phiên bản Module Mã hóa Go FIPS 140-3 đã giải quyết khi build với module đông lạnh bằng GOFIPS140.

#### [`crypto/mlkem`](/pkg/crypto/mlkem/)

Các phương thức mới [`DecapsulationKey768.Encapsulator`](/pkg/crypto/mlkem#DecapsulationKey768.Encapsulator) và
[`DecapsulationKey1024.Encapsulator`](/pkg/crypto/mlkem#DecapsulationKey1024.Encapsulator) triển khai interface
[`crypto.Decapsulator`](/pkg/crypto#Decapsulator) mới.

Các thao tác đóng gói và giải đóng gói hiện nhanh hơn khoảng 18%.

#### [`crypto/mlkem/mlkemtest`](/pkg/crypto/mlkem/mlkemtest/)

Gói mới [`crypto/mlkem/mlkemtest`](/pkg/crypto/mlkem/mlkemtest) cung cấp các hàm [`Encapsulate768`](/pkg/crypto/mlkem/mlkemtest#Encapsulate768) và
[`Encapsulate1024`](/pkg/crypto/mlkem/mlkemtest#Encapsulate1024) triển khai đóng gói ML-KEM không ngẫu nhiên (derandomized),
để dùng với các kiểm thử đáp án đã biết (known-answer test).

#### [`crypto/rand`](/pkg/crypto/rand/)

Tham số random trong [`Prime`](/pkg/crypto/rand#Prime) hiện bị bỏ qua.
Thay vào đó, nó hiện luôn dùng nguồn byte ngẫu nhiên bảo mật về mặt mã hóa.
Để kiểm thử xác định, dùng hàm mới [`testing/cryptotest.SetGlobalRandom`](/pkg/testing/cryptotest#SetGlobalRandom).
Cài đặt GODEBUG mới `cryptocustomrand=1` tạm thời khôi phục hành vi cũ.

#### [`crypto/rsa`](/pkg/crypto/rsa/)

Hàm mới [`EncryptOAEPWithOptions`](/pkg/crypto/rsa#EncryptOAEPWithOptions) cho phép chỉ định các hàm hash khác nhau
cho padding OAEP và tạo mặt nạ MGF1.

Tham số random trong [`GenerateKey`](/pkg/crypto/rsa#GenerateKey), [`GenerateMultiPrimeKey`](/pkg/crypto/rsa#GenerateMultiPrimeKey), và [`EncryptPKCS1v15`](/pkg/crypto/rsa#EncryptPKCS1v15) hiện bị bỏ qua.
Thay vào đó, chúng hiện luôn dùng nguồn byte ngẫu nhiên bảo mật về mặt mã hóa.
Để kiểm thử xác định, dùng hàm mới [`testing/cryptotest.SetGlobalRandom`](/pkg/testing/cryptotest#SetGlobalRandom).
Cài đặt GODEBUG mới `cryptocustomrand=1` tạm thời khôi phục hành vi cũ.

Nếu các trường [`PrivateKey`](/pkg/crypto/rsa#PrivateKey) bị sửa đổi sau khi gọi [`PrivateKey.Precompute`](/pkg/crypto/rsa#PrivateKey.Precompute),
[`PrivateKey.Validate`](/pkg/crypto/rsa#PrivateKey.Validate) hiện sẽ thất bại.

[`PrivateKey.D`](/pkg/crypto/rsa#PrivateKey.D) hiện được kiểm tra tính nhất quán với các giá trị được tính trước, ngay cả khi
nó không được sử dụng.

Padding mã hóa PKCS #1 v1.5 không an toàn (được triển khai bởi [`EncryptPKCS1v15`](/pkg/crypto/rsa#EncryptPKCS1v15),
[`DecryptPKCS1v15`](/pkg/crypto/rsa#DecryptPKCS1v15), và [`DecryptPKCS1v15SessionKey`](/pkg/crypto/rsa#DecryptPKCS1v15SessionKey)) hiện không còn được khuyến nghị (deprecated).

#### [`crypto/sha3`](/pkg/crypto/sha3/)

Giá trị không (zero value) của [`SHA3`](/pkg/crypto/sha3#SHA3) hiện là một instance SHA3-256 có thể dùng được, và giá trị không của [`SHAKE`](/pkg/crypto/sha3#SHAKE) hiện là một instance SHAKE256 có thể dùng được.

#### [`crypto/subtle`](/pkg/crypto/subtle)

Hàm [`WithDataIndependentTiming`](/pkg/crypto/subtle#WithDataIndependentTiming)
không còn khóa goroutine gọi vào luồng OS trong khi thực thi
hàm được truyền vào. Ngoài ra, bất kỳ goroutine nào được sinh ra trong quá trình
thực thi hàm được truyền vào và các goroutine con của chúng hiện kế thừa các thuộc tính của
`WithDataIndependentTiming` trong suốt vòng đời của chúng. Thay đổi này cũng ảnh hưởng đến cgo theo
các cách sau:

- Bất kỳ code C nào được gọi qua cgo từ bên trong hàm được truyền vào
  `WithDataIndependentTiming`, hoặc từ một goroutine được sinh ra bởi hàm được truyền vào
  `WithDataIndependentTiming` và các goroutine con của nó, cũng sẽ có
  data independent timing được bật trong suốt thời gian gọi. Nếu code C
  tắt data independent timing, nó sẽ được bật lại khi quay về Go.
- Nếu code C được gọi qua cgo, từ hàm được truyền vào
  `WithDataIndependentTiming` hoặc nơi khác, bật hoặc tắt data independent
  timing thì việc gọi vào Go sẽ giữ nguyên trạng thái đó trong suốt thời gian
  gọi.

#### [`crypto/tls`](/pkg/crypto/tls/)

Các trao đổi khóa hậu lượng tử hybrid [`SecP256r1MLKEM768`](/pkg/crypto/tls#SecP256r1MLKEM768) và [`SecP384r1MLKEM1024`](/pkg/crypto/tls#SecP384r1MLKEM1024)
hiện được bật mặc định. Chúng có thể bị tắt bằng cách đặt
[`Config.CurvePreferences`](/pkg/crypto/tls#Config.CurvePreferences) hoặc với cài đặt GODEBUG `tlssecpmlkem=0`.

Trường mới [`ClientHelloInfo.HelloRetryRequest`](/pkg/crypto/tls#ClientHelloInfo.HelloRetryRequest) cho biết liệu ClientHello
có được gửi để phản hồi tin nhắn HelloRetryRequest không. Trường mới
[`ConnectionState.HelloRetryRequest`](/pkg/crypto/tls#ConnectionState.HelloRetryRequest) cho biết liệu server
đã gửi HelloRetryRequest, hay client đã nhận HelloRetryRequest,
tùy thuộc vào vai trò kết nối.

Kiểu [`QUICConn`](/pkg/crypto/tls#QUICConn) được dùng bởi các triển khai QUIC bao gồm một sự kiện mới
để báo cáo lỗi TLS handshake.

Nếu [`Certificate.PrivateKey`](/pkg/crypto/tls#Certificate.PrivateKey) triển khai [`crypto.MessageSigner`](/pkg/crypto#MessageSigner), phương thức SignMessage của nó
được dùng thay vì Sign trong TLS 1.2 và mới hơn.

Các cài đặt GODEBUG sau được giới thiệu trong [Go 1.22](/doc/godebug#go-122)
và [Go 1.23](/doc/godebug#go-123) sẽ bị loại bỏ trong bản phát hành Go lớn tiếp theo.
Bắt đầu từ Go 1.27, hành vi mới sẽ áp dụng bất kể cài đặt GODEBUG hay phiên bản ngôn ngữ go.mod.

- `tlsunsafeekm`: [`ConnectionState.ExportKeyingMaterial`](/pkg/crypto/tls#ConnectionState.ExportKeyingMaterial) sẽ yêu cầu TLS 1.3 hoặc Extended Master Secret.
- `tlsrsakex`: các trao đổi khóa RSA-only cũ không có ECDH sẽ không được bật mặc định.
- `tls10server`: phiên bản TLS tối thiểu mặc định cho cả client và server sẽ là TLS 1.2.
- `tls3des`: các cipher suite mặc định sẽ không bao gồm 3DES.
- `x509keypairleaf`: [`X509KeyPair`](/pkg/crypto/tls#X509KeyPair) và [`LoadX509KeyPair`](/pkg/crypto/tls#LoadX509KeyPair) sẽ luôn điền trường [`Certificate.Leaf`](/pkg/crypto/tls#Certificate.Leaf).

#### [`crypto/x509`](/pkg/crypto/x509/)

Các kiểu [`ExtKeyUsage`](/pkg/crypto/x509#ExtKeyUsage) và [`KeyUsage`](/pkg/crypto/x509#KeyUsage) hiện có các phương thức `String` trả về
các tên OID tương ứng như được định nghĩa trong RFC 5280 và các registry khác.

Kiểu [`ExtKeyUsage`](/pkg/crypto/x509#ExtKeyUsage) hiện có một phương thức `OID` trả về OID tương ứng cho EKU.

Hàm mới [`OIDFromASN1OID`](/pkg/crypto/x509#OIDFromASN1OID) cho phép chuyển đổi một [`encoding/asn1.ObjectIdentifier`](/pkg/encoding/asn1#ObjectIdentifier) thành
một [`OID`](/pkg/crypto/x509#OID).

#### [`debug/elf`](/pkg/debug/elf/)

Các hằng số `R_LARCH_*` bổ sung từ [LoongArch ELF psABI v20250521](https://github.com/loongson/la-abi-specs/blob/v2.40/laelf.adoc)
(phiên bản toàn cục v2.40) được định nghĩa để dùng với các hệ thống LoongArch.

#### [`errors`](/pkg/errors/)

Hàm mới [`AsType`](/pkg/errors#AsType) là phiên bản generic của [`As`](/pkg/errors#As). Nó an toàn về kiểu, nhanh hơn,
và trong hầu hết các trường hợp, dễ sử dụng hơn.

#### [`fmt`](/pkg/fmt/)

<!-- go.dev/cl/708836 -->
Đối với các chuỗi không được định dạng, `fmt.Errorf("x")` hiện phân bổ ít hơn và nhìn chung khớp với
số phân bổ của `errors.New("x")`.

#### [`go/ast`](/pkg/go/ast/)

Hàm mới [`ParseDirective`](/pkg/go/ast#ParseDirective) phân tích cú pháp [comment
chỉ thị (directive comment)](/doc/comment#Syntax), là các comment như `//go:generate`.
Các công cụ mã nguồn có thể hỗ trợ các comment chỉ thị của riêng họ và API mới này
sẽ giúp họ triển khai cú pháp theo quy ước.

<!-- go.dev/issue/76395 -->
Trường mới [`BasicLit.ValueEnd`](/pkg/go/ast#BasicLit.ValueEnd) ghi lại vị trí kết thúc chính xác của
một literal để phương thức [`BasicLit.End`](/pkg/go/ast#BasicLit.End) hiện luôn có thể trả về
câu trả lời đúng. (Trước đây nó được tính bằng heuristic không chính xác
đối với các raw string literal đa dòng trong các tệp nguồn Windows,
do việc loại bỏ ký tự xuống dòng.)

Các chương trình cập nhật trường `ValuePos` của `BasicLit` được tạo bởi
parser có thể cần cập nhật hoặc xóa trường `ValueEnd` để
tránh sự khác biệt nhỏ trong output được định dạng.

#### [`go/token`](/pkg/go/token/)

Phương thức tiện ích mới [`File.End`](/pkg/go/token#File.End) trả về vị trí kết thúc của tệp.

#### [`go/types`](/pkg/go/types/)

Cài đặt GODEBUG `gotypesalias` được giới thiệu trong [Go 1.22](/doc/godebug#go-122)
sẽ bị loại bỏ trong bản phát hành Go lớn tiếp theo.
Bắt đầu từ Go 1.27, gói [`go/types`](/pkg/go/types) sẽ luôn tạo ra
[kiểu Alias](/pkg/go/types#Alias) để biểu diễn [type alias](/ref/spec#Type_declarations)
bất kể cài đặt GODEBUG hay phiên bản ngôn ngữ go.mod.

#### [`image/jpeg`](/pkg/image/jpeg/)

Bộ mã hóa và giải mã JPEG đã được thay thế bằng các triển khai mới, nhanh hơn và chính xác hơn.
Code mong đợi output bit-cho-bit cụ thể từ bộ mã hóa hoặc giải mã có thể cần được cập nhật.

#### [`io`](/pkg/io/)

<!-- go.dev/cl/722500 -->
[`ReadAll`](/pkg/io#ReadAll) hiện phân bổ ít bộ nhớ trung gian hơn và trả về một
slice cuối có kích thước tối thiểu. Nó thường nhanh hơn khoảng hai lần trong khi
thường phân bổ tổng cộng khoảng một nửa bộ nhớ, với lợi ích nhiều hơn cho các đầu vào lớn hơn.

#### [`log/slog`](/pkg/log/slog/)

Hàm [`NewMultiHandler`](/pkg/log/slog#NewMultiHandler) tạo một
[`MultiHandler`](/pkg/log/slog#MultiHandler) gọi tất cả các Handler đã cho.
Phương thức `Enabled` của nó báo cáo liệu có bất kỳ phương thức `Enabled` nào của các handler
trả về true không.
Các phương thức `Handle`, `WithAttrs` và `WithGroup` của nó gọi phương thức tương ứng
trên mỗi handler được bật.

#### [`net`](/pkg/net/)

Các phương thức mới của [`Dialer`](/pkg/net/#Dialer)
[`DialIP`](/pkg/net/#Dialer.DialIP),
[`DialTCP`](/pkg/net/#Dialer.DialTCP),
[`DialUDP`](/pkg/net/#Dialer.DialUDP), và
[`DialUnix`](/pkg/net/#Dialer.DialUnix)
cho phép dial các kiểu mạng cụ thể với giá trị context.

#### [`net/http`](/pkg/net/http/)

Trường mới
[`HTTP2Config.StrictMaxConcurrentRequests`](/pkg/net/http#HTTP2Config.StrictMaxConcurrentRequests)
kiểm soát liệu có nên mở một kết nối mới
khi kết nối HTTP/2 hiện có đã vượt quá giới hạn stream không.

Phương thức mới [`Transport.NewClientConn`](/pkg/net/http#Transport.NewClientConn) trả về một kết nối client
đến server HTTP.
Hầu hết người dùng nên tiếp tục dùng [`Transport.RoundTrip`](/pkg/net/http#Transport.RoundTrip) để thực hiện yêu cầu,
vốn quản lý một nhóm kết nối.
`NewClientConn` hữu ích cho người dùng cần triển khai quản lý kết nối của riêng họ.

[`Client`](/pkg/net/http#Client) hiện dùng và đặt cookie có phạm vi đến các URL với phần host khớp với
[`Request.Host`](/pkg/net/http#Request.Host) khi có sẵn.
Trước đây, host của địa chỉ kết nối luôn được dùng.
Các chuyển hướng dấu gạch chéo theo sau (trailing slash redirect) của [`ServeMux`](/pkg/net/http#ServeMux) hiện sử dụng HTTP status 307
(Temporary Redirect) thay vì 301 (Moved Permanently).

#### [`net/http/httptest`](/pkg/net/http/httptest/)

HTTP client được trả về bởi [`Server.Client`](/pkg/net/http/httptest#Server.Client) sẽ hiện chuyển hướng các yêu cầu cho
`example.com` và bất kỳ subdomain nào đến server đang được kiểm thử.

#### [`net/http/httputil`](/pkg/net/http/httputil/)

Trường cấu hình [`ReverseProxy.Director`](/pkg/net/http/httputil#ReverseProxy.Director) không còn được khuyến nghị (deprecated)
thay cho [`ReverseProxy.Rewrite`](/pkg/net/http/httputil#ReverseProxy.Rewrite).

Một client độc hại có thể xóa các header được thêm bởi hàm `Director`
bằng cách chỉ định các header đó là hop-by-hop. Vì không có cách nào để giải quyết
vấn đề này trong phạm vi API `Director`, chúng tôi đã thêm một hook
`Rewrite` mới trong Go 1.20. Các hook `Rewrite` được cung cấp cả
yêu cầu đến chưa được sửa đổi mà proxy nhận được và yêu cầu đi ra
sẽ được proxy gửi đi.

Vì hook `Director` về cơ bản là không an toàn, chúng tôi hiện đang không còn khuyến nghị (deprecate) nó.

#### [`net/netip`](/pkg/net/netip/)

Phương thức mới [`Prefix.Compare`](/pkg/net/netip#Prefix.Compare) so sánh hai prefix.

#### [`net/url`](/pkg/net/url/)

[`Parse`](/pkg/net/url#Parse) hiện từ chối các URL không đúng định dạng chứa dấu hai chấm trong subcomponent host,
chẳng hạn như `http://::1/` hoặc `http://localhost:80:80/`.
Các URL chứa địa chỉ IPv6 có dấu ngoặc, chẳng hạn như `http://[::1]/` vẫn được chấp nhận.
Cài đặt GODEBUG mới `urlstrictcolons=0` khôi phục hành vi cũ.

#### [`os`](/pkg/os/)

Phương thức mới [`Process.WithHandle`](/pkg/os#Process.WithHandle) cung cấp
quyền truy cập vào một handle tiến trình nội bộ trên các nền tảng được hỗ trợ (pidfd trên Linux 5.4
trở lên, Handle trên Windows).

Trên Windows, tham số `flag` của [`OpenFile`](/pkg/os#OpenFile) hiện có thể chứa bất kỳ tổ hợp nào của
các cờ tệp đặc thù Windows, chẳng hạn như `FILE_FLAG_OVERLAPPED` và
`FILE_FLAG_SEQUENTIAL_SCAN`, để kiểm soát hành vi caching tệp hoặc thiết bị,
chế độ truy cập và các cờ mục đích đặc biệt khác.

#### [`os/signal`](/pkg/os/signal/)

[`NotifyContext`](/pkg/os/signal#NotifyContext) hiện hủy context được trả về bằng [`context.CancelCauseFunc`](/pkg/context#CancelCauseFunc)
và một lỗi cho biết tín hiệu nào đã được nhận.

#### [`reflect`](/pkg/reflect/)

Các phương thức mới [`Type.Fields`](/pkg/reflect#Type.Fields),
[`Type.Methods`](/pkg/reflect#Type.Methods),
[`Type.Ins`](/pkg/reflect#Type.Ins)
và [`Type.Outs`](/pkg/reflect#Type.Outs)
trả về các iterator cho các trường của kiểu (đối với kiểu struct), phương thức,
tham số đầu vào và đầu ra (đối với kiểu hàm).

Tương tự, các phương thức mới [`Value.Fields`](/pkg/reflect#Value.Fields)
và [`Value.Methods`](/pkg/reflect#Value.Methods) trả về các iterator qua
các trường hoặc phương thức của một giá trị.
Mỗi lần lặp trả về thông tin kiểu ([`StructField`](/pkg/reflect#StructField) hoặc
[`Method`](/pkg/reflect#Method)) của một trường hoặc phương thức,
cùng với [`Value`](/pkg/reflect#Value) của trường hoặc phương thức đó.

#### [`runtime/metrics`](/pkg/runtime/metrics/)

Một số metric scheduler mới đã được thêm vào, bao gồm số lượng
goroutine ở các trạng thái khác nhau (đang chờ, có thể chạy, v.v.) dưới
prefix `/sched/goroutines`, số luồng OS mà runtime
biết đến với `/sched/threads:threads`, và tổng số
goroutine được tạo bởi chương trình với
`/sched/goroutines-created:goroutines`.

#### [`testing`](/pkg/testing/)

Các phương thức mới [`T.ArtifactDir`](/pkg/testing#T.ArtifactDir), [`B.ArtifactDir`](/pkg/testing#B.ArtifactDir), và [`F.ArtifactDir`](/pkg/testing#F.ArtifactDir)
trả về một thư mục để ghi các tệp output kiểm thử (artifact).

Khi cờ `-artifacts` được cung cấp cho `go test`,
thư mục này sẽ nằm dưới thư mục output
(được chỉ định bằng `-outputdir`, hoặc thư mục hiện tại theo mặc định).
Ngược lại, artifact được lưu trữ trong một thư mục tạm thời
sẽ bị xóa sau khi kiểm thử hoàn thành.

Lần gọi đầu tiên đến `ArtifactDir` khi `-artifacts` được cung cấp
ghi vị trí của thư mục vào log kiểm thử.

Ví dụ, trong một kiểm thử có tên `TestArtifacts`,
`t.ArtifactDir()` phát ra:

```
=== ARTIFACTS TestArtifacts /path/to/artifact/dir
```

<!-- go.dev/issue/73137 -->

Phương thức [`B.Loop`](/pkg/testing#B.Loop) không còn ngăn cản việc inline trong
thân vòng lặp, điều này có thể dẫn đến phân bổ không mong đợi và benchmark chậm hơn.
Với bản sửa lỗi này, chúng tôi kỳ vọng tất cả các benchmark có thể được chuyển đổi từ phong cách
[`B.N`](/pkg/testing#B) cũ sang phong cách `B.Loop` mới mà không có tác dụng phụ.
Trong thân vòng lặp `for b.Loop() { ... }`, các tham số lệnh gọi hàm, kết quả,
và các biến được gán vẫn được giữ sống, ngăn trình biên dịch
tối ưu hóa bỏ toàn bộ phần của benchmark.

#### [`testing/cryptotest`](/pkg/testing/cryptotest/)

Hàm mới [`SetGlobalRandom`](/pkg/testing/cryptotest#SetGlobalRandom) cấu hình một nguồn ngẫu nhiên mã hóa xác định toàn cục
trong suốt thời gian kiểm thử. Nó ảnh hưởng đến
`crypto/rand`, và tất cả các nguồn ngẫu nhiên mã hóa ngầm trong các gói
`crypto/...`.

#### [`time`](/pkg/time/)

Cài đặt GODEBUG `asynctimerchan` được giới thiệu trong [Go 1.23](/doc/godebug#go-123)
sẽ bị loại bỏ trong bản phát hành Go lớn tiếp theo.
Bắt đầu từ Go 1.27, gói [`time`](/pkg/time) sẽ luôn dùng các channel
không có buffer (đồng bộ) cho timer bất kể cài đặt GODEBUG hay phiên bản ngôn ngữ go.mod.

## Nền tảng {#ports}

### Darwin

<!-- go.dev/issue/75836 -->

Go 1.26 là bản phát hành cuối cùng chạy được trên macOS 12 Monterey. Go 1.27 sẽ yêu cầu macOS 13 Ventura trở lên.

### FreeBSD

<!-- go.dev/issue/76475 -->

Cổng freebsd/riscv64 (`GOOS=freebsd GOARCH=riscv64`) đã được đánh dấu là [bị hỏng](/wiki/PortingPolicy#broken-ports).
Xem [issue 76475](/issue/76475) để biết chi tiết.

### Windows

<!-- go.dev/issue/71671 -->

Như đã [thông báo](/doc/go1.25#windows) trong ghi chú phát hành Go 1.25, cổng 32-bit windows/arm
[đã bị hỏng](/doc/go1.24#windows) (`GOOS=windows`
`GOARCH=arm`) đã bị loại bỏ.

### PowerPC

<!-- go.dev/issue/76244 -->

Go 1.26 là bản phát hành cuối cùng hỗ trợ ELFv1 ABI trên cổng PowerPC 64-bit big-endian
trên Linux (`GOOS=linux` `GOARCH=ppc64`).
Nó sẽ chuyển sang ELFv2 ABI trong Go 1.27.
Vì cổng hiện không hỗ trợ liên kết với các đối tượng ELF khác,
chúng tôi kỳ vọng thay đổi này sẽ trong suốt đối với người dùng.

### RISC-V

<!-- CL 690497 -->

Cổng `linux/riscv64` hiện hỗ trợ trình phát hiện race.

### S390X

<!-- CL 719482 -->

Cổng `s390x` hiện hỗ trợ truyền đối số và kết quả hàm bằng thanh ghi.

### WebAssembly {#wasm}

<!-- CL 707855 -->

Trình biên dịch hiện vô điều kiện sử dụng các lệnh mở rộng dấu (sign extension)
và chuyển đổi dấu phẩy động sang số nguyên không bẫy lỗi (non-trapping floating-point to integer conversion).
Các tính năng này đã được chuẩn hóa từ ít nhất Wasm 2.0.
Các cài đặt `GOWASM` tương ứng, `signext` và `satconv`, hiện bị bỏ qua.

<!-- CL 683296 -->

Đối với các ứng dụng WebAssembly, runtime hiện quản lý các khối
bộ nhớ heap theo các gia số nhỏ hơn nhiều, dẫn đến
giảm đáng kể mức sử dụng bộ nhớ cho các ứng dụng có heap nhỏ hơn khoảng
16 MiB.
