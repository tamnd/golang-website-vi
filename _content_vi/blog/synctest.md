---
title: Kiểm thử mã đồng thời với testing/synctest
date: 2025-02-19
by:
- Damien Neil
tags:
- concurrency
- testing
summary: Go 1.24 có một gói thử nghiệm mới để hỗ trợ kiểm thử mã đồng thời.
template: true
---

Một trong những tính năng đặc trưng của Go là hỗ trợ tính đồng thời tích hợp sẵn.
Goroutine và channel là các nguyên thủy đơn giản và hiệu quả để
viết các chương trình đồng thời.

Tuy nhiên, kiểm thử các chương trình đồng thời có thể khó khăn và dễ mắc lỗi.

Trong Go 1.24, chúng tôi giới thiệu một gói thử nghiệm mới,
[`testing/synctest`](/pkg/testing/synctest),
để hỗ trợ kiểm thử mã đồng thời. Bài viết này sẽ giải thích động lực đằng sau
thử nghiệm này, minh họa cách sử dụng gói synctest, và thảo luận về tiềm năng của nó trong tương lai.

Trong Go 1.24, gói `testing/synctest` là thử nghiệm và
không thuộc phạm vi cam kết tương thích của Go.
Nó không hiển thị theo mặc định.
Để sử dụng, hãy biên dịch mã của bạn với `GOEXPERIMENT=synctest` được thiết lập trong môi trường.

## Kiểm thử chương trình đồng thời là khó

Để bắt đầu, hãy xem xét một ví dụ đơn giản.

Hàm [`context.AfterFunc`](/pkg/context#AfterFunc)
sắp xếp để một hàm được gọi trong goroutine riêng của nó sau khi một context bị hủy.
Đây là một bài kiểm thử có thể có cho `AfterFunc`:

{{raw `
    func TestAfterFunc(t *testing.T) {
        ctx, cancel := context.WithCancel(context.Background())

        calledCh := make(chan struct{}) // closed when AfterFunc is called
        context.AfterFunc(ctx, func() {
            close(calledCh)
        })

        // TODO: Assert that the AfterFunc has not been called.

        cancel()

        // TODO: Assert that the AfterFunc has been called.
    }
`}}

Chúng ta muốn kiểm tra hai điều kiện trong bài kiểm thử này:
Hàm không được gọi trước khi context bị hủy,
và hàm *được* gọi sau khi context bị hủy.

Kiểm tra một điều kiện phủ định trong một hệ thống đồng thời là khó.
Ta có thể dễ dàng kiểm tra rằng hàm chưa được gọi *lúc này*,
nhưng làm sao để kiểm tra rằng nó *sẽ không* được gọi?

Một cách tiếp cận phổ biến là chờ một khoảng thời gian nhất định trước khi
kết luận rằng sự kiện sẽ không xảy ra.
Hãy thử thêm một hàm trợ giúp vào bài kiểm thử của chúng ta để làm điều này.

{{raw `
    // funcCalled reports whether the function was called.
    funcCalled := func() bool {
        select {
        case <-calledCh:
            return true
        case <-time.After(10 * time.Millisecond):
            return false
        }
    }

    if funcCalled() {
        t.Fatalf("AfterFunc function called before context is canceled")
    }

    cancel()

    if !funcCalled() {
        t.Fatalf("AfterFunc function not called after context is canceled")
    }
`}}

Bài kiểm thử này chậm:
10 millisecond không phải là nhiều, nhưng nó sẽ tích lũy lại khi có nhiều bài kiểm thử.

Bài kiểm thử này cũng không ổn định:
10 millisecond là rất dài trên một máy tính nhanh,
nhưng không hiếm khi thấy các khoảng dừng kéo dài vài giây
trên các hệ thống
[CI](https://en.wikipedia.org/wiki/Continuous_integration)
dùng chung và bị quá tải.

Ta có thể làm bài kiểm thử ít không ổn định hơn với cái giá là làm nó chậm hơn,
và ta có thể làm nó ít chậm hơn với cái giá là làm nó không ổn định hơn,
nhưng ta không thể làm cho nó vừa nhanh vừa đáng tin cậy.

## Giới thiệu gói testing/synctest

Gói `testing/synctest` giải quyết vấn đề này.
Nó cho phép chúng ta viết lại bài kiểm thử này để vừa đơn giản, nhanh, và đáng tin cậy,
mà không cần thay đổi gì trong mã được kiểm thử.

Gói này chỉ chứa hai hàm: `Run` và `Wait`.

`Run` gọi một hàm trong một goroutine mới.
Goroutine này và bất kỳ goroutine nào được khởi động bởi nó
tồn tại trong một môi trường cô lập mà chúng ta gọi là một *bubble* (bong bóng).
`Wait` chờ cho đến khi mọi goroutine trong bubble của goroutine hiện tại
bị chặn bởi một goroutine khác trong cùng bubble.

Hãy viết lại bài kiểm thử trên bằng gói `testing/synctest`.

{{raw `
    func TestAfterFunc(t *testing.T) {
        synctest.Run(func() {
            ctx, cancel := context.WithCancel(context.Background())

            funcCalled := false
            context.AfterFunc(ctx, func() {
                funcCalled = true
            })

            synctest.Wait()
            if funcCalled {
                t.Fatalf("AfterFunc function called before context is canceled")
            }

            cancel()

            synctest.Wait()
            if !funcCalled {
                t.Fatalf("AfterFunc function not called after context is canceled")
            }
        })
    }
`}}

Bài kiểm thử này gần như giống hệt bài kiểm thử ban đầu của chúng ta,
nhưng chúng ta đã bọc bài kiểm thử trong một lời gọi `synctest.Run`
và gọi `synctest.Wait` trước khi kiểm tra xem hàm đã được gọi hay chưa.

Hàm `Wait` chờ cho đến khi mọi goroutine trong bubble của bên gọi bị chặn.
Khi nó trả về, chúng ta biết rằng gói context đã gọi hàm,
hoặc sẽ không gọi nó cho đến khi chúng ta thực hiện thêm hành động nào đó.

Bài kiểm thử này giờ đây vừa nhanh vừa đáng tin cậy.

Bài kiểm thử cũng đơn giản hơn:
chúng ta đã thay thế channel `calledCh` bằng một biến boolean.
Trước đây chúng ta cần dùng channel để tránh data race giữa
goroutine kiểm thử và goroutine `AfterFunc`,
nhưng hàm `Wait` giờ đây cung cấp sự đồng bộ hóa đó.

Bộ phát hiện race hiểu các lời gọi `Wait`,
và bài kiểm thử này vượt qua khi chạy với `-race`.
Nếu ta bỏ lời gọi `Wait` thứ hai,
bộ phát hiện race sẽ báo cáo đúng một data race trong bài kiểm thử.

## Kiểm thử thời gian

Mã đồng thời thường xuyên xử lý thời gian.

Kiểm thử mã làm việc với thời gian có thể khó.
Dùng thời gian thực trong bài kiểm thử gây ra các bài kiểm thử chậm và không ổn định,
như chúng ta đã thấy ở trên.
Dùng thời gian giả yêu cầu tránh các hàm của gói `time`,
và thiết kế mã được kiểm thử để hoạt động với
một đồng hồ giả tùy chọn.

Gói `testing/synctest` giúp việc kiểm thử mã sử dụng thời gian trở nên đơn giản hơn.

Các goroutine trong bubble được khởi động bởi `Run` sử dụng một đồng hồ giả.
Trong bubble, các hàm trong gói `time` hoạt động trên
đồng hồ giả. Thời gian tiến lên trong bubble khi tất cả các goroutine
bị chặn.

Để minh họa, hãy viết một bài kiểm thử cho hàm
[`context.WithTimeout`](/pkg/context#WithTimeout).
`WithTimeout` tạo ra một context con,
hết hạn sau một khoảng timeout nhất định.

{{raw `
    func TestWithTimeout(t *testing.T) {
        synctest.Run(func() {
            const timeout = 5 * time.Second
            ctx, cancel := context.WithTimeout(context.Background(), timeout)
            defer cancel()

            // Wait just less than the timeout.
            time.Sleep(timeout - time.Nanosecond)
            synctest.Wait()
            if err := ctx.Err(); err != nil {
                t.Fatalf("before timeout, ctx.Err() = %v; want nil", err)
            }

            // Wait the rest of the way until the timeout.
            time.Sleep(time.Nanosecond)
            synctest.Wait()
            if err := ctx.Err(); err != context.DeadlineExceeded {
                t.Fatalf("after timeout, ctx.Err() = %v; want DeadlineExceeded", err)
            }
        })
    }
`}}

Chúng ta viết bài kiểm thử này như thể đang làm việc với thời gian thực.
Điểm khác biệt duy nhất là chúng ta bọc hàm kiểm thử trong `synctest.Run`,
và gọi `synctest.Wait` sau mỗi lời gọi `time.Sleep` để chờ các timer của gói context
chạy xong.

## Chặn và bubble

Một khái niệm chính trong `testing/synctest` là bubble trở nên *bị chặn vĩnh cửu*.
Điều này xảy ra khi mọi goroutine trong bubble bị chặn,
và chỉ có thể được bỏ chặn bởi một goroutine khác trong bubble.

Khi một bubble bị chặn vĩnh cửu:

  - Nếu có một lời gọi `Wait` đang chờ, nó sẽ trả về.
  - Nếu không, thời gian tiến lên đến thời điểm tiếp theo có thể bỏ chặn một goroutine, nếu có.
  - Nếu không có, bubble bị deadlock và `Run` panic.

Một bubble không bị chặn vĩnh cửu nếu có bất kỳ goroutine nào bị chặn
nhưng có thể được đánh thức bởi một sự kiện từ bên ngoài bubble.

Danh sách đầy đủ các thao tác chặn vĩnh cửu một goroutine là:

  - gửi hoặc nhận trên một channel nil
  - gửi hoặc nhận bị chặn trên một channel được tạo trong cùng bubble
  - một câu lệnh select trong đó mọi nhánh đều bị chặn vĩnh cửu
  - `time.Sleep`
  - `sync.Cond.Wait`
  - `sync.WaitGroup.Wait`

### Mutex

Các thao tác trên `sync.Mutex` không bị chặn vĩnh cửu.

Các hàm thường xuyên thu thập một mutex toàn cục.
Ví dụ, một số hàm trong gói reflect
dùng một bộ nhớ cache toàn cục được bảo vệ bởi một mutex.
Nếu một goroutine trong một synctest bubble bị chặn khi thu thập
một mutex đang được giữ bởi một goroutine bên ngoài bubble,
nó không bị chặn vĩnh cửu vì nó bị chặn, nhưng sẽ được bỏ chặn
bởi một goroutine từ bên ngoài bubble của nó.

Vì mutex thường không được giữ trong thời gian dài,
chúng tôi đơn giản là không đưa chúng vào xem xét của `testing/synctest`.

### Channel

Channel được tạo trong một bubble hoạt động khác với channel được tạo bên ngoài.

Các thao tác trên channel chỉ bị chặn vĩnh cửu nếu channel đó thuộc bubble
(được tạo trong bubble).
Thao tác trên channel thuộc bubble từ bên ngoài bubble sẽ gây panic.

Những quy tắc này đảm bảo rằng một goroutine chỉ bị chặn vĩnh cửu khi
giao tiếp với các goroutine trong bubble của nó.

### I/O

Các thao tác I/O bên ngoài, chẳng hạn như đọc từ một kết nối mạng,
không bị chặn vĩnh cửu.

Đọc từ mạng có thể được bỏ chặn bởi các lần ghi từ bên ngoài bubble,
thậm chí từ các tiến trình khác.
Dù writer duy nhất cho một kết nối mạng cũng nằm trong cùng bubble,
runtime không thể phân biệt giữa kết nối đang chờ thêm dữ liệu
và kết nối mà kernel đã nhận dữ liệu và đang trong quá trình phân phối nó.

Kiểm thử một server hoặc client mạng với synctest thường
yêu cầu cung cấp một triển khai mạng giả.
Ví dụ, hàm [`net.Pipe`](/pkg/net#Pipe)
tạo ra một cặp `net.Conn` sử dụng kết nối mạng trong bộ nhớ
và có thể được dùng trong các bài kiểm thử synctest.

## Thời gian sống của bubble

Hàm `Run` khởi động một goroutine trong một bubble mới.
Nó trả về khi mọi goroutine trong bubble đã thoát.
Nó panic nếu bubble bị chặn vĩnh cửu
và không thể được bỏ chặn bằng cách tiến thời gian.

Yêu cầu mọi goroutine trong bubble thoát trước khi Run trả về
có nghĩa là các bài kiểm thử phải cẩn thận dọn dẹp bất kỳ goroutine nền nào
trước khi hoàn thành.

## Kiểm thử mã mạng

Hãy xem xét một ví dụ khác, lần này dùng gói `testing/synctest`
để kiểm thử một chương trình mạng.
Trong ví dụ này, chúng ta sẽ kiểm thử xử lý phản hồi 100 Continue của gói `net/http`.

Một HTTP client gửi request có thể kèm header "Expect: 100-continue"
để báo cho server biết rằng client có thêm dữ liệu cần gửi.
Server sau đó có thể phản hồi bằng phản hồi thông tin 100 Continue
để yêu cầu phần còn lại của request,
hoặc bằng trạng thái khác để báo cho client biết nội dung không cần thiết.
Ví dụ, một client đang tải lên một tệp lớn có thể dùng tính năng này để
xác nhận rằng server sẵn sàng chấp nhận tệp trước khi gửi nó.

Bài kiểm thử của chúng ta sẽ xác nhận rằng khi gửi header "Expect: 100-continue"
HTTP client không gửi nội dung request trước khi server
yêu cầu, và rằng nó gửi nội dung sau khi nhận được phản hồi
100 Continue.

Thông thường các bài kiểm thử về client và server giao tiếp có thể dùng
kết nối mạng loopback. Tuy nhiên, khi làm việc với `testing/synctest`,
thường ta muốn dùng kết nối mạng giả
để có thể phát hiện khi tất cả các goroutine đang bị chặn trên mạng.
Chúng ta bắt đầu bài kiểm thử này bằng cách tạo một `http.Transport` (một HTTP client) sử dụng
kết nối mạng trong bộ nhớ được tạo bởi [`net.Pipe`](/pkg/net#Pipe).

{{raw `
    func Test(t *testing.T) {
        synctest.Run(func() {
            srvConn, cliConn := net.Pipe()
            defer srvConn.Close()
            defer cliConn.Close()
            tr := &http.Transport{
                DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
                    return cliConn, nil
                },
                // Setting a non-zero timeout enables "Expect: 100-continue" handling.
                // Since the following test does not sleep,
                // we will never encounter this timeout,
                // even if the test takes a long time to run on a slow machine.
                ExpectContinueTimeout: 5 * time.Second,
            }
`}}

Chúng ta gửi một request trên transport này với header "Expect: 100-continue" được thiết lập.
Request được gửi trong một goroutine mới, vì nó sẽ không hoàn thành cho đến cuối bài kiểm thử.

{{raw `
            body := "request body"
            go func() {
                req, _ := http.NewRequest("PUT", "http://test.tld/", strings.NewReader(body))
                req.Header.Set("Expect", "100-continue")
                resp, err := tr.RoundTrip(req)
                if err != nil {
                    t.Errorf("RoundTrip: unexpected error %v", err)
                } else {
                    resp.Body.Close()
                }
            }()
`}}

Chúng ta đọc các header request được gửi bởi client.

{{raw `
            req, err := http.ReadRequest(bufio.NewReader(srvConn))
            if err != nil {
                t.Fatalf("ReadRequest: %v", err)
            }
`}}

Bây giờ chúng ta đến phần trọng tâm của bài kiểm thử.
Chúng ta muốn xác nhận rằng client sẽ chưa gửi nội dung request.

Chúng ta khởi động một goroutine mới sao chép nội dung được gửi đến server vào một `strings.Builder`,
chờ tất cả các goroutine trong bubble bị chặn, và xác minh rằng chúng ta chưa đọc gì
từ nội dung.

Nếu ta quên lời gọi `synctest.Wait`, bộ phát hiện race sẽ đúng đắn phàn nàn
về một data race, nhưng với `Wait` điều này là an toàn.

{{raw `
            var gotBody strings.Builder
            go io.Copy(&gotBody, req.Body)
            synctest.Wait()
            if got := gotBody.String(); got != "" {
                t.Fatalf("before sending 100 Continue, unexpectedly read body: %q", got)
            }
`}}

Chúng ta ghi phản hồi "100 Continue" cho client và xác minh rằng nó giờ đây gửi
nội dung request.

{{raw `
            srvConn.Write([]byte("HTTP/1.1 100 Continue\r\n\r\n"))
            synctest.Wait()
            if got := gotBody.String(); got != body {
                t.Fatalf("after sending 100 Continue, read body %q, want %q", got, body)
            }
`}}

Và cuối cùng, chúng ta hoàn thành bằng cách gửi phản hồi "200 OK" để kết thúc request.

Chúng ta đã khởi động nhiều goroutine trong bài kiểm thử này.
Lời gọi `synctest.Run` sẽ chờ tất cả chúng thoát trước khi trả về.

{{raw `
            srvConn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
        })
    }
`}}

Bài kiểm thử này có thể dễ dàng mở rộng để kiểm thử các hành vi khác,
chẳng hạn như xác minh rằng nội dung request không được gửi nếu server không yêu cầu,
hoặc rằng nó được gửi nếu server không phản hồi trong một khoảng timeout.

## Trạng thái của thử nghiệm

Chúng tôi đang giới thiệu [`testing/synctest`](/pkg/testing/synctest)
trong Go 1.24 dưới dạng một gói *thử nghiệm*.
Tùy thuộc vào phản hồi và kinh nghiệm,
chúng tôi có thể phát hành nó có hoặc không có sửa đổi,
tiếp tục thử nghiệm,
hoặc xóa nó trong phiên bản Go tương lai.

Gói này không hiển thị theo mặc định.
Để sử dụng, hãy biên dịch mã của bạn với `GOEXPERIMENT=synctest` được thiết lập trong môi trường.

Chúng tôi muốn nghe phản hồi của bạn!
Nếu bạn thử `testing/synctest`,
hãy báo cáo kinh nghiệm của bạn, tích cực hay tiêu cực,
tại [go.dev/issue/67434](/issue/67434).
