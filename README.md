个人主页右上角三个点
<img src="https://github.com/chengshihard/douyin-video-download/blob/main/images/IMG_5052.PNG?raw=true" width="375">
点击复制链接
<img src="https://github.com/chengshihard/douyin-video-download/blob/main/images/QQ%E6%88%AA%E5%9B%BE20210319185702.png?raw=true" width="375">
复制完粘贴到浏览器 F12打开调试 复制有数据列表的链接
<img src="https://github.com/chengshihard/douyin-video-download/blob/main/images/QQ%E6%88%AA%E5%9B%BE20210319185444.png?raw=true" width="375">

例如 https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=MS4wLjABAAAAKbhcZVC80794_3SNxHLzZEglkO5sDKRbZpN0TcbO2GA&count=21&max_cursor=0&aid=1128&_signature=SWImtgAAKTmhuATLVJI21EliJq&dytk=  
把链接中的max_cursor=0 修改成{$max_cursor}  
修改后链接是 https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=MS4wLjABAAAAKbhcZVC80794_3SNxHLzZEglkO5sDKRbZpN0TcbO2GA&count=21&max_cursor={$max_cursor}&aid=1128&_signature=SWImtgAAKTmhuATLVJI21EliJq&dytk=  

<img src="https://github.com/chengshihard/douyin-video-download/blob/main/images/QQ%E6%88%AA%E5%9B%BE20210319190950.png?raw=true" width="375">
path改成你本地保存视频的文件夹路径     
startTime是13位的时间戳，比如你要采集从2020-01-01 00:00:00开始的视频，就把2020-01-01 00:00:00转换成13位的时间戳填上即可    
requestUrl是上面修改过的url    

需要golang环境  
go run main.go   
等待视频下载完毕  
