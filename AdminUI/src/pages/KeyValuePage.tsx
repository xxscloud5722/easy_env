import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import CodeMirrorEditComponent from '@components/CodeMirrorEditComponent.tsx';
import { Button, Form, Input, message, Modal, Popconfirm, Space } from 'antd';
import React, { useRef, useState } from 'react';
import keyApi from '@/apis/key-api.ts';
import { Config } from '@/config';

type KeyValueItem = {
  id: string | undefined
  key: string
  value: string
  description: string
};

export default () => {
  const [messageApi, contextHolder] = message.useMessage();
  const tableRef = useRef<ActionType>();
  const [searchParams, setSearchParams] = useState({ keyword: '' });
  const [keyValue, setKeyValue] = useState({} as KeyValueItem);
  const setSearchParamValue = (key: string, value: string | number) => {
    const params = searchParams as any;
    params[key] = value;
    setSearchParams(params);
  };
  const [dialogModal, setDialogModal] = useState(false);
  const setFieldValue = (field: string, value: any) => {
    const info = { ...keyValue } as any;
    info[field] = value;
    setKeyValue({ ...info } as KeyValueItem);
  };
  const onConfirm = async () => {
    const result = await keyApi.save(keyValue);
    if (result.success) {
      messageApi.success('保存成功');
      tableRef.current?.reload?.();
      setDialogModal(false);
    } else {
      messageApi.error(result.message);
    }
  };
  const onRemove = async (key: string) => {
    const result = await keyApi.remove(key);
    if (result.success) {
      messageApi.success('删除成功');
      tableRef.current?.reload?.();
    } else {
      messageApi.error(result.message);
    }
  };
  const columns: ProColumns<KeyValueItem>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      width: '60px'
    },
    {
      title: '名称',
      dataIndex: 'key',
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
            setKeyValue(record);
            setDialogModal(true);
          }}>查看</a>
          <Popconfirm
            title="是否确认删除？"
            onConfirm={() => onRemove(record.key)}
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
    return keyApi.list<KeyValueItem>(searchParams.keyword)
      .then(result => {
        return result;
      });
  };
  return (
    <>
      {contextHolder}
      <ProTable<KeyValueItem>
        actionRef={tableRef}
        columns={columns}
        scroll={{ x: 'max-content' }}
        request={requestTable}
        size="small"
        rowKey="id"
        search={false}
        options={false}
        headerTitle={
          <Form layout="inline" style={{ width: '300px' }} onFinish={() => tableRef.current?.reload()}>
            <Form.Item>
              <Input placeholder="前缀匹配键值对" onChange={(e) => setSearchParamValue('keyword', e.target.value)}/>
            </Form.Item>
            <Form.Item>
              <Button type="primary" onClick={() => tableRef.current?.reload()}>
                查询
              </Button>
            </Form.Item>
          </Form>
        }
        toolBarRender={() => [
          <>
            <Button type="primary" onClick={() => {
              setKeyValue({} as KeyValueItem);
              setDialogModal(true);
            }}>
              创建键值对
            </Button>
          </>
        ]}
      />

      <Modal open={dialogModal} onCancel={() => setDialogModal(false)} onOk={onConfirm} title={
        keyValue.id === undefined ? '新建键值对' : '键值对管理'} width="560px" styles={{ body: { padding: '10px 0' } }}>
        <Form>
          <Form.Item label="名称">
            <Input value={keyValue.key} onChange={(e) => setFieldValue('key', e.target.value.trim())}/>
          </Form.Item>
          <Form.Item label="描述">
            <Input value={keyValue.description} onChange={(e) => setFieldValue('description', e.target.value)}/>
          </Form.Item>
          <Form.Item label="配置">
            <CodeMirrorEditComponent value={keyValue.value} onChange={(e) => setFieldValue('value', e)}/>
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};
