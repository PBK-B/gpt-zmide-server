/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/routers/routes.ts
 */
import { BrowserRouterProps } from 'react-router-dom';
import { HomeScreen, ApplicationScreen, SystemScreen, EmptyStateScreen } from '../screens'

export interface IRouter {
    path: string;
    redirect?: string;
    Component?: JSX.Element | React.FC<BrowserRouterProps> | (() => any);
    isFullPage?: boolean;
    children?: IRouter[];
}

const routes: IRouter[] = [
    {
        path: "/",
        Component: HomeScreen,
    },
    {
        path: "/app",
        Component: ApplicationScreen,
    },
    {
        path: "/system",
        Component: SystemScreen,
    },
    {
        path: "*",
        Component: EmptyStateScreen,
    }
];

export default routes;