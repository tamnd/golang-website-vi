---
title: Gói image/draw trong Go
date: 2011-09-29
by:
- Nigel Tao
tags:
- draw
- image
- libraries
- technical
summary: Giới thiệu về ghép ảnh trong Go bằng gói image/draw.
template: true
---

## Giới thiệu

[Gói image/draw](/pkg/image/draw/) chỉ định nghĩa một thao tác duy nhất:
vẽ một ảnh nguồn lên một ảnh đích, thông qua một ảnh mặt nạ tùy chọn.
Thao tác đơn giản này lại có khả năng ứng dụng đáng ngạc nhiên và có thể thực hiện nhiều tác vụ xử lý ảnh phổ biến một cách thanh lịch và hiệu quả.

Phép ghép ảnh được thực hiện theo từng điểm ảnh theo phong cách của thư viện đồ họa Plan 9 và phần mở rộng X Render.
Mô hình dựa trên bài báo kinh điển "Compositing Digital Images" của Porter và Duff, với thêm tham số mặt nạ:
`dst = (src IN mask) OP dst`.
Với mặt nạ hoàn toàn mờ đục, công thức này rút gọn thành công thức Porter-Duff gốc: `dst = src OP dst`.
Trong Go, ảnh mặt nạ nil tương đương với một ảnh mặt nạ có kích thước vô hạn và hoàn toàn mờ đục.

Bài báo Porter-Duff trình bày [12 toán tử ghép ảnh khác nhau](http://www.w3.org/TR/SVGCompositing/examples/compop-porterduff-examples.png),
nhưng với mặt nạ rõ ràng, trong thực tế chỉ cần 2 trong số đó:
source-over-destination và source.
Trong Go, hai toán tử này được biểu diễn bởi các hằng số `Over` và `Src`.
Toán tử `Over` thực hiện việc chồng lớp tự nhiên của ảnh nguồn lên ảnh đích:
sự thay đổi ở ảnh đích nhỏ hơn khi ảnh nguồn (sau khi áp mặt nạ) trong suốt hơn (tức là có alpha thấp hơn).
Toán tử `Src` chỉ đơn giản sao chép ảnh nguồn (sau khi áp mặt nạ) mà không quan tâm đến nội dung gốc của ảnh đích.
Với ảnh nguồn và mặt nạ hoàn toàn mờ đục, hai toán tử cho kết quả giống nhau, nhưng toán tử `Src` thường nhanh hơn.

## Căn chỉnh hình học

Phép ghép ảnh yêu cầu liên kết các điểm ảnh đích với các điểm ảnh nguồn và mặt nạ.
Rõ ràng là cần có ảnh đích, ảnh nguồn, ảnh mặt nạ và toán tử ghép, nhưng cũng cần chỉ định vùng hình chữ nhật nào của mỗi ảnh sẽ được sử dụng.
Không phải mọi thao tác vẽ đều ghi lên toàn bộ ảnh đích:
khi cập nhật một ảnh đang hoạt động, sẽ hiệu quả hơn nếu chỉ vẽ những phần đã thay đổi.
Không phải mọi thao tác vẽ đều đọc từ toàn bộ ảnh nguồn:
khi dùng sprite kết hợp nhiều ảnh nhỏ thành một ảnh lớn, chỉ cần một phần của ảnh.
Không phải mọi thao tác vẽ đều đọc từ toàn bộ mặt nạ:
một ảnh mặt nạ chứa các glyph phông chữ tương tự như sprite.
Do đó, việc vẽ cũng cần biết ba hình chữ nhật, một cho mỗi ảnh.
Vì mỗi hình chữ nhật có cùng chiều rộng và chiều cao, chỉ cần truyền vào một hình chữ nhật đích `r` và hai điểm `sp` và `mp`:
hình chữ nhật nguồn bằng `r` được dịch chuyển sao cho `r.Min` trong ảnh đích trùng với `sp` trong ảnh nguồn,
và tương tự với `mp`.
Vùng thực tế cũng được cắt theo bounds của mỗi ảnh trong không gian tọa độ tương ứng.

{{image "image-draw/20.png"}}

Hàm [`DrawMask`](/pkg/image/draw/#DrawMask) nhận bảy đối số, nhưng mặt nạ và điểm mặt nạ rõ ràng thường không cần thiết, vì vậy hàm [`Draw`](/pkg/image/draw/#Draw) chỉ nhận năm đối số:

	// Draw gọi DrawMask với mặt nạ nil.
	func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
	func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point,
	 mask image.Image, mp image.Point, op Op)

Ảnh đích phải có thể thay đổi được, vì vậy gói image/draw định nghĩa interface [`draw.Image`](/pkg/image/draw/#Image) có phương thức `Set`.

	type Image interface {
	    image.Image
	    Set(x, y int, c color.Color)
	}

## Tô màu một hình chữ nhật

Để tô một hình chữ nhật với màu đơn sắc, dùng nguồn `image.Uniform`.
Kiểu `ColorImage` tái diễn giải một `Color` như một `Image` có kích thước gần như vô hạn với màu đó.
Với những ai quen với thiết kế của thư viện draw trong Plan 9, không cần "repeat bit" rõ ràng trong các kiểu ảnh dựa trên slice của Go; khái niệm đó được tích hợp trong `Uniform`.

	// image.ZP là điểm gốc -- tọa độ (0, 0).
	draw.Draw(dst, r, &image.Uniform{c}, image.ZP, draw.Src)

Để khởi tạo một ảnh mới với màu xanh lam:

	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

Để đặt lại một ảnh về trong suốt (hoặc đen, nếu mô hình màu của ảnh đích không hỗ trợ độ trong suốt), dùng `image.Transparent`, là một `image.Uniform`:

	draw.Draw(m, m.Bounds(), image.Transparent, image.ZP, draw.Src)

{{image "image-draw/2a.png"}}

## Sao chép một ảnh

Để sao chép từ hình chữ nhật `sr` trong ảnh nguồn sang một hình chữ nhật bắt đầu từ điểm `dp` trong ảnh đích, chuyển đổi hình chữ nhật nguồn sang không gian tọa độ của ảnh đích:

	r := image.Rectangle{dp, dp.Add(sr.Size())}
	draw.Draw(dst, r, src, sr.Min, draw.Src)

Hoặc cách khác:

	r := sr.Sub(sr.Min).Add(dp)
	draw.Draw(dst, r, src, sr.Min, draw.Src)

Để sao chép toàn bộ ảnh nguồn, dùng `sr = src.Bounds()`.

{{image "image-draw/2b.png"}}

## Cuộn một ảnh

Cuộn một ảnh chỉ là sao chép ảnh lên chính nó, với hình chữ nhật đích và nguồn khác nhau.
Các ảnh đích và nguồn chồng lên nhau hoàn toàn hợp lệ, tương tự như hàm copy tích hợp sẵn của Go có thể xử lý các slice đích và nguồn chồng nhau.
Để cuộn ảnh m xuống 20 điểm ảnh:

	b := m.Bounds()
	p := image.Pt(0, 20)
	// Lưu ý rằng dù đối số thứ hai là b,
	// vùng thực tế nhỏ hơn do cắt xén.
	draw.Draw(m, b, m, b.Min.Add(p), draw.Src)
	dirtyRect := b.Intersect(image.Rect(b.Min.X, b.Max.Y-20, b.Max.X, b.Max.Y))

{{image "image-draw/2c.png"}}

## Chuyển đổi ảnh sang RGBA

Kết quả giải mã một định dạng ảnh có thể không phải là `image.RGBA`:
giải mã GIF cho ra `image.Paletted`,
giải mã JPEG cho ra `ycbcr.YCbCr`,
và kết quả giải mã PNG phụ thuộc vào dữ liệu ảnh.
Để chuyển đổi bất kỳ ảnh nào sang `image.RGBA`:

	b := src.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), src, b.Min, draw.Src)

{{image "image-draw/2d.png"}}

## Vẽ qua mặt nạ

Để vẽ một ảnh qua mặt nạ hình tròn với tâm `p` và bán kính `r`:

{{raw `
	type circle struct {
	    p image.Point
	    r int
	}

	func (c *circle) ColorModel() color.Model {
	    return color.AlphaModel
	}

	func (c *circle) Bounds() image.Rectangle {
	    return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
	}

	func (c *circle) At(x, y int) color.Color {
	    xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	    if xx*xx+yy*yy < rr*rr {
	        return color.Alpha{255}
	    }
	    return color.Alpha{0}
	}

	    draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{p, r}, image.ZP, draw.Over)
`}}

{{image "image-draw/2e.png"}}

## Vẽ glyph phông chữ

Để vẽ một glyph phông chữ màu xanh lam bắt đầu từ điểm `p`, vẽ với nguồn `image.ColorImage` và mặt nạ `image.Alpha`.
Để đơn giản, chúng tôi không thực hiện căn chỉnh hay render dưới mức điểm ảnh, cũng không bù trừ cho chiều cao của phông chữ so với đường cơ sở.

	src := &image.Uniform{color.RGBA{0, 0, 255, 255}}
	mask := theGlyphImageForAFont()
	mr := theBoundsFor(glyphIndex)
	draw.DrawMask(dst, mr.Sub(mr.Min).Add(p), src, image.ZP, mask, mr.Min, draw.Over)

{{image "image-draw/2f.png"}}

## Hiệu năng

Triển khai gói image/draw minh họa cách cung cấp một hàm xử lý ảnh vừa đa dụng vừa hiệu quả cho các trường hợp phổ biến.
Hàm `DrawMask` nhận các đối số kiểu interface, nhưng ngay lập tức thực hiện type assertion để xác định liệu các đối số có thuộc các kiểu struct cụ thể hay không, tương ứng với các thao tác phổ biến như vẽ một ảnh `image.RGBA` lên một ảnh khác, hoặc vẽ mặt nạ `image.Alpha` (chẳng hạn như glyph phông chữ) lên ảnh `image.RGBA`.
Nếu type assertion thành công, thông tin kiểu đó được dùng để chạy một triển khai chuyên biệt của thuật toán tổng quát.
Nếu các assertion thất bại, đường dẫn dự phòng sử dụng các phương thức `At` và `Set` tổng quát.
Các đường dẫn nhanh hoàn toàn chỉ là tối ưu hóa hiệu năng; ảnh đích kết quả vẫn giống nhau trong mọi trường hợp.
Trong thực tế, chỉ cần một số ít trường hợp đặc biệt để hỗ trợ các ứng dụng điển hình.
