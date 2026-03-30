---
title: Phát triển thư viện chuẩn Go với math/rand/v2
date: 2024-05-01
by:
- Russ Cox
summary: Go 1.22 bổ sung math/rand/v2 và vạch ra hướng phát triển cho thư viện chuẩn Go.
template: true
---

Kể từ khi Go 1 được [phát hành vào tháng 3 năm 2012](/blog/go1),
các thay đổi đối với thư viện chuẩn đã bị ràng buộc bởi
[cam kết tương thích](/doc/go1compat) của Go.
Nhìn chung, tính tương thích là một lợi ích lớn cho người dùng Go,
cung cấp nền tảng ổn định cho các hệ thống production,
tài liệu, hướng dẫn, sách vở và nhiều thứ khác.
Tuy nhiên, theo thời gian, chúng tôi đã nhận ra một số sai lầm trong các API gốc
mà không thể sửa một cách tương thích; trong các trường hợp khác,
các thực hành tốt nhất và quy ước đã thay đổi.
Chúng tôi cần một kế hoạch để thực hiện những thay đổi quan trọng, gây phá vỡ tương thích.

Bài viết này nói về gói [`math/rand/v2`](/pkg/math/rand/v2/) mới trong Go 1.22,
"v2" đầu tiên trong thư viện chuẩn.
Nó mang lại những cải tiến cần thiết cho API [`math/rand`](/pkg/math/rand/),
nhưng quan trọng hơn, nó là ví dụ mẫu về cách chúng tôi có thể
sửa đổi các gói thư viện chuẩn khác khi cần thiết.

(Trong Go, `math/rand` và `math/rand/v2` là hai gói khác nhau
với các đường dẫn import khác nhau.
Go 1 và mọi bản phát hành sau đó đều bao gồm `math/rand`; Go 1.22 đã thêm `math/rand/v2`.
Một chương trình Go có thể import một trong hai gói, hoặc cả hai.)

Bài viết này thảo luận về lý do cụ thể cho các thay đổi trong `math/rand/v2`
và sau đó [phản ánh các nguyên tắc chung](#principles) sẽ hướng dẫn
các phiên bản mới của các gói khác.

## Bộ sinh số ngẫu nhiên giả {#pseudo}

Trước khi xem xét `math/rand`, vốn là một API cho bộ sinh số ngẫu nhiên giả,
hãy dành một chút thời gian để hiểu điều đó có nghĩa là gì.

Bộ sinh số ngẫu nhiên giả là một chương trình xác định
tạo ra một chuỗi dài các số
có vẻ ngẫu nhiên từ một giá trị khởi tạo nhỏ,
mặc dù các số thực ra không ngẫu nhiên chút nào.
Trong trường hợp của `math/rand`, giá trị khởi tạo là một int64 duy nhất,
và thuật toán tạo ra một chuỗi các int64
sử dụng một biến thể của
[thanh ghi dịch phản hồi tuyến tính (LFSR)](https://en.wikipedia.org/wiki/Linear-feedback_shift_register).
Thuật toán dựa trên ý tưởng của George Marsaglia,
được chỉnh sửa bởi Don Mitchell và Jim Reeds,
và được tùy chỉnh thêm bởi Ken Thompson cho Plan 9 rồi Go.
Nó không có tên chính thức, vì vậy bài viết này gọi nó là bộ sinh Go 1.

Mục tiêu là các bộ sinh này nhanh,
có thể lặp lại, và đủ ngẫu nhiên để hỗ trợ mô phỏng,
xáo trộn và các trường hợp sử dụng phi mật mã học khác.
Khả năng lặp lại đặc biệt quan trọng cho các use case như
mô phỏng số hoặc kiểm thử ngẫu nhiên.
Ví dụ, một bộ kiểm thử ngẫu nhiên có thể chọn một giá trị khởi tạo
(có thể dựa trên thời gian hiện tại), tạo
một đầu vào kiểm thử ngẫu nhiên lớn, và lặp lại.
Khi bộ kiểm thử tìm thấy một lỗi, nó chỉ cần in giá trị khởi tạo
để cho phép lặp lại kiểm thử với đầu vào lớn cụ thể đó.

Khả năng lặp lại cũng quan trọng theo thời gian: với một giá trị khởi tạo cụ thể,
một phiên bản Go mới cần tạo ra cùng
chuỗi giá trị mà phiên bản cũ đã tạo.
Chúng tôi đã không nhận ra điều này khi phát hành Go 1;
thay vào đó, chúng tôi đã học được theo cách khó,
khi chúng tôi cố gắng thực hiện thay đổi trong Go 1.2
và nhận được báo cáo rằng chúng tôi đã làm hỏng một số bài kiểm thử
và các trường hợp sử dụng khác.
Vào thời điểm đó, chúng tôi quyết định tính tương thích Go 1 bao gồm
các đầu ra ngẫu nhiên cụ thể cho một giá trị khởi tạo nhất định
và [thêm một bài kiểm thử](/change/5aca0514941ce7dd0f3cea8d8ffe627dbcd542ca).

Mục tiêu không phải là để các bộ sinh kiểu này tạo ra
các số ngẫu nhiên phù hợp để tạo khóa mật mã
hoặc các bí mật quan trọng khác.
Vì giá trị khởi tạo chỉ có 63 bit,
bất kỳ đầu ra nào từ bộ sinh, dù dài đến đâu,
cũng chỉ chứa 63 bit entropy.
Ví dụ, sử dụng `math/rand` để tạo khóa AES 128-bit hoặc 256-bit
sẽ là một sai lầm nghiêm trọng,
vì khóa sẽ dễ bị tấn công brute force hơn.
Cho loại sử dụng đó, bạn cần một
bộ sinh số ngẫu nhiên mạnh về mặt mật mã học, như được cung cấp bởi [`crypto/rand`](/pkg/crypto/rand/).

Đó là đủ nền tảng để chúng ta có thể chuyển sang những gì cần
sửa trong gói `math/rand`.

## Vấn đề với `math/rand` {#problems}

Theo thời gian, chúng tôi nhận thấy ngày càng nhiều vấn đề với `math/rand`.
Nghiêm trọng nhất là những vấn đề sau.

### Thuật toán bộ sinh {#problem.generator}

Bản thân bộ sinh cần được thay thế.

Bản cài đặt ban đầu của Go, tuy đã sẵn sàng cho môi trường production, nhưng về nhiều mặt là một "bản phác thảo"
của toàn bộ hệ thống, đủ tốt để phục vụ làm nền tảng cho phát triển trong tương lai:
trình biên dịch và runtime được viết bằng C; bộ gom rác là một bộ thu gom bảo thủ, đơn luồng,
dừng toàn bộ thế giới; và các thư viện sử dụng các cài đặt cơ bản xuyên suốt.
Từ Go 1 đến khoảng Go 1.5, chúng tôi đã quay lại và vẽ "phiên bản hoàn chỉnh"
của từng thứ: chúng tôi chuyển đổi trình biên dịch và runtime sang Go; chúng tôi đã viết một bộ gom rác mới, chính xác, song song,
đồng thời với thời gian dừng ở mức microsecond; và chúng tôi thay thế
các cài đặt thư viện chuẩn bằng các thuật toán tinh vi, được tối ưu hóa hơn
khi cần thiết.

Thật không may, yêu cầu về khả năng lặp lại trong `math/rand`
có nghĩa là chúng tôi không thể thay thế bộ sinh ở đó mà không
làm phá vỡ tính tương thích.
Chúng tôi bị kẹt với bộ sinh Go 1,
khá nhanh (khoảng 1,8ns mỗi số trên Mac M3 của tôi)
nhưng duy trì trạng thái nội bộ gần 5 kilobyte.
Ngược lại, [họ bộ sinh PCG](https://www.pcg-random.org/) của Melissa O'Neill
tạo ra các số ngẫu nhiên tốt hơn trong khoảng 2,1ns mỗi số
với chỉ 16 byte trạng thái nội bộ.
Chúng tôi cũng muốn khám phá việc sử dụng
[mã hóa luồng ChaCha](https://cr.yp.to/chacha.html) của Daniel J. Bernstein
như một bộ sinh.
Một [bài viết tiếp theo](/blog/chacha8rand) thảo luận về bộ sinh đó cụ thể.

### Giao diện Source {#problem.source}

[Giao diện `rand.Source`](/pkg/math/rand/#Source) có vấn đề.
Giao diện đó định nghĩa
khái niệm về một bộ sinh số ngẫu nhiên cấp thấp tạo ra
các giá trị `int64` không âm:

{{raw `
	% go doc -src math/rand.Source
	package rand // import "math/rand"

	// A Source represents a source of uniformly-distributed
	// pseudo-random int64 values in the range [0, 1<<63).
	//
	// A Source is not safe for concurrent use by multiple goroutines.
	type Source interface {
		Int63() int64
		Seed(seed int64)
	}

	func NewSource(seed int64) Source
	%
`}}

(Trong comment doc, "[0, N)" ký hiệu
[khoảng nửa mở](https://en.wikipedia.org/wiki/Interval_(mathematics)#Definitions_and_terminology),
có nghĩa là phạm vi bao gồm 0 nhưng kết thúc ngay trước 2⁶³.)

[Kiểu `rand.Rand`](/pkg/math/rand/#Rand) bọc một `Source`
để cài đặt một tập hợp các thao tác phong phú hơn, chẳng hạn như
tạo ra [một số nguyên từ 0 đến N](/pkg/math/rand/#Rand.Intn),
tạo ra [các số dấu phẩy động](/pkg/math/rand/#Rand.Float64), và nhiều hơn nữa.

Chúng tôi đã định nghĩa giao diện `Source` để trả về một giá trị rút gọn 63-bit
thay vì uint64 vì đó là những gì bộ sinh Go 1 và
các bộ sinh được sử dụng rộng rãi khác tạo ra,
và nó khớp với quy ước được đặt ra bởi thư viện chuẩn C.
Nhưng đây là một sai lầm: các bộ sinh hiện đại hơn tạo ra các uint64 đầy đủ chiều rộng,
đó là một giao diện thuận tiện hơn.

Một vấn đề khác là phương thức `Seed` cứng nhắc với một hạt giống `int64`:
một số bộ sinh được khởi tạo bằng các giá trị lớn hơn,
và giao diện không cung cấp cách nào để xử lý điều đó.

### Trách nhiệm khởi tạo {#problem.seed}

Một vấn đề lớn hơn với `Seed` là trách nhiệm khởi tạo bộ sinh toàn cục không rõ ràng.
Hầu hết người dùng không sử dụng `Source` và `Rand` trực tiếp.
Thay vào đó, gói `math/rand` cung cấp một bộ sinh toàn cục
được truy cập qua các hàm cấp cao như [`Intn`](/pkg/math/rand/#Intn).
Theo sau thư viện chuẩn C, bộ sinh toàn cục mặc định
hoạt động như thể `Seed(1)` được gọi khi khởi động.
Điều này tốt cho khả năng lặp lại nhưng tệ cho các chương trình muốn
đầu ra ngẫu nhiên của chúng khác nhau từ lần chạy này sang lần chạy khác.
Tài liệu gói đề xuất sử dụng `rand.Seed(time.Now().UnixNano())` trong trường hợp đó,
để làm cho đầu ra của bộ sinh phụ thuộc vào thời gian,
nhưng mã nào nên làm điều này?

Có lẽ gói main nên chịu trách nhiệm về cách `math/rand` được khởi tạo:
sẽ không hay nếu các thư viện được import tự cấu hình trạng thái toàn cục,
vì các lựa chọn của chúng có thể xung đột với các thư viện khác hoặc gói main.
Nhưng điều gì xảy ra nếu một thư viện cần một số dữ liệu ngẫu nhiên và muốn sử dụng `math/rand`?
Điều gì sẽ xảy ra nếu gói main thậm chí không biết `math/rand` đang được sử dụng?
Chúng tôi nhận thấy trên thực tế nhiều thư viện thêm các hàm init
khởi tạo bộ sinh toàn cục với thời gian hiện tại, "chỉ để chắc chắn".

Các gói thư viện tự khởi tạo bộ sinh toàn cục gây ra một vấn đề mới.
Giả sử gói main import hai gói đều sử dụng `math/rand`:
gói A giả định bộ sinh toàn cục sẽ được khởi tạo bởi gói main,
nhưng gói B khởi tạo nó trong một hàm `init`.
Và giả sử gói main không tự khởi tạo bộ sinh.
Bây giờ hoạt động đúng đắn của gói A phụ thuộc vào sự trùng hợp ngẫu nhiên là gói B cũng được
import trong chương trình.
Nếu gói main ngừng import gói B, gói A sẽ ngừng nhận các giá trị ngẫu nhiên.
Chúng tôi đã quan sát thấy điều này xảy ra trên thực tế trong các cơ sở mã lớn.

Nhìn lại, rõ ràng là sai lầm khi làm theo thư viện chuẩn C:
tự động khởi tạo bộ sinh toàn cục sẽ loại bỏ sự nhầm lẫn
về ai khởi tạo nó, và người dùng sẽ ngừng bị ngạc nhiên bởi đầu ra có thể lặp lại
khi họ không muốn điều đó.

### Khả năng mở rộng {#problem.scale}

Bộ sinh toàn cục cũng không mở rộng tốt.
Vì các hàm cấp cao như [`rand.Intn`](/pkg/math/rand/#Intn)
có thể được gọi đồng thời từ nhiều goroutine,
cài đặt cần một khóa bảo vệ trạng thái bộ sinh được chia sẻ.
Trong sử dụng song song, việc mua và giải phóng khóa này tốn kém hơn
so với việc tạo số thực sự.
Sẽ hợp lý hơn nếu có trạng thái bộ sinh per-thread,
nhưng làm như vậy sẽ làm phá vỡ khả năng lặp lại
trong các chương trình không sử dụng `math/rand` đồng thời.

### Cài đặt `Rand` thiếu các tối ưu hóa quan trọng {#problem.rand}

[Kiểu `rand.Rand`](/pkg/math/rand/#Rand) bọc một `Source`
để cài đặt một tập hợp các thao tác phong phú hơn.
Ví dụ, đây là cài đặt Go 1 của `Int63n`, trả về
một số nguyên ngẫu nhiên trong phạm vi [0, `n`).

{{raw `
	func (r *Rand) Int63n(n int64) int64 {
		if n <= 0 {
			panic("invalid argument to Int63n")
		}
		max := int64((1<<63 - 1)  - (1<<63)%uint64(n))
		v := r.Int63()
		for v > max {
			v = r.Int63()
		}
		return v % n
	}
`}}

Việc chuyển đổi thực sự rất dễ dàng: `v % n`.
Tuy nhiên, không có thuật toán nào có thể chuyển đổi 2⁶³ giá trị có xác suất bằng nhau
thành `n` giá trị có xác suất bằng nhau trừ khi 2⁶³ là bội số của `n`:
nếu không, một số đầu ra nhất định sẽ xảy ra thường xuyên hơn
so với những đầu ra khác. (Ví dụ đơn giản hơn, hãy thử chuyển đổi 4 giá trị có xác suất bằng nhau thành 3.)
Mã tính toán `max` sao cho
`max+1` là bội số lớn nhất của `n` nhỏ hơn hoặc bằng 2⁶³,
và sau đó vòng lặp loại bỏ các giá trị ngẫu nhiên lớn hơn hoặc bằng `max+1`.
Loại bỏ các giá trị quá lớn này đảm bảo tất cả `n` đầu ra có xác suất bằng nhau.
Đối với `n` nhỏ, việc phải loại bỏ bất kỳ giá trị nào là hiếm;
việc loại bỏ trở nên phổ biến hơn và quan trọng hơn đối với các giá trị lớn hơn.
Ngay cả không có vòng lặp loại bỏ, hai phép toán modulo (chậm)
có thể làm cho việc chuyển đổi tốn kém hơn so với việc tạo giá trị ngẫu nhiên `v`
ngay từ đầu.

Năm 2018, [Daniel Lemire tìm ra thuật toán](https://arxiv.org/abs/1805.10941)
tránh phép chia hầu hết thời gian
(xem thêm [bài viết blog 2019 của ông](https://lemire.me/blog/2019/06/06/nearly-divisionless-random-integer-generation-on-various-systems/)).
Trong `math/rand`, áp dụng thuật toán của Lemire sẽ làm cho `Intn(1000)` nhanh hơn 20-30%,
nhưng chúng tôi không thể: thuật toán nhanh hơn tạo ra các giá trị khác so với chuyển đổi tiêu chuẩn,
làm phá vỡ khả năng lặp lại.

Các phương thức khác cũng chậm hơn mức có thể, bị ràng buộc bởi khả năng lặp lại.
Ví dụ, phương thức `Float64` có thể dễ dàng được tăng tốc khoảng 10%
nếu chúng tôi có thể thay đổi luồng giá trị được tạo ra.
(Đây là thay đổi chúng tôi đã cố gắng thực hiện trong Go 1.2 và phải thu hồi, như đã đề cập trước đó.)

### Sai lầm `Read` {#problem.read}

Như đã đề cập trước đó, `math/rand` không được thiết kế
và không phù hợp để tạo ra các bí mật mật mã học.
Gói `crypto/rand` thực hiện điều đó, và nguyên bản cơ bản của nó
là [hàm `Read`](/pkg/crypto/rand/#Read)
và biến [`Reader`](/pkg/crypto/rand/#Reader).

Năm 2015, chúng tôi đã chấp nhận một đề xuất để làm
`rand.Rand` cài đặt `io.Reader` nữa,
cùng với [thêm hàm `Read` cấp cao](/pkg/math/rand/#Read).
Điều này có vẻ hợp lý vào thời điểm đó,
nhưng nhìn lại chúng tôi đã không chú ý đủ đến
các khía cạnh kỹ thuật phần mềm của thay đổi này.
Bây giờ, nếu bạn muốn đọc dữ liệu ngẫu nhiên, bạn có
hai lựa chọn: `math/rand.Read` và `crypto/rand.Read`.
Nếu dữ liệu sẽ được sử dụng cho tài liệu khóa,
điều rất quan trọng là phải sử dụng `crypto/rand`,
nhưng bây giờ có thể sử dụng `math/rand` thay thế,
có thể gây ra hậu quả thảm khốc.

Các công cụ như `goimports` và `gopls` có xử lý đặc biệt
để đảm bảo chúng ưu tiên sử dụng `rand.Read` từ
`crypto/rand` thay vì `math/rand`, nhưng đó không phải là một bản sửa lỗi hoàn chỉnh.
Sẽ tốt hơn nếu xóa `Read` hoàn toàn.

## Sửa `math/rand` trực tiếp {#fix.v1}

Tạo một phiên bản chính mới, không tương thích của một gói không bao giờ là lựa chọn đầu tiên của chúng tôi:
phiên bản mới đó chỉ có lợi cho các chương trình chuyển sang nó,
bỏ lại tất cả cách sử dụng hiện có của phiên bản chính cũ.
Ngược lại, việc sửa một vấn đề trong gói hiện có có tác động lớn hơn nhiều,
vì nó sửa tất cả cách sử dụng hiện có.
Chúng tôi không bao giờ nên tạo ra `v2` mà không làm hết sức để sửa `v1`.
Trong trường hợp `math/rand`, chúng tôi đã có thể giải quyết một phần
một số vấn đề được mô tả ở trên:

- Go 1.8 đã giới thiệu một [giao diện `Source64` tùy chọn](/pkg/math/rand/#Uint64) với phương thức `Uint64`.
  Nếu một `Source` cũng cài đặt `Source64`, thì `Rand` sử dụng phương thức đó
  khi thích hợp.
  Mẫu "giao diện mở rộng" này cung cấp một cách tương thích (dù hơi khó xử)
  để sửa đổi một giao diện sau thực tế.

- Go 1.20 đã tự động khởi tạo bộ sinh cấp cao và
  đánh dấu [`rand.Seed`](/pkg/math/rand/#Seed) là deprecated.
  Mặc dù điều này có vẻ như một thay đổi không tương thích
  với trọng tâm của chúng tôi về khả năng lặp lại của luồng đầu ra,
  [chúng tôi đã lý luận](/issue/56319) rằng bất kỳ gói được import nào gọi [`rand.Int`](/pkg/math/rand/#Int)
  vào thời điểm init hoặc bên trong bất kỳ tính toán nào cũng sẽ
  thay đổi luồng đầu ra một cách có thể nhìn thấy, và chắc chắn việc thêm hoặc xóa
  một cuộc gọi như vậy không thể được coi là một thay đổi gây phá vỡ.
  Và nếu điều đó đúng, thì tự động khởi tạo cũng không tệ hơn,
  và nó sẽ loại bỏ nguồn dễ vỡ này cho các chương trình trong tương lai.
  Chúng tôi cũng đã thêm [cài đặt GODEBUG](/doc/godebug) để chọn
  quay lại hành vi cũ.
  Sau đó, chúng tôi đánh dấu `rand.Seed` cấp cao là [deprecated](/wiki/Deprecated).
  (Các chương trình cần khả năng lặp lại với khởi tạo vẫn có thể sử dụng
  `rand.New(rand.NewSource(seed))` để lấy một bộ sinh cục bộ
  thay vì sử dụng bộ sinh toàn cục.)

- Sau khi loại bỏ khả năng lặp lại của luồng đầu ra toàn cục,
  Go 1.20 cũng có thể làm cho bộ sinh toàn cục mở rộng tốt hơn
  trong các chương trình không gọi `rand.Seed`,
  thay thế bộ sinh Go 1 bằng một
  [bộ sinh wyrand per-thread](https://github.com/wangyi-fudan/wyhash) rất rẻ
  đã được sử dụng bên trong runtime Go. Điều này loại bỏ mutex toàn cục
  và làm cho các hàm cấp cao mở rộng tốt hơn nhiều.
  Các chương trình gọi `rand.Seed` sẽ quay lại
  bộ sinh Go 1 được bảo vệ bằng mutex.

- Chúng tôi đã có thể áp dụng tối ưu hóa của Lemire trong runtime Go,
  và chúng tôi cũng sử dụng nó bên trong [`rand.Shuffle`](/pkg/math/rand/#Shuffle),
  được cài đặt sau khi bài báo của Lemire được xuất bản.

- Mặc dù chúng tôi không thể xóa [`rand.Read`](/pkg/math/rand/#Read) hoàn toàn,
  Go 1.20 đã đánh dấu nó là [deprecated](/wiki/Deprecated) để ưu tiên
  `crypto/rand`.
  Từ đó chúng tôi đã nghe từ những người phát hiện ra rằng họ vô tình
  sử dụng `math/rand.Read` trong bối cảnh mật mã khi trình soạn thảo của họ
  đánh dấu việc sử dụng hàm deprecated.

Những bản sửa lỗi này không hoàn hảo và không đầy đủ nhưng cũng là những cải tiến thực sự
giúp ích cho tất cả người dùng của gói `math/rand` hiện có.
Để sửa chữa hoàn chỉnh hơn, chúng tôi cần chú ý đến `math/rand/v2`.

## Sửa phần còn lại trong `math/rand/v2` {#fix.v2}

Việc định nghĩa `math/rand/v2` mất
lập kế hoạch đáng kể,
sau đó là [thảo luận GitHub](/issue/60751)
và sau đó là [thảo luận đề xuất](/issue/61716).
Nó giống
với `math/rand` với các thay đổi gây phá vỡ sau đây
giải quyết các vấn đề được nêu ở trên:

- Chúng tôi đã xóa hoàn toàn bộ sinh Go 1, thay thế bằng hai bộ sinh mới,
  [PCG](/pkg/math/rand/v2/#PCG) và [ChaCha8](/pkg/math/rand/v2/#ChaCha8).
  Các kiểu mới được đặt tên theo thuật toán của chúng (tránh tên chung chung `NewSource`)
  để nếu một thuật toán quan trọng khác cần được thêm vào, nó sẽ phù hợp
  với quy ước đặt tên.

  Áp dụng đề xuất từ thảo luận đề xuất, các kiểu mới cài đặt
  các giao diện [`encoding.BinaryMarshaler`](/pkg/encoding/#BinaryMarshaler)
  và
  [`encoding.BinaryUnmarshaler`](/pkg/encoding/#BinaryUnmarshaler).

- Chúng tôi đã thay đổi giao diện `Source`, thay thế phương thức `Int63` bằng phương thức `Uint64`
  và xóa phương thức `Seed`. Các cài đặt hỗ trợ khởi tạo có thể cung cấp
  các phương thức cụ thể của riêng chúng, như [`PCG.Seed`](/pkg/math/rand/v2/#PCG.Seed) và
  [`ChaCha8.Seed`](/pkg/math/rand/v2/#ChaCha8.Seed).
  Lưu ý rằng hai phương thức này nhận các kiểu hạt giống khác nhau, và không phải là `int64` đơn lẻ.

- Chúng tôi đã xóa hàm `Seed` cấp cao: các hàm toàn cục như `Int` bây giờ chỉ có thể được sử dụng
  ở dạng tự động khởi tạo.

- Việc xóa `Seed` cấp cao cũng cho phép chúng tôi cứng nhắc hóa việc sử dụng các bộ sinh
  có thể mở rộng per-thread bởi các phương thức cấp cao,
  tránh kiểm tra GODEBUG tại mỗi lần sử dụng.

- Chúng tôi đã cài đặt tối ưu hóa của Lemire cho `Intn` và các hàm liên quan.
  API `rand.Rand` cụ thể bây giờ được khóa trong luồng giá trị đó,
  vì vậy chúng tôi sẽ không thể tận dụng bất kỳ tối ưu hóa nào chưa được khám phá,
  nhưng ít nhất chúng tôi đã cập nhật một lần nữa.
  Chúng tôi cũng đã cài đặt các tối ưu hóa `Float32` và `Float64` mà chúng tôi muốn sử dụng từ Go 1.2.

- Trong quá trình thảo luận đề xuất, một người đóng góp đã chỉ ra thiên lệch có thể phát hiện trong
  các cài đặt của `ExpFloat64` và `NormFloat64`.
  Chúng tôi đã sửa thiên lệch đó và khóa trong các luồng giá trị mới.

- `Perm` và `Shuffle` sử dụng các thuật toán xáo trộn khác nhau và tạo ra các luồng giá trị khác nhau,
  vì `Shuffle` xuất hiện sau và sử dụng một thuật toán nhanh hơn.
  Xóa hoàn toàn `Perm` sẽ khiến việc di chuyển khó khăn hơn cho người dùng.
  Thay vào đó, chúng tôi đã cài đặt `Perm` theo cách của `Shuffle`, điều này vẫn cho phép chúng tôi
  xóa một cài đặt.

- Chúng tôi đã đổi tên `Int31`, `Int63`, `Intn`, `Int31n` và `Int63n` thành
  `Int32`, `Int64`, `IntN`, `Int32N` và `Int64N`.
  Con số 31 và 63 trong các tên không cần thiết, hay gây nhầm lẫn, và N viết hoa là chuẩn mực hơn cho từ thứ hai
  trong tên trong Go.

- Chúng tôi đã thêm các hàm và phương thức cấp cao `Uint`, `Uint32`, `Uint64`, `UintN`, `Uint32N` và `Uint64N`.
  Chúng tôi cần thêm `Uint64` để cung cấp quyền truy cập trực tiếp vào chức năng cốt lõi `Source`,
  và có vẻ không nhất quán nếu không thêm các cái khác.

- Áp dụng một đề xuất khác từ thảo luận đề xuất,
  chúng tôi đã thêm một hàm generic cấp cao mới `N` giống như
  `Int64N` hoặc `Uint64N` nhưng hoạt động cho bất kỳ kiểu số nguyên nào.
  Trong API cũ, để tạo một khoảng thời gian ngẫu nhiên lên đến 5 giây,
  cần phải viết:

      d := time.Duration(rand.Int63n(int64(5*time.Second)))

  Sử dụng `N`, mã tương đương là:

      d := rand.N(5 * time.Second)

  `N` chỉ là hàm cấp cao; không có phương thức `N` trên `rand.Rand`
  vì không có phương thức generic trong Go.
  (Các phương thức generic cũng không có khả năng xuất hiện trong tương lai;
  chúng xung đột nhiều với các giao diện, và một cài đặt hoàn chỉnh
  sẽ đòi hỏi tạo mã tại runtime hoặc thực thi chậm.)

- Để giảm thiểu việc sử dụng sai `math/rand` trong bối cảnh mật mã học,
  chúng tôi đã làm cho `ChaCha8` trở thành bộ sinh mặc định được sử dụng trong các hàm toàn cục,
  và chúng tôi cũng đã thay đổi runtime Go để sử dụng nó (thay thế wyrand).
  Các chương trình vẫn được khuyến khích mạnh mẽ sử dụng `crypto/rand`
  để tạo ra các bí mật mật mã học,
  nhưng vô tình sử dụng `math/rand/v2` không thảm khốc như
  khi sử dụng `math/rand`.
  Thậm chí trong `math/rand`, các hàm toàn cục bây giờ sử dụng bộ sinh `ChaCha8` khi không được khởi tạo rõ ràng.

## Nguyên tắc phát triển thư viện chuẩn Go {#principles}

Như đã đề cập ở đầu bài, một trong những mục tiêu cho công việc này
là thiết lập các nguyên tắc và mẫu về cách chúng tôi tiếp cận tất cả
các gói v2 trong thư viện chuẩn.
Sẽ không có sự bùng nổ các gói v2
trong vài bản phát hành Go tiếp theo.
Thay vào đó, chúng tôi sẽ xử lý từng gói
một, đảm bảo chúng tôi đặt ra một tiêu chuẩn chất lượng sẽ tồn tại thêm một thập kỷ.
Nhiều gói sẽ không cần v2 chút nào.
Nhưng đối với những gói cần, cách tiếp cận của chúng tôi tóm gọn trong ba nguyên tắc.

Thứ nhất, một phiên bản mới, không tương thích của một gói sẽ sử dụng
`that/package/v2` làm đường dẫn import của nó,
theo
[semantic import versioning](https://research.swtch.com/vgo-import)
giống như một module v2 bên ngoài thư viện chuẩn sẽ làm.
Điều này cho phép sử dụng gói gốc và gói v2
cùng tồn tại trong một chương trình duy nhất,
điều quan trọng cho một
[chuyển đổi dần dần](/talks/2016/refactor.article) sang API mới.

Thứ hai, tất cả các thay đổi phải bắt nguồn từ
sự tôn trọng đối với cách sử dụng và người dùng hiện có:
chúng tôi không được giới thiệu những thay đổi không cần thiết,
dù là dưới dạng các thay đổi không cần thiết đối với một gói hiện có hay
một gói hoàn toàn mới phải được học thay thế.
Trên thực tế, điều đó có nghĩa là chúng tôi lấy gói hiện có
làm điểm xuất phát
và chỉ thực hiện những thay đổi có căn cứ tốt
và cung cấp giá trị biện minh cho chi phí cập nhật của người dùng.

Thứ ba, gói v2 không được bỏ lại người dùng v1.
Lý tưởng nhất là gói v2 có thể làm mọi thứ mà gói v1 có thể làm,
và khi v2 được phát hành, gói v1 nên được viết lại
để trở thành một wrapper mỏng quanh v2.
Điều này sẽ đảm bảo rằng các cách sử dụng hiện có của v1 tiếp tục được hưởng lợi
từ các bản sửa lỗi và tối ưu hóa hiệu suất trong v2.
Tất nhiên, vì v2 đang giới thiệu các thay đổi gây phá vỡ,
điều này không phải lúc nào cũng có thể, nhưng đó luôn là điều cần xem xét cẩn thận.
Đối với `math/rand/v2`, chúng tôi đã sắp xếp để các hàm v1 tự động khởi tạo
gọi bộ sinh v2, nhưng chúng tôi không thể chia sẻ mã khác
do vi phạm khả năng lặp lại.
Cuối cùng `math/rand` không phải là nhiều mã và không đòi hỏi
bảo trì thường xuyên, vì vậy sự trùng lặp có thể quản lý được.
Trong các bối cảnh khác, có thể đáng công sức hơn để tránh trùng lặp.
Ví dụ, trong
[thiết kế encoding/json/v2 (vẫn đang tiến hành)](/issue/63397),
mặc dù ngữ nghĩa mặc định và API được thay đổi,
gói cung cấp các nút điều chỉnh cấu hình
giúp có thể cài đặt API v1.
Khi chúng tôi cuối cùng phát hành `encoding/json/v2`,
`encoding/json` (v1) sẽ trở thành một wrapper mỏng quanh nó,
đảm bảo người dùng không di chuyển từ v1 vẫn
được hưởng lợi từ các tối ưu hóa và sửa lỗi bảo mật trong v2.

Một [bài viết tiếp theo](/blog/chacha8rand) trình bày bộ sinh `ChaCha8` chi tiết hơn.
