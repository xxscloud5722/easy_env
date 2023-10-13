import Fetch from '@/apis/api.ts';

class CommonApi extends Fetch {
  async getUserMenus() {
    return {
      data: [{
        id: '1',
        name: '键值对管理',
        link: '/kv'
      }, {
        id: '2',
        name: '脚本管理',
        link: '/script'
      }],
      success: true
    };
  }
}

export default new CommonApi();
