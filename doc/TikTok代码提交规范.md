# TikTok代码提交规范

### 获取仓库代码

```sh
git clone + 仓库链接
例：
git clone git@gitee.com:AloserL/tiktok.git
```

### 关于git的使用细则

```sh
第一次使用需要配置ssh，因为http经常性的崩坏，ssh的具体配置方法自行百度/问大佬。
配置名字和邮箱
例：
git config --global user.name AloserL
git config --global user.email xxx@xx.com
远程仓库链接(origin)设置
一般git clone下来的代码会自动设置origin，如果git remote -v显示的结果无仓库链接或者原为http的链接想修改的则使用如下代码
例：git remote add origin git@github.com:路由.git
```



### 修改代码

在修改代码前 要提交的分支下 建立一个新的分支，在新的分支改完代码后才能提交

新分支的命名方式：姓名(拼音)/原分支名/修改内容

如我要在master分支创建一个login接口，我的做法应该是

```sh
git checkout master # 切换原来的master分治
git pull   #拉取master分支在远程的代码
git checkout -b AloserL/master/add-login#在现在分支的基础上，创建新的分支
git push origin AloserL/master/add-login # 在远程仓库建立这个分支
git branch --set-upstream-to=origin/AloserL/master/add-login AloserL/master/add-login #本地分支远程分支做关联
```

### 代码提交

#### 提交代码

```sh
git add (文件1 文件2) || git add .(添加所有文件)
git commit -m '修改了什么'
git push
```

#### 如果自己的分支落后了远程的主分支

```sh
git checkout master
git pull   注：此操作会覆盖本地工作区
git checkout AloserL/master/add-login 
git push
```





##### ......如有其它疑问可以先自行百度，再询问大佬
