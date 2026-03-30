---
title: "Dịch vụ tối ưu hóa nội dung Chrome chạy trên Go"
company: Chrome
logoSrc: chrome.svg
logoSrcDark: chrome.svg
heroImgSrc: go_chrome_case_study.png
series: Case Studies
template: true
quote: |
  Google Chrome là trình duyệt web đơn giản hơn, bảo mật hơn và nhanh hơn bao giờ hết,
  với trí tuệ của Google được tích hợp sẵn.

  Trong nghiên cứu điển hình này, nhóm Chrome Optimization Guide
  chia sẻ cách họ thử nghiệm với Go, nhanh chóng làm quen và kế hoạch
  sử dụng Go trong tương lai.
---

Khi nghĩ đến sản phẩm Chrome, bạn có lẽ chỉ nghĩ đến trình duyệt được cài đặt bởi người dùng. Nhưng đằng sau, Chrome có một hạm đội backend phong phú. Trong số đó có dịch vụ Chrome Optimization Guide. Dịch vụ này là nền tảng quan trọng cho chiến lược trải nghiệm người dùng của Chrome, hoạt động trong đường dẫn quan trọng cho người dùng, và được triển khai bằng Go.

Dịch vụ Chrome Optimization Guide được thiết kế để mang sức mạnh của Google đến Chrome bằng cách cung cấp gợi ý cho trình duyệt đã cài đặt về các tối ưu hóa có thể thực hiện khi tải trang, cũng như khi nào chúng có thể được áp dụng hiệu quả nhất. Nó bao gồm sự kết hợp giữa server thời gian thực và phân tích log hàng loạt.

Tất cả người dùng Lite mode của Chrome nhận dữ liệu qua dịch vụ thông qua các cơ chế sau: push blob dữ liệu cung cấp gợi ý cho các site nổi tiếng trong khu vực địa lý của họ, check-in với server Google để lấy gợi ý cho các host mà người dùng cụ thể hay truy cập, và theo yêu cầu cho việc tải trang mà gợi ý chưa có trên thiết bị. Nếu dịch vụ Chrome Optimization Guide đột nhiên biến mất, người dùng có thể nhận thấy sự thay đổi đáng kể về tốc độ tải trang và lượng dữ liệu tiêu thụ khi duyệt web.

{{backgroundquote `
  author: Sophie Chang
  title: Software Engineer
  quote: |
    Given that Go was a success for us, we plan to continue to use
    it where appropriate
`}}

Khi nhóm kỹ thuật Chrome bắt đầu xây dựng dịch vụ, chỉ một vài thành viên quen với Go. Hầu hết nhóm quen thuộc với C++ hơn, nhưng họ thấy boilerplate phức tạp cần thiết để dựng server C++ quá nhiều. Nhóm chia sẻ rằng "[họ] khá có động lực học Go vì sự đơn giản, tốc độ làm quen nhanh và hệ sinh thái" và "[tinh thần phiêu lưu của họ] đã được đền đáp." Hàng triệu người dùng dựa vào dịch vụ này để có trải nghiệm Chrome tốt hơn, và việc chọn Go không phải là quyết định nhỏ. Sau kinh nghiệm cho đến nay, nhóm cũng chia sẻ rằng "vì Go là thành công với chúng tôi, chúng tôi có kế hoạch tiếp tục sử dụng nó khi thích hợp."

Ngoài nhóm Chrome Optimization Guide, các nhóm kỹ thuật trên toàn Google đã áp dụng Go trong quy trình phát triển. Đọc về cách các nhóm [Core Data Solutions](/solutions/google/coredata/) và [Firebase Hosting](/solutions/google/firebase/) sử dụng Go để xây dựng phần mềm nhanh, đáng tin cậy và hiệu quả ở quy mô lớn.

*Ghi chú biên tập: Nhóm Go xin cảm ơn Sophie Chang vì những đóng góp cho bài viết này.*
