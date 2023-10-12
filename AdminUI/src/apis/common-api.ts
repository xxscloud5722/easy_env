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

  async externalLinkSearch<T>(params: {}) {
    return this.pageTable<T>(await this.post<any>('/wechat/externalLinkHelper/getLinkList', undefined, this.pageParams(params)));
  }

  async shuntQrCodeSearch<T>(params: {}) {
    return this.pageTable<T>(await this.post<any>('/wechat/shuntQrCode/search', undefined, this.pageParams(params)));
  }

  async domainSearch<T>(params: {}) {
    return this.pageTable<T>(await this.post<any>('/wechat/domain/getDomainList', undefined, this.pageParams(params)));
  }

  async getEventLastList<T>(params: {}) {
    return this.get<T>('/wechat/event/getLastList', params);
  }

  async releaseEvent<T>(params: {}) {
    return this.post<T>('/wechat/event/release', undefined, params);
  }

  async getNextRegion<T>(params: {}) {
    return this.post<T>('/usercenter/sysArea/nextList', undefined, params);
  }

  async searchStaff<T>(params: {}) {
    return this.get<T>('/wechat/staff/searchStaff', params);
  }
}

export default new CommonApi();
