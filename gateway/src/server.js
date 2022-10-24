const express = require("express");
const { createProxyMiddleware } = require('http-proxy-middleware');
const app = express()

const cors = require("cors");

app.use(express.json())
    .use(express.urlencoded({ extended: true }))
    .use(cors())
    .use(express.static(__dirname + '/www'))
    .use('/api/v1/project', createProxyMiddleware({ target: 'http://project:3000', changeOrigin: true }))
    .use('/api/v1/contractor', createProxyMiddleware({ target: 'http://contractor:3000', changeOrigin: true }))
    .listen(3000, () => {
    console.log('server listening on port 3000');
});