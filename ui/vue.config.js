const fs = require('fs')
const path = require('path')
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');


const { useGuarkLockFile, checkBeforeBuild } = require('guark/build')


checkBeforeBuild()


module.exports =
{
	outputDir: process.env.GUARK_BUILD_DIR,
	pages:
	{
		index:
		{
			entry:    'src/main.js',
			template: path.resolve(__dirname, '/src/index.html'),
		},
	},

	productionSourceMap: process.env.NODE_ENV == 'production' ? false : true,

	configureWebpack:
	{
		devServer:
		{
			// After server started you should call useGuarkLockFile.
			after: (app, server, compiler) => compiler.hooks.done.tap("Guark", useGuarkLockFile)
		},
		optimization: {
			minimizer: [
//				new UglifyJsPlugin(),
			],
		},
	},

	chainWebpack: config =>
	{
		config
			.plugin('copy')
			.use(require('copy-webpack-plugin'), [[
			{
				from: path.resolve(__dirname, '../res/static'),
				to: process.env.GUARK_BUILD_DIR,
				toType: 'dir',
				ignore: ['.DS_Store']
			}]])
	}
}
