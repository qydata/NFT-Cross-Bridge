import { Route, Switch } from 'react-router-dom';
import HomePage from 'src/pages/HomePage';
import BridgePage from 'src/pages/BridgePage';
import StatusPage from 'src/pages/StatusPage';
import CallbackPage from 'src/pages/CallbackPage';
import LogoutPage from './pages/LogoutPage';
import MyNFTPage from './pages/MyNFTPage';

import 'antd/dist/antd.less';
import { ConfigProvider } from 'antd';
import zhCN from 'antd/es/locale/zh_CN';
import './App.css';

const App: React.FC = () => {
  return (
    <ConfigProvider locale={zhCN}>
      <Switch>
        <Route path='/bridge' component={BridgePage} />
        <Route path='/my-nft' component={MyNFTPage} />
        <Route path='/status' component={StatusPage} />
        <Route path='/callback' component={CallbackPage} />
        <Route path='/logout' component={LogoutPage} />
        <Route path='/' component={HomePage} />
      </Switch>
    </ConfigProvider>
  );
};

export default App;
