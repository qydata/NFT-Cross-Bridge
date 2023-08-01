import styled from 'styled-components';

const StatusPage = styled.div`
  .header {
    text-align: center;
    margin-bottom: 24px;
  }

  .connect-wallet-container {
    display: flex;
    padding-bottom: 16px;
    justify-content: flex-end;

    .ant-btn {
      width: 240px;
    }
  }

  .ant-table-thead .ant-table-cell {
    font-weight: bold;
  }

  @media all and (orientation: portrait) {
    /*竖屏*/
    .ant-table table {
      width: 200%;
      overflow: scroll;
    }
  }
  @media all and (orientation: landscape) {
    /*横屏*/
  }
`;
export default StatusPage;
