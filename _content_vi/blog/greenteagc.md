---
title: Bộ Thu Gom Rác Green Tea
date: 2025-10-29
by:
- Michael Knyszek
- Austin Clements
tags:
- garbage collection
- performance
summary: Go 1.25 bao gồm một bộ thu gom rác thử nghiệm mới, Green Tea.
template: true
---

<style type="text/css" scoped>
  .centered {
	position: relative;
	display: flex;
	flex-direction: column;
	align-items: center;
  }
  div.carousel {
	display: flex;
	width: 100%;
	height: auto;
	overflow-x: auto;
	scroll-snap-type: x mandatory;
	padding-bottom: 1.1em;
  }
  .hide-overflow {
	overflow-x: hidden !important;
  }
  button.scroll-button-left {
	left: 0;
	bottom: 0;
  }
  button.scroll-button-right {
	right: 0;
	bottom: 0;
  }
  button.scroll-button {
	position: absolute;
	font-size: 1em;
	font-family: inherit;
	font-style: oblique;
  }
  figure.carouselitem {
	display: flex;
	flex-direction: column;
	align-items: center;
	margin: 0;
	padding: 0;
	width: 100%;
	flex-shrink: 0;
	scroll-snap-align: start;
  }
  figure.carouselitem figcaption {
	display: table-caption;
	caption-side: top;
	text-align: left;
	width: 80%;
	height: auto;
	padding: 8px;
  }
  figure.captioned {
	display: flex;
	flex-direction: column;
	align-items: center;
	margin: 0 auto;
	padding: 0;
	width: 95%;
  }
  figure.captioned figcaption {
	display: table-caption;
	caption-side: top;
	text-align: center;
	font-style: oblique;
	height: auto;
	padding: 8px;
  }
  div.row {
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: center;
	width: 100%;
  }
</style>

<noscript>
    <center>
    <i>Để có trải nghiệm tốt nhất, hãy xem <a href="/blog/greenteagc">bài đăng blog này</a>
    trong trình duyệt có bật JavaScript.</i>
    </center>
</noscript>

Go 1.25 bao gồm một bộ thu gom rác thử nghiệm mới gọi là Green Tea,
có thể sử dụng bằng cách đặt `GOEXPERIMENT=greenteagc` vào lúc build.
Nhiều workload tốn khoảng 10% thời gian ít hơn trong bộ thu gom rác, nhưng một số
workload có thể giảm đến 40%!

Nó đã sẵn sàng cho môi trường sản xuất và đang được sử dụng tại Google, vì vậy chúng tôi khuyến khích bạn
thử nó.
Chúng tôi biết một số workload không hưởng lợi nhiều, hoặc thậm chí không hưởng lợi gì, vì vậy phản hồi của bạn
rất quan trọng để giúp chúng tôi tiến lên.
Dựa trên dữ liệu chúng tôi có hiện tại, chúng tôi có kế hoạch làm cho nó trở thành mặc định trong Go 1.26.

Để báo cáo bất kỳ vấn đề nào, [tạo issue mới](/issue/new).

Để báo cáo bất kỳ thành công nào, trả lời [issue Green Tea hiện có](
/issue/73581).

Sau đây là một bài đăng blog dựa trên bài nói chuyện GopherCon 2025 của Michael Knyszek.

{{video "https://www.youtube.com/embed/gPJkM95KpKo"}}

## Theo dõi quá trình thu gom rác

Trước khi thảo luận về Green Tea, hãy để chúng tôi đưa mọi người về cùng trang về thu gom rác.

### Đối tượng và con trỏ

Mục đích của thu gom rác là tự động thu hồi và tái sử dụng bộ nhớ
không còn được chương trình sử dụng nữa.

Để làm điều này, bộ thu gom rác Go quan tâm đến *đối tượng* và
*con trỏ*.

Trong ngữ cảnh của Go runtime, *đối tượng* là các giá trị Go có bộ nhớ cơ bản
được phân bổ từ heap.
Các đối tượng heap được tạo ra khi trình biên dịch Go không thể tìm ra cách khác để phân bổ
bộ nhớ cho một giá trị.
Ví dụ, đoạn code sau phân bổ một đối tượng heap đơn: bộ lưu trữ hỗ trợ
cho một slice của các con trỏ.

```
var x = make([]*int, 10) // global
```


Trình biên dịch Go không thể phân bổ bộ lưu trữ hỗ trợ của slice ở bất kỳ đâu ngoài heap,
vì rất khó, và thậm chí có thể là không thể, để nó biết `x` sẽ
tham chiếu đến đối tượng trong bao lâu.

*Con trỏ* chỉ là các số chỉ ra vị trí của một giá trị Go trong bộ nhớ,
và chúng là cách một chương trình Go tham chiếu đến các đối tượng.
Ví dụ, để lấy con trỏ đến đầu của đối tượng được phân bổ trong
đoạn code cuối cùng, chúng ta có thể viết:

```
&x[0] // 0xc000104000
```

### Thuật toán mark-sweep

Bộ thu gom rác Go tuân theo một chiến lược được gọi rộng rãi là *tracing garbage
collection*, có nghĩa là bộ thu gom rác theo dõi, hoặc truy tìm, các
con trỏ trong chương trình để xác định đối tượng nào chương trình vẫn đang sử dụng.

Cụ thể hơn, bộ thu gom rác Go triển khai thuật toán mark-sweep.
Điều này đơn giản hơn nhiều so với nghe có vẻ.
Hãy tưởng tượng các đối tượng và con trỏ như một loại đồ thị, theo nghĩa khoa học máy tính.
Đối tượng là các nút, con trỏ là các cạnh.

Thuật toán mark-sweep hoạt động trên đồ thị này, và như tên gợi ý,
tiến hành theo hai giai đoạn.

Trong giai đoạn đầu tiên, giai đoạn đánh dấu, nó duyệt đồ thị đối tượng từ các
cạnh nguồn được xác định rõ ràng gọi là *roots*.
Hãy nghĩ về các biến toàn cục và cục bộ.
Sau đó, nó *đánh dấu* mọi thứ nó tìm thấy trên đường đi là *đã thăm*, để tránh đi vào
vòng lặp.
Điều này tương tự như thuật toán flood đồ thị điển hình của bạn, như tìm kiếm theo chiều sâu hoặc
tìm kiếm theo chiều rộng.

Tiếp theo là giai đoạn sweep.
Bất kỳ đối tượng nào không được thăm trong quá trình duyệt đồ thị của chúng ta là không được sử dụng, hoặc *không thể tiếp cận*,
bởi chương trình.
Chúng ta gọi trạng thái này là không thể tiếp cận vì không thể với code Go an toàn thông thường
để truy cập bộ nhớ đó nữa, đơn giản thông qua ngữ nghĩa của ngôn ngữ.
Để hoàn thành giai đoạn sweep, thuật toán đơn giản là lặp qua tất cả các
nút chưa được thăm và đánh dấu bộ nhớ của chúng là tự do, để bộ phân bổ bộ nhớ có thể tái sử dụng nó.

### Chỉ vậy thôi sao?

Bạn có thể nghĩ tôi đang đơn giản hóa một chút ở đây.
Các bộ thu gom rác thường được gọi là *ma thuật* và *hộp đen*.
Và bạn sẽ đúng một phần, có nhiều phức tạp hơn.

Ví dụ, thuật toán này, trên thực tế, được thực thi đồng thời với
code Go thông thường của bạn.
Duyệt một đồ thị đang thay đổi bên dưới bạn đem lại những thách thức.
Chúng tôi cũng song song hóa thuật toán này, đây là một chi tiết sẽ xuất hiện lại
sau này.

Nhưng hãy tin tôi khi tôi nói rằng những chi tiết này hầu hết là tách biệt khỏi
thuật toán cốt lõi.
Thực sự chỉ là một graph flood đơn giản ở trung tâm.

### Ví dụ về graph flood

Hãy cùng xem qua một ví dụ.
Điều hướng qua trình chiếu bên dưới để theo dõi.

<noscript>
<i>Cuộn theo chiều ngang qua trình chiếu!</i>
<br />
<br />
Hãy xem xét xem với JavaScript được bật, sẽ thêm các nút "Trước" và "Tiếp theo".
Điều này cho phép bạn nhấp qua trình chiếu mà không có chuyển động cuộn,
sẽ làm nổi bật sự khác biệt giữa các sơ đồ tốt hơn.
<br />
<br />
</noscript>

<div class="centered">
<button type="button" id="marksweep-prev" class="scroll-button scroll-button-left" hidden disabled>← Prev</button>
<button type="button" id="marksweep-next" class="scroll-button scroll-button-right" hidden>Next →</button>
<div id="marksweep" class="carousel">
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-007.png" />
		<figcaption>
		Đây là sơ đồ của một số biến toàn cục và heap Go.
		Hãy phân tích nó từng phần.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-008.png" />
		<figcaption>
		Bên trái đây chúng ta có các roots của mình.
		Đây là các biến toàn cục x và y.
		Chúng sẽ là điểm xuất phát cho quá trình duyệt đồ thị của chúng ta.
		Vì chúng được tô màu xanh lam, theo chú thích tiện dụng của chúng ta ở góc dưới bên trái, chúng hiện đang trong danh sách công việc của chúng ta.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-009.png" />
		<figcaption>
		Ở phía bên phải, chúng ta có heap của mình.
		Hiện tại, mọi thứ trong heap của chúng ta được tô màu xám vì chúng ta chưa thăm bất kỳ phần nào của nó.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-010.png" />
		<figcaption>
		Mỗi một trong những hình chữ nhật này đại diện cho một đối tượng.
		Mỗi đối tượng được gán nhãn với kiểu của nó.
		Đối tượng cụ thể này là một đối tượng kiểu T, có định nghĩa kiểu ở phía trên bên trái.
		Nó có một con trỏ đến một mảng các children, và một số giá trị.
		Chúng ta có thể suy ra đây là một loại cấu trúc dữ liệu cây đệ quy nào đó.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-011.png" />
		<figcaption>
		Ngoài các đối tượng kiểu T, bạn cũng sẽ nhận thấy rằng chúng ta có các đối tượng mảng chứa *T.
		Chúng được trỏ đến bởi trường "children" của các đối tượng kiểu T.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-012.png" />
		<figcaption>
		Mỗi hình vuông bên trong hình chữ nhật đại diện cho 8 byte bộ nhớ.
		Một hình vuông có chấm là một con trỏ.
		Nếu nó có mũi tên, đó là con trỏ không nil trỏ đến một đối tượng khác.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-013.png" />
		<figcaption>
		Và nếu nó không có mũi tên tương ứng, thì đó là con trỏ nil.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-014.png" />
		<figcaption>
		Tiếp theo, những hình chữ nhật đứt nét này đại diện cho không gian trống, cái tôi sẽ gọi là "slot" trống. Chúng ta có thể đặt một đối tượng ở đó, nhưng hiện tại không có.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-015.png" />
		<figcaption>
		Bạn cũng sẽ nhận thấy rằng các đối tượng được nhóm lại với nhau bởi các hình chữ nhật góc tròn đứt nét có nhãn này.
		Mỗi cái trong số này đại diện cho một <i>trang</i>, là một
		khối bộ nhớ có kích thước cố định, liền tiếp và căn chỉnh.
		Trong Go, các trang có kích thước 8 KiB (bất kể kích thước trang bộ nhớ ảo phần cứng).
		Các trang này được gán nhãn A, B, C và D, và tôi sẽ gọi chúng như vậy.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-015.png" />
		<figcaption>
		Trong sơ đồ này, mỗi đối tượng được phân bổ như một phần của một trang nào đó.
		Giống như trong triển khai thực, mỗi trang ở đây chỉ chứa các đối tượng có một kích thước nhất định.
		Đây chỉ là cách heap Go được tổ chức.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-016.png" />
		<figcaption>
		Các trang cũng là cách chúng ta tổ chức metadata theo từng đối tượng.
		Ở đây bạn có thể thấy bảy ô, mỗi ô tương ứng với một trong bảy slot đối tượng trong trang A.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-016.png" />
		<figcaption>
		Mỗi ô đại diện cho một bit thông tin: liệu chúng ta đã thấy đối tượng trước đó chưa.
		Đây thực sự là cách runtime thực quản lý xem một đối tượng đã được thăm hay chưa, và đây sẽ là một chi tiết quan trọng sau này.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-017.png" />
		<figcaption>
		Đó là rất nhiều chi tiết, vì vậy cảm ơn bạn đã đọc theo.
		Tất cả điều này sẽ được đưa vào cuộc chơi sau này.
		Bây giờ, hãy chỉ xem graph flood của chúng ta áp dụng vào bức tranh này như thế nào.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-018.png" />
		<figcaption>
		Chúng ta bắt đầu bằng cách lấy một root ra khỏi danh sách công việc.
		Chúng ta đánh dấu nó màu đỏ để chỉ ra rằng nó hiện đang hoạt động.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-019.png" />
		<figcaption>
		Theo con trỏ của root đó, chúng ta tìm thấy một đối tượng kiểu T, mà chúng ta thêm vào danh sách công việc.
		Theo chú thích của chúng ta, chúng ta vẽ đối tượng màu xanh lam để chỉ ra rằng nó đang trong danh sách công việc.
		Cũng lưu ý rằng chúng ta đặt bit đã thấy tương ứng với đối tượng này trong metadata của chúng ta.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-020.png" />
		<figcaption>
		Tương tự với root tiếp theo.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-021.png" />
		<figcaption>
		Bây giờ chúng ta đã xử lý tất cả các roots, chúng ta còn lại hai đối tượng trong danh sách công việc.
		Hãy lấy một đối tượng ra khỏi danh sách công việc.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-022.png" />
		<figcaption>
		Những gì chúng ta sắp làm bây giờ là duyệt các con trỏ của các đối tượng, để tìm thêm đối tượng.
		Nhân tiện, chúng ta gọi việc duyệt các con trỏ của một đối tượng là "quét" đối tượng.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-023.png" />
		<figcaption>
		Chúng ta tìm thấy đối tượng mảng hợp lệ này...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-024.png" />
		<figcaption>
		... và thêm nó vào danh sách công việc.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-025.png" />
		<figcaption>
		Từ đây, chúng ta tiến hành đệ quy.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-026.png" />
		<figcaption>
		Chúng ta duyệt các con trỏ của mảng.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-027.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-028.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-029.png" />
		<figcaption>
		Tìm thêm một số đối tượng nữa...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-030.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-031.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-032.png" />
		<figcaption>
		Sau đó chúng ta duyệt các đối tượng mà đối tượng mảng đã tham chiếu!
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-033.png" />
		<figcaption>
		Và lưu ý rằng chúng ta vẫn phải duyệt qua tất cả các con trỏ, ngay cả khi chúng là nil.
		Chúng ta không biết trước liệu chúng có nil không.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-034.png" />
		<figcaption>
		Thêm một đối tượng nữa trên nhánh này...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-035.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-036.png" />
		<figcaption>
		Và bây giờ chúng ta đã đến nhánh kia, bắt đầu từ đối tượng trong trang A mà chúng ta tìm thấy trước đó từ một trong các roots.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-036.png" />
		<figcaption>
		Bạn có thể nhận thấy kỷ luật vào sau ra trước đối với danh sách công việc ở đây, cho thấy danh sách công việc của chúng ta là một stack, và do đó graph flood của chúng ta xấp xỉ tìm kiếm theo chiều sâu.
		Đây là có chủ đích, và phản ánh thuật toán graph flood thực tế trong Go runtime.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-037.png" />
		<figcaption>
		Hãy tiếp tục...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-038.png" />
		<figcaption>
		Tiếp theo chúng ta tìm thấy một đối tượng mảng khác...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-039.png" />
		<figcaption>
		Và duyệt nó...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-040.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-041.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-042.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-043.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-044.png" />
		<figcaption>
		Chỉ còn một đối tượng cuối cùng trong danh sách công việc của chúng ta...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-045.png" />
		<figcaption>
		Hãy quét nó...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-046.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-047.png" />
		<figcaption>
		Và chúng ta đã xong giai đoạn đánh dấu! Không có gì chúng ta đang tích cực làm việc và không còn gì trong danh sách công việc của chúng ta.
		Mỗi đối tượng được vẽ màu đen là có thể tiếp cận, và mỗi đối tượng được vẽ màu xám là không thể tiếp cận.
		Hãy sweep tất cả các đối tượng không thể tiếp cận, tất cả cùng một lúc.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/marksweep-048.png" />
		<figcaption>
		Chúng ta đã chuyển đổi những đối tượng đó thành các slot trống, sẵn sàng để chứa các đối tượng mới.
		</figcaption>
	</figure>
</div>
</div>

## Vấn đề

Sau tất cả những điều đó, tôi nghĩ chúng ta đã hiểu bộ thu gom rác Go thực sự đang làm gì.
Quá trình này có vẻ hoạt động tốt đủ ngày nay, vậy vấn đề là gì?

Thì ra chúng ta có thể dành *rất nhiều* thời gian để thực thi thuật toán cụ thể này trong một số
chương trình, và nó thêm chi phí đáng kể vào hầu hết mọi chương trình Go.
Không phải hiếm khi thấy các chương trình Go dành 20% hoặc nhiều hơn thời gian CPU của họ trong
bộ thu gom rác.

Hãy phân tích xem thời gian đó được dành ở đâu.

### Chi phí thu gom rác

Ở cấp độ cao, có hai phần trong chi phí của bộ thu gom rác.
Phần đầu tiên là tần suất nó chạy, và phần thứ hai là lượng công việc nó thực hiện mỗi lần chạy.
Nhân hai cái đó với nhau, và bạn sẽ có tổng chi phí của bộ thu gom rác.

<figure class="captioned">
	<figcaption>
	Tổng chi phí GC = Số chu kỳ GC &times; Chi phí trung bình mỗi chu kỳ GC
	</figcaption>
</figure>

Qua nhiều năm chúng ta đã giải quyết cả hai số hạng trong phương trình này, và để biết thêm về _tần suất_ bộ thu gom rác
chạy, xem [bài nói chuyện GopherCon EU 2022 của Michael](https://www.youtube.com/watch?v=07wduWyWx8M)
về giới hạn bộ nhớ.
[Hướng dẫn về bộ thu gom rác Go](/doc/gc-guide) cũng có nhiều điều để nói về chủ đề này,
và đáng xem nếu bạn muốn tìm hiểu sâu hơn.

Nhưng bây giờ hãy chỉ tập trung vào phần thứ hai, chi phí mỗi chu kỳ.

Từ nhiều năm xem xét kỹ lưỡng các profile CPU để cố gắng cải thiện hiệu năng, chúng tôi biết hai điều lớn
về bộ thu gom rác của Go.

Điều đầu tiên là khoảng 90% chi phí của bộ thu gom rác được dành cho việc đánh dấu,
và chỉ khoảng 10% là sweep.
Hóa ra sweep dễ tối ưu hóa hơn nhiều so với đánh dấu,
và Go đã có một sweeper rất hiệu quả trong nhiều năm.

Điều thứ hai là, trong thời gian dành cho đánh dấu, một phần đáng kể, thường ít nhất 35%, chỉ đơn giản
là bị _đình trệ_ khi truy cập bộ nhớ heap.
Điều này đã đủ tệ, nhưng nó hoàn toàn làm tắc nghẽn những gì làm cho CPU hiện đại
thực sự nhanh.

### "Một thảm họa vi kiến trúc"

"Tắc nghẽn" có nghĩa là gì trong ngữ cảnh này?
Các đặc điểm cụ thể của CPU hiện đại có thể khá phức tạp, vì vậy hãy sử dụng một phép loại suy.

Hãy tưởng tượng CPU đang lái xe trên một con đường, nơi con đường đó là chương trình của bạn.
CPU muốn tăng tốc lên tốc độ cao, và để làm được điều đó nó cần có thể nhìn xa phía trước,
và con đường cần thông thoáng.
Nhưng thuật toán graph flood giống như lái xe qua các đường phố thành phố đối với CPU.
CPU không thể nhìn quanh các góc và không thể dự đoán điều gì sẽ xảy ra tiếp theo.
Để tiến lên, nó liên tục phải chậm lại để rẽ, dừng lại ở đèn giao thông và tránh
người đi bộ.
Động cơ của bạn nhanh đến đâu không quan trọng vì bạn không bao giờ có cơ hội để chạy.

Hãy làm cho điều đó cụ thể hơn bằng cách nhìn lại ví dụ của chúng ta.
Tôi đã phủ lên heap ở đây con đường mà chúng ta đã đi.
Mỗi mũi tên từ trái sang phải đại diện cho một phần công việc quét mà chúng ta đã làm
và các mũi tên đứt nét cho thấy chúng ta đã nhảy giữa các phần công việc quét như thế nào.

<figure class="captioned">
	<img src="greenteagc/graphflood-path.png" />
	<figcaption>
	Con đường qua heap mà bộ thu gom rác đã đi trong ví dụ graph flood của chúng ta.
	</figcaption>
</figure>

Lưu ý rằng chúng ta đã nhảy khắp nơi trong bộ nhớ làm một ít công việc ở mỗi nơi.
Đặc biệt, chúng ta thường xuyên nhảy giữa các trang, và giữa các phần khác nhau của trang.

Các CPU hiện đại thực hiện rất nhiều bộ nhớ đệm.
Truy cập bộ nhớ chính có thể chậm hơn đến 100 lần so với truy cập bộ nhớ trong bộ nhớ đệm của chúng ta.
Bộ nhớ đệm CPU được điền bằng bộ nhớ được truy cập gần đây, và bộ nhớ gần với
bộ nhớ được truy cập gần đây.
Nhưng không có đảm bảo nào rằng bất kỳ hai đối tượng nào trỏ đến nhau cũng sẽ *gần nhau* trong bộ nhớ.
Graph flood không tính đến điều này.

Ghi chú bên: nếu chúng ta chỉ đình trệ khi lấy dữ liệu từ bộ nhớ chính, có thể không quá tệ.
CPU phát hành các yêu cầu bộ nhớ không đồng bộ, vì vậy ngay cả các yêu cầu chậm cũng có thể chồng chéo nếu CPU có thể nhìn
đủ xa phía trước.
Nhưng trong graph flood, mỗi phần công việc nhỏ, không thể đoán trước, và phụ thuộc cao vào
phần trước, vì vậy CPU buộc phải đợi hầu hết mỗi lần lấy bộ nhớ riêng lẻ.

Và thật không may cho chúng ta, vấn đề này chỉ đang trở nên tệ hơn.
Có một câu nói trong ngành rằng "chờ hai năm và code của bạn sẽ nhanh hơn."

Nhưng Go, là ngôn ngữ được thu gom rác dựa vào thuật toán mark-sweep, có nguy cơ ngược lại.
"Chờ hai năm và code của bạn sẽ chậm hơn."
Các xu hướng trong phần cứng CPU hiện đại đang tạo ra những thách thức mới cho hiệu năng bộ thu gom rác:

**Truy cập bộ nhớ không đồng đều.**
Một mặt, bộ nhớ giờ có xu hướng được liên kết với các tập con lõi CPU.
Các truy cập bởi các lõi CPU *khác* đến bộ nhớ đó chậm hơn trước.
Nói cách khác, chi phí của một lần truy cập bộ nhớ chính [phụ thuộc vào lõi CPU nào đang truy cập
nó](https://jprahman.substack.com/p/sapphire-rapids-core-to-core-latency).
Nó không đồng đều, vì vậy chúng ta gọi đây là truy cập bộ nhớ không đồng đều, hay viết tắt là NUMA.

**Giảm băng thông bộ nhớ.**
Băng thông bộ nhớ có sẵn mỗi CPU đang giảm dần theo thời gian.
Điều này chỉ có nghĩa là mặc dù chúng ta có nhiều lõi CPU hơn, mỗi lõi có thể gửi tương đối ít hơn
yêu cầu đến bộ nhớ chính, buộc các yêu cầu không được lưu trong bộ nhớ đệm phải chờ lâu hơn trước.

**Ngày càng nhiều lõi CPU.**
Ở trên, chúng ta đã xem xét thuật toán đánh dấu tuần tự, nhưng bộ thu gom rác thực sự thực hiện
thuật toán này song song.
Điều này mở rộng tốt đến một số lõi CPU hạn chế, nhưng hàng đợi chia sẻ của các đối tượng cần quét trở thành
nút cổ chai, ngay cả với thiết kế cẩn thận.

**Các tính năng phần cứng hiện đại.**
Phần cứng mới có các tính năng ưa thích như hướng dẫn vector, cho phép chúng ta hoạt động trên nhiều dữ liệu cùng một lúc.
Mặc dù điều này có tiềm năng tăng tốc đáng kể, nhưng không rõ ngay cách làm cho điều đó hoạt động cho
đánh dấu vì đánh dấu làm rất nhiều công việc không đều và thường nhỏ.

## Green Tea

Cuối cùng, điều này đưa chúng ta đến Green Tea, cách tiếp cận mới của chúng tôi với thuật toán mark-sweep.
Ý tưởng chính đằng sau Green Tea thật đơn giản đến kinh ngạc:

_Làm việc với các trang, không phải đối tượng._

Nghe có vẻ tầm thường, phải không?
Nhưng vẫn cần rất nhiều công sức để tìm ra cách sắp xếp thứ tự duyệt đồ thị đối tượng và những gì chúng ta cần
theo dõi để làm cho điều này hoạt động tốt trong thực tế.

Cụ thể hơn, điều này có nghĩa là:
* Thay vì quét các đối tượng, chúng ta quét toàn bộ trang.
* Thay vì theo dõi các đối tượng trong danh sách công việc, chúng ta theo dõi toàn bộ trang.
* Chúng ta vẫn cần đánh dấu các đối tượng ở cuối cùng, nhưng chúng ta sẽ theo dõi các đối tượng được đánh dấu cục bộ theo từng
  trang, thay vì trên toàn bộ heap.

### Ví dụ về Green Tea

Hãy xem điều này có nghĩa gì trong thực tế bằng cách nhìn lại heap ví dụ của chúng ta, nhưng lần này
chạy Green Tea thay vì graph flood đơn giản.

Như ở trên, điều hướng qua trình chiếu có chú thích để theo dõi.

<noscript>
<i>Cuộn theo chiều ngang qua trình chiếu!</i>
<br />
<br />
Hãy xem xét xem với JavaScript được bật, sẽ thêm các nút "Trước" và "Tiếp theo".
Điều này cho phép bạn nhấp qua trình chiếu mà không có chuyển động cuộn,
sẽ làm nổi bật sự khác biệt giữa các sơ đồ tốt hơn.
<br />
<br />
</noscript>

<div class="centered">
<button type="button" id="greentea-prev" class="scroll-button scroll-button-left" hidden disabled>← Prev</button>
<button type="button" id="greentea-next" class="scroll-button scroll-button-right" hidden>Next →</button>
<div id="greentea" class="carousel">
	<figure class="carouselitem">
		<img src="greenteagc/greentea-060.png" />
		<figcaption>
		Đây là cùng heap như trước, nhưng bây giờ với hai bit metadata mỗi đối tượng thay vì một.
		Một lần nữa, mỗi bit, hoặc ô, tương ứng với một trong các slot đối tượng trong trang.
		Tổng cộng, bây giờ chúng ta có mười bốn bit tương ứng với bảy slot trong trang A.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-060.png" />
		<figcaption>
		Các bit trên cùng đại diện cho điều tương tự như trước: liệu chúng ta đã thấy một con trỏ đến đối tượng chưa.
		Tôi sẽ gọi những bit này là các bit "đã thấy".
		Tập hợp bit phía dưới là mới.
		Những bit "đã quét" này theo dõi liệu chúng ta đã <i>quét</i> đối tượng chưa.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-060.png" />
		<figcaption>
		Phần metadata mới này là cần thiết vì, trong Green Tea, <b>danh sách công việc theo dõi các trang,
		không phải đối tượng</b>.
		Chúng ta vẫn cần theo dõi các đối tượng ở một mức nào đó, và đó là mục đích của những bit này.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-062.png" />
		<figcaption>
		Chúng ta bắt đầu giống như trước, duyệt các đối tượng từ các roots.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-063.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-064.png" />
		<figcaption>
		Nhưng lần này, thay vì đặt một đối tượng vào danh sách công việc,
		chúng ta đặt toàn bộ trang - trong trường hợp này là trang A - vào danh sách công việc,
		được chỉ ra bằng cách tô màu toàn bộ trang màu xanh lam.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-066.png" />
		<figcaption>
		Đối tượng chúng ta tìm thấy cũng có màu xanh lam để chỉ ra rằng khi chúng ta lấy trang này ra khỏi danh sách công việc, chúng ta sẽ cần nhìn vào đối tượng đó.
		Lưu ý rằng màu xanh lam của đối tượng trực tiếp phản ánh metadata trong trang A.
		Bit đã thấy tương ứng của nó được đặt, nhưng bit đã quét của nó thì không.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-069.png" />
		<figcaption>
		Chúng ta theo root tiếp theo, tìm thấy một đối tượng khác, và một lần nữa đặt toàn bộ trang - trang C - vào danh sách công việc và đặt bit đã thấy của đối tượng.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-071.png" />
		<figcaption>
		Chúng ta đã xong việc theo các roots, vì vậy chúng ta chuyển sang danh sách công việc và lấy trang A ra khỏi danh sách công việc.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-072.png" />
		<figcaption>
		Sử dụng các bit đã thấy và đã quét, chúng ta có thể biết có một đối tượng cần quét trên trang A.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-074.png" />
		<figcaption>
		Chúng ta quét đối tượng đó, theo các con trỏ của nó.
		Và kết quả là, chúng ta thêm trang B vào danh sách công việc, vì đối tượng đầu tiên trong trang A trỏ đến một đối tượng trong trang B.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-075.png" />
		<figcaption>
		Chúng ta đã xong với trang A.
		Tiếp theo chúng ta lấy trang C ra khỏi danh sách công việc.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-076.png" />
		<figcaption>
		Tương tự như trang A, có một đối tượng duy nhất trên trang C cần quét.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-078.png" />
		<figcaption>
		Chúng ta tìm thấy một con trỏ đến một đối tượng khác trong trang B.
		Trang B đã có trong danh sách công việc, vì vậy chúng ta không cần thêm gì vào danh sách công việc.
		Chúng ta chỉ cần đặt bit đã thấy cho đối tượng mục tiêu.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-079.png" />
		<figcaption>
		Bây giờ đến lượt trang B.
		Chúng ta đã tích lũy hai đối tượng cần quét trên trang B,
		và chúng ta có thể xử lý cả hai đối tượng này liên tiếp nhau, theo thứ tự bộ nhớ!
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-081.png" />
		<figcaption>
		Chúng ta duyệt các con trỏ của đối tượng đầu tiên...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-082.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-083.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-084.png" />
		<figcaption>
		Chúng ta tìm thấy một con trỏ đến một đối tượng trong trang A.
		Trang A trước đây đã có trong danh sách công việc, nhưng hiện tại không, vì vậy chúng ta đưa nó trở lại danh sách công việc.
		Không giống như thuật toán mark-sweep ban đầu, nơi bất kỳ đối tượng nào chỉ được thêm vào danh sách công việc
		nhiều nhất một lần mỗi giai đoạn đánh dấu hoàn chỉnh, trong Green Tea, một trang đã cho có thể xuất hiện lại trong danh sách công việc nhiều lần
		trong một giai đoạn đánh dấu.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-085.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-086.png" />
		<figcaption>
		Chúng ta quét đối tượng đã thấy thứ hai trong trang ngay sau đối tượng đầu tiên.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-087.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-088.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-089.png" />
		<figcaption>
		Chúng ta tìm thấy thêm một vài đối tượng trong trang A...
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-090.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-091.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-092.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-093.png" />
		<figcaption>
		Chúng ta đã quét xong trang B, vì vậy chúng ta kéo trang A ra khỏi danh sách công việc.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-094.png" />
		<figcaption>
		Lần này chúng ta chỉ cần quét ba đối tượng, không phải bốn,
		vì chúng ta đã quét đối tượng đầu tiên rồi.
		Chúng ta biết đối tượng nào cần quét bằng cách nhìn vào sự khác biệt giữa các bit "đã thấy" và "đã quét".
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-095.png" />
		<figcaption>
		Chúng ta sẽ quét những đối tượng này theo thứ tự.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-096.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-097.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-098.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-099.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-100.png" />
		<figcaption>
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-101.png" />
		<figcaption>
		Xong rồi! Không còn trang nào trong danh sách công việc và không có gì chúng ta đang tích cực xem xét.
		Lưu ý rằng metadata giờ được căn chỉnh gọn gàng, vì tất cả các đối tượng có thể tiếp cận đều được cả thấy và quét.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-101.png" />
		<figcaption>
		Bạn cũng có thể đã nhận thấy trong quá trình duyệt của chúng ta rằng thứ tự danh sách công việc hơi khác so với graph flood.
		Nơi graph flood có thứ tự vào sau ra trước, hay kiểu stack, ở đây chúng ta sử dụng thứ tự vào trước ra trước, hay kiểu hàng đợi, cho các trang trong danh sách công việc.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-101.png" />
		<figcaption>
		Điều này là có chủ đích.
		Chúng ta để các đối tượng đã thấy tích lũy trên mỗi trang trong khi trang nằm trong hàng đợi, để chúng ta có thể xử lý càng nhiều càng tốt trong một lần.
		Đó là cách chúng ta có thể xử lý nhiều đối tượng trên trang A cùng một lúc.
		Đôi khi sự lười biếng là một đức hạnh.
		</figcaption>
	</figure>
	<figure class="carouselitem">
		<img src="greenteagc/greentea-102.png" />
		<figcaption>
		Và cuối cùng chúng ta có thể quét đi các đối tượng chưa được thăm, như trước.
		</figcaption>
	</figure>
</div>
</div>

### Lên đường cao tốc

Hãy quay lại phép loại suy lái xe của chúng ta.
Cuối cùng chúng ta có lên đường cao tốc không?

Hãy nhớ lại bức tranh graph flood của chúng ta trước đó.

<figure class="captioned">
	<img src="greenteagc/graphflood-path2.png" />
	<figcaption>
	Con đường mà graph flood ban đầu đi qua heap yêu cầu 7 lần quét riêng biệt.
	</figcaption>
</figure>

Chúng ta đã nhảy rất nhiều, làm một ít công việc ở các nơi khác nhau.
Con đường được Green Tea đi trông rất khác.

<figure class="captioned">
	<img src="greenteagc/greentea-path.png" />
	<figcaption>
	Con đường được Green Tea đi chỉ yêu cầu 4 lần quét.
	</figcaption>
</figure>

Green Tea, ngược lại, thực hiện ít lần quét từ trái sang phải hơn nhưng dài hơn trên trang A và B.
Những mũi tên càng dài thì càng tốt, và với heap lớn hơn, hiệu ứng này có thể mạnh hơn nhiều.
*Đó* là điều kỳ diệu của Green Tea.

Đây cũng là cơ hội để chúng ta lên đường cao tốc.

Tất cả điều này cộng lại thành sự phù hợp tốt hơn với vi kiến trúc.
Chúng ta giờ có thể quét các đối tượng gần nhau với xác suất cao hơn nhiều, vì vậy
có khả năng cao hơn chúng ta có thể sử dụng bộ nhớ đệm và tránh bộ nhớ chính.
Tương tự, metadata theo từng trang có nhiều khả năng hơn trong bộ nhớ đệm.
Theo dõi các trang thay vì các đối tượng có nghĩa là danh sách công việc nhỏ hơn,
và ít áp lực hơn trên danh sách công việc có nghĩa là ít tranh chấp hơn và ít đình trệ CPU hơn.

Và nói đến đường cao tốc, chúng ta có thể đưa động cơ ẩn dụ của mình vào các số mà chúng ta chưa bao giờ có thể
trước đây, vì bây giờ chúng ta có thể sử dụng phần cứng vector!

### Tăng tốc vector

Nếu bạn chỉ quen thuộc mơ hồ với phần cứng vector, bạn có thể bị nhầm lẫn về cách chúng ta có thể sử dụng nó ở đây.
Nhưng ngoài các phép tính số học và lượng giác thông thường,
phần cứng vector gần đây hỗ trợ hai thứ có giá trị cho Green Tea:
các thanh ghi rất rộng, và các phép toán bit phức tạp.

Hầu hết các CPU x86 hiện đại hỗ trợ AVX-512, có các thanh ghi vector rộng 512 bit.
Đây đủ rộng để chứa tất cả metadata cho toàn bộ trang chỉ trong hai thanh ghi,
ngay trên CPU, cho phép Green Tea làm việc trên toàn bộ trang chỉ trong một vài hướng dẫn tuần tự.
Phần cứng vector từ lâu đã hỗ trợ các phép toán bit cơ bản trên toàn bộ thanh ghi vector, nhưng bắt đầu
với AMD Zen 4 và Intel Ice Lake, nó cũng hỗ trợ một hướng dẫn "dao đa năng" vector bit mới
cho phép một bước quan trọng của quá trình quét Green Tea được thực hiện chỉ trong vài chu kỳ CPU.
Cùng nhau, chúng cho phép chúng ta tăng tốc vòng lặp quét Green Tea.

Đây thậm chí không phải là tùy chọn cho graph flood, nơi chúng ta sẽ nhảy giữa các đối tượng quét
có tất cả các kích thước khác nhau.
Đôi khi bạn cần hai bit metadata và đôi khi bạn cần mười nghìn.
Đơn giản là không có đủ tính dự đoán được hoặc tính đều đặn để sử dụng phần cứng vector.

Nếu bạn muốn đào sâu vào một số chi tiết, hãy đọc tiếp!
Nếu không, hãy thoải mái bỏ qua phần [đánh giá](#evaluation).

#### Nhân quét AVX-512

Để hiểu quét GC AVX-512 trông như thế nào, hãy xem sơ đồ bên dưới.

<figure class="captioned">
	<img src="greenteagc/avx512.svg" />
	<figcaption>
	Nhân vector AVX-512 để quét.
	</figcaption>
</figure>

Có rất nhiều thứ đang diễn ra ở đây và chúng ta có thể lấp đầy toàn bộ bài đăng blog chỉ về cách điều này hoạt động.
Bây giờ, hãy chỉ phân tích nó ở cấp độ cao:

1. Đầu tiên chúng ta lấy các bit "đã thấy" và "đã quét" cho một trang.
   Nhớ lại, đây là một bit mỗi đối tượng trong trang, và tất cả các đối tượng trong một trang có cùng kích thước.

2. Tiếp theo, chúng ta so sánh hai tập bit.
   Hợp của chúng trở thành các bit "đã quét" mới, trong khi hiệu của chúng là bitmap "đối tượng hoạt động",
   cho chúng ta biết những đối tượng nào chúng ta cần quét trong lần đi qua trang này (so với các lần trước).

3. Chúng ta lấy hiệu của các bitmap và "mở rộng" nó, để thay vì một bit mỗi đối tượng,
   chúng ta có một bit mỗi từ (8 byte) của trang.
   Chúng ta gọi đây là bitmap "từ hoạt động".
   Ví dụ, nếu trang lưu trữ các đối tượng 6 từ (48 byte), mỗi bit trong bitmap đối tượng hoạt động
   sẽ được sao chép thành 6 bit trong bitmap từ hoạt động.
   Như sau:

<figure class="captioned">
	<div class="row"><pre>0 0 1 1 ...</pre> &rarr; <pre>000000 000000 111111 111111 ...</pre></div>
</figure>

4. Tiếp theo chúng ta lấy bitmap con trỏ/vô hướng cho trang.
   Ở đây cũng vậy, mỗi bit tương ứng với một từ (8 byte) của trang, và nó cho chúng ta biết liệu từ đó
   có lưu trữ một con trỏ không.
   Dữ liệu này được quản lý bởi bộ phân bổ bộ nhớ.

5. Bây giờ, chúng ta lấy giao điểm của bitmap con trỏ/vô hướng và bitmap từ hoạt động.
   Kết quả là "bitmap con trỏ hoạt động": một bitmap cho chúng ta biết vị trí của mỗi
   con trỏ trong toàn bộ trang có trong bất kỳ đối tượng đang hoạt động nào chúng ta chưa quét.

6. Cuối cùng, chúng ta có thể lặp qua bộ nhớ của trang và thu thập tất cả các con trỏ.
   Về mặt logic, chúng ta lặp qua mỗi bit được đặt trong bitmap con trỏ hoạt động,
   tải giá trị con trỏ tại từ đó, và ghi lại vào bộ đệm sẽ
   sau này được sử dụng để đánh dấu các đối tượng đã thấy và thêm các trang vào danh sách công việc.
   Sử dụng hướng dẫn vector, chúng ta có thể làm điều này 64 byte mỗi lần,
   chỉ trong một vài hướng dẫn.

Một phần làm cho điều này nhanh là hướng dẫn `VGF2P8AFFINEQB`,
một phần của phần mở rộng x86 "Galois Field New Instructions",
và dao đa năng thao tác bit mà chúng ta đề cập ở trên.
Đó là ngôi sao thực sự của buổi diễn, vì nó cho phép chúng ta thực hiện bước (3) trong nhân quét rất, rất
hiệu quả.
Nó thực hiện [biến đổi affine](https://en.wikipedia.org/wiki/Affine_transformation) theo bit,
xử lý mỗi byte trong vector như chính nó là một vector toán học của 8 bit
và nhân nó với một ma trận bit 8x8.
Tất cả điều này được thực hiện trên [trường Galois](https://en.wikipedia.org/wiki/Finite_field) `GF(2)`,
chỉ có nghĩa là phép nhân là AND và phép cộng là XOR.
Kết quả là chúng ta có thể định nghĩa một vài ma trận bit 8x8 cho mỗi
kích thước đối tượng thực hiện chính xác việc mở rộng bit 1:n mà chúng ta cần.

Để xem code assembly đầy đủ, xem [tệp này](https://cs.opensource.google/go/go/+/master:src/internal/runtime/gc/scan/scan_amd64.s;l=23;drc=041f564b3e6fa3f4af13a01b94db14c1ee8a42e0).
Các "expander" sử dụng các ma trận và hoán vị khác nhau cho mỗi lớp kích thước,
vì vậy chúng ở trong một [tệp riêng biệt](https://cs.opensource.google/go/go/+/master:src/internal/runtime/gc/scan/expand_amd64.s;drc=041f564b3e6fa3f4af13a01b94db14c1ee8a42e0)
được viết bởi một [trình tạo code](https://cs.opensource.google/go/go/+/master:src/internal/runtime/gc/scan/mkasm.go;drc=041f564b3e6fa3f4af13a01b94db14c1ee8a42e0).
Ngoài các hàm mở rộng, thực sự không có nhiều code.
Hầu hết nó được đơn giản hóa đáng kể bởi thực tế là chúng ta có thể thực hiện hầu hết các
phép toán ở trên trên dữ liệu nằm thuần túy trong các thanh ghi.
Và, hy vọng sớm thôi code assembly này [sẽ được thay thế bằng code Go](/issue/73787)!

Nhờ Austin Clements đã nghĩ ra quá trình này.
Nó cực kỳ thú vị, và cực kỳ nhanh!

### Đánh giá {#evaluation}

Vậy đó là cách nó hoạt động.
Nó thực sự giúp được bao nhiêu?

Có thể khá nhiều.
Ngay cả không có các cải tiến vector, chúng tôi thấy giảm chi phí CPU thu gom rác
từ 10% đến 40% trong bộ benchmark của chúng tôi.
Ví dụ, nếu một ứng dụng dành 10% thời gian trong bộ thu gom rác, thì đó
sẽ chuyển thành giảm CPU tổng thể từ 1% đến 4%, tùy thuộc vào đặc thù của
workload.
Mức giảm 10% thời gian CPU thu gom rác là cải tiến xấp xỉ phổ biến nhất.
(Xem [GitHub issue](/issue/73581) để biết một số chi tiết này.)

Chúng tôi đã triển khai Green Tea trong nội bộ Google, và chúng tôi thấy kết quả tương tự ở quy mô.

Chúng tôi vẫn đang triển khai các cải tiến vector,
nhưng benchmark và kết quả sớm cho thấy điều này sẽ mang lại thêm 10% giảm CPU GC.

Mặc dù hầu hết workload đều hưởng lợi ở một mức nào đó, có một số không.

Green Tea dựa trên giả thuyết rằng chúng ta có thể tích lũy đủ đối tượng để quét trên một
trang duy nhất trong một lần để bù đắp chi phí của quá trình tích lũy.
Điều này rõ ràng là đúng nếu heap có cấu trúc rất đều đặn: các đối tượng cùng kích thước ở
độ sâu tương tự trong đồ thị đối tượng.
Nhưng có một số workload thường yêu cầu chúng ta quét chỉ một đối tượng duy nhất mỗi trang mỗi lần.
Điều này có thể tệ hơn so với graph flood vì chúng ta có thể đang làm nhiều công việc hơn trước trong khi
cố gắng tích lũy các đối tượng trên các trang và thất bại.

Triển khai của Green Tea có trường hợp đặc biệt cho các trang chỉ có một đối tượng cần quét.
Điều này giúp giảm hồi quy, nhưng không loại bỏ hoàn toàn chúng.

Tuy nhiên, cần ít tích lũy theo từng trang hơn để vượt qua graph flood
so với bạn có thể mong đợi.
Một kết quả bất ngờ của công việc này là quét chỉ 2% trang mỗi lần
có thể mang lại cải tiến so với graph flood.

### Khả dụng

Green Tea đã có sẵn như một thử nghiệm trong bản phát hành Go 1.25 gần đây và có thể được bật
bằng cách đặt biến môi trường `GOEXPERIMENT` thành `greenteagc` vào lúc build.
Điều này không bao gồm tăng tốc vector được đề cập ở trên.

Chúng tôi kỳ vọng sẽ làm cho nó trở thành bộ thu gom rác mặc định trong Go 1.26, nhưng bạn vẫn có thể chọn không sử dụng
với `GOEXPERIMENT=nogreenteagc` vào lúc build.
Go 1.26 cũng sẽ thêm tăng tốc vector trên phần cứng x86 mới hơn, và bao gồm rất nhiều
điều chỉnh và cải tiến dựa trên phản hồi chúng tôi đã thu thập cho đến nay.

Nếu có thể, chúng tôi khuyến khích bạn thử tại Go tip-of-tree!
Nếu bạn thích sử dụng Go 1.25, chúng tôi vẫn rất muốn nghe phản hồi của bạn.
Xem [bình luận GitHub này](/issue/73581#issuecomment-2847696497) với một số chi tiết về
những gì chẩn đoán chúng tôi muốn xem, nếu bạn có thể chia sẻ, và các kênh ưa thích để
báo cáo phản hồi.

## Hành trình

Trước khi kết thúc bài đăng blog này, hãy dành một chút thời gian để nói về hành trình đã đưa chúng ta đến đây.
Yếu tố con người của công nghệ.

Cốt lõi của Green Tea có thể trông giống như một ý tưởng đơn giản, duy nhất.
Giống như tia lửa cảm hứng mà chỉ một người duy nhất có.

Nhưng điều đó hoàn toàn không đúng.
Green Tea là kết quả của công việc và ý tưởng từ nhiều người trong nhiều năm.
Một số người trong nhóm Go đã đóng góp vào các ý tưởng, bao gồm Michael Pratt, Cherry Mui, David
Chase và Keith Randall.
Những hiểu biết vi kiến trúc từ Yves Vandriessche, khi đó đang ở Intel, cũng đã giúp
định hướng khám phá thiết kế.
Có rất nhiều ý tưởng không hiệu quả, và có rất nhiều chi tiết cần tìm hiểu.
Chỉ để làm cho ý tưởng đơn giản, duy nhất này khả thi.

<figure class="captioned">
	<img src="greenteagc/timeline.png" />
	<figcaption>
	Một dòng thời gian mô tả một tập hợp con các ý tưởng chúng tôi đã thử theo hướng này trước khi đạt đến
	nơi chúng ta đang ở ngày nay.
	</figcaption>
</figure>

Hạt giống của ý tưởng này đi ngược lại đến năm 2018.
Điều buồn cười là mọi người trong nhóm đều nghĩ rằng người khác đã nghĩ ra ý tưởng ban đầu này.

Green Tea có tên vào năm 2024 khi Austin xây dựng một nguyên mẫu của phiên bản trước đó trong khi
lang thang qua các quán cà phê ở Nhật Bản và uống RẤT NHIỀU matcha!
Nguyên mẫu này cho thấy ý tưởng cốt lõi của Green Tea là khả thi.
Và từ đó chúng ta đã bắt đầu cuộc đua.

Trong suốt năm 2025, khi Michael triển khai và đưa Green Tea vào sản xuất, các ý tưởng đã phát triển và thay đổi thậm chí
hơn nữa.

Điều này đòi hỏi rất nhiều khám phá hợp tác vì Green Tea không chỉ là một thuật toán, mà là toàn bộ
không gian thiết kế.
Một không gian mà chúng tôi không nghĩ bất kỳ ai trong chúng ta có thể tự mình điều hướng.
Không đủ khi chỉ có ý tưởng, mà bạn cần tìm hiểu các chi tiết và chứng minh nó.
Và bây giờ chúng ta đã làm được điều đó, cuối cùng chúng ta có thể lặp lại.

Tương lai của Green Tea thật sáng sủa.

Một lần nữa, hãy thử nó bằng cách đặt `GOEXPERIMENT=greenteagc` và cho chúng tôi biết kết quả!
Chúng tôi thực sự hào hứng về công việc này và muốn nghe từ bạn!

<script src="greenteagc/carousel.js"></script>
