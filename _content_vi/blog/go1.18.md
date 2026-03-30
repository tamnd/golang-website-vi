---
title: Go 1.18 đã được phát hành!
date: 2022-03-15
by:
- The Go Team
summary: Go 1.18 bổ sung generics, fuzzing native, workspace mode, cải tiến hiệu năng, và nhiều hơn nữa.
---

Hôm nay nhóm Go vô cùng hào hứng phát hành Go 1.18,
bạn có thể tải về bằng cách truy cập [trang download](/dl/).

Go 1.18 là một bản phát hành lớn bao gồm các tính năng mới,
cải tiến hiệu năng, và thay đổi lớn nhất từ trước đến nay đối với ngôn ngữ.
Không ngoa khi nói rằng quá trình thiết kế một số phần của Go 1.18
đã bắt đầu hơn một thập kỷ trước khi chúng tôi lần đầu phát hành Go.

## Generics

Trong Go 1.18, chúng tôi giới thiệu hỗ trợ mới cho
[code generic dùng kiểu được tham số hóa](/blog/why-generics).
Hỗ trợ generics là tính năng được yêu cầu nhiều nhất trong Go,
và chúng tôi tự hào mang đến hỗ trợ generics đáp ứng nhu cầu của phần lớn người dùng hiện nay.
Các bản phát hành tiếp theo sẽ cung cấp hỗ trợ bổ sung cho một số
trường hợp sử dụng generics phức tạp hơn.
Chúng tôi khuyến khích bạn làm quen với tính năng mới này qua
[hướng dẫn generics](/doc/tutorial/generics),
và khám phá cách tốt nhất để dùng generics tối ưu hóa và đơn giản hóa code của bạn ngay hôm nay.
[Ghi chú phát hành](/doc/go1.18) có thêm chi tiết về cách dùng generics trong Go 1.18.

## Fuzzing

Với Go 1.18, Go là ngôn ngữ lớn đầu tiên tích hợp fuzzing
hoàn toàn vào toolchain chuẩn.
Giống như generics, fuzzing đã được thiết kế trong một thời gian dài,
và chúng tôi vui mừng chia sẻ nó với hệ sinh thái Go trong bản phát hành này.
Hãy xem
[hướng dẫn fuzzing](/doc/tutorial/fuzz)
để bắt đầu với tính năng mới này.

## Workspace

Go modules đã được áp dụng gần như phổ biến,
và người dùng Go báo cáo điểm hài lòng rất cao trong các khảo sát hàng năm của chúng tôi.
Trong khảo sát người dùng 2021, thách thức phổ biến nhất
người dùng gặp phải với modules là làm việc trên nhiều module.
Trong Go 1.18, chúng tôi đã giải quyết vấn đề này với
[Go workspace mode](/doc/tutorial/workspaces) mới,
giúp làm việc với nhiều module trở nên đơn giản.

## Cải thiện hiệu năng 20%

Người dùng Apple M1, ARM64 và PowerPC64, hãy vui mừng!
Go 1.18 bao gồm cải tiến hiệu năng CPU lên đến 20%
nhờ mở rộng quy ước gọi hàm theo register ABI của Go 1.17 sang các kiến trúc này.
Chỉ để nhấn mạnh quy mô của bản phát hành này: cải tiến hiệu năng 20%
lại chỉ là điểm nổi bật đứng thứ tư!

Để mô tả chi tiết hơn về tất cả những gì có trong 1.18,
hãy xem [ghi chú phát hành](/doc/go1.18).

Go 1.18 là một cột mốc quan trọng cho toàn bộ cộng đồng Go.
Chúng tôi muốn cảm ơn mọi người dùng Go đã báo cáo lỗi, gửi thay đổi, viết hướng dẫn,
hoặc giúp đỡ theo bất kỳ cách nào để Go 1.18 trở thành hiện thực.
Chúng tôi không thể làm được điều đó nếu không có bạn.
Cảm ơn.

Tận hưởng Go 1.18!
