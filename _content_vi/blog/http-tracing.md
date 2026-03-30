---
title: Giới thiệu HTTP Tracing
date: 2016-10-04
by:
- Jaana Burcu Dogan
tags:
- http
- technical
summary: Cách sử dụng HTTP tracing trong Go 1.7 để hiểu rõ hơn về các yêu cầu từ phía client.
template: true
---

## Giới thiệu

Trong Go 1.7, chúng tôi đã giới thiệu HTTP tracing, một tính năng cho phép thu thập thông tin chi tiết trong suốt vòng đời của một yêu cầu HTTP từ phía client.
Hỗ trợ cho HTTP tracing được cung cấp bởi gói [`net/http/httptrace`](/pkg/net/http/httptrace/).
Thông tin thu thập được có thể dùng để gỡ lỗi về độ trễ, giám sát dịch vụ, xây dựng các hệ thống thích ứng và nhiều mục đích khác.

## Các sự kiện HTTP

Gói `httptrace` cung cấp nhiều hook để thu thập thông tin trong quá trình thực hiện một HTTP round trip về nhiều loại sự kiện khác nhau. Các sự kiện đó bao gồm:

  - Tạo kết nối
  - Tái sử dụng kết nối
  - Tra cứu DNS
  - Ghi yêu cầu lên mạng
  - Đọc phản hồi

## Theo dõi sự kiện

Bạn có thể bật HTTP tracing bằng cách đặt một
[`*httptrace.ClientTrace`](/pkg/net/http/httptrace/#ClientTrace)
chứa các hàm hook vào [`context.Context`](/pkg/context/#Context) của yêu cầu.
Các triển khai [`http.RoundTripper`](/pkg/net/http/#RoundTripper)
báo cáo các sự kiện nội bộ bằng cách tìm kiếm `*httptrace.ClientTrace` trong context và gọi các hàm hook tương ứng.

Phạm vi theo dõi được giới hạn trong context của yêu cầu, và người dùng cần đặt `*httptrace.ClientTrace` vào context của yêu cầu trước khi bắt đầu gửi yêu cầu.

{{code "http-tracing/trace.go" `/START/` `/END/`}}

Trong quá trình round trip, `http.DefaultTransport` sẽ gọi từng hook khi có sự kiện xảy ra. Chương trình trên sẽ in ra thông tin DNS ngay khi việc tra cứu DNS hoàn tất. Tương tự, nó sẽ in thông tin kết nối khi một kết nối được thiết lập tới máy chủ của yêu cầu.

## Theo dõi với http.Client

Cơ chế theo dõi được thiết kế để theo dõi các sự kiện trong vòng đời của một lần `http.Transport.RoundTrip`. Tuy nhiên, một client có thể thực hiện nhiều round trip để hoàn thành một yêu cầu HTTP. Ví dụ, trong trường hợp chuyển hướng URL, các hook đã đăng ký sẽ được gọi nhiều lần tương ứng với số lần client theo chuyển hướng HTTP, tức là thực hiện nhiều yêu cầu.
Người dùng có trách nhiệm nhận biết các sự kiện đó ở cấp độ `http.Client`.
Chương trình dưới đây xác định yêu cầu hiện tại bằng cách dùng một wrapper `http.RoundTripper`.

{{code "http-tracing/client.go"}}

Chương trình sẽ theo chuyển hướng từ google.com đến www.google.com và in ra:

	Connection reused for https://google.com? false
	Connection reused for https://www.google.com/? false

Transport trong gói `net/http` hỗ trợ theo dõi cả yêu cầu HTTP/1 lẫn HTTP/2.

Nếu bạn là tác giả của một triển khai `http.RoundTripper` tùy chỉnh, bạn có thể hỗ trợ theo dõi bằng cách kiểm tra context của yêu cầu để tìm `*httptest.ClientTrace` và gọi các hook tương ứng khi sự kiện xảy ra.

## Kết luận

HTTP tracing là một bổ sung có giá trị cho Go dành cho những ai muốn gỡ lỗi độ trễ của yêu cầu HTTP và xây dựng các công cụ gỡ lỗi mạng cho lưu lượng gửi đi.
Bằng cách cung cấp tính năng mới này, chúng tôi hy vọng sẽ thấy các công cụ gỡ lỗi, đo hiệu năng và trực quan hóa HTTP từ cộng đồng, chẳng hạn như
[httpstat](https://github.com/davecheney/httpstat).
