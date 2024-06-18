# 【单例模式】心得

## 这个模式核心是什么？其实讲的是什么？

确保全局只有一个实例，方便资源共享。

## 给我什么启发？学到了什么？

1. 通过单例模式可以减少内存占用，减少资源占用。
2. 在只需要一个实例的时候，就可以考虑使用单例模式。

## 我以前写过的什么代码场景也许可以用这个重构？大概思路？

1. 数据库连接，日志记录
2. 用一个变量来记录是否已经创建过数据库连接、日志记录器，如果创建过，就直接返回，否则就创建实例