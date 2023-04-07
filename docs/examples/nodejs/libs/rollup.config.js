/*
 * @Author: Bin
 * @Date: 2023-04-01
 * @FilePath: /gpt-zmide-server/docs/examples/nodejs/libs/rollup.config.js
 */
import path from 'node:path'
import glob from 'glob'
import { URL, fileURLToPath } from 'node:url'
import process from 'node:process'

import nodeResolve from '@rollup/plugin-node-resolve'
import commonjs from '@rollup/plugin-commonjs'
import typescript from '@rollup/plugin-typescript'
import json from '@rollup/plugin-json'

const name = 'GPTServer'
const fileName = 'index'
const umdInput = 'src/index.ts'
const outputDir = 'dist/'
let sourcemap = false
let isProd = true
const inputList = Object.fromEntries(
    glob.sync('src/**/*.ts').map(file => [path.relative('src', file.slice(0, file.length - path.extname(file).length)), fileURLToPath(new URL(file, import.meta.url))])
)

if (process.env.ENV === 'development') {
    // 开发模式时
    isProd = false
    sourcemap = true
}

const getPlugins = ({ isProd = true } = {}) => {
    const plugins = [
        nodeResolve(),
        commonjs(),
        json(),
        typescript({ tsconfig: './tsconfig.json' }),
    ]
    return plugins
}



// CommonJS
// const cjsConfig = {
//     input: umdInput,
//     plugins: getPlugins({
//         isProd
//     }),
//     output: {
//         format: 'cjs',
//         sourcemap,
//         exports: 'named',
//         chunkFileNames: '_chunks/dep-[hash].js',
//         file: `${outputDir}${fileName}.cjs`,
//     },
// };


// ES module file
const esmConfig = {
    input: inputList,
    treeshake: false,
    plugins: getPlugins({
        isProd,
    }),
    output: {
        format: 'esm',
        sourcemap,
        chunkFileNames: '_chunks/dep-[hash].js',
        dir: outputDir,
    },
}

// Universal Module
const umdConfig = {
    input: umdInput,
    plugins: getPlugins({
        isProd,
    }),
    output: {
        name,
        format: 'umd',
        exports: 'named',
        sourcemap,
        file: `${outputDir}${fileName}.umd.js`,
    },
}

const umdMinConfig = {
    input: umdInput,
    plugins: getPlugins({
        isProd: true,
    }),
    output: {
        name,
        format: 'umd',
        exports: 'named',
        sourcemap,
        file: `${outputDir}${fileName}.umd.min.js`,
    },
}

export default [esmConfig, umdConfig, umdMinConfig]