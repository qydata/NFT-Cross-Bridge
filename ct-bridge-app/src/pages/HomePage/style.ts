import styled from 'styled-components';

const HomePageStyle = styled.div`
  text-align: center;
  .bridge {
    letter-spacing: 2px;
    font-size: 14px;
    white-space: pre;
    color: #0b6a62;
  }
  @media all and (orientation: portrait) {
    /*竖屏*/
    .bridge {
      transform: scale(0.35, 0.35);
      transform-origin: left;
    }
  }
  @media all and (orientation: landscape) {
    /*横屏*/
  }
  h1 {
    margin-top: 12px !important;
    margin-bottom: 32px;
  }
  .image-container {
    img {
      width: 50px;
      margin-left: 16px;
    }
  }
`;
export default HomePageStyle;
