---
title: Sử dụng Subtests và Sub-benchmarks
date: 2016-10-03
by:
- Marcel van Lohuizen
tags:
- testing
- hierarchy
- table-driven
- subtests
- sub-benchmarks
summary: Cách sử dụng subtests và sub-benchmarks mới trong Go 1.7.
template: true
---

## Giới thiệu

Trong Go 1.7, gói `testing` giới thiệu phương thức Run trên kiểu
[`T`](/pkg/testing/#T.Run) và
[`B`](/pkg/testing/#B.Run)
cho phép tạo subtests và sub-benchmarks.
Việc giới thiệu subtests và sub-benchmarks cho phép xử lý lỗi tốt hơn,
kiểm soát chi tiết về các bài kiểm tra nào cần chạy từ dòng lệnh,
kiểm soát tính song song, và thường dẫn đến mã đơn giản hơn và dễ bảo trì hơn.

## Kiến thức cơ bản về table-driven tests

Trước khi đi vào chi tiết, hãy thảo luận trước về một cách phổ biến
để viết tests trong Go.
Một loạt các kiểm tra liên quan có thể được triển khai bằng cách lặp qua một slice của
các test cases:

	func TestTime(t *testing.T) {
		testCases := []struct {
			gmt  string
			loc  string
			want string
		}{
			{"12:31", "Europe/Zuri", "13:31"},     // incorrect location name
			{"12:31", "America/New_York", "7:31"}, // should be 07:31
			{"08:08", "Australia/Sydney", "18:08"},
		}
		for _, tc := range testCases {
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatalf("could not load location %q", tc.loc)
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("In(%s, %s) = %s; want %s", tc.gmt, tc.loc, got, tc.want)
			}
		}
	}

Cách tiếp cận này, thường được gọi là table-driven tests, giảm lượng
mã lặp lại so với việc lặp lại cùng một mã cho mỗi test
và giúp dễ dàng thêm nhiều test cases hơn.

## Table-driven benchmarks

Trước Go 1.7, không thể sử dụng cùng cách tiếp cận table-driven cho
benchmarks.
Một benchmark kiểm tra hiệu suất của toàn bộ hàm, vì vậy việc lặp qua
các benchmarks sẽ chỉ đo tất cả chúng như một benchmark duy nhất.

Một giải pháp tạm thời phổ biến là định nghĩa các benchmark cấp cao nhất riêng biệt
mà mỗi cái gọi một hàm chung với các tham số khác nhau.
Ví dụ, trước phiên bản 1.7, các benchmarks của gói `strconv` cho `AppendFloat`
trông giống như thế này:

{{raw `
	func benchmarkAppendFloat(b *testing.B, f float64, fmt byte, prec, bitSize int) {
		dst := make([]byte, 30)
		b.ResetTimer() // Overkill here, but for illustrative purposes.
		for i := 0; i < b.N; i++ {
			AppendFloat(dst[:0], f, fmt, prec, bitSize)
		}
	}

	func BenchmarkAppendFloatDecimal(b *testing.B) { benchmarkAppendFloat(b, 33909, 'g', -1, 64) }
	func BenchmarkAppendFloat(b *testing.B)        { benchmarkAppendFloat(b, 339.7784, 'g', -1, 64) }
	func BenchmarkAppendFloatExp(b *testing.B)     { benchmarkAppendFloat(b, -5.09e75, 'g', -1, 64) }
	func BenchmarkAppendFloatNegExp(b *testing.B)  { benchmarkAppendFloat(b, -5.11e-95, 'g', -1, 64) }
	func BenchmarkAppendFloatBig(b *testing.B)     { benchmarkAppendFloat(b, 123456789123456789123456789, 'g', -1, 64) }
	...
`}}

Sử dụng phương thức `Run` có trong Go 1.7, cùng tập hợp benchmarks bây giờ được
biểu diễn như một benchmark cấp cao nhất duy nhất:

{{raw `
	func BenchmarkAppendFloat(b *testing.B) {
		benchmarks := []struct{
			name    string
			float   float64
			fmt     byte
			prec    int
			bitSize int
		}{
			{"Decimal", 33909, 'g', -1, 64},
			{"Float", 339.7784, 'g', -1, 64},
			{"Exp", -5.09e75, 'g', -1, 64},
			{"NegExp", -5.11e-95, 'g', -1, 64},
			{"Big", 123456789123456789123456789, 'g', -1, 64},
			...
		}
		dst := make([]byte, 30)
		for _, bm := range benchmarks {
			b.Run(bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					AppendFloat(dst[:0], bm.float, bm.fmt, bm.prec, bm.bitSize)
				}
			})
		}
	}
`}}

Mỗi lần gọi phương thức `Run` tạo ra một benchmark riêng biệt.
Một hàm benchmark bao quanh gọi phương thức `Run` chỉ chạy một lần và
không được đo.

Mã mới có nhiều dòng code hơn, nhưng dễ bảo trì hơn, dễ đọc hơn,
và nhất quán với cách tiếp cận table-driven thường được sử dụng để kiểm tra.
Hơn nữa, mã thiết lập chung bây giờ được chia sẻ giữa các lần chạy trong khi loại bỏ
nhu cầu đặt lại timer.

## Table-driven tests sử dụng subtests

Go 1.7 cũng giới thiệu phương thức `Run` để tạo subtests.
Test này là phiên bản được viết lại của ví dụ trước đó của chúng ta sử dụng subtests:

	func TestTime(t *testing.T) {
		testCases := []struct {
			gmt  string
			loc  string
			want string
		}{
			{"12:31", "Europe/Zuri", "13:31"},
			{"12:31", "America/New_York", "7:31"},
			{"08:08", "Australia/Sydney", "18:08"},
		}
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
				loc, err := time.LoadLocation(tc.loc)
				if err != nil {
					t.Fatal("could not load location")
				}
				gmt, _ := time.Parse("15:04", tc.gmt)
				if got := gmt.In(loc).Format("15:04"); got != tc.want {
					t.Errorf("got %s; want %s", got, tc.want)
				}
			})
		}
	}

Điều đầu tiên cần chú ý là sự khác biệt về đầu ra giữa hai cách triển khai.
Cách triển khai gốc in ra:

	--- FAIL: TestTime (0.00s)
		time_test.go:62: could not load location "Europe/Zuri"

Mặc dù có hai lỗi, việc thực thi test dừng lại khi gọi
`Fatalf` và test thứ hai không bao giờ chạy.

Cách triển khai sử dụng `Run` in ra cả hai:

	--- FAIL: TestTime (0.00s)
	    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
	    	time_test.go:84: could not load location
	    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
	    	time_test.go:88: got 07:31; want 7:31

`Fatal` và các hàm anh em của nó khiến một subtest bị bỏ qua nhưng không phải
parent của nó hoặc các subtests tiếp theo.

Một điều khác cần chú ý là các thông báo lỗi ngắn hơn trong cách triển khai mới.
Vì tên subtest xác định duy nhất subtest, không cần phải
xác định lại test trong các thông báo lỗi.

Có một số lợi ích khác khi sử dụng subtests hoặc sub-benchmarks,
như được làm rõ trong các phần sau.

## Chạy các test hoặc benchmark cụ thể

Cả subtests và sub-benchmarks đều có thể được chọn riêng lẻ trên dòng lệnh bằng cách sử dụng
[cờ `-run` hoặc `-bench`](/cmd/go/#hdr-Description_of_testing_flags).
Cả hai cờ đều nhận một danh sách các biểu thức chính quy được phân tách bằng dấu gạch chéo
khớp với các phần tương ứng của tên đầy đủ của subtest hoặc sub-benchmark.

Tên đầy đủ của một subtest hoặc sub-benchmark là một danh sách được phân tách bằng dấu gạch chéo của
tên của nó và tên của tất cả các parent của nó, bắt đầu từ cấp cao nhất.
Tên là tên hàm tương ứng cho các tests và benchmarks cấp cao nhất,
và là đối số đầu tiên cho `Run` trong các trường hợp khác.
Để tránh các vấn đề hiển thị và phân tích cú pháp, một tên được làm sạch bằng cách thay thế các khoảng trắng
bằng dấu gạch dưới và thoát các ký tự không in được.
Làm sạch tương tự cũng được áp dụng cho các biểu thức chính quy được truyền vào
cờ `-run` hoặc `-bench`.

Một vài ví dụ:

Chạy các tests sử dụng múi giờ ở Châu Âu:

	$ go test -run=TestTime/"in Europe"
	--- FAIL: TestTime (0.00s)
	    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
	    	time_test.go:85: could not load location

Chỉ chạy các tests cho các giờ sau buổi trưa:

	$ go test -run=Time/12:[0-9] -v
	=== RUN   TestTime
	=== RUN   TestTime/12:31_in_Europe/Zuri
	=== RUN   TestTime/12:31_in_America/New_York
	--- FAIL: TestTime (0.00s)
	    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
	    	time_test.go:85: could not load location
	    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
	    	time_test.go:89: got 07:31; want 7:31

Có thể hơi bất ngờ, việc sử dụng `-run=TestTime/New_York` sẽ không khớp với bất kỳ test nào.
Đó là vì dấu gạch chéo có trong tên địa điểm cũng được coi là
một dấu phân tách.
Thay vào đó hãy sử dụng:

	$ go test -run=Time//New_York
	--- FAIL: TestTime (0.00s)
	    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
	    	time_test.go:88: got 07:31; want 7:31

Chú ý `//` trong chuỗi được truyền vào `-run`.
Dấu `/` trong tên múi giờ `America/New_York` được xử lý như thể nó là
một dấu phân tách xuất phát từ một subtest.
Biểu thức chính quy đầu tiên của mẫu (`TestTime`) khớp với test
cấp cao nhất.
Biểu thức chính quy thứ hai (chuỗi rỗng) khớp với bất cứ thứ gì, trong trường hợp này
là giờ và phần lục địa của địa điểm.
Biểu thức chính quy thứ ba (`New_York`) khớp với phần thành phố của địa điểm.

Xử lý các dấu gạch chéo trong tên như dấu phân tách cho phép người dùng tái cấu trúc
phân cấp của các tests mà không cần thay đổi cách đặt tên.
Nó cũng đơn giản hóa các quy tắc thoát.
Người dùng nên thoát các dấu gạch chéo trong tên, ví dụ bằng cách thay thế chúng bằng
dấu gạch chéo ngược, nếu điều này gây ra vấn đề.

Một số thứ tự duy nhất được thêm vào tên tests không phải là duy nhất.
Vì vậy, người ta có thể chỉ truyền một chuỗi rỗng vào `Run`
nếu không có sơ đồ đặt tên rõ ràng cho subtests và subtests
có thể được xác định dễ dàng theo số thứ tự của chúng.

## Thiết lập và Dọn dẹp

Subtests và sub-benchmarks có thể được sử dụng để quản lý mã thiết lập và dọn dẹp chung:

{{raw `
	func TestFoo(t *testing.T) {
		// <setup code>
		t.Run("A=1", func(t *testing.T) { ... })
		t.Run("A=2", func(t *testing.T) { ... })
		t.Run("B=1", func(t *testing.T) {
			if !test(foo{B:1}) {
				t.Fail()
			}
		})
		// <tear-down code>
	}
`}}

Mã thiết lập và dọn dẹp sẽ chạy nếu bất kỳ subtests nào được bao gồm chạy
và sẽ chạy tối đa một lần.
Điều này áp dụng ngay cả khi bất kỳ subtest nào gọi `Skip`, `Fail`, hoặc `Fatal`.

## Kiểm soát Tính song song

Subtests cho phép kiểm soát chi tiết đối với tính song song.
Để hiểu cách sử dụng subtests theo cách đó,
điều quan trọng là phải hiểu ngữ nghĩa của các parallel tests.

Mỗi test được liên kết với một hàm test.
Một test được gọi là parallel test nếu hàm test của nó gọi phương thức Parallel
trên instance `testing.T` của nó.
Một parallel test không bao giờ chạy đồng thời với một sequential test và việc thực thi
của nó bị tạm dừng cho đến khi hàm test gọi nó, tức là hàm test của parent test,
đã trả về.
Cờ `-parallel` định nghĩa số lượng tối đa các parallel tests có thể chạy
song song.

Một test bị chặn cho đến khi hàm test của nó trả về và tất cả các subtests của nó
đã hoàn thành.
Điều này có nghĩa là các parallel tests được chạy bởi một sequential test sẽ
hoàn thành trước khi bất kỳ sequential test liên tiếp nào khác chạy.

Hành vi này giống hệt nhau đối với các tests được tạo bởi `Run` và các tests cấp cao nhất.
Trên thực tế, dưới bề mặt, các tests và benchmarks cấp cao nhất được triển khai như các subtests và
sub-benchmarks của một master test ẩn.

### Chạy một nhóm tests song song

Ngữ nghĩa ở trên cho phép chạy một nhóm tests song song với
nhau nhưng không với các parallel tests khác:

	func TestGroupedParallel(t *testing.T) {
		for _, tc := range testCases {
			tc := tc // capture range variable
			t.Run(tc.Name, func(t *testing.T) {
				t.Parallel()
				if got := foo(tc.in); got != tc.out {
					t.Errorf("got %v; want %v", got, tc.out)
				}
				...
			})
		}
	}

Test bên ngoài sẽ không hoàn thành cho đến khi tất cả các parallel tests được bắt đầu bởi `Run`
đã hoàn thành.
Kết quả là, không có parallel test nào khác có thể chạy song song với các parallel tests này.

Lưu ý rằng chúng ta cần lấy biến range để đảm bảo rằng `tc` được gắn vào
instance đúng.

### Dọn dẹp sau một nhóm parallel tests

Trong ví dụ trước, chúng ta đã sử dụng ngữ nghĩa để chờ một nhóm parallel
tests hoàn thành trước khi tiếp tục các tests khác.
Kỹ thuật tương tự có thể được sử dụng để dọn dẹp sau một nhóm parallel tests
chia sẻ tài nguyên chung:

{{raw `
	func TestTeardownParallel(t *testing.T) {
		// <setup code>
		// This Run will not return until its parallel subtests complete.
		t.Run("group", func(t *testing.T) {
			t.Run("Test1", parallelTest1)
			t.Run("Test2", parallelTest2)
			t.Run("Test3", parallelTest3)
		})
		// <tear-down code>
	}
`}}

Hành vi chờ một nhóm parallel tests giống hệt với
ví dụ trước.

## Kết luận

Việc bổ sung subtests và sub-benchmarks trong Go 1.7 cho phép bạn viết tests
và benchmarks có cấu trúc theo cách tự nhiên kết hợp tốt vào các công cụ hiện có.
Một cách để nghĩ về điều này là các phiên bản cũ hơn của gói testing có
phân cấp 1 cấp: package-level test được cấu trúc như một tập hợp
các tests và benchmarks riêng lẻ.
Bây giờ cấu trúc đó đã được mở rộng đến những tests và benchmarks riêng lẻ đó,
theo cách đệ quy.
Trên thực tế, trong triển khai, các tests và benchmarks cấp cao nhất được theo dõi
như thể chúng là subtests và sub-benchmarks của một master test và
benchmark ngầm: cách xử lý thực sự là giống nhau ở mọi cấp độ.

Khả năng của tests để định nghĩa cấu trúc này cho phép thực thi chi tiết
các test cases cụ thể, thiết lập và dọn dẹp chung, và kiểm soát tốt hơn việc song song hóa tests.
Chúng tôi rất mong được thấy những cách sử dụng khác mà mọi người tìm ra. Hãy tận hưởng.
