---
title: HTTP/2 Server Push
date: 2017-03-24
by:
- Jaana Burcu Dogan
- Tom Bergan
tags:
- http
- technical
summary: Cách sử dụng HTTP/2 server push để giảm thời gian tải trang.
template: true
---

## Giới thiệu

HTTP/2 được thiết kế để giải quyết nhiều thiếu sót của HTTP/1.x.
Các trang web hiện đại sử dụng nhiều tài nguyên: HTML, stylesheet,
script, hình ảnh và nhiều hơn nữa. Trong HTTP/1.x, mỗi tài nguyên này phải
được yêu cầu một cách tường minh. Đây có thể là một quá trình chậm chạp.
Trình duyệt bắt đầu bằng cách lấy HTML, sau đó tìm hiểu thêm về tài nguyên
dần dần khi nó phân tích và đánh giá trang. Vì máy chủ
phải đợi trình duyệt thực hiện từng yêu cầu, mạng thường
ở trạng thái rỗi và không được tận dụng.

Để cải thiện độ trễ, HTTP/2 đã giới thiệu _server push_, cho phép
máy chủ đẩy tài nguyên đến trình duyệt trước khi chúng được yêu cầu một cách tường minh.
Máy chủ thường biết nhiều tài nguyên bổ sung mà một
trang sẽ cần và có thể bắt đầu đẩy những tài nguyên đó khi nó phản hồi
yêu cầu ban đầu. Điều này cho phép máy chủ tận dụng đầy đủ một
mạng vốn không hoạt động và cải thiện thời gian tải trang.

{{image "h2push/serverpush.svg" 600}}

Ở cấp độ giao thức, HTTP/2 server push được điều khiển bởi các frame `PUSH_PROMISE`.
Một `PUSH_PROMISE` mô tả một yêu cầu mà máy chủ dự đoán
trình duyệt sẽ thực hiện trong tương lai gần. Ngay khi trình duyệt nhận được
một `PUSH_PROMISE`, nó biết rằng máy chủ sẽ giao tài nguyên.
Nếu trình duyệt sau đó phát hiện ra rằng nó cần tài nguyên này, nó sẽ
đợi push hoàn thành thay vì gửi yêu cầu mới.
Điều này giảm thời gian trình duyệt chờ đợi trên mạng.

## Server Push trong net/http

Go 1.8 đã giới thiệu hỗ trợ đẩy phản hồi từ một [`http.Server`](/pkg/net/http/#Server).
Tính năng này có sẵn nếu máy chủ đang chạy là máy chủ HTTP/2
và kết nối đến sử dụng HTTP/2. Trong bất kỳ HTTP handler nào,
bạn có thể kiểm tra xem http.ResponseWriter có hỗ trợ server push không bằng cách kiểm tra
xem nó có triển khai interface [`http.Pusher`](/pkg/net/http/#Pusher) mới không.

Ví dụ, nếu máy chủ biết rằng `app.js` sẽ được yêu cầu để
render trang, handler có thể khởi tạo một push nếu `http.Pusher`
có sẵn:

{{code "h2push/pusher.go" `/START/` `/END/`}}

Cuộc gọi Push tạo ra một yêu cầu tổng hợp cho `/app.js`,
tổng hợp yêu cầu đó thành một frame `PUSH_PROMISE`, sau đó chuyển tiếp
yêu cầu tổng hợp đến request handler của máy chủ, handler này sẽ
tạo ra phản hồi đã được đẩy. Đối số thứ hai cho Push chỉ định
các header bổ sung cần đưa vào `PUSH_PROMISE`. Ví dụ,
nếu phản hồi cho `/app.js` thay đổi theo Accept-Encoding,
thì `PUSH_PROMISE` nên bao gồm một giá trị Accept-Encoding:

{{code "h2push/pusher.go" `/START1/` `/END1/`}}

Một ví dụ hoạt động đầy đủ [có sẵn tại đây](https://cs.opensource.google/go/x/website/+/master:_content/blog/h2push/server/).

Nếu bạn chạy máy chủ và tải [https://localhost:8080](https://localhost:8080),
công cụ dành cho nhà phát triển của trình duyệt sẽ hiển thị rằng `app.js` và
`style.css` đã được đẩy bởi máy chủ.

{{image "h2push/networktimeline.png" 605}}

## Bắt đầu Push trước khi bạn phản hồi

Đây là ý tưởng tốt khi gọi phương thức Push trước khi gửi bất kỳ byte nào
của phản hồi. Nếu không, có thể vô tình tạo ra
các phản hồi trùng lặp. Ví dụ, giả sử bạn viết một phần của phản hồi HTML:

	<html>
	<head>
		<link rel="stylesheet" href="a.css">...

Sau đó bạn gọi Push("a.css", nil). Trình duyệt có thể phân tích phần
HTML này trước khi nhận được PUSH\_PROMISE của bạn, trong trường hợp đó trình duyệt
sẽ gửi yêu cầu cho `a.css` ngoài việc nhận
`PUSH_PROMISE` của bạn. Bây giờ máy chủ sẽ tạo ra hai phản hồi cho `a.css`.
Gọi Push trước khi ghi phản hồi hoàn toàn tránh khả năng này.

## Khi nào nên sử dụng Server Push

Hãy xem xét việc sử dụng server push bất cứ lúc nào đường truyền mạng của bạn ở trạng thái rỗi.
Vừa gửi xong HTML cho web app của bạn? Đừng lãng phí thời gian chờ đợi,
hãy bắt đầu đẩy các tài nguyên mà client của bạn sẽ cần. Bạn có đang nhúng
tài nguyên vào file HTML để giảm độ trễ không? Thay vì nhúng,
hãy thử đẩy. Redirect là thời điểm tốt khác để sử dụng push vì luôn luôn
có một round trip bị lãng phí trong khi client theo dõi redirect.
Có nhiều tình huống có thể sử dụng push - chúng ta mới chỉ bắt đầu.

Chúng tôi sẽ thiếu sót nếu không đề cập đến một vài lưu ý. Đầu tiên, bạn chỉ có thể
đẩy các tài nguyên mà máy chủ của bạn là có thẩm quyền - điều này có nghĩa là bạn không thể
đẩy các tài nguyên được lưu trữ trên máy chủ của bên thứ ba hoặc CDN. Thứ hai,
đừng đẩy tài nguyên trừ khi bạn tự tin rằng chúng thực sự cần thiết
bởi client, nếu không push của bạn sẽ lãng phí băng thông. Hệ quả là
tránh đẩy tài nguyên khi client có thể đã có
những tài nguyên đó trong bộ nhớ đệm. Thứ ba, cách tiếp cận ngây thơ là đẩy tất cả
tài nguyên trên trang của bạn thường làm hiệu năng tệ hơn. Khi không chắc, hãy đo lường.

Các liên kết sau đây là tài liệu bổ sung tốt:

  - [HTTP/2 Push: The Details](https://calendar.perfplanet.com/2016/http2-push-the-details/)
  - [Innovating with HTTP/2 Server Push](https://www.igvita.com/2013/06/12/innovating-with-http-2.0-server-push/)
  - [Cache-Aware Server Push in H2O](https://github.com/h2o/h2o/issues/421)
  - [The PRPL Pattern](https://developers.google.com/web/fundamentals/performance/prpl-pattern/)
  - [Rules of Thumb for HTTP/2 Push](https://docs.google.com/document/d/1K0NykTXBbbbTlv60t5MyJvXjqKGsCVNYHyLEXIxYMv0)
  - [Server Push in the HTTP/2 spec](https://tools.ietf.org/html/rfc7540#section-8.2)

## Kết luận

Với Go 1.8, thư viện chuẩn cung cấp hỗ trợ tích hợp sẵn cho HTTP/2
Server Push, cung cấp cho bạn nhiều linh hoạt hơn để tối ưu hóa các ứng dụng web của bạn.

Truy cập trang [demo HTTP/2 Server Push](https://http2.golang.org/serverpush) của chúng tôi
để xem nó hoạt động.
