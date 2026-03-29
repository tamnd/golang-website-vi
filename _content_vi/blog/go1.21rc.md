---
title: Release Candidate Go 1.21
date: 2023-06-21
by:
- Eli Bendersky, on behalf of the Go team
summary: Go 1.21 RC mang đến cải tiến ngôn ngữ, các gói thư viện chuẩn mới, PGO chính thức, tương thích xuôi và ngược trong toolchain, và build nhanh hơn.
---

Release Candidate (RC) đầu tiên của Go 1.21 có mặt hôm nay trên [trang download](/dl/#go1.21rc2)! Go 1.21 chứa đầy tính năng và cải tiến mới.
Tải RC (release candidate) cho phép bạn thử nghiệm sớm, kiểm thử trên
workload của mình, và báo cáo bất kỳ vấn đề nào trước khi phát hành chính thức (dự kiến vào tháng 8). Dưới đây là một số thay đổi và tính năng đáng chú ý trong Go 1.21; để xem danh sách đầy đủ,
tham khảo [ghi chú phát hành đầy đủ](https://tip.golang.org/doc/go1.21).

*(Xin lưu ý rằng RC đầu tiên của Go 1.21 có tên `go1.21rc2`
vì một lỗi được tìm thấy và sửa sau khi tag `go1.21rc1`)*

## Cải tiến công cụ

- Tính năng Profile Guided Optimization (PGO) chúng tôi [thông báo xem trước trong
  1.20](/blog/pgo-preview) nay đã chính thức sẵn dùng! Nếu có file tên
  `default.pgo` trong thư mục gói main, lệnh `go` sẽ
  dùng nó để bật PGO build. Xem [tài liệu PGO](/doc/pgo) để biết
  thêm chi tiết. Chúng tôi đã đo tác động của PGO trên nhiều chương trình Go và
  thấy cải tiến hiệu năng từ 2-7%.
- [Công cụ `go`](/cmd/go) nay hỗ trợ tương thích ngôn ngữ [ngược](https://tip.golang.org/doc/godebug)
  và [xuôi](/doc/toolchain).

## Thay đổi ngôn ngữ

- Các hàm built-in mới: [min, max](https://tip.golang.org/ref/spec#Min_and_max)
  và [clear](https://tip.golang.org/ref/spec#Clear).
- Một số cải tiến cho suy luận kiểu với hàm generic. Mô tả về
  [suy luận kiểu trong đặc tả](https://tip.golang.org/ref/spec#Type_inference)
  đã được mở rộng và làm rõ.
- Trong phiên bản Go tương lai, chúng tôi đang lên kế hoạch giải quyết một trong những lỗi phổ biến nhất
  trong lập trình Go:
  [bắt giữ biến vòng lặp](/wiki/CommonMistakes).
  Go 1.21 đi kèm với bản xem trước của tính năng này mà bạn có thể bật trong code
  bằng cách dùng biến môi trường. Xem [trang wiki LoopvarExperiment](/wiki/LoopvarExperiment)
  để biết thêm chi tiết.

## Bổ sung thư viện chuẩn

- Gói [log/slog](https://tip.golang.org/pkg/log/slog) mới cho structured logging.
- Gói [slices](https://tip.golang.org/pkg/slices) mới cho các thao tác phổ biến
  trên slice của bất kỳ kiểu phần tử nào. Bao gồm các hàm sắp xếp thường
  nhanh hơn và tiện lợi hơn gói [sort](https://tip.golang.org/pkg/sort).
- Gói [maps](https://tip.golang.org/pkg/maps) mới cho các thao tác phổ biến trên map
  của bất kỳ kiểu khóa hay phần tử nào.
- Gói [cmp](https://tip.golang.org/pkg/cmp) mới với các tiện ích mới để so sánh
  các giá trị có thứ tự.

## Cải thiện hiệu năng

Ngoài các cải tiến hiệu năng khi bật PGO:

- Trình biên dịch Go đã được rebuild với PGO bật cho 1.21, và kết quả là
  nó build các chương trình Go nhanh hơn 2-4%, tùy thuộc vào kiến trúc host.
- Do điều chỉnh bộ gom rác, một số ứng dụng có thể thấy giảm tới 40%
  tail latency.
- Thu thập trace với [runtime/trace](https://pkg.go.dev/runtime/trace) nay
  tốn ít chi phí CPU hơn đáng kể trên amd64 và arm64.

## Cổng mới đến WASI

Go 1.21 thêm cổng thử nghiệm cho [WebAssembly System Interface (WASI)](https://wasi.dev/),
Preview 1 (`GOOS=wasip1`, `GOARCH=wasm`).

Để hỗ trợ viết code WebAssembly (WASM) tổng quát hơn, trình biên dịch cũng
hỗ trợ một chỉ thị mới để import hàm từ WASM host:
`go:wasmimport`.

Hãy [tải Go 1.21 RC](/dl/#go1.21rc2) và thử xem! Nếu bạn phát hiện bất kỳ
vấn đề nào, hãy [tạo issue](/issue/new).
