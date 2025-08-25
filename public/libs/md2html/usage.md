/**
	* Markdown convert to html
	*
	<script src="md2html.min.js"></script>
	<script>
		md2Html(document.getElementById('md').value, $('#content'));
		md2Html('hello, **leanote**', document.getElementById('content2'), function allRendered(html) {
			alert(html);
		});
	</script>
*
* @author leanote.com
* @date 2015/04/11
*/