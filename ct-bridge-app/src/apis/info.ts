import axios from 'axios';
import setting from 'src/setting';

export const getInfo = async () => {
  const url = `${setting.API_URL}/v1/info`;
  const response = await axios.get<any>(url);
  return response.data;
};
