---
title: Hai bài nói chuyện Go gần đây
date: 2013-01-02
by:
- Andrew Gerrand
tags:
- talk
- video
- ethos
summary: "Hai bài nói chuyện về Go: \"Go: A Simple Programming Environment\" và \"Go: Code That Grows With Grace\"."
template: true
---

## Giới thiệu

Cuối năm ngoái tôi đã viết một vài bài nói chuyện về Go và trình bày chúng tại [Strange Loop](http://thestrangeloop.com/),
[Øredev](http://oredev.com), và nhiều địa điểm khác.
Các bài nói được thiết kế để mang lại cái nhìn thực tế về lập trình Go,
mỗi bài mô tả việc xây dựng một chương trình thực tế và thể hiện sức mạnh cũng như chiều sâu của ngôn ngữ Go,
các thư viện và công cụ của nó.

Dưới đây là những video mà theo tôi là bản ghi tốt nhất của các bài nói đó.

## Go: một môi trường lập trình đơn giản

Go là một ngôn ngữ đa năng thu hẹp khoảng cách giữa các ngôn ngữ kiểu tĩnh hiệu quả
và các ngôn ngữ động có năng suất cao.
Nhưng không chỉ là ngôn ngữ làm cho Go đặc biệt, Go còn có thư viện chuẩn rộng lớn và nhất quán,
cùng các công cụ mạnh mẽ nhưng đơn giản.

Bài nói này giới thiệu về Go, sau đó là một chuyến khảo sát một số chương trình thực tế
thể hiện sức mạnh, phạm vi và sự đơn giản của môi trường lập trình Go.

{{video "https://player.vimeo.com/video/53221558?badge=0" 500 281}}

Xem [bộ slide](/talks/2012/simple.slide) (dùng phím mũi tên trái và phải để điều hướng).

## Go: mã nguồn phát triển một cách duyên dáng

Một trong những mục tiêu thiết kế chính của Go là khả năng thích ứng của mã nguồn;
tức là có thể dễ dàng lấy một thiết kế đơn giản và mở rộng nó theo cách rõ ràng và tự nhiên.
Trong bài nói này, tôi mô tả một máy chủ "chat roulette" đơn giản để ghép cặp
các kết nối TCP đến,
rồi sử dụng các cơ chế đồng thời của Go, interface và thư viện chuẩn
để mở rộng nó với giao diện web và các tính năng khác.
Mặc dù chức năng của chương trình thay đổi đáng kể,
tính linh hoạt của Go vẫn giữ nguyên thiết kế ban đầu khi nó phát triển.

{{video "https://player.vimeo.com/video/53221560?badge=0" 500 281}}

Xem [bộ slide](/talks/2012/chat.slide) (dùng phím mũi tên trái và phải để điều hướng).
