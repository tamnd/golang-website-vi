---
title: Ngôn ngữ lập trình Go
summary: Go là ngôn ngữ lập trình mã nguồn mở giúp xây dựng các hệ thống bảo mật, có thể mở rộng một cách đơn giản.
template: true
---

{{$canShare := not googleCN}}

<section class="Hero bluebg">
  <div class="Hero-gridContainer">
    <div class="Hero-blurb">
      <h1>Xây dựng hệ thống đơn giản, bảo mật, có thể mở rộng với Go</h1>
      <ul class="Hero-blurbList">
        <li>
          <svg width="12" height="10" viewBox="0 0 12 10" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.8519 0.52594L3.89189 7.10404L1.14811 4.51081L0 5.59592L3.89189 9.27426L12 1.61105L10.8519 0.52594Z" fill="white" fill-opacity="0.87">
          </svg>
          Ngôn ngữ lập trình mã nguồn mở được Google hỗ trợ
        </li>
        <li>
          <svg width="12" height="10" viewBox="0 0 12 10" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.8519 0.52594L3.89189 7.10404L1.14811 4.51081L0 5.59592L3.89189 9.27426L12 1.61105L10.8519 0.52594Z" fill="white" fill-opacity="0.87">
          </svg>
          Dễ học và phù hợp cho làm việc nhóm
        </li>
        <li>
          <svg width="12" height="10" viewBox="0 0 12 10" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.8519 0.52594L3.89189 7.10404L1.14811 4.51081L0 5.59592L3.89189 9.27426L12 1.61105L10.8519 0.52594Z" fill="white" fill-opacity="0.87">
          </svg>
          Tính đồng thời tích hợp sẵn và thư viện chuẩn mạnh mẽ
        </li>
        <li>
          <svg width="12" height="10" viewBox="0 0 12 10" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.8519 0.52594L3.89189 7.10404L1.14811 4.51081L0 5.59592L3.89189 9.27426L12 1.61105L10.8519 0.52594Z" fill="white" fill-opacity="0.87">
          </svg>
          Hệ sinh thái đối tác, cộng đồng và công cụ phong phú
        </li>
      </ul>
    </div>
    <div class="Hero-actions">
      <div
        data-version=""
        class="js-latestGoVersion">
        <a class="Primary" href="/learn/" aria-label="Get Started" aria-describedby="getStarted-description" role="button">Bắt đầu</a>
        <a class="Secondary js-downloadBtn" href="/dl" aria-label="Download" aria-describedby="download-description" role="button">Tải xuống</a>
        <div class="screen-reader-only" id="getStarted-description" hidden>
          Mở cửa sổ mới với hướng dẫn Bắt đầu.
        </div>
        <div class="screen-reader-only" id="download-description" hidden>
          Mở cửa sổ mới để tải xuống Go.
        </div>
      </div>
      <div class="Hero-footnote">
        <p>
          Tải gói cho
          <a class="js-downloadWin">Windows 64-bit</a>,
          <a class="js-downloadMac">macOS</a>,
          <a class="js-downloadLinux">Linux</a>, và
          <a href="/dl/" aria-describedby="newwindow-description">nhiều hơn</a>
        </p>
        <p>
          Lệnh <code>go</code> theo mặc định tải xuống và xác thực
          module bằng Go module mirror và Go checksum database do
          Google vận hành. <a href="/dl" aria-describedby="newwindow-description">Tìm hiểu thêm.</a>
        </p>
      </div>
    </div>
    <div class="screen-reader-only" id="newwindow-description" hidden>
          Mở trong cửa sổ mới.
    </div>
    <div class="Hero-gopher">
      <img class="Hero-gopherLadder" src="/images/gophers/ladder.svg" alt="Go Gopher leo thang.">
    </div>
  </div>
</section>
<section class="WhoUses">
  <div class="WhoUses-gridContainer">
    <div class="WhoUses-header">
      <h2 class="WhoUses-headerH2">Các công ty dùng Go</h2>
      <p class="WhoUses-subheader">Các tổ chức trong mọi ngành đều dùng Go để xây dựng phần mềm và dịch vụ
        <a href="/solutions/" class="WhoUsesCaseStudyList-seeAll" aria-describedby="newwindow-description">
        Xem tất cả
       </a>
     </p>
    </div>
  <div class="WhoUsesCaseStudyList">
    <ul class="WhoUsesCaseStudyList-gridContainer">
    {{- range newest (pages "/solutions/*")}}{{if eq .series "Case Studies"}}
      {{- if .link }}
        {{- if .inLandingPageGrid }}
          <li class="WhoUsesCaseStudyList-caseStudy">
            <a href="{{.link}}" aria-label="Xem nghiên cứu điển hình của {{.company}}, (mở trong cửa sổ mới)" target="_blank" rel="noopener"
              class="WhoUsesCaseStudyList-caseStudyLink">
              <img
                loading="lazy"
                height="48"
                width="30%"
                src="/images/logos/{{.logoSrc}}"
                class="WhoUsesCaseStudyList-logo"
                alt="">
            </a>
          </li>
        {{- end}}
      {{- else}}
        <li class="WhoUsesCaseStudyList-caseStudy">
          <a href="{{.URL}}" aria-label="Xem nghiên cứu điển hình của {{.company}}, (mở trong cửa sổ mới)" class="WhoUsesCaseStudyList-caseStudyLink">
            <img
              loading="lazy"
              height="48"
              width="30%"
              src="/images/logos/{{.logoSrc}}"
              class="WhoUsesCaseStudyList-logo"
              alt="">
            <p>Xem nghiên cứu điển hình</p>
          </a>
        </li>
      {{- end}}
    {{- end}}
    {{- end}}
    </ul>
  </div>
</section>
<section class="TestimonialsGo">
  <div class="GoCarousel">
    <div class="GoCarousel-controlsContainer">
      <div class="GoCarousel-wrapper">
        <ul class="js-testimonialsGoQuotes TestimonialsGo-quotes">
          {{- range $index, $element := data "/testimonials.yaml"}}
            <li class="TestimonialsGo-quoteGroup GoCarousel-slide" id="quote_slide{{$index}}">
              <div class="TestimonialsGo-quoteSingleItem">
                <div class="TestimonialsGo-quoteSection">
                  <p class="TestimonialsGo-quote">{{raw .quote}}</p>
                  <div class="TestimonialsGo-author">— {{.name}},
                    <span class="NoWrapSpan">{{.title}}</span>
                    <span class="NoWrapSpan"> tại {{.company}}</span>
                  </div>
                </div>
              </div>
            </li>
          {{- end}}
        </ul>
      </div>
    <button class="js-testimonialsPrev GoCarousel-controlPrev" hidden>
      <i class="GoCarousel-icon material-icons">navigate_before</i>
    </button>
    <button class="js-testimonialsNext GoCarousel-controlNext">
      <i class="GoCarousel-icon material-icons">navigate_next</i>
    </button>
  </div>
  </div>
</section>
<section class="Playground">
  <div class="Playground-gridContainer">
    <div class="Playground-headerContainer">
      <h2 class="HomeSection-header">Thử Go</h2>
    </div>
    <div class="Playground-inputContainer">
      <div class="Playground-preContainer">
        Nhấn Esc để thoát khỏi trình soạn thảo.
      </div>
      <textarea class="Playground-input js-playgroundCodeEl" spellcheck="false" aria-label="Try Go" aria-describedby="editor-description" id="code">
// Bạn có thể chỉnh sửa code này!
// Click vào đây và bắt đầu gõ.
package main
import "fmt"
func main() {
  fmt.Println("Hello, 世界")
}</textarea>
    </div>
    <div class="screen-reader-only" id="editor-description" hidden>
      Nhấn Esc để thoát khỏi trình soạn thảo.
    </div>
    <div class="Playground-outputContainer js-playgroundOutputEl">
      <pre class="Playground-output"><noscript>Hello, 世界</noscript></pre>
    </div>
    <div class="Playground-controls">
      <select class="Playground-selectExample js-playgroundToysEl" aria-label="Code examples">
      <option value="hello.go">Hello, World!</option>
      <option value="life.go">Conway's Game of Life</option>
      <option value="fib.go">Fibonacci Closure</option>
      <option value="peano.go">Peano Integers</option>
      <option value="pi.go">Concurrent pi</option>
      <option value="sieve.go">Concurrent Prime Sieve</option>
      <option value="solitaire.go">Peg Solitaire Solver</option>
      <option value="tree.go">Tree Comparison</option>
      </select>
      <div class="Playground-buttons">
      <button class="Button Button--primary js-playgroundRunEl Playground-runButton" title="Chạy code này [shift-enter]">Chạy</button>
      <div class="Playground-secondaryButtons">
        {{- if $canShare}}
        <button class="Button js-playgroundShareEl Playground-button" title="Chia sẻ trên Go Playground">Chia sẻ</button>
        {{- end}}
        <a class="Button tour Playground-button" href="/tour/" title="Tour Go từ trình duyệt">Tour</a>
      </div>
      </div>
    </div>
  </div>
</section>
<section class="WhyGo">
  <div class="WhyGo-gridContainer">
    <div class="WhyGo-header">
      <h2 class="WhyGo-headerH2">Có thể làm gì với Go</h2>
      <p class="WhyGo-subheader">
        Dùng Go cho nhiều mục đích phát triển phần mềm khác nhau
      </p>
    </div>
    <ul class="WhyGo-reasons">
      {{- range first 4 (data "/resources.yaml")}}
        <li class="WhyGo-reason">
          <div class="WhyGo-reasonDetails">
            <div class="WhyGo-reasonIcon" role="presentation">
              <img class="DarkMode-img" src="{{.iconDark}}" alt="{{.iconName}}">
              <img class="LightMode-img" src="{{.icon}}" alt="{{.iconName}}">
            </div>
            <div class="WhyGo-reasonText">
              <h3 class="WhyGo-reasonTitle">{{.title}}</h3>
              <p>
                {{.description}}
              </p>
            </div>
          </div>
          <div class="WhyGo-reasonFooter">
            <div class="WhyGo-reasonPackages">
              <div class="WhyGo-reasonPackagesHeader">
                <img src="/images/icons/package.svg" alt="Packages.">
                Package phổ biến:
              </div>
              <ul class="WhyGo-reasonPackagesList">
                {{- range .packages }}
                  <li class="WhyGo-reasonPackage">
                    <a class="WhyGo-reasonLink" href="{{.url}}" target="_blank" rel="noopener">
                      {{.title}}
                    </a>
                  </li>
                  {{- end}}
              </ul>
            </div>
            <div class="WhyGo-reasonLearnMoreLink">
              <a href="{{.link}}" aria-describedby="newwindow-description">Tìm hiểu thêm
              <i class="material-icons WhyGo-forwardArrowIcon" aria-hidden="true">arrow_forward</i></a>
            </div>
          </div>
        </li>
      {{- end}}
      {{- if gt (len (data "resources.yaml")) 3}}
        <li class="WhyGo-reason">
          <div class="WhyGo-reasonShowMore">
            <div class="WhyGo-reasonShowMoreImgWrapper">
              <img
                class="WhyGo-reasonShowMoreImg"
                loading="lazy"
                height="148"
                width="229"
                src="/images/gophers/biplane.svg"
                alt="Go Gopher đang trượt ván.">
            </div>
            <div class="WhyGo-reasonShowMoreLink">
              <a href="/solutions/use-cases" aria-describedby="newwindow-description">Thêm trường hợp sử dụng
              <i class="material-icons
              WhyGo-forwardArrowIcon" aria-hidden="true">arrow_forward</i></a>
            </div>
          </div>
        </li>
      {{- end}}
    </ul>
  </div>
</section>
<section class="GettingStartedGo">
  <div class="GettingStartedGo-gridContainer">
    <div class="GettingStartedGo-header">
      <h2 class="GettingStartedGo-headerH2">Bắt đầu với Go</h2>
      <p class="GettingStartedGo-headerDesc">
        Khám phá kho tài nguyên học tập phong phú, bao gồm các hành trình có hướng dẫn, khóa học, sách và nhiều hơn nữa.
      </p>
      <div class="GettingStartedGo-ctas">
        <a class="GettingStartedGo-primaryCta" href="/learn/"aria-describedby="newwindow-description">Bắt đầu</a>
        <a href="/doc/install/" aria-describedby="newwindow-description">Tải xuống Go</a>
      </div>
    </div>
    <div class="GettingStartedGo-resourcesSection">
      <ul class="GettingStartedGo-resourcesList">
        <li class="GettingStartedGo-resourcesHeader">
          Tài nguyên để tự học
        </li>
        <li class="GettingStartedGo-resourceItem">
          <a href="/learn#guided-learning-journeys" class="GettingStartedGo-resourceItemTitle" aria-describedby="newwindow-description">
            Hành trình học có hướng dẫn
          </a>
          <div class="GettingStartedGo-resourceItemDescription">
            Hướng dẫn từng bước để làm quen
          </div>
        </li>
        <li class="GettingStartedGo-resourceItem">
          <a href="/learn#online-learning" class="GettingStartedGo-resourceItemTitle" aria-describedby="newwindow-description">
            Học trực tuyến
          </a>
          <div class="GettingStartedGo-resourceItemDescription">
            Duyệt tài nguyên và học theo tốc độ của bạn
          </div>
        </li>
        <li class="GettingStartedGo-resourceItem">
          <a href="/learn#featured-books" class="GettingStartedGo-resourceItemTitle" aria-describedby="newwindow-description">
            Sách nổi bật
          </a>
          <div class="GettingStartedGo-resourceItemDescription">
            Đọc qua các chương có cấu trúc và lý thuyết
          </div>
        </li>
        <li class="GettingStartedGo-resourceItem">
          <a href="/learn#self-paced-labs" class="GettingStartedGo-resourceItemTitle" aria-describedby="newwindow-description">
            Lab tự học trên Cloud
          </a>
          <div class="GettingStartedGo-resourceItemDescription">
            Bắt tay vào triển khai ứng dụng Go trên GCP
          </div>
        </li>
      </ul>
      <ul class="GettingStartedGo-resourcesList">
        <li class="GettingStartedGo-resourcesHeader">
          Đào tạo trực tiếp
        </li>
        {{- range first 4 (data "/learn/training.yaml")}}
          <li class="GettingStartedGo-resourceItem">
            <a href="{{.url}}" class="GettingStartedGo-resourceItemTitle" aria-describedby="newwindow-description">
              {{.title}}
            </a>
            <div class="GettingStartedGo-resourceItemDescription">
              {{.blurb}}
            </div>
          </li>
        {{- end}}
      </ul>
    </div>
  </div>
</section>
<script src="/js/index.js" defer></script>
