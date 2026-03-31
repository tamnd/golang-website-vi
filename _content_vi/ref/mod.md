<!--{
  "Template": true,
  "Title": "Tài liệu tham khảo về Go Modules"
}-->
<!-- TODO(golang.org/issue/33637): Write focused "guide" articles on specific
module topics and tasks. Link to those instead of the blog, which will probably
not be updated over time. -->

## Giới thiệu {#introduction}

Module là cách Go quản lý dependency.

Tài liệu này là hướng dẫn tham khảo chi tiết về hệ thống module của Go. Để bắt đầu tạo dự án Go, xem [Cách viết code Go](/doc/code.html). Để tìm hiểu về cách sử dụng module, chuyển đổi dự án sang module, và các chủ đề khác, xem loạt bài viết trên blog bắt đầu với [Sử dụng Go Modules](/blog/using-go-modules).

## Module, package và phiên bản {#modules-overview}

Một <dfn>module</dfn> là tập hợp các package được phát hành, quản lý phiên bản và phân phối cùng nhau. Module có thể được tải xuống trực tiếp từ kho lưu trữ quản lý phiên bản hoặc từ các máy chủ module proxy.

Một module được xác định bởi một [đường dẫn module](#glos-module-path), được khai báo trong [tệp `go.mod`](#go-mod-file), kèm theo thông tin về các dependency của module. <dfn>Thư mục gốc của module</dfn> là thư mục chứa tệp `go.mod`. <dfn>Module chính</dfn> là module chứa thư mục nơi lệnh `go` được gọi.

Mỗi <dfn>package</dfn> trong một module là tập hợp các tệp nguồn trong cùng một thư mục được biên dịch cùng nhau. <dfn>Đường dẫn package</dfn> là đường dẫn module ghép với thư mục con chứa package (tính từ thư mục gốc của module). Ví dụ, module `"golang.org/x/net"` chứa một package trong thư mục `"html"`. Đường dẫn của package đó là `"golang.org/x/net/html"`.

### Đường dẫn module {#module-path}

<dfn>Đường dẫn module</dfn> là tên chuẩn của một module, được khai báo bằng [chỉ thị `module`](#go-mod-file-module) trong [tệp `go.mod`](#glos-go-mod-file) của module. Đường dẫn module là tiền tố cho các đường dẫn package trong module đó.

Đường dẫn module nên mô tả cả chức năng của module lẫn nơi có thể tìm thấy nó. Thông thường, đường dẫn module bao gồm đường dẫn gốc của kho lưu trữ, một thư mục bên trong kho lưu trữ (thường để trống), và hậu tố phiên bản chính (chỉ dùng cho phiên bản chính từ 2 trở lên).

* <dfn>Đường dẫn gốc kho lưu trữ</dfn> là phần trong đường dẫn module tương ứng với thư mục gốc của kho lưu trữ quản lý phiên bản nơi module được phát triển. Hầu hết các module được định nghĩa ở thư mục gốc của kho lưu trữ, vì vậy đây thường là toàn bộ đường dẫn. Ví dụ, `golang.org/x/net` là đường dẫn gốc kho lưu trữ cho module cùng tên. Xem [Tìm kho lưu trữ theo đường dẫn module](#vcs-find) để biết cách lệnh `go` xác định vị trí kho lưu trữ qua các yêu cầu HTTP từ đường dẫn module.
* Nếu module không được định nghĩa ở thư mục gốc của kho lưu trữ, <dfn>thư mục con của module</dfn> là phần trong đường dẫn module đặt tên cho thư mục đó, không bao gồm hậu tố phiên bản chính. Đây cũng là tiền tố cho các thẻ phiên bản ngữ nghĩa. Ví dụ, module `golang.org/x/tools/gopls` nằm trong thư mục con `gopls` của kho lưu trữ có đường dẫn gốc `golang.org/x/tools`, nên module subdirectory của nó là `gopls`. Xem [Ánh xạ phiên bản đến commit](#vcs-version) và [Thư mục module trong kho lưu trữ](#vcs-dir).
* Nếu module được phát hành ở phiên bản chính từ 2 trở lên, đường dẫn module phải kết thúc bằng [hậu tố phiên bản chính](#major-version-suffixes) như `/v2`. Phần này có thể hoặc không phải là tên thư mục con. Ví dụ, module có đường dẫn `golang.org/x/repo/sub/v2` có thể nằm trong thư mục con `/sub` hoặc `/sub/v2` của kho lưu trữ `golang.org/x/repo`.

Nếu module có thể được các module khác phụ thuộc vào, các quy tắc này phải được tuân thủ để lệnh `go` có thể tìm và tải module xuống. Ngoài ra còn có một số [hạn chế từ vựng](#go-mod-file-ident) đối với các ký tự được phép trong đường dẫn module.

Một module sẽ không bao giờ được tải xuống làm dependency của module khác có thể sử dụng bất kỳ đường dẫn package hợp lệ nào làm đường dẫn module, nhưng cần tránh trùng với các đường dẫn có thể được sử dụng bởi các dependency của module hoặc thư viện chuẩn Go. Thư viện chuẩn Go sử dụng các đường dẫn package không chứa dấu chấm trong phần tử đường dẫn đầu tiên, và lệnh `go` không cố gắng phân giải những đường dẫn như vậy từ các máy chủ mạng. Các đường dẫn `example` và `test` được dành riêng cho người dùng: chúng sẽ không được dùng trong thư viện chuẩn và phù hợp để sử dụng trong các module độc lập, chẳng hạn những module được định nghĩa trong hướng dẫn, code ví dụ, hay được tạo ra và thao tác như một phần của bài kiểm thử.

### Phiên bản {#versions}

Một <dfn>phiên bản</dfn> xác định một ảnh chụp bất biến của module, có thể là một [bản phát hành](#glos-release-version) hoặc một [bản tiền phát hành](#glos-pre-release-version). Mỗi phiên bản bắt đầu bằng chữ `v`, tiếp theo là một phiên bản ngữ nghĩa. Xem [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html) để biết chi tiết về cách định dạng, diễn giải và so sánh phiên bản.

Tóm lại, một phiên bản ngữ nghĩa gồm ba số nguyên không âm (phiên bản chính, phụ và vá lỗi, từ trái sang phải) được phân cách bằng dấu chấm. Phiên bản vá lỗi có thể được theo sau bởi một chuỗi tiền phát hành tùy chọn bắt đầu bằng dấu gạch ngang. Chuỗi tiền phát hành hoặc phiên bản vá lỗi có thể được theo sau bởi một chuỗi siêu dữ liệu build bắt đầu bằng dấu cộng. Ví dụ, `v0.0.0`, `v1.12.134`, `v8.0.5-pre` và `v2.0.9+meta` là các phiên bản hợp lệ.

Mỗi phần của phiên bản cho biết liệu phiên bản đó có ổn định hay không và có tương thích với các phiên bản trước hay không.

* [Phiên bản chính](#glos-major-version) phải được tăng lên và phiên bản phụ cùng phiên bản vá lỗi phải được đặt về không sau khi có thay đổi không tương thích ngược với giao diện công khai hoặc chức năng được ghi lại của module, ví dụ sau khi một package bị xóa.
* [Phiên bản phụ](#glos-minor-version) phải được tăng lên và phiên bản vá lỗi phải được đặt về không sau khi có thay đổi tương thích ngược, ví dụ sau khi thêm một hàm mới.
* [Phiên bản vá lỗi](#glos-patch-version) phải được tăng lên sau khi có thay đổi không ảnh hưởng đến giao diện công khai của module, chẳng hạn như sửa lỗi hoặc tối ưu hóa.
* Hậu tố tiền phát hành chỉ ra rằng phiên bản là một [bản tiền phát hành](#glos-pre-release-version). Các phiên bản tiền phát hành được sắp xếp trước các phiên bản phát hành tương ứng. Ví dụ, `v1.2.3-pre` đứng trước `v1.2.3`.
* Hậu tố siêu dữ liệu build bị bỏ qua khi so sánh phiên bản. Lệnh go chấp nhận các phiên bản có siêu dữ liệu build và chuyển đổi chúng thành pseudo-version để duy trì thứ tự toàn phần giữa các phiên bản.
  * Hậu tố đặc biệt `+incompatible` biểu thị một phiên bản được phát hành trước khi chuyển sang module với phiên bản chính từ 2 trở lên (xem [Tương thích với kho lưu trữ không dùng module](#non-module-compat)).
  * Hậu tố đặc biệt `+dirty` được thêm vào thông tin phiên bản của một binary khi nó được build bằng Go toolchain 1.24 trở lên trong một kho lưu trữ VCS cục bộ hợp lệ có chứa các thay đổi chưa commit trong thư mục làm việc.

Một phiên bản được coi là không ổn định nếu phiên bản chính của nó là 0 hoặc nó có hậu tố tiền phát hành. Các phiên bản không ổn định không phải tuân theo các yêu cầu về tương thích. Ví dụ, `v0.2.0` có thể không tương thích với `v0.1.0`, và `v1.5.0-beta` có thể không tương thích với `v1.5.0`.

Go có thể truy cập các module trong hệ thống quản lý phiên bản bằng cách sử dụng thẻ, nhánh hoặc revision không tuân theo các quy ước này. Tuy nhiên, trong module chính, lệnh `go` sẽ tự động chuyển đổi tên revision không theo chuẩn này thành các phiên bản chuẩn. Lệnh `go` cũng sẽ xóa các hậu tố siêu dữ liệu build (ngoại trừ `+incompatible`) như một phần của quá trình này. Điều này có thể dẫn đến một [pseudo-version](#glos-pseudo-version), một phiên bản tiền phát hành mã hóa định danh revision (chẳng hạn như hash commit Git) và dấu thời gian từ hệ thống quản lý phiên bản. Ví dụ, lệnh `go get golang.org/x/net@daa7c041` sẽ chuyển đổi hash commit `daa7c041` thành pseudo-version `v0.0.0-20191109021931-daa7c04131f5`. Các phiên bản chuẩn là bắt buộc bên ngoài module chính, và lệnh `go` sẽ báo lỗi nếu một phiên bản không chuẩn như `master` xuất hiện trong tệp `go.mod`.

### Pseudo-version {#pseudo-versions}

Một <dfn>pseudo-version</dfn> là một [phiên bản](#glos-version) [tiền phát hành](#glos-pre-release-version) có định dạng đặc biệt, mã hóa thông tin về một revision cụ thể trong kho lưu trữ quản lý phiên bản. Ví dụ, `v0.0.0-20191109021931-daa7c04131f5` là một pseudo-version.

Pseudo-version có thể tham chiếu đến các revision không có [thẻ phiên bản ngữ nghĩa](#glos-semantic-version-tag) nào. Chúng có thể được dùng để kiểm thử các commit trước khi tạo thẻ phiên bản, ví dụ trên một nhánh phát triển.

Mỗi pseudo-version có ba phần:

* Tiền tố phiên bản cơ sở (`vX.0.0` hoặc `vX.Y.Z-0`), được lấy từ thẻ phiên bản ngữ nghĩa đứng trước revision hoặc là `vX.0.0` nếu không có thẻ như vậy.
* Dấu thời gian (`yyyymmddhhmmss`), là thời gian UTC khi revision được tạo. Trong Git, đây là thời gian commit, không phải thời gian tác giả.
* Định danh revision (`abcdefabcdef`), là tiền tố 12 ký tự của hash commit, hoặc trong Subversion là số revision được đệm bằng số không.

Mỗi pseudo-version có thể ở một trong ba dạng, tùy thuộc vào phiên bản cơ sở. Các dạng này đảm bảo rằng một pseudo-version so sánh cao hơn phiên bản cơ sở nhưng thấp hơn phiên bản có thẻ tiếp theo.

* `vX.0.0-yyyymmddhhmmss-abcdefabcdef` được dùng khi không có phiên bản cơ sở nào được biết đến. Giống như tất cả các phiên bản, phiên bản chính `X` phải khớp với [hậu tố phiên bản chính](#glos-major-version-suffix) của module.
* `vX.Y.Z-pre.0.yyyymmddhhmmss-abcdefabcdef` được dùng khi phiên bản cơ sở là một phiên bản tiền phát hành như `vX.Y.Z-pre`.
* `vX.Y.(Z+1)-0.yyyymmddhhmmss-abcdefabcdef` được dùng khi phiên bản cơ sở là một phiên bản phát hành như `vX.Y.Z`. Ví dụ, nếu phiên bản cơ sở là `v1.2.3`, một pseudo-version có thể là `v1.2.4-0.20191109021931-daa7c04131f5`.

Nhiều hơn một pseudo-version có thể tham chiếu đến cùng một commit bằng cách sử dụng các phiên bản cơ sở khác nhau. Điều này xảy ra tự nhiên khi một phiên bản thấp hơn được gắn thẻ sau khi một pseudo-version đã được ghi lại.

Các dạng này mang lại cho pseudo-version hai tính chất hữu ích:

* Pseudo-version có phiên bản cơ sở được biết sẽ được sắp xếp cao hơn các phiên bản đó nhưng thấp hơn các bản tiền phát hành khác của phiên bản sau.
* Pseudo-version có cùng tiền tố phiên bản cơ sở được sắp xếp theo thứ tự thời gian.

Lệnh `go` thực hiện một số kiểm tra để đảm bảo rằng tác giả module có quyền kiểm soát cách so sánh pseudo-version với các phiên bản khác và rằng pseudo-version tham chiếu đến các revision thực sự là một phần của lịch sử commit của module.

* Nếu một phiên bản cơ sở được chỉ định, phải có một thẻ phiên bản ngữ nghĩa tương ứng là tổ tiên của revision được mô tả bởi pseudo-version. Điều này ngăn các nhà phát triển bỏ qua [lựa chọn phiên bản tối thiểu](#glos-minimal-version-selection) bằng cách sử dụng một pseudo-version so sánh cao hơn tất cả các phiên bản có thẻ như `v1.999.999-99999999999999-daa7c04131f5`.
* Dấu thời gian phải khớp với dấu thời gian của revision. Điều này ngăn kẻ tấn công làm ngập [module proxy](#glos-module-proxy) với số lượng không giới hạn các pseudo-version giống nhau. Điều này cũng ngăn người dùng module thay đổi thứ tự tương đối của các phiên bản.
* Revision phải là tổ tiên của một trong các nhánh hoặc thẻ của kho lưu trữ module. Điều này ngăn kẻ tấn công tham chiếu đến các thay đổi hoặc pull request chưa được phê duyệt.

Pseudo-version không bao giờ cần phải gõ tay. Nhiều lệnh chấp nhận hash commit hoặc tên nhánh và sẽ tự động chuyển đổi chúng thành pseudo-version (hoặc phiên bản có thẻ nếu có). Ví dụ:

```
go get example.com/mod@master
go list -m -json example.com/mod@abcd1234
```

### Hậu tố phiên bản chính {#major-version-suffixes}

Bắt đầu từ phiên bản chính 2, đường dẫn module phải có <dfn>hậu tố phiên bản chính</dfn> như `/v2` khớp với phiên bản chính. Ví dụ, nếu một module có đường dẫn `example.com/mod` ở `v1.0.0`, nó phải có đường dẫn `example.com/mod/v2` ở phiên bản `v2.0.0`.

Hậu tố phiên bản chính thực thi [<dfn>quy tắc tương thích import</dfn>](https://research.swtch.com/vgo-import):

> Nếu một package cũ và một package mới có cùng đường dẫn import,
> package mới phải tương thích ngược với package cũ.

Theo định nghĩa, các package trong một phiên bản chính mới của module không tương thích ngược với các package tương ứng trong phiên bản chính trước. Do đó, bắt đầu từ `v2`, các package cần đường dẫn import mới. Điều này được thực hiện bằng cách thêm hậu tố phiên bản chính vào đường dẫn module. Vì đường dẫn module là tiền tố của đường dẫn import cho mỗi package trong module, việc thêm hậu tố phiên bản chính vào đường dẫn module cung cấp một đường dẫn import riêng biệt cho mỗi phiên bản không tương thích.

Hậu tố phiên bản chính không được phép ở phiên bản chính `v0` hoặc `v1`. Không cần thay đổi đường dẫn module giữa `v0` và `v1` vì các phiên bản `v0` không ổn định và không có đảm bảo tương thích. Ngoài ra, đối với hầu hết các module, `v1` tương thích ngược với phiên bản `v0` cuối cùng; phiên bản `v1` được coi là cam kết về tương thích, thay vì là dấu hiệu của các thay đổi không tương thích so với `v0`.

Là một trường hợp đặc biệt, các đường dẫn module bắt đầu bằng `gopkg.in/` phải luôn có hậu tố phiên bản chính, ngay cả ở `v0` và `v1`. Hậu tố phải bắt đầu bằng dấu chấm thay vì dấu gạch chéo (ví dụ, `gopkg.in/yaml.v2`).

Hậu tố phiên bản chính cho phép nhiều phiên bản chính của một module cùng tồn tại trong cùng một build. Điều này có thể cần thiết do [vấn đề dependency hình thoi](https://research.swtch.com/vgo-import#dependency_story). Thông thường, nếu một module được yêu cầu ở hai phiên bản khác nhau bởi các dependency bắc cầu, phiên bản cao hơn sẽ được sử dụng. Tuy nhiên, nếu hai phiên bản không tương thích, không phiên bản nào sẽ thỏa mãn tất cả các client. Vì các phiên bản không tương thích phải có số phiên bản chính khác nhau, chúng cũng phải có đường dẫn module khác nhau do hậu tố phiên bản chính. Điều này giải quyết xung đột: các module có hậu tố khác nhau được coi là các module riêng biệt, và các package của chúng, kể cả các package trong cùng thư mục con so với thư mục gốc module của chúng, đều là riêng biệt.

Nhiều dự án Go đã phát hành phiên bản ở `v2` hoặc cao hơn mà không sử dụng hậu tố phiên bản chính trước khi chuyển sang module (có lẽ trước khi module được giới thiệu). Các phiên bản này được chú thích bằng thẻ build `+incompatible` (ví dụ, `v2.0.0+incompatible`). Xem [Tương thích với kho lưu trữ không dùng module](#non-module-compat) để biết thêm thông tin.

### Phân giải package thành module {#resolve-pkg-mod}

Khi lệnh `go` tải một package bằng [đường dẫn package](#glos-package-path), nó cần xác định module nào cung cấp package đó.

Lệnh `go` bắt đầu bằng cách tìm kiếm [danh sách build](#glos-build-list) để tìm các module có đường dẫn là tiền tố của đường dẫn package. Ví dụ, nếu package `example.com/a/b` được import và module `example.com/a` nằm trong danh sách build, lệnh `go` sẽ kiểm tra xem `example.com/a` có chứa package trong thư mục `b` không. Ít nhất một tệp có phần mở rộng `.go` phải tồn tại trong một thư mục để thư mục đó được coi là một package. [Ràng buộc build](/pkg/go/build/#hdr-Build_Constraints) không được áp dụng cho mục đích này. Nếu chính xác một module trong danh sách build cung cấp package, module đó được sử dụng. Nếu không có module nào cung cấp package hoặc nếu hai hoặc nhiều module cung cấp package, lệnh `go` báo lỗi. Cờ `-mod=mod` hướng dẫn lệnh `go` cố gắng tìm các module mới cung cấp các package còn thiếu và cập nhật `go.mod` và `go.sum`. Các lệnh [`go get`](#go-get) và [`go mod tidy`](#go-mod-tidy) làm điều này tự động.

<!-- NOTE(golang.org/issue/27899): the go command reports an error when two
or more modules provide a package with the same path as above. In the future,
we may try to upgrade one (or all) of the colliding modules.
-->

Khi lệnh `go` tìm kiếm một module mới cho một đường dẫn package, nó kiểm tra biến môi trường `GOPROXY`, là danh sách các URL proxy được phân cách bằng dấu phẩy hoặc các từ khóa `direct` hoặc `off`. Một URL proxy chỉ ra rằng lệnh `go` nên liên hệ với một [module proxy](#glos-module-proxy) bằng [giao thức `GOPROXY`](#goproxy-protocol). `direct` chỉ ra rằng lệnh `go` nên [giao tiếp với hệ thống quản lý phiên bản](#vcs). `off` chỉ ra rằng không nên cố gắng giao tiếp. Các [biến môi trường](#environment-variables) `GOPRIVATE` và `GONOPROXY` cũng có thể được sử dụng để kiểm soát hành vi này.

Với mỗi mục trong danh sách `GOPROXY`, lệnh `go` yêu cầu phiên bản mới nhất của mỗi đường dẫn module có thể cung cấp package (tức là mỗi tiền tố của đường dẫn package). Với mỗi đường dẫn module được yêu cầu thành công, lệnh `go` sẽ tải module xuống ở phiên bản mới nhất và kiểm tra xem module có chứa package được yêu cầu không. Nếu một hoặc nhiều module chứa package được yêu cầu, module có đường dẫn dài nhất sẽ được sử dụng. Nếu một hoặc nhiều module được tìm thấy nhưng không có module nào chứa package được yêu cầu, một lỗi được báo cáo. Nếu không tìm thấy module nào, lệnh `go` thử mục tiếp theo trong danh sách `GOPROXY`. Nếu không còn mục nào, một lỗi được báo cáo.

Ví dụ, giả sử lệnh `go` đang tìm kiếm một module cung cấp package `golang.org/x/net/html` và `GOPROXY` được đặt thành `https://corp.example.com,https://proxy.golang.org`. Lệnh `go` có thể thực hiện các yêu cầu sau:

* Đến `https://corp.example.com/` (song song):
  * Yêu cầu phiên bản mới nhất của `golang.org/x/net/html`
  * Yêu cầu phiên bản mới nhất của `golang.org/x/net`
  * Yêu cầu phiên bản mới nhất của `golang.org/x`
  * Yêu cầu phiên bản mới nhất của `golang.org`
* Đến `https://proxy.golang.org/`, nếu tất cả các yêu cầu đến `https://corp.example.com/` đều thất bại với lỗi 404 hoặc 410:
  * Yêu cầu phiên bản mới nhất của `golang.org/x/net/html`
  * Yêu cầu phiên bản mới nhất của `golang.org/x/net`
  * Yêu cầu phiên bản mới nhất của `golang.org/x`
  * Yêu cầu phiên bản mới nhất của `golang.org`

Sau khi tìm được module phù hợp, lệnh `go` sẽ thêm một [yêu cầu](#go-mod-file-require) mới với đường dẫn và phiên bản của module mới vào tệp `go.mod` của module chính. Điều này đảm bảo rằng khi package tương tự được tải trong tương lai, cùng module đó sẽ được sử dụng ở cùng phiên bản. Nếu package được phân giải không được import bởi một package trong module chính, yêu cầu mới sẽ có comment `// indirect`.

## Tệp `go.mod` {#go-mod-file}

Một module được định nghĩa bởi một tệp văn bản được mã hóa UTF-8 có tên `go.mod` trong thư mục gốc của nó. Tệp `go.mod` theo hướng dòng. Mỗi dòng chứa một chỉ thị duy nhất, bao gồm một từ khóa theo sau là các đối số. Ví dụ:

```
module example.com/my/thing

go 1.23.0

require example.com/other/thing v1.0.2
require example.com/new/thing/v2 v2.3.4
exclude example.com/old/thing v1.2.3
replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
retract [v1.9.0, v1.9.5]
```

Từ khóa đứng đầu có thể được gom lại từ các dòng liền kề để tạo thành một khối, giống như trong lệnh import của Go.

```
require (
    example.com/new/thing/v2 v2.3.4
    example.com/old/thing v1.2.3
)
```

Tệp `go.mod` được thiết kế để con người đọc được và máy móc có thể ghi. Lệnh `go` cung cấp một số lệnh con để thay đổi tệp `go.mod`. Ví dụ, [`go get`](#go-get) có thể nâng cấp hoặc hạ cấp các dependency cụ thể. Các lệnh tải đồ thị module sẽ [tự động cập nhật](#go-mod-file-updates) `go.mod` khi cần. [`go mod edit`](#go-mod-edit) có thể thực hiện các chỉnh sửa cấp thấp. Gói [`golang.org/x/mod/modfile`](https://pkg.go.dev/golang.org/x/mod/modfile?tab=doc) có thể được các chương trình Go sử dụng để thực hiện các thay đổi tương tự theo cách lập trình.

Tệp `go.mod` là bắt buộc cho [module chính](#glos-main-module) và cho bất kỳ [module thay thế](#go-mod-file-replace) nào được chỉ định bằng đường dẫn tệp cục bộ. Tuy nhiên, một module không có tệp `go.mod` rõ ràng vẫn có thể được [yêu cầu](#go-mod-file-require) làm dependency hoặc được sử dụng làm module thay thế được chỉ định bằng đường dẫn module và phiên bản; xem [Tương thích với kho lưu trữ không dùng module](#non-module-compat).

### Các phần tử từ vựng {#go-mod-file-lexical}

Khi một tệp `go.mod` được phân tích cú pháp, nội dung của nó được chia thành một chuỗi các token. Có một số loại token: khoảng trắng, comment, dấu câu, từ khóa, định danh và chuỗi.

*Khoảng trắng* bao gồm dấu cách (U+0020), tab (U+0009), ký tự xuống dòng (U+000D) và ký tự xuống hàng (U+000A). Các ký tự khoảng trắng trừ ký tự xuống hàng không có tác dụng ngoài việc tách các token mà nếu không sẽ được ghép lại. Ký tự xuống hàng là các token quan trọng.

*Comment* bắt đầu bằng `//` và kéo dài đến cuối dòng. Comment kiểu `/* */` không được phép.

Các token *dấu câu* bao gồm `(`, `)` và `=>`.

*Từ khóa* phân biệt các loại chỉ thị khác nhau trong tệp `go.mod`. Các từ khóa được phép là `module`, `go`, `require`, `replace`, `exclude` và `retract`.

*Định danh* là các chuỗi ký tự không phải khoảng trắng, chẳng hạn như đường dẫn module hoặc phiên bản ngữ nghĩa.

*Chuỗi* là các chuỗi ký tự được trích dẫn. Có hai loại chuỗi: chuỗi được diễn giải bắt đầu và kết thúc bằng dấu nháy kép (`"`, U+0022) và chuỗi thô bắt đầu và kết thúc bằng dấu huyền (<code>&#x60;</code>, U+0060). Chuỗi được diễn giải có thể chứa các chuỗi thoát bao gồm dấu gạch chéo ngược (`\`, U+005C) theo sau là một ký tự khác. Dấu nháy kép được thoát (`\"`) không kết thúc một chuỗi được diễn giải. Giá trị không được trích dẫn của một chuỗi được diễn giải là chuỗi ký tự giữa các dấu nháy kép với mỗi chuỗi thoát được thay thế bằng ký tự theo sau dấu gạch chéo ngược (ví dụ, `\"` được thay thế bằng `"`, `\n` được thay thế bằng `n`). Ngược lại, giá trị không được trích dẫn của một chuỗi thô chỉ đơn giản là chuỗi ký tự giữa các dấu huyền; dấu gạch chéo ngược không có ý nghĩa đặc biệt trong chuỗi thô.

Định danh và chuỗi có thể thay thế cho nhau trong ngữ pháp `go.mod`.

### Đường dẫn module và phiên bản {#go-mod-file-ident}

Hầu hết các định danh và chuỗi trong tệp `go.mod` là đường dẫn module hoặc phiên bản.

Một đường dẫn module phải thỏa mãn các yêu cầu sau:

* Đường dẫn phải bao gồm một hoặc nhiều phần tử đường dẫn được phân cách bằng dấu gạch chéo (`/`, U+002F). Nó không được bắt đầu hoặc kết thúc bằng dấu gạch chéo.
* Mỗi phần tử đường dẫn là một chuỗi không rỗng bao gồm các chữ cái ASCII, chữ số ASCII và dấu câu ASCII hạn chế (`-`, `.`, `_` và `~`).
* Một phần tử đường dẫn không được bắt đầu hoặc kết thúc bằng dấu chấm (`.`, U+002E).
* Tiền tố phần tử đến dấu chấm đầu tiên không được là tên tệp dành riêng trên Windows, bất kể chữ hoa hay thường (`CON`, `com1`, `NuL`, v.v.).
* Tiền tố phần tử đến dấu chấm đầu tiên không được kết thúc bằng dấu ngã theo sau bởi một hoặc nhiều chữ số (như `EXAMPL~1.COM`).

Nếu đường dẫn module xuất hiện trong chỉ thị `require` và không bị thay thế, hoặc nếu đường dẫn module xuất hiện ở phía bên phải của chỉ thị `replace`, lệnh `go` có thể cần tải xuống các module với đường dẫn đó, và một số yêu cầu bổ sung phải được thỏa mãn.

* Phần tử đường dẫn đứng đầu (đến dấu gạch chéo đầu tiên, nếu có), theo quy ước là tên miền, chỉ được chứa các chữ cái ASCII thường, chữ số ASCII, dấu chấm (`.`, U+002E) và dấu gạch ngang (`-`, U+002D); nó phải chứa ít nhất một dấu chấm và không được bắt đầu bằng dấu gạch ngang.
* Đối với phần tử đường dẫn cuối có dạng `/vN` trong đó `N` trông như số (chữ số ASCII và dấu chấm), `N` không được bắt đầu bằng số không đứng đầu, không được là `/v1` và không được chứa bất kỳ dấu chấm nào.
  * Đối với các đường dẫn bắt đầu bằng `gopkg.in/`, yêu cầu này được thay thế bằng yêu cầu rằng đường dẫn phải tuân theo các quy ước của dịch vụ [gopkg.in](https://gopkg.in).

Các phiên bản trong tệp `go.mod` có thể là [chuẩn](#glos-canonical-version) hoặc không chuẩn.

Một phiên bản chuẩn bắt đầu bằng chữ `v`, theo sau là phiên bản ngữ nghĩa theo thông số kỹ thuật [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html). Xem [Phiên bản](#versions) để biết thêm thông tin.

Hầu hết các định danh và chuỗi khác có thể được sử dụng như phiên bản không chuẩn, mặc dù có một số hạn chế để tránh sự cố với hệ thống tệp, kho lưu trữ và [module proxy](#glos-module-proxy). Các phiên bản không chuẩn chỉ được phép trong tệp `go.mod` của module chính. Lệnh `go` sẽ cố gắng thay thế mỗi phiên bản không chuẩn bằng phiên bản chuẩn tương đương khi tự động [cập nhật](#go-mod-file-updates) tệp `go.mod`.

Ở những nơi đường dẫn module được liên kết với một phiên bản (như trong các chỉ thị `require`, `replace` và `exclude`), phần tử đường dẫn cuối phải nhất quán với phiên bản. Xem [Hậu tố phiên bản chính](#major-version-suffixes).

### Ngữ pháp {#go-mod-file-grammar}

Cú pháp `go.mod` được chỉ định dưới đây bằng Extended Backus-Naur Form (EBNF). Xem [phần Ký hiệu trong Đặc tả ngôn ngữ Go](/ref/spec#Notation) để biết chi tiết về cú pháp EBNF.

```
GoMod = { Directive } .
Directive = ModuleDirective |
            GoDirective |
            ToolDirective |
            IgnoreDirective |
            RequireDirective |
            ExcludeDirective |
            ReplaceDirective |
            RetractDirective .
```

Ký tự xuống hàng, định danh và chuỗi được ký hiệu lần lượt bằng `newline`, `ident` và `string`.

Đường dẫn module và phiên bản được ký hiệu bằng `ModulePath` và `Version`.

```
ModulePath = ident | string . /* see restrictions above */
Version = ident | string .    /* see restrictions above */
```

### Chỉ thị `module` {#go-mod-file-module}

Chỉ thị `module` định nghĩa [đường dẫn](#glos-module-path) của module chính. Tệp `go.mod` phải chứa chính xác một chỉ thị `module`.

```
ModuleDirective = "module" ( ModulePath | "(" newline ModulePath newline ")" ) newline .
```

Ví dụ:

```
module golang.org/x/net
```

#### Đánh dấu lỗi thời {#go-mod-file-module-deprecation}

Một module có thể được đánh dấu là lỗi thời trong một khối comment chứa chuỗi `Deprecated:` (phân biệt chữ hoa thường) ở đầu một đoạn văn. Thông báo lỗi thời bắt đầu sau dấu hai chấm và kéo dài đến cuối đoạn văn. Các comment có thể xuất hiện ngay trước chỉ thị `module` hoặc sau đó trên cùng một dòng.

Ví dụ:

```
// Deprecated: use example.com/mod/v2 instead.
module example.com/mod
```

Từ Go 1.17, [`go list -m -u`](#go-list-m) kiểm tra thông tin về tất cả các module lỗi thời trong [danh sách build](#glos-build-list). [`go get`](#go-get) kiểm tra các module lỗi thời cần thiết để build các package được đặt tên trên dòng lệnh.

Khi lệnh `go` lấy thông tin lỗi thời cho một module, nó tải tệp `go.mod` từ phiên bản khớp với [truy vấn phiên bản](#version-queries) `@latest` mà không xem xét [retraction](#go-mod-file-retract) hoặc [exclusion](#go-mod-file-exclude). Lệnh `go` tải danh sách [các phiên bản bị retract](#glos-retracted-version) từ cùng tệp `go.mod` đó.

Để đánh dấu một module là lỗi thời, tác giả có thể thêm comment `// Deprecated:` và gắn thẻ một bản phát hành mới. Tác giả có thể thay đổi hoặc xóa thông báo lỗi thời trong một bản phát hành cao hơn.

Việc đánh dấu lỗi thời áp dụng cho tất cả các phiên bản phụ của một module. Các phiên bản chính cao hơn `v2` được coi là các module riêng biệt cho mục đích này, vì [hậu tố phiên bản chính](#glos-major-version-suffix) của chúng cung cấp cho chúng các đường dẫn module riêng biệt.

Thông báo lỗi thời nhằm thông báo cho người dùng biết rằng module không còn được hỗ trợ và cung cấp hướng dẫn chuyển đổi, ví dụ, sang phiên bản chính mới nhất. Các phiên bản phụ và vá lỗi riêng lẻ không thể bị đánh dấu lỗi thời; [`retract`](#go-mod-file-retract) có thể phù hợp hơn cho trường hợp đó.

### Chỉ thị `go` {#go-mod-file-go}

Chỉ thị `go` chỉ ra rằng một module được viết dựa trên ngữ nghĩa của một phiên bản Go nhất định. Phiên bản phải là một [phiên bản Go hợp lệ](/doc/toolchain#version), chẳng hạn như `1.14`, `1.21rc1` hoặc `1.23.0`.

Chỉ thị `go` đặt phiên bản tối thiểu của Go cần thiết để sử dụng module này. Trước Go 1.21, chỉ thị này chỉ mang tính khuyến nghị; nay nó là yêu cầu bắt buộc: Go toolchain từ chối sử dụng các module khai báo phiên bản Go mới hơn.

Chỉ thị `go` là đầu vào để chọn Go toolchain nào sẽ chạy. Xem "[Go toolchains](/doc/toolchain)" để biết chi tiết.

Chỉ thị `go` ảnh hưởng đến việc sử dụng các tính năng ngôn ngữ mới:

* Đối với các package trong module, trình biên dịch từ chối việc sử dụng các tính năng ngôn ngữ được giới thiệu sau phiên bản được chỉ định bởi chỉ thị `go`. Ví dụ, nếu một module có chỉ thị `go 1.12`, các package của nó không được sử dụng literal số như `1_000_000`, vốn được giới thiệu trong Go 1.13.
* Nếu một phiên bản Go cũ hơn build một trong các package của module và gặp lỗi biên dịch, thông báo lỗi ghi chú rằng module được viết cho một phiên bản Go mới hơn. Ví dụ, giả sử một module có `go 1.13` và một package sử dụng literal số `1_000_000`. Nếu package đó được build với Go 1.12, trình biên dịch sẽ ghi chú rằng code được viết cho Go 1.13.

Chỉ thị `go` cũng ảnh hưởng đến hành vi của lệnh `go`:

* Ở `go 1.14` trở lên, [vendoring](#vendoring) tự động có thể được bật. Nếu tệp `vendor/modules.txt` tồn tại và nhất quán với `go.mod`, không cần sử dụng cờ `-mod=vendor` một cách rõ ràng.
* Ở `go 1.16` trở lên, mẫu package `all` chỉ khớp với các package được import bắc cầu bởi các package và bài kiểm thử trong [module chính](#glos-main-module). Đây là tập hợp package tương tự được giữ lại bởi [`go mod vendor`](#go-mod-vendor) kể từ khi module được giới thiệu. Ở các phiên bản thấp hơn, `all` cũng bao gồm các bài kiểm thử của các package được import bởi các package trong module chính, các bài kiểm thử của những package đó, v.v.
* Ở `go 1.17` trở lên:
   * Tệp `go.mod` bao gồm một [chỉ thị `require`](#go-mod-file-require) rõ ràng cho mỗi module cung cấp bất kỳ package nào được import bắc cầu bởi một package hoặc bài kiểm thử trong module chính. (Ở `go 1.16` trở xuống, [dependency gián tiếp](#glos-direct-dependency) chỉ được bao gồm nếu [lựa chọn phiên bản tối thiểu](#minimal-version-selection) nếu không sẽ chọn một phiên bản khác.) Thông tin bổ sung này cho phép [cắt tỉa đồ thị module](#graph-pruning) và [tải module lười biếng](#lazy-loading).
   * Vì có thể có nhiều dependency `// indirect` hơn trong các phiên bản `go` trước, các dependency gián tiếp được ghi lại trong một khối riêng trong tệp `go.mod`.
   * `go mod vendor` bỏ qua các tệp `go.mod` và `go.sum` cho các dependency được vendor. (Điều này cho phép các lệnh `go` trong các thư mục con của `vendor` xác định đúng module chính.)
   * `go mod vendor` ghi lại phiên bản `go` từ tệp `go.mod` của mỗi dependency trong `vendor/modules.txt`.
* Ở `go 1.21` trở lên:
   * Dòng `go` khai báo phiên bản tối thiểu bắt buộc của Go để sử dụng với module này.
   * Dòng `go` phải lớn hơn hoặc bằng dòng `go` của tất cả các dependency.
   * Lệnh `go` không còn cố gắng duy trì tương thích với phiên bản Go cũ hơn trước đó.
   * Lệnh `go` cẩn thận hơn trong việc giữ checksum của các tệp `go.mod` trong tệp `go.sum`.
<!-- If you update this list, also update /doc/modules/gomod-ref#go-notes. -->

Tệp `go.mod` có thể chứa nhiều nhất một chỉ thị `go`. Hầu hết các lệnh sẽ thêm một chỉ thị `go` với phiên bản Go hiện tại nếu chưa có.

Nếu chỉ thị `go` bị thiếu, `go 1.16` được giả định.

```
GoDirective = "go" GoVersion newline .
GoVersion = string | ident .  /* valid release version; see above */
```

Ví dụ:

```
go 1.23.0
```

### Chỉ thị `toolchain` {#go-mod-file-toolchain}

Chỉ thị `toolchain` khai báo một Go toolchain được đề xuất để sử dụng với module. Phiên bản của Go toolchain được đề xuất không thể nhỏ hơn phiên bản Go bắt buộc được khai báo trong chỉ thị `go`. Chỉ thị `toolchain` chỉ có hiệu lực khi module là module chính và phiên bản toolchain mặc định nhỏ hơn phiên bản toolchain được đề xuất.

Để đảm bảo khả năng tái tạo, lệnh `go` ghi tên toolchain của chính nó vào một dòng `toolchain` bất cứ lúc nào nó đang cập nhật phiên bản `go` trong tệp `go.mod` (thường trong quá trình `go get`).

Để biết chi tiết, xem "[Go toolchains](/doc/toolchain)".

```
ToolchainDirective = "toolchain" ToolchainName newline .
ToolchainName = string | ident .  /* valid toolchain name; see "Go toolchains" */
```

Ví dụ:

```
toolchain go1.21.0
```

### Chỉ thị `godebug` {#go-mod-file-godebug}

Chỉ thị `godebug` khai báo một [cài đặt GODEBUG](/doc/godebug) duy nhất để áp dụng khi module này là module chính. Có thể có nhiều hơn một dòng như vậy và chúng có thể được gom lại. Sẽ là lỗi nếu module chính đặt tên một khóa GODEBUG không tồn tại. Hiệu lực của `godebug key=value` như thể mọi gói chính được biên dịch đều chứa một tệp nguồn liệt kê `//go:debug key=value`.

```
GodebugDirective = "godebug" ( GodebugSpec | "(" newline { GodebugSpec } ")" newline ) .
GodebugSpec = GodebugKey "=" GodebugValue newline.
GodebugKey = GodebugChar { GodebugChar }.
GodebugValue = GodebugChar { GodebugChar }.
GodebugChar = any non-space character except , " ` ' (comma and quotes).
```

Ví dụ:

```
godebug default=go1.21
godebug (
	panicnil=1
	asynctimerchan=0
)
```

### Chỉ thị `require` {#go-mod-file-require}

Chỉ thị `require` khai báo phiên bản tối thiểu bắt buộc của một dependency module nhất định. Với mỗi phiên bản module bắt buộc, lệnh `go` tải tệp `go.mod` cho phiên bản đó và kết hợp các yêu cầu từ tệp đó. Sau khi tất cả các yêu cầu được tải, lệnh `go` phân giải chúng bằng cách sử dụng [lựa chọn phiên bản tối thiểu (MVS)](#minimal-version-selection) để tạo ra [danh sách build](#glos-build-list).

Lệnh `go` tự động thêm comment `// indirect` cho một số yêu cầu. Comment `// indirect` chỉ ra rằng không có package nào từ module bắt buộc được import trực tiếp bởi bất kỳ package nào trong [module chính](#glos-main-module).

Nếu [chỉ thị `go`](#go-mod-file-go) chỉ định `go 1.16` hoặc thấp hơn, lệnh `go` thêm yêu cầu gián tiếp khi phiên bản được chọn của một module cao hơn những gì đã được ngụ ý (bắc cầu) bởi các dependency khác của module chính. Điều đó có thể xảy ra do nâng cấp rõ ràng (`go get -u ./...`), xóa một số dependency khác trước đó đã áp đặt yêu cầu (`go mod tidy`), hoặc một dependency import một package mà không có yêu cầu tương ứng trong tệp `go.mod` của chính nó (chẳng hạn như một dependency hoàn toàn thiếu tệp `go.mod`).

Ở `go 1.17` trở lên, lệnh `go` thêm yêu cầu gián tiếp cho mỗi module cung cấp bất kỳ package nào được import (ngay cả [gián tiếp](#glos-indirect-dependency)) bởi một package hoặc bài kiểm thử trong module chính hoặc được truyền làm đối số cho `go get`. Các yêu cầu toàn diện hơn này cho phép [cắt tỉa đồ thị module](#graph-pruning) và [tải module lười biếng](#lazy-loading).

```
RequireDirective = "require" ( RequireSpec | "(" newline { RequireSpec } ")" newline ) .
RequireSpec = ModulePath Version newline .
```

Ví dụ:

```
require golang.org/x/net v1.2.3

require (
    golang.org/x/crypto v1.4.5 // indirect
    golang.org/x/text v1.6.7
)
```

### Chỉ thị `tool` {#go-mod-file-tool}

Kể từ Go 1.24, chỉ thị `tool` thêm một package làm dependency của module hiện tại. Nó cũng làm cho package đó có thể chạy bằng `go tool` khi thư mục làm việc hiện tại nằm trong module này, hoặc trong một workspace chứa module này.

Nếu package tool không nằm trong module hiện tại, phải có một chỉ thị `require` chỉ định phiên bản của tool cần sử dụng.

Mẫu meta `tool` phân giải thành danh sách các tool được định nghĩa trong `go.mod` của module hiện tại, hoặc trong chế độ workspace thì là tập hợp của tất cả các tool được định nghĩa trong tất cả các module trong workspace.

```
ToolDirective = "tool" ( ToolSpec | "(" newline { ToolSpec } ")" newline ) .
ToolSpec = ModulePath newline .
```

Ví dụ:

```
tool golang.org/x/tools/cmd/stringer

tool (
    example.com/module/cmd/a
    example.com/module/cmd/b
)
```

### Chỉ thị `ignore` {#go-mod-file-ignore}

Chỉ thị `ignore` khiến lệnh go bỏ qua các đường dẫn thư mục được phân cách bằng dấu gạch chéo, và bất kỳ tệp hoặc thư mục nào được chứa đệ quy trong chúng, khi khớp với các mẫu package.

Nếu đường dẫn bắt đầu bằng `./`, đường dẫn được diễn giải tương đối với thư mục gốc của module, và thư mục đó cùng bất kỳ thư mục hoặc tệp nào được chứa đệ quy trong nó sẽ bị bỏ qua khi khớp với các mẫu package.

Nếu không, bất kỳ thư mục nào có đường dẫn ở bất kỳ độ sâu nào trong module, và bất kỳ thư mục hoặc tệp nào được chứa đệ quy trong chúng sẽ bị bỏ qua.

```
IgnoreDirective = "ignore" ( IgnoreSpec | "(" newline { IgnoreSpec } ")" newline ) .
IgnoreSpec = RelativeFilePath newline .
RelativeFilePath = /* slash-separated relative file path */ .
```

Ví dụ
```
ignore ./node_modules

ignore (
    static
    content/html
    ./third_party/javascript
)
```

### Chỉ thị `exclude` {#go-mod-file-exclude}

Chỉ thị `exclude` ngăn một phiên bản module được tải bởi lệnh `go`.

Từ Go 1.16, nếu một phiên bản được tham chiếu bởi chỉ thị `require` trong bất kỳ tệp `go.mod` nào bị loại trừ bởi chỉ thị `exclude` trong tệp `go.mod` của module chính, yêu cầu đó bị bỏ qua. Điều này có thể khiến các lệnh như [`go get`](#go-get) và [`go mod tidy`](#go-mod-tidy) thêm các yêu cầu mới về các phiên bản cao hơn vào `go.mod`, với comment `// indirect` nếu thích hợp.

Trước Go 1.16, nếu một phiên bản bị loại trừ được tham chiếu bởi chỉ thị `require`, lệnh `go` liệt kê các phiên bản có sẵn cho module (như được hiển thị với [`go list -m -versions`](#go-list-m)) và tải phiên bản cao hơn không bị loại trừ tiếp theo thay thế. Điều này có thể dẫn đến lựa chọn phiên bản không xác định, vì phiên bản cao hơn tiếp theo có thể thay đổi theo thời gian. Cả phiên bản phát hành và tiền phát hành đều được xem xét cho mục đích này, nhưng pseudo-version thì không. Nếu
không có phiên bản nào cao hơn, lệnh `go` sẽ báo lỗi.

Chỉ thị `exclude` chỉ áp dụng trong tệp `go.mod` của module chính và bị bỏ qua
trong các module khác. Xem [Lựa chọn phiên bản tối giản](#minimal-version-selection)
để biết thêm chi tiết.

```
ExcludeDirective = "exclude" ( ExcludeSpec | "(" newline { ExcludeSpec } ")" newline ) .
ExcludeSpec = ModulePath Version newline .
```

Ví dụ:

```
exclude golang.org/x/net v1.2.3

exclude (
    golang.org/x/crypto v1.4.5
    golang.org/x/text v1.6.7
)
```

### Chỉ thị `replace` {#go-mod-file-replace}

Chỉ thị `replace` thay thế nội dung của một phiên bản cụ thể của một module,
hoặc tất cả các phiên bản của module đó, bằng nội dung tìm thấy ở nơi khác. Phần
thay thế có thể được chỉ định bằng một đường dẫn module và phiên bản khác, hoặc
bằng một đường dẫn tệp phụ thuộc vào nền tảng.

Nếu có một phiên bản ở phía trái của mũi tên (`=>`), chỉ phiên bản cụ thể đó
của module mới bị thay thế; các phiên bản khác vẫn được truy cập bình thường.
Nếu phiên bản bên trái bị bỏ qua, tất cả các phiên bản của module đều bị thay thế.

Nếu đường dẫn ở phía phải của mũi tên là đường dẫn tuyệt đối hoặc tương đối
(bắt đầu bằng `./` hoặc `../`), nó được hiểu là đường dẫn tệp cục bộ đến thư
mục gốc của module thay thế, và thư mục đó phải chứa tệp `go.mod`. Phiên bản
thay thế phải bị bỏ qua trong trường hợp này.

Nếu đường dẫn ở phía phải không phải là đường dẫn cục bộ, nó phải là một đường
dẫn module hợp lệ. Trong trường hợp này, phiên bản là bắt buộc. Phiên bản module
đó không được xuất hiện đồng thời trong danh sách build.

Bất kể phần thay thế được chỉ định bằng đường dẫn cục bộ hay đường dẫn module,
nếu module thay thế có tệp `go.mod`, chỉ thị `module` trong đó phải khớp với
đường dẫn module mà nó thay thế.

Chỉ thị `replace` chỉ áp dụng trong tệp `go.mod` của module chính và bị bỏ qua
trong các module khác. Xem [Lựa chọn phiên bản tối giản](#minimal-version-selection)
để biết thêm chi tiết.

Nếu có nhiều module chính, tệp `go.mod` của tất cả các module chính đều được áp
dụng. Các chỉ thị `replace` xung đột giữa các module chính không được phép và
phải được xóa hoặc ghi đè trong một [replace trong tệp `go.work`](#go-work-file-replace).

Lưu ý rằng chỉ thị `replace` đơn thuần không thêm module vào
[đồ thị module](#glos-module-graph). Cũng cần có một [chỉ thị `require`](#go-mod-file-require)
tham chiếu đến phiên bản module bị thay thế, trong tệp `go.mod` của module chính
hoặc tệp `go.mod` của một dependency. Chỉ thị `replace` sẽ không có tác dụng
nếu phiên bản module ở phía trái không được yêu cầu.

```
ReplaceDirective = "replace" ( ReplaceSpec | "(" newline { ReplaceSpec } ")" newline ) .
ReplaceSpec = ModulePath [ Version ] "=>" FilePath newline
            | ModulePath [ Version ] "=>" ModulePath Version newline .
FilePath = /* đường dẫn tệp tương đối hoặc tuyệt đối phụ thuộc nền tảng */
```

Ví dụ:

```
replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5

replace (
    golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
    golang.org/x/net => example.com/fork/net v1.4.5
    golang.org/x/net v1.2.3 => ./fork/net
    golang.org/x/net => ./fork/net
)
```

### Chỉ thị `retract` {#go-mod-file-retract}

Chỉ thị `retract` cho biết rằng một phiên bản hoặc một dải phiên bản của module
được định nghĩa bởi `go.mod` không nên được phụ thuộc vào. Chỉ thị `retract`
hữu ích khi một phiên bản được phát hành sớm hoặc phát hiện một vấn đề nghiêm
trọng sau khi phiên bản đó đã được phát hành. Các phiên bản bị thu hồi vẫn cần
có sẵn trong kho lưu trữ quản lý phiên bản và trên [module proxy](#glos-module-proxy)
để đảm bảo rằng các build phụ thuộc vào chúng không bị hỏng. Từ *retract* được
mượn từ tài liệu học thuật: một bài nghiên cứu bị thu hồi vẫn còn có sẵn, nhưng
nó có vấn đề và không nên là cơ sở cho công việc trong tương lai.

Khi một phiên bản module bị thu hồi, người dùng sẽ không tự động nâng cấp lên
nó bằng [`go get`](#go-get), [`go mod tidy`](#go-mod-tidy) hoặc các lệnh khác.
Các build phụ thuộc vào phiên bản bị thu hồi sẽ tiếp tục hoạt động, nhưng người
dùng sẽ được thông báo về việc thu hồi khi họ kiểm tra cập nhật bằng [`go list
-m -u`](#go-list-m) hoặc cập nhật một module liên quan bằng [`go get`](#go-get).

Để thu hồi một phiên bản, tác giả module nên thêm chỉ thị `retract` vào `go.mod`,
sau đó phát hành một phiên bản mới chứa chỉ thị đó. Phiên bản mới phải cao hơn
các phiên bản phát hành hoặc tiền phát hành khác; tức là [truy vấn phiên bản](#version-queries)
`@latest` phải phân giải đến phiên bản mới trước khi xem xét các lệnh thu hồi.
Lệnh `go` tải và áp dụng các lệnh thu hồi từ phiên bản được hiển thị bởi
`go list -m -retracted $modpath@latest` (trong đó `$modpath` là đường dẫn module).

Các phiên bản bị thu hồi bị ẩn khỏi danh sách phiên bản được in bởi [`go list -m
-versions`](#go-list-m) trừ khi dùng cờ `-retracted`. Các phiên bản bị thu hồi
bị loại trừ khi phân giải các truy vấn phiên bản như `@>=v1.2.3` hoặc `@latest`.

Một phiên bản chứa các lệnh thu hồi có thể tự thu hồi chính nó. Nếu phiên bản
phát hành hoặc tiền phát hành cao nhất của một module tự thu hồi, truy vấn
`@latest` sẽ phân giải đến phiên bản thấp hơn sau khi các phiên bản bị thu hồi
được loại trừ.

Ví dụ, hãy xem xét trường hợp tác giả của module `example.com/m` phát hành phiên
bản `v1.0.0` một cách nhầm lẫn. Để ngăn người dùng nâng cấp lên `v1.0.0`, tác
giả có thể thêm hai chỉ thị `retract` vào `go.mod`, sau đó gắn thẻ `v1.0.1` với
các lệnh thu hồi.

```
retract (
    v1.0.0 // Published accidentally.
    v1.0.1 // Contains retractions only.
)
```

Khi người dùng chạy `go get example.com/m@latest`, lệnh `go` đọc các lệnh thu
hồi từ `v1.0.1`, lúc này là phiên bản cao nhất. Cả `v1.0.0` và `v1.0.1` đều bị
thu hồi, vì vậy lệnh `go` sẽ nâng cấp (hoặc hạ cấp!) lên phiên bản cao nhất
tiếp theo, có thể là `v0.9.5`.

Chỉ thị `retract` có thể được viết với một phiên bản đơn (như `v1.0.0`) hoặc
với một khoảng đóng của các phiên bản có giới hạn trên và dưới, được phân cách
bởi `[` và `]` (như `[v1.1.0, v1.2.0]`). Một phiên bản đơn tương đương với một
khoảng trong đó giới hạn trên và dưới là như nhau. Giống như các chỉ thị khác,
nhiều chỉ thị `retract` có thể được nhóm lại trong một khối được phân cách bởi
`(` ở cuối dòng và `)` trên dòng riêng của nó.

Mỗi chỉ thị `retract` nên có một bình luận giải thích lý do thu hồi, dù điều
này không bắt buộc. Lệnh `go` có thể hiển thị các bình luận lý do trong cảnh
báo về các phiên bản bị thu hồi và trong đầu ra của `go list`. Bình luận lý do
có thể được viết ngay trên chỉ thị `retract` (không có dòng trắng ở giữa) hoặc
sau đó trên cùng một dòng. Nếu một bình luận xuất hiện trên một khối, nó áp dụng
cho tất cả các chỉ thị `retract` trong khối đó mà không có bình luận riêng.
Bình luận lý do có thể kéo dài nhiều dòng.

```
RetractDirective = "retract" ( RetractSpec | "(" newline { RetractSpec } ")" newline ) .
RetractSpec = ( Version | "[" Version "," Version "]" ) newline .
```

Ví dụ:

* Thu hồi tất cả các phiên bản từ `v1.0.0` đến `v1.9.9`:

```
retract v1.0.0
retract [v1.0.0, v1.9.9]
retract (
    v1.0.0
    [v1.0.0, v1.9.9]
)
```

* Quay lại chưa được phiên bản hóa sau khi phát hành sớm phiên bản `v1.0.0`:

```
retract [v0.0.0, v1.0.1] // assuming v1.0.1 contains this retraction.
```

* Xóa sạch một module bao gồm tất cả các pseudo-version và phiên bản được gắn thẻ:

```
retract [v0.0.0-0, v0.15.2]  // assuming v0.15.2 contains this retraction.
```

Chỉ thị `retract` được thêm vào Go 1.16. Go 1.15 trở xuống sẽ báo lỗi nếu một
chỉ thị `retract` được viết trong tệp `go.mod` của [module chính](#glos-main-module)
và sẽ bỏ qua các chỉ thị `retract` trong tệp `go.mod` của các dependency.

### Cập nhật tự động {#go-mod-file-updates}

Hầu hết các lệnh báo lỗi nếu `go.mod` thiếu thông tin hoặc không phản ánh chính
xác thực tế. Các lệnh [`go get`](#go-get) và [`go mod tidy`](#go-mod-tidy) có
thể được dùng để sửa hầu hết các vấn đề này. Ngoài ra, cờ `-mod=mod` có thể được
dùng với hầu hết các lệnh có nhận thức về module (`go build`, `go test`, v.v.)
để hướng dẫn lệnh `go` tự động sửa các vấn đề trong `go.mod` và `go.sum`.

Ví dụ, hãy xem xét tệp `go.mod` này:

```
module example.com/M

go 1.23.0

require (
    example.com/A v1
    example.com/B v1.0.0
    example.com/C v1.0.0
    example.com/D v1.2.3
    example.com/E dev
)

exclude example.com/D v1.2.3
```

Bản cập nhật được kích hoạt bằng `-mod=mod` viết lại các định danh phiên bản
không theo dạng chuẩn thành dạng semver [chuẩn](#glos-canonical-version), do đó
`v1` của `example.com/A` trở thành `v1.0.0`, và `dev` của `example.com/E` trở
thành pseudo-version cho commit mới nhất trên nhánh `dev`, chẳng hạn
`v0.0.0-20180523231146-b3f5c0f6e5f1`.

Bản cập nhật sửa đổi các yêu cầu để tôn trọng các loại trừ, do đó yêu cầu trên
`example.com/D v1.2.3` bị loại trừ được cập nhật để sử dụng phiên bản có sẵn
tiếp theo của `example.com/D`, có thể là `v1.2.4` hoặc `v1.3.0`.

Bản cập nhật loại bỏ các yêu cầu thừa hoặc gây hiểu nhầm. Ví dụ, nếu
`example.com/A v1.0.0` tự nó yêu cầu `example.com/B v1.2.0` và `example.com/C
v1.0.0`, thì yêu cầu `example.com/B v1.0.0` của `go.mod` là gây hiểu nhầm (bị
thay thế bởi nhu cầu `v1.2.0` của `example.com/A`), và yêu cầu `example.com/C
v1.0.0` của nó là thừa (được ngụ ý bởi nhu cầu cùng phiên bản của `example.com/A`),
vì vậy cả hai sẽ bị xóa. Nếu module chính chứa các package trực tiếp import
các package từ `example.com/B` hoặc `example.com/C`, thì các yêu cầu sẽ được
giữ lại nhưng được cập nhật lên các phiên bản thực tế đang được sử dụng.

Cuối cùng, bản cập nhật định dạng lại `go.mod` theo định dạng chuẩn, để các thay
đổi cơ học trong tương lai sẽ tạo ra các diff tối thiểu. Lệnh `go` sẽ không cập
nhật `go.mod` nếu chỉ cần thay đổi định dạng.

Vì đồ thị module định nghĩa ý nghĩa của các câu lệnh import, bất kỳ lệnh nào
tải các package cũng sử dụng `go.mod` và do đó có thể cập nhật nó, bao gồm
`go build`, `go get`, `go install`, `go list`, `go test`, `go mod tidy`.

Trong Go 1.15 trở xuống, cờ `-mod=mod` được bật theo mặc định, vì vậy các bản
cập nhật được thực hiện tự động. Kể từ Go 1.16, lệnh `go` hoạt động như thể
`-mod=readonly` được đặt thay thế: nếu cần bất kỳ thay đổi nào đối với `go.mod`,
lệnh `go` báo lỗi và đề xuất cách sửa.

## Lựa chọn phiên bản tối giản (MVS) {#minimal-version-selection}

Go sử dụng thuật toán gọi là <dfn>Lựa chọn phiên bản tối giản (MVS)</dfn> để
chọn một tập hợp các phiên bản module để sử dụng khi build các package. MVS được
mô tả chi tiết trong [Minimal Version Selection](https://research.swtch.com/vgo-mvs)
bởi Russ Cox.

Về mặt khái niệm, MVS hoạt động trên một đồ thị có hướng của các module, được
chỉ định bằng [các tệp `go.mod`](#glos-go-mod-file). Mỗi đỉnh trong đồ thị đại
diện cho một phiên bản module. Mỗi cạnh đại diện cho phiên bản tối thiểu được
yêu cầu của một dependency, được chỉ định bằng chỉ thị
[`require`](#go-mod-file-require). Đồ thị có thể được sửa đổi bởi các chỉ thị
[`exclude`](#go-mod-file-exclude) và [`replace`](#go-mod-file-replace) trong
tệp `go.mod` của (các) module chính và bởi các chỉ thị
[`replace`](#go-work-file-replace) trong tệp `go.work`.

MVS tạo ra [danh sách build](#glos-build-list) làm đầu ra, danh sách các phiên
bản module được sử dụng cho một build.

MVS bắt đầu từ các module chính (các đỉnh đặc biệt trong đồ thị không có phiên
bản) và duyệt đồ thị, theo dõi phiên bản được yêu cầu cao nhất của mỗi module.
Khi kết thúc quá trình duyệt, các phiên bản được yêu cầu cao nhất tạo thành danh
sách build: chúng là các phiên bản tối thiểu thỏa mãn tất cả các yêu cầu.

Danh sách build có thể được kiểm tra bằng lệnh [`go list -m
all`](#go-list-m). Khác với các hệ thống quản lý dependency khác, danh sách build
không được lưu trong tệp "lock". MVS là xác định, và danh sách build không thay
đổi khi các phiên bản mới của các dependency được phát hành, vì vậy MVS được sử
dụng để tính toán nó ở đầu mỗi lệnh có nhận thức về module.

Hãy xem xét ví dụ trong sơ đồ bên dưới. Module chính yêu cầu module A ở phiên
bản 1.2 trở lên và module B ở phiên bản 1.2 trở lên. A 1.2 và B 1.2 lần lượt
yêu cầu C 1.3 và C 1.4. C 1.3 và C 1.4 đều yêu cầu D 1.2.

![Đồ thị phiên bản module với các phiên bản đã ghé thăm được tô sáng](/doc/mvs/buildlist.svg "Đồ thị danh sách build MVS")

MVS ghé thăm và tải tệp `go.mod` cho mỗi phiên bản module được tô sáng màu xanh.
Khi kết thúc quá trình duyệt đồ thị, MVS trả về danh sách build chứa các phiên
bản được in đậm: A 1.2, B 1.2, C 1.4 và D 1.2. Lưu ý rằng các phiên bản B và D
cao hơn có sẵn nhưng MVS không chọn chúng, vì không có gì yêu cầu chúng.

### Thay thế {#mvs-replace}

Nội dung của một module (bao gồm tệp `go.mod` của nó) có thể được thay thế bằng
một [chỉ thị `replace`](#go-mod-file-replace) trong tệp `go.mod` của module chính
hoặc tệp `go.work` của workspace. Chỉ thị `replace` có thể áp dụng cho một phiên
bản cụ thể của module hoặc cho tất cả các phiên bản của module.

Các thay thế thay đổi đồ thị module, vì một module thay thế có thể có các
dependency khác với các phiên bản bị thay thế.

Hãy xem xét ví dụ bên dưới, trong đó C 1.4 đã được thay thế bằng R. R phụ thuộc
vào D 1.3 thay vì D 1.2, vì vậy MVS trả về danh sách build chứa A 1.2, B 1.2,
C 1.4 (được thay thế bằng R) và D 1.3.

![Đồ thị phiên bản module với một phần thay thế](/doc/mvs/replace.svg "Thay thế MVS")

### Loại trừ {#mvs-exclude}

Một module cũng có thể bị loại trừ ở các phiên bản cụ thể bằng cách dùng
[chỉ thị `exclude`](#go-mod-file-exclude) trong tệp `go.mod` của module chính.

Các loại trừ cũng thay đổi đồ thị module. Khi một phiên bản bị loại trừ, nó được
xóa khỏi đồ thị module, và các yêu cầu về nó được chuyển hướng đến phiên bản cao
hơn tiếp theo.

Hãy xem xét ví dụ bên dưới. C 1.3 đã bị loại trừ. MVS sẽ hoạt động như thể A 1.2
yêu cầu C 1.4 (phiên bản cao hơn tiếp theo) thay vì C 1.3.

![Đồ thị phiên bản module với một loại trừ](/doc/mvs/exclude.svg "MVS loại trừ")

### Nâng cấp {#mvs-upgrade}

Lệnh [`go get`](#go-get) có thể được dùng để nâng cấp một tập hợp các module.
Để thực hiện nâng cấp, lệnh `go` thay đổi đồ thị module trước khi chạy MVS bằng
cách thêm các cạnh từ các phiên bản đã ghé thăm đến các phiên bản được nâng cấp.

Hãy xem xét ví dụ bên dưới. Module B có thể được nâng cấp từ 1.2 lên 1.3, C có
thể được nâng cấp từ 1.3 lên 1.4, và D có thể được nâng cấp từ 1.2 lên 1.3.

![Đồ thị phiên bản module với các nâng cấp](/doc/mvs/upgrade.svg "Nâng cấp MVS")

Các nâng cấp (và hạ cấp) có thể thêm hoặc xóa các dependency gián tiếp. Trong
trường hợp này, E 1.1 và F 1.1 xuất hiện trong danh sách build sau khi nâng cấp,
vì E 1.1 được yêu cầu bởi B 1.3.

Để lưu giữ các nâng cấp, lệnh `go` cập nhật các yêu cầu trong `go.mod`. Nó sẽ
thay đổi yêu cầu trên B lên phiên bản 1.3. Nó cũng sẽ thêm các yêu cầu trên
C 1.4 và D 1.3 với các bình luận `// indirect`, vì các phiên bản đó sẽ không được
chọn nếu không có.

### Hạ cấp {#mvs-downgrade}

Lệnh [`go get`](#go-get) cũng có thể được dùng để hạ cấp một tập hợp các module.
Để thực hiện hạ cấp, lệnh `go` thay đổi đồ thị module bằng cách xóa các phiên
bản cao hơn các phiên bản bị hạ cấp. Nó cũng xóa các phiên bản của các module
khác phụ thuộc vào các phiên bản đã bị xóa, vì chúng có thể không tương thích
với các phiên bản đã hạ cấp của các dependency. Nếu module chính yêu cầu một
phiên bản module bị xóa bởi việc hạ cấp, yêu cầu đó được thay đổi thành phiên
bản trước đó chưa bị xóa. Nếu không có phiên bản trước nào có sẵn, yêu cầu đó
sẽ bị loại bỏ.

Hãy xem xét ví dụ bên dưới. Giả sử phát hiện một vấn đề với C 1.4, vì vậy chúng
ta hạ cấp xuống C 1.3. C 1.4 bị xóa khỏi đồ thị module. B 1.2 cũng bị xóa, vì
nó yêu cầu C 1.4 trở lên. Yêu cầu của module chính trên B được thay đổi thành 1.1.

![Đồ thị phiên bản module với hạ cấp](/doc/mvs/downgrade.svg "Hạ cấp MVS")

[`go get`](#go-get) cũng có thể xóa hoàn toàn các dependency, bằng cách dùng
hậu tố `@none` sau một đối số. Điều này hoạt động tương tự như hạ cấp. Tất cả
các phiên bản của module được đặt tên sẽ bị xóa khỏi đồ thị module.

## Cắt tỉa đồ thị module {#graph-pruning}

Nếu module chính ở `go 1.17` trở lên, [đồ thị module](#glos-module-graph) được
sử dụng cho [lựa chọn phiên bản tối giản](#minimal-version-selection) chỉ bao
gồm các yêu cầu _trực tiếp_ cho mỗi dependency module chỉ định `go 1.17` trở
lên trong tệp `go.mod` của chính nó, trừ khi phiên bản module đó cũng được yêu
cầu (bắc cầu) bởi một _dependency khác_ ở `go 1.16` trở xuống. (Các dependency
_bắc cầu_ của các dependency `go 1.17` bị _cắt tỉa_ khỏi đồ thị module.)

Vì tệp `go.mod` theo `go 1.17` bao gồm một [chỉ thị
require](#go-mod-file-require) cho mỗi dependency cần thiết để build bất kỳ
package hoặc test nào trong module đó, đồ thị module đã được cắt tỉa bao gồm tất
cả các dependency cần thiết để `go build` hoặc `go test` các package trong bất
kỳ dependency nào được yêu cầu tường minh bởi [module chính](#glos-main-module).
Một module _không_ cần thiết để build bất kỳ package hoặc test nào trong một
module nhất định không thể ảnh hưởng đến hành vi thời gian chạy của các package
của nó, vì vậy các dependency bị cắt tỉa khỏi đồ thị module sẽ chỉ gây ra sự
can thiệp giữa các module không liên quan đến nhau.

Các module mà yêu cầu của chúng đã bị cắt tỉa vẫn xuất hiện trong đồ thị module
và vẫn được báo cáo bởi `go list -m all`: các [phiên bản được chọn](#glos-selected-version)
của chúng đã biết và được xác định rõ ràng, và các package có thể được tải từ
các module đó (ví dụ, như là các dependency bắc cầu của các test được tải từ
các module khác). Tuy nhiên, vì lệnh `go` không thể dễ dàng xác định các
dependency nào của các module này được thỏa mãn, các đối số cho `go build` và
`go test` không thể bao gồm các package từ các module mà yêu cầu của chúng đã
bị cắt tỉa. [`go get`](#go-get) đưa module chứa mỗi package được đặt tên lên
thành một dependency tường minh, cho phép `go build` hoặc `go test` được gọi
trên package đó.

Vì Go 1.16 và trước đó không hỗ trợ cắt tỉa đồ thị module, toàn bộ bao đóng
bắc cầu của các dependency, bao gồm các dependency `go 1.17` bắc cầu, vẫn được
bao gồm cho mỗi module chỉ định `go 1.16` trở xuống. (Ở `go 1.16` trở xuống,
tệp `go.mod` chỉ bao gồm [các dependency trực tiếp](#glos-direct-dependency),
vì vậy một đồ thị lớn hơn nhiều phải được tải để đảm bảo tất cả các dependency
gián tiếp được bao gồm.)

Tệp [`go.sum`](#go-sum-files) được ghi bởi [`go mod tidy`](#go-mod-tidy) cho
một module theo mặc định bao gồm các checksum cần thiết bởi phiên bản Go _thấp
hơn một bậc_ so với phiên bản được chỉ định trong [chỉ thị `go`](#go-mod-file-go)
của nó. Vì vậy một module `go 1.17` bao gồm các checksum cần thiết cho đồ thị
module đầy đủ được tải bởi Go 1.16, nhưng một module `go 1.18` sẽ chỉ bao gồm
các checksum cần thiết cho đồ thị module đã được cắt tỉa được tải bởi Go 1.17.
Cờ `-compat` có thể được dùng để ghi đè phiên bản mặc định (ví dụ, để cắt tỉa
tệp `go.sum` tích cực hơn trong một module `go 1.17`).

Xem [tài liệu thiết kế](https://go.googlesource.com/proposal/+/master/design/36460-lazy-module-loading.md)
để biết thêm chi tiết.

### Tải module lười biếng {#lazy-loading}

Các yêu cầu toàn diện hơn được thêm vào để cắt tỉa đồ thị module cũng cho phép
một tối ưu hóa khác khi làm việc trong một module. Nếu module chính ở `go 1.17`
trở lên, lệnh `go` tránh tải đồ thị module đầy đủ cho đến khi (và trừ khi) cần
thiết. Thay vào đó, nó chỉ tải tệp `go.mod` của module chính, sau đó cố gắng
tải các package cần build chỉ bằng cách dùng các yêu cầu đó. Nếu một package
cần được import (ví dụ, một dependency của một test cho một package bên ngoài
module chính) không được tìm thấy trong các yêu cầu đó, thì phần còn lại của
đồ thị module sẽ được tải theo nhu cầu.

Nếu tất cả các package được import có thể được tìm thấy mà không cần tải đồ thị
module, lệnh `go` sau đó chỉ tải các tệp `go.mod` cho _các module chứa_ các
package đó, và các yêu cầu của chúng được kiểm tra so với các yêu cầu của module
chính để đảm bảo chúng nhất quán cục bộ. (Sự không nhất quán có thể phát sinh
do các lần merge kiểm soát phiên bản, các chỉnh sửa thủ công và các thay đổi
trong các module đã được [thay thế](#go-mod-file-replace) bằng các đường dẫn hệ
thống tệp cục bộ.)

## Workspace {#workspaces}

Một <dfn>workspace</dfn> là một tập hợp các module trên đĩa được sử dụng làm
các module chính khi chạy [lựa chọn phiên bản tối giản (MVS)](#minimal-version-selection).

Một workspace có thể được khai báo trong một [tệp `go.work`](#go-work-file) chỉ
định các đường dẫn tương đối đến các thư mục module của mỗi module trong workspace.
Khi không có tệp `go.work` nào tồn tại, workspace bao gồm module duy nhất chứa
thư mục hiện tại.

Hầu hết các lệnh con `go` làm việc với các module đều hoạt động trên tập hợp các
module được xác định bởi workspace hiện tại. `go mod init`, `go mod why`,
`go mod edit`, `go mod tidy`, `go mod vendor` và `go get` luôn hoạt động trên
một module chính duy nhất.

Một lệnh xác định xem nó có đang trong ngữ cảnh workspace hay không bằng cách
trước tiên kiểm tra biến môi trường `GOWORK`. Nếu `GOWORK` được đặt thành `off`,
lệnh sẽ ở trong ngữ cảnh một module. Nếu nó rỗng hoặc không được cung cấp, lệnh
sẽ tìm kiếm trong thư mục làm việc hiện tại, rồi các thư mục cha liên tiếp, để
tìm tệp `go.work`. Nếu tìm thấy tệp, lệnh sẽ hoạt động trong workspace mà nó
định nghĩa; nếu không, workspace sẽ chỉ bao gồm module chứa thư mục làm việc.
Nếu `GOWORK` đặt tên cho đường dẫn đến một tệp hiện có kết thúc bằng .work,
chế độ workspace sẽ được bật. Bất kỳ giá trị nào khác đều là lỗi. Bạn có thể
dùng lệnh `go env GOWORK` để xác định tệp `go.work` nào mà lệnh `go` đang sử
dụng. `go env GOWORK` sẽ rỗng nếu lệnh `go` không ở chế độ workspace.

### Tệp `go.work` {#go-work-file}

Một workspace được định nghĩa bởi một tệp văn bản mã hóa UTF-8 có tên `go.work`.
Tệp `go.work` theo hướng dòng. Mỗi dòng chứa một chỉ thị duy nhất, bao gồm một
từ khóa theo sau là các đối số. Ví dụ:

```
go 1.23.0

use ./my/first/thing
use ./my/second/thing

replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
```

Như trong các tệp `go.mod`, một từ khóa đứng đầu có thể được tách ra khỏi các
dòng liền kề để tạo một khối.

```
use (
    ./my/first/thing
    ./my/second/thing
)
```

Lệnh `go` cung cấp một số lệnh con để thao tác với tệp `go.work`.
[`go work init`](#go-work-init) tạo tệp `go.work` mới.
[`go work use`](#go-work-use) thêm các thư mục module vào tệp `go.work`.
[`go work edit`](#go-work-edit) thực hiện các chỉnh sửa cấp thấp. Gói
[`golang.org/x/mod/modfile`](https://pkg.go.dev/golang.org/x/mod/modfile?tab=doc)
có thể được sử dụng bởi các chương trình Go để thực hiện các thay đổi tương tự
theo cách lập trình.

Lệnh go sẽ duy trì tệp `go.work.sum` theo dõi các hash được sử dụng bởi workspace
mà không có trong các tệp go.sum của các module workspace tập thể.

Thông thường không nên commit tệp go.work vào hệ thống kiểm soát phiên bản, vì
hai lý do:

* Tệp `go.work` được check in có thể ghi đè tệp `go.work` riêng của nhà phát
  triển từ một thư mục cha, gây nhầm lẫn khi các chỉ thị `use` của họ không áp
  dụng.
* Tệp `go.work` được check in có thể khiến hệ thống tích hợp liên tục (CI) chọn
  và do đó kiểm tra các phiên bản không đúng của các dependency của module. Các
  hệ thống CI thường không được phép sử dụng tệp `go.work` để chúng có thể kiểm
  tra hành vi của module như khi nó được các module khác yêu cầu, trong đó tệp
  `go.work` bên trong module không có tác dụng.

Dù vậy, có một số trường hợp việc commit tệp `go.work` có ý nghĩa. Ví dụ, khi
các module trong một kho lưu trữ được phát triển dành riêng cho nhau nhưng không
cùng với các module bên ngoài, có thể không có lý do gì để nhà phát triển muốn
dùng sự kết hợp khác của các module trong một workspace. Trong trường hợp đó,
tác giả module nên đảm bảo các module riêng lẻ được kiểm tra và phát hành đúng
cách.

### Các phần tử từ vựng {#go-work-file-lexical}

Các phần tử từ vựng trong tệp `go.work` được định nghĩa theo cách hoàn toàn
giống [như đối với tệp `go.mod`](#go-mod-file-lexical).

### Ngữ pháp {#go-work-file-grammar}

Cú pháp `go.work` được chỉ định bên dưới bằng Ký hiệu Backus-Naur Mở rộng (EBNF).
Xem [phần Notation trong Đặc tả Ngôn ngữ Go](/ref/spec#Notation) để biết chi tiết
về cú pháp EBNF.

```
GoWork = { Directive } .
Directive = GoDirective |
            ToolchainDirective |
            UseDirective |
            ReplaceDirective .
```

Các ký tự xuống dòng, định danh và chuỗi ký tự được ký hiệu bằng `newline`,
`ident` và `string`.

Các đường dẫn module và phiên bản được ký hiệu bằng `ModulePath` và `Version`.
Các đường dẫn module và phiên bản được chỉ định theo cách hoàn toàn giống [như
đối với tệp `go.mod`](#go-mod-file-lexical).

```
ModulePath = ident | string . /* see restrictions above */
Version = ident | string .    /* see restrictions above */
```

### Chỉ thị `go` {#go-work-file-go}

Chỉ thị `go` là bắt buộc trong một tệp `go.work` hợp lệ. Phiên bản phải là một
phiên bản phát hành Go hợp lệ: một số nguyên dương theo sau là dấu chấm và một
số nguyên không âm (ví dụ, `1.18`, `1.19`).

Chỉ thị `go` cho biết phiên bản toolchain Go mà tệp `go.work` được thiết kế để
hoạt động cùng. Nếu có thay đổi đối với định dạng tệp `go.work`, các phiên bản
toolchain trong tương lai sẽ diễn giải tệp theo phiên bản được chỉ định của nó.

Một tệp `go.work` có thể chứa nhiều nhất một chỉ thị `go`.

```
GoDirective = "go" GoVersion newline .
GoVersion = string | ident .  /* valid release version; see above */
```

Ví dụ:

```
go 1.23.0
```

### Chỉ thị `toolchain` {#go-work-file-toolchain}

Chỉ thị `toolchain` khai báo một toolchain Go được đề xuất để dùng trong một
workspace. Nó chỉ có tác dụng khi toolchain mặc định cũ hơn toolchain được đề
xuất.

Để biết chi tiết, xem "[Go toolchains](/doc/toolchain)".

```
ToolchainDirective = "toolchain" ToolchainName newline .
ToolchainName = string | ident .  /* valid toolchain name; see "Go toolchains" */
```

Ví dụ:

```
toolchain go1.21.0
```

### Chỉ thị `godebug` {#go-work-file-godebug}

Chỉ thị `godebug` khai báo một [cài đặt GODEBUG](/doc/godebug) đơn để áp dụng
khi làm việc trong workspace này. Cú pháp và tác dụng giống như
[chỉ thị `godebug` của tệp `go.mod`](#go-mod-file-godebug). Khi workspace đang
được sử dụng, các chỉ thị `godebug` trong các tệp `go.mod` bị bỏ qua.


### Chỉ thị `use` {#go-work-file-use}

Một lệnh `use` thêm một module trên đĩa vào tập hợp các module chính trong một
workspace. Đối số của nó là một đường dẫn tương đối đến thư mục chứa tệp `go.mod`
của module. Chỉ thị `use` không thêm các module chứa trong các thư mục con của
thư mục đối số của nó. Các module đó có thể được thêm bằng thư mục chứa tệp
`go.mod` của chúng trong các chỉ thị `use` riêng biệt.

```
UseDirective = "use" ( UseSpec | "(" newline { UseSpec } ")" newline ) .
UseSpec = FilePath newline .
FilePath = /* platform-specific relative or absolute file path */

```

Ví dụ:

```
use ./mymod  // example.com/mymod

use (
    ../othermod
    ./subdir/thirdmod
)
```

### Chỉ thị `replace` {#go-work-file-replace}

Tương tự như chỉ thị `replace` trong tệp `go.mod`, chỉ thị `replace` trong tệp
`go.work` thay thế nội dung của một phiên bản cụ thể của một module, hoặc tất cả
các phiên bản của module, bằng nội dung tìm thấy ở nơi khác. Một replace ký tự
đại diện trong `go.work` ghi đè một `replace` theo phiên bản cụ thể trong tệp
`go.mod`.

Các chỉ thị `replace` trong tệp `go.work` ghi đè bất kỳ replace nào của cùng
module hoặc phiên bản module trong các module workspace.

```
ReplaceDirective = "replace" ( ReplaceSpec | "(" newline { ReplaceSpec } ")" newline ) .
ReplaceSpec = ModulePath [ Version ] "=>" FilePath newline
            | ModulePath [ Version ] "=>" ModulePath Version newline .
FilePath = /* platform-specific relative or absolute file path */
```

Ví dụ:

```
replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5

replace (
    golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
    golang.org/x/net => example.com/fork/net v1.4.5
    golang.org/x/net v1.2.3 => ./fork/net
    golang.org/x/net => ./fork/net
)
```

## Tương thích với các kho lưu trữ không phải module {#non-module-compat}

Để đảm bảo quá trình chuyển đổi suôn sẻ từ `GOPATH` sang module, lệnh `go` có
thể tải xuống và build các package trong chế độ nhận thức về module từ các kho
lưu trữ chưa di chuyển sang module bằng cách thêm [tệp `go.mod`](#glos-go-mod-file).

Khi lệnh `go` tải xuống một module ở một phiên bản nhất định [trực tiếp](#vcs)
từ một kho lưu trữ, nó tra cứu một URL kho lưu trữ cho đường dẫn module, ánh xạ
phiên bản đến một bản sửa đổi trong kho lưu trữ, sau đó trích xuất một tệp lưu
trữ của kho lưu trữ tại bản sửa đổi đó. Nếu [đường dẫn module](#glos-module-path)
bằng với [đường dẫn gốc kho lưu trữ](#glos-repository-root-path), và thư mục
gốc kho lưu trữ không chứa tệp `go.mod`, lệnh `go` tổng hợp một tệp `go.mod`
trong bộ nhớ cache module chứa một [chỉ thị `module`](#go-mod-file-module) và
không có gì khác. Vì các tệp `go.mod` tổng hợp không chứa [chỉ thị `require`](#go-mod-file-require)
cho các dependency của chúng, các module khác phụ thuộc vào chúng có thể cần các
chỉ thị `require` bổ sung (với các bình luận `// indirect`) để đảm bảo mỗi
dependency được tải ở cùng phiên bản trong mỗi lần build.

Khi lệnh `go` tải xuống một module từ một [proxy](#communicating-with-proxies),
nó tải xuống tệp `go.mod` riêng biệt so với phần còn lại của nội dung module.
Proxy được kỳ vọng sẽ phục vụ một tệp `go.mod` tổng hợp nếu module ban đầu
không có tệp đó.

### Các phiên bản `+incompatible` {#incompatible-versions}

Một module được phát hành ở phiên bản chính 2 trở lên phải có [hậu tố phiên bản
chính](#major-version-suffixes) phù hợp trên đường dẫn module của nó. Ví dụ,
nếu một module được phát hành ở `v2.0.0`, đường dẫn của nó phải có hậu tố `/v2`.
Điều này cho phép lệnh `go` xử lý nhiều phiên bản chính của một dự án như các
module riêng biệt, ngay cả khi chúng được phát triển trong cùng một kho lưu trữ.

Yêu cầu hậu tố phiên bản chính được giới thiệu khi hỗ trợ module được thêm vào
lệnh `go`, và nhiều kho lưu trữ đã gắn thẻ các bản phát hành với phiên bản chính
`2` trở lên trước đó. Để duy trì tương thích với các kho lưu trữ này, lệnh `go`
thêm hậu tố `+incompatible` vào các phiên bản có phiên bản chính 2 trở lên mà
không có tệp `go.mod`. `+incompatible` cho biết rằng một phiên bản là một phần
của cùng module với các phiên bản có số phiên bản chính thấp hơn; do đó, lệnh
`go` có thể tự động nâng cấp lên các phiên bản `+incompatible` cao hơn mặc dù
điều đó có thể làm hỏng build.

Hãy xem xét yêu cầu ví dụ bên dưới:

```
require example.com/m v4.1.2+incompatible
```

Phiên bản `v4.1.2+incompatible` tham chiếu đến [thẻ phiên bản ngữ nghĩa](#glos-semantic-version-tag)
`v4.1.2` trong kho lưu trữ cung cấp module `example.com/m`. Module phải nằm
trong thư mục gốc kho lưu trữ (tức là [đường dẫn gốc kho lưu trữ](#glos-module-path)
cũng phải là `example.com/m`), và tệp `go.mod` không được có. Module có thể có
các phiên bản với số phiên bản chính thấp hơn như `v1.5.2`, và lệnh `go` có thể
tự động nâng cấp lên `v4.1.2+incompatible` từ các phiên bản đó (xem [lựa chọn
phiên bản tối giản (MVS)](#minimal-version-selection) để biết thông tin về cách
nâng cấp hoạt động).

Một kho lưu trữ di chuyển sang module sau khi phiên bản `v2.0.0` được gắn thẻ
thường nên phát hành một phiên bản chính mới. Trong ví dụ trên, tác giả nên tạo
một module có đường dẫn `example.com/m/v5` và nên phát hành phiên bản `v5.0.0`.
Tác giả cũng nên cập nhật các import của các package trong module để sử dụng
tiền tố `example.com/m/v5` thay vì `example.com/m`. Xem [Go Modules: v2
and Beyond](/blog/v2-go-modules) để có ví dụ chi tiết hơn.

Lưu ý rằng hậu tố `+incompatible` không nên xuất hiện trên một thẻ trong kho
lưu trữ; thẻ như `v4.1.2+incompatible` sẽ bị bỏ qua. Hậu tố chỉ xuất hiện
trong các phiên bản được sử dụng bởi lệnh `go`. Xem [Ánh xạ phiên bản đến
commit](#vcs-version) để biết chi tiết về sự khác biệt giữa phiên bản và thẻ.

Lưu ý thêm rằng hậu tố `+incompatible` có thể xuất hiện trên
[pseudo-version](#glos-pseudo-version). Ví dụ,
`v2.0.1-20200722182040-012345abcdef+incompatible` có thể là một pseudo-version hợp lệ.

### Tương thích module tối giản {#minimal-module-compatibility}

Một module được phát hành ở phiên bản chính 2 trở lên được yêu cầu phải có [hậu
tố phiên bản chính](#glos-major-version-suffix) trên [đường dẫn module](#glos-module-path)
của nó. Module có thể hoặc không được phát triển trong [thư mục con phiên bản
chính](#glos-major-version-subdirectory) trong kho lưu trữ của nó. Điều này có
ảnh hưởng đến các package import các package trong module khi build trong chế độ
`GOPATH`.

Thông thường trong chế độ `GOPATH`, một package được lưu trữ trong một thư mục
khớp với [đường dẫn gốc kho lưu trữ](#glos-repository-root-path) của nó kết hợp
với thư mục của nó trong kho lưu trữ. Ví dụ, một package trong kho lưu trữ có
đường dẫn gốc `example.com/repo` trong thư mục con `sub` sẽ được lưu trữ trong
`$GOPATH/src/example.com/repo/sub` và sẽ được import là `example.com/repo/sub`.

Đối với một module có hậu tố phiên bản chính, người ta có thể mong đợi tìm thấy
package `example.com/repo/v2/sub` trong thư mục `$GOPATH/src/example.com/repo/v2/sub`.
Điều này sẽ đòi hỏi module phải được phát triển trong thư mục con `v2` của kho
lưu trữ của nó. Lệnh `go` hỗ trợ điều này nhưng không yêu cầu (xem [Ánh xạ
phiên bản đến commit](#vcs-version)).

Nếu một module *không* được phát triển trong thư mục con phiên bản chính, thì
thư mục của nó trong `GOPATH` sẽ không chứa hậu tố phiên bản chính, và các
package của nó có thể được import mà không có hậu tố phiên bản chính. Trong ví
dụ trên, package sẽ được tìm thấy trong thư mục `$GOPATH/src/example.com/repo/sub`
và sẽ được import là `example.com/repo/sub`.

Điều này tạo ra vấn đề cho các package được thiết kế để build trong cả chế độ
module và chế độ `GOPATH`: chế độ module yêu cầu hậu tố, trong khi chế độ
`GOPATH` thì không.

Để sửa điều này, <dfn>tương thích module tối giản</dfn> được thêm vào Go 1.11
và được backport sang Go 1.9.7 và 1.10.3. Khi một đường dẫn import được phân
giải đến một thư mục trong chế độ `GOPATH`:

* Khi phân giải một import có dạng `$modpath/$vn/$dir` trong đó:
  * `$modpath` là một đường dẫn module hợp lệ,
  * `$vn` là một hậu tố phiên bản chính,
  * `$dir` là một thư mục con có thể rỗng,
* Nếu tất cả các điều kiện sau đây là đúng:
  * Package `$modpath/$vn/$dir` không có trong bất kỳ [thư mục
    `vendor`](#glos-vendor-directory) liên quan nào.
  * Tệp `go.mod` có mặt trong cùng thư mục với tệp đang import hoặc trong bất
    kỳ thư mục cha nào lên đến gốc `$GOPATH/src`,
  * Không có thư mục `$GOPATH[i]/src/$modpath/$vn/$suffix` nào tồn tại (cho bất
    kỳ gốc `$GOPATH[i]` nào),
  * Tệp `$GOPATH[d]/src/$modpath/go.mod` tồn tại (cho một số gốc `$GOPATH[d]`)
    và khai báo đường dẫn module là `$modpath/$vn`,
* Thì import của `$modpath/$vn/$dir` được phân giải đến thư mục
  `$GOPATH[d]/src/$modpath/$dir`.

Các quy tắc này cho phép các package đã di chuyển sang module import các package
khác đã di chuyển sang module khi được build trong chế độ `GOPATH` ngay cả khi
thư mục con phiên bản chính không được sử dụng.

## Các lệnh nhận thức về module {#mod-commands}

Hầu hết các lệnh `go` có thể chạy trong *Chế độ nhận thức về module* hoặc *Chế
độ `GOPATH`*. Trong chế độ nhận thức về module, lệnh `go` sử dụng tệp `go.mod`
để tìm các dependency có phiên bản, và nó thường tải các package từ [bộ nhớ
cache module](#glos-module-cache), tải xuống các module nếu chúng bị thiếu. Trong
chế độ `GOPATH`, lệnh `go` bỏ qua các module; nó tìm trong [thư mục
`vendor`](#glos-vendor-directory) và trong `GOPATH` để tìm các dependency.

Kể từ Go 1.16, chế độ nhận thức về module được bật theo mặc định, bất kể tệp
`go.mod` có tồn tại hay không. Trong các phiên bản thấp hơn, chế độ nhận thức
về module được bật khi tệp `go.mod` có mặt trong thư mục hiện tại hoặc bất kỳ
thư mục cha nào.

Chế độ nhận thức về module có thể được kiểm soát bằng biến môi trường
`GO111MODULE`, có thể được đặt thành `on`, `off` hoặc `auto`.

* Nếu `GO111MODULE=off`, lệnh `go` bỏ qua tệp `go.mod` và chạy trong chế độ
  `GOPATH`.
* Nếu `GO111MODULE=on` hoặc không được đặt, lệnh `go` chạy trong chế độ nhận
  thức về module, ngay cả khi không có tệp `go.mod` nào. Không phải tất cả các
  lệnh đều hoạt động khi không có tệp `go.mod`: xem [Các lệnh module bên ngoài
  một module](#commands-outside).
* Nếu `GO111MODULE=auto`, lệnh `go` chạy trong chế độ nhận thức về module nếu
  tệp `go.mod` có mặt trong thư mục hiện tại hoặc bất kỳ thư mục cha nào. Trong
  Go 1.15 trở xuống, đây là hành vi mặc định. Các lệnh con `go mod` và
  `go install` với một [truy vấn phiên bản](#version-queries) chạy trong chế độ
  nhận thức về module ngay cả khi không có tệp `go.mod` nào.

Trong chế độ nhận thức về module, `GOPATH` không còn định nghĩa ý nghĩa của các
import trong quá trình build, nhưng nó vẫn lưu trữ các dependency đã tải xuống
(trong `GOPATH/pkg/mod`; xem [Bộ nhớ cache module](#module-cache)) và các lệnh
đã cài đặt (trong `GOPATH/bin`, trừ khi
`GOBIN` được đặt).

### Các lệnh build {#build-commands}

Tất cả các lệnh tải thông tin về package đều nhận biết module. Bao gồm:

* `go build`
* `go fix`
* `go generate`
* `go install`
* `go list`
* `go run`
* `go test`
* `go vet`

Khi chạy ở chế độ nhận biết module, các lệnh này sử dụng file `go.mod` để giải thích
các đường dẫn import được liệt kê trên dòng lệnh hoặc được viết trong các file nguồn Go. Các
lệnh này chấp nhận các cờ sau, chung cho tất cả các lệnh module.

* Cờ `-mod` kiểm soát việc `go.mod` có thể được cập nhật tự động hay không và
  liệu thư mục `vendor` có được sử dụng không.
  * `-mod=mod` yêu cầu lệnh `go` bỏ qua thư mục vendor và
     [tự động cập nhật](#go-mod-file-updates) `go.mod`, ví dụ khi một
     package được import không được cung cấp bởi bất kỳ module nào đã biết.
  * `-mod=readonly` yêu cầu lệnh `go` bỏ qua thư mục `vendor` và
    báo lỗi nếu `go.mod` cần được cập nhật.
  * `-mod=vendor` yêu cầu lệnh `go` sử dụng thư mục `vendor`. Ở chế độ
    này, lệnh `go` sẽ không sử dụng mạng hoặc cache module.
  * Theo mặc định, nếu [phiên bản `go`](#go-mod-file-go) trong `go.mod` là `1.14` hoặc
    cao hơn và thư mục `vendor` tồn tại, lệnh `go` hoạt động như thể
    `-mod=vendor` được sử dụng. Ngược lại, lệnh `go` hoạt động như thể
    `-mod=readonly` được sử dụng.
  * `go get` từ chối cờ này vì mục đích của lệnh là sửa đổi
    các dependency, điều này chỉ được phép bởi `-mod=mod`.
* Cờ `-modcacherw` yêu cầu lệnh `go` tạo các thư mục mới
  trong cache module với quyền đọc-ghi thay vì chỉ đọc. Khi
  cờ này được sử dụng nhất quán (thường bằng cách đặt
  `GOFLAGS=-modcacherw` trong môi trường hoặc bằng cách chạy
  `go env -w GOFLAGS=-modcacherw`), cache module có thể được xóa bằng
  các lệnh như `rm -r` mà không cần thay đổi quyền trước. Lệnh
  [`go clean -modcache`](#go-clean-modcache) có thể được sử dụng để xóa
  cache module, dù `-modcacherw` có được sử dụng hay không.
* Cờ `-modfile=file.mod` yêu cầu lệnh `go` đọc (và có thể
  ghi) một file thay thế thay vì `go.mod` trong thư mục gốc module. Tên
  file phải kết thúc bằng `.mod`. File có tên `go.mod` vẫn phải tồn tại
  để xác định thư mục gốc module, nhưng nó không được truy cập. Khi
  `-modfile` được chỉ định, một file `go.sum` thay thế cũng được sử dụng: đường dẫn của nó
  được lấy từ cờ `-modfile` bằng cách bỏ phần mở rộng `.mod` và
  thêm `.sum`.

### Vendoring {#vendoring}

Khi sử dụng module, lệnh `go` thường đáp ứng các dependency bằng cách
tải module từ nguồn vào cache module, sau đó tải
các package từ các bản sao đã tải đó. <dfn>Vendoring</dfn> có thể được sử dụng để cho phép
khả năng tương tác với các phiên bản Go cũ hơn, hoặc để đảm bảo rằng tất cả các file được sử dụng cho
một build được lưu trong một cây file duy nhất.

Lệnh [`go mod vendor`](#go-mod-vendor) tạo một thư mục có tên
`vendor` trong thư mục gốc của [module chính](#glos-main-module) chứa
các bản sao của tất cả các package cần thiết để build và kiểm thử các package trong module chính.
Các package chỉ được import bởi các bài kiểm thử của package ngoài module chính thì
không được đưa vào. Cũng giống như [`go mod tidy`](#go-mod-tidy) và các lệnh module khác,
[ràng buộc build](#glos-build-constraint) ngoại trừ `ignore` không
được xét khi tạo thư mục `vendor`.

`go mod vendor` cũng tạo file `vendor/modules.txt` chứa danh sách
các package đã được vendor và các phiên bản module mà chúng được sao chép từ đó. Khi
vendoring được bật, manifest này được sử dụng như một nguồn thông tin phiên bản module,
như được báo cáo bởi [`go list -m`](#go-list-m) và [`go version
-m`](#go-version-m). Khi lệnh `go` đọc `vendor/modules.txt`, nó kiểm tra
rằng các phiên bản module nhất quán với `go.mod`. Nếu `go.mod` đã thay đổi
kể từ khi `vendor/modules.txt` được tạo ra, lệnh `go` sẽ báo lỗi.
`go mod vendor` nên được chạy lại để cập nhật thư mục `vendor`.

Nếu thư mục `vendor` tồn tại trong thư mục gốc của module chính, nó
sẽ được sử dụng tự động nếu [phiên bản `go`](#go-mod-file-go) trong
[file `go.mod`](#glos-go-mod-file) của module chính là `1.14` hoặc cao hơn. Để bật
vendoring rõ ràng, gọi lệnh `go` với cờ `-mod=vendor`. Để
tắt vendoring, sử dụng cờ `-mod=readonly` hoặc `-mod=mod`.

Khi vendoring được bật, [các lệnh build](#build-commands) như `go build` và
`go test` tải package từ thư mục `vendor` thay vì truy cập
mạng hoặc cache module cục bộ. Lệnh [`go list -m`](#go-list-m) chỉ
in thông tin về các module được liệt kê trong `go.mod`. Các lệnh `go mod` như
[`go mod download`](#go-mod-download) và [`go mod tidy`](#go-mod-tidy) không
hoạt động khác khi vendoring được bật và vẫn sẽ tải module và
truy cập cache module. [`go get`](#go-get) cũng không hoạt động khác khi
vendoring được bật.

Không giống như [vendoring trong chế độ `GOPATH`](/s/go15vendor), lệnh `go`
bỏ qua các thư mục vendor ở các vị trí khác ngoài thư mục gốc của module chính.
Ngoài ra, vì các thư mục vendor trong các module khác không được
sử dụng, lệnh `go` không đưa vào các thư mục vendor khi build [các file zip module](#zip-files)
(nhưng xem các lỗi đã biết
[#31562](/issue/31562) và
[#37397](/issue/37397)).

### `go get` {#go-get}

Cách dùng:

```
go get [-d] [-t] [-u] [-tool] [build flags] [packages]
```

Ví dụ:

```
# Nâng cấp một module cụ thể.
$ go get golang.org/x/net

# Nâng cấp các module cung cấp package được import bởi các package trong module chính.
$ go get -u ./...

# Nâng cấp hoặc hạ cấp xuống một phiên bản cụ thể của module.
$ go get golang.org/x/text@v0.3.2

# Cập nhật lên commit trên nhánh master của module.
$ go get golang.org/x/text@master

# Xóa dependency vào module và hạ cấp các module yêu cầu nó
# xuống các phiên bản không yêu cầu nó.
$ go get golang.org/x/text@none

# Nâng cấp phiên bản Go tối thiểu bắt buộc cho module chính.
$ go get go

# Nâng cấp Go toolchain được đề xuất, giữ nguyên phiên bản Go tối thiểu.
$ go get toolchain

# Nâng cấp lên bản vá mới nhất của Go toolchain được đề xuất.
$ go get toolchain@patch
```

Lệnh `go get` cập nhật các dependency module trong [file
`go.mod`](#go-mod-file) cho [module chính](#glos-main-module), sau đó build và
cài đặt các package được liệt kê trên dòng lệnh.

Bước đầu tiên là xác định module nào cần cập nhật. `go get` chấp nhận danh sách
các package, pattern package và đường dẫn module làm đối số. Nếu một đối số
package được chỉ định, `go get` cập nhật module cung cấp package đó.
Nếu một pattern package được chỉ định (ví dụ: `all` hoặc một đường dẫn với ký tự đại diện `...`),
`go get` mở rộng pattern thành một tập hợp các package, sau đó cập nhật các
module cung cấp các package đó. Nếu một đối số đặt tên một module nhưng không đặt tên
package (ví dụ: module `golang.org/x/net` không có package nào trong thư mục gốc của nó),
`go get` sẽ cập nhật module nhưng sẽ không build một package. Nếu không có
đối số nào được chỉ định, `go get` hoạt động như thể `.` được chỉ định (package trong
thư mục hiện tại); điều này có thể được sử dụng cùng với cờ `-u` để cập nhật
các module cung cấp các package được import.

Mỗi đối số có thể bao gồm <dfn>hậu tố truy vấn phiên bản</dfn> chỉ ra
phiên bản mong muốn, như trong `go get golang.org/x/text@v0.3.0`. Một hậu tố truy vấn
phiên bản bao gồm ký hiệu `@` theo sau là một [truy vấn phiên bản](#version-queries),
có thể chỉ ra một phiên bản cụ thể (`v0.3.0`), tiền tố phiên bản (`v0.3`),
tên nhánh hoặc tag (`master`), revision (`1234abcd`), hoặc một trong các
truy vấn đặc biệt `latest`, `upgrade`, `patch`, hoặc `none`. Nếu không có phiên bản nào được đưa ra,
`go get` sử dụng truy vấn `@upgrade`.

Sau khi `go get` đã phân giải các đối số của nó thành các module và phiên bản cụ thể, `go
get` sẽ thêm, thay đổi, hoặc xóa [chỉ thị `require`](#go-mod-file-require) trong
file `go.mod` của module chính để đảm bảo các module vẫn ở phiên bản mong muốn
trong tương lai. Lưu ý rằng các phiên bản bắt buộc trong file `go.mod` là
*các phiên bản tối thiểu* và có thể tự động tăng khi các dependency mới được
thêm vào. Xem [Lựa chọn phiên bản tối thiểu (MVS)](#minimal-version-selection) để biết
chi tiết về cách các phiên bản được chọn và các xung đột được giải quyết bởi các lệnh
nhận biết module.

Các module khác có thể được nâng cấp khi một module được đặt tên trên dòng lệnh được thêm,
nâng cấp, hoặc hạ cấp nếu phiên bản mới của module được đặt tên yêu cầu các module khác
ở phiên bản cao hơn. Ví dụ: giả sử module `example.com/a` được
nâng cấp lên phiên bản `v1.5.0`, và phiên bản đó yêu cầu module `example.com/b`
ở phiên bản `v1.2.0`. Nếu module `example.com/b` hiện đang được yêu cầu ở phiên bản
`v1.1.0`, `go get example.com/a@v1.5.0` cũng sẽ nâng cấp `example.com/b` lên
`v1.2.0`.

![go get nâng cấp một yêu cầu bắc cầu](/doc/mvs/get-upgrade.svg)

Các module khác có thể bị hạ cấp khi một module được đặt tên trên dòng lệnh bị
hạ cấp hoặc bị xóa. Tiếp tục ví dụ trên, giả sử module
`example.com/b` bị hạ cấp xuống `v1.1.0`. Module `example.com/a` cũng sẽ bị
hạ cấp xuống một phiên bản yêu cầu `example.com/b` ở phiên bản `v1.1.0` hoặc
thấp hơn.

![go get hạ cấp một yêu cầu bắc cầu](/doc/mvs/get-downgrade.svg)

Một yêu cầu module có thể bị xóa bằng cách sử dụng hậu tố phiên bản `@none`. Đây là một
loại hạ cấp đặc biệt. Các module phụ thuộc vào module bị xóa sẽ bị
hạ cấp hoặc xóa khi cần thiết. Một yêu cầu module có thể bị xóa ngay cả khi một hoặc
nhiều package của nó được import bởi các package trong module chính. Trong trường hợp này,
lệnh build tiếp theo có thể thêm một yêu cầu module mới.

Nếu một module cần ở hai phiên bản khác nhau (được chỉ định rõ ràng trong các đối số dòng
lệnh hoặc để đáp ứng các nâng cấp và hạ cấp), `go get` sẽ báo một
lỗi.

Sau khi `go get` đã chọn một tập hợp phiên bản mới, nó kiểm tra xem bất kỳ phiên bản module
mới nào được chọn hoặc bất kỳ module nào cung cấp các package được đặt tên trên dòng lệnh
có phải là [bị thu hồi](#glos-retracted-version) hay
[bị deprecated](#glos-deprecated-module) không. `go get` in cảnh báo cho mỗi
phiên bản bị thu hồi hoặc module bị deprecated mà nó tìm thấy. [`go list -m -u
all`](#go-list-m) có thể được sử dụng để kiểm tra các lần thu hồi và deprecated trong tất cả
các dependency.

Sau khi `go get` cập nhật file `go.mod`, nó build các package được đặt tên
trên dòng lệnh. Các file thực thi sẽ được cài đặt trong thư mục được đặt tên bởi
biến môi trường `GOBIN`, mặc định là `$GOPATH/bin` hoặc
`$HOME/go/bin` nếu biến môi trường `GOPATH` không được đặt.

`go get` hỗ trợ các cờ sau:

* Cờ `-d` yêu cầu `go get` không build hoặc cài đặt package. Khi `-d` được
  sử dụng, `go get` sẽ chỉ quản lý các dependency trong `go.mod`. Sử dụng `go get`
  mà không có `-d` để build và cài đặt package đã bị deprecated (kể từ Go 1.17).
  Trong Go 1.18, `-d` sẽ luôn được bật.
* Cờ `-u` yêu cầu `go get` nâng cấp các module cung cấp package
  được import trực tiếp hoặc gián tiếp bởi các package được đặt tên trên dòng lệnh.
  Mỗi module được chọn bởi `-u` sẽ được nâng cấp lên phiên bản mới nhất của nó trừ khi
  nó đã được yêu cầu ở phiên bản cao hơn (một pre-release).
* Cờ `-u=patch` (không phải `-u patch`) cũng yêu cầu `go get` nâng cấp
  các dependency, nhưng `go get` sẽ nâng cấp mỗi dependency lên phiên bản vá mới nhất
  (tương tự như truy vấn phiên bản `@patch`).
* Cờ `-t` yêu cầu `go get` xem xét các module cần thiết để build các bài kiểm thử
  của các package được đặt tên trên dòng lệnh. Khi `-t` và `-u` được sử dụng cùng nhau,
  `go get` cũng sẽ cập nhật các dependency kiểm thử.
* Cờ `-insecure` không nên được sử dụng nữa. Nó cho phép `go get` phân giải
  các đường dẫn import tùy chỉnh và tải về từ các kho lưu trữ và proxy module sử dụng
  các giao thức không an toàn như HTTP. Biến [môi trường](#environment-variables) `GOINSECURE`
  cung cấp kiểm soát chi tiết hơn và
  nên được sử dụng thay thế.
* Cờ `-tool` hướng dẫn go thêm một dòng tool tương ứng vào `go.mod` cho mỗi
  package được liệt kê. Nếu `-tool` được sử dụng với `@none`, dòng đó sẽ bị xóa.

Kể từ Go 1.16, [`go install`](#go-install) là lệnh được khuyến nghị cho
việc build và cài đặt chương trình. Khi được sử dụng với hậu tố phiên bản (như
`@latest` hoặc `@v1.4.6`), `go install` build các package ở chế độ nhận biết module,
bỏ qua file `go.mod` trong thư mục hiện tại hoặc bất kỳ thư mục cha nào,
nếu có.

`go get` tập trung hơn vào quản lý các yêu cầu trong `go.mod`. Cờ `-d`
đã bị deprecated, và kể từ Go 1.18, nó luôn được bật.

### `go install` {#go-install}

Cách dùng:

```
go install [build flags] [packages]
```

Ví dụ:

```
# Cài đặt phiên bản mới nhất của chương trình,
# bỏ qua go.mod trong thư mục hiện tại (nếu có).
$ go install golang.org/x/tools/gopls@latest

# Cài đặt một phiên bản cụ thể của chương trình.
$ go install golang.org/x/tools/gopls@v0.6.4

# Cài đặt chương trình ở phiên bản được chọn bởi module trong thư mục hiện tại.
$ go install golang.org/x/tools/gopls

# Cài đặt tất cả chương trình trong một thư mục.
$ go install ./cmd/...
```

Lệnh `go install` build và cài đặt các package được đặt tên bởi các đường dẫn
trên dòng lệnh. Các file thực thi (các package `main`) được cài đặt vào
thư mục được đặt tên bởi biến môi trường `GOBIN`, mặc định là
`$GOPATH/bin` hoặc `$HOME/go/bin` nếu biến môi trường `GOPATH` không được đặt.
Các file thực thi trong `$GOROOT` được cài đặt trong `$GOROOT/bin` hoặc `$GOTOOLDIR` thay vì
`$GOBIN`. Các package không thực thi được build và cache nhưng không được cài đặt.

Kể từ Go 1.16, nếu các đối số có hậu tố phiên bản (như `@latest` hoặc
`@v1.0.0`), `go install` build các package ở chế độ nhận biết module, bỏ qua file
`go.mod` trong thư mục hiện tại hoặc bất kỳ thư mục cha nào nếu có
một. Điều này hữu ích để cài đặt các file thực thi mà không ảnh hưởng đến
các dependency của module chính.

Để loại bỏ sự mơ hồ về phiên bản module nào được sử dụng trong build, nếu bất kỳ
đối số nào có hậu tố phiên bản, các đối số phải đáp ứng các ràng buộc sau:

* Các đối số phải là đường dẫn package hoặc pattern package (với ký tự đại diện "`...`").
  Chúng không được là các package chuẩn (như `fmt`), meta-pattern (`std`, `cmd`,
  `all`, `work`, `tool`), hoặc các đường dẫn file tương đối hoặc tuyệt đối. Lưu ý rằng
  `go install tool` có thể được sử dụng mà không cần hậu tố phiên bản: xem bên dưới.
* Tất cả các đối số phải có cùng hậu tố phiên bản. Các truy vấn khác nhau không
  được phép, ngay cả khi chúng tham chiếu đến cùng một phiên bản.
* Tất cả các đối số phải tham chiếu đến các package trong cùng một module ở cùng một phiên bản.
* Các đối số đường dẫn package phải tham chiếu đến các package `main`. Các đối số pattern
  sẽ chỉ khớp với các package `main`.
* Không có module nào được coi là [module chính](#glos-main-module).
  * Nếu module chứa các package được đặt tên trên dòng lệnh có file `go.mod`,
    nó không được chứa các chỉ thị (`replace` và `exclude`) mà sẽ
    khiến nó được diễn giải khác đi nếu nó là module chính.
  * Module không được yêu cầu một phiên bản cao hơn của chính nó.
  * Các thư mục vendor không được sử dụng trong bất kỳ module nào. (Các thư mục vendor không
    được đưa vào [các file zip module](#zip-files), vì vậy `go install` không
    tải chúng xuống.)

Xem [Truy vấn phiên bản](#version-queries) để biết cú pháp truy vấn phiên bản được hỗ trợ.
Go 1.15 và thấp hơn không hỗ trợ sử dụng truy vấn phiên bản với `go install`.

Nếu các đối số không có hậu tố phiên bản, `go install` có thể chạy ở
chế độ nhận biết module hoặc chế độ `GOPATH`, tùy thuộc vào biến môi trường `GO111MODULE`
và sự hiện diện của file `go.mod`. Xem [Các lệnh
nhận biết module](#mod-commands) để biết chi tiết. Nếu chế độ nhận biết module được bật, `go
install` chạy trong ngữ cảnh của module chính, có thể khác với
module chứa package đang được cài đặt. Trong chế độ nhận biết module,
`go install tool` có thể được dùng từ một module để cài đặt tất cả các tool trong module đó.

### `go tool` {#go-tool}

Cách dùng:

```
go tool [-n] command [args...]
```

Ví dụ:

```
$ go tool golang.org/x/tools/cmd/stringer
$ go tool stringer
```

Trong chế độ module, lệnh `go tool` có thể được dùng để build và chạy các tool
được khai báo trong file `go.mod` sử dụng [chỉ thị `tool`](#go-mod-file-tool).
Lệnh có thể được chỉ định bằng đường dẫn package đầy đủ đến một tool được khai báo
bằng chỉ thị tool. Tên nhị phân mặc định của tool, là thành phần cuối của đường dẫn
package (không bao gồm hậu tố phiên bản chính), cũng có thể được dùng nếu nó là duy nhất
trong số các tool đã cài đặt.

### `go list -m` {#go-list-m}

Cách dùng:

```
go list -m [-u] [-retracted] [-versions] [list flags] [modules]
```

Ví dụ:

```
$ go list -m all
$ go list -m -versions example.com/m
$ go list -m -json example.com/m@latest
```

Cờ `-m` khiến `go list` liệt kê các module thay vì các package. Ở chế độ
này, các đối số cho `go list` có thể là các module, pattern module (chứa
ký tự đại diện `...`), [truy vấn phiên bản](#version-queries), hoặc pattern đặc biệt
`all`, khớp với tất cả các module trong [danh sách build](#glos-build-list). Nếu không có
đối số nào được chỉ định, [module chính](#glos-main-module) được liệt kê.

Khi liệt kê các module, cờ `-f` vẫn chỉ định một template định dạng áp dụng
cho một struct Go, nhưng bây giờ là struct `Module`:

```
type Module struct {
    Path       string        // module path
    Version    string        // module version
    Versions   []string      // available module versions
    Replace    *Module       // replaced by this module
    Time       *time.Time    // time version was created
    Update     *Module       // available update (with -u)
    Main       bool          // is this the main module?
    Indirect   bool          // module is only indirectly needed by main module
    Dir        string        // directory holding local copy of files, if any
    GoMod      string        // path to go.mod file describing module, if any
    GoVersion  string        // go version used in module
    Retracted  []string      // retraction information, if any (with -retracted or -u)
    Deprecated string        // deprecation message, if any (with -u)
    Error      *ModuleError  // error loading module
}

type ModuleError struct {
    Err string // the error itself
}
```

Đầu ra mặc định là in đường dẫn module và sau đó thông tin về
phiên bản và thay thế nếu có. Ví dụ: `go list -m all` có thể in:

```
example.com/main/module
golang.org/x/net v0.1.0
golang.org/x/text v0.3.0 => /tmp/text
rsc.io/pdf v0.1.1
```

Struct `Module` có phương thức `String` định dạng dòng đầu ra này, vì vậy
định dạng mặc định tương đương với {{raw "`-f '{{.String}}'`"}}.

Lưu ý rằng khi một module đã được thay thế, trường `Replace` của nó mô tả
module thay thế, và trường `Dir` của nó được đặt thành mã nguồn của module
thay thế, nếu có. (Nghĩa là nếu `Replace` khác nil, thì `Dir`
được đặt thành `Replace.Dir`, không có quyền truy cập vào mã nguồn đã được thay thế.)

Cờ `-u` thêm thông tin về các nâng cấp có sẵn. Khi phiên bản mới nhất
của một module đã cho mới hơn phiên bản hiện tại, `list -u` đặt trường `Update`
của module thành thông tin về module mới hơn. `list -u` cũng in
xem phiên bản hiện được chọn có [bị thu hồi](#glos-retracted-version) không
và liệu module có [bị deprecated](#go-mod-file-module-deprecation) không. Phương thức
`String` của module chỉ ra một nâng cấp có sẵn bằng cách định dạng phiên bản mới hơn
trong ngoặc vuông sau phiên bản hiện tại. Ví dụ: `go list -m -u all`
có thể in:

```
example.com/main/module
golang.org/x/old v1.9.9 (deprecated)
golang.org/x/net v0.1.0 (retracted) [v0.2.0]
golang.org/x/text v0.3.0 [v0.4.0] => /tmp/text
rsc.io/pdf v0.1.1 [v0.1.2]
```

(Đối với các công cụ, `go list -m -u -json all` có thể thuận tiện hơn để phân tích cú pháp.)

Cờ `-versions` khiến `list` đặt trường `Versions` của module thành một
danh sách tất cả các phiên bản đã biết của module đó, được sắp xếp theo
semantic versioning, từ thấp đến cao. Cờ này cũng thay đổi định dạng đầu ra mặc định
để hiển thị đường dẫn module theo sau là danh sách phiên bản được phân cách bằng dấu cách.
Các phiên bản bị thu hồi bị bỏ qua khỏi danh sách này trừ khi cờ `-retracted`
cũng được chỉ định.

Cờ `-retracted` yêu cầu `list` hiển thị các phiên bản bị thu hồi trong danh sách
được in với cờ `-versions` và xem xét các phiên bản bị thu hồi khi
phân giải [truy vấn phiên bản](#version-queries). Ví dụ: `go list -m
-retracted example.com/m@latest` hiển thị phiên bản release hoặc pre-release cao nhất
của module `example.com/m`, ngay cả khi phiên bản đó bị thu hồi.
Các [chỉ thị `retract`](#go-mod-file-retract) và
[deprecated](#go-mod-file-module-deprecation) được tải từ file `go.mod`
ở phiên bản này. Cờ `-retracted` được thêm vào trong Go 1.16.

Hàm template `module` nhận một đối số chuỗi duy nhất phải là một
đường dẫn module hoặc truy vấn và trả về module được chỉ định như một struct `Module`. Nếu
xảy ra lỗi, kết quả sẽ là một struct `Module` với trường `Error` khác nil.

### `go mod download` {#go-mod-download}

Cách dùng:

```
go mod download [-x] [-json] [-reuse=old.json] [modules]
```

Ví dụ:

```
$ go mod download
$ go mod download golang.org/x/mod@v0.2.0
```

Lệnh `go mod download` tải các module được đặt tên vào [cache
module](#glos-module-cache). Các đối số có thể là đường dẫn module hoặc pattern
module chọn các dependency của module chính hoặc [truy vấn
phiên bản](#version-queries) có dạng `path@version`. Không có đối số nào,
`download` áp dụng cho tất cả các dependency của [module chính](#glos-main-module).

Lệnh `go` sẽ tự động tải module khi cần thiết trong quá trình thực thi bình thường.
Lệnh `go mod download` hữu ích chủ yếu để điền trước cache
module hoặc để tải dữ liệu để được phục vụ bởi một [proxy
module](#glos-module-proxy).

Theo mặc định, `download` không ghi gì vào đầu ra tiêu chuẩn. Nó in các thông báo tiến trình
và lỗi vào lỗi tiêu chuẩn.

Cờ `-json` khiến `download` in một chuỗi các đối tượng JSON vào
đầu ra tiêu chuẩn, mô tả từng module được tải xuống (hoặc thất bại), tương ứng
với struct Go này:

```
type Module struct {
    Path     string // module path
    Query    string // version query corresponding to this version
    Version  string // module version
    Error    string // error loading module
    Info     string // absolute path to cached .info file
    GoMod    string // absolute path to cached .mod file
    Zip      string // absolute path to cached .zip file
    Dir      string // absolute path to cached source root directory
    Sum      string // checksum for path, version (as in go.sum)
    GoModSum string // checksum for go.mod (as in go.sum)
    Origin   any    // provenance of module
    Reuse    bool   // reuse of old module info is safe
}
```

Cờ `-x` khiến `download` in các lệnh mà `download` thực thi
vào lỗi tiêu chuẩn.

Cờ -reuse chấp nhận tên của file chứa đầu ra JSON của một lần gọi
'go mod download -json' trước đó. Lệnh go có thể sử dụng file này
để xác định rằng một module không thay đổi kể từ lần gọi trước
và tránh tải lại. Các module không được tải lại sẽ được đánh dấu
trong đầu ra mới bằng cách đặt trường Reuse thành true. Thông thường, cache module
cung cấp loại tái sử dụng này tự động; cờ -reuse có thể hữu ích
trên các hệ thống không bảo toàn cache module.

### `go mod edit` {#go-mod-edit}

Cách dùng:

```
go mod edit [editing flags] [-fmt|-print|-json] [go.mod]
```

Ví dụ:

```
# Thêm chỉ thị replace.
$ go mod edit -replace example.com/a@v1.0.0=./a

# Xóa chỉ thị replace.
$ go mod edit -dropreplace example.com/a@v1.0.0

# Đặt phiên bản go, thêm yêu cầu, và in file
# thay vì ghi nó vào đĩa.
$ go mod edit -go=1.14 -require=example.com/m@v1.0.0 -print

# Định dạng file go.mod.
$ go mod edit -fmt

# Định dạng và in một file .mod khác.
$ go mod edit -print tools.mod

# In biểu diễn JSON của file go.mod.
$ go mod edit -json
```

Lệnh `go mod edit` cung cấp giao diện dòng lệnh để chỉnh sửa và
định dạng các file `go.mod`, chủ yếu để sử dụng bởi các công cụ và script. `go mod edit`
chỉ đọc một file `go.mod`; nó không tra cứu thông tin về các module khác.
Theo mặc định, `go mod edit` đọc và ghi file `go.mod` của
module chính, nhưng một file đích khác có thể được chỉ định sau các cờ chỉnh sửa.

Các cờ chỉnh sửa chỉ định một chuỗi các thao tác chỉnh sửa.

* Cờ `-module` thay đổi đường dẫn của module (dòng module trong file `go.mod`).
* Cờ `-go=version` đặt phiên bản ngôn ngữ Go dự kiến.
* Các cờ `-require=path@version` và `-droprequire=path` thêm và bỏ một
  yêu cầu trên đường dẫn module và phiên bản đã cho. Lưu ý rằng `-require`
  ghi đè bất kỳ yêu cầu hiện có nào trên `path`. Các cờ này chủ yếu dành cho
  các công cụ hiểu biết về đồ thị module. Người dùng nên ưu tiên `go get
  path@version` hoặc `go get path@none`, thực hiện các điều chỉnh `go.mod` khác khi
  cần thiết để đáp ứng các ràng buộc được áp đặt bởi các module khác. Xem [`go
  get`](#go-get).
* Các cờ `-exclude=path@version` và `-dropexclude=path@version` thêm và bỏ
  một loại trừ cho đường dẫn module và phiên bản đã cho. Lưu ý rằng
  `-exclude=path@version` là no-op nếu loại trừ đó đã tồn tại.
* Cờ `-replace=old[@v]=new[@v]` thêm một thay thế của cặp đường dẫn module
  và phiên bản đã cho. Nếu `@v` trong `old@v` bị bỏ qua, một thay thế
  không có phiên bản ở phía bên trái được thêm vào, áp dụng cho tất cả các phiên bản của
  đường dẫn module cũ. Nếu `@v` trong `new@v` bị bỏ qua, đường dẫn mới phải là
  một thư mục gốc module cục bộ, không phải đường dẫn module. Lưu ý rằng `-replace`
  ghi đè bất kỳ thay thế dư thừa nào cho `old[@v]`, vì vậy bỏ qua `@v` sẽ xóa
  các thay thế cho các phiên bản cụ thể.
* Cờ `-dropreplace=old[@v]` xóa một thay thế của cặp đường dẫn module
  và phiên bản đã cho. Nếu `@v` được cung cấp, một thay thế với phiên bản đã cho
  bị xóa. Một thay thế hiện có không có phiên bản ở phía bên trái
  vẫn có thể thay thế module. Nếu `@v` bị bỏ qua, một thay thế không có
  phiên bản bị xóa.
* Các cờ `-retract=version` và `-dropretract=version` thêm và bỏ một
  thu hồi cho phiên bản đã cho, có thể là một phiên bản duy nhất (như
  `v1.2.3`) hoặc một khoảng (như `[v1.1.0,v1.2.0]`). Lưu ý rằng cờ `-retract`
  không thể thêm comment lý do cho chỉ thị `retract`. Các comment lý do
  được khuyến nghị và có thể được hiển thị bởi `go list -m -u` và các lệnh khác.
* Các cờ `-tool=path` và `-droptool=path` thêm và bỏ một chỉ thị `tool`
  cho các đường dẫn đã cho. Lưu ý rằng điều này sẽ không thêm các dependency cần thiết vào
  đồ thị build. Người dùng nên ưu tiên `go get -tool path` để thêm một công cụ, hoặc
  `go get -tool path@none` để xóa một công cụ.

Các cờ chỉnh sửa có thể được lặp lại. Các thay đổi được áp dụng theo thứ tự đã cho.

`go mod edit` có thêm các cờ kiểm soát đầu ra của nó.

* Cờ `-fmt` định dạng lại file `go.mod` mà không thực hiện các thay đổi khác.
  Việc định dạng lại này cũng được ngụ ý bởi bất kỳ sửa đổi nào khác sử dụng hoặc
  viết lại file `go.mod`. Cờ này chỉ cần thiết khi không có
  cờ nào khác được chỉ định, như trong `go mod edit -fmt`.
* Cờ `-print` in `go.mod` cuối cùng ở định dạng văn bản thay vì
  ghi nó trở lại vào đĩa.
* Cờ `-json` in `go.mod` cuối cùng ở định dạng JSON thay vì ghi
  nó trở lại vào đĩa ở định dạng văn bản. Đầu ra JSON tương ứng với các kiểu Go sau:

```
type Module struct {
    Path    string
    Version string
}

type GoMod struct {
    Module  ModPath
    Go      string
    Require []Require
    Exclude []Module
    Replace []Replace
    Retract []Retract
}

type ModPath struct {
    Path       string
    Deprecated string
}

type Require struct {
    Path     string
    Version  string
    Indirect bool
}

type Replace struct {
    Old Module
    New Module
}

type Retract struct {
    Low       string
    High      string
    Rationale string
}

type Tool struct {
    Path      string
}
```

Lưu ý rằng điều này chỉ mô tả bản thân file `go.mod`, không phải các module khác
được tham chiếu gián tiếp. Để có toàn bộ tập hợp các module có sẵn cho một build,
sử dụng `go list -m -json all`. Xem [`go list -m`](#go-list-m).

Ví dụ: một công cụ có thể lấy file `go.mod` như một cấu trúc dữ liệu bằng cách
phân tích đầu ra của `go mod edit -json` và sau đó thực hiện các thay đổi bằng cách gọi
`go mod edit` với `-require`, `-exclude`, v.v.

Các công cụ cũng có thể sử dụng package
[`golang.org/x/mod/modfile`](https://pkg.go.dev/golang.org/x/mod/modfile?tab=doc)
để phân tích cú pháp, chỉnh sửa và định dạng các file `go.mod`.

### `go mod graph` {#go-mod-graph}

Cách dùng:

```
go mod graph [-go=version]
```

Lệnh `go mod graph` in [đồ thị yêu cầu
module](#glos-module-graph) (với các thay thế được áp dụng) ở dạng văn bản. Ví dụ:

```
example.com/main example.com/a@v1.1.0
example.com/main example.com/b@v1.2.0
example.com/a@v1.1.0 example.com/b@v1.1.1
example.com/a@v1.1.0 example.com/c@v1.3.0
example.com/b@v1.1.0 example.com/c@v1.1.0
example.com/b@v1.2.0 example.com/c@v1.2.0
```

Mỗi đỉnh trong đồ thị module đại diện cho một phiên bản cụ thể của một module.
Mỗi cạnh trong đồ thị đại diện cho một yêu cầu về phiên bản tối thiểu của một
dependency.

`go mod graph` in các cạnh của đồ thị, mỗi cạnh một dòng. Mỗi dòng có hai
trường được phân cách bằng dấu cách: một phiên bản module và một trong các dependency của nó. Mỗi
phiên bản module được xác định như một chuỗi có dạng `path@version`. Module chính
không có hậu tố `@version`, vì nó không có phiên bản.

Cờ `-go` khiến `go mod graph` báo cáo đồ thị module như
được tải bởi phiên bản Go đã cho, thay vì phiên bản được chỉ ra bởi
[chỉ thị `go`](#go-mod-file-go) trong file `go.mod`.

Xem [Lựa chọn phiên bản tối thiểu (MVS)](#minimal-version-selection) để biết thêm
thông tin về cách các phiên bản được chọn. Xem thêm [`go list -m`](#go-list-m) để
in các phiên bản được chọn và [`go mod why`](#go-mod-why) để hiểu
tại sao một module lại cần thiết.

### `go mod init` {#go-mod-init}

Cách dùng:

```
go mod init [module-path]
```

Ví dụ:

```
go mod init
go mod init example.com/m
```

Lệnh `go mod init` khởi tạo và ghi một file `go.mod` mới vào
thư mục hiện tại, thực sự tạo một module mới bắt đầu tại thư mục
hiện tại. File `go.mod` không được tồn tại từ trước.

`init` chấp nhận một đối số tùy chọn, [đường dẫn module](#glos-module-path) cho
module mới. Xem [Đường dẫn module](#module-path) để biết hướng dẫn về cách chọn
một đường dẫn module. Nếu đối số đường dẫn module bị bỏ qua, `init` sẽ cố gắng
suy ra đường dẫn module bằng cách sử dụng các comment import trong các file `.go` và
thư mục hiện tại (nếu ở trong `GOPATH`).

### `go mod tidy` {#go-mod-tidy}

Cách dùng:

```
go mod tidy [-e] [-v] [-x] [-diff] [-go=version] [-compat=version]
```

`go mod tidy` đảm bảo rằng file `go.mod` khớp với mã nguồn trong
module. Nó thêm bất kỳ yêu cầu module nào còn thiếu cần thiết để build
các package và dependency của module hiện tại, và nó xóa các yêu cầu trên các module không
cung cấp bất kỳ package nào có liên quan. Nó cũng thêm bất kỳ entry nào còn thiếu vào
`go.sum` và xóa các entry không cần thiết.

Cờ `-e` (thêm vào trong Go 1.16) khiến `go mod tidy` cố gắng tiếp tục
mặc dù có lỗi gặp phải khi tải package.

Cờ `-v` khiến `go mod tidy` in thông tin về các module đã xóa
vào lỗi tiêu chuẩn.

Cờ `-x` khiến `go mod tidy` in các lệnh mà `tidy` thực thi.

Cờ `-diff` khiến `go mod tidy` không sửa đổi go.mod hoặc go.sum mà
thay vào đó in các thay đổi cần thiết dưới dạng unified diff. Nó thoát
với mã khác 0 nếu diff không rỗng.

`go mod tidy` hoạt động bằng cách tải tất cả các package trong [module
chính](#glos-main-module), tất cả các công cụ của nó, và tất cả các package mà chúng import,
đệ quy. Điều này bao gồm các package được import bởi các bài kiểm thử (bao gồm cả kiểm thử trong các module khác).
`go mod tidy` hoạt động như thể tất cả các build tag đều được bật, vì vậy nó sẽ
xem xét các file nguồn dành riêng cho nền tảng và các file yêu cầu build tag tùy chỉnh,
ngay cả khi những file nguồn đó thường không được build. Có một
ngoại lệ: build tag `ignore` không được bật, vì vậy một file có ràng buộc build
`// +build ignore` sẽ không được xem xét. Lưu ý rằng `go mod tidy`
sẽ không xem xét các package trong module chính trong các thư mục có tên là `testdata` hoặc
có tên bắt đầu bằng `.` hoặc `_` trừ khi những package đó được import rõ ràng
bởi các package khác.

Sau khi `go mod tidy` đã tải tập hợp package này, nó đảm bảo rằng mỗi module
cung cấp một hoặc nhiều package có chỉ thị `require` trong file `go.mod` của
module chính hoặc, nếu module chính ở `go 1.16` hoặc thấp hơn, được
yêu cầu bởi một module bắt buộc khác. `go mod tidy` sẽ thêm một yêu cầu trên
phiên bản mới nhất của mỗi module còn thiếu (xem [Truy vấn phiên bản](#version-queries)
để biết định nghĩa của phiên bản `latest`). `go mod tidy` sẽ xóa các chỉ thị `require`
cho các module không cung cấp bất kỳ package nào trong tập hợp được mô tả
ở trên.

`go mod tidy` cũng có thể thêm hoặc xóa các comment `// indirect` trên các chỉ thị `require`.
Một comment `// indirect` biểu thị một module không cung cấp
package được import bởi một package trong module chính. (Xem [chỉ thị
`require`](#go-mod-file-require) để biết thêm chi tiết về khi nào các
dependency và comment `// indirect` được thêm vào.)

Nếu cờ `-go` được đặt, `go mod tidy` sẽ cập nhật [chỉ thị
`go`](#go-mod-file-go) lên phiên bản được chỉ định, bật hoặc tắt
[cắt tỉa đồ thị module](#graph-pruning) và [tải module lười biếng](#lazy-loading)
(và thêm hoặc xóa các yêu cầu gián tiếp khi cần thiết) theo
phiên bản đó.

Theo mặc định, `go mod tidy` sẽ kiểm tra rằng các [phiên bản được
chọn](#glos-selected-version) của module không thay đổi khi đồ thị module
được tải bởi phiên bản Go ngay trước phiên bản được chỉ ra trong chỉ thị
`go`. Phiên bản được kiểm tra về khả năng tương thích cũng có thể được chỉ định
rõ ràng thông qua cờ `-compat`.

### `go mod vendor` {#go-mod-vendor}

Cách dùng:

```
go mod vendor [-e] [-v] [-o]
```

Lệnh `go mod vendor` tạo một thư mục có tên `vendor` trong thư mục gốc của
[module chính](#glos-main-module) chứa các bản sao của tất cả các package
cần thiết để hỗ trợ build và kiểm thử các package trong module chính. Các package
chỉ được import bởi các bài kiểm thử của package ngoài module chính thì không
được đưa vào. Cũng giống như [`go mod tidy`](#go-mod-tidy) và các lệnh module khác,
[ràng buộc build](#glos-build-constraint) ngoại trừ `ignore` không
được xét khi tạo thư mục `vendor`.

Khi vendoring được bật, lệnh `go` sẽ tải các package từ thư mục `vendor`
thay vì tải các module từ nguồn vào cache module
và sử dụng các package từ các bản sao đã tải đó. Xem [Vendoring](#vendoring)
để biết thêm thông tin.

`go mod vendor` cũng tạo file `vendor/modules.txt` chứa danh sách
các package đã được vendor và các phiên bản module mà chúng được sao chép từ đó. Khi
vendoring được bật, manifest này được sử dụng như một nguồn thông tin phiên bản module,
như được báo cáo bởi [`go list -m`](#go-list-m) và [`go version
-m`](#go-version-m). Khi lệnh `go` đọc `vendor/modules.txt`, nó kiểm tra
rằng các phiên bản module nhất quán với `go.mod`. Nếu `go.mod` thay đổi kể từ
khi `vendor/modules.txt` được tạo ra, `go mod vendor` nên được chạy lại.

Lưu ý rằng `go mod vendor` xóa thư mục `vendor` nếu nó tồn tại trước khi
tạo lại nó. Không nên thực hiện các thay đổi cục bộ đối với các package đã được vendor.
Lệnh `go` không kiểm tra rằng các package trong thư mục `vendor` không
bị sửa đổi, nhưng người ta có thể xác minh tính toàn vẹn của thư mục `vendor`
bằng cách chạy `go mod vendor` và kiểm tra rằng không có thay đổi nào được thực hiện.

Cờ `-e` (thêm vào trong Go 1.16) khiến `go mod vendor` cố gắng tiếp tục
mặc dù có lỗi gặp phải khi tải package.

Cờ `-v` khiến `go mod vendor` in tên các module và
package đã được vendor vào lỗi tiêu chuẩn.

Cờ `-o` (thêm vào trong Go 1.18) khiến `go mod vendor` xuất cây vendor
vào thư mục được chỉ định thay vì `vendor`. Đối số có thể là một
đường dẫn tuyệt đối hoặc đường dẫn tương đối với thư mục gốc module.

### `go mod verify` {#go-mod-verify}

Cách dùng:

```
go mod verify
```

`go mod verify` kiểm tra rằng các dependency của [module chính](#glos-main-module)
được lưu trữ trong [cache module](#glos-module-cache) không bị sửa đổi kể từ
khi chúng được tải xuống. Để thực hiện kiểm tra này, `go mod verify` băm từng
[file `.zip`](#zip-files) module đã tải và thư mục đã giải nén, sau đó
so sánh các băm đó với một băm được ghi lại khi module được
tải xuống lần đầu tiên. `go mod verify` kiểm tra từng module trong [danh sách
build](#glos-build-list) (có thể được in với [`go list -m
all`](#go-list-m)).

Nếu tất cả các module không bị sửa đổi, `go mod verify` in "all modules
verified". Ngược lại, nó báo cáo module nào đã bị thay đổi và thoát với
trạng thái khác 0.

Lưu ý rằng tất cả các lệnh nhận biết module xác minh rằng các băm trong file `go.sum`
của module chính khớp với các băm được ghi lại cho các module được tải xuống vào cache module.
Nếu một băm bị thiếu trong `go.sum` (ví dụ: vì module được
sử dụng lần đầu tiên), lệnh `go` xác minh băm của nó bằng cách sử dụng
[cơ sở dữ liệu checksum](#checksum-database) (trừ khi đường dẫn module được khớp bởi
`GOPRIVATE` hoặc `GONOSUMDB`). Xem [Xác thực module](#authenticating) để biết
chi tiết.

Ngược lại, `go mod verify` kiểm tra rằng các file `.zip` module và các thư mục đã giải nén
của chúng có các băm khớp với các băm được ghi lại trong cache module khi chúng
được tải xuống lần đầu tiên. Điều này hữu ích để phát hiện các thay đổi đối với các file trong
cache module *sau khi* một module đã được tải xuống và xác minh. `go mod verify`
không tải nội dung cho các module không có trong cache, và nó không sử dụng
các file `go.sum` để xác minh nội dung module. Tuy nhiên, `go mod verify` có thể tải xuống
các file `go.mod` để thực hiện [lựa chọn phiên bản
tối thiểu](#minimal-version-selection). Nó sẽ sử dụng `go.sum` để xác minh các
file đó, và nó có thể thêm các entry `go.sum` cho các băm bị thiếu.

### `go mod why` {#go-mod-why}

Cách dùng:

```
go mod why [-m] [-vendor] packages...
```

`go mod why` hiển thị đường đi ngắn nhất trong đồ thị import từ module chính đến
từng package được liệt kê.

Đầu ra là một chuỗi các đoạn, một đoạn cho mỗi package hoặc module được đặt tên trên
dòng lệnh, được phân cách bằng các dòng trống. Mỗi đoạn bắt đầu bằng một dòng comment
bắt đầu bằng `#` cho biết package hoặc module đích. Các dòng tiếp theo cho biết một
đường dẫn qua đồ thị import, mỗi dòng một package. Nếu package hoặc module
không được tham chiếu từ module chính, đoạn sẽ hiển thị một chú thích đơn
trong ngoặc đơn chỉ ra điều đó.

Ví dụ:

```
$ go mod why golang.org/x/text/language golang.org/x/text/encoding
# golang.org/x/text/language
rsc.io/quote
rsc.io/sampler
golang.org/x/text/language

# golang.org/x/text/encoding
(main module does not need package golang.org/x/text/encoding)
```

Cờ `-m` khiến `go mod why` xử lý các đối số của nó như một danh sách các module.
`go mod why` sẽ in một đường dẫn đến bất kỳ package nào trong mỗi module. Lưu ý rằng
ngay cả khi `-m` được sử dụng, `go mod why` truy vấn đồ thị package, không phải
đồ thị module được in bởi [`go mod graph`](#go-mod-graph).

Cờ `-vendor` khiến `go mod why` bỏ qua các import trong các bài kiểm thử của các package
ngoài module chính (như [`go mod vendor`](#go-mod-vendor) làm). Theo mặc định,
`go mod why` xem xét đồ thị của các package khớp với pattern `all`. Cờ này
không có hiệu lực sau Go 1.16 trong các module khai báo `go 1.16` hoặc cao hơn
(sử dụng [chỉ thị `go`](#go-mod-file-go) trong `go.mod`), vì ý nghĩa của
`all` thay đổi để khớp với tập hợp các package được khớp bởi `go mod vendor`.

### `go version -m` {#go-version-m}

Cách dùng:

```
go version [-m] [-v] [file ...]
```

Ví dụ:

```
# In phiên bản Go được sử dụng để build go.
$ go version

# In phiên bản Go được sử dụng để build một file thực thi cụ thể.
$ go version ~/go/bin/gopls

# In phiên bản Go và phiên bản module được sử dụng để build một file thực thi cụ thể.
$ go version -m ~/go/bin/gopls

# In phiên bản Go và phiên bản module được sử dụng để build các file thực thi trong một thư mục.
$ go version -m ~/go/bin/
```

`go version` báo cáo phiên bản Go được sử dụng để build từng file thực thi được đặt tên
trên dòng lệnh.

Nếu không có file nào được đặt tên trên dòng lệnh, `go version` in thông tin phiên bản của chính nó.

Nếu một thư mục được đặt tên, `go version` duyệt qua thư mục đó, đệ quy, tìm kiếm
các file nhị phân Go đã được nhận dạng và báo cáo phiên bản của chúng. Theo mặc định, `go
version` không báo cáo các file không được nhận dạng tìm thấy trong quá trình quét thư mục. Cờ
`-v` khiến nó báo cáo các file không được nhận dạng.

Cờ `-m` khiến `go version` in thông tin phiên bản module được nhúng của mỗi file thực thi,
khi có sẵn. Đối với mỗi file thực thi, `go version -m` in
một bảng với các cột được phân cách bằng tab như bảng bên dưới.

```
$ go version -m ~/go/bin/goimports
/home/jrgopher/go/bin/goimports: go1.14.3
        path    golang.org/x/tools/cmd/goimports
        mod     golang.org/x/tools      v0.0.0-20200518203908-8018eb2c26ba      h1:0Lcy64USfQQL6GAJma8BdHCgeofcchQj+Z7j0SXYAzU=
        dep     golang.org/x/mod        v0.2.0          h1:KU7oHjnv3XNWfa5COkzUifxZmxp1TyI7ImMXqFxLwvQ=
        dep     golang.org/x/xerrors    v0.0.0-20191204190536-9bdfabe68543      h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=
```

Định dạng của bảng có thể thay đổi trong tương lai. Thông tin tương tự có thể được
lấy từ
[`runtime/debug.ReadBuildInfo`](https://pkg.go.dev/runtime/debug?tab=doc#ReadBuildInfo).

Ý nghĩa của mỗi hàng trong bảng được xác định bởi từ trong cột đầu tiên.

* **`path`**: đường dẫn của package `main` được sử dụng để build file thực thi.
* **`mod`**: module chứa package `main`. Các cột là
  đường dẫn module, phiên bản và tổng kiểm tra, tương ứng. [Module
  chính](#glos-main-module) có phiên bản `(devel)` và không có tổng kiểm tra.
* **`dep`**: một module cung cấp một hoặc nhiều package được liên kết vào
  file thực thi. Cùng định dạng với `mod`.
* **`=>`**: một [thay thế](#go-mod-file-replace) cho module trên dòng trước.
  Nếu thay thế là một thư mục cục bộ, chỉ đường dẫn thư mục được
  liệt kê (không có phiên bản hoặc tổng kiểm tra). Nếu thay thế là một phiên bản module, đường dẫn,
  phiên bản và tổng kiểm tra được liệt kê, như với `mod` và `dep`. Một module đã thay thế không có
  tổng kiểm tra.

### `go clean -modcache` {#go-clean-modcache}

Cách dùng:

```
go clean [-modcache]
```

Cờ `-modcache` khiến [`go
clean`](/cmd/go/#hdr-Remove_object_files_and_cached_files) xóa toàn bộ
[cache module](#glos-module-cache), bao gồm cả mã nguồn đã giải nén của các
dependency theo phiên bản.

Đây thường là cách tốt nhất để xóa cache module. Theo mặc định, hầu hết các file
và thư mục trong cache module là chỉ đọc để ngăn các bài kiểm thử và trình soạn thảo
vô tình thay đổi các file sau khi chúng đã được
[xác thực](#authenticating). Thật không may, điều này khiến các lệnh như
`rm -r` thất bại, vì các file không thể bị xóa mà không trước tiên làm cho thư mục cha
của chúng có thể ghi.

Cờ `-modcacherw` (được chấp nhận bởi [`go
build`](/cmd/go/#hdr-Compile_packages_and_dependencies) và
các lệnh nhận biết module khác) khiến các thư mục mới trong cache module
có thể ghi. Để truyền `-modcacherw` cho tất cả các lệnh nhận biết module, thêm nó vào
biến `GOFLAGS`. `GOFLAGS` có thể được đặt trong môi trường hoặc với [`go env
-w`](/cmd/go/#hdr-Print_Go_environment_information). Ví dụ:
lệnh dưới đây đặt nó vĩnh viễn:

```
go env -w GOFLAGS=-modcacherw
```

`-modcacherw` nên được sử dụng một cách cẩn thận; các nhà phát triển nên cẩn thận không
thực hiện các thay đổi đối với các file trong cache module. [`go mod verify`](#go-mod-verify)
có thể được sử dụng để kiểm tra rằng các file trong cache khớp với các băm trong file `go.sum`
của module chính.

### Truy vấn phiên bản {#version-queries}

Một số lệnh cho phép bạn chỉ định một phiên bản của một module bằng cách sử dụng *truy vấn phiên bản*,
xuất hiện sau ký tự `@` theo sau một đường dẫn module hoặc package
trên dòng lệnh.

Ví dụ:

```
go get example.com/m@latest
go mod download example.com/m@master
go list -m -json example.com/m@e3702bed2
```

Một truy vấn phiên bản có thể là một trong những dạng sau:

* Một phiên bản semantic đầy đủ, như `v1.2.3`, chọn một
  phiên bản cụ thể. Xem [Phiên bản](#versions) để biết cú pháp.
* Một tiền tố phiên bản semantic, như `v1` hoặc `v1.2`, chọn phiên bản
  cao nhất có sẵn với tiền tố đó.
* Một so sánh phiên bản semantic, như {{raw "`<v1.2.3` hoặc `>=v1.5.6`"}}, chọn
  phiên bản có sẵn gần nhất với mục tiêu so sánh (phiên bản thấp nhất
  cho `>` và `>=`, và phiên bản cao nhất cho {{raw "`<` và `<=`"}}).
* Một định danh revision cho kho lưu trữ mã nguồn bên dưới, như tiền tố băm commit,
  thẻ revision, hoặc tên nhánh. Nếu revision được gắn thẻ với một
  phiên bản semantic, truy vấn này chọn phiên bản đó. Ngược lại, truy vấn này chọn
  một [pseudo-version](#glos-pseudo-version) cho commit bên dưới.
  Lưu ý rằng các nhánh và thẻ có tên được khớp bởi các truy vấn phiên bản khác
  không thể được chọn theo cách này. Ví dụ: truy vấn `v2` chọn
  phiên bản mới nhất bắt đầu bằng `v2`, không phải nhánh có tên là `v2`.
* Chuỗi `latest`, chọn phiên bản release cao nhất có sẵn. Nếu
  không có phiên bản release nào, `latest` chọn phiên bản pre-release cao nhất.
  Nếu không có phiên bản nào được gắn thẻ, `latest` chọn một pseudo-version cho
  commit ở đầu nhánh mặc định của kho lưu trữ.
* Chuỗi `upgrade`, giống như `latest` ngoại trừ nếu module hiện đang
  được yêu cầu ở phiên bản cao hơn phiên bản mà `latest` sẽ chọn
  (ví dụ: một pre-release), `upgrade` sẽ chọn phiên bản hiện tại.
* Chuỗi `patch`, chọn phiên bản mới nhất có sẵn với cùng
  số phiên bản major và minor như phiên bản hiện được yêu cầu. Nếu không có
  phiên bản nào hiện đang được yêu cầu, `patch` tương đương với `latest`. Kể từ
  Go 1.16, [`go get`](#go-get) yêu cầu phiên bản hiện tại khi sử dụng `patch`
  (nhưng cờ `-u=patch` không có yêu cầu này).

Ngoại trừ các truy vấn cho các phiên bản hoặc revision được đặt tên cụ thể, tất cả các truy vấn
xem xét các phiên bản có sẵn được báo cáo bởi `go list -m -versions` (xem [`go list
-m`](#go-list-m)). Danh sách này chỉ chứa các phiên bản được gắn thẻ, không phải các pseudo-version.
Các phiên bản module không được phép bởi [chỉ thị `exclude`](#go-mod-file-exclude) trong
[file `go.mod`](#glos-go-mod-file) của module chính không được xem xét.
Các phiên bản được bao phủ bởi [chỉ thị `retract`](#go-mod-file-retract) trong file `go.mod`
từ phiên bản `latest` của cùng một module cũng bị bỏ qua ngoại trừ khi
cờ `-retracted` được sử dụng với [`go list -m`](#go-list-m) và ngoại trừ khi
tải các chỉ thị `retract`.

[Các phiên bản release](#glos-release-version) được ưu tiên hơn các phiên bản pre-release. Ví dụ:
nếu có các phiên bản `v1.2.2` và `v1.2.3-pre`, truy vấn
`latest` sẽ chọn `v1.2.2`, mặc dù `v1.2.3-pre` cao hơn. Truy vấn
{{raw "`<v1.2.4`"}} cũng sẽ chọn `v1.2.2`, mặc dù `v1.2.3-pre` gần hơn
với `v1.2.4`. Nếu không có phiên bản release hoặc pre-release nào có sẵn, các truy vấn `latest`,
`upgrade` và `patch` sẽ chọn một pseudo-version cho commit
ở đầu nhánh mặc định của kho lưu trữ. Các truy vấn khác sẽ báo cáo
một lỗi.

### Các lệnh module bên ngoài một module {#commands-outside}

Các lệnh Go nhận biết module thường chạy trong ngữ cảnh của một [module
chính](#glos-main-module) được xác định bởi file `go.mod` trong thư mục làm việc
hoặc một thư mục cha. Một số lệnh có thể chạy ở chế độ nhận biết module mà không cần
file `go.mod`, nhưng hầu hết các lệnh hoạt động khác hoặc báo lỗi khi không có
file `go.mod` nào tồn tại.

Xem [Các lệnh nhận biết module](#mod-commands) để biết thông tin về việc bật và
tắt chế độ nhận biết module.

<table class="ModTable">
  <thead>
    <tr>
      <th>Lệnh</th>
      <th>Hành vi</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>
        <code>go build</code><br>
        <code>go doc</code><br>
        <code>go fix</code><br>
        <code>go fmt</code><br>
        <code>go generate</code><br>
        <code>go install</code><br>
        <code>go list</code><br>
        <code>go run</code><br>
        <code>go test</code><br>
        <code>go vet</code>
      </td>
      <td>
        Chỉ các package trong thư viện chuẩn và các package được chỉ định là
        các file <code>.go</code> trên dòng lệnh mới có thể được tải, import và
        build. Các package từ các module khác không thể được build, vì không có
        nơi nào để ghi lại các yêu cầu module và đảm bảo các build có tính tất định.
      </td>
    </tr>
    <tr>
      <td><code>go get</code></td>
      <td>
        Các package và file thực thi có thể được build và cài đặt như thường lệ. Lưu ý rằng
        không có module chính khi <code>go get</code> được chạy mà không có
        file <code>go.mod</code>, vì vậy các chỉ thị <code>replace</code> và
        <code>exclude</code> không được áp dụng.
      </td>
    </tr>
    <tr>
      <td><code>go list -m</code></td>
      <td>
        <a href="#version-queries">Truy vấn phiên bản</a> rõ ràng là bắt buộc
        cho hầu hết các đối số, ngoại trừ khi cờ <code>-versions</code> được sử dụng.
      </td>
    </tr>
    <tr>
      <td><code>go mod download</code></td>
      <td>
        <a href="#version-queries">Truy vấn phiên bản</a> rõ ràng là bắt buộc
        cho hầu hết các đối số.
      </td>
    </tr>
    <tr>
      <td><code>go mod edit</code></td>
      <td>Một đối số file rõ ràng là bắt buộc.</td>
    </tr>
    <tr>
      <td>
        <code>go mod graph</code><br>
        <code>go mod tidy</code><br>
        <code>go mod vendor</code><br>
        <code>go mod verify</code><br>
        <code>go mod why</code>
      </td>
      <td>
        Các lệnh này yêu cầu một file <code>go.mod</code> và sẽ báo
        lỗi nếu không có file nào tồn tại.
      </td>
    </tr>
  </tbody>
</table>

### `go work init` {#go-work-init}

Cách dùng:

```
go work init [moddirs]
```

Init khởi tạo và ghi một file go.work mới vào
thư mục hiện tại, thực sự tạo một workspace mới tại thư mục
hiện tại.

go work init tùy chọn chấp nhận các đường dẫn đến các module workspace như
đối số. Nếu đối số bị bỏ qua, một workspace trống không có
module nào sẽ được tạo.

Mỗi đường dẫn đối số được thêm vào một chỉ thị use trong file go.work. Phiên bản
go hiện tại cũng sẽ được liệt kê trong file go.work.

### `go work edit` {#go-work-edit}

Cách dùng:

```
go work edit [editing flags] [go.work]
```

Lệnh `go work edit` cung cấp giao diện dòng lệnh để chỉnh sửa `go.work`,
chủ yếu để sử dụng bởi các công cụ hoặc script. Nó chỉ đọc `go.work`;
nó không tra cứu thông tin về các module liên quan.
Nếu không có file nào được chỉ định, Edit tìm kiếm file `go.work` trong thư mục
hiện tại và các thư mục cha của nó.

Các cờ chỉnh sửa chỉ định một chuỗi các thao tác chỉnh sửa.
* Cờ `-fmt` định dạng lại file go.work mà không thực hiện các thay đổi khác.
  Việc định dạng lại này cũng được ngụ ý bởi bất kỳ sửa đổi nào khác sử dụng hoặc
  viết lại file `go.work`. Cờ này chỉ cần thiết khi không có các cờ
  khác được chỉ định, như trong 'go work edit `-fmt`'.
* Các cờ `-use=path` và `-dropuse=path`
  thêm và bỏ một chỉ thị use từ tập hợp các thư mục module của file `go.work`.
* Cờ `-replace=old[@v]=new[@v]` thêm một thay thế của cặp đường dẫn
  module và phiên bản đã cho. Nếu `@v` trong `old@v` bị bỏ qua, một
  thay thế không có phiên bản ở phía bên trái được thêm vào, áp dụng
  cho tất cả các phiên bản của đường dẫn module cũ. Nếu `@v` trong `new@v` bị bỏ qua,
  đường dẫn mới phải là một thư mục gốc module cục bộ, không phải module
  path. Lưu ý rằng `-replace` ghi đè bất kỳ thay thế dư thừa nào cho `old[@v]`,
  vì vậy bỏ qua `@v` sẽ xóa các thay thế hiện có cho các phiên bản cụ thể.
* Cờ `-dropreplace=old[@v]` xóa một thay thế của cặp đường dẫn
  module và phiên bản đã cho. Nếu `@v` bị bỏ qua, một thay thế không có
  phiên bản ở phía bên trái bị xóa.
* Cờ `-go=version` đặt phiên bản ngôn ngữ Go dự kiến.

Các cờ chỉnh sửa có thể được lặp lại. Các thay đổi được áp dụng theo thứ tự đã cho.

`go work edit` có thêm các cờ kiểm soát đầu ra của nó.

* Cờ -print in file go.work cuối cùng ở định dạng văn bản thay vì
  ghi nó trở lại go.mod.
* Cờ -json in file go.work cuối cùng ở định dạng JSON thay vì
  ghi nó trở lại go.mod. Đầu ra JSON tương ứng với các kiểu Go sau:

```
type Module struct {
    Path    string
    Version string
}

type GoWork struct {
    Go        string
    Directory []Directory
    Replace   []Replace
}

type Use struct {
    Path       string
    ModulePath string
}

type Replace struct {
    Old Module
    New Module
}
```

### `go work use` {#go-work-use}

Cách dùng:

```
go work use [-r] [moddirs]
```

Lệnh `go work use` cung cấp giao diện dòng lệnh để thêm
các thư mục, tùy chọn đệ quy, vào file `go.work`.

Một [chỉ thị `use`](#go-work-file-use) sẽ được thêm vào file `go.work` cho mỗi thư mục đối số
được liệt kê trên dòng lệnh trong file `go.work`, nếu nó tồn tại trên đĩa,
hoặc bị xóa khỏi file `go.work` nếu nó không tồn tại trên đĩa.

Cờ `-r` tìm kiếm đệ quy các module trong các thư mục đối số
và lệnh use hoạt động như thể mỗi thư mục đó
được chỉ định như là đối số.

### `go work sync` {#go-work-sync}

Cách dùng:

```
go work sync
```

Lệnh `go work sync` đồng bộ danh sách build của workspace trở lại
các module của workspace.

Danh sách build của workspace là tập hợp các phiên bản của tất cả các
module dependency (bắc cầu) được sử dụng để thực hiện build trong workspace. `go
work sync` tạo ra danh sách build đó bằng cách sử dụng thuật toán [Lựa chọn Phiên bản Tối thiểu
(MVS)](#glos-minimal-version-selection)
, và sau đó đồng bộ các phiên bản đó trở lại từng module
được chỉ định trong workspace (với các chỉ thị `use`).

Sau khi danh sách build workspace được tính toán, file `go.mod` cho mỗi
module trong workspace được viết lại với các dependency liên quan
đến module đó được nâng cấp để khớp với danh sách build workspace.
Lưu ý rằng [Lựa chọn Phiên bản Tối thiểu](#glos-minimal-version-selection)
đảm bảo rằng phiên bản của mỗi module trong danh sách build luôn
bằng hoặc cao hơn phiên bản trong mỗi module workspace.

## Proxy module {#module-proxy}

### Giao thức `GOPROXY` {#goproxy-protocol}

Một <dfn>proxy module</dfn> là một máy chủ HTTP có thể phản hồi các yêu cầu `GET`
cho các đường dẫn được chỉ định bên dưới. Các yêu cầu không có tham số truy vấn và không
có header cụ thể nào được yêu cầu, vì vậy ngay cả một site phục vụ từ một hệ thống file cố định
(bao gồm cả URL `file://`) cũng có thể là một proxy module.

Các phản hồi HTTP thành công phải có mã trạng thái 200 (OK). Các chuyển hướng (3xx)
được theo dõi. Các phản hồi với mã trạng thái 4xx và 5xx được coi là lỗi.
Các mã lỗi 404 (Not Found) và 410 (Gone) chỉ ra rằng
module hoặc phiên bản được yêu cầu không có sẵn trên proxy, nhưng nó có thể được tìm thấy
ở nơi khác. Các phản hồi lỗi nên có loại nội dung `text/plain` với
`charset` là `utf-8` hoặc `us-ascii`.

Lệnh `go` có thể được cấu hình để liên hệ với proxy hoặc máy chủ kiểm soát phiên bản
bằng cách sử dụng biến môi trường `GOPROXY`, chấp nhận một danh sách URL proxy.
Danh sách có thể bao gồm các từ khóa `direct` hoặc `off` (xem [Biến
môi trường](#environment-variables) để biết chi tiết). Các phần tử trong danh sách có thể được phân cách
bằng dấu phẩy (`,`) hoặc ống dẫn (`|`), xác định hành vi dự phòng khi lỗi. Khi một
URL được theo sau bởi dấu phẩy, lệnh `go` chuyển sang các nguồn sau chỉ
sau phản hồi 404 (Not Found) hoặc 410 (Gone). Khi một URL được theo sau bởi một
ống dẫn, lệnh `go` chuyển sang các nguồn sau sau bất kỳ lỗi nào, bao gồm
các lỗi không phải HTTP như thời gian chờ. Hành vi xử lý lỗi này cho phép proxy hoạt động
như một gatekeeper cho các module không xác định. Ví dụ: một proxy có thể phản hồi với
lỗi 403 (Forbidden) cho các module không có trong danh sách đã được chấp thuận (xem [Proxy riêng
phục vụ các module riêng](#private-module-proxy-private)).

Bảng dưới đây chỉ định các truy vấn mà một proxy module phải phản hồi. Cho mỗi
đường dẫn, `$base` là phần đường dẫn của URL proxy, `$module` là đường dẫn module, và
`$version` là một phiên bản. Ví dụ: nếu URL proxy là
`https://example.com/mod`, và máy khách đang yêu cầu file `go.mod` cho
module `golang.org/x/text` ở phiên bản `v0.3.2`, máy khách sẽ gửi một
yêu cầu `GET` cho `https://example.com/mod/golang.org/x/text/@v/v0.3.2.mod`.

Để tránh sự mơ hồ khi phục vụ từ các hệ thống file không phân biệt chữ hoa chữ thường,
các phần tử `$module` và `$version` được mã hóa theo chữ hoa bằng cách thay thế mỗi
chữ hoa bằng dấu chấm than theo sau là chữ thường tương ứng. Điều này cho phép các module `example.com/M` và `example.com/m` cùng
được lưu trên đĩa, vì cái trước được mã hóa là `example.com/!m`.

<table class="ModTable">
  <thead>
    <tr>
      <th>Đường dẫn</th>
      <th>Mô tả</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>$base/$module/@v/list</code></td>
      <td>
        Trả về danh sách các phiên bản đã biết của module đã cho dưới dạng văn bản thuần túy, mỗi phiên bản
        một dòng. Danh sách này không nên bao gồm các pseudo-version.
      </td>
    </tr>
    <tr>
      <td><code>$base/$module/@v/$version.info</code></td>
      <td>
        <p>
          Trả về metadata định dạng JSON về một phiên bản cụ thể của một module.
          Phản hồi phải là một đối tượng JSON tương ứng với cấu trúc dữ liệu Go
          bên dưới:
        </p>
        <pre>
type Info struct {
    Version string    // version string
    Time    time.Time // commit time
}
</pre>
        <p>
          Trường <code>Version</code> là bắt buộc và phải chứa một
          <a href="#glos-canonical-version">phiên bản chuẩn</a> hợp lệ (xem
          <a href="#versions">Phiên bản</a>). <code>$version</code> trong
          đường dẫn yêu cầu không cần phải là cùng một phiên bản hoặc thậm chí là một phiên bản hợp lệ;
          endpoint này có thể được sử dụng để tìm phiên bản cho tên nhánh
          hoặc định danh revision. Tuy nhiên, nếu <code>$version</code> là một
          phiên bản chuẩn với phiên bản major tương thích với
          <code>$module</code>, trường <code>Version</code> trong một phản hồi thành công
          phải giống nhau.
        </p>
        <p>
          Trường <code>Time</code> là tùy chọn. Nếu có, nó phải là một
          chuỗi ở định dạng RFC 3339. Nó chỉ ra thời gian khi phiên bản
          được tạo ra.
        </p>
        <p>
          Nhiều trường hơn có thể được thêm vào trong tương lai, vì vậy các tên khác được dành riêng.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>$base/$module/@v/$version.mod</code></td>
      <td>
        Trả về file <code>go.mod</code> cho một phiên bản cụ thể của một
        module. Nếu module không có file <code>go.mod</code> ở
        phiên bản được yêu cầu, một file chỉ chứa câu lệnh <code>module</code>
        với đường dẫn module được yêu cầu phải được trả về. Ngược lại,
        file <code>go.mod</code> gốc, không được sửa đổi phải được trả về.
      </td>
    </tr>
    <tr>
      <td><code>$base/$module/@v/$version.zip</code></td>
      <td>
        Trả về một file zip chứa nội dung của một phiên bản cụ thể của
        một module. Xem <a href="#zip-files">Các file zip module</a> để biết chi tiết
        về cách file zip này phải được định dạng.
      </td>
    </tr>
    <tr>
      <td><code>$base/$module/@latest</code></td>
      <td>
        Trả về metadata định dạng JSON về phiên bản mới nhất đã biết của một
        module ở cùng định dạng với
        <code>$base/$module/@v/$version.info</code>. Phiên bản mới nhất nên
        là phiên bản của module mà lệnh <code>go</code> nên sử dụng
        nếu <code>$base/$module/@v/list</code> trống hoặc không có phiên bản được liệt kê nào phù hợp.
        Endpoint này là tùy chọn và các proxy module không bắt buộc
        phải thực hiện nó.
      </td>
    </tr>
  </tbody>
</table>

Khi phân giải phiên bản mới nhất của một module, lệnh `go` sẽ yêu cầu
`$base/$module/@v/list`, sau đó, nếu không tìm thấy phiên bản phù hợp,
`$base/$module/@latest`. Lệnh `go` ưu tiên, theo thứ tự: phiên bản release cao nhất về mặt semantic,
phiên bản pre-release cao nhất về mặt semantic, và pseudo-version gần đây nhất theo thứ tự thời gian. Trong Go 1.12 và trước đó, lệnh `go`
coi các pseudo-version trong `$base/$module/@v/list` là các phiên bản pre-release,
nhưng điều này không còn đúng kể từ Go 1.13.

Một proxy module phải luôn phục vụ cùng nội dung cho các phản hồi thành công cho
các truy vấn `$base/$module/$version.mod` và `$base/$module/$version.zip`.
Nội dung này được [xác thực mã hóa](#authenticating)
bằng cách sử dụng [các file `go.sum`](#go-sum-files) và, theo mặc định,
[cơ sở dữ liệu checksum](#checksum-database).

Lệnh `go` cache hầu hết nội dung nó tải từ proxy module trong
cache module của nó ở `$GOPATH/pkg/mod/cache/download`. Ngay cả khi tải trực tiếp
từ các hệ thống kiểm soát phiên bản, lệnh `go` tổng hợp các file `info`,
`mod` và `zip` rõ ràng và lưu trữ chúng trong thư mục này, giống như nó đã
tải chúng trực tiếp từ proxy. Bố cục cache giống với không gian URL proxy,
vì vậy việc phục vụ `$GOPATH/pkg/mod/cache/download` tại (hoặc sao chép nó vào)
`https://example.com/proxy` sẽ cho phép người dùng truy cập các phiên bản module đã cache bằng cách
đặt `GOPROXY` thành `https://example.com/proxy`.

### Giao tiếp với proxy {#communicating-with-proxies}

Lệnh `go` có thể tải mã nguồn module và metadata từ một [proxy
module](#glos-module-proxy). Biến [môi trường](#environment-variables) `GOPROXY`
có thể được sử dụng để cấu hình proxy nào mà lệnh `go` có thể kết nối và liệu nó có thể
giao tiếp trực tiếp với [các hệ thống kiểm soát phiên bản](#vcs) không. Dữ liệu module đã tải
được lưu trong [cache module](#glos-module-cache). Lệnh `go` sẽ chỉ liên hệ proxy khi
nó cần thông tin chưa có trong cache.

Phần [Giao thức `GOPROXY`](#goproxy-protocol) mô tả các yêu cầu có thể
được gửi đến máy chủ `GOPROXY`. Tuy nhiên, cũng hữu ích khi hiểu
khi nào lệnh `go` thực hiện các yêu cầu này. Ví dụ: `go build` thực hiện
quy trình bên dưới:

* Tính toán [danh sách build](#glos-build-list) bằng cách đọc các [file
  `go.mod`](#glos-go-mod-file) và thực hiện [lựa chọn phiên bản
  tối thiểu (MVS)](#glos-minimal-version-selection).
* Đọc các package được đặt tên trên dòng lệnh và các package mà chúng import.
* Nếu một package không được cung cấp bởi bất kỳ module nào trong danh sách build, hãy tìm một module
  cung cấp nó. Thêm một yêu cầu module về phiên bản mới nhất của nó vào `go.mod`,
  và bắt đầu lại.
* Build các package sau khi mọi thứ được tải.

Khi lệnh `go` tính toán danh sách build, nó tải file `go.mod` cho
mỗi module trong [đồ thị module](#glos-module-graph). Nếu file `go.mod` không có
trong cache, lệnh `go` sẽ tải nó từ proxy bằng cách sử dụng yêu cầu
`$module/@v/$version.mod` (trong đó `$module` là đường dẫn module và
`$version` là phiên bản). Các yêu cầu này có thể được kiểm thử bằng một công cụ như
`curl`. Ví dụ: lệnh dưới đây tải file `go.mod` cho
`golang.org/x/mod` ở phiên bản `v0.2.0`:

```
$ curl https://proxy.golang.org/golang.org/x/mod/@v/v0.2.0.mod
module golang.org/x/mod

go 1.12

require (
    golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
    golang.org/x/tools v0.0.0-20191119224855-298f0cb1881e
    golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898
)
```

Để tải một package, lệnh `go` cần mã nguồn cho
module cung cấp nó. Mã nguồn module được phân phối trong các file `.zip`
được giải nén vào cache module. Nếu file `.zip` module không có trong cache,
lệnh `go` sẽ tải nó bằng cách sử dụng yêu cầu `$module/@v/$version.zip`.

```
$ curl -O https://proxy.golang.org/golang.org/x/mod/@v/v0.2.0.zip
$ unzip -l v0.2.0.zip | head
Archive:  v0.2.0.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
     1479  00-00-1980 00:00   golang.org/x/mod@v0.2.0/LICENSE
     1303  00-00-1980 00:00   golang.org/x/mod@v0.2.0/PATENTS
      559  00-00-1980 00:00   golang.org/x/mod@v0.2.0/README
       21  00-00-1980 00:00   golang.org/x/mod@v0.2.0/codereview.cfg
      214  00-00-1980 00:00   golang.org/x/mod@v0.2.0/go.mod
     1476  00-00-1980 00:00   golang.org/x/mod@v0.2.0/go.sum
     5224  00-00-1980 00:00   golang.org/x/mod@v0.2.0/gosumcheck/main.go
```

Lưu ý rằng các yêu cầu `.mod` và `.zip` là riêng biệt, mặc dù các file `go.mod`
thường được chứa trong các file `.zip`. Lệnh `go` có thể cần tải
các file `go.mod` cho nhiều module khác nhau, và các file `.mod` nhỏ hơn nhiều
so với các file `.zip`. Ngoài ra, nếu một dự án Go không có file `go.mod`,
proxy sẽ phục vụ một file `go.mod` tổng hợp chỉ chứa một [chỉ thị
`module`](#go-mod-file-module). Các file `go.mod` tổng hợp được tạo bởi
lệnh `go` khi tải từ một [hệ thống kiểm soát phiên bản](#vcs).

Nếu lệnh `go` cần tải một package không được cung cấp bởi bất kỳ module nào trong
danh sách build, nó sẽ cố gắng tìm một module mới cung cấp nó. Phần
[Phân giải một package thành một module](#resolve-pkg-mod) mô tả quá trình này. Tóm
lại, lệnh `go` yêu cầu thông tin về phiên bản mới nhất của mỗi
đường dẫn module có thể chứa package. Ví dụ: cho
package `golang.org/x/net/html`, lệnh `go` sẽ cố gắng tìm các phiên bản mới nhất
của các module `golang.org/x/net/html`, `golang.org/x/net`,
`golang.org/x/` và `golang.org`. Chỉ có `golang.org/x/net` thực sự tồn tại và
cung cấp package đó, vì vậy lệnh `go` sử dụng phiên bản mới nhất của
module đó. Nếu nhiều hơn một module cung cấp package, lệnh `go` sẽ sử dụng
module có đường dẫn dài nhất.

Khi lệnh `go` yêu cầu phiên bản mới nhất của một module, nó trước tiên gửi một
yêu cầu cho `$module/@v/list`. Nếu danh sách trống hoặc không có phiên bản nào được trả về
có thể được sử dụng, nó gửi yêu cầu cho `$module/@latest`. Sau khi một phiên bản
được chọn, lệnh `go` gửi yêu cầu `$module/@v/$version.info` để
lấy metadata. Sau đó nó có thể gửi các yêu cầu `$module/@v/$version.mod` và
`$module/@v/$version.zip` để tải file `go.mod` và mã nguồn.

```
$ curl https://proxy.golang.org/golang.org/x/mod/@v/list
v0.1.0
v0.2.0

$ curl https://proxy.golang.org/golang.org/x/mod/@v/v0.2.0.info
{"Version":"v0.2.0","Time":"2020-01-02T17:33:45Z"}
```

Sau khi tải file `.mod` hoặc `.zip`, lệnh `go` tính toán một
băm mã hóa và kiểm tra rằng nó khớp với một băm trong file `go.sum` của module chính.
Nếu băm không có trong `go.sum`, theo mặc định, lệnh `go`
lấy nó từ [cơ sở dữ liệu checksum](#checksum-database). Nếu
băm đã tính toán không khớp, lệnh `go` báo lỗi bảo mật và không
cài đặt file trong cache module. Các biến môi trường `GOPRIVATE` và `GONOSUMDB`
[environment variables](#environment-variables) có thể được sử dụng để tắt các yêu cầu
đến cơ sở dữ liệu checksum cho các module cụ thể. Biến môi trường `GOSUMDB` cũng có thể
được đặt thành `off` để tắt hoàn toàn các yêu cầu đến cơ sở dữ liệu checksum.
Xem [Xác thực module](#authenticating) để biết thêm
thông tin. Lưu ý rằng danh sách phiên bản và metadata phiên bản được trả về cho các yêu cầu `.info`
không được xác thực và có thể thay đổi theo thời gian.

### Phục vụ module trực tiếp từ proxy {#serving-from-proxy}

Hầu hết các module được phát triển và phục vụ từ một kho lưu trữ kiểm soát phiên bản. Trong
[chế độ trực tiếp](#glos-direct-mode), lệnh `go` tải module như vậy bằng
một công cụ kiểm soát phiên bản (xem [Các hệ thống kiểm soát phiên bản](#vcs)). Cũng có thể
phục vụ một module trực tiếp từ một proxy module. Điều này hữu ích cho các tổ chức
muốn phục vụ các module mà không để lộ các máy chủ kiểm soát phiên bản của họ và
cho các tổ chức sử dụng các công cụ kiểm soát phiên bản mà lệnh `go` không
hỗ trợ.

Khi lệnh `go` tải một module ở chế độ trực tiếp, trước tiên nó tra cứu
URL của máy chủ module bằng yêu cầu HTTP GET dựa trên đường dẫn module. Nó tìm
kiếm thẻ `<meta>` với tên `go-import` trong phản hồi HTML. Nội dung của thẻ
phải chứa [đường dẫn gốc
kho lưu trữ](#glos-repository-root-path), hệ thống kiểm soát phiên bản và URL,
được phân cách bằng dấu cách. Xem [Tìm kho lưu trữ cho đường dẫn module](#vcs-find) để biết
chi tiết.

Nếu hệ thống kiểm soát phiên bản là `mod`, lệnh `go` tải module
từ URL đã cho bằng cách sử dụng [giao thức `GOPROXY`](#goproxy-protocol).

Ví dụ: giả sử lệnh `go` đang cố gắng tải module
`example.com/gopher` ở phiên bản `v1.0.0`. Nó gửi yêu cầu đến
`https://example.com/gopher?go-get=1`. Máy chủ phản hồi bằng một tài liệu HTML
chứa thẻ:

```
<meta name="go-import" content="example.com/gopher mod https://modproxy.example.com">
```

Dựa trên phản hồi này, lệnh `go` tải module bằng cách gửi
các yêu cầu cho `https://modproxy.example.com/example.com/gopher/@v/v1.0.0.info`,
`v1.0.0.mod` và `v1.0.0.zip`.

Lưu ý rằng các module được phục vụ trực tiếp từ proxy không thể được tải bằng
`go get` ở chế độ GOPATH.

## Các hệ thống kiểm soát phiên bản {#vcs}

Lệnh `go` có thể tải mã nguồn module và metadata trực tiếp từ một
kho lưu trữ kiểm soát phiên bản. Tải một module từ
[proxy](#communicating-with-proxies) thường nhanh hơn, nhưng kết nối trực tiếp
đến một kho lưu trữ là cần thiết nếu proxy không có sẵn hoặc nếu kho lưu trữ
của module không thể truy cập được bởi proxy (thường đúng với các
kho lưu trữ riêng). Git, Subversion, Mercurial, Bazaar và Fossil được hỗ trợ. Một
công cụ kiểm soát phiên bản phải được cài đặt trong một thư mục trong `PATH` để lệnh
`go` sử dụng nó.

Để tải các module cụ thể từ kho lưu trữ nguồn thay vì proxy, hãy đặt
biến môi trường `GOPRIVATE` hoặc `GONOPROXY`. Để cấu hình lệnh `go`
tải tất cả các module trực tiếp từ kho lưu trữ nguồn, hãy đặt `GOPROXY`
thành `direct`. Xem [Biến môi trường](#environment-variables) để biết thêm
thông tin.

### Tìm kho lưu trữ cho đường dẫn module {#vcs-find}

Khi lệnh `go` tải một module ở chế độ `direct`, nó bắt đầu bằng việc xác định
vị trí kho lưu trữ chứa module đó.

Nếu đường dẫn module có một bổ ngữ VCS (một trong `.bzr`, `.fossil`, `.git`, `.hg`,
`.svn`) ở cuối một thành phần đường dẫn, lệnh `go` sẽ sử dụng tất cả mọi thứ cho đến
bổ ngữ đường dẫn đó làm URL kho lưu trữ. Ví dụ: cho module
`example.com/foo.git/bar`, lệnh `go` tải kho lưu trữ tại
`example.com/foo` bằng git, mong đợi tìm thấy module trong thư mục con `bar`. Lệnh `go` sẽ đoán giao thức để sử dụng dựa trên
các giao thức được hỗ trợ bởi công cụ kiểm soát phiên bản.

Nếu đường dẫn module không có bổ ngữ, lệnh `go` gửi một yêu cầu HTTP
`GET` đến một URL được lấy từ đường dẫn module với chuỗi truy vấn `?go-get=1`. Ví dụ: cho module `golang.org/x/mod`, lệnh `go` có thể
gửi các yêu cầu sau:

```
https://golang.org/x/mod?go-get=1 (preferred)
http://golang.org/x/mod?go-get=1  (fallback, only with GOINSECURE)
```

Lệnh `go` theo dõi các chuyển hướng nhưng nếu không bỏ qua các mã trạng thái phản hồi,
vì vậy máy chủ có thể phản hồi với 404 hoặc bất kỳ trạng thái lỗi nào khác. Biến môi trường
`GOINSECURE` có thể được đặt để cho phép dự phòng và chuyển hướng đến
HTTP không mã hóa cho các module cụ thể.

Máy chủ phải phản hồi bằng một tài liệu HTML chứa thẻ `<meta>` trong
`<head>` của tài liệu. Thẻ `<meta>` nên xuất hiện sớm trong tài liệu để
tránh gây nhầm lẫn cho trình phân tích cú pháp hạn chế của lệnh `go`. Đặc biệt, nó nên
xuất hiện trước bất kỳ JavaScript thô nào hoặc CSS. Thẻ `<meta>` phải có dạng:

```
<meta name="go-import" content="root-path vcs repo-url [subdirectory]">
```

`root-path` là đường dẫn gốc của kho lưu trữ, tức là phần trong đường dẫn module tương ứng với thư mục gốc của kho lưu trữ, hoặc tương ứng với `subdirectory` nếu có và đang dùng Go 1.25 trở lên (xem phần về `subdirectory` bên dưới). Giá trị này phải là tiền tố hoặc khớp chính xác với đường dẫn module được yêu cầu. Nếu không khớp chính xác, một yêu cầu khác sẽ được gửi đến tiền tố đó để xác minh các thẻ `<meta>` khớp nhau.

`vcs` là hệ thống quản lý phiên bản. Giá trị phải là một trong các công cụ liệt kê trong bảng bên dưới, hoặc từ khóa `mod` để yêu cầu lệnh `go` tải module từ URL đã cho bằng [giao thức `GOPROXY`](#goproxy-protocol). Xem [Phục vụ module trực tiếp từ proxy](#serving-from-proxy) để biết thêm chi tiết.

`repo-url` là URL của kho lưu trữ, bao gồm scheme và không chứa hậu tố `.vcs`. Các giao thức không bảo mật (như `http://` và `git://`) chỉ được dùng nếu đường dẫn module khớp với biến môi trường `GOINSECURE`.

`subdirectory`, nếu có, là thư mục con của kho lưu trữ (phân cách bằng dấu gạch chéo) mà `root-path` tương ứng, ghi đè mặc định là thư mục gốc của kho lưu trữ. Thẻ meta `go-import` cung cấp `subdirectory` chỉ được nhận diện từ Go 1.25 trở lên. Khi cố gắng phân giải module trên các phiên bản Go cũ hơn, thẻ meta này sẽ bị bỏ qua và quá trình phân giải sẽ thất bại nếu module không thể tìm thấy ở nơi khác.

<table id="vcs-support" class="ModTable">
  <thead>
    <tr>
      <th>Name</th>
      <th>Command</th>
      <th>GOVCS default</th>
      <th>Secure schemes</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Bazaar</td>
      <td><code>bzr</code></td>
      <td>Private only</td>
      <td><code>https</code>, <code>bzr+ssh</code></td>
    </tr>
    <tr>
      <td>Fossil</td>
      <td><code>fossil</code></td>
      <td>Private only</td>
      <td><code>https</code></td>
    </tr>
    <tr>
      <td>Git</td>
      <td><code>git</code></td>
      <td>Public and private</td>
      <td><code>https</code>, <code>git+ssh</code>, <code>ssh</code></td>
    </tr>
    <tr>
      <td>Mercurial</td>
      <td><code>hg</code></td>
      <td>Public and private</td>
      <td><code>https</code>, <code>ssh</code></td>
    </tr>
    <tr>
      <td>Subversion</td>
      <td><code>svn</code></td>
      <td>Private only</td>
      <td><code>https</code>, <code>svn+ssh</code></td>
    </tr>
  </tbody>
</table>

Ví dụ, xét lại `golang.org/x/mod`. Lệnh `go` gửi yêu cầu đến `https://golang.org/x/mod?go-get=1`. Máy chủ phản hồi bằng một tài liệu HTML chứa thẻ:

```
<meta name="go-import" content="golang.org/x/mod git https://go.googlesource.com/mod">
```

Từ phản hồi này, lệnh `go` sẽ sử dụng kho Git tại URL `https://go.googlesource.com/mod`.

GitHub và các dịch vụ lưu trữ phổ biến khác đều phản hồi các truy vấn `?go-get=1` cho mọi kho lưu trữ, nên thường không cần cấu hình máy chủ thêm cho các module được lưu trữ tại đó.

Sau khi tìm được URL kho lưu trữ, lệnh `go` sẽ clone kho lưu trữ vào bộ đệm module. Thông thường, lệnh `go` cố tránh tải dữ liệu không cần thiết từ kho lưu trữ. Tuy nhiên, các lệnh thực tế được sử dụng khác nhau tùy theo hệ thống quản lý phiên bản và có thể thay đổi theo thời gian. Với Git, lệnh `go` có thể liệt kê hầu hết các phiên bản hiện có mà không cần tải commit. Thường thì lệnh sẽ tải commit mà không tải các commit tổ tiên, nhưng đôi khi điều đó vẫn cần thiết.

### Ánh xạ phiên bản sang commit {#vcs-version}

Lệnh `go` có thể checkout một module trong kho lưu trữ ở một [phiên bản canonical](#glos-canonical-version) cụ thể như `v1.2.3`, `v2.4.0-beta`, hoặc `v3.0.0+incompatible`. Mỗi phiên bản module phải có một <dfn>thẻ semantic version</dfn> trong kho lưu trữ để chỉ ra revision nào cần được checkout cho phiên bản đó.

Nếu module được định nghĩa trong thư mục gốc của kho lưu trữ hoặc trong một thư mục con major version của thư mục gốc, thì mỗi tên thẻ phiên bản bằng với phiên bản tương ứng. Ví dụ, module `golang.org/x/text` được định nghĩa trong thư mục gốc của kho lưu trữ, nên phiên bản `v0.3.2` có thẻ `v0.3.2` trong kho lưu trữ đó. Điều này đúng với hầu hết các module.

Nếu module được định nghĩa trong một thư mục con của kho lưu trữ, tức là phần [thư mục con module](#glos-module-subdirectory) trong đường dẫn module không rỗng, thì mỗi tên thẻ phải có tiền tố là thư mục con của module, theo sau là dấu gạch chéo. Ví dụ, module `golang.org/x/tools/gopls` được định nghĩa trong thư mục con `gopls` của kho lưu trữ có đường dẫn gốc `golang.org/x/tools`. Phiên bản `v0.4.0` của module đó phải có thẻ tên `gopls/v0.4.0` trong kho lưu trữ đó.

Số major version trong thẻ semantic version phải nhất quán với hậu tố major version của đường dẫn module (nếu có). Ví dụ, thẻ `v1.0.0` có thể thuộc module `example.com/mod` nhưng không thuộc `example.com/mod/v2`, vì module đó sẽ có các thẻ như `v2.0.0`.

Một thẻ có major version `v2` trở lên có thể thuộc một module không có hậu tố major version nếu không có tệp `go.mod` và module nằm trong thư mục gốc của kho lưu trữ. Loại phiên bản này được ký hiệu với hậu tố `+incompatible`. Bản thân thẻ phiên bản không được có hậu tố này. Xem [Tương thích với các kho lưu trữ không dùng module](#non-module-compat).

Sau khi một thẻ được tạo, không nên xóa hoặc thay đổi nó sang revision khác. Các phiên bản được [xác thực](#authenticating) để đảm bảo các bản dựng an toàn và có thể tái tạo. Nếu thẻ bị sửa đổi, người dùng có thể gặp lỗi bảo mật khi tải xuống. Ngay cả sau khi thẻ bị xóa, nội dung của nó vẫn có thể còn trên các [module proxy](#glos-module-proxy).

### Ánh xạ pseudo-version sang commit {#vcs-pseudo}

Lệnh `go` có thể checkout một module trong kho lưu trữ ở một revision cụ thể, được mã hóa dưới dạng [pseudo-version](#glos-pseudo-version) như `v1.3.2-0.20191109021931-daa7c04131f5`.

12 ký tự cuối của pseudo-version (`daa7c04131f5` trong ví dụ trên) chỉ ra revision trong kho lưu trữ cần checkout. Ý nghĩa của chuỗi này tùy thuộc vào hệ thống quản lý phiên bản. Với Git và Mercurial, đây là tiền tố của commit hash. Với Subversion, đây là số revision được đệm bằng số không.

Trước khi checkout một commit, lệnh `go` xác minh rằng timestamp (`20191109021931` trong ví dụ trên) khớp với ngày commit. Lệnh cũng xác minh rằng phiên bản cơ sở (`v1.3.1`, phiên bản trước `v1.3.2` trong ví dụ trên) tương ứng với một thẻ semantic version là tổ tiên của commit đó. Các kiểm tra này đảm bảo tác giả module có toàn quyền kiểm soát cách pseudo-version so sánh với các phiên bản đã phát hành khác.

Xem [Pseudo-version](#pseudo-versions) để biết thêm thông tin.

### Ánh xạ nhánh và commit sang phiên bản {#vcs-branch}

Một module có thể được checkout ở một nhánh, thẻ hoặc revision cụ thể bằng cách dùng [truy vấn phiên bản](#version-queries).

```
go get example.com/mod@master
```

Lệnh `go` chuyển đổi các tên này thành [phiên bản canonical](#glos-canonical-version) có thể được dùng với [minimal version selection (MVS)](#minimal-version-selection). MVS phụ thuộc vào khả năng sắp xếp phiên bản theo thứ tự rõ ràng. Tên nhánh và revision không thể so sánh đáng tin cậy theo thời gian vì chúng phụ thuộc vào cấu trúc kho lưu trữ có thể thay đổi.

Nếu một revision được gắn thẻ với một hoặc nhiều thẻ semantic version như `v1.2.3`, thẻ có phiên bản hợp lệ cao nhất sẽ được dùng. Lệnh `go` chỉ xét các thẻ semantic version có thể thuộc về module đích; ví dụ, thẻ `v1.5.2` sẽ không được xét cho `example.com/mod/v2` vì major version không khớp với hậu tố đường dẫn module.

Nếu một revision không được gắn thẻ semantic version hợp lệ, lệnh `go` sẽ tạo một [pseudo-version](#glos-pseudo-version). Nếu revision đó có tổ tiên với thẻ semantic version hợp lệ, phiên bản tổ tiên cao nhất sẽ được dùng làm cơ sở cho pseudo-version. Xem [Pseudo-version](#pseudo-versions).

### Thư mục module trong kho lưu trữ {#vcs-dir}

Sau khi kho lưu trữ của module được checkout ở một revision cụ thể, lệnh `go` phải xác định thư mục chứa tệp `go.mod` của module (thư mục gốc của module).

Nhắc lại rằng một [đường dẫn module](#module-path) gồm ba phần: đường dẫn gốc kho lưu trữ (tương ứng với thư mục gốc kho lưu trữ), thư mục con module, và hậu tố major version (chỉ dành cho module phát hành ở `v2` trở lên).

Với hầu hết các module, đường dẫn module bằng với đường dẫn gốc kho lưu trữ, nên thư mục gốc của module chính là thư mục gốc của kho lưu trữ.

Module đôi khi được định nghĩa trong các thư mục con của kho lưu trữ. Điều này thường xảy ra với các kho lưu trữ lớn có nhiều thành phần cần được phát hành và tạo phiên bản độc lập. Module như vậy được tìm thấy trong thư mục con khớp với phần đường dẫn module sau đường dẫn gốc kho lưu trữ. Ví dụ, giả sử module `example.com/monorepo/foo/bar` nằm trong kho lưu trữ có đường dẫn gốc `example.com/monorepo`. Tệp `go.mod` của nó phải nằm trong thư mục con `foo/bar`.

Nếu module được phát hành ở major version `v2` trở lên, đường dẫn của nó phải có [hậu tố major version](#major-version-suffixes). Module có hậu tố major version có thể được định nghĩa trong một trong hai thư mục con: một thư mục có hậu tố và một thư mục không có. Ví dụ, giả sử phiên bản mới của module trên được phát hành với đường dẫn `example.com/monorepo/foo/bar/v2`. Tệp `go.mod` của nó có thể nằm trong `foo/bar` hoặc `foo/bar/v2`.

Các thư mục con có hậu tố major version được gọi là <dfn>thư mục con major version</dfn>. Chúng có thể được dùng để phát triển nhiều major version của một module trên cùng một nhánh. Điều này có thể không cần thiết khi việc phát triển nhiều major version diễn ra trên các nhánh riêng biệt. Tuy nhiên, thư mục con major version có một thuộc tính quan trọng: trong chế độ `GOPATH`, đường dẫn import của package khớp chính xác với các thư mục bên dưới `GOPATH/src`. Lệnh `go` cung cấp khả năng tương thích module tối thiểu trong chế độ `GOPATH` (xem [Tương thích với các kho lưu trữ không dùng module](#non-module-compat)), nên thư mục con major version không phải lúc nào cũng cần thiết để tương thích với các dự án được dựng trong chế độ `GOPATH`. Tuy nhiên, các công cụ cũ không hỗ trợ khả năng tương thích module tối thiểu có thể gặp vấn đề.

Sau khi lệnh `go` tìm được thư mục gốc của module, lệnh tạo một tệp `.zip` từ nội dung thư mục đó, rồi giải nén tệp `.zip` vào bộ đệm module. Xem [Ràng buộc về đường dẫn tệp và kích thước](#zip-path-size-constraints) để biết chi tiết về các tệp có thể được đưa vào tệp `.zip`. Nội dung của tệp `.zip` được [xác thực](#authenticating) trước khi giải nén vào bộ đệm module theo cách tương tự như khi tệp `.zip` được tải từ proxy.

Các tệp zip module không bao gồm nội dung của thư mục `vendor` hoặc bất kỳ module lồng nhau nào (các thư mục con chứa tệp `go.mod`). Điều này có nghĩa là module phải tránh tham chiếu đến các tệp bên ngoài thư mục của nó hoặc trong các module khác. Ví dụ, các mẫu [`//go:embed`](https://pkg.go.dev/embed#hdr-Directives) không được khớp với các tệp trong module lồng nhau. Hành vi này có thể hữu ích trong những tình huống các tệp không nên được đưa vào module. Ví dụ, nếu một kho lưu trữ có các tệp lớn được lưu trong thư mục `testdata`, tác giả module có thể thêm một tệp `go.mod` rỗng vào `testdata` để người dùng không cần tải các tệp đó. Tất nhiên, điều này có thể làm giảm độ phủ kiểm thử cho người dùng khi kiểm tra các dependency của họ.

### Trường hợp đặc biệt với tệp LICENSE {#vcs-license}

Khi lệnh `go` tạo tệp `.zip` cho một module không nằm trong thư mục gốc của kho lưu trữ, nếu module không có tệp `LICENSE` trong thư mục gốc của nó (bên cạnh `go.mod`), lệnh `go` sẽ sao chép tệp `LICENSE` từ thư mục gốc kho lưu trữ nếu tệp đó tồn tại ở cùng revision.

Trường hợp đặc biệt này cho phép cùng một tệp `LICENSE` áp dụng cho tất cả các module trong một kho lưu trữ. Điều này chỉ áp dụng cho các tệp có tên chính xác là `LICENSE`, không có phần mở rộng như `.txt`. Rất tiếc, không thể mở rộng điều này mà không phá vỡ các tổng kiểm tra mã hóa của các module hiện có; xem [Xác thực module](#authenticating). Các công cụ và trang web khác như [pkg.go.dev](https://pkg.go.dev) có thể nhận diện các tệp có tên khác.

Lưu ý rằng lệnh `go` không đưa các liên kết tượng trưng vào tệp `.zip` module; xem [Ràng buộc về đường dẫn tệp và kích thước](#zip-path-size-constraints). Do đó, nếu kho lưu trữ không có tệp `LICENSE` trong thư mục gốc, tác giả có thể tạo các bản sao tệp giấy phép trong các module được định nghĩa trong thư mục con để đảm bảo các tệp đó được đưa vào tệp `.zip` module.

### Kiểm soát công cụ quản lý phiên bản bằng `GOVCS` {#vcs-govcs}

Khả năng tải module bằng các lệnh quản lý phiên bản như `git` của lệnh `go` rất quan trọng với hệ sinh thái package phi tập trung, trong đó mã nguồn có thể được import từ bất kỳ máy chủ nào. Đây cũng là một rủi ro bảo mật tiềm ẩn nếu một máy chủ độc hại tìm cách khiến lệnh quản lý phiên bản được gọi chạy mã ngoài ý muốn.

Để cân bằng giữa tính năng và bảo mật, mặc định lệnh `go` chỉ dùng `git` và `hg` để tải mã từ các máy chủ công khai. Lệnh sẽ dùng bất kỳ [hệ thống quản lý phiên bản đã biết](#vcs-support) nào để tải mã từ các máy chủ riêng tư, được định nghĩa là các máy chủ lưu trữ package khớp với biến môi trường `GOPRIVATE`. Lý do chỉ cho phép Git và Mercurial là vì hai hệ thống này được chú ý nhiều nhất đến các vấn đề khi chạy như client của các máy chủ không tin cậy. Ngược lại, Bazaar, Fossil và Subversion chủ yếu được dùng trong môi trường đáng tin cậy, được xác thực và không được kiểm tra kỹ lưỡng như một bề mặt tấn công.

Các hạn chế về lệnh quản lý phiên bản chỉ áp dụng khi dùng truy cập trực tiếp vào hệ thống quản lý phiên bản để tải mã. Khi tải module từ proxy, lệnh `go` dùng [giao thức `GOPROXY`](#goproxy-protocol) thay thế, và điều này luôn được cho phép. Mặc định, lệnh `go` dùng Go module mirror ([proxy.golang.org](https://proxy.golang.org)) cho các module công khai và chỉ dùng quản lý phiên bản trực tiếp cho các module riêng tư hoặc khi mirror từ chối phục vụ một package công khai (thường vì lý do pháp lý). Do đó, người dùng vẫn có thể truy cập mã công khai từ các kho lưu trữ Bazaar, Fossil hoặc Subversion theo mặc định, vì các lượt tải đó dùng Go module mirror, nơi chịu rủi ro bảo mật khi chạy các lệnh quản lý phiên bản trong một sandbox tùy chỉnh.

Biến `GOVCS` có thể được dùng để thay đổi các hệ thống quản lý phiên bản được phép dùng cho các module cụ thể. Biến `GOVCS` được áp dụng khi dựng package trong cả chế độ module-aware lẫn chế độ GOPATH. Khi dùng module, các mẫu so khớp với đường dẫn module. Khi dùng GOPATH, các mẫu so khớp với đường dẫn import tương ứng với thư mục gốc của kho lưu trữ quản lý phiên bản.

Dạng chung của biến `GOVCS` là một danh sách các quy tắc `pattern:vcslist` phân cách bằng dấu phẩy. Pattern là một [mẫu glob](/pkg/path#Match) phải khớp với một hoặc nhiều phần đầu của đường dẫn module hoặc import. vcslist là danh sách các lệnh quản lý phiên bản được phép phân cách bằng dấu pipe, hoặc `all` để cho phép dùng bất kỳ lệnh đã biết nào, hoặc `off` để cấm tất cả. Lưu ý rằng nếu một module khớp với một mẫu có vcslist là `off`, nó vẫn có thể được tải nếu máy chủ gốc dùng scheme `mod`, yêu cầu lệnh go tải module bằng [giao thức `GOPROXY`](#goproxy-protocol). Mẫu khớp đầu tiên trong danh sách được áp dụng, ngay cả khi các mẫu sau cũng có thể khớp.

Ví dụ, xét:

```
GOVCS=github.com:git,evil.com:off,*:git|hg
```

Với cài đặt này, mã có đường dẫn module hoặc import bắt đầu bằng `github.com/` chỉ có thể dùng `git`; các đường dẫn trên `evil.com` không được dùng bất kỳ lệnh quản lý phiên bản nào, và tất cả các đường dẫn khác (`*` khớp với tất cả) chỉ có thể dùng `git` hoặc `hg`.

Các mẫu đặc biệt `public` và `private` khớp với các đường dẫn module hoặc import công khai và riêng tư. Một đường dẫn là riêng tư nếu nó khớp với biến `GOPRIVATE`; ngược lại là công khai.

Nếu không có quy tắc nào trong biến `GOVCS` khớp với một module hoặc đường dẫn import cụ thể, lệnh `go` áp dụng quy tắc mặc định, có thể được tóm tắt trong ký hiệu `GOVCS` là `public:git|hg,private:all`.

Để cho phép dùng bất kỳ hệ thống quản lý phiên bản nào cho bất kỳ package nào, dùng:

```
GOVCS=*:all
```

Để tắt hoàn toàn việc dùng quản lý phiên bản, dùng:

```
GOVCS=*:off
```

Lệnh [`go env -w`
](/cmd/go/#hdr-Print_Go_environment_information) có thể được dùng để đặt biến `GOVCS` cho các lần gọi lệnh go trong tương lai.

`GOVCS` được giới thiệu từ Go 1.16. Các phiên bản Go cũ hơn có thể dùng bất kỳ công cụ quản lý phiên bản đã biết nào cho bất kỳ module nào.

## Tệp zip module {#zip-files}

Các phiên bản module được phân phối dưới dạng tệp `.zip`. Hiếm khi cần tương tác trực tiếp với các tệp này, vì lệnh `go` tự động tạo, tải và giải nén chúng từ các [module proxy](#glos-module-proxy) và kho lưu trữ quản lý phiên bản. Tuy nhiên, việc hiểu về các tệp này vẫn hữu ích để nắm các ràng buộc về tương thích đa nền tảng hoặc khi triển khai một module proxy.

Lệnh [`go mod download`](#go-mod-download) tải tệp zip cho một hoặc nhiều module, rồi giải nén các tệp đó vào [bộ đệm module](#glos-module-cache). Tùy thuộc vào `GOPROXY` và các [biến môi trường](#environment-variables) khác, lệnh `go` có thể tải tệp zip từ proxy hoặc clone các kho lưu trữ mã nguồn và tạo tệp zip từ chúng. Cờ `-json` có thể được dùng để tìm vị trí của các tệp zip đã tải và nội dung đã giải nén của chúng trong bộ đệm module.

Package [`golang.org/x/mod/zip`](https://pkg.go.dev/golang.org/x/mod/zip?tab=doc) có thể được dùng để tạo, giải nén hoặc kiểm tra nội dung tệp zip theo cách lập trình.

### Ràng buộc về đường dẫn tệp và kích thước {#zip-path-size-constraints}

Có một số hạn chế về nội dung của các tệp zip module. Các ràng buộc này đảm bảo tệp zip có thể được giải nén an toàn và nhất quán trên nhiều nền tảng khác nhau.

* Tệp zip module có thể có kích thước tối đa 500 MiB. Tổng kích thước không nén của các tệp trong đó cũng bị giới hạn 500 MiB. Tệp `go.mod` bị giới hạn 16 MiB. Tệp `LICENSE` cũng bị giới hạn 16 MiB. Các giới hạn này nhằm giảm thiểu các cuộc tấn công từ chối dịch vụ nhắm vào người dùng, proxy và các thành phần khác của hệ sinh thái module. Các kho lưu trữ chứa hơn 500 MiB tệp trong cây thư mục module nên gắn thẻ phiên bản module tại các commit chỉ bao gồm các tệp cần thiết để dựng các package của module; video, model và các tài nguyên lớn khác thường không cần thiết cho việc dựng.
* Mỗi tệp trong tệp zip module phải bắt đầu bằng tiền tố `$module@$version/` trong đó `$module` là đường dẫn module và `$version` là phiên bản, ví dụ `golang.org/x/mod@v0.3.0/`. Đường dẫn module phải hợp lệ, phiên bản phải hợp lệ và canonical, và phiên bản phải khớp với hậu tố major version của đường dẫn module. Xem [Đường dẫn và phiên bản module](#go-mod-file-ident) để biết các định nghĩa và ràng buộc cụ thể.
* Chế độ tệp, timestamp và các metadata khác bị bỏ qua.
* Các thư mục rỗng (các mục có đường dẫn kết thúc bằng dấu gạch chéo) có thể được đưa vào tệp zip module nhưng không được giải nén. Lệnh `go` không đưa các thư mục rỗng vào tệp zip khi tạo.
* Các liên kết tượng trưng và các tệp không thông thường khác bị bỏ qua khi tạo tệp zip, vì chúng không di chuyển được giữa các hệ điều hành và hệ thống tệp, và không có cách di chuyển được để biểu diễn chúng trong định dạng tệp zip.
* Các tệp trong thư mục có tên `vendor` bị bỏ qua khi tạo tệp zip, vì các thư mục `vendor` bên ngoài module chính không bao giờ được dùng.
* Các tệp trong thư mục chứa tệp `go.mod`, ngoài thư mục gốc của module, bị bỏ qua khi tạo tệp zip, vì chúng không thuộc module đó. Lệnh `go` bỏ qua các thư mục con chứa tệp `go.mod` khi giải nén tệp zip.
* Không có hai tệp nào trong tệp zip được có đường dẫn bằng nhau khi so sánh theo Unicode case-folding (xem [`strings.EqualFold`](https://pkg.go.dev/strings?tab=doc#EqualFold)). Điều này đảm bảo tệp zip có thể được giải nén trên các hệ thống tệp không phân biệt chữ hoa chữ thường mà không xảy ra xung đột.
* Tệp `go.mod` có thể có hoặc không có trong thư mục cấp cao nhất (`$module@$version/go.mod`). Nếu có, tên phải là `go.mod` (toàn chữ thường). Các tệp tên `go.mod` không được phép ở bất kỳ thư mục nào khác.
* Tên tệp và thư mục trong module có thể gồm các chữ cái Unicode, chữ số ASCII, ký tự khoảng trắng ASCII (U+0020), và các ký tự dấu câu ASCII `!#$%&()+,-.=@[]^_{}~`. Lưu ý rằng đường dẫn package có thể không chứa tất cả các ký tự này. Xem [`module.CheckFilePath`](https://pkg.go.dev/golang.org/x/mod/module?tab=doc#CheckFilePath) và [`module.CheckImportPath`](https://pkg.go.dev/golang.org/x/mod/module?tab=doc#CheckImportPath) để biết sự khác biệt.
* Tên tệp hoặc thư mục tính đến dấu chấm đầu tiên không được là tên tệp đặc biệt trên Windows, bất kể chữ hoa hay chữ thường (`CON`, `com1`, `NuL`, v.v.).

## Module riêng tư {#private-modules}

Các module Go thường được phát triển và phân phối trên các máy chủ quản lý phiên bản và module proxy không có sẵn trên internet công khai. Lệnh `go` có thể tải và dựng module từ các nguồn riêng tư, dù thường cần một số cấu hình.

Các biến môi trường bên dưới có thể được dùng để cấu hình truy cập vào các module riêng tư. Xem [Biến môi trường](#environment-variables) để biết chi tiết. Xem thêm [Quyền riêng tư](#private-module-privacy) để biết thông tin về việc kiểm soát thông tin gửi đến các máy chủ công khai.

* `GOPROXY` — danh sách URL của module proxy. Lệnh `go` sẽ cố tải module từ từng máy chủ theo thứ tự. Từ khóa `direct` yêu cầu lệnh `go` tải module trực tiếp từ kho lưu trữ quản lý phiên bản nơi chúng được phát triển thay vì dùng proxy.
* `GOPRIVATE` — danh sách các mẫu glob về tiền tố đường dẫn module được coi là riêng tư. Đóng vai trò giá trị mặc định cho `GONOPROXY` và `GONOSUMDB`.
* `GONOPROXY` — danh sách các mẫu glob về tiền tố đường dẫn module không nên tải từ proxy. Lệnh `go` sẽ tải các module khớp trực tiếp từ kho lưu trữ quản lý phiên bản nơi chúng được phát triển, bất kể `GOPROXY`.
* `GONOSUMDB` — danh sách các mẫu glob về tiền tố đường dẫn module không nên kiểm tra bằng cơ sở dữ liệu checksum công khai [sum.golang.org](https://sum.golang.org).
* `GOINSECURE` — danh sách các mẫu glob về tiền tố đường dẫn module có thể được tải qua HTTP và các giao thức không bảo mật khác.

Các biến này có thể được đặt trong môi trường phát triển (ví dụ trong tệp `.profile`), hoặc đặt vĩnh viễn bằng [`go env -w`](/cmd/go/#hdr-Print_Go_environment_information).

Phần còn lại của mục này mô tả các mẫu phổ biến để cung cấp quyền truy cập vào các module proxy riêng tư và kho lưu trữ quản lý phiên bản.

### Proxy riêng tư phục vụ tất cả module {#private-module-proxy-all}

Một máy chủ proxy riêng tư trung tâm phục vụ tất cả module (công khai và riêng tư) mang lại khả năng kiểm soát cao nhất cho quản trị viên và đòi hỏi cấu hình ít nhất từ phía các nhà phát triển.

Để cấu hình lệnh `go` sử dụng máy chủ như vậy, đặt các biến môi trường sau, thay `https://proxy.corp.example.com` bằng URL proxy của bạn và `corp.example.com` bằng tiền tố module của bạn:

```
GOPROXY=https://proxy.corp.example.com
GONOSUMDB=corp.example.com
```

Cài đặt `GOPROXY` yêu cầu lệnh `go` chỉ tải module từ `https://proxy.corp.example.com`; lệnh `go` sẽ không kết nối đến các proxy hoặc kho lưu trữ quản lý phiên bản khác.

Cài đặt `GONOSUMDB` yêu cầu lệnh `go` không dùng cơ sở dữ liệu checksum công khai để xác thực các module có đường dẫn bắt đầu bằng `corp.example.com`.

Proxy chạy với cấu hình này có thể cần quyền đọc trên các máy chủ quản lý phiên bản riêng tư. Nó cũng cần truy cập internet công khai để tải các phiên bản mới của các module công khai.

Có một số triển khai máy chủ `GOPROXY` hiện có có thể dùng theo cách này. Một triển khai tối giản sẽ phục vụ các tệp từ thư mục [bộ đệm module](#glos-module-cache) và dùng [`go mod download`](#go-mod-download) (với cấu hình phù hợp) để lấy các module còn thiếu.

### Proxy riêng tư chỉ phục vụ module riêng tư {#private-module-proxy-private}

Một máy chủ proxy riêng tư có thể phục vụ các module riêng tư mà không phục vụ các module công khai. Lệnh `go` có thể được cấu hình để dùng các nguồn công khai khi cần cho các module không có trên máy chủ riêng tư.

Để cấu hình lệnh `go` hoạt động theo cách này, đặt các biến môi trường sau, thay `https://proxy.corp.example.com` bằng URL proxy và `corp.example.com` bằng tiền tố module:

```
GOPROXY=https://proxy.corp.example.com,https://proxy.golang.org,direct
GONOSUMDB=corp.example.com
```

Cài đặt `GOPROXY` yêu cầu lệnh `go` thử tải module từ `https://proxy.corp.example.com` trước. Nếu máy chủ đó phản hồi 404 (Not Found) hoặc 410 (Gone), lệnh `go` sẽ chuyển sang `https://proxy.golang.org`, rồi đến kết nối trực tiếp với kho lưu trữ.

Cài đặt `GONOSUMDB` yêu cầu lệnh `go` không dùng cơ sở dữ liệu checksum công khai để xác thực các module có đường dẫn bắt đầu bằng `corp.example.com`.

Lưu ý rằng proxy được dùng trong cấu hình này vẫn có thể kiểm soát quyền truy cập vào các module công khai, ngay cả khi nó không phục vụ chúng. Nếu proxy phản hồi một yêu cầu với mã lỗi khác 404 hoặc 410, lệnh `go` sẽ không chuyển sang các mục sau trong danh sách `GOPROXY`. Ví dụ, proxy có thể phản hồi 403 (Forbidden) cho một module có giấy phép không phù hợp hoặc có các lỗ hổng bảo mật đã biết.

### Truy cập trực tiếp vào module riêng tư {#private-module-proxy-direct}

Lệnh `go` có thể được cấu hình để bỏ qua các proxy công khai và tải các module riêng tư trực tiếp từ các máy chủ quản lý phiên bản. Điều này hữu ích khi không thể chạy máy chủ proxy riêng tư.

Để cấu hình lệnh `go` hoạt động theo cách này, đặt `GOPRIVATE`, thay `corp.example.com` bằng tiền tố module riêng tư:

```
GOPRIVATE=corp.example.com
```

Biến `GOPROXY` không cần thay đổi trong trường hợp này. Giá trị mặc định là `https://proxy.golang.org,direct`, yêu cầu lệnh `go` thử tải module từ `https://proxy.golang.org` trước, rồi chuyển sang kết nối trực tiếp nếu proxy đó phản hồi 404 (Not Found) hoặc 410 (Gone).

Cài đặt `GOPRIVATE` yêu cầu lệnh `go` không kết nối đến proxy hoặc cơ sở dữ liệu checksum cho các module bắt đầu bằng `corp.example.com`.

Một máy chủ HTTP nội bộ vẫn có thể cần thiết để [phân giải đường dẫn module thành URL kho lưu trữ](#vcs-find). Ví dụ, khi lệnh `go` tải module `corp.example.com/mod`, nó sẽ gửi yêu cầu GET đến `https://corp.example.com/mod?go-get=1` và tìm URL kho lưu trữ trong phản hồi. Để tránh yêu cầu này, hãy đảm bảo mỗi đường dẫn module riêng tư có hậu tố VCS (như `.git`) đánh dấu tiền tố gốc kho lưu trữ. Ví dụ, khi lệnh `go` tải module `corp.example.com/repo.git/mod`, nó sẽ clone kho Git tại `https://corp.example.com/repo.git` hoặc `ssh://corp.example.com/repo.git` mà không cần thêm yêu cầu bổ sung.

Các nhà phát triển cần quyền đọc trên các kho lưu trữ chứa module riêng tư. Điều này có thể được cấu hình trong các tệp cấu hình VCS toàn cục như `.gitconfig`. Tốt nhất là cấu hình các công cụ VCS để không cần nhắc xác thực tương tác. Mặc định, khi gọi Git, lệnh `go` tắt nhắc tương tác bằng cách đặt `GIT_TERMINAL_PROMPT=0`, nhưng vẫn tôn trọng các cài đặt tường minh.

### Truyền thông tin xác thực đến proxy riêng tư {#private-module-proxy-auth}

Lệnh `go` hỗ trợ HTTP [basic authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) khi giao tiếp với các máy chủ proxy.

Thông tin xác thực có thể được chỉ định trong tệp [`.netrc`](https://www.gnu.org/software/inetutils/manual/html_node/The-_002enetrc-file.html). Ví dụ, tệp `.netrc` chứa các dòng bên dưới sẽ cấu hình lệnh `go` kết nối đến máy `proxy.corp.example.com` với tên người dùng và mật khẩu đã cho.

```
machine proxy.corp.example.com
login jrgopher
password hunter2
```

Vị trí của tệp có thể được đặt bằng biến môi trường `NETRC`. Nếu `NETRC` không được đặt, lệnh `go` sẽ đọc `$HOME/.netrc` trên các nền tảng giống UNIX hoặc `%USERPROFILE%\_netrc` trên Windows.

Các trường trong `.netrc` được phân cách bằng khoảng trắng, tab và ký tự xuống dòng. Đáng tiếc, các ký tự này không thể dùng trong tên người dùng hoặc mật khẩu. Lưu ý rằng tên máy không thể là URL đầy đủ, nên không thể chỉ định tên người dùng và mật khẩu khác nhau cho các đường dẫn khác nhau trên cùng một máy.

Ngoài ra, thông tin xác thực có thể được chỉ định trực tiếp trong URL `GOPROXY`. Ví dụ:

```
GOPROXY=https://jrgopher:hunter2@proxy.corp.example.com
```

Hãy thận trọng khi dùng cách này: các biến môi trường có thể xuất hiện trong lịch sử shell và trong các bản ghi nhật ký.

### Truyền thông tin xác thực đến kho lưu trữ riêng tư {#private-module-repo-auth}

Lệnh `go` có thể tải module trực tiếp từ kho lưu trữ quản lý phiên bản. Điều này cần thiết cho các module riêng tư nếu không dùng proxy riêng tư. Xem [Truy cập trực tiếp vào module riêng tư](#private-module-proxy-direct) để cấu hình.

Lệnh `go` chạy các công cụ quản lý phiên bản như `git` khi tải module trực tiếp. Các công cụ này tự thực hiện xác thực, nên bạn có thể cần cấu hình thông tin xác thực trong tệp cấu hình của công cụ như `.gitconfig`.

Để đảm bảo điều này hoạt động suôn sẻ, hãy đảm bảo lệnh `go` dùng URL kho lưu trữ đúng và công cụ quản lý phiên bản không yêu cầu nhập mật khẩu tương tác. Lệnh `go` ưu tiên URL `https://` hơn các scheme khác như `ssh://` trừ khi scheme được chỉ định khi [tra cứu URL kho lưu trữ](#vcs-find). Riêng với các kho lưu trữ GitHub, lệnh `go` mặc định dùng `https://`.

<!-- TODO(golang.org/issue/26134): if this issue is fixed, we can remove the
mention of the special case for GitHub above. -->

Với hầu hết các máy chủ, bạn có thể cấu hình client xác thực qua HTTP. Ví dụ, GitHub hỗ trợ dùng [OAuth personal access token làm mật khẩu HTTP](https://docs.github.com/en/free-pro-team@latest/github/extending-github/git-automation-with-oauth-tokens). Bạn có thể lưu mật khẩu HTTP trong tệp `.netrc`, tương tự như khi [truyền thông tin xác thực đến proxy riêng tư](#private-module-proxy-auth).

Ngoài ra, bạn có thể viết lại URL `https://` sang scheme khác. Ví dụ, trong `.gitconfig`:

```
[url "git@github.com:"]
    insteadOf = https://github.com/
```

Để biết thêm thông tin, xem [Tại sao "go get" dùng HTTPS khi clone kho lưu trữ?](/doc/faq#git_https)

### Quyền riêng tư {#private-module-privacy}

Lệnh `go` có thể tải module và metadata từ các máy chủ module proxy và hệ thống quản lý phiên bản. Biến môi trường `GOPROXY` kiểm soát các máy chủ nào được dùng. Các biến môi trường `GOPRIVATE` và `GONOPROXY` kiểm soát các module nào được tải từ proxy.

Giá trị mặc định của `GOPROXY` là:

```
https://proxy.golang.org,direct
```

Với cài đặt này, khi lệnh `go` tải một module hoặc metadata module, nó sẽ gửi yêu cầu đến `proxy.golang.org` trước, một module proxy công khai do Google vận hành ([chính sách bảo mật](https://proxy.golang.org/privacy)). Xem [giao thức `GOPROXY`](#goproxy-protocol) để biết chi tiết về thông tin được gửi trong mỗi yêu cầu. Lệnh `go` không truyền thông tin nhận dạng cá nhân, nhưng truyền đường dẫn module đầy đủ được yêu cầu. Nếu proxy phản hồi 404 (Not Found) hoặc 410 (Gone), lệnh `go` sẽ cố kết nối trực tiếp đến hệ thống quản lý phiên bản cung cấp module đó. Xem [Hệ thống quản lý phiên bản](#vcs) để biết chi tiết.

Các biến môi trường `GOPRIVATE` hoặc `GONOPROXY` có thể được đặt thành danh sách các mẫu glob khớp với tiền tố module là riêng tư và không nên được yêu cầu từ bất kỳ proxy nào. Ví dụ:

```
GOPRIVATE=*.corp.example.com,*.research.example.com
```

`GOPRIVATE` chỉ đơn giản đóng vai trò mặc định cho `GONOPROXY` và `GONOSUMDB`, nên không cần đặt `GONOPROXY` trừ khi `GONOSUMDB` cần giá trị khác. Khi đường dẫn module khớp với `GONOPROXY`, lệnh `go` bỏ qua `GOPROXY` cho module đó và tải trực tiếp từ kho lưu trữ quản lý phiên bản của nó. Điều này hữu ích khi không có proxy nào phục vụ module riêng tư. Xem [Truy cập trực tiếp vào module riêng tư](#private-module-proxy-direct).

Nếu có [proxy đáng tin cậy phục vụ tất cả module](#private-module-proxy-all), thì không nên đặt `GONOPROXY`. Ví dụ, nếu `GOPROXY` được đặt thành một nguồn duy nhất, lệnh `go` sẽ không tải module từ các nguồn khác. `GONOSUMDB` vẫn nên được đặt trong trường hợp này.

```
GOPROXY=https://proxy.corp.example.com
GONOSUMDB=*.corp.example.com,*.research.example.com
```

Nếu có [proxy đáng tin cậy chỉ phục vụ module riêng tư](#private-module-proxy-private), không nên đặt `GONOPROXY`, nhưng cần chú ý đảm bảo proxy phản hồi với các mã trạng thái đúng. Ví dụ, xét cấu hình sau:

```
GOPROXY=https://proxy.corp.example.com,https://proxy.golang.org
GONOSUMDB=*.corp.example.com,*.research.example.com
```

Giả sử do nhầm lẫn, một nhà phát triển cố tải một module không tồn tại.

```
go mod download corp.example.com/secret-product/typo@latest
```

Lệnh `go` trước tiên yêu cầu module này từ `proxy.corp.example.com`. Nếu proxy đó phản hồi 404 (Not Found) hoặc 410 (Gone), lệnh `go` sẽ chuyển sang `proxy.golang.org`, truyền đường dẫn `secret-product` trong URL yêu cầu. Nếu proxy riêng tư phản hồi với mã lỗi khác, lệnh `go` in lỗi và không chuyển sang các nguồn khác.

Ngoài các proxy, lệnh `go` có thể kết nối đến cơ sở dữ liệu checksum để xác minh các hash mã hóa của các module không được liệt kê trong `go.sum`. Biến môi trường `GOSUMDB` đặt tên, URL và khóa công khai của cơ sở dữ liệu checksum. Giá trị mặc định của `GOSUMDB` là `sum.golang.org`, cơ sở dữ liệu checksum công khai do Google vận hành ([chính sách bảo mật](https://sum.golang.org/privacy)). Xem [Cơ sở dữ liệu checksum](#checksum-database) để biết chi tiết về thông tin được truyền với mỗi yêu cầu. Tương tự như proxy, lệnh `go` không truyền thông tin nhận dạng cá nhân, nhưng truyền đường dẫn module đầy đủ được yêu cầu, và cơ sở dữ liệu checksum không thể tính checksum cho các module không công khai.

Biến môi trường `GONOSUMDB` có thể được đặt thành các mẫu chỉ ra module nào là riêng tư và không nên được yêu cầu từ cơ sở dữ liệu checksum. `GOPRIVATE` đóng vai trò mặc định cho `GONOSUMDB` và `GONOPROXY`, nên không cần đặt `GONOSUMDB` trừ khi `GONOPROXY` cần giá trị khác.

Một proxy có thể [mirror cơ sở dữ liệu checksum](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#proxying-a-checksum-database). Nếu một proxy trong `GOPROXY` làm điều này, lệnh `go` sẽ không kết nối trực tiếp đến cơ sở dữ liệu checksum.

`GOSUMDB` có thể được đặt thành `off` để tắt hoàn toàn việc dùng cơ sở dữ liệu checksum. Với cài đặt này, lệnh `go` sẽ không xác thực các module đã tải trừ khi chúng đã có trong `go.sum`. Xem [Xác thực module](#authenticating).

## Bộ đệm module {#module-cache}

<dfn>Bộ đệm module</dfn> là thư mục nơi lệnh `go` lưu trữ các tệp module đã tải. Bộ đệm module khác với bộ đệm build, nơi chứa các package đã biên dịch và các artifact build khác.

Vị trí mặc định của bộ đệm module là `$GOPATH/pkg/mod`. Để dùng vị trí khác, đặt biến môi trường `GOMODCACHE` [environment variable](#environment-variables).

Bộ đệm module không có kích thước tối đa, và lệnh `go` không tự động xóa nội dung của nó.

Bộ đệm có thể được chia sẻ bởi nhiều dự án Go được phát triển trên cùng một máy. Lệnh `go` sẽ dùng cùng một bộ đệm bất kể vị trí của module chính. Nhiều phiên bản lệnh `go` có thể truy cập cùng một bộ đệm module cùng lúc một cách an toàn.

Lệnh `go` tạo các tệp và thư mục mã nguồn module trong bộ đệm với quyền chỉ đọc để tránh các thay đổi không cố ý vào module sau khi tải. Điều này có tác dụng phụ không mong muốn là làm cho bộ đệm khó xóa bằng các lệnh như `rm -rf`. Thay vào đó, bộ đệm có thể được xóa bằng [`go clean -modcache`](#go-clean-modcache). Ngoài ra, khi dùng cờ `-modcacherw`, lệnh `go` sẽ tạo các thư mục mới với quyền đọc-ghi. Điều này làm tăng nguy cơ các trình soạn thảo, bài kiểm thử và các chương trình khác sửa đổi tệp trong bộ đệm module. Lệnh [`go mod verify`](#go-mod-verify) có thể được dùng để phát hiện các sửa đổi đối với các dependency của module chính. Lệnh quét nội dung đã giải nén của từng dependency module và xác nhận chúng khớp với hash mong đợi trong `go.sum`.

Bảng dưới đây giải thích mục đích của hầu hết các tệp trong bộ đệm module. Một số tệp tạm thời (tệp khóa, thư mục tạm thời) bị bỏ qua. Với mỗi đường dẫn, `$module` là đường dẫn module và `$version` là phiên bản. Các đường dẫn kết thúc bằng dấu gạch chéo (`/`) là thư mục. Các chữ cái in hoa trong đường dẫn module và phiên bản được thoát bằng dấu chấm than (`Azure` được thoát thành `!azure`) để tránh xung đột trên các hệ thống tệp không phân biệt chữ hoa chữ thường.

<table class="ModTable">
  <thead>
    <tr>
      <th>Path</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>$module@$version/</code></td>
      <td>
        Directory containing extracted contents of a module <code>.zip</code>
        file. This serves as a module root directory for a downloaded module. It
        won't contain a <code>go.mod</code> file if the original module
        didn't have one.
      </td>
    </tr>
    <tr>
      <td><code>cache/download/</code></td>
      <td>
        Directory containing files downloaded from module proxies and files
        derived from <a href="#vcs">version control systems</a>. The layout of
        this directory follows the
        <a href="#goproxy-protocol"><code>GOPROXY</code> protocol</a>, so
        this directory may be used as a proxy when served by an HTTP file
        server or when referenced with a <code>file://</code> URL.
      </td>
    </tr>
    <tr>
      <td><code>cache/download/$module/@v/list</code></td>
      <td>
        List of known versions (see
        <a href="#goproxy-protocol"><code>GOPROXY</code> protocol</a>). This
        may change over time, so the <code>go</code> command usually fetches a
        new copy instead of re-using this file.
      <td>
        Danh sách các phiên bản đã biết (xem
        <a href="#goproxy-protocol">giao thức <code>GOPROXY</code></a>). Danh sách
        này có thể thay đổi theo thời gian, vì vậy lệnh <code>go</code> thường tải
        một bản mới thay vì tái sử dụng tệp này.
      </td>
    </tr>
    <tr>
      <td><code>cache/download/$module/@v/$version.info</code></td>
      <td>
        Siêu dữ liệu JSON về phiên bản. (xem
        <a href="#goproxy-protocol">giao thức <code>GOPROXY</code></a>). Dữ liệu
        này có thể thay đổi theo thời gian, vì vậy lệnh <code>go</code> thường tải
        một bản mới thay vì tái sử dụng tệp này.
      </td>
    </tr>
    <tr>
      <td><code>cache/download/$module/@v/$version.mod</code></td>
      <td>
        Tệp <code>go.mod</code> của phiên bản này (xem
        <a href="#goproxy-protocol">giao thức <code>GOPROXY</code></a>). Nếu
        module gốc không có tệp <code>go.mod</code>, đây là tệp tổng hợp không
        có yêu cầu nào.
      </td>
    </tr>
    <tr>
      <td><code>cache/download/$module/@v/$version.zip</code></td>
      <td>
        Nội dung nén của module (xem
        <a href="#goproxy-protocol">giao thức <code>GOPROXY</code></a> và
        <a href="#zip-files">Tệp zip module</a>).
      </td>
    </tr>
    <tr>
      <td><code>cache/download/$module/@v/$version.ziphash</code></td>
      <td>
        Giá trị băm mật mã của các tệp trong tệp <code>.zip</code>.
        Lưu ý rằng bản thân tệp <code>.zip</code> không được băm, vì vậy thứ tự
        tệp, nén, căn chỉnh và siêu dữ liệu không ảnh hưởng đến giá trị băm.
        Khi sử dụng một module, lệnh <code>go</code> xác minh rằng giá trị băm
        này khớp với dòng tương ứng trong
        <a href="#go-sum-files"><code>go.sum</code></a>. Lệnh
        <a href="#go-mod-verify"><code>go mod verify</code></a> kiểm tra
        rằng giá trị băm của các tệp <code>.zip</code> và thư mục đã giải nén
        khớp với các tệp này.
      </td>
    </tr>
    <tr>
      <td><code>cache/download/sumdb/</code></td>
      <td>
        Thư mục chứa các tệp được tải xuống từ một
        <a href="#checksum-database">cơ sở dữ liệu checksum</a> (thường là
        <code>sum.golang.org</code>).
      </td>
    </tr>
    <tr>
      <td><code>cache/vcs/</code></td>
      <td>
        Chứa các kho lưu trữ quản lý phiên bản đã được nhân bản dành cho các module
        được tải trực tiếp từ nguồn gốc. Tên thư mục là giá trị băm được mã hóa hex
        dựa trên kiểu kho lưu trữ và URL. Các kho lưu trữ được tối ưu hóa để
        tiết kiệm dung lượng đĩa. Ví dụ, các kho Git được nhân bản ở dạng bare
        và shallow khi có thể.
      </td>
    </tr>
  </tbody>
</table>

## Xác thực module {#authenticating}

Khi lệnh `go` tải xuống [tệp zip](#zip-files) của module hoặc tệp [`go.mod`
](#go-mod-file) vào [bộ nhớ đệm module](#module-cache), lệnh này tính toán
giá trị băm mật mã và so sánh với giá trị đã biết để xác minh tệp chưa bị
thay đổi kể từ lần đầu tải xuống. Lệnh `go` báo lỗi bảo mật nếu tệp tải
xuống không có giá trị băm đúng.

Với tệp `go.mod`, lệnh `go` tính giá trị băm từ nội dung tệp. Với tệp zip
của module, lệnh `go` tính giá trị băm từ tên và nội dung của các tệp trong
kho lưu trữ theo thứ tự xác định. Giá trị băm không bị ảnh hưởng bởi thứ
tự tệp, nén, căn chỉnh và các siêu dữ liệu khác. Xem
[`golang.org/x/mod/sumdb/dirhash`](https://pkg.go.dev/golang.org/x/mod/sumdb/dirhash?tab=doc)
để biết chi tiết về cách triển khai băm.

Lệnh `go` so sánh từng giá trị băm với dòng tương ứng trong tệp [`go.sum`
](#go-sum-files) của module chính. Nếu giá trị băm khác với giá trị trong
`go.sum`, lệnh `go` báo lỗi bảo mật và xóa tệp đã tải xuống mà không thêm
vào bộ nhớ đệm module.

Nếu tệp `go.sum` không tồn tại, hoặc nếu tệp không chứa giá trị băm cho
tệp đã tải xuống, lệnh `go` có thể xác minh giá trị băm bằng cách dùng
[cơ sở dữ liệu checksum](#checksum-database), nguồn cung cấp giá trị băm
toàn cầu cho các module công khai. Khi giá trị băm đã được xác minh, lệnh
`go` thêm nó vào `go.sum` và thêm tệp đã tải xuống vào bộ nhớ đệm module.
Nếu module là private (khớp với biến môi trường `GOPRIVATE` hoặc
`GONOSUMDB`) hoặc nếu cơ sở dữ liệu checksum bị vô hiệu hóa (bằng cách đặt
`GOSUMDB=off`), lệnh `go` chấp nhận giá trị băm và thêm tệp vào bộ nhớ đệm
module mà không xác minh.

Bộ nhớ đệm module thường được chia sẻ bởi tất cả các dự án Go trên một hệ
thống, và mỗi module có thể có tệp `go.sum` riêng với các giá trị băm khác
nhau. Để tránh phải tin tưởng các module khác, lệnh `go` xác minh giá trị
băm bằng `go.sum` của module chính mỗi khi truy cập tệp trong bộ nhớ đệm
module. Việc tính giá trị băm tệp zip tốn chi phí tính toán, vì vậy lệnh
`go` kiểm tra các giá trị băm được tính sẵn lưu cùng tệp zip thay vì tính
lại. Lệnh [`go mod verify`](#go-mod-verify) có thể dùng để kiểm tra rằng
các tệp zip và thư mục đã giải nén chưa bị thay đổi kể từ khi được thêm vào
bộ nhớ đệm module.

### Tệp go.sum {#go-sum-files}

Một module có thể có tệp văn bản tên `go.sum` trong thư mục gốc của nó, cạnh
tệp `go.mod`. Tệp `go.sum` chứa giá trị băm mật mã của các dependency trực
tiếp và gián tiếp của module. Khi lệnh `go` tải tệp `.mod` hoặc `.zip` của
module vào [bộ nhớ đệm module](#module-cache), lệnh này tính giá trị băm và
kiểm tra rằng giá trị băm khớp với giá trị tương ứng trong tệp `go.sum` của
module chính. `go.sum` có thể trống hoặc không tồn tại nếu module không có
dependency hoặc nếu tất cả dependency được thay thế bằng các thư mục cục bộ
bằng [chỉ thị `replace`](#go-mod-file-replace).

Mỗi dòng trong `go.sum` có ba trường phân cách bởi dấu cách: đường dẫn module,
phiên bản (có thể kết thúc bằng `/go.mod`), và giá trị băm.

* Đường dẫn module là tên của module mà giá trị băm thuộc về.
* Phiên bản là phiên bản của module mà giá trị băm thuộc về. Nếu phiên bản
  kết thúc bằng `/go.mod`, giá trị băm chỉ dành cho tệp `go.mod` của module;
  nếu không, giá trị băm dành cho các tệp trong tệp `.zip` của module.
* Cột giá trị băm gồm tên thuật toán (như `h1`) và giá trị băm mật mã được
  mã hóa base64, phân cách bởi dấu hai chấm (`:`). Hiện tại, SHA-256 (`h1`)
  là thuật toán băm duy nhất được hỗ trợ. Nếu phát hiện lỗ hổng bảo mật trong
  SHA-256 trong tương lai, hỗ trợ sẽ được bổ sung cho thuật toán khác (được
  đặt tên là `h2` và tiếp theo).

Tệp `go.sum` có thể chứa giá trị băm cho nhiều phiên bản của một module. Lệnh
`go` có thể cần tải các tệp `go.mod` từ nhiều phiên bản của một dependency để
thực hiện [lựa chọn phiên bản tối thiểu](#minimal-version-selection).
`go.sum` cũng có thể chứa giá trị băm cho các phiên bản module không còn cần
thiết (ví dụ sau khi nâng cấp). [`go mod tidy`](#go-mod-tidy) sẽ thêm các
giá trị băm còn thiếu và xóa các giá trị băm không cần thiết khỏi `go.sum`.

### Cơ sở dữ liệu checksum {#checksum-database}

Cơ sở dữ liệu checksum là nguồn toàn cầu cung cấp các dòng `go.sum`. Lệnh
`go` có thể dùng nó trong nhiều tình huống để phát hiện hành vi sai của các
proxy hoặc máy chủ gốc.

Cơ sở dữ liệu checksum cho phép nhất quán và tin cậy toàn cầu cho tất cả các
phiên bản module công khai. Nó khiến các proxy không đáng tin cậy trở nên khả
thi vì chúng không thể cung cấp mã sai mà không bị phát hiện. Nó cũng đảm bảo
rằng các bit liên kết với một phiên bản cụ thể không thay đổi từ ngày này sang
ngày khác, kể cả khi tác giả module sau đó thay đổi các tag trong kho lưu trữ.

Cơ sở dữ liệu checksum được phục vụ bởi [sum.golang.org](https://sum.golang.org),
do Google vận hành. Đây là một [Transparent
Log](https://research.swtch.com/tlog) (hay "Merkle Tree") của các giá trị băm
dòng `go.sum`, được hỗ trợ bởi [Trillian](https://github.com/google/trillian).
Ưu điểm chính của Merkle tree là các kiểm toán viên độc lập có thể xác minh
rằng nó chưa bị giả mạo, do đó đáng tin cậy hơn một cơ sở dữ liệu thông thường.

Lệnh `go` tương tác với cơ sở dữ liệu checksum bằng giao thức được mô tả ban
đầu trong [Đề xuất: Bảo mật Hệ sinh thái Module Go Công khai](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#checksum-database).

Bảng dưới đây chỉ định các truy vấn mà cơ sở dữ liệu checksum phải phản hồi.
Với mỗi đường dẫn, `$base` là phần đường dẫn của URL cơ sở dữ liệu checksum,
`$module` là đường dẫn module, và `$version` là phiên bản. Ví dụ, nếu URL cơ
sở dữ liệu checksum là `https://sum.golang.org`, và client đang yêu cầu bản
ghi cho module `golang.org/x/text` ở phiên bản `v0.3.2`, client sẽ gửi yêu
cầu `GET` tới
`https://sum.golang.org/lookup/golang.org/x/text@v0.3.2`.

Để tránh nhập nhằng khi phục vụ từ hệ thống tệp không phân biệt chữ hoa chữ
thường, các phần tử `$module` và `$version` được
[mã hóa chữ hoa](https://pkg.go.dev/golang.org/x/mod/module#EscapePath)
bằng cách thay mỗi chữ hoa bằng dấu chấm than theo sau là chữ thường tương
ứng. Điều này cho phép cả hai module `example.com/M` và `example.com/m` đều
được lưu trữ trên đĩa, vì module đầu được mã hóa thành `example.com/!m`.

Các phần đường dẫn được bao quanh bởi dấu ngoặc vuông, như `[.p/$W]`, biểu
thị các giá trị tùy chọn.

<table class="ModTable">
  <thead>
    <tr>
      <th>Đường dẫn</th>
      <th>Mô tả</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>$base/latest</code></td>
      <td>
        Trả về mô tả cây đã ký, được mã hóa cho log mới nhất. Mô tả đã ký này
        ở dạng
        <a href="https://pkg.go.dev/golang.org/x/mod/sumdb/note">note</a>,
        là văn bản đã được ký bởi một hoặc nhiều khóa máy chủ và có thể được
        xác minh bằng khóa công khai của máy chủ. Mô tả cây cung cấp kích thước
        của cây và giá trị băm của đầu cây tại kích thước đó. Cách mã hóa này
        được mô tả trong
        <code><a href="https://pkg.go.dev/golang.org/x/mod/sumdb/tlog#FormatTree">
        golang.org/x/mod/sumdb/tlog#FormatTree</a></code>.
      </td>
    </tr>
    <tr>
    <tr>
      <td><code>$base/lookup/$module@$version</code></td>
      <td>
        Trả về số thứ tự bản ghi log cho mục về <code>$module</code>
        ở <code>$version</code>, theo sau là dữ liệu cho bản ghi (tức là
        các dòng <code>go.sum</code> cho <code>$module</code> ở
        <code>$version</code>) và mô tả cây đã ký, được mã hóa chứa bản ghi.
      </td>
    </tr>
    <tr>
    <tr>
      <td><code>$base/tile/$H/$L/$K[.p/$W]</code></td>
      <td>
        Trả về một <a href="https://research.swtch.com/tlog#serving_tiles">log tile</a>,
        là một tập hợp các giá trị băm tạo thành một phần của log. Mỗi tile
        được xác định tại tọa độ hai chiều ở mức tile <code>$L</code>, thứ
        <code>$K</code> từ trái sang, với chiều cao tile là <code>$H</code>.
        Hậu tố tùy chọn <code>.p/$W</code> biểu thị một log tile không đầy đủ
        chỉ có <code>$W</code> giá trị băm. Client phải chuyển sang tải tile
        đầy đủ nếu tile không đầy đủ không được tìm thấy.
      </td>
    </tr>
    <tr>
    <tr>
      <td><code>$base/tile/$H/data/$K[.p/$W]</code></td>
      <td>
        Trả về dữ liệu bản ghi cho các giá trị băm lá trong
        <code>/tile/$H/0/$K[.p/$W]</code> (với phần tử đường dẫn <code>data</code>
        theo nghĩa đen).
      </td>
    </tr>
    <tr>
  </tbody>
</table>

Nếu lệnh `go` tham khảo cơ sở dữ liệu checksum, bước đầu tiên là lấy dữ liệu
bản ghi qua endpoint `/lookup`. Nếu phiên bản module chưa được ghi trong log,
cơ sở dữ liệu checksum sẽ cố gắng tải nó từ máy chủ gốc trước khi phản hồi.
Dữ liệu `/lookup` này cung cấp tổng kiểm tra cho phiên bản module này cùng
với vị trí của nó trong log, giúp client biết tile nào cần tải để thực hiện
bằng chứng. Lệnh `go` thực hiện bằng chứng "inclusion" (rằng một bản ghi cụ
thể tồn tại trong log) và bằng chứng "consistency" (rằng cây chưa bị giả mạo)
trước khi thêm các dòng `go.sum` mới vào tệp `go.sum` của module chính. Điều
quan trọng là dữ liệu từ `/lookup` không bao giờ được dùng mà không xác thực
nó trước với giá trị băm cây đã ký và xác thực giá trị băm cây đã ký với
dòng thời gian của client về các giá trị băm cây đã ký.

Các giá trị băm cây đã ký và các tile mới được phục vụ bởi cơ sở dữ liệu
checksum được lưu trong bộ nhớ đệm module, vì vậy lệnh `go` chỉ cần tải các
tile còn thiếu.

Lệnh `go` không cần kết nối trực tiếp với cơ sở dữ liệu checksum. Lệnh này
có thể yêu cầu tổng kiểm tra module qua một module proxy có [phản chiếu cơ sở
dữ liệu checksum](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#proxying-a-checksum-database)
và hỗ trợ giao thức trên. Điều này đặc biệt hữu ích cho các proxy nội bộ
của doanh nghiệp vốn chặn các yêu cầu ra ngoài tổ chức.

Biến môi trường `GOSUMDB` xác định tên cơ sở dữ liệu checksum cần dùng và
tùy chọn khóa công khai và URL của nó, ví dụ:

```
GOSUMDB="sum.golang.org"
GOSUMDB="sum.golang.org+<publickey>"
GOSUMDB="sum.golang.org+<publickey> https://sum.golang.org"
```

Lệnh `go` biết khóa công khai của `sum.golang.org`, và cũng biết rằng tên
`sum.golang.google.cn` (khả dụng ở nội địa Trung Quốc) kết nối tới cơ sở dữ
liệu checksum `sum.golang.org`; việc dùng bất kỳ cơ sở dữ liệu nào khác đòi
hỏi phải cung cấp khóa công khai một cách tường minh. URL mặc định là
`https://` theo sau là tên cơ sở dữ liệu.

`GOSUMDB` mặc định là `sum.golang.org`, cơ sở dữ liệu checksum Go do Google
vận hành. Xem https://sum.golang.org/privacy để biết chính sách bảo mật của
dịch vụ.

Nếu `GOSUMDB` được đặt thành `off`, hoặc nếu `go get` được gọi với cờ
`-insecure`, cơ sở dữ liệu checksum sẽ không được tham khảo, và tất cả các
module không được nhận dạng đều được chấp nhận, đánh đổi bằng việc từ bỏ bảo
đảm bảo mật về các lượt tải xuống có thể lặp lại đã được xác minh cho tất cả
module. Cách tốt hơn để bỏ qua cơ sở dữ liệu checksum cho các module cụ thể
là dùng biến môi trường `GOPRIVATE` hoặc `GONOSUMDB`. Xem
[Module private](#private-modules) để biết chi tiết.

Lệnh `go env -w` có thể được dùng để
[thiết lập các biến này](/pkg/cmd/go/#hdr-Print_Go_environment_information)
cho các lần gọi lệnh `go` trong tương lai.

## Biến môi trường {#environment-variables}

Hành vi module trong lệnh `go` có thể được cấu hình bằng các biến môi trường
liệt kê bên dưới. Danh sách này chỉ bao gồm các biến môi trường liên quan đến
module. Xem [`go help
environment`](/cmd/go/#hdr-Environment_variables) để biết danh sách tất cả
các biến môi trường được lệnh `go` nhận dạng.

<table class="ModTable">
  <thead>
    <tr>
      <th>Biến</th>
      <th>Mô tả</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>GO111MODULE</code></td>
      <td>
        <p>
          Kiểm soát liệu lệnh <code>go</code> có chạy ở chế độ nhận biết module
          hay chế độ <code>GOPATH</code>. Ba giá trị được nhận dạng:
        </p>
        <ul>
          <li>
            <code>off</code>: lệnh <code>go</code> bỏ qua các tệp
            <code>go.mod</code> và chạy ở chế độ <code>GOPATH</code>.
          </li>
          <li>
            <code>on</code> (hoặc không đặt): lệnh <code>go</code> chạy ở
            chế độ nhận biết module, kể cả khi không có tệp <code>go.mod</code>.
          </li>
          <li>
            <code>auto</code>: lệnh <code>go</code> chạy ở chế độ nhận biết
            module nếu có tệp <code>go.mod</code> trong thư mục hiện tại hoặc
            bất kỳ thư mục cha nào. Trong Go 1.15 và thấp hơn, đây là giá trị
            mặc định.
          </li>
        </ul>
        <p>
          Xem <a href="#mod-commands">Lệnh nhận biết module</a> để biết thêm
          thông tin.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>GOMODCACHE</code></td>
      <td>
        <p>
          Thư mục nơi lệnh <code>go</code> sẽ lưu các module đã tải xuống và
          các tệp liên quan. Xem <a href="#module-cache">Bộ nhớ đệm module</a>
          để biết chi tiết về cấu trúc thư mục này.
        </p>
        <p>
          Nếu <code>GOMODCACHE</code> không được đặt, nó mặc định là
          <code>$GOPATH/pkg/mod</code>.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>GOINSECURE</code></td>
      <td>
        <p>
          Danh sách phân cách bằng dấu phẩy của các mẫu glob (theo cú pháp của
          Go <a href="/pkg/path/#Match"><code>path.Match</code></a>) của các
          tiền tố đường dẫn module có thể luôn được tải theo cách không bảo
          mật. Chỉ áp dụng cho các dependency đang được tải trực tiếp.
        </p>
        <p>
          Không giống cờ <code>-insecure</code> trên <code>go get</code>,
          <code>GOINSECURE</code> không vô hiệu hóa việc xác thực cơ sở dữ liệu
          checksum module. Có thể dùng <code>GOPRIVATE</code> hoặc
          <code>GONOSUMDB</code> để đạt được điều đó.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>GONOPROXY</code></td>
      <td>
        <p>
          Danh sách phân cách bằng dấu phẩy của các mẫu glob (theo cú pháp của
          Go <a href="/pkg/path/#Match"><code>path.Match</code></a>) của các
          tiền tố đường dẫn module luôn nên được tải trực tiếp từ kho lưu trữ
          quản lý phiên bản, không qua các module proxy.
        </p>
        <p>
          Nếu <code>GONOPROXY</code> không được đặt, nó mặc định là
          <code>GOPRIVATE</code>. Xem
          <a href="#private-module-privacy">Quyền riêng tư</a>.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>GONOSUMDB</code></td>
      <td>
        <p>
          Danh sách phân cách bằng dấu phẩy của các mẫu glob (theo cú pháp của
          Go <a href="/pkg/path/#Match"><code>path.Match</code></a>) của các
          tiền tố đường dẫn module mà lệnh <code>go</code> không nên xác minh
          checksum bằng cơ sở dữ liệu checksum.
        </p>
        <p>
          Nếu <code>GONOSUMDB</code> không được đặt, nó mặc định là
          <code>GOPRIVATE</code>. Xem
          <a href="#private-module-privacy">Quyền riêng tư</a>.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>GOPATH</code></td>
      <td>
        <p>
          Trong chế độ <code>GOPATH</code>, biến <code>GOPATH</code> là danh
          sách các thư mục có thể chứa mã Go.
        </p>
        <p>
          Trong chế độ nhận biết module, <a href="#glos-module-cache">bộ nhớ
          đệm module</a> được lưu trong thư mục con <code>pkg/mod</code> của
          thư mục <code>GOPATH</code> đầu tiên. Mã nguồn module ngoài bộ nhớ
          đệm có thể được lưu trong bất kỳ thư mục nào.
        </p>
        <p>
          Nếu <code>GOPATH</code> không được đặt, nó mặc định là thư mục con
          <code>go</code> trong thư mục home của người dùng.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>GOPRIVATE</code></td>
      <td>
        Danh sách phân cách bằng dấu phẩy của các mẫu glob (theo cú pháp của
        Go <a href="/pkg/path/#Match"><code>path.Match</code></a>) của các tiền
        tố đường dẫn module nên được coi là private. <code>GOPRIVATE</code>
        là giá trị mặc định cho <code>GONOPROXY</code> và
        <code>GONOSUMDB</code>. Xem
        <a href="#private-module-privacy">Quyền riêng tư</a>. <code>GOPRIVATE</code>
        cũng xác định liệu một module có được coi là private đối với
        <code>GOVCS</code> hay không.
      </td>
    </tr>
    <tr>
      <td><code>GOPROXY</code></td>
      <td>
        <p>
          Danh sách URL module proxy, phân cách bằng dấu phẩy (<code>,</code>)
          hoặc dấu gạch đứng (<code>|</code>). Khi lệnh <code>go</code> tra cứu
          thông tin về module, nó liên hệ lần lượt từng proxy trong danh sách
          cho đến khi nhận được phản hồi thành công hoặc lỗi kết thúc. Một
          proxy có thể phản hồi với trạng thái 404 (Không tìm thấy) hoặc 410
          (Đã biến mất) để cho biết module không có trên máy chủ đó.
        </p>
        <p>
          Hành vi dự phòng lỗi của lệnh <code>go</code> được xác định bởi các
          ký tự phân cách giữa các URL. Nếu URL proxy theo sau là dấu phẩy,
          lệnh <code>go</code> chuyển sang URL tiếp theo sau lỗi 404 hoặc 410;
          tất cả lỗi khác được coi là kết thúc. Nếu URL proxy theo sau là dấu
          gạch đứng, lệnh <code>go</code> chuyển sang nguồn tiếp theo sau bất
          kỳ lỗi nào, kể cả lỗi không phải HTTP như timeout.
        </p>
        <p>
          URL <code>GOPROXY</code> có thể có các scheme <code>https</code>,
          <code>http</code>, hoặc <code>file</code>. Nếu URL không có scheme,
          <code>https</code> được giả định. Bộ nhớ đệm module có thể được dùng
          trực tiếp như một file proxy:
        </p>
        <pre>GOPROXY=file://$(go env GOMODCACHE)/cache/download</pre>
        <p>Hai từ khóa có thể dùng thay cho URL proxy:</p>
        <ul>
          <li>
            <code>off</code>: cấm tải module từ bất kỳ nguồn nào.
          </li>
          <li>
            <code>direct</code>: tải trực tiếp từ kho lưu trữ quản lý phiên
            bản thay vì dùng module proxy.
          </li>
        </ul>
        <p>
          <code>GOPROXY</code> mặc định là
          <code>https://proxy.golang.org,direct</code>. Với cấu hình đó, lệnh
          <code>go</code> trước tiên liên hệ mirror module Go do Google vận hành,
          sau đó chuyển sang kết nối trực tiếp nếu mirror không có module. Xem
          <a href="https://proxy.golang.org/privacy">https://proxy.golang.org/privacy</a>
          để biết chính sách bảo mật của mirror. Các biến môi trường
          <code>GOPRIVATE</code> và <code>GONOPROXY</code> có thể được đặt để
          ngăn các module cụ thể khỏi bị tải qua proxy. Xem
          <a href="#private-module-privacy">Quyền riêng tư</a> để biết thông
          tin về cấu hình proxy private.
        </p>
        <p>
          Xem <a href="#module-proxy">Module proxy</a> và
          <a href="#resolve-pkg-mod">Phân giải package thành module</a> để biết
          thêm thông tin về cách proxy được dùng.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>GOSUMDB</code></td>
      <td>
        <p>
          Xác định tên cơ sở dữ liệu checksum cần dùng và tùy chọn khóa công
          khai và URL của nó. Ví dụ:
        </p>
        <pre>
GOSUMDB="sum.golang.org"
GOSUMDB="sum.golang.org+&lt;publickey&gt;"
GOSUMDB="sum.golang.org+&lt;publickey&gt; https://sum.golang.org"
</pre>
        <p>
          Lệnh <code>go</code> biết khóa công khai của
          <code>sum.golang.org</code> và cũng biết rằng tên
          <code>sum.golang.google.cn</code> (khả dụng ở nội địa Trung Quốc)
          kết nối tới cơ sở dữ liệu <code>sum.golang.org</code>; việc dùng bất
          kỳ cơ sở dữ liệu nào khác đòi hỏi phải cung cấp khóa công khai một
          cách tường minh. URL mặc định là <code>https://</code> theo sau là
          tên cơ sở dữ liệu.
        </p>
        <p>
          <code>GOSUMDB</code> mặc định là <code>sum.golang.org</code>, cơ sở
          dữ liệu checksum Go do Google vận hành. Xem
          <a href="https://sum.golang.org/privacy">https://sum.golang.org/privacy</a>
          để biết chính sách bảo mật của dịch vụ.
        <p>
        <p>
          Nếu <code>GOSUMDB</code> được đặt thành <code>off</code> hoặc nếu
          <code>go get</code> được gọi với cờ <code>-insecure</code>, cơ sở
          dữ liệu checksum sẽ không được tham khảo, và tất cả các module không
          được nhận dạng đều được chấp nhận, đánh đổi bằng việc từ bỏ bảo đảm
          bảo mật về các lượt tải xuống có thể lặp lại đã được xác minh cho tất
          cả module. Cách tốt hơn để bỏ qua cơ sở dữ liệu checksum cho các
          module cụ thể là dùng biến môi trường <code>GOPRIVATE</code> hoặc
          <code>GONOSUMDB</code>.
        </p>
        <p>
          Xem <a href="#authenticating">Xác thực module</a> và
          <a href="#private-module-privacy">Quyền riêng tư</a> để biết thêm
          thông tin.
        </p>
      </td>
    </tr>
    <tr>
      <td><code>GOVCS</code></td>
      <td>
        <p>
          Kiểm soát tập hợp các công cụ quản lý phiên bản mà lệnh <code>go</code>
          có thể dùng để tải các module công khai và private (được xác định bởi
          liệu đường dẫn của chúng có khớp với mẫu trong <code>GOPRIVATE</code>)
          hoặc các module khác khớp với mẫu glob.
        </p>
        <p>
          Nếu <code>GOVCS</code> không được đặt, hoặc nếu module không khớp với
          bất kỳ mẫu nào trong <code>GOVCS</code>, lệnh <code>go</code> có thể
          dùng <code>git</code> và <code>hg</code> cho module công khai, hoặc
          bất kỳ công cụ quản lý phiên bản nào đã biết cho module private. Cụ
          thể, lệnh <code>go</code> hoạt động như thể <code>GOVCS</code> được
          đặt thành:
        </p>
        <pre>public:git|hg,private:all</pre>
        <p>
          Xem <a href="#vcs-govcs">Kiểm soát công cụ quản lý phiên bản với
          <code>GOVCS</code></a> để biết giải thích đầy đủ.
        </p>
      </td>
    </tr>
     <tr>
      <td><code>GOWORK</code></td>
      <td>
       <p>
        Biến môi trường `GOWORK` hướng dẫn lệnh `go` vào chế độ workspace
        bằng cách dùng tệp [`go.work`](#go-work-file) được cung cấp để định
        nghĩa workspace. Nếu `GOWORK` được đặt thành `off`, chế độ workspace
        bị vô hiệu hóa. Điều này có thể dùng để chạy lệnh `go` ở chế độ
        single-module: ví dụ, `GOWORK=off go build .` xây dựng package `.`
        ở chế độ single-module. Nếu `GOWORK` trống, lệnh `go` sẽ tìm kiếm
        tệp `go.work` như mô tả trong phần [Workspace](#workspaces).
       </p>
      </td>
    </tr>
  </tbody>
</table>

## Bảng thuật ngữ {#glossary}

<a id="glos-build-constraint"></a>
**build constraint:** Điều kiện xác định liệu tệp nguồn Go có được dùng khi
biên dịch package hay không. Build constraint có thể được biểu đạt bằng hậu
tố tên tệp (ví dụ, `foo_linux_amd64.go`) hoặc bằng các comment build
constraint (ví dụ, `// +build linux,amd64`). Xem [Build
Constraints](/pkg/go/build/#hdr-Build_Constraints).

<a id="glos-build-list"></a>
**build list:** Danh sách các phiên bản module sẽ được dùng cho một lệnh
build như `go build`, `go list`, hoặc `go test`. Build list được xác định từ
tệp [`go.mod`](#glos-go-mod-file) của [module chính](#glos-main-module) và
các tệp `go.mod` trong các module được yêu cầu bắc cầu bằng cách dùng [lựa
chọn phiên bản tối thiểu](#glos-minimal-version-selection). Build list chứa
các phiên bản cho tất cả module trong [đồ thị module](#glos-module-graph),
không chỉ những module liên quan đến một lệnh cụ thể.

<a id="glos-canonical-version"></a>
**canonical version:** Một [phiên bản](#glos-version) được định dạng đúng
không có hậu tố siêu dữ liệu build nào ngoài `+incompatible`. Ví dụ,
`v1.2.3` là canonical version, nhưng `v1.2.3+meta` thì không.

<a id="glos-current-module"></a>
**current module:** Từ đồng nghĩa với [module chính](#glos-main-module).

<a id="glos-deprecated-module"></a>
**deprecated module:** Module không còn được tác giả hỗ trợ (mặc dù các
phiên bản major được coi là các module riêng biệt cho mục đích này). Module
deprecated được đánh dấu bằng [comment deprecation](#go-mod-file-module-deprecation)
trong phiên bản mới nhất của tệp [`go.mod`](#glos-go-mod-file) của nó.

<a id="glos-direct-dependency"></a>
**direct dependency:** Package có đường dẫn xuất hiện trong [khai báo
`import`](/ref/spec#import_declarations) trong tệp nguồn `.go` của package
hoặc test trong [module chính](#glos-main-module), hoặc module chứa package
như vậy. (So sánh với [indirect dependency](#glos-indirect-dependency).)

<a id="glos-direct-mode"></a>
**direct mode:** Cài đặt của [biến môi trường](#environment-variables) khiến
lệnh `go` tải module trực tiếp từ [hệ thống quản lý phiên bản](#vcs), thay
vì từ [module proxy](#glos-module-proxy). `GOPROXY=direct` thực hiện điều này
cho tất cả module. `GOPRIVATE` và `GONOPROXY` thực hiện điều này cho các
module khớp với danh sách mẫu.

<a id="glos-go-mod-file"></a>
**`go.mod` file:** Tệp định nghĩa đường dẫn, yêu cầu và siêu dữ liệu khác
của module. Xuất hiện trong [thư mục gốc của module](#glos-module-root-directory).
Xem phần về [tệp `go.mod`](#go-mod-file).

<a id="glos-go-work-file"></a>
**`go.work` file:** Tệp định nghĩa tập hợp các module được dùng trong một
[workspace](#workspaces). Xem phần về
[tệp `go.work`](#go-work-file).

<a id="glos-import-path"></a>
**import path:** Chuỗi dùng để import package trong tệp nguồn Go. Đồng nghĩa
với [đường dẫn package](#glos-package-path).

<a id="glos-indirect-dependency"></a>
**indirect dependency:** Package được import bắc cầu bởi package hoặc test
trong [module chính](#glos-main-module), nhưng đường dẫn của nó không xuất
hiện trong bất kỳ [khai báo `import`](/ref/spec#import_declarations) nào
trong module chính; hoặc module xuất hiện trong [đồ thị module](#glos-module-graph)
nhưng không cung cấp bất kỳ package nào được module chính import trực tiếp.
(So sánh với [direct dependency](#glos-direct-dependency).)

<a id="glos-lazy-module-loading"></a>
**lazy module loading:** Thay đổi trong Go 1.17 tránh tải [đồ thị module](#glos-module-graph)
cho các lệnh không cần đến nó trong các module chỉ định `go 1.17` hoặc cao
hơn. Xem [Lazy module loading](#lazy-loading).

<a id="glos-main-module"></a>
**main module:** Module mà lệnh `go` được gọi trong đó. Module chính được
định nghĩa bởi tệp [`go.mod`](#glos-go-mod-file) trong thư mục hiện tại
hoặc thư mục cha. Xem [Module, package và phiên bản](#modules-overview).

<a id="glos-major-version"></a>
**major version:** Số đầu tiên trong semantic version (`1` trong `v1.2.3`).
Trong một bản phát hành có các thay đổi không tương thích, major version phải
được tăng lên, và minor và patch version phải được đặt về 0. Các semantic
version có major version 0 được coi là không ổn định.

<a id="glos-major-version-subdirectory"></a>
**major version subdirectory:** Thư mục con trong kho lưu trữ quản lý phiên
bản khớp với [hậu tố major version](#glos-major-version-suffix) của module,
nơi module có thể được định nghĩa. Ví dụ, module `example.com/mod/v2` trong
kho lưu trữ có [đường dẫn gốc](#glos-repository-root-path) `example.com/mod`
có thể được định nghĩa trong thư mục gốc kho lưu trữ hoặc thư mục con major
version `v2`. Xem [Thư mục module trong kho lưu trữ](#vcs-dir).

<a id="glos-major-version-suffix"></a>
**major version suffix:** Hậu tố đường dẫn module khớp với số major version.
Ví dụ, `/v2` trong `example.com/mod/v2`. Hậu tố major version là bắt buộc
từ `v2.0.0` trở lên và không được phép ở các phiên bản trước. Xem phần về
[Hậu tố major version](#major-version-suffixes).

<a id="glos-minimal-version-selection"></a>
**minimal version selection (MVS):** Thuật toán dùng để xác định các phiên
bản của tất cả module sẽ được dùng trong một build. Xem phần về
[Lựa chọn phiên bản tối thiểu](#minimal-version-selection) để biết chi tiết.

<a id="glos-minor-version"></a>
**minor version:** Số thứ hai trong semantic version (`2` trong `v1.2.3`).
Trong một bản phát hành có chức năng mới tương thích ngược, minor version
phải được tăng lên, và patch version phải được đặt về 0.

<a id="glos-module"></a>
**module:** Tập hợp các package được phát hành, đánh phiên bản và phân phối
cùng nhau.

<a id="glos-module-cache"></a>
**module cache:** Thư mục cục bộ lưu các module đã tải xuống, nằm tại
`GOPATH/pkg/mod`. Xem [Bộ nhớ đệm module](#module-cache).

<a id="glos-module-graph"></a>
**module graph:** Đồ thị có hướng của các yêu cầu module, có gốc tại [module
chính](#glos-main-module). Mỗi đỉnh trong đồ thị là một module; mỗi cạnh là
một phiên bản từ lệnh `require` trong tệp `go.mod` (chịu ảnh hưởng của các
lệnh `replace` và `exclude` trong tệp `go.mod` của module chính).

<a id="glos-module-graph-pruning"></a>
**module graph pruning:** Thay đổi trong Go 1.17 giảm kích thước đồ thị
module bằng cách bỏ qua các dependency bắc cầu của các module chỉ định `go
1.17` hoặc cao hơn. Xem [Module graph pruning](#graph-pruning).

<a id="glos-module-path"></a>
**module path:** Đường dẫn xác định module và đóng vai trò là tiền tố cho các
đường dẫn import package trong module. Ví dụ, `"golang.org/x/net"`.

<a id="glos-module-proxy"></a>
**module proxy:** Máy chủ web triển khai [giao thức
`GOPROXY`](#goproxy-protocol). Lệnh `go` tải thông tin phiên bản, tệp
`go.mod` và tệp zip module từ các module proxy.

<a id="glos-module-root-directory"></a>
**module root directory:** Thư mục chứa tệp `go.mod` định nghĩa module.

<a id="glos-module-subdirectory"></a>
**module subdirectory:** Phần [đường dẫn module](#glos-module-path) sau
[đường dẫn gốc kho lưu trữ](#glos-repository-root-path) cho biết thư mục
con nơi module được định nghĩa. Khi không trống, thư mục con module cũng là
tiền tố cho các [tag semantic version](#glos-semantic-version-tag). Thư mục
con module không bao gồm [hậu tố major version](#glos-major-version-suffix),
nếu có, kể cả khi module nằm trong [thư mục con major version](#glos-major-version-subdirectory).
Xem [Đường dẫn module](#module-path).

<a id="glos-package"></a>
**package:** Tập hợp các tệp nguồn trong cùng thư mục được biên dịch cùng
nhau. Xem [phần Packages](/ref/spec#Packages) trong Đặc tả ngôn ngữ Go.

<a id="glos-package-path"></a>
**package path:** Đường dẫn xác định duy nhất một package. Đường dẫn package
là [đường dẫn module](#glos-module-path) ghép với thư mục con trong module.
Ví dụ `"golang.org/x/net/html"` là đường dẫn package của package trong module
`"golang.org/x/net"` trong thư mục con `"html"`. Đồng nghĩa với
[import path](#glos-import-path).

<a id="glos-patch-version"></a>
**patch version:** Số thứ ba trong semantic version (`3` trong `v1.2.3`).
Trong một bản phát hành không có thay đổi nào đối với giao diện công khai
của module, patch version phải được tăng lên.

<a id="glos-pre-release-version"></a>
**pre-release version:** Phiên bản có dấu gạch ngang theo sau là một loạt
các định danh phân cách bằng dấu chấm ngay sau patch version, ví dụ,
`v1.2.3-beta4`. Các pre-release version được coi là không ổn định và không
được giả định là tương thích với các phiên bản khác. Một pre-release version
được sắp xếp trước release version tương ứng: `v1.2.3-pre` đứng trước
`v1.2.3`. Xem thêm [release version](#glos-release-version).

<a id="glos-pseudo-version"></a>
**pseudo-version:** Phiên bản mã hóa định danh sửa đổi (như hash commit Git)
và dấu thời gian từ hệ thống quản lý phiên bản. Ví dụ,
`v0.0.0-20191109021931-daa7c04131f5`. Được dùng cho [tương thích với kho lưu
trữ không phải module](#non-module-compat) và trong các tình huống khác khi
không có phiên bản được gắn tag.

<a id="glos-release-version"></a>
**release version:** Phiên bản không có hậu tố pre-release. Ví dụ, `v1.2.3`,
không phải `v1.2.3-pre`. Xem thêm [pre-release
version](#glos-pre-release-version).

<a id="glos-repository-root-path"></a>
**repository root path:** Phần [đường dẫn module](#glos-module-path) tương
ứng với thư mục gốc của kho lưu trữ quản lý phiên bản. Xem [Đường dẫn
module](#module-path).

<a id="glos-retracted-version"></a>
**retracted version:** Phiên bản không nên được phụ thuộc vào, hoặc vì nó
được công bố sớm hoặc vì phát hiện vấn đề nghiêm trọng sau khi công bố.
Xem [chỉ thị `retract`](#go-mod-file-retract).

<a id="glos-semantic-version-tag"></a>
**semantic version tag:** Tag trong kho lưu trữ quản lý phiên bản ánh xạ một
[phiên bản](#glos-version) tới một sửa đổi cụ thể. Xem [Ánh xạ phiên bản
tới commit](#vcs-version).

<a id="glos-selected-version"></a>
**selected version:** Phiên bản của một module nhất định được chọn bởi [lựa
chọn phiên bản tối thiểu](#minimal-version-selection). Phiên bản được chọn
là phiên bản cao nhất cho đường dẫn của module được tìm thấy trong [đồ thị
module](#glos-module-graph).

<a id="glos-vendor-directory"></a>
**vendor directory:** Thư mục tên `vendor` chứa các package từ các module
khác cần để build các package trong module chính. Được duy trì bằng
[`go mod vendor`](#go-mod-vendor). Xem [Vendoring](#vendoring).

<a id="glos-version"></a>
**version:** Định danh cho một snapshot bất biến của module, được viết là chữ
`v` theo sau là semantic version. Xem phần về [Phiên bản](#versions).

<a id="glos-workspace"></a>
**workspace:** Tập hợp các module trên đĩa được dùng làm các module chính khi
chạy [minimal version selection (MVS)](#minimal-version-selection).
Xem phần về [Workspace](#workspaces).
