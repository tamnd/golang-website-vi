---
title: Khớp ngôn ngữ và vùng miền trong Go
date: 2016-02-09
by:
- Marcel van Lohuizen
tags:
- language
- locale
- tag
- BCP
- 47
- matching
summary: Cách quốc tế hóa trang web của bạn với tính năng khớp ngôn ngữ và vùng miền của Go.
template: true
---

## Giới thiệu

Hãy xem xét một ứng dụng, chẳng hạn như một trang web, có hỗ trợ nhiều ngôn ngữ
trong giao diện người dùng.
Khi người dùng truy cập với danh sách các ngôn ngữ ưu tiên, ứng dụng phải
quyết định ngôn ngữ nào nên được sử dụng trong phần trình bày cho người dùng.
Điều này đòi hỏi phải tìm sự khớp tốt nhất giữa các ngôn ngữ mà ứng dụng hỗ trợ
và những ngôn ngữ người dùng ưu tiên.
Bài viết này giải thích tại sao đây là một quyết định khó và Go có thể giúp như thế nào.

## Thẻ ngôn ngữ

Thẻ ngôn ngữ, còn được gọi là định danh vùng miền, là các định danh có thể đọc bằng máy
cho ngôn ngữ và/hoặc phương ngữ đang được sử dụng.
Tài liệu tham khảo phổ biến nhất cho chúng là tiêu chuẩn IETF BCP 47, và đó là
tiêu chuẩn mà các thư viện Go tuân theo.
Dưới đây là một số ví dụ về thẻ ngôn ngữ BCP 47 và ngôn ngữ hoặc phương ngữ mà chúng
đại diện.

{{raw (file "matchlang/tags.html")}}

Dạng tổng quát của thẻ ngôn ngữ là
một mã ngôn ngữ ("en", "cmn", "zh", "nl", "az" ở trên)
theo sau là subtag tùy chọn cho chữ viết ("-Arab"),
vùng ("-US", "-BE", "-419"),
biến thể ("-oxendict" cho chính tả Từ điển Oxford),
và phần mở rộng ("-u-co-phonebk" để sắp xếp theo danh bạ điện thoại).
Dạng phổ biến nhất được giả định nếu subtag bị bỏ qua, ví dụ
"az-Latn-AZ" cho "az".

Việc sử dụng thẻ ngôn ngữ phổ biến nhất là chọn từ một tập hợp ngôn ngữ được hệ thống hỗ trợ
theo danh sách ưu tiên ngôn ngữ của người dùng, ví dụ
quyết định rằng người dùng ưu tiên tiếng Afrikaans sẽ được phục vụ tốt nhất (giả sử
tiếng Afrikaans không có sẵn) bằng cách hiển thị tiếng Hà Lan. Việc giải quyết các kết quả khớp như vậy
liên quan đến việc tham khảo dữ liệu về khả năng hiểu ngôn ngữ lẫn nhau.

Thẻ ngôn ngữ kết quả từ sự khớp này sau đó được sử dụng để lấy
các tài nguyên đặc thù theo ngôn ngữ như bản dịch, thứ tự sắp xếp,
và thuật toán viết hoa.
Điều này liên quan đến một loại khớp khác. Ví dụ, vì không có thứ tự sắp xếp cụ thể
cho tiếng Bồ Đào Nha, package collate có thể rơi về thứ tự sắp xếp
mặc định, hay ngôn ngữ "gốc".

## Tính phức tạp của việc khớp ngôn ngữ

Xử lý thẻ ngôn ngữ rất phức tạp.
Một phần là vì ranh giới của các ngôn ngữ con người không được định nghĩa rõ ràng
và một phần là do sự kế thừa của các tiêu chuẩn thẻ ngôn ngữ đang phát triển.
Trong phần này, chúng ta sẽ chỉ ra một số khía cạnh phức tạp của việc xử lý thẻ ngôn ngữ.

_Thẻ với các mã ngôn ngữ khác nhau có thể chỉ cùng một ngôn ngữ_

Vì lý do lịch sử và chính trị, nhiều mã ngôn ngữ đã thay đổi theo
thời gian, để lại các ngôn ngữ có cả mã cũ lẫn mã mới.
Nhưng ngay cả hai mã hiện tại cũng có thể chỉ cùng một ngôn ngữ.
Ví dụ, mã ngôn ngữ chính thức cho tiếng Phổ thông là "cmn", nhưng "zh" là
định danh được sử dụng phổ biến nhất cho ngôn ngữ này.
Mã "zh" chính thức được dành riêng cho cái gọi là ngôn ngữ vĩ mô, xác định
nhóm các ngôn ngữ Trung Hoa.
Thẻ cho các ngôn ngữ vĩ mô thường được sử dụng thay thế cho nhau với ngôn ngữ được nói nhiều nhất
trong nhóm.

_Chỉ khớp mã ngôn ngữ là không đủ_

Ví dụ, tiếng Azerbaijan ("az") được viết bằng các chữ viết khác nhau tùy theo
quốc gia nơi nó được nói: "az-Latn" cho Latin (chữ viết mặc định),
"az-Arab" cho tiếng Ả Rập, và "az-Cyrl" cho Cyrillic.
Nếu bạn thay thế "az-Arab" bằng chỉ "az", kết quả sẽ bằng chữ Latin và
có thể không thể hiểu được với người dùng chỉ biết dạng chữ Ả Rập.

Ngoài ra, các vùng khác nhau có thể ngụ ý các chữ viết khác nhau.
Ví dụ: "zh-TW" và "zh-SG" lần lượt ngụ ý việc sử dụng chữ Hán Phồn thể và
Giản thể. Một ví dụ khác, "sr" (tiếng Serbia) mặc định là chữ Cyrillic,
nhưng "sr-RU" (tiếng Serbia được viết ở Nga) ngụ ý chữ Latin!
Điều tương tự có thể nói về tiếng Kyrgyz và các ngôn ngữ khác.

Nếu bạn bỏ qua các subtag, bạn cũng có thể trình bày tiếng Hy Lạp cho người dùng.

_Kết quả khớp tốt nhất có thể là ngôn ngữ không có trong danh sách của người dùng_

Dạng viết phổ biến nhất của tiếng Na Uy ("nb") trông rất giống tiếng Đan Mạch.
Nếu tiếng Na Uy không có sẵn, tiếng Đan Mạch có thể là lựa chọn thứ hai tốt.
Tương tự, người dùng yêu cầu tiếng Đức Thụy Sĩ ("gsw") có thể sẽ hài lòng khi
được hiển thị tiếng Đức ("de"), mặc dù điều ngược lại thì xa vời.
Người dùng yêu cầu tiếng Duy Ngô Nhĩ có thể thích chuyển về tiếng Trung hơn là tiếng Anh.
Còn nhiều ví dụ khác.
Nếu ngôn ngữ được người dùng yêu cầu không được hỗ trợ, việc chuyển về tiếng Anh thường
không phải là điều tốt nhất nên làm.

_Lựa chọn ngôn ngữ quyết định nhiều hơn chỉ bản dịch_

Giả sử người dùng yêu cầu tiếng Đan Mạch, với tiếng Đức là lựa chọn thứ hai.
Nếu ứng dụng chọn tiếng Đức, không chỉ phải sử dụng bản dịch tiếng Đức
mà còn phải sử dụng thứ tự sắp xếp tiếng Đức (không phải tiếng Đan Mạch).
Nếu không, ví dụ, một danh sách động vật có thể sắp xếp "Bär" trước "Äffin".

Việc chọn ngôn ngữ được hỗ trợ dựa trên ngôn ngữ ưu tiên của người dùng giống như
một thuật toán bắt tay: trước tiên bạn xác định giao thức giao tiếp nào sẽ dùng (ngôn ngữ)
và sau đó bạn gắn bó với giao thức này trong tất cả các giao tiếp trong
suốt phiên làm việc.

_Sử dụng ngôn ngữ "cha" như dự phòng không đơn giản_

Giả sử ứng dụng của bạn hỗ trợ tiếng Bồ Đào Nha Angola ("pt-AO").
Các package trong [golang.org/x/text](https://golang.org/x/text), như collation và display, có thể không
có hỗ trợ cụ thể cho phương ngữ này.
Cách hành động đúng đắn trong những trường hợp như vậy là khớp với phương ngữ cha gần nhất.
Các ngôn ngữ được sắp xếp theo phân cấp, với mỗi ngôn ngữ cụ thể có một ngôn ngữ cha tổng quát hơn.
Ví dụ, cha của "en-GB-oxendict" là "en-GB", cha của nó là "en",
cha của nó là ngôn ngữ không xác định "und", còn được gọi là ngôn ngữ gốc.
Trong trường hợp collation, không có thứ tự sắp xếp cụ thể cho tiếng Bồ Đào Nha,
vì vậy package collate sẽ chọn thứ tự sắp xếp của ngôn ngữ gốc.
Cha gần nhất với tiếng Bồ Đào Nha Angola được hỗ trợ bởi package display là
tiếng Bồ Đào Nha châu Âu ("pt-PT") chứ không phải "pt" rõ ràng hơn, vốn ngụ ý
tiếng Bồ Đào Nha Brazil.

Nói chung, các mối quan hệ cha không đơn giản.
Để đưa ra thêm một vài ví dụ, cha của "es-CL" là "es-419", cha của
"zh-TW" là "zh-Hant", và cha của "zh-Hant" là "und".
Nếu bạn tính toán cha bằng cách đơn giản là xóa subtag, bạn có thể chọn một "phương ngữ"
không thể hiểu được với người dùng.

## Khớp ngôn ngữ trong Go

Package Go [golang.org/x/text/language](https://golang.org/x/text/language) cài đặt tiêu chuẩn BCP 47
cho thẻ ngôn ngữ và bổ sung hỗ trợ quyết định ngôn ngữ nào sẽ sử dụng
dựa trên dữ liệu được công bố trong Unicode Common Locale Data Repository (CLDR).

Đây là một chương trình mẫu, được giải thích bên dưới, khớp ngôn ngữ ưu tiên của người dùng
với các ngôn ngữ được ứng dụng hỗ trợ:

{{code "matchlang/complete.go"}}

### Tạo thẻ ngôn ngữ

Cách đơn giản nhất để tạo language.Tag từ chuỗi mã ngôn ngữ do người dùng cung cấp
là dùng language.Make.
Nó trích xuất thông tin có ý nghĩa ngay cả từ đầu vào không hợp lệ.
Ví dụ, "en-USD" sẽ cho kết quả "en" mặc dù USD không phải là subtag hợp lệ.

Make không trả về lỗi.
Thông thường người ta dùng ngôn ngữ mặc định khi xảy ra lỗi, vì vậy
điều này làm cho nó thuận tiện hơn. Dùng Parse để xử lý bất kỳ lỗi nào theo cách thủ công.

Header HTTP Accept-Language thường được dùng để truyền các ngôn ngữ mong muốn của người dùng.
Hàm ParseAcceptLanguage phân tích nó thành một slice thẻ ngôn ngữ,
được sắp xếp theo thứ tự ưu tiên.

Theo mặc định, package language không chuẩn hóa thẻ.
Ví dụ, nó không tuân theo khuyến nghị BCP 47 về loại bỏ chữ viết
nếu đó là lựa chọn phổ biến trong "đại đa số".
Nó cũng bỏ qua các khuyến nghị CLDR: "cmn" không được thay thế bằng "zh" và
"zh-Hant-HK" không được đơn giản hóa thành "zh-HK".
Việc chuẩn hóa thẻ có thể loại bỏ thông tin hữu ích về ý định của người dùng.
Việc chuẩn hóa được xử lý trong Matcher.
Một mảng đầy đủ các tùy chọn chuẩn hóa có sẵn nếu lập trình viên vẫn
muốn thực hiện điều đó.

### Khớp ngôn ngữ ưu tiên của người dùng với ngôn ngữ được hỗ trợ

Matcher khớp ngôn ngữ ưu tiên của người dùng với ngôn ngữ được hỗ trợ.
Người dùng được khuyến khích mạnh mẽ nên dùng nó nếu họ không muốn xử lý tất cả
những phức tạp của việc khớp ngôn ngữ.

Phương thức Match có thể truyền qua các cài đặt người dùng (từ phần mở rộng BCP 47) từ
các thẻ ưu tiên sang thẻ được hỗ trợ được chọn.
Vì vậy, điều quan trọng là thẻ được trả về bởi Match được sử dụng để lấy
các tài nguyên đặc thù theo ngôn ngữ.
Ví dụ, "de-u-co-phonebk" yêu cầu thứ tự danh bạ điện thoại cho tiếng Đức.
Phần mở rộng này bị bỏ qua để khớp, nhưng được package collate sử dụng để
chọn biến thể thứ tự sắp xếp tương ứng.

Matcher được khởi tạo với các ngôn ngữ được ứng dụng hỗ trợ, thường là
các ngôn ngữ có bản dịch.
Tập hợp này thường cố định, cho phép tạo matcher khi khởi động.
Matcher được tối ưu hóa để cải thiện hiệu suất của Match với chi phí
của việc khởi tạo.

Package language cung cấp một tập hợp các thẻ ngôn ngữ được sử dụng phổ biến nhất
có thể được dùng để định nghĩa tập hợp được hỗ trợ.
Người dùng thường không cần lo lắng về việc chọn thẻ chính xác cho các ngôn ngữ được hỗ trợ.
Ví dụ, AmericanEnglish ("en-US") có thể được dùng thay thế cho nhau với
tiếng Anh phổ biến hơn ("en"), mặc định là tiếng Mỹ.
Tất cả đều giống nhau với Matcher. Ứng dụng thậm chí có thể thêm cả hai, cho phép
tiếng lóng Mỹ cụ thể hơn cho "en-US".

### Ví dụ khớp

Hãy xem xét Matcher và danh sách các ngôn ngữ được hỗ trợ sau:

	var supported = []language.Tag{
		language.AmericanEnglish,    // en-US: first language is fallback
		language.German,             // de
		language.Dutch,              // nl
		language.Portuguese          // pt (defaults to Brazilian)
		language.EuropeanPortuguese, // pt-pT
		language.Romanian            // ro
		language.Serbian,            // sr (defaults to Cyrillic script)
		language.SerbianLatin,       // sr-Latn
		language.SimplifiedChinese,  // zh-Hans
		language.TraditionalChinese, // zh-Hant
	}
	var matcher = language.NewMatcher(supported)

Hãy xem kết quả khớp với danh sách ngôn ngữ được hỗ trợ này cho các ưu tiên người dùng khác nhau.

Với ưu tiên người dùng là "he" (tiếng Hebrew), kết quả khớp tốt nhất là "en-US" (tiếng Anh Mỹ).
Không có kết quả khớp tốt, vì vậy matcher sử dụng ngôn ngữ dự phòng (cái đầu tiên trong
danh sách được hỗ trợ).

Với ưu tiên người dùng là "hr" (tiếng Croatia), kết quả khớp tốt nhất là "sr-Latn" (tiếng Serbia
với chữ Latin), vì khi được viết bằng cùng một chữ viết, tiếng Serbia
và tiếng Croatia có thể hiểu lẫn nhau.

Với ưu tiên người dùng là "ru, mo" (tiếng Nga, rồi tiếng Moldova), kết quả khớp tốt nhất là
"ro" (tiếng Romania), vì tiếng Moldova hiện được phân loại chính thức là "ro-MD"
(tiếng Romania ở Moldova).

Với ưu tiên người dùng là "zh-TW" (tiếng Phổ thông ở Đài Loan), kết quả khớp tốt nhất là
"zh-Hant" (tiếng Phổ thông viết bằng chữ Hán Phồn thể), không phải "zh-Hans" (tiếng Phổ thông
viết bằng chữ Hán Giản thể).

Với ưu tiên người dùng là "af, ar" (tiếng Afrikaans, rồi tiếng Ả Rập), kết quả khớp tốt nhất là
"nl" (tiếng Hà Lan). Không có ưu tiên nào được hỗ trợ trực tiếp, nhưng tiếng Hà Lan là
kết quả khớp gần hơn đáng kể với tiếng Afrikaans so với ngôn ngữ dự phòng tiếng Anh với
cả hai.

Với ưu tiên người dùng là "pt-AO, id" (tiếng Bồ Đào Nha Angola, rồi tiếng Indonesia),
kết quả khớp tốt nhất là "pt-PT" (tiếng Bồ Đào Nha châu Âu), không phải "pt" (tiếng Bồ Đào Nha Brazil).

Với ưu tiên người dùng là "gsw-u-co-phonebk" (tiếng Đức Thụy Sĩ với thứ tự sắp xếp danh bạ điện thoại),
kết quả khớp tốt nhất là "de-u-co-phonebk" (tiếng Đức với thứ tự sắp xếp danh bạ điện thoại).
Tiếng Đức là kết quả khớp tốt nhất cho tiếng Đức Thụy Sĩ trong danh sách ngôn ngữ của máy chủ, và
tùy chọn thứ tự sắp xếp danh bạ điện thoại đã được truyền qua.

### Điểm độ tin cậy

Go sử dụng tính điểm độ tin cậy dạng thô với việc loại bỏ dựa trên quy tắc.
Kết quả khớp được phân loại là Exact (Chính xác), High (Cao, không chính xác nhưng không có sự mơ hồ đã biết), Low
(Thấp, có thể là kết quả khớp đúng nhưng không chắc), hoặc No (Không khớp).
Trong trường hợp có nhiều kết quả khớp, có một tập hợp các quy tắc phân định tie-breaking được
thực thi theo thứ tự.
Kết quả khớp đầu tiên được trả về trong trường hợp có nhiều kết quả khớp bằng nhau.
Các điểm độ tin cậy này có thể hữu ích, ví dụ, để từ chối các kết quả khớp tương đối yếu.
Chúng cũng được dùng để tính điểm, ví dụ, vùng hoặc chữ viết có khả năng nhất từ
một thẻ ngôn ngữ.

Các cài đặt trong các ngôn ngữ khác thường sử dụng tính điểm chi tiết hơn theo thang biến thiên.
Chúng tôi nhận thấy rằng việc sử dụng tính điểm thô trong cài đặt Go cuối cùng
đơn giản hơn để cài đặt, dễ bảo trì hơn và nhanh hơn, nghĩa là chúng tôi có thể
xử lý nhiều quy tắc hơn.

### Hiển thị các ngôn ngữ được hỗ trợ

Package [golang.org/x/text/language/display](https://golang.org/x/text/language/display) cho phép đặt tên cho thẻ ngôn ngữ
bằng nhiều ngôn ngữ khác nhau.
Nó cũng chứa "namer Self" để hiển thị thẻ bằng ngôn ngữ của chính nó.

Ví dụ:

{{code "matchlang/display.go" `/START/` `/END/`}}

in ra

	English              (English)
	French               (français)
	Dutch                (Nederlands)
	Flemish              (Vlaams)
	Simplified Chinese   (简体中文)
	Traditional Chinese  (繁體中文)
	Russian              (русский)

Trong cột thứ hai, hãy chú ý sự khác biệt trong cách viết hoa, phản ánh
các quy tắc của ngôn ngữ tương ứng.

## Kết luận

Nhìn thoáng qua, thẻ ngôn ngữ trông giống như dữ liệu có cấu trúc gọn gàng, nhưng vì
chúng mô tả các ngôn ngữ của con người, cấu trúc các mối quan hệ giữa các thẻ ngôn ngữ
thực ra khá phức tạp.
Thường rất hấp dẫn, đặc biệt với các lập trình viên nói tiếng Anh, để viết
code khớp ngôn ngữ đặc biệt chỉ dùng thao tác chuỗi trên thẻ ngôn ngữ.
Như mô tả ở trên, điều này có thể tạo ra kết quả tệ.

Package [golang.org/x/text/language](https://golang.org/x/text/language) của Go giải quyết vấn đề phức tạp này
trong khi vẫn cung cấp API đơn giản, dễ sử dụng. Hãy tận hưởng.
