import styled from 'styled-components';

const NFTDetailStyle = styled.div`
  .nft-detail-headar {
    display: flex;
    justify-content: space-between;
    margin-bottom: 8px;

    .anticon-close-circle {
      font-size: 20px;
    }
  }

  .detail-container {
    .ant-typography {
      margin: 0;
    }
    .detail {
      margin-bottom: 8px;
      ${(props: any) => (props.disabled ? 'color: gray;' : '')}
    }
  }

  .action-container {
    margin-top: 32px;
    display: flex;
    flex-direction: column;
    align-items: flex-end;

    .ant-btn {
      width: 150px;
    }
  }

  .register-info {
    margin-top: 8px;
    color: red;
  }
  @media all and (orientation: portrait) {
    /*竖屏*/
  }
  @media all and (orientation: landscape) {
    /*横屏*/
  }
`;
export default NFTDetailStyle;
