---
title: Go 1.9 đã được phát hành
date: 2017-08-24
by:
- Francesc Campoy
summary: Go 1.9 bổ sung type alias, bit intrinsic, tối ưu hóa, và nhiều hơn nữa.
template: true
---


Hôm nay nhóm Go vui mừng thông báo phát hành Go 1.9.
Bạn có thể tải về từ [trang download](/dl/).
Có nhiều thay đổi đối với ngôn ngữ, thư viện chuẩn, runtime và công cụ.
Bài viết này đề cập các thay đổi quan trọng và dễ nhìn thấy nhất.
Phần lớn công sức kỹ thuật trong bản phát hành này dành cho cải tiến runtime và công cụ,
điều này khiến thông báo kém thú vị hơn, nhưng đây vẫn là một bản phát hành xuất sắc.

Thay đổi quan trọng nhất đối với ngôn ngữ là việc giới thiệu type alias: một tính năng
được tạo ra để hỗ trợ sửa code theo từng bước. Khai báo type alias có dạng:

	type T1 = T2

Khai báo này giới thiệu tên alias `T1` cho kiểu `T2`, giống như cách `byte`
luôn là alias của `uint8`.
[Tài liệu thiết kế type alias](/design/18130-type-alias) và
[một bài viết về refactoring](/talks/2016/refactor.article) đề cập bổ sung này chi tiết hơn.

Gói [math/bits](/pkg/math/bits) mới cung cấp các hàm đếm bit và thao tác
cho số nguyên không dấu, được triển khai bằng các lệnh CPU đặc biệt khi có thể.
Ví dụ, trên hệ thống x86-64, `bits.TrailingZeros(x)` dùng lệnh
[BSF](https://pdos.csail.mit.edu/6.828/2010/readings/i386/BSF.htm).

Gói `sync` đã thêm kiểu [Map](/pkg/sync#Map) mới, an toàn cho truy cập đồng thời.
Bạn có thể đọc thêm từ tài liệu của nó và tìm hiểu thêm lý do nó được tạo ra từ
[lightning talk GopherCon 2017](https://www.youtube.com/watch?v=C1EtfDnsdDs)
([slides](https://github.com/gophercon/2017-talks/blob/master/lightningtalks/BryanCMills-AnOverviewOfSyncMap/An%20Overview%20of%20sync.Map.pdf)).
Nó không phải là bản thay thế chung cho kiểu map của Go; xem tài liệu để biết khi nào nên dùng nó.

Gói `testing` cũng có bổ sung. Phương thức `Helper` mới, được thêm vào cả
[testing.T](/pkg/testing#T.Helper) và [testing.B](/pkg/testing#B.Helper),
đánh dấu hàm gọi là hàm hỗ trợ kiểm thử.
Khi gói testing in thông tin file và dòng, nó hiển thị vị trí của lời gọi đến hàm hỗ trợ
thay vì một dòng trong chính hàm hỗ trợ.

Ví dụ, hãy xem test này:

{{code "go1.9/helper_test.go" `/package p/` `$`}}

Vì `failure` tự xác định là hàm hỗ trợ kiểm thử, thông báo lỗi in ra trong `Test` sẽ chỉ dòng 11,
nơi `failure` được gọi, thay vì dòng 7, nơi `failure` gọi `t.Fatal`.

Gói `time` nay theo dõi một cách trong suốt thời gian đơn điệu trong mỗi giá trị `Time`,
làm cho việc tính thời gian đã trôi qua giữa hai giá trị `Time` là thao tác an toàn khi có điều chỉnh đồng hồ.
Ví dụ, code này nay tính đúng thời gian đã trôi qua ngay cả qua một lần reset đồng hồ do giây nhuận:

	start := time.Now()
	f()
	elapsed := time.Since(start)

Xem [tài liệu gói](http://beta.golang.org/pkg/time/#hdr-Monotonic_Clocks) và
[tài liệu thiết kế](https://github.com/golang/proposal/blob/master/design/12914-monotonic.md) để biết chi tiết.

Cuối cùng, như một phần trong nỗ lực làm cho trình biên dịch Go nhanh hơn, Go 1.9 biên dịch các hàm trong một gói đồng thời.

Go 1.9 bao gồm nhiều bổ sung, cải tiến và sửa lỗi hơn nữa. Tìm tập thay đổi đầy đủ,
và thêm thông tin về các cải tiến được liệt kê ở trên, trong
[ghi chú phát hành Go 1.9](/doc/go1.9).

Để kỷ niệm bản phát hành, Nhóm người dùng Go trên toàn thế giới đang tổ chức
[tiệc phát hành](https://github.com/golang/cowg/blob/master/events/2017-08-go1.9-release-party.md).
