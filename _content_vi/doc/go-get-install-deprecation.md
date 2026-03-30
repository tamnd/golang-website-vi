<!--{
  "Title": "Ngừng hỗ trợ 'go get' để cài đặt tệp thực thi",
  "Path": "/doc/go-get-install-deprecation",
  "Breadcrumb": true
}-->

## Tổng quan

Bắt đầu từ Go 1.17, việc cài đặt tệp thực thi bằng `go get` đã bị ngừng hỗ trợ.
Thay vào đó, bạn có thể dùng `go install`.

Trong Go 1.18, `go get` sẽ không còn biên dịch gói nữa; lệnh này chỉ được dùng để thêm, cập nhật hoặc xóa dependency trong `go.mod`. Cụ thể, `go get` sẽ luôn hoạt động như khi cờ `-d` được bật.

## Nên dùng gì thay thế

Để cài đặt một tệp thực thi trong ngữ cảnh module hiện tại, dùng `go install` không kèm hậu tố phiên bản, như ví dụ dưới đây. Lệnh này áp dụng các yêu cầu phiên bản và các chỉ thị khác từ tệp `go.mod` trong thư mục hiện tại hoặc thư mục cha.

```
go install example.com/cmd
```

Để cài đặt một tệp thực thi mà không quan tâm đến module hiện tại, dùng `go install` *kèm* [hậu tố phiên bản](/ref/mod#version-queries) như `@v1.2.3` hoặc `@latest`, như ví dụ dưới đây. Khi dùng kèm hậu tố phiên bản, `go install` không đọc hoặc cập nhật tệp `go.mod` trong thư mục hiện tại hoặc thư mục cha.

```
# Cài đặt một phiên bản cụ thể.
go install example.com/cmd@v1.2.3

# Cài đặt phiên bản mới nhất hiện có.
go install example.com/cmd@latest
```

Để tránh nhập nhằng, khi `go install` được dùng kèm hậu tố phiên bản, tất cả các đối số phải trỏ đến gói `main` trong cùng một module và cùng một phiên bản. Nếu module đó có tệp `go.mod`, tệp này không được chứa các chỉ thị như `replace` hoặc `exclude` mà có thể làm nó được hiểu khác đi nếu nó là module chính. Thư mục `vendor` của module không được sử dụng.

Xem [`go install`](/ref/mod#go-install) để biết thêm chi tiết.

## Lý do thay đổi

Kể từ khi modules được giới thiệu, lệnh `go get` được dùng cho cả hai mục đích: cập nhật dependency trong `go.mod` và cài đặt lệnh. Sự kết hợp này thường gây nhầm lẫn và bất tiện: trong hầu hết các trường hợp, lập trình viên muốn cập nhật một dependency hoặc cài đặt một lệnh, nhưng không phải cả hai cùng lúc.

Kể từ Go 1.16, `go install` có thể cài đặt một lệnh tại một phiên bản được chỉ định trên dòng lệnh trong khi bỏ qua tệp `go.mod` trong thư mục hiện tại (nếu có). `go install` giờ đây nên được dùng để cài đặt lệnh trong hầu hết các trường hợp.

Khả năng biên dịch và cài đặt lệnh của `go get` giờ đây đã bị ngừng hỗ trợ, vì chức năng đó trùng lặp với `go install`. Loại bỏ chức năng này sẽ giúp `go get` nhanh hơn, vì lệnh này sẽ không còn biên dịch hoặc liên kết gói theo mặc định. Ngoài ra, `go get` cũng sẽ không báo lỗi khi cập nhật một gói không thể biên dịch cho nền tảng hiện tại.

Xem đề xuất [#40276](/issue/40276) để biết toàn bộ thảo luận.
