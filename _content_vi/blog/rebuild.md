---
title: Bộ công cụ Go hoàn toàn có thể tái tạo và xác minh
date: 2023-08-28
by:
- Russ Cox
summary: Go 1.21 là bộ công cụ Go đầu tiên có thể tái tạo hoàn toàn.
template: true
---

Một trong những lợi ích chính của phần mềm mã nguồn mở là bất kỳ ai cũng có thể đọc
mã nguồn và kiểm tra những gì nó làm.
Tuy nhiên, hầu hết phần mềm, ngay cả phần mềm mã nguồn mở,
được tải xuống dưới dạng các nhị phân đã được biên dịch,
vốn khó kiểm tra hơn nhiều.
Nếu một kẻ tấn công muốn thực hiện [tấn công chuỗi cung ứng](https://cloud.google.com/software-supply-chain-security/docs/attack-vectors)
trên một dự án mã nguồn mở,
cách ít bị chú ý nhất là thay thế các tệp nhị phân đang được phục vụ trong khi
giữ nguyên mã nguồn.

Cách tốt nhất để giải quyết loại tấn công này là làm cho các bản dựng phần mềm mã nguồn mở
_có thể tái tạo_,
có nghĩa là một bản dựng bắt đầu với cùng nguồn sẽ tạo ra cùng
đầu ra mỗi lần nó chạy.
Bằng cách đó, bất kỳ ai cũng có thể xác minh rằng các nhị phân được đăng không có thay đổi ẩn
bằng cách dựng từ nguồn xác thực và kiểm tra rằng các nhị phân được dựng lại
giống hệt bit-for-bit với các nhị phân được đăng.
Cách tiếp cận đó chứng minh các nhị phân không có backdoor hoặc các thay đổi khác không
có trong mã nguồn,
mà không cần phải tháo rời hoặc xem bên trong chúng chút nào.
Vì bất kỳ ai cũng có thể xác minh các nhị phân, các nhóm độc lập có thể dễ dàng phát hiện
và báo cáo các cuộc tấn công chuỗi cung ứng.

Khi bảo mật chuỗi cung ứng ngày càng quan trọng,
các bản dựng có thể tái tạo cũng vậy, vì chúng cung cấp cách đơn giản để xác minh
các nhị phân được đăng cho các dự án mã nguồn mở.

Go 1.21.0 là bộ công cụ Go đầu tiên với các bản dựng hoàn toàn có thể tái tạo.
Các bộ công cụ trước đây có thể tái tạo được,
nhưng chỉ với nỗ lực đáng kể, và có lẽ không ai làm:
họ chỉ tin tưởng rằng các nhị phân được đăng trên [go.dev/dl](/dl/) là đúng.
Bây giờ thật dễ dàng để "tin tưởng nhưng xác minh."

Bài viết này giải thích những gì cần thiết để làm cho các bản dựng có thể tái tạo,
xem xét nhiều thay đổi chúng tôi phải thực hiện với Go để làm cho các bộ công cụ Go có thể tái tạo,
và sau đó minh họa một trong những lợi ích của khả năng tái tạo bằng cách xác minh
gói Ubuntu cho Go 1.21.0.

## Làm cho một bản dựng có thể tái tạo {#how}

Máy tính thường là xác định, vì vậy bạn có thể nghĩ rằng tất cả các bản dựng sẽ
có thể tái tạo như nhau.
Điều đó chỉ đúng từ một quan điểm nhất định.
Hãy gọi một thông tin là _đầu vào liên quan_ khi đầu ra của
bản dựng có thể thay đổi tùy thuộc vào đầu vào đó.
Một bản dựng có thể tái tạo nếu nó có thể được lặp lại với tất cả các đầu vào liên quan giống nhau.
Thật không may, nhiều công cụ dựng hóa ra tích hợp các đầu vào mà chúng ta
thường không nhận ra là liên quan và có thể khó tái tạo
hoặc cung cấp làm đầu vào.
Hãy gọi một đầu vào là _đầu vào không có chủ đích_ khi nó hóa ra là liên quan
nhưng chúng ta không có ý định như vậy.

Đầu vào không có chủ đích phổ biến nhất trong các hệ thống dựng là thời gian hiện tại.
Nếu một bản dựng ghi một tệp thực thi vào đĩa, hệ thống tệp ghi lại thời gian hiện tại
như là thời gian sửa đổi của tệp thực thi.
Nếu bản dựng sau đó đóng gói tệp đó bằng một công cụ như "tar" hoặc "zip",
thời gian sửa đổi được ghi vào archive.
Chúng tôi chắc chắn không muốn bản dựng của chúng tôi thay đổi dựa trên thời gian hiện tại, nhưng nó thay đổi.
Vì vậy, thời gian hiện tại hóa ra là một đầu vào không có chủ đích cho bản dựng.
Tệ hơn nữa, hầu hết các chương trình không cho phép bạn cung cấp thời gian hiện tại như một đầu vào,
vì vậy không có cách nào để lặp lại bản dựng này.
Để sửa điều này, chúng tôi có thể đặt dấu thời gian trên các tệp được tạo thành Unix time 0
hoặc thành một thời gian cụ thể được đọc từ một trong các tệp nguồn của bản dựng.
Bằng cách đó, thời gian hiện tại không còn là đầu vào liên quan cho bản dựng nữa.

Các đầu vào liên quan phổ biến cho một bản dựng bao gồm:

  - phiên bản cụ thể của mã nguồn để dựng;
  - các phiên bản cụ thể của các dependency sẽ được đưa vào bản dựng;
  - hệ điều hành đang chạy bản dựng, có thể ảnh hưởng đến tên đường dẫn trong các nhị phân kết quả;
  - kiến trúc CPU trên hệ thống dựng,
    có thể ảnh hưởng đến tối ưu hóa nào trình biên dịch sử dụng hoặc bố cục của một số cấu trúc dữ liệu;
  - phiên bản trình biên dịch được sử dụng, cũng như các tùy chọn trình biên dịch được truyền cho nó, ảnh hưởng đến cách mã được biên dịch;
  - tên của thư mục chứa mã nguồn, có thể xuất hiện trong thông tin debug;
  - tên người dùng, tên nhóm, uid và gid của tài khoản chạy bản dựng, có thể xuất hiện trong metadata tệp trong một archive;
  - và nhiều hơn nữa.

Để có một bản dựng có thể tái tạo, mọi đầu vào liên quan phải có thể cấu hình trong bản dựng,
và sau đó các nhị phân phải được đăng cùng với một cấu hình rõ ràng
liệt kê mọi đầu vào liên quan.
Nếu bạn đã làm điều đó, bạn có một bản dựng có thể tái tạo. Chúc mừng!

Tuy nhiên chúng ta chưa xong. Nếu các nhị phân chỉ có thể tái tạo nếu bạn
trước tiên tìm một máy tính với kiến trúc đúng,
cài đặt một phiên bản hệ điều hành cụ thể,
phiên bản trình biên dịch, đặt mã nguồn vào đúng thư mục,
đặt danh tính người dùng của bạn đúng, v.v.,
điều đó có thể quá nhiều công việc trên thực tế để ai đó bận tâm.

Chúng tôi muốn các bản dựng không chỉ có thể tái tạo mà còn _dễ tái tạo_.
Để làm điều đó, chúng tôi cần xác định các đầu vào liên quan và sau đó,
thay vì ghi lại chúng, loại bỏ chúng.
Bản dựng rõ ràng phải phụ thuộc vào mã nguồn đang được dựng,
nhưng mọi thứ khác đều có thể loại bỏ.
Khi đầu vào liên quan duy nhất của bản dựng là mã nguồn của nó,
hãy gọi điều đó là _hoàn toàn có thể tái tạo_.

## Các bản dựng hoàn toàn có thể tái tạo cho Go {#go}

Kể từ Go 1.21, bộ công cụ Go hoàn toàn có thể tái tạo:
đầu vào liên quan duy nhất của nó là mã nguồn cho bản dựng đó.
Chúng tôi có thể dựng một bộ công cụ cụ thể (ví dụ: Go cho Linux/x86-64) trên một host Linux/x86-64,
hoặc một host Windows/ARM64, hoặc một host FreeBSD/386,
hoặc bất kỳ host nào khác hỗ trợ Go, và chúng tôi có thể sử dụng bất kỳ trình biên dịch bootstrap Go nào,
bao gồm cả việc bootstrap tất cả cách ngược lại đến cài đặt C của Go 1.4,
và chúng tôi có thể thay đổi bất kỳ chi tiết nào khác.
Không có gì trong số đó thay đổi các bộ công cụ được dựng.
Nếu chúng tôi bắt đầu với cùng mã nguồn bộ công cụ,
chúng tôi sẽ nhận được chính xác cùng các nhị phân bộ công cụ.

Khả năng tái tạo hoàn toàn này là kết quả của các nỗ lực có nguồn gốc ban đầu từ Go 1.10,
mặc dù hầu hết nỗ lực tập trung vào Go 1.20 và Go 1.21.
Phần này nêu bật một số đầu vào liên quan thú vị nhất mà chúng tôi đã loại bỏ.

### Khả năng tái tạo trong Go 1.10 {#go110}

Go 1.10 đã giới thiệu một cache dựng nhận biết nội dung quyết định liệu các mục tiêu
có cập nhật hay không dựa trên fingerprint của các đầu vào dựng thay vì thời gian sửa đổi tệp.
Vì bản thân bộ công cụ là một trong các đầu vào dựng đó,
và vì Go được viết bằng Go, [quá trình bootstrap](/s/go15bootstrap)
chỉ hội tụ nếu bản dựng bộ công cụ trên một máy đơn là có thể tái tạo.
Quá trình dựng bộ công cụ tổng thể trông như thế này:

<div class="image">
<img src="rebuild/bootstrap.png" srcset="rebuild/bootstrap.png 1x, rebuild/bootstrap@2x.png 2x" width="515" height="177">
</div>

Chúng tôi bắt đầu bằng cách dựng các nguồn cho bộ công cụ Go hiện tại sử dụng một phiên bản Go cũ hơn,
bộ công cụ bootstrap (Go 1.10 sử dụng Go 1.4, được viết bằng C;
Go 1.21 sử dụng Go 1.17).
Điều đó tạo ra "toolchain1", mà chúng tôi sử dụng để dựng mọi thứ lại,
tạo ra "toolchain2", mà chúng tôi sử dụng để dựng mọi thứ lại,
tạo ra "toolchain3".

Toolchain1 và toolchain2 đã được dựng từ cùng nguồn nhưng với
các cài đặt Go khác nhau (trình biên dịch và thư viện),
vì vậy các nhị phân của chúng nhất định sẽ khác nhau.
Tuy nhiên, nếu cả hai cài đặt Go đều không có lỗi,
cài đặt đúng, toolchain1 và toolchain2 nên hoạt động hoàn toàn giống nhau.
Đặc biệt, khi được cung cấp các nguồn Go 1.X,
đầu ra của toolchain1 (toolchain2) và đầu ra của toolchain2 (toolchain3)
nên giống hệt nhau,
có nghĩa là toolchain2 và toolchain3 nên giống hệt nhau.

Ít nhất, đó là ý tưởng. Để thực hiện điều đó trong thực tế đòi hỏi phải loại bỏ một vài đầu vào không có chủ đích:

**Ngẫu nhiên.** Việc lặp map và chạy công việc trong nhiều goroutine được tuần tự hóa
với khóa đều mang đến sự ngẫu nhiên trong thứ tự kết quả có thể được tạo ra.
Sự ngẫu nhiên này có thể làm cho bộ công cụ tạo ra một trong nhiều
đầu ra có thể khác nhau mỗi lần nó chạy.
Để làm cho bản dựng có thể tái tạo, chúng tôi phải tìm từng cái trong số này và sắp xếp
danh sách liên quan các mục trước khi sử dụng nó để tạo đầu ra.

**Thư viện Bootstrap.** Bất kỳ thư viện nào được trình biên dịch sử dụng có thể chọn
từ nhiều đầu ra đúng khác nhau có thể thay đổi đầu ra của nó từ một
phiên bản Go sang phiên bản tiếp theo.
Nếu sự thay đổi đầu ra thư viện đó gây ra sự thay đổi đầu ra trình biên dịch,
thì toolchain1 và toolchain2 sẽ không giống hệt về mặt ngữ nghĩa,
và toolchain2 và toolchain3 sẽ không giống hệt bit-for-bit.

Ví dụ chuẩn là gói [`sort`](/pkg/sort/),
có thể đặt các phần tử so sánh bằng nhau theo [bất kỳ thứ tự nào nó thích](/blog/compat#output).
Một trình phân bổ register có thể sắp xếp để ưu tiên các biến được sử dụng thường xuyên,
và linker sắp xếp các symbol trong phần dữ liệu theo kích thước.
Để hoàn toàn loại bỏ bất kỳ ảnh hưởng nào từ thuật toán sắp xếp,
hàm so sánh được sử dụng không bao giờ được báo cáo hai phần tử phân biệt là bằng nhau.
Trên thực tế, bất biến này hóa ra quá nặng nề để áp đặt lên mọi
sử dụng sort trong bộ công cụ,
vì vậy thay vào đó chúng tôi đã sắp xếp để sao chép gói `sort` Go 1.X vào cây nguồn
được trình bày cho trình biên dịch bootstrap.
Bằng cách đó, trình biên dịch sử dụng cùng thuật toán sort khi sử dụng bộ công cụ bootstrap
như khi được dựng với chính nó.

Một gói khác chúng tôi phải sao chép là [`compress/zlib`](/pkg/compress/zlib/),
vì linker ghi thông tin debug được nén,
và các tối ưu hóa cho thư viện nén có thể thay đổi đầu ra chính xác.
Theo thời gian, chúng tôi đã [thêm các gói khác vào danh sách đó](https://go.googlesource.com/go/+/go1.21.0/src/cmd/dist/buildtool.go#55).
Cách tiếp cận này có lợi ích bổ sung là cho phép trình biên dịch Go 1.X sử dụng
các API mới được thêm vào những gói đó ngay lập tức,
với chi phí là những gói đó phải được viết để biên dịch với các phiên bản Go cũ hơn.

### Khả năng tái tạo trong Go 1.20 {#go120}

Công việc trên Go 1.20 đã chuẩn bị cho cả các bản dựng dễ tái tạo và [quản lý bộ công cụ](toolchain)
bằng cách loại bỏ thêm hai đầu vào liên quan từ bản dựng bộ công cụ.

**Bộ công cụ C host.** Một số gói Go, đặc biệt là `net`,
mặc định [sử dụng `cgo`](cgo) trên hầu hết các hệ điều hành.
Trong một số trường hợp, chẳng hạn như macOS và Windows,
gọi các DLL hệ thống sử dụng `cgo` là cách đáng tin cậy duy nhất để phân giải tên host.
Khi chúng ta sử dụng `cgo`, tuy nhiên, chúng ta gọi bộ công cụ C host (có nghĩa là một cụ thể
trình biên dịch C và thư viện C),
và các bộ công cụ khác nhau có các thuật toán biên dịch và mã thư viện khác nhau,
tạo ra các đầu ra khác nhau.
Biểu đồ dựng cho một gói `cgo` trông như thế này:

<div class="image">
<img src="rebuild/cgo.png" srcset="rebuild/cgo.png 1x, rebuild/cgo@2x.png 2x" width="441" height="344">
</div>

Bộ công cụ C host do đó là một đầu vào liên quan cho `net.a` được biên dịch trước
đi kèm với bộ công cụ.
Đối với Go 1.20, chúng tôi quyết định sửa điều này bằng cách xóa `net.a` khỏi bộ công cụ.
Tức là, Go 1.20 đã ngừng vận chuyển các gói được biên dịch trước để khởi động cache dựng.
Bây giờ, lần đầu tiên một chương trình sử dụng gói `net`,
bộ công cụ Go biên dịch nó bằng cách sử dụng bộ công cụ C của hệ thống cục bộ và lưu kết quả vào cache.
Ngoài việc loại bỏ một đầu vào liên quan khỏi các bản dựng bộ công cụ và làm cho
tải xuống bộ công cụ nhỏ hơn,
không vận chuyển các gói được biên dịch trước cũng làm cho tải xuống bộ công cụ dễ chuyển đổi hơn.
Nếu chúng tôi dựng gói `net` trên một hệ thống với một bộ công cụ C và sau đó biên dịch
các phần khác của chương trình trên một hệ thống khác với một bộ công cụ C khác,
nói chung không có đảm bảo rằng hai phần có thể được link với nhau.

Một lý do chúng tôi đã vận chuyển gói `net` được biên dịch trước lần đầu tiên
là để cho phép dựng các chương trình sử dụng gói net ngay cả trên các hệ thống không có
bộ công cụ C được cài đặt.
Nếu không có gói được biên dịch trước, điều gì xảy ra trên những hệ thống đó? Câu trả lời
thay đổi theo hệ điều hành,
nhưng trong tất cả các trường hợp chúng tôi đã sắp xếp để bộ công cụ Go tiếp tục hoạt động tốt
để dựng các chương trình Go thuần túy mà không có bộ công cụ C host.

  - Trên macOS, chúng tôi đã viết lại gói net bằng cách sử dụng các cơ chế cơ bản mà cgo sẽ sử dụng,
    mà không có bất kỳ mã C thực sự nào.
    Điều này tránh gọi bộ công cụ C host nhưng vẫn phát ra một nhị phân
    tham chiếu đến các DLL hệ thống cần thiết.
    Cách tiếp cận này chỉ có thể thực hiện vì mọi Mac đều có cùng thư viện động được cài đặt.
    Làm cho gói net macOS không phải cgo sử dụng các DLL hệ thống cũng có nghĩa là
    các tệp thực thi macOS được biên dịch chéo giờ đây sử dụng các DLL hệ thống để truy cập mạng,
    giải quyết một yêu cầu tính năng lâu dài.

  - Trên Windows, gói net đã trực tiếp sử dụng các DLL mà không có mã C, vì vậy không cần thay đổi gì.

  - Trên các hệ thống Unix, chúng tôi không thể giả định một giao diện DLL cụ thể cho mã mạng,
    nhưng phiên bản Go thuần túy hoạt động tốt cho các hệ thống sử dụng cấu hình IP và DNS điển hình.
    Ngoài ra, cài đặt bộ công cụ C trên các hệ thống Unix dễ dàng hơn nhiều so với
    macOS và đặc biệt là Windows.
    Chúng tôi đã thay đổi lệnh `go` để bật hoặc tắt `cgo` tự động dựa trên
    việc hệ thống có bộ công cụ C được cài đặt hay không.
    Các hệ thống Unix không có bộ công cụ C sẽ quay lại phiên bản Go thuần túy của gói net,
    và trong những trường hợp hiếm hoi mà điều đó không đủ tốt,
    họ có thể cài đặt bộ công cụ C.

Sau khi bỏ các gói được biên dịch trước,
phần duy nhất của bộ công cụ Go vẫn phụ thuộc vào bộ công cụ C host
là các nhị phân được dựng bằng gói net,
cụ thể là lệnh `go`.
Với các cải tiến macOS, bây giờ có thể dựng những lệnh đó với `cgo` bị tắt,
loại bỏ hoàn toàn bộ công cụ C host như một đầu vào,
nhưng chúng tôi đã để lại bước cuối cùng đó cho Go 1.21.
(Chúng tôi không cảm thấy có đủ thời gian còn lại trong chu kỳ Go 1.20 để kiểm tra
thay đổi đó đúng cách.)

**Linker động host.** Khi các chương trình sử dụng `cgo` trên một hệ thống sử dụng các thư viện C được liên kết động,
các nhị phân kết quả chứa đường dẫn đến linker động của hệ thống,
điều gì đó như `/lib64/ld-linux-x86-64.so.2`.
Nếu đường dẫn sai, các nhị phân không chạy.
Thông thường mỗi tổ hợp hệ điều hành/kiến trúc có một câu trả lời đúng duy nhất
cho đường dẫn này.
Thật không may, các Linux dựa trên musl như Alpine Linux sử dụng một linker động khác
so với các Linux dựa trên glibc như Ubuntu.
Để làm cho Go chạy được trên Alpine Linux, quá trình bootstrap Go trông như thế này:

<div class="image">
<img src="rebuild/linker1.png" srcset="rebuild/linker1.png 1x, rebuild/linker1@2x.png 2x" width="480" height="209">
</div>

Chương trình bootstrap cmd/dist đã kiểm tra linker động của hệ thống cục bộ
và viết giá trị đó vào một tệp nguồn mới được biên dịch cùng với phần còn lại
của các nguồn linker,
thực sự cứng nhắc hóa mặc định đó vào bản thân linker.
Sau đó khi linker dựng một chương trình từ một tập hợp các gói được biên dịch,
nó sử dụng mặc định đó.
Kết quả là một bộ công cụ Go được dựng trên Alpine khác với một bộ công cụ được dựng trên Ubuntu:
cấu hình host là một đầu vào liên quan cho bản dựng bộ công cụ.
Đây là một vấn đề khả năng tái tạo nhưng cũng là một vấn đề về tính di động:
một bộ công cụ Go được dựng trên Alpine không dựng các nhị phân hoạt động hoặc thậm chí
không chạy trên Ubuntu, và ngược lại.

Đối với Go 1.20, chúng tôi đã thực hiện một bước tiến tới việc sửa vấn đề khả năng tái tạo bằng cách
thay đổi linker để tham khảo cấu hình host khi nó đang chạy,
thay vì có một mặc định được cứng nhắc hóa vào thời gian dựng bộ công cụ:

<div class="image">
<img src="rebuild/linker2.png" srcset="rebuild/linker2.png 1x, rebuild/linker2@2x.png 2x" width="450" height="175">
</div>

Điều này đã sửa tính di động của nhị phân linker trên Alpine Linux,
mặc dù không phải bộ công cụ tổng thể, vì lệnh `go` vẫn sử dụng gói
`net` và do đó `cgo` và do đó có một tham chiếu linker động trong nhị phân của chính nó.
Giống như trong phần trước, biên dịch lệnh `go` với `cgo`
bị tắt sẽ sửa điều này,
nhưng chúng tôi đã để lại thay đổi đó cho Go 1.21.
(Chúng tôi không cảm thấy có đủ thời gian còn lại trong chu kỳ Go 1.20 để kiểm tra
thay đổi đó đúng cách.)

### Khả năng tái tạo trong Go 1.21 {#go121}

Đối với Go 1.21, mục tiêu về khả năng tái tạo hoàn toàn đã trong tầm nhìn,
và chúng tôi đã xử lý các đầu vào liên quan còn lại, hầu hết nhỏ.

**Bộ công cụ C host và linker động.** Như đã thảo luận ở trên,
Go 1.20 đã thực hiện các bước quan trọng để loại bỏ bộ công cụ C host và linker
động như các đầu vào liên quan.
Go 1.21 đã hoàn thành việc loại bỏ các đầu vào liên quan này bằng cách dựng bộ công cụ
với `cgo` bị tắt.
Điều này cũng cải thiện tính di động của bộ công cụ:
Go 1.21 là bản phát hành Go đầu tiên mà bộ công cụ Go chuẩn chạy không sửa đổi
trên các hệ thống Alpine Linux.

Việc loại bỏ các đầu vào liên quan này đã cho phép biên dịch chéo một bộ công cụ Go
từ một hệ thống khác mà không mất bất kỳ chức năng nào.
Điều đó đến lượt cải thiện bảo mật chuỗi cung ứng của bộ công cụ Go:
chúng tôi bây giờ có thể dựng các bộ công cụ Go cho tất cả các hệ thống đích bằng cách sử dụng một hệ thống Linux/x86-64 đáng tin cậy,
thay vì cần phải sắp xếp một hệ thống đáng tin cậy riêng cho mỗi đích.
Kết quả là, Go 1.21 là bản phát hành đầu tiên bao gồm các nhị phân được đăng cho
tất cả các hệ thống tại [go.dev/dl/](/dl/).

**Thư mục nguồn.** Các chương trình Go bao gồm các đường dẫn đầy đủ trong metadata runtime và debug,
để khi một chương trình bị crash hoặc được chạy trong debugger,
stack trace bao gồm đường dẫn đầy đủ đến tệp nguồn,
không chỉ là tên của tệp trong một thư mục không xác định.
Thật không may, bao gồm đường dẫn đầy đủ làm cho thư mục nơi mã nguồn
được lưu trữ là một đầu vào liên quan cho bản dựng.
Để sửa điều này, Go 1.21 đã thay đổi các bản dựng bộ công cụ phát hành để cài đặt các lệnh
như trình biên dịch sử dụng `go install -trimpath`,
thay thế thư mục nguồn bằng đường dẫn module của mã.
Nếu một trình biên dịch được phát hành bị crash, stack trace sẽ in các đường dẫn như `cmd/compile/main.go`
thay vì `/home/user/go/src/cmd/compile/main.go`.
Vì các đường dẫn đầy đủ sẽ tham chiếu đến một thư mục trên một máy khác dù sao,
việc viết lại này không mất mát gì.
Mặt khác, đối với các bản dựng không phải phát hành,
chúng tôi giữ đường dẫn đầy đủ, để khi các nhà phát triển đang làm việc trên trình biên dịch tự gây ra crash,
các IDE và các công cụ khác đọc những crash đó có thể dễ dàng tìm thấy tệp nguồn đúng.

**Hệ điều hành host.** Các đường dẫn trên hệ thống Windows được phân cách bằng dấu gạch chéo ngược,
như `cmd\compile\main.go`.
Các hệ thống khác sử dụng dấu gạch chéo xuôi, như `cmd/compile/main.go`.
Mặc dù các phiên bản Go trước đó đã chuẩn hóa hầu hết các đường dẫn này để sử dụng dấu gạch chéo xuôi,
một sự không nhất quán đã trở lại, gây ra các bản dựng bộ công cụ hơi khác nhau trên Windows.
Chúng tôi đã tìm và sửa lỗi.

**Kiến trúc host.** Go chạy trên nhiều hệ thống ARM và có thể phát ra
mã sử dụng thư viện phần mềm cho phép tính dấu phẩy động (SWFP) hoặc sử dụng phần cứng
các lệnh dấu phẩy động (HWFP).
Các bộ công cụ mặc định theo một trong hai chế độ sẽ nhất thiết khác nhau.
Như chúng ta đã thấy với linker động trước đó,
quá trình bootstrap Go đã kiểm tra hệ thống dựng để đảm bảo rằng
bộ công cụ kết quả hoạt động trên hệ thống đó.
Vì lý do lịch sử, quy tắc là "giả định SWFP trừ khi bản dựng đang
chạy trên hệ thống ARM với phần cứng dấu phẩy động",
với các bộ công cụ được biên dịch chéo giả định SWFP.
Đại đa số các hệ thống ARM ngày nay có phần cứng dấu phẩy động,
vì vậy điều này đã tạo ra một sự khác biệt không cần thiết giữa các bộ công cụ được biên dịch tự nhiên và
các bộ công cụ được biên dịch chéo,
và như một phức tạp thêm, các bản dựng Windows ARM luôn giả định HWFP,
làm cho quyết định phụ thuộc vào hệ điều hành.
Chúng tôi đã thay đổi quy tắc thành "giả định HWFP trừ khi bản dựng đang chạy trên
hệ thống ARM không có phần cứng dấu phẩy động".
Bằng cách này, biên dịch chéo và các bản dựng trên các hệ thống ARM hiện đại tạo ra các bộ công cụ giống hệt nhau.

**Logic đóng gói.** Tất cả mã để tạo các archive bộ công cụ thực sự
chúng tôi đăng để tải xuống sống trong một kho Git riêng,
golang.org/x/build, và các chi tiết chính xác về cách các archive được đóng gói thay đổi theo thời gian.
Nếu bạn muốn tái tạo những archive đó,
bạn cần có phiên bản đúng của kho lưu trữ đó.
Chúng tôi đã loại bỏ đầu vào liên quan này bằng cách di chuyển mã để đóng gói các archive
vào cây nguồn Go chính, dưới dạng `cmd/distpack`.
Kể từ Go 1.21, nếu bạn có nguồn cho một phiên bản Go nhất định,
bạn cũng có nguồn để đóng gói các archive.
Kho lưu trữ golang.org/x/build không còn là đầu vào liên quan nữa.

**ID người dùng.** Các archive tar chúng tôi đăng để tải xuống đã được dựng từ một
bản phân phối được ghi vào hệ thống tệp,
và sử dụng [`tar.FileInfoHeader`](/pkg/archive/tar/#FileInfoHeader) sao chép
các ID người dùng và nhóm từ hệ thống tệp vào tệp tar,
làm cho người dùng chạy bản dựng là một đầu vào liên quan.
Chúng tôi đã thay đổi mã lưu trữ để xóa những ID này.

**Thời gian hiện tại.** Như với ID người dùng, các archive tar và zip chúng tôi đăng
để tải xuống đã được dựng bằng cách sao chép thời gian sửa đổi tệp hệ thống vào các archive,
làm cho thời gian hiện tại là một đầu vào liên quan.
Chúng tôi có thể đã xóa thời gian, nhưng chúng tôi nghĩ nó sẽ trông ngạc nhiên
và thậm chí có thể làm hỏng một số công cụ để sử dụng thời gian zero của Unix hoặc MS-DOS.
Thay vào đó, chúng tôi đã thay đổi tệp go/VERSION được lưu trữ trong kho lưu trữ để thêm
thời gian liên quan đến phiên bản đó:

	$ cat go1.21.0/VERSION
	go1.21.0
	time 2023-08-04T20:14:06Z
	$

Các trình đóng gói bây giờ sao chép thời gian từ tệp VERSION khi ghi tệp vào archive,
thay vì sao chép thời gian sửa đổi tệp cục bộ.

**Khóa ký mật mã.** Bộ công cụ Go cho macOS sẽ không chạy trên
hệ thống người dùng cuối trừ khi chúng tôi ký các nhị phân với khóa ký được Apple chấp thuận.
Chúng tôi sử dụng một hệ thống nội bộ để ký chúng với khóa ký của Google,
và rõ ràng chúng tôi không thể chia sẻ khóa bí mật đó để cho phép người khác
tái tạo các nhị phân đã ký.
Thay vào đó, chúng tôi đã viết một trình xác minh có thể kiểm tra liệu hai nhị phân có giống hệt nhau không
ngoại trừ chữ ký của chúng.

**Trình đóng gói dành riêng cho OS.** Chúng tôi sử dụng các công cụ Xcode `pkgbuild` và `productbuild`
để tạo trình cài đặt macOS PKG có thể tải xuống,
và chúng tôi sử dụng WiX để tạo trình cài đặt Windows MSI có thể tải xuống.
Chúng tôi không muốn người xác minh cần cùng các phiên bản chính xác của những công cụ đó,
vì vậy chúng tôi đã sử dụng cùng cách tiếp cận như cho các khóa ký mật mã,
viết một trình xác minh có thể nhìn bên trong các gói và kiểm tra rằng các tệp
bộ công cụ chính xác như mong đợi.

## Xác minh các bộ công cụ Go {#verify}

Chỉ làm cho các bộ công cụ Go có thể tái tạo một lần là không đủ.
Chúng tôi muốn đảm bảo chúng vẫn có thể tái tạo,
và chúng tôi muốn đảm bảo người khác có thể tái tạo chúng dễ dàng.

Để giữ bản thân trung thực, chúng tôi bây giờ dựng tất cả các bản phân phối Go trên cả một
hệ thống Linux/x86-64 đáng tin cậy và một hệ thống Windows/x86-64.
Ngoại trừ kiến trúc, hai hệ thống hầu như không có gì chung.
Hai hệ thống phải tạo ra các archive giống hệt nhau bit-for-bit, nếu không thì chúng tôi
không tiến hành bản phát hành.

Để cho phép người khác xác minh rằng chúng tôi trung thực,
chúng tôi đã viết và xuất bản một trình xác minh,
[`golang.org/x/build/cmd/gorebuild`](https://pkg.go.dev/golang.org/x/build/cmd/gorebuild).
Chương trình đó sẽ bắt đầu với mã nguồn trong kho Git của chúng tôi và dựng lại
các phiên bản Go hiện tại, kiểm tra rằng chúng khớp với các archive được đăng trên [go.dev/dl](/dl/).
Hầu hết các archive được yêu cầu khớp bit-for-bit.
Như đã đề cập ở trên, có ba ngoại lệ khi sử dụng kiểm tra lỏng lẻo hơn:

  - Tệp macOS tar.gz được mong đợi khác, nhưng sau đó trình xác minh so sánh nội dung bên trong.
    Các bản sao được dựng lại và được đăng phải chứa cùng tệp,
    và tất cả các tệp phải khớp chính xác, ngoại trừ các nhị phân thực thi.
    Các nhị phân thực thi phải khớp chính xác sau khi xóa chữ ký code.

  - Trình cài đặt macOS PKG không được dựng lại. Thay vào đó,
    trình xác minh đọc các tệp bên trong trình cài đặt PKG và kiểm tra rằng chúng
    khớp chính xác với macOS tar.gz,
    một lần nữa sau khi xóa chữ ký code.
    Về lâu dài, việc tạo PKG đủ đơn giản để có thể
    được thêm vào cmd/distpack,
    nhưng trình xác minh vẫn phải phân tích tệp PKG để chạy so sánh
    nhị phân thực thi bỏ qua chữ ký.

  - Trình cài đặt Windows MSI không được dựng lại.
    Thay vào đó, trình xác minh gọi chương trình Linux `msiextract` để trích xuất
    các tệp bên trong và kiểm tra rằng chúng khớp chính xác với tệp zip Windows được dựng lại.
    Về lâu dài, có thể việc tạo MSI có thể được thêm vào cmd/distpack,
    và sau đó trình xác minh có thể sử dụng so sánh MSI bit-for-bit.

Chúng tôi chạy `gorebuild` hàng đêm, đăng kết quả tại [go.dev/rebuild](/rebuild),
và tất nhiên bất kỳ ai khác cũng có thể chạy nó.

## Xác minh bộ công cụ Go của Ubuntu {#ubuntu}

Các bản dựng dễ tái tạo của bộ công cụ Go có nghĩa là các nhị phân
trong các bộ công cụ được đăng trên go.dev phải khớp với các nhị phân được đưa vào các hệ thống đóng gói khác,
ngay cả khi những người đóng gói đó dựng từ nguồn.
Ngay cả khi những người đóng gói đã biên dịch với các cấu hình khác nhau hoặc các thay đổi khác,
các bản dựng dễ tái tạo vẫn sẽ làm cho việc tái tạo nhị phân của họ dễ dàng.
Để minh họa điều này, hãy tái tạo gói Ubuntu `golang-1.21`
phiên bản `1.21.0-1` cho Linux/x86-64.

Để bắt đầu, chúng ta cần tải xuống và trích xuất các gói Ubuntu,
là [archive ar(1)](https://linux.die.net/man/1/ar) chứa các archive tar nén bằng zstd:

{{raw `
	$ mkdir deb
	$ cd deb
	$ curl -LO http://mirrors.kernel.org/ubuntu/pool/main/g/golang-1.21/golang-1.21-src_1.21.0-1_all.deb
	$ ar xv golang-1.21-src_1.21.0-1_all.deb
	x - debian-binary
	x - control.tar.zst
	x - data.tar.zst
	$ unzstd < data.tar.zst | tar xv
	...
	x ./usr/share/go-1.21/src/archive/tar/common.go
	x ./usr/share/go-1.21/src/archive/tar/example_test.go
	x ./usr/share/go-1.21/src/archive/tar/format.go
	x ./usr/share/go-1.21/src/archive/tar/fuzz_test.go
	...
	$
`}}

Đó là archive nguồn. Bây giờ là archive nhị phân amd64:

{{raw `
	$ rm -f debian-binary *.zst
	$ curl -LO http://mirrors.kernel.org/ubuntu/pool/main/g/golang-1.21/golang-1.21-go_1.21.0-1_amd64.deb
	$ ar xv golang-1.21-src_1.21.0-1_all.deb
	x - debian-binary
	x - control.tar.zst
	x - data.tar.zst
	$ unzstd < data.tar.zst | tar xv | grep -v '/$'
	...
	x ./usr/lib/go-1.21/bin/go
	x ./usr/lib/go-1.21/bin/gofmt
	x ./usr/lib/go-1.21/go.env
	x ./usr/lib/go-1.21/pkg/tool/linux_amd64/addr2line
	x ./usr/lib/go-1.21/pkg/tool/linux_amd64/asm
	x ./usr/lib/go-1.21/pkg/tool/linux_amd64/buildid
	...
	$
`}}

Ubuntu tách cây Go thông thường thành hai nửa,
trong /usr/share/go-1.21 và /usr/lib/go-1.21.
Hãy ghép chúng lại:

	$ mkdir go-ubuntu
	$ cp -R usr/share/go-1.21/* usr/lib/go-1.21/* go-ubuntu
	cp: cannot overwrite directory go-ubuntu/api with non-directory usr/lib/go-1.21/api
	cp: cannot overwrite directory go-ubuntu/misc with non-directory usr/lib/go-1.21/misc
	cp: cannot overwrite directory go-ubuntu/pkg/include with non-directory usr/lib/go-1.21/pkg/include
	cp: cannot overwrite directory go-ubuntu/src with non-directory usr/lib/go-1.21/src
	cp: cannot overwrite directory go-ubuntu/test with non-directory usr/lib/go-1.21/test
	$

Các lỗi đang phàn nàn về việc sao chép symlink, điều chúng ta có thể bỏ qua.

Bây giờ chúng ta cần tải xuống và trích xuất nguồn Go upstream:

	$ curl -LO https://go.googlesource.com/go/+archive/refs/tags/go1.21.0.tar.gz
	$ mkdir go-clean
	$ cd go-clean
	$ curl -L https://go.googlesource.com/go/+archive/refs/tags/go1.21.0.tar.gz | tar xzv
	...
	x src/archive/tar/common.go
	x src/archive/tar/example_test.go
	x src/archive/tar/format.go
	x src/archive/tar/fuzz_test.go
	...
	$

Để bỏ qua một số thử và sai, hóa ra Ubuntu dựng Go với `GO386=softfloat`,
buộc sử dụng dấu phẩy động phần mềm khi biên dịch cho x86 32-bit,
và strip (xóa bảng symbol) các nhị phân ELF kết quả.
Hãy bắt đầu với bản dựng `GO386=softfloat`:

	$ cd src
	$ GOOS=linux GO386=softfloat ./make.bash -distpack
	Building Go cmd/dist using /Users/rsc/sdk/go1.17.13. (go1.17.13 darwin/amd64)
	Building Go toolchain1 using /Users/rsc/sdk/go1.17.13.
	Building Go bootstrap cmd/go (go_bootstrap) using Go toolchain1.
	Building Go toolchain2 using go_bootstrap and Go toolchain1.
	Building Go toolchain3 using go_bootstrap and Go toolchain2.
	Building commands for host, darwin/amd64.
	Building packages and commands for target, linux/amd64.
	Packaging archives for linux/amd64.
	distpack: 818d46ede85682dd go1.21.0.src.tar.gz
	distpack: 4fcd8651d084a03d go1.21.0.linux-amd64.tar.gz
	distpack: eab8ed80024f444f v0.0.1-go1.21.0.linux-amd64.zip
	distpack: 58528cce1848ddf4 v0.0.1-go1.21.0.linux-amd64.mod
	distpack: d8da1f27296edea4 v0.0.1-go1.21.0.linux-amd64.info
	---
	Installed Go for linux/amd64 in /Users/rsc/deb/go-clean
	Installed commands in /Users/rsc/deb/go-clean/bin
	*** You need to add /Users/rsc/deb/go-clean/bin to your PATH.
	$

Điều đó đã để lại gói chuẩn trong `pkg/distpack/go1.21.0.linux-amd64.tar.gz`.
Hãy giải nén nó và strip các nhị phân để khớp với Ubuntu:

	$ cd ../..
	$ tar xzvf go-clean/pkg/distpack/go1.21.0.linux-amd64.tar.gz
	x go/CONTRIBUTING.md
	x go/LICENSE
	x go/PATENTS
	x go/README.md
	x go/SECURITY.md
	x go/VERSION
	...
	$ elfstrip go/bin/* go/pkg/tool/linux_amd64/*
	$

Bây giờ chúng ta có thể diff bộ công cụ Go chúng ta đã tạo trên Mac với bộ công cụ Go mà Ubuntu ship:

{{raw `
	$ diff -r go go-ubuntu
	Only in go: CONTRIBUTING.md
	Only in go: LICENSE
	Only in go: PATENTS
	Only in go: README.md
	Only in go: SECURITY.md
	Only in go: codereview.cfg
	Only in go: doc
	Only in go: lib
	Binary files go/misc/chrome/gophertool/gopher.png and go-ubuntu/misc/chrome/gophertool/gopher.png differ
	Only in go-ubuntu/pkg/tool/linux_amd64: dist
	Only in go-ubuntu/pkg/tool/linux_amd64: distpack
	Only in go/src: all.rc
	Only in go/src: clean.rc
	Only in go/src: make.rc
	Only in go/src: run.rc
	diff -r go/src/syscall/mksyscall.pl go-ubuntu/src/syscall/mksyscall.pl
	1c1
	< #!/usr/bin/env perl
	---
	> #! /usr/bin/perl
	...
	$
`}}

Chúng ta đã tái tạo thành công các tệp thực thi của gói Ubuntu và xác định
tập hợp đầy đủ các thay đổi vẫn còn:

  - Nhiều tệp metadata và tệp hỗ trợ đã bị xóa.
  - Tệp `gopher.png` đã được sửa đổi. Khi kiểm tra kỹ hơn, hai tệp
    giống hệt nhau ngoại trừ một dấu thời gian nhúng mà Ubuntu đã cập nhật.
    Có thể các script đóng gói của Ubuntu đã nén lại png với một công cụ
    ghi lại dấu thời gian ngay cả khi không thể cải thiện nén hiện có.
  - Các nhị phân `dist` và `distpack`, được dựng trong quá trình bootstrap nhưng
    không được đưa vào các archive chuẩn,
    đã được đưa vào gói Ubuntu.
  - Các script dựng Plan 9 (`*.rc`) đã bị xóa, mặc dù các script dựng Windows (`*.bat`) vẫn còn.
  - `mksyscall.pl` và bảy Perl script khác không được hiển thị đã có header thay đổi.

Đặc biệt lưu ý rằng chúng ta đã tái cấu trúc các nhị phân bộ công cụ bit-for-bit:
chúng không xuất hiện trong diff chút nào.
Tức là, chúng ta đã chứng minh rằng các nhị phân Go của Ubuntu tương ứng chính xác với
nguồn Go upstream.

Tốt hơn nữa, chúng ta đã chứng minh điều này mà không sử dụng bất kỳ phần mềm Ubuntu nào:
những lệnh này được chạy trên Mac, và [`unzstd`](https://github.com/rsc/tmp/blob/master/unzstd/)
và [`elfstrip`](https://github.com/rsc/tmp/blob/master/elfstrip/) là các chương trình Go ngắn.
Một kẻ tấn công tinh vi có thể chèn mã độc vào gói Ubuntu
bằng cách thay đổi các công cụ tạo gói.
Nếu họ làm vậy, việc tái tạo gói Go Ubuntu từ nguồn sạch bằng cách sử dụng
những công cụ độc hại đó vẫn sẽ tạo ra các bản sao giống hệt nhau bit-for-bit của
các gói độc hại.
Cuộc tấn công này sẽ vô hình với loại dựng lại đó,
giống như [cuộc tấn công trình biên dịch của Ken Thompson](https://dl.acm.org/doi/10.1145/358198.358210).
Việc xác minh các gói Ubuntu mà không sử dụng bất kỳ phần mềm Ubuntu nào là một
kiểm tra mạnh mẽ hơn nhiều.
Các bản dựng hoàn toàn có thể tái tạo của Go, không phụ thuộc vào các
chi tiết không có chủ đích như hệ điều hành host,
kiến trúc host và bộ công cụ C host, là những gì làm cho kiểm tra mạnh mẽ hơn này có thể.

(Như một ghi chú lịch sử, Ken Thompson đã nói với tôi một lần rằng cuộc tấn công của ông
thực sự đã bị phát hiện,
vì bản dựng trình biên dịch ngừng có thể tái tạo.
Nó có một lỗi: một hằng số chuỗi trong backdoor được thêm vào trình biên dịch
được xử lý không hoàn hảo và phát triển thêm một byte NUL mỗi lần trình biên dịch biên dịch chính nó.
Cuối cùng ai đó nhận thấy bản dựng không thể tái tạo và cố gắng tìm nguyên nhân bằng cách biên dịch ra assembly.
Backdoor của trình biên dịch không tái tạo chính nó vào đầu ra assembly chút nào,
vì vậy việc assemble đầu ra đó đã xóa backdoor.)

## Kết luận

Các bản dựng có thể tái tạo là một công cụ quan trọng để tăng cường chuỗi cung ứng mã nguồn mở.
Các framework như [SLSA](https://slsa.dev/) tập trung vào xuất xứ và một
chuỗi giám sát phần mềm có thể được sử dụng để thông báo quyết định về tin tưởng.
Các bản dựng có thể tái tạo bổ sung cho cách tiếp cận đó bằng cách cung cấp một cách để xác minh
rằng sự tin tưởng được đặt đúng chỗ.

Khả năng tái tạo hoàn toàn (khi các tệp nguồn là đầu vào liên quan duy nhất của bản dựng)
chỉ có thể với các chương trình tự dựng chính mình,
như bộ công cụ trình biên dịch.
Đây là một mục tiêu cao cả nhưng xứng đáng chính xác vì các bộ công cụ trình biên dịch tự lưu trữ
thông thường khó xác minh.
Khả năng tái tạo hoàn toàn của Go có nghĩa là,
giả sử những người đóng gói không sửa đổi mã nguồn,
mỗi lần đóng gói lại Go 1.21.0 cho Linux/x86-64 (thay thế bằng hệ thống yêu thích của bạn)
dưới bất kỳ hình thức nào sẽ phân phối chính xác cùng các nhị phân,
ngay cả khi tất cả đều dựng từ nguồn.
Chúng ta đã thấy rằng điều này không hoàn toàn đúng cho Ubuntu Linux,
nhưng khả năng tái tạo hoàn toàn vẫn cho phép chúng ta tái tạo đóng gói Ubuntu
bằng cách sử dụng một hệ thống rất khác, không phải Ubuntu.

Lý tưởng nhất là tất cả phần mềm mã nguồn mở được phân phối dưới dạng nhị phân sẽ có các bản dựng dễ tái tạo.
Trên thực tế, như chúng ta đã thấy trong bài viết này,
rất dễ để các đầu vào không có chủ đích rò rỉ vào các bản dựng.
Đối với các chương trình Go không cần `cgo`, một bản dựng có thể tái tạo đơn giản như
biên dịch với `CGO_ENABLED=0 go build -trimpath`.
Tắt `cgo` loại bỏ bộ công cụ C host như một đầu vào liên quan,
và `-trimpath` loại bỏ thư mục hiện tại.
Nếu chương trình của bạn cần `cgo`, bạn cần sắp xếp một phiên bản
bộ công cụ C host cụ thể trước khi chạy `go build`,
chẳng hạn như bằng cách chạy bản dựng trong một máy ảo hoặc image container cụ thể.

Vượt ra ngoài Go, dự án [Reproducible Builds](https://reproducible-builds.org/)
nhằm cải thiện khả năng tái tạo của tất cả phần mềm mã nguồn mở và là một điểm khởi đầu tốt
để biết thêm thông tin về cách làm cho các bản dựng phần mềm của riêng bạn có thể tái tạo.
