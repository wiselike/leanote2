(function() {

	(function(){
	  if (!window.console) window.console = {};
	  var ks=['log','info','warn','error']; for (var i=0;i<ks.length;i++) if(!console[ks[i]]) console[ks[i]]=function(){};
	  if (!Function.prototype.bind) Function.prototype.bind = function (o){ var f=this,s=[].slice,p=s.call(arguments,1),N=function(){},
		b=function(){ return f.apply(this instanceof N? this:o, p.concat(s.call(arguments))); }; if (f.prototype){ N.prototype=f.prototype; b.prototype=new N(); } return b; };
	  if (!Object.assign) Object.assign = function(t){ if(t==null) throw new TypeError; var to=Object(t);
		for(var i=1;i<arguments.length;i++){ var s=arguments[i]; if(s!=null) for(var k in s) if(Object.prototype.hasOwnProperty.call(s,k)) to[k]=s[k]; } return to; };
	  if (window.NodeList && !NodeList.prototype.forEach) NodeList.prototype.forEach = Array.prototype.forEach;
	})();

	// Create the converter and the editor
	var converter = new Markdown.Converter();
	var options = {
		_DoItalicsAndBold: function(text) {
			// Restore original markdown implementation
			text = text.replace(/(\*\*|__)(?=\S)(.+?[*_]*)(?=\S)\1/g,
				"<strong>$2</strong>");
			text = text.replace(/(\*|_)(?=\S)(.+?)(?=\S)\1/g,
				"<em>$2</em>");
			return text;
		}
	};
	converter.setOptions(options);

	function loadJs(src, callback) {
		var _doc = document.getElementsByTagName('head')[0];
		var script = document.createElement('script');
		script.setAttribute('type', 'text/javascript');
		script.setAttribute('src', src);
		_doc.appendChild(script);
		script.onload = script.onreadystatechange = function() {
			if(!this.readyState || this.readyState=='loaded' || this.readyState=='complete'){
				callback && callback();
			}
			script.onload = script.onreadystatechange = null;
		}
	}

	function _each(list, callback) {
		if(list && list.length > 0) {
			for(var i = 0; i < list.length; i++) {
				callback(list[i]);
			}
		}
	}
	function _has(obj, key) {
		return hasOwnProperty.call(obj, key);
	};

	// markdown extra
	function initMarkdownExtra() {
		// Create the converter and the editor
		// var converter = new Markdown.Converter();
		var options = {
			_DoItalicsAndBold: function(text) {
				// Restore original markdown implementation
				text = text.replace(/(\*\*|__)(?=\S)(.+?[*_]*)(?=\S)\1/g,
					"<strong>$2</strong>");
				text = text.replace(/(\*|_)(?=\S)(.+?)(?=\S)\1/g,
					"<em>$2</em>");
				return text;
			}
		};
		converter.setOptions(options);

		//================
		// markdown exstra

		var markdownExtra = {};
		markdownExtra.config = {
			extensions: [
				"fenced_code_gfm",
				"tables",
				"def_list",
				"attr_list",
				"footnotes",
				"smartypants",
				"strikethrough",
				"newlines"
			],
			intraword: true,
			comments: true,
			highlighter: "highlight"
		};
		var extraOptions = {
			extensions: markdownExtra.config.extensions,
			highlighter: "prettify"
		};

		if(markdownExtra.config.intraword === true) {
			var converterOptions = {
				_DoItalicsAndBold: function(text) {
					text = text.replace(/([^\w*]|^)(\*\*|__)(?=\S)(.+?[*_]*)(?=\S)\2(?=[^\w*]|$)/g, "$1<strong>$3</strong>");
					text = text.replace(/([^\w*]|^)(\*|_)(?=\S)(.+?)(?=\S)\2(?=[^\w*]|$)/g, "$1<em>$3</em>");
					// Redo bold to handle _**word**_
					text = text.replace(/([^\w*]|^)(\*\*|__)(?=\S)(.+?[*_]*)(?=\S)\2(?=[^\w*]|$)/g, "$1<strong>$3</strong>");
					return text;
				}
			};
			converter.setOptions(converterOptions);
		}

		if(markdownExtra.config.comments === true) {
			converter.hooks.chain("postConversion", function(text) {
				return text.replace(/<!--.*?-->/g, function(wholeMatch) {
					return wholeMatch.replace(/^<!---(.+?)-?-->$/, ' <span class="comment label label-danger">$1</span> ');
				});
			});
		}

		Markdown.Extra.init(converter, extraOptions);
	}

	//==============
	// toc start

	function initToc() {
		var toc = {};
		toc.config = {
			marker: "\\[(TOC|toc)\\]",
			maxDepth: 6,
			button: true,
		};

		// TOC element description
		function TocElement(tagName, anchor, text) {
			this.tagName = tagName;
			this.anchor = anchor;
			this.text = text;
			this.children = [];
		}
		TocElement.prototype.childrenToString = function() {
			if(this.children.length === 0) {
				return "";
			}
			var result = "<ul>\n";
			_each(this.children, function(child) {
				result += child.toString();
			});
			result += "</ul>\n";
			return result;
		};
		TocElement.prototype.toString = function() {
			var result = "<li>";
			if(this.anchor && this.text) {
				result += '<a href="#' + this.anchor + '">' + this.text + '</a>';
			}
			result += this.childrenToString() + "</li>\n";
			return result;
		};

		// Transform flat list of TocElement into a tree
		function groupTags(array, level) {
			level = level || 1;
			var tagName = "H" + level;
			var result = [];

			var currentElement;
			function pushCurrentElement() {
				if(currentElement !== undefined) {
					if(currentElement.children.length > 0) {
						currentElement.children = groupTags(currentElement.children, level + 1);
					}
					result.push(currentElement);
				}
			}

			_each(array, function(element) {
				if(element.tagName != tagName) {
					if(level !== toc.config.maxDepth) {
						if(currentElement === undefined) {
							currentElement = new TocElement();
						}
						currentElement.children.push(element);
					}
				}
				else {
					pushCurrentElement();
					currentElement = element;
				}
			});
			pushCurrentElement();
			return result;
		}

		var utils = {};
		var nonWordChars = new RegExp('[^\\p{L}\\p{N}-]', 'g');
		utils.slugify = function(text) {
			return text.toLowerCase().replace(/\s/g, '-') // Replace spaces with -
				.replace(nonWordChars, '') // Remove all non-word chars
				.replace(/\-\-+/g, '-') // Replace multiple - with single -
				.replace(/^-+/, '') // Trim - from start of text
				.replace(/-+$/, ''); // Trim - from end of text
			};

			// Build the TOC
			var previewContentsElt;
			function buildToc(previewContentsElt) {
				var anchorList = {};
				function createAnchor(element) {
					var id = element.id || utils.slugify(element.textContent) || 'title';
					var anchor = id;
					var index = 0;
					while (_has(anchorList, anchor)) {
						anchor = id + "-" + (++index);
					}
					anchorList[anchor] = true;
					// Update the id of the element
					element.id = anchor;
					return anchor;
				}

				var elementList = [];
				_each(previewContentsElt.querySelectorAll('h1, h2, h3, h4, h5, h6'), function(elt) {
					elementList.push(new TocElement(elt.tagName, createAnchor(elt), elt.textContent));
				});
				elementList = groupTags(elementList);
				return '<div class="toc">\n<ul>\n' + elementList.join("") + '</ul>\n</div>\n';
			}

			toc.convert = function(previewContentsElt) {
				var tocExp = new RegExp("^\\s*" + toc.config.marker + "\\s*$");
				var tocEltList = document.querySelectorAll('.table-of-contents, .toc');
				var htmlToc = buildToc(previewContentsElt);
				// Replace toc paragraphs
				_each(previewContentsElt.getElementsByTagName('p'), function(elt) {
					if(tocExp.test(elt.innerHTML)) {
						elt.innerHTML = htmlToc;
					}
				});
				// Add toc in the TOC button
				_each(tocEltList, function(elt) {
					elt.innerHTML = htmlToc;
				});
			}

			return toc;
	}

	//===========
	// mathjax
	function initMathJax() {
		// 配置
		MathJax.Hub.Config({
			skipStartupTypeset: true,
			"HTML-CSS": {
				preferredFont: "TeX",
				availableFonts: [
					"STIX",
					"TeX"
				],
				linebreaks: {
					automatic: true
				},
				EqnChunk: 10,
				imageFont: null
			},
			tex2jax: { inlineMath: [["$","$"],["\\\\(","\\\\)"]], displayMath: [["$$","$$"],["\\[","\\]"]], processEscapes: true },
			TeX: {
				noUndefined: {
					attributes: {
						mathcolor: "red",
						mathbackground: "#FFEEEE",
						mathsize: "90%"
					}
				},
				Safe: {
					allow: {
						URLs: "safe",
						classes: "safe",
						cssIDs: "safe",
						styles: "safe",
						fontsize: "all"
					}
				}
			},
			messageStyle: "none"
		});

		var mathJax = {};
		mathJax.config = {
			tex    : "{}",
			tex2jax: '{ inlineMath: [["$","$"],["\\\\\\\\(","\\\\\\\\)"]], displayMath: [["$$","$$"],["\\\\[","\\\\]"]], processEscapes: true }'
		};

		mathJax.init = function(p) {
			converter.hooks.chain("preConversion", removeMath);
			converter.hooks.chain("postConversion", replaceMath);
		};

		// From math.stackexchange.com...

		//
		//  The math is in blocks i through j, so
		//    collect it into one block and clear the others.
		//  Replace &, <, and > by named entities.
		//  For IE, put <br> at the ends of comments since IE removes \n.
		//  Clear the current math positions and store the index of the
		//    math, then push the math string onto the storage array.
		//
		function processMath(i, j, unescape) {
			var block = blocks.slice(i, j + 1).join("")
				.replace(/&/g, "&amp;")
				.replace(/</g, "&lt;")
				.replace(/>/g, "&gt;");
			for(HUB.Browser.isMSIE && (block = block.replace(/(%[^\n]*)\n/g, "$1<br/>\n")); j > i;)
				blocks[j] = "", j--;
			blocks[i] = "@@" + math.length + "@@";
			unescape && (block = unescape(block));
			math.push(block);
			start = end = last = null;
		}

		function removeMath(text) {
			if(!text) {
				return;
			}
			start = end = last = null;
			math = [];
			var unescape;
			if(/`/.test(text)) {
				text = text.replace(/~/g, "~T").replace(/(^|[^\\])(`+)([^\n]*?[^`\n])\2(?!`)/gm, function(text) {
					return text.replace(/\$/g, "~D")
				});
				unescape = function(text) {
					return text.replace(/~([TD])/g,
						function(match, n) {
							return {T: "~", D: "$"}[n]
						})
					};
				} else {
					unescape = function(text) {
						return text
					};
				}

				//
				//  The pattern for math delimiters and special symbols
				//    needed for searching for math in the page.
				//
				var splitDelimiter = /(\$\$?|\\(?:begin|end)\{[a-z]*\*?\}|\\[\\{}$]|[{}]|(?:\n\s*)+|@@\d+@@)/i;
				var split;

				if(3 === "aba".split(/(b)/).length) {
					split = function(text, delimiter) {
						return text.split(delimiter)
					};
				} else {
					split = function(text, delimiter) {
						var b = [], c;
						if(!delimiter.global) {
							c = delimiter.toString();
							var d = "";
							c = c.replace(/^\/(.*)\/([im]*)$/, function(a, c, b) {
								d = b;
								return c
							});
							delimiter = RegExp(c, d + "g")
						}
						for(var e = delimiter.lastIndex = 0; c = delimiter.exec(text);) {
							b.push(text.substring(e, c.index));
							b.push.apply(b, c.slice(1));
							e = c.index + c[0].length;
						}
						b.push(text.substring(e));
						return b
					};
				}

				blocks = split(text.replace(/\r\n?/g, "\n"), splitDelimiter);
				for(var i = 1, m = blocks.length; i < m; i += 2) {
					var block = blocks[i];
					if("@" === block.charAt(0)) {
						//
						//  Things that look like our math markers will get
						//  stored and then retrieved along with the math.
						//
						blocks[i] = "@@" + math.length + "@@";
						math.push(block)
					} else if(start) {
						// Ignore inline maths that are actually multiline (fixes #136)
						if(end == inline && block.charAt(0) == '\n') {
							if(last) {
								i = last;
								processMath(start, i, unescape);
							}
							start = end = last = null;
							braces = 0;
						}
						//
						//  If we are in math, look for the end delimiter,
						//    but don't go past double line breaks, and
						//    and balance braces within the math.
						//
						else if(block === end) {
							if(braces) {
								last = i
							} else {
								processMath(start, i, unescape)
							}
						} else {
							if(block.match(/\n.*\n/)) {
								if(last) {
									i = last;
									processMath(start, i, unescape);
								}
								start = end = last = null;
								braces = 0;
							} else {
								if("{" === block) {
									braces++
								} else {
									"}" === block && braces && braces--
								}
							}
						}
					} else {
						if(block === inline || "$$" === block) {
							start = i;
							end = block;
							braces = 0;
						} else {
							if("begin" === block.substr(1, 5)) {
								start = i;
								end = "\\end" + block.substr(6);
								braces = 0;
							}
						}
					}

				}
				last && processMath(start, last, unescape);
				return unescape(blocks.join(""));
			}

				//
				//  Put back the math strings that were saved,
				//    and clear the math array (no need to keep it around).
				//
				function replaceMath(text) {
					text = text.replace(/@@(\d+)@@/g, function(match, n) {
						return math[n]
					});
					math = null;
					return text
				}

				//
				//  This is run to restart MathJax after it has finished
				//    the previous run (that may have been canceled)
				//
				function startMJ(toElem, callback) {
					var preview = toElem;
					pending = false;
					HUB.cancelTypeset = false;
					HUB.Queue([
						"Typeset",
						HUB,
						preview
					]);
					// 执行完后, 再执行
					HUB.Queue(function() {
						callback && callback();
					});
				}

				var ready = false, pending = false, preview = null, inline = "$", blocks, start, end, last, braces, math, HUB = MathJax.Hub;

				//
				//  Runs after initial typeset
				//
				HUB.Queue(function() {
					ready = true;
					HUB.processUpdateTime = 50;
					HUB.Config({"HTML-CSS": {EqnChunk: 10, EqnChunkFactor: 1}, SVG: {EqnChunk: 10, EqnChunkFactor: 1}})
				});

			mathJax.init();
		return {
			convert: startMJ
		}
	}

	function initUml() {
		//===========
		// uml
		var umlDiagrams = {};
		umlDiagrams.config = {
			flowchartOptions: [
				'{',
				'   "line-width": 2,',
				'   "font-family": "sans-serif",',
				'   "font-weight": "normal"',
				'}'
			].join('\n')
		};

		var _loadUmlJs = false;

		// callback 执行完后执行
		umlDiagrams.convert = function(target, callback) {
			var previewContentsElt = target;

			var sequenceElems = previewContentsElt.querySelectorAll('.prettyprint > .language-sequence');
			var flowElems = previewContentsElt.querySelectorAll('.prettyprint > .language-flow');

			function convert() {
				_each(sequenceElems, function(elt) {
					try {
						var diagram = Diagram.parse(elt.textContent);
						var preElt = elt.parentNode;
						var containerElt = crel('div', {
							class: 'sequence-diagram'
						});
						preElt.parentNode.replaceChild(containerElt, preElt);
						diagram.drawSVG(containerElt, {
							theme: 'simple'
						});
					}
					catch(e) {
						console.trace(e);
					}
				});
				_each(flowElems, function(elt) {
					try {

						var chart = flowchart.parse(elt.textContent);
						var preElt = elt.parentNode;
						var containerElt = crel('div', {
							class: 'flow-chart'
						});
						preElt.parentNode.replaceChild(containerElt, preElt);
						chart.drawSVG(containerElt, JSON.parse(umlDiagrams.config.flowchartOptions));
					}
					catch(e) {
						console.error(e);
					}
				});

				callback && callback();
			}

			if(sequenceElems.length > 0 || flowElems.length > 0) {
				if(!_loadUmlJs) {
					loadJs('/public/libs/md2html/uml.js', function() {
						_loadUmlJs = true;
						convert();
					});
				} else {
					convert();
				}
			} else {
				callback && callback();
			}
		};

		return umlDiagrams;
	}

	// ===== Mermaid（ES5 + 单一路径，wkhtml 友好） =====

	function normalizeMermaidCodeBlocks(root){
	  var sel='.prettyprint > .language-mermaid, pre > code.language-mermaid, code.mermaid';
	  var codeNodes=root.querySelectorAll(sel);
	  var i;
	  for(i=0;i<codeNodes.length;i++){
	    var node=codeNodes[i];
	    var wrap=(node.parentNode && node.parentNode.nodeName==='PRE') ? node.parentNode : node;
	    var div=document.createElement('div'); div.className='mermaid';
	    div.textContent=(node.textContent||'').replace(/^\s+|\s+$/g,'');
	    if (wrap.parentNode) wrap.parentNode.replaceChild(div, wrap);
	  }
	}

	function hasMermaidAPI(){
	  return (window.mermaid && window.mermaid.mermaidAPI && typeof window.mermaid.mermaidAPI.render==='function') ||
	         (window.mermaidAPI && typeof window.mermaidAPI.render==='function') ||
	         (window.mermaid && typeof window.mermaid.render==='function');
	}

	function ensureMermaid(cb){
	  if (hasMermaidAPI()) { cb && cb(); return; }
	  loadJs('public/libs/md2html/mermaid-7.1.2.min.js', function(){
	    // 旧 QtWebKit 上偶尔 onload 先触发，再挂载到 window：轮询等待
	    var t=0, id=setInterval(function(){
	      if (hasMermaidAPI() || ++t>40) { clearInterval(id); cb && cb(); }
	    }, 50);
	  });
	}

	function renderMermaidOnce(root, done){
	  var all=root.querySelectorAll('div.mermaid'); var targets=[], i;
	  for(i=0;i<all.length;i++){
	    var el=all[i];
	    if (el.getAttribute('data-processed')==='true') continue;
	    if (el.querySelector('svg')) { el.setAttribute('data-processed','true'); continue; }
	    if (!el.getAttribute('data-graph')) {
	      el.setAttribute('data-graph', (el.textContent||'').replace(/^\s+|\s+$/g,''));
	    }
	    targets.push(el);
	  }
	  if (!targets.length) { done && done(); return; }

	  if (window.mermaid && typeof window.mermaid.initialize==='function') {
	    try { window.mermaid.initialize({ startOnLoad:false, securityLevel:'strict' }); } catch(e){}
	  }

	  var ctx=null, renderFn=null;
	  if (window.mermaid && window.mermaid.mermaidAPI && window.mermaid.mermaidAPI.render) {
	    ctx=window.mermaid.mermaidAPI; renderFn=ctx.render;
	  } else if (window.mermaidAPI && window.mermaidAPI.render) {
	    ctx=window.mermaidAPI; renderFn=ctx.render;
	  } else if (window.mermaid && window.mermaid.render) {
	    ctx=window.mermaid; renderFn=ctx.render;
	  } else {
	    done && done(); return;
	  }

	  var doneCnt=0, total=targets.length;
	  for(i=0;i<total;i++){
	    (function(el, idx){
	      var id='mmd-'+idx+'-'+(+new Date());
	      var code=el.getAttribute('data-graph')||'';
	      if (!/^(graph|flowchart|sequenceDiagram|classDiagram|stateDiagram|erDiagram|gantt|journey|pie)\b/.test(code)) {
	        el.setAttribute('data-processed','true'); mark(); return;
	      }
	      try {
	        renderFn.call(ctx, id, code, function(svg){
	          el.innerHTML=svg;
	          el.setAttribute('data-processed','true');
	          mark();
	        }, el);
	      } catch(e){
	        el.setAttribute('data-processed','true');
	        mark();
	      }
	    })(targets[i], i);
	  }
	  function mark(){ if(++doneCnt>=total){ done && done(); } }
	}

	function fixMermaidSvgSizeForWkhtml(root){
	  var svgs = (root||document).querySelectorAll('.mermaid > svg');
	  var i;
	  for (i = 0; i < svgs.length; i++) {
	    var svg = svgs[i];
	    var vb = svg.getAttribute('viewBox'); // "0 0 w h"
	    if (!vb) continue;
	    var m = vb.match(/^\s*0\s+0\s+([\d.]+)\s+([\d.]+)\s*$/);
	    if (!m) continue;
	    var w = m[1], h = m[2];
	    svg.setAttribute('width',  Math.min(+w, 780) + 'px');
	    svg.setAttribute('height', Math.min(+h, 19000) + 'px');
	    svg.style.maxWidth = 'none';
	  }
	}

	// extra是实时的, 同步进行
	initMarkdownExtra();

	var m;
	window.md2Html = function(mdText, toElem, callback) {
		var _umlEnd = false;
		var _mathJaxEnd = false;
		var _mermaidEnd = true; // 默认 true；检测到 mermaid 再置为 false

		function _tryFinish() {
			if (_umlEnd && _mathJaxEnd && _mermaidEnd) {
				// Mermaid 尺寸修正（若有）
				try { fixMermaidSvgSizeForWkhtml(toElem); } catch(e){}
				callback && callback(toElem.innerHTML);
				try { window.status='done'; } catch(e){}
			}
		}

		// 如果是jQuery对象
		if(!toElem['querySelectorAll'] && toElem['get']) {
			toElem = toElem.get(0);
		}
		function _go(mdText, toElem) {
			var htmlParsed = converter.makeHtml(mdText);
			toElem.innerHTML = htmlParsed;

			// 同步执行
			var toc = initToc();
			toc.convert(toElem);

			// 规范化 Mermaid 代码块
			normalizeMermaidCodeBlocks(toElem);

			// 异步执行：UML
			var umlDiagrams = initUml();
			umlDiagrams.convert(toElem, function() {
				_umlEnd = true;
				_tryFinish();
			});

			// 异步执行：Mermaid（若存在）
			if (toElem.querySelector('.mermaid')) {
				_mermaidEnd = false;
				ensureMermaid(function(){
					renderMermaidOnce(toElem, function(){
						_mermaidEnd = true;
						_tryFinish();
					});
				});
			}

			_tryFinish();
		}

		// 表示有mathjax?
		// 加载mathJax
		if(mdText.indexOf('$') !== -1) {
			loadJs("/public/libs/MathJax/MathJax.js?config=TeX-AMS_HTML", function() {
				if(!m) {
					m = initMathJax(); // 注意：不要再用 var m，避免遮蔽
				}
				// 放到后面, 不然removeMathJax()不运行, bug
				_go(mdText, toElem);
				m.convert(toElem, function() {
					_mathJaxEnd = true;
					_tryFinish();
				});
			});
		} else {
			_mathJaxEnd = true;
			_go(mdText, toElem);
		}
	}

})();
