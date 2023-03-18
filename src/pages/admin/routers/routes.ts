/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/routers/routes.ts
 */
import { HomeScreen, ApplicationScreen } from '../screens'

const routes = [
    {
        path: "/",
        component: HomeScreen,
    },
    {
        path: "/app",
        component: ApplicationScreen,
    }
];

export default routes;