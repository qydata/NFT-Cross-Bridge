import Spin from 'antd/lib/spin';

const CallbackPage: React.FC = () => {
  return (
    <div
      style={{
        display: 'flex',
        width: '100%',
        height: '100vh',
        justifyContent: 'center',
        alignItems: 'center'
      }}
    >
      <Spin size='large' />
    </div>
  );
};

export default CallbackPage;
