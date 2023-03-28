/*
 * @Author: Bin
 * @Date: 2023-03-18
 * @FilePath: /gpt-zmide-server/src/pages/admin/screens/index.tsx
 */
import Home from './home'
import Application from './application'
import System from './system'
import Chat from './chat'
import EmptyState from './empty'

const HomeScreen = <Home />
const ApplicationScreen = <Application />
const EmptyStateScreen = <EmptyState />
const SystemScreen = <System />
const ChatScreen = <Chat />

export {
    HomeScreen,
    ApplicationScreen,
    SystemScreen,
    ChatScreen,
    EmptyStateScreen
}