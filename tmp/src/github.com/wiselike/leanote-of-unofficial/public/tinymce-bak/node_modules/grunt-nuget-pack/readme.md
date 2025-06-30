grunt-nuget-pack
==================

This is a grunt task for creating nuget packages without using the nuget.exe. This makes it easier to build nuget packages on non windows environment or continuous integration services like Travis CI. It's created to produce packages for JavaScript projects not .NET libraries.

### Parameters

#### id

Type: `string` ***Required***

The unique identifier for the package. This is the package name that is shown when packages are listed using the Package Manager Console. These are also used when installing a package using the Install-Package command within the Package Manager Console. Package IDs may not contain any spaces or characters that are invalid in an URL. In general, they follow the same rules as .NET namespaces do. So Foo.Bar is a valid ID, Foo! and Foo Bar are not.

#### version

Type: `string` ***Required***

The version of the package, in a format like 1.2.3.

#### title

Type: `string`

The human-friendly title of the package displayed in the Manage NuGet Packages dialog. If none is specified, the ID is used instead.

#### authors

Type: `string` ***Required***

A comma-separated list of authors of the package code.

#### owners

Type: `string`

A comma-separated list of the package creators. This is often the same list as in authors. This is ignored when uploading the package to the NuGet.org Gallery.
description	A long description of the package. This shows up in the right pane of the Add Package Dialog as well as in the Package Manager Console when listing packages using the Get-Package command.

#### description

Type: `string` ***Required***

A long description of the package. This shows up in the right pane of the Add Package Dialog as well as in the Package Manager Console when listing packages using the Get-Package command.

#### releaseNotes

Type: `string`

A description of the changes made in each release of the package. This field only shows up when the _Updates_ tab is selected and the package is an update to a previously installed package. It is displayed where the Description would normally be displayed.

#### summary

Type: `string`

A short description of the package. If specified, this shows up in the middle pane of the Add Package Dialog. If not specified, a truncated version of the description is used instead.

#### language

Type: `string`

The locale ID for the package, such as en-us.

#### projectUrl

Type: `string`

A URL for the home page of the package.

#### iconUrl

Type: `string`

A URL for the image to use as the icon for the package in the Manage NuGet Packages dialog box. This should be a 32x32-pixel .png file that has a transparent background.

#### licenseUrl

Type: `string`

A link to the license that the package is under.

#### copyright

Type: `string`

Copyright details for the package.

#### requireLicenseAcceptance

Type: `boolean`

A Boolean value that specifies whether the client needs to ensure that the package license (described by licenseUrl) is accepted before the package is installed.

#### dependencies

Type: `array`

The list of dependencies for the package. For more information, see Specifying Dependencies later in this document.

Dependencies should have the id and a [version range](http://docs.nuget.org/docs/reference/versioning#Specifying_Version_Ranges_in_.nuspec_Files)

##### Example of usage
```
dependencies: [
	{id: "dependency1", version: "(1.0,)"},
	{id: "dependency2", version: "1.0"}
]
```

#### tags

Type: `string`

A space-delimited list of tags and keywords that describe the package. This information is used to help make sure users can find the package using searches in the Add Package Reference dialog box or filtering in the Package Manager Console window.

#### excludes

Type: `array`

Array with glob patterns to exclude from the output package. Try to make the glob pattern as exact as possible for performance reasons.

#### outputDir

Type: `string`

A path where the nupkg file is to be created. If the specified directory doesn't exist it will be created.

#### baseDir

Type: `string`

A path that should be treated as the base dir for the contents of the package. This defaults to the current working directory.

### Example of usage

#### Simple example

Simple example showing the essential options this will add the files specified to the "content" directory if the nupkg.

```
grunt.initConfig({
	nugetpack: {
		simple: {
			options: {
				id: "MyPackage",
				version: "1.0.0",
				authors: "Some author",
				description: "Description of my package."
			},

			src: [
				"lib/**/*.js",
				"readme.md"
			]
		}
	}
});
```

#### Advanced example

More advanced example showing all the options notice that you need to specify "/content" as the base path in the package for script resources.

```
grunt.initConfig({
	nugetpack: {
		advanced: {
			options: {
				id: "MyPackage",
				version: "1.0.0",
				authors: "Author1",
				owners: "Owner1",
				description: "Description of my package.",
				releaseNotes: "Release notes for my package.",
				summary: "Summary of my package.",
				language: "en-us",
				projectUrl: "http://www.host.com/",
				iconUrl: "http://www.host.com/icon.png",
				licenseUrl: "http://www.host.com/license",
				copyright: "Copyright details",
				requireLicenseAcceptance: true,
				dependencies: [
					{id: "dependency1", version: "(1.0,)"},
					{id: "dependency2", version: "1.0"}
				],
				tags: "tag1 tag2 tag3",
				excludes: ["js/**/*.dev.js"],
				outputDir: "out"
			},

			files: [
				{src: "somefile.md", dest: "/content/somepath/readme.md"},
				{src: "somedir", dest: "/content/newpath/newdir"}
			]
		}
	}
});
```