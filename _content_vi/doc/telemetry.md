---
title: "Go Telemetry"
layout: article
breadcrumb: true
date: 2024-02-07:00:00Z
template: true
---

<style>
.DocInfo {
  background-color: var(--color-background-info);
  padding: 1.5rem 2rem 1.5rem 4rem;
  border-left: 0.875rem solid var(--color-border);
  position: relative;
}
.DocInfo:before {
  content: "ⓘ";
  position: absolute;
  top: 1rem;
  left: 1rem;
  font-size: 2rem;
}
</style>

Mục lục:

 [Bối cảnh](#background)\
 [Tổng quan](#overview)\
 [Cấu hình](#config)\
 [Bộ đếm](#counters)\
 [Báo cáo và tải lên](#reports)\
 [Biểu đồ](#charts) \
 [Đề xuất Telemetry](#proposals)\
 [Nhắc nhở qua IDE](#ide) \
 [Câu hỏi thường gặp](#faq)

## Bối cảnh {#background}

Go telemetry là cách để các chương trình công cụ Go thu thập dữ liệu về hiệu suất và mức sử dụng của chúng. Ở đây "công cụ Go" có nghĩa là các công cụ dành cho lập trình viên được duy trì bởi nhóm Go, bao gồm lệnh `go` và các công cụ bổ sung như máy chủ ngôn ngữ Go [`gopls`] hoặc công cụ bảo mật Go [`govulncheck`]. Go telemetry chỉ được dùng cho các chương trình do nhóm Go duy trì và các dependency được chọn của chúng như [Delve].

Theo mặc định, dữ liệu telemetry chỉ được lưu trên máy tính cục bộ, nhưng người dùng có thể chọn tải lên một tập con được phê duyệt của dữ liệu telemetry lên [telemetry.go.dev].
Dữ liệu được tải lên giúp nhóm Go cải thiện ngôn ngữ Go và các công cụ của nó, bằng cách giúp chúng tôi hiểu mức sử dụng và các sự cố.

Từ "telemetry" đã mang những hàm ý tiêu cực trong thế giới phần mềm mã nguồn mở, trong nhiều trường hợp là xứng đáng. Tuy nhiên, việc đo lường trải nghiệm người dùng là một yếu tố quan trọng của kỹ thuật phần mềm hiện đại, và các nguồn dữ liệu như các vấn đề GitHub hoặc các khảo sát hàng năm là các chỉ số thô và chậm trễ, không đủ cho các loại câu hỏi mà nhóm Go cần có thể trả lời.
Go telemetry được thiết kế để giúp các chương trình trong công cụ thu thập dữ liệu hữu ích về độ tin cậy, hiệu suất và mức sử dụng của chúng, trong khi vẫn duy trì tính minh bạch và quyền riêng tư mà người dùng mong đợi từ dự án Go. Để tìm hiểu thêm về quy trình thiết kế và động lực đằng sau telemetry, hãy xem [các bài viết blog về telemetry](https://research.swtch.com/telemetry).
Để tìm hiểu thêm về telemetry và quyền riêng tư, hãy xem [chính sách quyền riêng tư của telemetry](https://telemetry.go.dev/privacy).

Trang này giải thích cách Go telemetry hoạt động một cách chi tiết. Để có câu trả lời nhanh cho các câu hỏi thường gặp, hãy xem phần [Câu hỏi thường gặp](#faq).

<div class="DocInfo">
Sử dụng Go 1.23 trở lên, để <strong>bật</strong> tải lên dữ liệu telemetry
cho nhóm Go, hãy chạy:
<pre>
go telemetry on
</pre>
Để tắt hoàn toàn telemetry, bao gồm cả việc thu thập cục bộ, hãy chạy:
<pre>
go telemetry off
</pre>
Để khôi phục về chế độ mặc định chỉ lưu cục bộ, hãy chạy:
<pre>
go telemetry local
</pre>
Trước Go 1.23, điều này cũng có thể được thực hiện với lệnh
<code>golang.org/x/telemetry/cmd/gotelemetry</code>. Xem phần <a
href="#config">Cấu hình</a> để biết thêm chi tiết.
</div>

## Tổng quan {#overview}

Go telemetry sử dụng ba kiểu dữ liệu cốt lõi:

- [_Bộ đếm_](#counters) là các số đếm nhẹ của các sự kiện được đặt tên, được thêm vào chương trình công cụ. Nếu việc thu thập được bật (chế độ [mode](#config) là **local** hoặc **on**), các bộ đếm được ghi vào tệp ánh xạ bộ nhớ trong hệ thống tệp cục bộ.
- [_Báo cáo_](#reports) là các bản tóm tắt tổng hợp của bộ đếm cho một tuần nhất định.
  Nếu tải lên được bật (chế độ [mode](#config) là **on**), các báo cáo cho [bộ đếm được phê duyệt](#proposals) sẽ được tải lên [telemetry.go.dev], nơi chúng có thể truy cập công khai.
- [_Biểu đồ_](#charts) tóm tắt các báo cáo đã tải lên cho tất cả người dùng.
  Biểu đồ có thể xem tại [telemetry.go.dev].

Tất cả dữ liệu và cấu hình Go telemetry cục bộ được lưu trong thư mục
<code>[os.UserConfigDir()](/pkg/os#UserConfigDir)/go/telemetry</code>.
Dưới đây, chúng tôi sẽ gọi thư mục này là `<gotelemetry>`.

Sơ đồ dưới đây minh họa luồng dữ liệu này.

<div class="image">
  <center>
    <img max-width="800px" src="/doc/telemetry/dataflow.png" />
  </center>
</div>

Trong phần còn lại của tài liệu này, chúng ta sẽ khám phá các thành phần của sơ đồ này. Nhưng trước tiên, hãy tìm hiểu thêm về cấu hình kiểm soát nó.

## Cấu hình {#config}

Hành vi của Go telemetry được kiểm soát bởi một giá trị duy nhất: _chế độ_ telemetry.
Các giá trị có thể cho `mode` là `local` (mặc định), `on`, hoặc `off`:

- Khi `mode` là `local`, dữ liệu telemetry được thu thập và lưu trên máy tính cục bộ, nhưng không bao giờ được tải lên máy chủ từ xa.
- Khi `mode` là `on`, dữ liệu được thu thập, và có thể được tải lên tùy thuộc vào [việc lấy mẫu](#uploads).
- Khi `mode` là `off`, dữ liệu không được thu thập cũng không được tải lên.

Với Go 1.23 trở lên, các lệnh sau tương tác với chế độ telemetry:

- `go telemetry`: xem chế độ hiện tại.
- `go telemetry on`: đặt chế độ thành `on`.
- `go telemetry off`: đặt chế độ thành `off`.
- `go telemetry local`: đặt chế độ thành `local`.

Thông tin về cấu hình telemetry cũng có sẵn thông qua các biến môi trường Go chỉ đọc:

- `go env GOTELEMETRY` báo cáo chế độ telemetry.
- `go env GOTELEMETRYDIR` báo cáo thư mục chứa cấu hình và dữ liệu telemetry.

Lệnh [`gotelemetry`](/pkg/golang.org/x/telemetry/cmd/gotelemetry) cũng có thể được dùng để cấu hình chế độ telemetry, cũng như để kiểm tra dữ liệu telemetry cục bộ. Dùng lệnh này để cài đặt nó:

```
go install golang.org/x/telemetry/cmd/gotelemetry@latest
```

Để biết thông tin sử dụng đầy đủ của công cụ dòng lệnh `gotelemetry`, hãy xem [tài liệu gói](/pkg/golang.org/x/telemetry/cmd/gotelemetry) của nó.

## Bộ đếm {#counters}

Như đã đề cập ở trên, Go telemetry được thêm vào thông qua _bộ đếm_. Bộ đếm có hai biến thể: bộ đếm cơ bản và bộ đếm stack.

### Bộ đếm cơ bản

_Bộ đếm cơ bản_ là giá trị có thể tăng với tên mô tả sự kiện mà nó đếm. Ví dụ, bộ đếm `gopls/client:vscode` ghi số lần phiên `gopls` được khởi tạo bởi VS Code. Cùng với bộ đếm này, chúng tôi có thể có `gopls/client:neovim`, `gopls/client:eglot`, v.v., để ghi các phiên với các trình soạn thảo hoặc client ngôn ngữ khác nhau. Nếu bạn dùng nhiều trình soạn thảo trong tuần, bạn có thể ghi dữ liệu bộ đếm sau:

    gopls/client:vscode 8
    gopls/client:neovim 5
    gopls/client:eglot  2

Khi các bộ đếm liên quan theo cách này, đôi khi chúng tôi gọi phần trước `:` là _tên biểu đồ_ (`gopls/client` trong trường hợp này), và phần sau `:` là _tên bucket_ (`vscode`). Chúng ta sẽ thấy tại sao điều này quan trọng khi thảo luận về [biểu đồ](#charts).

Bộ đếm cơ bản cũng có thể đại diện cho một _histogram_. Ví dụ, bộ đếm {{raw
`<code>gopls/completion/latency:&lt;50ms</code>`}} ghi số lần tự động hoàn thành mất ít hơn 50ms.

{{raw `
<pre>
gopls/completion/latency:&lt;10ms
gopls/completion/latency:&lt;50ms
gopls/completion/latency:&lt;100ms
...
</pre>
`}}

Mẫu này để ghi dữ liệu histogram là một quy ước: không có gì đặc biệt về tên bucket {{raw `<code>&lt;50ms</code>`}}. Các loại bộ đếm này thường được dùng để đo hiệu suất.

### Bộ đếm stack

_Bộ đếm stack_ là bộ đếm cũng ghi call stack hiện tại của chương trình công cụ Go khi số đếm được tăng. Ví dụ, bộ đếm stack `crash/crash` ghi call stack khi một chương trình công cụ bị crash:

    crash/crash
    golang.org/x/tools/gopls/internal/golang.hoverBuiltin:+22
    golang.org/x/tools/gopls/internal/golang.Hover:+94
    golang.org/x/tools/gopls/internal/server.Hover:+42
    ...

Bộ đếm stack thường đo các sự kiện khi các bất biến chương trình bị vi phạm.
Ví dụ phổ biến nhất là crash, nhưng một ví dụ khác là bộ đếm stack `gopls/bug`, đếm các tình huống bất thường được lập trình viên xác định trước, chẳng hạn như một panic được khôi phục hoặc một lỗi "không thể xảy ra". Bộ đếm stack chỉ bao gồm tên và số dòng của các hàm trong các chương trình công cụ Go. Chúng không bao gồm bất kỳ thông tin nào về đầu vào người dùng, chẳng hạn như tên hoặc nội dung mã nguồn của người dùng.

Bộ đếm stack có thể giúp theo dõi các lỗi hiếm gặp hoặc phức tạp mà không được báo cáo bằng các phương tiện khác. Kể từ khi giới thiệu bộ đếm `gopls/bug`, chúng tôi đã tìm thấy [hàng chục trường hợp](https://github.com/golang/go/issues?q=label%3Agopls%2Ftelemetry-wins) của mã "không thể đạt được" đã được đạt trong thực tế, và việc theo dõi các ngoại lệ này đã dẫn đến việc phát hiện (và sửa) nhiều lỗi hiển thị với người dùng mà người dùng không nhận ra hoặc quá khó để báo cáo. Đặc biệt với kiểm thử trước khi phát hành, bộ đếm stack có thể giúp chúng tôi cải thiện sản phẩm hiệu quả hơn so với khi không có tự động hóa.

### Tệp bộ đếm

Tất cả dữ liệu bộ đếm được ghi vào thư mục `<gotelemetry>/local`, trong các tệp được đặt tên theo schema sau:

```
[program name]@[program version]-[go version]-[GOOS]-[GOARCH]-[date].v1.count
```

- **Tên chương trình** là tên cơ sở của đường dẫn gói chương trình, như được báo cáo bởi [debug.BuildInfo].
- **Phiên bản chương trình** và **phiên bản go** cũng được báo cáo bởi [debug.BuildInfo].
- Các giá trị **GOOS** và **GOARCH** được báo cáo bởi
  [`runtime.GOOS`](/pkg/runtime#GOOS) và
  [`runtime.GOARCH`](/pkg/runtime#GOARCH).
- **Ngày** là ngày tệp bộ đếm được tạo, ở định dạng `YYYY-MM-DD`.

Các tệp này được ánh xạ bộ nhớ vào mỗi instance đang chạy của các chương trình được thêm công cụ đo. Việc sử dụng tệp ánh xạ bộ nhớ có nghĩa là ngay cả khi chương trình lập tức bị crash, hoặc nhiều bản sao của công cụ được thêm công cụ đo đang chạy đồng thời, các bộ đếm vẫn được ghi một cách an toàn.

## Báo cáo và tải lên {#reports}

Khoảng một lần mỗi tuần, dữ liệu bộ đếm được tổng hợp thành các báo cáo có tên `<date>.json` trong thư mục `<gotelemetry>/local`. Các báo cáo này tổng hợp tất cả các số đếm của tuần trước, được nhóm theo cùng các định danh chương trình được dùng cho tệp bộ đếm (tên chương trình, phiên bản chương trình, phiên bản go, GOOS và GOARCH).

Các báo cáo cục bộ có thể được xem như biểu đồ với lệnh [`gotelemetry view`](/pkg/golang.org/x/telemetry/cmd/gotelemetry).
Đây là một ví dụ tóm tắt của bộ đếm `gopls/completion/latency`:

<div class="image">
  <center>
    <img max-width="800px" src="/doc/telemetry/gopls-latency.png" />
  </center>
</div>

### Tải lên {#uploads}

Nếu việc tải lên telemetry được bật, quá trình báo cáo hàng tuần cũng sẽ tạo các báo cáo chứa tập con các bộ đếm có trong [cấu hình tải lên](https://telemetry.go.dev/config). Các bộ đếm này phải được phê duyệt bởi quy trình xem xét công khai được mô tả trong phần tiếp theo. Sau khi được tải lên thành công, một bản sao của các báo cáo đã tải lên được lưu trong thư mục `<gotelemetry>/upload`.

Khi đủ người dùng chọn tải lên dữ liệu telemetry, quá trình tải lên sẽ bỏ ngẫu nhiên việc tải lên cho một phần các báo cáo, để giảm lượng thu thập và tăng quyền riêng tư trong khi vẫn duy trì ý nghĩa thống kê.

## Biểu đồ {#charts}

Ngoài việc chấp nhận tải lên, trang web [telemetry.go.dev] làm cho dữ liệu đã tải lên có thể truy cập công khai. Mỗi ngày, các báo cáo đã tải lên được xử lý thành hai đầu ra, có sẵn trên trang chủ [telemetry.go.dev].

- Các báo cáo _hợp nhất_ hợp nhất các bộ đếm từ tất cả các lần tải lên nhận được vào ngày đó.
- _Biểu đồ_ vẽ đồ thị dữ liệu đã tải lên như được chỉ định trong [cấu hình biểu đồ], được tạo ra như một phần của quy trình đề xuất. Nhớ lại từ cuộc thảo luận về [bộ đếm](#counters) rằng tên bộ đếm như `foo:bar` được phân tách thành tên biểu đồ `foo` và tên bucket `bar`. Mỗi biểu đồ tổng hợp các bộ đếm có cùng tên biểu đồ vào các bucket tương ứng.

Biểu đồ được chỉ định theo định dạng của gói [chartconfig]. Ví dụ, đây là cấu hình biểu đồ cho biểu đồ `gopls/client`.

    title: Editor Distribution
    counter: gopls/client:{vscode,vscodium,vscode-insiders,code-server,eglot,govim,neovim,coc.nvim,sublimetext,other}
    description: measure editor distribution for gopls users.
    type: partition
    issue: https://go.dev/issue/61038
    issue: https://go.dev/issue/62214 # add vscode-insiders
    program: golang.org/x/tools/gopls
    version: v0.13.0 # temporarily back-version to demonstrate config generation.

Cấu hình này mô tả biểu đồ cần tạo, liệt kê tập hợp các bộ đếm cần tổng hợp, và chỉ định các phiên bản chương trình mà biểu đồ áp dụng. Ngoài ra, [quy trình đề xuất](#proposals) yêu cầu một đề xuất được chấp nhận phải được liên kết với biểu đồ. Đây là biểu đồ kết quả từ cấu hình đó:

<div class="image">
  <center>
    <img src="/doc/telemetry/gopls-clients.png" />
  </center>
</div>

## Quy trình đề xuất telemetry {#proposals}

Các thay đổi đối với cấu hình tải lên hoặc tập hợp biểu đồ trên [telemetry.go.dev] phải đi qua _quy trình đề xuất telemetry_, được thiết kế để đảm bảo tính minh bạch xung quanh các thay đổi đối với cấu hình telemetry.

Đáng chú ý là thực tế không có sự phân biệt giữa cấu hình tải lên và cấu hình biểu đồ trong quy trình này. Cấu hình tải lên được biểu thị theo các tổng hợp mà chúng tôi muốn hiển thị trên telemetry.go.dev, dựa trên nguyên tắc rằng chúng tôi chỉ nên thu thập dữ liệu mà chúng tôi muốn _xem_.

Quy trình đề xuất như sau:

1. Người đề xuất tạo CL sửa đổi [config.txt] của gói [chartconfig] để chứa các tổng hợp bộ đếm mới mong muốn.
2. Người đề xuất nộp [đề xuất] để hợp nhất CL này.
3. Khi thảo luận về vấn đề được giải quyết, đề xuất được thành viên nhóm Go phê duyệt hoặc từ chối.
4. Một quy trình tự động tái tạo cấu hình tải lên để cho phép tải lên các bộ đếm cần thiết cho biểu đồ mới. Quy trình này cũng sẽ thường xuyên thêm các phiên bản mới của các chương trình liên quan vào cấu hình tải lên khi chúng được phát hành.

Để được phê duyệt, các biểu đồ mới không được mang thông tin nhạy cảm của người dùng, và ngoài ra phải vừa hữu ích vừa khả thi. Để hữu ích, biểu đồ phải phục vụ một mục đích cụ thể, với các kết quả có thể hành động. Để khả thi, phải có thể thu thập đáng tin cậy dữ liệu cần thiết, và các phép đo kết quả phải có ý nghĩa thống kê. Để chứng minh tính khả thi, người đề xuất có thể được yêu cầu thêm công cụ đo bộ đếm vào chương trình mục tiêu và thu thập chúng cục bộ trước.

Toàn bộ tập hợp các đề xuất như vậy có sẵn tại [dự án đề xuất](https://github.com/orgs/golang/projects/29) trên GitHub.

## Nhắc nhở qua IDE {#ide}

Để telemetry có thể trả lời các loại câu hỏi chúng tôi muốn hỏi, tập hợp người dùng chọn tải lên không cần phải lớn, khoảng 16.000 người tham gia sẽ cho phép các phép đo có ý nghĩa thống kê ở mức độ chi tiết mong muốn. Tuy nhiên, vẫn có chi phí để tập hợp mẫu lành mạnh này: chúng tôi cần hỏi một số lượng lớn lập trình viên Go xem họ có muốn tham gia không.

Hơn nữa, ngay cả khi một số lượng lớn người dùng chọn tham gia _ngay bây giờ_ (có thể sau khi đọc một bài viết blog Go), những người dùng đó có thể bị lệch hướng về các lập trình viên Go có kinh nghiệm, và theo thời gian mẫu ban đầu đó sẽ càng lệch hơn.
Ngoài ra, khi mọi người thay máy tính, họ phải chủ động chọn tham gia lại. Trong loạt bài viết blog telemetry, điều này được gọi là ["chi phí chiến dịch"](https://research.swtch.com/telemetry-opt-in#campaign) của mô hình chọn tham gia.

Để giúp giữ mẫu người dùng tham gia luôn mới, máy chủ ngôn ngữ Go [`gopls`] hỗ trợ lời nhắc yêu cầu người dùng chọn tham gia Go telemetry.
Đây là giao diện từ VS Code:

<div class="image">
  <center>
    <img width="600px" src="/doc/telemetry/prompt.png" />
  </center>
</div>

Nếu người dùng chọn "Yes", [chế độ](#config) telemetry của họ sẽ được đặt thành `on`, giống như khi họ đã chạy [`gotelemetry on`](/pkg/golang.org/x/telemetry/cmd/gotelemetry). Theo cách này, việc chọn tham gia trở nên dễ dàng nhất có thể, và chúng tôi có thể liên tục tiếp cận một mẫu lớn và phân tầng của các lập trình viên Go.

## Câu hỏi thường gặp {#faq}

**H: Làm thế nào để bật hoặc tắt Go telemetry?**

T: Dùng lệnh `gotelemetry`, có thể được cài đặt với `go install
golang.org/x/telemetry/cmd/gotelemetry@latest`. Chạy `gotelemetry off` để tắt mọi thứ, kể cả việc thu thập cục bộ. Chạy `gotelemetry on` để bật mọi thứ, bao gồm tải lên các bộ đếm được phê duyệt lên [telemetry.go.dev]. Xem phần [Cấu hình](#config) để biết thêm.

**H: Dữ liệu cục bộ được lưu ở đâu?**

T: Trong thư mục <code>[os.UserConfigDir()](/pkg/os#UserConfigDir)/go/telemetry</code>.

**H: Dữ liệu được tải lên bao lâu một lần, nếu tôi chọn tham gia?**

T: Khoảng một lần mỗi tuần.

**H: Dữ liệu nào được tải lên, nếu tôi chọn tham gia?**

T: Chỉ các bộ đếm được liệt kê trong [cấu hình tải lên](https://telemetry.go.dev/config) mới có thể được tải lên. Điều này được tạo ra từ [cấu hình biểu đồ], có thể dễ đọc hơn.

**H: Bộ đếm được thêm vào cấu hình tải lên như thế nào?**

T: Thông qua [quy trình đề xuất công khai](#proposals).

**H: Tôi có thể xem dữ liệu telemetry đã được tải lên ở đâu?**

T: Dữ liệu đã tải lên có sẵn dưới dạng biểu đồ hoặc bản tóm tắt đã hợp nhất tại [telemetry.go.dev].

**H: Mã nguồn cho Go telemetry ở đâu?**

T: Tại [golang.org/x/telemetry](/pkg/golang.org/x/telemetry).

[`gopls`]: /pkg/golang.org/x/tools/gopls
[`govulncheck`]: /pkg/golang.org/x/vuln/cmd/govulncheck
[Delve]: /pkg/github.com/go-delve/delve#section-readme
[debug.BuildInfo]: /pkg/runtime/debug#BuildInfo
[proposal]: /issue/new?assignees=&labels=Telemetry-Proposal&projects=golang%2F29&template=12-telemetry.yml&title=x%2Ftelemetry%2Fconfig%3A+proposal+title
[telemetry.go.dev]: https://telemetry.go.dev
[chartconfig]: /pkg/golang.org/x/telemetry/internal/chartconfig
[config.txt]: https://go.googlesource.com/telemetry/+/refs/heads/master/internal/chartconfig/config.txt
