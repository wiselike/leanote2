(function () {
    'use strict';
    var script = window.document.querySelector("script:last-child");
    var path = script.src.split('?')[0]; // remove any ?query
    var pluginDir = path.split('/').slice(0, -1).join('/');

    tinymce.PluginManager.add('ace', function (editor, url) {
        var $this = this;
        var $ = editor.$;

        editor.contentCSS.push(pluginDir + '/' + "style.css");

        $this.escapePreHtml = function (content) {
            return content.replace(/</g, '&lt;').replace(/>/g, '&gt;');
        }
        $this.checkElementIsAcePreTag = function (element) {
            if (element.nodeName.toLowerCase() == "pre" && element.className.indexOf('ace-editor') != -1) {
                return true;
            }
            return false;
        }

        $this.openDialog = function () {
            var $node = null;
            var modalTitle;
            var mode = "html";
            var width = "100%";
            var height = "200px";
            var code = "";
            if ($this.checkElementIsAcePreTag(editor.selection.getNode())) {
                modalTitle = "Edit Code";
                $node = $(editor.selection.getNode());
                if ($node.attr("data-ace-mode")) {
                    mode = $node.attr("data-ace-mode");
                }
                if ($node.attr("data-ace-width")) {
                    width = $node.attr("data-ace-width");
                }
                if ($node.attr("data-ace-height")) {
                    height = $node.attr("data-ace-height");
                }
                code = editor.selection.getNode().innerHTML;
            } else {
                $node = null;
                modalTitle = "Insert Code";
            }

            $this.dialog = tinymce.activeEditor.windowManager.openUrl({
                title: modalTitle,
                url: pluginDir + '/' + 'dialog.html',
                buttons: [{
                        type: 'custom',
                        name: 'save',
                        text: 'Save',
                        primary: true,
                        align: 'end'
                    },
                    {
                        type: 'cancel',
                        name: 'cancel',
                        text: 'Close',
                        primary: false,
                        align: 'end'
                    }
                ],
                onAction: function (dialogApi, trigger) {
                    switch (trigger.name) {
                        case 'save':
                            dialogApi.sendMessage({
                                mceAction: "save",
                            });
                            break;
                    }
                },
                onMessage: function (dialogApi, data) {
                    switch (data.mceAction) {
                        case 'init':
                            dialogApi.sendMessage({
                                mceAction: "init",
                                mode: mode,
                                code: code,
                                width: width,
                                height: height,
                            });
                            break;
                        case 'save':
                            code = data.code;
                            mode = data.mode;
                            width = data.width;
                            height = data.height;
                            if (code && code != "") {
                                var preTag = "<pre class=\"ace-editor\" data-ace-mode=\"" + mode + "\" data-ace-width=\"" + width + "\" data-ace-height=\"" + height + "\" contenteditable=\"false\" style=\"width: " + width + "; height: " + height + ";\">" + $this.escapePreHtml(code) + "</pre><p></p>";
                                if ($node !== null) {
                                    $node[0].outerHTML = preTag;
                                } else {
                                    editor.insertContent(preTag);
                                }
                            }
                            $this.dialog.close();
                            break;
                    }
                }
            });

        };


        editor.on('dblclick', function (e) {
            if ($this.checkElementIsAcePreTag(e.target)) {
                $this.openDialog();
            }
        });

        editor.ui.registry.addButton('ace', {
            text: 'ACE Editor',
            onAction: function () {
                $this.openDialog();
            }
        });
        // Adds a menu item to the tools menu
        editor.ui.registry.addMenuItem('ace', {
            text: 'ACE Editor',
            context: 'insert',
            onAction: function () {
                // Open window with a specific url
                $this.openDialog();
            }
        });
    });
}());