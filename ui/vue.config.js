const { useGuarkLockFile, checkBeforeBuild } = require('guark/build')

checkBeforeBuild()

module.exports =
{
	outputDir: process.env.GUARK_BUILD_DIR,
	productionSourceMap: process.env.NODE_ENV == 'production' ? false : true,
	configureWebpack:
	{
		devServer:
		{
			// After server started you should call useGuarkLockFile.
			after: (app, server, compiler) => compiler.hooks.done.tap("Guark", useGuarkLockFile)
		},
	},
}
