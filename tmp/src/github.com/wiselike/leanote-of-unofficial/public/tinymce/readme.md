TinyMCE For Leanote
==========================================

本仓库来自[Tinymce](https://github.com/tinymce/tinymce), 在此基础上作了修改. 同样都基于LGPL协议.

## 分支

* master 作为Leanote web版编辑器
* desktop-app 作为Leanote桌面版编辑器

## 常用命令

* grunt minify

打包成一个
* grunt bundle --themes leanote --plugins autolink,link,leaui_image,leaui_mindmap,lists,hr,paste,searchreplace,leanote_nav,leanote_code,tabfocus,table,directionality,textcolor

## 注意

先要删除tinymce.dev.js, tinymce.js, tinymce.min.js 不然 
grunt amdlc:core不会生成这几个文件!!!

可能是因为没有变化, 所以不再生成!

------------------------------------------

Building TinyMCE
-----------------
1. Install Node.js
2. Open a console and go to the project directory
3. Write "npm i -g grunt-cli" to install the grunt command line tool globally.
4. Write "npm i" to install all package dependencies.
4. Build TinyMCE by writing "grunt"

Build tasks
------------
`grunt`
Lints, minified, unit tests and creates release packages for TinyMCE.

`grunt minify`
Minifies all JS and CSS files.

`grunt test`
Runs all qunit tests on PhantomJS.

`grunt lint`
Runs all source files though various JS linters.

`grunt sc-test`
Runs all qunit tests on Saucelabs.

`grunt watch`
Watches for source code changes and triggers rebuilds and linting.

`grunt --help`
Displays the various build tasks.

Bundle themes and plugins into a single file
---------------------------------------------
`grunt bundle --themes modern --plugins table,paste`
Minifies the core, adds the modern theme and adds the table and paste plugin into tinymce.min.js.

Contributing to the TinyMCE project
------------------------------------
You can read more about how to contribute to this project at [http://www.tinymce.com/develop/contributing.php](http://www.tinymce.com/develop/contributing.php)

[![Build Status](https://travis-ci.org/tinymce/tinymce.png?branch=master)](https://travis-ci.org/tinymce/tinymce)
