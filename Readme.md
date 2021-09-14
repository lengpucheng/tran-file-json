<center><h1>Tran-File-Json(TFJ)</h1></center>

TFJ可以将指定文件夹或文件编码成一个文本文件😊将一个编码的文本文件还原成一个文件夹或者文件🤭可用于仅能传输文本的场景传输其他类型文件

# 本地使用
```shell
./tfj -o <cmd> -f <file> -t <targer>
```
参数如下：
+ `cmd`:执行的命令，有`save`和`load`两种
+ `file`:文件路径，**必填**保存或者加载的文件路径，需要保存或加载的文件路径
+ `targer`:目标路径，保存或加载后的目标文件路径，若为保存则会将指定目录文件编码成一个文本文件保存到指定的目录，若为加载则会将-f指定的文件加载并还原到-t指定的目录

若未windows机器则命令为
```shell
.\tfj.exe -o <cmd> -f <file> -t <targer>
```
# 设置为系统命令
linux下将文件拷贝至`tfj /usr/bin/`
```shell
cp tfj /usr/bin/
```
windows下设置path环境变量
```shell
PATH=%PATH%;<tfj.exe路径>
```
设置完毕后可以直接使用如下方式执行（windows和linux相同）
```shell
tfj -o <cmd> -f <file> -t <targer>
```