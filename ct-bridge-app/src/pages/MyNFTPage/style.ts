import styled from 'styled-components';

const MyNFTPageStyle = styled.div`
  .header {
    text-align: center;
    margin-bottom: 24px;
  }
  .loading-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 40px;
  }

  .search-container {
    display: flex;
    justify-content: space-between;
    margin-bottom: 16px;
  }

  .no-data-container {
    display: flex;
    justify-content: center;
    height: 100%;
    padding-top: 40px;
    align-items: center;
  }

  .ant-card {
    border-radius: 8px;
    width: 250px;
  }
  .ant-card-cover {
    border-radius: 8px;
  }
  .ant-radio-group {
    font-weight: bold;
    &.ant-radio-group-solid
      .ant-radio-button-wrapper-checked:not(
        .ant-radio-button-wrapper-disabled
      ) {
      background: #3fafac;
      font-family: monospace;
      border-color: #3fafac;
      &:hover {
        background: #2c9794;
        border-color: #2c9794;
        color: white;
      }
    }
    .ant-radio-button-wrapper-checked:not(
        .ant-radio-button-wrapper-disabled
      )::before {
      background: #3fafac;
    }
    .ant-radio-button-wrapper:hover {
      color: #2c9794;
    }
    .ant-radio-button-wrapper:first-child {
      border-radius: 10px 0 0 10px;
    }
    .ant-radio-button-wrapper:last-child {
      border-radius: 0 10px 10px 0;
    }
  }
  @media all and (orientation: portrait) {
    /*竖屏*/
    .ant-card {
      width: 100%;
      margin-bottom: 10px;
    }
  }
  @media all and (orientation: landscape) {
    /*横屏*/
  }
`;

export default MyNFTPageStyle;
