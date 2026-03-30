---
title: "Nhóm Firebase Hosting mở rộng quy mô với Go như thế nào"
company: Firebase
logoSrc: firebase.svg
logoSrcDark: firebase.svg
heroImgSrc: go_firebase_case_study.png
series: Case Studies
template: true
quote: |
  Firebase là nền tảng di động của Google giúp bạn nhanh chóng phát triển các ứng dụng
  chất lượng cao và phát triển doanh nghiệp.

  Nhóm Firebase Hosting đã chia sẻ hành trình sử dụng Go, bao gồm việc di chuyển
  backend từ Node.js, sự dễ dàng trong việc giới thiệu các nhà phát triển Go mới, và
  cách Go đã giúp họ mở rộng quy mô.
---

Nhóm Firebase Hosting cung cấp dịch vụ lưu trữ web tĩnh cho khách hàng Google Cloud.
Họ cung cấp một máy chủ web tĩnh đứng sau mạng phân phối nội dung toàn cầu,
và cung cấp cho người dùng các công cụ dễ sử dụng. Nhóm cũng phát triển các tính năng
từ tải lên tệp trang web đến đăng ký tên miền đến theo dõi mức sử dụng.

Trước khi gia nhập Google, stack công nghệ của Firebase Hosting được viết bằng Node.js.
Nhóm bắt đầu sử dụng Go khi họ cần tương tác với nhiều dịch vụ Google khác.
Họ quyết định dùng Go để giúp mở rộng quy mô dễ dàng và hiệu quả, biết rằng "tính
đồng thời sẽ tiếp tục là nhu cầu lớn." Họ "tự tin rằng Go sẽ có hiệu suất tốt hơn"
và "thích việc Go súc tích hơn" các ngôn ngữ khác mà họ đang cân nhắc, theo lời
Michael Bleigh, một kỹ sư phần mềm trong nhóm.

Bắt đầu với một dịch vụ nhỏ viết bằng Go, nhóm đã di chuyển toàn bộ backend qua
một loạt các bước. Nhóm liên tục xác định các tính năng lớn muốn triển khai và,
trong quá trình đó, viết lại chúng bằng Go và chuyển sang Google Cloud cùng hệ thống
quản lý cụm nội bộ của Google. **Hiện tại nhóm Firebase Hosting đã thay thế 100%
code backend Node.js bằng Go.**

Kinh nghiệm viết Go của nhóm bắt đầu với một kỹ sư. "Thông qua việc học hỏi
ngang hàng và Go nhìn chung dễ bắt đầu, tất cả mọi người trong nhóm giờ đều có
kinh nghiệm phát triển Go," Bleigh cho biết. Họ nhận thấy rằng trong khi phần lớn
người mới vào nhóm chưa có kinh nghiệm với Go, "hầu hết họ có thể làm việc hiệu quả
sau vài tuần."

"Dùng Go, dễ dàng thấy code được tổ chức như thế nào và code làm gì," Bleigh nói
thay mặt nhóm. "Go nhìn chung rất dễ đọc và dễ hiểu. Cách xử lý lỗi, receiver
và interface của ngôn ngữ đều dễ hiểu nhờ các idiom trong ngôn ngữ."

Tính đồng thời tiếp tục là trọng tâm của nhóm khi họ mở rộng quy mô. Robert Rossney,
một kỹ sư phần mềm, chia sẻ rằng "Go giúp rất dễ dàng đặt tất cả những thứ đồng thời
khó khăn vào một chỗ, và ở những nơi khác nó được trừu tượng hóa." Rossney cũng đề cập
đến lợi ích của việc sử dụng ngôn ngữ được xây dựng với tính đồng thời trong đầu,
nói rằng "cũng có nhiều cách để làm đồng thời trong Go. Chúng tôi đã phải học khi
nào mỗi cách là tốt nhất, cách xác định khi nào một bài toán là bài toán đồng thời,
cách debug, nhưng điều đó xuất phát từ việc bạn thực sự có thể viết các pattern này
trong code Go."

{{backgroundquote `
  author: Robert Rossney
  title: Software Engineer
  quote: |
    Nhìn chung, không có lúc nào trong nhóm mà chúng tôi cảm thấy bực bội với Go,
    nó chỉ giải phóng bạn và để bạn làm việc.
`}}

Hàng trăm nghìn khách hàng lưu trữ trang web của họ với Firebase Hosting,
nghĩa là code Go được dùng để phục vụ hàng tỷ request mỗi ngày. "Lượng khách hàng
và lưu lượng truy cập của chúng tôi đã tăng gấp đôi nhiều lần kể từ khi di chuyển
sang Go mà không bao giờ cần các tối ưu hóa tinh chỉnh," Bleigh chia sẻ. Với Go,
nhóm đã thấy cải thiện hiệu suất cả trong phần mềm và trong nhóm, với mức tăng
năng suất xuất sắc. "Nhìn chung," Rossney đề cập, "...không có lúc nào trong nhóm
mà chúng tôi cảm thấy bực bội với Go, nó chỉ giải phóng bạn và để bạn làm việc."

Ngoài nhóm Firebase Hosting, các nhóm kỹ thuật trên toàn Google đã áp dụng Go
trong quy trình phát triển của họ. Đọc về cách nhóm [Core Data
Solutions](/solutions/google/coredata/) và [Chrome](/solutions/google/chrome/)
sử dụng Go để xây dựng phần mềm nhanh, đáng tin cậy và hiệu quả ở quy mô lớn.
