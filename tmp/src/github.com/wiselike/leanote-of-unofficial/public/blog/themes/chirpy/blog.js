$(function () {
	// 初始化tooltip
	$('[data-toggle="tooltip"]').tooltip();

	// back-to-top
	$(window).scroll(() => {
		if ($(this).scrollTop() > 50 && $("#sidebar-trigger").css("display") === "none") {
			$("#back-to-top").fadeIn();
		} else {
			$("#back-to-top").fadeOut();
		}
	});

	$("#back-to-top").click(() => {
		$("body,html").animate({ scrollTop: 0 }, 800);
		return false;
	});

	// search
	const input = $("#search-input");
	const btnClear = $("#search-cleaner");
	const searchWrapper = $("#search-wrapper");
	if (input.val() !== "") {
		if (!btnClear.hasClass("visible")) { btnClear.addClass("visible"); }
	}
	input.on("keyup", function (e) {
		if (input.val() !== "") {
			if (!btnClear.hasClass("visible")) { btnClear.addClass("visible"); }
		}
	});
	btnClear.on("click", function () {
		input.val("");
		input.focus(); searchWrapper.addClass("input-focus");
		btnClear.removeClass("visible");
	});

	// Expand or close the sidebar in mobile screens.
	const sidebarUtil = (function () {
		const ATTR_DISPLAY = "sidebar-display";
		let isExpanded = false;
		const body = $("body");
		return {
			toggle() {
				if (isExpanded === false) {
					body.attr(ATTR_DISPLAY, "");
				} else {
					body.removeAttr(ATTR_DISPLAY);
				}
				isExpanded = !isExpanded;
			}
		};
	}());
	$("#sidebar-trigger").click(sidebarUtil.toggle);
	$("#mask").click(sidebarUtil.toggle);

	//check box
	/* hide browser default checkbox */
	$("input[type=checkbox]").addClass("unloaded");
	/* create checked checkbox */
	$("input[type=checkbox][checked]").before("<i class=\"fas fa-check-circle checked\"></i>");
	/* create normal checkbox */
	$("input[type=checkbox]:not([checked])").before("<i class=\"far fa-circle\"></i>");
});

function blogPostInit() {
	const btnSelector = '.code-header>button';
	const ATTR_TIMEOUT = 'timeout';
	const TIMEOUT = 2000; // in milliseconds

	// Toc初始化
	var navSelector = "#toc";
	Toc.init($(navSelector)
		// $nav: $(navSelector)
		// $scope: $("h1, h2, h3, h4, h5, h6"),
	);
	$("body").scrollspy({
		target: $(navSelector)
	});

	function isLocked(node) {
		if ($(node)[0].hasAttribute(ATTR_TIMEOUT)) {
			let timeout = $(node).attr(ATTR_TIMEOUT);
			if (Number(timeout) > Date.now()) {
				return true;
			}
		}
		return false;
	}

	function lock(node) {
		$(node).attr(ATTR_TIMEOUT, Date.now() + TIMEOUT);
	}

	function unlock(node) {
		$(node).removeAttr(ATTR_TIMEOUT);
	}

	// clipboard 相关
	/* --- Copy code block --- */
	// Initial the clipboard.js object
	const clipboard = new ClipboardJS(btnSelector, {
		target(trigger) {
			return trigger.parentNode.nextElementSibling; // clipboard button 的父节点的下一个节点就是代码块
		}
	});

	$(btnSelector).tooltip({
		trigger: 'hover',
		placement: 'left'
	});

	clipboard.on('success', (e) => {
		e.clearSelection();

		let trigger = $(e.trigger);
		let ICON_DEFAULT = trigger.children().attr('class'); // 获取原先的Icon
		if (isLocked(trigger)) {
			return;
		}

		trigger.children().attr('class', 'fas fa-check'); // set Success Icon
		trigger.attr('data-original-title', "已复制").tooltip('show'); // show tooltip
		lock(trigger);

		setTimeout(() => {
			trigger.tooltip('hide').removeAttr('data-original-title'); // hide tooltip
			trigger.children().attr('class', ICON_DEFAULT); // resume Icon
			unlock(trigger);
		}, TIMEOUT);

	});

	/* --- Copy post share link --- */
	var clipboard_link = new ClipboardJS('.clipboard');
	clipboard_link.on("success", function (e) {
		let target = $(e.trigger); //clipboard.js 添加已复制提示效果 https://www.veidc.com/21341.html
		if (isLocked(target)) {
			return;
		}
		// Switch tooltip title
		const defaultTitle = target.attr('data-original-title');
		target.attr("data-original-title", "已复制").tooltip('show');
		lock(target);
		setTimeout(() => {
			target.attr('data-original-title', defaultTitle).tooltip(defaultTitle.length > 0 ? 'show' : 'hide');
			unlock(target);
		}, 2000); // 2000 ms
		e.clearSelection();
	});
	clipboard_link.on("error", function (e) {
		alert("复制出错");
	});

	/* Mermaid */
	let initTheme = "default";
	if ($("html[mode=dark]").length > 0
		|| ($("html[mode]").length == 0 && window.matchMedia("(prefers-color-scheme: dark)").matches)
	) {
		initTheme = "dark";
	}

	// Mermaid Markdown converts to HTML
	$("pre").has("code.language-mermaid").each(function () {
		let svgCode = $(this).children().html(); // 执行这里的时候还没有添加行号等信息
		// var svgCode = "";
		// $(this).children("ol").children("li").each(function () { //放到 prttyInit 后面执行，已经有行号了
		//   svgCode = svgCode + $(this).text() + "\n";
		// });
		$(this).addClass("unloaded");
		$(this).after(`<div class=\"mermaid\">${svgCode}</div>`);
	});

	let mermaidConf = {
		theme: initTheme  /* <default|dark|forest|neutral> */
	};
	mermaid.initialize(mermaidConf);

	/* ChartJs */
	$("pre").has("code.language-chart").each(function () {
		let svgCode = $(this).children().html();
		let jsonObject = JSON.parse(svgCode);
		let containerElt = document.createElement('canvas');
		let ctx = containerElt.getContext('2d'); // canvas.getContext 测试浏览器是否支持 canvas，返回CanvasRenderingContext2D对象
		new Chart(ctx, jsonObject);
		$(this).replaceWith(containerElt); // $(selectorA).replaceWitj(dom); $(this).after(containerElt); $(this).remove();
	});

	/* Magnific Popup */
	$("img").each(function () {
		$(this).attr("href", $(this).attr("src"));
	});
	$("img").magnificPopup({type:'image'});
}
