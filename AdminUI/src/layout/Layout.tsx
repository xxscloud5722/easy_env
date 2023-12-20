import React, { FC, useEffect, useState } from 'react';
import { ProConfigProvider, ProLayout } from '@ant-design/pro-components';
import { Button, ConfigProvider, Input, Modal, theme } from 'antd';
import { css } from '@emotion/css';
import { Session } from 'beer-network/session';
import { BrowserRouter, Link, Navigate, Route, Routes } from 'react-router-dom';
import ImageLayout01 from '@assets/layout_01.png';
import ImageLayout02 from '@assets/layout_02.png';
import ImageLayout03 from '@assets/layout_03.png';
import ImageLogo from '@assets/logo.svg';
import KeyValuePage from '@/pages/KeyValuePage.tsx';

const App: FC = () => {
  const [pathname, setPathname] = useState('/');
  const [dialogTokenModal, setDialogTokenModal] = useState(false);
  const [token, setToken] = useState('');
  const onConfirmToken = () => {
    if (token === '') {
      return;
    }
    localStorage.setItem('login_user', JSON.stringify({ token }));
    setDialogTokenModal(false);
    window.location.reload();
  };
  useEffect(() => {
    if (!(window.location.host.indexOf('127.0.0.1') > -1 || window.location.host.indexOf('localhost') > -1) && !Session.checkAccessTokenExist()) {
      setDialogTokenModal(true);
    }
  }, []);
  return (
    <ConfigProvider theme={{
      algorithm: theme.defaultAlgorithm,
      token: {
        borderRadius: 4
      },
      components: {
        Select: {},
        Divider: {
          orientationMargin: 0
        },
        Table: {
          motion: false
        },
        Tree: {
          directoryNodeSelectedBg: '#efefef',
          directoryNodeSelectedColor: '#333'
        }
      }
    }}>
      <ProConfigProvider hashed={false}>
        <BrowserRouter basename={import.meta.env.VITE_ROUTE_BASE?.toString()}>
          <ProLayout
            fixSiderbar={true}
            layout="mix"
            contentWidth="Fluid"
            splitMenus={true}
            bgLayoutImgList={[
              {
                src: ImageLayout01,
                left: 85,
                bottom: 100,
                height: '303px'
              },
              {
                src: ImageLayout02,
                bottom: -68,
                right: -45,
                height: '303px'
              },
              {
                src: ImageLayout03,
                bottom: 0,
                left: 0,
                width: '331px'
              }
            ]}
            token={{
              bgLayout: '#f5f7fa',
              sider: {
                colorMenuBackground: '#fff',
                colorTextMenu: 'rgb(87, 89, 102)',
                colorBgMenuItemSelected: '#ededed'
              },
              pageContainer: {
                paddingInlinePageContainerContent: 0,
                paddingBlockPageContainerContent: 0
              },
              header: {
                colorBgMenuItemSelected: 'rgba(255,255,255,0.04)',
                colorBgHeader: '#1f2227',
                colorHeaderTitle: '#fff',
                colorTextMenu: '#b6babf',
                colorTextMenuActive: '#fff',
                colorTextMenuSelected: '#fff',
                colorBgMenuElevated: '#1f2227',
                colorTextRightActionsItem: '#fff',
                colorBgRightActionsItemHover: '#333'
              }
            }}
            siderMenuType="sub"
            fixedHeader={true}
            location={{
              pathname
            }}
            menuItemRender={(props, defaultDom) => {
              if (props.isUrl || !props.path) {
                return defaultDom;
              }
              return <Link to={props.path} onClick={() => {
                setPathname(props.path || '/');
              }}>{defaultDom}</Link>;
            }}
            headerContentRender={(_, defaultDom) => {
              return <>
                {defaultDom}
              </>;
            }}
            menuExtraRender={() => (
              <div className={css`
                  display: flex;
                  height: 36px;
                  align-items: center;
                  padding-bottom: 6px;

                  & > img {
                      width: 24px;
                      height: 24px;
                      margin-right: 8px;
                  }

                  & > p {
                      font-size: 16px;
                      font-weight: 500;
                      min-width: 80px;
                      overflow: hidden;
                  }
              `}>
                <img src={ImageLogo} alt=""/>
              </div>
            )}
            logo={<div className={css`
                display: flex;
                align-items: center;

                & > img {
                    height: 30px;
                }
            `}>
              <img src={ImageLogo} alt=""/>
            </div>}
            title=""
          >
            <div style={{ padding: '20px 20px 20px 20px' }}>
              <Routes>
                <Route path="/" Component={() => <Navigate to="/kv"/>}/>
                <Route path="/kv" Component={KeyValuePage}/>
              </Routes>
            </div>
            <Modal
              open={dialogTokenModal}
              title="输入令牌"
              width="420px"
              destroyOnClose={false}
              closeIcon={false}
              cancelText={null}
              styles={{ body: { padding: '10px 0' } }}
              footer={[
                <Button key="submit" onClick={() => onConfirmToken()} type="primary">确定</Button>
              ]}>
              <Input
                placeholder="请输入令牌"
                value={token}
                onChange={(e) => setToken(e.target.value)}/>
            </Modal>
          </ProLayout>
        </BrowserRouter>
      </ProConfigProvider>
    </ConfigProvider>
  );
};
export default App;
