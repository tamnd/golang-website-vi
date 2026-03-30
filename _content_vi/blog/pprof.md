---
title: Profiling chương trình Go
date: 2011-06-24
by:
- Russ Cox, July 2011; updated by Shenghou Ma, May 2013
tags:
- benchmark
- pprof
- profiling
- technical
summary: Cách sử dụng profiler tích hợp của Go để hiểu và tối ưu hóa các chương trình của bạn.
template: true
---


Tại Scala Days 2011, Robert Hundt đã trình bày một bài báo có tiêu đề
[Loop Recognition in C++/Java/Go/Scala.](http://research.google.com/pubs/pub37122.html)
Bài báo triển khai một thuật toán tìm vòng lặp cụ thể, chẳng hạn như bạn có thể dùng
trong một pass phân tích luồng của compiler, bằng C++, Go, Java, Scala, và sau đó sử dụng
các chương trình đó để rút ra kết luận về các vấn đề hiệu suất điển hình trong các ngôn ngữ này.
Chương trình Go được trình bày trong bài báo đó chạy khá chậm, khiến nó trở thành
một cơ hội tuyệt vời để thể hiện cách sử dụng các công cụ profiling của Go để đưa một
chương trình chậm trở nên nhanh hơn.

_Bằng cách sử dụng các công cụ profiling của Go để xác định và sửa chữa các điểm nghẽn cụ thể, chúng ta có thể làm cho chương trình tìm vòng lặp Go chạy nhanh hơn một bậc độ lớn và sử dụng ít bộ nhớ hơn 6 lần._
(Cập nhật: Do các tối ưu hóa gần đây của `libstdc++` trong `gcc`, mức giảm bộ nhớ hiện là 3.7 lần.)

Bài báo của Hundt không chỉ định phiên bản nào của các công cụ C++, Go, Java và Scala
mà ông đã sử dụng.
Trong bài đăng blog này, chúng tôi sẽ sử dụng snapshot hàng tuần gần nhất của compiler Go `6g`
và phiên bản `g++` đi kèm với bản phân phối Ubuntu Natty.
(Chúng tôi sẽ không sử dụng Java hay Scala, vì chúng tôi không thành thạo trong việc viết
các chương trình hiệu quả bằng một trong hai ngôn ngữ đó, nên việc so sánh sẽ không công bằng.
Vì C++ là ngôn ngữ nhanh nhất trong bài báo, các so sánh ở đây với C++ là đủ.)
(Cập nhật: Trong bài đăng cập nhật này, chúng tôi sẽ sử dụng snapshot phát triển gần nhất
của compiler Go trên amd64 và phiên bản mới nhất của `g++` -- 4.8.0, được
phát hành vào tháng 3 năm 2013.)

	$ go version
	go version devel +08d20469cc20 Tue Mar 26 08:27:18 2013 +0100 linux/amd64
	$ g++ --version
	g++ (GCC) 4.8.0
	Copyright (C) 2013 Free Software Foundation, Inc.
	...
	$

Các chương trình được chạy trên máy tính với CPU Core i7-2600 3.4GHz và 16 GB
RAM chạy kernel Gentoo Linux 3.8.4-gentoo.
Máy đang chạy với tính năng điều chỉnh tần số CPU bị vô hiệu hóa qua

	$ sudo bash
	# for i in /sys/devices/system/cpu/cpu[0-7]
	do
	    echo performance > $i/cpufreq/scaling_governor
	done
	#

Chúng tôi đã lấy [các chương trình benchmark của Hundt](https://github.com/hundt98847/multi-language-bench)
bằng C++ và Go, kết hợp mỗi cái vào một tệp nguồn duy nhất và loại bỏ tất cả trừ một
dòng đầu ra.
Chúng tôi sẽ tính thời gian chương trình bằng tiện ích `time` của Linux với định dạng hiển thị thời gian người dùng,
thời gian hệ thống, thời gian thực và mức sử dụng bộ nhớ tối đa:

	$ cat xtime
	#!/bin/sh
	/usr/bin/time -f '%Uu %Ss %er %MkB %C' "$@"
	$

	$ make havlak1cc
	g++ -O3 -o havlak1cc havlak1.cc
	$ ./xtime ./havlak1cc
	# of loops: 76002 (total 3800100)
	loop-0, nest: 0, depth: 0
	17.70u 0.05s 17.80r 715472kB ./havlak1cc
	$

	$ make havlak1
	go build havlak1.go
	$ ./xtime ./havlak1
	# of loops: 76000 (including 1 artificial root node)
	25.05u 0.11s 25.20r 1334032kB ./havlak1
	$

Chương trình C++ chạy trong 17.80 giây và sử dụng 700 MB bộ nhớ.
Chương trình Go chạy trong 25.20 giây và sử dụng 1302 MB bộ nhớ.
(Những phép đo này khó đối chiếu với những phép đo trong bài báo, nhưng
mục đích của bài đăng này là khám phá cách sử dụng `go tool pprof`, không phải để tái hiện
các kết quả từ bài báo.)

Để bắt đầu tinh chỉnh chương trình Go, chúng ta phải kích hoạt profiling.
Nếu code sử dụng hỗ trợ benchmark của [package testing Go](/pkg/testing/),
chúng ta có thể sử dụng các cờ `-cpuprofile` và `-memprofile` tiêu chuẩn của gotest.
Trong một chương trình độc lập như thế này, chúng ta phải import `runtime/pprof` và thêm một vài
dòng code:

	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	func main() {
	    flag.Parse()
	    if *cpuprofile != "" {
	        f, err := os.Create(*cpuprofile)
	        if err != nil {
	            log.Fatal(err)
	        }
	        pprof.StartCPUProfile(f)
	        defer pprof.StopCPUProfile()
	    }
	    ...

Code mới định nghĩa một cờ tên `cpuprofile`, gọi
[thư viện cờ Go](/pkg/flag/) để phân tích các cờ dòng lệnh,
và sau đó, nếu cờ `cpuprofile` đã được đặt trên dòng lệnh,
[bắt đầu CPU profiling](/pkg/runtime/pprof/#StartCPUProfile)
chuyển hướng đến tệp đó.
Profiler yêu cầu một lời gọi cuối cùng đến
[`StopCPUProfile`](/pkg/runtime/pprof/#StopCPUProfile) để
xả bất kỳ lần ghi đang chờ xử lý nào vào tệp trước khi chương trình thoát; chúng ta sử dụng `defer`
để đảm bảo điều này xảy ra khi `main` trả về.

Sau khi thêm code đó, chúng ta có thể chạy chương trình với cờ `-cpuprofile` mới
và sau đó chạy `go tool pprof` để phân tích profile.

	$ make havlak1.prof
	./havlak1 -cpuprofile=havlak1.prof
	# of loops: 76000 (including 1 artificial root node)
	$ go tool pprof havlak1 havlak1.prof
	Welcome to pprof!  For help, type 'help'.
	(pprof)

Chương trình `go tool pprof` là một biến thể nhỏ của
[profiler C++ `pprof` của Google](https://github.com/gperftools/gperftools).
Lệnh quan trọng nhất là `topN`, hiển thị `N` mẫu hàng đầu trong profile:

	(pprof) top10
	Total: 2525 samples
	     298  11.8%  11.8%      345  13.7% runtime.mapaccess1_fast64
	     268  10.6%  22.4%     2124  84.1% main.FindLoops
	     251   9.9%  32.4%      451  17.9% scanblock
	     178   7.0%  39.4%      351  13.9% hash_insert
	     131   5.2%  44.6%      158   6.3% sweepspan
	     119   4.7%  49.3%      350  13.9% main.DFS
	      96   3.8%  53.1%       98   3.9% flushptrbuf
	      95   3.8%  56.9%       95   3.8% runtime.aeshash64
	      95   3.8%  60.6%      101   4.0% runtime.settype_flush
	      88   3.5%  64.1%      988  39.1% runtime.mallocgc

Khi CPU profiling được bật, chương trình Go dừng lại khoảng 100 lần mỗi giây
và ghi lại một mẫu bao gồm các bộ đếm chương trình trên stack
của goroutine đang thực thi hiện tại.
Profile có 2525 mẫu, vì vậy nó đã chạy trong hơn 25 giây một chút.
Trong đầu ra `go tool pprof`, có một hàng cho mỗi hàm xuất hiện trong
một mẫu.
Hai cột đầu hiển thị số mẫu mà hàm đang chạy
(trái ngược với việc chờ một hàm được gọi trả về), dưới dạng số đếm thô và dưới dạng phần trăm của tổng mẫu.
Hàm `runtime.mapaccess1_fast64` đang chạy trong 298 mẫu, hay 11.8%.
Đầu ra `top10` được sắp xếp theo số mẫu này.
Cột thứ ba hiển thị tổng đang chạy trong danh sách:
ba hàng đầu chiếm 32.4% các mẫu.
Cột thứ tư và thứ năm hiển thị số mẫu mà hàm xuất hiện
(đang chạy hoặc chờ một hàm được gọi trả về).
Hàm `main.FindLoops` đang chạy trong 10.6% các mẫu, nhưng nó có mặt trên
call stack (nó hoặc các hàm nó gọi đang chạy) trong 84.1% các mẫu.

Để sắp xếp theo cột thứ tư và thứ năm, hãy sử dụng cờ `-cum` (cho cumulative):

	(pprof) top5 -cum
	Total: 2525 samples
	       0   0.0%   0.0%     2144  84.9% gosched0
	       0   0.0%   0.0%     2144  84.9% main.main
	       0   0.0%   0.0%     2144  84.9% runtime.main
	       0   0.0%   0.0%     2124  84.1% main.FindHavlakLoops
	     268  10.6%  10.6%     2124  84.1% main.FindLoops
	(pprof) top5 -cum

Trên thực tế, tổng của `main.FindLoops` và `main.main` lẽ ra phải là 100%, nhưng
mỗi mẫu stack chỉ bao gồm 100 stack frame dưới cùng; trong khoảng một phần tư
các mẫu, hàm `main.DFS` đệ quy sâu hơn 100 frame so với `main.main`
nên trace đầy đủ đã bị cắt ngắn.

Các mẫu stack trace chứa nhiều dữ liệu thú vị hơn về mối quan hệ lời gọi hàm
so với những gì danh sách văn bản có thể hiển thị.
Lệnh `web` viết đồ thị dữ liệu profile dưới định dạng SVG và mở nó trong trình duyệt web.
(Cũng có lệnh `gv` viết PostScript và mở nó trong Ghostview.
Đối với cả hai lệnh, bạn cần [graphviz](http://www.graphviz.org/) được cài đặt.)

	(pprof) web

Một đoạn nhỏ của
[đồ thị đầy đủ](https://rawgit.com/rsc/benchgraffiti/master/havlak/havlak1.svg) trông như thế này:

{{image "pprof/havlak1a-75.png"}}

Mỗi hộp trong đồ thị tương ứng với một hàm duy nhất, và các hộp được định kích thước
theo số mẫu mà hàm đang chạy.
Một cạnh từ hộp X đến hộp Y chỉ ra rằng X gọi Y; con số dọc theo cạnh là
số lần lời gọi đó xuất hiện trong một mẫu.
Nếu một lời gọi xuất hiện nhiều lần trong một mẫu, chẳng hạn như trong các lời gọi hàm đệ quy,
mỗi lần xuất hiện đều tính vào trọng số cạnh.
Điều đó giải thích cho 21342 trên cạnh tự gọi từ `main.DFS` đến chính nó.

Chỉ nhìn qua, chúng ta có thể thấy rằng chương trình dành nhiều thời gian cho các
thao tác hash, tương ứng với việc sử dụng các giá trị `map` của Go.
Chúng ta có thể yêu cầu `web` chỉ sử dụng các mẫu bao gồm một hàm cụ thể, chẳng hạn như
`runtime.mapaccess1_fast64`, để làm sạch một số nhiễu từ đồ thị:

	(pprof) web mapaccess1

{{image "pprof/havlak1-hash_lookup-75.png"}}

Nếu nhìn kỹ, chúng ta có thể thấy rằng các lời gọi đến `runtime.mapaccess1_fast64` đang được
thực hiện bởi `main.FindLoops` và `main.DFS`.

Bây giờ chúng ta đã có ý tưởng sơ bộ về bức tranh toàn cảnh, đã đến lúc phóng to vào một
hàm cụ thể.
Hãy xem `main.DFS` trước, chỉ vì nó là một hàm ngắn hơn:

	(pprof) list DFS
	Total: 2525 samples
	ROUTINE ====================== main.DFS in /home/rsc/g/benchgraffiti/havlak/havlak1.go
	   119    697 Total samples (flat / cumulative)
	     3      3  240: func DFS(currentNode *BasicBlock, nodes []*UnionFindNode, number map[*BasicBlock]int, last []int, current int) int {
	     1      1  241:     nodes[current].Init(currentNode, current)
	     1     37  242:     number[currentNode] = current
	     .      .  243:
	     1      1  244:     lastid := current
	    89     89  245:     for _, target := range currentNode.OutEdges {
	     9    152  246:             if number[target] == unvisited {
	     7    354  247:                     lastid = DFS(target, nodes, number, last, lastid+1)
	     .      .  248:             }
	     .      .  249:     }
	     7     59  250:     last[number[currentNode]] = lastid
	     1      1  251:     return lastid
	(pprof)

Danh sách hiển thị code nguồn của hàm `DFS` (thực ra, cho mọi hàm
khớp với biểu thức chính quy `DFS`).
Ba cột đầu là số mẫu được lấy khi đang chạy dòng đó, số
mẫu được lấy khi đang chạy dòng đó hoặc trong code được gọi từ dòng đó, và
số dòng trong tệp.
Lệnh liên quan `disasm` hiển thị phân rã hợp ngữ của hàm thay vì danh sách nguồn;
khi có đủ mẫu, điều này có thể giúp bạn xem lệnh nào tốn kém.
Lệnh `weblist` kết hợp hai chế độ: nó hiển thị
[danh sách nguồn trong đó nhấp vào một dòng sẽ hiển thị phân rã hợp ngữ](https://rawgit.com/rsc/benchgraffiti/master/havlak/havlak1.html).

Vì chúng ta đã biết rằng thời gian đang dành cho các tra cứu map được triển khai bởi các
hàm runtime hash, chúng ta quan tâm nhất đến cột thứ hai.
Một phần lớn thời gian được dành cho các lời gọi đệ quy đến `DFS` (dòng 247), như sẽ được
mong đợi từ một duyệt đệ quy.
Ngoại trừ đệ quy, có vẻ như thời gian đang dành cho các truy cập vào
map `number` ở các dòng 242, 246 và 250.
Đối với tra cứu cụ thể đó, map không phải là lựa chọn hiệu quả nhất.
Giống như trong một compiler, các cấu trúc basic block có số thứ tự duy nhất
được gán cho chúng.
Thay vì sử dụng `map[*BasicBlock]int` chúng ta có thể sử dụng `[]int`, một slice được lập chỉ mục bởi
số block.
Không có lý do gì để dùng map khi một mảng hoặc slice sẽ làm được.

Việc thay đổi `number` từ map sang slice yêu cầu chỉnh sửa bảy dòng trong chương trình
và cắt giảm thời gian chạy gần hai lần:

	$ make havlak2
	go build havlak2.go
	$ ./xtime ./havlak2
	# of loops: 76000 (including 1 artificial root node)
	16.55u 0.11s 16.69r 1321008kB ./havlak2
	$

(Xem [diff giữa `havlak1` và `havlak2`](https://github.com/rsc/benchgraffiti/commit/58ac27bcac3ffb553c29d0b3fb64745c91c95948))

Chúng ta có thể chạy lại profiler để xác nhận rằng `main.DFS` không còn là một phần
đáng kể của thời gian chạy:

	$ make havlak2.prof
	./havlak2 -cpuprofile=havlak2.prof
	# of loops: 76000 (including 1 artificial root node)
	$ go tool pprof havlak2 havlak2.prof
	Welcome to pprof!  For help, type 'help'.
	(pprof)
	(pprof) top5
	Total: 1652 samples
	     197  11.9%  11.9%      382  23.1% scanblock
	     189  11.4%  23.4%     1549  93.8% main.FindLoops
	     130   7.9%  31.2%      152   9.2% sweepspan
	     104   6.3%  37.5%      896  54.2% runtime.mallocgc
	      98   5.9%  43.5%      100   6.1% flushptrbuf
	(pprof)

Mục nhập `main.DFS` không còn xuất hiện trong profile nữa, và phần còn lại của thời gian
chạy chương trình cũng đã giảm.
Bây giờ chương trình đang dành phần lớn thời gian để cấp phát bộ nhớ và thu gom rác
(`runtime.mallocgc`, vừa cấp phát vừa chạy thu gom rác định kỳ,
chiếm 54.2% thời gian).
Để tìm hiểu tại sao bộ gom rác chạy nhiều như vậy, chúng ta phải tìm hiểu những gì đang
cấp phát bộ nhớ.
Một cách là thêm memory profiling vào chương trình.
Chúng ta sẽ sắp xếp để nếu cờ `-memprofile` được cung cấp, chương trình dừng sau một
lần lặp của việc tìm vòng lặp, ghi một memory profile và thoát:

	var memprofile = flag.String("memprofile", "", "write memory profile to this file")
	...

		FindHavlakLoops(cfgraph, lsgraph)
		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
			    log.Fatal(err)
			}
			pprof.WriteHeapProfile(f)
			f.Close()
			return
		}

Chúng ta gọi chương trình với cờ `-memprofile` để ghi một profile:

	$ make havlak3.mprof
	go build havlak3.go
	./havlak3 -memprofile=havlak3.mprof
	$

(Xem [diff từ havlak2](https://github.com/rsc/benchgraffiti/commit/b78dac106bea1eb3be6bb3ca5dba57c130268232))

Chúng ta sử dụng `go tool pprof` theo cách hoàn toàn giống nhau. Bây giờ các mẫu chúng ta đang xem xét là
các cấp phát bộ nhớ, không phải các tick đồng hồ.

	$ go tool pprof havlak3 havlak3.mprof
	Adjusting heap profiles for 1-in-524288 sampling rate
	Welcome to pprof!  For help, type 'help'.
	(pprof) top5
	Total: 82.4 MB
	    56.3  68.4%  68.4%     56.3  68.4% main.FindLoops
	    17.6  21.3%  89.7%     17.6  21.3% main.(*CFG).CreateNode
	     8.0   9.7%  99.4%     25.6  31.0% main.NewBasicBlockEdge
	     0.5   0.6% 100.0%      0.5   0.6% itab
	     0.0   0.0% 100.0%      0.5   0.6% fmt.init
	(pprof)

Lệnh `go tool pprof` báo cáo rằng `FindLoops` đã cấp phát khoảng
56.3 trong số 82.4 MB đang được sử dụng; `CreateNode` chiếm thêm 17.6 MB.
Để giảm overhead, memory profiler chỉ ghi lại thông tin cho khoảng
một block trên mỗi nửa megabyte được cấp phát ("tốc độ lấy mẫu 1-in-524288"), vì vậy đây
là các xấp xỉ của số đếm thực tế.

Để tìm các cấp phát bộ nhớ, chúng ta có thể liệt kê những hàm đó.

{{raw `
	(pprof) list FindLoops
	Total: 82.4 MB
	ROUTINE ====================== main.FindLoops in /home/rsc/g/benchgraffiti/havlak/havlak3.go
	  56.3   56.3 Total MB (flat / cumulative)
	...
	   1.9    1.9  268:     nonBackPreds := make([]map[int]bool, size)
	   5.8    5.8  269:     backPreds := make([][]int, size)
	     .      .  270:
	   1.9    1.9  271:     number := make([]int, size)
	   1.9    1.9  272:     header := make([]int, size, size)
	   1.9    1.9  273:     types := make([]int, size, size)
	   1.9    1.9  274:     last := make([]int, size, size)
	   1.9    1.9  275:     nodes := make([]*UnionFindNode, size, size)
	     .      .  276:
	     .      .  277:     for i := 0; i < size; i++ {
	   9.5    9.5  278:             nodes[i] = new(UnionFindNode)
	     .      .  279:     }
	...
	     .      .  286:     for i, bb := range cfgraph.Blocks {
	     .      .  287:             number[bb.Name] = unvisited
	  29.5   29.5  288:             nonBackPreds[i] = make(map[int]bool)
	     .      .  289:     }
	...
`}}

Có vẻ như điểm nghẽn hiện tại giống với điểm trước: sử dụng map khi
các cấu trúc dữ liệu đơn giản hơn là đủ.
`FindLoops` đang cấp phát khoảng 29.5 MB maps.

Ngoài ra, nếu chúng ta chạy `go tool pprof` với cờ `--inuse_objects`, nó sẽ
báo cáo số lượng cấp phát thay vì kích thước:

{{raw `
	$ go tool pprof --inuse_objects havlak3 havlak3.mprof
	Adjusting heap profiles for 1-in-524288 sampling rate
	Welcome to pprof!  For help, type 'help'.
	(pprof) list FindLoops
	Total: 1763108 objects
	ROUTINE ====================== main.FindLoops in /home/rsc/g/benchgraffiti/havlak/havlak3.go
	720903 720903 Total objects (flat / cumulative)
	...
	     .      .  277:     for i := 0; i < size; i++ {
	311296 311296  278:             nodes[i] = new(UnionFindNode)
	     .      .  279:     }
	     .      .  280:
	     .      .  281:     // Step a:
	     .      .  282:     //   - initialize all nodes as unvisited.
	     .      .  283:     //   - depth-first traversal and numbering.
	     .      .  284:     //   - unreached BB's are marked as dead.
	     .      .  285:     //
	     .      .  286:     for i, bb := range cfgraph.Blocks {
	     .      .  287:             number[bb.Name] = unvisited
	409600 409600  288:             nonBackPreds[i] = make(map[int]bool)
	     .      .  289:     }
	...
	(pprof)
`}}

Vì khoảng 200.000 map chiếm 29.5 MB, có vẻ như cấp phát map ban đầu
mất khoảng 150 byte.
Điều đó hợp lý khi một map được sử dụng để chứa các cặp key-value, nhưng không hợp lý khi một map
được sử dụng như một proxy cho một tập hợp đơn giản, như ở đây.

Thay vì sử dụng map, chúng ta có thể sử dụng một slice đơn giản để liệt kê các phần tử.
Trong tất cả trừ một trường hợp nơi map đang được sử dụng, thuật toán không thể
chèn một phần tử trùng lặp.
Trong trường hợp còn lại, chúng ta có thể viết một biến thể đơn giản của hàm built-in `append`:

	func appendUnique(a []int, x int) []int {
	    for _, y := range a {
	        if x == y {
	            return a
	        }
	    }
	    return append(a, x)
	}

Ngoài việc viết hàm đó, việc thay đổi chương trình Go để sử dụng slice thay vì
map chỉ yêu cầu thay đổi một vài dòng code.

	$ make havlak4
	go build havlak4.go
	$ ./xtime ./havlak4
	# of loops: 76000 (including 1 artificial root node)
	11.84u 0.08s 11.94r 810416kB ./havlak4
	$

(Xem [diff từ havlak3](https://github.com/rsc/benchgraffiti/commit/245d899f7b1a33b0c8148a4cd147cb3de5228c8a))

Hiện chúng ta đang nhanh hơn 2.11 lần so với lúc bắt đầu. Hãy xem lại CPU profile.

	$ make havlak4.prof
	./havlak4 -cpuprofile=havlak4.prof
	# of loops: 76000 (including 1 artificial root node)
	$ go tool pprof havlak4 havlak4.prof
	Welcome to pprof!  For help, type 'help'.
	(pprof) top10
	Total: 1173 samples
	     205  17.5%  17.5%     1083  92.3% main.FindLoops
	     138  11.8%  29.2%      215  18.3% scanblock
	      88   7.5%  36.7%       96   8.2% sweepspan
	      76   6.5%  43.2%      597  50.9% runtime.mallocgc
	      75   6.4%  49.6%       78   6.6% runtime.settype_flush
	      74   6.3%  55.9%       75   6.4% flushptrbuf
	      64   5.5%  61.4%       64   5.5% runtime.memmove
	      63   5.4%  66.8%      524  44.7% runtime.growslice
	      51   4.3%  71.1%       51   4.3% main.DFS
	      50   4.3%  75.4%      146  12.4% runtime.MCache_Alloc
	(pprof)

Hiện tại cấp phát bộ nhớ và việc thu gom rác hệ quả (`runtime.mallocgc`)
chiếm 50.9% thời gian chạy của chúng ta.
Một cách khác để xem xét tại sao hệ thống đang thu gom rác là xem xét các
cấp phát đang gây ra việc thu gom, những cấp phát dành nhiều thời gian nhất trong
`mallocgc`:

	(pprof) web mallocgc

{{image "pprof/havlak4a-mallocgc.png"}}

Khó để biết những gì đang xảy ra trong đồ thị đó, vì có nhiều node với
số mẫu nhỏ làm che khuất những node lớn.
Chúng ta có thể yêu cầu `go tool pprof` bỏ qua các node không chiếm ít nhất 10% của
các mẫu:

	$ go tool pprof --nodefraction=0.1 havlak4 havlak4.prof
	Welcome to pprof!  For help, type 'help'.
	(pprof) web mallocgc

{{image "pprof/havlak4a-mallocgc-trim.png"}}

Bây giờ chúng ta có thể dễ dàng theo các mũi tên dày để thấy rằng `FindLoops` đang kích hoạt
hầu hết việc thu gom rác.
Nếu chúng ta liệt kê `FindLoops`, chúng ta có thể thấy rằng nhiều trong số đó nằm ngay ở đầu:

{{raw `
	(pprof) list FindLoops
	...
	     .      .  270: func FindLoops(cfgraph *CFG, lsgraph *LSG) {
	     .      .  271:     if cfgraph.Start == nil {
	     .      .  272:             return
	     .      .  273:     }
	     .      .  274:
	     .      .  275:     size := cfgraph.NumNodes()
	     .      .  276:
	     .    145  277:     nonBackPreds := make([][]int, size)
	     .      9  278:     backPreds := make([][]int, size)
	     .      .  279:
	     .      1  280:     number := make([]int, size)
	     .     17  281:     header := make([]int, size, size)
	     .      .  282:     types := make([]int, size, size)
	     .      .  283:     last := make([]int, size, size)
	     .      .  284:     nodes := make([]*UnionFindNode, size, size)
	     .      .  285:
	     .      .  286:     for i := 0; i < size; i++ {
	     2     79  287:             nodes[i] = new(UnionFindNode)
	     .      .  288:     }
	...
	(pprof)
`}}

Mỗi lần `FindLoops` được gọi, nó cấp phát một số cấu trúc bookkeeping có kích thước đáng kể.
Vì benchmark gọi `FindLoops` 50 lần, chúng tích lũy thành một lượng rác đáng kể,
do đó là một lượng công việc đáng kể cho bộ gom rác.

Có một ngôn ngữ được thu gom rác không có nghĩa là bạn có thể bỏ qua các vấn đề cấp phát bộ nhớ.
Trong trường hợp này, một giải pháp đơn giản là giới thiệu một cache để mỗi lời gọi đến `FindLoops`
tái sử dụng bộ nhớ lưu trữ của lần gọi trước khi có thể.
(Thực ra, trong bài báo của Hundt, ông giải thích rằng chương trình Java cần chính xác thay đổi này để
đạt được hiệu suất hợp lý, nhưng ông không thực hiện thay đổi tương tự trong
các triển khai thu gom rác khác.)

Chúng ta sẽ thêm một cấu trúc `cache` toàn cục:

	var cache struct {
	    size int
	    nonBackPreds [][]int
	    backPreds [][]int
	    number []int
	    header []int
	    types []int
	    last []int
	    nodes []*UnionFindNode
	}

và sau đó có `FindLoops` tham chiếu nó như một sự thay thế cho cấp phát:

{{raw `
	if cache.size < size {
	    cache.size = size
	    cache.nonBackPreds = make([][]int, size)
	    cache.backPreds = make([][]int, size)
	    cache.number = make([]int, size)
	    cache.header = make([]int, size)
	    cache.types = make([]int, size)
	    cache.last = make([]int, size)
	    cache.nodes = make([]*UnionFindNode, size)
	    for i := range cache.nodes {
	        cache.nodes[i] = new(UnionFindNode)
	    }
	}

	nonBackPreds := cache.nonBackPreds[:size]
	for i := range nonBackPreds {
	    nonBackPreds[i] = nonBackPreds[i][:0]
	}
	backPreds := cache.backPreds[:size]
	for i := range nonBackPreds {
	    backPreds[i] = backPreds[i][:0]
	}
	number := cache.number[:size]
	header := cache.header[:size]
	types := cache.types[:size]
	last := cache.last[:size]
	nodes := cache.nodes[:size]
`}}

Một biến toàn cục như vậy là thực hành kỹ thuật không tốt, tất nhiên: điều đó có nghĩa là
các lời gọi đồng thời đến `FindLoops` hiện không an toàn.
Hiện tại, chúng tôi đang thực hiện những thay đổi tối thiểu có thể để hiểu điều gì
quan trọng đối với hiệu suất của chương trình; thay đổi này đơn giản và phản ánh
code trong triển khai Java.
Phiên bản cuối cùng của chương trình Go sẽ sử dụng một instance `LoopFinder` riêng biệt để
theo dõi bộ nhớ này, khôi phục khả năng sử dụng đồng thời.

	$ make havlak5
	go build havlak5.go
	$ ./xtime ./havlak5
	# of loops: 76000 (including 1 artificial root node)
	8.03u 0.06s 8.11r 770352kB ./havlak5
	$

(Xem [diff từ havlak4](https://github.com/rsc/benchgraffiti/commit/2d41d6d16286b8146a3f697dd4074deac60d12a4))

Còn nhiều điều chúng ta có thể làm để dọn dẹp chương trình và làm cho nó nhanh hơn, nhưng không có
gì trong số đó yêu cầu các kỹ thuật profiling mà chúng ta chưa hiển thị.
Danh sách công việc được sử dụng trong vòng lặp bên trong có thể được tái sử dụng qua các lần lặp và qua
các lời gọi đến `FindLoops`, và nó có thể được kết hợp với "node pool" riêng biệt được tạo ra
trong lần chạy đó.
Tương tự, bộ nhớ lưu trữ đồ thị vòng lặp có thể được tái sử dụng trên mỗi lần lặp thay vì được cấp phát lại.
Ngoài những thay đổi hiệu suất này,
[phiên bản cuối cùng](https://github.com/rsc/benchgraffiti/blob/master/havlak/havlak6.go)
được viết theo phong cách Go thành thục, sử dụng các cấu trúc dữ liệu và phương thức.
Các thay đổi về phong cách chỉ có tác động nhỏ đến thời gian chạy: thuật toán và
các ràng buộc không thay đổi.

Phiên bản cuối cùng chạy trong 2.29 giây và sử dụng 351 MB bộ nhớ:

	$ make havlak6
	go build havlak6.go
	$ ./xtime ./havlak6
	# of loops: 76000 (including 1 artificial root node)
	2.26u 0.02s 2.29r 360224kB ./havlak6
	$

Đó là nhanh hơn 11 lần so với chương trình chúng ta bắt đầu.
Ngay cả khi chúng ta tắt tái sử dụng đồ thị vòng lặp được tạo ra, để bộ nhớ duy nhất được cache
là bookkeeping tìm vòng lặp, chương trình vẫn chạy nhanh hơn 6.7 lần so với bản gốc
và sử dụng ít bộ nhớ hơn 1.5 lần.

	$ ./xtime ./havlak6 -reuseloopgraph=false
	# of loops: 76000 (including 1 artificial root node)
	3.69u 0.06s 3.76r 797120kB ./havlak6 -reuseloopgraph=false
	$

Tất nhiên, việc so sánh chương trình Go này với chương trình C++ gốc không còn công bằng nữa,
vì chương trình đó sử dụng các cấu trúc dữ liệu không hiệu quả như `set` trong khi `vector` sẽ
phù hợp hơn.
Như một kiểm tra tính hợp lệ, chúng tôi đã dịch chương trình Go cuối cùng sang
[code C++ tương đương](https://github.com/rsc/benchgraffiti/blob/master/havlak/havlak6.cc).
Thời gian thực thi của nó tương tự với chương trình Go:

	$ make havlak6cc
	g++ -O3 -o havlak6cc havlak6.cc
	$ ./xtime ./havlak6cc
	# of loops: 76000 (including 1 artificial root node)
	1.99u 0.19s 2.19r 387936kB ./havlak6cc

Chương trình Go chạy gần nhanh bằng chương trình C++.
Vì chương trình C++ sử dụng delete và cấp phát tự động thay vì một cache tường minh, chương trình C++ ngắn hơn một chút và dễ viết hơn, nhưng không đáng kể:

	$ wc havlak6.cc; wc havlak6.go
	 401 1220 9040 havlak6.cc
	 461 1441 9467 havlak6.go
	$

(Xem [havlak6.cc](https://github.com/rsc/benchgraffiti/blob/master/havlak/havlak6.cc)
và [havlak6.go](https://github.com/rsc/benchgraffiti/blob/master/havlak/havlak6.go))

Các benchmark chỉ tốt bằng các chương trình chúng đo lường.
Chúng tôi đã sử dụng `go tool pprof` để nghiên cứu một chương trình Go không hiệu quả và sau đó cải thiện
hiệu suất của nó theo một bậc độ lớn và giảm mức sử dụng bộ nhớ của nó đi 3.7 lần.
Một so sánh tiếp theo với một chương trình C++ được tối ưu hóa tương đương cho thấy Go có thể cạnh tranh
với C++ khi các lập trình viên cẩn thận về lượng rác được tạo ra
bởi các vòng lặp bên trong.

Các nguồn chương trình, nhị phân Linux x86-64 và các profile được sử dụng để viết bài đăng này
có sẵn trong [dự án benchgraffiti trên GitHub](https://github.com/rsc/benchgraffiti/).

Như đã đề cập ở trên, [`go test`](/cmd/go/#Test_packages) đã bao gồm
các cờ profiling này: định nghĩa một
[hàm benchmark](/pkg/testing/) và bạn đã sẵn sàng.
Ngoài ra còn có một giao diện HTTP tiêu chuẩn cho dữ liệu profiling. Trong một HTTP server, thêm

	import _ "net/http/pprof"

sẽ cài đặt các handler cho một vài URL dưới `/debug/pprof/`.
Sau đó bạn có thể chạy `go tool pprof` với một đối số duy nhất, đó là URL đến dữ liệu profiling của máy chủ của bạn, và nó sẽ tải xuống và kiểm tra một profile trực tiếp.

	go tool pprof http://localhost:6060/debug/pprof/profile   # 30-second CPU profile
	go tool pprof http://localhost:6060/debug/pprof/heap      # heap profile
	go tool pprof http://localhost:6060/debug/pprof/block     # goroutine blocking profile

Profile chặn goroutine sẽ được giải thích trong một bài đăng tương lai. Hãy đón chờ.
