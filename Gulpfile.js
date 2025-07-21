var gulp = require('gulp');
var clean = require('gulp-clean');
var uglify = require('gulp-uglify');
var rename = require('gulp-rename');
var minifyHtml = require("gulp-minify-html");
var minifycss = require('gulp-minify-css');
var concat = require('gulp-concat');
var replace = require('gulp-replace');
var inject = require('gulp-inject');
var gulpSequence = require('gulp-sequence');
var footer = require('gulp-footer');
var fs = require('fs');
var path = require('path');

var leanoteBase = './';
var base = leanoteBase + '/public'; // public base
var noteDev = leanoteBase + '/app/views/note/note-dev.html';
var noteProBase = leanoteBase + '/app/views/note';
var messagesPath = leanoteBase + 'messages';

// gulp-uglify 会并发-异步地压缩每一个 src 文件，谁先压完就谁先往下游传递。
// 所以一旦在 uglify() 之后再前使用 concat() 插件，输出可能不会按原本的 src([...]) 顺序来。
// 解决方式是先 concat() 再 uglify()，虽慢但稳。
// 先合并，再压缩就不怕打乱文件顺序，造成函数执行异常。

// gulp3里task是全并行执行的，所以这里需要特别注意任务的执行顺序依赖问题；
// gulp3任务的父子依赖关系，只由 task() 函数区分。并且由它保证先子再父的task执行顺序；

// / 合并requirejs和markdown为一个文件
gulp.task('concatMarkdownJsV2', function() {
    var jss = [
        'libs/require.js',
        'md/main-v2.js',
    ];

    for(var i in jss) {
        jss[i] = base + '/' + jss[i];
    }

    return gulp.src(jss)
        .pipe(concat('markdown-v2.min.js'))
        .pipe(uglify())
        .pipe(gulp.dest(base + '/js'));
});

// mincss
gulp.task('minifycss', function() {
    gulp.src(base + '/libs/bootstrap/bootstrap.css')
        .pipe(rename({suffix: '.min'}))
        .pipe(minifycss())
        .pipe(gulp.dest(base + '/libs/bootstrap'));

    gulp.src(base + '/fonts/font-awesome-4.2.0/css/font-awesome.css')
        .pipe(rename({suffix: '-min'}))
        .pipe(minifycss())
        .pipe(gulp.dest(base + '/fonts/font-awesome-4.2.0/css'));

    gulp.src(base + '/css/zTreeStyle/zTreeStyle.css')
        .pipe(rename({suffix: '-min'}))
        .pipe(minifycss())
        .pipe(gulp.dest(base + '/css/zTreeStyle'));

    gulp.src(base + '/md/themes/default.css')
        .pipe(rename({suffix: '-min'}))
        .pipe(minifycss())
        .pipe(gulp.dest(base + '/md/themes'));

    gulp.src(base + '/js/contextmenu/css/contextmenu.css')
        .pipe(rename({suffix: '-min'}))
        .pipe(minifycss())
        .pipe(gulp.dest(base + '/js/contextmenu/css'));

    gulp.src(base + '/album/css/style.css')
        .pipe(rename({suffix: '-min'}))
        .pipe(minifycss())
        .pipe(gulp.dest(base + '/album/css'));

    // theme TODO
    var as = ['default', 'simple', 'writting', 'writting-overwrite', 'mobile'];
    /*
    for(var i = 0; i < as.length; ++i) {
        gulp.src(base + '/css/theme/' + as[i] + '.css')
            .pipe(minifycss())
            .pipe(gulp.dest(base + '/css/theme'));
    }
    */
});

gulp.task('fileupload', function() {
	var jss = [
        'jquery.iframe-transport.js',
        'jquery.ui.widget.js',
		'jquery.fileupload.js'
    ];

    for(var i in jss) {
        jss[i] = base + '/js/plugins/fileupload/' + jss[i];
    }

    return gulp.src(jss)
        .pipe(concat('fileupload.min.js'))
        .pipe(footer("if (window.define) {define('fileupload', [], function() {});}"))
        .pipe(uglify())
        .pipe(gulp.dest(base + '/js/plugins'));
});

// plugins压缩
gulp.task('plugins', ['fileupload'], function() {
    // 所有js合并成一个
     var jss = [
        'note_info',
        'tips',
        'history',
        'attachment_upload',
        'editor_drop_paste',
        'main'
    ];

    for(var i in jss) {
        jss[i] = base + '/js/plugins/' + jss[i] + '.js';
    }
    jss.push(base + '/js/plugins/fileupload.min.js');

    return gulp.src(jss)
        .pipe(concat('plugins.min.js'))
        .pipe(uglify())
        .pipe(gulp.dest(base + '/js/plugins'));
});

// 只获取需要js i18n的key
gulp.task('i18n', function() {
    var keys = {};
    var reg = /getMsg\((['"])((?:(?!\1).|\\\1)*?)\1/g;
    var reg2 = /msg: ?"?([0-9a-zA-Z]*)"?/g;
    function getKey(data) {
        while ((ret = reg.exec(data)) !== null) {
            keys[ret[2].replace(/\\(['"])/g, '$1')] = 1;
        }

        while(ret2 = reg2.exec(data)) {
            keys[ret2[1]] = 1;
        }
    }
    // 先获取需要的key
    function ls(ff) { 
        var files = fs.readdirSync(ff);  
        for(fn in files) {  
            var fname = ff + path.sep + files[fn];  
            var stat = fs.lstatSync(fname);  
            if(stat.isDirectory() == true) {
                ls(fname);
            } 
            else {
                if ((fname.indexOf('.html') > 0 || fname.indexOf('.js') > 0)) {
                    // console.log(fname);
                    // if (fname.indexOf('min.js') < 0) {
                        var data = fs.readFileSync(fname, "utf-8");
                        // 得到getMsg里的key
                        getKey(data);
                    // }
                }
            }  
        }  
    }

    ls(base + '/admin');
    ls(base + '/blog');
    ls(base + '/md');
    ls(base + '/js');
    ls(base + '/album');
    ls(base + '/libs');
    ls(base + '/member');
    ls(base + '/tinymce');

    ls(leanoteBase + '/app/views');

    console.log('used keys parsed OK');
    var langs = {}; // zh-cn: 1

    // msg.zh
    function getAllMsgs(fname) {
        var msg = {};

        var data = fs.readFileSync(fname, "utf-8");
        var lines = data.split('\n');
        for (var i = 0; i < lines.length; ++i) {
            var line = lines[i];
            // 忽略注释
            if (line[0] == '#' || line[1] == '#') {
                continue;
            }
            var lineArr = line.split('=');
            if (lineArr.length >= 2) {
               var key = lineArr[0];
               lineArr.shift();
               msg[key] = lineArr.join('=');
               // msg[lineArr[0]] = lineArr[1];
            }
        }
        return msg;
    }

    // 得到所有的语言的后缀
    // 返回{en-us: 1, }
    function getAllLangs() {
        var langs = {};
        var files = fs.readdirSync(messagesPath);  
        for(fn in files) {
            var fname = files[fn]; 
            if (fname.indexOf('-') > 0) {
                langs[fname] = 1;
            }
        }
        return langs;
    }

    // msg.zh, msg.js
    function genI18nJsFile(targetFilename, lang, fromFilenames, keys) {
        var msgs = {};
        fromFilenames.forEach(function (name) {
            var tmpMsgs = getAllMsgs(leanoteBase + '/messages/' + lang + '/' + name + '.conf');
            for (var i in tmpMsgs) {
                msgs[i] = tmpMsgs[i];
            }
        });

        var toMsgs = {};
        for (var i in msgs) {
            // 只要需要的
            if (i in keys) {
                toMsgs[i] = msgs[i];
            }
        }
        var str = 'var MSG=' + JSON.stringify(toMsgs) + ';';
        str += 'function getMsg(key, data) {' +
            'var msg = MSG[key];' +
            'if(msg) {' +
                'if(data) {' +
                    'if(!isArray(data)) {data = [data];}' + 
                    'for(var i = 0; i < data.length; ++i) {' + 
                        'msg = msg.replace("%s", data[i]);' + 
                    '}' + 
                '}' + 
                'return msg;' + 
            '}' + 
             'return key;' + 
        '}';

        // 写入到文件中
        var toFilename = targetFilename + '.' + lang + '.js';
        fs.writeFileSync(base + '/js/i18n/' + toFilename, str);
    }

    function genTinymceLang(lang) {
        var msgs = getAllMsgs(leanoteBase + 'messages/' + lang + '/tinymce_editor.conf');
        var str = 'tinymce.addI18n("' + lang + '",' + JSON.stringify(msgs) + ');';
        fs.writeFileSync(base + '/tinymce/langs/' + lang + '.js', str);
    }

    var langs = getAllLangs();
    for (var lang in langs) {
        genI18nJsFile('blog', lang, ['blog'], keys);
        genI18nJsFile('msg', lang, ['msg', 'member', 'markdown', 'album'], keys);

        genTinymceLang(lang);
    }
    
});

// 合并album需要的js
gulp.task('concatAlbumJs', function() {
    var jss = [
        //'libs/jquery/jquery.min.js',       // 这些基础库在之前已经打包过了，
        //'libs/bootstrap/bootstrap.min.js', // 不单独使用album的话不用重复再打包
        //'js/plugins/fileupload.min.js',
        'libs/jquery/jquery.pagination.js',
        'album/js/main.js',
    ];

    for(var i in jss) {
        jss[i] = base + '/' + jss[i];
    }

    return gulp.src(jss)
        .pipe(concat('album.all.js'))
        .pipe(uglify())
        .pipe(gulp.dest(base + '/album/js'));
});

gulp.task('object_id', function() {
    return gulp.src(base + '/js/object_id.js')
        .pipe(uglify())
        .pipe(rename({suffix: '-min'}))
        .pipe(gulp.dest(base + '/js/'));
});

// 合并Js, 依据：note-dev.html里的<!-- dev -->代码块
gulp.task('concatDepJs', ['object_id'], function() {
    var jss = [
        'libs/jquery/jquery.min.js',
        'libs/jquery/jquery.ztree.all-3.5-min.js',
        'js/object_id-min.js',
    ];

    for(var i in jss) {
        jss[i] = base + '/' + jss[i];
    }

    return gulp.src(jss)
        // .pipe(uglify()) // 压缩
        .pipe(concat('dep.min.js'))
        .pipe(gulp.dest(base + '/js'));
});

gulp.task('contextmenu', function() {
    return gulp.src(base + '/js/contextmenu/jquery.contextmenu.js')
        .pipe(uglify())
        .pipe(rename({suffix: '-min'}))
        .pipe(gulp.dest(base + '/js/contextmenu/'));
});

// 合并Js, 依据：note-dev.html里的<!-- dev -->代码块
gulp.task('concatAppJs', ['contextmenu'], function() {
    var jss = [
        'libs/jquery/jQuery-slimScroll-1.3.0/jquery.slimscroll-min.js',
        'js/contextmenu/jquery.contextmenu-min.js',
        'js/common.js',
        'libs/bootstrap/bootstrap.min.js',
        'js/app/note.js',
        'js/app/tag.js',
        'js/app/notebook.js',
        'js/app/share.js',
        'js/app/page.js',
    ];

    for(var i in jss) {
        jss[i] = base + '/' + jss[i];
    }

    return gulp.src(jss)
        .pipe(concat('app.min.js'))
		.pipe(uglify())
        .pipe(gulp.dest(base + '/js'));
});

// note-dev.html -> note.html, 替换css, js
gulp.task('devToProHtml', function() {
    return gulp
        .src(noteDev)
        .pipe(replace(/<!-- dev -->[.\s\S]+?<!-- \/dev -->/g, '')) // 把dev 去掉
        .pipe(replace(/<!-- pro_dep_js -->/, '<script src="/js/dep.min.js"></script>')) // 替换
        .pipe(replace(/<!-- pro_app_js -->/, '<script src="/js/app.min.js"></script>')) // 替换
        .pipe(replace(/<!-- pro_markdown_js -->/, '<script src="/js/markdown-v2.min.js"></script>')) // 替换
        .pipe(replace('/tinymce/tinymce.js', '/tinymce/tinymce.full.min.js')) // 替换
        .pipe(replace('/js/contextmenu/css/contextmenu.css', '/js/contextmenu/css/contextmenu-min.css'))
        .pipe(replace(/<!-- pro_tinymce_init_js -->/, "var tinyMCEPreInit = {base: '/public/tinymce', suffix: '.min'};")) // 替换
        .pipe(replace(/plugins\/main.js/, "plugins/plugins.min.js")) // 替换
        // 连续两个空行换成一个空行
        .pipe(replace(/\r\n\r\n/g, '\r\n'))
        .pipe(replace(/\r\n\r\n/g, '\r\n'))
        .pipe(replace(/\r\n\r\n/g, '\r\n'))
        .pipe(replace(/\r\n\r\n/g, '\r\n'))
        .pipe(replace(/\r\n\r\n/g, '\r\n'))
        .pipe(replace(/\r\n\r\n/g, '\r\n'))
        .pipe(replace('console.log(o);', ''))
        .pipe(replace('console.trace(o);', ''))
        // .pipe(minifyHtml()) // 不行, 压缩后golang报错
        .pipe(rename('note.html'))
        .pipe(gulp.dest(noteProBase));
});

gulp.task('concat', ['concatMarkdownJsV2', 'concatDepJs', 'concatAppJs']);
gulp.task('html', ['devToProHtml']);
gulp.task('default', ['minifycss', 'plugins', 'i18n', 'concatAlbumJs', 'concat', 'html']);
