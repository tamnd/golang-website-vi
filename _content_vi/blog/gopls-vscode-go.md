---
title: Gopls được bật mặc định trong extension VS Code Go
date: 2021-02-01
by:
- Go tools team
tags:
- tools
- gopls
summary: Gopls, cung cấp các tính năng IDE cho Go cho nhiều trình soạn thảo, giờ được dùng mặc định trong VS Code Go.
template: true
---


Chúng tôi vui mừng thông báo rằng extension VS Code Go giờ bật [language server gopls](https://github.com/golang/tools/blob/master/gopls/README.md)
theo mặc định, để cung cấp các tính năng IDE mạnh mẽ hơn và hỗ trợ tốt hơn cho
module Go.

{{image "gopls/features.gif" 635}}
_(`gopls` cung cấp các tính năng IDE, như tự động hoàn thành thông minh, trợ giúp chữ ký, tái cấu trúc và tìm kiếm ký hiệu workspace.)_

Khi [module Go](using-go-modules) được
phát hành hai năm trước, chúng đã thay đổi hoàn toàn cảnh quan của công cụ nhà phát triển Go. Các công cụ như `goimports` và `godef` trước đây phụ thuộc vào thực tế
rằng code được lưu trữ trong `$GOPATH` của bạn. Khi nhóm Go bắt đầu viết lại các
công cụ này để hoạt động với module, chúng tôi ngay lập tức nhận ra rằng chúng tôi cần một cách tiếp cận có hệ thống hơn để thu hẹp khoảng cách.

Kết quả là, chúng tôi bắt đầu làm việc trên một
[language server](https://microsoft.github.io/language-server-protocol/) Go duy nhất,
`gopls`, cung cấp các tính năng IDE, như tự động hoàn thành, định dạng và
chẩn đoán cho bất kỳ frontend trình soạn thảo tương thích nào. Server liên tục và thống nhất này là một [thay đổi căn bản](https://www.youtube.com/watch?v=EFJfdWzBHwE&t=1s) so với các tập hợp công cụ dòng lệnh trước đó.

Ngoài làm việc trên `gopls`, chúng tôi đã tìm kiếm các cách khác để tạo ra một
hệ sinh thái ổn định của công cụ trình soạn thảo. Năm ngoái, nhóm Go đã tiếp nhận trách nhiệm về
[extension Go cho VS Code](/blog/vscode-go). Là một phần của công việc này, chúng tôi đã làm mịn tích hợp của extension với language server - tự động hóa
cập nhật `gopls`, sắp xếp lại và làm rõ cài đặt `gopls`, cải thiện quy trình khắc phục sự cố và thu thập phản hồi qua khảo sát. Chúng tôi cũng đã
tiếp tục nuôi dưỡng một cộng đồng người dùng và cộng tác viên tích cực đã
giúp chúng tôi cải thiện sự ổn định, hiệu năng và trải nghiệm người dùng của extension Go.

## Thông báo

Ngày 28 tháng 1 đánh dấu một cột mốc quan trọng trong cả hành trình `gopls` và VS Code Go,
khi `gopls` giờ được bật mặc định trong extension Go cho VS Code.

Trước khi chuyển đổi này, chúng tôi đã dành nhiều thời gian lặp đi lặp lại trên thiết kế, bộ tính năng
và trải nghiệm người dùng của `gopls`, tập trung vào việc cải thiện hiệu năng và
sự ổn định. Trong hơn một năm, `gopls` đã là mặc định trong hầu hết các plugin cho
Vim, Emacs và các trình soạn thảo khác. Chúng tôi đã có 24 bản phát hành `gopls`, và chúng tôi
vô cùng biết ơn người dùng của mình vì đã liên tục cung cấp phản hồi và
báo cáo sự cố cho từng bản phát hành.

Chúng tôi cũng đã dành thời gian để làm mịn trải nghiệm người dùng mới. Chúng tôi hy vọng rằng VS
Code Go với `gopls` sẽ trực quan với các thông báo lỗi rõ ràng, nhưng nếu bạn có
câu hỏi hoặc cần điều chỉnh một số cấu hình, bạn có thể tìm câu trả lời
trong [tài liệu cập nhật của chúng tôi](https://github.com/golang/vscode-go/blob/master/README.md).
Chúng tôi cũng đã ghi lại [một screencast](https://www.youtube.com/watch?v=1MXIGYrMk80)
để giúp bạn bắt đầu, cũng như
[animation](https://github.com/golang/vscode-go/blob/master/docs/features.md)
để hiển thị một số tính năng khó tìm.

Gopls là cách tốt nhất để làm việc với code Go, đặc biệt với module Go.
Với sự xuất hiện của Go 1.16 sắp tới, trong đó module được bật mặc định,
người dùng VS Code Go sẽ có trải nghiệm tốt nhất có thể ngay từ đầu.

Tuy nhiên, việc chuyển đổi này không có nghĩa là `gopls` đã hoàn chỉnh. Chúng tôi sẽ tiếp tục
làm việc về sửa lỗi, tính năng mới và sự ổn định chung. Lĩnh vực tập trung tiếp theo của chúng tôi sẽ là cải thiện trải nghiệm người dùng khi [làm việc với nhiều module](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md).
Phản hồi từ cơ sở người dùng lớn hơn của chúng tôi sẽ giúp định hướng các bước tiếp theo.

## Vậy, bạn nên làm gì?

Nếu bạn dùng VS Code, bạn không cần làm gì.
Khi bạn nhận được bản cập nhật VS Code Go tiếp theo, `gopls` sẽ được bật tự động.

Nếu bạn dùng trình soạn thảo khác, bạn có thể đã đang dùng `gopls`. Nếu chưa, xem
[hướng dẫn người dùng `gopls`](https://github.com/golang/tools/blob/master/gopls/README.md)
để tìm hiểu cách bật `gopls` trong trình soạn thảo ưa thích của bạn. Language Server
Protocol đảm bảo rằng `gopls` sẽ tiếp tục cung cấp các tính năng tương tự cho mọi
trình soạn thảo.

Nếu `gopls` không hoạt động với bạn, hãy xem [hướng dẫn khắc phục sự cố chi tiết](https://github.com/golang/vscode-go/blob/master/docs/troubleshooting.md)
của chúng tôi và tạo issue. Nếu cần, bạn luôn có thể [vô hiệu hóa `gopls` trong VS Code](https://github.com/golang/vscode-go/blob/master/docs/settings.md#gouselanguageserver).

## Cảm ơn

Với người dùng hiện tại, cảm ơn vì đã kiên nhẫn với chúng tôi khi chúng tôi viết lại lớp cache
lần thứ ba. Với người dùng mới, chúng tôi mong được nghe báo cáo trải nghiệm và phản hồi của bạn.

Cuối cùng, không có cuộc thảo luận nào về công cụ Go hoàn chỉnh mà không đề cập đến
những đóng góp quý giá của cộng đồng công cụ Go. Cảm ơn vì các cuộc thảo luận dài,
báo cáo lỗi chi tiết, kiểm thử tích hợp, và quan trọng nhất,
cảm ơn vì những đóng góp tuyệt vời. Các tính năng `gopls` thú vị nhất
đến từ những người đóng góp mã nguồn mở đầy nhiệt huyết, và chúng tôi đánh giá cao
sự cần cù và tận tụy của bạn.

## Tìm hiểu thêm

Xem [screencast](https://www.youtube.com/watch?v=1MXIGYrMk80) để
hướng dẫn cách bắt đầu với `gopls` và VS Code Go, và xem
[VS Code Go README](https://github.com/golang/vscode-go/blob/master/README.md)
để biết thêm thông tin.

Nếu bạn muốn đọc về `gopls` chi tiết hơn, xem
[`gopls` README](https://github.com/golang/tools/blob/master/gopls/README.md).
