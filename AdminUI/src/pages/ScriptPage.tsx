import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import { Button, message, Modal } from 'antd';
import React, { useRef, useState } from 'react';
import keyApi from '@/apis/key-api.ts';

type ScriptItem = {
  id: string;
  key: string;
  desc: string;
};

export default () => {
  const [_messageApi, contextHolder] = message.useMessage();
  const tableRef = useRef<ActionType>();
  const [_tableParams, setTableParams] = useState({});
  const [labelSelectModal, setLabelSelectModal] = useState(false);
  const columns: ProColumns<ScriptItem>[] = [
    {
      title: '序号',
      dataIndex: 'nickName',
      width: '60px'
    },
    {
      title: '脚本名称',
      dataIndex: 'companyName',
      width: '240px',
      hideInTable: true
    },
    {
      title: '描述',
      ellipsis: true,
      dataIndex: '_panel',
      hideInSearch: true
    }
  ];

  const requestTable = async (params = {}) => {
    // 清除所有选择
    tableRef.current?.clearSelected?.();
    // 设置搜索参数
    setTableParams(params);
    return keyApi.list<ScriptItem>(params)
      .then(result => {
        return result;
      });
  };
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
        toolBarRender={() => [
          <>
            <Button type="primary" onClick={() => {
            }}>
              创建脚本
            </Button>
          </>
        ]}
      />

      <Modal open={labelSelectModal} width={800} onCancel={() => setLabelSelectModal(false)} onOk={() => {
      }}>

      </Modal>
    </>
  );
};
