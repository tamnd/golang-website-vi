---
title: "Go cho dịch vụ đám mây và mạng"
linkTitle: "Đám mây & Dịch vụ mạng"
description: "Với hệ sinh thái công cụ và API mạnh mẽ trên các nhà cung cấp đám mây lớn, việc xây dựng dịch vụ bằng Go chưa bao giờ dễ dàng hơn."
date: 2019-10-04T15:26:31-04:00
series: Use Cases
template: true
icon:
  file: cloud-green.svg
  alt: cloud icon
iconDark:
  file: cloud-white.svg
  alt: cloud icon
---

## Tổng quan {#overview .sectionHeading}

<div class="UseCase-halfColumn">
    <h3>Go giúp doanh nghiệp xây dựng và mở rộng hệ thống điện toán đám mây</h3>
    <p>Khi ứng dụng và quá trình xử lý chuyển lên đám mây, tính đồng thời trở thành vấn đề rất lớn. Các hệ thống điện toán đám mây, về bản chất, chia sẻ và mở rộng tài nguyên. Việc điều phối truy cập vào tài nguyên dùng chung ảnh hưởng đến mọi ứng dụng xử lý trên đám mây, và đòi hỏi các ngôn ngữ lập trình "được thiết kế đặc biệt để phát triển ứng dụng đồng thời độ tin cậy cao."</p>
  </div>

{{quote `
  author: Ruchi Malik
  title: developer at Choozle
  link: https://builtin.com/software-engineering-perspectives/golang-advantages
  quote: |
    Go makes it very easy to scale as a company. This is very important because, as our engineering team grows, each service can be managed by a different unit.
`}}

## Lợi ích chính {#key-benefits .sectionHeading}

### Cân bằng giữa tốc độ phát triển và hiệu năng máy chủ

Go được tạo ra để giải quyết chính xác những nhu cầu đồng thời này cho ứng dụng có quy mô lớn, microservice và phát triển đám mây. Thực tế, hơn 75% dự án trong Cloud Native Computing Foundation được viết bằng Go.

Go giúp giảm sự đánh đổi này nhờ thời gian build nhanh hỗ trợ phát triển lặp lại, giảm thiểu bộ nhớ và mức sử dụng CPU. Các server xây dựng bằng Go có thời gian khởi động tức thì và chi phí vận hành thấp hơn trong môi trường trả theo lưu lượng và serverless.

### Giải quyết thách thức với đám mây hiện đại, cung cấp API chuẩn idiomatic

Go giải quyết nhiều thách thức mà lập trình viên gặp phải với đám mây hiện đại: cung cấp API chuẩn idiomatic và tính đồng thời tích hợp để tận dụng bộ xử lý đa nhân. Độ trễ thấp và không cần điều chỉnh thông số của Go là điểm cân bằng tuyệt vời giữa hiệu năng và năng suất, trao cho đội kỹ thuật quyền lựa chọn và quyền di chuyển.

## Trường hợp sử dụng {#use-case .sectionHeading}

### Dùng Go cho điện toán đám mây

Điểm mạnh của Go nổi bật nhất khi xây dựng dịch vụ. Tốc độ và hỗ trợ đồng thời tích hợp mang lại dịch vụ nhanh và hiệu quả, trong khi kiểu tĩnh, công cụ mạnh mẽ và sự nhấn mạnh vào sự đơn giản và dễ đọc giúp xây dựng code đáng tin cậy và dễ bảo trì.

Go có hệ sinh thái mạnh mẽ hỗ trợ phát triển dịch vụ. [Thư viện chuẩn](/pkg/) bao gồm các package cho nhu cầu phổ biến như HTTP server và client, phân tích JSON/XML, cơ sở dữ liệu SQL, và nhiều chức năng bảo mật/mã hóa, trong khi runtime Go bao gồm các công cụ [phát hiện race](/doc/articles/race_detector.html), [benchmarking](/pkg/testing/#hdr-Benchmarks)/profiling, sinh code và phân tích code tĩnh.

Các nhà cung cấp đám mây lớn ([GCP](https://cloud.google.com/go/home), [AWS](https://aws.amazon.com/sdk-for-go/), [Azure](https://docs.microsoft.com/en-us/azure/go/)) đều có API Go cho dịch vụ của họ, và các thư viện mã nguồn mở phổ biến hỗ trợ công cụ API ([Swagger](https://github.com/go-swagger/go-swagger)), truyền tải ([protocol buffers](https://github.com/golang/protobuf), [gRPC](https://grpc.io/docs/quickstart/go/)), giám sát ([OpenCensus](https://godoc.org/go.opencensus.io)), Object-Relational Mapping ([gORM](https://gorm.io/)), và xác thực ([JWT](https://github.com/dgrijalva/jwt-go)). Cộng đồng mã nguồn mở cũng cung cấp nhiều framework dịch vụ, bao gồm [Go Kit](https://gokit.io/), [Go Micro](https://micro.mu/docs/go-micro.html) và [Gizmo](https://github.com/nytimes/gizmo), có thể là điểm khởi đầu tuyệt vời.

### Công cụ Go cho điện toán đám mây

{{toolsblurbs `
  - title: Docker
    url: https://www.docker.com/
    iconSrc: /images/logos/docker.svg
    paragraphs:
      - Docker là nền tảng-dịch-vụ phân phối phần mềm trong container. Container đóng gói phần mềm, thư viện và file cấu hình, được Docker Engine lưu trữ và chạy bởi một kernel hệ điều hành duy nhất (sử dụng ít tài nguyên hệ thống hơn máy ảo).
      - Lập trình viên đám mây dùng Docker để quản lý code Go và hỗ trợ nhiều nền tảng, vì Docker hỗ trợ quy trình phát triển và triển khai.
  - title: Kubernetes
    url: https://kubernetes.io/
    iconSrc: /images/logos/kubernetes.svg
    paragraphs:
      - Kubernetes là hệ thống điều phối container mã nguồn mở, viết bằng Go, để tự động hóa triển khai ứng dụng web. Ứng dụng web thường được đóng gói trong container với các dependency và cấu hình. Kubernetes giúp triển khai và quản lý các container đó ở quy mô lớn. Lập trình viên đám mây dùng Kubernetes để build, phân phối và mở rộng ứng dụng containerized nhanh chóng, quản lý sự phức tạp ngày càng tăng thông qua các API kiểm soát cách container chạy.
`}}

{{projects `
  - company: Google
    url: https://cloud.google.com/go
    logoSrc: google-cloud.svg
    logoSrcDark: google-cloud.svg
    desc: Google Cloud uses Go across its ecosystem of products and tools, including Kubernetes, gVisor, Knative, Istio, and Anthos. Go is fully supported on Google Cloud across all APIs and runtimes.
    ctas:
      - text: Go on Google Cloud Platform
        url: https://cloud.google.com/go
  - company: Capital One
    url: https://www.capitalone.com/
    logoSrc: capitalone_light.svg
    logoSrcDark: capitalone_dark.svg
    desc: Capital One uses Go to power the Credit Offers API, a critical service. The engineering team is also building their serverless architecture with Go, citing Go's speed and simplicity, and mentioning that "[they] didn't want to go serverless without Go."
    ctas:
      - text: Credit Offers API
        url: https://medium.com/capital-one-tech/a-serverless-and-go-journey-credit-offers-api-74ef1f9fde7f
  - company: Dropbox
    url: https://www.dropbox.com/
    logoSrc: dropbox.svg
    logoSrcDark: dropbox.svg
    desc: Dropbox was built on Python, but in 2013 decided to migrate their performance-critical backends to Go. Today, most of the company's infrastructure is written in Go.
    ctas:
      - text: Dropbox libraries
        url: https://dropbox.tech/infrastructure/open-sourcing-our-go-libraries
  - company: Mercado Libre
    url: https://www.mercadolibre.com.ar/
    logoSrc: mercadolibre_light.svg
    logoSrcDark: mercadolibre_dark.svg
    desc: MercadoLibre uses Go to scale its eCommerce platform. Go produces efficient code that readily scales as MercadoLibre's online commerce grows. Go improves their productivity while streamlining and expanding MercadoLibre services.
    ctas:
      - text: MercadoLibre & Go
        url: /solutions/mercadolibre
  - company: The New York Times
    url: https://www.nytimes.com/
    logoSrc: the-new-york-times-icon.svg
    logoSrcDark: the-new-york-times-icon.svg
    desc: The New York Times adopted Go "to build better back-end services". As the usage of Go expanded with in the company they felt the need to create a toolkit to "to help developers quickly configure and build microservice APIs and pubsub daemons", which they have open sourced.
    ctas:
      - text: NYTimes - Gizmo
        url: https://open.nytimes.com/introducing-gizmo-aa7ea463b208
      - text: Gizmo GitHub
        url: https://github.com/nytimes/gizmo
  - company: Twitch
    url: https://www.twitch.tv/
    logoSrc: twitch.svg
    logoSrcDark: twitch.svg
    desc: Twitch uses Go to power many of its busiest systems that serve live video and chat to millions of users.
    ctas:
      - text: Go's march to low-latency GC
        url: https://blog.twitch.tv/en/2016/07/05/gos-march-to-low-latency-gc-a6fa96f06eb7/
  - company: Uber
    url: https://www.uber.com/
    logoSrc: uber_light.svg
    logoSrcDark: uber_dark.svg
    desc: Uber uses Go to power several of its critical services that impact the experience of millions of drivers and passengers around the world. From their real-time analytics engine, AresDB, to their microservice for Geo-querying, Geofence, and their resource scheduler, Peloton.
    ctas:
      - text: AresDB
        url: https://eng.uber.com/aresdb/
      - text: Geofence
        url: https://eng.uber.com/go-geofence/
      - text: Peloton
        url:  https://eng.uber.com/open-sourcing-peloton/
`}}

## Bắt đầu {#get-started .sectionHeading}

### Sách Go dành cho điện toán đám mây

{{books `
  - title: Building Microservices with Go
    url: https://www.amazon.com/Building-Microservices-Go-efficient-microservices/dp/1786468662/
    thumbnail: /images/books/building-microservices-with-go.jpg
  - title: Hands-On Software Architecture with Golang
    url: https://www.amazon.com/dp/1788622596/ref=cm_sw_r_tw_dp_U_x_-aZWDbS8PD7R4
    thumbnail: /images/books/hands-on-software-architecture-with-golang.jpg
  - title: Building RESTful Web services with Go
    url: https://www.amazon.com/Building-RESTful-Web-services-gracefully-ebook/dp/B072QB8KL1
    thumbnail: /images/books/building-restful-web-services-with-go.jpg
  - title: Mastering Go Web Services
    url: https://www.amazon.com/Mastering-Web-Services-Nathan-Kozyra/dp/178398130X
    thumbnail: /images/books/mastering-go-web-services.jpg
`}}

{{libraries `
  - title: Web framework
    viewMoreUrl: https://pkg.go.dev/search?q=web+framework
    items:
      - text: Echo
        url: https://echo.labstack.com/
        desc: Web framework Go hiệu năng cao, mở rộng được và tối giản
      - text: Flamingo
        url: https://www.flamingo.me/
        desc: Framework mã nguồn mở nhanh dựa trên Go với kiến trúc sạch và có thể mở rộng
      - text: Gin
        url: https://gin-gonic.com/
        desc: Web framework viết bằng Go, với API kiểu martini
      - text: Gorilla
        url: https://www.gorillatoolkit.org/
        desc: Bộ công cụ web cho ngôn ngữ lập trình Go
  - title: Router
    viewMoreUrl: https://pkg.go.dev/search?q=http%20router
    items:
      - text: net/http
        url: https://pkg.go.dev/net/http
        desc: Package HTTP thư viện chuẩn
      - text: julienschmidt/httprouter
        url: https://pkg.go.dev/github.com/julienschmidt/httprouter?tab=overview
        desc: HTTP request router nhẹ hiệu năng cao
      - text: gorilla/mux
        url: https://pkg.go.dev/github.com/gorilla/mux?tab=overview
        desc: HTTP router mạnh mẽ và URL matcher để xây dựng Go web server
      - text: Chi
        url: https://pkg.go.dev/github.com/go-chi/chi?tab=overview
        desc: Router nhẹ, idiomatic và có thể kết hợp để xây dựng dịch vụ HTTP Go
  - title: Template engine
    viewMoreUrl: https://pkg.go.dev/search?q=templates
    items:
      - text: html/template
        url: https://pkg.go.dev/html/template
        desc: Template engine HTML thư viện chuẩn
      - text: flosch/pongo2
        url: https://pkg.go.dev/github.com/flosch/pongo2?tab=overview
        desc: Ngôn ngữ template theo cú pháp Django
  - title: Cơ sở dữ liệu & Driver
    viewMoreUrl: https://pkg.go.dev/search?q=database%20OR%20sql
    items:
      - text: database/sql
        url: https://pkg.go.dev/database/sql
        desc: Giao diện thư viện chuẩn với hỗ trợ driver cho MySQL, Postgres, Oracle, MS SQL, BigQuery và hầu hết cơ sở dữ liệu SQL
      - text: mongo-driver/mongo
        url: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo?tab=overview
        desc: Driver Go chính thức cho MongoDB
      - text: elastic/go-elasticsearch
        url: https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8?tab=overview
        desc: Elasticsearch client cho Go
      - text: GORM
        url: https://gorm.io/
        desc: Thư viện ORM cho Go
      - text: Bleve
        url: https://blevesearch.com/
        desc: Tìm kiếm toàn văn bản và lập chỉ mục cho Go
      - text: CockroachDB
        url: https://www.cockroachlabs.com/
        desc: Cơ sở dữ liệu thế hệ mới, được thiết kế cho đám mây để cung cấp SQL phân tán có khả năng phục hồi, nhất quán ở quy mô lớn
  - title: Thư viện Web
    viewMoreUrl: https://pkg.go.dev/search?q=web
    items:
      - text: markbates/goth
        url: https://pkg.go.dev/github.com/markbates/goth?tab=overview
        desc: Xác thực cho ứng dụng web
      - text: jinzhu/gorm
        url: https://pkg.go.dev/github.com/jinzhu/gorm?tab=overview
        desc: Thư viện ORM cho Go
      - text: dgrijalva/jwt-go
        url: https://pkg.go.dev/github.com/dgrijalva/jwt-go?tab=overview
        desc: Triển khai Go của JSON web token
  - title: Dự án khác
    items:
      - text: gopherjs
        url: https://pkg.go.dev/github.com/gopherjs/gopherjs?tab=overview
        desc: Trình biên dịch từ Go sang JavaScript cho phép lập trình viên viết code frontend bằng Go để chạy trên tất cả các trình duyệt.
`}}
