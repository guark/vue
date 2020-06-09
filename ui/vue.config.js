const fs = require('fs')
const path = require('path')
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');


// const { useGuarkLockFile, checkBeforeBuild } = require('guark/build')


// if (err = checkBeforeBuild()) {
// 		console.error(err.message)
// 		process.exit(err.code)
// }
if (process.env.NODE_ENV == 'production' && !process.env.GUARK_BUILD_DIR) {

	console.error('ERROR: Guark build dir not provided!')
	process.exit(1)
}


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
			// Roger that.
			// after: (app, server, compiler) => compiler.hooks.done.tap("Guark", useGuarkLockFile)
			after: function(app, server, compiler)
			{
				compiler.hooks.done.tap('Guark', (params) =>
				{
					fs.writeFile(`${process.cwd()}/guark.lock`, 'Do not delete me!', function (err)
					{
						if (err) throw err;
					})
				})
			}
		},
		optimization: {
			minimizer: [
				new UglifyJsPlugin(),
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