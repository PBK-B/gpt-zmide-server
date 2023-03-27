/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/screens/index.tsx
 */
import Home from './home'
import Application from './application'
import System from './system'
import EmptyState from './empty'

const HomeScreen = <Home />
const ApplicationScreen = <Application />
const EmptyStateScreen = <EmptyState />
const SystemScreen = <System />

export { HomeScreen, ApplicationScreen, SystemScreen, EmptyStateScreen }