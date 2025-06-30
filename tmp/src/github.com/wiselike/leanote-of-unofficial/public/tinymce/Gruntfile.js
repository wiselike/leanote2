module.exports = function(grunt) {
	var packageData = grunt.file.readJSON("package.json");
	var changelogLine = grunt.file.read("changelog.txt").toString().split("\n")[0];
	packageData.version = /^Version ([0-9xabrc.]+)/.exec(changelogLine)[1];
	packageData.date = /^Version [^\(]+\(([^\)]+)\)/.exec(changelogLine)[1];

	grunt.initConfig({
		pkg: packageData,

		eslint: {
			options: {
				config: ".eslintrc"
			},

			core: ["classes/**/*.js"],

			plugins: [
				"plugins/*/plugin.js",
				"plugins/*/classes/**/*.js",
				"!plugins/paste/plugin.js",
				"!plugins/table/plugin.js",
				"!plugins/spellchecker/plugin.js"
			],

			themes: ["themes/*/theme.js"]
		},

		jshint: {
			core: ["classes/**/*.js"],

			plugins: [
				"plugins/*/plugin.js",
				"plugins/*/classes/**/*.js",
				"!plugins/paste/plugin.js",
				"!plugins/table/plugin.js",
				"!plugins/spellchecker/plugin.js"
			],

			themes: ["themes/*/theme.js"]
		},

		jscs: {
			options: {
				config: ".jscsrc"
			},

			core: ["**/*.js"],

			plugins: [
				"plugins/*/plugin.js",
				"plugins/*/classes/**.js",
				"!plugins/paste/plugin.js",
				"!plugins/table/plugin.js",
				"!plugins/spellchecker/plugin.js"
			],

			themes: ["themes/*/theme.js"]
		},

		qunit: {
			core: {
				options: {
					urls: [
						"tests/index.html"
					]
				}
			}
		},

		amdlc: {
			core: {
				options: {
					version: packageData.version,
					releaseDate: packageData.date,
					baseDir: "classes",
					rootNS: "tinymce",
					outputSource: "tinymce.js",
					outputMinified: "tinymce.min.js",
					outputDev: "tinymce.dev.js",
					verbose: true,
					expose: "public",
					compress: true,

					from: [
						"dom/DomQuery.js",
						"EditorManager.js",
						"LegacyInput.js",
						"util/XHR.js",
						"util/JSONRequest.js",
						"util/JSONP.js",
						"util/LocalStorage.js",
						"Compat.js",
						"ui/*.js"
					]
				}
			},

			"core-jquery": {
				options: {
					moduleOverrides: {
						"tinymce/dom/Sizzle": "classes/dom/Sizzle.jQuery.js"
					},
					version: packageData.version,
					releaseDate: packageData.date,
					baseDir: "classes",
					rootNS: "tinymce",
					outputSource: "tinymce.jquery.js",
					outputMinified: "tinymce.jquery.min.js",
					outputDev: "tinymce.jquery.dev.js",
					verbose: true,
					expose: "public",
					compress: true,

					from: [
						"dom/DomQuery.js",
						"EditorManager.js",
						"LegacyInput.js",
						"util/XHR.js",
						"util/JSONRequest.js",
						"util/JSONP.js",
						"util/LocalStorage.js",
						"Compat.js",
						"ui/*.js"
					]
				}
			},

			"paste-plugin": {
				options: {
					baseDir: "plugins/paste/classes",
					rootNS: "tinymce.pasteplugin",
					outputSource: "plugins/paste/plugin.js",
					outputMinified: "plugins/paste/plugin.min.js",
					outputDev: "plugins/paste/plugin.dev.js",
					verbose: true,
					expose: "public",
					compress: true,

					from: "Plugin.js"
				}
			},

			"table-plugin": {
				options: {
					baseDir: "plugins/table/classes",
					rootNS: "tinymce.tableplugin",
					outputSource: "plugins/table/plugin.js",
					outputMinified: "plugins/table/plugin.min.js",
					outputDev: "plugins/table/plugin.dev.js",
					verbose: true,
					expose: "public",
					compress: true,

					from: "Plugin.js"
				}
			},

			"spellchecker-plugin": {
				options: {
					baseDir: "plugins/spellchecker/classes",
					rootNS: "tinymce.spellcheckerplugin",
					outputSource: "plugins/spellchecker/plugin.js",
					outputMinified: "plugins/spellchecker/plugin.min.js",
					outputDev: "plugins/spellchecker/plugin.dev.js",
					verbose: true,
					expose: "public",
					compress: true,

					from: "Plugin.js"
				}
			}
		},

		skin: {
			modern: {
				options: {
					prepend: [
						"Variables.less",
						"Reset.less",
						"Mixins.less",
						"Animations.less",
						"TinyMCE.less"
					],
					append: ["Icons.less"],
					importFrom: "tinymce.js",
					path: "skins",
					ext: ".modern.dev.less"
				}
			},

			ie7: {
				options: {
					prepend: [
						"Variables.less",
						"Reset.less",
						"Mixins.less",
						"Animations.less",
						"TinyMCE.less"
					],
					append: ["Icons.Ie7.less"],
					importFrom: "tinymce.js",
					path: "skins",
					ext: ".ie7.dev.less"
				}
			}
		},

		less: {
			modern: {
				options: {
					cleancss: true,
					strictImports: true
				},

				expand: true,
				src: ["skins/**/skin.modern.dev.less"],
				ext: ".min.css"
			},

			ie7: {
				options: {
					compress: true,
					strictImports: true,
					ieCompat: true
				},

				expand: true,
				src: ["skins/**/skin.ie7.dev.less"],
				ext: ".ie7.min.css"
			},

			content: {
				options: {
					cleancss: true,
					strictImports: true
				},

				rename: function(dest, src) {
					return src.toLowerCase();
				},

				expand: true,
				src: ["skins/**/Content.less"],
				ext: ".min.css"
			},

			"content-inline": {
				options: {
					cleancss: true,
					strictImports: true
				},

				rename: function(dest, src) {
					return src.toLowerCase();
				},

				expand: true,
				src: ["skins/**/Content.Inline.less"],
				ext: ".inline.min.css"
			}
		},

		uglify: {
			options: {
				beautify: {
					ascii_only: true
				}
			},

			themes: {
				src: ["themes/*/theme.js"],
				expand: true,
				ext: ".min.js"
			},

			plugins: {
				src: ["plugins/*/plugin.js"],
				expand: true,
				ext: ".min.js"
			},

			"jquery-plugin": {
				src: ["classes/jquery.tinymce.js"],
				dest: "jquery.tinymce.min.js"
			}
		},

		moxiezip: {
			production: {
				options: {
					baseDir: "tinymce",

					excludes: [
						"plugins/moxiemanager",
						"plugins/compat3x",
						"plugins/visualblocks/img",
						"plugins/*/classes/**",
						"plugins/*/plugin.js",
						"plugins/*/plugin.dev.js",
						"themes/*/theme.js",
						"skins/*/*.less",
						"skins/*/fonts/*.json",
						"skins/*/fonts/*.dev.svg",
						"skins/*/fonts/readme.md",
						"readme.md"
					],

					to: "tmp/tinymce_<%= pkg.version %>.zip"
				},

				src: [
					"langs",
					"plugins",
					"skins",
					"themes",
					"tinymce.min.js",
					"license.txt",
					"changelog.txt",
					"LICENSE.TXT",
					"readme.md"
				]
			},

			jquery: {
				options: {
					baseDir: "tinymce",

					excludes: [
						"plugins/moxiemanager",
						"plugins/compat3x",
						"plugins/visualblocks/img",
						"plugins/*/classes/**",
						"plugins/*/plugin.js",
						"plugins/*/plugin.dev.js",
						"themes/*/theme.js",
						"skins/*/*.less",
						"skins/*/fonts/*.json",
						"skins/*/fonts/*.dev.svg",
						"skins/*/fonts/readme.md",
						"readme.md"
					],

					pathFilter: function(args) {
						if (args.zipFilePath == "tinymce.jquery.min.js") {
							args.zipFilePath = "tinymce.min.js";
						}
					},

					to: "tmp/tinymce_<%= pkg.version %>_jquery.zip"
				},

				src: [
					"langs",
					"plugins",
					"skins",
					"themes",
					"tinymce.jquery.min.js",
					"jquery.tinymce.min.js",
					"license.txt",
					"changelog.txt",
					"LICENSE.TXT",
					"readme.md"
				]
			},

			development: {
				options: {
					baseDir: "tinymce",

					excludes: [
						"tinymce.full.min.js",
						"plugins/moxiemanager",
						"js/tests/.jshintrc"
					],

					to: "tmp/tinymce_<%= pkg.version %>_dev.zip"
				},

				src: [
					"js",
					"tests",
					"tools",
					"changelog.txt",
					"LICENSE.TXT",
					"Gruntfile.js",
					"readme.md",
					"package.json",
					".eslintrc",
					".jscsrc",
					".jshintrc"
				]
			},

			component: {
				options: {
					excludes: [
						"plugins/moxiemanager",
						"plugins/example",
						"plugins/example_dependency",
						"plugins/compat3x",
						"plugins/visualblocks/img",
						"plugins/*/classes/**",
						"plugins/*/plugin.dev.js",
						"skins/*/*.less",
						"skins/*/fonts/*.json",
						"skins/*/fonts/*.dev.svg",
						"skins/*/fonts/readme.md",
						"readme.md"
					],

					pathFilter: function(args) {
						if (args.zipFilePath.indexOf("") === 0) {
							args.zipFilePath = args.zipFilePath.substr("".length);
						}
					},

					onBeforeSave: function(zip) {
						function jsonToBuffer(json) {
							return new Buffer(JSON.stringify(json, null, '\t'));
						}

						zip.addData("bower.json", jsonToBuffer({
							"name": "tinymce",
							"version": packageData.version,
							"description": "Web based JavaScript HTML WYSIWYG editor control.",
							"license": "http://www.tinymce.com/license",
							"keywords": ["editor", "wysiwyg", "tinymce", "richtext", "javascript", "html"],
							"homepage": "http://www.tinymce.com",
							"main": "tinymce.min.js",
							"ignore": ["readme.md", "composer.json", "package.json"]
						}));

						zip.addData("package.json", jsonToBuffer({
							"name": "tinymce",
							"version": packageData.version,
							"description": "Web based JavaScript HTML WYSIWYG editor control.",
							"license": "LGPL-2.1",
							"keywords": ["editor", "wysiwyg", "tinymce", "richtext", "javascript", "html"],
							"bugs": {"url": "http://www.tinymce.com/develop/bugtracker.php"}
						}));

						zip.addData("composer.json", jsonToBuffer({
							"name": "tinymce/tinymce",
							"version": packageData.version,
							"description": "Web based JavaScript HTML WYSIWYG editor control.",
							"license": ["LGPL-2.1"],
							"keywords": ["editor", "wysiwyg", "tinymce", "richtext", "javascript", "html"],
							"homepage": "http://www.tinymce.com",
							"type": "component",
							"extra": {
								"component": {
									"scripts": [
										"tinymce.js",
										"plugins/*/plugin.js",
										"themes/*/theme.js"
									],
									"files": [
										"tinymce.min.js",
										"plugins/*/plugin.min.js",
										"themes/*/theme.min.js",
										"skins/**"
									]
								}
							},
							"archive": {
								"exclude": ["readme.md", "bower.js", "package.json"]
							}
						}));
					},

					to: "tmp/tinymce_<%= pkg.version %>_component.zip"
				},

				src: [
					"skins",
					"plugins",
					"themes",
					"tinymce.js",
					"tinymce.min.js",
					"jquery.tinymce.min.js",
					"tinymce.jquery.js",
					"tinymce.jquery.min.js",
					"license.txt",
					"changelog.txt"
				]
			}
		},

		connect: {
			server: {
				options: {
					port: 9999
				}
			}
		},

		"saucelabs-qunit": {
			all: {
				options: {
					urls: ["127.0.0.1:9999/tests/index.html?min=true"],
					testname: "TinyMCE QUnit Tests",
					browsers: [
						{browserName: "firefox", platform: "XP"},
						{browserName: "googlechrome", platform: "XP"},
						{browserName: "firefox", platform: "Linux"},
						{browserName: "googlechrome", platform: "Linux"},
						{browserName: "internet explorer", platform: "XP", version: "8"},
						{browserName: "internet explorer", platform: "Windows 7", version: "9"},
						{browserName: "internet explorer", platform: "Windows 7", version: "10"},
						{browserName: "internet explorer", platform: "Windows 7", version: "11"},
						{browserName: "safari", platform: "OS X 10.9", version: "7"},
						{browserName: "safari", platform: "OS X 10.8", version: "6"}
					]
				}
			}
		},

		nugetpack: {
			main: {
				options: {
					id: "TinyMCE",
					version: packageData.version,
					authors: "Moxiecode Systems AB",
					owners: "Moxiecode Systems AB",
					description: "The best WYSIWYG editor! TinyMCE is a platform independent web based Javascript HTML WYSIWYG editor " +
						"control released as Open Source under LGPL by Moxiecode Systems AB. TinyMCE has the ability to convert HTML " +
						"TEXTAREA fields or other HTML elements to editor instances. TinyMCE is very easy to integrate " +
						"into other Content Management Systems.",
					releaseNotes: "Release notes for my package.",
					summary: "TinyMCE is a platform independent web based Javascript HTML WYSIWYG editor " +
						"control released as Open Source under LGPL by Moxiecode Systems AB.",
					projectUrl: "http://www.tinymce.com/",
					iconUrl: "http://www.tinymce.com/favicon.ico",
					licenseUrl: "http://www.tinymce.com/license",
					requireLicenseAcceptance: true,
					tags: "Editor TinyMCE HTML HTMLEditor",
					excludes: [
						"skins/**/*.dev.svg",
						"skins/**/*.less",
						"plugins/**/classes",
						"plugins/**/*.dev.js"
					],
					outputDir: "tmp"
				},

				files: [
					{src: "langs", dest: "/content/scripts/tinymce/langs"},
					{src: "plugins", dest: "/content/scripts/tinymce/plugins"},
					{src: "themes", dest: "/content/scripts/tinymce/themes"},
					{src: "skins", dest: "/content/scripts/tinymce/skins"},
					{src: "tinymce.js", dest: "/content/scripts/tinymce/tinymce.js"},
					{src: "tinymce.min.js", dest: "/content/scripts/tinymce/tinymce.min.js"},
					{src: "license.txt", dest: "/content/scripts/tinymce/license.txt"}
				]
			},

			jquery: {
				options: {
					id: "TinyMCE.jQuery",
					version: packageData.version,
					authors: "Moxiecode Systems AB",
					owners: "Moxiecode Systems AB",
					description: "The best WYSIWYG editor! TinyMCE is a platform independent web based Javascript HTML WYSIWYG editor " +
						"control released as Open Source under LGPL by Moxiecode Systems AB. TinyMCE has the ability to convert HTML " +
						"TEXTAREA fields or other HTML elements to editor instances. TinyMCE is very easy to integrate " +
						"into other Content Management Systems.",
					releaseNotes: "Release notes for my package.",
					summary: "TinyMCE is a platform independent web based Javascript HTML WYSIWYG editor " +
						"control released as Open Source under LGPL by Moxiecode Systems AB.",
					projectUrl: "http://www.tinymce.com/",
					iconUrl: "http://www.tinymce.com/favicon.ico",
					licenseUrl: "http://www.tinymce.com/license",
					requireLicenseAcceptance: true,
					tags: "Editor TinyMCE HTML HTMLEditor",
					excludes: [
						"skins/**/*.dev.svg",
						"skins/**/*.less",
						"plugins/**/classes",
						"plugins/**/*.dev.js"
					],
					outputDir: "tmp"
				},

				files: [
					{src: "langs", dest: "/content/scripts/tinymce/langs"},
					{src: "plugins", dest: "/content/scripts/tinymce/plugins"},
					{src: "themes", dest: "/content/scripts/tinymce/themes"},
					{src: "skins", dest: "/content/scripts/tinymce/skins"},
					{src: "tinymce.js", dest: "/content/scripts/tinymce/tinymce.js"},
					{src: "tinymce.min.js", dest: "/content/scripts/tinymce/tinymce.min.js"},
					{src: "jquery.tinymce.min.js", dest: "/content/scripts/tinymce/jquery.tinymce.min.js"},
					{src: "license.txt", dest: "/content/scripts/tinymce/license.txt"}
				]
			}
		},

		bundle: {
			minified: {
				options: {
					themesDir: "themes",
					pluginsDir: "plugins",
					pluginFileName: "plugin.min.js",
					themeFileName: "theme.min.js",
					outputPath: "tinymce.full.min.js"
				},

				src: [
					"tinymce.min.js"
				]
			},

			source: {
				options: {
					themesDir: "themes",
					pluginsDir: "plugins",
					pluginFileName: "plugin.js",
					themeFileName: "theme.js",
					outputPath: "tinymce.full.js"
				},

				src: [
					"tinymce.js"
				]
			}
		},

		clean: {
			release: ["tmp"],

			core: [
				"tinymce*",
				"*.min.js",
				"*.dev.js"
			],

			plugins: [
				"plugins/**/*.min.js",
				"plugins/**/*.dev.js",
				"plugins/table/plugin.js",
				"plugins/paste/plugin.js",
				"plugins/spellchecker/plugin.js"
			],

			skins: [
				"skins/**/*.min.css",
				"skins/**/*.dev.less"
			],

			npm: [
				"node_modules",
				"npm-debug.log"
			],

			saucelabs: [
				"?sc.log",
				"sc_*.log"
			]
		},

		watch: {
			core: {
				files: ["classes/**/*.js"],
				tasks: ["eslint:core", "jshint:core", "jscs:core", "amdlc:core", "amdlc:core-jquery", "skin"],
				options: {
					spawn: false
				}
			},

			plugins: {
				files: ["plugins/**/*.js"],
				tasks: [
					"eslint:plugins", "jshint:plugins", "jscs:plugins", "amdlc:paste-plugin",
					"amdlc:table-plugin", "amdlc:spellchecker-plugin", "uglify:plugins"
				],
				options: {
					spawn: false
				}
			},

			themes: {
				files: ["themes/**/*.js"],
				tasks: ["eslint:themes", "jshint:themes", "jscs:themes", "uglify:themes"],
				options: {
					spawn: false
				}
			},

			skins: {
				files: ["skins/**/*"],
				tasks: ["less"],
				options: {
					spawn: false
				}
			}
		}
	});

	require("load-grunt-tasks")(grunt);
	grunt.loadTasks("tools/tasks");

	grunt.registerTask("lint", ["eslint", "jshint", "jscs"]);
	grunt.registerTask("minify", ["amdlc", "uglify", "skin", "less"]);
	grunt.registerTask("test", ["qunit"]);
	grunt.registerTask("sc-test", ["connect", "clean:saucelabs", "saucelabs-qunit"]);
	grunt.registerTask("default", ["lint", "minify", "test", "clean:release", "moxiezip", "nugetpack"]);
};