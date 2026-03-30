---
title: "Giao diện dòng lệnh (CLI)"
linkTitle: "Giao diện dòng lệnh (CLI)"
description: "Với nhiều package mã nguồn mở phổ biến và thư viện chuẩn mạnh mẽ, hãy dùng Go để tạo các CLI nhanh và tinh tế."
date: 2019-10-04T15:26:31-04:00
series: Use Cases
template: true
icon:
  file: clis-green.svg
  alt: CLI icon
iconDark:
  file: clis-white.svg
  alt: CLI icon
---

## Tổng quan {#overview .sectionHeading}

### Lập trình viên CLI ưa chuộng Go vì tính di động, hiệu năng và dễ tạo

Giao diện dòng lệnh (CLI), khác với giao diện đồ họa (GUI), chỉ hoạt động trên văn bản. Các ứng dụng đám mây và hạ tầng chủ yếu dùng CLI nhờ khả năng tự động hóa và truy cập từ xa dễ dàng.

## Lợi ích chính {#key-benefits .sectionHeading}

### Tận dụng thời gian biên dịch nhanh để xây dựng chương trình khởi động ngay và chạy trên mọi hệ thống

Các lập trình viên CLI nhận thấy Go rất lý tưởng để thiết kế ứng dụng. Go biên dịch rất nhanh thành một binary duy nhất, hoạt động nhất quán trên nhiều nền tảng và có cộng đồng phát triển mạnh. Từ một chiếc laptop Windows hay Mac, lập trình viên có thể build chương trình Go cho hàng chục kiến trúc và hệ điều hành mà Go hỗ trợ chỉ trong vài giây, không cần các build farm phức tạp. Không ngôn ngữ biên dịch nào khác có thể đóng gói di động và nhanh chóng như vậy. Ứng dụng Go được đóng gói thành một binary độc lập, giúp việc cài đặt trở nên cực kỳ đơn giản.

Cụ thể, **các chương trình viết bằng Go chạy trên mọi hệ thống mà không cần thư viện, runtime hay dependency nào**. Và **các chương trình viết bằng Go có thời gian khởi động tức thì**, tương đương C hay C++ nhưng không thể đạt được với các ngôn ngữ lập trình khác.

## Trường hợp sử dụng {#use-case .sectionHeading}

### Dùng Go để xây dựng CLI tinh tế

{{backgroundquote `
  author: Steve Domino
  title: senior engineer and architect at Strala
  link: https://medium.com/@skdomino/writing-better-clis-one-snake-at-a-time-d22e50e60056
  quote: |
    I was tasked with building our CLI tool and found two really great projects, Cobra and Viper, which make building CLI's easy. Individually they are very powerful, very flexible and very good at what they do. But together they will help you show your next CLI who is boss!
`}}

{{backgroundquote `
  author: Francesc Campoy
  title: VP of product at DGraph Labs and producer of Just For Func videos
  link: https://www.youtube.com/watch?v=WvWPGVKLvR4
  quote: |
    Cobra is a great product to write small tools or even large ones. It's more of a framework than a library, because when you call the binary that would create a skeleton, then you would be adding code in between."
`}}

Khi phát triển CLI bằng Go, hai công cụ được sử dụng rộng rãi là: Cobra và Viper.

{{pkg "github.com/spf13/cobra" "Cobra"}} vừa là thư viện để tạo các ứng dụng CLI hiện đại mạnh mẽ, vừa là chương trình để sinh ra ứng dụng và CLI trong Go. Cobra hỗ trợ hầu hết các ứng dụng Go phổ biến như CoreOS, Delve, Docker, Dropbox, Git Lfs, Hugo, Kubernetes và [nhiều ứng dụng khác](https://pkg.go.dev/github.com/spf13/cobra?tab=importedby). Với tính năng hỗ trợ trợ giúp lệnh, tự động hoàn thành và tài liệu tích hợp, "[nó] giúp việc viết tài liệu cho từng lệnh trở nên thực sự đơn giản," theo lời [Alex Ellis](https://blog.alexellis.io/5-keys-to-a-killer-go-cli/), người sáng lập OpenFaaS.


{{pkg "github.com/spf13/viper" "Viper"}} là giải pháp cấu hình toàn diện cho ứng dụng Go, được thiết kế để xử lý nhu cầu và định dạng cấu hình ngay trong ứng dụng. Cobra và Viper được thiết kế để hoạt động cùng nhau.

Viper [hỗ trợ cấu trúc lồng nhau](https://scene-si.org/2017/04/20/managing-configuration-with-viper/) trong cấu hình, cho phép lập trình viên CLI quản lý cấu hình cho nhiều phần của ứng dụng lớn. Viper cũng cung cấp tất cả các công cụ cần thiết để dễ dàng xây dựng ứng dụng twelve-factor.

"Nếu bạn không muốn làm lộn xộn dòng lệnh, hoặc bạn đang làm việc với dữ liệu nhạy cảm mà bạn không muốn xuất hiện trong lịch sử, thì nên dùng biến môi trường. Để làm điều này, bạn có thể dùng Viper," [Geudens gợi ý](https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html).

{{projects `
  - company: Comcast
    url: https://xfinity.com/
    logoSrc: comcast.svg
    logoSrcDark: comcast.svg
    desc: Comcast uses Go for a CLI client used to publish and subscribe to its high-traffic sites. The company also supports an open source client library which is written in Go - designed for working with Apache Pulsar.
    ctas:
      - text: Client library for Apache Pulsar
        url: https://github.com/Comcast/pulsar-client-go
      - text: Pulsar CLI Client
        url: https://github.com/Comcast/pulsar-client-go/blob/master/cli/main.go
  - company: GitHub
    url: https://github.com/
    logoSrc: github.svg
    logoSrcDark: github.svg
    desc: GitHub uses Go for a command-line tool that makes it easier to work with GitHub, wrapping git in order to extend it with extra features and commands.
    ctas:
      - text: GitHub command-line tool
        url: https://github.com/cli/cli
  - company: Hugo
    url: https://gohugo.io/
    logoSrc: hugo.svg
    logoSrcDark: hugo.svg
    desc: Hugo is one of the most popular Go CLI applications powering thousands of sites, including this one. One reason for its popularity is its ease of install thanks to Go. Hugo author Bjørn Erik Pedersen writes "The single binary takes most of the pain out of installation and upgrades."
    ctas:
      - text: Hugo Website
        url: https://gohugo.io/
  - company: Kubernetes
    url: https://kubernetes.com/
    logoSrc: kubernetes.svg
    logoSrcDark: kubernetes.svg
    desc: Kubernetes is one of the most popular Go CLI applications. Kubernetes Creator, Joe Beda, said that for writing Kubernetes, "Go was the only logical choice". Calling Go "the sweet spot" between low level languages like C++ and high level languages like Python.
    ctas:
      - text: Kubernetes + Go
        url: https://blog.gopheracademy.com/birthday-bash-2014/kubernetes-go-crazy-delicious/
  - company: MongoDB
    url: https://mongodb.com/
    logoSrc: mongodb.svg
    logoSrcDark: mongodb.svg
    desc: MongoDB chose to implement their Backup CLI Tool in Go citing Go's "C-like syntax, strong standard library, the resolution of concurrency problems via goroutines, and painless multi-platform distribution" as reasons.
    ctas:
      - text: MongoDB Backup Service
        url: https://www.mongodb.com/blog/post/go-agent-go
  - company: Netflix
    url: https://netflix.com/
    logoSrc: netflix.svg
    logoSrcDark: netflix.svg
    desc: Netflix uses Go to build the CLI application ChaosMonkey, an application responsible for randomly terminating instances in production to ensure that engineers implement their services to be resilient to instance failures.
    ctas:
      - text: Netflix Techblog Article
        url: https://medium.com/netflix-techblog/application-data-caching-using-ssds-5bf25df851ef
  - company: Stripe
    url: https://stripe.com/
    logoSrc: stripe.svg
    logoSrcDark: stripe.svg
    desc: Stripe uses Go for the Stripe CLI aimed to help build, test, and manage a Stripe integration right from the terminal.
    ctas:
      - text: Stripe CLI
        url: https://github.com/stripe/stripe-cli
  - company: Uber
    url: https://uber.com/
    logoSrc: uber.svg
    logoSrcDark: uber.svg
    desc: Uber uses Go for several CLI tools, including the CLI API for Jaeger, a distributed tracing system used for monitoring microservice distributed systems.
    ctas:
      - text: CLI API for Jaeger
        url: https://www.jaegertracing.io/docs/1.14/cli/
`}}

## Bắt đầu {#get-started .sectionHeading}

### Sách Go dành cho việc tạo CLI

{{books `
  - title: Powerful Command-Line Applications in Go
    url: https://www.amazon.com/Powerful-Command-Line-Applications-Go-Maintainable/dp/168050696X
    thumbnail: /images/books/powerful-command-line-applications-in-go.jpg
  - title: Go in Action
    url: https://www.amazon.com/Go-Action-William-Kennedy/dp/1617291781
    thumbnail: /images/books/go-in-action.jpg
  - title: The Go Programming Language
    url: https://www.gopl.io/
    thumbnail: /images/learn/go-programming-language-book.png
  - title: Go Programming Blueprints
    url: https://github.com/matryer/goblueprints
    thumbnail: /images/learn/go-programming-blueprints.png
`}}

{{libraries `
  - title: Thư viện CLI
    viewMoreUrl: https://pkg.go.dev/search?q=command%20line%20OR%20CLI
    items:
      - text: spf13/cobra
        url: https://pkg.go.dev/github.com/spf13/cobra?tab=overview
        desc: Thư viện để tạo các ứng dụng CLI hiện đại mạnh mẽ và chương trình sinh ứng dụng CLI trong Go
      - text: spf13/viper
        url: https://pkg.go.dev/github.com/spf13/viper?tab=overview
        desc: Giải pháp cấu hình toàn diện cho ứng dụng Go, xử lý nhu cầu và định dạng cấu hình ngay trong ứng dụng
      - text: urfave/cli
        url: https://pkg.go.dev/github.com/urfave/cli?tab=overview
        desc: Framework tối giản để tạo và tổ chức ứng dụng dòng lệnh Go
      - text: delve
        url: https://pkg.go.dev/github.com/go-delve/delve?tab=overview
        desc: Công cụ đơn giản và mạnh mẽ cho lập trình viên dùng debugger cấp nguồn trong ngôn ngữ biên dịch
      - text: chzyer/readline
        url: https://pkg.go.dev/github.com/chzyer/readline?tab=overview
        desc: Triển khai Go thuần túy cung cấp hầu hết tính năng của GNU Readline (giấy phép MIT)
      - text: dixonwille/wmenu
        url: https://pkg.go.dev/github.com/dixonwille/wmenu?tab=overview
        desc: Cấu trúc menu dễ sử dụng cho ứng dụng CLI để nhắc người dùng lựa chọn
      - text: spf13/pflag
        url: https://pkg.go.dev/github.com/spf13/pflag?tab=overview
        desc: Thay thế drop-in cho package flag của Go, triển khai cờ kiểu POSIX/GNU
      - text: golang/glog
        url: https://pkg.go.dev/github.com/golang/glog?tab=overview
        desc: Log thực thi phân cấp cho Go
      - text: go-prompt
        url: https://pkg.go.dev/github.com/c-bata/go-prompt?tab=overview
        desc: Thư viện xây dựng prompt tương tác mạnh mẽ, giúp tạo công cụ dòng lệnh đa nền tảng dễ dàng hơn.
`}}
