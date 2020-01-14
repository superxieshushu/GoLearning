- Go 语言25个关键字
```
    break default func interface select
    case defer go map struct
    chan else goto package switch
    const fallthrough if range type
    continue for import return var
```

- 变量的生命周期：如果将指向短生命周期对象的指针保存到具有长生命周期的对象中，
  特别是保存到全局变量时，会阻止对短生命周期对象的垃圾回收（从而可能影响程序的性
  能）。
  
- 位操作的有什么使用场景？