---
title: "Nhóm Core Data Solutions của Google dùng Go như thế nào"
company: Core Data
logoSrc: google.svg
logoSrcDark: google.svg
heroImgSrc: go_core_data_case_study.png
series: Case Studies
template: true
quote: |
  Google là công ty công nghệ có sứ mệnh tổ chức thông tin của thế giới
  và làm cho nó có thể truy cập và hữu ích cho tất cả mọi người.

  Trong nghiên cứu điển hình này, nhóm Core Data Solutions của Google chia sẻ
  hành trình với Go, bao gồm quyết định viết lại dịch vụ lập chỉ mục web bằng Go,
  tận dụng tính đồng thời tích hợp của Go, và quan sát cách Go giúp cải thiện
  quy trình phát triển.
authors:
  - Prasanna Meda, Software Engineer, Core Data Solutions
---

Sứ mệnh của Google là "tổ chức thông tin của thế giới và làm cho nó có thể truy cập và hữu ích cho tất cả mọi người." Một trong những nhóm chịu trách nhiệm tổ chức thông tin đó là nhóm Core Data Solutions của Google. Nhóm này, trong số những thứ khác, duy trì các dịch vụ lập chỉ mục trang web trên toàn cầu. Các dịch vụ lập chỉ mục web này giúp hỗ trợ các sản phẩm như Google Search bằng cách cập nhật và toàn diện các kết quả tìm kiếm, và chúng được viết bằng Go.

Năm 2015, để theo kịp quy mô của Google, nhóm chúng tôi cần viết lại stack lập chỉ mục từ một binary nguyên khối duy nhất viết bằng C++ thành nhiều thành phần trong kiến trúc microservice. Chúng tôi quyết định viết lại nhiều dịch vụ lập chỉ mục bằng Go, mà hiện nay chúng tôi sử dụng để vận hành phần lớn kiến trúc.

{{backgroundquote `
  author: Minjae Hwang
  title: Software Engineer
  quote: |
    Go's built-in concurrency is a natural fit because engineers on the team are
    encouraged to use concurrency and parallel algorithms.
`}}

Khi chọn ngôn ngữ, nhóm chúng tôi thấy một số tính năng của Go làm nó đặc biệt phù hợp. Ví dụ, tính đồng thời tích hợp của Go rất phù hợp vì kỹ sư trong nhóm được khuyến khích sử dụng thuật toán đồng thời và song song. Kỹ sư cũng thấy rằng "code Go tự nhiên hơn," cho phép họ dành thời gian tập trung vào logic nghiệp vụ và phân tích thay vì quản lý bộ nhớ và tối ưu hiệu năng.

Viết code đơn giản hơn nhiều khi viết bằng Go, vì nó giúp giảm gánh nặng nhận thức trong quá trình phát triển. Ví dụ, khi làm việc với C++, IDE tinh vi có thể "hiển thị source code không có lỗi biên dịch khi thực ra có" trong khi "trong Go, [code] sẽ luôn biên dịch khi [IDE] nói code không có lỗi biên dịch," MinJae Hwang, kỹ sư phần mềm trong nhóm Core Data Solutions, cho biết. Giảm các điểm ma sát nhỏ trong quá trình phát triển, như rút ngắn chu kỳ sửa lỗi biên dịch, giúp nhóm chúng tôi phát hành nhanh hơn trong quá trình viết lại ban đầu, và đã giúp duy trì chi phí bảo trì thấp.

"Khi trong C++ và tôi muốn sử dụng thêm package, tôi phải viết các phần như header. Khi viết bằng Go, **công cụ tích hợp cho phép tôi sử dụng package dễ dàng hơn. Tốc độ phát triển của tôi nhanh hơn nhiều,**" Hwang cũng chia sẻ.

Với cú pháp ngôn ngữ đơn giản và hỗ trợ của các công cụ Go, một số thành viên trong nhóm thấy dễ viết code Go hơn nhiều. Chúng tôi cũng thấy Go làm rất tốt việc kiểm tra kiểu tĩnh và một số nguyên tắc cơ bản của Go, như lệnh godoc, đã giúp nhóm xây dựng văn hóa kỷ luật hơn xung quanh việc viết tài liệu.

{{backgroundquote `
  author: Prasanna Meda
  title: Software Engineer
  quote: |
    ...Google's web indexing was re-architected within a year. More impressively,
    most developers on the team were rewriting in Go while also learning it.
`}}

Làm việc trên sản phẩm được sử dụng nhiều như vậy trên toàn thế giới không phải là nhiệm vụ nhỏ và quyết định dùng Go của nhóm chúng tôi không đơn giản, nhưng làm vậy đã giúp chúng tôi di chuyển nhanh hơn. Kết quả là, lập chỉ mục web của Google được tái kiến trúc trong vòng một năm. Ấn tượng hơn, hầu hết lập trình viên trong nhóm đang viết lại bằng Go trong khi cũng đang học nó.

Ngoài nhóm Core Data Solutions, các nhóm kỹ thuật trên toàn Google đã áp dụng Go trong quy trình phát triển. Đọc về cách các nhóm [Chrome](/solutions/google/chrome/) và [Firebase Hosting](/solutions/google/firebase/) sử dụng Go để xây dựng phần mềm nhanh, đáng tin cậy và hiệu quả ở quy mô lớn.
