# gorm指南

### 模型定义

gorm通过将go结构体 映射到数据库来简化数据库交互。

##### 定义模型

模型是通过普通结构体定义的，这些结构体可以包含具有基本go类型，指针或这些类型的别名，甚至是自定义类型（只需要实现database/sql包中的Scanner和Valuer接口）

- 示例

  ```go
  type User strcut {
      ID uint //
      Name straing
  	Email
  }
  ```

  