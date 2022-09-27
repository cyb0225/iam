# golang 新增知识

### errors
`官方包及第三方包`

Is 
As
Wrap (fmt %w)
Unwrap

### errGroup
Group 结构体内部封装了 context 和 waitGroup， 用于在一个协程返回错误时，将group的其余协程也关闭。
