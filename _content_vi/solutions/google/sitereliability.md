---
title: "Vận hành Google Production: Nhóm Site Reliability Engineering của Google dùng Go như thế nào"
company: Google Site Reliability Engineering (SRE)
logoSrc: sitereliability.svg
logoSrcDark: sitereliability.svg
heroImgSrc: go_sitereliability_case_study.png
series: Case Studies
quote: |
  Nhóm Site Reliability Engineering của Google có sứ mệnh bảo vệ, cung cấp và cải tiến
  phần mềm và hệ thống đằng sau tất cả các dịch vụ công khai của Google, bao gồm
  Google Search, Ads, Gmail, Android, YouTube và App Engine, cùng nhiều dịch vụ khác,
  với ánh mắt luôn theo dõi tính sẵn sàng, độ trễ, hiệu suất và dung lượng của chúng.

  Họ đã chia sẻ kinh nghiệm xây dựng các hệ thống quản lý production cốt lõi với Go,
  xuất phát từ kinh nghiệm với Python và C++.
authors:
  - Pierre Palatin, Site Reliability Engineer
template: true
---

Google vận hành một số lượng nhỏ các dịch vụ rất lớn. Những dịch vụ đó được hỗ trợ
bởi cơ sở hạ tầng toàn cầu bao gồm mọi thứ mà nhà phát triển cần: hệ thống lưu trữ,
bộ cân bằng tải, mạng, ghi log, giám sát và nhiều thứ khác. Tuy nhiên, đó không phải
là hệ thống tĩnh, và không thể là. Kiến trúc phát triển, các sản phẩm và ý tưởng mới
được tạo ra, các phiên bản mới phải được triển khai, config được đẩy, schema cơ sở dữ liệu
được cập nhật, và nhiều hơn nữa. Chúng tôi triển khai các thay đổi cho hệ thống hàng
chục lần mỗi giây.

Do quy mô này và nhu cầu quan trọng về độ tin cậy, Google đã tiên phong trong Site
Reliability Engineering (SRE), một vai trò mà nhiều công ty khác đã áp dụng từ đó.
"SRE là điều bạn có được khi coi vận hành như thể đó là bài toán phần mềm.
Sứ mệnh của chúng tôi là bảo vệ, cung cấp và cải tiến phần mềm và hệ thống đằng sau
tất cả các dịch vụ công khai của Google với ánh mắt luôn theo dõi tính sẵn sàng,
độ trễ, hiệu suất và dung lượng của chúng."
— [Site Reliability Engineering (SRE)](https://sre.google/).

{{backgroundquote `
  quote: |
    Go hứa hẹn một điểm ngọt giữa hiệu suất và khả năng đọc mà không ngôn ngữ nào
    trong số các ngôn ngữ kia [Python và C++] có thể cung cấp.
`}}

Vào năm 2013-2014, nhóm SRE của Google nhận ra rằng cách tiếp cận của chúng tôi
với quản lý production không còn phù hợp nữa theo nhiều cách. Chúng tôi đã tiến xa
hơn shell script, nhưng quy mô của chúng tôi có quá nhiều thành phần chuyển động
và phức tạp đến mức cần một cách tiếp cận mới. Chúng tôi xác định rằng cần chuyển
sang mô hình khai báo của production, gọi là "Prodspec", điều khiển một control plane
chuyên dụng, gọi là "Annealing".

Khi chúng tôi bắt đầu những dự án đó, Go vừa trở thành lựa chọn khả thi cho
các dịch vụ quan trọng tại Google. Hầu hết kỹ sư quen thuộc hơn với Python và C++,
cả hai đều là lựa chọn hợp lệ. Tuy nhiên, Go thu hút sự quan tâm của chúng tôi.
Sức hấp dẫn của tính mới lạ chắc chắn là một yếu tố. Nhưng quan trọng hơn, Go
hứa hẹn một điểm ngọt giữa hiệu suất và khả năng đọc mà không ngôn ngữ nào trong số
các ngôn ngữ kia có thể cung cấp. Chúng tôi bắt đầu một thử nghiệm nhỏ với Go cho
một số phần ban đầu của Annealing và Prodspec. Khi các dự án tiến triển, những phần
ban đầu đó viết bằng Go thấy mình ở trung tâm. Chúng tôi hài lòng với Go, sự đơn
giản của nó ngày càng phù hợp với chúng tôi, hiệu suất ở đó, và các primitive đồng
thời sẽ khó thay thế.

{{backgroundquote `
  quote: |
    Hiện tại phần lớn production của Google được quản lý và duy trì bởi các hệ thống
    của chúng tôi viết bằng Go.
`}}

Không bao giờ có yêu cầu hay bắt buộc phải sử dụng Go, nhưng chúng tôi không có
mong muốn quay lại Python hay C++. Go phát triển tự nhiên trong Annealing và Prodspec.
Đó là lựa chọn đúng đắn, và do đó hiện là ngôn ngữ được ưa chuộng của chúng tôi.
Hiện tại phần lớn production của Google được quản lý và duy trì bởi các hệ thống
của chúng tôi viết bằng Go.

Sức mạnh của việc có một ngôn ngữ đơn giản trong những dự án đó rất khó đánh giá hết.
Đã có những trường hợp một tính năng nào đó thực sự thiếu, chẳng hạn như khả năng
thực thi trong code rằng một cấu trúc phức tạp không được thay đổi. Nhưng với mỗi
trường hợp như vậy, chắc chắn đã có hàng chục hay hàng trăm trường hợp mà sự đơn
giản đã giúp ích.

{{backgroundquote `
  quote: |
    Sự đơn giản của Go có nghĩa là code dễ theo dõi, dù là để phát hiện lỗi trong
    quá trình review hay khi cố gắng xác định chính xác điều gì đã xảy ra trong
    một sự cố dịch vụ.
`}}

Ví dụ, Annealing ảnh hưởng đến nhiều nhóm và dịch vụ khác nhau, nghĩa là chúng tôi
phụ thuộc nhiều vào sự đóng góp từ khắp công ty. Sự đơn giản của Go giúp những người
bên ngoài nhóm của chúng tôi có thể thấy lý do tại sao một phần nào đó không hoạt
động với họ, và thường tự cung cấp bản sửa lỗi hoặc tính năng. Điều này cho phép
chúng tôi phát triển nhanh chóng.

Prodspec và Annealing chịu trách nhiệm về một số thành phần khá quan trọng. Sự đơn
giản của Go có nghĩa là code dễ theo dõi, dù là để phát hiện lỗi trong quá trình
review hay khi cố gắng xác định chính xác điều gì đã xảy ra trong một sự cố dịch vụ.

Hiệu suất của Go và hỗ trợ đồng thời cũng rất quan trọng cho công việc của chúng tôi.
Vì mô hình production của chúng tôi là khai báo, chúng tôi có xu hướng thao tác
nhiều dữ liệu có cấu trúc, mô tả production đang là gì và nên là gì. Chúng tôi có
các dịch vụ lớn nên dữ liệu có thể tăng lớn, thường làm cho xử lý tuần tự thuần túy
không đủ hiệu quả.

Chúng tôi đang thao tác dữ liệu này theo nhiều cách và nhiều nơi. Không phải là vấn đề
có một người thông minh nghĩ ra phiên bản song song của thuật toán của chúng tôi.
Đó là vấn đề của sự song song thông thường, tìm điểm nghẽn tiếp theo và song song
hóa đoạn code đó. Và Go cho phép chính xác điều đó.

Kết quả từ thành công với Go, chúng tôi hiện dùng Go cho mọi phát triển mới
cho Prodspec và Annealing.

Ngoài nhóm Site Reliability Engineering, các nhóm kỹ thuật trên toàn Google đã áp dụng
Go trong quy trình phát triển của họ. Đọc về cách nhóm
[Core Data Solutions](/solutions/google/coredata/),
[Firebase Hosting](/solutions/google/firebase/) và
[Chrome](/solutions/google/chrome/) sử dụng Go để xây dựng phần mềm nhanh, đáng
tin cậy và hiệu quả ở quy mô lớn.
