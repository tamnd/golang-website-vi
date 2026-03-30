---
title: "Bắt đầu"
breadcrumbTitle: "Học"
template: true
---

<section class="Learn-hero">
  <div class="Container">
    <div class="Learn-heroInner">
      <div class="Learn-heroContent">
        {{breadcrumbs .}}
        <h1>Cài đặt phiên bản Go mới nhất</h1>
        <p>
          Cài đặt phiên bản Go mới nhất. Để biết hướng dẫn tải về và cài đặt
          trình biên dịch, công cụ và thư viện Go,
          <a href="/doc/install" target="_blank" rel="noopener">
            xem tài liệu cài đặt.
          </a>
        </p>
        <div class="Learn-heroAction">
          <div
            data-version=""
            class="js-latestGoVersion"
          >
            <a
              class="js-downloadBtn"
              href="/dl"
              target="_blank"
              rel="noopener"
            >
              <span class="GoVersionSpan">Tải về</span>
            </a>
          </div>
        </div>
      </div>
      <div class="Learn-heroGopher">
        <img src="/images/gophers/motorcycle.svg" alt="Go Gopher đang lái xe máy">
      </div>
    </div>
  </div>
</section>

<div class="Learn-columns">
  <aside class="Learn-sidebar">
    <nav class="LearnNav">
      <a class="active" href="#selected-tutorials">
        <svg width="5" height="5" viewBox="0 0 5 5" fill="none" xmlns="http://www.w3.org/2000/svg"><circle cx="2.5" cy="2.5" r="2.5" fill="#007F9F"/></svg>
        <span>Hướng dẫn được chọn lọc</span>
      </a>
      <a href="#guided-learning-journeys">
      <svg width="5" height="5" viewBox="0 0 5 5" fill="none" xmlns="http://www.w3.org/2000/svg"><circle cx="2.5" cy="2.5" r="2.5" fill="#007F9F"/></svg>
      <span>Lộ trình học có hướng dẫn</span>
      </a>
      <a href="#self-paced-labs">
      <svg width="5" height="5" viewBox="0 0 5 5" fill="none" xmlns="http://www.w3.org/2000/svg"><circle cx="2.5" cy="2.5" r="2.5" fill="#007F9F"/></svg>
      <span>Qwiklabs</span>
      </a>
      <a href="#tutorials">
      <svg width="5" height="5" viewBox="0 0 5 5" fill="none" xmlns="http://www.w3.org/2000/svg"><circle cx="2.5" cy="2.5" r="2.5" fill="#007F9F"/></svg>
      <span>Hướng dẫn</span>
      </a>
      <a href="#training">
      <svg width="5" height="5" viewBox="0 0 5 5" fill="none" xmlns="http://www.w3.org/2000/svg"><circle cx="2.5" cy="2.5" r="2.5" fill="#007F9F"/></svg>
      <span>Đào tạo</span>
      </a>
      <a href="#featured-books">
      <svg width="5" height="5" viewBox="0 0 5 5" fill="none" xmlns="http://www.w3.org/2000/svg"><circle cx="2.5" cy="2.5" r="2.5" fill="#007F9F"/></svg>
      <span>Sách</span>
      </a>
    </nav>
  </aside>
  <div class="Learn-body">
  <section id="selected-tutorials" class="Learn-tutorials">
    <div class="Container">
      <div class="Learn-learningResourcesHeader">
          <h3>Hướng dẫn được chọn lọc</h3>
          <p>Mới bắt đầu với Go và chưa biết bắt đầu từ đâu?</p>
      </div>
      <div class="LearnGo-gridContainer">
        <ul class="Learn-cardList">
        {{- range first 3 (data "quickstart.yaml")}}
            <li class="Learn-card">
            {{- template "learn-card" . }}
            </li>
        {{- end}}
        </ul>
      </div>
    </div>
  </section>

  <section id="guided-learning-journeys" class="Learn-guided">
    <div class="Container">
      <div class="Learn-learningResourcesHeader">
        <h3>Lộ trình học có hướng dẫn</h3>
        <p>Đã nắm được kiến thức cơ bản và muốn tìm hiểu thêm?</p>
      </div>
      <div class="LearnGo-gridContainer">
        <ul class="Learn-cardList">
          {{- range first 4 (data "guided.yaml")}}
            <li class="Learn-card">
              {{- template "learn-card" .}}
            </li>
          {{- end}}
        </ul>
      </div>
    </div>
  </section>

  <section id="self-paced-labs" class="Learn-selfPaced">
    <div class="Container">
      <div class="Learn-learningResourcesHeader">
        <h3>Qwiklabs</h3>
        <p>Tham quan có hướng dẫn các chương trình Go</p>
      </div>
      <div class="LearnGo-gridContainer">
        <ul class="Learn-cardList">
          {{- range first 3 (data "cloud.yaml")}}
          <li class="Learn-card">
            {{- template "learn-self-paced-card" .}}
          </li>
          </li>
          {{- end}}
        </ul>
      </div>
    </div>
  </section>

  <section id="tutorials" class="Learn-tutorials">
    <div class="Container">
      <div class="Learn-learningResourcesHeader">
        <h3>Hướng dẫn</h3>
        <p></p>
      </div>
      <div class="LearnGo-gridContainer">
        <ul class="Learn-cardList">
          {{- range first 3 (data "tutorials.yaml") }}
            <li class="Learn-card">
              {{- template "learn-card" .}}
            </li>
          {{- end}}
        </ul>
      </div>
    </div>
  </section>

  <section id="training" class="Learn-inPersonTraining">
    <div class="Container">
      <div class="Learn-learningResourcesHeader">
        <h3>Đào tạo</h3>
        <p>Tham quan có hướng dẫn các chương trình Go</p>
      </div>
      <div class="LearnGo-gridContainer">
        <ul class="Learn-inPersonList">
          {{- range first 4 (data "training.yaml")}}
          <li class="Learn-inPerson">
            <p class="Learn-inPersonTitle">
              <a href="{{.url}}">{{.title}} </a>
            </p>
            <p class="Learn-inPersonBlurb">{{.blurb}}</p>
          </li>
          {{- end}}
        </ul>
      </div>
    </div>
  </section>

  <section id="featured-books" class="Learn-books">
    <div class="Container">
      <div class="Learn-learningResourcesHeader">
        <h3>Sách</h3>
        <p></p>
      </div>
      <div class="LearnGo-gridContainer">
        <ul class="Learn-cardList Learn-bookList">
          {{- range first 5 (data "books.yaml")}}
            <li class="Learn-card Learn-book">
              {{template "learn-book" .}}
            </li>
          {{- end}}
        </ul>
      </div>
    </div>
  </section>
  </div>
</div>

<script async src="/js/jumplinks.js"></script>

{{define "learn-card"}}
<div class="Card">
  <div class="Card-inner">
    {{- if .thumbnailDark}}
    <div
      class="Card-thumbnail DarkMode-img"
      style="background-image: url('{{.thumbnailDark}}')"
    ></div>
    {{- else if .thumbnail}}
    <div
      class="Card-thumbnail DarkMode-img"
      style="background-image: url('{{.thumbnail}}')"
    ></div>
    {{- end}}
    {{- if .thumbnail}}
    <div
      class="Card-thumbnail LightMode-img"
      style="background-image: url('{{.thumbnail}}')"
    ></div>
    {{- end}}
    <div class="Card-content">
      <div class="Card-contentTitle">{{.title}}</div>
      <p class="Card-contentBody Card-lineClamp">{{raw .content}}</p>
      <div class="Card-contentCta">
        <a href="{{.url}}" target="_blank">
          <span>{{.cta}}</span>
        </a>
      </div>
    </div>
  </div>
</div>
{{- end}}

{{define "learn-self-paced-card"}}
<div class="Card">
  <a href="{{.url}}" target="_blank" rel="noopener">
    <div class="Card-inner">
      {{- if .thumbnail}}
      <div
        class="Card-thumbnail"
        style="background-image: url('{{.thumbnail}}')"
      ></div>
      {{- end}}
      <div class="Card-content">
        <div class="Card-contentTitle">{{.title}}</div>
        <div class="Card-selfPacedFooter">
          <div class="Card-selfPacedCredits">
            <span>{{ .length }}</span> •
            <span>{{.credits}} Credits</span>
          </div>
          <div class="Card-selfPacedRating">
            <div class="Card-starRating" style="width: {{ .rating }}rem"></div>
          </div>
        </div>
      </div>
    </div>
  </a>
</div>
{{- end}}

{{define "learn-book"}}
<div class="Book">
  <a href="{{.url}}" target="_blank" rel="noopener">
    <div class="Book-inner">
      {{- if .thumbnail}}
      <div class="Book-thumbnail">
        <img alt="{{.title}} thumbnail." src="{{.thumbnail}}" />
      </div>
      {{- end}}
      <div class="Book-content">
        <p class="Book-eyebrow">{{.eyebrow}}</p>
        <p class="Book-title">{{.title}}</p>
        <p class="Book-description">{{.description}}</p>
        <div class="Book-cta">
          <span>xem sách</span>
        </div>
      </div>
    </div>
  </a>
</div>
{{- end}}
