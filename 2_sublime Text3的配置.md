Sublime Text
这里将介绍Sublime Text 3（以下简称Sublime）+ GoSublime + gocode的组合，那么为什么选择这个组合呢？

自动化提示代码,如下图所示

<img src='https://github.com/KenNaNa/build-web-application-with-golang/raw/master/zh/images/1.4.sublime1.png?raw=true'/>

图1.5 sublime自动化提示界面

保存的时候自动格式化代码，让您编写的代码更加美观，符合Go的标准。

支持项目管理
<img src='https://github.com/KenNaNa/build-web-application-with-golang/raw/master/zh/images/1.4.sublime2.png?raw=true'/>


图1.6 sublime项目管理界面

支持语法高亮

Sublime Text 3可免费使用，只是保存次数达到一定数量之后就会提示是否购买，点击取消继续用，和正式注册版本没有任何区别。

接下来就开始讲如何安装，下载 Sublime

根据自己相应的系统下载相应的版本，然后打开Sublime，对于不熟悉Sublime的同学可以先看一下这篇文章Sublime Text 全程指南或者sublime text3入门教程

打开之后安装 Package Control：Ctrl+` 打开命令行，执行如下代码：
适用于 Sublime Text 3：

import  urllib.request,os;pf='Package Control.sublime-package';ipp=sublime.installed_packages_path();urllib.request.install_opener(urllib.request.build_opener(urllib.request.ProxyHandler()));open(os.path.join(ipp,pf),'wb').write(urllib.request.urlopen('http://sublime.wbond.net/'+pf.replace(' ','%20')).read())
适用于 Sublime Text 2：

import  urllib2,os;pf='Package Control.sublime-package';ipp=sublime.installed_packages_path();os.makedirs(ipp)ifnotos.path.exists(ipp)elseNone;urllib2.install_opener(urllib2.build_opener(urllib2.ProxyHandler()));open(os.path.join(ipp,pf),'wb').write(urllib2.urlopen('http://sublime.wbond.net/'+pf.replace(' ','%20')).read());print('Please restart Sublime Text to finish installation')
这个时候重启一下Sublime，可以发现在在菜单栏多了一个如下的栏目，说明Package Control已经安装成功了。

<img src="https://github.com/KenNaNa/build-web-application-with-golang/raw/master/zh/images/1.4.sublime3.png?raw=true"/>


图1.7 sublime包管理
安装完之后就可以安装Sublime的插件了。需安装GoSublime、SidebarEnhancements和Go Build，安装插件之后记得重启Sublime生效，Ctrl+Shift+p打开Package Controll 输入pcip（即“Package Control: Install Package”的缩写）。
这个时候看左下角显示正在读取包数据，完成之后出现如下界面
<img src='https://github.com/KenNaNa/build-web-application-with-golang/raw/master/zh/images/1.4.sublime4.png?raw=true'/>


图1.8 sublime安装插件界面
这个时候输入GoSublime，按确定就开始安装了。同理应用于SidebarEnhancements和Go Build。

安装 gocode

go get -u github.com/nsf/gocode

gocode 将会安装在默认`$GOBIN`
另外建议安装gotests(生成测试代码):

先在sublime安装gotests插件,再运行:
go get -u -v github.com/cweill/gotests/...
验证是否安装成功，你可以打开Sublime，打开main.go，看看语法是不是高亮了，输入import是不是自动化提示了，import "fmt"之后，输入fmt.是不是自动化提示有函数了。
如果已经出现这个提示，那说明你已经安装完成了，并且完成了自动提示。

如果没有出现这样的提示，一般就是你的$PATH没有配置正确。你可以打开终端，输入gocode，是不是能够正确运行，如果不行就说明$PATH没有配置正确。 (针对XP)有时候在终端能运行成功,但sublime无提示或者编译解码错误,请安装sublime text3和convert utf8插件试一试

MacOS下已经设置了$GOROOT, $GOPATH, $GOBIN，还是没有自动提示怎么办。
请在sublime中使用command + 9， 然后输入env检查$PATH, GOROOT, $GOPATH, $GOBIN等变量， 如果没有请采用下面的方法。

首先建立下面的连接， 然后从Terminal中直接启动sublime

ln -s /Applications/Sublime\ Text\ 2.app/Contents/SharedSupport/bin/subl /usr/local/bin/sublime
