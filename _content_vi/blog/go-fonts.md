---
title: Font chữ Go
date: 2016-11-16
by:
- Nigel Tao
- Chuck Bigelow
- Rob Pike
summary: Giới thiệu bộ font chữ Go, được thiết kế bởi Bigelow & Holmes.
template: true
---

<link rel="stylesheet" href="/css/fonts.css">

## Thông báo

Bộ công cụ giao diện người dùng thử nghiệm đang được xây dựng tại
[`golang.org/x/exp/shiny`](https://godoc.org/golang.org/x/exp/shiny)
bao gồm một số thành phần văn bản, nhưng gặp một vấn đề trong quá trình kiểm thử:
Nên dùng font chữ nào?
Câu hỏi này dẫn đến thông báo hôm nay:
phát hành bộ font chữ TrueType [WGL4](https://en.wikipedia.org/wiki/Windows_Glyph_List_4) chất lượng cao,
được xưởng thiết kế chữ [Bigelow & Holmes](http://bigelowandholmes.typepad.com/) tạo ra đặc biệt cho dự án Go.

Bộ font, tự nhiên được đặt tên là Go, bao gồm kiểu chữ proportional và fixed-width với các biến thể normal,
bold và italic.
Các font này đã được kiểm thử cho các ứng dụng kỹ thuật, đặc biệt là lập trình.
Code Go trông đặc biệt đẹp khi hiển thị bằng font Go, đúng với tên gọi, với những thứ như
ký tự dấu câu dễ phân biệt và các toán tử được căn chỉnh và sắp xếp nhất quán:

{{image "go-fonts/go-font-code.png" 519}}

Tính năng đáng chú ý nhất của font Go có lẽ là giấy phép sử dụng:
Chúng được cấp phép theo cùng giấy phép mã nguồn mở như phần mềm còn lại của dự án Go,
một sự sắp xếp hiếm có đối với một bộ font chất lượng cao.

Dưới đây là mẫu của các font proportional...

{{image "go-fonts/go-regular.png" 600}}

và font monospaced:

{{image "go-fonts/go-mono.png" 600}}

## Cách sử dụng

Nếu bạn chỉ muốn các file TTF, hãy chạy

	git clone https://go.googlesource.com/image

và sao chép chúng từ thư mục `image/font/gofont/ttfs` sau đó.
Nếu bạn muốn dùng Go (font chữ) với Go (phần mềm), mỗi font được cung cấp qua một gói riêng.
Để dùng font Go Regular trong chương trình, import `golang.org/x/image/font/gofont/goregular` và viết:

	font, err := truetype.Parse(goregular.TTF)

Gói [`github.com/golang/freetype/truetype`](https://godoc.org/github.com/golang/freetype/truetype)
cung cấp hàm [`truetype.Parse`](https://godoc.org/github.com/golang/freetype/truetype#Parse) hiện tại.
Ngoài ra, đang có công việc thêm gói TrueType dưới `golang.org/x`
cũng được cấp phép theo cùng giấy phép mã nguồn mở như phần mềm còn lại của dự án Go.

Chúng tôi để bạn tự khám phá một số thuộc tính bất thường khác của font,
nhưng để có cái nhìn tổng quan về thiết kế font, chúng tôi đã nhờ Chuck Bigelow cung cấp một số thông tin nền.
Phần còn lại của bài đăng blog này là phản hồi của ông.

## Ghi chú về các font, bởi Chuck Bigelow

Font Go được chia thành hai bộ, Go proportional vốn là
sans-serif, và Go Mono vốn là slab-serif.

## Font proportional Go

### Sans-serif

Font proportional Go là sans-serif, giống một số font phổ biến
dùng để hiển thị trên màn hình. Có một số bằng chứng cho thấy một số kiểu
sans-serif ở kích thước nhỏ và độ phân giải thấp trên màn hình có thể dễ đọc hơn
một chút so với các kiểu có serif tương đương, trong khi ở kích thước lớn,
không có sự khác biệt đáng kể về khả năng đọc giữa sans và
có serif, ít nhất là trong cặp được kiểm thử. [1] (Các số trong ngoặc
tham chiếu đến danh sách tài liệu tham khảo ở cuối bài.)

### Phong cách

Font sans-serif Go theo phong cách "humanist" thay vì "grotesque".
Đây là sự phân biệt mang tính lịch sử, không phải đánh giá thẩm mỹ.
Các font sans-serif được dùng rộng rãi như Helvetica và Arial được gọi là
grotesque vì một kiểu chữ sans-serif đầu thế kỷ 19
có tên "Grotesque", và cái tên này trở thành tên chung.

Hình dạng của các font grotesque hiện đại như Helvetica được tạo tác tỉ mỉ,
với các hình dạng mượt mà và đồng nhất.

Font sans-serif humanist bắt nguồn từ chữ viết tay Humanist
và các font đầu thời Phục hưng Ý, vẫn còn lưu lại dấu vết tinh tế
của thư pháp viết bằng bút lông. Có một số bằng chứng cho thấy
font humanist dễ đọc hơn font grotesque. [2]

### Chữ nghiêng (Italics)

Chữ nghiêng proportional Go có cùng số liệu chiều rộng như
các font roman. Chữ nghiêng Go là phiên bản xiên của roman, với
một ngoại lệ đáng chú ý: chữ 'a' thường ở chữ nghiêng được thiết kế lại thành
dạng một câu chuyện cursive để hài hòa với các hình dạng bowl của
tập b d g p q, trong đó các dạng đứng cũng thích nghi tốt với
độ nghiêng. Việc thêm 'a' cursive làm cho chữ nghiêng trông sinh động hơn
so với roman chỉ đơn giản là nghiêng. Một số nhà thiết kế chữ cho rằng
chữ nghiêng roman xiên sans-serif được ưa thích hơn là chữ nghiêng "cursive" thực sự,
một phần vì lịch sử và thiết kế. [3]

{{image "go-fonts/abdgpq-proportional.png"}}

### Chiều cao x (x-height)

Chiều cao x của một kiểu chữ là chiều cao của chữ 'x' thường so với
kích thước thân chữ. Chiều cao x của font Go là 53,0% kích thước thân chữ, cao hơn
một chút so với chiều cao x của Helvetica (52,3%) hoặc Arial (51,9%),
nhưng sự khác biệt thường không nhận thấy ở kích thước đọc bình thường.
Các nhà thiết kế chữ cho rằng chiều cao x lớn hơn góp phần vào khả năng đọc tốt hơn
ở kích thước nhỏ và trên màn hình. Một nghiên cứu về "kích thước in"
(đặc biệt là chiều cao x) và việc đọc lưu ý rằng các kiểu chữ dùng trên
màn hình và kích thước nhỏ thường có chiều cao x lớn. [4]

### Tiêu chuẩn DIN về khả năng đọc

Tiêu chuẩn DIN 1450 của Đức về khả năng đọc gần đây đề xuất
một số tính năng giúp font dễ đọc, bao gồm việc phân biệt
các hình dạng chữ để giảm nhầm lẫn. Font Go tuân thủ
tiêu chuẩn 1450 bằng cách phân biệt cẩn thận số không với chữ O hoa;
số 1 với chữ I hoa (eye) và chữ l thường (ell); số 5 với
chữ S hoa; và số 8 với chữ B hoa. Hình dạng của các bowl của
b d p q tuân theo các bất đối xứng tự nhiên của chữ viết tay Phục hưng dễ đọc,
giúp phân biệt để giảm nhầm lẫn. [5]

### Độ đậm (Weights)

Font proportional Go có ba độ đậm: Normal, Medium
và Bold. Độ đậm Normal đủ mạnh để duy trì sự rõ nét trên màn hình
phát sáng, vốn thường làm mờ các nét và độ dày của chữ. Độ đậm Medium có độ dày
thân chữ gấp 1,25 lần Normal, để trông vững chắc hơn trên màn hình sáng hoặc cho
người dùng ưa font dày. Độ đậm Bold có độ dày thân chữ
gấp 1,5 lần Normal, đủ đậm để phân biệt với độ đậm normal.
Các font Go này có trọng số CSS là 400, 500
và 600. Mặc dù CSS quy định "Bold" là trọng số 700 và 600
là Semibold hoặc Demibold, trọng số số của Go khớp với
tiến trình thực tế của tỷ lệ độ dày thân chữ:
Normal:Medium = 400:500; Normal:Bold = 400:600. Tên độ đậm Bold
khớp với việc dùng "Bold" như trọng số bold tương ứng thông thường
của font normal. Thảo luận thêm về mối quan hệ giữa
độ dày thân chữ, tên độ đậm và đánh số CSS có trong [6].

### Bộ ký tự WGL4

Bộ ký tự WGL4, ban đầu được phát triển bởi Microsoft, thường được
dùng như tiêu chuẩn bộ ký tự không chính thức. WGL4 bao gồm ký tự Latin
Tây và Đông Âu cùng tiếng Hy Lạp hiện đại và
Cyrillic, với các ký hiệu, dấu hiệu và ký tự đồ họa bổ sung,
tổng cộng hơn 650 ký tự. Font WGL4 Go có thể
dùng để soạn thảo nhiều ngôn ngữ. [7]

### Tương thích số liệu với Arial và Helvetica

Font sans-serif Go gần như tương thích số liệu với
ký tự Helvetica hoặc Arial tiêu chuẩn. Văn bản đặt trong Go chiếm
gần cùng không gian như văn bản trong Helvetica hoặc Arial (ở cùng
kích thước), nhưng Go có diện mạo và kết cấu khác vì phong cách
humanist. Một số chữ Go có tính năng DIN rộng hơn các chữ tương ứng
trong Helvetica hoặc Arial, vì vậy một số văn bản đặt trong Go có thể chiếm
nhiều không gian hơn đôi chút.

## Font Mono Go

### Monospaced

Font Mono Go là monospaced, mỗi chữ có cùng chiều rộng với
các chữ khác. Font monospaced đã được dùng trong lập trình
từ buổi đầu của điện toán và vẫn được dùng rộng rãi vì sự đều đặn
kiểu máy chữ trong khoảng cách của chúng giúp văn bản căn thẳng theo cột và
hàng, một phong cách cũng thấy trong các chữ khắc Hy Lạp từ thế kỷ 5 TCN.
(Người Hy Lạp cổ đại không có máy chữ hay bàn phím máy tính,
nhưng họ có các nhà toán học vĩ đại và cảm quan tuyệt vời về đối xứng
và hoa văn đã định hình bảng chữ cái của họ.)

### Slab-serif

Font Mono Go có serif dạng slab, mang lại vẻ ngoài vững chắc.

### Phong cách

Hình dạng chữ cái cơ bản của Go Mono, giống như font sans-serif Go,
bắt nguồn từ chữ viết tay humanist, nhưng kiểu monospaced và slab serif
có xu hướng che khuất các kết nối lịch sử và phong cách.

### Chữ nghiêng (Italics)

Go Mono Italics là phiên bản xiên của roman, với ngoại lệ
là chữ 'a' thường ở chữ nghiêng được thiết kế lại thành dạng cursive một câu chuyện
để hài hòa với các hình dạng bowl của b d g p q. Chữ 'a' cursive làm
cho chữ nghiêng trông sinh động hơn so với roman đơn giản chỉ nghiêng. Như với nhiều
font sans-serif, người ta cho rằng font slab-serif roman xiên có thể
dễ đọc hơn chữ nghiêng "cursive" thực sự.

{{image "go-fonts/abdgpq-mono.png"}}

### Chiều cao x (x-height)

Font Mono Go có cùng chiều cao x với font sans-serif Go, 53% của
kích thước thân chữ. Go Mono trông lớn hơn Courier gần 18%, vì Courier có
chiều cao x bằng 45% kích thước thân chữ. Tuy nhiên Go Mono có cùng chiều rộng
với Courier, vì vậy vẻ ngoài lớn hơn đạt được mà không mất đi sự tiết kiệm về
số ký tự trên mỗi dòng.

### Tiêu chuẩn DIN về khả năng đọc

Font Mono Go tuân thủ tiêu chuẩn DIN 1450 bằng cách phân biệt
số không với chữ O hoa; số 1 với chữ I hoa (eye) và chữ l thường (ell);
số 5 với chữ S hoa; và số 8 với chữ B hoa. Hình dạng của các bowl của
b d p q tuân theo các bất đối xứng tự nhiên của chữ viết tay Phục hưng dễ đọc,
giúp phân biệt và giảm nhầm lẫn.

### Độ đậm (Weights)

Font Mono Go có hai độ đậm: Normal và Bold. Thân chữ độ đậm normal
giống như trong Go Normal và do đó duy trì sự rõ nét trên màn hình phát sáng
ngược, vốn có xu hướng làm mờ các nét chữ và độ dày thân. Độ dày thân chữ bold
dày hơn 1,5 lần so với độ đậm normal, vì vậy Bold Mono có cùng độ dày thân chữ
như Bold Go proportional. Vì chiều rộng chữ của monospaced bold giống với chiều rộng
của monospaced normal, Mono bold trông đậm hơn một chút so với
Go Bold proportional, vì nhiều pixel đen hơn được đặt vào cùng diện tích.

### Tương thích số liệu với font monospaced phổ biến

Go Mono tương thích số liệu với Courier và các font monospaced khác
khớp với chiều rộng chữ máy chữ "Pica" là 10 ký tự trên một inch thẳng
tại 12 point. Ở 10 point, font Go Mono đặt 12 ký tự trên mỗi inch. Các font
TrueType có thể tỉ lệ, tất nhiên, vì vậy Go Mono có thể đặt ở bất kỳ kích thước nào.

### Bộ ký tự WGL4

Font Mono Go cung cấp bộ ký tự WGL4 thường được dùng như
tiêu chuẩn bộ ký tự không chính thức. WGL4 bao gồm ký tự Latin Tây và Đông
Âu cùng tiếng Hy Lạp hiện đại và Cyrillic, với
các ký hiệu, dấu hiệu và ký tự đồ họa bổ sung. Hơn 650 ký tự
của bộ WGL4 Go có thể dùng cho nhiều ngôn ngữ.

## Tài liệu tham khảo

[1] Morris, R. A., Aquilante, K., Yager, D., & Bigelow, C.
(2002, May). P-13: Serifs Slow RSVP Reading at Very Small Sizes,
but Don't Matter at Larger Sizes.
In SID Symposium Digest of Technical Papers (Vol.
33, No. 1, pp. 244-247). Blackwell Publishing Ltd.

[2] Bryan Reimer et al. (2014) "Assessing the impact of typeface design
in a text-rich automotive user interface",
Ergonomics, 57:11, 1643-1658.
http://www.tandfonline.com/doi/abs/10.1080/00140139.2014.940000

[3] Adrian Frutiger - Typefaces: The Complete Works.
H. Osterer and P. Stamm, editors. Birkhäuser,
Basel, 2009, page 257.

[4] Legge, G. E., & Bigelow, C. A. (2011).
Does print size matter for reading? A review of findings from vision science and typography.
Journal of Vision, 11(5), 8-8. http://jov.arvojournals.org/article.aspx?articleid=2191906

[5] Charles Bigelow. "Oh, oh, zero!" TUGboat, Volume 34 (2013), No. 2.
https://tug.org/TUGboat/tb34-2/tb107bigelow-zero.pdf
https://tug.org/TUGboat/tb34-2/tb107bigelow-wang.pdf

[6] "Lucida Basic Font Weights" Bigelow & Holmes.
http://lucidafonts.com/pages/facts

[7] WGL4 language coverage: Afrikaans, Albanian, Asu, Basque,
Belarusian, Bemba, Bena, Bosnian, Bulgarian, Catalan, Chiga,
Colognian, Cornish, Croatian, Czech, Danish, Embu, English, Esperanto,
Estonian, Faroese, Filipino, Finnish, French, Friulian, Galician,
Ganda, German, Greek, Gusii, Hungarian, Icelandic, Inari Sami,
Indonesian, Irish, Italian, Jola-Fonyi, Kabuverdianu, Kalaallisut,
Kalenjin, Kamba, Kikuyu, Kinyarwanda, Latvian, Lithuanian, Lower
Sorbian, Luo, Luxembourgish, Luyia, Macedonian, Machame, Makhuwa-Meetto,
Makonde, Malagasy, Malay, Maltese, Manx, Meru, Morisyen, North
Ndebele, Northern Sami, Norwegian Bokmål, Norwegian Nynorsk, Nyankole,
Oromo, Polish, Portuguese, Romanian, Romansh, Rombo, Rundi, Russian,
Rwa, Samburu, Sango, Sangu, Scottish Gaelic, Sena, Serbian, Shambala,
Shona, Slovak, Slovenian, Soga, Somali, Spanish, Swahili, Swedish,
Swiss German, Taita, Teso, Turkish, Turkmen, Upper Sorbian, Vunjo,
Walser, Welsh, Zulu

## Jabberwocky trong Go Regular

Từ [en.wikipedia.org/wiki/Jabberwocky](https://en.wikipedia.org/wiki/Jabberwocky):

{{image "go-fonts/go-font-jabberwocky.png" 500}}

Không có phiên bản tiếng Hy Lạp nào được liệt kê. Thay vào đó, một pangram từ [clagnut.com/blog/2380/#Greek](http://clagnut.com/blog/2380/#Greek):

{{image "go-fonts/go-font-greek.png" 530}}
