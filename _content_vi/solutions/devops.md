---
title: "DevOps & Site Reliability Engineering"
linkTitle: "DevOps & Site Reliability Engineering"
description: "Với thời gian build nhanh, cú pháp gọn gàng, trình định dạng và sinh tài liệu tự động, Go được thiết kế để hỗ trợ cả DevOps và SRE."
date: 2019-10-03T17:16:43-04:00
series: Use Cases
template: true
books:
icon:
  file: devops-green.svg
  alt: ops icon
iconDark:
  file: devops-white.svg
  alt: ops icon
---

## Tổng quan {#overview .sectionHeading}

### Go giúp doanh nghiệp tự động hóa và mở rộng quy mô

Các nhóm Development Operations (DevOps) giúp tổ chức kỹ thuật tự động hóa tác vụ và cải thiện quy trình tích hợp liên tục, phân phối và triển khai liên tục (CI/CD). DevOps có thể phá vỡ các rào cản phát triển và triển khai công cụ cùng tự động hóa để nâng cao phát triển, triển khai và hỗ trợ phần mềm.

Site Reliability Engineering (SRE) ra đời tại Google để làm cho "các site quy mô lớn của công ty trở nên đáng tin cậy, hiệu quả và có khả năng mở rộng hơn," [Silvia Fressard viết](https://opensource.com/article/18/10/what-site-reliability-engineer), một chuyên gia tư vấn DevOps độc lập. "Và các phương pháp họ phát triển đáp ứng tốt nhu cầu của Google đến mức các công ty công nghệ lớn khác như Amazon và Netflix cũng áp dụng." SRE đòi hỏi sự kết hợp giữa kỹ năng phát triển và vận hành, và "[trao quyền cho lập trình viên phần mềm](https://stackify.com/site-reliability-engineering/) quản lý hoạt động hàng ngày của ứng dụng trong môi trường production."

Go phục vụ cả hai vai trò DevOps và SRE, từ thời gian build nhanh và cú pháp gọn gàng đến hỗ trợ bảo mật và độ tin cậy. Các tính năng đồng thời và mạng của Go cũng làm cho nó lý tưởng cho các công cụ quản lý triển khai đám mây, dễ dàng hỗ trợ tự động hóa trong khi mở rộng về tốc độ và khả năng bảo trì code khi cơ sở hạ tầng phát triển.

Các nhóm DevOps/SRE viết phần mềm từ script nhỏ, đến giao diện dòng lệnh (CLI), đến tự động hóa và dịch vụ phức tạp, và bộ tính năng của Go mang lại lợi ích cho mọi tình huống.

## Lợi ích chính {#key-benefits .sectionHeading}

### Dễ dàng xây dựng script nhỏ với thư viện chuẩn mạnh mẽ và kiểu tĩnh của Go
Thời gian build và khởi động nhanh của Go. Thư viện chuẩn phong phú của Go, bao gồm các package cho nhu cầu phổ biến như HTTP, I/O file, thời gian, biểu thức chính quy, exec và định dạng JSON/CSV, giúp nhóm DevOps/SRE tập trung thẳng vào logic nghiệp vụ. Thêm vào đó, hệ thống kiểu tĩnh và xử lý lỗi rõ ràng của Go làm cho ngay cả script nhỏ cũng trở nên mạnh mẽ hơn.

### Triển khai CLI nhanh chóng với thời gian build nhanh của Go
Mọi site reliability engineer đều đã viết script "dùng một lần" rồi trở thành CLI được hàng chục kỹ sư khác sử dụng mỗi ngày. Và các script tự động hóa triển khai nhỏ trở thành dịch vụ quản lý triển khai. Với Go, DevOps/SRE ở vị trí tốt để thành công khi phạm vi phần mềm chắc chắn mở rộng. Bắt đầu với Go đặt bạn vào vị trí tốt để thành công khi điều đó xảy ra.

### Mở rộng và duy trì ứng dụng lớn hơn với bộ nhớ thấp và sinh tài liệu của Go
Bộ gom rác của Go có nghĩa là nhóm DevOps/SRE không phải lo lắng về quản lý bộ nhớ. Và bộ sinh tài liệu tự động của Go (godoc) làm cho code tự tài liệu hóa, giảm chi phí bảo trì và thiết lập các phương pháp tốt nhất ngay từ đầu.

{{projects `
  - company: Docker
    url: https://docker.com/
    logoSrc: docker.svg
    logoSrcDark: docker.svg
    desc: Docker is a software-as-a-service (SaaS) product, written in Go, that DevOps/SRE teams leverage to "drive secure automation and deployment at massive scale," supporting their CI/CD efforts.
    ctas:
      - text: Docker CI/CD
        url: https://www.docker.com/solutions/cicd
  - company: Drone
    url: https://github.com/drone
    logoSrc: drone.svg
    logoSrcDark: drone.svg
    desc: Drone is a Continuous Delivery system built on container technology, written in Go, that uses a simple YAML configuration file, a superset of docker-compose, to define and execute Pipelines inside Docker containers.
    ctas:
      - text: Drone
        url: https://github.com/drone
  - company: etcd
    url: https://github.com/etcd-io/etcd
    logoSrc: etcd.svg
    logoSrcDark: etcd.svg
    desc: etcd is a strongly consistent, distributed key-value store that provides a reliable way to store data that needs to be accessed by a distributed system or cluster of machines, and it's written in Go.
    ctas:
      - text: etcd
        url: https://github.com/etcd-io/etcd
  - company: IBM
    url: https://ibm.com/
    logoSrc: ibm.svg
    logoSrcDark: ibm.svg
    desc: IBM's DevOps teams use Go through Docker and Kubernetes, plus other DevOps and CI/CD tools written in Go. The company also supports connection to it's messaging middleware through a Go-specific API.
    ctas:
      - text: IBM Applications in Golang
        url: https://developer.ibm.com/messaging/2019/02/05/simplified-ibm-mq-applications-golang/
  - company: Netflix
    url: https://netflix.com/
    logoSrc: netflix.svg
    logoSrcDark: netflix.svg
    desc: Netflix uses Go to handle large scale data caching, with a service called Rend, which manages globally replicated storage for personalization data.
    ctas:
      - text: Application Data Caching
        url: https://medium.com/netflix-techblog/application-data-caching-using-ssds-5bf25df851ef
      - text: Rend
        url: https://github.com/netflix/rend
  - company: Microsoft
    url: https://microsoft.com/
    logoSrc: microsoft_light.svg
    logoSrcDark: microsoft_dark.svg
    desc: Microsoft uses Go in Azure Red Hat OpenShift services. This Microsoft solution provides DevOps teams with OpenShift clusters to maintain regulatory compliance and focus on application development.
    ctas:
      - text: OpenShift
        url: https://azure.microsoft.com/en-us/services/openshift/
  - company: Terraform
    url: https://terraform.io/
    logoSrc: terraform-icon.svg
    logoSrcDark: terraform-icon.svg
    desc: Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently. It supports a number of cloud providers such as AWS, IBM Cloud, GCP, and Microsoft Azure - and it's written in Go.
    ctas:
      - text: Terraform
        url: https://www.terraform.io/intro/index.html
  - company: Prometheus
    url: https://github.com/prometheus/prometheus
    logoSrc: prometheus.svg
    logoSrcDark: prometheus.svg
    desc: Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud. Most Prometheus components are written in Go, making them easy to build and deploy as static binaries.
    ctas:
      - text: Prometheus
        url: https://github.com/prometheus/prometheus
  - company: YouTube
    url: https://youtube.com/
    logoSrc: youtube.svg
    logoSrcDark: youtube.svg
    desc: YouTube uses Go with Vitess (now part of PlanetScale), its database clustering system for horizontal scaling of MySQL through generalized sharding. Since 2011 it's been a core component of YouTube's database infrastructure, and has grown to encompass tens of thousands of MySQL nodes.
    ctas:
      - text: Vitess
        url: https://github.com/vitessio/vitess
`}}

## Bắt đầu {#get-started .sectionHeading}

### Sách Go về DevOps & SRE

{{books `
  - title: Go Programming for Network Operations
    url: https://www.amazon.com/Go-Programming-Network-Operations-Automation-ebook/dp/B07JKKN34L/ref=sr_1_16
    thumbnail: /images/books/go-programming-for-network-operations.jpg
  - title: Go Programming Blueprints
    url: https://github.com/matryer/goblueprints
    thumbnail: /images/learn/go-programming-blueprints.png
  - title: Go in Action
    url: https://www.amazon.com/Go-Action-William-Kennedy/dp/1617291781
    thumbnail: /images/books/go-in-action.jpg
  - title: The Go Programming Language
    url: https://www.gopl.io/
    thumbnail: /images/learn/go-programming-language-book.png
`}}

{{libraries `
  - title: Giám sát và tracing
    viewMoreUrl: https://pkg.go.dev/search?q=tracing
    items:
      - text: open-telemetry/opentelemetry-go
        url: https://pkg.go.dev/go.opentelemetry.io/otel
        desc: API và công cụ đo lường trung lập với nhà cung cấp để giám sát và distributed tracing
      - text: jaegertracing/jaeger-client-go
        url: https://pkg.go.dev/github.com/jaegertracing/jaeger-client-go?tab=overview
        desc: Hệ thống distributed tracing mã nguồn mở được phát triển bởi Uber
      - text: grafana/grafana
        url: https://pkg.go.dev/github.com/grafana/grafana?tab=overview
        desc: Nền tảng mã nguồn mở để giám sát và quan sát
      - text: istio/istio
        url: https://pkg.go.dev/github.com/istio/istio?tab=overview
        desc: Service mesh mã nguồn mở và nền tảng có thể tích hợp
  - title: Thư viện CLI
    viewMoreUrl: https://pkg.go.dev/search?q=command%20line%20OR%20CLI
    items:
      - text: spf13/cobra
        url: https://pkg.go.dev/github.com/spf13/cobra?tab=overview
        desc: Thư viện tạo ứng dụng CLI hiện đại mạnh mẽ và chương trình sinh ứng dụng CLI trong Go
      - text: spf13/viper
        url: https://pkg.go.dev/github.com/spf13/viper?tab=overview
        desc: Giải pháp cấu hình toàn diện cho ứng dụng Go, xử lý nhu cầu và định dạng cấu hình ngay trong ứng dụng
      - text: urfave/cli
        url: https://pkg.go.dev/github.com/urfave/cli?tab=overview
        desc: Framework tối giản để tạo và tổ chức ứng dụng dòng lệnh Go
  - title: Dự án khác
    items:
      - text: golang-migrate/migrate
        url: https://pkg.go.dev/github.com/golang-migrate/migrate?tab=overview
        desc: Công cụ migration cơ sở dữ liệu viết bằng Go
`}}
