/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/apis/axios.ts
 */
import axios from 'axios';

import { configure } from 'axios-hooks';

const initAxios = (token?: string) => {
	axios.defaults.baseURL = '/';
	axios.defaults.headers.common['Accept'] = 'application/json';

	// 将 axios 对象配置到 axios-hooks
	configure({ axios });
};

export { default as axios } from 'axios';

export * from 'axios';
export * from 'axios-hooks';

export default axios;
export { initAxios };