// 插件, 不是立刻就需要的功能
// 1. 上传, 粘贴图片2
// 2. 笔记信息
// 3. 历史记录
// 4. 附件
requirejs.config({
	paths: {
		// life
		// 'editor_drop_paste': 'js/plugins/editor_drop_paste', // 改为使用tinymce里的插件，md编辑器也先禁用此功能
		'attachment_upload': 'js/plugins/attachment_upload',

		// 'jquery.ui.widget': 'js/plugins/fileupload/jquery.ui.widget',
		// 'iframe-transport': 'js/plugins/fileupload/jquery.iframe-transport',
		'fileupload': 'js/plugins/fileupload.min',

		'note_info': 'js/plugins/note_info',
		'tips': 'js/plugins/tips',
		'history': 'js/plugins/history',
	},
	shim: {
		// life
		// 'fileupload': {deps: ['jquery.ui.widget', 'iframe-transport']},
	}
});

// 异步加载
setTimeout(function () {
	// 小异步
	// require(["editor_drop_paste"]);
	require(["attachment_upload"]);

	require(['note_info']);

	// 大异步
	setTimeout(function () {
		require(['tips']);
		require(['history']);
	}, 10);
});