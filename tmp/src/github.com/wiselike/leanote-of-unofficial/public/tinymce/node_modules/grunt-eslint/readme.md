# grunt-eslint [![Build Status](https://travis-ci.org/sindresorhus/grunt-eslint.svg?branch=master)](https://travis-ci.org/sindresorhus/grunt-eslint)

> Validate files with [ESLint](https://github.com/nzakas/eslint)

![screenshot](screenshot.png)


## Install

```sh
$ npm install --save-dev grunt-eslint
```


## Usage

```js
require('load-grunt-tasks')(grunt); // npm install --save-dev load-grunt-tasks

grunt.initConfig({
	eslint: {
		target: ['file.js']
	}
});

grunt.registerTask('default', ['eslint']);
```


## Examples

### Custom config and rules

```js
grunt.initConfig({
	eslint: {
		options: {
			config: 'conf/eslint.json',
			rulesdir: ['conf/rules']
		},
		target: ['file.js']
	}
});
```

### Custom formatter

```js
grunt.initConfig({
	eslint: {
		options: {
			format: require('eslint-tap')
		},
		target: ['file.js']
	}
});
```


## Options

### config

Type: `string`  
Default: [built-in eslint.json](conf/eslint.json)

Path to your [ESLint config file](https://github.com/nzakas/eslint/blob/master/docs/rules/README.md) (`eslint.json`).

### rulesdir

Type: `array`  
Default: [built-in rules directory](https://github.com/nzakas/eslint/tree/master/lib/rules)

Paths to directories with custom rules. Your custom rules will be used in addition to the built-in ones.

Recommended read: [Working with Rules](https://github.com/nzakas/eslint/blob/master/docs/developer-guide/working-with-rules.md)

### format

Type: `string`  
Default: `'stylish'`

Name of a [built-in formatter](https://github.com/nzakas/eslint/tree/master/lib/formatters) or path to a custom one.

Some formatters you might find useful: [eslint-json](https://github.com/sindresorhus/eslint-json), [eslint-tap](https://github.com/sindresorhus/eslint-tap).


## License

MIT © [Sindre Sorhus](http://sindresorhus.com)
