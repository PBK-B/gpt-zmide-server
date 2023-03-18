/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/apis/index.ts
 */
import axiosInstance, { initAxios } from './axios';

export * from 'axios';
export { default as axios } from 'axios';
export { axiosInstance, initAxios };