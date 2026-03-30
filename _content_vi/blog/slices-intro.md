---
title: "Go Slices: cách dùng và cơ chế hoạt động bên trong"
date: 2011-01-05
by:
- Andrew Gerrand
tags:
- slice
- technical
summary: Cách sử dụng slice trong Go và cơ chế hoạt động bên trong của chúng.
template: true
---

## Giới thiệu

Kiểu slice của Go cung cấp một phương tiện tiện lợi và hiệu quả để làm việc với
các dãy dữ liệu có kiểu xác định.
Slice tương tự như mảng trong các ngôn ngữ khác,
nhưng có một số đặc điểm khác biệt.
Bài viết này sẽ xem xét slice là gì và cách chúng được sử dụng.

## Mảng

Kiểu slice là một lớp trừu tượng được xây dựng trên kiểu mảng của Go,
vì vậy để hiểu slice, trước tiên ta phải hiểu mảng.

Định nghĩa kiểu mảng chỉ định độ dài và kiểu phần tử.
Ví dụ, kiểu `[4]int` biểu diễn một mảng gồm bốn số nguyên.
Kích thước của mảng là cố định; độ dài là một phần của kiểu (`[4]int` và `[5]int` là hai kiểu
khác nhau và không tương thích với nhau).
Mảng có thể được truy cập theo chỉ số theo cách thông thường, vì vậy biểu thức `s[n]` truy cập
phần tử thứ n, bắt đầu từ không.

	var a [4]int
	a[0] = 1
	i := a[0]
	// i == 1

Mảng không cần được khởi tạo tường minh;
giá trị không (zero value) của một mảng là một mảng sẵn sàng sử dụng với các phần tử đều được khởi tạo về không:

	// a[2] == 0, giá trị không của kiểu int

Biểu diễn trong bộ nhớ của `[4]int` chỉ đơn giản là bốn giá trị nguyên được sắp xếp tuần tự:

{{image "slices-intro/slice-array.png"}}

Mảng trong Go là các giá trị. Một biến mảng đại diện cho toàn bộ mảng;
nó không phải là con trỏ tới phần tử đầu tiên của mảng (như trong C).
Điều này có nghĩa là khi bạn gán hoặc truyền một giá trị mảng, bạn sẽ tạo ra
một bản sao toàn bộ nội dung của nó.
(Để tránh sao chép, bạn có thể truyền một _con trỏ_ tới mảng,
nhưng khi đó đó là con trỏ tới mảng, không phải bản thân mảng.) Một cách nghĩ về
mảng là chúng giống như một struct nhưng với các trường được truy cập bằng chỉ số thay vì tên:
một giá trị tổng hợp có kích thước cố định.

Một mảng literal có thể được khai báo như sau:

	b := [2]string{"Penn", "Teller"}

Hoặc, bạn có thể để trình biên dịch đếm số phần tử cho bạn:

	b := [...]string{"Penn", "Teller"}

Trong cả hai trường hợp, kiểu của `b` là `[2]string`.

## Slice

Mảng có chỗ dùng của chúng, nhưng chúng khá kém linh hoạt,
nên bạn không thấy chúng xuất hiện nhiều trong mã Go.
Slice, ngược lại, xuất hiện ở khắp nơi. Chúng được xây dựng trên nền mảng để cung cấp sức mạnh và sự tiện lợi vượt trội.

Đặc tả kiểu cho một slice là `[]T`,
trong đó `T` là kiểu của các phần tử trong slice.
Khác với kiểu mảng, kiểu slice không có độ dài cụ thể.

Một slice literal được khai báo giống như một mảng literal, ngoại trừ việc bỏ qua số lượng phần tử:

	letters := []string{"a", "b", "c", "d"}

Một slice có thể được tạo bằng hàm tích hợp sẵn `make`, có chữ ký:

	func make([]T, len, cap) []T

trong đó T là kiểu phần tử của slice cần tạo.
Hàm `make` nhận một kiểu, một độ dài,
và một dung lượng tùy chọn.
Khi được gọi, `make` cấp phát một mảng và trả về một slice tham chiếu tới mảng đó.

	var s []byte
	s = make([]byte, 5, 5)
	// s == []byte{0, 0, 0, 0, 0}

Khi đối số dung lượng bị bỏ qua, nó mặc định bằng độ dài đã chỉ định.
Đây là phiên bản ngắn gọn hơn của đoạn mã trên:

	s := make([]byte, 5)

Độ dài và dung lượng của một slice có thể được kiểm tra bằng các hàm tích hợp `len` và `cap`.

	len(s) == 5
	cap(s) == 5

Hai phần tiếp theo thảo luận về mối quan hệ giữa độ dài và dung lượng.

Giá trị không của một slice là `nil`. Cả hai hàm `len` và `cap` đều trả về 0 cho một slice nil.

Một slice cũng có thể được tạo bằng cách "cắt" một slice hoặc mảng hiện có.
Việc cắt được thực hiện bằng cách chỉ định một khoảng nửa mở với hai chỉ số phân cách bởi dấu hai chấm.
Ví dụ, biểu thức `b[1:4]` tạo ra một slice bao gồm các phần tử
từ 1 đến 3 của `b` (các chỉ số của slice kết quả sẽ từ 0 đến 2).

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// b[1:4] == []byte{'o', 'l', 'a'}, chia sẻ cùng vùng lưu trữ với b

Chỉ số bắt đầu và kết thúc của biểu thức slice là tùy chọn; chúng mặc định là không và độ dài của slice tương ứng:

	// b[:2] == []byte{'g', 'o'}
	// b[2:] == []byte{'l', 'a', 'n', 'g'}
	// b[:] == b

Đây cũng là cú pháp để tạo một slice từ một mảng:

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // một slice tham chiếu tới vùng lưu trữ của x

## Cơ chế hoạt động bên trong của slice

Một slice là một bộ mô tả của một đoạn mảng.
Nó bao gồm một con trỏ tới mảng, độ dài của đoạn,
và dung lượng của nó (độ dài tối đa của đoạn).

{{image "slices-intro/slice-struct.png"}}

Biến `s` của chúng ta, được tạo trước đó bằng `make([]byte, 5)`, có cấu trúc như sau:

{{image "slices-intro/slice-1.png"}}

Độ dài là số phần tử được tham chiếu bởi slice.
Dung lượng là số phần tử trong mảng bên dưới (bắt đầu
từ phần tử được tham chiếu bởi con trỏ slice).
Sự khác biệt giữa độ dài và dung lượng sẽ được làm rõ khi chúng ta đi qua
các ví dụ tiếp theo.

Khi chúng ta cắt `s`, hãy quan sát các thay đổi trong cấu trúc dữ liệu slice và mối quan hệ của chúng với mảng bên dưới:

	s = s[2:4]

{{image "slices-intro/slice-2.png"}}

Việc cắt slice không sao chép dữ liệu của slice. Nó tạo ra một giá trị slice mới
trỏ đến mảng gốc.
Điều này làm cho các thao tác trên slice hiệu quả như việc thao tác các chỉ số mảng.
Do đó, việc sửa đổi các _phần tử_ (không phải bản thân slice) của một slice được cắt lại
sẽ sửa đổi các phần tử của slice gốc:

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}

Trước đó chúng ta đã cắt `s` xuống độ dài nhỏ hơn dung lượng của nó. Ta có thể mở rộng s đến dung lượng của nó bằng cách cắt lại:

	s = s[:cap(s)]

{{image "slices-intro/slice-3.png"}}

Một slice không thể được mở rộng vượt quá dung lượng của nó.
Cố gắng làm vậy sẽ gây ra panic lúc chạy,
cũng như khi truy cập chỉ số ngoài phạm vi của một slice hoặc mảng.
Tương tự, slice không thể được cắt lại xuống dưới không để truy cập các phần tử trước đó trong mảng.

## Mở rộng slice (hàm copy và append)

Để tăng dung lượng của một slice, ta phải tạo một slice mới lớn hơn
và sao chép nội dung của slice gốc vào đó.
Kỹ thuật này là cách các triển khai mảng động trong các ngôn ngữ khác
hoạt động đằng sau hậu trường.
Ví dụ sau nhân đôi dung lượng của `s` bằng cách tạo một slice mới `t`,
sao chép nội dung của `s` vào `t`,
và sau đó gán giá trị slice `t` cho `s`:

	t := make([]byte, len(s), (cap(s)+1)*2) // +1 trong trường hợp cap(s) == 0
	for i := range s {
	        t[i] = s[i]
	}
	s = t

Phần vòng lặp của thao tác phổ biến này được đơn giản hóa bởi hàm tích hợp copy.
Như tên gợi ý, copy sao chép dữ liệu từ một slice nguồn sang một slice đích.
Nó trả về số phần tử đã được sao chép.

	func copy(dst, src []T) int

Hàm `copy` hỗ trợ sao chép giữa các slice có độ dài khác nhau
(nó sẽ chỉ sao chép tối đa số lượng phần tử nhỏ hơn).
Ngoài ra, `copy` có thể xử lý các slice nguồn và đích chia sẻ
cùng mảng bên dưới,
xử lý đúng các slice chồng chéo nhau.

Sử dụng `copy`, chúng ta có thể đơn giản hóa đoạn mã trên:

	t := make([]byte, len(s), (cap(s)+1)*2)
	copy(t, s)
	s = t

Một thao tác thường gặp là thêm dữ liệu vào cuối slice.
Hàm này nối các phần tử byte vào một slice byte,
mở rộng slice nếu cần thiết, và trả về giá trị slice đã được cập nhật:

	func AppendByte(slice []byte, data ...byte) []byte {
	    m := len(slice)
	    n := m + len(data)
	    if n > cap(slice) { // nếu cần thiết, cấp phát lại
	        // cấp phát gấp đôi dung lượng cần thiết, để dự phòng tăng trưởng.
	        newSlice := make([]byte, (n+1)*2)
	        copy(newSlice, slice)
	        slice = newSlice
	    }
	    slice = slice[0:n]
	    copy(slice[m:n], data)
	    return slice
	}

Có thể sử dụng `AppendByte` như sau:

	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	// p == []byte{2, 3, 5, 7, 11, 13}

Các hàm như `AppendByte` hữu ích vì chúng cho phép kiểm soát hoàn toàn
cách slice được mở rộng.
Tùy thuộc vào đặc điểm của chương trình,
có thể mong muốn cấp phát theo từng khối nhỏ hoặc lớn hơn,
hoặc đặt giới hạn trên cho kích thước khi cấp phát lại.

Nhưng hầu hết các chương trình không cần kiểm soát hoàn toàn,
vì vậy Go cung cấp hàm tích hợp `append` phù hợp cho hầu hết các mục đích;
nó có chữ ký

	func append(s []T, x ...T) []T

Hàm `append` nối các phần tử `x` vào cuối slice `s`,
và mở rộng slice nếu cần thêm dung lượng.

	a := make([]int, 1)
	// a == []int{0}
	a = append(a, 1, 2, 3)
	// a == []int{0, 1, 2, 3}

Để nối một slice vào một slice khác, hãy dùng `...` để mở rộng đối số thứ hai thành danh sách đối số.

	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...) // tương đương với "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}

Vì giá trị không của một slice (`nil`) hoạt động như một slice có độ dài bằng không,
bạn có thể khai báo một biến slice và sau đó nối vào nó trong một vòng lặp:

	// Filter trả về một slice mới chỉ chứa
	// các phần tử của s thỏa mãn fn()
	func Filter(s []int, fn func(int) bool) []int {
	    var p []int // == nil
	    for _, v := range s {
	        if fn(v) {
	            p = append(p, v)
	        }
	    }
	    return p
	}

## Một bẫy tiềm ẩn

Như đã đề cập trước đó, việc cắt lại một slice không tạo bản sao của mảng bên dưới.
Toàn bộ mảng sẽ được giữ trong bộ nhớ cho đến khi nó không còn được tham chiếu nữa.
Đôi khi điều này có thể khiến chương trình giữ toàn bộ dữ liệu trong bộ nhớ trong khi
chỉ cần một phần nhỏ.

Ví dụ, hàm `FindDigits` này tải một tệp vào bộ nhớ và tìm kiếm
nhóm các chữ số liên tiếp đầu tiên trong tệp đó,
trả về chúng dưới dạng một slice mới.

	var digitRegexp = regexp.MustCompile("[0-9]+")

	func FindDigits(filename string) []byte {
	    b, _ := ioutil.ReadFile(filename)
	    return digitRegexp.Find(b)
	}

Đoạn mã này hoạt động đúng như mô tả, nhưng `[]byte` được trả về trỏ vào
mảng chứa toàn bộ nội dung tệp.
Vì slice tham chiếu tới mảng gốc,
chừng nào slice còn được giữ thì bộ gom rác không thể giải phóng mảng đó;
vài byte hữu ích của tệp giữ toàn bộ nội dung trong bộ nhớ.

Để khắc phục vấn đề này, ta có thể sao chép dữ liệu cần thiết vào một slice mới trước khi trả về:

	func CopyDigits(filename string) []byte {
	    b, _ := ioutil.ReadFile(filename)
	    b = digitRegexp.Find(b)
	    c := make([]byte, len(b))
	    copy(c, b)
	    return c
	}

Một phiên bản ngắn gọn hơn của hàm này có thể được xây dựng bằng cách sử dụng `append`.
Phần này được dành lại như bài tập cho người đọc.

## Đọc thêm

[Effective Go](/doc/effective_go.html) có phần thảo luận chuyên sâu về [slices](/doc/effective_go.html#slices)
và [arrays](/doc/effective_go.html#arrays),
và [đặc tả ngôn ngữ](/doc/go_spec.html) của Go
định nghĩa [slices](/doc/go_spec.html#Slice_types) và
các [hàm](/doc/go_spec.html#Length_and_capacity)
[hỗ trợ](/doc/go_spec.html#Making_slices_maps_and_channels)
[liên quan](/doc/go_spec.html#Appending_and_copying_slices).
