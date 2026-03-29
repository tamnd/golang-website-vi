---
title: Go 1.22 đã được phát hành!
date: 2024-02-06
by:
- Eli Bendersky, on behalf of the Go team
summary: Go 1.22 cải tiến vòng lặp for, mang đến chức năng thư viện chuẩn mới và cải thiện hiệu năng.
template: true
---

Hôm nay nhóm Go vô cùng hào hứng phát hành Go 1.22,
bạn có thể tải về từ [trang download](/dl/).

Go 1.22 đi kèm với một số tính năng và cải tiến quan trọng. Dưới đây là
một số thay đổi đáng chú ý; để xem danh sách đầy đủ, tham khảo [ghi chú
phát hành](/doc/go1.22).

## Thay đổi ngôn ngữ

Vấn đề lâu năm về vòng lặp "for" với việc vô tình chia sẻ biến vòng lặp
giữa các lần lặp đã được giải quyết. Từ Go 1.22, đoạn code dưới đây
sẽ in "a", "b" và "c" theo thứ tự nào đó:

{{raw `
	func main() {
		done := make(chan bool)

		values := []string{"a", "b", "c"}
		for _, v := range values {
			go func() {
				fmt.Println(v)
				done <- true
			}()
		}

		// chờ tất cả goroutine hoàn thành trước khi thoát
		for _ = range values {
			<-done
		}
	}
`}}

Để biết thêm thông tin về thay đổi này và các công cụ giúp giữ code không bị hỏng
vô tình, xem [bài đăng blog về biến vòng lặp](/blog/loopvar-preview) trước đó.

Thay đổi ngôn ngữ thứ hai là hỗ trợ range trên số nguyên:

{{raw `
	package main

	import "fmt"

	func main() {
		for i := range 10 {
			fmt.Println(10 - i)
		}
		fmt.Println("go1.22 has lift-off!")
	}
`}}

Các giá trị của `i` trong chương trình đếm ngược này từ 0 đến 9, bao gồm. Để biết thêm
chi tiết, tham khảo [đặc tả](/ref/spec#For_range).

## Cải thiện hiệu năng

Tối ưu hóa bộ nhớ trong Go runtime cải thiện hiệu năng CPU từ 1-3%, trong khi
cũng giảm bộ nhớ sử dụng của hầu hết các chương trình Go khoảng 1%.

Trong Go 1.21, [chúng tôi đã đưa vào](/blog/pgo) tối ưu hóa dựa trên hồ sơ thực thi (PGO) cho trình biên dịch Go
và chức năng này tiếp tục được cải thiện. Một trong những tối ưu hóa
được thêm trong 1.22 là devirtualization cải tiến, cho phép dispatch tĩnh cho nhiều
lời gọi phương thức interface hơn. Hầu hết các chương trình sẽ thấy cải tiến từ 2-14% với
PGO bật.

## Bổ sung thư viện chuẩn

- Gói [math/rand/v2](/pkg/math/rand/v2) mới
  cung cấp API gọn gàng, nhất quán hơn và dùng các thuật toán tạo số giả ngẫu nhiên
  chất lượng cao hơn, nhanh hơn. Xem
  [đề xuất](/issue/61716) để biết thêm chi tiết.
- Các pattern dùng bởi [net/http.ServeMux](/pkg/net/http#ServeMux)
  nay chấp nhận method và wildcard.

  Ví dụ, router chấp nhận pattern như `GET /task/{id}/`, khớp với
  chỉ các request `GET` và bắt giữ giá trị của segment `{id}`
  trong một map có thể truy cập qua giá trị [Request](/pkg/net/http#Request).
- Kiểu `Null[T]` mới trong [database/sql](/pkg/database/sql) cung cấp
  cách scan các cột nullable.
- Hàm `Concat` được thêm vào gói [slices](/pkg/slices), để
  nối nhiều slice của bất kỳ kiểu nào.

---

Cảm ơn tất cả những người đã đóng góp cho bản phát hành này bằng cách viết code và
tài liệu, báo cáo lỗi, chia sẻ phản hồi, và kiểm thử các release
candidate. Nỗ lực của bạn giúp đảm bảo Go 1.22 ổn định nhất có thể.
Như thường lệ, nếu bạn phát hiện bất kỳ vấn đề nào, hãy [tạo issue](/issue/new).

Tận hưởng Go 1.22!
