---
title: Gói image trong Go
date: 2011-09-21
by:
- Nigel Tao
tags:
- image
- libraries
- technical
summary: Giới thiệu về xử lý ảnh 2D với gói image trong Go.
template: true
---

## Giới thiệu

Các gói [image](/pkg/image/) và [image/color](/pkg/image/color/) định nghĩa một số kiểu dữ liệu:
`color.Color` và `color.Model` mô tả màu sắc,
`image.Point` và `image.Rectangle` mô tả hình học 2D cơ bản,
và `image.Image` kết hợp hai khái niệm trên để biểu diễn một lưới hình chữ nhật các màu sắc.
Một [bài viết riêng](/doc/articles/image_draw.html) trình bày về ghép ảnh bằng gói [image/draw](/pkg/image/draw/).

## Màu sắc và Mô hình màu

[Color](/pkg/image/color/#Color) là một interface định nghĩa tập phương thức tối thiểu của bất kỳ kiểu nào có thể coi là một màu sắc: kiểu đó có thể chuyển đổi thành các giá trị đỏ, xanh lá, xanh lam và alpha.
Việc chuyển đổi có thể bị mất mát, chẳng hạn khi chuyển từ không gian màu CMYK hoặc YCbCr.

	type Color interface {
	    // RGBA trả về các giá trị đỏ, xanh lá, xanh lam và alpha đã nhân với alpha.
	    // Mỗi giá trị nằm trong khoảng [0, 0xFFFF], nhưng được biểu diễn
	    // bằng uint32 để khi nhân với hệ số pha trộn lên đến 0xFFFF sẽ không
	    // bị tràn số.
	    RGBA() (r, g, b, a uint32)
	}

Có ba điểm tinh tế quan trọng về các giá trị trả về.
Thứ nhất, các kênh đỏ, xanh lá và xanh lam đã được nhân với alpha:
màu đỏ bão hòa hoàn toàn nhưng trong suốt 25% được biểu diễn bằng RGBA trả về r bằng 75%.
Thứ hai, các kênh có phạm vi hiệu dụng 16 bit:
100% đỏ được biểu diễn bằng RGBA trả về r là 65535, không phải 255, để việc chuyển đổi từ CMYK hoặc YCbCr ít bị mất mát hơn.
Thứ ba, kiểu trả về là `uint32`, dù giá trị tối đa là 65535, để đảm bảo rằng nhân hai giá trị với nhau sẽ không bị tràn số.
Phép nhân đó xảy ra khi pha trộn hai màu theo mặt nạ alpha từ màu thứ ba, theo phong cách đại số kinh điển của [Porter và Duff](https://en.wikipedia.org/wiki/Alpha_compositing):

{{raw `
	dstr, dstg, dstb, dsta := dst.RGBA()
	srcr, srcg, srcb, srca := src.RGBA()
	_, _, _, m := mask.RGBA()
	const M = 1<<16 - 1
	// Giá trị đỏ kết quả là sự pha trộn giữa dstr và srcr, nằm trong [0, M].
	// Phép tính cho xanh lá, xanh lam và alpha tương tự.
	dstr = (dstr*(M-m) + srcr*m) / M
`}}

Dòng cuối của đoạn code đó sẽ phức tạp hơn nếu dùng màu không nhân với alpha, đó là lý do tại sao `Color` sử dụng các giá trị đã nhân với alpha.

Gói image/color cũng định nghĩa nhiều kiểu cụ thể triển khai interface `Color`.
Ví dụ, [`RGBA`](/pkg/image/color/#RGBA) là một struct biểu diễn màu "8 bit mỗi kênh" kinh điển.

	type RGBA struct {
	    R, G, B, A uint8
	}

Lưu ý rằng trường `R` của `RGBA` là màu đỏ đã nhân với alpha 8 bit trong khoảng [0, 255].
`RGBA` thỏa mãn interface `Color` bằng cách nhân giá trị đó với 0x101 để tạo ra màu đỏ đã nhân với alpha 16 bit trong khoảng [0, 65535].
Tương tự, kiểu struct [`NRGBA`](/pkg/image/color/#NRGBA) biểu diễn màu 8 bit chưa nhân với alpha, như được dùng trong định dạng ảnh PNG.
Khi thao tác trực tiếp với các trường của `NRGBA`, các giá trị là chưa nhân với alpha, nhưng khi gọi phương thức `RGBA`, các giá trị trả về đã được nhân với alpha.

Một [`Model`](/pkg/image/color/#Model) đơn giản là thứ có thể chuyển đổi `Color` sang các `Color` khác, có thể bị mất mát.
Ví dụ, `GrayModel` có thể chuyển đổi bất kỳ `Color` nào thành [`Gray`](/pkg/image/color/#Gray) không bão hòa.
`Palette` có thể chuyển đổi bất kỳ `Color` nào thành một màu trong bảng màu giới hạn.

	type Model interface {
	    Convert(c Color) Color
	}

	type Palette []Color

## Điểm và Hình chữ nhật

Một [`Point`](/pkg/image/#Point) là tọa độ (x, y) trên lưới số nguyên, với các trục tăng dần sang phải và xuống dưới.
Nó không phải là điểm ảnh hay ô lưới. Một `Point` không có chiều rộng, chiều cao hay màu sắc cố hữu, nhưng các hình minh họa dưới đây dùng một hình vuông nhỏ có màu.

	type Point struct {
	    X, Y int
	}

{{image "image/image-package-01.png"}}

	p := image.Point{2, 1}

Một [`Rectangle`](/pkg/image/#Rectangle) là hình chữ nhật song song với trục trên lưới số nguyên, được xác định bởi hai điểm `Point` ở góc trên-trái và góc dưới-phải.
Một `Rectangle` cũng không có màu sắc cố hữu, nhưng các hình minh họa dưới đây vẽ đường viền hình chữ nhật bằng đường mảnh có màu và chỉ ra các `Point` `Min` và `Max` của chúng.

	type Rectangle struct {
	    Min, Max Point
	}

Để thuận tiện, `image.Rect(x0, y0, x1, y1)` tương đương với `image.Rectangle{image.Point{x0, y0}, image.Point{x1, y1}}`, nhưng dễ gõ hơn nhiều.

Một `Rectangle` bao gồm góc trên-trái và không bao gồm góc dưới-phải.
Với `Point p` và `Rectangle r`, `p.In(r)` khi và chỉ khi {{raw "`r.Min.X <= p.X && p.X < r.Max.X`"}}, và tương tự với `Y`.
Điều này tương tự như cách một slice `s[i0:i1]` bao gồm đầu thấp nhưng không bao gồm đầu cao.
(Khác với mảng và slice, một `Rectangle` thường có gốc tọa độ khác không.)

{{image "image/image-package-02.png"}}

	r := image.Rect(2, 1, 5, 5)
	// Dx và Dy trả về chiều rộng và chiều cao của hình chữ nhật.
	fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r)) // in ra 3 4 false

Cộng một `Point` vào một `Rectangle` sẽ dịch chuyển `Rectangle` đó.
Các điểm và hình chữ nhật không bị giới hạn trong góc phần tư dưới-phải.

{{image "image/image-package-03.png"}}

	r := image.Rect(2, 1, 5, 5).Add(image.Pt(-4, -2))
	fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r)) // in ra 3 4 true

Giao của hai hình chữ nhật cho ra một hình chữ nhật khác, có thể rỗng.

{{image "image/image-package-04.png"}}

	r := image.Rect(0, 0, 4, 3).Intersect(image.Rect(2, 2, 5, 5))
	// Size trả về chiều rộng và chiều cao của hình chữ nhật dưới dạng Point.
	fmt.Printf("%#v\n", r.Size()) // in ra image.Point{X:2, Y:1}

Các điểm và hình chữ nhật được truyền và trả về theo giá trị.
Một hàm nhận đối số `Rectangle` sẽ hiệu quả không kém gì một hàm nhận hai đối số `Point`, hoặc bốn đối số `int`.

## Ảnh

Một [Image](/pkg/image/#Image) ánh xạ mỗi ô lưới trong một `Rectangle` đến một `Color` từ một `Model`.
"Điểm ảnh tại (x, y)" đề cập đến màu của ô lưới được xác định bởi các điểm (x, y), (x+1, y), (x+1, y+1) và (x, y+1).

	type Image interface {
	    // ColorModel trả về mô hình màu của Image.
	    ColorModel() color.Model
	    // Bounds trả về miền mà At có thể trả về màu khác không.
	    // Bounds không nhất thiết phải chứa điểm (0, 0).
	    Bounds() Rectangle
	    // At trả về màu của điểm ảnh tại (x, y).
	    // At(Bounds().Min.X, Bounds().Min.Y) trả về điểm ảnh trên-trái của lưới.
	    // At(Bounds().Max.X-1, Bounds().Max.Y-1) trả về điểm ảnh dưới-phải.
	    At(x, y int) color.Color
	}

Một lỗi phổ biến là giả định rằng bounds của một `Image` bắt đầu từ (0, 0).
Ví dụ, một GIF hoạt hình chứa một chuỗi các Image, và mỗi `Image` sau ảnh đầu tiên thường chỉ chứa dữ liệu điểm ảnh cho vùng đã thay đổi, và vùng đó không nhất thiết bắt đầu từ (0, 0).
Cách đúng để duyệt qua các điểm ảnh của `Image` m là:

{{raw `
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
	 for x := b.Min.X; x < b.Max.X; x++ {
	  doStuffWith(m.At(x, y))
	 }
	}
`}}

Các triển khai `Image` không nhất thiết phải dựa trên slice dữ liệu điểm ảnh trong bộ nhớ.
Ví dụ, một [`Uniform`](/pkg/image/#Uniform) là một `Image` với bounds cực lớn và màu đồng nhất, có biểu diễn trong bộ nhớ chỉ là màu đó.

	type Uniform struct {
	    C color.Color
	}

Tuy nhiên, thông thường các chương trình sẽ muốn một ảnh dựa trên slice.
Các kiểu struct như [`RGBA`](/pkg/image/#RGBA) và [`Gray`](/pkg/image/#Gray) (mà các gói khác gọi là `image.RGBA` và `image.Gray`) chứa các slice dữ liệu điểm ảnh và triển khai interface `Image`.

	type RGBA struct {
	    // Pix chứa các điểm ảnh của ảnh theo thứ tự R, G, B, A. Điểm ảnh tại
	    // (x, y) bắt đầu tại Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	    Pix []uint8
	    // Stride là khoảng cách Pix (theo byte) giữa các điểm ảnh liền kề theo chiều dọc.
	    Stride int
	    // Rect là bounds của ảnh.
	    Rect Rectangle
	}

Các kiểu này cũng cung cấp phương thức `Set(x, y int, c color.Color)` cho phép thay đổi ảnh từng điểm ảnh một.

	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	m.Set(5, 5, color.RGBA{255, 0, 0, 255})

Nếu bạn đọc hoặc ghi nhiều dữ liệu điểm ảnh, việc truy cập trực tiếp vào trường `Pix` của kiểu struct có thể hiệu quả hơn, nhưng cũng phức tạp hơn.

Các triển khai `Image` dựa trên slice cũng cung cấp phương thức `SubImage`, trả về một `Image` được hỗ trợ bởi cùng mảng dữ liệu.
Việc thay đổi các điểm ảnh của ảnh con sẽ ảnh hưởng đến các điểm ảnh của ảnh gốc, tương tự như thay đổi nội dung của slice con `s[i0:i1]` sẽ ảnh hưởng đến nội dung của slice gốc `s`.

{{image "image/image-package-05.png"}}

	m0 := image.NewRGBA(image.Rect(0, 0, 8, 5))
	m1 := m0.SubImage(image.Rect(1, 2, 5, 5)).(*image.RGBA)
	fmt.Println(m0.Bounds().Dx(), m1.Bounds().Dx()) // in ra 8, 4
	fmt.Println(m0.Stride == m1.Stride)             // in ra true

Với mã cấp thấp làm việc trên trường `Pix` của ảnh, cần lưu ý rằng duyệt qua `Pix` có thể ảnh hưởng đến các điểm ảnh ngoài bounds của ảnh.
Trong ví dụ trên, các điểm ảnh được bao phủ bởi `m1.Pix` được tô màu xanh lam.
Các mã cấp cao hơn, chẳng hạn như phương thức `At` và `Set` hoặc [gói image/draw](/pkg/image/draw/), sẽ cắt xén các thao tác của chúng theo bounds của ảnh.

## Định dạng ảnh

Thư viện gói chuẩn hỗ trợ nhiều định dạng ảnh phổ biến như GIF, JPEG và PNG.
Nếu bạn biết định dạng của tệp ảnh nguồn, bạn có thể giải mã trực tiếp từ một [`io.Reader`](/pkg/io/#Reader).

	import (
	 "image/jpeg"
	 "image/png"
	 "io"
	)

	// convertJPEGToPNG chuyển đổi từ JPEG sang PNG.
	func convertJPEGToPNG(w io.Writer, r io.Reader) error {
	 img, err := jpeg.Decode(r)
	 if err != nil {
	  return err
	 }
	 return png.Encode(w, img)
	}

Nếu bạn có dữ liệu ảnh không rõ định dạng, hàm [`image.Decode`](/pkg/image/#Decode) có thể phát hiện định dạng.
Tập hợp các định dạng được nhận dạng được xây dựng tại thời điểm chạy và không giới hạn ở những định dạng trong thư viện gói chuẩn.
Một gói định dạng ảnh thường đăng ký định dạng của mình trong hàm init, và gói main sẽ "import gạch dưới" gói đó chỉ để lấy hiệu ứng phụ của việc đăng ký định dạng.

	import (
	 "image"
	 "image/png"
	 "io"

	 _ "code.google.com/p/vp8-go/webp"
	 _ "image/jpeg"
	)

	// convertToPNG chuyển đổi từ bất kỳ định dạng nào được nhận dạng sang PNG.
	func convertToPNG(w io.Writer, r io.Reader) error {
	 img, _, err := image.Decode(r)
	 if err != nil {
	  return err
	 }
	 return png.Encode(w, img)
	}
