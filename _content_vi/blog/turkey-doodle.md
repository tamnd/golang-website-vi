---
title: "Từ không đến Go: ra mắt trên trang chủ Google trong 24 giờ"
date: 2011-12-13
by:
- Reinaldo Aguiar
tags:
- appengine
- google
- guest
summary: Cách Go giúp ra mắt Google Doodle cho Lễ Tạ Ơn 2011.
template: true
---

## Giới thiệu

_Bài viết này được viết bởi Reinaldo Aguiar, kỹ sư phần mềm tại nhóm Search của Google. Ông chia sẻ trải nghiệm phát triển chương trình Go đầu tiên của mình và đưa nó lên phục vụ hàng triệu người dùng, tất cả chỉ trong một ngày!_

Gần đây tôi được cơ hội cộng tác vào một "dự án 20%" nhỏ nhưng rất nổi bật:
[Google Doodle Lễ Tạ Ơn 2011](http://www.google.com/logos/2011/thanksgiving.html).
Doodle này có hình ảnh một con gà tây được tạo ra bằng cách kết hợp ngẫu nhiên các kiểu đầu,
cánh, lông và chân khác nhau.
Người dùng có thể tùy chỉnh nó bằng cách nhấp vào các bộ phận khác nhau của gà tây.
Tính tương tác này được triển khai trên trình duyệt bằng sự kết hợp của JavaScript,
CSS và tất nhiên là HTML, tạo ra các con gà tây ngay lập tức.

{{image "turkey-doodle/image00.png"}}

Khi người dùng tạo ra một con gà tây được cá nhân hóa, nó có thể được chia sẻ với bạn bè
và gia đình bằng cách đăng lên Google+.
Nhấp vào nút "Share" (không được chụp ở đây) sẽ tạo trong luồng Google+ của người dùng
một bài đăng chứa ảnh chụp của con gà tây.
Ảnh chụp là một hình ảnh duy nhất khớp với con gà tây mà người dùng đã tạo.

Với 13 lựa chọn cho mỗi trong số 8 bộ phận của gà tây (đầu,
cặp chân, lông riêng biệt, v.v.) có hơn 800 triệu
hình ảnh ảnh chụp có thể được tạo ra.
Tính toán trước tất cả chúng rõ ràng là không khả thi.
Thay vào đó, chúng tôi phải tạo ra các ảnh chụp ngay lập tức.
Kết hợp vấn đề đó với yêu cầu về khả năng mở rộng tức thì và tính sẵn sàng cao,
lựa chọn nền tảng là hiển nhiên: Google App Engine!

Điều tiếp theo chúng tôi cần quyết định là sử dụng runtime App Engine nào.
Các tác vụ xử lý ảnh bị giới hạn bởi CPU, vì vậy hiệu năng là yếu tố quyết định trong trường hợp này.

Để đưa ra quyết định sáng suốt, chúng tôi đã chạy một bài kiểm tra.
Chúng tôi nhanh chóng chuẩn bị một vài ứng dụng demo tương đương cho
[Python 2.7 runtime](http://code.google.com/appengine/docs/python/python27/newin27.html) mới
(cung cấp [PIL](http://www.pythonware.com/products/pil/),
một thư viện xử lý ảnh dựa trên C) và Go runtime.
Mỗi ứng dụng tạo ra một hình ảnh được ghép từ nhiều hình ảnh nhỏ,
mã hóa hình ảnh đó dưới dạng JPEG và gửi dữ liệu JPEG làm phản hồi HTTP.
Ứng dụng Python 2.7 phục vụ các yêu cầu với độ trễ trung vị là 65 mili giây,
trong khi ứng dụng Go chạy với độ trễ trung vị chỉ 32 mili giây.

Vì vậy bài toán này có vẻ là cơ hội hoàn hảo để thử nghiệm Go runtime.

Tôi không có kinh nghiệm gì về Go trước đó và thời gian rất eo hẹp:
hai ngày để sẵn sàng cho môi trường production.
Điều này thật đáng sợ, nhưng tôi coi đó là cơ hội để kiểm tra Go từ một góc độ khác,
thường bị bỏ qua:
tốc độ phát triển.
Một người không có kinh nghiệm Go có thể tiếp thu và xây dựng thứ gì đó
có hiệu năng tốt và có khả năng mở rộng nhanh đến mức nào?

## Thiết kế

Cách tiếp cận là mã hóa trạng thái của gà tây vào URL, vẽ và mã hóa ảnh chụp ngay lập tức.

Nền tảng của mọi doodle là hình nền:

{{image "turkey-doodle/image01.jpg"}}

Một URL yêu cầu hợp lệ có thể trông như thế này:
`http://google-turkey.appspot.com/thumb/20332620][http://google-turkey.appspot.com/thumb/20332620`

Chuỗi chữ và số theo sau "/thumb/" chỉ ra (dưới dạng thập lục phân)
lựa chọn nào sẽ được vẽ cho từng thành phần bố cục,
như được minh họa bởi hình ảnh này:

{{image "turkey-doodle/image03.png"}}

Trình xử lý yêu cầu của chương trình phân tích URL để xác định thành phần nào
được chọn cho mỗi bộ phận,
vẽ các hình ảnh phù hợp lên trên hình nền,
và phục vụ kết quả dưới dạng JPEG.

Nếu xảy ra lỗi, một hình ảnh mặc định sẽ được phục vụ.
Không có ích gì khi phục vụ trang lỗi vì người dùng sẽ không bao giờ thấy nó,
trình duyệt hầu như chắc chắn đang tải URL này vào một thẻ hình ảnh.

## Triển khai

Ở phạm vi package, chúng tôi khai báo một số cấu trúc dữ liệu để mô tả các thành phần của gà tây,
vị trí của các hình ảnh tương ứng,
và nơi chúng nên được vẽ trên hình nền.

	var (
	    // dirs maps each layout element to its location on disk.
	    dirs = map[string]string{
	        "h": "img/heads",
	        "b": "img/eyes_beak",
	        "i": "img/index_feathers",
	        "m": "img/middle_feathers",
	        "r": "img/ring_feathers",
	        "p": "img/pinky_feathers",
	        "f": "img/feet",
	        "w": "img/wing",
	    }

	    // urlMap maps each URL character position to
	    // its corresponding layout element.
	    urlMap = [...]string{"b", "h", "i", "m", "r", "p", "f", "w"}

	    // layoutMap maps each layout element to its position
	    // on the background image.
	    layoutMap = map[string]image.Rectangle{
	        "h": {image.Pt(109, 50), image.Pt(166, 152)},
	        "i": {image.Pt(136, 21), image.Pt(180, 131)},
	        "m": {image.Pt(159, 7), image.Pt(201, 126)},
	        "r": {image.Pt(188, 20), image.Pt(230, 125)},
	        "p": {image.Pt(216, 48), image.Pt(258, 134)},
	        "f": {image.Pt(155, 176), image.Pt(243, 213)},
	        "w": {image.Pt(169, 118), image.Pt(250, 197)},
	        "b": {image.Pt(105, 104), image.Pt(145, 148)},
	    }
	)

Hình học của các điểm trên được tính toán bằng cách đo vị trí và kích thước thực tế
của từng thành phần bố cục trong hình ảnh.

Tải hình ảnh từ đĩa trong mỗi yêu cầu sẽ là sự lặp lại lãng phí,
vì vậy chúng tôi tải tất cả 106 hình ảnh (13 \* 8 phần tử + 1 nền + 1 mặc định) vào
các biến toàn cục khi nhận yêu cầu đầu tiên.

	var (
	    // elements maps each layout element to its images.
	    elements = make(map[string][]*image.RGBA)

	    // backgroundImage contains the background image data.
	    backgroundImage *image.RGBA

	    // defaultImage is the image that is served if an error occurs.
	    defaultImage *image.RGBA

	    // loadOnce is used to call the load function only on the first request.
	    loadOnce sync.Once
	)

	// load reads the various PNG images from disk and stores them in their
	// corresponding global variables.
	func load() {
	    defaultImage = loadPNG(defaultImageFile)
	    backgroundImage = loadPNG(backgroundImageFile)
	    for dirKey, dir := range dirs {
	        paths, err := filepath.Glob(dir + "/*.png")
	        if err != nil {
	            panic(err)
	        }
	        for _, p := range paths {
	            elements[dirKey] = append(elements[dirKey], loadPNG(p))
	        }
	    }
	}

Các yêu cầu được xử lý theo một trình tự đơn giản:

  - Phân tích URL yêu cầu, giải mã giá trị thập phân của từng ký tự trong đường dẫn.

  - Tạo một bản sao của hình nền làm nền cho hình ảnh cuối cùng.

  - Vẽ từng thành phần hình ảnh lên hình nền sử dụng layoutMap để xác định nơi chúng nên được vẽ.

  - Mã hóa hình ảnh dưới dạng JPEG.

  - Trả hình ảnh cho người dùng bằng cách ghi JPEG trực tiếp vào HTTP response writer.

Nếu có lỗi xảy ra, chúng tôi phục vụ defaultImage cho người dùng và ghi
lỗi vào bảng điều khiển App Engine để phân tích sau.

Đây là mã cho trình xử lý yêu cầu với chú thích giải thích:

{{raw `
<pre>
func handler(w http.ResponseWriter, r *http.Request) {
    // <a href="/blog/defer-panic-and-recover.html">Defer</a> a function to recover from any panics.
    // When recovering from a panic, log the error condition to
    // the App Engine dashboard and send the default image to the user.
    defer func() {
        if err := recover(); err != nil {
            c := appengine.NewContext(r)
            c.Errorf("%s", err)
            c.Errorf("%s", "Traceback: %s", r.RawURL)
            if defaultImage != nil {
                w.Header().Set("Content-type", "image/jpeg")
                jpeg.Encode(w, defaultImage, &imageQuality)
            }
        }
    }()

    // Load images from disk on the first request.
    loadOnce.Do(load)

    // Make a copy of the background to draw into.
    bgRect := backgroundImage.Bounds()
    m := image.NewRGBA(bgRect.Dx(), bgRect.Dy())
    draw.Draw(m, m.Bounds(), backgroundImage, image.ZP, draw.Over)

    // Process each character of the request string.
    code := strings.ToLower(r.URL.Path[len(prefix):])
    for i, p := range code {
        // Decode hex character p in place.
        if p &lt; 'a' {
            // it's a digit
            p = p - '0'
        } else {
            // it's a letter
            p = p - 'a' + 10
        }

        t := urlMap[i]    // element type by index
        em := elements[t] // element images by type
        if p >= len(em) {
            panic(fmt.Sprintf("element index out of range %s: "+
                "%d >= %d", t, p, len(em)))
        }

        // Draw the element to m,
        // using the layoutMap to specify its position.
        draw.Draw(m, layoutMap[t], em[p], image.ZP, draw.Over)
    }

    // Encode JPEG image and write it as the response.
    w.Header().Set("Content-type", "image/jpeg")
    w.Header().Set("Cache-control", "public, max-age=259200")
    jpeg.Encode(w, m, &imageQuality)
}
</pre>
`}}

Để ngắn gọn, tôi đã bỏ qua một số hàm trợ giúp trong các danh sách mã này.
Xem [mã nguồn](http://code.google.com/p/go-thanksgiving/source/browse/) để biết đầy đủ.

## Hiệu năng

{{image "turkey-doodle/image02.png"}}

Biểu đồ này, lấy trực tiếp từ bảng điều khiển App Engine, cho thấy độ trễ yêu cầu trung bình trong khi ra mắt.
Như bạn có thể thấy, ngay cả khi tải cao nó cũng không bao giờ vượt quá 60 ms,
với độ trễ trung vị là 32 mili giây.
Đây là kết quả cực kỳ nhanh, xét đến việc trình xử lý yêu cầu của chúng tôi đang thực hiện
thao tác và mã hóa hình ảnh ngay lập tức.

## Kết luận

Tôi thấy cú pháp của Go trực quan, đơn giản và gọn gàng.
Tôi đã làm việc nhiều với các ngôn ngữ thông dịch trong quá khứ,
và mặc dù Go là ngôn ngữ kiểu tĩnh và biên dịch,
việc viết ứng dụng này lại cảm giác giống như làm việc với một ngôn ngữ
động được thông dịch hơn.

Máy chủ phát triển đi kèm với [SDK](http://code.google.com/appengine/downloads.html#Google_App_Engine_SDK_for_Go)
nhanh chóng biên dịch lại chương trình sau bất kỳ thay đổi nào,
vì vậy tôi có thể lặp đi lặp lại nhanh chóng như khi dùng ngôn ngữ thông dịch.
Nó cũng cực kỳ đơn giản, mất chưa đầy một phút để thiết lập môi trường phát triển của tôi.

Tài liệu tuyệt vời của Go cũng giúp tôi hoàn thiện nhanh chóng.
Các tài liệu được tạo từ mã nguồn,
vì vậy tài liệu của mỗi hàm liên kết trực tiếp đến mã nguồn liên quan.
Điều này không chỉ cho phép nhà phát triển hiểu rất nhanh một hàm cụ thể
làm gì mà còn khuyến khích nhà phát triển đào sâu vào triển khai package,
giúp dễ dàng học phong cách và quy ước lập trình tốt hơn.

Khi viết ứng dụng này, tôi chỉ sử dụng ba tài nguyên:
[Ví dụ Hello World Go của App Engine](http://code.google.com/appengine/docs/go/gettingstarted/helloworld.html),
[tài liệu package Go](/pkg/),
và [một bài đăng blog giới thiệu package Draw](/blog/go-imagedraw-package).
Nhờ vào vòng lặp phát triển nhanh được tạo ra bởi máy chủ phát triển và
chính ngôn ngữ,
tôi đã có thể tiếp thu ngôn ngữ và xây dựng một bộ tạo doodle siêu nhanh,
sẵn sàng cho môi trường production, trong chưa đầy 24 giờ.

Tải mã nguồn đầy đủ của ứng dụng (bao gồm hình ảnh) tại [dự án Google Code](http://code.google.com/p/go-thanksgiving/source/browse/).

Xin gửi lời cảm ơn đặc biệt đến Guillermo Real và Ryan Germick, những người đã thiết kế doodle này.
