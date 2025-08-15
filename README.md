# Leanote 非官方版本
* 基于官方二次开发维护
* 实时同步官方最新补丁，并同时合入新特性功能
* 必要时，将新特性功能推送到官方
* 包含部分实验性功能，并始终保持向前兼容
* 完整的列出所有变更记录
* [提供docker下的一键部署实施方法](https://github.com/wiselike/leanote-of-unofficial-nolicensed/wiki/docker-deploy-method-docker一键部署方法--Linux)
* [以docker镜像的方式，提供完整的开发环境，欢迎共同来开发](https://github.com/wiselike/leanote-of-unofficial-nolicensed/wiki/how-to-build-in-docker-docker编译环境搭建方法--Linux)

# Leanote of non-official version
* for secondary development only
* update with official patches
* push to mainline as needed
* with experimental features
* with full ChangeLogs
* [provide deploy method by Docker Container](https://github.com/wiselike/leanote-of-unofficial-nolicensed/wiki/docker-deploy-method-docker一键部署方法--Linux)
* [provide full develop environment by Docker Container](https://github.com/wiselike/leanote-of-unofficial-nolicensed/wiki/how-to-build-in-docker-docker编译环境搭建方法--Linux)

# 注意事项
* 代码合并请求，必须一次commit一个独立完整功能，请不要随意PR。
* 拒绝一次PR合入多个特性功能或者故障修复。
* 若代码无法review，只能拒绝合入，请配合，谢谢。

# ChangeLogs
1. searched from https://github.com/leanote/leanote  
		with git(d58fd64)[gen tmp tool without revel] on 15 Aug 2021
2. patched https://github.com/ctaoist/leanote/commit/2cee584f793e21c7469e8701874d1548bee1be17
		which comes from https://github.com/leanote/leanote/compare/c4bb20fd129e63edd14bc7ecd229bbad3b13bcb7..450deb09bdf1ebc47ea31b0ed209b8d85492f7fa
		and https://github.com/leanote/leanote/pull/933/commits/92db56f4f141e477dbd1fa01232ea2c6536fe027  
3. patched https://github.com/ctaoist/leanote/commit/c5c19e32e0cb892fe35178a14dfe927049f5b3a9
4. patched https://github.com/ctaoist/leanote/commit/c2c4a5536301132a78594c2311d1dbd0d957b304
5. 自研的优化
6. patched "markdown编辑器增加字数统计功能" https://github.com/ctaoist/leanote/commit/297ca0c3ef15db680a7fe395b0283497dd768b2d and https://github.com/ctaoist/leanote/commit/7060829c7ab015431d05a529c4f2d31822992f15
7. 自研：修改配置文件，改默认语言为中文
8. 自研：添加自定义的git忽略文件
9. 自研：整理node图片，按标题来存放，以便于到服务器上检索维护
10. 自研：修复Site's URL设置后，却不同步配置文件，导致重启后会失效的问题
11. 自研：添加在配置文件中自定义note的图片、附件存放路径
12. 自研：修改默认note历史数为5，并且添加app.conf配置文件可配。优化历史记录新增删除算法。修改note历史顺序，与官方原生不兼容，如使用，会自动删除之前的旧历史，无其他副影响
13. 自研：将所有配置参数，调整为从系统全局变量中读取，而不是每次都从文件中读。优化了读取速度和效率
14. 自研：使用gofmt格式化所有go代码，不对源码做任何手动改动
15. 自研：禁用github.io，改为使用本地css文件
16. 自研：禁用demo账号，自己用的话demo没有必要存在啊，直接用admin不就行啦
17. 自研：修复无法退出登录的故障
18. 自研：修正保存note历史记录的算法，调整note自动保存到历史记录的功能，用起来更顺畅
19. **上传原始package.json文件里定义的项目GPLv2 license**
20. 自研：前端实现博客置顶设置
21. 优化note的字数统计功能
22. 自研：修复移动端界面的博客图标显示异常
23. 自研：改进验证码登录流程，降低爆破的可能性
24. 自研：添加图片备份文件夹，防止图片丢失
25. 自研：屏蔽首页的广告页，改为直接跳转为登录或者note页
26. 自研：清理数据库中冗余数据，将chirpy主题(非自研)合入为默认主题之一
27. 自研：修复发送邮件的中文标题乱码故障
28. 自研：在个人中心->账户信息->Email栏目增加用户邮箱地址修改的功能；修复邮件发送错误提示故障
29. 自研: 用户名允许长度放宽为2位
30. 自研：允许admin用户名的修改，并实时更新到app.conf配置文件，且无需重启服务
31. 自研：整理node附件，按标题来存放，以便于直接到服务器上检索维护
32. 自研：将防止图片丢失的图片备份文件夹backup-origins按用户user-id和“年”来创建文件夹分隔
33. 自研：“历史记录功能”一系列调整：  
		1. 改进单个文章的“历史记录功能”：加宽显示列，可美观显示10个以上历史记录;  
		2. 修正记录算法，仅记录历史，不再记录当前页，并且不会再丢历史记录  
		3. 优化历史记录数据库存取算法  
34. 自研：新增删除单条历史记录的按钮，用户可手动删除历史记录了
35. 自研：修复文章移动/复制的问题。任意子目录下的文章，想移动或复制时，对应栏目没有变成灰色、鼠标不可点选。这会导致，全局查找某文章后，就不知道这篇文章是在哪个子笔记本下的了。现在已修复。
36. 自研：修复在写作模式下，置顶按钮显示错误的bug。并添加置顶按钮的blog页面跳转功能(同blog按钮一样，鼠标点击后跳转到blog页面)。
36. 自研：[修复图片、附件时，会同时插入域名到笔记中的故障](https://github.com/wiselike/leanote-of-unofficial/pull/3)，这会导致后续没法修改网站域名
37. 自研：修复“未编辑笔记，笔记的更新时间却被刷新了”的故障。
38. 核心：新增**一键容器部署方法**、**自编译环境搭建**方法
39. 自研：修复手机端“公开博客、博客置顶”按钮的重叠显示问题；修复手机端笔记编辑按钮点不动的问题
40. 自研：切换笔记视图时，同时切换更新笔记本视图
41. 自研：修复服务端后端整理笔记的图片和附件后，图片和附件的编号不连续的问题
42. 自研：重构图片、附件在服务端存储算法，详见 [#6](https://github.com/wiselike/leanote-of-unofficial/issues/6) [#7](https://github.com/wiselike/leanote-of-unofficial/issues/7)  
		1. 相册控件无法看到也无法编辑默认相册了，默认相册变为自行管控，无法人为干预。  
		2. 在相册控件里相册里的图片，只做图片存储使用。一旦插入到笔记中，就会自动复制一份出来。图片原始文件会保存到文章标题的文件夹下，imageID则会记录到默认相册。  
		3. 从笔记删除图片或不引用图片，则文章标题的文件夹下的图片原始文件会自动被清理，默认相册里的imageID也会自动被清理；  
		4. 从相册控件里删除相册里的图片，也不会影响到笔记里的图片显示。  
		5. 附件按笔记标题存放，复制笔记会同步复制附件，删除笔记也会同步删除附件。  
43. 自研：在Makefile中，新增github-release的发布方法
44. 自研：使用[simple-formater](https://github.com/wiselike/simple-formater)格式化所有js/css/html代码，不对源码做任何手动改动
44. 自研：更新第三方revel组件版本至v1.3.0
44. 自研：更新js组件版本：jquery-1.9.0->3.6.0，require-2.1.14->2.3.7，bootstrap-3.0.2->3.4.1，ace-1.2.3->1.2.9
45. 自研：修复由于网络等其他原因导致当前编写的笔记实际保存失败时，仍然页面上仍然提示“保存成功”的问题。
46. 自研：修复中文tag的翻译问题，并解决按标签查找笔记的功能。
47. 自研：修复调整笔记本视图的多个问题：  
		1. 修复笔记本视图的宽度调整不生效，刷新页面后又会被自动重置的问题；  
		2. 修复连续多次鼠标右键点击后不同笔记本后，笔记本上的阴影不消散的问题；  
		3. 修复鼠标悬停到之前已激活的笔记本时，阴影反而变亮的情况；  
		4. 修复2、3级菜单的缩进方式，避免笔记本变更菜单级数的拖动时，缩进不能也跟着及时调整的问题；  
		5. 修复“废纸篓”里的笔记，还会显示笔记本路径链的问题；  
48. 自研：新增笔记本视图的多个特性：  
		1. 调整笔记本的右键菜单，放开右键笔记本展开限制，右键可以看到本notebook下的子notebook了；  
		2. 增加可以正常选择复制或移动到的目标笔记本的子笔记本里去；  
		3. 给当前笔记本的路径链，增加下划线的提示。用户可以快速知晓，当前笔记的笔记本所在路径链回溯；  
49. 自研：调整tinymce的gulp，确保每次格式化和生成代码，自动创建出来的文件的md5值不变。这有助于git的文件版本管理。
50. 自研：增加可以为每个笔记本单独设置排序和视图显示。设置排序和视图显示时候，粒度变为每个笔记本独立进行控制，而不会影响到其他所有笔记本。









