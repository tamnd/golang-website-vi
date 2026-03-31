---
title: "MercadoLibre lớn mạnh cùng Go"
company: MercadoLibre
logoSrc: mercadolibre_light.svg
logoSrcDark: mercadolibre_dark.svg
heroImgSrc: go_mercadolibre_case_study_logo.png
carouselImgSrc: go_mercadolibre_case_study.png
date: 2019-11-10T16:26:31-04:00
series: Case Studies
quote: Go cung cấp code sạch, hiệu quả, dễ mở rộng cùng với sự tăng trưởng thương mại trực tuyến của MercadoLibre, và tăng năng suất lập trình viên bằng cách cho phép kỹ sư phục vụ lượng khán giả ngày càng tăng với ít code hơn.
template: true
---

{{pullquote `
  author: Eric Kohan
  title: Software Engineering Manager
  company: MercadoLibre
  quote: |
    I think that **the tour of Go is by far the best introduction to a language that I've seen**, It's really simple and it gives you a fair overview of probably 80 percent of the language. When we want to get developers to learn Go, and to get to production fast, we tell them to start with the tour of Go.
`}}

## Go giúp hệ sinh thái tích hợp thu hút lập trình viên và mở rộng thương mại điện tử

MercadoLibre, Inc. là hệ sinh thái thương mại trực tuyến lớn nhất Mỹ Latinh và hiện diện tại 18 quốc gia. Được thành lập năm 1999 và đặt trụ sở tại Argentina, công ty đã chuyển sang Go để giúp mở rộng và hiện đại hóa hệ sinh thái. Go cung cấp code sạch, hiệu quả, dễ mở rộng cùng với sự tăng trưởng thương mại trực tuyến của MercadoLibre, và tăng năng suất lập trình viên bằng cách cho phép kỹ sư phục vụ lượng khán giả ngày càng tăng với ít code hơn.

### MercadoLibre dùng Go để mở rộng quy mô

Vào năm 2015, trong nội bộ MercadoLibre có cảm giác ngày càng rõ rằng framework API hiện có của họ, trên Groovy và Grails, đang đạt giới hạn và công ty cần nền tảng khác để tiếp tục mở rộng. Nền tảng của MercadoLibre đang (và tiếp tục) mở rộng theo cấp số nhân, tạo ra rất nhiều công việc thêm cho lập trình viên: Cả Groovy và Grails đều đòi hỏi nhiều quyết định từ lập trình viên và Groovy là ngôn ngữ lập trình động. Đây không phải là sự kết hợp tốt cho tăng trưởng mở rộng nhanh chóng, vì MercadoLibre cần lập trình viên rất giàu kinh nghiệm trong môi trường tốn nhiều tài nguyên này để phát triển và điều chỉnh để đạt hiệu năng mong muốn. Thời gian thực thi test chậm, thời gian build và triển khai chậm. Do đó, nhu cầu hiệu quả code và khả năng mở rộng trở nên quan trọng như nhu cầu tốc độ phát triển code.

### Go cải thiện hiệu quả hệ thống

Như ví dụ về đóng góp của Go cho hiệu quả mạng, nhóm API cốt lõi xây dựng và duy trì các API lớn nhất tại trung tâm giải pháp microservice của công ty. Nhóm này tạo ra các user API, được sử dụng bởi MercadoLibre Marketplace, nền tảng FinTech MercadoPago, giải pháp vận chuyển và logistics của MercadoLibre, và các giải pháp khác. Với mức dịch vụ cao được yêu cầu bởi các giải pháp này, user API trung bình có từ tám đến mười triệu request mỗi phút, nhóm sử dụng Go để phục vụ chúng ở dưới mười mili giây mỗi request.

Nhóm API cũng triển khai container Docker, một sản phẩm SaaS cũng viết bằng Go, để ảo hóa môi trường phát triển và dễ dàng triển khai microservice qua Docker Engine. Hệ thống này hỗ trợ các API quan trọng, lớn hơn xử lý **hơn 20 triệu request mỗi phút bằng Go.**

Một API đã sử dụng các nguyên thủy đồng thời của Go để ghép ID hiệu quả từ nhiều dịch vụ. Nhóm đã thực hiện điều này chỉ với vài dòng code Go, và thành công của API này đã thuyết phục nhóm API cốt lõi di chuyển ngày càng nhiều microservice sang Go. Kết quả cuối cùng cho MercadoLibre là cải thiện hiệu quả chi phí và thời gian phản hồi hệ thống.

### Go cho khả năng mở rộng

Về mặt lịch sử, phần lớn stack của công ty dựa trên Grails và Groovy được hỗ trợ bởi cơ sở dữ liệu quan hệ. Tuy nhiên framework lớn với nhiều lớp này sớm gặp phải vấn đề về khả năng mở rộng.

Chuyển đổi kiến trúc kế thừa đó sang Go như framework mới rất mỏng để xây dựng API đã đơn giản hóa các lớp trung gian đó và mang lại lợi ích hiệu năng lớn. Ví dụ, một dịch vụ Go lớn hiện có thể **chạy 70.000 request mỗi máy chỉ với 20 MB RAM.**

{{backgroundquote `
  author: Eric Kohan
  title: Software Engineering Manager
  company: MercadoLibre
  quote: |
    Go was just marvelous for us. It's very powerful
    and very easy to learn, and with backend infrastructure, has been great for us in terms of scalability.
`}}

Dùng **Go cho phép MercadoLibre giảm số server** họ sử dụng cho dịch vụ này xuống còn một phần tám so với ban đầu (từ 32 server xuống còn bốn), cộng với mỗi server có thể hoạt động với ít năng lượng hơn (ban đầu bốn nhân CPU, giờ xuống còn hai nhân CPU). Với Go, công ty **loại bỏ 88% server và giảm CPU trên các server còn lại xuống một nửa**, tạo ra tiết kiệm chi phí khổng lồ.

Nằm giữa lập trình viên và nhà cung cấp đám mây, MercadoLibre sử dụng nền tảng gọi là Fury, công cụ platform-as-a-service để build, triển khai, giám sát và quản lý dịch vụ theo cách cloud-agnostic. Kết quả là, bất kỳ nhóm nào muốn tạo dịch vụ mới trong Go đều có quyền truy cập vào các template đã được chứng minh cho nhiều loại dịch vụ, và có thể nhanh chóng tạo kho lưu trữ trong GitHub với code khởi đầu, Docker image cho dịch vụ và pipeline triển khai. Kết quả cuối cùng là hệ thống cho phép kỹ sư tập trung vào xây dựng dịch vụ sáng tạo trong khi tránh được các giai đoạn tẻ nhạt khi thiết lập dự án mới, đồng thời chuẩn hóa hiệu quả các pipeline build và triển khai.

Ngày nay, **khoảng một nửa lưu lượng truy cập của Mercadolibre được xử lý bởi ứng dụng Go.**

### MercadoLibre dùng Go cho lập trình viên

Lingua franca lập trình cho cơ sở hạ tầng của MercadoLibre hiện là Go và Java. Mọi ứng dụng, mọi chương trình, mọi microservice đều được lưu trữ trong kho lưu trữ GitHub riêng, và công ty sử dụng thêm kho lưu trữ GitHub chứa bộ công cụ để giải quyết vấn đề mới và cho phép client tương tác với dịch vụ.

Các bộ công cụ Go và Java phong phú và được quản lý tốt này cho phép lập trình viên phát triển ứng dụng mới nhanh chóng với hỗ trợ tuyệt vời. Ngoài ra, trong cộng đồng hơn 2.800 lập trình viên, MercadoLibre có nhiều nhóm nội bộ có sẵn để chat và hướng dẫn về triển khai Go, dù ở các trung tâm phát triển khác nhau hay các quốc gia khác nhau. Công ty cũng thúc đẩy các nhóm làm việc nội bộ để cung cấp các buổi đào tạo cho lập trình viên Go mới của MercadoLibre, và tổ chức Go meetup cho các lập trình viên bên ngoài để giúp xây dựng cộng đồng lập trình viên Go Mỹ Latinh rộng hơn.

### Go như công cụ tuyển dụng

Sự ủng hộ Go của MercadoLibre cũng đã trở thành công cụ tuyển dụng mạnh mẽ cho công ty. MercadoLibre là một trong những công ty đầu tiên sử dụng Go ở Argentina, và có lẽ là lớn nhất ở Mỹ Latinh sử dụng ngôn ngữ này rộng rãi trong production. Có trụ sở tại Buenos Aires với nhiều startup và công ty công nghệ đang nổi lên gần đó, việc MercadoLibre áp dụng Go đã định hình thị trường cho lập trình viên trên vùng Pampas.

{{backgroundquote `
  author: Eric Kohan
  title: Software Engineering Manager
  company: MercadoLibre
  quote: |
    We really see eye-to-eye with the larger philosophy of the language. We love Go's simplicity, and we find that having its very explicit error handling has been a gain for developers because it results in safer, more stable code in production.
`}}

Buenos Aires ngày nay là thị trường rất cạnh tranh cho lập trình viên, cung cấp nhiều lựa chọn việc làm, và nhu cầu cao về công nghệ trong khu vực thúc đẩy mức lương cao, phúc lợi tốt và khả năng chọn lọc khi chọn nhà tuyển dụng. Do đó, MercadoLibre, như tất cả nhà tuyển dụng kỹ sư và lập trình viên trong khu vực, nỗ lực cung cấp môi trường làm việc thú vị và con đường sự nghiệp mạnh. Go đã chứng minh là yếu tố khác biệt quan trọng cho MercadoLibre: công ty tổ chức workshop Go cho lập trình viên bên ngoài để họ đến học Go, và khi họ thích những gì đang làm và những người họ nói chuyện, họ nhanh chóng nhận ra MercadoLibre là nơi làm việc hấp dẫn.

### Go trao quyền cho lập trình viên

MercadoLibre sử dụng Go vì sự đơn giản với hệ thống ở quy mô lớn, nhưng sự đơn giản đó cũng là lý do lập trình viên của công ty yêu thích Go.

Công ty cũng sử dụng các trang web như [Go by Example](https://gobyexample.com/) và [Effective Go](/doc/effective_go.html) để giáo dục lập trình viên mới, và chia sẻ các API nội bộ tiêu biểu viết bằng Go để tăng tốc độ hiểu biết và thành thạo. Lập trình viên MercadoLibre nhận được tài nguyên cần thiết để tiếp nhận ngôn ngữ, sau đó tận dụng kỹ năng và sự nhiệt tình của chính họ để bắt đầu lập trình.

{{backgroundquote `
  author: Federico Martin Roasio
  title: Technical Project Lead
  company: MercadoLibre
  quote: |
    Go has been great for writing business logic, and we are the team that writes those APIs.
`}}

MercadoLibre tận dụng cú pháp biểu cảm và sạch của Go để giúp lập trình viên dễ dàng viết chương trình chạy hiệu quả trên các nền tảng đám mây hiện đại. Và trong khi tốc độ phát triển mang lại hiệu quả chi phí cho công ty, lập trình viên cá nhân được hưởng lợi từ đường cong học tập nhanh mà Go mang lại. Không chỉ các kỹ sư giàu kinh nghiệm của MercadoLibre có thể xây dựng ứng dụng quan trọng rất nhanh với Go, mà ngay cả kỹ sư cấp đầu vào cũng đã viết được dịch vụ mà trong các ngôn ngữ khác, MercadoLibre chỉ tin tưởng với lập trình viên cấp cao hơn. Ví dụ, một tập hợp user API quan trọng xử lý gần mười triệu request mỗi phút được phát triển bởi kỹ sư phần mềm đầu vào, nhiều người chỉ biết về lập trình từ các khóa học gần đây tại đại học. Tương tự, MercadoLibre đã thấy các lập trình viên đã thành thạo các ngôn ngữ lập trình khác (như Java, .NET hay Ruby) học Go đủ nhanh để bắt đầu viết dịch vụ production chỉ trong vài tuần.

Với Go, **thời gian build của MercadoLibre nhanh gấp ba lần (3x)** và **bộ test chạy nhanh hơn đáng kinh ngạc 24 lần**. Điều này có nghĩa là lập trình viên của công ty có thể thực hiện thay đổi, sau đó build và test thay đổi đó nhanh hơn nhiều so với trước.

Và việc giảm thời gian chạy bộ test của MercadoLibre từ 90 giây xuống **chỉ còn 3 giây với Go** là một lợi ích khổng lồ cho lập trình viên, cho phép họ duy trì tập trung (và context) trong khi các test nhanh hơn nhiều hoàn thành.

Tận dụng thành công này, MercadoLibre cam kết không chỉ giáo dục liên tục cho lập trình viên mà còn giáo dục Go liên tục. Công ty gửi các lãnh đạo kỹ thuật chủ chốt đến GopherCon và các sự kiện Go khác mỗi năm, nhóm cơ sở hạ tầng và bảo mật của MercadoLibre khuyến khích tất cả nhóm phát triển cập nhật phiên bản Go, và công ty có nhóm phát triển _Go-meli-toolkit_: Thư viện Go hoàn chỉnh để giao tiếp với tất cả dịch vụ do Fury cung cấp.

### Bắt đầu với Go cho doanh nghiệp của bạn

Giống như MercadoLibre bắt đầu với dự án bằng chứng khái niệm để triển khai Go, hàng chục doanh nghiệp lớn khác cũng đang áp dụng Go.

Có hơn một triệu lập trình viên dùng Go trên toàn thế giới, trải rộng ngân hàng và thương mại, game và truyền thông, công nghệ và các ngành khác, tại các doanh nghiệp đa dạng như [American Express](/solutions/americanexpress), [PayPal](/solutions/paypal), Capital One, Dropbox, IBM, Monzo, New York Times, Salesforce, Square, Target, Twitch, Uber và tất nhiên Google.

Để tìm hiểu thêm về cách Go có thể giúp doanh nghiệp của bạn xây dựng phần mềm đáng tin cậy và có khả năng mở rộng như tại MercadoLibre, hãy ghé thăm [go.dev](/) ngay hôm nay.
