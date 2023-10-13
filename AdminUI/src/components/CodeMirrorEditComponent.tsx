import React, { FC, useEffect } from 'react';
import CodeMirror from '@uiw/react-codemirror';
import { css } from '@emotion/css';
import { theme } from 'antd';

export declare type CodeMirrorEditComponentProps = {
  value?: string | undefined
  onChange?: (value: string | undefined) => void
}
const App: FC<CodeMirrorEditComponentProps> = (props) => {

  const { token } = theme.useToken();

  useEffect(() => {
  }, []);
  return <>
    <div className={css`
      border: 1px solid #ddd;
      border-radius: ${token.borderRadius}px;
      overflow: hidden;

      & .cm-focused {
        outline: none;
      }
    `}>
      <CodeMirror
        value={props.value}
        basicSetup={{
          lineNumbers: true
        }}
        height="200px"
        onChange={(e) => props?.onChange?.(e)}
        //extensions={[s]}
      />
    </div>
  </>;
};
export default App;
