import React, { FC, useEffect, useState } from 'react';
import { MenuDataItem, ProConfigProvider, ProLayout } from '@ant-design/pro-components';
import { ConfigProvider, Input, Modal, theme } from 'antd';
import { css } from '@emotion/css';
import { BrowserRouter, Link, Navigate, Route, Routes } from 'react-router-dom';
import ImageLayout01 from '@assets/layout_01.png';
import ImageLayout02 from '@assets/layout_02.png';
import ImageLayout03 from '@assets/layout_03.png';
import ImageLogo from '@assets/logo.svg';
import commonApi from '@/apis/common-api.ts';
import KeyValuePage from '@/pages/KeyValuePage.tsx';
import ScriptPage from '@/pages/ScriptPage.tsx';
import session from '@/apis/session.ts';

const App: FC = () => {
  const [pathname, setPathname] = useState('/');
  const requestMenuItem = async (_params: Record<string, any>, _defaultMenuData: MenuDataItem[]): Promise<MenuDataItem[]> => {
    const menusResponse = await commonApi.getUserMenus();
    const recursion = (menus: any[]): MenuDataItem[] => {
      return (menus || []).map(menu => {
        return {
          key: menu.id,
          name: menu.name,
          path: menu.link,
          children: recursion(menu.child)
        };
      });
    };
    if (menusResponse.success) {
      return recursion(menusResponse.data);
    }
    return [];
  };
  const [dialogTokenModal, setDialogTokenModal] = useState(false);
  const [token, setToken] = useState('');
  const onConfirmToken = () => {
    session.saveToken(token);
    setDialogTokenModal(false);
    window.location.reload();
  };
  useEffect(() => {
    if (session.getToken() === undefined) {
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
            menu={{
              request: requestMenuItem
            }}
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
                <Route path="/script" Component={ScriptPage}/>
              </Routes>
            </div>
            <Modal
              open={dialogTokenModal}
              title="输入令牌"
              width="420px"
              destroyOnClose={false}
              closeIcon={false}
              cancelText={null}
              onOk={() => onConfirmToken()}
              styles={{ body: { padding: '10px 0' } }}
            >
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
