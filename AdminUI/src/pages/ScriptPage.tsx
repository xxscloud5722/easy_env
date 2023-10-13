import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import { css } from '@emotion/css';
import { Button, Dropdown, Form, Input, MenuProps, message, Modal, Popconfirm, Space, Tree } from 'antd';
import { DataNode } from 'antd/es/tree';
import { BlockOutlined, FolderFilled } from '@ant-design/icons';
import CodeMirrorEditComponent from '@components/CodeMirrorEditComponent.tsx';
import React, { useEffect, useRef, useState } from 'react';
import scriptApi from '@/apis/script-api.ts';
import { Config } from '@/config';

const { DirectoryTree } = Tree;

type ScriptItem = {
  id: number
  path: string
  name: string
  description: string
  value: string
};

type DirectoryItem = {
  id: number
  parentId: number
  name: string
};

export default () => {
  const [messageApi, contextHolder] = message.useMessage();
  const tableRef = useRef<ActionType>();
  const [dialogModal, setDialogModal] = useState(false);
  const [dialogDirectoryModal, setDialogDirectoryModal] = useState(false);
  const [script, setScript] = useState({} as ScriptItem);
  const [expandedKeys, setExpandedKeys] = useState([] as number[]);
  const [directoryItems, setDirectoryItems] = useState([] as DirectoryItem[]);
  const [treeData, setTreeData] = useState([] as DataNode[]);
  const [currentDirectory, setCurrentDirectory] = useState('ROOT');
  const [directory, setDirectory] = useState({} as { name: string, parentId: number, id: number });
  const setFieldValue = (field: string, value: any) => {
    const info = { ...script } as any;
    info[field] = value;
    setScript({ ...info } as ScriptItem);
  };
  const onConfirm = async () => {
    const result = await scriptApi.save(script);
    if (result.success) {
      messageApi.success('保存成功');
      tableRef.current?.reload?.();
      setDialogModal(false);
    } else {
      messageApi.error(result.message);
    }
  };
  const onRemove = async (id: number) => {
    const result = await scriptApi.remove(id);
    if (result.success) {
      messageApi.success('删除成功');
      tableRef.current?.reload?.();
    } else {
      messageApi.error(result.message);
    }
  };

  const onRemoveDirectory = async (id: number) => {
    Modal.confirm({
      title: '确认删除目录？',
      content: (
        <>
          <p>删除目录之后不可恢复</p>
          <p>请先清理子目录以及脚本后才可删除</p>
        </>
      ),
      okText: '确定',
      cancelText: '取消',
      onOk: async () => {
        const result = await scriptApi.removeDirectory(id);
        if (result.success) {
          messageApi.success('删除成功');
          setExpandedKeys(expandedKeys.filter(item => item !== id));
          await requestDirectory();
        } else {
          messageApi.error(result.message);
        }
      }
    });
  };
  const onConfirmDirectory = async () => {
    const result = directory.id === undefined || directory.id <= 0
      ? await scriptApi.createDirectory(directory) : await scriptApi.renameDirectory(directory);
    if (result.success) {
      messageApi.success('保存成功');
      await requestDirectory();
      setDialogDirectoryModal(false);
    } else {
      messageApi.error(result.message);
    }
  };
  const onRenameDirectory = async (id: number, parentId: number, name: string) => {
    setDialogDirectoryModal(true);
    setDirectory({
      ...directory,
      id,
      parentId,
      name
    });
  };

  const onCreateDirectory = async (id: number) => {
    setDialogDirectoryModal(true);
    setDirectory({
      ...directory,
      id: 0,
      parentId: id,
      name: ''
    });
  };

  const columns: ProColumns<ScriptItem>[] = [
    {
      title: '名称',
      dataIndex: 'name',
      width: '200px'
    },
    {
      title: '描述',
      ellipsis: true,
      dataIndex: 'description',
      width: '420px'
    },
    {
      title: '操作',
      ellipsis: true,
      dataIndex: '_panel',
      hideInSearch: true,
      render: (_, record) => (
        <Space size={15}>
          <a onClick={() => {
            setScript(record as ScriptItem);
            setDialogModal(true);
          }}>查看</a>
          <Popconfirm
            title="是否确认删除？"
            onConfirm={() => onRemove(record.id)}
            okText="是"
            cancelText="否"
          >
            <a style={{ color: Config.colors.danger }}>删除</a>
          </Popconfirm>
        </Space>
      )
    }
  ];

  const requestTable = async () => {
    return scriptApi.list<ScriptItem>({ path: currentDirectory })
      .then(result => {
        return result;
      });
  };

  const requestDirectory = async () => {
    const recursion = (value: DirectoryItem, items: DirectoryItem[]): DataNode[] => {
      return items.filter(item => item.parentId === value.id)
        .map(item => {
          const menuProps: MenuProps = {
            items: [
              {
                label: <span onClick={() => onRenameDirectory(item.id, item.parentId, item.name)}>重命名目录</span>,
                key: '1'
              },
              {
                label: <span onClick={() => onCreateDirectory(item.id)}>新建子目录</span>,
                key: '2'
              },
              {
                type: 'divider'
              },
              {
                label: <span style={{ color: Config.colors.danger }} onClick={() => onRemoveDirectory(item.id)}>删除目录</span>,
                key: '3'
              }
            ]
          };
          return {
            title: <>
              <div style={{
                display: 'flex',
                alignItems: 'center',
                height: '28px',
                fontSize: '14px',
                color: '#202d40'
              }}>
                <FolderFilled style={{
                  margin: '-1px 6px 0 0',
                  fontSize: '18px',
                  color: '#ffd94b'
                }}/>
                <div style={{
                  flex: 1,
                  width: '100%'
                }}>{item.name}</div>
                <Dropdown menu={menuProps} trigger={['click']}>
                  <div className="dropdown-a" style={{ display: 'none' }}>
                    <a style={{ marginRight: '6px' }} className={css`
                      height: 22px;
                      width: 22px;
                      display: block;

                      & > .icon {
                        margin: 1px 0 0 3px;
                      }

                      :hover {
                        background: #e0e0e0;
                        border-radius: 2px;
                      }
                    `} onClick={(e) => e.preventDefault()}>
                      <BlockOutlined className="icon" style={{
                        fontSize: '16px',
                        color: '#888'
                      }}/>
                    </a>
                  </div>
                </Dropdown>
              </div>
            </>,
            key: item.id,
            children: recursion(item, items)
          } as DataNode;
        });
    };
    scriptApi.directoryList<DirectoryItem[]>()
      .then(result => {
        return result.data;
      })
      .then(directory => {
        setDirectoryItems(directory);
        const directoryItem: DirectoryItem = directory[0];
        const menuProps: MenuProps = {
          items: [
            {
              label: <span onClick={() => onCreateDirectory(directoryItem.id)}>新建子目录</span>,
              key: '1'
            }
          ]
        };

        setTreeData([{
          title: <div style={{
            display: 'flex',
            alignItems: 'center',
            height: '28px',
            fontSize: '14px',
            color: '#202d40'
          }}>
            <FolderFilled style={{
              margin: '-1px 6px 0 0',
              fontSize: '18px',
              color: '#ffd94b'
            }}/>
            <div style={{
              flex: 1,
              width: '100%'
            }}>{directoryItem.name}</div>
            <Dropdown menu={menuProps} trigger={['click']}>
              <div className="dropdown-a" style={{ display: 'none' }}>
                <a style={{ marginRight: '6px' }} className={css`
                  height: 22px;
                  width: 22px;
                  display: block;

                  & > .icon {
                    margin: 1px 0 0 3px;
                  }

                  :hover {
                    background: #e0e0e0;
                    border-radius: 2px;
                  }
                `} onClick={(e) => e.preventDefault()}>
                  <BlockOutlined className="icon" style={{
                    fontSize: '16px',
                    color: '#888'
                  }}/>
                </a>
              </div>
            </Dropdown>
          </div>,
          key: directoryItem.id,
          children: recursion(directoryItem, directory)
        } as DataNode]);
      });
  };
  useEffect(() => {
    tableRef?.current?.reload?.();
  }, [currentDirectory]);
  useEffect(() => {
    requestDirectory()
      .then();
  }, []);
  return (
    <>
      {contextHolder}
      <ProTable<ScriptItem>
        actionRef={tableRef}
        columns={columns}
        scroll={{ x: 'max-content' }}
        request={requestTable}
        size="small"
        rowKey="id"
        search={false}
        options={false}
        tableStyle={{ minHeight: '75vh' }}
        tableRender={(_, dom) => (
          <div style={{
            display: 'flex',
            width: '100%',
            gap: '15px',
            minHeight: '400px'
          }}>
            <div style={{
              minWidth: '320px',
              padding: '0 10px 10px 10px',
              background: '#fff'
            }}>
              <p style={{
                textIndent: '5px',
                fontSize: '16px',
                padding: '15px 0',
                borderBottom: '1px  solid #eee'
              }}>脚本目录</p>
              <DirectoryTree
                blockNode
                showLine
                expandedKeys={expandedKeys}
                onExpand={(expandedKeys, {
                  node,
                  nativeEvent
                }) => {
                  const tag = (nativeEvent.target as HTMLElement).tagName.toLowerCase();
                  const icon = (nativeEvent.target as HTMLElement).getAttribute('data-icon');
                  let flag: boolean;
                  if (tag === 'svg') {
                    flag = (icon === 'minus-square' || icon === 'plus-square');
                  } else if (tag === 'path') {
                    flag = icon === undefined || icon === null;
                  } else {
                    flag = tag === 'div' || tag === 'span';
                  }
                  if (flag) {
                    if (node.children !== undefined && node.children.length > 0) {
                      setExpandedKeys(expandedKeys as number[]);
                    }
                  }
                }}
                onSelect={(selectData) => {
                  const recursion = (id: number): string[] => {
                    const item = directoryItems.find(item => item.id === id);
                    if (item !== undefined) {
                      return [...recursion(item.parentId), item.name];
                    }
                    return [];
                  };
                  const paths = recursion(selectData[0] as number);
                  setCurrentDirectory(paths.join('/'));
                }}
                autoExpandParent={true}
                showIcon={false}
                treeData={treeData}
                className={css`
                  & .anticon {
                    margin-top: 6px;
                  }

                  & .anticon.ant-tree-switcher-line-icon {
                    color: #a0a0a0 !important;
                  }

                  & .ant-tree-treenode-leaf-last .ant-tree-switcher-leaf-line:before {
                    height: 14px !important;
                  }


                  & .ant-tree-switcher-leaf-line:after {
                    margin-top: 2px;
                  }
                `}
              />
            </div>
            <div style={{ flex: 1 }}>
              {dom}
            </div>
          </div>
        )}
        toolBarRender={() => [
          <>
            {currentDirectory !== 'ROOT' ? <Button type="primary" onClick={() => {
              setScript({ path: currentDirectory } as ScriptItem);
              setDialogModal(true);
            }}>
              创建脚本
            </Button> : undefined}
          </>
        ]}
      />

      <Modal open={dialogModal} onCancel={() => setDialogModal(false)} onOk={() => onConfirm()} title={
        script.id === undefined ? '新建脚本' : '脚本管理'} width="560px" styles={{ body: { padding: '10px 0' } }}>
        <Form>
          <Form.Item label="名称">
            <Input value={script.name} onChange={(e) => setFieldValue('name', e.target.value.trim())}/>
          </Form.Item>
          <Form.Item label="路径">
            <Input disabled={true} value={script.path}/>
          </Form.Item>
          <Form.Item label="描述">
            <Input value={script.description} onChange={(e) => setFieldValue('description', e.target.value)}/>
          </Form.Item>
          <Form.Item label="配置">
            <CodeMirrorEditComponent value={script.value} onChange={(e) => setFieldValue('value', e)}/>
          </Form.Item>
        </Form>
      </Modal>

      <Modal
        open={dialogDirectoryModal}
        title={directory.id === undefined || directory.id <= 0 ? '新建目录' : '修改目录'}
        width="420px"
        onCancel={() => setDialogDirectoryModal(false)}
        onOk={() => onConfirmDirectory()}
        styles={{ body: { padding: '10px 0' } }}
      >
        <Input
          placeholder="请输入目录名称"
          value={directory.name}
          onChange={(e) => setDirectory({
            ...directory,
            name: e.target.value
          })}/>
      </Modal>
    </>
  );
};
