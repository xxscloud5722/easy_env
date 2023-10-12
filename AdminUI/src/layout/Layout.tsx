import React, { FC } from 'react';
import { MenuDataItem, ProConfigProvider, ProLayout } from '@ant-design/pro-components';
import { ConfigProvider, theme } from 'antd';
import { css } from '@emotion/css';
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom';
import ImageLayout01 from '@assets/layout_01.png';
import ImageLayout02 from '@assets/layout_02.png';
import ImageLayout03 from '@assets/layout_03.png';
import ImageModule from '@assets/module.svg';
import ImageLogo from '@assets/logo.svg';
import commonApi from '@/apis/common-api.ts';
import KeyValuePage from '@/pages/KeyValuePage.tsx';
import ScriptPage from '@/pages/ScriptPage.tsx';

const App: FC = () => {

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

  return (
    <ConfigProvider theme={{
      algorithm: theme.defaultAlgorithm,
      token: {
        // TODO 需要UI 给处主题色
        // colorPrimary: '#006aff',
        // colorPrimaryBg: '#e6f0ff',
        // colorLink: '#006aff',
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
        <BrowserRouter>
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
            menuItemRender={(props, defaultDom) => {
              if (props.isUrl || !props.path) {
                return defaultDom;
              }
              return <Link to={props.path}>{defaultDom}</Link>;
            }}
            headerContentRender={(_, defaultDom) => {
              return <>
                {defaultDom}
              </>;
            }}
            menuExtraRender={(props) => (
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
                <img src={ImageModule} alt=""/>
                {props.collapsed ? undefined : <p>企微助手</p>}
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
                <Route path="/kv" Component={KeyValuePage}/>
                <Route path="/script" Component={ScriptPage}/>
              </Routes>
            </div>
          </ProLayout>
        </BrowserRouter>
      </ProConfigProvider>
    </ConfigProvider>
  );
};
export default App;
