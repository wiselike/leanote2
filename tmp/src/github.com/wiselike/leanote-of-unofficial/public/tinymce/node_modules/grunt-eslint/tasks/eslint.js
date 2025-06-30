'use strict';
var path = require('path');
var chalk = require('chalk');
var eslint = require('eslint').cli;

module.exports = function (grunt) {
	grunt.registerMultiTask('eslint', 'Validate files with ESLint', function () {
		var options = this.options();

		if (this.filesSrc.length === 0) {
			grunt.log.writeln(chalk.magenta('Couldn\'t find any files to validate.'));
			return;
		}

		options._ = this.filesSrc; // set positional arguments
		options.config = options.config ? path.resolve(options.config) : '';

		try {
			return eslint.execute(options) === 0;
		} catch (err) {
			grunt.warn(err);
		}
	});
};
