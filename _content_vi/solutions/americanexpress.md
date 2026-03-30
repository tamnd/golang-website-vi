---
title: "American Express dùng Go cho thanh toán & điểm thưởng"
company: American Express
logoSrc: american-express.svg
logoSrcDark: american-express.svg
heroImgSrc: go_amex_case_study_logo.png
carouselImgSrc: go_amex_case_study.png
date: 2019-12-19
series: Case Studies
template: true
quote: Go cung cấp cho American Express tốc độ và khả năng mở rộng cần thiết cho cả mạng thanh toán và mạng điểm thưởng.
---

{{pullquote `
  author: Glen Balliet
  title: Engineering Director of loyalty platforms
  company: American Express
  quote: |
    What makes Go different from other programming languages is cognitive load. You can do more with less code, which makes it easier to reason about and understand the code that you do end up writing.

    The majority of Go code ends up looking quite similar, so, even if you're working with a completely new codebase, you can get up and running pretty quickly.
`}}

## Go cải thiện microservice và tăng năng suất

Được thành lập năm 1850, American Express là công ty thanh toán tích hợp toàn cầu cung cấp sản phẩm thẻ tín dụng và thẻ ghi nợ, dịch vụ mua bán và xử lý cho merchant, dịch vụ mạng và dịch vụ du lịch.

Hệ thống xử lý thanh toán của American Express được phát triển qua lịch sử lâu dài và đã được cập nhật qua nhiều thế hệ kiến trúc. Quan trọng nhất trong mọi lần cập nhật, xử lý thanh toán cần phải nhanh, đặc biệt ở khối lượng giao dịch rất lớn, với khả năng phục hồi được xây dựng xuyên suốt các hệ thống phải tuân thủ các tiêu chuẩn bảo mật và quy định. Với Go, American Express đạt được tốc độ và khả năng mở rộng cần thiết cho cả mạng thanh toán và điểm thưởng.

### Hiện đại hóa hệ thống American Express

American Express hiểu rằng bối cảnh ngôn ngữ lập trình đang thay đổi mạnh mẽ. Các hệ thống hiện có của công ty được xây dựng có mục đích cho tính đồng thời cao và độ trễ thấp, nhưng biết rằng những hệ thống đó sẽ được tái nền tảng hóa trong tương lai gần. Nhóm nền tảng thanh toán quyết định dành thời gian xác định ngôn ngữ nào lý tưởng nhất cho nhu cầu phát triển của American Express.

Các nhóm nền tảng thanh toán và điểm thưởng tại American Express nằm trong số những nhóm đầu tiên bắt đầu đánh giá Go. Các nhóm này tập trung vào microservice, định tuyến giao dịch và cân bằng tải, và cần hiện đại hóa kiến trúc. Nhiều lập trình viên American Express quen với khả năng của ngôn ngữ và muốn thử nghiệm Go cho các ứng dụng đồng thời cao và độ trễ thấp (như load balancer giao dịch tùy chỉnh). Với mục tiêu này, các nhóm bắt đầu vận động lãnh đạo cấp cao để triển khai Go trên nền tảng thanh toán American Express.

"Chúng tôi muốn tìm ngôn ngữ tối ưu để viết ứng dụng nhanh và hiệu quả cho xử lý thanh toán," Benjamin Cane, phó chủ tịch và kỹ sư chính tại American Express, cho biết. "Để làm vậy, chúng tôi bắt đầu một cuộc đua ngôn ngữ lập trình nội bộ với mục tiêu xem ngôn ngữ nào phù hợp nhất với nhu cầu thiết kế và hiệu năng của chúng tôi."

### So sánh ngôn ngữ

Để đánh giá, nhóm của Cane chọn xây dựng một microservice bằng bốn ngôn ngữ lập trình khác nhau. Sau đó so sánh bốn ngôn ngữ về tốc độ/hiệu năng, công cụ, kiểm thử và dễ phát triển.

Với dịch vụ, họ quyết định dùng bộ chuyển đổi ISO8583 sang JSON. ISO8583 là tiêu chuẩn quốc tế cho giao dịch tài chính và được sử dụng phổ biến trong mạng thanh toán của American Express. Về ngôn ngữ lập trình, họ chọn so sánh C++, Go, Java và Node.js. Ngoại trừ Go, tất cả các ngôn ngữ này đã được sử dụng tại American Express.

Về tốc độ, Go đạt hiệu năng đứng thứ hai với 140.000 request mỗi giây. Go cho thấy nó vượt trội khi dùng cho microservice backend.

Mặc dù Go có thể không phải ngôn ngữ nhanh nhất được kiểm thử, nhưng công cụ mạnh mẽ của nó giúp củng cố kết quả tổng thể. Framework kiểm thử tích hợp, khả năng profiling và công cụ benchmarking của Go đã gây ấn tượng với nhóm. "Dễ viết test hiệu quả trong Go," Cane nói. "Các tính năng benchmarking và profiling giúp đơn giản hóa việc tối ưu ứng dụng. Kết hợp với thời gian build nhanh, Go giúp dễ dàng viết code được kiểm thử tốt và được tối ưu."

Cuối cùng, Go được nhóm chọn là ngôn ngữ ưu tiên để xây dựng microservice hiệu năng cao. Công cụ, framework kiểm thử, hiệu năng và sự đơn giản của ngôn ngữ đều là các yếu tố đóng góp chính.

### Go cho cơ sở hạ tầng

"Nhiều dịch vụ của chúng tôi đang chạy trong container Docker trong nền tảng đám mây nội bộ dựa trên Kubernetes," Cane cho biết. Kubernetes là hệ thống điều phối container mã nguồn mở viết bằng Go. Nó cung cấp các cụm máy chủ để chạy khối lượng công việc dựa trên container, đặc biệt là container Docker. Docker là sản phẩm phần mềm, cũng viết bằng Go, sử dụng ảo hóa cấp hệ điều hành để cung cấp môi trường thực thi phần mềm di động gọi là container.

American Express cũng thu thập số liệu ứng dụng thông qua Prometheus, bộ công cụ giám sát và cảnh báo mã nguồn mở viết bằng Go. Prometheus thu thập và tổng hợp sự kiện và số liệu thời gian thực để giám sát và cảnh báo.

Ba giải pháp Go này, Kubernetes, Docker và Prometheus, đã giúp hiện đại hóa cơ sở hạ tầng American Express.

### Cải thiện hiệu năng với Go

Ngày nay, hàng chục lập trình viên đang lập trình với Go tại American Express, với hầu hết làm việc trên các nền tảng được thiết kế cho tính khả dụng cao và hiệu năng cao.

"Công cụ luôn là lĩnh vực nhu cầu quan trọng cho codebase cũ của chúng tôi," Cane nói. "Chúng tôi thấy Go có công cụ xuất sắc, cộng với framework kiểm thử, benchmarking và profiling tích hợp. Dễ viết ứng dụng hiệu quả và có khả năng phục hồi."

{{backgroundquote `
  author: Benjamin Cane
  title: Vice President and Principal Engineer
  company: American Express
  quote: |
    After working on Go, most of our developers don't want to go back to other languages.
`}}

American Express chỉ mới bắt đầu thấy lợi ích của Go. Ví dụ, Go được thiết kế từ đầu với tính đồng thời trong tâm trí, sử dụng "goroutine" nhẹ hơn luồng hệ điều hành nặng hơn, giúp thực tế để tạo hàng trăm nghìn goroutine trong cùng không gian địa chỉ. Dùng goroutine, American Express đã thấy cải thiện về số liệu hiệu năng trong xử lý giao dịch thời gian thực.

Bộ gom rác của Go cũng là cải tiến lớn so với các ngôn ngữ khác, cả về hiệu năng và dễ phát triển. "Chúng tôi thấy kết quả bộ gom rác trong Go tốt hơn nhiều so với các ngôn ngữ khác, và bộ gom rác cho xử lý giao dịch thời gian thực là điều quan trọng," Cane nói. "Điều chỉnh bộ gom rác trong các ngôn ngữ khác có thể rất phức tạp. Với Go bạn không cần điều chỉnh gì cả."

Để tìm hiểu thêm, đọc ["Choosing Go at American Express"](https://americanexpress.io/choosing-go/), đi sâu hơn về việc American Express áp dụng Go.

### Bắt đầu với Go cho doanh nghiệp của bạn

Giống như American Express dùng Go để hiện đại hóa mạng thanh toán và điểm thưởng, hàng chục doanh nghiệp lớn khác cũng đang áp dụng Go.

Có hơn một triệu lập trình viên dùng Go trên toàn thế giới, trải rộng ngân hàng và thương mại, game và truyền thông, công nghệ và các ngành khác, tại các doanh nghiệp đa dạng như [PayPal](/solutions/paypal), [Mercado Libre](/solutions/mercadolibre), Capital One, Dropbox, IBM, Monzo, New York Times, Salesforce, Square, Target, Twitch, Uber và tất nhiên Google.

Để tìm hiểu thêm về cách Go có thể giúp doanh nghiệp của bạn xây dựng phần mềm đáng tin cậy và có khả năng mở rộng như tại American Express, hãy ghé thăm [go.dev](/) ngay hôm nay.
