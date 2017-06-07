### 如何使用

把config.toml.temp复制并另命名为config.toml,修改其中的对应的sqlite,mysql,postgres环境配置，没有则不修改

进入当前目录执行**go test -v .**

    cp config.toml.temp config.toml
    go test -v .

