---
title: Lập trình đám mây đa nền tảng với Go Cloud
date: 2018-07-24
by:
- Eno Compton
- Cassandra Salisbury
summary: Giới thiệu Go Cloud, thư viện lập trình đám mây đa nền tảng viết bằng Go.
---

## Giới thiệu

Hôm nay, nhóm Go tại Google phát hành một dự án mã nguồn mở mới,
[Go Cloud](https://github.com/google/go-cloud),
một thư viện và bộ công cụ để phát triển ứng dụng trên [nền tảng đám mây mở](https://cloud.google.com/open-cloud/).
Với dự án này, chúng tôi muốn biến Go thành ngôn ngữ được ưu tiên cho các lập trình viên
xây dựng ứng dụng đám mây đa nền tảng.

Bài viết này giải thích lý do chúng tôi khởi động dự án, cơ chế hoạt động của Go Cloud, và cách tham gia đóng góp.

## Tại sao cần lập trình đám mây đa nền tảng? Tại sao lại là bây giờ?

Chúng tôi ước tính hiện có [hơn một triệu](https://research.swtch.com/gophercount)
lập trình viên Go trên toàn thế giới.
Go là nền tảng của nhiều dự án hạ tầng đám mây quan trọng nhất,
bao gồm Kubernetes, Istio và Docker.
Các công ty như Lyft, Capital One, Netflix và [nhiều đơn vị khác](/wiki/GoUsers)
đang dùng Go trong môi trường production.
Qua nhiều năm, chúng tôi nhận thấy lập trình viên yêu thích Go cho phát triển đám mây
nhờ hiệu năng cao, năng suất tốt, hỗ trợ đồng thời sẵn có và độ trễ thấp.

Trong khuôn khổ công việc hỗ trợ sự tăng trưởng nhanh của Go,
chúng tôi đã phỏng vấn nhiều nhóm làm việc với Go để hiểu cách họ
sử dụng ngôn ngữ và cách hệ sinh thái Go có thể cải thiện hơn nữa.
Một chủ đề phổ biến ở nhiều tổ chức là nhu cầu chạy ứng dụng được trên nhiều nhà cung cấp đám mây.
Các nhóm này muốn triển khai ứng dụng vững chắc trong môi trường [đa đám mây](https://en.wikipedia.org/wiki/Cloud_computing#Multicloud)
và [đám mây lai](https://en.wikipedia.org/wiki/Cloud_computing#Hybrid_cloud),
đồng thời chuyển workload giữa các nhà cung cấp mà không cần thay đổi nhiều code.

Để đạt được điều đó, một số nhóm cố tách ứng dụng khỏi các API riêng của từng nhà cung cấp
nhằm tạo ra code đơn giản và dễ di chuyển hơn.
Tuy nhiên áp lực ngắn hạn về việc ra tính năng khiến nhiều nhóm thường bỏ qua
các nỗ lực dài hạn nhằm đảm bảo tính di động.
Kết quả là phần lớn ứng dụng Go chạy trên đám mây bị gắn chặt với nhà cung cấp đám mây ban đầu.

Go Cloud là một giải pháp thay thế: một tập hợp các API đám mây chung và mở,
giúp viết ứng dụng đám mây đơn giản và di động hơn.
Go Cloud còn đặt nền móng cho một hệ sinh thái các thư viện đám mây đa nền tảng
được xây dựng trên các API chung này.
Go Cloud giúp các nhóm vừa đạt được mục tiêu phát triển tính năng
vừa giữ được sự linh hoạt lâu dài cho kiến trúc đa đám mây và đám mây lai.
Ứng dụng Go Cloud cũng có thể chuyển sang nhà cung cấp đám mây phù hợp nhất với nhu cầu.

## Go Cloud là gì?

Chúng tôi đã xác định các dịch vụ phổ biến được ứng dụng đám mây sử dụng và tạo ra
các API chung để hoạt động trên nhiều nhà cung cấp.
Hiện tại, Go Cloud ra mắt với blob storage,
truy cập cơ sở dữ liệu MySQL, cấu hình runtime
và một HTTP server đã được cài sẵn tính năng ghi log request, tracing và health checking.
Go Cloud hỗ trợ Google Cloud Platform (GCP) và Amazon Web Services (AWS).
Chúng tôi dự kiến cộng tác với các đối tác trong ngành đám mây và cộng đồng Go để sớm
bổ sung hỗ trợ cho các nhà cung cấp khác.

Go Cloud hướng đến việc phát triển các API chung, không phụ thuộc nhà cung cấp, cho các dịch vụ được dùng nhiều nhất
để việc triển khai ứng dụng Go lên một đám mây khác trở nên đơn giản.
Go Cloud cũng tạo nền tảng để các dự án mã nguồn mở khác viết thư viện đám mây
hoạt động được trên nhiều nhà cung cấp.
Phản hồi từ cộng đồng, từ mọi loại lập trình viên ở mọi trình độ,
sẽ định hướng mức độ ưu tiên cho các API tương lai trong Go Cloud.

## Cơ chế hoạt động ra sao?

Cốt lõi của Go Cloud là tập hợp các API chung cho lập trình đám mây đa nền tảng.
Hãy xem ví dụ về sử dụng blob storage.
Bạn có thể dùng kiểu chung [`*blob.Bucket`](https://godoc.org/github.com/google/go-cloud/blob#Bucket)
để sao chép một tệp từ đĩa cục bộ lên nhà cung cấp đám mây.
Bắt đầu bằng cách mở một S3 bucket dùng [gói s3blob](https://godoc.org/github.com/google/go-cloud/blob/s3blob):

	// setupBucket mở một AWS bucket.
	func setupBucket(ctx context.Context) (*blob.Bucket, error) {
		// Lấy thông tin xác thực AWS.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-2"),
		})
		if err != nil {
			return nil, err
		}
		// Mở một handle tới s3://go-cloud-bucket.
		return s3blob.OpenBucket(ctx, sess, "go-cloud-bucket")
	}

Khi chương trình đã có `*blob.Bucket`, nó có thể tạo `*blob.Writer`,
triển khai interface `io.Writer`.
Từ đó, chương trình dùng `*blob.Writer` để ghi dữ liệu vào bucket,
kiểm tra xem `Close` có trả về lỗi hay không.

	ctx := context.Background()
	b, err := setupBucket(ctx)
	if err != nil {
		log.Fatalf("Failed to open bucket: %v", err)
	}
	data, err := ioutil.ReadFile("gopher.png")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	w, err := b.NewWriter(ctx, "gopher.png", nil)
	if err != nil {
		log.Fatalf("Failed to obtain writer: %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to bucket: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close: %v", err)
	}

Chú ý rằng logic sử dụng bucket không hề đề cập đến AWS S3.
Go Cloud biến việc chuyển đổi cloud storage thành chuyện chỉ cần thay hàm
dùng để mở `*blob.Bucket`.
Ứng dụng có thể dùng Google Cloud Storage thay thế bằng cách tạo
`*blob.Bucket` qua [`gcsblob.OpenBucket`](https://godoc.org/github.com/google/go-cloud/blob/gcsblob#OpenBucket)
mà không cần sửa code sao chép file:

	// setupBucket mở một GCS bucket.
	func setupBucket(ctx context.Context) (*blob.Bucket, error) {
		// Mở GCS bucket.
		creds, err := gcp.DefaultCredentials(ctx)
		if err != nil {
			return nil, err
		}
		c, err := gcp.NewHTTPClient(gcp.DefaultTransport(), gcp.CredentialsTokenSource(creds))
		if err != nil {
			return nil, err
		}
		// Mở một handle tới gs://go-cloud-bucket.
		return gcsblob.OpenBucket(ctx, "go-cloud-bucket", c)
	}

Mặc dù các bước để truy cập bucket trên các nhà cung cấp khác nhau,
kiểu được ứng dụng sử dụng vẫn là một: `*blob.Bucket`.
Điều này tách biệt code ứng dụng khỏi code đặc thù của từng đám mây.
Để tăng khả năng tương thích với các thư viện Go hiện có,
Go Cloud tận dụng các interface quen thuộc như `io.Writer`,
`io.Reader` và `*sql.DB`.

Code khởi tạo để truy cập dịch vụ đám mây thường theo một mẫu:
các tầng trừu tượng cao hơn được xây dựng từ các tầng cơ bản hơn.
Bạn có thể tự viết code này, nhưng Go Cloud tự động hóa việc đó bằng **Wire**,
một công cụ sinh code khởi tạo đặc thù cho từng đám mây.
[Tài liệu Wire](https://github.com/google/go-cloud/tree/master/wire)
giải thích cách cài đặt và sử dụng, còn [ví dụ Guestbook](https://github.com/google/go-cloud/tree/master/samples/guestbook)
minh họa Wire trong thực tế.

## Làm thế nào để tham gia và tìm hiểu thêm?

Để bắt đầu, chúng tôi đề nghị bạn làm theo [hướng dẫn](https://github.com/google/go-cloud/tree/master/samples/tutorial)
rồi tự thử xây dựng một ứng dụng.
Nếu bạn đang dùng AWS hoặc GCP, bạn có thể thử chuyển một phần ứng dụng hiện có sang Go Cloud.
Nếu bạn đang dùng nhà cung cấp đám mây khác hoặc dịch vụ tại chỗ,
bạn có thể mở rộng Go Cloud để hỗ trợ nó bằng cách triển khai các interface driver
(như [`driver.Bucket`](https://godoc.org/github.com/google/go-cloud/blob/driver#Bucket)).

Chúng tôi đánh giá cao mọi phản hồi của bạn về trải nghiệm sử dụng.
Việc phát triển [Go Cloud](https://github.com/google/go-cloud) được thực hiện trên GitHub.
Chúng tôi mong đợi sự đóng góp, bao gồm cả pull request.
[Tạo issue](https://github.com/google/go-cloud/issues/new) để cho chúng tôi biết
điều gì cần cải thiện hoặc các API nào dự án nên hỗ trợ trong tương lai.
Để cập nhật và thảo luận về dự án,
hãy tham gia [danh sách gửi thư của dự án](https://groups.google.com/forum/#!forum/go-cloud).

Dự án yêu cầu người đóng góp ký cùng Thỏa thuận Cấp phép Đóng góp (CLA)
như dự án Go.
Đọc [hướng dẫn đóng góp](https://github.com/google/go-cloud/blob/master/CONTRIBUTING.md) để biết thêm chi tiết.
Xin lưu ý, Go Cloud tuân thủ [Quy tắc Ứng xử](https://github.com/google/go-cloud/blob/master/CODE_OF_CONDUCT.md) của Go.

Cảm ơn bạn đã dành thời gian tìm hiểu về Go Cloud.
Chúng tôi rất hào hứng được cùng bạn biến Go thành ngôn ngữ được ưu tiên cho các lập trình viên
xây dựng ứng dụng đám mây đa nền tảng.
