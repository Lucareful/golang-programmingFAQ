# channel

## 无缓冲 chan
 
- 无缓冲的意思就是channel只能保存一个数据，当往无缓冲的 channel 写入第2个数据的时候协程会被阻塞，直到channel中的第1个数据被取走，才会唤醒被阻塞的协程。

- :heavy_exclamation_mark:使用无缓冲 chan 需要读和写同时准备好不然就会出现死锁 deadlock! 

```go
	ch := make(chan int)
	go func() {
		{
			ch <- 0
		}
	}()
	go func() {
		{
			ch <- 1
		}
	}()
	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()
```

## 带缓冲 chan

- 缓冲channel的使用方式跟无缓冲channel一样，区别就是往channel写输入数据的时候，如果缓存队列还没满，是不会阻塞写操作。

```go
ch := make(chan int, 100)
```

例如：上面创建了的channel缓冲队列大小是100，如果写入到channel中，还未被取走的数据大于100，就会阻塞写操作。