学习笔记

Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

Q: sql.ErrNoRows不应该Wrap到上层，应为dao作为一个封装库，其中一个功能就是屏蔽底层依赖，而不应该将底层的sql信息暴露给上层。而需要依据业务需求，返回nil对象，rows为0或者新建dao.ErrNoRows错误。