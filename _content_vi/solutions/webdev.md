---
title: "Go cho phát triển web"
linkTitle: "Phát triển Web"
description: "Với hiệu năng bộ nhớ được cải thiện và hỗ trợ nhiều IDE, Go tạo ra các ứng dụng web nhanh và có khả năng mở rộng cao."
date: 2019-10-04T15:26:31-04:00
series: Use Cases
template: true
books:
icon:
  file: webdev-green.svg
  alt: web dev icon
iconDark:
  file: webdev-white.svg
  alt: web dev icon
---

## Tổng quan {#overview .sectionHeading}

### Go mang lại tốc độ, bảo mật và công cụ thân thiện với lập trình viên cho ứng dụng Web

Go được thiết kế để giúp lập trình viên phát triển nhanh các ứng dụng web có thể mở rộng và bảo mật. Go đi kèm với web server dễ sử dụng, bảo mật và hiệu năng cao, đồng thời có thư viện template web riêng. Go hỗ trợ xuất sắc tất cả các công nghệ mới nhất, từ [HTTP/2](https://pkg.go.dev/net/http) đến cơ sở dữ liệu như [MySQL](https://pkg.go.dev/mod/github.com/go-sql-driver/mysql), [MongoDB](https://pkg.go.dev/mod/go.mongodb.org/mongo-driver) và [Elasticsearch](https://pkg.go.dev/mod/github.com/elastic/go-elasticsearch/v8), đến các tiêu chuẩn mã hóa mới nhất bao gồm [TLS 1.3](https://pkg.go.dev/crypto/tls). Ứng dụng web Go chạy nguyên bản trên [Google App Engine](https://cloud.google.com/appengine/) và [Google Cloud Run](https://cloud.google.com/run/) (để mở rộng dễ dàng) hoặc trên bất kỳ môi trường, đám mây hay hệ điều hành nào nhờ khả năng di động cực cao của Go.

## Lợi ích chính {#key-benefits .sectionHeading}

### Triển khai đa nền tảng với tốc độ kỷ lục

Đối với doanh nghiệp, Go được ưu tiên vì triển khai đa nền tảng nhanh chóng. Với goroutine, biên dịch gốc và không gian tên package dựa trên URI, code Go biên dịch thành một binary nhỏ duy nhất, không có dependency, rất nhanh.

### Tận dụng hiệu năng sẵn có của Go để mở rộng dễ dàng

Tigran Bayburtsyan, đồng sáng lập và CTO của Hexact Inc., tóm tắt năm lý do chính mà công ty ông chuyển sang Go:

-   **Biên dịch thành một binary duy nhất** - "Dùng static linking, Go thực sự kết hợp tất cả thư viện dependency và module thành một file binary duy nhất dựa trên loại OS và kiến trúc."

-   **Hệ thống kiểu tĩnh** - "Hệ thống kiểu thực sự quan trọng cho ứng dụng quy mô lớn."

-   **Hiệu năng** - "Go hoạt động tốt hơn nhờ mô hình đồng thời và khả năng mở rộng CPU. Bất cứ khi nào cần xử lý một request nội bộ, chúng tôi thực hiện với Goroutine riêng biệt, rẻ hơn 10 lần về tài nguyên so với Python Thread."

-   **Không cần web framework** - "Trong hầu hết các trường hợp, bạn thực sự không cần thư viện bên thứ ba nào."

-   **Hỗ trợ IDE và debug tuyệt vời** - "Sau khi viết lại tất cả dự án sang Go, chúng tôi có ít hơn 64% code so với trước đây."


{{projects `
  - company: Caddy
    url: https://caddyserver.com/
    logoSrc: caddy.svg
    logoSrcDark: caddy.svg
    desc: Caddy 2 is a powerful, enterprise-ready, open source web server with automatic HTTPS written in Go. Caddy offers greater memory safety than servers written in C. A hardened TLS stack powered by the Go standard library serves a significant portion of all Internet traffic.
    ctas:
      - text: Caddy 2
        url: https://caddyserver.com/
  - company: Cloudflare
    url: https://www.cloudflare.com/en-gb/
    logoSrc: cloudflare-icon.svg
    logoSrcDark: cloudflare-icon.svg
    desc: Cloudflare speeds up and protects millions of websites, APIs, SaaS services, and other properties connected to the Internet. "Go is at the heart of CloudFlare's services including handling compression for high-latency HTTP connections, our entire DNS infrastructure, SSL, load testing and more."
    ctas:
      - text: Cloudflare and Go
        url: https://blog.cloudflare.com/what-weve-been-doing-with-go/
  - company: gov.uk
    url: https://gov.uk/
    logoSrc: govuk_light.svg
    logoSrcDark: govuk_dark.svg
    desc: The simplicity and safety of the Go language were a good fit for the United Kingdom's government's HTTP infrastructure, and some brief experiments with the excellent net/http package convinced web developers they were on the right track. "In particular, Go's concurrency model makes it absurdly easy to build performant I/O-bound applications."
    ctas:
      - text: Building a new router for gov.uk
        url: https://technology.blog.gov.uk/2013/12/05/building-a-new-router-for-gov-uk/
      - text: Using Go in government
        url: https://technology.blog.gov.uk/2014/11/14/using-go-in-government/
  - company: Hugo
    url: https://gohugo.io/
    logoSrc: hugo.svg
    logoSrcDark: hugo.svg
    desc: Hugo is a fast and modern website engine written in Go, and designed to make website creation fun again. Websites built with Hugo are extremely fast and secure and can be hosted anywhere without any dependencies.
    ctas:
      - text: Hugo
        url: https://gohugo.io/
  - company: Mattermost
    url: https://mattermost.com/
    logoSrc: mattermost_light.svg
    logoSrcDark: mattermost_dark.svg
    desc: Mattermost is a flexible, open source messaging platform that enables secure team collaboration. It's written in Go and React.
    ctas:
      - text: Mattermost
        url: https://mattermost.com/
  - company: Medium
    url: https://medium.org/
    logoSrc: medium_light.svg
    logoSrcDark: medium_dark.svg
    desc: Medium uses Go to power their social graph, their image server and several auxiliary services. "We've found Go very easy to build, package, and deploy. We like the type-safety without the verbosity and JVM tuning of Java."
    ctas:
      - text: Medium's Go Services
        url: https://medium.engineering/how-medium-goes-social-b7dbefa6d413
  - company: The Economist
    url: https://economist.com/
    logoSrc: economist.svg
    logoSrcDark: economist.svg
    desc: The Economist needed more flexibility to deliver content to increasingly diverse digital channels. Services written in Go were a key component of the new system that would enable The Economist to deliver scalable, high performing services and quickly iterate new products. "Overall, it was determined that Go was the language best designed for usability and efficiency in a distributed, cloud-based system."
    ctas:
      - text: The Economist's Go microservices
        url: https://www.infoq.com/articles/golang-the-economist/
`}}

## Bắt đầu {#get-started .sectionHeading}

### Sách Go về phát triển web

{{books `
  - title: Web Development with Go
    url: https://www.amazon.com/Web-Development-Go-Building-Scalable-ebook/dp/B01JCOC6Z6
    thumbnail: /images/books/web-development-with-go.jpg
  - title: Go Web Programming
    url: https://www.amazon.com/Web-Programming-Sau-Sheong-Chang/dp/1617292567
    thumbnail: /images/books/go-web-programming.jpg
  - title: "Web Development Cookbook: Build full-stack web applications with Go"
    url: https://www.amazon.com/Web-Development-Cookbook-full-stack-applications-ebook/dp/B077TVQ28W
    thumbnail: /images/books/go-web-development-cookbook.jpg
  - title: Building RESTful Web services with Go
    url: https://www.amazon.com/Building-RESTful-Web-services-gracefully-ebook/dp/B072QB8KL1
    thumbnail: /images/books/building-restful-web-services-with-go.jpg
  - title: Mastering Go Web Services
    url: https://www.amazon.com/Mastering-Web-Services-Nathan-Kozyra-ebook/dp/B00W5GUKL6
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

### Khóa học
* [Học tạo ứng dụng web bằng Go](https://www.usegolang.com), khóa học trực tuyến có tính phí

### Dự án
*   {{pkg "github.com/gopherjs/gopherjs" "gopherjs"}}, trình biên dịch từ Go sang JavaScript cho phép lập trình viên viết code frontend bằng Go để chạy trên tất cả các trình duyệt.
*   [Hugo](https://gohugo.io/), framework xây dựng website nhanh nhất thế giới
*   [Mattermost](https://mattermost.com/), nền tảng nhắn tin mã nguồn mở linh hoạt cho phép cộng tác nhóm bảo mật
*   [Caddy](https://caddyserver.com/), web server mã nguồn mở mạnh mẽ, sẵn sàng cho doanh nghiệp với HTTPS tự động viết bằng Go
