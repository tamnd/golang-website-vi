---
title: Map trong Go thực tế
date: 2013-02-06
by:
- Andrew Gerrand
tags:
- map
- technical
summary: Cách sử dụng và khi nào nên dùng map trong Go.
template: true
---

## Giới thiệu

Một trong những cấu trúc dữ liệu hữu ích nhất trong khoa học máy tính là bảng băm (hash table).
Có nhiều cài đặt bảng băm với các đặc tính khác nhau,
nhưng nhìn chung chúng đều cung cấp thao tác tìm kiếm, thêm và xóa nhanh.
Go cung cấp kiểu map tích hợp sẵn cài đặt một bảng băm.

## Khai báo và khởi tạo

Kiểu map trong Go trông như sau:

	map[KeyType]ValueType

trong đó `KeyType` có thể là bất kỳ kiểu nào [có thể so sánh được](/ref/spec#Comparison_operators)
(sẽ nói thêm về điều này sau),
và `ValueType` có thể là bất kỳ kiểu nào, kể cả một map khác!

Biến `m` này là một map có key kiểu string và giá trị kiểu int:

	var m map[string]int

Các kiểu map là kiểu tham chiếu, giống như con trỏ hay slice,
vì vậy giá trị của `m` ở trên là `nil`;
nó không trỏ đến một map đã được khởi tạo.
Map nil hoạt động giống như map rỗng khi đọc,
nhưng các thao tác ghi vào map nil sẽ gây ra runtime panic; đừng làm vậy.
Để khởi tạo một map, hãy dùng hàm tích hợp `make`:

	m = make(map[string]int)

Hàm `make` cấp phát và khởi tạo một cấu trúc dữ liệu bảng băm
và trả về một giá trị map trỏ đến nó.
Chi tiết cụ thể của cấu trúc dữ liệu đó là chi tiết cài đặt của
runtime và không được quy định bởi ngôn ngữ.
Trong bài viết này, chúng ta sẽ tập trung vào _việc sử dụng_ map,
chứ không phải cách cài đặt của chúng.

## Làm việc với map

Go cung cấp cú pháp quen thuộc để làm việc với map. Câu lệnh này gán key `"route"` giá trị `66`:

	m["route"] = 66

Câu lệnh này lấy giá trị được lưu trữ dưới key `"route"` và gán nó vào biến mới i:

	i := m["route"]

Nếu key được yêu cầu không tồn tại, chúng ta nhận được _giá trị zero_ của kiểu value.
Trong trường hợp này kiểu value là `int`, vì vậy giá trị zero là `0`:

	j := m["root"]
	// j == 0

Hàm tích hợp `len` trả về số phần tử trong map:

	n := len(m)

Hàm tích hợp `delete` xóa một mục khỏi map:

	delete(m, "route")

Hàm `delete` không trả về gì, và sẽ không làm gì nếu key được chỉ định không tồn tại.

Phép gán hai giá trị kiểm tra sự tồn tại của một key:

	i, ok := m["route"]

Trong câu lệnh này, giá trị đầu tiên (`i`) được gán giá trị được lưu trữ dưới key `"route"`.
Nếu key đó không tồn tại, `i` là giá trị zero của kiểu value (`0`).
Giá trị thứ hai (`ok`) là một `bool` có giá trị `true` nếu key tồn tại trong
map, và `false` nếu không.

Để kiểm tra key mà không lấy giá trị, dùng dấu gạch dưới thay cho giá trị đầu tiên:

	_, ok := m["route"]

Để duyệt qua nội dung của một map, dùng từ khóa `range`:

	for key, value := range m {
	    fmt.Println("Key:", key, "Value:", value)
	}

Để khởi tạo một map với một số dữ liệu, dùng map literal:

	commits := map[string]int{
	    "rsc": 3711,
	    "r":   2138,
	    "gri": 1908,
	    "adg": 912,
	}

Cú pháp tương tự có thể được sử dụng để khởi tạo một map rỗng, tương đương về chức năng với việc dùng hàm `make`:

	m = map[string]int{}

## Khai thác giá trị zero

Sẽ rất tiện lợi khi việc truy xuất map trả về giá trị zero khi key không có mặt.

Ví dụ, một map của các giá trị boolean có thể được sử dụng như cấu trúc dữ liệu dạng set
(nhớ rằng giá trị zero của kiểu boolean là false).
Ví dụ này duyệt qua một danh sách liên kết của các `Node` và in giá trị của chúng.
Nó sử dụng một map của các con trỏ `Node` để phát hiện vòng lặp trong danh sách.

{{code "maps/list.go" `/START/` `/END/`}}

Biểu thức `visited[n]` là `true` nếu `n` đã được thăm,
hoặc `false` nếu `n` không có trong map;
không cần dùng dạng hai giá trị để kiểm tra sự hiện diện của `n` trong map;
giá trị zero mặc định giúp chúng ta làm điều đó.

Một ví dụ khác về giá trị zero hữu ích là map của slice.
Việc append vào một slice nil chỉ cần cấp phát một slice mới,
vì vậy chỉ cần một dòng để append một giá trị vào map của slice;
không cần kiểm tra xem key có tồn tại không.
Trong ví dụ sau, slice people được điền với các giá trị `Person`.
Mỗi `Person` có `Name` và một slice của Likes.
Ví dụ tạo một map để liên kết mỗi sở thích với một slice của những người thích nó.

{{code "maps/people.go" `/START1/` `/END1/`}}

Để in danh sách những người thích cheese:

{{code "maps/people.go" `/START2/` `/END2/`}}

Để in số người thích bacon:

{{code "maps/people.go" `/bacon/`}}

Lưu ý rằng vì cả range lẫn len đều xử lý slice nil như slice có độ dài zero,
hai ví dụ cuối này sẽ hoạt động ngay cả khi không ai thích cheese hay bacon (dù
điều đó có vẻ khó xảy ra).

## Các kiểu key

Như đã đề cập trước đó, key của map có thể là bất kỳ kiểu nào có thể so sánh được.
[Đặc tả ngôn ngữ](/ref/spec#Comparison_operators)
định nghĩa điều này chính xác,
nhưng tóm lại, các kiểu có thể so sánh là boolean,
số, chuỗi, con trỏ, channel và kiểu interface,
cũng như struct hay array chỉ chứa các kiểu đó.
Đáng chú ý là slice, map và hàm không có trong danh sách;
các kiểu này không thể được so sánh bằng `==`,
và không thể được dùng làm key của map.

Rõ ràng là string, int và các kiểu cơ bản khác nên có thể dùng làm key map,
nhưng có lẽ bất ngờ là các struct key.
Struct có thể được dùng để key dữ liệu theo nhiều chiều.
Ví dụ, map của map này có thể được dùng để đếm lượt truy cập trang web theo quốc gia:

	hits := make(map[string]map[string]int)

Đây là map của string tới (map của `string` tới `int`).
Mỗi key của map ngoài là đường dẫn đến một trang web với map trong của riêng nó.
Mỗi key map trong là mã quốc gia hai chữ cái.
Biểu thức này lấy số lần người Úc đã tải trang tài liệu:

	n := hits["/doc/"]["au"]

Thật không may, cách tiếp cận này trở nên cồng kềnh khi thêm dữ liệu,
vì với bất kỳ key ngoài nào bạn phải kiểm tra xem map trong có tồn tại không,
và tạo nó nếu cần:

	func add(m map[string]map[string]int, path, country string) {
	    mm, ok := m[path]
	    if !ok {
	        mm = make(map[string]int)
	        m[path] = mm
	    }
	    mm[country]++
	}
	add(hits, "/doc/", "au")

Mặt khác, thiết kế dùng một map đơn với key là struct loại bỏ tất cả sự phức tạp đó:

	type Key struct {
	    Path, Country string
	}
	hits := make(map[Key]int)

Khi một người Việt Nam truy cập trang chủ,
việc tăng (và có thể tạo) bộ đếm tương ứng chỉ cần một dòng:

	hits[Key{"/", "vn"}]++

Và việc xem có bao nhiêu người Thụy Sĩ đã đọc đặc tả cũng đơn giản tương tự:

	n := hits[Key{"/ref/spec", "ch"}]

## Concurrency

[Map không an toàn cho việc sử dụng đồng thời](/doc/faq#atomic_maps):
không xác định điều gì xảy ra khi bạn đọc và ghi vào chúng đồng thời.
Nếu bạn cần đọc và ghi vào map từ các goroutine đang thực thi đồng thời,
các thao tác phải được kiểm soát bởi một số cơ chế đồng bộ hóa.
Một cách phổ biến để bảo vệ map là dùng [sync.RWMutex](/pkg/sync/#RWMutex).

Câu lệnh này khai báo một biến `counter` là một anonymous struct
chứa một map và `sync.RWMutex` nhúng vào.

	var counter = struct{
	    sync.RWMutex
	    m map[string]int
	}{m: make(map[string]int)}

Để đọc từ counter, lấy read lock:

	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)

Để ghi vào counter, lấy write lock:

	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

## Thứ tự duyệt

Khi duyệt qua map bằng vòng lặp range,
thứ tự duyệt không được chỉ định và không đảm bảo sẽ giống nhau
giữa các lần duyệt khác nhau.
Nếu bạn cần thứ tự duyệt ổn định, bạn phải duy trì một cấu trúc dữ liệu riêng chỉ định thứ tự đó.
Ví dụ này dùng một slice key đã sắp xếp riêng biệt để in một `map[int]string` theo thứ tự key:

	import "sort"

	var m map[int]string
	var keys []int
	for k := range m {
	    keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
	    fmt.Println("Key:", k, "Value:", m[k])
	}
