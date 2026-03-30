---
title: PayPal dùng Go để hiện đại hóa và mở rộng quy mô
date: 2020-06-01
company: PayPal
logoSrc: paypal.svg
logoSrcDark: paypal.svg
heroImgSrc: go_paypal_case_study_logo.png
carouselImgSrc: go_paypal_case_study.png
series: Case Studies
quote: Giá trị của Go trong việc tạo ra code sạch, hiệu quả, dễ dàng mở rộng cùng với triển khai phần mềm, làm cho ngôn ngữ này phù hợp để hỗ trợ mục tiêu của PayPal.
template: true
---

{{pullquote `
  author: Bala Natarajan
  title: <span class="NoWrapSpan">Sr. Director of Engineering,</span>&nbsp;<span class="NoWrapSpan">Developer Experience</span>
  company: PayPal
  quote: |
    Since our NoSQL and DB proxy used quite a bit of system details in a multi-threaded mode, the code got complex managing the different conditions, given that Go provides channels and routines to deal with complexity, we were able to structure the code to meet our requirements.
`}}

## Cơ sở hạ tầng code mới được xây dựng trên Go

PayPal được tạo ra để dân chủ hóa dịch vụ tài chính và trao quyền cho cá nhân và doanh nghiệp tham gia và phát triển trong nền kinh tế toàn cầu. Trung tâm của nỗ lực này là Nền tảng Thanh toán PayPal, sử dụng kết hợp công nghệ độc quyền và bên thứ ba để xử lý giao dịch hiệu quả và an toàn giữa hàng triệu merchant và người tiêu dùng trên toàn thế giới. Khi Nền tảng Thanh toán ngày càng lớn và phức tạp hơn, PayPal tìm cách hiện đại hóa hệ thống và giảm thời gian đưa ứng dụng mới ra thị trường.

Giá trị của Go trong việc tạo ra code sạch, hiệu quả, dễ mở rộng cùng với triển khai phần mềm làm cho ngôn ngữ này phù hợp để hỗ trợ mục tiêu của PayPal.

Trung tâm của Nền tảng Xử lý Thanh toán là cơ sở dữ liệu NoSQL độc quyền mà PayPal đã phát triển bằng C++. Tuy nhiên, sự phức tạp của code làm giảm đáng kể khả năng lập trình viên phát triển nền tảng. Bố cục code đơn giản, goroutine (luồng thực thi nhẹ) và channel (đường ống kết nối các goroutine đồng thời) của Go làm cho Go trở thành lựa chọn tự nhiên cho nhóm phát triển NoSQL để đơn giản hóa và hiện đại hóa nền tảng.

Như một bằng chứng khái niệm, nhóm phát triển dành sáu tháng học Go và tái triển khai hệ thống NoSQL từ đầu bằng Go, trong thời gian đó họ cũng cung cấp thông tin chi tiết về cách triển khai Go rộng rãi hơn tại PayPal. Đến nay, ba mươi phần trăm cluster đã được di chuyển để sử dụng cơ sở dữ liệu NoSQL mới.

## Dùng Go để đơn giản hóa cho quy mô

Khi nền tảng của PayPal ngày càng phức tạp hơn, Go cung cấp cách để dễ dàng đơn giản hóa sự phức tạp của việc tạo và chạy phần mềm ở quy mô lớn. Ngôn ngữ cung cấp cho PayPal thư viện tuyệt vời và công cụ nhanh, cộng với tính đồng thời, bộ gom rác và type safety.

Với Go, PayPal cho phép lập trình viên dành nhiều thời gian hơn để nhìn vào code và suy nghĩ chiến lược, bằng cách giải phóng họ khỏi sự ồn ào của phát triển C++ và Java.

Sau thành công của hệ thống NoSQL được viết lại này, nhiều nhóm nền tảng và nội dung trong PayPal bắt đầu áp dụng Go. Nhóm hiện tại của Natarajan chịu trách nhiệm về các pipeline build, test và release của PayPal, tất cả được xây dựng bằng Go. Công ty có trang trại build và test lớn được quản lý hoàn toàn bằng cơ sở hạ tầng Go để hỗ trợ build-as-a-service (và test-as-a-service) cho lập trình viên trên toàn công ty.

  <img
    loading="lazy"
    width="607"
    height="289"
    class=""
    alt="Go gopher factory"
    src="/images/gophers/factory.png">

## Hiện đại hóa hệ thống PayPal với Go

Với các khả năng điện toán phân tán mà PayPal yêu cầu, Go là ngôn ngữ phù hợp để làm mới hệ thống. PayPal cần lập trình có tính đồng thời và song song, biên dịch để hiệu năng cao và di động cao, và mang lại cho lập trình viên lợi ích của kiến trúc mã nguồn mở module, có thể kết hợp, Go đã cung cấp tất cả những điều đó và nhiều hơn để giúp PayPal hiện đại hóa hệ thống.

Bảo mật và khả năng hỗ trợ là vấn đề quan trọng tại PayPal, và các pipeline vận hành của công ty ngày càng bị thống trị bởi Go vì sự sạch sẽ và tính module của ngôn ngữ giúp họ đạt được những mục tiêu này. Việc triển khai Go của PayPal tạo ra nền tảng sáng tạo cho lập trình viên, cho phép họ tạo ra phần mềm đơn giản, hiệu quả và đáng tin cậy ở quy mô lớn cho các thị trường toàn cầu của PayPal.

Khi PayPal tiếp tục hiện đại hóa cơ sở hạ tầng mạng được định nghĩa bởi phần mềm (SDN) với Go, họ thấy lợi ích hiệu năng ngoài code dễ bảo trì hơn. Ví dụ, Go hiện hỗ trợ router, load balancer và ngày càng nhiều hệ thống production.

{{backgroundquote `
  author: Bala Natarajan
  title: Sr. Director of Engineering
  quote: |
    In our tightly managed environments where we run Go code, we have seen a CPU reduction of approximately ten percent with cleaner and maintainable code.
`}}

## Go tăng năng suất lập trình viên

Là hoạt động toàn cầu, PayPal cần các nhóm phát triển hiệu quả trong việc quản lý hai loại quy mô: quy mô production, đặc biệt là các hệ thống đồng thời tương tác với nhiều server khác (như dịch vụ đám mây); và quy mô phát triển, đặc biệt là codebase lớn được phát triển bởi nhiều lập trình viên phối hợp (như phát triển mã nguồn mở).

PayPal tận dụng Go để giải quyết các vấn đề quy mô này. Lập trình viên của công ty được hưởng lợi từ khả năng của Go kết hợp sự dễ dàng lập trình của ngôn ngữ được thông dịch, kiểu động với hiệu quả và an toàn của ngôn ngữ kiểu tĩnh, biên dịch. Khi PayPal hiện đại hóa hệ thống, hỗ trợ cho điện toán mạng và đa nhân là rất quan trọng. Go không chỉ cung cấp hỗ trợ đó mà cung cấp nhanh chóng, chỉ mất nhiều nhất vài giây để biên dịch một file thực thi lớn trên một máy tính.

Hiện có hơn 100 lập trình viên Go tại PayPal, và các lập trình viên tương lai chọn áp dụng Go sẽ dễ dàng hơn khi được phê duyệt ngôn ngữ nhờ nhiều triển khai thành công đã trong production tại công ty.

Quan trọng nhất, lập trình viên PayPal đã tăng năng suất với Go. Cơ chế đồng thời của Go đã giúp dễ dàng viết các chương trình tận dụng tối đa máy đa nhân và máy mạng của PayPal. Lập trình viên dùng Go cũng được hưởng lợi từ việc biên dịch nhanh sang mã máy và ứng dụng được hưởng tiện ích của bộ gom rác và sức mạnh của phản chiếu thời gian chạy.

## Tăng tốc thời gian ra thị trường của PayPal

Ngôn ngữ hạng nhất tại PayPal ngày nay là Java và Node, với Go chủ yếu được dùng làm ngôn ngữ cơ sở hạ tầng. Mặc dù Go có thể không bao giờ thay thế Node.js cho một số ứng dụng, Natarajan đang nỗ lực đưa Go lên hàng ngôn ngữ hạng nhất tại PayPal.

Qua nỗ lực của ông, PayPal cũng đang đánh giá việc chuyển sang Google Kubernetes Engine (GKE) để tăng tốc thời gian ra thị trường cho sản phẩm mới. GKE là môi trường sẵn sàng production được quản lý để triển khai ứng dụng containerized, mang đến các đổi mới mới nhất của Google về năng suất lập trình viên, vận hành tự động và linh hoạt mã nguồn mở.

Đối với PayPal, triển khai lên GKE sẽ cho phép phát triển và lặp nhanh bằng cách giúp PayPal dễ dàng hơn trong việc triển khai, cập nhật và quản lý ứng dụng và dịch vụ. Thêm vào đó PayPal sẽ thấy dễ dàng hơn khi chạy Machine Learning, GPU đa năng, High-Performance Computing và các khối lượng công việc khác được hưởng lợi từ bộ tăng tốc phần cứng chuyên dụng được GKE hỗ trợ.

Quan trọng nhất với PayPal, sự kết hợp phát triển Go và GKE cho phép công ty mở rộng quy mô dễ dàng để đáp ứng nhu cầu, vì autoscaling Kubernetes sẽ cho phép PayPal xử lý nhu cầu người dùng tăng đối với dịch vụ, giữ chúng khả dụng khi cần thiết nhất, rồi thu hẹp trong các giai đoạn yên tĩnh để tiết kiệm chi phí.

## Bắt đầu với Go cho doanh nghiệp của bạn

Câu chuyện của PayPal không phải duy nhất; hàng chục doanh nghiệp lớn khác đang khám phá cách Go giúp họ phát hành phần mềm đáng tin cậy nhanh hơn. Có hơn một triệu lập trình viên dùng Go trên toàn thế giới, trải rộng ngân hàng và thương mại, game và truyền thông, công nghệ và các ngành khác, tại các doanh nghiệp đa dạng như [American Express](/solutions/americanexpress), [Mercado Libre](/solutions/mercadolibre), Capital One, Dropbox, IBM, Monzo, New York Times, Salesforce, Square, Target, Twitch, Uber và tất nhiên Google.

Để tìm hiểu thêm về cách Go có thể giúp doanh nghiệp của bạn xây dựng phần mềm đáng tin cậy và có khả năng mở rộng như tại PayPal, hãy ghé thăm [go.dev](/) ngay hôm nay.
