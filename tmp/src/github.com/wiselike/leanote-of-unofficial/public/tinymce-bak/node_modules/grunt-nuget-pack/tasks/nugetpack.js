var Package = require("../lib/Package");
var path = require("path");

module.exports = function(grunt) {
	grunt.registerMultiTask("nugetpack", "Creates nupkg packages.", function() {
		var done = this.async(), target = grunt.config([this.name, this.target]);
		var packageFilePath, baseDir, package;

		if (typeof target.options != "object") {
			grunt.fail.fatal("Required meta information not specified.");
		}

		package = new Package(target.options);

		if (target.options.baseDir) {
			baseDir = target.options.baseDir;
		} else {
			baseDir = process.cwd();
		}

		if (target.options.excludes) {
			package.addExcludes(grunt.file.expand(target.options.excludes));
		}

		this.files.forEach(function(filePair) {
			filePair.src.forEach(function(src) {
				var dest = filePair.dest;

				if (!dest) {
					if (path.resolve(src).indexOf(path.resolve(baseDir)) !== 0) {
						grunt.fail.fatal("Path for file: " + src + " isn't within the baseDir: " + baseDir);
					}

					dest = path.join("content", path.relative(baseDir, src));
				}

				try {
					package.addFile(src, dest);
				} catch (ex) {
					grunt.fail.fatal(ex);
				}
			});
		});

		packageFilePath = target.options.id + "." + target.options.version + ".nupkg";
		if (target.options.outputDir) {
			grunt.file.mkdir(target.options.outputDir);
			packageFilePath = path.join(target.options.outputDir, packageFilePath);
		}

		try {
			package.saveAs(packageFilePath, done);
			grunt.log.ok("Created nupkg file:", packageFilePath);
		} catch (ex) {
			grunt.fail.fatal(ex);
		}
	});
};
