---
template: true
title: Ghi chú phát hành Go 1
---

## Giới thiệu về Go 1 {#introduction}

Go phiên bản 1, gọi tắt là Go 1, định nghĩa một ngôn ngữ và một bộ thư viện cốt lõi
cung cấp nền tảng ổn định để xây dựng các sản phẩm, dự án và ấn phẩm đáng tin cậy.

Động lực chính của Go 1 là sự ổn định cho người dùng. Mọi người có thể
viết chương trình Go và kỳ vọng rằng chúng sẽ tiếp tục biên dịch và chạy mà không cần thay đổi,
trong nhiều năm, kể cả trong các môi trường production như Google App Engine.
Tương tự, mọi người có thể viết sách về Go, nêu rõ phiên bản Go mà cuốn sách
mô tả, và số phiên bản đó vẫn còn có ý nghĩa sau nhiều năm.

Mã nguồn biên dịch được trong Go 1 sẽ, với một vài ngoại lệ, tiếp tục biên dịch và
chạy suốt vòng đời của phiên bản đó, ngay cả khi chúng tôi phát hành cập nhật và sửa lỗi
như Go phiên bản 1.1, 1.2 và tiếp theo. Ngoài những sửa lỗi nghiêm trọng, các thay đổi
đối với ngôn ngữ và thư viện trong các bản phát hành tiếp theo của Go 1 có thể
bổ sung chức năng nhưng sẽ không làm hỏng các chương trình Go 1 hiện có.
[Tài liệu tương thích Go 1](go1compat.html)
giải thích chi tiết hơn về các hướng dẫn tương thích.

Go 1 là biểu hiện của Go như hiện đang được sử dụng, không phải là việc xem xét lại toàn bộ
ngôn ngữ. Chúng tôi tránh thiết kế các tính năng mới mà thay vào đó tập trung vào việc khắc phục
các vấn đề, sự không nhất quán và cải thiện tính di động. Có một số
thay đổi đối với ngôn ngữ và các gói Go mà chúng tôi đã cân nhắc trong một thời gian và
đã làm mẫu thử nhưng chưa phát hành, chủ yếu vì chúng đáng kể và
không tương thích ngược. Go 1 là cơ hội để đưa chúng ra, điều này
có lợi cho dài hạn, nhưng cũng có nghĩa là Go 1 giới thiệu sự không tương thích
với các chương trình cũ. May mắn thay, công cụ `go` `fix` có thể
tự động hóa phần lớn công việc cần thiết để đưa các chương trình lên chuẩn Go 1.

Tài liệu này phác thảo những thay đổi chính trong Go 1 sẽ ảnh hưởng đến các lập trình viên
đang cập nhật mã nguồn hiện có; điểm tham chiếu của nó là bản phát hành trước đó, r60 (được gắn thẻ là
r60.3). Nó cũng giải thích cách cập nhật mã nguồn từ r60 để chạy trên Go 1.

## Thay đổi về ngôn ngữ {#language}

### Append {#append}

Hàm variadic tiền khai báo `append` giúp dễ dàng mở rộng một slice
bằng cách thêm các phần tử vào cuối.
Một cách dùng phổ biến là thêm byte vào cuối một byte slice khi tạo đầu ra.
Tuy nhiên, `append` không cung cấp cách để nối một chuỗi vào một `[]byte`,
đây là một trường hợp phổ biến khác.

{{code "/doc/progs/go1.go" `/greeting := ..byte/` `/append.*hello/`}}

Tương tự với tính năng tương tự của `copy`, Go 1
cho phép một chuỗi được nối (theo từng byte) trực tiếp vào một
byte slice, giảm ma sát giữa chuỗi và byte slice.
Phép chuyển đổi không còn cần thiết nữa:

{{code "/doc/progs/go1.go" `/append.*world/`}}

_Cập nhật_:
Đây là tính năng mới, vì vậy mã nguồn hiện có không cần thay đổi.

### Close {#close}

Hàm tiền khai báo `close` cung cấp cơ chế
để một bên gửi báo hiệu rằng sẽ không còn gửi thêm giá trị nào nữa.
Nó quan trọng đối với việc triển khai các vòng lặp `for` `range`
trên channel và hữu ích trong các tình huống khác.
Một phần do thiết kế và một phần do điều kiện race có thể xảy ra,
nó chỉ được dùng bởi goroutine gửi trên channel,
không phải goroutine nhận dữ liệu.
Tuy nhiên, trước Go 1 không có kiểm tra lúc biên dịch rằng `close`
đang được dùng đúng cách.

Để thu hẹp khoảng cách này, ít nhất là một phần, Go 1 không cho phép `close` trên các channel chỉ nhận.
Cố gắng đóng một channel như vậy là lỗi lúc biên dịch.

{{code "/doc/progs/go1.go" `/STARTCLOSE/` `/ENDCLOSE/`}}

_Cập nhật_:
Mã nguồn hiện có cố gắng đóng một channel chỉ nhận đã là lỗi
ngay cả trước Go 1 và cần được sửa. Trình biên dịch sẽ
từ chối mã nguồn như vậy.

### Composite literals {#literals}

Trong Go 1, một composite literal của kiểu mảng, slice hoặc map có thể bỏ qua
đặc tả kiểu cho các bộ khởi tạo phần tử nếu chúng là kiểu con trỏ.
Cả bốn khởi tạo trong ví dụ này đều hợp lệ; cái cuối cùng không hợp lệ trước Go 1.

{{code "/doc/progs/go1.go" `/type Date struct/` `/STOP/`}}

_Cập nhật_:
Thay đổi này không ảnh hưởng đến mã nguồn hiện có, nhưng lệnh
`gofmt` `-s` áp dụng lên mã nguồn hiện có
sẽ, trong số các việc khác, bỏ qua các kiểu phần tử tường minh ở bất cứ đâu được phép.

### Goroutine trong init {#init}

Ngôn ngữ cũ định nghĩa rằng các câu lệnh `go` được thực thi trong quá trình khởi tạo tạo ra goroutine nhưng chúng không bắt đầu chạy cho đến khi quá trình khởi tạo của toàn bộ chương trình hoàn tất.
Điều này tạo ra sự vụng về ở nhiều nơi và, trên thực tế, hạn chế tính hữu ích
của cấu trúc `init`:
nếu một gói khác có thể sử dụng thư viện trong quá trình khởi tạo, thư viện
buộc phải tránh dùng goroutine.
Thiết kế này được thực hiện vì lý do đơn giản và an toàn nhưng,
khi sự tin tưởng của chúng tôi vào ngôn ngữ tăng lên, điều đó có vẻ không cần thiết.
Chạy goroutine trong quá trình khởi tạo không phức tạp hay không an toàn hơn chạy chúng trong quá trình thực thi bình thường.

Trong Go 1, mã nguồn sử dụng goroutine có thể được gọi từ
các hàm `init` và các biểu thức khởi tạo toàn cục
mà không gây ra deadlock.

{{code "/doc/progs/go1.go" `/PackageGlobal/` `/^}/`}}

_Cập nhật_:
Đây là tính năng mới, vì vậy mã nguồn hiện có không cần thay đổi,
mặc dù có thể mã nguồn phụ thuộc vào việc goroutine không khởi động trước `main` sẽ bị hỏng.
Không có mã nguồn như vậy trong kho lưu trữ chuẩn.

### Kiểu rune {#rune}

Đặc tả ngôn ngữ cho phép kiểu `int` rộng 32 hoặc 64 bit, nhưng các triển khai hiện tại đặt `int` thành 32 bit ngay cả trên các nền tảng 64 bit.
Sẽ tốt hơn nếu `int` là 64 bit trên các nền tảng 64 bit.
(Có những hệ quả quan trọng đối với việc đánh chỉ số các slice lớn.)
Tuy nhiên, thay đổi này sẽ lãng phí không gian khi xử lý các ký tự Unicode với
ngôn ngữ cũ vì kiểu `int` cũng được dùng để giữ các điểm mã Unicode: mỗi điểm mã sẽ lãng phí thêm 32 bit lưu trữ nếu `int` tăng từ 32 bit lên 64.

Để làm cho việc chuyển sang `int` 64 bit khả thi,
Go 1 giới thiệu một kiểu cơ bản mới, `rune`, để biểu diễn
các điểm mã Unicode riêng lẻ.
Nó là bí danh cho `int32`, tương tự như `byte`
là bí danh cho `uint8`.

Các ký tự literal như `'a'`, `'語'` và `'\u0345'`
bây giờ có kiểu mặc định là `rune`,
tương tự như `1.0` có kiểu mặc định là `float64`.
Do đó, một biến được khởi tạo bằng một hằng ký tự sẽ
có kiểu `rune` trừ khi được chỉ định khác.

Các thư viện đã được cập nhật để dùng `rune` thay vì `int`
khi phù hợp. Ví dụ, các hàm `unicode.ToLower` và
các hàm liên quan bây giờ nhận và trả về kiểu `rune`.

{{code "/doc/progs/go1.go" `/STARTRUNE/` `/ENDRUNE/`}}

_Cập nhật_:
Hầu hết mã nguồn sẽ không bị ảnh hưởng vì suy luận kiểu từ
các bộ khởi tạo `:=` đưa vào kiểu mới một cách im lặng, và nó lan truyền
từ đó.
Một số mã nguồn có thể gặp lỗi kiểu mà một phép chuyển đổi đơn giản sẽ giải quyết.

### Kiểu error {#error}

Go 1 giới thiệu kiểu built-in mới, `error`, có định nghĩa như sau:

	    type error interface {
	        Error() string
	    }

Vì hệ quả của kiểu này đều nằm trong thư viện gói,
nó được thảo luận [bên dưới](#errors).

### Xóa khỏi map {#delete}

Trong ngôn ngữ cũ, để xóa mục có khóa `k` từ map `m`, người ta viết câu lệnh,

	    m[k] = value, false

Cú pháp này là một trường hợp đặc biệt kỳ lạ, phép gán hai-vào-một duy nhất.
Nó yêu cầu truyền một giá trị (thường bị bỏ qua) được đánh giá nhưng bị loại bỏ,
cộng với một boolean gần như luôn là hằng số `false`.
Nó hoạt động nhưng kỳ lạ và là điểm tranh cãi.

Trong Go 1, cú pháp đó đã biến mất; thay vào đó có một
hàm built-in mới, `delete`. Lệnh gọi

{{code "/doc/progs/go1.go" `/delete\(m, k\)/`}}

sẽ xóa mục map được lấy bởi biểu thức `m[k]`.
Không có giá trị trả về. Xóa một mục không tồn tại là no-op.

_Cập nhật_:
Chạy `go` `fix` sẽ chuyển đổi các biểu thức dạng `m[k] = value,
false` thành `delete(m, k)` khi rõ ràng là
giá trị bị bỏ qua có thể được loại bỏ một cách an toàn khỏi chương trình và
`false` tham chiếu đến hằng boolean được định nghĩa sẵn.
Công cụ fix sẽ đánh dấu các cách dùng khác của cú pháp này để lập trình viên kiểm tra.

### Duyệt map {#iteration}

Đặc tả ngôn ngữ cũ không định nghĩa thứ tự duyệt đối với map,
và trên thực tế nó khác nhau trên các nền tảng phần cứng.
Điều này làm cho các test duyệt qua map dễ bị hỏng và không di động, với
tính chất khó chịu là một test có thể luôn vượt qua trên một máy nhưng bị hỏng trên máy khác.

Trong Go 1, thứ tự các phần tử được thăm khi duyệt
qua một map bằng câu lệnh `for` `range`
được định nghĩa là không thể đoán trước, ngay cả khi cùng một vòng lặp được chạy nhiều
lần với cùng một map.
Mã nguồn không nên giả định rằng các phần tử được thăm theo bất kỳ thứ tự cụ thể nào.

Thay đổi này có nghĩa là mã nguồn phụ thuộc vào thứ tự duyệt rất có thể sẽ bị hỏng sớm và được sửa trước khi trở thành vấn đề.
Quan trọng không kém, nó cho phép triển khai map đảm bảo cân bằng map tốt hơn ngay cả khi chương trình sử dụng vòng lặp range để chọn một phần tử từ map.

{{code "/doc/progs/go1.go" `/Sunday/` `/^	}/`}}

_Cập nhật_:
Đây là một thay đổi mà các công cụ không thể giúp đỡ. Hầu hết mã nguồn hiện có
sẽ không bị ảnh hưởng, nhưng một số chương trình có thể bị hỏng hoặc hoạt động sai; chúng tôi
khuyến nghị kiểm tra thủ công tất cả các câu lệnh range trên map để
xác minh chúng không phụ thuộc vào thứ tự duyệt. Có một vài ví dụ như vậy
trong kho lưu trữ chuẩn; chúng đã được sửa.
Lưu ý rằng việc phụ thuộc vào thứ tự duyệt đã là sai ngay từ đầu, vì
thứ tự đó không được chỉ định. Thay đổi này quy định hóa tính không thể đoán trước.

### Phép gán nhiều giá trị {#multiple_assignment}

Đặc tả ngôn ngữ từ lâu đã đảm bảo rằng trong các phép gán
tất cả các biểu thức vế phải đều được đánh giá trước khi bất kỳ biểu thức vế trái nào được gán.
Để đảm bảo hành vi có thể dự đoán,
Go 1 tinh chỉnh thêm đặc tả.

Nếu vế trái của câu lệnh gán
chứa các biểu thức yêu cầu đánh giá, chẳng hạn như
các lệnh gọi hàm hoặc thao tác đánh chỉ số mảng, chúng đều sẽ được thực hiện
theo quy tắc từ trái sang phải thông thường trước khi bất kỳ biến nào được gán
giá trị của chúng. Sau khi mọi thứ được đánh giá, các phép gán thực tế
diễn ra theo thứ tự từ trái sang phải.

Các ví dụ này minh họa hành vi.

{{code "/doc/progs/go1.go" `/sa :=/` `/then sc.0. = 2/`}}

_Cập nhật_:
Đây là một thay đổi mà các công cụ không thể giúp đỡ, nhưng khả năng hỏng là thấp.
Không có mã nguồn nào trong kho lưu trữ chuẩn bị hỏng bởi thay đổi này, và mã nguồn
phụ thuộc vào hành vi không được chỉ định trước đây đã sai rồi.

### Return và biến bị che khuất {#shadowing}

Một lỗi phổ biến là dùng `return` (không có đối số) sau khi gán cho một biến có cùng tên với biến kết quả nhưng không phải là cùng biến.
Tình huống này được gọi là _shadowing_: biến kết quả đã bị che khuất bởi biến khác có cùng tên được khai báo trong phạm vi bên trong.

Trong các hàm có giá trị trả về được đặt tên,
trình biên dịch Go 1 không cho phép câu lệnh return không có đối số nếu bất kỳ giá trị trả về được đặt tên nào bị che khuất tại điểm của câu lệnh return.
(Nó không phải là một phần của đặc tả, vì đây là một lĩnh vực chúng tôi vẫn đang khám phá;
tình huống tương tự như trình biên dịch từ chối các hàm không kết thúc bằng câu lệnh return tường minh.)

Hàm này trả về ngầm một giá trị trả về bị che khuất và sẽ bị trình biên dịch từ chối:

{{code "/doc/progs/go1.go" `/^func Bug/` `/^}$/`}}

_Cập nhật_:
Mã nguồn che khuất các giá trị trả về theo cách này sẽ bị trình biên dịch từ chối và cần được sửa thủ công.
Một vài trường hợp xuất hiện trong kho lưu trữ chuẩn hầu hết là lỗi.

### Sao chép struct có trường không xuất {#unexported}

Ngôn ngữ cũ không cho phép một gói tạo bản sao của một giá trị struct chứa các trường không xuất thuộc về gói khác.
Tuy nhiên, có một ngoại lệ bắt buộc cho một method receiver;
ngoài ra, các triển khai của `copy` và `append` chưa bao giờ tuân thủ hạn chế này.

Go 1 sẽ cho phép các gói sao chép các giá trị struct chứa các trường không xuất từ các gói khác.
Ngoài việc giải quyết sự không nhất quán,
thay đổi này cho phép một loại API mới: một gói có thể trả về một giá trị mờ đục mà không cần dùng đến con trỏ hoặc interface.
Các triển khai mới của `time.Time` và
`reflect.Value` là ví dụ về các kiểu tận dụng thuộc tính mới này.

Ví dụ, nếu gói `p` bao gồm các định nghĩa,

	    type Struct struct {
	        Public int
	        secret int
	    }
	    func NewStruct(a int) Struct {  // Lưu ý: không phải con trỏ.
	        return Struct{a, f(a)}
	    }
	    func (s Struct) String() string {
	        return fmt.Sprintf("{%d (secret %d)}", s.Public, s.secret)
	    }

một gói import `p` có thể gán và sao chép các giá trị kiểu
`p.Struct` tùy ý.
Phía sau, các trường không xuất sẽ được gán và sao chép giống như
khi chúng được xuất,
nhưng mã nguồn phía khách hàng sẽ không bao giờ biết về chúng. Mã nguồn

	    import "p"

	    myStruct := p.NewStruct(23)
	    copyOfMyStruct := myStruct
	    fmt.Println(myStruct, copyOfMyStruct)

sẽ cho thấy rằng trường secret của struct đã được sao chép sang giá trị mới.

_Cập nhật_:
Đây là tính năng mới, vì vậy mã nguồn hiện có không cần thay đổi.

### So sánh bằng {#equality}

Trước Go 1, ngôn ngữ không định nghĩa so sánh bằng trên các giá trị struct và mảng.
Điều này có nghĩa là,
trong số các thứ khác, struct và mảng không thể được dùng làm khóa map.
Mặt khác, Go định nghĩa so sánh bằng trên các giá trị hàm và map.
So sánh bằng hàm là vấn đề khi có closure
(khi nào hai closure bằng nhau?)
trong khi so sánh bằng map so sánh các con trỏ, không phải nội dung của map, điều này thường
không phải là điều người dùng muốn.

Go 1 giải quyết những vấn đề này.
Đầu tiên, struct và mảng có thể được so sánh bằng và bất bằng
(`==` và `!=`),
và do đó có thể được dùng làm khóa map,
miễn là chúng được cấu thành từ các phần tử mà so sánh bằng cũng được định nghĩa,
sử dụng so sánh từng phần tử.

{{code "/doc/progs/go1.go" `/type Day struct/` `/Printf/`}}

Thứ hai, Go 1 loại bỏ định nghĩa so sánh bằng cho các giá trị hàm,
ngoại trừ so sánh với `nil`.
Cuối cùng, so sánh bằng map cũng biến mất, ngoại trừ so sánh với `nil`.

Lưu ý rằng so sánh bằng vẫn chưa được định nghĩa cho slice, vì
phép tính nhìn chung là không khả thi. Cũng lưu ý rằng các toán tử
so sánh có thứ tự (< <=
`>` `>=`) vẫn chưa được định nghĩa cho
struct và mảng.

_Cập nhật_:
So sánh bằng struct và mảng là tính năng mới, vì vậy mã nguồn hiện có không cần thay đổi.
Mã nguồn hiện có phụ thuộc vào so sánh bằng hàm hoặc map sẽ
bị trình biên dịch từ chối và cần được sửa thủ công.
Ít chương trình sẽ bị ảnh hưởng, nhưng việc sửa có thể yêu cầu một số
thiết kế lại.

## Phân cấp gói {#packages}

Go 1 giải quyết nhiều thiếu sót trong thư viện chuẩn cũ và
dọn dẹp một số gói, làm cho chúng nhất quán nội bộ hơn
và di động hơn.

Phần này mô tả cách các gói đã được sắp xếp lại trong Go 1.
Một số đã di chuyển, một số đã được đổi tên, một số đã bị xóa.
Các gói mới được mô tả trong các phần sau.

### Phân cấp gói {#hierarchy}

Go 1 có phân cấp gói được sắp xếp lại để nhóm các mục liên quan
vào các thư mục con. Ví dụ, `utf8` và
`utf16` bây giờ chiếm các thư mục con của `unicode`.
Ngoài ra, [một số gói](#subrepo) đã được chuyển vào
các kho phụ của
[`code.google.com/p/go`](https://code.google.com/p/go)
trong khi [một số khác](#deleted) đã bị xóa hoàn toàn.

<table class="codetable" frame="border" summary="Moved packages">
<colgroup align="left" width="60%"></colgroup>
<colgroup align="left" width="40%"></colgroup>
<tbody><tr>
<th align="left">Đường dẫn cũ</th>
<th align="left">Đường dẫn mới</th>
</tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>asn1</td> <td>encoding/asn1</td></tr>
<tr><td>csv</td> <td>encoding/csv</td></tr>
<tr><td>gob</td> <td>encoding/gob</td></tr>
<tr><td>json</td> <td>encoding/json</td></tr>
<tr><td>xml</td> <td>encoding/xml</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>exp/template/html</td> <td>html/template</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>big</td> <td>math/big</td></tr>
<tr><td>cmath</td> <td>math/cmplx</td></tr>
<tr><td>rand</td> <td>math/rand</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>http</td> <td>net/http</td></tr>
<tr><td>http/cgi</td> <td>net/http/cgi</td></tr>
<tr><td>http/fcgi</td> <td>net/http/fcgi</td></tr>
<tr><td>http/httptest</td> <td>net/http/httptest</td></tr>
<tr><td>http/pprof</td> <td>net/http/pprof</td></tr>
<tr><td>mail</td> <td>net/mail</td></tr>
<tr><td>rpc</td> <td>net/rpc</td></tr>
<tr><td>rpc/jsonrpc</td> <td>net/rpc/jsonrpc</td></tr>
<tr><td>smtp</td> <td>net/smtp</td></tr>
<tr><td>url</td> <td>net/url</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>exec</td> <td>os/exec</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>scanner</td> <td>text/scanner</td></tr>
<tr><td>tabwriter</td> <td>text/tabwriter</td></tr>
<tr><td>template</td> <td>text/template</td></tr>
<tr><td>template/parse</td> <td>text/template/parse</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>utf8</td> <td>unicode/utf8</td></tr>
<tr><td>utf16</td> <td>unicode/utf16</td></tr>
</tbody></table>

Lưu ý rằng tên gói cho các gói `cmath` cũ và
`exp/template/html` đã thay đổi thành `cmplx`
và `template`.

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật tất cả các import và đổi tên gói cho các gói
vẫn còn trong kho lưu trữ chuẩn. Các chương trình import các gói
không còn trong kho lưu trữ chuẩn sẽ cần được chỉnh sửa thủ công.

### Cây gói exp {#exp}

Vì chúng không được tiêu chuẩn hóa, các gói trong thư mục `exp` sẽ không có trong
các bản phân phối phát hành Go 1 chuẩn, mặc dù chúng sẽ có dưới dạng mã nguồn
trong [kho lưu trữ](https://code.google.com/p/go/) cho
các nhà phát triển muốn dùng chúng.

Một số gói đã chuyển vào `exp` tại thời điểm phát hành Go 1:

  - `ebnf`
  - `html`<sup>†</sup>
  - `go/types`

(<sup>†</sup>Các kiểu `EscapeString` và `UnescapeString` vẫn
còn trong gói `html`.)

Tất cả các gói này đều có sẵn dưới các tên giống nhau, với tiền tố `exp/`: `exp/ebnf` v.v.

Ngoài ra, kiểu `utf8.String` đã được chuyển vào gói riêng của nó, `exp/utf8string`.

Cuối cùng, lệnh `gotype` hiện nằm trong `exp/gotype`, trong khi
`ebnflint` hiện nằm trong `exp/ebnflint`.
Nếu chúng được cài đặt, chúng hiện nằm trong `$GOROOT/bin/tool`.

_Cập nhật_:
Mã nguồn sử dụng các gói trong `exp` sẽ cần được cập nhật thủ công,
hoặc được biên dịch từ một cài đặt có sẵn `exp`.
Công cụ `go` `fix` hoặc trình biên dịch sẽ phàn nàn về những cách dùng như vậy.

### Cây gói old {#old}

Vì chúng không được khuyến nghị, các gói trong thư mục `old` sẽ không có trong
các bản phân phối phát hành Go 1 chuẩn, mặc dù chúng sẽ có dưới dạng mã nguồn cho
các nhà phát triển muốn dùng chúng.

Các gói ở vị trí mới của chúng là:

  - `old/netchan`

_Cập nhật_:
Mã nguồn sử dụng các gói hiện nằm trong `old` sẽ cần được cập nhật thủ công,
hoặc được biên dịch từ một cài đặt có sẵn `old`.
Công cụ `go` `fix` sẽ cảnh báo về những cách dùng như vậy.

### Các gói đã bị xóa {#deleted}

Go 1 xóa hoàn toàn một số gói:

  - `container/vector`
  - `exp/datafmt`
  - `go/typechecker`
  - `old/regexp`
  - `old/template`
  - `try`

và cả lệnh `gotry`.

_Cập nhật_:
Mã nguồn sử dụng `container/vector` nên được cập nhật để dùng
slice trực tiếp. Xem
[Go Language Community Wiki](https://code.google.com/p/go-wiki/wiki/SliceTricks) để biết một số gợi ý.
Mã nguồn sử dụng các gói khác (gần như sẽ bằng không) sẽ cần được xem xét lại.

### Các gói chuyển sang kho phụ {#subrepo}

Go 1 đã chuyển một số gói vào các kho lưu trữ khác, thường là kho phụ của
[kho lưu trữ Go chính](https://code.google.com/p/go/).
Bảng này liệt kê các đường dẫn import cũ và mới:

<table class="codetable" frame="border" summary="Sub-repositories">
<colgroup align="left" width="40%"></colgroup>
<colgroup align="left" width="60%"></colgroup>
<tbody><tr>
<th align="left">Cũ</th>
<th align="left">Mới</th>
</tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>crypto/bcrypt</td> <td>code.google.com/p/go.crypto/bcrypt</td></tr>
<tr><td>crypto/blowfish</td> <td>code.google.com/p/go.crypto/blowfish</td></tr>
<tr><td>crypto/cast5</td> <td>code.google.com/p/go.crypto/cast5</td></tr>
<tr><td>crypto/md4</td> <td>code.google.com/p/go.crypto/md4</td></tr>
<tr><td>crypto/ocsp</td> <td>code.google.com/p/go.crypto/ocsp</td></tr>
<tr><td>crypto/openpgp</td> <td>code.google.com/p/go.crypto/openpgp</td></tr>
<tr><td>crypto/openpgp/armor</td> <td>code.google.com/p/go.crypto/openpgp/armor</td></tr>
<tr><td>crypto/openpgp/elgamal</td> <td>code.google.com/p/go.crypto/openpgp/elgamal</td></tr>
<tr><td>crypto/openpgp/errors</td> <td>code.google.com/p/go.crypto/openpgp/errors</td></tr>
<tr><td>crypto/openpgp/packet</td> <td>code.google.com/p/go.crypto/openpgp/packet</td></tr>
<tr><td>crypto/openpgp/s2k</td> <td>code.google.com/p/go.crypto/openpgp/s2k</td></tr>
<tr><td>crypto/ripemd160</td> <td>code.google.com/p/go.crypto/ripemd160</td></tr>
<tr><td>crypto/twofish</td> <td>code.google.com/p/go.crypto/twofish</td></tr>
<tr><td>crypto/xtea</td> <td>code.google.com/p/go.crypto/xtea</td></tr>
<tr><td>exp/ssh</td> <td>code.google.com/p/go.crypto/ssh</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>image/bmp</td> <td>code.google.com/p/go.image/bmp</td></tr>
<tr><td>image/tiff</td> <td>code.google.com/p/go.image/tiff</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>net/dict</td> <td>code.google.com/p/go.net/dict</td></tr>
<tr><td>net/websocket</td> <td>code.google.com/p/go.net/websocket</td></tr>
<tr><td>exp/spdy</td> <td>code.google.com/p/go.net/spdy</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>encoding/git85</td> <td>code.google.com/p/go.codereview/git85</td></tr>
<tr><td>patch</td> <td>code.google.com/p/go.codereview/patch</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>exp/wingui</td> <td>code.google.com/p/gowingui</td></tr>
</tbody></table>

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật các import của các gói này để dùng các đường dẫn import mới.
Các cài đặt phụ thuộc vào các gói này sẽ cần cài đặt chúng bằng
lệnh `go get`.

## Thay đổi lớn đối với thư viện {#major}

Phần này mô tả những thay đổi đáng kể đối với các thư viện cốt lõi, những thay đổi
ảnh hưởng đến nhiều chương trình nhất.

### Kiểu error và gói errors {#errors}

Việc đặt `os.Error` trong gói `os` chủ yếu mang tính lịch sử: các lỗi xuất hiện lần đầu
khi triển khai gói `os`, và chúng có vẻ liên quan đến hệ thống vào lúc đó.
Kể từ đó, rõ ràng là các lỗi mang tính cơ bản hơn hệ điều hành. Ví dụ, sẽ tốt nếu dùng `Errors` trong các gói mà `os` phụ thuộc vào, như `syscall`.
Ngoài ra, việc có `Error` trong `os` tạo ra nhiều dependency vào `os` mà lẽ ra không cần thiết.

Go 1 giải quyết những vấn đề này bằng cách giới thiệu kiểu interface built-in `error` và gói `errors` riêng biệt (tương tự như `bytes` và `strings`) chứa các hàm tiện ích.
Nó thay thế `os.NewError` bằng
[`errors.New`](/pkg/errors/#New),
đặt lỗi vào vị trí trung tâm hơn trong môi trường.

Để phương thức `String` được sử dụng rộng rãi không vô tình thỏa mãn
interface `error`, interface `error` thay vào đó dùng
tên `Error` cho phương thức đó:

	    type error interface {
	        Error() string
	    }

Thư viện `fmt` tự động gọi `Error`, cũng như đã làm với
`String`, để in các giá trị lỗi dễ dàng.

{{code "/doc/progs/go1.go" `/START ERROR EXAMPLE/` `/END ERROR EXAMPLE/`}}

Tất cả các gói chuẩn đã được cập nhật để dùng interface mới; `os.Error` cũ đã biến mất.

Một gói mới, [`errors`](/pkg/errors/), chứa hàm

	func New(text string) error

để chuyển đổi một chuỗi thành lỗi. Nó thay thế `os.NewError` cũ.

{{code "/doc/progs/go1.go" `/ErrSyntax/`}}

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật hầu hết mã nguồn bị ảnh hưởng bởi thay đổi.
Mã nguồn định nghĩa các kiểu lỗi với phương thức `String` sẽ cần được cập nhật
thủ công để đổi tên phương thức thành `Error`.

### Lỗi system call {#errno}

Gói `syscall` cũ, có trước `os.Error`
(và hầu hết mọi thứ khác),
trả về lỗi dưới dạng giá trị `int`.
Đến lượt mình, gói `os` chuyển tiếp nhiều lỗi trong số này, chẳng hạn như
`EINVAL`, nhưng dùng một tập hợp lỗi khác nhau trên mỗi nền tảng.
Hành vi này thật khó chịu và không di động.

Trong Go 1, gói
[`syscall`](/pkg/syscall/)
thay vào đó trả về một `error` cho lỗi system call.
Trên Unix, việc triển khai được thực hiện bằng kiểu
[`syscall.Errno`](/pkg/syscall/#Errno)
thỏa mãn `error` và thay thế `os.Errno` cũ.

Các thay đổi ảnh hưởng đến `os.EINVAL` và các hằng liên quan được
mô tả [ở nơi khác](#os).

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật hầu hết mã nguồn bị ảnh hưởng bởi thay đổi.
Dù vậy, hầu hết mã nguồn nên dùng gói `os`
thay vì `syscall` và do đó sẽ không bị ảnh hưởng.

### Time {#time}

Thời gian luôn là thách thức để hỗ trợ tốt trong ngôn ngữ lập trình.
Gói `time` cũ của Go có đơn vị `int64`, không có
an toàn kiểu thực sự,
và không phân biệt giữa thời điểm tuyệt đối và khoảng thời gian.

Một trong những thay đổi toàn diện nhất trong thư viện Go 1 là
một thiết kế lại hoàn toàn của gói
[`time`](/pkg/time/).
Thay vì một số nguyên nanosecond dưới dạng `int64`,
và một kiểu `*time.Time` riêng biệt để xử lý các
đơn vị của con người như giờ và năm,
bây giờ có hai kiểu cơ bản:
[`time.Time`](/pkg/time/#Time)
(một giá trị, vì vậy `*` đã biến mất), biểu diễn một thời điểm;
và [`time.Duration`](/pkg/time/#Duration),
biểu diễn một khoảng thời gian.
Cả hai đều có độ phân giải nanosecond.
Một `Time` có thể biểu diễn bất kỳ thời điểm nào trong quá khứ xa xôi và tương lai xa, trong khi một `Duration` có thể
trải dài cộng hoặc trừ khoảng 290 năm.
Có các phương thức trên các kiểu này, cộng với một số
hằng khoảng thời gian được xác định trước hữu ích như `time.Second`.

Trong số các phương thức mới có những thứ như
[`Time.Add`](/pkg/time/#Time.Add),
cộng một `Duration` vào một `Time`, và
[`Time.Sub`](/pkg/time/#Time.Sub),
trừ hai `Time` để tạo ra một `Duration`.

Thay đổi ngữ nghĩa quan trọng nhất là epoch Unix (ngày 1 tháng 1 năm 1970) bây giờ
chỉ liên quan đến những hàm và phương thức đề cập đến Unix:
[`time.Unix`](/pkg/time/#Unix)
và các phương thức [`Unix`](/pkg/time/#Time.Unix)
và [`UnixNano`](/pkg/time/#Time.UnixNano)
của kiểu `Time`.
Cụ thể,
[`time.Now`](/pkg/time/#Now)
trả về giá trị `time.Time` thay vì, trong API cũ, một số nguyên đếm nanosecond kể từ epoch Unix.

{{code "/doc/progs/go1.go" `/sleepUntil/` `/^}/`}}

Các kiểu, phương thức và hằng số mới đã được lan truyền qua
tất cả các gói chuẩn sử dụng thời gian, chẳng hạn như `os` và
biểu diễn của nó về dấu thời gian tệp.

_Cập nhật_:
Công cụ `go` `fix` sẽ cập nhật nhiều cách dùng của gói `time` cũ để dùng các
kiểu và phương thức mới, mặc dù nó không thay thế các giá trị như `1e9`
biểu diễn nanosecond mỗi giây.
Ngoài ra, vì có thay đổi kiểu trong một số giá trị xuất hiện,
một số biểu thức được công cụ fix viết lại có thể yêu cầu
chỉnh sửa thủ công thêm; trong những trường hợp như vậy, phần viết lại sẽ bao gồm
hàm hoặc phương thức đúng cho chức năng cũ, nhưng
có thể có kiểu sai hoặc yêu cầu phân tích thêm.

## Thay đổi nhỏ đối với thư viện {#minor}

Phần này mô tả các thay đổi nhỏ hơn, chẳng hạn như những thay đổi đối với các gói ít được dùng hơn hoặc ảnh hưởng đến
ít chương trình ngoài nhu cầu chạy `go` `fix`.
Danh mục này bao gồm các gói mới trong Go 1.
Tổng thể chúng cải thiện tính di động, quy chuẩn hóa hành vi, và
làm cho các interface hiện đại hơn và mang phong cách Go hơn.

### Gói archive/zip {#archive_zip}

Trong Go 1, [`*zip.Writer`](/pkg/archive/zip/#Writer) không còn
có phương thức `Write`. Sự hiện diện của nó là một sai lầm.

_Cập nhật_:
Mã nguồn bị ảnh hưởng ít ỏi sẽ bị bắt bởi trình biên dịch và phải được cập nhật thủ công.

### Gói bufio {#bufio}

Trong Go 1, các hàm [`bufio.NewReaderSize`](/pkg/bufio/#NewReaderSize)
và
[`bufio.NewWriterSize`](/pkg/bufio/#NewWriterSize)
không còn trả về lỗi với kích thước không hợp lệ.
Nếu đối số kích thước quá nhỏ hoặc không hợp lệ, nó được điều chỉnh.

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật các lệnh gọi gán lỗi cho \_.
Các lệnh gọi không được sửa sẽ bị bắt bởi trình biên dịch và phải được cập nhật thủ công.

### Các gói compress/flate, compress/gzip và compress/zlib {#compress}

Trong Go 1, các hàm `NewWriterXxx` trong
[`compress/flate`](/pkg/compress/flate),
[`compress/gzip`](/pkg/compress/gzip) và
[`compress/zlib`](/pkg/compress/zlib)
đều trả về `(*Writer, error)` nếu chúng nhận mức nén,
và `*Writer` nếu không. Các kiểu `Compressor` và `Decompressor` của gói `gzip`
đã được đổi tên thành `Writer` và `Reader`. Kiểu
`WrongValueError` của gói `flate`
đã bị xóa.

_Cập nhật_
Chạy `go` `fix` sẽ cập nhật các tên cũ và các lệnh gọi gán lỗi cho \_.
Các lệnh gọi không được sửa sẽ bị bắt bởi trình biên dịch và phải được cập nhật thủ công.

### Các gói crypto/aes và crypto/des {#crypto_aes_des}

Trong Go 1, phương thức `Reset` đã bị xóa. Go không đảm bảo
rằng bộ nhớ không được sao chép và do đó phương thức này gây nhầm lẫn.

Các kiểu cụ thể cho cipher `*aes.Cipher`, `*des.Cipher`,
và `*des.TripleDESCipher` đã bị xóa để ủng hộ
`cipher.Block`.

_Cập nhật_:
Xóa các lệnh gọi Reset. Thay thế các cách dùng kiểu cipher cụ thể bằng
cipher.Block.

### Gói crypto/elliptic {#crypto_elliptic}

Trong Go 1, [`elliptic.Curve`](/pkg/crypto/elliptic/#Curve)
đã được tạo thành một interface để cho phép các triển khai thay thế. Các tham số curve
đã được chuyển sang cấu trúc
[`elliptic.CurveParams`](/pkg/crypto/elliptic/#CurveParams).

_Cập nhật_:
Người dùng hiện tại của `*elliptic.Curve` sẽ cần thay đổi thành
chỉ `elliptic.Curve`. Các lệnh gọi đến `Marshal`,
`Unmarshal` và `GenerateKey` bây giờ là các hàm
trong `crypto/elliptic` nhận một `elliptic.Curve`
làm đối số đầu tiên của chúng.

### Gói crypto/hmac {#crypto_hmac}

Trong Go 1, các hàm cụ thể cho hash, chẳng hạn như `hmac.NewMD5`, đã
bị xóa khỏi `crypto/hmac`. Thay vào đó, `hmac.New` nhận
một hàm trả về `hash.Hash`, chẳng hạn như `md5.New`.

_Cập nhật_:
Chạy `go` `fix` sẽ thực hiện các thay đổi cần thiết.

### Gói crypto/x509 {#crypto_x509}

Trong Go 1, hàm
[`CreateCertificate`](/pkg/crypto/x509/#CreateCertificate)
và phương thức
[`CreateCRL`](/pkg/crypto/x509/#Certificate.CreateCRL)
trong `crypto/x509` đã được sửa đổi để nhận một
`interface{}` ở nơi trước đây chúng nhận `*rsa.PublicKey`
hoặc `*rsa.PrivateKey`. Điều này sẽ cho phép các thuật toán khóa công khai khác
được triển khai trong tương lai.

_Cập nhật_:
Không cần thay đổi nào.

### Gói encoding/binary {#encoding_binary}

Trong Go 1, hàm `binary.TotalSize` đã được thay thế bằng
[`Size`](/pkg/encoding/binary/#Size),
nhận một đối số `interface{}` thay vì
một `reflect.Value`.

_Cập nhật_:
Mã nguồn bị ảnh hưởng ít ỏi sẽ bị bắt bởi trình biên dịch và phải được cập nhật thủ công.

### Gói encoding/xml {#encoding_xml}

Trong Go 1, gói [`xml`](/pkg/encoding/xml/)
đã được đưa gần hơn với thiết kế của các gói marshaling khác như
[`encoding/gob`](/pkg/encoding/gob/).

Kiểu `Parser` cũ được đổi tên thành
[`Decoder`](/pkg/encoding/xml/#Decoder) và có phương thức
[`Decode`](/pkg/encoding/xml/#Decoder.Decode) mới. Một kiểu
[`Encoder`](/pkg/encoding/xml/#Encoder) cũng đã được giới thiệu.

Các hàm [`Marshal`](/pkg/encoding/xml/#Marshal)
và [`Unmarshal`](/pkg/encoding/xml/#Unmarshal)
bây giờ làm việc với các giá trị `[]byte`. Để làm việc với các stream,
dùng các kiểu [`Encoder`](/pkg/encoding/xml/#Encoder)
và [`Decoder`](/pkg/encoding/xml/#Decoder) mới.

Khi marshaling hoặc unmarshaling các giá trị, định dạng của các flag được hỗ trợ trong
field tag đã thay đổi để gần hơn với gói
[`json`](/pkg/encoding/json)
(`` `xml:"name,flag"` ``). Việc khớp giữa field tag, tên field
và tên thuộc tính XML và phần tử bây giờ phân biệt chữ hoa chữ thường.
Field tag `XMLName`, nếu có, cũng phải khớp với tên
của phần tử XML đang được marshal.

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật hầu hết cách dùng của gói ngoại trừ một số lệnh gọi đến
`Unmarshal`. Cần cẩn thận đặc biệt với field tag,
vì công cụ fix sẽ không cập nhật chúng và nếu không được sửa thủ công chúng sẽ
hoạt động sai trong im lặng ở một số trường hợp. Ví dụ, cũ
`"attr"` bây giờ được viết `",attr"` trong khi
`"attr"` thuần túy vẫn hợp lệ nhưng có ý nghĩa khác.

### Gói expvar {#expvar}

Trong Go 1, hàm `RemoveAll` đã bị xóa.
Hàm `Iter` và phương thức Iter trên `*Map` đã
được thay thế bởi
[`Do`](/pkg/expvar/#Do)
và
[`(*Map).Do`](/pkg/expvar/#Map.Do).

_Cập nhật_:
Hầu hết mã nguồn dùng `expvar` sẽ không cần thay đổi. Mã nguồn hiếm dùng
`Iter` có thể được cập nhật để truyền một closure cho `Do` để đạt được hiệu ứng tương tự.

### Gói flag {#flag}

Trong Go 1, interface [`flag.Value`](/pkg/flag/#Value) đã thay đổi một chút.
Phương thức `Set` bây giờ trả về một `error` thay vì
một `bool` để biểu thị thành công hay thất bại.

Ngoài ra còn có một loại flag mới, `Duration`, để hỗ trợ các giá trị đối số
chỉ định khoảng thời gian.
Các giá trị cho các flag như vậy phải được đưa ra kèm đơn vị, giống như `time.Duration`
định dạng chúng: `10s`, `1h30m`, v.v.

{{code "/doc/progs/go1.go" `/timeout/`}}

_Cập nhật_:
Các chương trình triển khai flag của riêng chúng sẽ cần sửa thủ công nhỏ để cập nhật
phương thức `Set` của chúng.
Flag `Duration` là mới và không ảnh hưởng đến mã nguồn hiện có.

### Các gói go/\* {#go}

Một số gói trong `go` có API được sửa đổi nhẹ.

Kiểu `Mode` cụ thể đã được giới thiệu cho các flag chế độ cấu hình
trong các gói
[`go/scanner`](/pkg/go/scanner/),
[`go/parser`](/pkg/go/parser/),
[`go/printer`](/pkg/go/printer/) và
[`go/doc`](/pkg/go/doc/).

Các chế độ `AllowIllegalChars` và `InsertSemis` đã bị xóa
khỏi gói [`go/scanner`](/pkg/go/scanner/). Chúng chủ yếu
hữu ích cho việc quét văn bản khác với các file mã nguồn Go. Thay vào đó, gói
[`text/scanner`](/pkg/text/scanner/) nên được dùng
cho mục đích đó.

[`ErrorHandler`](/pkg/go/scanner/#ErrorHandler) được cung cấp
cho phương thức [`Init`](/pkg/go/scanner/#Scanner.Init) của scanner bây giờ
chỉ là một hàm đơn giản thay vì một interface. Kiểu `ErrorVector` đã
bị xóa để ủng hộ kiểu [`ErrorList`](/pkg/go/scanner/#ErrorList) (đã tồn tại),
và các phương thức `ErrorVector` đã được di chuyển sang. Thay vì nhúng
một `ErrorVector` vào một client của scanner, bây giờ một client nên duy trì
một `ErrorList`.

Tập hợp các hàm phân tích cú pháp được cung cấp bởi gói [`go/parser`](/pkg/go/parser/)
đã được thu gọn thành hàm phân tích cú pháp chính
[`ParseFile`](/pkg/go/parser/#ParseFile), và một vài
hàm tiện ích [`ParseDir`](/pkg/go/parser/#ParseDir)
và [`ParseExpr`](/pkg/go/parser/#ParseExpr).

Gói [`go/printer`](/pkg/go/printer/) hỗ trợ thêm
chế độ cấu hình [`SourcePos`](/pkg/go/printer/#Mode);
nếu được đặt, printer sẽ phát ra các comment `//line` sao cho đầu ra được tạo
chứa thông tin vị trí mã nguồn gốc. Kiểu mới
[`CommentedNode`](/pkg/go/printer/#CommentedNode) có thể được
dùng để cung cấp các comment liên kết với một
[`ast.Node`](/pkg/go/ast/#Node) tùy ý (cho đến nay chỉ
[`ast.File`](/pkg/go/ast/#File) mang thông tin comment).

Tên kiểu của gói [`go/doc`](/pkg/go/doc/) đã được
đơn giản hóa bằng cách xóa hậu tố `Doc`: `PackageDoc`
bây giờ là `Package`, `ValueDoc` là `Value`, v.v.
Ngoài ra, tất cả các kiểu bây giờ đều nhất quán có field `Name` (hoặc `Names`,
trong trường hợp kiểu `Value`) và `Type.Factories` đã trở thành
`Type.Funcs`.
Thay vì gọi `doc.NewPackageDoc(pkg, importpath)`,
tài liệu cho một gói được tạo với:

	    doc.New(pkg, importpath, mode)

trong đó tham số `mode` mới chỉ định chế độ hoạt động:
nếu được đặt thành [`AllDecls`](/pkg/go/doc/#AllDecls), tất cả các khai báo
(không chỉ các khai báo được xuất) được xem xét.
Hàm `NewFileDoc` đã bị xóa, và hàm
`CommentText` đã trở thành phương thức
[`Text`](/pkg/go/ast/#CommentGroup.Text) của
[`ast.CommentGroup`](/pkg/go/ast/#CommentGroup).

Trong gói [`go/token`](/pkg/go/token/), phương thức
[`token.FileSet`](/pkg/go/token/#FileSet) `Files`
(trước đây trả về một channel của `*token.File`) đã được thay thế
bằng iterator [`Iterate`](/pkg/go/token/#FileSet.Iterate) nhận
một đối số hàm thay thế.

Trong gói [`go/build`](/pkg/go/build/), API
gần như đã được thay thế hoàn toàn.
Gói vẫn tính toán thông tin gói Go
nhưng nó không chạy build: các kiểu `Cmd` và `Script`
đã biến mất.
(Để build mã nguồn, hãy dùng lệnh
[`go`](/cmd/go/) mới thay thế.)
Kiểu `DirInfo` bây giờ được đặt tên là
[`Package`](/pkg/go/build/#Package).
`FindTree` và `ScanDir` được thay thế bởi
[`Import`](/pkg/go/build/#Import)
và
[`ImportDir`](/pkg/go/build/#ImportDir).

_Cập nhật_:
Mã nguồn dùng các gói trong `go` sẽ phải được cập nhật thủ công;
trình biên dịch sẽ từ chối các cách dùng không chính xác. Các template được sử dụng kết hợp với bất kỳ
kiểu `go/doc` nào có thể cần sửa thủ công; các field được đổi tên sẽ dẫn đến
lỗi thời gian chạy.

### Gói hash {#hash}

Trong Go 1, định nghĩa của [`hash.Hash`](/pkg/hash/#Hash) bao gồm
một phương thức mới, `BlockSize`. Phương thức mới này được dùng chủ yếu trong
các thư viện mã hóa.

Phương thức `Sum` của interface
[`hash.Hash`](/pkg/hash/#Hash) bây giờ nhận một
đối số `[]byte`, mà giá trị hash sẽ được nối thêm vào.
Hành vi trước đây có thể được tái tạo bằng cách thêm đối số `nil` vào lệnh gọi.

_Cập nhật_:
Các triển khai hiện có của `hash.Hash` sẽ cần thêm phương thức
`BlockSize`. Các hash xử lý đầu vào một byte tại một thời điểm có thể
triển khai `BlockSize` để trả về 1.
Chạy `go` `fix` sẽ cập nhật các lệnh gọi đến phương thức `Sum` của các
triển khai khác nhau của `hash.Hash`.

_Cập nhật_:
Vì chức năng của gói là mới, không cần cập nhật nào.

### Gói http {#http}

Trong Go 1, gói [`http`](/pkg/net/http/) được tái cấu trúc,
đặt một số tiện ích vào thư mục con
[`httputil`](/pkg/net/http/httputil/).
Những phần này chỉ hiếm khi cần bởi các client HTTP.
Các mục bị ảnh hưởng là:

  - ClientConn
  - DumpRequest
  - DumpRequestOut
  - DumpResponse
  - NewChunkedReader
  - NewChunkedWriter
  - NewClientConn
  - NewProxyClientConn
  - NewServerConn
  - NewSingleHostReverseProxy
  - ReverseProxy
  - ServerConn

Field `Request.RawURL` đã bị xóa; nó là một tạo phẩm lịch sử.

Các hàm `Handle` và `HandleFunc`,
và các phương thức cùng tên của `ServeMux`,
bây giờ gây panic nếu cố gắng đăng ký cùng một pattern hai lần.

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật một số chương trình bị ảnh hưởng ngoại trừ
các cách dùng `RawURL`, phải được sửa thủ công.

### Gói image {#image}

Gói [`image`](/pkg/image/) đã có một số
thay đổi nhỏ, sắp xếp lại và đổi tên.

Hầu hết mã xử lý màu đã được chuyển vào gói riêng của nó,
[`image/color`](/pkg/image/color/).
Đối với các phần tử đã di chuyển, xuất hiện một sự đối xứng; ví dụ,
mỗi pixel của
[`image.RGBA`](/pkg/image/#RGBA)
là một
[`color.RGBA`](/pkg/image/color/#RGBA).

Gói `image/ycbcr` cũ đã được hợp nhất, với một số
đổi tên, vào các gói
[`image`](/pkg/image/)
và
[`image/color`](/pkg/image/color/).

Kiểu `image.ColorImage` cũ vẫn còn trong gói `image`
nhưng đã được đổi tên thành
[`image.Uniform`](/pkg/image/#Uniform),
trong khi `image.Tiled` đã bị xóa.

Bảng này liệt kê các đổi tên.

<table class="codetable" frame="border" summary="image renames">
<colgroup align="left" width="50%"></colgroup>
<colgroup align="left" width="50%"></colgroup>
<tbody><tr>
<th align="left">Cũ</th>
<th align="left">Mới</th>
</tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>image.Color</td> <td>color.Color</td></tr>
<tr><td>image.ColorModel</td> <td>color.Model</td></tr>
<tr><td>image.ColorModelFunc</td> <td>color.ModelFunc</td></tr>
<tr><td>image.PalettedColorModel</td> <td>color.Palette</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>image.RGBAColor</td> <td>color.RGBA</td></tr>
<tr><td>image.RGBA64Color</td> <td>color.RGBA64</td></tr>
<tr><td>image.NRGBAColor</td> <td>color.NRGBA</td></tr>
<tr><td>image.NRGBA64Color</td> <td>color.NRGBA64</td></tr>
<tr><td>image.AlphaColor</td> <td>color.Alpha</td></tr>
<tr><td>image.Alpha16Color</td> <td>color.Alpha16</td></tr>
<tr><td>image.GrayColor</td> <td>color.Gray</td></tr>
<tr><td>image.Gray16Color</td> <td>color.Gray16</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>image.RGBAColorModel</td> <td>color.RGBAModel</td></tr>
<tr><td>image.RGBA64ColorModel</td> <td>color.RGBA64Model</td></tr>
<tr><td>image.NRGBAColorModel</td> <td>color.NRGBAModel</td></tr>
<tr><td>image.NRGBA64ColorModel</td> <td>color.NRGBA64Model</td></tr>
<tr><td>image.AlphaColorModel</td> <td>color.AlphaModel</td></tr>
<tr><td>image.Alpha16ColorModel</td> <td>color.Alpha16Model</td></tr>
<tr><td>image.GrayColorModel</td> <td>color.GrayModel</td></tr>
<tr><td>image.Gray16ColorModel</td> <td>color.Gray16Model</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>ycbcr.RGBToYCbCr</td> <td>color.RGBToYCbCr</td></tr>
<tr><td>ycbcr.YCbCrToRGB</td> <td>color.YCbCrToRGB</td></tr>
<tr><td>ycbcr.YCbCrColorModel</td> <td>color.YCbCrModel</td></tr>
<tr><td>ycbcr.YCbCrColor</td> <td>color.YCbCr</td></tr>
<tr><td>ycbcr.YCbCr</td> <td>image.YCbCr</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>ycbcr.SubsampleRatio444</td> <td>image.YCbCrSubsampleRatio444</td></tr>
<tr><td>ycbcr.SubsampleRatio422</td> <td>image.YCbCrSubsampleRatio422</td></tr>
<tr><td>ycbcr.SubsampleRatio420</td> <td>image.YCbCrSubsampleRatio420</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>image.ColorImage</td> <td>image.Uniform</td></tr>
</tbody></table>

Các hàm `New` của gói image
([`NewRGBA`](/pkg/image/#NewRGBA),
[`NewRGBA64`](/pkg/image/#NewRGBA64), v.v.)
nhận một [`image.Rectangle`](/pkg/image/#Rectangle) làm đối số
thay vì bốn số nguyên.

Cuối cùng, có các biến `color.Color` được định nghĩa sẵn mới
[`color.Black`](/pkg/image/color/#Black),
[`color.White`](/pkg/image/color/#White),
[`color.Opaque`](/pkg/image/color/#Opaque)
và
[`color.Transparent`](/pkg/image/color/#Transparent).

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật hầu hết mã nguồn bị ảnh hưởng bởi thay đổi.

### Gói log/syslog {#log_syslog}

Trong Go 1, hàm [`syslog.NewLogger`](/pkg/log/syslog/#NewLogger)
trả về cả một lỗi lẫn một `log.Logger`.

_Cập nhật_:
Mã nguồn bị ảnh hưởng ít ỏi sẽ bị bắt bởi trình biên dịch và phải được cập nhật thủ công.

### Gói mime {#mime}

Trong Go 1, hàm [`FormatMediaType`](/pkg/mime/#FormatMediaType)
của gói `mime` đã được đơn giản hóa để làm cho nó
nhất quán với
[`ParseMediaType`](/pkg/mime/#ParseMediaType).
Bây giờ nó nhận `"text/html"` thay vì `"text"` và `"html"`.

_Cập nhật_:
Mã nguồn bị ảnh hưởng ít ỏi sẽ bị bắt bởi trình biên dịch và phải được cập nhật thủ công.

### Gói net {#net}

Trong Go 1, các phương thức `SetTimeout`,
`SetReadTimeout` và `SetWriteTimeout` khác nhau
đã được thay thế bằng
[`SetDeadline`](/pkg/net/#IPConn.SetDeadline),
[`SetReadDeadline`](/pkg/net/#IPConn.SetReadDeadline) và
[`SetWriteDeadline`](/pkg/net/#IPConn.SetWriteDeadline),
tương ứng. Thay vì nhận một giá trị timeout tính bằng nanosecond áp dụng cho bất kỳ hoạt động nào trên kết nối, các phương thức mới đặt một
deadline tuyệt đối (dưới dạng giá trị `time.Time`) sau đó
đọc và ghi sẽ timeout và không còn block.

Ngoài ra còn có các hàm mới
[`net.DialTimeout`](/pkg/net/#DialTimeout)
để đơn giản hóa việc timeout khi kết nối đến một địa chỉ mạng và
[`net.ListenMulticastUDP`](/pkg/net/#ListenMulticastUDP)
để cho phép UDP multicast lắng nghe đồng thời qua nhiều listener.
Hàm `net.ListenMulticastUDP` thay thế các phương thức
`JoinGroup` và `LeaveGroup` cũ.

_Cập nhật_:
Mã nguồn dùng các phương thức cũ sẽ không biên dịch được và phải được cập nhật thủ công.
Sự thay đổi ngữ nghĩa làm cho công cụ fix khó cập nhật tự động.

### Gói os {#os}

Hàm `Time` đã bị xóa; người gọi nên dùng
kiểu [`Time`](/pkg/time/#Time) từ gói
`time`.

Hàm `Exec` đã bị xóa; người gọi nên dùng
`Exec` từ gói `syscall`, khi có sẵn.

Hàm `ShellExpand` đã được đổi tên thành [`ExpandEnv`](/pkg/os/#ExpandEnv).

Hàm [`NewFile`](/pkg/os/#NewFile)
bây giờ nhận một `uintptr` fd, thay vì một `int`.
Phương thức [`Fd`](/pkg/os/#File.Fd) trên các file bây giờ
cũng trả về một `uintptr`.

Không còn các hằng lỗi như `EINVAL`
trong gói `os`, vì tập hợp các giá trị khác nhau với
hệ điều hành cơ bản. Có các hàm di động mới như
[`IsPermission`](/pkg/os/#IsPermission)
để kiểm tra các thuộc tính lỗi phổ biến, cộng với một vài giá trị lỗi mới
với tên mang phong cách Go hơn, chẳng hạn như
[`ErrPermission`](/pkg/os/#ErrPermission)
và
[`ErrNotExist`](/pkg/os/#ErrNotExist).

Hàm `Getenverror` đã bị xóa. Để phân biệt
giữa một biến môi trường không tồn tại và một chuỗi rỗng,
dùng [`os.Environ`](/pkg/os/#Environ) hoặc
[`syscall.Getenv`](/pkg/syscall/#Getenv).

Phương thức [`Process.Wait`](/pkg/os/#Process.Wait) đã
bỏ đối số tùy chọn của nó và các hằng liên quan đã biến mất
khỏi gói.
Ngoài ra, hàm `Wait` đã biến mất; chỉ phương thức của
kiểu `Process` vẫn còn.

Kiểu `Waitmsg` được trả về bởi
[`Process.Wait`](/pkg/os/#Process.Wait)
đã được thay thế bằng kiểu
[`ProcessState`](/pkg/os/#ProcessState)
di động hơn với các phương thức accessor để lấy thông tin về
process.
Do thay đổi đối với `Wait`, giá trị `ProcessState`
luôn mô tả một process đã thoát.
Các mối lo ngại về tính di động đã đơn giản hóa interface theo những cách khác, nhưng các giá trị được trả về bởi
[`ProcessState.Sys`](/pkg/os/#ProcessState.Sys) và
[`ProcessState.SysUsage`](/pkg/os/#ProcessState.SysUsage)
có thể được type-assert sang các cấu trúc dữ liệu cụ thể cho hệ thống bên dưới như
[`syscall.WaitStatus`](/pkg/syscall/#WaitStatus) và
[`syscall.Rusage`](/pkg/syscall/#Rusage) trên Unix.

_Cập nhật_:
Chạy `go` `fix` sẽ bỏ đối số zero vào `Process.Wait`.
Tất cả các thay đổi khác sẽ bị bắt bởi trình biên dịch và phải được cập nhật thủ công.

#### Kiểu os.FileInfo {#os_fileinfo}

Go 1 định nghĩa lại kiểu [`os.FileInfo`](/pkg/os/#FileInfo),
thay đổi nó từ struct thành interface:

	    type FileInfo interface {
	        Name() string       // tên cơ sở của file
	        Size() int64        // độ dài tính bằng byte
	        Mode() FileMode     // các bit chế độ file
	        ModTime() time.Time // thời gian sửa đổi
	        IsDir() bool        // viết tắt cho Mode().IsDir()
	        Sys() interface{}   // nguồn dữ liệu cơ bản (có thể trả về nil)
	    }

Thông tin chế độ file đã được chuyển vào một kiểu con gọi là
[`os.FileMode`](/pkg/os/#FileMode),
một kiểu số nguyên đơn giản với các phương thức `IsDir`, `Perm` và `String`.

Các chi tiết cụ thể của hệ thống về chế độ file và thuộc tính như (trên Unix)
i-number đã bị xóa hoàn toàn khỏi `FileInfo`.
Thay vào đó, gói `os` của mỗi hệ điều hành cung cấp một
triển khai của interface `FileInfo`, có
phương thức `Sys` trả về
biểu diễn metadata file cụ thể cho hệ thống.
Ví dụ, để khám phá i-number của một file trên hệ thống Unix, giải nén
`FileInfo` như thế này:

	    fi, err := os.Stat("hello.go")
	    if err != nil {
	        log.Fatal(err)
	    }
	    // Kiểm tra đây là file Unix.
	    unixStat, ok := fi.Sys().(*syscall.Stat_t)
	    if !ok {
	        log.Fatal("hello.go: not a Unix file")
	    }
	    fmt.Printf("file i-number: %d\n", unixStat.Ino)

Giả sử (điều này không khôn ngoan) rằng `"hello.go"` là một file Unix,
biểu thức i-number có thể được rút gọn thành

	    fi.Sys().(*syscall.Stat_t).Ino

Đại đa số cách dùng `FileInfo` chỉ cần các phương thức
của interface chuẩn.

Gói `os` không còn chứa các wrapper cho các lỗi POSIX
như `ENOENT`.
Đối với một vài chương trình cần xác minh các điều kiện lỗi cụ thể, bây giờ có
các hàm boolean
[`IsExist`](/pkg/os/#IsExist),
[`IsNotExist`](/pkg/os/#IsNotExist)
và
[`IsPermission`](/pkg/os/#IsPermission).

{{code "/doc/progs/go1.go" `/os\.Open/` `/}/`}}

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật mã nguồn dùng phiên bản tương đương cũ của API `os.FileInfo`
và `os.FileMode` hiện tại.
Mã nguồn cần chi tiết file cụ thể cho hệ thống sẽ cần được cập nhật thủ công.
Mã nguồn dùng các giá trị lỗi POSIX cũ từ gói `os`
sẽ không biên dịch được và cũng cần được cập nhật thủ công.

### Gói os/signal {#os_signal}

Gói `os/signal` trong Go 1 thay thế hàm
`Incoming`, trả về một channel
nhận tất cả các tín hiệu đến,
bằng hàm có chọn lọc `Notify`, yêu cầu
gửi các tín hiệu cụ thể trên một channel hiện có.

_Cập nhật_:
Mã nguồn phải được cập nhật thủ công.
Bản dịch nguyên văn của

	c := signal.Incoming()

là

	c := make(chan os.Signal, 1)
	signal.Notify(c) // yêu cầu tất cả các tín hiệu

nhưng hầu hết mã nguồn nên liệt kê các tín hiệu cụ thể mà chúng muốn xử lý:

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT)

### Gói path/filepath {#path_filepath}

Trong Go 1, hàm [`Walk`](/pkg/path/filepath/#Walk) của
gói `path/filepath`
đã được thay đổi để nhận giá trị hàm kiểu
[`WalkFunc`](/pkg/path/filepath/#WalkFunc)
thay vì giá trị interface `Visitor`.
`WalkFunc` hợp nhất việc xử lý cả file và thư mục.

	    type WalkFunc func(path string, info os.FileInfo, err error) error

Hàm `WalkFunc` sẽ được gọi ngay cả đối với các file hoặc thư mục không thể mở được;
trong những trường hợp như vậy, đối số lỗi sẽ mô tả sự thất bại.
Nếu nội dung của một thư mục cần bị bỏ qua,
hàm nên trả về giá trị [`filepath.SkipDir`](/pkg/path/filepath/#pkg-variables)

{{code "/doc/progs/go1.go" `/STARTWALK/` `/ENDWALK/`}}

_Cập nhật_:
Thay đổi đơn giản hóa hầu hết mã nguồn nhưng có hệ quả tinh tế, vì vậy các chương trình bị ảnh hưởng
sẽ cần được cập nhật thủ công.
Trình biên dịch sẽ bắt mã nguồn dùng interface cũ.

### Gói regexp {#regexp}

Gói [`regexp`](/pkg/regexp/) đã được viết lại.
Nó có cùng interface nhưng đặc tả của các biểu thức chính quy
mà nó hỗ trợ đã thay đổi từ dạng "egrep" cũ sang dạng của
[RE2](https://code.google.com/p/re2/).

_Cập nhật_:
Mã nguồn dùng gói nên kiểm tra thủ công các biểu thức chính quy của nó.

### Gói runtime {#runtime}

Trong Go 1, phần lớn API được xuất bởi gói
`runtime` đã bị xóa để ủng hộ
chức năng được cung cấp bởi các gói khác.
Mã nguồn dùng interface `runtime.Type`
hoặc các triển khai kiểu cụ thể của nó nên
bây giờ dùng gói [`reflect`](/pkg/reflect/).
Mã nguồn dùng `runtime.Semacquire` hoặc `runtime.Semrelease`
nên dùng channel hoặc các trừu tượng trong gói [`sync`](/pkg/sync/).
Các hàm `runtime.Alloc`, `runtime.Free`,
và `runtime.Lookup`, một API không an toàn được tạo ra để
gỡ lỗi bộ cấp phát bộ nhớ, không có sự thay thế.

Trước đây, `runtime.MemStats` là một biến toàn cục chứa
thống kê về cấp phát bộ nhớ, và các lệnh gọi đến `runtime.UpdateMemStats`
đảm bảo rằng nó được cập nhật.
Trong Go 1, `runtime.MemStats` là một kiểu struct, và mã nguồn nên dùng
[`runtime.ReadMemStats`](/pkg/runtime/#ReadMemStats)
để lấy thống kê hiện tại.

Gói thêm một hàm mới,
[`runtime.NumCPU`](/pkg/runtime/#NumCPU), trả về số lượng CPU có sẵn
để thực thi song song, theo báo cáo của nhân hệ điều hành.
Giá trị của nó có thể thông báo cho việc đặt `GOMAXPROCS`.
Các hàm `runtime.Cgocalls` và `runtime.Goroutines`
đã được đổi tên thành `runtime.NumCgoCall` và `runtime.NumGoroutine`.

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật mã nguồn cho các đổi tên hàm.
Mã nguồn khác sẽ cần được cập nhật thủ công.

### Gói strconv {#strconv}

Trong Go 1, gói
[`strconv`](/pkg/strconv/)
đã được làm lại đáng kể để làm cho nó mang phong cách Go hơn và ít phong cách C hơn,
mặc dù `Atoi` vẫn còn (nó tương tự như
`int(ParseInt(x, 10, 0))`, cũng như
`Itoa(x)` (`FormatInt(int64(x), 10)`).
Ngoài ra còn có các biến thể mới của một số hàm nối vào byte slice thay vì
trả về chuỗi, để cho phép kiểm soát việc cấp phát.

Bảng này tóm tắt các đổi tên; xem
[tài liệu gói](/pkg/strconv/)
để biết chi tiết đầy đủ.

<table class="codetable" frame="border" summary="strconv renames">
<colgroup align="left" width="50%"></colgroup>
<colgroup align="left" width="50%"></colgroup>
<tbody><tr>
<th align="left">Lệnh gọi cũ</th>
<th align="left">Lệnh gọi mới</th>
</tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Atob(x)</td> <td>ParseBool(x)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Atof32(x)</td> <td>ParseFloat(x, 32)§</td></tr>
<tr><td>Atof64(x)</td> <td>ParseFloat(x, 64)</td></tr>
<tr><td>AtofN(x, n)</td> <td>ParseFloat(x, n)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Atoi(x)</td> <td>Atoi(x)</td></tr>
<tr><td>Atoi(x)</td> <td>ParseInt(x, 10, 0)§</td></tr>
<tr><td>Atoi64(x)</td> <td>ParseInt(x, 10, 64)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Atoui(x)</td> <td>ParseUint(x, 10, 0)§</td></tr>
<tr><td>Atoui64(x)</td> <td>ParseUint(x, 10, 64)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Btoi64(x, b)</td> <td>ParseInt(x, b, 64)</td></tr>
<tr><td>Btoui64(x, b)</td> <td>ParseUint(x, b, 64)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Btoa(x)</td> <td>FormatBool(x)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Ftoa32(x, f, p)</td> <td>FormatFloat(float64(x), f, p, 32)</td></tr>
<tr><td>Ftoa64(x, f, p)</td> <td>FormatFloat(x, f, p, 64)</td></tr>
<tr><td>FtoaN(x, f, p, n)</td> <td>FormatFloat(x, f, p, n)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Itoa(x)</td> <td>Itoa(x)</td></tr>
<tr><td>Itoa(x)</td> <td>FormatInt(int64(x), 10)</td></tr>
<tr><td>Itoa64(x)</td> <td>FormatInt(x, 10)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Itob(x, b)</td> <td>FormatInt(int64(x), b)</td></tr>
<tr><td>Itob64(x, b)</td> <td>FormatInt(x, b)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Uitoa(x)</td> <td>FormatUint(uint64(x), 10)</td></tr>
<tr><td>Uitoa64(x)</td> <td>FormatUint(x, 10)</td></tr>
<tr>
<td colspan="2"><hr></hr></td>
</tr>
<tr><td>Uitob(x, b)</td> <td>FormatUint(uint64(x), b)</td></tr>
<tr><td>Uitob64(x, b)</td> <td>FormatUint(x, b)</td></tr>
</tbody></table>

_Cập nhật_:
Chạy `go` `fix` sẽ cập nhật hầu hết mã nguồn bị ảnh hưởng bởi thay đổi.
\
§ `Atoi` vẫn còn nhưng `Atoui` và `Atof32` thì không, vì vậy
chúng có thể yêu cầu
một phép ép kiểu phải được thêm thủ công; công cụ `go` `fix` sẽ cảnh báo về điều đó.

### Các gói template {#templates}

Các gói `template` và `exp/template/html` đã được chuyển đến
[`text/template`](/pkg/text/template/) và
[`html/template`](/pkg/html/template/).
Quan trọng hơn, interface với các gói này đã được đơn giản hóa.
Ngôn ngữ template vẫn như cũ, nhưng khái niệm "template set" đã biến mất
và các hàm và phương thức của các gói đã thay đổi theo đó,
thường bằng cách loại bỏ.

Thay vì các tập hợp, một đối tượng `Template`
có thể chứa nhiều định nghĩa template được đặt tên,
trên thực tế xây dựng
các không gian tên cho việc gọi template.
Một template có thể gọi bất kỳ template nào khác được liên kết với nó, nhưng chỉ những
template được liên kết với nó.
Cách đơn giản nhất để liên kết các template là phân tích cú pháp chúng cùng nhau, điều gì đó
được thực hiện dễ hơn với cấu trúc mới của các gói.

_Cập nhật_:
Các import sẽ được cập nhật bởi công cụ fix.
Các cách dùng template đơn sẽ phần lớn không bị ảnh hưởng.
Mã nguồn dùng nhiều template kết hợp sẽ cần được cập nhật thủ công.
Các [ví dụ](/pkg/text/template/#pkg-examples) trong
tài liệu cho `text/template` có thể cung cấp hướng dẫn.

### Gói testing {#testing}

Gói testing có một kiểu, `B`, được truyền làm đối số cho các hàm benchmark.
Trong Go 1, `B` có các phương thức mới, tương tự như của `T`, cho phép
ghi nhật ký và báo cáo lỗi.

{{code "/doc/progs/go1.go" `/func.*Benchmark/` `/^}/`}}

_Cập nhật_:
Mã nguồn hiện có không bị ảnh hưởng, mặc dù các benchmark dùng `println`
hoặc `panic` nên được cập nhật để dùng các phương thức mới.

### Gói testing/script {#testing_script}

Gói testing/script đã bị xóa. Nó là một mảnh vụn không cần thiết.

_Cập nhật_:
Không có mã nguồn nào có khả năng bị ảnh hưởng.

### Gói unsafe {#unsafe}

Trong Go 1, các hàm
`unsafe.Typeof`, `unsafe.Reflect`,
`unsafe.Unreflect`, `unsafe.New` và
`unsafe.NewArray` đã bị xóa;
chúng trùng lặp chức năng an toàn hơn được cung cấp bởi
gói [`reflect`](/pkg/reflect/).

_Cập nhật_:
Mã nguồn dùng các hàm này phải được viết lại để dùng
gói [`reflect`](/pkg/reflect/).
Các thay đổi đối với [encoding/gob](/change/2646dc956207) và [thư viện protocol buffer](https://code.google.com/p/goprotobuf/source/detail?r=5340ad310031)
có thể hữu ích như các ví dụ.

### Gói url {#url}

Trong Go 1, một số field từ kiểu [`url.URL`](/pkg/net/url/#URL)
đã bị xóa hoặc thay thế.

Phương thức [`String`](/pkg/net/url/#URL.String) bây giờ
xây dựng lại một chuỗi URL được mã hóa một cách có thể dự đoán bằng cách dùng tất cả các
field của `URL` khi cần thiết. Chuỗi kết quả cũng sẽ không còn
có mật khẩu được escape.

Field `Raw` đã bị xóa. Trong hầu hết các trường hợp, phương thức `String`
có thể được dùng thay thế.

Field `RawUserinfo` cũ được thay thế bởi field `User`
kiểu [`*net.Userinfo`](/pkg/net/url/#Userinfo).
Các giá trị kiểu này có thể được tạo ra bằng các hàm mới [`net.User`](/pkg/net/url/#User)
và [`net.UserPassword`](/pkg/net/url/#UserPassword).
Các hàm `EscapeUserinfo` và `UnescapeUserinfo`
cũng đã biến mất.

Field `RawAuthority` đã bị xóa. Thông tin tương tự có sẵn
trong các field `Host` và `User`.

Field `RawPath` và phương thức `EncodedPath` đã
bị xóa. Thông tin đường dẫn trong các URL có gốc (với dấu gạch chéo sau
schema) bây giờ chỉ có sẵn ở dạng đã giải mã trong field `Path`.
Đôi khi, dữ liệu được mã hóa có thể cần thiết để lấy thông tin
bị mất trong quá trình giải mã. Những trường hợp này phải được xử lý bằng cách truy cập
dữ liệu mà URL được xây dựng từ đó.

Các URL có đường dẫn không có gốc, chẳng hạn như `"mailto:dev@golang.org?subject=Hi"`,
cũng được xử lý khác đi. Field boolean `OpaquePath` đã bị
xóa và một field chuỗi `Opaque` mới được giới thiệu để giữ đường dẫn
được mã hóa cho các URL như vậy. Trong Go 1, URL được trích dẫn phân tích cú pháp thành:

	    URL{
	        Scheme: "mailto",
	        Opaque: "dev@golang.org",
	        RawQuery: "subject=Hi",
	    }

Một phương thức mới [`RequestURI`](/pkg/net/url/#URL.RequestURI) đã được
thêm vào `URL`.

Hàm `ParseWithReference` đã được đổi tên thành `ParseWithFragment`.

_Cập nhật_:
Mã nguồn dùng các field cũ sẽ không biên dịch được và phải được cập nhật thủ công.
Các thay đổi ngữ nghĩa làm cho công cụ fix khó cập nhật tự động.

## Lệnh go {#cmd_go}

Go 1 giới thiệu [lệnh go](/cmd/go/), một công cụ để tìm nạp,
build và cài đặt các gói và lệnh Go. Lệnh `go`
loại bỏ makefile, thay vào đó dùng mã nguồn Go để tìm các dependency và
xác định các điều kiện build. Hầu hết các chương trình Go hiện có sẽ không còn cần
makefile để được build.

Xem [Cách viết mã Go](/doc/code.html) để có phần giới thiệu về
lệnh `go` và [tài liệu lệnh go](/cmd/go/)
để biết chi tiết đầy đủ.

_Cập nhật_:
Các dự án phụ thuộc vào cơ sở hạ tầng build dựa trên makefile cũ của dự án Go
(`Make.pkg`, `Make.cmd`, v.v.) nên
chuyển sang dùng lệnh `go` để build mã Go và, nếu
cần thiết, viết lại makefile của chúng để thực hiện bất kỳ nhiệm vụ build phụ trợ nào.

## Lệnh cgo {#cmd_cgo}

Trong Go 1, [lệnh cgo](/cmd/cgo)
dùng một file `_cgo_export.h` khác,
được tạo ra cho các gói có chứa các dòng `//export`.
File `_cgo_export.h` bây giờ bắt đầu bằng comment preamble C,
để các định nghĩa hàm được xuất có thể dùng các kiểu được định nghĩa ở đó.
Điều này có tác dụng biên dịch preamble nhiều lần, vì vậy một
gói dùng `//export` không được đặt các định nghĩa hàm
hoặc khởi tạo biến trong preamble C.

## Các bản phát hành đóng gói {#releases}

Một trong những thay đổi quan trọng nhất liên quan đến Go 1 là sự có sẵn của
các bản phân phối có thể tải xuống được đóng gói sẵn.
Chúng có sẵn cho nhiều kết hợp kiến trúc và hệ điều hành
(bao gồm cả Windows) và danh sách sẽ phát triển.
Chi tiết cài đặt được mô tả trên
trang [Bắt đầu](/doc/install), trong khi
các bản phân phối chính chúng được liệt kê trên
[trang tải xuống](/dl/).
