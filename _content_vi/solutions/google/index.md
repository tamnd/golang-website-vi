---
title: 'Sử dụng Go tại Google'
date: 2020-08-27
company: Google
logoSrc: google.svg
logoSrcDark: google.svg
heroImgSrc: go_core_data_case_study.png
carouselImgSrc: go_google_case_study_carousel.png
series: Case Studies
type: solutions
template: true
description: |-
  Google là công ty công nghệ có sứ mệnh tổ chức thông tin của thế giới
  và làm cho nó có thể truy cập và hữu ích cho tất cả mọi người.

  Go được tạo ra tại Google vào năm 2007 để cải thiện năng suất lập trình trong
  thời đại máy tính đa nhân có kết nối mạng và codebase lớn. Ngày nay, hơn 10
  năm kể từ khi công bố công khai vào năm 2009, việc sử dụng Go trong Google đã
  phát triển rất nhiều.
quote: Go được tạo ra tại Google năm 2007, và từ đó, các nhóm kỹ thuật trên toàn Google đã áp dụng Go để xây dựng sản phẩm và dịch vụ ở quy mô khổng lồ.

---

{{pullquote `
  author: Rob Pike
  quote: |
    Go started in September 2007 when Robert Griesemer, Ken Thompson, and I began
    discussing a new language to address the engineering challenges we and our
    colleagues at Google were facing in our daily work.

    When we first released Go to the public in November 2009, we didn't know if the
    language would be widely adopted or if it might influence future languages.
    Looking back from 2020, Go has succeeded in both ways: it is widely used both
    inside and outside Google, and its approaches to network concurrency and
    software engineering have had a noticeable effect on other languages and their
    tools.

    Go has turned out to have a much broader reach than we had ever expected. Its
    growth in the industry has been phenomenal, and it has powered many projects at
    Google.
`}}

Các câu chuyện dưới đây là một mẫu nhỏ trong nhiều cách Go được sử dụng tại Google.

### Nhóm Core Data Solutions của Google dùng Go như thế nào

Sứ mệnh của Google là "tổ chức thông tin của thế giới và làm cho nó có thể truy cập và hữu ích cho tất cả mọi người." Một trong những nhóm chịu trách nhiệm tổ chức thông tin đó là nhóm Core Data Solutions của Google. Nhóm này, trong số những thứ khác, duy trì các dịch vụ lập chỉ mục trang web trên toàn cầu. Các dịch vụ lập chỉ mục web này giúp hỗ trợ các sản phẩm như Google Search bằng cách cập nhật và toàn diện các kết quả tìm kiếm, và chúng được viết bằng Go.

[Tìm hiểu thêm](/solutions/google/coredata/)

---

### Dịch vụ tối ưu hóa nội dung Chrome chạy trên Go

Khi nghĩ đến sản phẩm Chrome, bạn có lẽ chỉ nghĩ đến trình duyệt được cài đặt bởi người dùng. Nhưng đằng sau, Chrome có một hạm đội backend phong phú. Trong số đó có dịch vụ Chrome Optimization Guide. Dịch vụ này là nền tảng quan trọng cho chiến lược trải nghiệm người dùng của Chrome, hoạt động trong đường dẫn quan trọng cho người dùng, và được triển khai bằng Go.

[Tìm hiểu thêm](/solutions/google/chrome/)

---

### Nhóm Firebase Hosting mở rộng quy mô như thế nào với Go

Nhóm Firebase Hosting cung cấp dịch vụ lưu trữ web tĩnh cho khách hàng Google Cloud. Họ cung cấp web host tĩnh nằm sau mạng phân phối nội dung toàn cầu và cung cấp cho người dùng các công cụ dễ sử dụng. Nhóm cũng phát triển các tính năng từ tải lên file site đến đăng ký tên miền đến theo dõi mức sử dụng.

[Tìm hiểu thêm](/solutions/google/firebase/)

---

### Vận hành production Google: Nhóm SRE của Google dùng Go như thế nào

Google vận hành một số lượng nhỏ các dịch vụ rất lớn. Các dịch vụ đó được cung cấp bởi cơ sở hạ tầng toàn cầu bao gồm mọi thứ cần thiết: hệ thống lưu trữ, load balancer, mạng, logging, giám sát và nhiều hơn nữa. Tuy nhiên, đây không phải là hệ thống tĩnh và không thể là. Kiến trúc phát triển, các sản phẩm và ý tưởng mới được tạo ra, các phiên bản mới phải được triển khai, cấu hình được push, schema cơ sở dữ liệu được cập nhật và nhiều hơn nữa. Chúng tôi triển khai các thay đổi cho hệ thống hàng chục lần mỗi giây.

[Tìm hiểu thêm](/solutions/google/sitereliability/)
