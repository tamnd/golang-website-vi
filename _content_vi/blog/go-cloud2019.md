---
title: Điểm mới trong Go Cloud Development Kit
date: 2019-03-04
by:
- The Go Cloud Development Kit team at Google
summary: Các thay đổi gần đây trong Go Cloud Development Kit (Go CDK).
---

## Giới thiệu

Tháng 7 năm ngoái, chúng tôi đã [giới thiệu](/blog/go-cloud) [Go Cloud Development Kit](https://gocloud.dev)
(trước đây gọi đơn giản là "Go Cloud"),
một dự án mã nguồn mở xây dựng các thư viện và công cụ nhằm cải thiện trải nghiệm
phát triển ứng dụng đám mây với Go.
Chúng tôi đã đạt được nhiều tiến bộ kể từ đó, xin cảm ơn những người đóng góp đầu tiên!
Chúng tôi mong được mở rộng cộng đồng người dùng và người đóng góp cho Go CDK,
và rất hào hứng được làm việc chặt chẽ với các nhóm dùng sớm.

## API đa nền tảng

Sáng kiến đầu tiên của chúng tôi là một tập hợp các API đa nền tảng cho các dịch vụ đám mây phổ biến.
Bạn viết ứng dụng dùng các API này,
sau đó triển khai lên bất kỳ kết hợp nhà cung cấp nào,
bao gồm AWS, GCP, Azure, tại chỗ, hoặc trên một máy phát triển duy nhất để kiểm thử.
Có thể thêm nhà cung cấp mới bằng cách triển khai một interface.

Các API đa nền tảng này rất phù hợp nếu bất kỳ điều nào sau đây đúng với bạn:

  - Bạn phát triển ứng dụng đám mây ở môi trường cục bộ.
  - Bạn có ứng dụng tại chỗ muốn chuyển lên đám mây (vĩnh viễn hoặc trong quá trình di chuyển).
  - Bạn cần khả năng chạy trên nhiều nhà cung cấp đám mây.
  - Bạn đang tạo ứng dụng Go mới sẽ sử dụng dịch vụ đám mây.

Khác với cách tiếp cận truyền thống khi bạn phải viết lại code ứng dụng cho từng nhà cung cấp,
với Go CDK bạn chỉ cần viết code ứng dụng một lần dùng các API đa nền tảng của chúng tôi
để truy cập các dịch vụ liệt kê dưới đây.
Sau đó, bạn có thể chạy ứng dụng trên bất kỳ đám mây nào được hỗ trợ với thay đổi cấu hình tối thiểu.

Tập API hiện tại bao gồm:

  - [blob](https://godoc.org/gocloud.dev/blob),
    để lưu trữ dữ liệu dạng blob.
    Các nhà cung cấp được hỗ trợ gồm: AWS S3, Google Cloud Storage (GCS),
    Azure Storage, filesystem và in-memory.
  - [pubsub](https://godoc.org/gocloud.dev/pubsub) để publish/subscribe
    tin nhắn lên một topic.
    Các nhà cung cấp được hỗ trợ gồm: Amazon SNS/SQS,
    Google Pub/Sub, Azure Service Bus, RabbitMQ và in-memory.
  - [runtimevar](https://godoc.org/gocloud.dev/runtimevar),
    để theo dõi các biến cấu hình bên ngoài.
    Các nhà cung cấp được hỗ trợ gồm AWS Parameter Store,
    Google Runtime Configurator, etcd và filesystem.
  - [secrets](https://godoc.org/gocloud.dev/secrets),
    để mã hóa/giải mã.
    Các nhà cung cấp được hỗ trợ gồm AWS KMS, GCP KMS,
    Hashicorp Vault và khóa đối xứng cục bộ.
  - Các helper để kết nối với các nhà cung cấp cloud SQL. Hỗ trợ AWS RDS và Google Cloud SQL.
  - Chúng tôi cũng đang phát triển API lưu trữ tài liệu (ví dụ: MongoDB, DynamoDB, Firestore).

## Phản hồi

Chúng tôi hy vọng bạn cũng hào hứng với Go CDK như chúng tôi. Hãy xem [godoc](https://godoc.org/gocloud.dev) của chúng tôi,
đi qua [hướng dẫn](https://github.com/google/go-cloud/tree/master/samples/tutorial),
và dùng Go CDK trong ứng dụng của bạn.
Chúng tôi rất muốn nghe ý tưởng của bạn về các API và nhà cung cấp API khác mà bạn muốn thấy.

Nếu bạn đang tìm hiểu sâu về Go CDK, hãy chia sẻ trải nghiệm với chúng tôi:

  - Điều gì diễn ra tốt?
  - Có điểm nào khó khăn khi dùng API không?
  - API bạn dùng có thiếu tính năng nào không?
  - Gợi ý để cải thiện tài liệu.

Để gửi phản hồi, bạn có thể:

  - Tạo issue trên [kho lưu trữ GitHub công khai](https://github.com/google/go-cloud/issues/new/choose).
  - Gửi email tới [go-cdk-feedback@google.com](mailto:go-cdk-feedback@google.com).
  - Đăng lên [nhóm Google công khai](https://groups.google.com/forum/#!forum/go-cloud).

Cảm ơn!
