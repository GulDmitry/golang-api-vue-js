const webpack = require('webpack');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

process.env.NODE_ENV = process.env.NODE_ENV || 'development';

const config = {
    entry: [
        'babel-polyfill',
        './static/js/app.js'
    ],
    output: {
        filename: 'bundle.js',
        path: 'static/assets/',
        publicPath: '/assets/'
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                loader: 'eslint-loader',
                enforce: 'pre',
                exclude: /node_modules/,
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'babel-loader'
            },
            {
                test: /\.css$/,
                loader: ExtractTextPlugin.extract({
                    fallbackLoader: 'style-loader',
                    loader: 'css-loader'
                })
            },
            {
                test: /\.less$/,
                loader: ExtractTextPlugin.extract({
                    fallbackLoader: 'style-loader',
                    loader: 'css-loader!less-loader'
                })
            },
            {
                test: /\.png$/,
                loader: 'url-loader?limit=100000'
            },
            {
                test: /\.jpg$/,
                loader: 'file-loader'
            },
            {test: /\.(woff2?|svg)$/, loader: 'url-loader?limit=10000'},
            {test: /\.(ttf|eot)$/, loader: 'file-loader'}
        ]
    },
    plugins: [
        new ExtractTextPlugin({
            filename: 'bundle.css'
        }),
        new webpack.ProvidePlugin({
            $: 'jquery',
            jQuery: 'jquery'
        }),
        new webpack.NoEmitOnErrorsPlugin()
    ],
    externals: {},
    resolve: {extensions: ['.js']}
};

if (process.env.ENV === 'production') {
    config.devtool = 'source-map';
    config.plugins.push(
        new webpack.optimize.UglifyJsPlugin({
            sourceMap: true,
            beautify: false,
            comments: false,
            compress: {
                sequences: true,
                booleans: true,
                loops: true,
                unused: true,
                warnings: true,
                drop_console: true,
                unsafe: true
            }
        }),
        new webpack.LoaderOptionsPlugin({
            minimize: true
        })
    );

} else {
    // config.plugins.push(
    //     new webpack.HotModuleReplacementPlugin()
    // );
}

module.exports = config;
